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

package mocklicensemanager

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

// configurationName is a parsed configuration name
type configurationName struct {
	Project       *projects.ProjectData
	Location      string
	Configuration string
}

func (n *configurationName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/configurations/" + n.Configuration
}

// parseConfigurationName parses a string into a configurationName.
// The expected form is projects/<projectID>/locations/<region>/configurations/<id>
func (s *MockService) parseConfigurationName(name string) (*configurationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "configurations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &configurationName{
			Project:       project,
			Location:      tokens[3],
			Configuration: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is malformed", name)
	}
}
