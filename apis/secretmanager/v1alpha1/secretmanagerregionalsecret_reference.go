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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &SecretManagerRegionalSecretRef{}

// SecretManagerRegionalSecretRef defines the resource reference to SecretManagerRegionalSecret.
type SecretManagerRegionalSecretRef struct {
	// A reference to an externally managed SecretManagerRegionalSecret resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/secrets/{{secret}}".
	External string `json:"external,omitempty"`

	// The name of a SecretManagerRegionalSecret resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SecretManagerRegionalSecret resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&SecretManagerRegionalSecretRef{})
}

func (r *SecretManagerRegionalSecretRef) GetGVK() schema.GroupVersionKind {
	return SecretManagerRegionalSecretGVK
}

func (r *SecretManagerRegionalSecretRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SecretManagerRegionalSecretRef) GetExternal() string {
	return r.External
}

func (r *SecretManagerRegionalSecretRef) SetExternal(external string) {
	r.External = external
}

func (r *SecretManagerRegionalSecretRef) ValidateExternal(external string) error {
	id := &SecretManagerRegionalSecretIdentity{}
	return id.FromExternal(external)
}

func (r *SecretManagerRegionalSecretRef) ParseExternalToIdentity() (interface{}, error) {
	id := &SecretManagerRegionalSecretIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SecretManagerRegionalSecretRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromSecretManagerRegionalSecretSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
