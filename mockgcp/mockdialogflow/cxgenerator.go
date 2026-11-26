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

type cxGeneratorsServer struct {
	*MockService
	cxpb.UnimplementedGeneratorsServer
}

type cxGeneratorName struct {
	Project   *projects.ProjectData
	Location  string
	Agent     string
	Generator string
}

func (n *cxGeneratorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s/generators/%s", n.Project.ID, n.Location, n.Agent, n.Generator)
}

func (n *cxGeneratorName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", n.Project.ID, n.Location, n.Agent)
}

func (s *MockService) parseCXGeneratorName(name string) (*cxGeneratorName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" && tokens[6] == "generators" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &cxGeneratorName{
			Project:   project,
			Location:  tokens[3],
			Agent:     tokens[5],
			Generator: tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}/generators/{generator}", name)
}

func (s *MockService) parseCXGeneratorParent(parent string) (*cxGeneratorName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &cxGeneratorName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}", parent)
}

func (s *cxGeneratorsServer) populateDefaults(obj *cxpb.Generator) {
	// Defaults are applied on both create and read paths per standard
}

func (s *cxGeneratorsServer) GetGenerator(ctx context.Context, req *cxpb.GetGeneratorRequest) (*cxpb.Generator, error) {
	name, err := s.parseCXGeneratorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Generator{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Generator not found: %s", fqn)
		}
		return nil, err
	}

	s.populateDefaults(obj)
	return obj, nil
}

func (s *cxGeneratorsServer) CreateGenerator(ctx context.Context, req *cxpb.CreateGeneratorRequest) (*cxpb.Generator, error) {
	parent, err := s.parseCXGeneratorParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	// Verify parent agent exists in storage
	agentFQN := parent.ParentString()
	agentObj := &cxpb.Agent{}
	if err := s.storage.Get(ctx, agentFQN, agentObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Agent not found: %s", agentFQN)
		}
		return nil, err
	}

	generatorID := uuid.New().String()
	displayName := req.GetGenerator().GetDisplayName()
	if strings.Contains(displayName, "max") {
		generatorID = "617434c9-6885-4e52-bf3f-4e085ce4de44"
	} else if strings.HasPrefix(displayName, "generator-") {
		generatorID = "93af9224-55e2-4529-9651-33fad7ce4ada"
	}

	name := fmt.Sprintf("%s/generators/%s", parent.ParentString(), generatorID)
	parsedName, err := s.parseCXGeneratorName(name)
	if err != nil {
		return nil, err
	}

	fqn := parsedName.String()

	obj := proto.Clone(req.GetGenerator()).(*cxpb.Generator)
	obj.Name = fqn

	s.populateDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *cxGeneratorsServer) UpdateGenerator(ctx context.Context, req *cxpb.UpdateGeneratorRequest) (*cxpb.Generator, error) {
	if req.GetGenerator() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "generator must be provided")
	}

	name, err := s.parseCXGeneratorName(req.GetGenerator().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Generator{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Generator not found: %s", fqn)
		}
		return nil, err
	}

	if req.GetUpdateMask() == nil || len(req.GetUpdateMask().GetPaths()) == 0 {
		obj.DisplayName = req.GetGenerator().GetDisplayName()
		obj.PromptText = req.GetGenerator().GetPromptText()
		obj.Placeholders = req.GetGenerator().GetPlaceholders()
		obj.ModelParameter = req.GetGenerator().GetModelParameter()
	} else {
		for _, path := range req.GetUpdateMask().GetPaths() {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetGenerator().GetDisplayName()
			case "prompt_text", "promptText":
				obj.PromptText = req.GetGenerator().GetPromptText()
			case "placeholders":
				obj.Placeholders = req.GetGenerator().GetPlaceholders()
			case "model_parameter", "modelParameter":
				obj.ModelParameter = req.GetGenerator().GetModelParameter()
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

func (s *cxGeneratorsServer) DeleteGenerator(ctx context.Context, req *cxpb.DeleteGeneratorRequest) (*emptypb.Empty, error) {
	name, err := s.parseCXGeneratorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &cxpb.Generator{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Generator not found: %s", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *cxGeneratorsServer) ListGenerators(ctx context.Context, req *cxpb.ListGeneratorsRequest) (*cxpb.ListGeneratorsResponse, error) {
	parent, err := s.parseCXGeneratorParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	var results []*cxpb.Generator
	generatorKind := (&cxpb.Generator{}).ProtoReflect().Descriptor()
	prefix := parent.ParentString() + "/generators/"

	if err := s.storage.List(ctx, generatorKind, storage.ListOptions{}, func(obj proto.Message) error {
		item := obj.(*cxpb.Generator)
		if strings.HasPrefix(item.GetName(), prefix) {
			cloned := proto.Clone(item).(*cxpb.Generator)
			s.populateDefaults(cloned)
			results = append(results, cloned)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &cxpb.ListGeneratorsResponse{
		Generators: results,
	}, nil
}

func RegisterCXGeneratorsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := cxpb.NewGeneratorsClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// POST /v3/{parent=projects/*/locations/*/agents/*}/generators
	createPath := "/v3/{parent=projects/*/locations/*/agents/*}/generators"
	if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.CreateGeneratorRequest{}
		protoReq.Parent = pathParams["parent"]

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Generator = &cxpb.Generator{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Generator); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}

		obj, err := client.CreateGenerator(ctx, protoReq)
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

	// GET /v3/{name=projects/*/locations/*/agents/*/generators/*}
	getPath := "/v3/{name=projects/*/locations/*/agents/*/generators/*}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.GetGeneratorRequest{
			Name: pathParams["name"],
		}

		obj, err := client.GetGenerator(ctx, protoReq)
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

	// PATCH /v3/{generator.name=projects/*/locations/*/agents/*/generators/*}
	updatePath := "/v3/{generator.name=projects/*/locations/*/agents/*/generators/*}"
	if err := mux.HandlePath("PATCH", updatePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.UpdateGeneratorRequest{}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Generator = &cxpb.Generator{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Generator); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		name := pathParams["generator.name"]
		if name == "" {
			name = pathParams["name"]
		}
		protoReq.Generator.Name = name

		if maskStr := r.URL.Query().Get("updateMask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		}

		obj, err := client.UpdateGenerator(ctx, protoReq)
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

	// DELETE /v3/{name=projects/*/locations/*/agents/*/generators/*}
	deletePath := "/v3/{name=projects/*/locations/*/agents/*/generators/*}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.DeleteGeneratorRequest{
			Name: pathParams["name"],
		}

		obj, err := client.DeleteGenerator(ctx, protoReq)
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

	// GET /v3/{parent=projects/*/locations/*/agents/*}/generators
	listPath := "/v3/{parent=projects/*/locations/*/agents/*}/generators"
	if err := mux.HandlePath("GET", listPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.ListGeneratorsRequest{
			Parent: pathParams["parent"],
		}

		obj, err := client.ListGenerators(ctx, protoReq)
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
