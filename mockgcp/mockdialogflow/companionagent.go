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
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb_v2 "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
)

type companionAgentsServer struct {
	*MockService
	pb_v2.UnimplementedCompanionAgentsServer
}

type companionAgentName struct {
	Project        *projects.ProjectData
	Location       string
	CompanionAgent string
}

func (n *companionAgentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/companionAgents/%s", n.Project.ID, n.Location, n.CompanionAgent)
}

func (s *MockService) parseCompanionAgentName(name string) (*companionAgentName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "companionAgents" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &companionAgentName{
			Project:        project,
			Location:       tokens[3],
			CompanionAgent: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/companionAgents/{companionAgent}", name)
}

func (s *MockService) parseCompanionAgentParent(parent string) (*companionAgentName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &companionAgentName{
			Project:  project,
			Location: tokens[3],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}", parent)
}

func (s *companionAgentsServer) GetCompanionAgent(ctx context.Context, req *pb_v2.GetCompanionAgentRequest) (*pb_v2.CompanionAgent, error) {
	name, err := s.parseCompanionAgentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb_v2.CompanionAgent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "CompanionAgent with ID %s not found in projects/%s/location/%s/companionAgents/%s.", name.CompanionAgent, name.Project.ID, name.Location, name.CompanionAgent)
		}
		return nil, err
	}

	return obj, nil
}

func (s *companionAgentsServer) CreateCompanionAgent(ctx context.Context, req *pb_v2.CreateCompanionAgentRequest) (*pb_v2.CompanionAgent, error) {
	parent, err := s.parseCompanionAgentParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	companionAgentID := req.GetCompanionAgentId()
	if companionAgentID == "" {
		if req.GetCompanionAgent().GetName() != "" {
			name, err := s.parseCompanionAgentName(req.GetCompanionAgent().GetName())
			if err != nil {
				return nil, err
			}
			companionAgentID = name.CompanionAgent
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "companionAgentId must be provided")
		}
	}

	name := &companionAgentName{
		Project:        parent.Project,
		Location:       parent.Location,
		CompanionAgent: companionAgentID,
	}

	fqn := name.String()

	obj := proto.Clone(req.GetCompanionAgent()).(*pb_v2.CompanionAgent)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *companionAgentsServer) UpdateCompanionAgent(ctx context.Context, req *pb_v2.UpdateCompanionAgentRequest) (*pb_v2.CompanionAgent, error) {
	reqName := req.GetCompanionAgent().GetName()
	name, err := s.parseCompanionAgentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb_v2.CompanionAgent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "CompanionAgent with ID %s not found in projects/%s/location/%s/companionAgents/%s.", name.CompanionAgent, name.Project.ID, name.Location, name.CompanionAgent)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetCompanionAgent().GetDisplayName()
		obj.Description = req.GetCompanionAgent().GetDescription()
		obj.SkillConfigs = req.GetCompanionAgent().GetSkillConfigs()
		obj.CesToolSpecs = req.GetCompanionAgent().GetCesToolSpecs()
		obj.ToolsetTools = req.GetCompanionAgent().GetToolsetTools()
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetCompanionAgent().GetDisplayName()
			case "description":
				obj.Description = req.GetCompanionAgent().GetDescription()
			case "skill_configs", "skillConfigs":
				obj.SkillConfigs = req.GetCompanionAgent().GetSkillConfigs()
			case "ces_tool_specs", "cesToolSpecs":
				obj.CesToolSpecs = req.GetCompanionAgent().GetCesToolSpecs()
			case "toolset_tools", "toolsetTools":
				obj.ToolsetTools = req.GetCompanionAgent().GetToolsetTools()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *companionAgentsServer) DeleteCompanionAgent(ctx context.Context, req *pb_v2.DeleteCompanionAgentRequest) (*emptypb.Empty, error) {
	name, err := s.parseCompanionAgentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb_v2.CompanionAgent{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "CompanionAgent with ID %s not found in projects/%s/location/%s/companionAgents/%s.", name.CompanionAgent, name.Project.ID, name.Location, name.CompanionAgent)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *companionAgentsServer) ListCompanionAgents(ctx context.Context, req *pb_v2.ListCompanionAgentsRequest) (*pb_v2.ListCompanionAgentsResponse, error) {
	parent, err := s.parseCompanionAgentParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%s/locations/%s/companionAgents/", parent.Project.ID, parent.Location)

	response := &pb_v2.ListCompanionAgentsResponse{}
	companionAgentKind := (&pb_v2.CompanionAgent{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, companionAgentKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		companionAgent := obj.(*pb_v2.CompanionAgent)
		response.CompanionAgents = append(response.CompanionAgents, companionAgent)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func RegisterCompanionAgentsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := pb_v2.NewCompanionAgentsClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	for _, prefix := range []string{"/v2", "/v2beta1"} {
		// POST {prefix}/{parent=projects/*/locations/*}/companionAgents
		createPath := prefix + "/{parent=projects/*/locations/*}/companionAgents"
		if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.CreateCompanionAgentRequest{}
			protoReq.Parent = pathParams["parent"]
			protoReq.CompanionAgentId = r.URL.Query().Get("companionAgentId")

			newReader, berr := utilities.IOReaderFactory(r.Body)
			if berr != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
				return
			}
			protoReq.CompanionAgent = &pb_v2.CompanionAgent{}
			if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.CompanionAgent); err != nil && err != io.EOF {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
				return
			}

			obj, err := client.CreateCompanionAgent(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, obj, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// GET {prefix}/{name=projects/*/locations/*/companionAgents/*}
		getPath := prefix + "/{name=projects/*/locations/*/companionAgents/*}"
		if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.GetCompanionAgentRequest{
				Name: pathParams["name"],
			}

			obj, err := client.GetCompanionAgent(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, obj, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// PATCH {prefix}/{name=projects/*/locations/*/companionAgents/*}
		patchPath := prefix + "/{name=projects/*/locations/*/companionAgents/*}"
		if err := mux.HandlePath("PATCH", patchPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.UpdateCompanionAgentRequest{}
			protoReq.CompanionAgent = &pb_v2.CompanionAgent{}

			newReader, berr := utilities.IOReaderFactory(r.Body)
			if berr != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
				return
			}
			if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.CompanionAgent); err != nil && err != io.EOF {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
				return
			}
			if protoReq.CompanionAgent.Name == "" {
				protoReq.CompanionAgent.Name = pathParams["name"]
			}

			updateMask := r.URL.Query().Get("updateMask")
			if updateMask == "" {
				updateMask = r.URL.Query().Get("update_mask")
			}
			if updateMask != "" {
				protoReq.UpdateMask = &fieldmaskpb.FieldMask{Paths: strings.Split(updateMask, ",")}
			}

			obj, err := client.UpdateCompanionAgent(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, obj, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// DELETE {prefix}/{name=projects/*/locations/*/companionAgents/*}
		deletePath := prefix + "/{name=projects/*/locations/*/companionAgents/*}"
		if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.DeleteCompanionAgentRequest{
				Name: pathParams["name"],
			}

			resp, err := client.DeleteCompanionAgent(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, resp, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// GET {prefix}/{parent=projects/*/locations/*}/companionAgents
		listPath := prefix + "/{parent=projects/*/locations/*}/companionAgents"
		if err := mux.HandlePath("GET", listPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.ListCompanionAgentsRequest{
				Parent: pathParams["parent"],
			}

			resp, err := client.ListCompanionAgents(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, resp, forwardResponseOptions...)
		}); err != nil {
			return err
		}
	}

	return nil
}
