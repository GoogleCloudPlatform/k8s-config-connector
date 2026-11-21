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

package mockgeminidataanalytics

import (
	"context"
	"fmt"
	"time"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/geminidataanalytics/apiv1/geminidataanalyticspb"
)

type GeminiDataAnalyticsV1 struct {
	*MockService
	pb.UnimplementedDataAgentServiceServer
}

func (s *GeminiDataAnalyticsV1) GetDataAgent(ctx context.Context, req *pb.GetDataAgentRequest) (*pb.DataAgent, error) {
	name, err := s.parseDataAgentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataAgent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *GeminiDataAnalyticsV1) CreateDataAgent(ctx context.Context, req *pb.CreateDataAgentRequest) (*longrunningpb.Operation, error) {
	dataAgentID := req.DataAgentId
	if dataAgentID == "" {
		dataAgentID = fmt.Sprintf("agent-%d", time.Now().UnixNano())
	}
	reqName := req.Parent + "/dataAgents/" + dataAgentID
	name, err := s.parseDataAgentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.DataAgent)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *GeminiDataAnalyticsV1) UpdateDataAgent(ctx context.Context, req *pb.UpdateDataAgentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDataAgentName(req.GetDataAgent().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataAgent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, req.GetDataAgent(), paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	now := timestamppb.Now()
	obj.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *GeminiDataAnalyticsV1) DeleteDataAgent(ctx context.Context, req *pb.DeleteDataAgentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDataAgentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataAgent{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := timestamppb.Now()
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
