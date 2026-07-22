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

// +tool:mockgcp-support
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.BackupChannel

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

func (s *BackupForGKEV1) GetBackupChannel(ctx context.Context, req *pb.GetBackupChannelRequest) (*pb.BackupChannel, error) {
	name, err := s.parseBackupChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupChannel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackupChannel %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupForGKEV1) CreateBackupChannel(ctx context.Context, req *pb.CreateBackupChannelRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupChannels/%s", req.GetParent(), req.GetBackupChannelId())
	name, err := s.parseBackupChannelName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.GetBackupChannel())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.BackupChannelID

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)

		destProjToken := strings.Split(obj.DestinationProject, "/")
		if len(destProjToken) == 2 && destProjToken[0] == "projects" {
			obj.DestinationProjectId = destProjToken[1]
		}

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *BackupForGKEV1) UpdateBackupChannel(ctx context.Context, req *pb.UpdateBackupChannelRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupChannelName(req.GetBackupChannel().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupChannel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetBackupChannel().GetDescription()
		case "labels":
			obj.Labels = req.GetBackupChannel().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q is not supported", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

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
		lroMetadata.EndTime = timestamppb.New(now)
		updatedObj := &pb.BackupChannel{}
		if err := s.storage.Get(ctx, fqn, updatedObj); err != nil {
			return nil, err
		}
		return updatedObj, nil
	})
}

func (s *BackupForGKEV1) DeleteBackupChannel(ctx context.Context, req *pb.DeleteBackupChannelRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deletedObj := &pb.BackupChannel{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackupChannel %q not found", fqn)
		}
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
		lroMetadata.EndTime = timestamppb.New(now)
		return &emptypb.Empty{}, nil
	})
}

type backupChannelName struct {
	Project         *projects.ProjectData
	Location        string
	BackupChannelID string
}

func (n *backupChannelName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupChannels/%s", n.Project.ID, n.Location, n.BackupChannelID)
}

func (s *MockService) parseBackupChannelName(name string) (*backupChannelName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupChannels" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupChannelName{
			Project:         project,
			Location:        tokens[3],
			BackupChannelID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid format for a backup channel", name)
}
