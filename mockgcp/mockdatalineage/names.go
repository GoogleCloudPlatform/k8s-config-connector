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

package mockdatalineage

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type processName struct {
	Project  *projects.ProjectData
	Location string
	Process  string
}

func (n *processName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/processes/%s", n.Project.Number, n.Location, n.Process)
}

// parseProcessName parses a string into a processName.
// The expected form is projects/<projectID>/locations/<location>/processes/<process>
func (s *MockService) parseProcessName(name string) (*processName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "processes" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &processName{
			Project:  project,
			Location: tokens[3],
			Process:  tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type locationName struct {
	Project  *projects.ProjectData
	Location string
}

func (n *locationName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s", n.Project.Number, n.Location)
}

// parseLocationName parses a string into a locationName.
// The expected form is projects/<projectID>/locations/<location>
func (s *MockService) parseLocationName(name string) (*locationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &locationName{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", name)
}
