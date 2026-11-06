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

package mockfinancialservices

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
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "cloud.google.com/go/financialservices/apiv1/financialservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *amlServer) GetModel(ctx context.Context, req *pb.GetModelRequest) (*pb.Model, error) {
	name, err := s.parseModelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Model{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *amlServer) CreateModel(ctx context.Context, req *pb.CreateModelRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/models/" + req.ModelId
	name, err := s.parseModelName(reqName)
	if err != nil {
		return nil, err
	}

	instanceObj := &pb.Instance{}
	if err := s.storage.Get(ctx, name.Instance.String(), instanceObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name.Instance.String())
		}
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Model).(*pb.Model)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Model_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Instance.Project.ID, name.Instance.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *amlServer) UpdateModel(ctx context.Context, req *pb.UpdateModelRequest) (*longrunning.Operation, error) {
	name, err := s.parseModelName(req.Model.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Model{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj = proto.Clone(req.Model).(*pb.Model)
	} else {
		if err := fields.UpdateByFieldMask(obj, req.Model, paths); err != nil {
			return nil, err
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Instance.Project.ID, name.Instance.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *amlServer) DeleteModel(ctx context.Context, req *pb.DeleteModelRequest) (*longrunning.Operation, error) {
	name, err := s.parseModelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	oldObj := &pb.Model{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Instance.Project.ID, name.Instance.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *amlServer) ListModels(ctx context.Context, req *pb.ListModelsRequest) (*pb.ListModelsResponse, error) {
	parent, err := s.parseInstanceName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListModelsResponse{}
	findPrefix := parent.String() + "/models/"
	modelKind := (&pb.Model{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, modelKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		model := obj.(*pb.Model)
		response.Models = append(response.Models, model)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
