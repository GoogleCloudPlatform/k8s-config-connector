// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockbigqueryreservation

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ReservationV1) CreateReservationGroup(ctx context.Context, req *pb.CreateReservationGroupRequest) (*pb.ReservationGroup, error) {
	reqName := fmt.Sprintf("%s/reservationGroups/%s", req.Parent, req.ReservationGroupId)

	name, err := s.parseReservationGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.ReservationGroup).(*pb.ReservationGroup)
	obj.Name = fqn
	obj.CreationTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ReservationV1) GetReservationGroup(ctx context.Context, req *pb.GetReservationGroupRequest) (*pb.ReservationGroup, error) {
	name, err := s.parseReservationGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ReservationGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No reservation group found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *ReservationV1) UpdateReservationGroup(ctx context.Context, req *pb.UpdateReservationGroupRequest) (*pb.ReservationGroup, error) {
	reqName := req.ReservationGroup.Name
	name, err := s.parseReservationGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	actual := &pb.ReservationGroup{}
	if err := s.storage.Get(ctx, fqn, actual); err != nil {
		return nil, err
	}

	obj := proto.Clone(actual).(*pb.ReservationGroup)
	paths := req.UpdateMask.GetPaths()
	if len(paths) == 0 {
		// If no update mask is specified, update all fields.
		obj.ParentGroup = req.ReservationGroup.ParentGroup
	} else {
		for _, path := range paths {
			switch path {
			case "parent_group":
				obj.ParentGroup = req.ReservationGroup.ParentGroup
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask contains invalid path: %q", path)
			}
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ReservationV1) DeleteReservationGroup(ctx context.Context, req *pb.DeleteReservationGroupRequest) (*emptypb.Empty, error) {
	name, err := s.parseReservationGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.ReservationGroup{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No reservation group found.")
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ReservationV1) ListReservationGroups(ctx context.Context, req *pb.ListReservationGroupsRequest) (*pb.ListReservationGroupsResponse, error) {
	parent, err := s.parseReservationParent(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/reservationGroups/", parent.Project.ID, parent.Location)

	var items []*pb.ReservationGroup
	if err := s.storage.List(ctx, (&pb.ReservationGroup{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: prefix,
	}, func(msg proto.Message) error {
		items = append(items, msg.(*pb.ReservationGroup))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListReservationGroupsResponse{
		ReservationGroups: items,
	}, nil
}

type reservationGroupName struct {
	Project          *projects.ProjectData
	Location         string
	ReservationGroup string
}

func (n *reservationGroupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/reservationGroups/%s", n.Project.ID, n.Location, n.ReservationGroup)
}

func (s *MockService) parseReservationGroupName(name string) (*reservationGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "reservationGroups" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ReservationGroup name %q", name)
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	return &reservationGroupName{
		Project:          project,
		Location:         tokens[3],
		ReservationGroup: tokens[5],
	}, nil
}

func (s *MockService) parseReservationParent(parent string) (*reservationParent, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q", parent)
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	return &reservationParent{
		Project:  project,
		Location: tokens[3],
	}, nil
}

type reservationParent struct {
	Project  *projects.ProjectData
	Location string
}
