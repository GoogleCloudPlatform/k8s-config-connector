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

package mockdesigncenter

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type parentName struct {
	Project  *projects.ProjectData
	Location string
}

func (n *parentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (s *MockService) parseParent(parent string) (*parentName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &parentName{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

type spaceName struct {
	Project  *projects.ProjectData
	Location string
	Space    string
}

func (n *spaceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s", n.Project.ID, n.Location, n.Space)
}

func (n *spaceName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (s *MockService) parseSpaceName(name string) (*spaceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &spaceName{
			Project:  project,
			Location: tokens[3],
			Space:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type catalogName struct {
	Project  *projects.ProjectData
	Location string
	Space    string
	Catalog  string
}

func (n *catalogName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/catalogs/%s", n.Project.ID, n.Location, n.Space, n.Catalog)
}

func (n *catalogName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s", n.Project.ID, n.Location, n.Space)
}

func (s *MockService) parseCatalogName(name string) (*catalogName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" && tokens[6] == "catalogs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &catalogName{
			Project:  project,
			Location: tokens[3],
			Space:    tokens[5],
			Catalog:  tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type catalogTemplateName struct {
	Project         *projects.ProjectData
	Location        string
	Space           string
	Catalog         string
	CatalogTemplate string
}

func (n *catalogTemplateName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/catalogs/%s/templates/%s", n.Project.ID, n.Location, n.Space, n.Catalog, n.CatalogTemplate)
}

func (n *catalogTemplateName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/catalogs/%s", n.Project.ID, n.Location, n.Space, n.Catalog)
}

func (s *MockService) parseCatalogTemplateName(name string) (*catalogTemplateName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" && tokens[6] == "catalogs" && tokens[8] == "templates" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &catalogTemplateName{
			Project:         project,
			Location:        tokens[3],
			Space:           tokens[5],
			Catalog:         tokens[7],
			CatalogTemplate: tokens[9],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type catalogTemplateRevisionName struct {
	Project                 *projects.ProjectData
	Location                string
	Space                   string
	Catalog                 string
	CatalogTemplate         string
	CatalogTemplateRevision string
}

func (n *catalogTemplateRevisionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/catalogs/%s/templates/%s/revisions/%s", n.Project.ID, n.Location, n.Space, n.Catalog, n.CatalogTemplate, n.CatalogTemplateRevision)
}

func (n *catalogTemplateRevisionName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/catalogs/%s/templates/%s", n.Project.ID, n.Location, n.Space, n.Catalog, n.CatalogTemplate)
}

func (s *MockService) parseCatalogTemplateRevisionName(name string) (*catalogTemplateRevisionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 12 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" && tokens[6] == "catalogs" && tokens[8] == "templates" && tokens[10] == "revisions" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &catalogTemplateRevisionName{
			Project:                 project,
			Location:                tokens[3],
			Space:                   tokens[5],
			Catalog:                 tokens[7],
			CatalogTemplate:         tokens[9],
			CatalogTemplateRevision: tokens[11],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type applicationTemplateName struct {
	Project             *projects.ProjectData
	Location            string
	Space               string
	ApplicationTemplate string
}

func (n *applicationTemplateName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/applicationTemplates/%s", n.Project.ID, n.Location, n.Space, n.ApplicationTemplate)
}

func (n *applicationTemplateName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s", n.Project.ID, n.Location, n.Space)
}

func (s *MockService) parseApplicationTemplateName(name string) (*applicationTemplateName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" && tokens[6] == "applicationTemplates" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &applicationTemplateName{
			Project:             project,
			Location:            tokens[3],
			Space:               tokens[5],
			ApplicationTemplate: tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type policyName struct {
	Project             *projects.ProjectData
	Location            string
	Space               string
	ApplicationTemplate string
	Policy              string
}

func (n *policyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/applicationTemplates/%s/policies/%s", n.Project.ID, n.Location, n.Space, n.ApplicationTemplate, n.Policy)
}

func (n *policyName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/applicationTemplates/%s", n.Project.ID, n.Location, n.Space, n.ApplicationTemplate)
}

func (s *MockService) parsePolicyName(name string) (*policyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" && tokens[6] == "applicationTemplates" && tokens[8] == "policies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &policyName{
			Project:             project,
			Location:            tokens[3],
			Space:               tokens[5],
			ApplicationTemplate: tokens[7],
			Policy:              tokens[9],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type applicationName struct {
	Project     *projects.ProjectData
	Location    string
	Space       string
	Application string
}

func (n *applicationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/applications/%s", n.Project.ID, n.Location, n.Space, n.Application)
}

func (n *applicationName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s", n.Project.ID, n.Location, n.Space)
}

func (s *MockService) parseApplicationName(name string) (*applicationName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" && tokens[6] == "applications" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &applicationName{
			Project:     project,
			Location:    tokens[3],
			Space:       tokens[5],
			Application: tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type componentName struct {
	Project             *projects.ProjectData
	Location            string
	Space               string
	ApplicationTemplate string
	Component           string
}

func (n *componentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/applicationTemplates/%s/components/%s", n.Project.ID, n.Location, n.Space, n.ApplicationTemplate, n.Component)
}

func (n *componentName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/applicationTemplates/%s", n.Project.ID, n.Location, n.Space, n.ApplicationTemplate)
}

func (s *MockService) parseComponentName(name string) (*componentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" && tokens[6] == "applicationTemplates" && tokens[8] == "components" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &componentName{
			Project:             project,
			Location:            tokens[3],
			Space:               tokens[5],
			ApplicationTemplate: tokens[7],
			Component:           tokens[9],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type connectionName struct {
	Project             *projects.ProjectData
	Location            string
	Space               string
	ApplicationTemplate string
	Component           string
	Connection          string
}

func (n *connectionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/applicationTemplates/%s/components/%s/connections/%s", n.Project.ID, n.Location, n.Space, n.ApplicationTemplate, n.Component, n.Connection)
}

func (n *connectionName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/spaces/%s/applicationTemplates/%s/components/%s", n.Project.ID, n.Location, n.Space, n.ApplicationTemplate, n.Component)
}

func (s *MockService) parseConnectionName(name string) (*connectionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 12 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "spaces" && tokens[6] == "applicationTemplates" && tokens[8] == "components" && tokens[10] == "connections" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &connectionName{
			Project:             project,
			Location:            tokens[3],
			Space:               tokens[5],
			ApplicationTemplate: tokens[7],
			Component:           tokens[9],
			Connection:          tokens[11],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
