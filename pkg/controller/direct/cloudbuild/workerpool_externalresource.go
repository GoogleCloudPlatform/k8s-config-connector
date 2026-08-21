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

package cloudbuild

import (
	"fmt"
	"strings"
)

const (
	serviceDomain = "//cloudbuild.googleapis.com"
)

type CloudBuildWorkerPoolIdentity struct {
	project    string
	location   string
	workerpool string
}

// Parent builds a CloudBuildWorkerPool parent of the format projects/<project>/locations/<location>
func (c *CloudBuildWorkerPoolIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", c.project, c.location)
}

// FullyQualifiedName builds a CloudBuildWorkerPool resource of the format projects/<project>/locations/<location>/workerPools/<workerPool>
func (c *CloudBuildWorkerPoolIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", c.project, c.location, c.workerpool)
}

// AsExternalRef builds a externalRef from a CloudBuildWorkerPoolIdentity
func (c *CloudBuildWorkerPoolIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a CloudBuildWorkerPoolIdentity from a externalRef
func asID(externalRef string) (*CloudBuildWorkerPoolIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "workerPools" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/workerPools/<workerPool>, got %s",
			serviceDomain, externalRef)
	}
	return &CloudBuildWorkerPoolIdentity{
		project:    tokens[1],
		location:   tokens[3],
		workerpool: tokens[5],
	}, nil
}

// BuildID builds a CloudBuildWorkerPoolIdentity from resource components.
func BuildID(project, location, workerpool string) *CloudBuildWorkerPoolIdentity {
	return &CloudBuildWorkerPoolIdentity{
		project:    project,
		location:   location,
		workerpool: workerpool,
	}
}
