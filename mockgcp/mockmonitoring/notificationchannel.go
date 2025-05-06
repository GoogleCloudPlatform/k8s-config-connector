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

// +tool:mockgcp-support
// proto.service: google.monitoring.v3.NotificationChannelService
// proto.message: google.monitoring.v3.NotificationChannel

package mockmonitoring

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/golang/protobuf/ptypes/empty"
)

type NotificationChannelService struct {
	*MockService
	pb.UnimplementedNotificationChannelServiceServer
}

func (s *NotificationChannelService) GetNotificationChannel(ctx context.Context, req *pb.GetNotificationChannelRequest) (*pb.NotificationChannel, error) {
	name, err := s.parseNotificationChannelName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NotificationChannel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *NotificationChannelService) ListNotificationChannels(ctx context.Context, req *pb.ListNotificationChannelsRequest) (*pb.ListNotificationChannelsResponse, error) {
	name, err := s.parseNotificationChannelName(req.GetName() + "/notificationChannels/" + "placeholder")
	if err != nil {
		return nil, err
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	var notificationChannels []*pb.NotificationChannel

	findKind := (&pb.NotificationChannel{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		notificationChannel := obj.(*pb.NotificationChannel)
		notificationChannels = append(notificationChannels, notificationChannel)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListNotificationChannelsResponse{
		NotificationChannels: notificationChannels,
		TotalSize:            int32(len(notificationChannels)),
	}, nil
}

func (s *NotificationChannelService) CreateNotificationChannel(ctx context.Context, req *pb.CreateNotificationChannelRequest) (*pb.NotificationChannel, error) {
	now := time.Now()

	channelID := fmt.Sprintf("%d", now.UnixNano())

	reqName := req.GetName() + "/notificationChannels/" + channelID
	name, err := s.parseNotificationChannelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.NotificationChannel).(*pb.NotificationChannel)
	obj.CreationRecord = &pb.MutationRecord{
		MutateTime: timestamppb.New(now),
	}
	obj.MutationRecords = append(obj.MutationRecords, &pb.MutationRecord{
		MutateTime: timestamppb.New(now),
	})
	obj.Name = fqn
	obj.Enabled = wrapperspb.Bool(true)

	redactPasswords(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func redactPasswords(obj *pb.NotificationChannel) {
	labels := obj.GetLabels()
	if labels["password"] != "" {
		labels["password"] = "*********"
	}
	obj.Labels = labels
}

func (s *NotificationChannelService) UpdateNotificationChannel(ctx context.Context, req *pb.UpdateNotificationChannelRequest) (*pb.NotificationChannel, error) {
	name, err := s.parseNotificationChannelName(req.GetNotificationChannel().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.NotificationChannel{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := proto.Clone(existing).(*pb.NotificationChannel)
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "description":
			updated.Description = req.GetNotificationChannel().GetDescription()
		case "display_name":
			updated.DisplayName = req.GetNotificationChannel().GetDisplayName()
		case "enabled":
			updated.Enabled = req.GetNotificationChannel().GetEnabled()
		case "labels":
			updated.Labels = req.GetNotificationChannel().GetLabels()
		case "type":
			updated.Type = req.GetNotificationChannel().GetType()
		case "user_labels":
			updated.UserLabels = req.GetNotificationChannel().GetUserLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock (full update_mask=%v)", path, req.GetUpdateMask())
		}
	}

	redactPasswords(updated)

	updated.MutationRecords = append(updated.MutationRecords, &pb.MutationRecord{
		MutateTime: timestamppb.New(now),
	})

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *NotificationChannelService) DeleteNotificationChannel(ctx context.Context, req *pb.DeleteNotificationChannelRequest) (*empty.Empty, error) {
	name, err := s.parseNotificationChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.NotificationChannel{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type NotificationChannelName struct {
	Project                 *projects.ProjectData
	NotificationChannelName string
}

func (n *NotificationChannelName) String() string {
	return fmt.Sprintf("projects/%s/notificationChannels/%s", n.Project.ID, n.NotificationChannelName)
}

// parseNotificationChannelName parses a string into a NotificationChannelName.
// The expected form is projects/[PROJECT_ID_OR_NUMBER]/NotificationChannels/[NotificationChannel_ID]
func (s *MockService) parseNotificationChannelName(name string) (*NotificationChannelName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "notificationChannels" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &NotificationChannelName{
			Project:                 project,
			NotificationChannelName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
