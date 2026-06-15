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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var VertexAITensorboardGVK = schema.GroupVersionKind{
	Group:   "vertexai.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "VertexAITensorboard",
}

var _ refs.Ref = &VertexAITensorboardRef{}

// VertexAITensorboardRef is a reference to a VertexAITensorboard.
type VertexAITensorboardRef struct {
	// A reference to an externally managed VertexAITensorboard resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/tensorboards/{{tensorboardID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAITensorboard resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAITensorboard resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAITensorboardRef{}, nil)
}

func (r *VertexAITensorboardRef) GetGVK() schema.GroupVersionKind {
	return VertexAITensorboardGVK
}

func (r *VertexAITensorboardRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAITensorboardRef) GetExternal() string {
	return r.External
}

func (r *VertexAITensorboardRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

var VertexAITensorboardIdentityFormat = gcpurls.Template[VertexAITensorboardIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/tensorboards/{tensorboard}")

type VertexAITensorboardIdentity struct {
	Project     string
	Location    string
	Tensorboard string
}

func (i *VertexAITensorboardIdentity) String() string {
	return VertexAITensorboardIdentityFormat.ToString(*i)
}

func (i *VertexAITensorboardIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAITensorboardIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAITensorboard external=%q was not known (use %s): %w", ref, VertexAITensorboardIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAITensorboard external=%q was not known (use %s)", ref, VertexAITensorboardIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *VertexAITensorboardRef) ValidateExternal(ref string) error {
	id := &VertexAITensorboardIdentity{}
	return id.FromExternal(ref)
}

func (r *VertexAITensorboardRef) ParseExternalToIdentity() (identity.Identity, error) {
	if r.External == "" {
		return nil, fmt.Errorf("external is empty")
	}
	id := &VertexAITensorboardIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAITensorboardRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		name, _, _ := unstructured.NestedString(u.Object, "status", "name")
		return name
	}
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, fallback)
}
