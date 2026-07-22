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

package v1alpha1

import (
	"context"
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &StorageBucketObjectRef{}
var StorageBucketObjectGVK = GroupVersion.WithKind("StorageBucketObject")

// StorageBucketObjectRef is a reference to a StorageBucketObject.
type StorageBucketObjectRef struct {
	// A reference to an externally managed StorageBucketObject resource.
	// Should be in the format "projects/{{projectID}}/buckets/{{bucket}}/objects/{{object}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a StorageBucketObject resource.
	Name string `json:"name,omitempty"`

	// The namespace of a StorageBucketObject resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func (r *StorageBucketObjectRef) GetGVK() schema.GroupVersionKind {
	return StorageBucketObjectGVK
}

func (r *StorageBucketObjectRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *StorageBucketObjectRef) GetExternal() string {
	return r.External
}

func (r *StorageBucketObjectRef) SetExternal(ref string) {
	r.External = ref
}

func (r *StorageBucketObjectRef) ValidateExternal(ref string) error {
	id := &StorageBucketObjectIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *StorageBucketObjectRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", StorageBucketObjectGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
