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

package mockbigqueryreservation

import (
	"context"
	"strings"

	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type reservationGroupName struct {
	Project          *projects.ProjectData
	Location         string
	ReservationGroup string
}

func (n *reservationGroupName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/reservationGroups/" + n.ReservationGroup
}

// parseReservationGroupName parses a string into a reservationGroupName.
// The expected form is projects/<projectId>/locations/<location>/reservationGroups/<reservationGroupID>
func (s *MockService) parseReservationGroupName(name string) (*reservationGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "reservationGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &reservationGroupName{
			Project:          project,
			Location:         tokens[3],
			ReservationGroup: tokens[5],
		}
		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ReservationV1) CreateReservationGroup(ctx context.Context, req *pb.CreateReservationGroupRequest) (*pb.ReservationGroup, error) {
	reqName := req.Parent + "/reservationGroups/" + req.ReservationGroupId
	name, err := s.parseReservationGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ReservationGroup).(*pb.ReservationGroup)
	if obj == nil {
		obj = &pb.ReservationGroup{}
	}
	obj.Name = fqn
	// TODO: Real GCP returns creationTime and updateTime, but they are not defined in the pb.ReservationGroup proto.

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
			return nil, status.Errorf(codes.NotFound, "Reservation group %s not found.", fqn)
		}
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
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ReservationV1) ListReservationGroups(ctx context.Context, req *pb.ListReservationGroupsRequest) (*pb.ListReservationGroupsResponse, error) {
	parent := req.GetParent()

	response := &pb.ListReservationGroupsResponse{}

	kind := (&pb.ReservationGroup{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		rg := obj.(*pb.ReservationGroup)
		name, err := s.parseReservationGroupName(rg.Name)
		if err != nil {
			return err
		}
		expectedParent := "projects/" + name.Project.ID + "/locations/" + name.Location
		if expectedParent == parent {
			response.ReservationGroups = append(response.ReservationGroups, rg)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
