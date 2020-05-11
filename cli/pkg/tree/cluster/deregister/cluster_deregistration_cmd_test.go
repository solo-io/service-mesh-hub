package deregister_test

import (
	"context"
	"fmt"
	"os"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	cli_mocks "github.com/solo-io/service-mesh-hub/cli/pkg/mocks"
	cli_test "github.com/solo-io/service-mesh-hub/cli/pkg/test"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/deregister"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	mock_registration "github.com/solo-io/service-mesh-hub/pkg/clients/cluster-registration/mocks"
	"github.com/solo-io/service-mesh-hub/pkg/env"
	mock_kubeconfig "github.com/solo-io/service-mesh-hub/pkg/kubeconfig/mocks"
	mock_core "github.com/solo-io/service-mesh-hub/test/mocks/clients/discovery.zephyr.solo.io/v1alpha1"
	mock_kubernetes_core "github.com/solo-io/service-mesh-hub/test/mocks/clients/kubernetes/core/v1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("ClusterDeregistrationCmd", func() {
	var (
		ctrl                            *gomock.Controller
		ctx                             context.Context
		secretClient                    *mock_kubernetes_core.MockSecretClient
		mockKubeLoader                  *mock_kubeconfig.MockKubeLoader
		meshctl                         *cli_test.MockMeshctl
		configVerifier                  *cli_mocks.MockMasterKubeConfigVerifier
		mockClusterDeregistrationClient *mock_registration.MockClusterDeregistrationClient
		mockKubeClusterClient           *mock_core.MockKubernetesClusterClient
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()

		secretClient = mock_kubernetes_core.NewMockSecretClient(ctrl)
		mockKubeLoader = mock_kubeconfig.NewMockKubeLoader(ctrl)
		configVerifier = cli_mocks.NewMockMasterKubeConfigVerifier(ctrl)
		mockKubeClusterClient = mock_core.NewMockKubernetesClusterClient(ctrl)
		mockClusterDeregistrationClient = mock_registration.NewMockClusterDeregistrationClient(ctrl)
		meshctl = &cli_test.MockMeshctl{
			KubeClients: common.KubeClients{
				SecretClient:                secretClient,
				ClusterDeregistrationClient: mockClusterDeregistrationClient,
				KubeClusterClient:           mockKubeClusterClient,
			},
			Clients: common.Clients{
				MasterClusterVerifier: configVerifier,
			},
			MockController: ctrl,
			KubeLoader:     mockKubeLoader,
			Ctx:            ctx,
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	var expectVerifyMasterCluster = func(localKubeConfigPath string) {
		configVerifier.EXPECT().Verify(localKubeConfigPath, "").Return(nil)
		localKubeConfig := &rest.Config{}
		mockKubeLoader.EXPECT().GetRestConfigForContext(localKubeConfigPath, "").Return(localKubeConfig, nil)
	}

	It("should work", func() {
		remoteClusterName := "remote-cluster-name"
		localKubeConfigPath := "~/.kube/master-config"
		os.Setenv("KUBECONFIG", localKubeConfigPath)
		defer os.Setenv("KUBECONFIG", "")
		expectVerifyMasterCluster(localKubeConfigPath)
		kubeCluster := &v1alpha1.KubernetesCluster{}
		mockKubeClusterClient.
			EXPECT().
			GetKubernetesCluster(ctx, client.ObjectKey{Name: remoteClusterName, Namespace: env.GetWriteNamespace()}).
			Return(kubeCluster, nil)
		mockClusterDeregistrationClient.EXPECT().Deregister(ctx, kubeCluster).Return(nil)

		stdout, err := meshctl.Invoke(fmt.Sprintf("cluster deregister --remote-cluster-name %s", remoteClusterName))
		Expect(err).NotTo(HaveOccurred())
		Expect(stdout).To(Equal(fmt.Sprintf("Successfully deregistered cluster %s.\n", remoteClusterName)))
	})

	It("should work", func() {
		remoteClusterName := "remote-cluster-name"
		localKubeConfigPath := "~/.kube/master-config"
		os.Setenv("KUBECONFIG", localKubeConfigPath)
		defer os.Setenv("KUBECONFIG", "")
		expectVerifyMasterCluster(localKubeConfigPath)
		testErr := eris.New("test error")
		mockKubeClusterClient.
			EXPECT().
			GetKubernetesCluster(ctx, client.ObjectKey{Name: remoteClusterName, Namespace: env.GetWriteNamespace()}).
			Return(nil, testErr)

		stdout, err := meshctl.Invoke(fmt.Sprintf("cluster deregister --remote-cluster-name %s", remoteClusterName))
		Expect(err).To(testutils.HaveInErrorChain(deregister.KubeClusterNotFound(remoteClusterName, testErr)))
		Expect(stdout).To(Equal(fmt.Sprintf("Error deregistering cluster %s.\n", remoteClusterName)))
	})
})
