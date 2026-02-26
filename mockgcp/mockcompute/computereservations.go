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

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type ReservationsV1 struct {
	*MockService
	pb.UnimplementedReservationsServer
}

func (s *ReservationsV1) Get(ctx context.Context, req *pb.GetReservationRequest) (*pb.Reservation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/reservations/" + req.GetReservation()
	name, err := s.parseZonalReservationName(reqName)
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

func (s *ReservationsV1) Insert(ctx context.Context, req *pb.InsertReservationRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/reservations/" + req.GetReservationResource().GetName()
	name, err := s.parseZonalReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetReservationResource()).(*pb.Reservation)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#reservation")
	obj.Zone = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	obj.Status = PtrTo("READY")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ReservationsV1) Update(ctx context.Context, req *pb.UpdateReservationRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/reservations/" + req.GetReservation()
	name, err := s.parseZonalReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Reservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetReservationResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ReservationsV1) Delete(ctx context.Context, req *pb.DeleteReservationRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/reservations/" + req.GetReservation()
	name, err := s.parseZonalReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := (&pb.Reservation{})
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type zonalReservationName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalReservationName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/reservations/" + n.Name
}

func (s *MockService) parseZonalReservationName(name string) (*zonalReservationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "reservations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &zonalReservationName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
