// Copyright 2022 Google LLC
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

package mockedgenetwork

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type networkName struct {
	Project   *projects.ProjectData
	Location  string
	Zone      string
	NetworkId string
}

func (n *networkName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/zones/" + n.Zone + "/networks/" + n.NetworkId
}

// parseNetworkName parses a string into a networkName.
// The expected form is projects/<projectID>/locations/<region>/zones/<zone>/networks/<networkId>
func (s *MockService) parseNetworkName(name string) (*networkName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "zones" && tokens[6] == "networks" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &networkName{
			Project:   project,
			Location:  tokens[3],
			Zone:      tokens[5],
			NetworkId: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type subnetName struct {
	Project  *projects.ProjectData
	Location string
	Zone     string
	SubnetId string
}

func (n *subnetName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/zones/" + n.Zone + "/subnets/" + n.SubnetId
}

// parseSubnetName parses a string into a subnetName.
// The expected form is projects/<projectID>/locations/<region>/zones/<zone>/subnets/<subnetId>
func (s *MockService) parseSubnetName(name string) (*subnetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "zones" && tokens[6] == "subnets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &subnetName{
			Project:  project,
			Location: tokens[3],
			Zone:     tokens[5],
			SubnetId: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
