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

func deriveToolDisplayName(obj *pb.Tool) string {
	if obj.GetClientFunction() != nil {
		return obj.GetClientFunction().GetName()
	}
	if obj.GetOpenApiTool() != nil {
		return obj.GetOpenApiTool().GetName()
	}
	if obj.GetGoogleSearchTool() != nil {
		return obj.GetGoogleSearchTool().GetName()
	}
	if obj.GetConnectorTool() != nil {
		return obj.GetConnectorTool().GetName()
	}
	if obj.GetDataStoreTool() != nil {
		return obj.GetDataStoreTool().GetName()
	}
	if obj.GetPythonFunction() != nil {
		return obj.GetPythonFunction().GetName()
	}
	if obj.GetMcpTool() != nil {
		return obj.GetMcpTool().GetName()
	}
	if obj.GetFileSearchTool() != nil {
		return obj.GetFileSearchTool().GetName()
	}
	if obj.GetSystemTool() != nil {
		return obj.GetSystemTool().GetName()
	}
	if obj.GetAgentTool() != nil {
		return obj.GetAgentTool().GetName()
	}
	if obj.GetWidgetTool() != nil {
		return obj.GetWidgetTool().GetName()
	}
	return ""
}

func (s *AgentServiceV1) GetTool(ctx context.Context, req *pb.GetToolRequest) (*pb.Tool, error) {
	name, err := s.parseToolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Tool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Tool %s does not exist in the app.", name.ToolID)
		}
		return nil, err
	}
	result := proto.Clone(obj).(*pb.Tool)
	result.Etag = "abcdef0123A="
	return result, nil
}

func (s *AgentServiceV1) CreateTool(ctx context.Context, req *pb.CreateToolRequest) (*pb.Tool, error) {
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

	toolID := req.GetToolId()
	if toolID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "toolId is required")
	}

	name := &toolName{
		Project:  parent.Project,
		Location: parent.Location,
		AppID:    parent.AppID,
		ToolID:   toolID,
	}

	fqn := name.String()

	existing := &pb.Tool{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
	}

	obj := proto.Clone(req.GetTool()).(*pb.Tool)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="
	if obj.DisplayName == "" {
		obj.DisplayName = deriveToolDisplayName(obj)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Real GCP CreateTool returns the tool without etag
	result := proto.Clone(obj).(*pb.Tool)
	result.Etag = ""
	return result, nil
}

func (s *AgentServiceV1) UpdateTool(ctx context.Context, req *pb.UpdateToolRequest) (*pb.Tool, error) {
	reqTool := req.GetTool()
	if reqTool == nil {
		return nil, status.Errorf(codes.InvalidArgument, "tool is required")
	}

	name, err := s.parseToolName(reqTool.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Tool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Tool %s does not exist in the app.", name.ToolID)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Merge(obj, reqTool)
	} else {
		if err := fields.UpdateByFieldMask(obj, reqTool, paths); err != nil {
			return nil, err
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())
	if obj.DisplayName == "" {
		obj.DisplayName = deriveToolDisplayName(obj)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Real GCP UpdateTool returns etag
	result := proto.Clone(obj).(*pb.Tool)
	result.Etag = "abcdef0123A="
	return result, nil
}

func (s *AgentServiceV1) DeleteTool(ctx context.Context, req *pb.DeleteToolRequest) (*emptypb.Empty, error) {
	name, err := s.parseToolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Tool{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: Tool %s does not exist in the app.", name.ToolID)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, &pb.Tool{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AgentServiceV1) ListTools(ctx context.Context, req *pb.ListToolsRequest) (*pb.ListToolsResponse, error) {
	parent, err := s.parseAppName(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListToolsResponse{}
	findPrefix := fmt.Sprintf("projects/%s/locations/%s/apps/%s/tools/", parent.Project.ID, parent.Location, parent.AppID)

	toolKind := (&pb.Tool{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, toolKind, storage.ListOptions{}, func(obj proto.Message) error {
		item := obj.(*pb.Tool)
		if strings.HasPrefix(item.GetName(), findPrefix) {
			itemCopy := proto.Clone(item).(*pb.Tool)
			response.Tools = append(response.Tools, itemCopy)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
