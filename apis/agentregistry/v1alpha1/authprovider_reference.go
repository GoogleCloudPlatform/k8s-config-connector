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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var AuthProviderGVK = schema.GroupVersionKind{
	Group:   "agentregistry.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "AuthProvider",
}

var _ refs.Ref = &AuthProviderRef{}

// AuthProviderRef is a reference to a GCP AuthProvider.
type AuthProviderRef struct {
	// A reference to an externally managed AuthProvider resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/authProviders/{{authProviderID}}".
	External string `json:"external,omitempty"`

	// The name of an AuthProvider resource.
	Name string `json:"name,omitempty"`

	// The namespace of an AuthProvider resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&AuthProviderRef{})
}

func (r *AuthProviderRef) GetGVK() schema.GroupVersionKind {
	return AuthProviderGVK
}

func (r *AuthProviderRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *AuthProviderRef) GetExternal() string {
	return r.External
}

func (r *AuthProviderRef) SetExternal(external string) {
	r.External = external
}

func (r *AuthProviderRef) ValidateExternal(ref string) error {
	id := &AuthProviderIdentity{}
	return id.FromExternal(ref)
}

func (r *AuthProviderRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &AuthProviderIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *AuthProviderRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

// AuthProviderIdentity is the identity of a GCP AuthProvider.
type AuthProviderIdentity struct {
	ProjectID      string
	Location       string
	AuthProviderID string
}

func (i *AuthProviderIdentity) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/authProviders/%s", i.ProjectID, i.Location, i.AuthProviderID)
}

func (i *AuthProviderIdentity) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "authProviders" {
		return fmt.Errorf("format of AuthProvider external=%q was not known (use projects/{{projectID}}/locations/{{location}}/authProviders/{{authProviderID}})", external)
	}
	i.ProjectID = tokens[1]
	i.Location = tokens[3]
	i.AuthProviderID = tokens[5]
	return nil
}
