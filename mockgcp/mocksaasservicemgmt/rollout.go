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

package mocksaasservicemgmt

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type rolloutName struct {
	Project  *projects.ProjectData
	Location string
	Rollout  string
}

func (n *rolloutName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/rollouts/" + n.Rollout
}

func (s *MockService) parseRolloutName(name string) (*rolloutName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "rollouts" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &rolloutName{
			Project:  project,
			Location: tokens[3],
			Rollout:  tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SaasRolloutsServer) GetRollout(ctx context.Context, req *pb.GetRolloutRequest) (*pb.Rollout, error) {
	name, err := s.parseRolloutName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Rollout{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SaasRolloutsServer) CreateRollout(ctx context.Context, req *pb.CreateRolloutRequest) (*pb.Rollout, error) {
	reqName := req.Parent + "/rollouts/" + req.RolloutId
	name, err := s.parseRolloutName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Rollout).(*pb.Rollout)
	obj.Name = fqn

	if obj.Control == nil {
		obj.Control = &pb.RolloutControl{
			Action: pb.RolloutAction_ROLLOUT_ACTION_RUN,
		}
	} else if obj.Control.Action == pb.RolloutAction_ROLLOUT_ACTION_UNSPECIFIED {
		obj.Control.Action = pb.RolloutAction_ROLLOUT_ACTION_RUN
	}

	if obj.UnitFilter != "" {
		obj.EffectiveUnitFilter = obj.UnitFilter
	}

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = "mock-etag"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
		}
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Rollout)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	return response, nil
}

func (s *SaasRolloutsServer) UpdateRollout(ctx context.Context, req *pb.UpdateRolloutRequest) (*pb.Rollout, error) {
	reqName := req.GetRollout().GetName()
	name, err := s.parseRolloutName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Rollout{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.Rollout, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update Rollout fields: %v", err)
		}
	} else {
		desired := req.GetRollout()
		obj.RolloutOrchestrationStrategy = desired.RolloutOrchestrationStrategy
		obj.UnitFilter = desired.UnitFilter
		obj.Control = desired.Control
		obj.Labels = desired.Labels
		obj.Annotations = desired.Annotations
	}

	if obj.UnitFilter != "" {
		obj.EffectiveUnitFilter = obj.UnitFilter
	}

	now := timestamppb.Now()
	obj.UpdateTime = now
	obj.Etag = "mock-etag-updated"

	if obj.Control != nil && obj.Control.Action == pb.RolloutAction_ROLLOUT_ACTION_RUN {
		obj.RootRollout = fqn
		if obj.StartTime == nil {
			obj.StartTime = now
		}
		obj.State = pb.Rollout_ROLLOUT_STATE_RUNNING
		obj.StateMessage = "Root Rollout set to RUNNING, child rollouts are being created."
		obj.StateTransitionTime = now
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Rollout)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	response.RootRollout = ""
	response.StartTime = nil
	response.State = pb.Rollout_ROLLOUT_STATE_UNSPECIFIED
	response.StateMessage = ""
	response.StateTransitionTime = nil
	return response, nil
}

func (s *SaasRolloutsServer) DeleteRollout(ctx context.Context, req *pb.DeleteRolloutRequest) (*emptypb.Empty, error) {
	name, err := s.parseRolloutName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Rollout{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SaasRolloutsServer) ListRollouts(ctx context.Context, req *pb.ListRolloutsRequest) (*pb.ListRolloutsResponse, error) {
	response := &pb.ListRolloutsResponse{}
	findKind := (&pb.Rollout{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: req.Parent}, func(obj proto.Message) error {
		item := obj.(*pb.Rollout)
		response.Rollouts = append(response.Rollouts, item)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
