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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeURLMapRef{}

var ComputeURLMapGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeURLMap",
}

// ComputeURLMapRef is a reference to a GCP ComputeURLMap.
type ComputeURLMapRef struct {
	// A reference to an externally managed ComputeURLMap resource.
	// Should be in the format "projects/{{projectID}}/global/urlMaps/{{urlMapID}}"
	// or "projects/{{projectID}}/regions/{{region}}/urlMaps/{{urlMapID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeURLMap resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeURLMap resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeURLMapRef{})
}

func (r *ComputeURLMapRef) GetGVK() schema.GroupVersionKind {
	return ComputeURLMapGVK
}

func (r *ComputeURLMapRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeURLMapRef) GetExternal() string {
	return r.External
}

func (r *ComputeURLMapRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeURLMapRef) ValidateExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	id := &ComputeURLMapIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return err
	}
	return nil
}

func (r *ComputeURLMapRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeURLMapIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeURLMapRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			trimmed := apirefs.TrimComputeURIPrefix(selfLink)
			id := &ComputeURLMapIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}

		obj, err := common.ToStructuredType[*ComputeURLMap](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeURLMapSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var _ refs.Ref = &UrlmapResourceRef{}

func (r *UrlmapResourceRef) GetGVK() schema.GroupVersionKind {
	return ComputeBackendServiceGVK
}

func (r *UrlmapResourceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *UrlmapResourceRef) GetExternal() string {
	return r.External
}

func (r *UrlmapResourceRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *UrlmapResourceRef) ValidateExternal(ref string) error {
	return nil
}

func (r *UrlmapResourceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		return nil
	}
	if r.Name == "" {
		return nil
	}
	ns := r.Namespace
	if ns == "" {
		ns = defaultNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: ns}

	// Try ComputeBackendService first
	backendService := &unstructured.Unstructured{}
	backendService.SetGroupVersionKind(ComputeBackendServiceGVK)
	if err := reader.Get(ctx, key, backendService); err == nil {
		selfLink, _, _ := unstructured.NestedString(backendService.Object, "status", "selfLink")
		if selfLink != "" {
			r.SetExternal(selfLink)
			return nil
		}
		externalRef, _, _ := unstructured.NestedString(backendService.Object, "status", "externalRef")
		if externalRef != "" {
			r.SetExternal(externalRef)
			return nil
		}
		projectID, err := refs.ResolveProjectID(ctx, reader, backendService)
		if err != nil {
			return fmt.Errorf("cannot resolve project ID for referenced %s %v: %w", backendService.GetKind(), key, err)
		}
		location, _, _ := unstructured.NestedString(backendService.Object, "spec", "location")
		if location == "" {
			location = "global"
		}
		name := backendService.GetName()
		if location == "global" {
			r.SetExternal(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/backendServices/%s", projectID, name))
		} else {
			r.SetExternal(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/backendServices/%s", projectID, location, name))
		}
		return nil
	} else if !apierrors.IsNotFound(err) {
		return fmt.Errorf("reading referenced %s %s: %w", ComputeBackendServiceGVK, key, err)
	}

	// Try ComputeBackendBucket
	backendBucket := &unstructured.Unstructured{}
	backendBucket.SetGroupVersionKind(ComputeBackendBucketGVK)
	if err := reader.Get(ctx, key, backendBucket); err == nil {
		selfLink, _, _ := unstructured.NestedString(backendBucket.Object, "status", "selfLink")
		if selfLink != "" {
			r.SetExternal(selfLink)
			return nil
		}
		externalRef, _, _ := unstructured.NestedString(backendBucket.Object, "status", "externalRef")
		if externalRef != "" {
			r.SetExternal(externalRef)
			return nil
		}
		projectID, err := refs.ResolveProjectID(ctx, reader, backendBucket)
		if err != nil {
			return fmt.Errorf("cannot resolve project ID for referenced %s %v: %w", backendBucket.GetKind(), key, err)
		}
		name := backendBucket.GetName()
		r.SetExternal(fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/global/backendBuckets/%s", projectID, name))
		return nil
	} else if !apierrors.IsNotFound(err) {
		return fmt.Errorf("reading referenced %s %s: %w", ComputeBackendBucketGVK, key, err)
	}

	return k8s.NewReferenceNotFoundError(ComputeBackendServiceGVK, key)
}
