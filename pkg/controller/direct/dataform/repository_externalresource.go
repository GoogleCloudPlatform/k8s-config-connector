/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dataform

import (
	"fmt"
	"strings"
)

const (
	serviceDomain = "//dataform.googleapis.com"
)

type DataformRepositoryIdentity struct {
	project  string
	location string
	dataform string
}

// Parent builds a DataformRepository parent of the format projects/<project>/locations/<location>
func (c *DataformRepositoryIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", c.project, c.location)
}

// FullyQualifiedName builds a DataformRepository resource of the format projects/<project>/locations/<location>/repositories/<dataformRepository>
func (c *DataformRepositoryIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", c.project, c.location, c.dataform)
}

// AsExternalRef builds a externalRef from a DataformRepositoryIdentity
func (c *DataformRepositoryIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a DataformRepositoryIdentity from a externalRef
func asID(externalRef string) (*DataformRepositoryIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "repositories" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/repositories/<dataform_repository>, got %s",
			serviceDomain, externalRef)
	}
	return &DataformRepositoryIdentity{
		project:  tokens[1],
		location: tokens[3],
		dataform: tokens[5],
	}, nil
}

// BuildID builds a DataformRepositoryIdentity from resource components.
func BuildID(project, location, dataform string) *DataformRepositoryIdentity {
	return &DataformRepositoryIdentity{
		project:  project,
		location: location,
		dataform: dataform,
	}
}
