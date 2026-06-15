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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ModelArmorTemplateRef{}

// ModelArmorTemplateRef represents a reference to a ModelArmorTemplate resource.
type ModelArmorTemplateRef struct {
	// A reference to an externally managed ModelArmorTemplate resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/templates/{{template}}".
	External string `json:"external,omitempty"`

	// The name of a ModelArmorTemplate resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ModelArmorTemplate resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ModelArmorTemplateRef{}, &ModelArmorTemplate{})
}

func (r *ModelArmorTemplateRef) GetGVK() schema.GroupVersionKind {
	return ModelArmorTemplateGVK
}

func (r *ModelArmorTemplateRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ModelArmorTemplateRef) GetExternal() string {
	return r.External
}

func (r *ModelArmorTemplateRef) SetExternal(external string) {
	r.External = external
}

func (r *ModelArmorTemplateRef) ValidateExternal(ref string) error {
	id := &ModelArmorTemplateIdentity{}
	return id.FromExternal(ref)
}

func (r *ModelArmorTemplateRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ModelArmorTemplateIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ModelArmorTemplateRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromModelArmorTemplateSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
