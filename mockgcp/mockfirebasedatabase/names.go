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

package mockfirebasedatabase

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

// instanceName is a parsed instance name
type instanceName struct {
	Project    *projects.ProjectData
	Location   string
	InstanceID string
}

func (s *instanceName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/instances/%s", s.Project.Number, s.Location, s.InstanceID)
}

// parseInstanceName parses a string into an instanceName.
// The expected form is projects/<projectID|Number>/locations/<location>/instances/<instanceID>
func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		in := &instanceName{
			Project:    project,
			Location:   tokens[3],
			InstanceID: tokens[5],
		}

		return in, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is malformed", name)
}
