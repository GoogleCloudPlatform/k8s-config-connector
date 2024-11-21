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

package mockalloydb

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type clusterName struct {
	Project     *projects.ProjectData
	Location    string
	ClusterName string
}

func (n *clusterName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clusters/" + n.ClusterName
}

func (n *clusterName) ProjectAndLocation() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
}

// parseClusterName parses a string into an alloyDBClusterName.
// The expected form is projects/<projectID>/locations/<region>/clusters/<AlloyDBClusterName>
func (s *MockService) parseClusterName(name string) (*clusterName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &clusterName{
			Project:     project,
			Location:    tokens[3],
			ClusterName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type instanceName struct {
	Project      *projects.ProjectData
	Location     string
	ClusterName  string
	InstanceName string
}

func (n *instanceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clusters/" + n.ClusterName + "/instances/" + n.InstanceName
}

func (n *instanceName) ProjectAndLocation() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
}

// parseInstanceName parses a string into an alloyDBInstanceName.
// The expected form is projects/<projectID>/locations/<region>/clusters/<cluster>/instances/<AlloyDBInstanceName>
func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" && tokens[6] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceName{
			Project:      project,
			Location:     tokens[3],
			ClusterName:  tokens[5],
			InstanceName: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
