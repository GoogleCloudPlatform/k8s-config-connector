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

package mockvisionai

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type applicationServer struct {
	*MockService
	pb.UnimplementedAppPlatformServer
}

func (s *applicationServer) GetApplication(ctx context.Context, req *pb.GetApplicationRequest) (*pb.Application, error) {
	name, err := s.parseApplicationName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Application{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return proto.Clone(obj).(*pb.Application), nil
}

func (s *applicationServer) ListApplications(ctx context.Context, req *pb.ListApplicationsRequest) (*pb.ListApplicationsResponse, error) {
	response := &pb.ListApplicationsResponse{}
	findPrefix := req.GetParent() + "/applications/"
	applicationKind := (&pb.Application{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, applicationKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		app := obj.(*pb.Application)
		response.Applications = append(response.Applications, app)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *applicationServer) CreateApplication(ctx context.Context, req *pb.CreateApplicationRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/applications/%s", req.GetParent(), req.GetApplicationId())
	name, err := s.parseApplicationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetApplication()).(*pb.Application)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Application_CREATED
	if obj.ApplicationConfigs != nil {
		obj.ApplicationConfigs = &pb.ApplicationConfigs{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
		}
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		retObj := proto.Clone(obj).(*pb.Application)
		return retObj, nil
	})
}

func (s *applicationServer) UpdateApplication(ctx context.Context, req *pb.UpdateApplicationRequest) (*longrunning.Operation, error) {
	name, err := s.parseApplicationName(req.GetApplication().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Application{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Application)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = fields.ComputeImpliedFieldMask(ctx, req.GetApplication(), "name", "create_time", "update_time")
	}

	if err := fields.UpdateByFieldMask(updated, req.GetApplication(), paths); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating fields: %v", err)
	}

	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)
	if updated.ApplicationConfigs != nil {
		updated.ApplicationConfigs = &pb.ApplicationConfigs{}
	}
	if updated.RuntimeInfo == nil {
		updated.RuntimeInfo = &pb.Application_ApplicationRuntimeInfo{}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		retObj := proto.Clone(updated).(*pb.Application)
		if retObj.RuntimeInfo == nil {
			retObj.RuntimeInfo = &pb.Application_ApplicationRuntimeInfo{}
		}
		return retObj, nil
	})
}

func (s *applicationServer) DeleteApplication(ctx context.Context, req *pb.DeleteApplicationRequest) (*longrunning.Operation, error) {
	name, err := s.parseApplicationName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Application{}
	if err := s.storage.Get(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	deleted := &pb.Application{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
