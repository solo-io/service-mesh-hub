// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for networking.enterprise.mesh.gloo.solo.io/v1alpha1 resources

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for WasmDeployment

func (in *WasmDeployment) DeepCopyInto(out *WasmDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *WasmDeployment) DeepCopy() *WasmDeployment {
	if in == nil {
		return nil
	}
	out := new(WasmDeployment)
	in.DeepCopyInto(out)
	return out
}

func (in *WasmDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *WasmDeploymentList) DeepCopyInto(out *WasmDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WasmDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *WasmDeploymentList) DeepCopy() *WasmDeploymentList {
	if in == nil {
		return nil
	}
	out := new(WasmDeploymentList)
	in.DeepCopyInto(out)
	return out
}

func (in *WasmDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}