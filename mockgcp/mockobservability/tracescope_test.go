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

package mockobservability

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/observability/apiv1/observabilitypb"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

func setupMockObservabilityService() (*MockService, *observabilityServer) {
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
	server := &observabilityServer{MockService: s}
	s.observabilityServer = server
	return s, server
}

func TestTraceScope_Lifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockObservabilityService()
	parent := "projects/test-project/locations/global"
	traceScopeID := "scope-1"
	expectedName := "projects/test-project/locations/global/traceScopes/scope-1"

	// 1. Create
	createReq := &pb.CreateTraceScopeRequest{
		Parent:       parent,
		TraceScopeId: traceScopeID,
		TraceScope: &pb.TraceScope{
			Description: "Initial description",
			ResourceNames: []string{
				"projects/test-project",
			},
		},
	}
	created, err := server.CreateTraceScope(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateTraceScope failed: %v", err)
	}
	if created.GetName() != expectedName {
		t.Errorf("expected name %q, got %q", expectedName, created.GetName())
	}
	if created.GetDescription() != "Initial description" {
		t.Errorf("expected description %q, got %q", "Initial description", created.GetDescription())
	}
	if created.GetCreateTime() == nil || created.GetUpdateTime() == nil {
		t.Errorf("expected timestamps to be populated")
	}

	// 2. Get
	getReq := &pb.GetTraceScopeRequest{Name: expectedName}
	got, err := server.GetTraceScope(ctx, getReq)
	if err != nil {
		t.Fatalf("GetTraceScope failed: %v", err)
	}
	if got.GetName() != expectedName {
		t.Errorf("expected name %q, got %q", expectedName, got.GetName())
	}
	if got.GetDescription() != "Initial description" {
		t.Errorf("expected description %q, got %q", "Initial description", got.GetDescription())
	}

	// 3. List
	listReq := &pb.ListTraceScopesRequest{Parent: parent}
	listResp, err := server.ListTraceScopes(ctx, listReq)
	if err != nil {
		t.Fatalf("ListTraceScopes failed: %v", err)
	}
	if len(listResp.GetTraceScopes()) != 1 {
		t.Fatalf("expected 1 trace scope, got %d", len(listResp.GetTraceScopes()))
	}
	if listResp.GetTraceScopes()[0].GetName() != expectedName {
		t.Errorf("expected listed scope name %q, got %q", expectedName, listResp.GetTraceScopes()[0].GetName())
	}

	// 4. Update
	updateReq := &pb.UpdateTraceScopeRequest{
		TraceScope: &pb.TraceScope{
			Name:        expectedName,
			Description: "Updated description",
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"description"},
		},
	}
	updated, err := server.UpdateTraceScope(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateTraceScope failed: %v", err)
	}
	if updated.GetDescription() != "Updated description" {
		t.Errorf("expected updated description %q, got %q", "Updated description", updated.GetDescription())
	}
	// Verify resourceNames was preserved
	if len(updated.GetResourceNames()) != 1 || updated.GetResourceNames()[0] != "projects/test-project" {
		t.Errorf("expected resourceNames to be preserved, got %v", updated.GetResourceNames())
	}

	// 5. Delete
	deleteReq := &pb.DeleteTraceScopeRequest{Name: expectedName}
	if _, err := server.DeleteTraceScope(ctx, deleteReq); err != nil {
		t.Fatalf("DeleteTraceScope failed: %v", err)
	}

	// 6. Get after Delete -> 404
	_, err = server.GetTraceScope(ctx, getReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got %v", err)
	}
}

func TestTraceScope_ValidationErrors(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockObservabilityService()

	// Missing traceScopeId
	_, err := server.CreateTraceScope(ctx, &pb.CreateTraceScopeRequest{
		Parent: "projects/test-project/locations/global",
		TraceScope: &pb.TraceScope{
			ResourceNames: []string{"projects/test-project"},
		},
	})
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for missing traceScopeId, got %v", err)
	}

	// Invalid parent format
	_, err = server.CreateTraceScope(ctx, &pb.CreateTraceScopeRequest{
		Parent:       "invalid-parent",
		TraceScopeId: "scope-1",
		TraceScope:   &pb.TraceScope{},
	})
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for invalid parent, got %v", err)
	}

	// Missing parent project
	_, err = server.CreateTraceScope(ctx, &pb.CreateTraceScopeRequest{
		Parent:       "projects/non-existent-project/locations/global",
		TraceScopeId: "scope-1",
		TraceScope:   &pb.TraceScope{},
	})
	if status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound for non-existent project, got %v", err)
	}

	// Nil TraceScope
	_, err = server.CreateTraceScope(ctx, &pb.CreateTraceScopeRequest{
		Parent:       "projects/test-project/locations/global",
		TraceScopeId: "scope-1",
	})
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for nil TraceScope, got %v", err)
	}

	// Get non-existent TraceScope
	_, err = server.GetTraceScope(ctx, &pb.GetTraceScopeRequest{
		Name: "projects/test-project/locations/global/traceScopes/non-existent",
	})
	if status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound for non-existent TraceScope, got %v", err)
	}

	// Delete non-existent TraceScope
	_, err = server.DeleteTraceScope(ctx, &pb.DeleteTraceScopeRequest{
		Name: "projects/test-project/locations/global/traceScopes/non-existent",
	})
	if status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound for non-existent TraceScope, got %v", err)
	}
}

func TestTraceScope_DefensiveCopy(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockObservabilityService()
	parent := "projects/test-project/locations/global"
	traceScopeID := "scope-defensive"
	expectedName := "projects/test-project/locations/global/traceScopes/scope-defensive"

	_, err := server.CreateTraceScope(ctx, &pb.CreateTraceScopeRequest{
		Parent:       parent,
		TraceScopeId: traceScopeID,
		TraceScope: &pb.TraceScope{
			Description:   "Original Description",
			ResourceNames: []string{"projects/test-project"},
		},
	})
	if err != nil {
		t.Fatalf("CreateTraceScope failed: %v", err)
	}

	listResp, err := server.ListTraceScopes(ctx, &pb.ListTraceScopesRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListTraceScopes failed: %v", err)
	}
	if len(listResp.GetTraceScopes()) != 1 {
		t.Fatalf("expected 1 trace scope, got %d", len(listResp.GetTraceScopes()))
	}

	// Mutate returned list item
	listResp.GetTraceScopes()[0].Description = "Mutated In-Memory"

	// Fetch again via GetTraceScope and verify storage was not mutated
	got, err := server.GetTraceScope(ctx, &pb.GetTraceScopeRequest{Name: expectedName})
	if err != nil {
		t.Fatalf("GetTraceScope failed: %v", err)
	}
	if got.GetDescription() != "Original Description" {
		t.Errorf("storage was mutated in-place; expected %q, got %q", "Original Description", got.GetDescription())
	}
}
