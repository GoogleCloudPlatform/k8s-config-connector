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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &VertexAIExtensionRef{}

var VertexAIExtensionGVK = schema.GroupVersionKind{
	Group:   "vertexai.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "VertexAIExtension",
}

// VertexAIExtensionRef is a reference to a GCP VertexAIExtension.
type VertexAIExtensionRef struct {
	// A reference to an externally managed VertexAIExtension resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/extensions/{{extensionID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIExtension resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIExtension resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAIExtensionRef{})
}

func (r *VertexAIExtensionRef) GetGVK() schema.GroupVersionKind {
	return VertexAIExtensionGVK
}

func (r *VertexAIExtensionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIExtensionRef) GetExternal() string {
	return r.External
}

func (r *VertexAIExtensionRef) SetExternal(ref string) {
	r.External = ref
}

func (r *VertexAIExtensionRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("VertexAIExtension external %q must start with 'projects/'", ref)
	}
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[2] != "locations" || parts[4] != "extensions" {
		return fmt.Errorf("VertexAIExtension external %q must be in format projects/{project}/locations/{location}/extensions/{extension}", ref)
	}
	return nil
}

func (r *VertexAIExtensionRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAIExtensionIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIExtensionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		name, _, _ := unstructured.NestedString(u.Object, "status", "name")
		return name
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// +k8s:deepcopy-gen=false
type VertexAIExtensionIdentity struct {
	Project   string
	Location  string
	Extension string
}

var _ identity.Identity = &VertexAIExtensionIdentity{}

func (i *VertexAIExtensionIdentity) Host() string {
	return "aiplatform.googleapis.com"
}

func (i *VertexAIExtensionIdentity) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/extensions/%s", i.Project, i.Location, i.Extension)
}

func (i *VertexAIExtensionIdentity) FromExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") {
		return fmt.Errorf("VertexAIExtension external %q must start with 'projects/'", ref)
	}
	parts := strings.Split(ref, "/")
	if len(parts) != 6 || parts[2] != "locations" || parts[4] != "extensions" {
		return fmt.Errorf("VertexAIExtension external %q must be in format projects/{project}/locations/{location}/extensions/{extension}", ref)
	}
	i.Project = parts[1]
	i.Location = parts[3]
	i.Extension = parts[5]
	return nil
}
