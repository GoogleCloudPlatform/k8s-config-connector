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

package mockcloudbuild

import (
	"context"
	"strconv"
	"strings"
	"testing"

	pbv2 "cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
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

func setupMockCloudBuildService() (*MockService, *CloudBuildConnectionServer) {
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
	s.v2 = &CloudBuildConnectionServer{MockService: s}
	return s, s.v2
}

func TestParseConnectionName(t *testing.T) {
	s, _ := setupMockCloudBuildService()

	tests := []struct {
		name        string
		input       string
		wantErr     bool
		wantProject string
		wantLoc     string
		wantID      string
	}{
		{
			name:        "valid project ID",
			input:       "projects/test-project/locations/us-central1/connections/my-conn",
			wantErr:     false,
			wantProject: "test-project",
			wantLoc:     "us-central1",
			wantID:      "my-conn",
		},
		{
			name:        "valid project number",
			input:       "projects/1234567890/locations/us-central1/connections/my-conn",
			wantErr:     false,
			wantProject: "test-project",
			wantLoc:     "us-central1",
			wantID:      "my-conn",
		},
		{
			name:    "invalid format too short",
			input:   "projects/test-project/connections/my-conn",
			wantErr: true,
		},
		{
			name:    "invalid format empty connection ID",
			input:   "projects/test-project/locations/us-central1/connections/",
			wantErr: true,
		},
		{
			name:    "invalid prefix",
			input:   "folders/123/locations/us-central1/connections/my-conn",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cn, err := s.parseConnectionName(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q, got nil", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for %q: %v", tt.input, err)
			}
			if cn.Project.ID != tt.wantProject {
				t.Errorf("Project.ID = %q, want %q", cn.Project.ID, tt.wantProject)
			}
			if cn.Location != tt.wantLoc {
				t.Errorf("Location = %q, want %q", cn.Location, tt.wantLoc)
			}
			if cn.ConnectionID != tt.wantID {
				t.Errorf("ConnectionID = %q, want %q", cn.ConnectionID, tt.wantID)
			}
		})
	}
}

func TestCloudBuildConnectionLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockCloudBuildService()

	parent := "projects/test-project/locations/us-central1"
	connID := "test-connection"
	name := parent + "/connections/" + connID

	// 1. Get before create -> 404 NotFound
	_, err := server.GetConnection(ctx, &pbv2.GetConnectionRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got: %v", err)
	}
	if !strings.Contains(err.Error(), "Resource '"+name+"' was not found") {
		t.Errorf("expected error message to contain 'Resource %q was not found', got %v", name, err)
	}

	// 2. Create connection
	createReq := &pbv2.CreateConnectionRequest{
		Parent:       parent,
		ConnectionId: connID,
		Connection: &pbv2.Connection{
			ConnectionConfig: &pbv2.Connection_GithubConfig{
				GithubConfig: &pbv2.GitHubConfig{
					AuthorizerCredential: &pbv2.OAuthCredential{},
				},
			},
			Annotations: map[string]string{
				"env": "test",
			},
		},
	}
	op, err := server.CreateConnection(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateConnection failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from CreateConnection")
	}

	// 3. Create again -> 409 AlreadyExists
	_, err = server.CreateConnection(ctx, createReq)
	if status.Code(err) != codes.AlreadyExists {
		t.Fatalf("expected AlreadyExists on duplicate create, got: %v", err)
	}
	if !strings.Contains(err.Error(), "Resource '"+name+"' already exists") {
		t.Errorf("expected error message to contain 'Resource %q already exists', got %v", name, err)
	}

	// 4. Get connection
	conn, err := server.GetConnection(ctx, &pbv2.GetConnectionRequest{Name: name})
	if err != nil {
		t.Fatalf("GetConnection failed: %v", err)
	}
	if conn.Name != name {
		t.Errorf("expected Name %q, got %q", name, conn.Name)
	}
	if conn.Annotations["env"] != "test" {
		t.Errorf("expected annotation env=test, got %v", conn.Annotations)
	}
	if conn.GetInstallationState() == nil {
		t.Fatalf("expected InstallationState to be populated, got nil")
	}
	if conn.GetInstallationState().Stage != pbv2.InstallationState_PENDING_USER_OAUTH {
		t.Errorf("expected PENDING_USER_OAUTH stage, got %v", conn.GetInstallationState().Stage)
	}
	if conn.Etag == "" {
		t.Errorf("expected non-empty Etag")
	}

	// 5. Get connection using project number
	connNum, err := server.GetConnection(ctx, &pbv2.GetConnectionRequest{
		Name: "projects/1234567890/locations/us-central1/connections/" + connID,
	})
	if err != nil {
		t.Fatalf("GetConnection with project number failed: %v", err)
	}
	if connNum.Name != name {
		t.Errorf("expected Name %q, got %q", name, connNum.Name)
	}

	// 6. List connections
	listResp, err := server.ListConnections(ctx, &pbv2.ListConnectionsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListConnections failed: %v", err)
	}
	if len(listResp.GetConnections()) != 1 {
		t.Fatalf("expected 1 connection, got %d", len(listResp.GetConnections()))
	}

	// 7. List connections with project number
	listRespNum, err := server.ListConnections(ctx, &pbv2.ListConnectionsRequest{
		Parent: "projects/1234567890/locations/us-central1",
	})
	if err != nil {
		t.Fatalf("ListConnections with project number failed: %v", err)
	}
	if len(listRespNum.GetConnections()) != 1 {
		t.Fatalf("expected 1 connection with project number, got %d", len(listRespNum.GetConnections()))
	}

	// 8. List connections with invalid parent
	_, err = server.ListConnections(ctx, &pbv2.ListConnectionsRequest{Parent: "invalid-parent"})
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for invalid parent, got %v", err)
	}

	// 9. Update connection without update_mask -> InvalidArgument
	_, err = server.UpdateConnection(ctx, &pbv2.UpdateConnectionRequest{
		Connection: &pbv2.Connection{Name: name},
	})
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for missing update_mask, got %v", err)
	}

	// 10. Update non-existent connection -> NotFound
	_, err = server.UpdateConnection(ctx, &pbv2.UpdateConnectionRequest{
		Connection: &pbv2.Connection{Name: parent + "/connections/nonexistent"},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"disabled"}},
	})
	if status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound for non-existent connection update, got %v", err)
	}

	// 11. Update connection successfully
	updateReq := &pbv2.UpdateConnectionRequest{
		Connection: &pbv2.Connection{
			Name:     name,
			Disabled: true,
			Annotations: map[string]string{
				"env": "production",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"disabled", "annotations"}},
	}
	op, err = server.UpdateConnection(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateConnection failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from UpdateConnection")
	}

	// Verify update
	conn, err = server.GetConnection(ctx, &pbv2.GetConnectionRequest{Name: name})
	if err != nil {
		t.Fatalf("GetConnection after update failed: %v", err)
	}
	if !conn.Disabled {
		t.Errorf("expected Disabled=true")
	}
	if conn.Annotations["env"] != "production" {
		t.Errorf("expected annotation env=production, got %v", conn.Annotations)
	}

	// 12. Delete connection
	op, err = server.DeleteConnection(ctx, &pbv2.DeleteConnectionRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteConnection failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from DeleteConnection")
	}

	// 13. Delete again -> NotFound
	_, err = server.DeleteConnection(ctx, &pbv2.DeleteConnectionRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound on DeleteConnection of deleted resource, got %v", err)
	}
	if !strings.Contains(err.Error(), "Resource '"+name+"' was not found") {
		t.Errorf("expected error message to contain 'Resource %q was not found', got %v", name, err)
	}

	// 14. Get after delete -> NotFound
	_, err = server.GetConnection(ctx, &pbv2.GetConnectionRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound on GetConnection of deleted resource, got %v", err)
	}
}
