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

package dataform

import (
	"fmt"
	"strings"
)

type ReleaseConfigIdentity struct {
	project       string
	location      string
	repository    string
	releaseConfig string
}

func (c *ReleaseConfigIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", c.project, c.location, c.repository)
}

func (c *ReleaseConfigIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/releaseConfigs/%s", c.project, c.location, c.repository, c.releaseConfig)
}

func (c *ReleaseConfigIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

func asReleaseConfigID(externalRef string) (*ReleaseConfigIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "repositories" || tokens[6] != "releaseConfigs" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/repositories/<repository>/releaseConfigs/<release_config>, got %s",
			serviceDomain, externalRef)
	}
	return &ReleaseConfigIdentity{
		project:       tokens[1],
		location:      tokens[3],
		repository:    tokens[5],
		releaseConfig: tokens[7],
	}, nil
}

func BuildReleaseConfigID(project, location, repository, releaseConfig string) *ReleaseConfigIdentity {
	return &ReleaseConfigIdentity{
		project:       project,
		location:      location,
		repository:    repository,
		releaseConfig: releaseConfig,
	}
}
