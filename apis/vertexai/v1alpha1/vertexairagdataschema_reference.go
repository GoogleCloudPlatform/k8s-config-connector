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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ v1beta1.ExternalRef = &VertexAIRagDataSchemaRef{}

var VertexAIRagDataSchemaRefGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "VertexAIRagDataSchema",
}

func init() {
	v1beta1.Register(&VertexAIRagDataSchemaRef{})
	v1beta1.Register(&VertexAIRagCorpusRef{})
}

// VertexAIRagDataSchemaRef is a reference to a GCP VertexAIRagDataSchema.
type VertexAIRagDataSchemaRef struct {
	// A reference to an externally managed VertexAIRagDataSchema resource. Should be in the format "projects/{{project}}/locations/{{location}}/ragCorpora/{{ragCorpus}}/ragDataSchemas/{{ragDataSchema}}"
	External string `json:"external,omitempty"`

	// The name of a VertexAIRagDataSchema resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIRagDataSchema resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *VertexAIRagDataSchemaRef) GetGVK() schema.GroupVersionKind {
	return VertexAIRagDataSchemaRefGVK
}

func (r *VertexAIRagDataSchemaRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{Namespace: r.Namespace, Name: r.Name}
}

func (r *VertexAIRagDataSchemaRef) GetExternal() string {
	return r.External
}

func (r *VertexAIRagDataSchemaRef) SetExternal(external string) {
	r.External = external
}

func (r *VertexAIRagDataSchemaRef) ValidateExternal(external string) error {
	return (&VertexAIRagDataSchemaIdentity{}).FromExternal(external)
}

func (r *VertexAIRagDataSchemaRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &VertexAIRagDataSchemaIdentity{}
	if err := identity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *VertexAIRagDataSchemaRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return v1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromVertexAIRagDataSchemaSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	})
}

// VertexAIRagCorpusRef

var _ v1beta1.ExternalRef = &VertexAIRagCorpusRef{}

var VertexAIRagCorpusGVK = schema.GroupVersionKind{
	Group:   "vertexai.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "VertexAIRagCorpus",
}

// VertexAIRagCorpusRef is a reference to a VertexAIRagCorpus.
type VertexAIRagCorpusRef struct {
	// A reference to an externally managed VertexAIRagCorpus resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/ragCorpora/{{ragCorpusID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIRagCorpus resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIRagCorpus resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *VertexAIRagCorpusRef) GetGVK() schema.GroupVersionKind {
	return VertexAIRagCorpusGVK
}

func (r *VertexAIRagCorpusRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{Namespace: r.Namespace, Name: r.Name}
}

func (r *VertexAIRagCorpusRef) GetExternal() string {
	return r.External
}

func (r *VertexAIRagCorpusRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

var VertexAIRagCorpusIdentityFormat = gcpurls.Template[VertexAIRagCorpusIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/ragCorpora/{ragCorpus}")

type VertexAIRagCorpusIdentity struct {
	Project   string
	Location  string
	RagCorpus string
}

func (i *VertexAIRagCorpusIdentity) String() string {
	return VertexAIRagCorpusIdentityFormat.ToString(*i)
}

func (i *VertexAIRagCorpusIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIRagCorpusIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIRagCorpus external=%q was not known (use %s): %w", ref, VertexAIRagCorpusIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIRagCorpus external=%q was not known (use %s)", ref, VertexAIRagCorpusIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *VertexAIRagCorpusRef) ValidateExternal(ref string) error {
	id := &VertexAIRagCorpusIdentity{}
	return id.FromExternal(ref)
}

func (r *VertexAIRagCorpusRef) ParseExternalToIdentity() (identity.Identity, error) {
	if r.External == "" {
		return nil, fmt.Errorf("external is empty")
	}
	id := &VertexAIRagCorpusIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIRagCorpusRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		name, _, _ := unstructured.NestedString(u.Object, "status", "name")
		return name
	}
	return v1beta1.NormalizeWithFallback(ctx, reader, r, otherNamespace, fallback)
}
