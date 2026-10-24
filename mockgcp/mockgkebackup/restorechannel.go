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
// proto.message: google.cloud.gkebackup.v1.RestoreChannel

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *BackupForGKEV1) GetRestoreChannel(ctx context.Context, req *pb.GetRestoreChannelRequest) (*pb.RestoreChannel, error) {
	name, err := s.parseRestoreChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.RestoreChannel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupForGKEV1) CreateRestoreChannel(ctx context.Context, req *pb.CreateRestoreChannelRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/restoreChannels/%s", req.GetParent(), req.GetRestoreChannelId())
	name, err := s.parseRestoreChannelName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetRestoreChannel()).(*pb.RestoreChannel)
	if obj.DestinationProject == fmt.Sprintf("projects/%s", name.Project.ID) || obj.DestinationProject == name.Project.ID {
		return nil, status.Errorf(codes.InvalidArgument, "source and destination project cannot be same: invalid argument")
	}
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.RestoreChannelID

	destProjToken := strings.Split(obj.DestinationProject, "/")
	if len(destProjToken) == 2 && destProjToken[0] == "projects" {
		projectID := destProjToken[1]
		obj.DestinationProjectId = projectID
		if projectID == "shared-reservations-project" {
			obj.DestinationProject = "projects/1061767337696"
		} else if destProject, err := s.Projects.GetProjectByID(projectID); err == nil {
			obj.DestinationProject = fmt.Sprintf("projects/%d", destProject.Number)
		}
	}

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

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *BackupForGKEV1) UpdateRestoreChannel(ctx context.Context, req *pb.UpdateRestoreChannelRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRestoreChannelName(req.GetRestoreChannel().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.RestoreChannel{}
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
			obj.Description = req.GetRestoreChannel().GetDescription()
		case "labels":
			obj.Labels = req.GetRestoreChannel().GetLabels()
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
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		updatedObj := &pb.RestoreChannel{}
		if err := s.storage.Get(ctx, fqn, updatedObj); err != nil {
			return nil, err
		}
		return updatedObj, nil
	})
}

func (s *BackupForGKEV1) DeleteRestoreChannel(ctx context.Context, req *pb.DeleteRestoreChannelRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRestoreChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deletedObj := &pb.RestoreChannel{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return &emptypb.Empty{}, nil
	})
}

func (s *BackupForGKEV1) ListRestoreChannels(ctx context.Context, req *pb.ListRestoreChannelsRequest) (*pb.ListRestoreChannelsResponse, error) {
	response := &pb.ListRestoreChannelsResponse{}
	findKind := (&pb.RestoreChannel{}).ProtoReflect().Descriptor()
	prefix := req.GetParent() + "/restoreChannels/"
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.RestoreChannel)
		response.RestoreChannels = append(response.RestoreChannels, item)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

type restoreChannelName struct {
	Project          *projects.ProjectData
	Location         string
	RestoreChannelID string
}

func (n *restoreChannelName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/restoreChannels/%s", n.Project.ID, n.Location, n.RestoreChannelID)
}

func (s *MockService) parseRestoreChannelName(name string) (*restoreChannelName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "restoreChannels" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &restoreChannelName{
			Project:          project,
			Location:         tokens[3],
			RestoreChannelID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid format for a restore channel", name)
}
