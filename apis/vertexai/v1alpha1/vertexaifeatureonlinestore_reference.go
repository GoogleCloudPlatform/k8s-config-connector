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

var VertexAIFeatureOnlineStoreGVK = schema.GroupVersionKind{
	Group:   "aiplatform.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "VertexAIFeatureOnlineStore",
}

var _ refs.Ref = &VertexAIFeatureOnlineStoreRef{}

// VertexAIFeatureOnlineStoreRef is a reference to a VertexAIFeatureOnlineStore.
type VertexAIFeatureOnlineStoreRef struct {
	// A reference to an externally managed VertexAIFeatureOnlineStore resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/featureOnlineStores/{{featureOnlineStoreID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIFeatureOnlineStore resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIFeatureOnlineStore resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAIFeatureOnlineStoreRef{}, nil)
}

func (r *VertexAIFeatureOnlineStoreRef) GetGVK() schema.GroupVersionKind {
	return VertexAIFeatureOnlineStoreGVK
}

func (r *VertexAIFeatureOnlineStoreRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIFeatureOnlineStoreRef) GetExternal() string {
	return r.External
}

func (r *VertexAIFeatureOnlineStoreRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

var VertexAIFeatureOnlineStoreIdentityFormat = gcpurls.Template[VertexAIFeatureOnlineStoreIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/featureOnlineStores/{featureOnlineStore}")

type VertexAIFeatureOnlineStoreIdentity struct {
	Project            string
	Location           string
	FeatureOnlineStore string
}

func (i *VertexAIFeatureOnlineStoreIdentity) String() string {
	return VertexAIFeatureOnlineStoreIdentityFormat.ToString(*i)
}

func (i *VertexAIFeatureOnlineStoreIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIFeatureOnlineStoreIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIFeatureOnlineStore external=%q was not known (use %s): %w", ref, VertexAIFeatureOnlineStoreIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIFeatureOnlineStore external=%q was not known (use %s)", ref, VertexAIFeatureOnlineStoreIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (r *VertexAIFeatureOnlineStoreRef) ValidateExternal(ref string) error {
	id := &VertexAIFeatureOnlineStoreIdentity{}
	return id.FromExternal(ref)
}

func (r *VertexAIFeatureOnlineStoreRef) ParseExternalToIdentity() (identity.Identity, error) {
	if r.External == "" {
		return nil, fmt.Errorf("external is empty")
	}
	id := &VertexAIFeatureOnlineStoreIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIFeatureOnlineStoreRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		name, _, _ := unstructured.NestedString(u.Object, "status", "name")
		return name
	}
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, fallback)
}
