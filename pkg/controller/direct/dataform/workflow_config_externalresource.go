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

type WorkflowConfigIdentity struct {
	project        string
	location       string
	repository     string
	workflowConfig string
}

func (c *WorkflowConfigIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", c.project, c.location, c.repository)
}

func (c *WorkflowConfigIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s/workflowConfigs/%s", c.project, c.location, c.repository, c.workflowConfig)
}

func (c *WorkflowConfigIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

func asWorkflowConfigID(externalRef string) (*WorkflowConfigIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "repositories" || tokens[6] != "workflowConfigs" {
		return nil, fmt.Errorf("externalRef should be %s/projects/<project>/locations/<location>/repositories/<repository>/workflowConfigs/<workflow_config>, got %s",
			serviceDomain, externalRef)
	}
	return &WorkflowConfigIdentity{
		project:        tokens[1],
		location:       tokens[3],
		repository:     tokens[5],
		workflowConfig: tokens[7],
	}, nil
}

func BuildWorkflowConfigID(project, location, repository, workflowConfig string) *WorkflowConfigIdentity {
	return &WorkflowConfigIdentity{
		project:        project,
		location:       location,
		repository:     repository,
		workflowConfig: workflowConfig,
	}
}
