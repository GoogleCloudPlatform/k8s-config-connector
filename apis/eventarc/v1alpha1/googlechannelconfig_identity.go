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

// GoogleChannelConfigIdentity defines the resource reference to EventarcGoogleChannelConfig, which "External" field
// holds the GCP identifier for the KRM object.
type GoogleChannelConfigIdentity struct {
	parent *GoogleChannelConfigParent
	id     string
}

func (i *GoogleChannelConfigIdentity) String() string {
	// Use the singular name 'googleChannelConfig' from the proto definition
	// The resource ID *is* expected by the API.
	return i.parent.String() + "/" + i.id
}

func (i *GoogleChannelConfigIdentity) ID() string {
	return i.id
}

func (i *GoogleChannelConfigIdentity) Parent() *GoogleChannelConfigParent {
	return i.parent
}

type GoogleChannelConfigParent struct {
	ProjectID string
	Location  string
}

func (p *GoogleChannelConfigParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a GoogleChannelConfigIdentity from the Config Connector GoogleChannelConfig object.
func NewGoogleChannelConfigIdentity(ctx context.Context, reader client.Reader, obj *EventarcGoogleChannelConfig) (*GoogleChannelConfigIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	if resourceID != "googleChannelConfig" {
		return nil, fmt.Errorf("resourceID or .metadata.name can only be 'googleChannelConfig'")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseGoogleChannelConfigExternal(externalRef)
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
	return &GoogleChannelConfigIdentity{
		parent: &GoogleChannelConfigParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseGoogleChannelConfigExternal(external string) (parent *GoogleChannelConfigParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	// Expected format based on proto pattern `projects/{project}/locations/{location}/googleChannelConfig` and KCC convention of adding resource ID:
	// projects/{projectID}/locations/{location}/googleChannelConfig/{googlechannelconfigID}
	if len(tokens) != 5 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "googleChannelConfig" {
		return nil, "", fmt.Errorf("format of EventarcGoogleChannelConfig external=%q was not known (use projects/{projectID}/locations/{location}/googleChannelConfig/{googlechannelconfigID})", external)
	}
	parent = &GoogleChannelConfigParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[4]
	return parent, resourceID, nil
}
