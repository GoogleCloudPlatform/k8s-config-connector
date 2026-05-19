// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeNodeGroupRef{}

func init() {
	refsv1beta1.Register(&ComputeNodeGroupRef{})
}

// ComputeNodeGroupRef defines the resource reference to ComputeNodeGroup, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeNodeGroupRef struct {
	// A reference to an externally managed ComputeNodeGroup resource. Should be in the format "projects/{{projectID}}/zones/{{zone}}/nodeGroups/{{nodeGroupID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeNodeGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeNodeGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeNodeGroupRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNodeGroup",
	}
}

func (r *ComputeNodeGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeNodeGroupRef) GetExternal() string {
	return r.External
}

func (r *ComputeNodeGroupRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeNodeGroupRef) ValidateExternal(ref string) error {
	id := &ComputeNodeGroupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeNodeGroupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeNodeGroupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeNodeGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		zone, _, _ := unstructured.NestedString(u.Object, "spec", "zone")
		if zone == "" {
			return ""
		}
		resourceID, _ := refsv1beta1.GetResourceID(u)

		id := &ComputeNodeGroupIdentity{
			Project:   projectID,
			Zone:      zone,
			NodeGroup: resourceID,
		}
		return id.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
