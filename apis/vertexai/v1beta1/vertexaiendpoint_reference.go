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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var VertexAIEndpointGVK = schema.GroupVersionKind{
	Group:   "vertexai.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "VertexAIEndpoint",
}

var _ refs.Ref = &VertexAIEndpointRef{}

// VertexAIEndpointRef is a reference to a VertexAIEndpoint.
type VertexAIEndpointRef struct {
	// A reference to an externally managed VertexAIEndpoint resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/endpoints/{{endpointID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIEndpoint resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIEndpoint resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAIEndpointRef{}, nil)
}

func (r *VertexAIEndpointRef) GetGVK() schema.GroupVersionKind {
	return VertexAIEndpointGVK
}

func (r *VertexAIEndpointRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIEndpointRef) GetExternal() string {
	return r.External
}

func (r *VertexAIEndpointRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAIEndpointRef) ValidateExternal(ref string) error {
	id := &VertexAIEndpointIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *VertexAIEndpointRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAIEndpointIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIEndpointRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// Verify the resource is ready
		ready := false
		conditions, found, err := unstructured.NestedSlice(u.Object, "status", "conditions")
		if err == nil && found {
			for _, cond := range conditions {
				condMap, ok := cond.(map[string]any)
				if !ok {
					continue
				}
				if condMap["type"] == "Ready" && condMap["status"] == "True" {
					ready = true
					break
				}
			}
		}
		if !ready {
			return ""
		}

		// Resource is ready, construct the identity
		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}

		region, _, _ := unstructured.NestedString(u.Object, "spec", "region")
		if region == "" {
			return ""
		}

		resourceID, _ := refs.GetResourceID(u)
		if resourceID == "" {
			return ""
		}

		id := &VertexAIEndpointIdentity{
			Project:  projectID,
			Location: region,
			Endpoint: resourceID,
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
