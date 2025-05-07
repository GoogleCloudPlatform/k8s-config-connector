// Copyright 2023 Google LLC
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
// proto.service: google.cloud.resourcemanager.v3.TagValues
// proto.message: google.cloud.resourcemanager.v3.TagValue

package mockresourcemanager

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *TagValues) GetTagValue(ctx context.Context, req *pb.GetTagValueRequest) (*pb.TagValue, error) {
	name, err := s.parseTagValueName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TagValue{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "tagValue %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading tagValue: %v", err)
		}
	}

	// We should verify that this is part of on of our projects, but ... it's a mock

	return obj, nil
}

func (s *TagValues) GetNamespacedTagValue(ctx context.Context, req *pb.GetNamespacedTagValueRequest) (*pb.TagValue, error) {
	namespacedName := req.GetName()
	var tagValues []*pb.TagValue

	tagValueKind := (&pb.TagValue{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tagValueKind, storage.ListOptions{}, func(obj proto.Message) error {
		tagValue := obj.(*pb.TagValue)
		if tagValue.GetNamespacedName() == namespacedName {
			tagValues = append(tagValues, tagValue)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if len(tagValues) == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "Permission denied on resource '%s' (or it may not exist).", req.GetName())
	}
	if len(tagValues) > 1 {
		return nil, status.Error(codes.Internal, "found multiple matching values")
	}
	return tagValues[0], nil
}

func (s *TagValues) CreateTagValue(ctx context.Context, req *pb.CreateTagValueRequest) (*longrunningpb.Operation, error) {
	parentName, err := s.parseTagKeyName(req.GetTagValue().GetParent())
	if err != nil {
		return nil, err
	}

	parentTagKey := &pb.TagKey{}
	if err := s.storage.Get(ctx, parentName.String(), parentTagKey); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "tagKey %q not found", req.GetTagValue().GetParent())
		} else {
			return nil, status.Errorf(codes.Internal, "error reading tagKey: %v", err)
		}
	}

	namespacedName := ""

	// We should verify that this is part of on of our projects, but ... it's a mock
	tagKeyParent := parentTagKey.GetParent()
	if strings.HasPrefix(tagKeyParent, "projects/") {
		projectName, err := projects.ParseProjectName(tagKeyParent)
		if err != nil {
			return nil, err
		}
		project, err := s.projectsInternal.GetProject(projectName)
		if err != nil {
			return nil, err
		}
		namespacedName = project.ID + "/" + parentTagKey.GetShortName() + "/" + req.GetTagValue().GetShortName()
	} else if strings.HasPrefix(tagKeyParent, "organizations/") {
		// TODO: Set namespacedName: {organizationId}/{tag_key_short_name}/{tag_value_short_name}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", tagKeyParent)
	}

	if req.ValidateOnly {
		return nil, fmt.Errorf("ValidateOnly not yet implemented")
	}

	name := &tagValueName{
		ID: time.Now().UnixNano(),
	}

	fqn := name.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.TagValue).(*pb.TagValue)

	obj.CreateTime = now
	obj.UpdateTime = now
	obj.NamespacedName = namespacedName
	obj.Etag = base64.StdEncoding.EncodeToString(computeEtag(obj))
	obj.Name = fqn
	obj.Parent = parentTagKey.Name

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating tagValue: %v", err)
	}

	metadata := &pb.CreateTagValueMetadata{}
	return s.operations.StartLRO(ctx, "", metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *TagValues) UpdateTagValue(ctx context.Context, req *pb.UpdateTagValueRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetTagValue().GetName()
	name, err := s.parseTagValueName(reqName)
	if err != nil {
		return nil, err
	}

	if req.ValidateOnly {
		return nil, fmt.Errorf("ValidateOnly not yet implemented")
	}

	fqn := name.String()
	obj := &pb.TagValue{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "tagValue %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading tagValue: %v", err)
	}

	// We should verify that this is part of on of our projects, but ... it's a mock

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetTagValue().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	if len(paths) == 0 {
		obj.Description = req.GetTagValue().GetDescription()
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating tagValue: %v", err)
	}

	// Operation is not actually async
	lro, err := s.operations.DoneLRO(ctx, "", nil, obj)
	if err != nil {
		return nil, err
	}
	lro.Name = "" // Does not return name
	return lro, nil
}

func (s *TagValues) DeleteTagValue(ctx context.Context, req *pb.DeleteTagValueRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseTagValueName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TagValue{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "tagValue %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting tagValue: %v", err)
		}
	}

	// We should verify that this is part of on of our projects, but ... it's a mock

	metadata := &pb.DeleteTagValueMetadata{}
	return s.operations.StartLRO(ctx, "", metadata, func() (proto.Message, error) {
		return deleted, nil
	})
}

type tagValueName struct {
	ID int64
}

func (n *tagValueName) String() string {
	return fmt.Sprintf("tagValues/%d", n.ID)
}

// parseTagValueName parses a string into a tagValueName.
// The expected form is tagValues/<tagvalueName>
func (s *MockService) parseTagValueName(name string) (*tagValueName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "tagValues" {
		n, err := strconv.ParseInt(tokens[1], 10, 64)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (bad id)", name)
		}
		name := &tagValueName{
			ID: n,
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
