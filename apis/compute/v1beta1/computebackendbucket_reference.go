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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeBackendBucketRef{}

// ComputeBackendBucketRef is a reference to a ComputeBackendBucket.
type ComputeBackendBucketRef struct {
	// A reference to an externally managed StorageBucketObject resource.
	// Should be in the format "projects/{{project}}/zones/{{zone}}/backendBuckets/{{backendBucket}}".
	External string `json:"external,omitempty"`

	// The name of a StorageBucketObject resource.
	Name string `json:"name,omitempty"`

	// The namespace of a StorageBucketObject resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeBackendBucketRef) GetGVK() schema.GroupVersionKind {
	return ComputeBackendBucketGVK
}

func (r *ComputeBackendBucketRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *ComputeBackendBucketRef) GetExternal() string {
	return r.External
}

func (r *ComputeBackendBucketRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeBackendBucketRef) ValidateExternal(ref string) error {
	id := &ComputeBackendBucketIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeBackendBucketRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			trimmed := apirefs.TrimComputeURIPrefix(selfLink)
			id := &ComputeAddressIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}

		obj, err := common.ToStructuredType[*ComputeAddress](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeAddressSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
