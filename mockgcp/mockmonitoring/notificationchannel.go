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
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/v3"
	"github.com/golang/protobuf/ptypes/empty"
)

type NotificationsChannelService struct {
	*MockService
	pb.UnimplementedNotificationChannelServiceServer
}

func (s *NotificationsChannelService) GetNotificationChannel(ctx context.Context, req *pb.GetNotificationChannelRequest) (*pb.NotificationChannel, error) {
	name, err := s.parseNotificationChannelName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NotificationChannel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "notificationChannel %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading notificationChannel: %v", err)
		}
	}

	return redactNotificationChannel(obj), nil
}

func redactNotificationChannel(obj *pb.NotificationChannel) *pb.NotificationChannel {
	// Fields containing sensitive information like authentication tokens or contact info are only partially populated on retrieval.
	redacted := proto.Clone(obj).(*pb.NotificationChannel)
	for k, v := range redacted.Labels {
		switch k {
		case "password":
			redacted.Labels[k] = strings.Repeat("*", len(v))
		}
	}
	return redacted
}

func (s *NotificationsChannelService) CreateNotificationChannel(ctx context.Context, req *pb.CreateNotificationChannelRequest) (*pb.NotificationChannel, error) {
	now := time.Now()

	channelID := fmt.Sprintf("%d", now.UnixNano())

	reqName := req.GetName() + "/notificationChannels/" + channelID
	name, err := s.parseNotificationChannelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.NotificationChannel).(*pb.NotificationChannel)
	obj.Name = fqn

	creationRecord := &pb.MutationRecord{
		MutateTime: timestamppb.New(now),
	}

	obj.CreationRecord = creationRecord
	obj.MutationRecords = []*pb.MutationRecord{
		creationRecord,
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating notificationChannel: %v", err)
	}

	return redactNotificationChannel(obj), nil
}

func (s *NotificationsChannelService) UpdateNotificationChannel(ctx context.Context, req *pb.UpdateNotificationChannelRequest) (*pb.NotificationChannel, error) {
	name, err := s.parseNotificationChannelName(req.GetNotificationChannel().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.NotificationChannel{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "notificationChannel %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading notificationChannel: %v", err)
		}
	}

	now := time.Now()

	updated := proto.Clone(req.NotificationChannel).(*pb.NotificationChannel)
	updated.CreationRecord = existing.CreationRecord
	updated.MutationRecords = append(existing.MutationRecords, &pb.MutationRecord{
		MutateTime: timestamppb.New(now),
	})

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating notificationChannel: %v", err)
	}

	return redactNotificationChannel(updated), nil
}

func (s *NotificationsChannelService) DeleteNotificationChannel(ctx context.Context, req *pb.DeleteNotificationChannelRequest) (*empty.Empty, error) {
	name, err := s.parseNotificationChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.NotificationChannel{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "notificationChannel %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting notificationChannel: %v", err)
		}
	}

	return &empty.Empty{}, nil
}

type notificationChannelName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *notificationChannelName) String() string {
	return "projects/" + n.Project.ID + "/notificationChannels/" + n.Name
}

// parseNotificationChannelName parses a string into a notificationChannelName.
// The expected form is projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]
func (s *MockService) parseNotificationChannelName(name string) (*notificationChannelName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "notificationChannels" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &notificationChannelName{
			Project: project,
			Name:    tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
