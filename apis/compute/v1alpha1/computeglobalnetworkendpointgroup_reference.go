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
// See the License for the_identity.go specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeGlobalNetworkEndpointGroupRef{}

var ComputeGlobalNetworkEndpointGroupGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "ComputeGlobalNetworkEndpointGroup",
}

// ComputeGlobalNetworkEndpointGroupRef is a reference to a ComputeGlobalNetworkEndpointGroup.
type ComputeGlobalNetworkEndpointGroupRef struct {
	// A reference to an externally managed ComputeGlobalNetworkEndpointGroup resource.
	// Should be in the format "projects/{{projectID}}/global/networkEndpointGroups/{{networkEndpointGroupID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeGlobalNetworkEndpointGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeGlobalNetworkEndpointGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeGlobalNetworkEndpointGroupRef{})
}

func (r *ComputeGlobalNetworkEndpointGroupRef) GetGVK() schema.GroupVersionKind {
	return ComputeGlobalNetworkEndpointGroupGVK
}

func (r *ComputeGlobalNetworkEndpointGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeGlobalNetworkEndpointGroupRef) GetExternal() string {
	return r.External
}

func (r *ComputeGlobalNetworkEndpointGroupRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeGlobalNetworkEndpointGroupRef) ValidateExternal(ref string) error {
	id := &ComputeGlobalNetworkEndpointGroupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeGlobalNetworkEndpointGroupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeGlobalNetworkEndpointGroupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeGlobalNetworkEndpointGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			return selfLink
		}

		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}

		resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		if resourceID == "" {
			resourceID = u.GetName()
		}

		id := &ComputeGlobalNetworkEndpointGroupIdentity{
			Project:              projectID,
			NetworkEndpointGroup: resourceID,
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
