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

func (s *amlServer) populateBacktestResultDefaults(obj *pb.BacktestResult) {
	if obj.State == pb.BacktestResult_STATE_UNSPECIFIED {
		obj.State = pb.BacktestResult_ACTIVE
	}
	// TODO: proto3 zero-value gap for backtest_periods if omitted by user
}

func (s *amlServer) GetBacktestResult(ctx context.Context, req *pb.GetBacktestResultRequest) (*pb.BacktestResult, error) {
	name, err := s.parseBacktestResultName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.BacktestResult{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	s.populateBacktestResultDefaults(obj)

	return obj, nil
}

func (s *amlServer) CreateBacktestResult(ctx context.Context, req *pb.CreateBacktestResultRequest) (*longrunning.Operation, error) {
	reqId := req.GetBacktestResultId()
	if reqId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "backtest_result_id is required")
	}
	reqName := req.GetParent() + "/backtestResults/" + reqId
	name, err := s.parseBacktestResultName(reqName)
	if err != nil {
		return nil, err
	}

	parentName, err := s.parseInstanceName(req.GetParent())
	if err != nil {
		return nil, err
	}
	parentObj := &pb.Instance{}
	if err := s.storage.Get(ctx, parentName.String(), parentObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parentName.String())
		}
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetBacktestResult()).(*pb.BacktestResult)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateBacktestResultDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *amlServer) UpdateBacktestResult(ctx context.Context, req *pb.UpdateBacktestResultRequest) (*longrunning.Operation, error) {
	name, err := s.parseBacktestResultName(req.GetBacktestResult().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.BacktestResult{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj = proto.Clone(req.GetBacktestResult()).(*pb.BacktestResult)
	} else {
		if err := fields.UpdateByFieldMask(obj, req.GetBacktestResult(), paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update fields: %v", err)
		}
	}
	obj.Name = fqn
	obj.UpdateTime = timestamppb.New(now)
	s.populateBacktestResultDefaults(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *amlServer) DeleteBacktestResult(ctx context.Context, req *pb.DeleteBacktestResultRequest) (*longrunning.Operation, error) {
	name, err := s.parseBacktestResultName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	oldObj := &pb.BacktestResult{}
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

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *amlServer) ListBacktestResults(ctx context.Context, req *pb.ListBacktestResultsRequest) (*pb.ListBacktestResultsResponse, error) {
	parentName, err := s.parseInstanceName(req.GetParent())
	if err != nil {
		return nil, err
	}
	response := &pb.ListBacktestResultsResponse{}
	findPrefix := parentName.String() + "/backtestResults/"
	kind := (&pb.BacktestResult{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		res := proto.Clone(obj).(*pb.BacktestResult)
		s.populateBacktestResultDefaults(res)
		response.BacktestResults = append(response.BacktestResults, res)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
