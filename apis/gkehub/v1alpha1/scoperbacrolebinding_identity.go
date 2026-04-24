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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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

func NewGKEHubScopeRBACRoleBindingIdentity(project, location, scopeID, rbacRoleBindingID string) *GKEHubScopeRBACRoleBindingIdentity {
	return &GKEHubScopeRBACRoleBindingIdentity{
		ProjectID:         project,
		Location:          location,
		ScopeID:           scopeID,
		RBACRoleBindingID: rbacRoleBindingID,
	}
}

func (obj *GKEHubScopeRBACRoleBinding) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	if obj.Spec.ProjectRef == nil {
		return nil, fmt.Errorf("spec.projectRef is required")
	}
	project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := project.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if obj.Spec.ScopeRef == nil {
		return nil, fmt.Errorf("spec.scopeRef is required")
	}
	scopeRef, err := ResolveGKEHubScopeRef(ctx, reader, obj, obj.Spec.ScopeRef)
	if err != nil {
		return nil, err
	}
	scopeID := scopeRef.ID()

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	location := direct.ValueOf(obj.Spec.Location)
	if location == "" {
		location = "global"
	}

	return NewGKEHubScopeRBACRoleBindingIdentity(projectID, location, scopeID, resourceID), nil
}
