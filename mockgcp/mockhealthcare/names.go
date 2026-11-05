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

package mockhealthcare

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type datasetName struct {
	Project  *projects.ProjectData
	Location string
	Dataset  string
}

func (n *datasetName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/datasets/%s", n.Project.ID, n.Location, n.Dataset)
}

func (s *MockService) parseDatasetName(name string) (*datasetName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "datasets" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &datasetName{
			Project:  project,
			Location: tokens[3],
			Dataset:  tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type consentStoreName struct {
	Dataset      *datasetName
	ConsentStore string
}

func (n *consentStoreName) String() string {
	return fmt.Sprintf("%s/consentStores/%s", n.Dataset.String(), n.ConsentStore)
}

func (s *MockService) parseConsentStoreName(name string) (*consentStoreName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "datasets" && tokens[6] == "consentStores" {
		dataset, err := s.parseDatasetName(strings.Join(tokens[0:6], "/"))
		if err != nil {
			return nil, err
		}
		return &consentStoreName{
			Dataset:      dataset,
			ConsentStore: tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
