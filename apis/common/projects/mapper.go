// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package projects

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
)

type ProjectMapper struct {
	client *resourcemanager.ProjectsClient
}

func NewProjectMapper(client *resourcemanager.ProjectsClient) *ProjectMapper {
	return &ProjectMapper{
		client: client,
	}
}

func (m *ProjectMapper) ReplaceProjectNumberWithID(ctx context.Context, projectID string) (string, error) {
	if _, err := strconv.ParseInt(projectID, 10, 64); err != nil {
		// Not a project number, no need to map
		return projectID, nil
	}

	req := &resourcemanagerpb.GetProjectRequest{
		Name: "projects/" + projectID,
	}
	project, err := m.client.GetProject(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error getting project %q: %w", req.Name, err)
	}
	return project.ProjectId, nil
}

// LookupProjectNumber retrieves the project number for a given project ID.
// If the project ID is actually a project number, it returns it directly.
func (m *ProjectMapper) LookupProjectNumber(ctx context.Context, projectID string) (int64, error) {
	// Check if the project number is already a valid integer
	// If not, we need to look it up
	projectNumber, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		req := &resourcemanagerpb.GetProjectRequest{
			Name: "projects/" + projectID,
		}
		project, err := m.client.GetProject(ctx, req)
		if err != nil {
			return 0, fmt.Errorf("error getting project %q: %w", req.Name, err)
		}
		n, err := strconv.ParseInt(strings.TrimPrefix(project.Name, "projects/"), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("error parsing project number for %q: %w", project.Name, err)
		}
		projectNumber = n
	}
	return projectNumber, nil
}
