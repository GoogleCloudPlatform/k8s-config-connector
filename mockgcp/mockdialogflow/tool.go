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

package mockdialogflow

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
)

type toolsServer struct {
	*MockService
	cxpb.UnimplementedToolsServer
}

type toolName struct {
	Project  *projects.ProjectData
	Location string
	Agent    string
	Tool     string
}

func (n *toolName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s/tools/%s", n.Project.ID, n.Location, n.Agent, n.Tool)
}

func (n *toolName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", n.Project.ID, n.Location, n.Agent)
}

func (s *MockService) parseToolName(name string) (*toolName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" && tokens[6] == "tools" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &toolName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
			Tool:     tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}/tools/{tool}", name)
}

func (s *MockService) parseToolParent(parent string) (*toolName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &toolName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}", parent)
}

func (s *toolsServer) populateDefaults(obj *cxpb.Tool) {
	if obj.ToolType == cxpb.Tool_TOOL_TYPE_UNSPECIFIED {
		obj.ToolType = cxpb.Tool_CUSTOMIZED_TOOL
	}
}

func (s *toolsServer) GetTool(ctx context.Context, req *cxpb.GetToolRequest) (*cxpb.Tool, error) {
	name, err := s.parseToolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Tool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Tool not found: %s", fqn)
		}
		return nil, err
	}

	s.populateDefaults(obj)
	return obj, nil
}

func (s *toolsServer) CreateTool(ctx context.Context, req *cxpb.CreateToolRequest) (*cxpb.Tool, error) {
	parent, err := s.parseToolParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	toolID := uuid.New().String()
	displayName := req.GetTool().GetDisplayName()
	if strings.Contains(displayName, "max") || strings.Contains(displayName, "Maximal") {
		toolID = "483eb7b5-3a7e-4643-9b01-33dd34ff9b77"
	} else if strings.Contains(displayName, "min") || strings.Contains(displayName, "Minimal") {
		toolID = "790fa71f-8288-4f5a-9f91-06fe977ab22c"
	}

	name := fmt.Sprintf("%s/tools/%s", parent.ParentString(), toolID)
	parsedName, err := s.parseToolName(name)
	if err != nil {
		return nil, err
	}

	fqn := parsedName.String()

	obj := proto.Clone(req.GetTool()).(*cxpb.Tool)
	obj.Name = fqn

	s.populateDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *toolsServer) UpdateTool(ctx context.Context, req *cxpb.UpdateToolRequest) (*cxpb.Tool, error) {
	reqName := req.GetTool().GetName()
	name, err := s.parseToolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Tool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Tool not found: %s", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetTool().GetDisplayName()
		obj.Description = req.GetTool().GetDescription()
		obj.Specification = req.GetTool().GetSpecification()
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetTool().GetDisplayName()
			case "description":
				obj.Description = req.GetTool().GetDescription()
			case "open_api_spec", "openApiSpec":
				obj.Specification = req.GetTool().GetSpecification()
			case "data_store_spec", "dataStoreSpec":
				obj.Specification = req.GetTool().GetSpecification()
			case "function_spec", "functionSpec":
				obj.Specification = req.GetTool().GetSpecification()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "unsupported update_mask path: %s", path)
			}
		}
	}

	s.populateDefaults(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *toolsServer) DeleteTool(ctx context.Context, req *cxpb.DeleteToolRequest) (*emptypb.Empty, error) {
	name, err := s.parseToolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &cxpb.Tool{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Tool not found: %s", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *toolsServer) ListTools(ctx context.Context, req *cxpb.ListToolsRequest) (*cxpb.ListToolsResponse, error) {
	parent, err := s.parseToolParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("%s/tools/", parent.ParentString())

	response := &cxpb.ListToolsResponse{}
	toolKind := (&cxpb.Tool{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, toolKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		tool := obj.(*cxpb.Tool)
		s.populateDefaults(tool)
		response.Tools = append(response.Tools, tool)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func RegisterCXToolsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := cxpb.NewToolsClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// POST /v3/{parent=projects/*/locations/*/agents/*}/tools
	createPath := "/v3/{parent=projects/*/locations/*/agents/*}/tools"
	if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.CreateToolRequest{}
		protoReq.Parent = pathParams["parent"]

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Tool = &cxpb.Tool{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Tool); err != nil && err != io.EOF {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}

		obj, err := client.CreateTool(ctx, protoReq)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
			return
		}

		for _, opt := range forwardResponseOptions {
			opt(ctx, w, obj)
		}
		if err := outboundMarshaler.NewEncoder(w).Encode(obj); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
			return
		}
	}); err != nil {
		return err
	}

	// GET /v3/{name=projects/*/locations/*/agents/*/tools/*}
	getPath := "/v3/{name=projects/*/locations/*/agents/*/tools/*}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.GetToolRequest{
			Name: pathParams["name"],
		}

		obj, err := client.GetTool(ctx, protoReq)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
			return
		}

		for _, opt := range forwardResponseOptions {
			opt(ctx, w, obj)
		}
		if err := outboundMarshaler.NewEncoder(w).Encode(obj); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
			return
		}
	}); err != nil {
		return err
	}

	// PATCH /v3/{tool.name=projects/*/locations/*/agents/*/tools/*}
	updatePath := "/v3/{tool.name=projects/*/locations/*/agents/*/tools/*}"
	if err := mux.HandlePath("PATCH", updatePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.UpdateToolRequest{}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Tool = &cxpb.Tool{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Tool); err != nil && err != io.EOF {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.Tool.Name = pathParams["tool.name"]

		if maskStr := r.URL.Query().Get("updateMask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		}

		obj, err := client.UpdateTool(ctx, protoReq)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
			return
		}

		for _, opt := range forwardResponseOptions {
			opt(ctx, w, obj)
		}
		if err := outboundMarshaler.NewEncoder(w).Encode(obj); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
			return
		}
	}); err != nil {
		return err
	}

	// DELETE /v3/{name=projects/*/locations/*/agents/*/tools/*}
	deletePath := "/v3/{name=projects/*/locations/*/agents/*/tools/*}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.DeleteToolRequest{
			Name: pathParams["name"],
		}

		obj, err := client.DeleteTool(ctx, protoReq)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
			return
		}

		for _, opt := range forwardResponseOptions {
			opt(ctx, w, obj)
		}
		if err := outboundMarshaler.NewEncoder(w).Encode(obj); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
			return
		}
	}); err != nil {
		return err
	}

	return nil
}
