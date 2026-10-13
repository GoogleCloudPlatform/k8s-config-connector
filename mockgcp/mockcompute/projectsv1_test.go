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

package mockcompute

import (
	"context"
	"strconv"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
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
	if _, err := strconv.ParseInt(idOrNumber, 10, 64); err == nil {
		if p, ok := m.projectsByNumber[idOrNumber]; ok {
			return p, nil
		}
	}
	return m.GetProjectByID(idOrNumber)
}

func TestProjectMetadataLifecycle(t *testing.T) {
	ctx := context.Background()
	st := storage.NewInMemoryStorage()
	projectData := &projects.ProjectData{
		ID:     "test-project",
		Number: 123456789012,
	}
	projectStore := &mockProjectStore{
		projectsByID:     map[string]*projects.ProjectData{"test-project": projectData},
		projectsByNumber: map[string]*projects.ProjectData{"123456789012": projectData},
	}
	env := &common.MockEnvironment{
		Projects: projectStore,
	}
	svc := New(env, st).(*MockService)
	server := &ProjectsV1{MockService: svc}

	// 1. Get Project initially (populates default project metadata)
	proj, err := server.Get(ctx, &pb.GetProjectRequest{
		Project: "test-project",
	})
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}
	if proj.GetName() != "test-project" {
		t.Errorf("Get() name = %q, want %q", proj.GetName(), "test-project")
	}
	if proj.GetCommonInstanceMetadata() == nil {
		t.Fatalf("Get() CommonInstanceMetadata is nil")
	}
	initialFingerprint := proj.GetCommonInstanceMetadata().GetFingerprint()
	if initialFingerprint == "" {
		t.Errorf("Get() expected non-empty fingerprint")
	}

	// 2. SetCommonInstanceMetadata
	op, err := server.SetCommonInstanceMetadata(ctx, &pb.SetCommonInstanceMetadataProjectRequest{
		Project: "test-project",
		MetadataResource: &pb.Metadata{
			Fingerprint: &initialFingerprint,
			Items: []*pb.Items{
				{
					Key:   PtrTo("foo"),
					Value: PtrTo("bar"),
				},
				{
					Key:   PtrTo("baz"),
					Value: PtrTo("bat"),
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("SetCommonInstanceMetadata() unexpected error: %v", err)
	}
	if op.GetOperationType() != "compute.projects.setCommonInstanceMetadata" {
		t.Errorf("operationType = %q, want %q", op.GetOperationType(), "compute.projects.setCommonInstanceMetadata")
	}

	// 3. Verify updated metadata
	projUpdated, err := server.Get(ctx, &pb.GetProjectRequest{
		Project: "test-project",
	})
	if err != nil {
		t.Fatalf("Get() after update error: %v", err)
	}
	items := projUpdated.GetCommonInstanceMetadata().GetItems()
	if len(items) != 2 {
		t.Fatalf("items length = %d, want 2", len(items))
	}
	newFingerprint := projUpdated.GetCommonInstanceMetadata().GetFingerprint()
	if newFingerprint == initialFingerprint {
		t.Errorf("expected fingerprint to change after update")
	}

	// 4. Mismatched fingerprint should fail with FailedPrecondition
	wrongFingerprint := "wrong-fingerprint"
	_, err = server.SetCommonInstanceMetadata(ctx, &pb.SetCommonInstanceMetadataProjectRequest{
		Project: "test-project",
		MetadataResource: &pb.Metadata{
			Fingerprint: &wrongFingerprint,
			Items: []*pb.Items{
				{Key: PtrTo("key"), Value: PtrTo("val")},
			},
		},
	})
	if status.Code(err) != codes.FailedPrecondition {
		t.Errorf("expected codes.FailedPrecondition, got: %v", err)
	}
}
