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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeBackendBucketRef{}

// ComputeBackendBucketRef defines the resource reference to ComputeBackendBucket, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeBackendBucketRef struct {
	// A reference to an externally managed ComputeBackendBucket resource.
	// Should be of the format `projects/{{project}}/global/backendBuckets/{{backendBucket}}`.
	External string `json:"external,omitempty"`

	// The `name` of a `ComputeBackendBucket` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `ComputeBackendBucket` resource.
	Namespace string `json:"namespace,omitempty"`
}

// GetGVK returns the GroupVersionKind of the referenced resource.
func (r *ComputeBackendBucketRef) GetGVK() schema.GroupVersionKind {
	return ComputeBackendBucketGVK
}

// GetNamespacedName returns the NamespacedName of the referenced resource.
func (r *ComputeBackendBucketRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}

// GetExternal returns the External field value of the ComputeBackendBucketRef.
func (r *ComputeBackendBucketRef) GetExternal() string {
	return r.External
}

// SetExternal sets the External field value of the ComputeBackendBucketRef.
func (r *ComputeBackendBucketRef) SetExternal(ref string) {
	r.External = ref
}

// ValidateExternal returns nil if the external string is valid.
func (r *ComputeBackendBucketRef) ValidateExternal(ref string) error {
	id := &ComputeBackendBucketIdentity{}
	return id.FromExternal(ref)
}

// Normalize ensures the "External" reference (in string format) is set for a given Ref.
func (r *ComputeBackendBucketRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

var _ refs.Ref = &BackendBucketBucketRef{}

// GetGVK returns the GroupVersionKind of the referenced resource.
func (r *BackendBucketBucketRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "storage.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "StorageBucket",
	}
}

// GetNamespacedName returns the NamespacedName of the referenced resource.
func (r *BackendBucketBucketRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}

// GetExternal returns the External field value of the BackendBucketBucketRef.
func (r *BackendBucketBucketRef) GetExternal() string {
	return r.External
}

// SetExternal sets the External field value of the BackendBucketBucketRef.
func (r *BackendBucketBucketRef) SetExternal(ref string) {
	r.External = ref
}

// ValidateExternal returns nil if the external string is valid.
func (r *BackendBucketBucketRef) ValidateExternal(ref string) error {
	// StorageBucket external representation is just the bucket name.
	return nil
}

// Normalize ensures the "External" reference (in string format) is set for a given Ref.
func (r *BackendBucketBucketRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
