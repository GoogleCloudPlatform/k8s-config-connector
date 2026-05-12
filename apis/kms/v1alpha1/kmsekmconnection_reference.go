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

var _ refs.Ref = &KMSEKMConnectionRef{}

// KMSEKMConnectionRef represents a reference to a KMSEKMConnection.
type KMSEKMConnectionRef struct {
	// A reference to an externally managed KMSEKMConnection resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/ekmConnections/{{ekmConnection}}"
	External string `json:"external,omitempty"`

	// The name of a KMSEKMConnection resource.
	Name string `json:"name,omitempty"`

	// The namespace of a KMSEKMConnection resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&KMSEKMConnectionRef{})
}

func (r *KMSEKMConnectionRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   KMSEKMConnectionGVK.Group,
		Version: KMSEKMConnectionGVK.Version,
		Kind:    KMSEKMConnectionGVK.Kind,
	}
}

func (r *KMSEKMConnectionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *KMSEKMConnectionRef) GetExternal() string {
	return r.External
}

func (r *KMSEKMConnectionRef) SetExternal(external string) {
	r.External = external
}

func (r *KMSEKMConnectionRef) ValidateExternal(external string) error {
	id := &KMSEKMConnectionIdentity{}
	return id.FromExternal(external)
}

func (r *KMSEKMConnectionRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &KMSEKMConnectionIdentity{}
	err := id.FromExternal(r.External)
	return id, err
}

func (r *KMSEKMConnectionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromKMSEKMConnectionSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
