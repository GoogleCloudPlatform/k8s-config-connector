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

var _ identity.Identity = &GKEHubScopeIdentity{}
var _ identity.Resource = &GKEHubScope{}

var scopeURL = gcpurls.Template[GKEHubScopeIdentity](
	"gkehub.googleapis.com",
	"projects/{projectID}/locations/{location}/scopes/{scopeID}",
)

// GKEHubScopeIdentity defines the resource reference to GKEHubScope, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type GKEHubScopeIdentity struct {
	ProjectID string
	Location  string
	ScopeID   string
}

func (i *GKEHubScopeIdentity) String() string {
	return scopeURL.ToString(*i)
}

func (i *GKEHubScopeIdentity) ID() string {
	return i.ScopeID
}

func (i *GKEHubScopeIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.ProjectID, i.Location)
}

func (i *GKEHubScopeIdentity) FromExternal(external string) error {
	out, match, err := scopeURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of GKEHubScope external=%q was not known (use %s)", external, scopeURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func NewGKEHubScopeIdentity(project, location, scopeID string) *GKEHubScopeIdentity {
	return &GKEHubScopeIdentity{
		ProjectID: project,
		Location:  location,
		ScopeID:   scopeID,
	}
}

func (i *GKEHubScopeIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}

func (i *GKEHubScopeIdentity) DefaultLocationState(location string) {
	if i.Location == "" {
		i.Location = location
	}
}

func (obj *GKEHubScope) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := project.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	location := direct.ValueOf(obj.Spec.Location)
	if location == "" {
		location = "global"
	}

	return NewGKEHubScopeIdentity(projectID, location, resourceID), nil
}
