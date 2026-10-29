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

package mocksaasservicemgmt

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
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

func setupMockRolloutsService() (*MockService, *SaasRolloutsServer) {
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
	}
	s.saasRolloutsServer = &SaasRolloutsServer{MockService: s}
	return s, s.saasRolloutsServer
}

func TestRolloutKindLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockRolloutsService()

	parent := "projects/test-project/locations/us-central1"
	rolloutKindID := "test-rolloutkind"
	name := parent + "/rolloutKinds/" + rolloutKindID

	// 1. Get before creation -> 404 NotFound
	_, err := server.GetRolloutKind(ctx, &pb.GetRolloutKindRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent rollout kind, got: %v", err)
	}

	// 2. Create rollout kind
	createReq := &pb.CreateRolloutKindRequest{
		Parent:        parent,
		RolloutKindId: rolloutKindID,
		RolloutKind: &pb.RolloutKind{
			RolloutOrchestrationStrategy: "Google.Cloud.Simple.AllAtOnce",
			UnitKind:                     "projects/test-project/locations/us-central1/unitKinds/test-unit-kind",
			Labels: map[string]string{
				"env": "test",
			},
		},
	}
	created, err := server.CreateRolloutKind(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateRolloutKind failed: %v", err)
	}
	if created.GetName() != name {
		t.Errorf("expected name %q, got %q", name, created.GetName())
	}

	// 3. Get after creation
	got, err := server.GetRolloutKind(ctx, &pb.GetRolloutKindRequest{Name: name})
	if err != nil {
		t.Fatalf("GetRolloutKind failed: %v", err)
	}
	if got.GetName() != name {
		t.Errorf("expected name %q, got %q", name, got.GetName())
	}
	if got.GetRolloutOrchestrationStrategy() != "Google.Cloud.Simple.AllAtOnce" {
		t.Errorf("expected strategy Google.Cloud.Simple.AllAtOnce, got %v", got.GetRolloutOrchestrationStrategy())
	}
	if got.GetLabels()["env"] != "test" {
		t.Errorf("expected label env=test, got %v", got.GetLabels())
	}

	// 4. Update rollout kind
	updateReq := &pb.UpdateRolloutKindRequest{
		RolloutKind: &pb.RolloutKind{
			Name:                         name,
			RolloutOrchestrationStrategy: "Google.Cloud.Simple.OneLocationAtATime",
			UnitFilter:                   "unit.labels.tier == 'staging'",
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"rollout_orchestration_strategy", "unit_filter"}},
	}
	updated, err := server.UpdateRolloutKind(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateRolloutKind failed: %v", err)
	}
	if updated.GetRolloutOrchestrationStrategy() != "Google.Cloud.Simple.OneLocationAtATime" {
		t.Errorf("expected updated strategy Google.Cloud.Simple.OneLocationAtATime, got %v", updated.GetRolloutOrchestrationStrategy())
	}

	got, err = server.GetRolloutKind(ctx, &pb.GetRolloutKindRequest{Name: name})
	if err != nil {
		t.Fatalf("GetRolloutKind after update failed: %v", err)
	}
	if got.GetUnitFilter() != "unit.labels.tier == 'staging'" {
		t.Errorf("expected unitFilter unit.labels.tier == 'staging', got %v", got.GetUnitFilter())
	}
	// labels should still be preserved
	if got.GetLabels()["env"] != "test" {
		t.Errorf("expected label env=test preserved, got %v", got.GetLabels())
	}

	// 5. Delete rollout kind
	_, err = server.DeleteRolloutKind(ctx, &pb.DeleteRolloutKindRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteRolloutKind failed: %v", err)
	}

	// 6. Get after delete -> 404 NotFound
	_, err = server.GetRolloutKind(ctx, &pb.GetRolloutKindRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}
}

func TestParseRolloutKindName(t *testing.T) {
	s, _ := setupMockRolloutsService()

	validName := "projects/test-project/locations/us-central1/rolloutKinds/test-rolloutkind"
	parsed, err := s.parseRolloutKindName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid rollout kind name: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "us-central1" || parsed.RolloutKind != "test-rolloutkind" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1",
		"projects/test-project/locations/us-central1/rolloutKinds",
		"invalid/test-project/locations/us-central1/rolloutKinds/test-rolloutkind",
	}
	for _, name := range invalidNames {
		if _, err := s.parseRolloutKindName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}
