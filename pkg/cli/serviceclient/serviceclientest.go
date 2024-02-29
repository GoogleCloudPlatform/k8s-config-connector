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

package serviceclient

import (
	"fmt"
	"testing"

	resourcemanager "google.golang.org/api/cloudresourcemanager/v1"
)

type mockServiceClient struct {
	t *testing.T
}

func NewMockServiceClient(t *testing.T) mockServiceClient { //nolint:revive
	return mockServiceClient{
		t: t,
	}
}

func (m *mockServiceClient) GetProjectFromProjectIDOrNumber(projectIDOrNumber string) (*resourcemanager.Project, error) {
	returnedProjectID := ""
	switch projectIDOrNumber {
	case "1234567890":
		returnedProjectID = "project-id-1"
	default:
		return nil, fmt.Errorf("mock does not have a project id for %v, please add a new mapping", projectIDOrNumber)
	}
	return &resourcemanager.Project{
		ProjectId: returnedProjectID,
	}, nil
}
