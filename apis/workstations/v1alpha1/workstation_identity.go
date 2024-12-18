// Copyright 2024 Google LLC
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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// WorkstationIdentity defines the resource reference to Workstation.
type WorkstationIdentity struct {
	parent *WorkstationParent
	id     string
}

func (i *WorkstationIdentity) String() string {
	return i.parent.String() + "/workstations/" + i.id
}

func (i *WorkstationIdentity) ID() string {
	return i.id
}

func (i *WorkstationIdentity) Parent() *WorkstationParent {
	return i.parent
}

type WorkstationParent struct {
	ProjectID string
	Location  string
	Cluster   string
	Config    string
}

func (p *WorkstationParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/workstationClusters/" + p.Cluster + "/workstationConfigs/" + p.Config
}

// New builds a WorkstationIdentity from the Config Connector Workstation object.
func NewWorkstationIdentity(ctx context.Context, reader client.Reader, obj *Workstation) (*WorkstationIdentity, error) {
	// Get Parent
	configRef := obj.Spec.Parent
	if configRef == nil {
		return nil, fmt.Errorf("no parent config")
	}
	configExternal, err := configRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve config: %w", err)
	}
	configParent, config, err := ParseWorkstationConfigExternal(configExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse external config: %w", err)
	}
	projectID := configParent.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := configParent.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	cluster := configParent.Cluster
	if cluster == "" {
		return nil, fmt.Errorf("cannot resolve cluster")
	}

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
		actualParent, actualResourceID, err := ParseWorkstationExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualParent.Cluster != cluster {
			return nil, fmt.Errorf("spec.cluster changed, expect %s, got %s", actualParent.Cluster, cluster)
		}
		if actualParent.Config != config {
			return nil, fmt.Errorf("spec.config changed, expect %s, got %s", actualParent.Config, config)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &WorkstationIdentity{
		parent: &WorkstationParent{
			ProjectID: projectID,
			Location:  location,
			Cluster:   cluster,
			Config:    config,
		},
		id: resourceID,
	}, nil
}

func ParseWorkstationExternal(external string) (parent *WorkstationParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 10 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "workstationClusters" || tokens[6] != "workstationConfigs" || tokens[8] != "workstations" {
		return nil, "", fmt.Errorf("format of Workstation external=%q was not known (use projects/{{projectID}}/locations/{{location}}/workstationClusters/{{workstationclusterID}}/workstationConfigs/{{workstationconfigID}}/workstations/{{workstationID}})", external)
	}
	parent = &WorkstationParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		Cluster:   tokens[5],
		Config:    tokens[7],
	}
	resourceID = tokens[9]
	return parent, resourceID, nil
}
