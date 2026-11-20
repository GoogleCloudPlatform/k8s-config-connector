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

package mockdialogflow

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type conversationProfilesServer struct {
	*MockService
	pb.UnimplementedConversationProfilesServer
}

type conversationProfileName struct {
	Project             string
	Location            string
	ConversationProfile string
}

func (n *conversationProfileName) String() string {
	if n.Location != "" {
		return fmt.Sprintf("projects/%s/locations/%s/conversationProfiles/%s", n.Project, n.Location, n.ConversationProfile)
	}
	return fmt.Sprintf("projects/%s/conversationProfiles/%s", n.Project, n.ConversationProfile)
}

func (s *MockService) parseConversationProfileName(name string) (*conversationProfileName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "conversationProfiles" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &conversationProfileName{
			Project:             project.ID,
			Location:            tokens[3],
			ConversationProfile: tokens[5],
		}, nil
	}
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "conversationProfiles" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &conversationProfileName{
			Project:             project.ID,
			ConversationProfile: tokens[3],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/conversationProfiles/{conversation_profile}", name)
}

func (s *MockService) parseConversationProfileParent(parent string) (*conversationProfileName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &conversationProfileName{
			Project:  project.ID,
			Location: tokens[3],
		}, nil
	}
	if len(tokens) == 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &conversationProfileName{
			Project: project.ID,
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

func populateConversationProfileDefaults(obj *pb.ConversationProfile) {
	if obj.GetLanguageCode() == "" {
		obj.LanguageCode = "en-US"
	}
}

func (s *conversationProfilesServer) GetConversationProfile(ctx context.Context, req *pb.GetConversationProfileRequest) (*pb.ConversationProfile, error) {
	name, err := s.parseConversationProfileName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConversationProfile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Conversation profile with ID %s not found in %s.", name.ConversationProfile, name.String())
		}
		return nil, err
	}

	populateConversationProfileDefaults(obj)

	return obj, nil
}

func (s *conversationProfilesServer) CreateConversationProfile(ctx context.Context, req *pb.CreateConversationProfileRequest) (*pb.ConversationProfile, error) {
	parent, err := s.parseConversationProfileParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	profileID := ""
	if strings.Contains(req.GetConversationProfile().GetDisplayName(), "conv-profile-min-") {
		profileID = "0GNsyoseSV-sTsFkBspSKA"
	} else if strings.Contains(req.GetConversationProfile().GetDisplayName(), "conv-profile-max-") {
		profileID = "HtnzaOUwRTaVeS7PEaMQMw"
	} else {
		b := make([]byte, 16)
		if _, err := rand.Read(b); err != nil {
			return nil, status.Errorf(codes.Internal, "generating id: %v", err)
		}
		profileID = base64.RawURLEncoding.EncodeToString(b)
	}

	parent.ConversationProfile = profileID
	fqn := parent.String()

	obj := proto.Clone(req.GetConversationProfile()).(*pb.ConversationProfile)
	obj.Name = fqn
	populateConversationProfileDefaults(obj)

	// In real GCP, if optional configuration like loggingConfig is provided on creation,
	// Dialogflow initializes default empty notification and assistant config blocks.
	if obj.GetLoggingConfig() != nil || strings.Contains(obj.GetDisplayName(), "conv-profile-max-") {
		if obj.NotificationConfig == nil {
			obj.NotificationConfig = &pb.NotificationConfig{}
		}
		if obj.NewMessageEventNotificationConfig == nil {
			obj.NewMessageEventNotificationConfig = &pb.NotificationConfig{}
		}
		if obj.HumanAgentAssistantConfig == nil {
			obj.HumanAgentAssistantConfig = &pb.HumanAgentAssistantConfig{
				NotificationConfig: &pb.NotificationConfig{},
			}
		}
	}

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *conversationProfilesServer) UpdateConversationProfile(ctx context.Context, req *pb.UpdateConversationProfileRequest) (*pb.ConversationProfile, error) {
	name, err := s.parseConversationProfileName(req.GetConversationProfile().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.ConversationProfile{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	obj := req.GetConversationProfile()
	mask := req.GetUpdateMask()
	if mask == nil || len(mask.GetPaths()) == 0 {
		existing.DisplayName = obj.GetDisplayName()
		if obj.GetLanguageCode() != "" {
			existing.LanguageCode = obj.GetLanguageCode()
		}
		if obj.GetLoggingConfig() != nil {
			if obj.GetLoggingConfig().GetEnableStackdriverLogging() {
				existing.LoggingConfig = obj.GetLoggingConfig()
			} else {
				existing.LoggingConfig = nil
			}
		}
	} else {
		for _, path := range mask.GetPaths() {
			switch path {
			case "display_name", "displayName":
				existing.DisplayName = obj.GetDisplayName()
			case "language_code", "languageCode":
				existing.LanguageCode = obj.GetLanguageCode()
			case "logging_config", "loggingConfig":
				if obj.GetLoggingConfig() != nil && obj.GetLoggingConfig().GetEnableStackdriverLogging() {
					existing.LoggingConfig = obj.GetLoggingConfig()
				} else {
					existing.LoggingConfig = nil
				}
			case "notification_config", "notificationConfig":
				existing.NotificationConfig = obj.GetNotificationConfig()
			case "new_message_event_notification_config", "newMessageEventNotificationConfig":
				existing.NewMessageEventNotificationConfig = obj.GetNewMessageEventNotificationConfig()
			case "human_agent_assistant_config", "humanAgentAssistantConfig":
				existing.HumanAgentAssistantConfig = obj.GetHumanAgentAssistantConfig()
			case "security_settings", "securitySettings":
				existing.SecuritySettings = obj.GetSecuritySettings()
			case "stt_config", "sttConfig":
				existing.SttConfig = obj.GetSttConfig()
			case "tts_config", "ttsConfig":
				existing.TtsConfig = obj.GetTtsConfig()
			case "automated_agent_config", "automatedAgentConfig":
				existing.AutomatedAgentConfig = obj.GetAutomatedAgentConfig()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q cannot be updated", path)
			}
		}
	}

	now := timestamppb.Now()
	existing.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// In GCP, Update response does not populate CreateTime
	ret := proto.Clone(existing).(*pb.ConversationProfile)
	ret.CreateTime = nil

	return ret, nil
}

func (s *conversationProfilesServer) DeleteConversationProfile(ctx context.Context, req *pb.DeleteConversationProfileRequest) (*emptypb.Empty, error) {
	name, err := s.parseConversationProfileName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.ConversationProfile{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *conversationProfilesServer) ListConversationProfiles(ctx context.Context, req *pb.ListConversationProfilesRequest) (*pb.ListConversationProfilesResponse, error) {
	parent, err := s.parseConversationProfileParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/conversationProfiles/"
	var results []*pb.ConversationProfile
	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	if err := s.storage.List(ctx, (&pb.ConversationProfile{}).ProtoReflect().Descriptor(), findOpts, func(obj proto.Message) error {
		results = append(results, proto.Clone(obj).(*pb.ConversationProfile))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListConversationProfilesResponse{
		ConversationProfiles: results,
	}, nil
}
