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

package mockresourcemanager

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type TagBindings struct {
	*MockService
	pb.UnimplementedTagBindingsServer
}

func (s *TagBindings) ListTagBindings(ctx context.Context, req *pb.ListTagBindingsRequest) (*pb.ListTagBindingsResponse, error) {
	if req.GetParent() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}
	parent, err := s.normalizeParent(ctx, req.GetParent())
	if err != nil {
		return nil, err
	}

	var matches []*pb.TagBinding

	// We should verify that this is part of one of our projects, but ... it's a mock

	tagBindingKind := (&pb.TagBinding{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tagBindingKind, storage.ListOptions{}, func(obj proto.Message) error {
		tagBinding := obj.(*pb.TagBinding)
		klog.Infof("found=%+v", tagBinding)
		if tagBinding.GetParent() == parent {
			// NamespacedName is not returned, apparently
			matches = append(matches, &pb.TagBinding{
				Name:     tagBinding.Name,
				Parent:   tagBinding.Parent,
				TagValue: tagBinding.TagValue,
			})
		}
		return nil
	}); err != nil {
		return nil, err
	}

	klog.Infof("req=%+v, matches=%+v", req, matches)

	return &pb.ListTagBindingsResponse{TagBindings: matches}, nil
}

func (s *TagBindings) normalizeParent(ctx context.Context, parent string) (string, error) {
	if strings.HasPrefix(parent, "//cloudresourcemanager.googleapis.com/projects/") {
		projectName, err := projects.ParseProjectName(strings.TrimPrefix(parent, "//cloudresourcemanager.googleapis.com/"))
		if err != nil {
			return "", err
		}

		project, err := s.projectsInternal.GetProject(projectName)
		if err != nil {
			return "", err
		}
		// This is normalized to a project number
		parent = fmt.Sprintf("//cloudresourcemanager.googleapis.com/projects/%d", project.Number)
	} else {
		return "", status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}
	return parent, nil
}

func (s *TagBindings) CreateTagBinding(ctx context.Context, req *pb.CreateTagBindingRequest) (*longrunningpb.Operation, error) {
	parent, err := s.normalizeParent(ctx, req.GetTagBinding().GetParent())
	if err != nil {
		return nil, err
	}
	req.TagBinding.Parent = parent

	tagValueName, err := s.parseTagValueName(req.GetTagBinding().GetTagValue())
	if err != nil {
		return nil, err
	}

	tagValue := &pb.TagValue{}
	if err := s.storage.Get(ctx, tagValueName.String(), tagValue); err != nil {
		return nil, err
	}

	req.TagBinding.TagValueNamespacedName = tagValue.NamespacedName

	if req.ValidateOnly {
		return nil, fmt.Errorf("ValidateOnly not yet implemented")
	}

	name := &tagBindingName{
		TagValueID: tagValueName.ID,
		Parent:     parent,
	}
	req.TagBinding.Name = name.String()

	fqn := name.String()

	obj := proto.Clone(req.TagBinding).(*pb.TagBinding)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op, err := s.operations.DoneLRO(ctx, "", nil, obj)
	if err != nil {
		return nil, err
	}
	op.Name = "" // Name is empty for some reason on this API
	return op, nil
}

func (s *TagBindings) DeleteTagBinding(ctx context.Context, req *pb.DeleteTagBindingRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseTagBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// We should verify that this is part of one of our projects, but ... it's a mock
	deleted := &pb.TagBinding{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op, err := s.operations.DoneLRO(ctx, "", nil, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	op.Name = "" // Name is empty for some reason on this API
	return op, nil
}

type tagBindingName struct {
	Parent     string
	TagValueID int64
}

func (n *tagBindingName) String() string {
	return fmt.Sprintf("tagBindings/%s/tagValues/%d", url.PathEscape(n.Parent), n.TagValueID)
}

// parseTagBindingName parses a string into a tagBindingName.
// The expected form is tagBindings/<parent>/tagValues/<tagValueId>
func (s *MockService) parseTagBindingName(name string) (*tagBindingName, error) {
	tokens := strings.Split(name, "/")

	klog.Infof("%q => tokens %+v", name, tokens)
	if len(tokens) == 4 && tokens[0] == "tagBindings" && tokens[2] == "tagValues" {
		tagValueID, err := strconv.ParseInt(tokens[3], 10, 64)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (bad tagValueID)", name)
		}
		parent, err := url.PathUnescape(tokens[1])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (bad parent)", name)
		}
		name := &tagBindingName{
			Parent:     parent,
			TagValueID: tagValueID,
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
