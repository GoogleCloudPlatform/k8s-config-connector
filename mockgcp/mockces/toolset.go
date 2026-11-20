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

package mockces

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/ces/apiv1/cespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *AgentServiceV1) GetToolset(ctx context.Context, req *pb.GetToolsetRequest) (*pb.Toolset, error) {
	name, err := s.parseToolsetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Toolset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Toolset %s does not exist in the app.", name.ToolsetID)
		}
		return nil, err
	}
	result := proto.Clone(obj).(*pb.Toolset)
	result.Etag = "abcdef0123A="
	return result, nil
}

func (s *AgentServiceV1) CreateToolset(ctx context.Context, req *pb.CreateToolsetRequest) (*pb.Toolset, error) {
	parent, err := s.parseAppName(req.GetParent())
	if err != nil {
		return nil, err
	}

	parentApp := &pb.App{}
	if err := s.storage.Get(ctx, parent.String(), parentApp); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Parent resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	toolsetID := req.GetToolsetId()
	if toolsetID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "toolsetId is required")
	}

	name := &toolsetName{
		Project:   parent.Project,
		Location:  parent.Location,
		AppID:     parent.AppID,
		ToolsetID: toolsetID,
	}

	fqn := name.String()

	existing := &pb.Toolset{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
	}

	obj := proto.Clone(req.GetToolset()).(*pb.Toolset)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Real GCP CreateToolset returns the toolset without etag
	result := proto.Clone(obj).(*pb.Toolset)
	result.Etag = ""
	return result, nil
}

func (s *AgentServiceV1) UpdateToolset(ctx context.Context, req *pb.UpdateToolsetRequest) (*pb.Toolset, error) {
	reqToolset := req.GetToolset()
	if reqToolset == nil {
		return nil, status.Errorf(codes.InvalidArgument, "toolset is required")
	}

	name, err := s.parseToolsetName(reqToolset.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Toolset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Toolset %s does not exist in the app.", name.ToolsetID)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, reqToolset)
	} else {
		if err := fields.UpdateByFieldMask(obj, reqToolset, paths); err != nil {
			return nil, err
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Real GCP UpdateToolset returns etag
	result := proto.Clone(obj).(*pb.Toolset)
	result.Etag = "abcdef0123A="
	return result, nil
}

func (s *AgentServiceV1) DeleteToolset(ctx context.Context, req *pb.DeleteToolsetRequest) (*emptypb.Empty, error) {
	name, err := s.parseToolsetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Toolset{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Toolset %s does not exist in the app.", name.ToolsetID)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, &pb.Toolset{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AgentServiceV1) ListToolsets(ctx context.Context, req *pb.ListToolsetsRequest) (*pb.ListToolsetsResponse, error) {
	parent, err := s.parseAppName(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListToolsetsResponse{}
	findPrefix := fmt.Sprintf("projects/%s/locations/%s/apps/%s/toolsets/", parent.Project.ID, parent.Location, parent.AppID)

	toolsetKind := (&pb.Toolset{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, toolsetKind, storage.ListOptions{}, func(obj proto.Message) error {
		item := obj.(*pb.Toolset)
		if strings.HasPrefix(item.GetName(), findPrefix) {
			itemCopy := proto.Clone(item).(*pb.Toolset)
			response.Toolsets = append(response.Toolsets, itemCopy)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
