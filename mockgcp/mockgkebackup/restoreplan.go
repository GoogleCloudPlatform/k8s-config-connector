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

//go:build mock
// +build mock

// +tool:mockgcp-support
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.RestorePlan

package mockgkebackup

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkebackup/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *BackupForGKEV1) GetRestorePlan(ctx context.Context, req *pb.GetRestorePlanRequest) (*pb.RestorePlan, error) {
	name, err := s.parseRestorePlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RestorePlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "RestorePlan %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupForGKEV1) CreateRestorePlan(ctx context.Context, req *pb.CreateRestorePlanRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/restorePlans/%s", req.GetParent(), req.GetRestorePlanId())
	name, err := s.parseRestorePlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetRestorePlan()).(*pb.RestorePlan)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.RestorePlanID
	obj.Etag = fields.ComputeWeakEtag(obj)
	obj.State = pb.RestorePlan_STATE_UNSPECIFIED

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		obj.State = pb.RestorePlan_READY
		obj.StateReason = "Resource has been created successfully."
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *BackupForGKEV1) UpdateRestorePlan(ctx context.Context, req *pb.UpdateRestorePlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRestorePlanName(req.GetRestorePlan().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.RestorePlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetRestorePlan().GetDescription()
		case "labels":
			obj.Labels = req.GetRestorePlan().GetLabels()
		case "restore_config":
			obj.RestoreConfig = req.GetRestorePlan().GetRestoreConfig()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *BackupForGKEV1) DeleteRestorePlan(ctx context.Context, req *pb.DeleteRestorePlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRestorePlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.RestorePlan{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type restorePlanName struct {
	Project       *projects.ProjectData
	Location      string
	RestorePlanID string
}

func (n *restorePlanName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s", n.Project.ID, n.Location, n.RestorePlanID)
}

// parseRestorePlanName parses a string into an restorePlanName.
// The expected form is `projects/*/locations/*/restorePlans/*`.
func (s *MockService) parseRestorePlanName(name string) (*restorePlanName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "restorePlans" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &restorePlanName{
			Project:       project,
			Location:      tokens[3],
			RestorePlanID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
