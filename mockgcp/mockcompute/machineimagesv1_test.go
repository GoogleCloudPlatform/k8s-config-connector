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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newTestMachineImagesServer() (*MachineImagesV1, context.Context) {
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
	server := &MachineImagesV1{MockService: svc}
	return server, ctx
}

func TestMachineImagesV1_InsertValidation(t *testing.T) {
	// Arrange
	server, ctx := newTestMachineImagesServer()

	// Act
	_, errNilResource := server.Insert(ctx, &pb.InsertMachineImageRequest{
		Project: "test-project",
	})
	_, errEmptyName := server.Insert(ctx, &pb.InsertMachineImageRequest{
		Project:              "test-project",
		MachineImageResource: &pb.MachineImage{},
	})

	// Assert
	if status.Code(errNilResource) != codes.InvalidArgument {
		t.Errorf("Insert(nil resource) code = %v, want %v", status.Code(errNilResource), codes.InvalidArgument)
	}
	if status.Code(errEmptyName) != codes.InvalidArgument {
		t.Errorf("Insert(empty name) code = %v, want %v", status.Code(errEmptyName), codes.InvalidArgument)
	}
}

func TestMachineImagesV1_Lifecycle(t *testing.T) {
	// Arrange
	server, ctx := newTestMachineImagesServer()
	insertReq := &pb.InsertMachineImageRequest{
		Project: "test-project",
		MachineImageResource: &pb.MachineImage{
			Name:        PtrTo("my-machine-image"),
			Description: PtrTo("Test machine image"),
			Labels: map[string]string{
				"env": "test",
			},
		},
		SourceInstance: PtrTo("projects/test-project/zones/us-central1-a/instances/test-instance"),
	}

	// Act
	op, err := server.Insert(ctx, insertReq)

	// Assert
	if err != nil {
		t.Fatalf("Insert() unexpected error: %v", err)
	}
	if op.GetOperationType() != "insert" {
		t.Errorf("Insert() operationType = %q, want 'insert'", op.GetOperationType())
	}

	// Arrange (Get)
	getReq := &pb.GetMachineImageRequest{
		Project:      "test-project",
		MachineImage: "my-machine-image",
	}

	// Act (Get)
	got, err := server.Get(ctx, getReq)

	// Assert (Get)
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}
	if got.GetName() != "my-machine-image" {
		t.Errorf("Get() name = %q, want 'my-machine-image'", got.GetName())
	}
	if got.GetStatus() != "READY" {
		t.Errorf("Get() status = %q, want 'READY'", got.GetStatus())
	}
	if got.GetLabelFingerprint() == "" {
		t.Errorf("Get() expected non-empty labelFingerprint")
	}

	// Arrange (List)
	listReq := &pb.ListMachineImagesRequest{
		Project: "test-project",
	}

	// Act (List)
	listResp, err := server.List(ctx, listReq)

	// Assert (List)
	if err != nil {
		t.Fatalf("List() unexpected error: %v", err)
	}
	if len(listResp.GetItems()) != 1 {
		t.Fatalf("List() got %d items, want 1", len(listResp.GetItems()))
	}
	if listResp.GetItems()[0].GetName() != "my-machine-image" {
		t.Errorf("List() items[0].Name = %q, want 'my-machine-image'", listResp.GetItems()[0].GetName())
	}

	// Arrange (Delete)
	deleteReq := &pb.DeleteMachineImageRequest{
		Project:      "test-project",
		MachineImage: "my-machine-image",
	}

	// Act (Delete)
	delOp, err := server.Delete(ctx, deleteReq)

	// Assert (Delete)
	if err != nil {
		t.Fatalf("Delete() unexpected error: %v", err)
	}
	if delOp.GetOperationType() != "delete" {
		t.Errorf("Delete() operationType = %q, want 'delete'", delOp.GetOperationType())
	}

	// Arrange (Verify Deleted)
	getDeletedReq := &pb.GetMachineImageRequest{
		Project:      "test-project",
		MachineImage: "my-machine-image",
	}

	// Act (Verify Deleted)
	_, errNotFound := server.Get(ctx, getDeletedReq)

	// Assert (Verify Deleted)
	if status.Code(errNotFound) != codes.NotFound {
		t.Errorf("Get() deleted resource code = %v, want %v", status.Code(errNotFound), codes.NotFound)
	}
}

func TestMachineImagesV1_SetLabelsOptimisticConcurrency(t *testing.T) {
	// Arrange
	server, ctx := newTestMachineImagesServer()
	_, err := server.Insert(ctx, &pb.InsertMachineImageRequest{
		Project: "test-project",
		MachineImageResource: &pb.MachineImage{
			Name: PtrTo("concurrency-image"),
		},
	})
	if err != nil {
		t.Fatalf("Insert() unexpected error: %v", err)
	}
	current, err := server.Get(ctx, &pb.GetMachineImageRequest{
		Project:      "test-project",
		MachineImage: "concurrency-image",
	})
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}

	// Act: attempt update with invalid fingerprint
	_, errMismatch := server.SetLabels(ctx, &pb.SetLabelsMachineImageRequest{
		Project:  "test-project",
		Resource: "concurrency-image",
		GlobalSetLabelsRequestResource: &pb.GlobalSetLabelsRequest{
			LabelFingerprint: PtrTo("invalid-fingerprint"),
			Labels: map[string]string{
				"k": "v",
			},
		},
	})

	// Assert
	if status.Code(errMismatch) != codes.FailedPrecondition {
		t.Errorf("SetLabels(invalid fingerprint) code = %v, want %v", status.Code(errMismatch), codes.FailedPrecondition)
	}

	// Arrange: valid fingerprint
	validFingerprint := current.GetLabelFingerprint()

	// Act: update with correct fingerprint
	op, errValid := server.SetLabels(ctx, &pb.SetLabelsMachineImageRequest{
		Project:  "test-project",
		Resource: "concurrency-image",
		GlobalSetLabelsRequestResource: &pb.GlobalSetLabelsRequest{
			LabelFingerprint: PtrTo(validFingerprint),
			Labels: map[string]string{
				"env": "prod",
			},
		},
	})

	// Assert
	if errValid != nil {
		t.Fatalf("SetLabels(valid fingerprint) unexpected error: %v", errValid)
	}
	if op.GetOperationType() != "setLabels" {
		t.Errorf("SetLabels() operationType = %q, want 'setLabels'", op.GetOperationType())
	}

	updated, err := server.Get(ctx, &pb.GetMachineImageRequest{
		Project:      "test-project",
		MachineImage: "concurrency-image",
	})
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}
	if updated.GetLabels()["env"] != "prod" {
		t.Errorf("Get() labels['env'] = %q, want 'prod'", updated.GetLabels()["env"])
	}
}
