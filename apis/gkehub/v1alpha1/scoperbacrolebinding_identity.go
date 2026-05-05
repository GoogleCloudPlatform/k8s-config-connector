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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.Identity   = &GKEHubScopeRBACRoleBindingIdentity{}
	_ identity.IdentityV2 = &GKEHubScopeRBACRoleBindingIdentity{}
	_ identity.Resource   = &GKEHubScopeRBACRoleBinding{}

	rbacRoleBindingURL = gcpurls.Template[GKEHubScopeRBACRoleBindingIdentity](
		"gkehub.googleapis.com",
		"projects/{projectID}/locations/{location}/scopes/{scopeID}/rbacrolebindings/{rbacRoleBindingID}",
	)
)

// GKEHubScopeRBACRoleBindingIdentity defines the resource reference to GKEHubScopeRBACRoleBinding, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type GKEHubScopeRBACRoleBindingIdentity struct {
	ProjectID         string
	Location          string
	ScopeID           string
	RBACRoleBindingID string
}

func (i *GKEHubScopeRBACRoleBindingIdentity) String() string {
	return rbacRoleBindingURL.ToString(*i)
}

func (i *GKEHubScopeRBACRoleBindingIdentity) ID() string {
	return i.RBACRoleBindingID
}

func (i *GKEHubScopeRBACRoleBindingIdentity) Host() string {
	return rbacRoleBindingURL.Host()
}

func (i *GKEHubScopeRBACRoleBindingIdentity) Parent() *GKEHubScopeIdentity {
	return NewGKEHubScopeIdentity(i.ProjectID, i.Location, i.ScopeID)
}

func (i *GKEHubScopeRBACRoleBindingIdentity) FromExternal(external string) error {
	out, match, err := rbacRoleBindingURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of GKEHubScopeRBACRoleBinding external=%q was not known (use %s)", external, rbacRoleBindingURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func NewGKEHubScopeRBACRoleBindingIdentity(project, location, scope, rbacRoleBinding string) *GKEHubScopeRBACRoleBindingIdentity {
	return &GKEHubScopeRBACRoleBindingIdentity{
		ProjectID:         project,
		Location:          location,
		ScopeID:           scope,
		RBACRoleBindingID: rbacRoleBinding,
	}
}

func (obj *GKEHubScopeRBACRoleBinding) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	if obj.Spec.ScopeRef == nil {
		return nil, fmt.Errorf("spec.scopeRef is required")
	}
	scopeRef, err := ResolveGKEHubScopeRef(ctx, reader, obj, obj.Spec.ScopeRef)
	if err != nil {
		return nil, err
	}

	resourceID := direct.ValueOf(obj.Spec.RBACRoleBindingID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	// We use a single parent (ScopeRef) instead of grandparent (ProjectRef, Location).
	// The ScopeRef contains the project and location.
	return NewGKEHubScopeRBACRoleBindingIdentity(scopeRef.ProjectID, scopeRef.Location, scopeRef.ScopeID, resourceID), nil
}
