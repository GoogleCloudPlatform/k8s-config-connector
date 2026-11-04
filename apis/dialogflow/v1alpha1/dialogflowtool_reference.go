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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var VertexAIExtensionGVK = schema.GroupVersionKind{
	Group:   "vertexai.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "VertexAIExtension",
}

var (
	_ refsv1beta1.Ref         = &VertexAIExtensionRef{}
	_ refsv1beta1.ExternalRef = &VertexAIExtensionRef{}
)

// VertexAIExtensionRef is a reference to a VertexAIExtension.
type VertexAIExtensionRef struct {
	// A reference to an externally managed VertexAIExtension resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/extensions/{{extensionID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIExtension resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIExtension resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&VertexAIExtensionRef{}, nil)
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

func (r *VertexAIExtensionRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

var VertexAIExtensionIdentityFormat = gcpurls.Template[VertexAIExtensionIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/extensions/{extension}")

type VertexAIExtensionIdentity struct {
	Project   string
	Location  string
	Extension string
}

func (i *VertexAIExtensionIdentity) String() string {
	return VertexAIExtensionIdentityFormat.ToString(*i)
}

func (i *VertexAIExtensionIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIExtensionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIExtension external=%q was not known (use %s): %w", ref, VertexAIExtensionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIExtension external=%q was not known (use %s)", ref, VertexAIExtensionIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (r *VertexAIExtensionRef) ValidateExternal(ref string) error {
	id := &VertexAIExtensionIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
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
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, nil)
}
