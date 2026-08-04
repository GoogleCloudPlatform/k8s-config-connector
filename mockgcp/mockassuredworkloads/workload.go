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

// +tool:mockgcp-support
// proto.service: google.cloud.assuredworkloads.v1.AssuredWorkloadsService
// proto.message: google.cloud.assuredworkloads.v1.Workload

package mockassuredworkloads

import (
	"context"
	"fmt"
	"time"

	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *AssuredWorkloadsV1) GetWorkload(ctx context.Context, req *pb.GetWorkloadRequest) (*pb.Workload, error) {
	fqn := req.GetName()

	obj := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AssuredWorkloadsV1) CreateWorkload(ctx context.Context, req *pb.CreateWorkloadRequest) (*longrunning.Operation, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}

	workloadId := "123456" // predictable server-generated ID
	fqn := fmt.Sprintf("%s/workloads/%s", parent, workloadId)

	now := time.Now()

	obj := proto.Clone(req.GetWorkload()).(*pb.Workload)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.Etag = "mockgcp-etag-1"

	// Set output only fields
	obj.KajEnrollmentState = pb.Workload_KAJ_ENROLLMENT_STATE_COMPLETE
	obj.SaaEnrollmentResponse = &pb.Workload_SaaEnrollmentResponse{
		SetupStatus: pb.Workload_SaaEnrollmentResponse_STATUS_COMPLETE.Enum(),
	}

	if len(obj.ResourceSettings) > 0 {
		var resources []*pb.Workload_ResourceInfo
		for _, settings := range obj.ResourceSettings {
			resources = append(resources, &pb.Workload_ResourceInfo{
				ResourceId:   123456789,
				ResourceType: settings.ResourceType,
			})
		}
		obj.Resources = resources
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := parent
	lroMetadata := &pb.CreateWorkloadOperationMetadata{
		CreateTime:       timestamppb.New(now),
		DisplayName:      obj.DisplayName,
		Parent:           parent,
		ComplianceRegime: obj.ComplianceRegime,
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *AssuredWorkloadsV1) UpdateWorkload(ctx context.Context, req *pb.UpdateWorkloadRequest) (*pb.Workload, error) {
	reqWorkload := req.GetWorkload()
	if reqWorkload == nil {
		return nil, status.Errorf(codes.InvalidArgument, "workload is required")
	}

	fqn := reqWorkload.GetName()
	obj := &pb.Workload{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = reqWorkload.GetDisplayName()
		case "labels":
			obj.Labels = reqWorkload.GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.Etag = fmt.Sprintf("mockgcp-etag-%d", time.Now().UnixNano())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *AssuredWorkloadsV1) DeleteWorkload(ctx context.Context, req *pb.DeleteWorkloadRequest) (*emptypb.Empty, error) {
	fqn := req.GetName()

	deletedObj := &pb.Workload{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AssuredWorkloadsV1) ListWorkloads(ctx context.Context, req *pb.ListWorkloadsRequest) (*pb.ListWorkloadsResponse, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}

	var workloads []*pb.Workload
	if err := s.storage.List(ctx, (&pb.Workload{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: parent + "/workloads/",
	}, func(obj proto.Message) error {
		workload := obj.(*pb.Workload)
		workloads = append(workloads, workload)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListWorkloadsResponse{
		Workloads: workloads,
	}, nil
}
