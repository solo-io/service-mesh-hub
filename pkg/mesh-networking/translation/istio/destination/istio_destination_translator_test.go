package destination

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	mock_output "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio/mocks"
	mock_reporting "github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting/mocks"
	mock_authorizationpolicy "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/destination/authorizationpolicy/mocks"
	mock_destinationrule "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/destination/destinationrule/mocks"
	mock_virtualservice "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/destination/virtualservice/mocks"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/security/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("IstioDestinationTranslator", func() {
	var (
		ctrl                              *gomock.Controller
		mockDestinationRuleTranslator     *mock_destinationrule.MockTranslator
		mockVirtualServiceTranslator      *mock_virtualservice.MockTranslator
		mockAuthorizationPolicyTranslator *mock_authorizationpolicy.MockTranslator
		mockOutputs                       *mock_output.MockBuilder
		mockReporter                      *mock_reporting.MockReporter
		istioDestinationTranslator        Translator
		ctx                               = context.TODO()
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockDestinationRuleTranslator = mock_destinationrule.NewMockTranslator(ctrl)
		mockVirtualServiceTranslator = mock_virtualservice.NewMockTranslator(ctrl)
		mockAuthorizationPolicyTranslator = mock_authorizationpolicy.NewMockTranslator(ctrl)
		mockOutputs = mock_output.NewMockBuilder(ctrl)
		mockReporter = mock_reporting.NewMockReporter(ctrl)
		istioDestinationTranslator = &translator{
			ctx:                   ctx,
			destinationRules:      mockDestinationRuleTranslator,
			virtualServices:       mockVirtualServiceTranslator,
			authorizationPolicies: mockAuthorizationPolicyTranslator,
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should translate", func() {
		destination := &v1alpha2.Destination{
			Spec: v1alpha2.DestinationSpec{
				Mesh: &v1.ObjectRef{
					Name:      "hello",
					Namespace: "world",
				},
			},
		}

		in := input.NewInputLocalSnapshotManualBuilder("").
			AddMeshes([]*v1alpha2.Mesh{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      destination.Spec.GetMesh().GetName(),
						Namespace: destination.Spec.GetMesh().GetNamespace(),
					},
					Spec: v1alpha2.MeshSpec{
						Type: &v1alpha2.MeshSpec_Istio_{
							Istio: &v1alpha2.MeshSpec_Istio{},
						},
					},
				},
			}).
			Build()

		vs := &v1alpha3.VirtualService{}
		dr := &v1alpha3.DestinationRule{}
		ap := &v1beta1.AuthorizationPolicy{}

		mockDestinationRuleTranslator.
			EXPECT().
			Translate(ctx, in, destination, nil, mockReporter).
			Return(dr)
		mockVirtualServiceTranslator.
			EXPECT().
			Translate(ctx, in, destination, nil, mockReporter).
			Return(vs)
		mockAuthorizationPolicyTranslator.
			EXPECT().
			Translate(in, destination, mockReporter).
			Return(ap)
		mockOutputs.
			EXPECT().
			AddVirtualServices(vs)
		mockOutputs.
			EXPECT().
			AddDestinationRules(dr)
		mockOutputs.
			EXPECT().
			AddAuthorizationPolicies(ap)

		istioDestinationTranslator.Translate(in, destination, mockOutputs, mockReporter)
	})
})
