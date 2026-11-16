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

type unitKindName struct {
	Project  *projects.ProjectData
	Location string
	UnitKind string
}

func (n *unitKindName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/unitKinds/" + n.UnitKind
}

func (s *MockService) parseUnitKindName(name string) (*unitKindName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "unitKinds" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &unitKindName{
			Project:  project,
			Location: tokens[3],
			UnitKind: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SaasDeploymentsServer) GetUnitKind(ctx context.Context, req *pb.GetUnitKindRequest) (*pb.UnitKind, error) {
	name, err := s.parseUnitKindName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UnitKind{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SaasDeploymentsServer) CreateUnitKind(ctx context.Context, req *pb.CreateUnitKindRequest) (*pb.UnitKind, error) {
	reqName := req.Parent + "/unitKinds/" + req.UnitKindId
	name, err := s.parseUnitKindName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.UnitKind).(*pb.UnitKind)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = "mock-etag"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.UnitKind)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	return response, nil
}

func (s *SaasDeploymentsServer) UpdateUnitKind(ctx context.Context, req *pb.UpdateUnitKindRequest) (*pb.UnitKind, error) {
	reqName := req.GetUnitKind().GetName()
	name, err := s.parseUnitKindName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UnitKind{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.UnitKind, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update UnitKind fields: %v", err)
		}
	} else {
		desired := req.GetUnitKind()
		obj.DefaultRelease = desired.DefaultRelease
		obj.Dependencies = desired.Dependencies
		obj.InputVariableMappings = desired.InputVariableMappings
		obj.OutputVariableMappings = desired.OutputVariableMappings
		obj.Saas = desired.Saas
		obj.Labels = desired.Labels
		obj.Annotations = desired.Annotations
	}
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = "mock-etag-updated"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.UnitKind)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	return response, nil
}

func (s *SaasDeploymentsServer) DeleteUnitKind(ctx context.Context, req *pb.DeleteUnitKindRequest) (*emptypb.Empty, error) {
	name, err := s.parseUnitKindName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UnitKind{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
