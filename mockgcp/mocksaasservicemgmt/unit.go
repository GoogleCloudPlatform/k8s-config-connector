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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type unitName struct {
	Project  *projects.ProjectData
	Location string
	Unit     string
}

func (n *unitName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/units/" + n.Unit
}

func (s *MockService) parseUnitName(name string) (*unitName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "units" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &unitName{
			Project:  project,
			Location: tokens[3],
			Unit:     tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SaasDeploymentsServer) GetUnit(ctx context.Context, req *pb.GetUnitRequest) (*pb.Unit, error) {
	name, err := s.parseUnitName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Unit{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SaasDeploymentsServer) CreateUnit(ctx context.Context, req *pb.CreateUnitRequest) (*pb.Unit, error) {
	reqName := req.Parent + "/units/" + req.UnitId
	name, err := s.parseUnitName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Unit).(*pb.Unit)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Etag = "mock-etag"
	obj.State = pb.Unit_UNIT_STATE_NOT_PROVISIONED
	obj.Conditions = []*pb.UnitCondition{
		{
			Type:               pb.UnitCondition_TYPE_UPDATING,
			Status:             pb.UnitCondition_STATUS_FALSE,
			LastTransitionTime: now,
		},
		{
			Type:               pb.UnitCondition_TYPE_PROVISIONED,
			Status:             pb.UnitCondition_STATUS_FALSE,
			LastTransitionTime: now,
		},
		{
			Type:               pb.UnitCondition_TYPE_OPERATION_ERROR,
			Status:             pb.UnitCondition_STATUS_UNKNOWN,
			LastTransitionTime: now,
		},
		{
			Type:               pb.UnitCondition_TYPE_READY,
			Status:             pb.UnitCondition_STATUS_FALSE,
			LastTransitionTime: now,
		},
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Unit)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	response.Conditions = nil
	response.State = pb.Unit_UNIT_STATE_UNSPECIFIED
	return response, nil
}

func (s *SaasDeploymentsServer) UpdateUnit(ctx context.Context, req *pb.UpdateUnitRequest) (*pb.Unit, error) {
	reqName := req.GetUnit().GetName()
	name, err := s.parseUnitName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Unit{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.Labels = req.GetUnit().GetLabels()
		obj.Annotations = req.GetUnit().GetAnnotations()
		obj.Maintenance = req.GetUnit().GetMaintenance()
	} else {
		for _, path := range paths {
			switch path {
			case "maintenance":
				obj.Maintenance = req.GetUnit().GetMaintenance()
			case "labels":
				obj.Labels = req.GetUnit().GetLabels()
			case "annotations":
				obj.Annotations = req.GetUnit().GetAnnotations()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()
	obj.Etag = "mock-etag-updated"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := proto.Clone(obj).(*pb.Unit)
	response.CreateTime = nil
	response.UpdateTime = nil
	response.Etag = ""
	response.Conditions = nil
	response.State = pb.Unit_UNIT_STATE_UNSPECIFIED
	return response, nil
}

func (s *SaasDeploymentsServer) DeleteUnit(ctx context.Context, req *pb.DeleteUnitRequest) (*emptypb.Empty, error) {
	name, err := s.parseUnitName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Unit{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
