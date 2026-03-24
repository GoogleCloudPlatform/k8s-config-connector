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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"

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

var _ identity.Identity = &GKEHubNamespaceIdentity{}
var _ identity.Resource = &GKEHubNamespace{}

func (c *GKEHubNamespace) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	projectID := ""
	if c.Spec.ProjectRef != nil {
		project, err := refs.ResolveProject(ctx, reader, c.GetNamespace(), c.Spec.ProjectRef)
		if err != nil {
			return nil, err
		}
		if project != nil {
			projectID = project.ProjectID
		}
	}
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project for GKEHubNamespace")
	}

	location := "global"
	if c.Spec.Location != nil {
		location = *c.Spec.Location
	}

	scopeID := ""
	if c.Spec.ScopeRef != nil {
		if c.Spec.ScopeRef.External != nil {
			scopeIdentity := &GKEHubScopeIdentity{}
			if err := scopeIdentity.FromExternal(*c.Spec.ScopeRef.External); err != nil {
				return nil, err
			}
			scopeID = scopeIdentity.ID()
		} else {
			var err error
			scopeID, err = resolveScopeID(ctx, reader, c.GetNamespace(), c.Spec.ScopeRef)
			if err != nil {
				return nil, err
			}
		}
	}
	if scopeID == "" {
		return nil, fmt.Errorf("cannot resolve scope for GKEHubNamespace")
	}

	resourceID := ""
	if c.Spec.ResourceID != nil {
		resourceID = *c.Spec.ResourceID
	}
	if resourceID == "" {
		resourceID = c.GetName()
	}

	id := NewGKEHubNamespaceIdentity(projectID, location, scopeID, resourceID)

	if c.Status.ExternalRef != nil && *c.Status.ExternalRef != "" {
		statusID := &GKEHubNamespaceIdentity{}
		if err := statusID.FromExternal(*c.Status.ExternalRef); err != nil {
			return nil, err
		}
		if statusID.String() != id.String() {
			return nil, fmt.Errorf("existing externalRef %q does not match identity resolved from spec %q", *c.Status.ExternalRef, id.String())
		}
	}

	return id, nil
}
