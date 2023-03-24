// Copyright 2022 Google LLC
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

package projects

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectStore struct {
	mutex            sync.Mutex
	projectsByID     map[string]*ProjectData
	projectsByNumber map[int64]*ProjectData
}

type ProjectData struct {
	Number int64
	ID     string
}

func NewProjectStore() *ProjectStore {
	return &ProjectStore{
		projectsByID:     make(map[string]*ProjectData),
		projectsByNumber: make(map[int64]*ProjectData),
	}
}

func projectNotFoundError(project string) error {
	// This error follows a very specific format
	// For privacy reasons we don't want to reveal if the project exists.
	// Terraform also string-matches against the error(!!!)

	msg := fmt.Sprintf("Project '%s' not found or permission denied.", project)

	return status.Error(codes.PermissionDenied, msg)
}

func (s *ProjectStore) GetProjectByID(projectID string) (*ProjectData, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	project := s.projectsByID[projectID]
	if project == nil {
		project = &ProjectData{
			Number: 123, // TODO: Generate unique project number (and maybe require projects to be created)
			ID:     projectID,
		}
		s.projectsByID[project.ID] = project
		s.projectsByNumber[project.Number] = project
	}

	if project == nil {
		return nil, projectNotFoundError(projectID)
	}
	return project, nil
}

func (s *ProjectStore) GetProjectByNumber(projectNumberAsString string) (*ProjectData, error) {
	projectNumber, err := strconv.ParseInt(projectNumberAsString, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid project number %q", projectNumberAsString)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	project := s.projectsByNumber[projectNumber]
	if project == nil {
		// Terraform passes the project ID as 0000000 and expects that back in the error, not 0 (!!!)
		return nil, projectNotFoundError(projectNumberAsString)
	}

	return project, nil
}

type ProjectName struct {
	Project string
}

func (n *ProjectName) String() string {
	return "projects/" + n.Project
}

func ParseProjectName(name string) (*ProjectName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		name := &ProjectName{
			Project: tokens[1],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
