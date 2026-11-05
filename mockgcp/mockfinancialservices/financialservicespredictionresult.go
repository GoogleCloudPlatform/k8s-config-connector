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

func (s *amlServer) GetPredictionResult(ctx context.Context, req *pb.GetPredictionResultRequest) (*pb.PredictionResult, error) {
	name, err := s.parsePredictionResultName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.PredictionResult{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *amlServer) CreatePredictionResult(ctx context.Context, req *pb.CreatePredictionResultRequest) (*longrunning.Operation, error) {
	if req.GetPredictionResult() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "prediction_result is required")
	}

	reqName := req.Parent + "/predictionResults/" + req.PredictionResultId
	name, err := s.parsePredictionResultName(reqName)
	if err != nil {
		return nil, err
	}

	// Verify parent Instance exists
	parentObj := &pb.Instance{}
	if err := s.storage.Get(ctx, name.Instance.String(), parentObj); err != nil {
		if apierrors.IsNotFound(err) || status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name.Instance.String())
		}
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.PredictionResult).(*pb.PredictionResult)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.PredictionResult_ACTIVE

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

func (s *amlServer) UpdatePredictionResult(ctx context.Context, req *pb.UpdatePredictionResultRequest) (*longrunning.Operation, error) {
	if req.GetPredictionResult() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "prediction_result is required")
	}

	name, err := s.parsePredictionResultName(req.GetPredictionResult().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.PredictionResult{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		createTime := obj.CreateTime
		state := obj.State
		obj = proto.Clone(req.PredictionResult).(*pb.PredictionResult)
		obj.Name = fqn
		obj.CreateTime = createTime
		obj.State = state
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.PredictionResult.Labels
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not updatable", path)
			}
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

func (s *amlServer) DeletePredictionResult(ctx context.Context, req *pb.DeletePredictionResultRequest) (*longrunning.Operation, error) {
	name, err := s.parsePredictionResultName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	oldObj := &pb.PredictionResult{}
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
