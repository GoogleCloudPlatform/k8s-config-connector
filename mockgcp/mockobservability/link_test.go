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
	"time"

	pb "cloud.google.com/go/observability/apiv1/observabilitypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockProjectStore struct {
	projectsByID map[string]*projects.ProjectData
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
	return m.GetProjectByIDOrNumber(projectNumberAsString)
}

func (m *mockProjectStore) GetProjectByIDOrNumber(projectIDOrNumber string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByID[projectIDOrNumber]; ok {
		return p, nil
	}
	if n, err := strconv.ParseInt(projectIDOrNumber, 10, 64); err == nil {
		p := &projects.ProjectData{
			Number: n,
			ID:     projectIDOrNumber,
		}
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectIDOrNumber)
}

func TestLinkLifecycle(t *testing.T) {
	ctx := context.Background()
	st := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{
		Projects: &mockProjectStore{
			projectsByID: map[string]*projects.ProjectData{
				"test-project": {
					Number: 1234567890,
					ID:     "test-project",
				},
			},
		},
	}

	svc := New(env, st).(*MockService)
	s := svc.observabilityServer

	parent := "projects/test-project/locations/us-central1/buckets/_Trace/datasets/Spans"
	linkID := "test-link"
	fqn := parent + "/links/" + linkID

	// 1. Get non-existent -> 404
	_, err := s.GetLink(ctx, &pb.GetLinkRequest{Name: fqn})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got %v", err)
	}

	// 2. Create
	createReq := &pb.CreateLinkRequest{
		Parent: parent,
		LinkId: linkID,
		Link: &pb.Link{
			DisplayName: "Test Display Name",
			Description: "Test Description",
		},
	}
	op, err := s.CreateLink(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateLink failed: %v", err)
	}
	if _, err := svc.operations.Wait(ctx, op.Name, 2*time.Second); err != nil {
		t.Fatalf("failed waiting for CreateLink operation: %v", err)
	}

	// 3. Get existing
	got, err := s.GetLink(ctx, &pb.GetLinkRequest{Name: fqn})
	if err != nil {
		t.Fatalf("GetLink failed: %v", err)
	}
	if got.Name != fqn {
		t.Errorf("got name %q, want %q", got.Name, fqn)
	}
	if got.DisplayName != "Test Display Name" {
		t.Errorf("got displayName %q, want %q", got.DisplayName, "Test Display Name")
	}
	if got.Description != "Test Description" {
		t.Errorf("got description %q, want %q", got.Description, "Test Description")
	}
	if got.CreateTime == nil {
		t.Errorf("expected createTime to be populated")
	}

	// 4. Update
	updateReq := &pb.UpdateLinkRequest{
		Link: &pb.Link{
			Name:        fqn,
			DisplayName: "Updated Display Name",
			Description: "Updated Description",
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"displayName", "description"},
		},
	}
	op, err = s.UpdateLink(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateLink failed: %v", err)
	}
	if _, err := svc.operations.Wait(ctx, op.Name, 2*time.Second); err != nil {
		t.Fatalf("failed waiting for UpdateLink operation: %v", err)
	}

	got, err = s.GetLink(ctx, &pb.GetLinkRequest{Name: fqn})
	if err != nil {
		t.Fatalf("GetLink after update failed: %v", err)
	}
	if got.DisplayName != "Updated Display Name" {
		t.Errorf("got displayName %q, want %q", got.DisplayName, "Updated Display Name")
	}
	if got.Description != "Updated Description" {
		t.Errorf("got description %q, want %q", got.Description, "Updated Description")
	}

	// 5. List
	listResp, err := s.ListLinks(ctx, &pb.ListLinksRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListLinks failed: %v", err)
	}
	if len(listResp.Links) != 1 {
		t.Errorf("got %d links, want 1", len(listResp.Links))
	}

	// 6. Delete
	op, err = s.DeleteLink(ctx, &pb.DeleteLinkRequest{Name: fqn})
	if err != nil {
		t.Fatalf("DeleteLink failed: %v", err)
	}
	if _, err := svc.operations.Wait(ctx, op.Name, 2*time.Second); err != nil {
		t.Fatalf("failed waiting for DeleteLink operation: %v", err)
	}

	// 7. Get after delete -> 404
	_, err = s.GetLink(ctx, &pb.GetLinkRequest{Name: fqn})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got %v", err)
	}
}
