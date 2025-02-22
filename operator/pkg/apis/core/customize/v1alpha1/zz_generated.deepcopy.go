//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerResourceSpec) DeepCopyInto(out *ContainerResourceSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerResourceSpec.
func (in *ContainerResourceSpec) DeepCopy() *ContainerResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResource) DeepCopyInto(out *ControllerResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResource.
func (in *ControllerResource) DeepCopy() *ControllerResource {
	if in == nil {
		return nil
	}
	out := new(ControllerResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceList) DeepCopyInto(out *ControllerResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceList.
func (in *ControllerResourceList) DeepCopy() *ControllerResourceList {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceSpec) DeepCopyInto(out *ControllerResourceSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerResourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceSpec.
func (in *ControllerResourceSpec) DeepCopy() *ControllerResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResourceStatus) DeepCopyInto(out *ControllerResourceStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResourceStatus.
func (in *ControllerResourceStatus) DeepCopy() *ControllerResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedControllerResource) DeepCopyInto(out *NamespacedControllerResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedControllerResource.
func (in *NamespacedControllerResource) DeepCopy() *NamespacedControllerResource {
	if in == nil {
		return nil
	}
	out := new(NamespacedControllerResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedControllerResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedControllerResourceList) DeepCopyInto(out *NamespacedControllerResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespacedControllerResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedControllerResourceList.
func (in *NamespacedControllerResourceList) DeepCopy() *NamespacedControllerResourceList {
	if in == nil {
		return nil
	}
	out := new(NamespacedControllerResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespacedControllerResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedControllerResourceSpec) DeepCopyInto(out *NamespacedControllerResourceSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerResourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedControllerResourceSpec.
func (in *NamespacedControllerResourceSpec) DeepCopy() *NamespacedControllerResourceSpec {
	if in == nil {
		return nil
	}
	out := new(NamespacedControllerResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedControllerResourceStatus) DeepCopyInto(out *NamespacedControllerResourceStatus) {
	*out = *in
	in.CommonStatus.DeepCopyInto(&out.CommonStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedControllerResourceStatus.
func (in *NamespacedControllerResourceStatus) DeepCopy() *NamespacedControllerResourceStatus {
	if in == nil {
		return nil
	}
	out := new(NamespacedControllerResourceStatus)
	in.DeepCopyInto(out)
	return out
}
