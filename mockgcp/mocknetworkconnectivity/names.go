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

package mocknetworkconnectivity

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type hubName struct {
	Project *projects.ProjectData
	HubID   string
}

func (n *hubName) String() string {
	return "projects/" + n.Project.ID + "/locations/global/hubs/" + n.HubID
}

// StringWithProjectNumber returns the hub name with project number instead of project ID.
func (n *hubName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/global/hubs/%s", n.Project.Number, n.HubID)
}

func (s *MockService) parseHubName(name string) (*hubName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[3] == "global" && tokens[4] == "hubs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &hubName{
			Project: project,
			HubID:   tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type spokeName struct {
	Project  *projects.ProjectData
	Location string
	SpokeID  string
}

func (n *spokeName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/spokes/" + n.SpokeID
}

// StringWithProjectNumber returns the spoke name with project number instead of project ID.
func (n *spokeName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/spokes/%s", n.Project.Number, n.Location, n.SpokeID)
}

func (s *MockService) parseSpokeName(name string) (*spokeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spokes" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &spokeName{
			Project:  project,
			Location: tokens[3],
			SpokeID:  tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
