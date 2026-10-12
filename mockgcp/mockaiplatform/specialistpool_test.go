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

package mockaiplatform

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
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

func setupMockSpecialistPoolService() (*MockService, *specialistPoolService) {
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
	return s, &specialistPoolService{MockService: s}
}

func TestSpecialistPoolLifecycle(t *testing.T) {
	ctx := context.Background()
	_, service := setupMockSpecialistPoolService()

	parent := "projects/test-project/locations/us-central1"

	// Create
	createReq := &pb.CreateSpecialistPoolRequest{
		Parent: parent,
		SpecialistPool: &pb.SpecialistPool{
			DisplayName: "test-pool",
			SpecialistManagerEmails: []string{
				"manager@example.com",
			},
			SpecialistWorkerEmails: []string{
				"worker@example.com",
			},
		},
	}

	createOp, err := service.CreateSpecialistPool(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateSpecialistPool failed: %v", err)
	}
	if createOp == nil {
		t.Fatalf("expected non-nil LRO from CreateSpecialistPool")
	}

	// List
	listReq := &pb.ListSpecialistPoolsRequest{
		Parent: parent,
	}
	listResp, err := service.ListSpecialistPools(ctx, listReq)
	if err != nil {
		t.Fatalf("ListSpecialistPools failed: %v", err)
	}
	if len(listResp.SpecialistPools) != 1 {
		t.Fatalf("expected 1 specialist pool, got %d", len(listResp.SpecialistPools))
	}

	created := listResp.SpecialistPools[0]
	if created.DisplayName != "test-pool" {
		t.Errorf("expected displayName 'test-pool', got %q", created.DisplayName)
	}
	if created.SpecialistManagersCount != 1 {
		t.Errorf("expected specialistManagersCount 1, got %d", created.SpecialistManagersCount)
	}

	// Get
	getReq := &pb.GetSpecialistPoolRequest{
		Name: created.Name,
	}
	getResp, err := service.GetSpecialistPool(ctx, getReq)
	if err != nil {
		t.Fatalf("GetSpecialistPool failed: %v", err)
	}
	if getResp.DisplayName != "test-pool" {
		t.Errorf("expected displayName 'test-pool', got %q", getResp.DisplayName)
	}

	// Update
	updateReq := &pb.UpdateSpecialistPoolRequest{
		SpecialistPool: &pb.SpecialistPool{
			Name:        created.Name,
			DisplayName: "test-pool-updated",
			SpecialistManagerEmails: []string{
				"manager1@example.com",
				"manager2@example.com",
			},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"displayName", "specialistManagerEmails"},
		},
	}
	updateOp, err := service.UpdateSpecialistPool(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateSpecialistPool failed: %v", err)
	}
	if updateOp == nil {
		t.Fatalf("expected non-nil LRO from UpdateSpecialistPool")
	}

	getResp, err = service.GetSpecialistPool(ctx, getReq)
	if err != nil {
		t.Fatalf("GetSpecialistPool after update failed: %v", err)
	}
	if getResp.DisplayName != "test-pool-updated" {
		t.Errorf("expected updated displayName 'test-pool-updated', got %q", getResp.DisplayName)
	}
	if getResp.SpecialistManagersCount != 2 {
		t.Errorf("expected updated specialistManagersCount 2, got %d", getResp.SpecialistManagersCount)
	}

	// Delete
	deleteReq := &pb.DeleteSpecialistPoolRequest{
		Name: created.Name,
	}
	deleteOp, err := service.DeleteSpecialistPool(ctx, deleteReq)
	if err != nil {
		t.Fatalf("DeleteSpecialistPool failed: %v", err)
	}
	if deleteOp == nil {
		t.Fatalf("expected non-nil LRO from DeleteSpecialistPool")
	}

	// Verify Deleted
	_, err = service.GetSpecialistPool(ctx, getReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound error after delete, got %v", err)
	}
}

func TestParseSpecialistPoolName(t *testing.T) {
	service, _ := setupMockSpecialistPoolService()

	validName := "projects/test-project/locations/us-central1/specialistPools/1234567890"
	parsed, err := service.parseSpecialistPoolName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid name: %v", err)
	}
	if parsed.Location != "us-central1" {
		t.Errorf("expected location us-central1, got %s", parsed.Location)
	}
	if parsed.SpecialistPool != "1234567890" {
		t.Errorf("expected pool id 1234567890, got %s", parsed.SpecialistPool)
	}

	invalidNumericID := "projects/test-project/locations/us-central1/specialistPools/notanumber"
	_, err = service.parseSpecialistPoolName(invalidNumericID)
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for non-numeric ID, got %v", err)
	}

	invalidFormat := "projects/test-project/specialistPools/1234567890"
	_, err = service.parseSpecialistPoolName(invalidFormat)
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for malformed name, got %v", err)
	}
}
