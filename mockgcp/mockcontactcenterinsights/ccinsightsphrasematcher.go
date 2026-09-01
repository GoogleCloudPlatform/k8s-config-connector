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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
)

func (s *ContactCenterInsightsServer) GetPhraseMatcher(ctx context.Context, req *pb.GetPhraseMatcherRequest) (*pb.PhraseMatcher, error) {
	name, err := s.parsePhraseMatcherName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PhraseMatcher{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) CreatePhraseMatcher(ctx context.Context, req *pb.CreatePhraseMatcherRequest) (*pb.PhraseMatcher, error) {
	idStr := generatePhraseMatcherID(req.PhraseMatcher.DisplayName)
	reqName := req.Parent + "/phraseMatchers/" + idStr
	name, err := s.parsePhraseMatcherName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.PhraseMatcher).(*pb.PhraseMatcher)
	obj.Name = fqn

	if obj.DisplayName == "" {
		obj.DisplayName = idStr
	}

	now := timestamppb.Now()
	obj.ActivationUpdateTime = now
	obj.RevisionCreateTime = now
	obj.RevisionId = generateRevisionID(req.PhraseMatcher.DisplayName)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) UpdatePhraseMatcher(ctx context.Context, req *pb.UpdatePhraseMatcherRequest) (*pb.PhraseMatcher, error) {
	reqName := req.GetPhraseMatcher().GetName()
	name, err := s.parsePhraseMatcherName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PhraseMatcher{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.PhraseMatcher, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update CCInsightsPhraseMatcher fields: %v", err)
		}
	} else {
		obj = proto.Clone(req.PhraseMatcher).(*pb.PhraseMatcher)
	}

	obj.Name = fqn

	now := timestamppb.Now()
	obj.UpdateTime = now
	obj.ActivationUpdateTime = now
	obj.RevisionCreateTime = now

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) DeletePhraseMatcher(ctx context.Context, req *pb.DeletePhraseMatcherRequest) (*emptypb.Empty, error) {
	name, err := s.parsePhraseMatcherName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PhraseMatcher{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func generatePhraseMatcherID(displayName string) string {
	if displayName == "" {
		return "12568581401334663530"
	}
	return "6492803406424493961"
}

func generateRevisionID(displayName string) string {
	if displayName == "" {
		return "12568581401334663577"
	}
	return "10120555725210755114"
}
