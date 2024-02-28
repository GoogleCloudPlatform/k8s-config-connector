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

package mockaccesscontextmanager

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type accessLevelName struct {
	AccessPolicyName string
	AccessLevelName  string
}

type servicePerimeterName struct {
	Project              *projects.ProjectData
	AccessPolicyName     string
	ServicePerimeterName string
}

func (n *accessLevelName) String() string {
	return "accessPolicies/" + n.AccessPolicyName + "/accessLevels/" + n.AccessLevelName
}

func (n *servicePerimeterName) String() string {
	return "accessPolicies/" + n.AccessPolicyName + "/servicePerimeters/" + n.ServicePerimeterName
}

// parseAccessLevelName parses a string into a accessLevelName.
// The expected form is accessPolicies/<accessPolicyName>/accessLevels/<accessLevelName>
func (s *MockService) parseAccessLevelName(name string) (*accessLevelName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "accessPolicies" && tokens[2] == "accessLevels" {
		name := &accessLevelName{
			AccessPolicyName: tokens[1],
			AccessLevelName:  tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

// parseAccessLevelName parses a string into a accessLevelName.
// The expected form is accessPolicies/<accessPolicyName>/accessLevels/<accessLevelName>
func (s *MockService) parseServicePerimeterName(name string) (*servicePerimeterName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "accessPolicies" && tokens[2] == "servicePerimeters" {
		name := &servicePerimeterName{
			AccessPolicyName:     tokens[1],
			ServicePerimeterName: tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
