/*
Copyright 2026.

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

package cloudbuild

import (
	"fmt"
	"strings"
)

type CloudBuildTriggerIdentity struct {
	project  string
	location string
	trigger  string
}

// Parent builds a CloudBuildTrigger parent of the format projects/<project>/locations/<location>
func (c *CloudBuildTriggerIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", c.project, c.location)
}

// FullyQualifiedName builds a CloudBuildTrigger resource of the format projects/<project>/locations/<location>/triggers/<trigger>
func (c *CloudBuildTriggerIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/triggers/%s", c.project, c.location, c.trigger)
}

// AsExternalRef builds a externalRef from a CloudBuildTriggerIdentity
func (c *CloudBuildTriggerIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asTriggerID builds a CloudBuildTriggerIdentity from a externalRef
func asTriggerID(externalRef string) (*CloudBuildTriggerIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "triggers" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/triggers/<trigger>, got %s",
			serviceDomain, externalRef)
	}
	return &CloudBuildTriggerIdentity{
		project:  tokens[1],
		location: tokens[3],
		trigger:  tokens[5],
	}, nil
}

// BuildTriggerID builds a CloudBuildTriggerIdentity from resource components.
func BuildTriggerID(project, location, trigger string) *CloudBuildTriggerIdentity {
	return &CloudBuildTriggerIdentity{
		project:  project,
		location: location,
		trigger:  trigger,
	}
}
