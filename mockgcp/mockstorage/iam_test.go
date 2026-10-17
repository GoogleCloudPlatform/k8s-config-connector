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

package mockstorage

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
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

func (m *mockProjectStore) GetProjectByIDOrNumber(idOrNumber string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByID[idOrNumber]; ok {
		return p, nil
	}
	if p, ok := m.projectsByNumber[idOrNumber]; ok {
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", idOrNumber)
}

func (m *mockProjectStore) GetOrCreateProject(projectName *projects.ProjectName) (*projects.ProjectData, error) {
	return m.GetProject(projectName)
}

func (m *mockProjectStore) FindProject(ctx context.Context, idOrNumber string) (*projects.ProjectData, error) {
	return m.GetProjectByIDOrNumber(idOrNumber)
}

func setupTestService() (*MockService, *storage.InMemoryStorage) {
	memStorage := storage.NewInMemoryStorage()
	projectStore := &mockProjectStore{
		projectsByID: map[string]*projects.ProjectData{
			"test-project": {
				ID:     "test-project",
				Number: 123456789,
			},
		},
		projectsByNumber: map[string]*projects.ProjectData{
			"123456789": {
				ID:     "test-project",
				Number: 123456789,
			},
		},
	}
	env := &common.MockEnvironment{
		Projects: projectStore,
	}
	svc := &MockService{
		MockEnvironment: env,
		storage:         memStorage,
		operations:      operations.NewOperationsService(memStorage),
	}
	return svc, memStorage
}

func TestBucketIAM(t *testing.T) {
	svc, memStorage := setupTestService()
	ctx := context.Background()

	bucketName := "test-bucket"

	// 1. GET IAM on non-existent bucket -> 404
	{
		req := httptest.NewRequest(http.MethodGet, "/storage/v1/b/"+bucketName+"/iam", nil)
		w := httptest.NewRecorder()
		handled := svc.handleBucketIAM(w, req)
		if !handled {
			t.Fatalf("expected handleBucketIAM to return true")
		}
		if w.Code != http.StatusNotFound {
			t.Fatalf("expected status 404, got %d", w.Code)
		}
	}

	// Create bucket in storage
	bucketObj := &pb.Bucket{
		Name:          PtrTo(bucketName),
		ProjectNumber: PtrTo(uint64(123456789)),
	}
	if err := memStorage.Create(ctx, "buckets/"+bucketName, bucketObj); err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}

	// 2. GET default IAM on existing bucket -> 200 with default bindings
	{
		req := httptest.NewRequest(http.MethodGet, "/storage/v1/b/"+bucketName+"/iam?optionsRequestedPolicyVersion=3", nil)
		w := httptest.NewRecorder()
		handled := svc.handleBucketIAM(w, req)
		if !handled {
			t.Fatalf("expected handleBucketIAM to return true")
		}
		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", w.Code)
		}

		var resp map[string]any
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}
		if resp["kind"] != "storage#policy" {
			t.Errorf("expected kind 'storage#policy', got %v", resp["kind"])
		}
		bindings, ok := resp["bindings"].([]any)
		if !ok || len(bindings) != 2 {
			t.Fatalf("expected 2 default bindings, got %v", resp["bindings"])
		}
	}

	// 3. PUT updated IAM policy -> 200 with updated bindings
	{
		putPayload := map[string]any{
			"bindings": []map[string]any{
				{
					"role":    "roles/storage.admin",
					"members": []string{"serviceAccount:test-sa@project.iam.gserviceaccount.com"},
				},
				{
					"role":    "roles/storage.legacyBucketOwner",
					"members": []string{"projectEditor:test-project", "projectOwner:test-project"},
				},
			},
			"version": 3,
		}
		putBytes, _ := json.Marshal(putPayload)
		req := httptest.NewRequest(http.MethodPut, "/storage/v1/b/"+bucketName+"/iam", bytes.NewReader(putBytes))
		w := httptest.NewRecorder()
		handled := svc.handleBucketIAM(w, req)
		if !handled {
			t.Fatalf("expected handleBucketIAM to return true")
		}
		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d: %s", w.Code, w.Body.String())
		}

		var resp map[string]any
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}
		bindings, ok := resp["bindings"].([]any)
		if !ok || len(bindings) != 2 {
			t.Fatalf("expected 2 bindings, got %v", resp["bindings"])
		}
	}

	// 4. GET after PUT -> returns updated policy
	{
		req := httptest.NewRequest(http.MethodGet, "/storage/v1/b/"+bucketName+"/iam", nil)
		w := httptest.NewRecorder()
		handled := svc.handleBucketIAM(w, req)
		if !handled {
			t.Fatalf("expected handleBucketIAM to return true")
		}
		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", w.Code)
		}

		var resp map[string]any
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}
		bindings := resp["bindings"].([]any)
		firstBinding := bindings[0].(map[string]any)
		if firstBinding["role"] != "roles/storage.admin" {
			t.Errorf("expected role 'roles/storage.admin', got %v", firstBinding["role"])
		}
	}
}
