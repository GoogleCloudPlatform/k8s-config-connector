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

package mockgkehub

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type membershipName struct {
	Project     *projects.ProjectData
	Location    string
	Memberships string
}

func (n *membershipName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/memberships/" + n.Memberships
}

func (s *MockService) parseMembershipName(name string) (*membershipName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "memberships" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &membershipName{
			Project:     project,
			Location:    tokens[3],
			Memberships: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type featureName struct {
	Project  *projects.ProjectData
	Location string
	Features string
}

func (n *featureName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/features/" + n.Features
}

// parseFeatureName parses a string into a featureName.
// The expected form is projects/<projectID>/locations/<region>/features/<featureName>
func (s *MockService) parseFeatureName(name string) (*featureName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "features" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &featureName{
			Project:  project,
			Location: tokens[3],
			Features: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
