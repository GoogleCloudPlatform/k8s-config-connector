// Copyright 2023 Google LLC
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

package mockcertificatemanager

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type certificateName struct {
	Project         *projects.ProjectData
	Location        string
	CertificateName string
}

func (n *certificateName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/certificates/" + n.CertificateName
}

// parseCertificateName parses a string into a certificateName.
// The expected form is `projects/*/locations/*/certificates/*`.
func (s *MockService) parseCertificateName(name string) (*certificateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "certificates" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &certificateName{
			Project:         project,
			Location:        tokens[3],
			CertificateName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type certificateMapName struct {
	Project            *projects.ProjectData
	Location           string
	CertificateMapName string
}

func (n *certificateMapName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/certificateMaps/" + n.CertificateMapName
}

// parseCertificateMapName parses a string into a certificateMapName.
// The expected form is `projects/*/locations/*/certificateMaps/*`.
func (s *MockService) parseCertificateMapName(name string) (*certificateMapName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "certificateMaps" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &certificateMapName{
			Project:            project,
			Location:           tokens[3],
			CertificateMapName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
