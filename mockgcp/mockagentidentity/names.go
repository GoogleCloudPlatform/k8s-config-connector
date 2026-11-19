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

package mockagentidentity

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
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
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

type authProviderName struct {
	Project      *projects.ProjectData
	Location     string
	AuthProvider string
}

func (n *authProviderName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/authProviders/%s", n.Project.ID, n.Location, n.AuthProvider)
}

func (n *authProviderName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (s *MockService) parseAuthProviderName(name string) (*authProviderName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "authProviders" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &authProviderName{
			Project:      project,
			Location:     tokens[3],
			AuthProvider: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
