// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"context"
	"github.com/solo-io/go-utils/installutils/helminstall"
	"github.com/solo-io/reporting-client/pkg/client"
	"github.com/solo-io/service-mesh-hub/cli/pkg"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/config"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/exec"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/files"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/interactive"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/kube"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/resource_printing"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/usage"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check/healthcheck"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check/status"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/deregister"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/register"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/register/csr"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/create"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/create/access-control-policy"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/create/traffic-policy"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/create/virtualmesh"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/demo"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/demo/cleanup"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/demo/init"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/describe"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/describe/description"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get/cluster"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get/mesh"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get/service"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get/virtual_mesh"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get/vmcsr"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get/workload"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/install"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh/install"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh/install/istio/operator"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/config_lookup"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/crd"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/helm"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/upgrade"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/upgrade/assets"
	version2 "github.com/solo-io/service-mesh-hub/cli/pkg/tree/version"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/version/server"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/apiextensions.k8s.io/v1beta1"
	v1_2 "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/apps/v1"
	"github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	v1alpha1_3 "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	v1alpha1_2 "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/auth"
	"github.com/solo-io/service-mesh-hub/pkg/clients/kubernetes/discovery"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
	"github.com/solo-io/service-mesh-hub/pkg/selector"
	"github.com/solo-io/service-mesh-hub/pkg/version"
	"github.com/spf13/cobra"
	"io"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Injectors from wire.go:

func DefaultKubeClientsFactory(masterConfig *rest.Config, writeNamespace string) (*common.KubeClients, error) {
	clientset, err := v1.ClientsetFromConfigProvider(masterConfig)
	if err != nil {
		return nil, err
	}
	secretClient := v1.SecretClientFromClientsetProvider(clientset)
	serviceAccountClient := v1.ServiceAccountClientFromClientsetProvider(clientset)
	remoteAuthorityConfigCreator := auth.NewRemoteAuthorityConfigCreator(secretClient, serviceAccountClient)
	kubernetesClientset, err := kubernetes.NewForConfig(masterConfig)
	if err != nil {
		return nil, err
	}
	rbacClient := auth.RbacClientProvider(kubernetesClientset)
	remoteAuthorityManager := auth.NewRemoteAuthorityManager(serviceAccountClient, rbacClient)
	clusterAuthorization := auth.NewClusterAuthorization(remoteAuthorityConfigCreator, remoteAuthorityManager)
	helmClient := helminstall.DefaultHelmClient()
	installer := install.HelmInstallerProvider(helmClient, kubernetesClientset)
	v1alpha1Clientset, err := v1alpha1.ClientsetFromConfigProvider(masterConfig)
	if err != nil {
		return nil, err
	}
	kubernetesClusterClient := v1alpha1.KubernetesClusterClientFromClientsetProvider(v1alpha1Clientset)
	namespaceClient := v1.NamespaceClientFromClientsetProvider(clientset)
	serverVersionClient := kubernetes_discovery.NewGeneratedServerVersionClient(kubernetesClientset)
	podClient := v1.PodClientFromClientsetProvider(clientset)
	meshServiceClient := v1alpha1.MeshServiceClientFromClientsetProvider(v1alpha1Clientset)
	clients := healthcheck.ClientsProvider(namespaceClient, serverVersionClient, podClient, meshServiceClient)
	v1Clientset, err := v1_2.ClientsetFromConfigProvider(masterConfig)
	if err != nil {
		return nil, err
	}
	deploymentClient := v1_2.DeploymentClientFromClientsetProvider(v1Clientset)
	imageNameParser := docker.NewImageNameParser()
	deployedVersionFinder := version.NewDeployedVersionFinder(deploymentClient, imageNameParser)
	customResourceDefinitionClientFromConfigFactory := v1beta1.CustomResourceDefinitionClientFromConfigFactoryProvider()
	crdRemover := crd_uninstall.NewCrdRemover(customResourceDefinitionClientFromConfigFactory)
	fileReader := files.NewDefaultFileReader()
	converter := kube.NewConverter(fileReader)
	uninstallClients := common.UninstallClientsProvider(crdRemover, converter)
	inMemoryRESTClientGetterFactory := common_config.NewInMemoryRESTClientGetterFactory()
	uninstallerFactory := helm_uninstall.NewUninstallerFactory()
	kubeConfigLookup := config_lookup.NewKubeConfigLookup(kubernetesClusterClient, secretClient, converter)
	secretClientFactory := v1.SecretClientFactoryProvider()
	dynamicClientGetter := config_lookup.NewDynamicClientGetter(kubeConfigLookup)
	clusterDeregistrationClient := deregister.NewClusterDeregistrationClient(crdRemover, inMemoryRESTClientGetterFactory, uninstallerFactory, kubeConfigLookup, kubernetesClusterClient, secretClient, secretClientFactory, dynamicClientGetter)
	clientset2, err := v1alpha1_2.ClientsetFromConfigProvider(masterConfig)
	if err != nil {
		return nil, err
	}
	virtualMeshCertificateSigningRequestClient := v1alpha1_2.VirtualMeshCertificateSigningRequestClientFromClientsetProvider(clientset2)
	meshClient := v1alpha1.MeshClientFromClientsetProvider(v1alpha1Clientset)
	clientset3, err := v1alpha1_3.ClientsetFromConfigProvider(masterConfig)
	if err != nil {
		return nil, err
	}
	virtualMeshClient := v1alpha1_3.VirtualMeshClientFromClientsetProvider(clientset3)
	trafficPolicyClient := v1alpha1_3.TrafficPolicyClientFromClientsetProvider(clientset3)
	accessControlPolicyClient := v1alpha1_3.AccessControlPolicyClientFromClientsetProvider(clientset3)
	meshWorkloadClient := v1alpha1.MeshWorkloadClientFromClientsetProvider(v1alpha1Clientset)
	deploymentClientFactory := v1_2.DeploymentClientFactoryProvider()
	resourceSelector := selector.NewResourceSelector(meshServiceClient, meshWorkloadClient, deploymentClientFactory, dynamicClientGetter)
	resourceDescriber := description.NewResourceDescriber(trafficPolicyClient, accessControlPolicyClient, resourceSelector)
	kubeClients := common.KubeClientsProvider(clusterAuthorization, installer, helmClient, kubernetesClusterClient, clients, deployedVersionFinder, customResourceDefinitionClientFromConfigFactory, secretClient, namespaceClient, uninstallClients, inMemoryRESTClientGetterFactory, clusterDeregistrationClient, kubeConfigLookup, virtualMeshCertificateSigningRequestClient, meshServiceClient, meshClient, virtualMeshClient, resourceDescriber, resourceSelector, trafficPolicyClient, accessControlPolicyClient, meshWorkloadClient)
	return kubeClients, nil
}

func DefaultClientsFactory(opts *options.Options) (*common.Clients, error) {
	kubeLoader := common_config.DefaultKubeLoaderProvider(opts)
	imageNameParser := docker.NewImageNameParser()
	serverVersionClient := server.DefaultServerVersionClientProvider(opts, kubeLoader, imageNameParser)
	githubAssetClient := upgrade_assets.DefaultGithubAssetClient()
	assetHelper := upgrade_assets.NewAssetHelper(githubAssetClient)
	masterKubeConfigVerifier := common_config.NewMasterKubeConfigVerifier(kubeLoader)
	unstructuredKubeClientFactory := kube.NewUnstructuredKubeClientFactory()
	deploymentClient := server.NewDeploymentClient(kubeLoader, opts)
	installerManifestBuilder := operator.NewInstallerManifestBuilder()
	operatorManagerFactory := operator.NewOperatorManagerFactory()
	istioClients := common.IstioClientsProvider(installerManifestBuilder, operatorManagerFactory)
	statusClientFactory := status.StatusClientFactoryProvider()
	healthCheckSuite := healthcheck.DefaultHealthChecksProvider()
	csrAgentInstallerFactory := csr.NewCsrAgentInstallerFactory()
	clusterRegistrationClients := common.ClusterRegistrationClientsProvider(csrAgentInstallerFactory)
	fileReader := files.NewDefaultFileReader()
	converter := kube.NewConverter(fileReader)
	clients := common.ClientsProvider(serverVersionClient, assetHelper, masterKubeConfigVerifier, unstructuredKubeClientFactory, deploymentClient, istioClients, statusClientFactory, healthCheckSuite, clusterRegistrationClients, converter)
	return clients, nil
}

func InitializeCLI(ctx context.Context, out io.Writer, in io.Reader) *cobra.Command {
	optionsOptions := options.NewOptionsProvider()
	client := usage.DefaultUsageReporterProvider()
	kubeClientsFactory := DefaultKubeClientsFactoryProvider()
	clientsFactory := DefaultClientsFactoryProvider()
	kubeLoader := common_config.DefaultKubeLoaderProvider(optionsOptions)
	registrationCmd := register.ClusterRegistrationCmd(ctx, kubeClientsFactory, clientsFactory, optionsOptions, out, kubeLoader)
	clusterCommand := cluster.ClusterRootCmd(registrationCmd)
	versionCommand := version2.VersionCmd(out, clientsFactory, optionsOptions)
	imageNameParser := docker.NewImageNameParser()
	fileReader := files.NewDefaultFileReader()
	meshInstallCommand := mesh_install.MeshInstallRootCmd(clientsFactory, optionsOptions, out, in, kubeLoader, imageNameParser, fileReader)
	meshCommand := mesh.MeshRootCmd(meshInstallCommand)
	upgradeCommand := upgrade.UpgradeCmd(ctx, optionsOptions, out, clientsFactory)
	installCommand := install.InstallCmd(ctx, optionsOptions, kubeClientsFactory, clientsFactory, kubeLoader, out)
	uninstallCommand := uninstall.UninstallCmd(ctx, out, optionsOptions, kubeClientsFactory, kubeLoader)
	prettyPrinter := status.NewPrettyPrinter()
	jsonPrinter := status.NewJsonPrinter()
	checkCommand := check.CheckCmd(ctx, out, optionsOptions, kubeClientsFactory, clientsFactory, kubeLoader, prettyPrinter, jsonPrinter)
	tableBuilder := table_printing.DefaultTableBuilder()
	meshPrinter := table_printing.NewMeshPrinter(tableBuilder)
	meshServicePrinter := table_printing.NewMeshServicePrinter(tableBuilder)
	meshWorkloadPrinter := table_printing.NewMeshWorkloadPrinter(tableBuilder)
	kubernetesClusterPrinter := table_printing.NewKubernetesClusterPrinter(tableBuilder)
	trafficPolicyPrinter := table_printing.NewTrafficPolicyPrinter(tableBuilder)
	accessControlPolicyPrinter := table_printing.NewAccessControlPolicyPrinter(tableBuilder)
	virtualMeshPrinter := table_printing.NewVirtualMeshPrinter(tableBuilder)
	virtualMeshCSRPrinter := table_printing.NewVirtualMeshMCSRPrinter(tableBuilder)
	resourcePrinter := resource_printing.NewResourcePrinter()
	printers := common.PrintersProvider(meshPrinter, meshServicePrinter, meshWorkloadPrinter, kubernetesClusterPrinter, trafficPolicyPrinter, accessControlPolicyPrinter, virtualMeshPrinter, virtualMeshCSRPrinter, resourcePrinter)
	describeCommand := describe.DescribeCmd(ctx, kubeLoader, kubeClientsFactory, printers, optionsOptions, out)
	runner := exec.NewShellRunner(in, out)
	initCmd := demo_init.DemoInitCmd(ctx, runner)
	cleanupCmd := cleanup.DemoCleanupCmd(ctx, runner)
	demoCommand := demo.DemoRootCmd(initCmd, cleanupCmd)
	getMeshCommand := get_mesh.GetMeshRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getWorkloadCommand := get_workload.GetWorkloadRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getServiceCommand := get_service.GetServiceRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getClusterCommand := get_cluster.GetClusterRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getVirtualMeshCommand := get_virtual_mesh.GetVirtualMeshRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getVirtualMeshCSRCommand := get_vmcsr.GetVirtualMeshCSRRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getCommand := get.GetRootCommand(getMeshCommand, getWorkloadCommand, getServiceCommand, getClusterCommand, getVirtualMeshCommand, getVirtualMeshCSRCommand, optionsOptions)
	interactivePrompt := interactive.NewSurveyInteractivePrompt()
	createVirtualMeshCmd := virtualmesh.CreateVirtualMeshCommand(ctx, out, optionsOptions, kubeLoader, kubeClientsFactory, interactivePrompt, printers)
	createTrafficPolicyCmd := traffic_policy.CreateTrafficPolicyCommand(ctx, out, optionsOptions, kubeLoader, kubeClientsFactory, interactivePrompt, printers)
	createAccessControlPolicyCmd := access_control_policy.CreateAccessControlPolicyCommand(ctx, out, optionsOptions, kubeLoader, kubeClientsFactory, interactivePrompt, printers)
	createRootCmd := create.CreateRootCommand(optionsOptions, createVirtualMeshCmd, createTrafficPolicyCmd, createAccessControlPolicyCmd)
	command := cli.BuildCli(ctx, optionsOptions, client, clusterCommand, versionCommand, meshCommand, upgradeCommand, installCommand, uninstallCommand, checkCommand, describeCommand, demoCommand, getCommand, createRootCmd)
	return command
}

func InitializeCLIWithMocks(ctx context.Context, out io.Writer, in io.Reader, usageClient client.Client, kubeClientsFactory common.KubeClientsFactory, clientsFactory common.ClientsFactory, kubeLoader common_config.KubeLoader, imageNameParser docker.ImageNameParser, fileReader files.FileReader, kubeconfigConverter kube.Converter, printers common.Printers, runner exec.Runner, interactivePrompt interactive.InteractivePrompt) *cobra.Command {
	optionsOptions := options.NewOptionsProvider()
	registrationCmd := register.ClusterRegistrationCmd(ctx, kubeClientsFactory, clientsFactory, optionsOptions, out, kubeLoader)
	clusterCommand := cluster.ClusterRootCmd(registrationCmd)
	versionCommand := version2.VersionCmd(out, clientsFactory, optionsOptions)
	meshInstallCommand := mesh_install.MeshInstallRootCmd(clientsFactory, optionsOptions, out, in, kubeLoader, imageNameParser, fileReader)
	meshCommand := mesh.MeshRootCmd(meshInstallCommand)
	upgradeCommand := upgrade.UpgradeCmd(ctx, optionsOptions, out, clientsFactory)
	installCommand := install.InstallCmd(ctx, optionsOptions, kubeClientsFactory, clientsFactory, kubeLoader, out)
	uninstallCommand := uninstall.UninstallCmd(ctx, out, optionsOptions, kubeClientsFactory, kubeLoader)
	prettyPrinter := status.NewPrettyPrinter()
	jsonPrinter := status.NewJsonPrinter()
	checkCommand := check.CheckCmd(ctx, out, optionsOptions, kubeClientsFactory, clientsFactory, kubeLoader, prettyPrinter, jsonPrinter)
	describeCommand := describe.DescribeCmd(ctx, kubeLoader, kubeClientsFactory, printers, optionsOptions, out)
	initCmd := demo_init.DemoInitCmd(ctx, runner)
	cleanupCmd := cleanup.DemoCleanupCmd(ctx, runner)
	demoCommand := demo.DemoRootCmd(initCmd, cleanupCmd)
	getMeshCommand := get_mesh.GetMeshRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getWorkloadCommand := get_workload.GetWorkloadRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getServiceCommand := get_service.GetServiceRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getClusterCommand := get_cluster.GetClusterRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getVirtualMeshCommand := get_virtual_mesh.GetVirtualMeshRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getVirtualMeshCSRCommand := get_vmcsr.GetVirtualMeshCSRRootCommand(ctx, out, printers, kubeClientsFactory, kubeLoader, optionsOptions)
	getCommand := get.GetRootCommand(getMeshCommand, getWorkloadCommand, getServiceCommand, getClusterCommand, getVirtualMeshCommand, getVirtualMeshCSRCommand, optionsOptions)
	createVirtualMeshCmd := virtualmesh.CreateVirtualMeshCommand(ctx, out, optionsOptions, kubeLoader, kubeClientsFactory, interactivePrompt, printers)
	createTrafficPolicyCmd := traffic_policy.CreateTrafficPolicyCommand(ctx, out, optionsOptions, kubeLoader, kubeClientsFactory, interactivePrompt, printers)
	createAccessControlPolicyCmd := access_control_policy.CreateAccessControlPolicyCommand(ctx, out, optionsOptions, kubeLoader, kubeClientsFactory, interactivePrompt, printers)
	createRootCmd := create.CreateRootCommand(optionsOptions, createVirtualMeshCmd, createTrafficPolicyCmd, createAccessControlPolicyCmd)
	command := cli.BuildCli(ctx, optionsOptions, usageClient, clusterCommand, versionCommand, meshCommand, upgradeCommand, installCommand, uninstallCommand, checkCommand, describeCommand, demoCommand, getCommand, createRootCmd)
	return command
}
