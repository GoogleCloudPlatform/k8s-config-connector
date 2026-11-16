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
	"testing"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setupMockModelMonitoringService() (*MockService, *modelMonitoringService) {
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
	return s, &modelMonitoringService{MockService: s}
}

func TestModelMonitorLifecycle(t *testing.T) {
	ctx := context.Background()
	_, service := setupMockModelMonitoringService()

	parent := "projects/test-project/locations/us-central1"
	monitorID := "test-monitor"
	nameWithProjectID := parent + "/modelMonitors/" + monitorID
	expectedFQN := "projects/1234567890/locations/us-central1/modelMonitors/" + monitorID

	// 1. Get before create -> NotFound
	getReq := &pb.GetModelMonitorRequest{
		Name: nameWithProjectID,
	}
	if _, err := service.GetModelMonitor(ctx, getReq); status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got: %v", err)
	}

	// 2. Create
	createReq := &pb.CreateModelMonitorRequest{
		Parent:         parent,
		ModelMonitorId: monitorID,
		ModelMonitor: &pb.ModelMonitor{
			DisplayName: "My Monitor",
			ModelMonitoringSchema: &pb.ModelMonitoringSchema{
				FeatureFields: []*pb.ModelMonitoringSchema_FieldSchema{
					{
						Name:     "feature1",
						DataType: "string",
					},
				},
			},
			ModelMonitoringTarget: &pb.ModelMonitor_ModelMonitoringTarget{
				Source: &pb.ModelMonitor_ModelMonitoringTarget_VertexModel{
					VertexModel: &pb.ModelMonitor_ModelMonitoringTarget_VertexModelSource{
						Model:          "projects/test-project/locations/us-central1/models/test-model",
						ModelVersionId: "1",
					},
				},
			},
		},
	}

	createOp, err := service.CreateModelMonitor(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateModelMonitor failed: %v", err)
	}
	if createOp == nil {
		t.Fatalf("expected non-nil LRO from CreateModelMonitor")
	}

	// 3. Create duplicate -> FailedPrecondition
	if _, err := service.CreateModelMonitor(ctx, createReq); status.Code(err) != codes.FailedPrecondition {
		t.Fatalf("expected FailedPrecondition on duplicate create, got: %v", err)
	}

	// 4. Get after create (using project ID and project Number)
	got, err := service.GetModelMonitor(ctx, getReq)
	if err != nil {
		t.Fatalf("GetModelMonitor failed: %v", err)
	}
	if got.Name != expectedFQN {
		t.Errorf("expected Name %q, got %q", expectedFQN, got.Name)
	}
	if got.DisplayName != "My Monitor" {
		t.Errorf("expected DisplayName %q, got %q", "My Monitor", got.DisplayName)
	}

	getReqNum := &pb.GetModelMonitorRequest{
		Name: expectedFQN,
	}
	gotNum, err := service.GetModelMonitor(ctx, getReqNum)
	if err != nil {
		t.Fatalf("GetModelMonitor with project number failed: %v", err)
	}
	if gotNum.Name != expectedFQN {
		t.Errorf("expected Name %q, got %q", expectedFQN, gotNum.Name)
	}

	// 5. List
	listReq := &pb.ListModelMonitorsRequest{
		Parent: parent,
	}
	listResp, err := service.ListModelMonitors(ctx, listReq)
	if err != nil {
		t.Fatalf("ListModelMonitors failed: %v", err)
	}
	if len(listResp.ModelMonitors) != 1 {
		t.Fatalf("expected 1 monitor, got %d", len(listResp.ModelMonitors))
	}

	// 6. Update
	updateReq := &pb.UpdateModelMonitorRequest{
		ModelMonitor: &pb.ModelMonitor{
			Name:        nameWithProjectID,
			DisplayName: "My Monitor Updated",
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"displayName"},
		},
	}
	updateOp, err := service.UpdateModelMonitor(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateModelMonitor failed: %v", err)
	}
	if updateOp == nil {
		t.Fatalf("expected non-nil LRO from UpdateModelMonitor")
	}

	// Verify update
	gotUpdated, err := service.GetModelMonitor(ctx, getReq)
	if err != nil {
		t.Fatalf("GetModelMonitor after update failed: %v", err)
	}
	if gotUpdated.DisplayName != "My Monitor Updated" {
		t.Errorf("expected DisplayName 'My Monitor Updated', got %q", gotUpdated.DisplayName)
	}

	// 7. Delete
	deleteReq := &pb.DeleteModelMonitorRequest{
		Name: nameWithProjectID,
	}
	deleteOp, err := service.DeleteModelMonitor(ctx, deleteReq)
	if err != nil {
		t.Fatalf("DeleteModelMonitor failed: %v", err)
	}
	if deleteOp == nil {
		t.Fatalf("expected non-nil LRO from DeleteModelMonitor")
	}

	// 8. Get after delete -> NotFound
	if _, err := service.GetModelMonitor(ctx, getReq); status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got: %v", err)
	}
}

func TestModelMonitorValidation(t *testing.T) {
	ctx := context.Background()
	s, service := setupMockModelMonitoringService()

	// Invalid name
	if _, err := s.parseModelMonitorName("invalid/name"); status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for invalid name, got: %v", err)
	}

	// Invalid parent
	if _, err := s.parseModelMonitorParent("invalid/parent"); status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for invalid parent, got: %v", err)
	}

	// Invalid update field
	invalidUpdateReq := &pb.UpdateModelMonitorRequest{
		ModelMonitor: &pb.ModelMonitor{
			Name: "projects/test-project/locations/us-central1/modelMonitors/non-existent",
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"invalidField"},
		},
	}
	// Target does not exist -> NotFound
	if _, err := service.UpdateModelMonitor(ctx, invalidUpdateReq); status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound, got: %v", err)
	}
}
