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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ReservationsV1) Insert(ctx context.Context, req *pb.InsertReservationRequest) (*pb.Operation, error) {
	if req.GetReservationResource() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "reservation_resource is required")
	}
	reqName := "projects/" + req.GetProject() + "/zones/" + req.GetZone() + "/reservations/" + req.GetReservationResource().GetName()
	name, err := s.parseZonalReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetReservationResource()).(*pb.Reservation)

	if obj.GetShareSettings().GetShareType() == "SPECIFIC_PROJECTS" {
		if len(obj.GetShareSettings().GetProjectMap()) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "project_map is required when share_type is SPECIFIC_PROJECTS")
		}
	}

	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#reservation")
	obj.Zone = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))
	obj.Status = PtrTo("READY")
	if obj.SpecificReservation != nil {
		obj.SpecificReservation.InUseCount = PtrTo(int64(0))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	if req.GetReservationResource() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "reservation_resource is required")
	}

	update := req.GetReservationResource()
	if update.ShareSettings != nil {
		if obj.ShareSettings == nil {
			obj.ShareSettings = &pb.ShareSettings{}
		}
		if update.ShareSettings.ShareType != nil {
			obj.ShareSettings.ShareType = update.ShareSettings.ShareType
		}
		if update.ShareSettings.ProjectMap != nil {
			if obj.ShareSettings.ProjectMap == nil {
				obj.ShareSettings.ProjectMap = make(map[string]*pb.ShareSettingsProjectConfig)
			}
			for k, v := range update.ShareSettings.ProjectMap {
				obj.ShareSettings.ProjectMap[k] = v
			}
		}
	}

	// For other fields, we use proto.Merge
	updateCopy := proto.Clone(update).(*pb.Reservation)
	updateCopy.ShareSettings = nil
	// Preserve immutable fields
	updateCopy.Zone = nil
	updateCopy.SelfLink = nil
	updateCopy.Id = nil
	updateCopy.Kind = nil
	updateCopy.Status = nil
	updateCopy.CreationTimestamp = nil
	proto.Merge(obj, updateCopy)

	if obj.GetShareSettings() != nil && obj.GetShareSettings().GetShareType() == "SPECIFIC_PROJECTS" {
		if len(obj.GetShareSettings().GetProjectMap()) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "project_map is required when share_type is SPECIFIC_PROJECTS")
		}
	}

	obj.Status = PtrTo("READY")
	if obj.SpecificReservation != nil {
		obj.SpecificReservation.InUseCount = PtrTo(int64(0))
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
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
