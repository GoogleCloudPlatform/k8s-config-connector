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

package mockdeveloperconnect

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type insightsConfigName struct {
	Project        *projects.ProjectData
	Location       string
	InsightsConfig string
}

func (n *insightsConfigName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/insightsConfigs/" + n.InsightsConfig
}

// parseInsightsConfigName parses a string into an insightsConfigName.
// The expected form is projects/<projectID>/locations/<location>/insightsConfigs/<insightsConfig>
func (s *MockService) parseInsightsConfigName(name string) (*insightsConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "insightsConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &insightsConfigName{
			Project:        project,
			Location:       tokens[3],
			InsightsConfig: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
