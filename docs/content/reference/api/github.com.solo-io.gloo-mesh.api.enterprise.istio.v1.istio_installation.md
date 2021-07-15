
---

title: "istio_installation.proto"

---

## Package : `istio.enterprise.mesh.gloo.solo.io`



<a name="top"></a>

<a name="API Reference for istio_installation.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## istio_installation.proto


## Table of Contents
  - [IstioInstallationSpec](#istio.enterprise.mesh.gloo.solo.io.IstioInstallationSpec)
  - [IstioInstallationStatus](#istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus)
  - [IstioInstallationStatus.IstioOperatorStatus](#istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus)
  - [IstioInstallationStatus.StatusesEntry](#istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.StatusesEntry)

  - [IstioInstallationStatus.IstioOperatorStatus.State](#istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus.State)






<a name="istio.enterprise.mesh.gloo.solo.io.IstioInstallationSpec"></a>

### IstioInstallationSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterName | string |  | The cluster where the IstioOperators should be installed |
  | istioOperatorRefs | [][core.skv2.solo.io.ObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ObjectRef" >}}) | repeated | References for the IstioOperators that should be installed in the managed cluster |
  





<a name="istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus"></a>

### IstioInstallationStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observedGeneration | int64 |  | The most recent generation observed in the the IstioInstallation metadata. If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh has not processed the most recent version of this resource. |
  | statuses | [][istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.StatusesEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.enterprise.istio.v1.istio_installation#istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.StatusesEntry" >}}) | repeated | The status of the installation for each IstioOperator that should be applied. |
  





<a name="istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus"></a>

### IstioInstallationStatus.IstioOperatorStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observedGeneration | int64 |  | The most recent generation observed in the the IstioOperator metadata. If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh has not processed the most recent version of this resource. |
  | state | [istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus.State]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.enterprise.istio.v1.istio_installation#istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus.State" >}}) |  | The current state of the IstioOperator |
  | message | string |  | A human readable message about the current state of the IstioOperator |
  





<a name="istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.StatusesEntry"></a>

### IstioInstallationStatus.StatusesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | [istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.enterprise.istio.v1.istio_installation#istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus" >}}) |  |  |
  




 <!-- end messages -->


<a name="istio.enterprise.mesh.gloo.solo.io.IstioInstallationStatus.IstioOperatorStatus.State"></a>

### IstioInstallationStatus.IstioOperatorStatus.State
The state of a IstioOperator installation

| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 | Waiting for resources to be reconciled |
| RECONCILING | 1 | In the process of reconciling Istio resources on to the managed cluster |
| HEALTHY | 2 | All Istio components were installed successfully and they are healthy |
| ERROR | 3 | One or more installed Istio component(s) in an error state |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

