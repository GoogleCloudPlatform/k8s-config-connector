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

package mockaiplatform

import (
	"context"
	"testing"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setupMockFeatureOnlineStoreAdminService() (*MockService, *featureOnlineStoreAdminService) {
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
	return s, &featureOnlineStoreAdminService{MockService: s}
}

func TestParseFeatureViewName(t *testing.T) {
	s, _ := setupMockFeatureOnlineStoreAdminService()

	validName := "projects/test-project/locations/us-central1/featureOnlineStores/fos1/featureViews/fv1"
	name, err := s.parseFeatureViewName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid name: %v", err)
	}
	if name.Project.ID != "test-project" {
		t.Errorf("expected project ID test-project, got %s", name.Project.ID)
	}
	if name.Location != "us-central1" {
		t.Errorf("expected location us-central1, got %s", name.Location)
	}
	if name.FeatureOnlineStoreID != "fos1" {
		t.Errorf("expected feature online store ID fos1, got %s", name.FeatureOnlineStoreID)
	}
	if name.FeatureViewID != "fv1" {
		t.Errorf("expected feature view ID fv1, got %s", name.FeatureViewID)
	}
	expectedFQN := "projects/1234567890/locations/us-central1/featureOnlineStores/fos1/featureViews/fv1"
	if name.String() != expectedFQN {
		t.Errorf("expected FQN %s, got %s", expectedFQN, name.String())
	}

	invalidName := "projects/test-project/locations/us-central1/invalid/fv1"
	_, err = s.parseFeatureViewName(invalidName)
	if err == nil {
		t.Fatalf("expected error parsing invalid name, got nil")
	}
}

func TestFeatureViewLifecycle(t *testing.T) {
	ctx := context.Background()
	_, service := setupMockFeatureOnlineStoreAdminService()

	parent := "projects/test-project/locations/us-central1/featureOnlineStores/fos1"
	viewID := "fv1"
	fullName := parent + "/featureViews/" + viewID

	// 1. Get non-existent view -> NotFound
	_, err := service.GetFeatureView(ctx, &pb.GetFeatureViewRequest{Name: fullName})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent view, got %v", err)
	}

	// 2. Create view without parent FeatureOnlineStore -> NotFound
	createReq := &pb.CreateFeatureViewRequest{
		Parent:        parent,
		FeatureViewId: viewID,
		FeatureView: &pb.FeatureView{
			Source: &pb.FeatureView_BigQuerySource_{
				BigQuerySource: &pb.FeatureView_BigQuerySource{
					Uri:             "bq://test-project.dataset1.table1",
					EntityIdColumns: []string{"entity_id"},
				},
			},
		},
	}
	_, err = service.CreateFeatureView(ctx, createReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound creating view without parent store, got %v", err)
	}

	// 3. Create parent FeatureOnlineStore
	_, err = service.CreateFeatureOnlineStore(ctx, &pb.CreateFeatureOnlineStoreRequest{
		Parent:               "projects/test-project/locations/us-central1",
		FeatureOnlineStoreId: "fos1",
		FeatureOnlineStore:   &pb.FeatureOnlineStore{},
	})
	if err != nil {
		t.Fatalf("failed to create parent FeatureOnlineStore: %v", err)
	}

	// 4. Create FeatureView -> Success
	op, err := service.CreateFeatureView(ctx, createReq)
	if err != nil {
		t.Fatalf("failed to create FeatureView: %v", err)
	}
	if op == nil || op.Name == "" {
		t.Fatalf("expected non-empty operation from CreateFeatureView")
	}

	// 5. Get FeatureView -> Success
	view, err := service.GetFeatureView(ctx, &pb.GetFeatureViewRequest{Name: fullName})
	if err != nil {
		t.Fatalf("failed to get FeatureView: %v", err)
	}
	if view.GetBigQuerySource().GetUri() != "bq://test-project.dataset1.table1" {
		t.Errorf("unexpected uri: %s", view.GetBigQuerySource().GetUri())
	}

	// 6. Update FeatureView (first time adding syncConfig) -> returns async LRO
	updateOp, err := service.UpdateFeatureView(ctx, &pb.UpdateFeatureViewRequest{
		FeatureView: &pb.FeatureView{
			Name: fullName,
			SyncConfig: &pb.FeatureView_SyncConfig{
				Cron: "0 0 * * *",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"sync_config"}},
	})
	if err != nil {
		t.Fatalf("failed to update FeatureView: %v", err)
	}
	if updateOp.Done {
		t.Errorf("expected Done to be false for initial syncConfig addition")
	}

	// 7. Update FeatureView (updating existing syncConfig) -> returns Done LRO
	doneOp, err := service.UpdateFeatureView(ctx, &pb.UpdateFeatureViewRequest{
		FeatureView: &pb.FeatureView{
			Name: fullName,
			SyncConfig: &pb.FeatureView_SyncConfig{
				Cron: "0 12 * * *",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"sync_config"}},
	})
	if err != nil {
		t.Fatalf("failed to update FeatureView again: %v", err)
	}
	if !doneOp.Done {
		t.Errorf("expected Done to be true for modifying existing syncConfig")
	}

	// 8. List FeatureViews
	listResp, err := service.ListFeatureViews(ctx, &pb.ListFeatureViewsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("failed to list FeatureViews: %v", err)
	}
	if len(listResp.FeatureViews) != 1 {
		t.Errorf("expected 1 FeatureView, got %d", len(listResp.FeatureViews))
	}

	// 9. Delete FeatureView
	delOp, err := service.DeleteFeatureView(ctx, &pb.DeleteFeatureViewRequest{Name: fullName})
	if err != nil {
		t.Fatalf("failed to delete FeatureView: %v", err)
	}
	if !delOp.Done {
		t.Errorf("expected delete operation to be Done")
	}

	// 10. Get FeatureView after deletion -> NotFound
	_, err = service.GetFeatureView(ctx, &pb.GetFeatureViewRequest{Name: fullName})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got %v", err)
	}
}
