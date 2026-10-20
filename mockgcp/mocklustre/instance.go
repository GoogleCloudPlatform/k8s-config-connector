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
	"reflect"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/lustre/apiv1/lustrepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
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
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *lustreServer) ListInstances(ctx context.Context, req *pb.ListInstancesRequest) (*pb.ListInstancesResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	location := tokens[3]

	res := &pb.ListInstancesResponse{}
	var prefix string
	if location == "-" {
		prefix = fmt.Sprintf("projects/%s/locations/", project.ID)
	} else {
		prefix = fmt.Sprintf("projects/%s/locations/%s/instances", project.ID, location)
	}

	if err := s.storage.List(ctx, (&pb.Instance{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		instance := obj.(*pb.Instance)
		res.Instances = append(res.Instances, instance)
		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *lustreServer) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	instanceID := req.GetInstanceId()
	if instanceID == "" && req.GetInstance() != nil && req.GetInstance().GetName() != "" {
		parts := strings.Split(req.GetInstance().GetName(), "/")
		instanceID = parts[len(parts)-1]
	}
	if instanceID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "instance_id must be provided")
	}

	reqName := fmt.Sprintf("%s/instances/%s", req.GetParent(), instanceID)
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Resource %q already exists", fqn)
	}

	inst := req.GetInstance()
	if inst == nil {
		return nil, status.Errorf(codes.InvalidArgument, "instance is required")
	}

	if inst.GetFilesystem() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "filesystem must be provided")
	}

	now := time.Now()

	obj := proto.Clone(inst).(*pb.Instance)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Instance_CREATING

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		return mutateObject(ctx, s.storage, fqn, func(obj *pb.Instance) error {
			obj.State = pb.Instance_ACTIVE
			if obj.MountPoint == "" {
				obj.MountPoint = fmt.Sprintf("10.0.0.2@tcp:/%s", obj.Filesystem)
			}
			return nil
		})
	})
}

func (s *lustreServer) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunning.Operation, error) {
	inst := req.GetInstance()
	if inst == nil {
		return nil, status.Errorf(codes.InvalidArgument, "instance is required")
	}

	name, err := s.parseInstanceName(inst.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, inst, paths); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating by field mask: %v", err)
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updatedObj := proto.Clone(obj).(*pb.Instance)
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return updatedObj, nil
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
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type instanceName struct {
	Project  *projects.ProjectData
	Location string
	Instance string
}

func (n *instanceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", n.Project.ID, n.Location, n.Instance)
}

func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &instanceName{
			Project:  project,
			Location: tokens[3],
			Instance: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func mutateObject[T proto.Message](ctx context.Context, storage storage.Storage, fqn string, mutator func(obj T) error) (T, error) {
	var nilT T

	typeT := reflect.TypeOf(nilT)
	obj := reflect.New(typeT.Elem()).Interface().(T)
	if err := storage.Get(ctx, fqn, obj); err != nil {
		return nilT, err
	}

	if err := mutator(obj); err != nil {
		return nilT, err
	}

	if err := storage.Update(ctx, fqn, obj); err != nil {
		return nilT, err
	}

	return obj, nil
}
