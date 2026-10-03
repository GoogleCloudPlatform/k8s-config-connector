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

package mockmapmanagement

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/maps/mapmanagement/apiv2beta/mapmanagementpb"
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

func setupMockService() (*MockService, *MapManagementServer) {
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
	return svc, svc.mapManagementServer
}

func TestParseStyleConfigName(t *testing.T) {
	svc, _ := setupMockService()

	tests := []struct {
		name        string
		input       string
		wantProject int64
		wantStyleID string
		wantErr     bool
	}{
		{
			name:        "valid name with project number",
			input:       "projects/1234567890/styleConfigs/abc123style",
			wantProject: 1234567890,
			wantStyleID: "abc123style",
			wantErr:     false,
		},
		{
			name:        "valid name with project id",
			input:       "projects/test-project/styleConfigs/abc123style",
			wantProject: 1234567890,
			wantStyleID: "abc123style",
			wantErr:     false,
		},
		{
			name:    "invalid format",
			input:   "projects/test-project/invalid/abc123style",
			wantErr: true,
		},
		{
			name:    "too few tokens",
			input:   "projects/test-project",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := svc.parseStyleConfigName(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("parseStyleConfigName(%q) error = %v, wantErr = %v", tc.input, err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if got.Project.Number != tc.wantProject {
				t.Errorf("Project.Number = %d, want %d", got.Project.Number, tc.wantProject)
			}
			if got.StyleConfig != tc.wantStyleID {
				t.Errorf("StyleConfig = %q, want %q", got.StyleConfig, tc.wantStyleID)
			}
			if gotStr := got.String(); gotStr != "projects/1234567890/styleConfigs/abc123style" {
				t.Errorf("String() = %q, want projects/1234567890/styleConfigs/abc123style", gotStr)
			}
		})
	}
}

func TestStyleConfigCRUD(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockService()

	// 1. Create StyleConfig
	createReq := &pb.CreateStyleConfigRequest{
		Parent: "projects/test-project",
		StyleConfig: &pb.StyleConfig{
			DisplayName: "my-test-style",
			Description: "A style for testing",
		},
	}
	created, err := server.CreateStyleConfig(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateStyleConfig failed: %v", err)
	}
	if created.DisplayName != "my-test-style" {
		t.Errorf("DisplayName = %q, want %q", created.DisplayName, "my-test-style")
	}
	if created.Description != "A style for testing" {
		t.Errorf("Description = %q, want %q", created.Description, "A style for testing")
	}
	if created.JsonStyleSheet != "null" {
		t.Errorf("JsonStyleSheet = %q, want %q", created.JsonStyleSheet, "null")
	}
	if created.StyleId == "" {
		t.Errorf("StyleId should not be empty")
	}
	if created.CreateTime == nil || created.UpdateTime == nil {
		t.Errorf("CreateTime and UpdateTime must be set")
	}

	// 2. Get StyleConfig
	getReq := &pb.GetStyleConfigRequest{
		Name: created.Name,
	}
	got, err := server.GetStyleConfig(ctx, getReq)
	if err != nil {
		t.Fatalf("GetStyleConfig failed: %v", err)
	}
	if got.Name != created.Name {
		t.Errorf("GetStyleConfig returned name = %q, want %q", got.Name, created.Name)
	}

	// 3. Update StyleConfig
	updateReq := &pb.UpdateStyleConfigRequest{
		StyleConfig: &pb.StyleConfig{
			Name:        created.Name,
			DisplayName: "my-test-style-updated",
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"display_name"},
		},
	}
	updated, err := server.UpdateStyleConfig(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateStyleConfig failed: %v", err)
	}
	if updated.DisplayName != "my-test-style-updated" {
		t.Errorf("DisplayName = %q, want %q", updated.DisplayName, "my-test-style-updated")
	}
	// Description should still be retained from previous version
	if updated.Description != "A style for testing" {
		t.Errorf("Description = %q, want %q", updated.Description, "A style for testing")
	}

	// 4. List StyleConfigs
	listReq := &pb.ListStyleConfigsRequest{
		Parent: "projects/test-project",
	}
	listResp, err := server.ListStyleConfigs(ctx, listReq)
	if err != nil {
		t.Fatalf("ListStyleConfigs failed: %v", err)
	}
	if len(listResp.StyleConfigs) != 1 {
		t.Fatalf("ListStyleConfigs returned %d items, want 1", len(listResp.StyleConfigs))
	}

	// 5. Delete StyleConfig
	deleteReq := &pb.DeleteStyleConfigRequest{
		Name: created.Name,
	}
	if _, err := server.DeleteStyleConfig(ctx, deleteReq); err != nil {
		t.Fatalf("DeleteStyleConfig failed: %v", err)
	}

	// 6. Verify NotFound after delete
	if _, err := server.GetStyleConfig(ctx, getReq); status.Code(err) != codes.NotFound {
		t.Fatalf("GetStyleConfig after delete expected NotFound, got %v", err)
	}
}
