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

package mockfinancialservices

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/financialservices/apiv1/financialservicespb"
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

func setupMockFinancialServicesService() (*MockService, *amlServer) {
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
	return s, &amlServer{MockService: s}
}

func TestFinancialServicesInstanceLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockFinancialServicesService()

	parent := "projects/test-project/locations/us-central1"
	instanceID := "test-instance"
	name := parent + "/instances/" + instanceID

	// 1. Get before creation -> 404 NotFound
	_, err := server.GetInstance(ctx, &pb.GetInstanceRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent instance, got: %v", err)
	}

	// 2. Create instance
	createReq := &pb.CreateInstanceRequest{
		Parent:     parent,
		InstanceId: instanceID,
		Instance: &pb.Instance{
			KmsKey: "projects/test-project/locations/us-central1/keyRings/test-kr/cryptoKeys/test-key",
			Labels: map[string]string{
				"env": "test",
			},
		},
	}
	op, err := server.CreateInstance(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateInstance failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from CreateInstance")
	}

	// 3. Get after creation -> Active
	inst, err := server.GetInstance(ctx, &pb.GetInstanceRequest{Name: name})
	if err != nil {
		t.Fatalf("GetInstance failed: %v", err)
	}
	if inst.GetName() != name {
		t.Errorf("expected name %q, got %q", name, inst.GetName())
	}
	if inst.GetState() != pb.Instance_ACTIVE {
		t.Errorf("expected state ACTIVE, got %v", inst.GetState())
	}
	if inst.GetLabels()["env"] != "test" {
		t.Errorf("expected label env=test, got %v", inst.GetLabels())
	}

	// 4. Update instance labels
	updateReq := &pb.UpdateInstanceRequest{
		Instance: &pb.Instance{
			Name: name,
			Labels: map[string]string{
				"env": "updated",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"labels"}},
	}
	op, err = server.UpdateInstance(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateInstance failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from UpdateInstance")
	}

	inst, err = server.GetInstance(ctx, &pb.GetInstanceRequest{Name: name})
	if err != nil {
		t.Fatalf("GetInstance after update failed: %v", err)
	}
	if inst.GetLabels()["env"] != "updated" {
		t.Errorf("expected label env=updated, got %v", inst.GetLabels())
	}

	// 5. List instances
	listResp, err := server.ListInstances(ctx, &pb.ListInstancesRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListInstances failed: %v", err)
	}
	if len(listResp.GetInstances()) != 1 {
		t.Fatalf("expected 1 instance, got %d", len(listResp.GetInstances()))
	}

	// 6. Delete instance
	op, err = server.DeleteInstance(ctx, &pb.DeleteInstanceRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteInstance failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from DeleteInstance")
	}

	// 7. Get after delete -> 404 NotFound
	_, err = server.GetInstance(ctx, &pb.GetInstanceRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}
}

func TestParseInstanceName(t *testing.T) {
	s, _ := setupMockFinancialServicesService()

	validName := "projects/test-project/locations/us-central1/instances/test-instance"
	parsed, err := s.parseInstanceName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid instance name: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "us-central1" || parsed.Instance != "test-instance" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1",
		"projects/test-project/locations/us-central1/instances",
		"invalid/test-project/locations/us-central1/instances/test-instance",
	}
	for _, name := range invalidNames {
		if _, err := s.parseInstanceName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}
