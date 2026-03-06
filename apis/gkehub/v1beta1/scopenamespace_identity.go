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

var scopeNamespaceURL = gcpurls.Template[GKEHubNamespaceIdentity](
	"gkehub.googleapis.com",
	"projects/{projectID}/locations/{location}/scopes/{scopeID}/namespaces/{scopeNamespaceID}",
)

// GKEHubNamespaceIdentity defines the resource reference to GKEHubNamespace, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type GKEHubNamespaceIdentity struct {
	ProjectID        string
	Location         string
	ScopeID          string
	ScopeNamespaceID string
}

func (i *GKEHubNamespaceIdentity) String() string {
	return scopeNamespaceURL.ToString(*i)
}

func (i *GKEHubNamespaceIdentity) ID() string {
	return i.ScopeNamespaceID
}

func (i *GKEHubNamespaceIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s/scopes/%s", i.ProjectID, i.Location, i.ScopeID)
}

func (i *GKEHubNamespaceIdentity) FromExternal(external string) error {
	out, match, err := scopeNamespaceURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of GKEHubNamespace external=%q was not known (use %s)", external, scopeNamespaceURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func NewGKEHubNamespaceIdentity(project, location, scopeID, scopeNamespaceID string) *GKEHubNamespaceIdentity {
	return &GKEHubNamespaceIdentity{
		ProjectID:        project,
		Location:         location,
		ScopeID:          scopeID,
		ScopeNamespaceID: scopeNamespaceID,
	}
}

func (i *GKEHubNamespaceIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}

func (i *GKEHubNamespaceIdentity) DefaultLocationState(location string) {
	if i.Location == "" {
		i.Location = location
	}
}
