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

package mockartifactregistry

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type artifactRegistryName struct {
	Project    *projects.ProjectData
	Location   string
	Repository string
}

func (n *artifactRegistryName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/repositories/" + n.Repository
}

// parseArtifactRegistryName parses a string into a artifactRegistryName.
// The expected form is projects/<projectID>/locations/<region>/repositories/<repositoryName>
func (s *MockService) parseArtifactRegistryName(name string) (*artifactRegistryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "repositories" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &artifactRegistryName{
			Project:    project,
			Location:   tokens[3],
			Repository: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type ruleName struct {
	Project    *projects.ProjectData
	Location   string
	Repository string
	Rule       string
}

func (n *ruleName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/repositories/" + n.Repository + "/rules/" + n.Rule
}

func (n *ruleName) Parent() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/repositories/" + n.Repository
}

// parseRuleName parses a string into a ruleName.
// The expected form is projects/<projectID>/locations/<region>/repositories/<repositoryName>/rules/<ruleName>
func (s *MockService) parseRuleName(name string) (*ruleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "repositories" && tokens[6] == "rules" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &ruleName{
			Project:    project,
			Location:   tokens[3],
			Repository: tokens[5],
			Rule:       tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
