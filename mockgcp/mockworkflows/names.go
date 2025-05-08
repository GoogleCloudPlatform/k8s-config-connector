// Copyright 2025 Google LLC
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

package mockworkflows

// +tool:mockgcp-service
// http.host: workflows.googleapis.com
// proto.service: google.cloud.workflows.v1.Workflows

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type workflowName struct {
	Project  *projects.ProjectData
	Location string
	Workflow string
}

func (n *workflowName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/workflows/%s", n.Project.ID, n.Location, n.Workflow)
}

func (s *MockService) parseWorkflowName(name string) (*workflowName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workflows" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &workflowName{
			Project:  project,
			Location: tokens[3],
			Workflow: tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
