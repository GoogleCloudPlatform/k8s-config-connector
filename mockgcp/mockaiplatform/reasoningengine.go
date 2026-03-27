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
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
)

type reasoningEngineService struct {
	*MockService
	pb.UnimplementedReasoningEngineServiceServer
}

func (s *reasoningEngineService) GetReasoningEngine(ctx context.Context, req *pb.GetReasoningEngineRequest) (*pb.ReasoningEngine, error) {
	name, err := s.parseReasoningEngineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ReasoningEngine{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The ReasoningEngine does not exist.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *reasoningEngineService) CreateReasoningEngine(ctx context.Context, req *pb.CreateReasoningEngineRequest) (*longrunning.Operation, error) {
	// Generate a random ID if not provided (though the API doesn't seem to support providing it)
	reasoningEngineId := fmt.Sprintf("%d", time.Now().UnixNano())

	reqName := req.Parent + "/reasoningEngines/" + reasoningEngineId
	name, err := s.parseReasoningEngineName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.ReasoningEngine).(*pb.ReasoningEngine)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateReasoningEngineOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *reasoningEngineService) UpdateReasoningEngine(ctx context.Context, req *pb.UpdateReasoningEngineRequest) (*longrunning.Operation, error) {
	name, err := s.parseReasoningEngineName(req.GetReasoningEngine().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.ReasoningEngine{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "display_name":
			obj.DisplayName = req.GetReasoningEngine().GetDisplayName()
		case "description":
			obj.Description = req.GetReasoningEngine().GetDescription()
		case "spec":
			obj.Spec = req.GetReasoningEngine().GetSpec()
		case "etag":
			obj.Etag = req.GetReasoningEngine().GetEtag()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateReasoningEngineOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *reasoningEngineService) DeleteReasoningEngine(ctx context.Context, req *pb.DeleteReasoningEngineRequest) (*longrunning.Operation, error) {
	name, err := s.parseReasoningEngineName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.ReasoningEngine{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type ReasoningEngineName struct {
	Project           *projects.ProjectData
	Location          string
	ReasoningEngineID string
}

func (n *ReasoningEngineName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", n.Project.ID, n.Location, n.ReasoningEngineID)
}

// parseReasoningEngineName parses a string into a ReasoningEngineName.
func (s *MockService) parseReasoningEngineName(name string) (*ReasoningEngineName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "reasoningEngines" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &ReasoningEngineName{
			Project:           project,
			Location:          tokens[3],
			ReasoningEngineID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
