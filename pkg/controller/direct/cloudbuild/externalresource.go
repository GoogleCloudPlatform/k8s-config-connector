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
	serviceDomain = "//cloudbuild.googleapis.com/"
)

type CloudBuildWorkerPoolIDentity struct {
	project    string
	location   string
	workerpool string
}

// Parent builds a CloudBuildWorkerPool parent of the format projects/<project>/locations/<location>
func (c *CloudBuildWorkerPoolIDentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", c.project, c.location)
}

// FullyQualifiedName builds a CloudBuildWorkerPool resource of the format projects/<project>/locations/<location>/workerPools/<workerPool>
func (c *CloudBuildWorkerPoolIDentity) FullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", c.project, c.location, c.workerpool)
}

// fromID builds a externalRef from a CloudBuildWorkerPoolIDentity
func asExternalRef(id *CloudBuildWorkerPoolIDentity) *string {
	e := fmt.Sprintf("%sprojects/%s/locations/%s/workerPools/%s", serviceDomain, id.project, id.location, id.workerpool)
	return &e
}

// asID builds a CloudBuildWorkerPoolIDentity from a externalRef
func asID(externalRef string) (*CloudBuildWorkerPoolIDentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef shall has prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain)
	tokens := strings.Split(path, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "workerPools" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/workerPools/<workerPool>, got %s",
			serviceDomain, externalRef)
	}
	return &CloudBuildWorkerPoolIDentity{
		project:    tokens[1],
		location:   tokens[3],
		workerpool: tokens[5],
	}, nil
}

// fromRaw builds a CloudBuildWorkerPoolIDentity from resource components.
func fromRaw(project, location, workerpool string) *CloudBuildWorkerPoolIDentity {
	return &CloudBuildWorkerPoolIDentity{
		project:    project,
		location:   location,
		workerpool: workerpool,
	}
}
