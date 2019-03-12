package clientset

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClientsetFromContext", func() {
	It("initializes all the clients in the Clientset struct", func() {

		cs, err := ClientsetFromContext(context.TODO())
		Expect(err).NotTo(HaveOccurred())
		Expect(cs.Input).NotTo(BeNil())
		Expect(cs.Input.RoutingRule).NotTo(BeNil())
		Expect(cs.Input.Upstream).NotTo(BeNil())
		Expect(cs.Input.Install).NotTo(BeNil())
		Expect(cs.Input.Mesh).NotTo(BeNil())
		Expect(cs.Input.MeshGroup).NotTo(BeNil())
		Expect(cs.Input.SecurityRule).NotTo(BeNil())
		Expect(cs.Input.TlsSecret).NotTo(BeNil())
		Expect(cs.Discovery).NotTo(BeNil())
		Expect(cs.Discovery.Pod).NotTo(BeNil())
	})
})

var _ = Describe("IstioClientsFromContext", func() {
	It("initializes all the clients in the Clientset struct", func() {

		istio, err := IstioFromContext(context.TODO())
		Expect(err).NotTo(HaveOccurred())
		Expect(istio).NotTo(BeNil())
		Expect(istio.RbacConfig).NotTo(BeNil())
		Expect(istio.ServiceRole).NotTo(BeNil())
		Expect(istio.MeshPolicy).NotTo(BeNil())
		Expect(istio.DestinationRule).NotTo(BeNil())
		Expect(istio.ServiceRoleBinding).NotTo(BeNil())
		Expect(istio.VirtualService).NotTo(BeNil())

	})
})
