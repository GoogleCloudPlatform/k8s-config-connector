// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"context"

	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeInstanceTemplateRef{}

func init() {
	refs.Register(&ComputeInstanceTemplateRef{}, nil)
	refs.Register(&VersionsInstanceTemplateRef{}, nil)
	refs.Register(&ComputeHealthCheckRef{}, nil)
}

func (r *ComputeInstanceTemplateRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeInstanceTemplate",
	}
}

func (r *ComputeInstanceTemplateRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeInstanceTemplateRef) GetExternal() string {
	return r.External
}

func (r *ComputeInstanceTemplateRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeInstanceTemplateRef) ValidateExternal(ref string) error {
	return nil
}

func (r *ComputeInstanceTemplateRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			return apirefs.TrimComputeURIPrefix(selfLink)
		}
		return ""
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var _ refs.Ref = &VersionsInstanceTemplateRef{}

func (r *VersionsInstanceTemplateRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeInstanceTemplate",
	}
}

func (r *VersionsInstanceTemplateRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VersionsInstanceTemplateRef) GetExternal() string {
	return r.External
}

func (r *VersionsInstanceTemplateRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VersionsInstanceTemplateRef) ValidateExternal(ref string) error {
	return nil
}

func (r *VersionsInstanceTemplateRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			return apirefs.TrimComputeURIPrefix(selfLink)
		}
		return ""
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var _ refs.Ref = &ComputeHealthCheckRef{}

func (r *ComputeHealthCheckRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeHealthCheck",
	}
}

func (r *ComputeHealthCheckRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeHealthCheckRef) GetExternal() string {
	return r.External
}

func (r *ComputeHealthCheckRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeHealthCheckRef) ValidateExternal(ref string) error {
	return nil
}

func (r *ComputeHealthCheckRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			return apirefs.TrimComputeURIPrefix(selfLink)
		}
		return ""
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
