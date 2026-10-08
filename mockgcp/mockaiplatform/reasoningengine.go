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
// proto.service: google.cloud.aiplatform.v1.ReasoningEngineService
// proto.message: google.cloud.aiplatform.v1.ReasoningEngine

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

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

	fqn := name.StorageKey()

	obj := &pb.ReasoningEngine{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.CloneOf(obj)
	result.Name = name.String()
	return result, nil
}

func (s *reasoningEngineService) CreateReasoningEngine(ctx context.Context, req *pb.CreateReasoningEngineRequest) (*longrunning.Operation, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	reqName := req.Parent + "/reasoningEngines/" + id
	name, err := s.parseReasoningEngineName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.CloneOf(req.ReasoningEngine)
	obj.Name = name.String()

	if obj.Spec == nil {
		obj.Spec = &pb.ReasoningEngineSpec{}
	}

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateReasoningEngineOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.StringWithProjectNumber()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		result := proto.CloneOf(obj)
		result.CreateTime = nil
		result.UpdateTime = nil
		result.Name = name.StringWithProjectNumber()
		return result, nil
	})
}

func (s *reasoningEngineService) UpdateReasoningEngine(ctx context.Context, req *pb.UpdateReasoningEngineRequest) (*longrunning.Operation, error) {
	name, err := s.parseReasoningEngineName(req.GetReasoningEngine().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := &pb.ReasoningEngine{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "displayName", "display_name":
			obj.DisplayName = req.GetReasoningEngine().GetDisplayName()
		case "description":
			obj.Description = req.GetReasoningEngine().GetDescription()
		case "labels":
			obj.Labels = req.GetReasoningEngine().GetLabels()
		case "spec":
			obj.Spec = req.GetReasoningEngine().GetSpec()
		case "encryptionSpec", "encryption_spec":
			obj.EncryptionSpec = req.GetReasoningEngine().GetEncryptionSpec()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if obj.Spec == nil {
		obj.Spec = &pb.ReasoningEngineSpec{}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.UpdateReasoningEngineOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.StringWithProjectNumber()
	result := proto.CloneOf(obj)
	result.Name = name.StringWithProjectNumber()
	return s.operations.DoneLRO(ctx, opPrefix, op, result)
}

func (s *reasoningEngineService) DeleteReasoningEngine(ctx context.Context, req *pb.DeleteReasoningEngineRequest) (*longrunning.Operation, error) {
	name, err := s.parseReasoningEngineName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
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

func (s *reasoningEngineService) ListReasoningEngines(ctx context.Context, req *pb.ListReasoningEnginesRequest) (*pb.ListReasoningEnginesResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}
	projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
	if err != nil {
		return nil, err
	}
	project, err := s.Projects.GetProject(projectName)
	if err != nil {
		return nil, err
	}
	location := tokens[3]

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/reasoningEngines/", project.Number, location)

	var reasoningEngines []*pb.ReasoningEngine
	reasoningEngineKind := (&pb.ReasoningEngine{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, reasoningEngineKind, storage.ListOptions{}, func(obj proto.Message) error {
		re := obj.(*pb.ReasoningEngine)
		if strings.HasPrefix(re.GetName(), findPrefix) {
			reasoningEngines = append(reasoningEngines, re)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListReasoningEnginesResponse{
		ReasoningEngines: reasoningEngines,
	}, nil
}

type ReasoningEngineName struct {
	Project           *projects.ProjectData
	Location          string
	ReasoningEngineID string
}

func (n *ReasoningEngineName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/reasoningEngines/%s", n.Project.ID, n.Location, n.ReasoningEngineID)
}

func (n *ReasoningEngineName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/reasoningEngines/%s", n.Project.Number, n.Location, n.ReasoningEngineID)
}

func (n *ReasoningEngineName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/reasoningEngines/%s", n.Project.Number, n.Location, n.ReasoningEngineID)
}

// parseReasoningEngineName parses a string into a ReasoningEngineName.
// The expected form is projects/<projectID_or_number>/locations/<location>/reasoningEngines/<reasoningEngineID>
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

		reasoningEngineID := tokens[5]
		if !isNumeric(reasoningEngineID) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid reasoning engine id: %s", reasoningEngineID)
		}

		return &ReasoningEngineName{
			Project:           project,
			Location:          tokens[3],
			ReasoningEngineID: reasoningEngineID,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
