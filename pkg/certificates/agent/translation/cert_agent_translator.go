package translation

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/agent/input"
	"github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/agent/output/certagent"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/certificates/agent/utils"
	"github.com/solo-io/gloo-mesh/pkg/certificates/common/secrets"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// the map key used to store the agent's private key in a kube Secret
	privateKeySecretKey = "private-key"
)

func agentLabels() map[string]string {
	labels := map[string]string{
		fmt.Sprintf("agent.%v", v1.SchemeGroupVersion.Group): defaults.GetPodNamespace(),
	}
	if agentCluster := defaults.GetAgentCluster(); agentCluster != "" {
		labels[metautils.AgentLabelKey] = agentCluster
	}
	return labels
}

func PrivateKeySecretType() corev1.SecretType {
	return corev1.SecretType(fmt.Sprintf("%s/generated_private_key", v1.SchemeGroupVersion.Group))
}

func IssuedCertificateSecretType() corev1.SecretType {
	return corev1.SecretType(fmt.Sprintf("%s/issued_certificate", v1.SchemeGroupVersion.Group))
}

//go:generate mockgen -source ./cert_agent_translator.go -destination mocks/translator.go

// These functions correspond to issued certiticate statuses
// PENDING
// REQUESTED
// ISSUED
// FINISHED
type Translator interface {
	// This function is called when the IssuedCertiticate is first created, it is meant to create the CSR
	// and return the bytes directly
	// If both return values are nil, it signals to the reconciler to ignore this resource.
	IssuedCertiticatePending(
		ctx context.Context,
		issuedCertificate *v1.IssuedCertificate,
		inputs input.Snapshot,
		outputs certagent.Builder,
	) ([]byte, error)

	// This function is called when the IssuedCertiticate has been REQUESTED, this means that the
	// CertificateRequest has been successfully created, and is passed in along with the
	// IssuedCertificate
	// This function may be called multiple times in a row, as long as the IssuedCertificate is in a
	// REQUESTED state
	// If err == nil, and continueIterating == false it signals to the reconciler to ignore this resource.
	IssuedCertificateRequested(
		ctx context.Context,
		issuedCertificate *v1.IssuedCertificate,
		certificateRequest *v1.CertificateRequest,
		inputs input.Snapshot,
		outputs certagent.Builder,
	) (continueIterating bool, err error)

	// This function is called when the IssuedCertiticate has been ISSUED, this means that the
	// CertificateRequest has been fulfilled by the relevant party. By this stage the intermediate
	// cert should already have made it to it's destination.
	// If err == nil, and continueIterating == false it signals to the reconciler to ignore this resource.
	IssuedCertificateIssued(
		ctx context.Context,
		issuedCertificate *v1.IssuedCertificate,
		inputs input.Snapshot,
		outputs certagent.Builder,
	) (continueIterating bool, err error)
}

func NewCertAgentTranslator(localClient client.Client) Translator {
	return &certAgentTranslator{
		localClient: localClient,
	}
}

type certAgentTranslator struct {
	localClient client.Client
}

func (c *certAgentTranslator) IssuedCertiticatePending(
	ctx context.Context,
	issuedCertificate *v1.IssuedCertificate,
	inputs input.Snapshot,
	outputs certagent.Builder,
) ([]byte, error) {

	// If IssuedCertificateSecret is nil, it means the keypair isn't meant to be written to a secret
	if issuedCertificate.Spec.GetIssuedCertificateSecret() == nil {
		return nil, nil
	}

	// create a new private key
	privateKey, err := utils.GeneratePrivateKey()
	if err != nil {
		return nil, err
	}
	outputs.AddSecrets(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      issuedCertificate.Name,
			Namespace: issuedCertificate.Namespace,
			Labels:    agentLabels(),
		},
		Data: map[string][]byte{privateKeySecretKey: privateKey},
		Type: PrivateKeySecretType(),
	})

	// create certificate request for private key
	csrBytes, err := utils.GenerateCertificateSigningRequest(
		issuedCertificate.Spec.Hosts,
		issuedCertificate.Spec.Org,
		privateKey,
	)
	if err != nil {
		return nil, err
	}
	return csrBytes, err
}

func (c *certAgentTranslator) IssuedCertificateRequested(
	ctx context.Context,
	issuedCertificate *v1.IssuedCertificate,
	certificateRequest *v1.CertificateRequest,
	inputs input.Snapshot,
	outputs certagent.Builder,
) (continueIterating bool, err error) {

	// If IssuedCertificateSecret is nil, it means the keypair isn't meant to be written to a secret
	if issuedCertificate.Spec.GetIssuedCertificateSecret() == nil {
		return false, nil
	}

	// retrieve private key
	privateKeySecret, err := inputs.Secrets().Find(issuedCertificate)
	if err != nil {
		return true, err
	}
	privateKey := privateKeySecret.Data[privateKeySecretKey]

	if len(privateKey) == 0 {
		return true, eris.Errorf("invalid private key found, no data provided")
	}

	switch certificateRequest.Status.State {
	case v1.CertificateRequestStatus_PENDING:
		contextutils.LoggerFrom(ctx).Infof("waiting for certificate request %v to be signed by Issuer", sets.Key(certificateRequest))

		// add secret and certrequest to output to prevent them from being GC'ed
		outputs.AddSecrets(privateKeySecret)
		outputs.AddCertificateRequests(certificateRequest)

		// if the certificate signing request has not been
		// fulfilled, return and wait for the next reconcile
		return true, nil
	case v1.CertificateRequestStatus_FAILED:
		return true, eris.Errorf("certificate request %v failed to be signed by Issuer: %v", sets.Key(certificateRequest), certificateRequest.Status.Error)
	}

	signedCert := certificateRequest.Status.SignedCertificate
	signingRootCA := certificateRequest.Status.SigningRootCa

	issuedCertificateData := secrets.IntermediateCAData{
		RootCAData: secrets.RootCAData{
			RootCert: signingRootCA,
		},
		CertChain:    utils.AppendRootCerts(signedCert, signingRootCA),
		CaCert:       signedCert,
		CaPrivateKey: privateKey,
	}

	// finally, create the secret
	issuedCertificateSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      issuedCertificate.Spec.IssuedCertificateSecret.Name,
			Namespace: issuedCertificate.Spec.IssuedCertificateSecret.Namespace,
			Labels:    agentLabels(),
		},
		Data: issuedCertificateData.ToSecretData(),
		Type: IssuedCertificateSecretType(),
	}
	outputs.AddSecrets(issuedCertificateSecret)
	return true, nil
}

func (c *certAgentTranslator) IssuedCertificateIssued(
	ctx context.Context,
	issuedCertificate *v1.IssuedCertificate,
	inputs input.Snapshot,
	outputs certagent.Builder,
) (continueIterating bool, err error) {

	// If IssuedCertificateSecret is nil, it means the keypair isn't meant to be written to a secret
	if issuedCertificate.Spec.GetIssuedCertificateSecret() == nil {
		return false, nil
	}

	// ensure issued cert secret exists, if not, return an error (restart the workflow)
	if issuedCertificateSecret, err := inputs.Secrets().Find(
		issuedCertificate.Spec.GetIssuedCertificateSecret(),
	); err != nil {
		return true, err
	} else {
		// add secret output to prevent it from being GC'ed
		outputs.AddSecrets(issuedCertificateSecret)
	}
	return true, nil
}
