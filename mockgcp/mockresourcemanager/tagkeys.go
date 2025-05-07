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
// proto.service: google.cloud.resourcemanager.v3.TagKeys
// proto.message: google.cloud.resourcemanager.v3.TagKey

package mockresourcemanager

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *TagKeys) GetTagKey(ctx context.Context, req *pb.GetTagKeyRequest) (*pb.TagKey, error) {
	name, err := s.parseTagKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TagKey{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// We should verify that this is part of on of our projects, but ... it's a mock

	return obj, nil
}

func (s *TagKeys) ListTagKeys(ctx context.Context, req *pb.ListTagKeysRequest) (*pb.ListTagKeysResponse, error) {

	findParent := ""
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByIDOrNumber(req.Parent)
		if err != nil {
			return nil, err
		}

		findParent = fmt.Sprintf("projects/%d", project.Number)
	} else {
		return nil, fmt.Errorf("parent %q is not valid for mock", req.GetParent())
	}

	var tagKeys []*pb.TagKey

	tagKeyKind := (&pb.TagKey{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tagKeyKind, storage.ListOptions{}, func(obj proto.Message) error {
		tagKey := obj.(*pb.TagKey)
		if tagKey.Parent == findParent {
			tagKeys = append(tagKeys, tagKey)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListTagKeysResponse{
		TagKeys: tagKeys,
	}, nil
}

func (s *TagKeys) GetNamespacedTagKey(ctx context.Context, req *pb.GetNamespacedTagKeyRequest) (*pb.TagKey, error) {
	namespacedName := req.GetName()
	var tagKeys []*pb.TagKey

	tagKeyKind := (&pb.TagKey{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tagKeyKind, storage.ListOptions{}, func(obj proto.Message) error {
		tagKey := obj.(*pb.TagKey)
		if tagKey.GetNamespacedName() == namespacedName {
			tagKeys = append(tagKeys, tagKey)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if len(tagKeys) == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "Permission denied on resource '%s' (or it may not exist).", req.GetName())
	}
	if len(tagKeys) > 1 {
		return nil, status.Error(codes.Internal, "found multiple matching keys")
	}
	return tagKeys[0], nil
}

func (s *TagKeys) CreateTagKey(ctx context.Context, req *pb.CreateTagKeyRequest) (*longrunningpb.Operation, error) {
	var namespacedName string

	parent := req.GetTagKey().GetParent()
	if strings.HasPrefix(parent, "projects/") {
		projectName, err := projects.ParseProjectName(parent)
		if err != nil {
			return nil, err
		}
		project, err := s.projectsInternal.GetProject(projectName)
		if err != nil {
			return nil, err
		}
		namespacedName = project.ID + "/" + req.GetTagKey().GetShortName()
		// Parent is normalized to the project number
		req.GetTagKey().Parent = fmt.Sprintf("projects/%d", project.Number)
	} else if strings.HasPrefix(parent, "organizations/") {
		// We should check that the org exists, permissions etc, but ... it's a mock
		namespacedName = strings.TrimPrefix(parent, "organizations/") + "/" + req.GetTagKey().GetShortName()
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}

	if req.ValidateOnly {
		return nil, fmt.Errorf("ValidateOnly not yet implemented")
	}

	name := &tagKeyName{
		ID: time.Now().UnixNano(),
	}

	fqn := name.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.TagKey).(*pb.TagKey)

	obj.CreateTime = now
	obj.UpdateTime = now
	obj.NamespacedName = namespacedName
	obj.Etag = base64.StdEncoding.EncodeToString(computeEtag(obj))
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreateTagKeyMetadata{}
	return s.operations.StartLRO(ctx, "", metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *TagKeys) UpdateTagKey(ctx context.Context, req *pb.UpdateTagKeyRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetTagKey().GetName()
	name, err := s.parseTagKeyName(reqName)
	if err != nil {
		return nil, err
	}

	if req.ValidateOnly {
		return nil, fmt.Errorf("ValidateOnly not yet implemented")
	}

	fqn := name.String()
	obj := &pb.TagKey{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// We should verify that this is part of on of our projects, but ... it's a mock

	// Fields to be updated. The mask may only contain `description` or
	// `etag`. If omitted entirely, both `description` and `etag` are assumed to
	// be significant.
	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetTagKey().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	if len(paths) == 0 {
		obj.Description = req.GetTagKey().GetDescription()
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// LRO is immediately done
	return s.operations.DoneLRO(ctx, "", nil, obj)
}

func (s *TagKeys) DeleteTagKey(ctx context.Context, req *pb.DeleteTagKeyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseTagKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// We should verify that this is part of on of our projects, but ... it's a mock
	metadata := &pb.DeleteTagKeyMetadata{}
	return s.operations.StartLRO(ctx, "", metadata, func() (proto.Message, error) {
		// TagKey must not have any child TagValues
		tagValueKind := (&pb.TagValue{}).ProtoReflect().Descriptor()
		if err := s.storage.List(ctx, tagValueKind, storage.ListOptions{}, func(obj proto.Message) error {
			tagValue := obj.(*pb.TagValue)
			if tagValue.GetParent() == fqn {
				return status.Errorf(codes.FailedPrecondition,
					"TagKey: %s has child TagValues. Please list all TagValues under this key and delete them before retrying TagKey deletion.",
					fqn)
			}
			return nil
		}); err != nil {
			return nil, err
		}

		deleted := &pb.TagKey{}
		if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
			return nil, err
		}

		return deleted, nil
	})
}

type tagKeyName struct {
	ID int64
}

func (n *tagKeyName) String() string {
	return fmt.Sprintf("tagKeys/%d", n.ID)
}

// parseTagKeyName parses a string into a tagKeyName.
// The expected form is tagKeys/<tagkeyName>
func (s *MockService) parseTagKeyName(name string) (*tagKeyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "tagKeys" {

		n, err := strconv.ParseInt(tokens[1], 10, 64)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (bad id)", name)
		}
		name := &tagKeyName{
			ID: n,
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func computeEtag(obj proto.Message) []byte {
	// TODO: Do we risk exposing internal fields?  Doesn't matter on a mock, I guess
	b, err := proto.Marshal(obj)
	if err != nil {
		klog.Fatalf("failed to marshal proto object: %v", err)
	}
	hash := md5.Sum(b)
	return hash[:]
}
