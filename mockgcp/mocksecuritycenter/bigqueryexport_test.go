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

package mocksecuritycenter

import (
	"context"
	"testing"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestBigQueryExportLifecycle(t *testing.T) {
	ctx := context.Background()
	env := &common.MockEnvironment{}
	st := storage.NewInMemoryStorage()
	svc := New(env, st).(*MockService)
	server := &SecurityCenterServer{
		MockService: svc,
	}

	parent := "organizations/123456789/locations/global"
	exportID := "my-export"
	createdFQN := "organizations/123456789/locations/global/bigQueryExports/my-export"

	// 1. Create
	createReq := &pb.CreateBigQueryExportRequest{
		Parent:           parent,
		BigQueryExportId: exportID,
		BigQueryExport: &pb.BigQueryExport{
			Description: "Initial description",
			Filter:      "category=\"OPEN_FIREWALL\"",
			Dataset:     "projects/test-project/datasets/my_dataset",
		},
	}
	created, err := server.CreateBigQueryExport(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateBigQueryExport failed: %v", err)
	}
	if created.Name != createdFQN {
		t.Errorf("expected Name %q, got %q", createdFQN, created.Name)
	}
	if created.Description != "Initial description" {
		t.Errorf("expected Description %q, got %q", "Initial description", created.Description)
	}
	if created.Principal == "" {
		t.Errorf("expected non-empty Principal")
	}

	// 2. Get
	getReq := &pb.GetBigQueryExportRequest{
		Name: createdFQN,
	}
	got, err := server.GetBigQueryExport(ctx, getReq)
	if err != nil {
		t.Fatalf("GetBigQueryExport failed: %v", err)
	}
	if got.Name != createdFQN {
		t.Errorf("expected Name %q, got %q", createdFQN, got.Name)
	}

	// 3. List
	listReq := &pb.ListBigQueryExportsRequest{
		Parent: parent,
	}
	listResp, err := server.ListBigQueryExports(ctx, listReq)
	if err != nil {
		t.Fatalf("ListBigQueryExports failed: %v", err)
	}
	if len(listResp.BigQueryExports) != 1 {
		t.Fatalf("expected 1 export, got %d", len(listResp.BigQueryExports))
	}

	// 4. Update
	updateReq := &pb.UpdateBigQueryExportRequest{
		BigQueryExport: &pb.BigQueryExport{
			Name:        createdFQN,
			Description: "Updated description",
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"description"},
		},
	}
	updated, err := server.UpdateBigQueryExport(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateBigQueryExport failed: %v", err)
	}
	if updated.Description != "Updated description" {
		t.Errorf("expected Description 'Updated description', got %q", updated.Description)
	}
	if updated.Filter != "category=\"OPEN_FIREWALL\"" {
		t.Errorf("expected Filter preserved, got %q", updated.Filter)
	}

	// Update with invalid mask path
	invalidUpdateReq := &pb.UpdateBigQueryExportRequest{
		BigQueryExport: &pb.BigQueryExport{
			Name: createdFQN,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"invalid_field"},
		},
	}
	if _, err := server.UpdateBigQueryExport(ctx, invalidUpdateReq); status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument, got %v", err)
	}

	// 5. Delete
	deleteReq := &pb.DeleteBigQueryExportRequest{
		Name: createdFQN,
	}
	if _, err := server.DeleteBigQueryExport(ctx, deleteReq); err != nil {
		t.Fatalf("DeleteBigQueryExport failed: %v", err)
	}

	// Verify not found after delete
	if _, err := server.GetBigQueryExport(ctx, getReq); status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound after delete, got %v", err)
	}
}
