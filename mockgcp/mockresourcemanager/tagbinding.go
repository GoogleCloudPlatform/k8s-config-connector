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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *TagBindingsServer) normalizeParent(parent string) (string, error) {
	tokens := strings.Split(parent, "/projects/")

	if len(tokens) != 2 {
		// Not a project parent, return as-is.
		return parent, nil
	}
	pTokens := strings.Split(tokens[1], "/")
	if len(pTokens) < 1 {
		return parent, nil
	}
	projectIdentifier := pTokens[0]
	project, err := s.Projects.GetProjectByIDOrNumber(projectIdentifier)
	if err != nil {
		return parent, nil
	}
	return strings.Replace(parent, pTokens[0], fmt.Sprintf("%d", project.Number), 1), nil
}

func (s *TagBindingsServer) CreateTagBinding(ctx context.Context, req *pb.CreateTagBindingRequest) (*lropb.Operation, error) {
	tagValue, err := s.tagValues.GetTagValue(ctx, &pb.GetTagValueRequest{
		Name: req.GetTagBinding().GetTagValue(),
	})
	if err != nil {
		return nil, err
	}

	obj := proto.Clone(req.TagBinding).(*pb.TagBinding)
	obj.TagValue = tagValue.Name
	obj.TagValueNamespacedName = tagValue.NamespacedName

	normalizedParent, err := s.normalizeParent(req.GetTagBinding().GetParent())
	if err != nil {
		return nil, err
	}
	obj.Parent = normalizedParent

	// The name in the response should use the original parent string (with project ID).
	obj.Name = fmt.Sprintf("tagBindings/%s/tagValues/%s", url.PathEscape(normalizedParent), strings.TrimPrefix(tagValue.Name, "tagValues/"))

	// The internal storage key should use the normalized parent (with project number).
	fqn := fmt.Sprintf("tagBindings/%s/tagValues/%s", url.PathEscape(normalizedParent), strings.TrimPrefix(tagValue.Name, "tagValues/"))
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, "", nil, obj)
}

func (s *TagBindingsServer) DeleteTagBinding(ctx context.Context, req *pb.DeleteTagBindingRequest) (*lropb.Operation, error) {
	deleted := &pb.TagBinding{}

	name := req.GetName()
	// The `name` field is URL-encoded, but different clients do different things.
	// gcloud seems to double-encode, terraform seems to single-encode.
	// Try to unescape twice.
	if unescaped, err := url.PathUnescape(name); err == nil {
		name = unescaped
	}

	name = strings.TrimPrefix(name, "tagBindings/")
	tokens := strings.Split(name, "/tagValues/")
	if len(tokens) == 2 {
		normalizedParent, err := s.normalizeParent(tokens[0])
		if err != nil {
			return nil, err
		}
		name = "tagBindings/" + url.PathEscape(normalizedParent) + "/tagValues/" + tokens[1]
	}

	if err := s.storage.Delete(ctx, name, deleted); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, "", nil, &emptypb.Empty{})
}

func (s *TagBindingsServer) ListTagBindings(ctx context.Context, req *pb.ListTagBindingsRequest) (*pb.ListTagBindingsResponse, error) {
	findParent, err := s.normalizeParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	var bindings []*pb.TagBinding

	tagBindingKind := (&pb.TagBinding{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tagBindingKind, storage.ListOptions{}, func(obj proto.Message) error {
		tagBinding := obj.(*pb.TagBinding)
		if tagBinding.Parent == findParent {
			clone := proto.Clone(tagBinding).(*pb.TagBinding)
			clone.TagValueNamespacedName = "" // Not returned in list
			bindings = append(bindings, clone)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListTagBindingsResponse{
		TagBindings: bindings,
	}, nil
}
