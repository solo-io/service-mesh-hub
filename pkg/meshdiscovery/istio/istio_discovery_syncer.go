package istio

import (
	"context"
	"strings"

	"github.com/solo-io/supergloo/pkg/translator/utils"
	"go.uber.org/zap"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/supergloo/pkg/api/external/istio/authorization/v1alpha1"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

// the subset of the api extensions client that we need
type CrdGetter interface {
	Get(name string, options metav1.GetOptions) (*v1beta1.CustomResourceDefinition, error)
}

type istioDiscoverySyncer struct {
	writeNamespace   string
	meshReconciler   v1.MeshReconciler
	meshPolicyClient v1alpha1.MeshPolicyClient
	crdGetter        CrdGetter
}

func NewIstioDiscoverySyncer(writeNamespace string, meshReconciler v1.MeshReconciler, meshPolicyClient v1alpha1.MeshPolicyClient, crdGetter CrdGetter) v1.DiscoverySyncer {
	return &istioDiscoverySyncer{writeNamespace: writeNamespace, meshReconciler: meshReconciler, meshPolicyClient: meshPolicyClient, crdGetter: crdGetter}
}

var discoveryLabels = map[string]string{
	"discovered_by": "istio-mesh-discovery",
}

func (i *istioDiscoverySyncer) Sync(ctx context.Context, snap *v1.DiscoverySnapshot) error {
	ctx = contextutils.WithLogger(ctx, "istio-mesh-discovery")
	logger := contextutils.LoggerFrom(ctx)
	logger.Infow("begin sync",
		zap.Int("Upstreams", len(snap.Upstreams)),
		zap.Int("Deployments", len(snap.Deployments)),
		zap.Int("Tlssecrets", len(snap.Tlssecrets)),
		zap.Int("Configmaps", len(snap.Configmaps)),
		zap.Int("Pods", len(snap.Pods)),
	)
	defer logger.Infow("end sync")
	logger.Debugf("full snapshot: %v", snap)

	istioMeshes, err := i.desiredMeshes(ctx, snap)
	if err != nil {
		return err
	}

	istioMeshes.Each(func(element *v1.Mesh) {
		element.Metadata.Labels = discoveryLabels
	})

	if err := i.meshReconciler.Reconcile(i.writeNamespace, istioMeshes, onlyUpdateDiscoveryMetadata, clients.ListOpts{Ctx: ctx, Selector: discoveryLabels}); err != nil {
		return err
	}

	return nil
}

type pilotDeployment struct {
	version, namespace string
}

func (i *istioDiscoverySyncer) desiredMeshes(ctx context.Context, snap *v1.DiscoverySnapshot) (v1.MeshList, error) {
	meshPolicyCrdRegistered, err := detectMeshPolicyCrd(i.crdGetter)
	if err != nil {
		return nil, err
	}
	if !meshPolicyCrdRegistered {
		return nil, nil
	}

	pilots, err := detectPilotDeployments(snap.Deployments)
	if err != nil {
		return nil, err
	}
	if len(pilots) == 0 {
		return nil, nil
	}

	globalMtlsEnabled := func() bool {
		// https://istio.io/docs/tasks/security/authn-policy/#globally-enabling-istio-mutual-tls
		defaultMeshPolicy, err := i.meshPolicyClient.Read("default", clients.ReadOpts{Ctx: ctx})
		if err != nil {
			return false
		}
		for _, peer := range defaultMeshPolicy.GetPeers() {
			return peer.GetMtls() != nil
		}
		return false
	}()

	injectedPods, err := detectInjectedIstioPods(snap.Pods)
	if err != nil {
		return nil, err
	}

	var istioMeshes v1.MeshList
	for _, pilot := range pilots {
		var autoInjectionEnabled bool
		sidecarInjector, err := snap.Deployments.Find(pilot.namespace, "istio-sidecar-injector")
		if err == nil && (sidecarInjector.Spec.Replicas == nil || *sidecarInjector.Spec.Replicas > 0) {
			autoInjectionEnabled = true
		}

		var smiEnabled bool
		smiAdapter, err := snap.Deployments.Find(pilot.namespace, "smi-adapter-istio")
		if err == nil && (smiAdapter.Spec.Replicas == nil || *smiAdapter.Spec.Replicas > 0) {
			smiEnabled = true
		}

		// https://istio.io/docs/tasks/security/plugin-ca-cert/#plugging-in-the-existing-certificate-and-key
		var rootCa *core.ResourceRef
		customRootCa, err := snap.Tlssecrets.Find(pilot.namespace, "cacerts")
		if err == nil {
			root := customRootCa.Metadata.Ref()
			rootCa = &root
		}

		mtlsConfig := &v1.MtlsConfig{
			MtlsEnabled:     globalMtlsEnabled,
			RootCertificate: rootCa,
		}

		meshUpstreams := func() []*core.ResourceRef {
			injectedUpstreams := utils.UpstreamsForPods(injectedPods[pilot.namespace], snap.Upstreams)
			var usRefs []*core.ResourceRef
			for _, us := range injectedUpstreams {
				ref := us.Metadata.Ref()
				usRefs = append(usRefs, &ref)
			}
			return usRefs
		}()

		istioMesh := &v1.Mesh{
			Metadata: core.Metadata{
				Name:      pilot.namespace + "-istio",
				Namespace: i.writeNamespace,
			},
			MeshType: &v1.Mesh_Istio{
				Istio: &v1.IstioMesh{
					InstallationNamespace: pilot.namespace,
					Version:               pilot.version,
				},
			},
			MtlsConfig: mtlsConfig,
			SmiEnabled: smiEnabled,
			DiscoveryMetadata: &v1.DiscoveryMetadata{
				EnableAutoInject: autoInjectionEnabled,
				MtlsConfig:       mtlsConfig,
				Upstreams:        meshUpstreams,
			},
		}
		istioMeshes = append(istioMeshes, istioMesh)
	}

	return istioMeshes, nil
}

func detectPilotDeployments(deployments kubernetes.DeploymentList) ([]pilotDeployment, error) {
	var pilots []pilotDeployment
	for _, deployment := range deployments {
		for _, container := range deployment.Spec.Template.Spec.Containers {
			if strings.Contains(container.Image, "istio/pilot") {
				split := strings.Split(container.Image, ":")
				if len(split) != 2 {
					return nil, errors.Errorf("invalid or unexpected image format for pilot: %v", container.Image)
				}
				pilots = append(pilots, pilotDeployment{version: split[1], namespace: deployment.Namespace})
			}
		}
	}
	return pilots, nil
}

func detectMeshPolicyCrd(crdGetter CrdGetter) (bool, error) {
	_, err := crdGetter.Get(v1alpha1.MeshPolicyCrd.FullName(), metav1.GetOptions{})
	if err == nil {
		return true, nil
	}
	if kubeerrs.IsNotFound(err) {
		return false, nil
	}
	return false, err
}

func detectInjectedIstioPods(pods kubernetes.PodList) (map[string]kubernetes.PodList, error) {
	injectedPods := make(map[string]kubernetes.PodList)
findInjectedPods:
	for _, pod := range pods {
		for _, container := range pod.Spec.Containers {
			if container.Name == "istio-proxy" {
				for i, arg := range container.Args {
					if arg == "--discoveryAddress" {
						if i == len(container.Args) {
							return nil, errors.Errorf("invalid args for istio-proxy sidecar for pod %v.%v", pod.Namespace, pod.Name)
						}
						discoveryAddress := container.Args[i+1]
						discoveryService := strings.Split(discoveryAddress, ":")[0]
						discoveryNamespace := strings.TrimPrefix(discoveryService, "istio-pilot.")
						injectedPods[discoveryNamespace] = append(injectedPods[discoveryNamespace], pod)
						continue findInjectedPods
					}
				}
			}
		}
	}
	return injectedPods, nil
}

// after the first write, i.e. on any update
// we are only interested in updating discovery metadata at this point
// do not want to overwrite user-modified config settings
// smi is also included here as it's not intended to be user-writeable
func onlyUpdateDiscoveryMetadata(original, desired *v1.Mesh) (b bool, e error) {
	if desired.DiscoveryMetadata.Equal(original.DiscoveryMetadata) && desired.SmiEnabled == original.SmiEnabled {
		return false, nil
	}
	original.DiscoveryMetadata = desired.DiscoveryMetadata
	original.SmiEnabled = desired.SmiEnabled
	*desired = *original
	return true, nil
}
