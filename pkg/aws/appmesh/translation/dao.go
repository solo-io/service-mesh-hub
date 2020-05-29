package translation

import (
	"context"

	aws2 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appmesh"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	appmesh2 "github.com/solo-io/service-mesh-hub/pkg/aws/appmesh"
	"github.com/solo-io/service-mesh-hub/pkg/collections/sets"
	"github.com/solo-io/service-mesh-hub/pkg/selector"
	"k8s.io/apimachinery/pkg/labels"
)

type appmeshTranslationDao struct {
	meshServiceClient    zephyr_discovery.MeshServiceClient
	meshWorkloadClient   zephyr_discovery.MeshWorkloadClient
	acpClient            zephyr_networking.AccessControlPolicyClient
	resourceSelector     selector.ResourceSelector
	appmeshClientFactory appmesh2.AppmeshClientFactory
}

func NewAppmeshAccessControlDao(
	meshServiceClient zephyr_discovery.MeshServiceClient,
	meshWorkloadClient zephyr_discovery.MeshWorkloadClient,
	resourceSelector selector.ResourceSelector,
	appmeshClientFactory appmesh2.AppmeshClientFactory,
	acpClient zephyr_networking.AccessControlPolicyClient,
) AppmeshTranslationDao {
	return &appmeshTranslationDao{
		meshServiceClient:    meshServiceClient,
		meshWorkloadClient:   meshWorkloadClient,
		resourceSelector:     resourceSelector,
		appmeshClientFactory: appmeshClientFactory,
		acpClient:            acpClient,
	}
}

func (a *appmeshTranslationDao) GetAllServiceWorkloadPairsForMesh(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
) (map[*zephyr_discovery.MeshService][]*zephyr_discovery.MeshWorkload,
	map[*zephyr_discovery.MeshWorkload][]*zephyr_discovery.MeshService,
	error) {
	meshServices, err := a.listMeshServicesForMesh(ctx, mesh)
	if err != nil {
		return nil, nil, err
	}
	meshWorkloads, err := a.listMeshWorkloadsForMesh(ctx, mesh)
	if err != nil {
		return nil, nil, err
	}
	serviceToWorkloads, workloadToServices := a.buildServiceToWorkloadMaps(meshServices, meshWorkloads)
	return serviceToWorkloads, workloadToServices, nil
}

func (a *appmeshTranslationDao) GetWorkloadsToAllUpstreamServices(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
) (map[*zephyr_discovery.MeshWorkload]sets.MeshServiceSet, error) {
	meshServices, err := a.listMeshServicesForMesh(ctx, mesh)
	if err != nil {
		return nil, err
	}
	meshWorkloads, err := a.listMeshWorkloadsForMesh(ctx, mesh)
	if err != nil {
		return nil, err
	}
	meshServiceSet := sets.NewMeshServiceSet(meshServices...)
	workloadsToAllUpstreamServices := map[*zephyr_discovery.MeshWorkload]sets.MeshServiceSet{}
	for _, meshWorkload := range meshWorkloads {
		workloadsToAllUpstreamServices[meshWorkload] = meshServiceSet
	}
	return workloadsToAllUpstreamServices, nil
}

func (a *appmeshTranslationDao) GetServicesWithACP(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
) (sets.MeshServiceSet, error) {
	meshServices, err := a.listMeshServicesForMesh(ctx, mesh)
	if err != nil {
		return nil, err
	}
	servicesInMesh := sets.NewMeshServiceSet(meshServices...)
	acpList, err := a.acpClient.ListAccessControlPolicy(ctx)
	if err != nil {
		return nil, err
	}
	services := sets.NewMeshServiceSet()
	for _, acp := range acpList.Items {
		acpServices, err := a.resourceSelector.GetAllMeshServicesByServiceSelector(ctx, acp.Spec.GetDestinationSelector())
		if err != nil {
			return nil, err
		}
		acpServicesSet := sets.NewMeshServiceSet(acpServices...)
		services.Insert(servicesInMesh.Intersection(acpServicesSet).List()...)
	}
	return services, nil
}

func (a *appmeshTranslationDao) GetWorkloadsToUpstreamServicesWithACP(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
) (sets.MeshWorkloadSet, map[*zephyr_discovery.MeshWorkload]sets.MeshServiceSet, error) {
	meshWorkloads, err := a.listMeshWorkloadsForMesh(ctx, mesh)
	if err != nil {
		return nil, nil, err
	}
	workloadsInMesh := sets.NewMeshWorkloadSet(meshWorkloads...)
	workloadsToUpstreamServices := map[*zephyr_discovery.MeshWorkload]sets.MeshServiceSet{}
	acpList, err := a.acpClient.ListAccessControlPolicy(ctx)
	if err != nil {
		return nil, nil, err
	}
	declaredWorkloads := sets.NewMeshWorkloadSet()
	for _, acp := range acpList.Items {
		workloads, err := a.resourceSelector.GetMeshWorkloadsByIdentitySelector(ctx, acp.Spec.GetSourceSelector())
		upstreamServices, err := a.resourceSelector.GetAllMeshServicesByServiceSelector(ctx, acp.Spec.GetDestinationSelector())
		if err != nil {
			return nil, nil, err
		}
		for _, workload := range workloads {
			upstreamServicesSet, ok := workloadsToUpstreamServices[workload]
			if !ok {
				workloadsToUpstreamServices[workload] = sets.NewMeshServiceSet(upstreamServices...)
			} else {
				upstreamServicesSet.Insert(upstreamServices...)
			}
		}
		acpWorkloadsSet := sets.NewMeshWorkloadSet(workloads...)
		declaredWorkloads.Insert(workloadsInMesh.Intersection(acpWorkloadsSet).List()...)
	}
	return declaredWorkloads, workloadsToUpstreamServices, nil
}

func (a *appmeshTranslationDao) buildServiceToWorkloadMaps(
	meshServices []*zephyr_discovery.MeshService,
	meshWorkloads []*zephyr_discovery.MeshWorkload,
) (map[*zephyr_discovery.MeshService][]*zephyr_discovery.MeshWorkload,
	map[*zephyr_discovery.MeshWorkload][]*zephyr_discovery.MeshService) {
	serviceToWorkloads := map[*zephyr_discovery.MeshService][]*zephyr_discovery.MeshWorkload{}
	workloadToServices := map[*zephyr_discovery.MeshWorkload][]*zephyr_discovery.MeshService{}
	for _, workload := range meshWorkloads {
		workloadToServices[workload] = nil
		for _, service := range meshServices {
			if isServiceBackedByWorkload(service, workload) {
				serviceToWorkloads[service] = append(serviceToWorkloads[service], workload)
				workloadToServices[workload] = append(workloadToServices[workload], service)
			}
		}
	}
	for _, service := range meshServices {
		_, ok := serviceToWorkloads[service]
		if !ok {
			serviceToWorkloads[service] = nil
		}
	}
	return serviceToWorkloads, workloadToServices
}

func (a *appmeshTranslationDao) listMeshServicesForMesh(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
) ([]*zephyr_discovery.MeshService, error) {
	meshServiceList, err := a.meshServiceClient.ListMeshService(ctx)
	if err != nil {
		return nil, err
	}
	var meshServices []*zephyr_discovery.MeshService
	for _, meshService := range meshServiceList.Items {
		meshService := meshService
		if meshService.Spec.GetMesh().GetName() != mesh.GetName() || meshService.Spec.GetMesh().GetNamespace() != mesh.GetNamespace() {
			continue
		}
		meshServices = append(meshServices, &meshService)
	}
	return meshServices, nil
}

func (a *appmeshTranslationDao) listMeshWorkloadsForMesh(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
) ([]*zephyr_discovery.MeshWorkload, error) {
	meshWorkloadList, err := a.meshWorkloadClient.ListMeshWorkload(ctx)
	if err != nil {
		return nil, err
	}
	var meshWorkloads []*zephyr_discovery.MeshWorkload
	for _, meshWorkload := range meshWorkloadList.Items {
		meshWorkload := meshWorkload
		if meshWorkload.Spec.GetMesh().GetName() != mesh.GetName() || meshWorkload.Spec.GetMesh().GetNamespace() != mesh.GetNamespace() {
			continue
		}
		meshWorkloads = append(meshWorkloads, &meshWorkload)
	}
	return meshWorkloads, nil
}

func (a *appmeshTranslationDao) EnsureVirtualService(
	mesh *zephyr_discovery.Mesh,
	virtualServiceData *appmesh.VirtualServiceData,
) error {
	appmeshClient, err := a.appmeshClientFactory(mesh)
	if err != nil {
		return err
	}
	return appmeshClient.EnsureVirtualService(virtualServiceData)
}

func (a *appmeshTranslationDao) EnsureVirtualRouter(
	mesh *zephyr_discovery.Mesh,
	virtualRouter *appmesh.VirtualRouterData,
) error {
	appmeshClient, err := a.appmeshClientFactory(mesh)
	if err != nil {
		return err
	}
	return appmeshClient.EnsureVirtualRouter(virtualRouter)
}

func (a *appmeshTranslationDao) EnsureRoute(
	mesh *zephyr_discovery.Mesh,
	route *appmesh.RouteData,
) error {
	appmeshClient, err := a.appmeshClientFactory(mesh)
	if err != nil {
		return err
	}
	return appmeshClient.EnsureRoute(route)
}

func (a *appmeshTranslationDao) EnsureVirtualNode(
	mesh *zephyr_discovery.Mesh,
	virtualNode *appmesh.VirtualNodeData,
) error {
	appmeshClient, err := a.appmeshClientFactory(mesh)
	if err != nil {
		return err
	}
	return appmeshClient.EnsureVirtualNode(virtualNode)
}

func (a *appmeshTranslationDao) ReconcileVirtualServices(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
	virtualServices []*appmesh.VirtualServiceData,
) error {
	appmeshClient, err := a.appmeshClientFactory(mesh)
	if err != nil {
		return err
	}
	meshName := aws2.String(mesh.Spec.GetAwsAppMesh().GetName())
	return appmeshClient.ReconcileVirtualServices(ctx, meshName, virtualServices)
}

func (a *appmeshTranslationDao) ReconcileVirtualRouters(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
	virtualRouters []*appmesh.VirtualRouterData,
) error {
	appmeshClient, err := a.appmeshClientFactory(mesh)
	if err != nil {
		return err
	}
	meshName := aws2.String(mesh.Spec.GetAwsAppMesh().GetName())
	return appmeshClient.ReconcileVirtualRouters(ctx, meshName, virtualRouters)
}

func (a *appmeshTranslationDao) ReconcileRoutes(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
	routes []*appmesh.RouteData,
) error {
	appmeshClient, err := a.appmeshClientFactory(mesh)
	if err != nil {
		return err
	}
	meshName := aws2.String(mesh.Spec.GetAwsAppMesh().GetName())
	return appmeshClient.ReconcileRoutes(ctx, meshName, routes)
}

func (a *appmeshTranslationDao) ReconcileVirtualNodes(
	ctx context.Context,
	mesh *zephyr_discovery.Mesh,
	virtualNodes []*appmesh.VirtualNodeData,
) error {
	appmeshClient, err := a.appmeshClientFactory(mesh)
	if err != nil {
		return err
	}
	meshName := aws2.String(mesh.Spec.GetAwsAppMesh().GetName())
	return appmeshClient.ReconcileVirtualNodes(ctx, meshName, virtualNodes)
}

func isServiceBackedByWorkload(
	meshService *zephyr_discovery.MeshService,
	meshWorkload *zephyr_discovery.MeshWorkload,
) bool {
	if len(meshService.Spec.GetKubeService().GetWorkloadSelectorLabels()) == 0 ||
		len(meshWorkload.Spec.GetKubeController().GetLabels()) == 0 {
		return false
	}
	return labels.AreLabelsInWhiteList(
		meshService.Spec.GetKubeService().GetWorkloadSelectorLabels(),
		meshWorkload.Spec.GetKubeController().GetLabels(),
	)
}
