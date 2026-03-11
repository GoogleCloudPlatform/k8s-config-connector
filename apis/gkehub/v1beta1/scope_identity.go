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

package v1beta1

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

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
