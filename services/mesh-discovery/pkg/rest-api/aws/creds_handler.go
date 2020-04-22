package aws

import (
	"context"
	"time"

	"github.com/rotisserie/eris"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/aws_creds"
	"github.com/solo-io/service-mesh-hub/pkg/logging"
	rest_api "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/rest-api"
	k8s_core_types "k8s.io/api/core/v1"
)

const (
	ReconcileIntervalSeconds = 1
)

var (
	AppMeshClientNotFound = func(apiName string) error {
		return eris.Errorf("AppMesh client not found for API: %s", apiName)
	}
)

type awsCredsHandler struct {
	secretConverter          aws_creds.SecretAwsCredsConverter
	appMeshReconcilerFactory AppMeshReconcilerFactory
}

type AwsCredsHandler rest_api.RestAPICredsHandler

func NewAwsCredsHandler(
	secretAwsCredsConverter aws_creds.SecretAwsCredsConverter,
	appMeshReconcilerFactory AppMeshReconcilerFactory,
) AwsCredsHandler {
	return &awsCredsHandler{
		secretConverter:          secretAwsCredsConverter,
		appMeshReconcilerFactory: appMeshReconcilerFactory,
	}
}

func (a *awsCredsHandler) RestAPIAdded(ctx context.Context, secret *k8s_core_types.Secret) error {
	logger := logging.BuildEventLogger(ctx, logging.GenericEvent, secret)
	// Periodically run API Reconciler to ensure AppMesh state is consistent with SMH
	ticker := time.NewTicker(ReconcileIntervalSeconds * time.Second)
	err := a.appMeshReconcilerFactory(secret).Reconcile(ctx)
	if err != nil {
		logger.Error(err)
	}
	go func() {
		for {
			select {
			case <-ticker.C:
				logger.Debugf("Reconciling AppMesh with secret %s.%s", secret.GetName(), secret.GetNamespace())
				err := a.appMeshReconcilerFactory(secret).Reconcile(ctx)
				if err != nil {
					logger.Error(err)
				}
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
	return nil
}

func (a *awsCredsHandler) RestAPIRemoved(ctx context.Context, apiName string) error {
	//logger := logging.BuildEventLogger(ctx, logging.DeleteEvent, nil)
	//logger.Warn("Secret removal not implemented for AwsCredsHandler")
	return nil
}
