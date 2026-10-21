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
	s.v1 = &CloudBuildV1{MockService: s}
	s.v2 = &CloudBuildConnectionServer{MockService: s}
	return s, s.v2
}

func TestParseConnectionName(t *testing.T) {
	s, _ := setupMockCloudBuildService()

	validName := "projects/test-project/locations/us-central1/connections/test-conn"
	parsed, err := s.parseConnectionName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid connection name: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "us-central1" || parsed.ConnectionID != "test-conn" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	validWithProjectNumber := "projects/1234567890/locations/us-central1/connections/test-conn"
	parsed, err = s.parseConnectionName(validWithProjectNumber)
	if err != nil {
		t.Fatalf("unexpected error parsing valid connection name with project number: %v", err)
	}
	if parsed.Project.Number != 1234567890 || parsed.Location != "us-central1" || parsed.ConnectionID != "test-conn" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1",
		"projects/test-project/locations/us-central1/connections",
		"projects//locations/us-central1/connections/test-conn",
		"invalid/test-project/locations/us-central1/connections/test-conn",
	}
	for _, name := range invalidNames {
		if _, err := s.parseConnectionName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}

func TestCloudBuildConnectionLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockCloudBuildService()

	parent := "projects/test-project/locations/us-central1"
	connID := "my-conn"
	fqn := "projects/test-project/locations/us-central1/connections/my-conn"

	// 1. Get nonexistent returns NotFound
	_, err := server.GetConnection(ctx, &pbv2.GetConnectionRequest{Name: fqn})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got: %v", err)
	}
	if expected := "Resource '" + fqn + "' was not found"; err.Error() != status.Errorf(codes.NotFound, "%s", expected).Error() {
		t.Errorf("expected error %q, got %q", expected, err.Error())
	}

	// 2. Update nonexistent returns NotFound
	_, err = server.UpdateConnection(ctx, &pbv2.UpdateConnectionRequest{
		Connection: &pbv2.Connection{Name: fqn},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"disabled"}},
	})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound on update, got: %v", err)
	}

	// 3. Delete nonexistent returns NotFound
	_, err = server.DeleteConnection(ctx, &pbv2.DeleteConnectionRequest{Name: fqn})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound on delete, got: %v", err)
	}

	// 4. CreateConnection
	createReq := &pbv2.CreateConnectionRequest{
		Parent:       parent,
		ConnectionId: connID,
		Connection: &pbv2.Connection{
			ConnectionConfig: &pbv2.Connection_GithubConfig{
				GithubConfig: &pbv2.GitHubConfig{
					AuthorizerCredential: &pbv2.OAuthCredential{
						OauthTokenSecretVersion: "projects/test-project/secrets/my-secret/versions/1",
					},
				},
			},
		},
	}
	op, err := server.CreateConnection(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateConnection failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil operation")
	}

	// 5. Duplicate create returns AlreadyExists
	_, err = server.CreateConnection(ctx, createReq)
	if status.Code(err) != codes.AlreadyExists {
		t.Fatalf("expected AlreadyExists on duplicate create, got: %v", err)
	}
	if expected := "Resource '" + fqn + "' already exists"; err.Error() != status.Errorf(codes.AlreadyExists, "%s", expected).Error() {
		t.Errorf("expected error %q, got %q", expected, err.Error())
	}

	// 6. GetConnection returns created resource with defaulted fields
	got, err := server.GetConnection(ctx, &pbv2.GetConnectionRequest{Name: fqn})
	if err != nil {
		t.Fatalf("GetConnection failed: %v", err)
	}
	if got.Name != fqn {
		t.Errorf("expected name %q, got %q", fqn, got.Name)
	}
	if got.Etag == "" {
		t.Errorf("expected non-empty etag")
	}
	if got.CreateTime == nil || got.UpdateTime == nil {
		t.Errorf("expected non-nil createTime and updateTime")
	}
	if got.InstallationState == nil {
		t.Fatalf("expected non-nil installationState")
	}
	if got.InstallationState.Stage != pbv2.InstallationState_PENDING_USER_OAUTH {
		t.Errorf("expected PENDING_USER_OAUTH, got %v", got.InstallationState.Stage)
	}

	// 7. UpdateConnection
	updateReq := &pbv2.UpdateConnectionRequest{
		Connection: &pbv2.Connection{
			Name:     fqn,
			Disabled: true,
			Annotations: map[string]string{
				"env": "test",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"disabled", "annotations"}},
	}
	op, err = server.UpdateConnection(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateConnection failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil operation on update")
	}

	// Verify updated fields
	got, err = server.GetConnection(ctx, &pbv2.GetConnectionRequest{Name: fqn})
	if err != nil {
		t.Fatalf("GetConnection after update failed: %v", err)
	}
	if !got.Disabled {
		t.Errorf("expected disabled to be true")
	}
	if got.Annotations["env"] != "test" {
		t.Errorf("expected annotation env=test, got %v", got.Annotations)
	}

	// 8. ListConnections with project ID
	listResp, err := server.ListConnections(ctx, &pbv2.ListConnectionsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListConnections failed: %v", err)
	}
	if len(listResp.Connections) != 1 {
		t.Fatalf("expected 1 connection, got %d", len(listResp.Connections))
	}

	// ListConnections with project number
	parentWithNumber := "projects/1234567890/locations/us-central1"
	listResp, err = server.ListConnections(ctx, &pbv2.ListConnectionsRequest{Parent: parentWithNumber})
	if err != nil {
		t.Fatalf("ListConnections with project number failed: %v", err)
	}
	if len(listResp.Connections) != 1 {
		t.Fatalf("expected 1 connection, got %d", len(listResp.Connections))
	}

	// 9. DeleteConnection
	op, err = server.DeleteConnection(ctx, &pbv2.DeleteConnectionRequest{Name: fqn})
	if err != nil {
		t.Fatalf("DeleteConnection failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil operation on delete")
	}

	// 10. Subsequent Get returns NotFound
	_, err = server.GetConnection(ctx, &pbv2.GetConnectionRequest{Name: fqn})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got: %v", err)
	}
}
