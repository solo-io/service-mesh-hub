// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for ratelimit.api.solo.io.mesh.gloo.solo.io/v1alpha1 resources

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for RateLimiterConfig

func (in *RateLimiterConfig) DeepCopyInto(out *RateLimiterConfig) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *RateLimiterConfig) DeepCopy() *RateLimiterConfig {
	if in == nil {
		return nil
	}
	out := new(RateLimiterConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *RateLimiterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *RateLimiterConfigList) DeepCopyInto(out *RateLimiterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimiterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *RateLimiterConfigList) DeepCopy() *RateLimiterConfigList {
	if in == nil {
		return nil
	}
	out := new(RateLimiterConfigList)
	in.DeepCopyInto(out)
	return out
}

func (in *RateLimiterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
