package e2e_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/test/e2e"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = FDescribe("SMH e2e", func() {
	var (
		ctx           = context.Background()
		appNamespace  = "bookinfo"
		dynamicClient client.Client
		//accessPolicyClient v1alpha2.AccessPolicyClient
		virtualMeshClient v1alpha2.VirtualMeshClient
	)

	BeforeEach(func() {
		var err error
		dynamicClient, err = client.New(e2e.GetEnv().Management.Config, client.Options{})
		//accessPolicyClient = v1alpha2.NewAccessPolicyClient(dynamicClient)
		virtualMeshClient = v1alpha2.NewVirtualMeshClient(dynamicClient)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
	})

	It("should enforce access policy when enabled", func() {
		virtualMesh := &v1alpha2.VirtualMesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "enforcement-enabled",
				Namespace: appNamespace,
			},
			Spec: v1alpha2.VirtualMeshSpec{
				Meshes: []*v1.ObjectRef{
					{
						Name:      "istiod-istio-system-master-cluster",
						Namespace: "service-mesh-hub",
					},
				},
				EnforceAccessControl: v1alpha2.VirtualMeshSpec_ENABLED,
			},
		}
		err := virtualMeshClient.CreateVirtualMesh(ctx, virtualMesh)
		Expect(err).ToNot(HaveOccurred())

	})
})
