package istio

import (
	"context"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	istiov1 "github.com/solo-io/supergloo/pkg/api/external/istio/encryption/v1"
	"github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/pkg/translator/kube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type EncryptionSyncer struct {
	Kube         kubernetes.Interface
	SecretClient istiov1.IstioCacertsSecretClient
	ctx          context.Context
}

const (
	defaultIstioNamespace            = "istio-system"
	customRootCertificateSecretName  = "cacerts"
	defaultRootCertificateSecretName = "istio.default"
	istioLabelKey                    = "istio"
	citadelLabelValue                = "citadel"
)

func (s *EncryptionSyncer) Sync(ctx context.Context, snap *v1.TranslatorSnapshot) error {
	ctx = contextutils.WithLogger(ctx, "istio-encryption-syncer")
	logger := contextutils.LoggerFrom(ctx)
	logger.Infof("begin sync %v (%v meshes, %v certs)", snap.Hash(),
		len(snap.Meshes), len(snap.Istiocerts))
	defer logger.Infof("end sync %v", snap.Hash())
	logger.Debugf("%v", snap)

	s.ctx = ctx
	for _, mesh := range snap.Meshes.List() {
		if err := s.syncMesh(ctx, mesh, snap); err != nil {
			return err
		}
	}
	return nil
}

func (s *EncryptionSyncer) syncMesh(ctx context.Context, mesh *v1.Mesh, snap *v1.TranslatorSnapshot) error {
	if mesh.TargetMesh == nil {
		// ilackarms (todo): reporting for this error
		return errors.Errorf("invalid mesh %v: target_mesh required", mesh.Metadata.Ref())
	}
	if mesh.TargetMesh.MeshType != v1.MeshType_ISTIO {
		return nil
	}
	encryption := mesh.Encryption
	if encryption == nil {
		return nil
	}
	encryptionSecret := encryption.Secret
	if encryptionSecret == nil {
		return nil
	}
	secretList := snap.Istiocerts.List()
	sourceSecret, err := secretList.Find(encryptionSecret.Namespace, encryptionSecret.Name)
	if err != nil {
		return errors.Wrapf(err, "Error finding secret referenced in mesh config (%s:%s)",
			encryptionSecret.Namespace, encryptionSecret.Name)
	}
	// this is where custom root certs will live once configured, if not found existingSecret will be nil
	existingSecret, _ := secretList.Find(defaultIstioNamespace, customRootCertificateSecretName)

	return s.syncSecret(ctx, sourceSecret, existingSecret)
}

func (s *EncryptionSyncer) syncSecret(ctx context.Context, sourceSecret, existingSecret *istiov1.IstioCacertsSecret) error {
	if err := validateTlsSecret(sourceSecret); err != nil {
		return errors.Wrapf(err, "invalid secret %v", sourceSecret.Metadata.Ref())
	}
	istioSecret := resources.Clone(sourceSecret).(*istiov1.IstioCacertsSecret)
	if existingSecret == nil {
		istioSecret.Metadata = core.Metadata{
			Namespace: defaultIstioNamespace,
			Name:      defaultRootCertificateSecretName,
		}
		if _, err := s.SecretClient.Write(istioSecret, clients.WriteOpts{
			Ctx: ctx,
		}); err != nil {
			return errors.Wrapf(err, "creating tool tls secret %v for istio", istioSecret.Metadata.Ref())
		}
		return nil
	}

	// move secret over to destination name/namespace
	istioSecret.SetMetadata(existingSecret.Metadata)
	istioSecret.Metadata.Annotations["created_by"] = "supergloo"
	// nothing to do
	if istioSecret.Equal(existingSecret) {
		return nil
	}
	if _, err := s.SecretClient.Write(istioSecret, clients.WriteOpts{
		Ctx: ctx,
	}); err != nil {
		return errors.Wrapf(err, "updating tool tls secret %v for istio", istioSecret.Metadata.Ref())
	}
	return nil
}

func validateTlsSecret(secret *istiov1.IstioCacertsSecret) error {
	if secret.RootCert == "" {
		return errors.Errorf("Root cert is missing.")
	}
	if secret.CaKey == "" {
		return errors.Errorf("Private key is missing.")
	}
	if secret.CertChain == "" {
		return errors.Errorf("Cert Chain is missing.")
	}
	if secret.CaCert == "" {
		return errors.Errorf("Cert Chain is missing.")
	}
	return nil
}

func (s *EncryptionSyncer) deleteIstioDefaultSecret() error {
	// Using Kube API directly cause we don't expect this secret to be tagged and it should be mostly a one-time op
	return s.Kube.CoreV1().Secrets(defaultIstioNamespace).Delete(defaultRootCertificateSecretName, &metav1.DeleteOptions{})
}

func (s *EncryptionSyncer) restartCitadel() error {
	selector := make(map[string]string)
	selector[istioLabelKey] = citadelLabelValue
	return kube.RestartPods(s.Kube, defaultIstioNamespace, selector)
}
