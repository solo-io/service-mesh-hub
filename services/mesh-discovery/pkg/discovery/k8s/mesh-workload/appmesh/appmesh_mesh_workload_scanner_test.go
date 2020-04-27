package appmesh_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_discovery_types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/service-mesh-hub/pkg/env"
	mesh_workload "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/k8s/mesh-workload"
	"github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/k8s/mesh-workload/appmesh"
	mock_mesh_workload "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/k8s/mesh-workload/mocks"
	mock_core "github.com/solo-io/service-mesh-hub/test/mocks/clients/discovery.zephyr.solo.io/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("AppmeshMeshWorkloadScanner", func() {
	var (
		ctrl                *gomock.Controller
		ctx                 context.Context
		mockOwnerFetcher    *mock_mesh_workload.MockOwnerFetcher
		meshWorkloadScanner mesh_workload.MeshWorkloadScanner
		mockMeshClient      *mock_core.MockMeshClient
		namespace           = "namespace"
		clusterName         = "clusterName"
		deploymentName      = "deployment-name"
		deploymentKind      = "deployment-kind"
		meshName            = "mesh-name-1"
		region              = "us-east-1"
		pod                 = &corev1.Pod{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Image: "appmesh-proxy",
						Env: []corev1.EnvVar{
							{
								Name:  appmesh.AppMeshVirtualNodeEnvVarName,
								Value: fmt.Sprintf("mesh/%s/virtualNode/virtualNodeName", meshName),
							},
							{
								Name:  appmesh.AppMeshRegionEnvVarName,
								Value: region,
							},
						},
					},
				},
			},
			ObjectMeta: metav1.ObjectMeta{ClusterName: clusterName, Namespace: namespace},
		}
		deployment = &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{Name: deploymentName, Namespace: namespace},
			TypeMeta:   metav1.TypeMeta{Kind: deploymentKind},
		}
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		mockOwnerFetcher = mock_mesh_workload.NewMockOwnerFetcher(ctrl)
		mockMeshClient = mock_core.NewMockMeshClient(ctrl)
		meshWorkloadScanner = appmesh.NewAppMeshWorkloadScanner(mockOwnerFetcher, mockMeshClient)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should scan pod and disambiguate multiple AppMesh Meshes", func() {
		mockOwnerFetcher.EXPECT().GetDeployment(ctx, pod).Return(deployment, nil)
		meshList := &zephyr_discovery.MeshList{
			Items: []zephyr_discovery.Mesh{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      meshName,
						Namespace: "mesh-namespace",
					},
					Spec: zephyr_discovery_types.MeshSpec{
						Cluster: &zephyr_core_types.ResourceRef{Name: clusterName},
						MeshType: &zephyr_discovery_types.MeshSpec_AwsAppMesh_{
							AwsAppMesh: &zephyr_discovery_types.MeshSpec_AwsAppMesh{
								Name:   meshName,
								Region: region,
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "mesh-name-2",
						Namespace: "mesh-namespace",
					},
					Spec: zephyr_discovery_types.MeshSpec{
						Cluster: &zephyr_core_types.ResourceRef{Name: clusterName},
						MeshType: &zephyr_discovery_types.MeshSpec_AwsAppMesh_{
							AwsAppMesh: &zephyr_discovery_types.MeshSpec_AwsAppMesh{
								Name:   "some-other-mesh",
								Region: "some-other-region",
							},
						},
					},
				},
			},
		}
		mockMeshClient.EXPECT().ListMesh(ctx).Return(meshList, nil)
		expectedMeshWorkload := &zephyr_discovery.MeshWorkload{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("appmesh-%s-%s-%s", deploymentName, namespace, clusterName),
				Namespace: env.GetWriteNamespace(),
				Labels:    appmesh.DiscoveryLabels(),
			},
			Spec: zephyr_discovery_types.MeshWorkloadSpec{
				KubeController: &zephyr_discovery_types.MeshWorkloadSpec_KubeController{
					KubeControllerRef: &zephyr_core_types.ResourceRef{
						Name:      deployment.Name,
						Namespace: deployment.Namespace,
						Cluster:   clusterName,
					},
					Labels:             nil,
					ServiceAccountName: "",
				},
				Mesh: &zephyr_core_types.ResourceRef{
					Name:      meshList.Items[0].GetName(),
					Namespace: meshList.Items[0].GetNamespace(),
					Cluster:   clusterName,
				},
			},
		}
		meshWorkload, err := meshWorkloadScanner.ScanPod(ctx, pod, clusterName)
		Expect(err).NotTo(HaveOccurred())
		Expect(meshWorkload).To(Equal(expectedMeshWorkload))
	})

	It("should return nil if not appmesh injected pod", func() {
		nonAppMeshPod := &corev1.Pod{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{Image: "random-image"},
				},
			},
			ObjectMeta: metav1.ObjectMeta{ClusterName: clusterName, Namespace: namespace},
		}
		meshWorkload, err := meshWorkloadScanner.ScanPod(ctx, nonAppMeshPod, clusterName)
		Expect(err).NotTo(HaveOccurred())
		Expect(meshWorkload).To(BeNil())
	})

	It("should return error if error fetching deployment", func() {
		expectedErr := eris.New("error")
		mockOwnerFetcher.EXPECT().GetDeployment(ctx, pod).Return(nil, expectedErr)
		_, err := meshWorkloadScanner.ScanPod(ctx, pod, clusterName)
		Expect(err).To(testutils.HaveInErrorChain(expectedErr))
	})
})
