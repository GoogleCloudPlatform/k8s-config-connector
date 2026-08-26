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
			return nil, status.Errorf(codes.NotFound, "No PhraseMatcher found for name %s.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) CreatePhraseMatcher(ctx context.Context, req *pb.CreatePhraseMatcherRequest) (*pb.PhraseMatcher, error) {
	name, err := s.parsePhraseMatcherName(req.PhraseMatcher.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.PhraseMatcher).(*pb.PhraseMatcher)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.RevisionId = "rev-12345678"
	obj.RevisionCreateTime = now
	obj.ActivationUpdateTime = now
	obj.UpdateTime = now

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

	now := timestamppb.Now()
	obj.UpdateTime = now
	obj.RevisionId = "rev-87654321"

	if req.UpdateMask == nil || len(req.UpdateMask.Paths) == 0 {
		obj = proto.Clone(req.PhraseMatcher).(*pb.PhraseMatcher)
		obj.Name = fqn
		obj.UpdateTime = now
		obj.RevisionId = "rev-87654321"
	} else {
		if err := fields.UpdateByFieldMask(obj, req.PhraseMatcher, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update fields by mask: %v", err)
		}
	}

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
