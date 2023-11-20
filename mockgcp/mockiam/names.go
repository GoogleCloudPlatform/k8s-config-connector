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

package mockiam

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type serviceAccountName struct {
	Project *projects.ProjectData
	Email   string
}

func (n *serviceAccountName) String() string {
	return "projects/" + n.Project.ID + "/serviceAccounts/" + n.Email
}

func (s *MockService) parseServiceAccountName(ctx context.Context, name string) (*serviceAccountName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "serviceAccounts" {
		projectID := tokens[1]
		email := tokens[3]

		// Using `-` as a wildcard for the `PROJECT_ID` will infer the project from
		// the account. The `ACCOUNT` value can be the `email` address or the
		// `unique_id` of the service account.
		if projectID == "-" {
			tokens := strings.Split(email, "@")
			if len(tokens) == 2 && strings.HasSuffix(tokens[1], ServiceAccountSuffix) {
				projectID = strings.TrimSuffix(tokens[1], ServiceAccountSuffix)
			} else {
				// Infer from the account
				uniqueID, err := strconv.ParseInt(email, 10, 64)
				if err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "name %q not known", name)
				}

				projectNumber := uniqueID >> 32
				project, err := s.projects.GetProjectByNumber(strconv.FormatInt(projectNumber, 10))
				if err != nil {
					return nil, err
				}

				return &serviceAccountName{
					Project: project,
					Email:   email,
				}, nil
			}
		}

		project, err := s.projects.GetProjectByID(projectID)
		if err != nil {
			return nil, err
		}

		name := &serviceAccountName{
			Project: project,
			Email:   tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type workloadIdentityPoolName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *workloadIdentityPoolName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/workloadIdentityPools/" + n.Name
}

func (s *MockService) parseWorkloadIdentityPoolName(name string) (*workloadIdentityPoolName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workloadIdentityPools" {
		projectID := tokens[1]
		location := tokens[3]
		resource := tokens[5]

		project, err := s.projects.GetProjectByID(projectID)
		if err != nil {
			return nil, err
		}

		return &workloadIdentityPoolName{
			Project:  project,
			Location: location,
			Name:     resource,
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type workloadIdentityPoolProviderName struct {
	Project      *projects.ProjectData
	Location     string
	ProviderName string
	PoolName     string
}

func (n *workloadIdentityPoolProviderName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/workloadIdentityPools/" + n.PoolName + "/providers/" + n.ProviderName
}

func (s *MockService) parseWorkloadIdentityPoolProviderName(name string) (*workloadIdentityPoolProviderName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workloadIdentityPools" && tokens[6] == "providers" {
		projectID := tokens[1]
		location := tokens[3]
		poolName := tokens[5]
		providerName := tokens[7]

		project, err := s.projects.GetProjectByID(projectID)
		if err != nil {
			return nil, err
		}

		return &workloadIdentityPoolProviderName{
			Project:      project,
			Location:     location,
			PoolName:     poolName,
			ProviderName: providerName,
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
