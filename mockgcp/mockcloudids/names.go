// Copyright 2024 Google LLC
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

package mockcloudids

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

// endpoint is a parsed endpoint name
type endpoint struct {
	Project  *projects.ProjectData
	Location string
	Endpoint string
}

func (e *endpoint) String() string {
	return "projects/" + e.Project.ID + "/locations/" + e.Location + "/endpoints/" + e.Endpoint
}

// parseEndpointName parses a string into a endpoint.
// The expected form is projects/<projectID>/locations/<region>/endpoints/<endpoint>
func (s *MockService) parseEndpointName(name string) (*endpoint, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "endpoints" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &endpoint{
			Project:  project,
			Location: tokens[3],
			Endpoint: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is malformed", name)
	}
}
