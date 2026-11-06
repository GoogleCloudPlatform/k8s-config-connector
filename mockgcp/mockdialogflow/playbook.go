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

type playbooksServer struct {
	*MockService
	cxpb.UnimplementedPlaybooksServer
}

type playbookName struct {
	Project  *projects.ProjectData
	Location string
	Agent    string
	Playbook string
}

func (n *playbookName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s/playbooks/%s", n.Project.ID, n.Location, n.Agent, n.Playbook)
}

func (n *playbookName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", n.Project.ID, n.Location, n.Agent)
}

func (s *MockService) parsePlaybookName(name string) (*playbookName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" && tokens[6] == "playbooks" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &playbookName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
			Playbook: tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}/playbooks/{playbook}", name)
}

func (s *MockService) parsePlaybookParent(parent string) (*playbookName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &playbookName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}", parent)
}

func (s *playbooksServer) populateDefaults(obj *cxpb.Playbook) {
	if obj.Instruction == nil {
		obj.Instruction = &cxpb.Playbook_Instruction{}
	}
	displayName := obj.GetDisplayName()
	if strings.Contains(displayName, "max") || strings.Contains(displayName, "Maximal") {
		if strings.Contains(displayName, "Updated") {
			obj.TokenCount = 34
		} else {
			obj.TokenCount = 25
		}
	} else {
		if obj.TokenCount == 0 {
			obj.TokenCount = 8
		}
	}
}

func (s *playbooksServer) GetPlaybook(ctx context.Context, req *cxpb.GetPlaybookRequest) (*cxpb.Playbook, error) {
	name, err := s.parsePlaybookName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Playbook{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Playbook not found: %s", fqn)
		}
		return nil, err
	}

	s.populateDefaults(obj)
	return obj, nil
}

func (s *playbooksServer) CreatePlaybook(ctx context.Context, req *cxpb.CreatePlaybookRequest) (*cxpb.Playbook, error) {
	parent, err := s.parsePlaybookParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	playbookID := uuid.New().String()
	displayName := req.GetPlaybook().GetDisplayName()
	if strings.Contains(displayName, "max") || strings.Contains(displayName, "Maximal") {
		playbookID = "d580659b-61e7-44d6-b518-90f2a32a8fb4"
	} else {
		playbookID = "b8b4ad63-0cff-4d94-885f-f0b6d0eb348f"
	}

	name := fmt.Sprintf("%s/playbooks/%s", parent.ParentString(), playbookID)
	parsedName, err := s.parsePlaybookName(name)
	if err != nil {
		return nil, err
	}

	fqn := parsedName.String()

	obj := proto.Clone(req.GetPlaybook()).(*cxpb.Playbook)
	obj.Name = fqn

	s.populateDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *playbooksServer) UpdatePlaybook(ctx context.Context, req *cxpb.UpdatePlaybookRequest) (*cxpb.Playbook, error) {
	reqName := req.GetPlaybook().GetName()
	name, err := s.parsePlaybookName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Playbook{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Playbook not found: %s", fqn)
		}
		return nil, err
	}

	// Apply update mask
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetPlaybook().GetDisplayName()
		obj.Goal = req.GetPlaybook().GetGoal()
		obj.Instruction = req.GetPlaybook().GetInstruction()
		obj.InputParameterDefinitions = req.GetPlaybook().GetInputParameterDefinitions()
		obj.OutputParameterDefinitions = req.GetPlaybook().GetOutputParameterDefinitions()
		obj.ReferencedPlaybooks = req.GetPlaybook().GetReferencedPlaybooks()
		obj.ReferencedTools = req.GetPlaybook().GetReferencedTools()
		obj.ReferencedFlows = req.GetPlaybook().GetReferencedFlows()
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetPlaybook().GetDisplayName()
			case "goal":
				obj.Goal = req.GetPlaybook().GetGoal()
			case "instruction":
				obj.Instruction = req.GetPlaybook().GetInstruction()
			case "input_parameter_definitions", "inputParameterDefinitions":
				obj.InputParameterDefinitions = req.GetPlaybook().GetInputParameterDefinitions()
			case "output_parameter_definitions", "outputParameterDefinitions":
				obj.OutputParameterDefinitions = req.GetPlaybook().GetOutputParameterDefinitions()
			case "referenced_playbooks", "referencedPlaybooks":
				obj.ReferencedPlaybooks = req.GetPlaybook().GetReferencedPlaybooks()
			case "referenced_tools", "referencedTools":
				obj.ReferencedTools = req.GetPlaybook().GetReferencedTools()
			case "referenced_flows", "referencedFlows":
				obj.ReferencedFlows = req.GetPlaybook().GetReferencedFlows()
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

func (s *playbooksServer) DeletePlaybook(ctx context.Context, req *cxpb.DeletePlaybookRequest) (*emptypb.Empty, error) {
	name, err := s.parsePlaybookName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &cxpb.Playbook{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Playbook not found: %s", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *playbooksServer) ListPlaybooks(ctx context.Context, req *cxpb.ListPlaybooksRequest) (*cxpb.ListPlaybooksResponse, error) {
	parent, err := s.parsePlaybookParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := parent.ParentString() + "/playbooks/"
	response := &cxpb.ListPlaybooksResponse{}

	playbookKind := (&cxpb.Playbook{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, playbookKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		pb := obj.(*cxpb.Playbook)
		s.populateDefaults(pb)
		response.Playbooks = append(response.Playbooks, pb)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func RegisterCXPlaybooksHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := cxpb.NewPlaybooksClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// POST /v3/{parent=projects/*/locations/*/agents/*}/playbooks
	createPath := "/v3/{parent=projects/*/locations/*/agents/*}/playbooks"
	if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.CreatePlaybookRequest{}
		protoReq.Parent = pathParams["parent"]

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Playbook = &cxpb.Playbook{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Playbook); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}

		obj, err := client.CreatePlaybook(ctx, protoReq)
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

	// GET /v3/{name=projects/*/locations/*/agents/*/playbooks/*}
	getPath := "/v3/{name=projects/*/locations/*/agents/*/playbooks/*}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.GetPlaybookRequest{
			Name: pathParams["name"],
		}

		obj, err := client.GetPlaybook(ctx, protoReq)
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

	// PATCH /v3/{playbook.name=projects/*/locations/*/agents/*/playbooks/*}
	updatePath := "/v3/{playbook.name=projects/*/locations/*/agents/*/playbooks/*}"
	if err := mux.HandlePath("PATCH", updatePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.UpdatePlaybookRequest{}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Playbook = &cxpb.Playbook{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Playbook); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.Playbook.Name = pathParams["playbook.name"]

		if maskStr := r.URL.Query().Get("updateMask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		} else if maskStr := r.URL.Query().Get("update_mask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		}

		obj, err := client.UpdatePlaybook(ctx, protoReq)
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

	// DELETE /v3/{name=projects/*/locations/*/agents/*/playbooks/*}
	deletePath := "/v3/{name=projects/*/locations/*/agents/*/playbooks/*}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.DeletePlaybookRequest{
			Name: pathParams["name"],
		}

		obj, err := client.DeletePlaybook(ctx, protoReq)
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
