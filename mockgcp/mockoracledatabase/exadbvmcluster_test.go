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

package mockoracledatabase

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type mockProjectStore struct {
	projectsByID map[string]*projects.ProjectData
}

func (m *mockProjectStore) GetProject(projectName *projects.ProjectName) (*projects.ProjectData, error) {
	return m.GetProjectByID(projectName.ProjectID)
}

func (m *mockProjectStore) GetProjectByID(projectID string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByID[projectID]; ok {
		return p, nil
	}
	return &projects.ProjectData{
		Number: 123456789,
		ID:     projectID,
	}, nil
}

func (m *mockProjectStore) GetProjectByNumber(projectNumberAsString string) (*projects.ProjectData, error) {
	return &projects.ProjectData{
		Number: 123456789,
		ID:     "project-123456789",
	}, nil
}

func (m *mockProjectStore) GetProjectByIDOrNumber(projectIDOrNumber string) (*projects.ProjectData, error) {
	return m.GetProjectByID(projectIDOrNumber)
}

func TestParseExadbVmClusterName(t *testing.T) {
	env := &common.MockEnvironment{
		Projects: &mockProjectStore{
			projectsByID: map[string]*projects.ProjectData{
				"test-project": {
					Number: 123456789,
					ID:     "test-project",
				},
			},
		},
	}
	st := storage.NewInMemoryStorage()
	svc := New(env, st).(*MockService)
	server := &oracleDatabaseServer{MockService: svc}

	validName := "projects/test-project/locations/us-east4/exadbVmClusters/cluster-1"
	parsed, err := server.parseExadbVmClusterName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid name: %v", err)
	}
	if parsed.Project.ID != "test-project" {
		t.Errorf("expected project ID 'test-project', got %q", parsed.Project.ID)
	}
	if parsed.Location != "us-east4" {
		t.Errorf("expected location 'us-east4', got %q", parsed.Location)
	}
	if parsed.ExadbVmCluster != "cluster-1" {
		t.Errorf("expected cluster 'cluster-1', got %q", parsed.ExadbVmCluster)
	}
	if parsed.String() != validName {
		t.Errorf("expected %q, got %q", validName, parsed.String())
	}

	invalidName := "projects/test-project/exadbVmClusters/cluster-1"
	if _, err := server.parseExadbVmClusterName(invalidName); err == nil {
		t.Errorf("expected error parsing invalid name, got nil")
	}
}
