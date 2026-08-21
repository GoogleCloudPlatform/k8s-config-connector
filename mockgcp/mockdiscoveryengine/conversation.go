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
// proto.service: google.cloud.discoveryengine.v1.ConversationalSearchService
// proto.message: google.cloud.discoveryengine.v1.Conversation

package mockdiscoveryengine

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

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *conversationalSearchService) CreateConversation(ctx context.Context, req *pb.CreateConversationRequest) (*pb.Conversation, error) {
	name, err := s.parseConversationName(req.GetConversation().GetName())
	if err != nil {
		return nil, err
	}
	now := time.Now()

	fqn := name.String()
	obj := proto.Clone(req.GetConversation()).(*pb.Conversation)
	obj.Name = fqn
	if obj.State == pb.Conversation_STATE_UNSPECIFIED {
		obj.State = pb.Conversation_IN_PROGRESS
	}
	if obj.UserPseudoId == "" {
		obj.UserPseudoId = name.Conversation
	}
	obj.StartTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *conversationalSearchService) DeleteConversation(ctx context.Context, req *pb.DeleteConversationRequest) (*emptypb.Empty, error) {
	name, err := s.parseConversationName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Conversation{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *conversationalSearchService) UpdateConversation(ctx context.Context, req *pb.UpdateConversationRequest) (*pb.Conversation, error) {
	name, err := s.parseConversationName(req.GetConversation().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Conversation{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Conversation %q not found", name)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) > 0 {
		for _, path := range paths {
			switch path {
			case "state":
				obj.State = req.GetConversation().GetState()
			case "user_pseudo_id", "userPseudoId":
				obj.UserPseudoId = req.GetConversation().GetUserPseudoId()
			case "messages":
				obj.Messages = req.GetConversation().GetMessages()
			}
		}
	} else {
		// Simple merge of updated fields
		proto.Merge(obj, req.GetConversation())
	}
	obj.Name = fqn

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *conversationalSearchService) GetConversation(ctx context.Context, req *pb.GetConversationRequest) (*pb.Conversation, error) {
	name, err := s.parseConversationName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Conversation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.Internal, "Internal error encountered. Please try again. If the issue persists, please contact our support team.")
		}
		return nil, err
	}
	return obj, nil
}

type conversationName struct {
	Project      *projects.ProjectData
	Location     string
	Collection   string
	DataStore    string
	Conversation string
}

func (n *conversationName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/conversations/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.Conversation)
}

func (s *MockService) parseConversationName(name string) (*conversationName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "conversations" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &conversationName{
			Project:      project,
			Location:     tokens[3],
			Collection:   tokens[5],
			DataStore:    tokens[7],
			Conversation: tokens[9],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid conversation name: %q", name)
}
