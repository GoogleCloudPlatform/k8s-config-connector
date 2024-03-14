// Copyright 2024 Google LLC
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

package mockworkstations

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/workstations/v1beta"
	"github.com/google/uuid"
)

func (s *workstationsServer) GetWorkstationCluster(ctx context.Context, req *pb.GetWorkstationClusterRequest) (*pb.WorkstationCluster, error) {
	name, err := s.parseWorkstationClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkstationCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *workstationsServer) CreateWorkstationCluster(ctx context.Context, req *pb.CreateWorkstationClusterRequest) (*longrunning.Operation, error) {
	reqName := req.GetParent() + "/workstationClusters/" + req.GetWorkstationClusterId()
	name, err := s.parseWorkstationClusterName(reqName)
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validate only is not yet implemented in mockgcp")
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.WorkstationCluster).(*pb.WorkstationCluster)
	obj.Name = fqn

	obj.ControlPlaneIp = "10.2.0.2"
	obj.PrivateClusterConfig = &pb.WorkstationCluster_PrivateClusterConfig{}
	obj.Uid = uuid.New().String()

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// // Many fields are not populated in the LRO result
		// result := proto.Clone(obj).(*pb.WorkstationCluster)
		// result.CreateTime = nil
		// result.UpdateTime = nil
		// result.Etag = ""

		// return result, nil
		return obj, nil
	})
}

func (s *workstationsServer) UpdateWorkstationCluster(ctx context.Context, req *pb.UpdateWorkstationClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkstationClusterName(req.GetWorkstationCluster().GetName())
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validateOnly is not yet implemented in mockgcp")
	}
	if req.GetAllowMissing() {
		return nil, status.Errorf(codes.Unimplemented, "allowMissing is not yet implemented in mockgcp")
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.WorkstationCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetWorkstationCluster().GetDisplayName()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// // Many fields are not populated in the LRO result
		// result := proto.Clone(obj).(*pb.WorkstationCluster)
		// result.CreateTime = nil
		// result.UpdateTime = nil
		// result.Etag = ""

		// return result, nil
		return obj, nil
	})
}

func (s *workstationsServer) DeleteWorkstationCluster(ctx context.Context, req *pb.DeleteWorkstationClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkstationClusterName(req.GetName())
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validateOnly is not yet implemented in mockgcp")
	}
	if req.GetForce() {
		return nil, status.Errorf(codes.Unimplemented, "force is not yet implemented in mockgcp")
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.WorkstationCluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// // Many fields are not populated in the LRO result
		// result := proto.Clone(obj).(*pb.WorkstationCluster)
		// result.CreateTime = nil
		// result.UpdateTime = nil
		// result.Etag = ""

		// return result, nil
		return deleted, nil
	})
}

type WorkstationClusterName struct {
	Project              *projects.ProjectData
	Location             string
	WorkstationClusterID string
}

func (n *WorkstationClusterName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", n.Project.ID, n.Location, n.WorkstationClusterID)
}

// parseWorkstationClusterName parses a string into a WorkstationClusterName.
// The expected form is projects/*/locations/*/workstationClusters/*
func (s *MockService) parseWorkstationClusterName(name string) (*WorkstationClusterName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workstationClusters" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.GetProjects().GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &WorkstationClusterName{
			Project:              project,
			Location:             tokens[3],
			WorkstationClusterID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
