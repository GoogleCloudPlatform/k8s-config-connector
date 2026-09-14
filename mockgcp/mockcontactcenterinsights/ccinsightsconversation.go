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
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
)

func (s *ContactCenterInsightsServer) GetConversation(ctx context.Context, req *pb.GetConversationRequest) (*pb.Conversation, error) {
	name, err := s.parseConversationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Conversation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No conversation found for project %d and conversation ID %s.", name.Project.Number, name.Conversation)
		}
		return nil, err
	}

	populateConversationDefaults(obj)

	return obj, nil
}

func (s *ContactCenterInsightsServer) CreateConversation(ctx context.Context, req *pb.CreateConversationRequest) (*pb.Conversation, error) {
	reqName := req.Parent + "/conversations/" + req.ConversationId
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

	populateConversationDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) UpdateConversation(ctx context.Context, req *pb.UpdateConversationRequest) (*pb.Conversation, error) {
	reqName := req.GetConversation().GetName()
	name, err := s.parseConversationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Conversation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No conversation found for project %d and conversation ID %s.", name.Project.Number, name.Conversation)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if obj.Labels == nil {
			obj.Labels = make(map[string]string)
		}
		if err := fields.UpdateByFieldMask(obj, req.Conversation, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update CCInsightsConversation fields: %v", err)
		}
	} else {
		obj = proto.Clone(req.Conversation).(*pb.Conversation)
	}

	obj.Name = fqn

	now := timestamppb.Now()
	obj.UpdateTime = now

	populateConversationDefaults(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) DeleteConversation(ctx context.Context, req *pb.DeleteConversationRequest) (*emptypb.Empty, error) {
	name, err := s.parseConversationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Conversation{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No conversation found for project %d and conversation ID %s.", name.Project.Number, name.Conversation)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func populateConversationDefaults(obj *pb.Conversation) {
	if obj.StartTime == nil {
		obj.StartTime = timestamppb.Now()
	}
	obj.Duration = &durationpb.Duration{Seconds: 4}
	obj.TurnCount = 2

	if obj.Transcript == nil {
		obj.Transcript = &pb.Conversation_Transcript{
			TranscriptSegments: []*pb.Conversation_Transcript_TranscriptSegment{
				{
					ChannelTag:   1,
					LanguageCode: "en-US",
					MessageTime:  timestamppb.New(time.Unix(1, 0)),
					SegmentParticipant: &pb.ConversationParticipant{
						Role: pb.ConversationParticipant_Role(3),
						Participant: &pb.ConversationParticipant_UserId{
							UserId: "1",
						},
					},
					Text: "Hello, I am calling about my recent invoice.",
					Words: []*pb.Conversation_Transcript_TranscriptSegment_WordInfo{
						{Word: "Hello,"},
						{Word: "I"},
						{Word: "am"},
						{Word: "calling"},
						{Word: "about"},
						{Word: "my"},
						{Word: "recent"},
						{Word: "invoice."},
					},
				},
				{
					ChannelTag:   2,
					LanguageCode: "en-US",
					MessageTime:  timestamppb.New(time.Unix(5, 0)),
					SegmentParticipant: &pb.ConversationParticipant{
						Role: pb.ConversationParticipant_Role(1),
						Participant: &pb.ConversationParticipant_UserId{
							UserId: "2",
						},
					},
					Text: "I can certainly help with that. Can I have your account number?",
					Words: []*pb.Conversation_Transcript_TranscriptSegment_WordInfo{
						{Word: "I"},
						{Word: "can"},
						{Word: "certainly"},
						{Word: "help"},
						{Word: "with"},
						{Word: "that."},
						{Word: "Can"},
						{Word: "I"},
						{Word: "have"},
						{Word: "your"},
						{Word: "account"},
						{Word: "number?"},
					},
				},
			},
		}
	}

	if obj.AgentId == "" {
		obj.AgentId = "2"
	}
	if obj.LanguageCode == "" {
		obj.LanguageCode = "en-US"
	}
	if obj.QualityMetadata == nil {
		obj.QualityMetadata = &pb.Conversation_QualityMetadata{
			AgentInfo: []*pb.Conversation_QualityMetadata_AgentInfo{
				{
					AgentId:   "2",
					AgentType: pb.ConversationParticipant_Role(1),
				},
			},
		}
	}

	// Real GCP does not return ObfuscatedUserId in responses.
	obj.ObfuscatedUserId = ""
}
