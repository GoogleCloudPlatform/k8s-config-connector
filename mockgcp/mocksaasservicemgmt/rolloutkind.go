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
)

type SaasRolloutsServer struct {
	*MockService
	pb.UnimplementedSaasRolloutsServer
}

type rolloutKindName struct {
	Project     *projects.ProjectData
	Location    string
	RolloutKind string
}

func (n *rolloutKindName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/rolloutKinds/" + n.RolloutKind
}

func (s *MockService) parseRolloutKindName(name string) (*rolloutKindName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "rolloutKinds" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &rolloutKindName{
			Project:     project,
			Location:    tokens[3],
			RolloutKind: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SaasRolloutsServer) GetRolloutKind(ctx context.Context, req *pb.GetRolloutKindRequest) (*pb.RolloutKind, error) {
	name, err := s.parseRolloutKindName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RolloutKind{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SaasRolloutsServer) CreateRolloutKind(ctx context.Context, req *pb.CreateRolloutKindRequest) (*pb.RolloutKind, error) {
	reqName := req.Parent + "/rolloutKinds/" + req.RolloutKindId
	name, err := s.parseRolloutKindName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.RolloutKind).(*pb.RolloutKind)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = "mock-etag"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.RolloutKind)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	return response, nil
}

func (s *SaasRolloutsServer) UpdateRolloutKind(ctx context.Context, req *pb.UpdateRolloutKindRequest) (*pb.RolloutKind, error) {
	reqName := req.GetRolloutKind().GetName()
	name, err := s.parseRolloutKindName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RolloutKind{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.RolloutKind, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update RolloutKind fields: %v", err)
		}
	} else {
		desired := req.GetRolloutKind()
		obj.UnitKind = desired.UnitKind
		obj.RolloutOrchestrationStrategy = desired.RolloutOrchestrationStrategy
		obj.UnitFilter = desired.UnitFilter
		obj.UpdateUnitKindStrategy = desired.UpdateUnitKindStrategy
		obj.ErrorBudget = desired.ErrorBudget
		obj.Labels = desired.Labels
		obj.Annotations = desired.Annotations
	}
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = "mock-etag-updated"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.RolloutKind)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	return response, nil
}

func (s *SaasRolloutsServer) DeleteRolloutKind(ctx context.Context, req *pb.DeleteRolloutKindRequest) (*emptypb.Empty, error) {
	name, err := s.parseRolloutKindName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RolloutKind{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
