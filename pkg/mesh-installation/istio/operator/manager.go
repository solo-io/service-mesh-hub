package operator

import (
	"strings"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/versionutils"
	"github.com/solo-io/service-mesh-hub/pkg/container-runtime/docker"
	k8s_apps "k8s.io/api/apps/v1"
)

var (
	FailedToParseInstallManifest = func(err error) error {
		return eris.Wrap(err, "Failed to parse the templated Mesh operator manifest.")
	}
	FailedToInstallOperator = func(err error) error {
		return eris.Wrap(err, "Failed to install the Mesh installation operator. Artifacts from the failed installation will be cleaned up.")
	}
	FailedToCleanFailedInstallation = func(err error) error {
		return eris.Wrap(err, "Failed to clean up the failed Mesh operator installation.")
	}
	UnrecognizedOperatorInstance    = eris.New("This instance of the Mesh operator is not recognized - cannot verify its version matches what you specified")
	FailedToGenerateInstallManifest = func(err error) error {
		return eris.Wrap(err, "Install manifest template failed to render. This shouldn't happen")
	}
	FailedToValidateExistingOperator = func(err error, clusterName, namespace string) error {
		return eris.Wrapf(err, "Failed to validate existing Mesh operator deployment in cluster %s namespace %s", clusterName, namespace)
	}
	FailedToDetermineOperatorVersion = func(current string, minimumSupportedVersion versionutils.Version) error {
		return eris.Errorf("Failed to determine whether the current operator running at version %s is the minimum version %s", current, minimumSupportedVersion.String())
	}
	IncompatibleOperatorVersion = func(current string, minimumSupportedVersion string) error {
		return eris.Errorf("Found istio operator running at version %s, while %s is the minimum supported version", current, minimumSupportedVersion)
	}
	FailedToCheckIfOperatorExists = func(err error, clusterName, namespace string) error {
		return eris.Wrapf(err, "Failed to check whether the Mesh operator is already installed to cluster %s in namespace %s", clusterName, namespace)
	}

	istioVersionToMinimumOperatorVersion = map[IstioVersion]versionutils.Version{
		IstioVersion1_5: {
			Major: 1,
			Minor: 5,
			Patch: 1,
		},
		IstioVersion1_6: {
			Major: 1,
			Minor: 6,
			Patch: 0,
		},
	}
)

type OperatorManagerFactory func(
	operatorDao OperatorDao,
	manifestBuilder InstallerManifestBuilder,
	imageNameParser docker.ImageNameParser,
) OperatorManager

func NewOperatorManagerFactory() OperatorManagerFactory {
	return NewManager
}

func NewManager(
	operatorDao OperatorDao,
	manifestBuilder InstallerManifestBuilder,
	imageNameParser docker.ImageNameParser,
) OperatorManager {
	return &manager{
		operatorDao:     operatorDao,
		manifestBuilder: manifestBuilder,
		imageNameParser: imageNameParser,
	}
}

type manager struct {
	manifestBuilder InstallerManifestBuilder
	operatorDao     OperatorDao
	imageNameParser docker.ImageNameParser
}

func (m *manager) InstallOperatorApplication(installationOptions *InstallationOptions) error {
	manifest, err := m.InstallDryRun(installationOptions)
	if err != nil {
		return err
	}

	err = m.operatorDao.ApplyManifest(m.getNamespaceFromOptions(installationOptions), manifest)
	if err != nil {
		return FailedToInstallOperator(err)
	}

	return nil
}

func (m *manager) InstallDryRun(installationOptions *InstallationOptions) (manifest string, err error) {
	namespace := m.getNamespaceFromOptions(installationOptions)

	applicationManifest, err := m.manifestBuilder.BuildOperatorDeploymentManifest(installationOptions.IstioVersion, namespace, installationOptions.CreateNamespace)
	if err != nil {
		return "", FailedToGenerateInstallManifest(err)
	}

	manifest += "\n---\n" + applicationManifest

	var operatorConfig string
	if installationOptions.InstallationProfile != "" {
		operatorConfig, err = m.manifestBuilder.BuildOperatorConfigurationWithProfile(installationOptions.InstallationProfile, namespace)
		if err != nil {
			return "", err
		}
	} else if installationOptions.CustomIstioOperatorSpec != "" {
		operatorConfig = installationOptions.CustomIstioOperatorSpec
	} else {
		operatorConfig, err = m.manifestBuilder.BuildOperatorConfigurationWithProfile(DefaultIstioProfile, namespace)
		if err != nil {
			return "", err
		}
	}

	manifest += "\n---\n" + operatorConfig

	return manifest, nil
}

func (m *manager) ValidateOperatorNamespace(istioVersion IstioVersion, operatorName, operatorNamespace, clusterName string) (installNeeded bool, err error) {
	operatorDeployment, err := m.operatorDao.FindOperatorDeployment(operatorName, operatorNamespace)
	if err != nil {
		return false, FailedToCheckIfOperatorExists(err, clusterName, operatorNamespace)
	} else if operatorDeployment == nil {
		return true, nil
	}

	err = m.validateExistingOperatorDeployment(istioVersion, operatorDeployment)
	if err != nil {
		return false, FailedToValidateExistingOperator(err, clusterName, operatorNamespace)
	}

	return false, nil
}

func (*manager) getNamespaceFromOptions(options *InstallationOptions) string {
	namespace := options.InstallNamespace
	if namespace == "" {
		namespace = DefaultIstioOperatorNamespace
	}

	return namespace
}

func (m *manager) validateExistingOperatorDeployment(istioVersion IstioVersion, deployment *k8s_apps.Deployment) error {
	containers := deployment.Spec.Template.Spec.Containers
	if len(containers) != 1 {
		return UnrecognizedOperatorInstance
	}

	image, err := m.imageNameParser.Parse(containers[0].Image)
	if err != nil {
		return err
	}

	actualImageVersion := image.Tag
	if actualImageVersion == "" {
		actualImageVersion = image.Digest
	}

	if !strings.HasPrefix(actualImageVersion, "v") {
		actualImageVersion = "v" + actualImageVersion
	}

	version, err := versionutils.ParseVersion(actualImageVersion)
	if err != nil {
		return err
	}

	minimumSupportedVersion, ok := istioVersionToMinimumOperatorVersion[istioVersion]
	if !ok {
		return eris.Errorf("Istio version %s does not have a minimum supported version; this is unexpected", string(istioVersion))
	}
	greater, ok := version.IsGreaterThan(minimumSupportedVersion)
	if !ok {
		return FailedToDetermineOperatorVersion(version.String(), minimumSupportedVersion)
	}

	if !(greater || version.Equals(&minimumSupportedVersion)) {
		return IncompatibleOperatorVersion(version.String(), minimumSupportedVersion.String())
	}

	return nil
}
