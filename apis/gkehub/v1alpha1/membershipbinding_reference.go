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

var _ refs.Ref = &GKEHubMembershipBindingRef{}

// GKEHubMembershipBindingRef is a reference to a GKEHubMembershipBinding.
type GKEHubMembershipBindingRef struct {
	// A reference to an externally managed GKEHubMembershipBinding resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/memberships/{{membershipID}}/bindings/{{membershipBindingID}}".
	External string `json:"external,omitempty"`

	// The name of a GKEHubMembershipBinding resource.
	Name string `json:"name,omitempty"`

	// The namespace of a GKEHubMembershipBinding resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&GKEHubMembershipBindingRef{})
}

func (r *GKEHubMembershipBindingRef) GetGVK() schema.GroupVersionKind {
	return GKEHubMembershipBindingGVK
}

func (r *GKEHubMembershipBindingRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *GKEHubMembershipBindingRef) GetExternal() string {
	return r.External
}

func (r *GKEHubMembershipBindingRef) SetExternal(ref string) {
	r.External = ref
}

func (r *GKEHubMembershipBindingRef) ValidateExternal(ref string) error {
	id := &GKEHubMembershipBindingIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *GKEHubMembershipBindingRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &GKEHubMembershipBindingIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *GKEHubMembershipBindingRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromGKEHubMembershipBindingSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func ResolveGKEHubMembershipBindingRef(ctx context.Context, reader client.Reader, obj client.Object, ref *GKEHubMembershipBindingRef) (*GKEHubMembershipBindingIdentity, error) {
	if ref == nil {
		return nil, nil
	}
	if err := ref.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	id := &GKEHubMembershipBindingIdentity{}
	if err := id.FromExternal(ref.External); err != nil {
		return nil, err
	}
	return id, nil
}
