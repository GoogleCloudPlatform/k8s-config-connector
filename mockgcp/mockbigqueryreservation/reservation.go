// Copyright 2025 Google LLC
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
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/reservation/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReservationV1 struct {
	*MockService
	pb.UnimplementedReservationServiceServer
}

func (s *ReservationV1) CreateReservation(ctx context.Context, req *pb.CreateReservationRequest) (*pb.Reservation, error) {
	var reqName string
	if req.ReservationId != "" {
		reqName = req.Parent + "/reservations/" + req.ReservationId
	} else if req.Reservation.Name != "" {
		reqName = req.Reservation.Name
	} else {
		// reqName = req.Parent + "/connections/" + uuid.New().String()
		// Using fixed UUID to test "acquire" in spec.resourceID. This also fix the dynamic uuid value in the `x-goog-request-params` header.
		reqName = req.Parent + "/reservations/" + "91389360-641d-541e-2kfzymot3v66w6q"
	}

	name, err := s.parseReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Reservation).(*pb.Reservation)
	obj.Name = fqn
	obj.CreationTime = &timestamppb.Timestamp{
		Seconds: now.Unix(),
	}
	obj.UpdateTime = &timestamppb.Timestamp{
		Seconds: now.Unix(),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ReservationV1) GetReservation(ctx context.Context, req *pb.GetReservationRequest) (*pb.Reservation, error) {
	name, err := s.parseReservationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Reservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Reservation %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ReservationV1) UpdateReservation(ctx context.Context, req *pb.UpdateReservationRequest) (*pb.Reservation, error) {
	name, err := s.parseReservationName(req.Reservation.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Reservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := fields.UpdateByFieldMask(obj, req.Reservation, req.UpdateMask.Paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	obj.UpdateTime = &timestamppb.Timestamp{
		Seconds: now.Unix(),
	}
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ReservationV1) DeleteReservation(ctx context.Context, req *pb.DeleteReservationRequest) (*empty.Empty, error) {
	name, err := s.parseReservationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.Reservation{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type reservationName struct {
	Project    *projects.ProjectData
	Location   string
	ResourceID string
}

func (n *reservationName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/reservations/" + n.ResourceID
}

// parseReservationName parses a string into a reservationName.
// The expected form is projects/<projectId>/locations/<location>/reservations/<reservationID>
func (s *MockService) parseReservationName(name string) (*reservationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "reservations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &reservationName{
			Project:    project,
			Location:   tokens[3],
			ResourceID: tokens[5],
		}
		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
