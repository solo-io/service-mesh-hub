// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"os"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

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
		namespace1      string
		namespace2      string
		name1, name2    = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg             *rest.Config
		kube            kubernetes.Interface
		emitter         DiscoveryEmitter
		podClient       github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.PodClient
		upstreamClient  gloo_solo_io.UpstreamClient
		configMapClient github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		kube = helpers.MustKubeClient()
		err := kubeutils.CreateNamespacesInParallel(kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
		cfg, err = kubeutils.GetConfig("", "")
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
		// ConfigMap Constructor
		configMapClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}

		configMapClient, err = github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMapClient(configMapClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewDiscoveryEmitter(podClient, upstreamClient, configMapClient)
	})
	AfterEach(func() {
		err := kubeutils.DeleteNamespacesInParallelBlocking(kube, namespace1, namespace2)
		Expect(err).NotTo(HaveOccurred())
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

		var snap *DiscoverySnapshot

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
			ConfigMap
		*/

		assertSnapshotconfigmaps := func(expectconfigmaps github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList, unexpectconfigmaps github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectconfigmaps {
						if _, err := snap.Configmaps.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectconfigmaps {
						if _, err := snap.Configmaps.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := configMapClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := configMapClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		configMap1a, err := configMapClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMap(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		configMap1b, err := configMapClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMap(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotconfigmaps(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap1a, configMap1b}, nil)
		configMap2a, err := configMapClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMap(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		configMap2b, err := configMapClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMap(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotconfigmaps(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap1a, configMap1b, configMap2a, configMap2b}, nil)

		err = configMapClient.Delete(configMap2a.GetMetadata().Namespace, configMap2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = configMapClient.Delete(configMap2b.GetMetadata().Namespace, configMap2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotconfigmaps(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap1a, configMap1b}, github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap2a, configMap2b})

		err = configMapClient.Delete(configMap1a.GetMetadata().Namespace, configMap1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = configMapClient.Delete(configMap1b.GetMetadata().Namespace, configMap1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotconfigmaps(nil, github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap1a, configMap1b, configMap2a, configMap2b})
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

		var snap *DiscoverySnapshot

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
			ConfigMap
		*/

		assertSnapshotconfigmaps := func(expectconfigmaps github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList, unexpectconfigmaps github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectconfigmaps {
						if _, err := snap.Configmaps.Find(expected.GetMetadata().Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectconfigmaps {
						if _, err := snap.Configmaps.Find(unexpected.GetMetadata().Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := configMapClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := configMapClient.List(namespace2, clients.ListOpts{})
					combined := append(nsList1, nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		configMap1a, err := configMapClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMap(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		configMap1b, err := configMapClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMap(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotconfigmaps(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap1a, configMap1b}, nil)
		configMap2a, err := configMapClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMap(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		configMap2b, err := configMapClient.Write(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewConfigMap(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotconfigmaps(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap1a, configMap1b, configMap2a, configMap2b}, nil)

		err = configMapClient.Delete(configMap2a.GetMetadata().Namespace, configMap2a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = configMapClient.Delete(configMap2b.GetMetadata().Namespace, configMap2b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotconfigmaps(github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap1a, configMap1b}, github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap2a, configMap2b})

		err = configMapClient.Delete(configMap1a.GetMetadata().Namespace, configMap1a.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = configMapClient.Delete(configMap1b.GetMetadata().Namespace, configMap1b.GetMetadata().Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotconfigmaps(nil, github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.ConfigMapList{configMap1a, configMap1b, configMap2a, configMap2b})
	})
})
