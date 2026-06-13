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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeInstanceGroupRef{}

var ComputeInstanceGroupGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeInstanceGroup",
}

// ComputeInstanceGroupRef is a reference to a ComputeInstanceGroup.
type ComputeInstanceGroupRef struct {
	// A reference to an externally managed ComputeInstanceGroup resource.
	// Should be in the format "projects/{{projectID}}/zones/{{zone}}/instanceGroups/{{instanceGroupID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeInstanceGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeInstanceGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeInstanceGroupRef{})
}

func (r *ComputeInstanceGroupRef) GetGVK() schema.GroupVersionKind {
	return ComputeInstanceGroupGVK
}

func (r *ComputeInstanceGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeInstanceGroupRef) GetExternal() string {
	return r.External
}

func (r *ComputeInstanceGroupRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeInstanceGroupRef) ValidateExternal(ref string) error {
	id := &ComputeInstanceGroupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeInstanceGroupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeInstanceGroupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeInstanceGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
