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

package mockmapmanagement

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type mapConfigName struct {
	Project   *projects.ProjectData
	MapConfig string
}

func (n *mapConfigName) String() string {
	return fmt.Sprintf("projects/%d/mapConfigs/%s", n.Project.Number, n.MapConfig)
}

func (s *MockService) parseMapConfigName(name string) (*mapConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "mapConfigs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &mapConfigName{
			Project:   project,
			MapConfig: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is malformed", name)
}
