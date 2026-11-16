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
)

func (s *amlServer) GetEngineConfig(ctx context.Context, req *pb.GetEngineConfigRequest) (*pb.EngineConfig, error) {
	name, err := s.parseEngineConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.EngineConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *amlServer) CreateEngineConfig(ctx context.Context, req *pb.CreateEngineConfigRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/engineConfigs/" + req.EngineConfigId
	name, err := s.parseEngineConfigName(reqName)
	if err != nil {
		return nil, err
	}

	parentInstance := &pb.Instance{}
	if err := s.storage.Get(ctx, name.Instance.String(), parentInstance); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name.Instance.String())
		}
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.EngineConfig).(*pb.EngineConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.EngineConfig_ACTIVE

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

func (s *amlServer) UpdateEngineConfig(ctx context.Context, req *pb.UpdateEngineConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseEngineConfigName(req.EngineConfig.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.EngineConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	proto.Merge(obj, req.EngineConfig)
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

func (s *amlServer) DeleteEngineConfig(ctx context.Context, req *pb.DeleteEngineConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseEngineConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.EngineConfig{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
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
