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

package mockapihub

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type apiName struct {
	Project  *projects.ProjectData
	Location string
	ApiName  string
}

func (n *apiName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/apis/" + n.ApiName
}

// parseApiName parses a string into an apiName.
// The expected form is projects/<projectID>/locations/<location>/apis/<apiName>
func (s *MockService) parseApiName(name string) (*apiName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "apis" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &apiName{
			Project:  project,
			Location: tokens[3],
			ApiName:  tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type attributeName struct {
	Project       *projects.ProjectData
	Location      string
	AttributeName string
}

func (n *attributeName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/attributes/" + n.AttributeName
}

// parseAttributeName parses a string into an attributeName.
// The expected form is projects/<projectID>/locations/<location>/attributes/<attributeName>
func (s *MockService) parseAttributeName(name string) (*attributeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "attributes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &attributeName{
			Project:       project,
			Location:      tokens[3],
			AttributeName: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type deploymentName struct {
	Project        *projects.ProjectData
	Location       string
	DeploymentName string
}

func (n *deploymentName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/deployments/" + n.DeploymentName
}

// parseDeploymentName parses a string into a deploymentName.
// The expected form is projects/<projectID>/locations/<location>/deployments/<deploymentName>
func (s *MockService) parseDeploymentName(name string) (*deploymentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deployments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &deploymentName{
			Project:        project,
			Location:       tokens[3],
			DeploymentName: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type externalApiName struct {
	Project         *projects.ProjectData
	Location        string
	ExternalApiName string
}

func (n *externalApiName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/externalApis/" + n.ExternalApiName
}

// parseExternalApiName parses a string into an externalApiName.
// The expected form is projects/<projectID>/locations/<location>/externalApis/<externalApiName>
func (s *MockService) parseExternalApiName(name string) (*externalApiName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "externalApis" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &externalApiName{
			Project:         project,
			Location:        tokens[3],
			ExternalApiName: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type dependencyName struct {
	Project        *projects.ProjectData
	Location       string
	DependencyName string
}

func (n *dependencyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/dependencies/" + n.DependencyName
}

// parseDependencyName parses a string into a dependencyName.
// The expected form is projects/<projectID>/locations/<location>/dependencies/<dependencyName>
func (s *MockService) parseDependencyName(name string) (*dependencyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dependencies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &dependencyName{
			Project:        project,
			Location:       tokens[3],
			DependencyName: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
