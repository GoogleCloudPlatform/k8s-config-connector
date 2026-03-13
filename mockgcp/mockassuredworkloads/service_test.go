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

package mockassuredworkloads_test

import (
	"context"
	"net"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/assuredworkloads/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockassuredworkloads"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestService(t *testing.T) {
	s := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{}
	service := mockassuredworkloads.New(env, s)

	// Start a local gRPC server
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	service.Register(server)
	go func() {
		if err := server.Serve(lis); err != nil {
			// t.Errorf might not be safe here, but we can log
		}
	}()
	defer server.Stop()

	// Connect to the server
	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAssuredWorkloadsServiceClient(conn)

	// Test GetWorkload (should fail as not found)
	ctx := context.Background()
	_, err = client.GetWorkload(ctx, &pb.GetWorkloadRequest{Name: "organizations/123/locations/us-central1/workloads/456"})
	if err == nil {
		t.Error("expected error, got nil")
	}

	// Test CreateWorkload
	req := &pb.CreateWorkloadRequest{
		Parent: "organizations/123/locations/us-central1",
		Workload: &pb.Workload{
			DisplayName: "test-workload",
		},
	}
	op, err := client.CreateWorkload(ctx, req)
	if err != nil {
		t.Fatalf("CreateWorkload failed: %v", err)
	}
	if op.Name == "" {
		t.Error("expected operation name, got empty")
	}
}
