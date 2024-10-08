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

package workstations

import (
	"fmt"
	"strings"
)

// The Identifier for ConfigConnector to track the WorkstationCluster resource from the GCP service.
type WorkstationClusterIdentity struct {
	Parent             parent
	WorkstationCluster string
}

type parent struct {
	Project  string
	Location string
}

func (p *parent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", p.Project, p.Location)
}

// FullyQualifiedName returns both parent and resource ID in the full url format.
func (c *WorkstationClusterIdentity) FullyQualifiedName() string {
	// TODO(user): Edit the URL path
	return fmt.Sprintf("%s/workstationClusters/%s", c.Parent.String(), c.WorkstationCluster)
}

// AsExternalRef builds a externalRef from a WorkstationCluster
func (c *WorkstationClusterIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a WorkstationClusterIdentity from a `status.externalRef`
func asID(externalRef string) (*WorkstationClusterIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "workstationClusters" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/workstationClusters/<WorkstationCluster>, got %s",
			serviceDomain, externalRef)
	}
	return &WorkstationClusterIdentity{
		Parent:             parent{Project: tokens[1], Location: tokens[3]},
		WorkstationCluster: tokens[5],
	}, nil
}

// BuildID builds the ID for ConfigConnector to track the WorkstationCluster resource from the GCP service.
func BuildID(project, location, resourceID string) *WorkstationClusterIdentity {
	return &WorkstationClusterIdentity{
		Parent:             parent{Project: project, Location: location},
		WorkstationCluster: resourceID,
	}
}
