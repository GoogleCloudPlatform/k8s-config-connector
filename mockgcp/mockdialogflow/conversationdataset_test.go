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

package mockdialogflow

import (
	"context"
	"strconv"
	"testing"

	pb_v2 "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

func setupMockConversationDatasetService() (*MockService, *conversationDatasetsServer) {
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
	svc := New(env, st).(*MockService)
	return svc, &conversationDatasetsServer{MockService: svc}
}

func TestParseConversationDatasetName(t *testing.T) {
	s, _ := setupMockConversationDatasetService()

	validName := "projects/test-project/locations/global/conversationDatasets/dataset-1"
	parsed, err := s.parseConversationDatasetName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid conversation dataset name: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "global" || parsed.ConversationDataset != "dataset-1" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/global",
		"projects/test-project/locations/global/conversationDatasets",
		"invalid/test-project/locations/global/conversationDatasets/dataset-1",
	}
	for _, name := range invalidNames {
		if _, err := s.parseConversationDatasetName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}

func TestParseConversationDatasetParent(t *testing.T) {
	s, _ := setupMockConversationDatasetService()

	validParent := "projects/test-project/locations/global"
	parsed, err := s.parseConversationDatasetParent(validParent)
	if err != nil {
		t.Fatalf("unexpected error parsing valid conversation dataset parent: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "global" {
		t.Errorf("unexpected parsed parent: %+v", parsed)
	}

	invalidParents := []string{
		"",
		"projects/test-project",
		"invalid/test-project/locations/global",
		"projects/test-project/locations/global/extra",
	}
	for _, parent := range invalidParents {
		if _, err := s.parseConversationDatasetParent(parent); err == nil {
			t.Errorf("expected error for invalid parent %q, got nil", parent)
		}
	}
}

func TestConversationDatasetLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockConversationDatasetService()

	parent := "projects/test-project/locations/global"

	// 1. List initially empty
	listResp, err := server.ListConversationDatasets(ctx, &pb_v2.ListConversationDatasetsRequest{
		Parent: parent,
	})
	if err != nil {
		t.Fatalf("ListConversationDatasets failed: %v", err)
	}
	if len(listResp.GetConversationDatasets()) != 0 {
		t.Fatalf("expected 0 datasets, got %d", len(listResp.GetConversationDatasets()))
	}

	// 2. Create dataset
	createReq := &pb_v2.CreateConversationDatasetRequest{
		Parent: parent,
		ConversationDataset: &pb_v2.ConversationDataset{
			DisplayName: "test-dataset",
			Description: "test description",
		},
	}
	op, err := server.CreateConversationDataset(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateConversationDataset failed: %v", err)
	}
	if op.GetDone() {
		t.Errorf("expected initial operation done=false, got done=true")
	}
	metadata := &pb_v2.CreateConversationDatasetOperationMetadata{}
	if err := op.GetMetadata().UnmarshalTo(metadata); err != nil {
		t.Fatalf("failed to unmarshal operation metadata: %v", err)
	}
	createdName := metadata.GetConversationDataset()
	if createdName == "" {
		t.Fatalf("expected non-empty conversation dataset name in metadata")
	}

	// Verify operation in storage is marked done
	getOp, err := server.operations.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.GetName()})
	if err != nil {
		t.Fatalf("GetOperation failed: %v", err)
	}
	if !getOp.GetDone() {
		t.Errorf("expected finished operation done=true, got done=false")
	}

	// 3. Get created dataset
	getResp, err := server.GetConversationDataset(ctx, &pb_v2.GetConversationDatasetRequest{
		Name: createdName,
	})
	if err != nil {
		t.Fatalf("GetConversationDataset failed: %v", err)
	}
	if getResp.GetName() != createdName {
		t.Errorf("expected name %q, got %q", createdName, getResp.GetName())
	}
	if getResp.GetDisplayName() != "test-dataset" {
		t.Errorf("expected displayName 'test-dataset', got %q", getResp.GetDisplayName())
	}
	if getResp.GetDescription() != "test description" {
		t.Errorf("expected description 'test description', got %q", getResp.GetDescription())
	}
	if getResp.GetCreateTime() == nil {
		t.Errorf("expected createTime to be set")
	}
	if getResp.GetSatisfiesPzi() != false || getResp.GetSatisfiesPzs() != false {
		t.Errorf("expected satisfiesPzi/satisfiesPzs to be false")
	}
	if getResp.GetConversationInfo() == nil {
		t.Errorf("expected conversationInfo to be initialized")
	}

	// 4. List returns the dataset
	listResp, err = server.ListConversationDatasets(ctx, &pb_v2.ListConversationDatasetsRequest{
		Parent: parent,
	})
	if err != nil {
		t.Fatalf("ListConversationDatasets failed: %v", err)
	}
	if len(listResp.GetConversationDatasets()) != 1 {
		t.Fatalf("expected 1 dataset, got %d", len(listResp.GetConversationDatasets()))
	}

	// 5. Delete dataset
	delOp, err := server.DeleteConversationDataset(ctx, &pb_v2.DeleteConversationDatasetRequest{
		Name: createdName,
	})
	if err != nil {
		t.Fatalf("DeleteConversationDataset failed: %v", err)
	}
	if !delOp.GetDone() {
		t.Errorf("expected delete operation done=true, got %v", delOp.GetDone())
	}

	// 6. Get after delete returns NotFound
	_, err = server.GetConversationDataset(ctx, &pb_v2.GetConversationDatasetRequest{
		Name: createdName,
	})
	if status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound error after deletion, got: %v", err)
	}

	// 7. Delete non-existent returns NotFound
	_, err = server.DeleteConversationDataset(ctx, &pb_v2.DeleteConversationDatasetRequest{
		Name: createdName,
	})
	if status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound error on deleting non-existent dataset, got: %v", err)
	}
}
