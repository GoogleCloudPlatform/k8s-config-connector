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
// proto.service: google.cloud.lustre.v1.Lustre
// proto.message: google.cloud.lustre.v1.Instance

package mocklustre

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/lustre/apiv1/lustrepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type lustreServer struct {
	*MockService
	pb.UnimplementedLustreServer
}

func (s *lustreServer) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *lustreServer) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	instanceID := req.GetInstanceId()
	if instanceID == "" {
		tokens := strings.Split(req.GetInstance().GetName(), "/")
		instanceID = tokens[len(tokens)-1]
	}
	if instanceID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "instanceId is required")
	}

	fqn := fmt.Sprintf("%s/instances/%s", parent.String(), instanceID)

	now := time.Now()

	obj := proto.Clone(req.GetInstance()).(*pb.Instance)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Instance_ACTIVE

	// Mount point IP allocation
	ip := "192.168.84.3"
	if strings.Contains(instanceID, "maximal") {
		ip = "192.168.122.2"
	}
	obj.MountPoint = fmt.Sprintf("%s@tcp:/%s", ip, obj.GetFilesystem())

	if obj.GetPerUnitStorageThroughput() == 0 {
		obj.PerUnitStorageThroughput = 1000
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := parent.String()
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *lustreServer) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.GetInstance().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", name)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetInstance().GetDescription()
		case "labels":
			obj.Labels = req.GetInstance().GetLabels()
		case "capacity_gib", "capacityGib":
			obj.CapacityGib = req.GetInstance().GetCapacityGib()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *lustreServer) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", name)
		}
		return nil, err
	}

	now := time.Now()
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *lustreServer) ListInstances(ctx context.Context, req *pb.ListInstancesRequest) (*pb.ListInstancesResponse, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListInstancesResponse{}
	prefix := parent.String() + "/instances/"

	kind := (&pb.Instance{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		instance := obj.(*pb.Instance)
		response.Instances = append(response.Instances, instance)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
