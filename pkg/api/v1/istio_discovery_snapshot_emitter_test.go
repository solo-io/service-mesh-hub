// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"os"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	istio_authentication_v1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/authorization/v1alpha1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	kuberc "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/test/helpers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1       string
		namespace2       string
		name1, name2     = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg              *rest.Config
		kube             kubernetes.Interface
		emitter          IstioDiscoveryEmitter
		meshClient       MeshClient
		installClient    InstallClient
		podClient        github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodClient
		upstreamClient   gloo_solo_io.UpstreamClient
		meshPolicyClient istio_authentication_v1alpha1.MeshPolicyClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		kube = helpers.MustKubeClient()
		err := kubeutils.CreateNamespacesInParallel(kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		// Mesh Constructor
		meshClientFactory := &factory.KubeResourceClientFactory{
			Crd:         MeshCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}

		meshClient, err = NewMeshClient(meshClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Install Constructor
		installClientFactory := &factory.KubeResourceClientFactory{
			Crd:         InstallCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}

		installClient, err = NewInstallClient(installClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Pod Constructor
		podClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		podClient, err = github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPodClient(podClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Upstream Constructor
		upstreamClientFactory := &factory.KubeResourceClientFactory{
			Crd:         gloo_solo_io.UpstreamCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}

		upstreamClient, err = gloo_solo_io.NewUpstreamClient(upstreamClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// MeshPolicy Constructor
		meshPolicyClientFactory := &factory.KubeResourceClientFactory{
			Crd:         istio_authentication_v1alpha1.MeshPolicyCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}

		meshPolicyClient, err = istio_authentication_v1alpha1.NewMeshPolicyClient(meshPolicyClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewIstioDiscoveryEmitter(meshClient, installClient, podClient, upstreamClient, meshPolicyClient)
	})
	AfterEach(func() {
		err := kubeutils.DeleteNamespacesInParallelBlocking(kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
		meshPolicyClient.Delete(name1, clients.DeleteOpts{})
		meshPolicyClient.Delete(name2, clients.DeleteOpts{})
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *IstioDiscoverySnapshot

		/*
			Mesh
		*/

		assertSnapshotMeshes := func(expectMeshes MeshList, unexpectMeshes MeshList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMeshes {
						if _, err := snap.Meshes.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMeshes {
						if _, err := snap.Meshes.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := meshClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := meshClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		mesh1a, err := meshClient.Write(NewMesh(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mesh1b, err := meshClient.Write(NewMesh(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b}, nil)
		mesh2a, err := meshClient.Write(NewMesh(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mesh2b, err := meshClient.Write(NewMesh(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b, mesh2a, mesh2b}, nil)

		err = meshClient.Delete(mesh2a.GetMetadata().Namespace, mesh2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = meshClient.Delete(mesh2b.GetMetadata().Namespace, mesh2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b}, MeshList{mesh2a, mesh2b})

		err = meshClient.Delete(mesh1a.GetMetadata().Namespace, mesh1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = meshClient.Delete(mesh1b.GetMetadata().Namespace, mesh1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(nil, MeshList{mesh1a, mesh1b, mesh2a, mesh2b})

		/*
			Install
		*/

		assertSnapshotInstalls := func(expectInstalls InstallList, unexpectInstalls InstallList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectInstalls {
						if _, err := snap.Installs.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectInstalls {
						if _, err := snap.Installs.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := installClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := installClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		install1a, err := installClient.Write(NewInstall(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		install1b, err := installClient.Write(NewInstall(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotInstalls(InstallList{install1a, install1b}, nil)
		install2a, err := installClient.Write(NewInstall(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		install2b, err := installClient.Write(NewInstall(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotInstalls(InstallList{install1a, install1b, install2a, install2b}, nil)

		err = installClient.Delete(install2a.GetMetadata().Namespace, install2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = installClient.Delete(install2b.GetMetadata().Namespace, install2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotInstalls(InstallList{install1a, install1b}, InstallList{install2a, install2b})

		err = installClient.Delete(install1a.GetMetadata().Namespace, install1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = installClient.Delete(install1b.GetMetadata().Namespace, install1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotInstalls(nil, InstallList{install1a, install1b, install2a, install2b})

		/*
			Pod
		*/

		assertSnapshotpods := func(expectpods github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList, unexpectpods github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectpods {
						if _, err := snap.Pods.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectpods {
						if _, err := snap.Pods.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := podClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := podClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		pod1a, err := podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		pod1b, err := podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotpods(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod1a, pod1b}, nil)
		pod2a, err := podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		pod2b, err := podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotpods(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod1a, pod1b, pod2a, pod2b}, nil)

		err = podClient.Delete(pod2a.GetMetadata().Namespace, pod2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = podClient.Delete(pod2b.GetMetadata().Namespace, pod2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotpods(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod1a, pod1b}, github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod2a, pod2b})

		err = podClient.Delete(pod1a.GetMetadata().Namespace, pod1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = podClient.Delete(pod1b.GetMetadata().Namespace, pod1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotpods(nil, github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod1a, pod1b, pod2a, pod2b})

		/*
			Upstream
		*/

		assertSnapshotUpstreams := func(expectUpstreams gloo_solo_io.UpstreamList, unexpectUpstreams gloo_solo_io.UpstreamList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreams {
						if _, err := snap.Upstreams.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreams {
						if _, err := snap.Upstreams.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := upstreamClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := upstreamClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		upstream1a, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream1b, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b}, nil)
		upstream2a, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream2b, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b}, nil)

		err = upstreamClient.Delete(upstream2a.GetMetadata().Namespace, upstream2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2b.GetMetadata().Namespace, upstream2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b}, gloo_solo_io.UpstreamList{upstream2a, upstream2b})

		err = upstreamClient.Delete(upstream1a.GetMetadata().Namespace, upstream1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.GetMetadata().Namespace, upstream1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(nil, gloo_solo_io.UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b})

		/*
			MeshPolicy
		*/

		assertSnapshotMeshpolicies := func(expectMeshpolicies istio_authentication_v1alpha1.MeshPolicyList, unexpectMeshpolicies istio_authentication_v1alpha1.MeshPolicyList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMeshpolicies {
						if _, err := snap.Meshpolicies.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMeshpolicies {
						if _, err := snap.Meshpolicies.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					combined, _ := meshPolicyClient.List(clients.ListOpts{})
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		meshPolicy1a, err := meshPolicyClient.Write(istio_authentication_v1alpha1.NewMeshPolicy(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshpolicies(istio_authentication_v1alpha1.MeshPolicyList{meshPolicy1a}, nil)
		meshPolicy2a, err := meshPolicyClient.Write(istio_authentication_v1alpha1.NewMeshPolicy(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshpolicies(istio_authentication_v1alpha1.MeshPolicyList{meshPolicy1a, meshPolicy2a}, nil)

		err = meshPolicyClient.Delete(meshPolicy2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshpolicies(istio_authentication_v1alpha1.MeshPolicyList{meshPolicy1a}, istio_authentication_v1alpha1.MeshPolicyList{meshPolicy2a})

		err = meshPolicyClient.Delete(meshPolicy1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshpolicies(nil, istio_authentication_v1alpha1.MeshPolicyList{meshPolicy1a, meshPolicy2a})
	})
	It("tracks snapshots on changes to any resource using AllNamespace", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{""}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *IstioDiscoverySnapshot

		/*
			Mesh
		*/

		assertSnapshotMeshes := func(expectMeshes MeshList, unexpectMeshes MeshList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMeshes {
						if _, err := snap.Meshes.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMeshes {
						if _, err := snap.Meshes.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := meshClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := meshClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		mesh1a, err := meshClient.Write(NewMesh(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mesh1b, err := meshClient.Write(NewMesh(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b}, nil)
		mesh2a, err := meshClient.Write(NewMesh(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mesh2b, err := meshClient.Write(NewMesh(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b, mesh2a, mesh2b}, nil)

		err = meshClient.Delete(mesh2a.GetMetadata().Namespace, mesh2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = meshClient.Delete(mesh2b.GetMetadata().Namespace, mesh2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b}, MeshList{mesh2a, mesh2b})

		err = meshClient.Delete(mesh1a.GetMetadata().Namespace, mesh1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = meshClient.Delete(mesh1b.GetMetadata().Namespace, mesh1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(nil, MeshList{mesh1a, mesh1b, mesh2a, mesh2b})

		/*
			Install
		*/

		assertSnapshotInstalls := func(expectInstalls InstallList, unexpectInstalls InstallList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectInstalls {
						if _, err := snap.Installs.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectInstalls {
						if _, err := snap.Installs.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := installClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := installClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		install1a, err := installClient.Write(NewInstall(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		install1b, err := installClient.Write(NewInstall(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotInstalls(InstallList{install1a, install1b}, nil)
		install2a, err := installClient.Write(NewInstall(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		install2b, err := installClient.Write(NewInstall(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotInstalls(InstallList{install1a, install1b, install2a, install2b}, nil)

		err = installClient.Delete(install2a.GetMetadata().Namespace, install2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = installClient.Delete(install2b.GetMetadata().Namespace, install2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotInstalls(InstallList{install1a, install1b}, InstallList{install2a, install2b})

		err = installClient.Delete(install1a.GetMetadata().Namespace, install1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = installClient.Delete(install1b.GetMetadata().Namespace, install1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotInstalls(nil, InstallList{install1a, install1b, install2a, install2b})

		/*
			Pod
		*/

		assertSnapshotpods := func(expectpods github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList, unexpectpods github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectpods {
						if _, err := snap.Pods.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectpods {
						if _, err := snap.Pods.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := podClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := podClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		pod1a, err := podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		pod1b, err := podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotpods(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod1a, pod1b}, nil)
		pod2a, err := podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		pod2b, err := podClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewPod(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotpods(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod1a, pod1b, pod2a, pod2b}, nil)

		err = podClient.Delete(pod2a.GetMetadata().Namespace, pod2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = podClient.Delete(pod2b.GetMetadata().Namespace, pod2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotpods(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod1a, pod1b}, github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod2a, pod2b})

		err = podClient.Delete(pod1a.GetMetadata().Namespace, pod1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = podClient.Delete(pod1b.GetMetadata().Namespace, pod1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotpods(nil, github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodList{pod1a, pod1b, pod2a, pod2b})

		/*
			Upstream
		*/

		assertSnapshotUpstreams := func(expectUpstreams gloo_solo_io.UpstreamList, unexpectUpstreams gloo_solo_io.UpstreamList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectUpstreams {
						if _, err := snap.Upstreams.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectUpstreams {
						if _, err := snap.Upstreams.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := upstreamClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := upstreamClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		upstream1a, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream1b, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b}, nil)
		upstream2a, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		upstream2b, err := upstreamClient.Write(gloo_solo_io.NewUpstream(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b}, nil)

		err = upstreamClient.Delete(upstream2a.GetMetadata().Namespace, upstream2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream2b.GetMetadata().Namespace, upstream2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(gloo_solo_io.UpstreamList{upstream1a, upstream1b}, gloo_solo_io.UpstreamList{upstream2a, upstream2b})

		err = upstreamClient.Delete(upstream1a.GetMetadata().Namespace, upstream1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = upstreamClient.Delete(upstream1b.GetMetadata().Namespace, upstream1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotUpstreams(nil, gloo_solo_io.UpstreamList{upstream1a, upstream1b, upstream2a, upstream2b})

		/*
			MeshPolicy
		*/

		assertSnapshotMeshpolicies := func(expectMeshpolicies istio_authentication_v1alpha1.MeshPolicyList, unexpectMeshpolicies istio_authentication_v1alpha1.MeshPolicyList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMeshpolicies {
						if _, err := snap.Meshpolicies.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMeshpolicies {
						if _, err := snap.Meshpolicies.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					combined, _ := meshPolicyClient.List(clients.ListOpts{})
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		meshPolicy1a, err := meshPolicyClient.Write(istio_authentication_v1alpha1.NewMeshPolicy(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshpolicies(istio_authentication_v1alpha1.MeshPolicyList{meshPolicy1a}, nil)
		meshPolicy2a, err := meshPolicyClient.Write(istio_authentication_v1alpha1.NewMeshPolicy(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshpolicies(istio_authentication_v1alpha1.MeshPolicyList{meshPolicy1a, meshPolicy2a}, nil)

		err = meshPolicyClient.Delete(meshPolicy2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshpolicies(istio_authentication_v1alpha1.MeshPolicyList{meshPolicy1a}, istio_authentication_v1alpha1.MeshPolicyList{meshPolicy2a})

		err = meshPolicyClient.Delete(meshPolicy1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshpolicies(nil, istio_authentication_v1alpha1.MeshPolicyList{meshPolicy1a, meshPolicy2a})
	})
})
