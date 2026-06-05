// Copyright 2025 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// BackendAuthenticationConfigIdentity defines the resource reference to NetworkSecurityBackendAuthenticationConfig, which "External" field
// holds the GCP identifier for the KRM object.
type BackendAuthenticationConfigIdentity struct {
	parent *BackendAuthenticationConfigParent
	id     string
}

func (i *BackendAuthenticationConfigIdentity) String() string {
	return i.parent.String() + "/backendAuthenticationConfigs/" + i.id
}

func (i *BackendAuthenticationConfigIdentity) ID() string {
	return i.id
}

func (i *BackendAuthenticationConfigIdentity) Parent() *BackendAuthenticationConfigParent {
	return i.parent
}

type BackendAuthenticationConfigParent struct {
	ProjectID string
	Location  string
}

func (p *BackendAuthenticationConfigParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a BackendAuthenticationConfigIdentity from the Config Connector BackendAuthenticationConfig object.
func NewBackendAuthenticationConfigIdentity(ctx context.Context, reader client.Reader, obj *NetworkSecurityBackendAuthenticationConfig) (*BackendAuthenticationConfigIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := common.ValueOf(obj.Spec.Location)

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseBackendAuthenticationConfigExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &BackendAuthenticationConfigIdentity{
		parent: &BackendAuthenticationConfigParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseBackendAuthenticationConfigExternal(external string) (parent *BackendAuthenticationConfigParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || strings.ToLower(tokens[4]) != "backendauthenticationconfigs" {
		return nil, "", fmt.Errorf("format of NetworkSecurityBackendAuthenticationConfig external=%q was not known (use projects/{{projectID}}/locations/{{location}}/backendAuthenticationConfigs/{{backendauthenticationconfigID}})", external)
	}
	parent = &BackendAuthenticationConfigParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
