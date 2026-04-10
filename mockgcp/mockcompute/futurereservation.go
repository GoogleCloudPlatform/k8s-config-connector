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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pbv1 "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1beta"
	"google.golang.org/protobuf/proto"
)

type FutureReservationsV1Beta1 struct {
	*MockService
	pb.UnimplementedFutureReservationsServer
}

func (s *FutureReservationsV1Beta1) Get(ctx context.Context, req *pb.GetFutureReservationRequest) (*pb.FutureReservation, error) {
	name, err := s.newFutureReservationName(req.GetProject(), req.GetZone(), req.GetFutureReservation())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *FutureReservationsV1Beta1) Insert(ctx context.Context, req *pb.InsertFutureReservationRequest) (*pb.Operation, error) {
	name, err := s.newFutureReservationName(req.GetProject(), req.GetZone(), req.GetFutureReservationResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetFutureReservationResource()).(*pb.FutureReservation)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#futureReservation")
	obj.Zone = PtrTo(buildComputeSelfLink(ctx, name.ZoneName()))
	obj.Status = &pb.FutureReservationStatus{
		ProcurementStatus: PtrTo("APPROVED"),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op, err := s.newLRO(ctx, name.Project.ID)
	if err != nil {
		return nil, err
	}
	return s.convertOperationToV1Beta(op), nil
}

func (s *FutureReservationsV1Beta1) Update(ctx context.Context, req *pb.UpdateFutureReservationRequest) (*pb.Operation, error) {
	name, err := s.newFutureReservationName(req.GetProject(), req.GetZone(), req.GetFutureReservation())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Simple update logic
	newObj := proto.Clone(req.GetFutureReservationResource()).(*pb.FutureReservation)
	newObj.Id = obj.Id
	newObj.SelfLink = obj.SelfLink
	newObj.CreationTimestamp = obj.CreationTimestamp
	newObj.Kind = obj.Kind
	newObj.Zone = obj.Zone
	newObj.Status = obj.Status

	if err := s.storage.Update(ctx, fqn, newObj); err != nil {
		return nil, err
	}

	op, err := s.newLRO(ctx, name.Project.ID)
	if err != nil {
		return nil, err
	}
	return s.convertOperationToV1Beta(op), nil
}

func (s *FutureReservationsV1Beta1) Delete(ctx context.Context, req *pb.DeleteFutureReservationRequest) (*pb.Operation, error) {
	name, err := s.newFutureReservationName(req.GetProject(), req.GetZone(), req.GetFutureReservation())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FutureReservation{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op, err := s.newLRO(ctx, name.Project.ID)
	if err != nil {
		return nil, err
	}
	return s.convertOperationToV1Beta(op), nil
}

func (s *FutureReservationsV1Beta1) convertOperationToV1Beta(v1Op *pbv1.Operation) *pb.Operation {
	if v1Op == nil {
		return nil
	}
	// This is a bit hacky but should work since they have the same structure
	v1OpData, err := proto.Marshal(v1Op)
	if err != nil {
		panic(err)
	}
	v1betaOp := &pb.Operation{}
	if err := proto.Unmarshal(v1OpData, v1betaOp); err != nil {
		panic(err)
	}
	return v1betaOp
}

type futureReservationName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *futureReservationName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/futureReservations/" + n.Name
}

func (n *futureReservationName) ZoneName() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone
}

func (s *MockService) newFutureReservationName(project string, zone string, name string) (*futureReservationName, error) {
	projectObj, err := s.Projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}

	return &futureReservationName{
		Project: projectObj,
		Zone:    zone,
		Name:    name,
	}, nil
}
