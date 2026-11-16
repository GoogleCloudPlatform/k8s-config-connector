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
	"testing"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setupMockSaasDeploymentsService() (*MockService, *SaasDeploymentsServer) {
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
	s.saasDeploymentsServer = &SaasDeploymentsServer{MockService: s}
	return s, s.saasDeploymentsServer
}

func TestSaasLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockSaasDeploymentsService()

	parent := "projects/test-project/locations/us-central1"
	saasID := "test-saas"
	expectedName := parent + "/saas/" + saasID

	// 1. Get before create -> NotFound
	_, err := server.GetSaas(ctx, &pb.GetSaasRequest{Name: expectedName})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got %v", err)
	}

	// 2. Create Saas with invalid parent -> InvalidArgument
	_, err = server.CreateSaas(ctx, &pb.CreateSaasRequest{
		Parent: "invalid-parent",
		SaasId: saasID,
		Saas:   &pb.Saas{},
	})
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument, got %v", err)
	}

	// 3. Create Saas with non-existent project -> NotFound
	_, err = server.CreateSaas(ctx, &pb.CreateSaasRequest{
		Parent: "projects/nonexistent/locations/us-central1",
		SaasId: saasID,
		Saas:   &pb.Saas{},
	})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for nonexistent project, got %v", err)
	}

	// 4. Create Saas successfully
	created, err := server.CreateSaas(ctx, &pb.CreateSaasRequest{
		Parent: parent,
		SaasId: saasID,
		Saas: &pb.Saas{
			Locations: []*pb.Location{
				{Name: "us-central1"},
			},
			Labels: map[string]string{
				"env": "dev",
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateSaas failed: %v", err)
	}

	if created.Name != expectedName {
		t.Errorf("expected name %q, got %q", expectedName, created.Name)
	}
	if len(created.Locations) != 1 || created.Locations[0].Name != "us-central1" {
		t.Errorf("unexpected locations: %+v", created.Locations)
	}
	if created.CreateTime != nil || created.UpdateTime != nil || created.Etag != "" {
		t.Errorf("CreateSaas response should strip CreateTime, UpdateTime, Etag")
	}

	// 5. Get Saas after create
	got, err := server.GetSaas(ctx, &pb.GetSaasRequest{Name: expectedName})
	if err != nil {
		t.Fatalf("GetSaas failed: %v", err)
	}
	if got.CreateTime == nil || got.UpdateTime == nil || got.Etag == "" {
		t.Errorf("GetSaas response should have CreateTime, UpdateTime, Etag")
	}

	// 6. Update Saas with updateMask
	updated, err := server.UpdateSaas(ctx, &pb.UpdateSaasRequest{
		Saas: &pb.Saas{
			Name: expectedName,
			Locations: []*pb.Location{
				{Name: "us-central1"},
				{Name: "us-east1"},
			},
			Labels: map[string]string{
				"env": "prod",
			},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"locations", "labels"},
		},
	})
	if err != nil {
		t.Fatalf("UpdateSaas failed: %v", err)
	}
	if len(updated.Locations) != 2 {
		t.Errorf("expected 2 locations, got %d", len(updated.Locations))
	}
	if updated.Labels["env"] != "prod" {
		t.Errorf("expected env=prod, got %q", updated.Labels["env"])
	}
	if updated.BlueprintRepo != "" || updated.State != pb.Saas_STATE_TYPE_UNSPECIFIED || len(updated.Conditions) != 0 {
		t.Errorf("UpdateSaas response should strip blueprintRepo, state, conditions")
	}

	// 7. Get Saas after update has blueprintRepo, state, conditions
	gotAfterUpdate, err := server.GetSaas(ctx, &pb.GetSaasRequest{Name: expectedName})
	if err != nil {
		t.Fatalf("GetSaas after update failed: %v", err)
	}
	if gotAfterUpdate.State != pb.Saas_STATE_ACTIVE {
		t.Errorf("expected state ACTIVE, got %v", gotAfterUpdate.State)
	}
	if len(gotAfterUpdate.Conditions) == 0 || gotAfterUpdate.Conditions[0].Type != pb.SaasCondition_TYPE_READY {
		t.Errorf("expected READY condition, got %+v", gotAfterUpdate.Conditions)
	}
	expectedRepo := "us-central1-docker.pkg.dev/test-project/" + saasID + "-00000000"
	if gotAfterUpdate.BlueprintRepo != expectedRepo {
		t.Errorf("expected blueprintRepo %q, got %q", expectedRepo, gotAfterUpdate.BlueprintRepo)
	}

	// 8. List Saas
	listResp, err := server.ListSaas(ctx, &pb.ListSaasRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListSaas failed: %v", err)
	}
	if len(listResp.Saas) != 1 {
		t.Fatalf("expected 1 saas, got %d", len(listResp.Saas))
	}

	// 9. Delete Saas
	_, err = server.DeleteSaas(ctx, &pb.DeleteSaasRequest{Name: expectedName})
	if err != nil {
		t.Fatalf("DeleteSaas failed: %v", err)
	}

	// 10. Get Saas after delete -> NotFound
	_, err = server.GetSaas(ctx, &pb.GetSaasRequest{Name: expectedName})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got %v", err)
	}
}
