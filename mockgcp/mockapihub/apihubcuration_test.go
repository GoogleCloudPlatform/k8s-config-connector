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

package mockapihub

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockProjectStore struct {
	projectsByID     map[string]*projects.ProjectData
	projectsByNumber map[string]*projects.ProjectData
}

func (m *mockProjectStore) GetProject(projectName *projects.ProjectName) (*projects.ProjectData, error) {
	return m.GetProjectByIDOrNumber(projectName.OriginalValue)
}

func (m *mockProjectStore) GetProjectByID(projectID string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByID[projectID]; ok {
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectID)
}

func (m *mockProjectStore) GetProjectByNumber(projectNumberAsString string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByNumber[projectNumberAsString]; ok {
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectNumberAsString)
}

func (m *mockProjectStore) GetProjectByIDOrNumber(projectIDOrNumber string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByID[projectIDOrNumber]; ok {
		return p, nil
	}
	if p, ok := m.projectsByNumber[projectIDOrNumber]; ok {
		return p, nil
	}
	if n, err := strconv.ParseInt(projectIDOrNumber, 10, 64); err == nil {
		p := &projects.ProjectData{
			Number: n,
			ID:     "project-" + projectIDOrNumber,
		}
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectIDOrNumber)
}

func setupMockApiHubService() (*MockService, *ApiHubServer) {
	store := &mockProjectStore{
		projectsByID: map[string]*projects.ProjectData{
			"test-project": {
				Number: 1234567890,
				ID:     "test-project",
			},
		},
		projectsByNumber: map[string]*projects.ProjectData{
			"1234567890": {
				Number: 1234567890,
				ID:     "test-project",
			},
		},
	}
	env := &common.MockEnvironment{
		Projects: store,
	}
	st := storage.NewInMemoryStorage()
	s := &MockService{
		MockEnvironment: env,
		storage:         st,
		operations:      operations.NewOperationsService(st),
	}
	server := &ApiHubServer{MockService: s}
	s.apiHubServer = server
	return s, server
}

func TestParseCurationName(t *testing.T) {
	s, _ := setupMockApiHubService()

	validName := "projects/test-project/locations/us-central1/curations/test-curation"
	parsed, err := s.parseCurationName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid curation name: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "us-central1" || parsed.CurationName != "test-curation" {
		t.Errorf("unexpected parsed curation name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1",
		"projects/test-project/locations/us-central1/curations",
		"invalid/test-project/locations/us-central1/curations/test-curation",
	}
	for _, name := range invalidNames {
		if _, err := s.parseCurationName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}

func TestCurationLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockApiHubService()

	parent := "projects/test-project/locations/us-central1"
	curationID := "test-curation"
	name := parent + "/curations/" + curationID

	// 1. Get before creation -> 404 NotFound
	_, err := server.GetCuration(ctx, &pb.GetCurationRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent curation, got: %v", err)
	}

	// 2. Create curation
	createReq := &pb.CreateCurationRequest{
		Parent:     parent,
		CurationId: curationID,
		Curation: &pb.Curation{
			DisplayName: "Test Curation",
			Description: "Initial description",
			Endpoint: &pb.Endpoint{
				EndpointDetails: &pb.Endpoint_ApplicationIntegrationEndpointDetails{
					ApplicationIntegrationEndpointDetails: &pb.ApplicationIntegrationEndpointDetails{
						TriggerId: "api_trigger/trigger1",
						Uri:       "https://integrations.googleapis.com/v1/projects/test-project/locations/us-central1/integrations/int-1:execute",
					},
				},
			},
		},
	}
	created, err := server.CreateCuration(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateCuration failed: %v", err)
	}
	if created.GetName() != name {
		t.Errorf("expected name %q, got %q", name, created.GetName())
	}
	if created.GetDisplayName() != "Test Curation" {
		t.Errorf("expected displayName 'Test Curation', got %q", created.GetDisplayName())
	}
	if created.GetDescription() != "Initial description" {
		t.Errorf("expected description 'Initial description', got %q", created.GetDescription())
	}
	if created.GetCreateTime() == nil || created.GetUpdateTime() == nil {
		t.Errorf("expected CreateTime and UpdateTime to be populated")
	}

	// 3. Get curation -> 200 OK
	getObj, err := server.GetCuration(ctx, &pb.GetCurationRequest{Name: name})
	if err != nil {
		t.Fatalf("GetCuration failed: %v", err)
	}
	if getObj.GetName() != name {
		t.Errorf("expected name %q, got %q", name, getObj.GetName())
	}

	// 4. Update curation
	updateReq := &pb.UpdateCurationRequest{
		Curation: &pb.Curation{
			Name:        name,
			DisplayName: "Updated Curation",
			Description: "Updated description",
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"display_name", "description"},
		},
	}
	updated, err := server.UpdateCuration(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateCuration failed: %v", err)
	}
	if updated.GetDisplayName() != "Updated Curation" {
		t.Errorf("expected displayName 'Updated Curation', got %q", updated.GetDisplayName())
	}
	if updated.GetDescription() != "Updated description" {
		t.Errorf("expected description 'Updated description', got %q", updated.GetDescription())
	}

	// 5. List curations
	listResp, err := server.ListCurations(ctx, &pb.ListCurationsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListCurations failed: %v", err)
	}
	if len(listResp.GetCurations()) != 1 {
		t.Fatalf("expected 1 curation in list, got %d", len(listResp.GetCurations()))
	}
	if listResp.GetCurations()[0].GetName() != name {
		t.Errorf("expected listed curation name %q, got %q", name, listResp.GetCurations()[0].GetName())
	}

	// 6. Delete curation
	_, err = server.DeleteCuration(ctx, &pb.DeleteCurationRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteCuration failed: %v", err)
	}

	// 7. Get after delete -> 404 NotFound
	_, err = server.GetCuration(ctx, &pb.GetCurationRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}
}
