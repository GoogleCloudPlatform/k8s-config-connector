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

package mockcontactcenterinsights

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

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
)

type ContactCenterInsightsV1 struct {
	*MockService
	pb.UnimplementedContactCenterInsightsServer
}

func (s *ContactCenterInsightsV1) GetConversation(ctx context.Context, req *pb.GetConversationRequest) (*pb.Conversation, error) {
	name, err := s.parseConversationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Conversation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsV1) CreateConversation(ctx context.Context, req *pb.CreateConversationRequest) (*pb.Conversation, error) {
	conversationID := req.ConversationId
	if conversationID == "" {
		conversationID = fmt.Sprintf("conv-%d", time.Now().UnixNano())
	}
	reqName := req.Parent + "/conversations/" + conversationID
	name, err := s.parseConversationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Conversation).(*pb.Conversation)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsV1) UpdateConversation(ctx context.Context, req *pb.UpdateConversationRequest) (*pb.Conversation, error) {
	reqName := req.Conversation.GetName()
	name, err := s.parseConversationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Conversation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Update fields based on update mask or full clone if no mask
	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		for _, path := range req.UpdateMask.Paths {
			normalizedPath := strings.ToLower(strings.ReplaceAll(path, "_", ""))
			switch normalizedPath {
			case "expiretime":
				obj.Expiration = req.Conversation.Expiration
			case "ttl":
				obj.Expiration = req.Conversation.Expiration
			case "languagecode":
				obj.LanguageCode = req.Conversation.LanguageCode
			case "agentid":
				obj.AgentId = req.Conversation.AgentId
			case "labels":
				obj.Labels = req.Conversation.Labels
			case "medium":
				obj.Medium = req.Conversation.Medium
			case "obfuscateduserid":
				obj.ObfuscatedUserId = req.Conversation.ObfuscatedUserId
			}
		}
	} else {
		req.Conversation.Name = fqn
		req.Conversation.CreateTime = obj.CreateTime
		obj = proto.Clone(req.Conversation).(*pb.Conversation)
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsV1) DeleteConversation(ctx context.Context, req *pb.DeleteConversationRequest) (*emptypb.Empty, error) {
	name, err := s.parseConversationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Conversation{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
