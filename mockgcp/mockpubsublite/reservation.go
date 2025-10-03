// Copyright 2023 Google LLC
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
// proto.service: google.cloud.pubsublite.v1.AdminService
// proto.message: google.cloud.pubsublite.v1.Reservation

package mockpubsublite

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/pubsublite/v1"
)

func (s *PubSubLiteV1) GetReservation(ctx context.Context, req *pb.GetReservationRequest) (*pb.Reservation, error) {
	name, err := s.parseReservationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Reservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PubSubLiteV1) CreateReservation(ctx context.Context, req *pb.CreateReservationRequest) (*pb.Reservation, error) {
	reqName := req.Parent + "/reservations/" + req.ReservationId
	name, err := s.parseReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Reservation).(*pb.Reservation)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PubSubLiteV1) UpdateReservation(ctx context.Context, req *pb.UpdateReservationRequest) (*pb.Reservation, error) {
	reqName := req.GetReservation().GetName()

	name, err := s.parseReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Reservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "throughputCapacity":
			obj.ThroughputCapacity = req.GetReservation().GetThroughputCapacity()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PubSubLiteV1) DeleteReservation(ctx context.Context, req *pb.DeleteReservationRequest) (*emptypb.Empty, error) {
	name, err := s.parseReservationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Reservation{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type reservationName struct {
	Project         *projects.ProjectData
	Location        string
	ReservationName string
}

func (n *reservationName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/reservations/%s", n.Project.Number, n.Location, n.ReservationName)
}

func (s *MockService) parseReservationName(name string) (*reservationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "reservations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &reservationName{
			Project:         project,
			Location:        tokens[3],
			ReservationName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
