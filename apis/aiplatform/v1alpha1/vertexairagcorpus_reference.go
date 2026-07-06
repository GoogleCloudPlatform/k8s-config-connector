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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	VertexAIIndexGVK         = schema.GroupVersionKind{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIIndex"}
	VertexAIIndexEndpointGVK = schema.GroupVersionKind{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIIndexEndpoint"}
	VertexAIEndpointGVK      = schema.GroupVersionKind{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIEndpoint"}
)

// VertexAIRagCorpusRef is a reference to a GCP VertexAIRagCorpus.
type VertexAIRagCorpusRef struct {
	// A reference to an externally managed VertexAIRagCorpus resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/ragCorpora/{{ragCorpusID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIRagCorpus resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIRagCorpus resource.
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &VertexAIRagCorpusRef{}

func (r *VertexAIRagCorpusRef) GetGVK() schema.GroupVersionKind {
	return VertexAIRagCorpusGVK
}

func (r *VertexAIRagCorpusRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIRagCorpusRef) GetExternal() string {
	return r.External
}

func (r *VertexAIRagCorpusRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAIRagCorpusRef) ValidateExternal(ref string) error {
	id := &VertexAIRagCorpusIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *VertexAIRagCorpusRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAIRagCorpusIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIRagCorpusRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		structuredObj, err := common.ToStructuredType[*VertexAIRagCorpus](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromVertexAIRagCorpusSpec(ctx, reader, structuredObj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// VertexAIIndexEndpointRef is a reference to a GCP VertexAIIndexEndpoint.
type VertexAIIndexEndpointRef struct {
	// A reference to an externally managed VertexAIIndexEndpoint resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/indexEndpoints/{{indexEndpointID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIIndexEndpoint resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIIndexEndpoint resource.
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &VertexAIIndexEndpointRef{}

func (r *VertexAIIndexEndpointRef) GetGVK() schema.GroupVersionKind {
	return VertexAIIndexEndpointGVK
}

func (r *VertexAIIndexEndpointRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIIndexEndpointRef) GetExternal() string {
	return r.External
}

func (r *VertexAIIndexEndpointRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAIIndexEndpointRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("external reference format %q is not known; expected projects/<project>/locations/<location>/indexEndpoints/<indexEndpointID>", ref)
	}
	return nil
}

func (r *VertexAIIndexEndpointRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := refs.GetResourceID(u)
		if err != nil {
			return ""
		}
		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		region, _, _ := unstructured.NestedString(u.Object, "spec", "region")
		if region == "" {
			return ""
		}
		return fmt.Sprintf("projects/%s/locations/%s/indexEndpoints/%s", projectID, region, resourceID)
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// VertexAIIndexRef is a reference to a GCP VertexAIIndex.
type VertexAIIndexRef struct {
	// A reference to an externally managed VertexAIIndex resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/indexes/{{indexID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIIndex resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIIndex resource.
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &VertexAIIndexRef{}

func (r *VertexAIIndexRef) GetGVK() schema.GroupVersionKind {
	return VertexAIIndexGVK
}

func (r *VertexAIIndexRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIIndexRef) GetExternal() string {
	return r.External
}

func (r *VertexAIIndexRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAIIndexRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("external reference format %q is not known; expected projects/<project>/locations/<location>/indexes/<indexID>", ref)
	}
	return nil
}

func (r *VertexAIIndexRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := refs.GetResourceID(u)
		if err != nil {
			return ""
		}
		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		region, _, _ := unstructured.NestedString(u.Object, "spec", "region")
		if region == "" {
			return ""
		}
		return fmt.Sprintf("projects/%s/locations/%s/indexes/%s", projectID, region, resourceID)
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// VertexAIEndpointRef is a reference to a GCP VertexAIEndpoint.
type VertexAIEndpointRef struct {
	// A reference to an externally managed VertexAIEndpoint resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/endpoints/{{endpointID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIEndpoint resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIEndpoint resource.
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &VertexAIEndpointRef{}

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
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("external reference format %q is not known; expected projects/<project>/locations/<location>/endpoints/<endpointID>", ref)
	}
	return nil
}

func (r *VertexAIEndpointRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := refs.GetResourceID(u)
		if err != nil {
			return ""
		}
		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		region, _, _ := unstructured.NestedString(u.Object, "spec", "region")
		if region == "" {
			return ""
		}
		return fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", projectID, region, resourceID)
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func init() {
	refs.Register(&VertexAIRagCorpusRef{}, &VertexAIRagCorpus{})
}
