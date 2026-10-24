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

package mockvisionai

import (
	"context"
	"strconv"
	"testing"
	"time"

	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
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

func setupMockVisionAIService() (*MockService, *applicationServer) {
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
	return s, &applicationServer{MockService: s}
}

func TestApplicationLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockVisionAIService()

	parent := "projects/test-project/locations/us-central1"
	appID := "test-app"
	name := parent + "/applications/" + appID

	// 1. Initial Get -> 404 NotFound
	_, err := server.GetApplication(ctx, &pb.GetApplicationRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got: %v", err)
	}

	// 2. Create -> Success
	app := &pb.Application{
		DisplayName: "Test App",
		Description: "A test application",
		ApplicationConfigs: &pb.ApplicationConfigs{
			EventDeliveryConfig: &pb.ApplicationConfigs_EventDeliveryConfig{
				Channel:                 "//pubsub.googleapis.com/projects/test-project/topics/test-topic",
				MinimalDeliveryInterval: durationpb.New(10 * time.Second),
			},
		},
	}
	op, err := server.CreateApplication(ctx, &pb.CreateApplicationRequest{
		Parent:        parent,
		ApplicationId: appID,
		Application:   app,
	})
	if err != nil {
		t.Fatalf("CreateApplication failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected operation, got nil")
	}

	// 3. Create duplicate -> 409 AlreadyExists
	_, err = server.CreateApplication(ctx, &pb.CreateApplicationRequest{
		Parent:        parent,
		ApplicationId: appID,
		Application:   app,
	})
	if status.Code(err) != codes.AlreadyExists {
		t.Fatalf("expected AlreadyExists, got: %v", err)
	}

	// 4. Get -> Success
	got, err := server.GetApplication(ctx, &pb.GetApplicationRequest{Name: name})
	if err != nil {
		t.Fatalf("GetApplication failed: %v", err)
	}
	if got.GetName() != name {
		t.Errorf("expected name %q, got %q", name, got.GetName())
	}
	if got.GetDisplayName() != "Test App" {
		t.Errorf("expected displayName %q, got %q", "Test App", got.GetDisplayName())
	}
	if got.GetState() != pb.Application_CREATED {
		t.Errorf("expected state CREATED, got %v", got.GetState())
	}

	// 5. List -> Success
	listResp, err := server.ListApplications(ctx, &pb.ListApplicationsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListApplications failed: %v", err)
	}
	if len(listResp.GetApplications()) != 1 {
		t.Fatalf("expected 1 application, got %d", len(listResp.GetApplications()))
	}
	if listResp.GetApplications()[0].GetName() != name {
		t.Errorf("expected name %q, got %q", name, listResp.GetApplications()[0].GetName())
	}

	// 6. Update -> Success
	updatedApp := &pb.Application{
		Name:        name,
		DisplayName: "Updated App",
	}
	updateOp, err := server.UpdateApplication(ctx, &pb.UpdateApplicationRequest{
		Application: updatedApp,
		UpdateMask:  &field_mask.FieldMask{Paths: []string{"display_name"}},
	})
	if err != nil {
		t.Fatalf("UpdateApplication failed: %v", err)
	}
	if updateOp == nil {
		t.Fatalf("expected operation, got nil")
	}

	// Verify update took effect
	got, err = server.GetApplication(ctx, &pb.GetApplicationRequest{Name: name})
	if err != nil {
		t.Fatalf("GetApplication failed after update: %v", err)
	}
	if got.GetDisplayName() != "Updated App" {
		t.Errorf("expected updated displayName %q, got %q", "Updated App", got.GetDisplayName())
	}

	// 7. Delete -> Success
	delOp, err := server.DeleteApplication(ctx, &pb.DeleteApplicationRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteApplication failed: %v", err)
	}
	if delOp == nil {
		t.Fatalf("expected operation, got nil")
	}

	// 8. Get after delete -> 404 NotFound
	_, err = server.GetApplication(ctx, &pb.GetApplicationRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}

	// 9. Delete after delete -> 404 NotFound
	_, err = server.DeleteApplication(ctx, &pb.DeleteApplicationRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound on second delete, got: %v", err)
	}
}

func TestParseApplicationName(t *testing.T) {
	s, _ := setupMockVisionAIService()

	// Valid with project ID
	parsed, err := s.parseApplicationName("projects/test-project/locations/us-central1/applications/test-app")
	if err != nil {
		t.Fatalf("expected valid parse with project ID, got: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "us-central1" || parsed.Application != "test-app" {
		t.Errorf("unexpected parsed result: %+v", parsed)
	}

	// Valid with project Number
	parsed, err = s.parseApplicationName("projects/1234567890/locations/us-central1/applications/test-app")
	if err != nil {
		t.Fatalf("expected valid parse with project number, got: %v", err)
	}
	if parsed.Project.ID != "test-project" {
		t.Errorf("expected project ID 'test-project', got %q", parsed.Project.ID)
	}

	// Invalid names
	invalidNames := []string{
		"",
		"invalid",
		"projects/test-project",
		"projects/test-project/locations/us-central1",
		"projects/test-project/locations/us-central1/applications",
		"invalid/test-project/locations/us-central1/applications/test-app",
	}
	for _, name := range invalidNames {
		if _, err := s.parseApplicationName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}
