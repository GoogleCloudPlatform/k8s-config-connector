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

package mockassuredworkloads

import (
	"context"
	"fmt"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/assuredworkloads/v1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AssuredWorkloadsV1 struct {
	*MockService
	pb.UnimplementedAssuredWorkloadsServiceServer
}

func (s *AssuredWorkloadsV1) GetWorkload(ctx context.Context, req *pb.GetWorkloadRequest) (*pb.Workload, error) {
	name, err := s.parseWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AssuredWorkloadsV1) CreateWorkload(ctx context.Context, req *pb.CreateWorkloadRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/workloads/" + "dummy" // WorkloadID is not in the request, it is generated or in req.Workload?
	// Actually, looking at the proto, CreateWorkloadRequest has a Workload, but where is the ID?
	// Some services have a Separate ID field.

	// Let's check CreateWorkloadRequest in the proto.
	/*
		message CreateWorkloadRequest {
		  string parent = 1;
		  Workload workload = 2;
		  string external_id = 3;
		}
	*/
	// It has external_id.

	workloadID := req.ExternalId
	if workloadID == "" {
		workloadID = fmt.Sprintf("workload-%d", time.Now().UnixNano())
	}

	reqName = req.Parent + "/workloads/" + workloadID
	name, err := s.parseWorkloadName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Workload).(*pb.Workload)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	// Assured Workloads often create a folder.
	// We might want to simulate that.
	obj.Resources = []*pb.Workload_ResourceInfo{
		{
			ResourceId:   123456789,
			ResourceType: pb.Workload_ResourceInfo_CONSUMER_FOLDER,
		},
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, fqn, nil, obj)
}

func (s *AssuredWorkloadsV1) UpdateWorkload(ctx context.Context, req *pb.UpdateWorkloadRequest) (*pb.Workload, error) {
	obj := req.GetWorkload()
	if obj == nil {
		return nil, status.Error(codes.InvalidArgument, "workload is required")
	}

	name, err := s.parseWorkloadName(obj.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	if updateMask == nil {
		updateMask = &fieldmaskpb.FieldMask{Paths: []string{"display_name", "labels"}}
	}

	for _, path := range updateMask.Paths {
		switch path {
		case "display_name", "displayName":
			existing.DisplayName = obj.DisplayName
		case "labels":
			existing.Labels = obj.Labels
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not updatable", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *AssuredWorkloadsV1) DeleteWorkload(ctx context.Context, req *pb.DeleteWorkloadRequest) (*emptypb.Empty, error) {
	name, err := s.parseWorkloadName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Workload{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AssuredWorkloadsV1) ListWorkloads(ctx context.Context, req *pb.ListWorkloadsRequest) (*pb.ListWorkloadsResponse, error) {
	// Simple implementation
	return nil, status.Error(codes.Unimplemented, "ListWorkloads")
}
