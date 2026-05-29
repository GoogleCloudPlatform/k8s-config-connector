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

package mocknetworksecurity

import (
	"context"
	"time"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/networksecurity/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type InterceptServer struct {
	*MockService
	pb.UnimplementedInterceptServer
}

func (s *InterceptServer) CreateInterceptEndpointGroup(ctx context.Context, req *pb.CreateInterceptEndpointGroupRequest) (*longrunningpb.Operation, error) {
	name := req.Parent + "/interceptEndpointGroups/" + req.InterceptEndpointGroupId
	fqn := name

	obj := proto.Clone(req.InterceptEndpointGroup).(*pb.InterceptEndpointGroup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb.InterceptEndpointGroup_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name,
		Verb:       "create",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *InterceptServer) GetInterceptEndpointGroup(ctx context.Context, req *pb.GetInterceptEndpointGroupRequest) (*pb.InterceptEndpointGroup, error) {
	fqn := req.Name
	obj := &pb.InterceptEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *InterceptServer) UpdateInterceptEndpointGroup(ctx context.Context, req *pb.UpdateInterceptEndpointGroupRequest) (*longrunningpb.Operation, error) {
	fqn := req.GetInterceptEndpointGroup().GetName()
	obj := &pb.InterceptEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.GetInterceptEndpointGroup()).(*pb.InterceptEndpointGroup)
	updated.CreateTime = obj.CreateTime
	updated.UpdateTime = timestamppb.New(time.Now())
	updated.State = pb.InterceptEndpointGroup_ACTIVE

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, fqn, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(updated)
		return result, nil
	})
}

func (s *InterceptServer) DeleteInterceptEndpointGroup(ctx context.Context, req *pb.DeleteInterceptEndpointGroupRequest) (*longrunningpb.Operation, error) {
	fqn := req.Name
	obj := &pb.InterceptEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, fqn, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &longrunningpb.Operation{}, nil
	})
}

func (s *InterceptServer) CreateInterceptDeploymentGroup(ctx context.Context, req *pb.CreateInterceptDeploymentGroupRequest) (*longrunningpb.Operation, error) {
	name := req.Parent + "/interceptDeploymentGroups/" + req.InterceptDeploymentGroupId
	fqn := name

	obj := proto.Clone(req.InterceptDeploymentGroup).(*pb.InterceptDeploymentGroup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb.InterceptDeploymentGroup_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name,
		Verb:       "create",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *InterceptServer) GetInterceptDeploymentGroup(ctx context.Context, req *pb.GetInterceptDeploymentGroupRequest) (*pb.InterceptDeploymentGroup, error) {
	fqn := req.Name
	obj := &pb.InterceptDeploymentGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}
