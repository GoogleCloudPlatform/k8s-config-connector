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

package mocksecretmanager

import (
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type projectStore struct {
	mutex            sync.Mutex
	projectsByID     map[string]*projectData
	projectsByNumber map[int64]*projectData
}
type projectData struct {
	Number int64
	ID     string
}

func newProjectStore() *projectStore {
	return &projectStore{
		projectsByID:     make(map[string]*projectData),
		projectsByNumber: make(map[int64]*projectData),
	}
}

func (s *projectStore) getProject(projectID string) *projectData {
	projectNumber, err := strconv.ParseInt(projectID, 10, 64)
	if err == nil {
		return s.getProjectByNumber(projectNumber)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	project := s.projectsByID[projectID]
	if project == nil {
		project = &projectData{
			Number: 123, // TODO: Fix projectid
			ID:     projectID,
		}
		s.projectsByID[project.ID] = project
		s.projectsByNumber[project.Number] = project
	}

	return project
}

func (s *projectStore) getProjectByNumber(projectNumber int64) *projectData {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	project := s.projectsByNumber[projectNumber]
	return project
}

type projectName struct {
	Project string
}

func (n *projectName) String() string {
	return "projects/" + n.Project
}

func parseProjectName(name string) (*projectName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		name := &projectName{
			Project: tokens[1],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
