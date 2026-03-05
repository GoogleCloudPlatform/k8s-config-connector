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
		project, err := s.Projects.GetProjectByID(tokens[1])
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
		project, err := s.Projects.GetProjectByID(tokens[1])
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

type dnsAuthorizationName struct {
	Project              *projects.ProjectData
	Location             string
	DNSAuthorizationName string
}

func (n *dnsAuthorizationName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/dnsAuthorizations/" + n.DNSAuthorizationName
}

func (s *MockService) parseDNSAuthorizationName(name string) (*dnsAuthorizationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dnsAuthorizations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &dnsAuthorizationName{
			Project:              project,
			Location:             tokens[3],
			DNSAuthorizationName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type certificateIssuanceConfigName struct {
	Project                       *projects.ProjectData
	Location                      string
	CertificateIssuanceConfigName string
}

func (n *certificateIssuanceConfigName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/certificateIssuanceConfigs/" + n.CertificateIssuanceConfigName
}

func (s *MockService) parseCertificateIssuanceConfigName(name string) (*certificateIssuanceConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "certificateIssuanceConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &certificateIssuanceConfigName{
			Project:                       project,
			Location:                      tokens[3],
			CertificateIssuanceConfigName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type certificateMapEntryName struct {
	Project                 *projects.ProjectData
	Location                string
	CertificateMap          string
	CertificateMapEntryName string
}

func (n *certificateMapEntryName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/certificateMaps/" + n.CertificateMap + "/certificateMapEntries/" + n.CertificateMapEntryName
}

func (s *MockService) parseCertificateMapEntryName(name string) (*certificateMapEntryName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "certificateMaps" && tokens[6] == "certificateMapEntries" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &certificateMapEntryName{
			Project:                 project,
			Location:                tokens[3],
			CertificateMap:          tokens[5],
			CertificateMapEntryName: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
