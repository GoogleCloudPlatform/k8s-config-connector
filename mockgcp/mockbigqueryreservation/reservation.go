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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

	if obj.SecondaryLocation != "" {
		if obj.Edition == pb.Edition_ENTERPRISE_PLUS {
			obj.PrimaryLocation = name.Location
			obj.OriginalPrimaryLocation = obj.PrimaryLocation
		} else {
			return nil, status.Error(codes.InvalidArgument, "secondary_location can only be specified for ENTERPRISE_PLUS edition")
		}
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

	updateSecondaryLocation := false
	for _, p := range req.UpdateMask.Paths {
		if p == "secondaryLocation" {
			updateSecondaryLocation = true
			break
		}
	}

	if updateSecondaryLocation {
		if obj.Edition != pb.Edition_ENTERPRISE_PLUS {
			return nil, status.Error(codes.InvalidArgument, "secondary_location can only be specified for ENTERPRISE_PLUS edition")
		}
		if obj.SecondaryLocation != "" && req.Reservation.SecondaryLocation != "" {
			return nil, status.Error(codes.InvalidArgument, "Changing the secondary location of an existing failover reservation is not supported. Please remove the existing secondary location before setting a new secondary location.")
		}
	}

	if err := fields.UpdateByFieldMask(obj, req.Reservation, req.UpdateMask.Paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	if updateSecondaryLocation {
		// handle the case of changing from failover to non-failover reservation
		if obj.SecondaryLocation == "" {
			obj = setFieldfailoverNonFailover(obj)
		} else {
			obj = setFieldNonfailoverFailover(obj, name.Location)
		}
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

// For Enterprise_plus, user can switch between failover and non-failover reservations
// by setting or unsetting the secondary_location.
// When user unsets the secondary_location, setFieldfailoverNonFailover removes
// the primary_location, original_primary_location and secondary_location in the response.
func setFieldfailoverNonFailover(obj *pb.Reservation) *pb.Reservation {
	if obj.OriginalPrimaryLocation != "" && obj.PrimaryLocation != "" {
		obj.OriginalPrimaryLocation = ""
		obj.PrimaryLocation = ""
	}
	return obj
}

func setFieldNonfailoverFailover(obj *pb.Reservation, location string) *pb.Reservation {
	if obj.OriginalPrimaryLocation == "" && obj.PrimaryLocation == "" {
		obj.OriginalPrimaryLocation = location
		obj.PrimaryLocation = location
	}
	return obj
}

type assignmentName struct {
	Reservation *reservationName
	ResourceID  string
}

// parseAssignmentName parses a string into a assignmentName.
// The expected form is projects/<projectId>/locations/<location>/reservations/<reservationID>/assignments/<assignmentID>
func (s *MockService) parseAssignmentName(name string) (*assignmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "reservations" && tokens[6] == "assignments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &assignmentName{
			Reservation: &reservationName{
				Project:    project,
				Location:   tokens[3],
				ResourceID: tokens[5],
			},
			ResourceID: tokens[7],
		}
		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (n *assignmentName) String() string {
	return "projects/" + n.Reservation.Project.ID + "/locations/" + n.Reservation.Location + "/reservations/" + n.Reservation.ResourceID + "/assignments/" + n.ResourceID
}

// Creates an assignment object which allows the given project to submit jobs
// of a certain type using slots from the specified reservation.
func (s *ReservationV1) CreateAssignment(ctx context.Context, req *pb.CreateAssignmentRequest) (*pb.Assignment, error) {

	// reqName = req.Parent + "/assignments/" + uuid.New().String()
	// Using fixed UUID to test "acquire" in spec.resourceID. This also fix the dynamic uuid value in the `x-goog-request-params` header.
	reqName := req.Parent + "/assignments/" + "87687860-6689-5789-1dfzymot3v66w7f"

	name, err := s.parseAssignmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Assignment).(*pb.Assignment)
	obj.Name = fqn
	obj.State = pb.Assignment_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// Lists assignments.
// Only explicitly created assignments will be returned.
func (s *ReservationV1) ListAssignments(ctx context.Context, req *pb.ListAssignmentsRequest) (*pb.ListAssignmentsResponse, error) {
	parent, err := s.parseReservationName(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListAssignmentsResponse{}

	assignmentKind := (&pb.Assignment{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, assignmentKind, storage.ListOptions{}, func(obj proto.Message) error {
		assignment := obj.(*pb.Assignment)
		var name *assignmentName
		name, err = s.parseAssignmentName(assignment.Name)
		if err != nil {
			return err
		}
		if name.Reservation.String() == parent.String() {
			response.Assignments = append(response.Assignments, assignment)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

// Moves an assignment under a new reservation.
//
// This differs from removing an existing assignment and recreating a new one
// by providing a transactional change that ensures an assignee always has an
// associated reservation.
func (s *ReservationV1) MoveAssignment(ctx context.Context, req *pb.MoveAssignmentRequest) (*pb.Assignment, error) {
	name, err := s.parseAssignmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Assignment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Rebuild name for the Assignment
	obj.Name = req.DestinationId + "/assignments/" + "27687860-6459-5709-1dfzymot3v66w8h"
	// Delete and recreate
	if err := s.storage.Delete(ctx, fqn, &pb.Assignment{}); err != nil {
		return nil, err
	}
	if err := s.storage.Create(ctx, obj.Name, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// Deletes a assignment. No expansion will happen.
func (s *ReservationV1) DeleteAssignment(ctx context.Context, req *pb.DeleteAssignmentRequest) (*empty.Empty, error) {
	name, err := s.parseAssignmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.Assignment{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
