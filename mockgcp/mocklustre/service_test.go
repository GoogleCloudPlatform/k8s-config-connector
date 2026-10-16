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

package mocklustre_test

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	pb "cloud.google.com/go/lustre/apiv1/lustrepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocklustre"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type staticProjectStore struct {
	project *projects.ProjectData
}

func (s *staticProjectStore) GetProject(name *projects.ProjectName) (*projects.ProjectData, error) {
	return s.project, nil
}

func (s *staticProjectStore) GetProjectByID(id string) (*projects.ProjectData, error) {
	if id == s.project.ID {
		return s.project, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %s not found", id)
}

func (s *staticProjectStore) GetProjectByNumber(num string) (*projects.ProjectData, error) {
	return s.project, nil
}

func (s *staticProjectStore) GetProjectByIDOrNumber(idOrNum string) (*projects.ProjectData, error) {
	return s.project, nil
}

func TestServiceCreation(t *testing.T) {
	env := &common.MockEnvironment{}
	st := storage.NewInMemoryStorage()
	svc := mocklustre.New(env, st)
	if svc == nil {
		t.Fatalf("expected New() to return a non-nil MockService")
	}
	expectedHosts := svc.ExpectedHosts()
	if len(expectedHosts) == 0 || expectedHosts[0] != "lustre.googleapis.com" {
		t.Errorf("unexpected expected hosts: %v", expectedHosts)
	}
}

func TestLustreInstanceCRUD(t *testing.T) {
	ctx := context.Background()
	st := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{
		Projects: &staticProjectStore{
			project: &projects.ProjectData{ID: "test-project", Number: 123456789},
		},
	}

	svc := mocklustre.New(env, st)
	server, ok := svc.(pb.LustreServer)
	if !ok {
		t.Fatalf("expected MockService to implement pb.LustreServer")
	}

	parent := "projects/test-project/locations/us-central1-a"
	instanceID := "test-lustre"
	instanceName := parent + "/instances/" + instanceID

	// 1. Create Instance
	createReq := &pb.CreateInstanceRequest{
		Parent:     parent,
		InstanceId: instanceID,
		Instance: &pb.Instance{
			Filesystem:               "testfs",
			CapacityGib:              18000,
			Network:                  "projects/test-project/global/networks/default",
			PerUnitStorageThroughput: 250,
			Description:              "Test Lustre Instance",
			Labels: map[string]string{
				"env": "test",
			},
		},
	}

	op, err := server.CreateInstance(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateInstance failed: %v", err)
	}
	if op == nil || op.Name == "" {
		t.Fatalf("expected operation with valid name from CreateInstance, got: %v", op)
	}

	// 2. Get Instance (callback already transitioned instance to ACTIVE)
	getReq := &pb.GetInstanceRequest{
		Name: instanceName,
	}
	got, err := server.GetInstance(ctx, getReq)
	if err != nil {
		t.Fatalf("GetInstance failed: %v", err)
	}
	if got.Name != instanceName {
		t.Errorf("expected name %q, got %q", instanceName, got.Name)
	}
	if got.Filesystem != "testfs" {
		t.Errorf("expected filesystem 'testfs', got %q", got.Filesystem)
	}
	if got.CapacityGib != 18000 {
		t.Errorf("expected capacity 18000, got %d", got.CapacityGib)
	}
	if got.State != pb.Instance_ACTIVE {
		t.Errorf("expected state ACTIVE, got %v", got.State)
	}
	if got.MountPoint != "10.0.0.2@tcp:/testfs" {
		t.Errorf("expected mountPoint '10.0.0.2@tcp:/testfs', got %q", got.MountPoint)
	}

	// 3. Duplicate Create should fail with AlreadyExists
	_, err = server.CreateInstance(ctx, createReq)
	if status.Code(err) != codes.AlreadyExists {
		t.Errorf("expected AlreadyExists, got: %v", err)
	}

	// 4. Update Instance (expand capacity to 27000)
	updateReq := &pb.UpdateInstanceRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"capacity_gib"},
		},
		Instance: &pb.Instance{
			Name:        instanceName,
			CapacityGib: 27000,
		},
	}
	updateOp, err := server.UpdateInstance(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateInstance failed: %v", err)
	}
	if updateOp == nil || updateOp.Name == "" {
		t.Fatalf("expected operation with valid name from UpdateInstance, got: %v", updateOp)
	}

	updated, err := server.GetInstance(ctx, getReq)
	if err != nil {
		t.Fatalf("GetInstance after update failed: %v", err)
	}
	if updated.CapacityGib != 27000 {
		t.Errorf("expected capacity 27000, got %d", updated.CapacityGib)
	}
	if updated.Filesystem != "testfs" {
		t.Errorf("expected filesystem preserved as 'testfs', got %q", updated.Filesystem)
	}

	// 5. List Instances
	listReq := &pb.ListInstancesRequest{
		Parent: parent,
	}
	listResp, err := server.ListInstances(ctx, listReq)
	if err != nil {
		t.Fatalf("ListInstances failed: %v", err)
	}
	if len(listResp.Instances) != 1 {
		t.Errorf("expected 1 instance in list, got %d", len(listResp.Instances))
	}

	// List with location wildcard "-"
	listWildcardReq := &pb.ListInstancesRequest{
		Parent: "projects/test-project/locations/-",
	}
	listWildcardResp, err := server.ListInstances(ctx, listWildcardReq)
	if err != nil {
		t.Fatalf("ListInstances with wildcard failed: %v", err)
	}
	if len(listWildcardResp.Instances) != 1 {
		t.Errorf("expected 1 instance in wildcard list, got %d", len(listWildcardResp.Instances))
	}

	// 6. Delete Instance
	deleteReq := &pb.DeleteInstanceRequest{
		Name: instanceName,
	}
	deleteOp, err := server.DeleteInstance(ctx, deleteReq)
	if err != nil {
		t.Fatalf("DeleteInstance failed: %v", err)
	}
	if deleteOp == nil || deleteOp.Name == "" {
		t.Fatalf("expected operation with valid name from DeleteInstance, got: %v", deleteOp)
	}

	// 7. Get after Delete should return NotFound
	_, err = server.GetInstance(ctx, getReq)
	if status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound after delete, got: %v", err)
	}
}
