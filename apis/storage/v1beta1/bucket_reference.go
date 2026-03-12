// Copyright 2025 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &StorageBucketRef{}
var StorageBucketGVK = GroupVersion.WithKind("StorageBucket")

func init() {
	refsv1beta1.Register(&StorageBucketRef{})
}

// StorageBucketRef defines the resource reference to StorageBucket, which "External" field
// holds the GCP identifier for the KRM object.
type StorageBucketRef struct {
	// For backward compatibility, we are not enforcing the external format.

	// A reference to an externally-managed StorageBucket resource.
	External string `json:"external,omitempty"`

	// The name of a StorageBucket resource.
	Name string `json:"name,omitempty"`

	// The namespace of a StorageBucket resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *StorageBucketRef) GetGVK() schema.GroupVersionKind {
	return StorageBucketGVK
}

func (r *StorageBucketRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *StorageBucketRef) GetExternal() string {
	return r.External
}

func (r *StorageBucketRef) SetExternal(ref string) {
	r.External = ref
}

func (r *StorageBucketRef) ValidateExternal(ref string) error {
	id := &StorageBucketIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *StorageBucketRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &StorageBucketIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *StorageBucketRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// Backward compatible to Terraform/DCL based resource, which does not have status.externalRef.
		resourceID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return ""
		}

		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}

		return fmt.Sprintf("projects/%s/buckets/%s", projectID, resourceID)
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
