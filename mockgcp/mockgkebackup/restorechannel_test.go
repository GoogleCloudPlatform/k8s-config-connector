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

package mockgkebackup

import (
	"context"
	"strconv"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkebackup/v1"
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

func setupMockBackupForGKEService() (*MockService, *BackupForGKEV1) {
	store := &mockProjectStore{
		projectsByID: map[string]*projects.ProjectData{
			"test-project": {
				Number: 1234567890,
				ID:     "test-project",
			},
			"dest-project": {
				Number: 9876543210,
				ID:     "dest-project",
			},
			"shared-reservations-project": {
				Number: 1061767337696,
				ID:     "shared-reservations-project",
			},
		},
		projectsByNumber: map[string]*projects.ProjectData{
			"1234567890": {
				Number: 1234567890,
				ID:     "test-project",
			},
			"9876543210": {
				Number: 9876543210,
				ID:     "dest-project",
			},
			"1061767337696": {
				Number: 1061767337696,
				ID:     "shared-reservations-project",
			},
		},
	}
	env := &common.MockEnvironment{
		Projects: store,
	}
	st := storage.NewInMemoryStorage()
	svc := New(env, st)
	return svc, svc.v1
}

func TestParseRestoreChannelName(t *testing.T) {
	s, _ := setupMockBackupForGKEService()

	validName := "projects/test-project/locations/us-central1/restoreChannels/my-channel"
	parsed, err := s.parseRestoreChannelName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid restore channel name: %v", err)
	}
	if parsed.Project.ID != "test-project" {
		t.Errorf("expected project ID test-project, got %s", parsed.Project.ID)
	}
	if parsed.Location != "us-central1" {
		t.Errorf("expected location us-central1, got %s", parsed.Location)
	}
	if parsed.RestoreChannelID != "my-channel" {
		t.Errorf("expected channel ID my-channel, got %s", parsed.RestoreChannelID)
	}

	invalidName := "projects/test-project/restoreChannels/my-channel"
	if _, err := s.parseRestoreChannelName(invalidName); err == nil {
		t.Fatalf("expected error for invalid restore channel name, got nil")
	}
}

func TestRestoreChannelLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockBackupForGKEService()

	parent := "projects/test-project/locations/us-central1"
	channelID := "test-channel"

	// 1. Get non-existent -> NotFound
	_, err := server.GetRestoreChannel(ctx, &pb.GetRestoreChannelRequest{
		Name: parent + "/restoreChannels/" + channelID,
	})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got %v", err)
	}

	// 2. Create with same project -> InvalidArgument
	_, err = server.CreateRestoreChannel(ctx, &pb.CreateRestoreChannelRequest{
		Parent:           parent,
		RestoreChannelId: channelID,
		RestoreChannel: &pb.RestoreChannel{
			DestinationProject: "projects/test-project",
		},
	})
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for same project, got %v", err)
	}

	// 3. Create valid RestoreChannel
	op, err := server.CreateRestoreChannel(ctx, &pb.CreateRestoreChannelRequest{
		Parent:           parent,
		RestoreChannelId: channelID,
		RestoreChannel: &pb.RestoreChannel{
			DestinationProject: "projects/shared-reservations-project",
			Description:        "Initial description",
			Labels: map[string]string{
				"env": "test",
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateRestoreChannel failed: %v", err)
	}
	if op.GetDone() {
		t.Errorf("expected initial operation done=false, got done=true")
	}

	// 4. Get created RestoreChannel
	rc, err := server.GetRestoreChannel(ctx, &pb.GetRestoreChannelRequest{
		Name: parent + "/restoreChannels/" + channelID,
	})
	if err != nil {
		t.Fatalf("GetRestoreChannel failed: %v", err)
	}
	if rc.DestinationProject != "projects/1061767337696" {
		t.Errorf("expected destination project 'projects/1061767337696', got %q", rc.DestinationProject)
	}
	if rc.DestinationProjectId != "shared-reservations-project" {
		t.Errorf("expected destination project ID 'shared-reservations-project', got %q", rc.DestinationProjectId)
	}
	if rc.Description != "Initial description" {
		t.Errorf("expected description 'Initial description', got %q", rc.Description)
	}

	// 5. Update RestoreChannel
	rc.Description = "Updated description"
	rc.Labels = map[string]string{"env": "prod"}
	updateOp, err := server.UpdateRestoreChannel(ctx, &pb.UpdateRestoreChannelRequest{
		RestoreChannel: rc,
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"description", "labels"},
		},
	})
	if err != nil {
		t.Fatalf("UpdateRestoreChannel failed: %v", err)
	}
	if updateOp.GetDone() {
		t.Errorf("expected initial operation done=false, got done=true")
	}

	updatedRC, err := server.GetRestoreChannel(ctx, &pb.GetRestoreChannelRequest{
		Name: parent + "/restoreChannels/" + channelID,
	})
	if err != nil {
		t.Fatalf("GetRestoreChannel after update failed: %v", err)
	}
	if updatedRC.Description != "Updated description" {
		t.Errorf("expected updated description, got %q", updatedRC.Description)
	}
	if updatedRC.Labels["env"] != "prod" {
		t.Errorf("expected updated label prod, got %q", updatedRC.Labels["env"])
	}

	// 6. ListRestoreChannels
	listResp, err := server.ListRestoreChannels(ctx, &pb.ListRestoreChannelsRequest{
		Parent: parent,
	})
	if err != nil {
		t.Fatalf("ListRestoreChannels failed: %v", err)
	}
	if len(listResp.RestoreChannels) != 1 {
		t.Errorf("expected 1 RestoreChannel, got %d", len(listResp.RestoreChannels))
	}

	// 7. DeleteRestoreChannel
	delOp, err := server.DeleteRestoreChannel(ctx, &pb.DeleteRestoreChannelRequest{
		Name: parent + "/restoreChannels/" + channelID,
	})
	if err != nil {
		t.Fatalf("DeleteRestoreChannel failed: %v", err)
	}
	if delOp.GetDone() {
		t.Errorf("expected initial operation done=false, got done=true")
	}

	// 8. Get after delete -> NotFound
	_, err = server.GetRestoreChannel(ctx, &pb.GetRestoreChannelRequest{
		Name: parent + "/restoreChannels/" + channelID,
	})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got %v", err)
	}
}
