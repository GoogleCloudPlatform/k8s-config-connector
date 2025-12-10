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

func (s *TagBindingsServer) normalizeParent(parent string) (string, error) {
	if suffix, ok := strings.CutPrefix(parent, "//cloudresourcemanager.googleapis.com/projects/"); ok {
		project, err := s.Projects.GetProjectByIDOrNumber(suffix)
		if err != nil {
			return "", err
		}
		projectWithNumber := fmt.Sprintf("//cloudresourcemanager.googleapis.com/projects/%d", project.Number)
		return projectWithNumber, nil
	} else {
		return parent, nil
	}
}

func (s *TagBindingsServer) CreateTagBinding(ctx context.Context, req *pb.CreateTagBindingRequest) (*lropb.Operation, error) {
	var tagValue *pb.TagValue

	// For methods that support TagValue namespaced name, only one of
	// tag_value_namespaced_name or tag_value may be filled. Requests with both
	// fields will be rejected.
	if req.GetTagBinding().GetTagValue() != "" && req.GetTagBinding().GetTagValueNamespacedName() != "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: both TagValue and TagValueNamespacedName are set")
	}

	if req.GetTagBinding().GetTagValue() != "" {
		found, err := s.tagValues.GetTagValue(ctx, &pb.GetTagValueRequest{
			Name: req.GetTagBinding().GetTagValue(),
		})
		if err != nil {
			return nil, err
		}
		tagValue = found
	} else if req.GetTagBinding().GetTagValueNamespacedName() != "" {
		found, err := s.tagValues.GetNamespacedTagValue(ctx, &pb.GetNamespacedTagValueRequest{
			Name: req.GetTagBinding().GetTagValueNamespacedName(),
		})
		if err != nil {
			return nil, err
		}
		tagValue = found
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: must specify TagValue or TagValueNamespacedName")
	}

	obj := proto.Clone(req.TagBinding).(*pb.TagBinding)
	obj.TagValue = tagValue.Name
	obj.TagValueNamespacedName = tagValue.NamespacedName

	parent := req.GetTagBinding().GetParent()
	normalizedParent, err := s.normalizeParent(parent)
	if err != nil {
		return nil, err
	}
	obj.Parent = normalizedParent

	obj.Name = fmt.Sprintf("tagBindings/%s/tagValues/%s", url.PathEscape(obj.Parent), strings.TrimPrefix(tagValue.Name, "tagValues/"))

	fqn := obj.Name
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lro, err := s.operations.DoneLRO(ctx, "", nil, obj)
	if err != nil {
		return nil, err
	}
	lro.Name = "" // Does not return LRO name
	return lro, nil
}

func (s *TagBindingsServer) DeleteTagBinding(ctx context.Context, req *pb.DeleteTagBindingRequest) (*lropb.Operation, error) {
	deleted := &pb.TagBinding{}

	name := req.GetName()
	// The `name` field is URL-encoded, but different clients do different things.
	// gcloud seems to double-encode, terraform seems to single-encode.
	// Try to unescape twice.
	if unescaped, err := url.PathUnescape(name); err == nil {
		name = unescaped
		if unescaped, err := url.PathUnescape(name); err == nil {
			name = unescaped
		}
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

	lro, err := s.operations.DoneLRO(ctx, "", nil, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	lro.Name = "" // Does not return LRO name
	return lro, nil
}

func (s *TagBindingsServer) ListTagBindings(ctx context.Context, req *pb.ListTagBindingsRequest) (*pb.ListTagBindingsResponse, error) {
	var err error
	findParent := req.GetParent()
	if strings.Contains(req.GetParent(), "//cloudresourcemanager.googleapis.com/projects/") {
		findParent, err = s.normalizeParent(req.GetParent())
		if err != nil {
			return nil, err
		}
	}

	var bindings []*pb.TagBinding

	tagBindingKind := (&pb.TagBinding{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tagBindingKind, storage.ListOptions{}, func(obj proto.Message) error {
		tagBinding := obj.(*pb.TagBinding)
		if tagBinding.Parent == findParent {
			tagBinding.TagValueNamespacedName = "" // Not returned in list

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
