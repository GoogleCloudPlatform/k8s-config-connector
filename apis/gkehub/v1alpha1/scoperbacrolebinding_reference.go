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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ v1beta1.Ref = &GKEHubScopeRBACRoleBindingRef{}

// GKEHubScopeRBACRoleBindingRef defines the resource reference to GKEHubScopeRBACRoleBinding.
type GKEHubScopeRBACRoleBindingRef struct {
	/* A reference to an externally managed GKEHubScopeRBACRoleBinding resource.
	   Should be in the format "projects/{{projectID}}/locations/{{location}}/scopes/{{scopeID}}/rbacrolebindings/{{rbacRoleBindingID}}". */
	External string `json:"external,omitempty"`

	/* The name of a GKEHubScopeRBACRoleBinding resource. */
	Name string `json:"name,omitempty"`

	/* The namespace of a GKEHubScopeRBACRoleBinding resource. */
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	v1beta1.Register(&GKEHubScopeRBACRoleBindingRef{})
}

// GetGVK returns the GroupVersionKind.
func (r *GKEHubScopeRBACRoleBindingRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   GKEHubScopeRBACRoleBindingGVK.Group,
		Version: GKEHubScopeRBACRoleBindingGVK.Version,
		Kind:    GKEHubScopeRBACRoleBindingGVK.Kind,
	}
}

// GetNamespacedName returns the namespaced name.
func (r *GKEHubScopeRBACRoleBindingRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// GetExternal returns the external string.
func (r *GKEHubScopeRBACRoleBindingRef) GetExternal() string {
	return r.External
}

// SetExternal sets the external string.
func (r *GKEHubScopeRBACRoleBindingRef) SetExternal(external string) {
	r.External = external
}

// ValidateExternal validates the external string.
func (r *GKEHubScopeRBACRoleBindingRef) ValidateExternal(external string) error {
	id := &GKEHubScopeRBACRoleBindingIdentity{}
	return id.FromExternal(external)
}

// ParseExternalToIdentity parses the external string to an identity.
func (r *GKEHubScopeRBACRoleBindingRef) ParseExternalToIdentity(external string) (identity.Identity, error) {
	id := &GKEHubScopeRBACRoleBindingIdentity{}
	err := id.FromExternal(external)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// Normalize parses the reference to a canonical string.
func (r *GKEHubScopeRBACRoleBindingRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return v1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		obj := &GKEHubScopeRBACRoleBinding{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
			return ""
		}
		id, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return ""
		}
		return id.String()
	})
}

// ResolveGKEHubScopeRBACRoleBindingRef resolves the GKEHubScopeRBACRoleBinding reference.
func ResolveGKEHubScopeRBACRoleBindingRef(ctx context.Context, reader client.Reader, obj client.Object, ref *GKEHubScopeRBACRoleBindingRef) (*GKEHubScopeRBACRoleBindingIdentity, error) {
	if ref == nil {
		return nil, nil
	}
	if err := ref.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	id := &GKEHubScopeRBACRoleBindingIdentity{}
	if err := id.FromExternal(ref.External); err != nil {
		return nil, err
	}
	return id, nil
}
