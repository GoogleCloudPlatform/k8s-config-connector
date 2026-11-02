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
)

func TestGcpUserAccessBinding_UniqueGroupKey(t *testing.T) {
	ctx := context.Background()
	inMemoryStorage := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{}
	mockService := New(env, inMemoryStorage).(*MockService)
	server := &AccessContextManagerV1{MockService: mockService}

	req1 := &pb.CreateGcpUserAccessBindingRequest{
		Parent: "organizations/12345",
		GcpUserAccessBinding: &pb.GcpUserAccessBinding{
			GroupKey:     "group-1",
			AccessLevels: []string{"accessPolicies/123/accessLevels/level1"},
		},
	}
	op1, err := server.CreateGcpUserAccessBinding(ctx, req1)
	if err != nil {
		t.Fatalf("CreateGcpUserAccessBinding failed: %v", err)
	}
	if op1 == nil || !op1.GetDone() {
		t.Fatalf("expected completed operation, got %v", op1)
	}

	// Creating duplicate with same groupKey in the same organization must return codes.AlreadyExists
	req2 := &pb.CreateGcpUserAccessBindingRequest{
		Parent: "organizations/12345",
		GcpUserAccessBinding: &pb.GcpUserAccessBinding{
			GroupKey:     "group-1",
			AccessLevels: []string{"accessPolicies/123/accessLevels/level2"},
		},
	}
	_, err = server.CreateGcpUserAccessBinding(ctx, req2)
	if status.Code(err) != codes.AlreadyExists {
		t.Fatalf("expected AlreadyExists for duplicate groupKey, got: %v", err)
	}

	// Creating same groupKey in a different organization should succeed
	req3 := &pb.CreateGcpUserAccessBindingRequest{
		Parent: "organizations/67890",
		GcpUserAccessBinding: &pb.GcpUserAccessBinding{
			GroupKey:     "group-1",
			AccessLevels: []string{"accessPolicies/456/accessLevels/level1"},
		},
	}
	op3, err := server.CreateGcpUserAccessBinding(ctx, req3)
	if err != nil {
		t.Fatalf("CreateGcpUserAccessBinding in another org failed: %v", err)
	}
	if op3 == nil || !op3.GetDone() {
		t.Fatalf("expected completed operation, got %v", op3)
	}
}

func TestGcpUserAccessBinding_DeleteLROName(t *testing.T) {
	ctx := context.Background()
	inMemoryStorage := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{}
	mockService := New(env, inMemoryStorage).(*MockService)
	server := &AccessContextManagerV1{MockService: mockService}

	// Deleting a non-existent binding should return an LRO with unified name format
	delReq := &pb.DeleteGcpUserAccessBindingRequest{
		Name: "organizations/12345/gcpUserAccessBindings/notfound123",
	}
	op, err := server.DeleteGcpUserAccessBinding(ctx, delReq)
	if err != nil {
		t.Fatalf("DeleteGcpUserAccessBinding failed: %v", err)
	}
	expectedPrefix := "operations/organizations/12345/gcpUserAccessBindings/notfound123/delete/"
	if !strings.HasPrefix(op.GetName(), expectedPrefix) {
		t.Errorf("expected LRO name with prefix %q, got %q", expectedPrefix, op.GetName())
	}
}
