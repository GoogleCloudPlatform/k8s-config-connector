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

package mockces

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type parentName struct {
	Project  *projects.ProjectData
	Location string
}

func (n *parentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (s *MockService) parseParent(parent string) (*parentName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &parentName{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

type appName struct {
	Project  *projects.ProjectData
	Location string
	AppID    string
}

func (n *appName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/apps/%s", n.Project.ID, n.Location, n.AppID)
}

func (s *MockService) parseAppName(name string) (*appName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "apps" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &appName{
			Project:  project,
			Location: tokens[3],
			AppID:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
