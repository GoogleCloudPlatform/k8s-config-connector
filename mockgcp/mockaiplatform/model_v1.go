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

package mockaiplatform

import (
	"context"

	v1pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/protobuf/proto"
)

type modelServiceV1 struct {
	*MockService
	v1pb.UnimplementedModelServiceServer
	v1beta1 *modelService
}

func convertProto[T proto.Message](src proto.Message, dest T) (T, error) {
	b, err := proto.Marshal(src)
	if err != nil {
		return dest, err
	}
	if err := proto.Unmarshal(b, dest); err != nil {
		return dest, err
	}
	return dest, nil
}

func (s *modelServiceV1) UploadModel(ctx context.Context, req *v1pb.UploadModelRequest) (*longrunningpb.Operation, error) {
	betaReq, err := convertProto(req, &pb.UploadModelRequest{})
	if err != nil {
		return nil, err
	}
	return s.v1beta1.UploadModel(ctx, betaReq)
}

func (s *modelServiceV1) UpdateModel(ctx context.Context, req *v1pb.UpdateModelRequest) (*v1pb.Model, error) {
	betaReq, err := convertProto(req, &pb.UpdateModelRequest{})
	if err != nil {
		return nil, err
	}
	betaResp, err := s.v1beta1.UpdateModel(ctx, betaReq)
	if err != nil {
		return nil, err
	}
	return convertProto(betaResp, &v1pb.Model{})
}

func (s *modelServiceV1) GetModel(ctx context.Context, req *v1pb.GetModelRequest) (*v1pb.Model, error) {
	betaReq, err := convertProto(req, &pb.GetModelRequest{})
	if err != nil {
		return nil, err
	}
	betaResp, err := s.v1beta1.GetModel(ctx, betaReq)
	if err != nil {
		return nil, err
	}
	return convertProto(betaResp, &v1pb.Model{})
}

func (s *modelServiceV1) ListModels(ctx context.Context, req *v1pb.ListModelsRequest) (*v1pb.ListModelsResponse, error) {
	betaReq, err := convertProto(req, &pb.ListModelsRequest{})
	if err != nil {
		return nil, err
	}
	betaResp, err := s.v1beta1.ListModels(ctx, betaReq)
	if err != nil {
		return nil, err
	}
	return convertProto(betaResp, &v1pb.ListModelsResponse{})
}

func (s *modelServiceV1) DeleteModel(ctx context.Context, req *v1pb.DeleteModelRequest) (*longrunningpb.Operation, error) {
	betaReq, err := convertProto(req, &pb.DeleteModelRequest{})
	if err != nil {
		return nil, err
	}
	return s.v1beta1.DeleteModel(ctx, betaReq)
}
