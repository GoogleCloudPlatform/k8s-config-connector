// Copyright 2025 Google LLC
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
// proto.service: google.cloud.resourcemanager.v3.TagBindings
// proto.message: google.cloud.resourcemanager.v3.TagBinding

package mockresourcemanager

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	lropb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *TagBindingsServer) normalizeParent(ctx context.Context, parent string) (string, error) {
	if suffix, ok := strings.CutPrefix(parent, "//cloudresourcemanager.googleapis.com/projects/"); ok {
		project, err := s.Projects.GetProjectByIDOrNumber(suffix)
		if err != nil {
			return "", err
		}
		projectWithNumber := fmt.Sprintf("//cloudresourcemanager.googleapis.com/projects/%d", project.Number)
		return projectWithNumber, nil
	} else {
		return "", status.Errorf(codes.InvalidArgument, "invalid parent")
	}
}

func (s *TagBindingsServer) CreateTagBinding(ctx context.Context, req *pb.CreateTagBindingRequest) (*lropb.Operation, error) {
	tagValue, err := s.tagValues.GetTagValue(ctx, &pb.GetTagValueRequest{
		Name: req.GetTagBinding().GetTagValue(),
	})
	if err != nil {
		return nil, err
	}

	// obj := &pb.TagBinding{}
	obj := proto.Clone(req.TagBinding).(*pb.TagBinding)
	obj.TagValue = tagValue.Name
	obj.TagValueNamespacedName = tagValue.NamespacedName

	parent := req.GetTagBinding().GetParent()
	obj.Parent = parent
	if strings.Contains(parent, "//cloudresourcemanager.googleapis.com/projects/") {
		normalizedParent, err := s.normalizeParent(ctx, parent)
		if err != nil {
			return nil, err
		}
		obj.Parent = normalizedParent
	}

	obj.Name = fmt.Sprintf("tagBindings/%s/tagValues/%s", url.PathEscape(obj.Parent), strings.TrimPrefix(tagValue.Name, "tagValues/"))

	fqn := obj.Name
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, "", nil, obj)
}

func (s *TagBindingsServer) DeleteTagBinding(ctx context.Context, req *pb.DeleteTagBindingRequest) (*lropb.Operation, error) {
	deleted := &pb.TagBinding{}

	name := req.GetName()

	// The name is of the form `tagBindings/{parent}/tagValues/{tag_value}`
	// The parent part of the name has been URL-decoded by the framework, so slashes are represented as `/` not `%2F`.
	// We need to parse this carefully, re-encode the parent, and then look it up.
	if !strings.HasPrefix(name, "tagBindings/") {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name, expected prefix 'tagBindings/': %q", name)
	}
	// This gives us "{parent}/tagValues/{tag_value}"
	suffix := strings.TrimPrefix(name, "tagBindings/")

	parts := strings.Split(suffix, "/tagValues/")
	if len(parts) != 2 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name, expected '/tagValues/' separator: %q", name)
	}

	parent := parts[0]
	tagValueID := parts[1]

	// We need to normalize the parent to match the key used during creation.
	// For projects, this turns the project ID into a project number.
	// For other resources, it should be a pass-through.
	if strings.Contains(parent, "//cloudresourcemanager.googleapis.com/projects/") {
		normalizedParent, err := s.normalizeParent(ctx, parent)
		if err != nil {
			return nil, err
		}
		parent = normalizedParent
	}

	// Reconstruct the name with the *escaped* normalized parent to find it in storage.
	// This is the critical step to match the format used in CreateTagBinding.
	name = fmt.Sprintf("tagBindings/%s/tagValues/%s", url.PathEscape(parent), tagValueID)

	if err := s.storage.Delete(ctx, name, deleted); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, "", nil, &emptypb.Empty{})
}

func (s *TagBindingsServer) ListTagBindings(ctx context.Context, req *pb.ListTagBindingsRequest) (*pb.ListTagBindingsResponse, error) {
	var err error
	findParent := req.GetParent()
	if strings.Contains(req.GetParent(), "//cloudresourcemanager.googleapis.com/projects/") {
		findParent, err = s.normalizeParent(ctx, req.GetParent())
		if err != nil {
			return nil, err
		}
	}

	var bindings []*pb.TagBinding

	tagBindingKind := (&pb.TagBinding{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tagBindingKind, storage.ListOptions{}, func(obj proto.Message) error {
		tagBinding := obj.(*pb.TagBinding)
		if tagBinding.Parent == findParent {
			tagBinding.TagValueNamespacedName = "" // Not returned in list?

			bindings = append(bindings, tagBinding)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListTagBindingsResponse{
		TagBindings: bindings,
	}, nil
}
