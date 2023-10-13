// Copyright 2023 Google LLC
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

package mockresourcemanager

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type projectName struct {
	Project *projects.ProjectData
}

func (n *projectName) String() string {
	return "projects/" + n.Project.ID
}

// parseProjectName parses a string into a projectName.
func (s *MockService) parseProjectName(name string) (*projectName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "projects" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &projectName{
			Project: project,
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
