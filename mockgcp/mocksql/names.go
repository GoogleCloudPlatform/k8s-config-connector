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

package mocksql

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type blueGreenDeploymentName struct {
	Project                 *projects.ProjectData
	Location                string
	BlueGreenDeploymentName string
}

func (n *blueGreenDeploymentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/blueGreenDeployments/%s", n.Project.ID, n.Location, n.BlueGreenDeploymentName)
}

func (s *MockService) parseBlueGreenDeploymentName(name string) (*blueGreenDeploymentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "blueGreenDeployments" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &blueGreenDeploymentName{
			Project:                 project,
			Location:                tokens[3],
			BlueGreenDeploymentName: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
