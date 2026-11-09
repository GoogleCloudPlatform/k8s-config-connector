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

package mockfinancialservices

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

// instanceName is a parsed instance name
type instanceName struct {
	Project  *projects.ProjectData
	Location string
	Instance string
}

func (n *instanceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/instances/" + n.Instance
}

// parseInstanceName parses a string into an instanceName.
// The expected form is projects/<projectID>/locations/<region>/instances/<id>
func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceName{
			Project:  project,
			Location: tokens[3],
			Instance: tokens[5],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is malformed", name)
}

// backtestResultName is a parsed backtestResult name
type backtestResultName struct {
	Project        *projects.ProjectData
	Location       string
	Instance       string
	BacktestResult string
}

func (n *backtestResultName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/instances/" + n.Instance + "/backtestResults/" + n.BacktestResult
}

func (n *backtestResultName) Parent() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/instances/" + n.Instance
}

// parseBacktestResultName parses a string into a backtestResultName.
// The expected form is projects/<projectID>/locations/<region>/instances/<instanceID>/backtestResults/<id>
func (s *MockService) parseBacktestResultName(name string) (*backtestResultName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" && tokens[6] == "backtestResults" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backtestResultName{
			Project:        project,
			Location:       tokens[3],
			Instance:       tokens[5],
			BacktestResult: tokens[7],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is malformed", name)
}
