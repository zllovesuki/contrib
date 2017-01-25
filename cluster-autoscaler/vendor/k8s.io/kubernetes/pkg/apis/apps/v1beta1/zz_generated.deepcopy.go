// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StatefulSet, InType: reflect.TypeOf(&StatefulSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StatefulSetList, InType: reflect.TypeOf(&StatefulSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StatefulSetSpec, InType: reflect.TypeOf(&StatefulSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StatefulSetStatus, InType: reflect.TypeOf(&StatefulSetStatus{})},
	)
}

func DeepCopy_v1beta1_StatefulSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSet)
		out := out.(*StatefulSet)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_StatefulSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_StatefulSetStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_StatefulSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetList)
		out := out.(*StatefulSetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]StatefulSet, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_StatefulSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_StatefulSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetSpec)
		out := out.(*StatefulSetSpec)
		*out = *in
		if in.Replicas != nil {
			in, out := &in.Replicas, &out.Replicas
			*out = new(int32)
			**out = **in
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.VolumeClaimTemplates != nil {
			in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
			*out = make([]api_v1.PersistentVolumeClaim, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_PersistentVolumeClaim(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_StatefulSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetStatus)
		out := out.(*StatefulSetStatus)
		*out = *in
		if in.ObservedGeneration != nil {
			in, out := &in.ObservedGeneration, &out.ObservedGeneration
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}
