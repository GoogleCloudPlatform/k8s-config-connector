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
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type capacityCommitmentName struct {
	Project    *projects.ProjectData
	Location   string
	ResourceID string
}

func (n *capacityCommitmentName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/capacityCommitments/" + n.ResourceID
}

// parseCapacityCommitmentName parses a string into a capacityCommitmentName.
// The expected form is projects/<projectId>/locations/<location>/capacityCommitments/<capacityCommitmentID>
func (s *MockService) parseCapacityCommitmentName(name string) (*capacityCommitmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "capacityCommitments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &capacityCommitmentName{
			Project:    project,
			Location:   tokens[3],
			ResourceID: tokens[5],
		}
		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ReservationV1) CreateCapacityCommitment(ctx context.Context, req *pb.CreateCapacityCommitmentRequest) (*pb.CapacityCommitment, error) {
	var reqName string
	if req.CapacityCommitmentId != "" {
		reqName = req.Parent + "/capacityCommitments/" + req.CapacityCommitmentId
	} else if req.CapacityCommitment.Name != "" {
		reqName = req.CapacityCommitment.Name
	} else {
		reqName = req.Parent + "/capacityCommitments/" + "71389360-641d-541e-2kfzymot3v66w6q"
	}

	name, err := s.parseCapacityCommitmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.CapacityCommitment).(*pb.CapacityCommitment)
	obj.Name = fqn
	obj.State = pb.CapacityCommitment_ACTIVE
	obj.CommitmentStartTime = &timestamppb.Timestamp{
		Seconds: now.Unix(),
	}
	obj.CommitmentEndTime = &timestamppb.Timestamp{
		Seconds: now.Add(time.Hour).Unix(),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ReservationV1) GetCapacityCommitment(ctx context.Context, req *pb.GetCapacityCommitmentRequest) (*pb.CapacityCommitment, error) {
	name, err := s.parseCapacityCommitmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CapacityCommitment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: CapacityCommitment %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ReservationV1) UpdateCapacityCommitment(ctx context.Context, req *pb.UpdateCapacityCommitmentRequest) (*pb.CapacityCommitment, error) {
	name, err := s.parseCapacityCommitmentName(req.CapacityCommitment.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CapacityCommitment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := fields.UpdateByFieldMask(obj, req.CapacityCommitment, req.UpdateMask.Paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ReservationV1) DeleteCapacityCommitment(ctx context.Context, req *pb.DeleteCapacityCommitmentRequest) (*emptypb.Empty, error) {
	name, err := s.parseCapacityCommitmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.CapacityCommitment{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ReservationV1) ListCapacityCommitments(ctx context.Context, req *pb.ListCapacityCommitmentsRequest) (*pb.ListCapacityCommitmentsResponse, error) {
	parent := req.GetParent()

	response := &pb.ListCapacityCommitmentsResponse{}

	kind := (&pb.CapacityCommitment{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		cc := obj.(*pb.CapacityCommitment)
		name, err := s.parseCapacityCommitmentName(cc.Name)
		if err != nil {
			return err
		}
		expectedParent := "projects/" + name.Project.ID + "/locations/" + name.Location
		if expectedParent == parent {
			response.CapacityCommitments = append(response.CapacityCommitments, cc)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
