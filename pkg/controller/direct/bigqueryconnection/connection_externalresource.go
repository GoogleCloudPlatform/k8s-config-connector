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

package bigqueryconnection

import (
	"fmt"
	"strings"
)

// The Identifier for ConfigConnector to track the BigQueryConnectionConnection resource from the GCP service.
type BigQueryConnectionConnectionIdentity struct {
	Parent             *parent
	serviceGeneratedID string
}

type parent struct {
	Project  string
	Location string
}

func (p *parent) String() string {
	return "projects/" + p.Project + "/locations/" + p.Location
}

// FullyQualifiedName returns both parent and resource ID in the full url format.
func (c *BigQueryConnectionConnectionIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("%s/connections/%s", c.Parent, c.serviceGeneratedID)
}

// AsExternalRef builds a externalRef from a BigQueryConnectionConnection
func (c *BigQueryConnectionConnectionIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.Parent.String() + "/connections/" + c.serviceGeneratedID
	return &e
}

// asID builds a BigQueryConnectionConnectionIdentity from a `status.externalRef`
func asID(externalRef string) (*BigQueryConnectionConnectionIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/connections/<Connection>, got %s",
			serviceDomain, externalRef)
	}
	return &BigQueryConnectionConnectionIdentity{
		Parent:             &parent{Project: tokens[1], Location: tokens[3]},
		serviceGeneratedID: tokens[5],
	}, nil
}

// BuildID builds the ID for ConfigConnector to track the BigQueryConnectionConnection resource from the GCP service.
func BuildIDWithServiceGeneratedID(project, location, serviceGeneratedID string) *BigQueryConnectionConnectionIdentity {
	return &BigQueryConnectionConnectionIdentity{
		Parent:             &parent{Project: project, Location: location},
		serviceGeneratedID: serviceGeneratedID,
	}
}

func ParseNameFromGCP(fullyQualifiedName string) string {
	tokens := strings.Split(fullyQualifiedName, "/")
	return tokens[5]
}
