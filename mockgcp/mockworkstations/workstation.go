// Copyright 2024 Google LLC
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

package mockworkstations

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/workstations/v1beta"
	"github.com/google/uuid"
)

func (s *workstationsServer) GetWorkstation(ctx context.Context, req *pb.GetWorkstationRequest) (*pb.Workstation, error) {
	name, err := s.parseWorkstationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Workstation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *workstationsServer) CreateWorkstation(ctx context.Context, req *pb.CreateWorkstationRequest) (*longrunning.Operation, error) {
	reqName := req.GetParent() + "/workstations/" + req.GetWorkstationId()
	name, err := s.parseWorkstationName(reqName)
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validate only is not yet implemented in mockgcp")
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Workstation).(*pb.Workstation)
	obj.Name = fqn

	obj.Uid = uuid.New().String()

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *workstationsServer) UpdateWorkstation(ctx context.Context, req *pb.UpdateWorkstationRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkstationName(req.GetWorkstation().GetName())
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validateOnly is not yet implemented in mockgcp")
	}
	if req.GetAllowMissing() {
		return nil, status.Errorf(codes.Unimplemented, "allowMissing is not yet implemented in mockgcp")
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Workstation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetWorkstation().GetDisplayName()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *workstationsServer) DeleteWorkstation(ctx context.Context, req *pb.DeleteWorkstationRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkstationName(req.GetName())
	if err != nil {
		return nil, err
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validateOnly is not yet implemented in mockgcp")
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Workstation{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
		RequestedCancellation: false,
		ApiVersion:            "v1beta",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type WorkstationName struct {
	WorkstationConfigName
	WorkstationID string
}

func (n *WorkstationName) String() string {
	return n.WorkstationConfigName.String() + "/workstations/" + n.WorkstationID
}

// parseWorkstationName parses a string into a WorkstationName.
// The expected form is projects/*/locations/*/workstationClusters/*/workstationConfigs/*/workstations/*
func (s *MockService) parseWorkstationName(name string) (*WorkstationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[8] == "workstations" {
		configName, err := s.parseWorkstationConfigName(strings.Join(tokens[0:8], "/"))
		if err != nil {
			return nil, err
		}

		name := &WorkstationName{
			WorkstationConfigName: *configName,
			WorkstationID:         tokens[9],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
