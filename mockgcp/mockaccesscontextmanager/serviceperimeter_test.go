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

package mockaccesscontextmanager

import (
	"context"
	"strings"
	"testing"

	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestServicePerimeter_CRUD(t *testing.T) {
	ctx := context.Background()
	inMemoryStorage := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{}
	mockService := New(env, inMemoryStorage).(*MockService)
	server := &AccessContextManagerV1{MockService: mockService}

	fqn := "accessPolicies/12345/servicePerimeters/perimeter_1"
	createReq := &pb.CreateServicePerimeterRequest{
		Parent: "accessPolicies/12345",
		ServicePerimeter: &pb.ServicePerimeter{
			Name:  fqn,
			Title: "perimeter 1",
			Status: &pb.ServicePerimeterConfig{
				Resources:          []string{"projects/11111"},
				RestrictedServices: []string{"storage.googleapis.com"},
			},
		},
	}
	op, err := server.CreateServicePerimeter(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateServicePerimeter failed: %v", err)
	}
	if op == nil || !op.GetDone() {
		t.Fatalf("expected completed operation, got %v", op)
	}
	if !strings.HasPrefix(op.GetName(), "operations/"+fqn+"/create/") {
		t.Errorf("unexpected create LRO name: %q", op.GetName())
	}

	got, err := server.GetServicePerimeter(ctx, &pb.GetServicePerimeterRequest{Name: fqn})
	if err != nil {
		t.Fatalf("GetServicePerimeter failed: %v", err)
	}
	if got.GetTitle() != "perimeter 1" || len(got.GetStatus().GetRestrictedServices()) != 1 {
		t.Fatalf("unexpected ServicePerimeter after create: %v", got)
	}

	updateReq := &pb.UpdateServicePerimeterRequest{
		ServicePerimeter: &pb.ServicePerimeter{
			Name:  fqn,
			Title: "perimeter 1 updated",
			Status: &pb.ServicePerimeterConfig{
				RestrictedServices: []string{"storage.googleapis.com", "bigquery.googleapis.com"},
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"title", "status.restricted_services"},
		},
	}
	updateOp, err := server.UpdateServicePerimeter(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateServicePerimeter failed: %v", err)
	}
	if !strings.HasPrefix(updateOp.GetName(), "operations/"+fqn+"/update/") {
		t.Errorf("unexpected update LRO name: %q", updateOp.GetName())
	}

	gotAfterUpdate, err := server.GetServicePerimeter(ctx, &pb.GetServicePerimeterRequest{Name: fqn})
	if err != nil {
		t.Fatalf("GetServicePerimeter after update failed: %v", err)
	}
	if gotAfterUpdate.GetTitle() != "perimeter 1 updated" || len(gotAfterUpdate.GetStatus().GetRestrictedServices()) != 2 {
		t.Fatalf("unexpected ServicePerimeter after update: %v", gotAfterUpdate)
	}
	// Verify untouched status subfields (resources) were preserved when updating status.restricted_services
	if len(gotAfterUpdate.GetStatus().GetResources()) != 1 || gotAfterUpdate.GetStatus().GetResources()[0] != "projects/11111" {
		t.Fatalf("expected status.resources to be preserved, got: %v", gotAfterUpdate.GetStatus().GetResources())
	}

	listResp, err := server.ListServicePerimeters(ctx, &pb.ListServicePerimetersRequest{Parent: "accessPolicies/12345"})
	if err != nil {
		t.Fatalf("ListServicePerimeters failed: %v", err)
	}
	if len(listResp.GetServicePerimeters()) != 1 {
		t.Fatalf("expected 1 ServicePerimeter in list, got %d", len(listResp.GetServicePerimeters()))
	}

	delOp, err := server.DeleteServicePerimeter(ctx, &pb.DeleteServicePerimeterRequest{Name: fqn})
	if err != nil {
		t.Fatalf("DeleteServicePerimeter failed: %v", err)
	}
	if !strings.HasPrefix(delOp.GetName(), "operations/"+fqn+"/delete/") {
		t.Errorf("unexpected delete LRO name: %q", delOp.GetName())
	}

	_, err = server.GetServicePerimeter(ctx, &pb.GetServicePerimeterRequest{Name: fqn})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got: %v", err)
	}
}

func TestServicePerimeter_DeleteLROName(t *testing.T) {
	ctx := context.Background()
	inMemoryStorage := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{}
	mockService := New(env, inMemoryStorage).(*MockService)
	server := &AccessContextManagerV1{MockService: mockService}

	// Deleting a non-existent ServicePerimeter should return an LRO with unified name format
	delReq := &pb.DeleteServicePerimeterRequest{
		Name: "accessPolicies/12345/servicePerimeters/notfound123",
	}
	op, err := server.DeleteServicePerimeter(ctx, delReq)
	if err != nil {
		t.Fatalf("DeleteServicePerimeter failed: %v", err)
	}
	expectedPrefix := "operations/accessPolicies/12345/servicePerimeters/notfound123/delete/"
	if !strings.HasPrefix(op.GetName(), expectedPrefix) {
		t.Errorf("expected LRO name with prefix %q, got %q", expectedPrefix, op.GetName())
	}
}
