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

package mockgeminidataanalytics

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/geminidataanalytics/apiv1beta/geminidataanalyticspb"
)

type GeminiDataAnalyticsV1beta struct {
	*MockService
	pb.UnimplementedDataChatServiceServer
}

func (s *GeminiDataAnalyticsV1beta) GetConversation(ctx context.Context, req *pb.GetConversationRequest) (*pb.Conversation, error) {
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

func (s *GeminiDataAnalyticsV1beta) CreateConversation(ctx context.Context, req *pb.CreateConversationRequest) (*pb.Conversation, error) {
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

	obj := proto.CloneOf(req.Conversation)
	obj.Name = fqn

	// Populate output-only fields
	now := timestamppb.Now()
	obj.CreateTime = now
	obj.LastUsedTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GeminiDataAnalyticsV1beta) DeleteConversation(ctx context.Context, req *pb.DeleteConversationRequest) (*emptypb.Empty, error) {
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
