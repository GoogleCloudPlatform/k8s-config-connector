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

// +tool:mockgcp-support
// proto.service: google.cloud.compute.v1beta.FutureReservations
// proto.message: google.cloud.compute.v1beta.FutureReservation

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	pbv1beta "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1beta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type FutureReservationsV1beta struct {
	*MockService
	pbv1beta.UnimplementedFutureReservationsServer
}

func (s *FutureReservationsV1beta) Get(ctx context.Context, req *pbv1beta.GetFutureReservationRequest) (*pbv1beta.FutureReservation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/futureReservations/%s", req.GetProject(), req.GetZone(), req.GetFutureReservation())
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pbv1beta.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *FutureReservationsV1beta) Insert(ctx context.Context, req *pbv1beta.InsertFutureReservationRequest) (*pbv1beta.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/futureReservations/%s", req.GetProject(), req.GetZone(), req.GetFutureReservationResource().GetName())
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetFutureReservationResource()).(*pbv1beta.FutureReservation)
	obj.Id = proto.Uint64(s.generateID())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s/futureReservations/%d", name.Project, name.Zone, *obj.Id)))
	obj.Kind = PtrTo("compute#futureReservation")
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Zone = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project, name.Zone)))
	// // Set fields during insert/creation as expected by golden file
	obj.AutoCreatedReservationsDeleteTime = PtrTo("2025-12-31T07:00:00Z")
	obj.PlanningStatus = PtrTo("DRAFT")
	obj.SpecificReservationRequired = PtrTo(false)
	obj.Status = &pbv1beta.FutureReservationStatus{
		ProcurementStatus: PtrTo("DRAFTING"),
	}
	// Clear user input fields that GCP doesn't record in real response
	obj.AutoDeleteAutoCreatedReservations = nil
	obj.Description = nil

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pbv1beta.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FutureReservationsV1beta) Update(ctx context.Context, req *pbv1beta.UpdateFutureReservationRequest) (*pbv1beta.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/futureReservations/%s", req.GetProject(), req.GetZone(), req.GetFutureReservation())
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pbv1beta.FutureReservation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	proto.Merge(obj, req.GetFutureReservationResource())

	// Set fields during updates as expected by golden file
	obj.AutoDeleteAutoCreatedReservations = nil
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.Status == nil {
		obj.Status = &pbv1beta.FutureReservationStatus{}
	}
	obj.Status.ExistingMatchingUsageInfo = &pbv1beta.FutureReservationStatusExistingMatchingUsageInfo{
		Count:     PtrTo(int64(0)),
		Timestamp: PtrTo("2025-08-04T04:26:32.088Z"),
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pbv1beta.Operation{
		OperationType: PtrTo("update"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *FutureReservationsV1beta) Delete(ctx context.Context, req *pbv1beta.DeleteFutureReservationRequest) (*pbv1beta.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/futureReservations/%s", req.GetProject(), req.GetZone(), req.GetFutureReservation())
	name, err := s.parseFutureReservationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pbv1beta.FutureReservation{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pbv1beta.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type futureReservationName struct {
	Project             string
	Zone                string
	FutureReservationID string
}

func (f *futureReservationName) String() string {
	return fmt.Sprintf("projects/%s/zones/%s/futureReservations/%s", f.Project, f.Zone, f.FutureReservationID)
}

// parseFutureReservationName parses a string into a futureReservationName.
// The expected form is `projects/*/zones/*/futureReservations/*`.
func (s *MockService) parseFutureReservationName(name string) (*futureReservationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "futureReservations" {
		name := &futureReservationName{
			Project:             tokens[1],
			Zone:                tokens[3],
			FutureReservationID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
