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

package mockvisionai

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type applicationName struct {
	Project     *projects.ProjectData
	Location    string
	Application string
}

func (n *applicationName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/applications/" + n.Application
}

// parseApplicationName parses a string into an applicationName.
// Expected format: projects/<projectID>/locations/<location>/applications/<applicationID>
func (s *MockService) parseApplicationName(name string) (*applicationName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "applications" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &applicationName{
			Project:     project,
			Location:    tokens[3],
			Application: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is malformed", name)
}
