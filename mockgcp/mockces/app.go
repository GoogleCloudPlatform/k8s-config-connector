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

package mockces

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/ces/apiv1/cespb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type AgentServiceV1 struct {
	*MockService
	pb.UnimplementedAgentServiceServer
}

func (s *AgentServiceV1) populateDefaultsForApp(obj *pb.App) {
	if obj.TimeZoneSettings == nil || obj.TimeZoneSettings.TimeZone == "" {
		obj.TimeZoneSettings = &pb.TimeZoneSettings{
			TimeZone: "America/Los_Angeles",
		}
	}
	if len(obj.PredefinedVariableDeclarations) == 0 {
		obj.PredefinedVariableDeclarations = []*pb.App_VariableDeclaration{
			{
				Name:        "current_date",
				Description: "Current time represented in the ISO 8601 format along with its corresponding IANA Time Zone.",
				Schema: &pb.Schema{
					Type: pb.Schema_STRING,
				},
			},
			{
				Name:        "_session",
				Description: "Contains session specific variables like modality, channel, recording control, and language code.",
				Schema: &pb.Schema{
					Type: pb.Schema_OBJECT,
					Properties: map[string]*pb.Schema{
						"channel": {
							Type: pb.Schema_STRING,
						},
						"language_code": {
							Type: pb.Schema_STRING,
						},
						"modality": {
							Type: pb.Schema_STRING,
						},
						"recording_control": {
							Type: pb.Schema_STRING,
						},
					},
				},
			},
		}
	}
	if obj.LoggingSettings == nil {
		obj.LoggingSettings = &pb.LoggingSettings{
			ConversationLoggingSettings: &pb.ConversationLoggingSettings{},
		}
	} else if obj.LoggingSettings.ConversationLoggingSettings == nil {
		obj.LoggingSettings.ConversationLoggingSettings = &pb.ConversationLoggingSettings{}
	}
	// TODO: ConversationLoggingSettings in real GCP defaults retentionWindow to "31536000s",
	// but retentionWindow is not present in cloud.google.com/go/ces/apiv1/cespb.ConversationLoggingSettings.
}

func (s *AgentServiceV1) GetApp(ctx context.Context, req *pb.GetAppRequest) (*pb.App, error) {
	name, err := s.parseAppName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.App{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	s.populateDefaultsForApp(obj)
	return obj, nil
}

func (s *AgentServiceV1) CreateApp(ctx context.Context, req *pb.CreateAppRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	appID := req.GetAppId()
	if appID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "appId is required")
	}

	name := &appName{
		Project:  parent.Project,
		Location: parent.Location,
		AppID:    appID,
	}

	fqn := name.String()

	existing := &pb.App{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
	}

	obj := proto.Clone(req.GetApp()).(*pb.App)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="
	s.populateDefaultsForApp(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *AgentServiceV1) UpdateApp(ctx context.Context, req *pb.UpdateAppRequest) (*pb.App, error) {
	reqApp := req.GetApp()
	if reqApp == nil {
		return nil, status.Errorf(codes.InvalidArgument, "app is required")
	}

	name, err := s.parseAppName(reqApp.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.App{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, reqApp)
	} else {
		if err := fields.UpdateByFieldMask(obj, reqApp, paths); err != nil {
			return nil, err
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = "abcdef0123A="
	s.populateDefaultsForApp(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AgentServiceV1) DeleteApp(ctx context.Context, req *pb.DeleteAppRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAppName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.App{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		if err := s.storage.Delete(ctx, fqn, &pb.App{}); err != nil {
			return nil, err
		}
		return &emptypb.Empty{}, nil
	})
}

func (s *AgentServiceV1) ListApps(ctx context.Context, req *pb.ListAppsRequest) (*pb.ListAppsResponse, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListAppsResponse{}
	findPrefix := fmt.Sprintf("projects/%s/locations/%s/apps/", parent.Project.ID, parent.Location)

	appKind := (&pb.App{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, appKind, storage.ListOptions{}, func(obj proto.Message) error {
		item := obj.(*pb.App)
		if strings.HasPrefix(item.GetName(), findPrefix) {
			s.populateDefaultsForApp(item)
			response.Apps = append(response.Apps, item)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
