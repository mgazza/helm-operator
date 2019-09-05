// +build !ignore_autogenerated

/*
Copyright 2018-2019 The Flux CD contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartFileSelector) DeepCopyInto(out *ChartFileSelector) {
	*out = *in
	if in.Optional != nil {
		in, out := &in.Optional, &out.Optional
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartFileSelector.
func (in *ChartFileSelector) DeepCopy() *ChartFileSelector {
	if in == nil {
		return nil
	}
	out := new(ChartFileSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSource) DeepCopyInto(out *ChartSource) {
	*out = *in
	if in.GitChartSource != nil {
		in, out := &in.GitChartSource, &out.GitChartSource
		*out = new(GitChartSource)
		**out = **in
	}
	if in.RepoChartSource != nil {
		in, out := &in.RepoChartSource, &out.RepoChartSource
		*out = new(RepoChartSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSource.
func (in *ChartSource) DeepCopy() *ChartSource {
	if in == nil {
		return nil
	}
	out := new(ChartSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSourceSelector) DeepCopyInto(out *ExternalSourceSelector) {
	*out = *in
	if in.Optional != nil {
		in, out := &in.Optional, &out.Optional
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSourceSelector.
func (in *ExternalSourceSelector) DeepCopy() *ExternalSourceSelector {
	if in == nil {
		return nil
	}
	out := new(ExternalSourceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitChartSource) DeepCopyInto(out *GitChartSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitChartSource.
func (in *GitChartSource) DeepCopy() *GitChartSource {
	if in == nil {
		return nil
	}
	out := new(GitChartSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRelease) DeepCopyInto(out *HelmRelease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRelease.
func (in *HelmRelease) DeepCopy() *HelmRelease {
	if in == nil {
		return nil
	}
	out := new(HelmRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmRelease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseCondition) DeepCopyInto(out *HelmReleaseCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseCondition.
func (in *HelmReleaseCondition) DeepCopy() *HelmReleaseCondition {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseList) DeepCopyInto(out *HelmReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmRelease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseList.
func (in *HelmReleaseList) DeepCopy() *HelmReleaseList {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseSpec) DeepCopyInto(out *HelmReleaseSpec) {
	*out = *in
	in.ChartSource.DeepCopyInto(&out.ChartSource)
	if in.ValueFileSecrets != nil {
		in, out := &in.ValueFileSecrets, &out.ValueFileSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ValuesFrom != nil {
		in, out := &in.ValuesFrom, &out.ValuesFrom
		*out = make([]ValuesFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.HelmValues.DeepCopyInto(&out.HelmValues)
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	in.Rollback.DeepCopyInto(&out.Rollback)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseSpec.
func (in *HelmReleaseSpec) DeepCopy() *HelmReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseStatus) DeepCopyInto(out *HelmReleaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HelmReleaseCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseStatus.
func (in *HelmReleaseStatus) DeepCopy() *HelmReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepoChartSource) DeepCopyInto(out *RepoChartSource) {
	*out = *in
	if in.ChartPullSecret != nil {
		in, out := &in.ChartPullSecret, &out.ChartPullSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepoChartSource.
func (in *RepoChartSource) DeepCopy() *RepoChartSource {
	if in == nil {
		return nil
	}
	out := new(RepoChartSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rollback) DeepCopyInto(out *Rollback) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rollback.
func (in *Rollback) DeepCopy() *Rollback {
	if in == nil {
		return nil
	}
	out := new(Rollback)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValuesFromSource) DeepCopyInto(out *ValuesFromSource) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(corev1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalSourceRef != nil {
		in, out := &in.ExternalSourceRef, &out.ExternalSourceRef
		*out = new(ExternalSourceSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ChartFileRef != nil {
		in, out := &in.ChartFileRef, &out.ChartFileRef
		*out = new(ChartFileSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValuesFromSource.
func (in *ValuesFromSource) DeepCopy() *ValuesFromSource {
	if in == nil {
		return nil
	}
	out := new(ValuesFromSource)
	in.DeepCopyInto(out)
	return out
}
