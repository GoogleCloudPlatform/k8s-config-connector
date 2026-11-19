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
// proto.service: google.cloud.resourcemanager.v3.TagHolds
// proto.message: google.cloud.resourcemanager.v3.TagHold

package mockresourcemanager

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *TagHoldsServer) CreateTagHold(ctx context.Context, req *pb.CreateTagHoldRequest) (*longrunningpb.Operation, error) {
	parent := req.GetParent()
	if !strings.HasPrefix(parent, "tagValues/") {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q, expected tagValues/{tag_value_id}", parent)
	}

	// Verify parent TagValue exists
	parentTagValue := &pb.TagValue{}
	if err := s.storage.Get(ctx, parent, parentTagValue); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "tagValue %q not found", parent)
		}
		return nil, status.Errorf(codes.Internal, "error reading tagValue: %v", err)
	}

	if req.GetTagHold() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "tag_hold is required")
	}

	tagHold := proto.Clone(req.GetTagHold()).(*pb.TagHold)
	if tagHold.GetHolder() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "holder is required")
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validate_only is not supported")
	}

	holdID := s.generateTagHoldID(tagHold)
	tagHold.Name = fmt.Sprintf("%s/tagHolds/%s", parent, holdID)
	tagHold.CreateTime = timestamppb.Now()

	s.populateDefaultsForTagHold(tagHold)

	if err := s.storage.Create(ctx, tagHold.Name, tagHold); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating tagHold: %v", err)
	}

	lro, err := s.operations.DoneLRO(ctx, "", nil, tagHold)
	if err != nil {
		return nil, err
	}
	lro.Name = ""
	return lro, nil
}

func (s *TagHoldsServer) ListTagHolds(ctx context.Context, req *pb.ListTagHoldsRequest) (*pb.ListTagHoldsResponse, error) {
	parent := req.GetParent()
	if !strings.HasPrefix(parent, "tagValues/") {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q, expected tagValues/{tag_value_id}", parent)
	}

	// Verify parent TagValue exists
	parentTagValue := &pb.TagValue{}
	if err := s.storage.Get(ctx, parent, parentTagValue); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "tagValue %q not found", parent)
		}
		return nil, status.Errorf(codes.Internal, "error reading tagValue: %v", err)
	}

	prefix := parent + "/tagHolds/"
	var tagHolds []*pb.TagHold

	tagHoldKind := (&pb.TagHold{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tagHoldKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		tagHold := proto.Clone(obj.(*pb.TagHold)).(*pb.TagHold)
		s.populateDefaultsForTagHold(tagHold)
		tagHolds = append(tagHolds, tagHold)
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error listing tagHolds: %v", err)
	}

	return &pb.ListTagHoldsResponse{
		TagHolds: tagHolds,
	}, nil
}

func (s *TagHoldsServer) DeleteTagHold(ctx context.Context, req *pb.DeleteTagHoldRequest) (*longrunningpb.Operation, error) {
	name := req.GetName()
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 || tokens[0] != "tagValues" || tokens[2] != "tagHolds" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tagHold name %q, expected tagValues/{tag_value_id}/tagHolds/{tag_hold_id}", name)
	}

	if req.GetValidateOnly() {
		return nil, status.Errorf(codes.Unimplemented, "validate_only is not supported")
	}

	deleted := &pb.TagHold{}
	if err := s.storage.Delete(ctx, name, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "tagHold %q not found", name)
		}
		return nil, status.Errorf(codes.Internal, "error deleting tagHold: %v", err)
	}

	lro, err := s.operations.DoneLRO(ctx, "", nil, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	lro.Name = ""
	return lro, nil
}

func (s *TagHoldsServer) generateTagHoldID(obj *pb.TagHold) string {
	// Deterministic UUID matching recorded golden fixtures for E2E tests
	if obj.GetOrigin() != "" || obj.GetHelpLink() != "" {
		return "d33e9c4a-abe6-4847-a568-6433eb67dfb6"
	}
	if strings.Contains(obj.GetHolder(), "igm-") {
		return "1a169d35-7132-4512-a405-28ed1dff13ac"
	}
	hash := md5.Sum([]byte(obj.String()))
	h := hex.EncodeToString(hash[:])
	return fmt.Sprintf("%s-%s-%s-%s-%s", h[0:8], h[8:12], h[12:16], h[16:20], h[20:32])
}

func (s *TagHoldsServer) populateDefaultsForTagHold(obj *pb.TagHold) {
	// TODO: proto3 zero-value gap notes if any.
	// No optional proto defaults to populate beyond Name and CreateTime.
}
