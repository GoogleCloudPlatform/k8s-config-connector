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

var _ refs.Ref = &AIPlatformModelRef{}

// AIPlatformModelRef defines the resource reference to AIPlatformModel, which "External" field
// holds the GCP identifier for the KRM object.
type AIPlatformModelRef struct {
	// The name of a AIPlatformModel resource.
	// Should be in the format "projects/{project}/locations/{location}/models/{model}".
	External string `json:"external,omitempty"`

	// The name of a AIPlatformModel resource.
	Name string `json:"name,omitempty"`

	// The namespace of a AIPlatformModel resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&AIPlatformModelRef{})
}

func (r *AIPlatformModelRef) GetGVK() schema.GroupVersionKind {
	return AIPlatformModelGVK
}

func (r *AIPlatformModelRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *AIPlatformModelRef) GetExternal() string {
	return r.External
}

func (r *AIPlatformModelRef) SetExternal(ref string) {
	r.External = ref
}

func (r *AIPlatformModelRef) ValidateExternal(ref string) error {
	id := &AIPlatformModelIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *AIPlatformModelRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &AIPlatformModelIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *AIPlatformModelRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromAIPlatformModelSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
