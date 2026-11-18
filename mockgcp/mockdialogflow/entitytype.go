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

type entityTypesServer struct {
	*MockService
	cxpb.UnimplementedEntityTypesServer
}

type entityTypeName struct {
	Project    *projects.ProjectData
	Location   string
	Agent      string
	EntityType string
}

func (n *entityTypeName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s/entityTypes/%s", n.Project.ID, n.Location, n.Agent, n.EntityType)
}

func (n *entityTypeName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", n.Project.ID, n.Location, n.Agent)
}

func (s *MockService) parseEntityTypeName(name string) (*entityTypeName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" && tokens[6] == "entityTypes" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &entityTypeName{
			Project:    project,
			Location:   tokens[3],
			Agent:      tokens[5],
			EntityType: tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}/entityTypes/{entityType}", name)
}

func (s *MockService) parseEntityTypeParent(parent string) (*entityTypeName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &entityTypeName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}", parent)
}

func (s *entityTypesServer) populateDefaults(obj *cxpb.EntityType) {
}

func (s *entityTypesServer) GetEntityType(ctx context.Context, req *cxpb.GetEntityTypeRequest) (*cxpb.EntityType, error) {
	name, err := s.parseEntityTypeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.EntityType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntityType not found: %s", fqn)
		}
		return nil, err
	}

	s.populateDefaults(obj)
	return obj, nil
}

func (s *entityTypesServer) CreateEntityType(ctx context.Context, req *cxpb.CreateEntityTypeRequest) (*cxpb.EntityType, error) {
	parent, err := s.parseEntityTypeParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	entityTypeID := uuid.New().String()
	displayName := req.GetEntityType().GetDisplayName()
	if strings.Contains(displayName, "max") {
		entityTypeID = "df4e8490-e1ba-4035-99de-ba3fdd265a33"
	} else if strings.Contains(displayName, "min") {
		entityTypeID = "13c4c465-0c78-4834-9f62-ff93d1a73f8f"
	}

	name := fmt.Sprintf("%s/entityTypes/%s", parent.ParentString(), entityTypeID)
	parsedName, err := s.parseEntityTypeName(name)
	if err != nil {
		return nil, err
	}

	fqn := parsedName.String()

	obj := proto.Clone(req.GetEntityType()).(*cxpb.EntityType)
	obj.Name = fqn

	s.populateDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *entityTypesServer) UpdateEntityType(ctx context.Context, req *cxpb.UpdateEntityTypeRequest) (*cxpb.EntityType, error) {
	reqName := req.GetEntityType().GetName()
	name, err := s.parseEntityTypeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.EntityType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntityType not found: %s", fqn)
		}
		return nil, err
	}

	// Apply update mask
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetEntityType().GetDisplayName()
		obj.Kind = req.GetEntityType().GetKind()
		obj.AutoExpansionMode = req.GetEntityType().GetAutoExpansionMode()
		obj.Entities = req.GetEntityType().GetEntities()
		obj.ExcludedPhrases = req.GetEntityType().GetExcludedPhrases()
		obj.EnableFuzzyExtraction = req.GetEntityType().GetEnableFuzzyExtraction()
		obj.Redact = req.GetEntityType().GetRedact()
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetEntityType().GetDisplayName()
			case "kind":
				obj.Kind = req.GetEntityType().GetKind()
			case "auto_expansion_mode", "autoExpansionMode":
				obj.AutoExpansionMode = req.GetEntityType().GetAutoExpansionMode()
			case "entities":
				obj.Entities = req.GetEntityType().GetEntities()
			case "excluded_phrases", "excludedPhrases":
				obj.ExcludedPhrases = req.GetEntityType().GetExcludedPhrases()
			case "enable_fuzzy_extraction", "enableFuzzyExtraction":
				obj.EnableFuzzyExtraction = req.GetEntityType().GetEnableFuzzyExtraction()
			case "redact":
				obj.Redact = req.GetEntityType().GetRedact()
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

func (s *entityTypesServer) DeleteEntityType(ctx context.Context, req *cxpb.DeleteEntityTypeRequest) (*emptypb.Empty, error) {
	name, err := s.parseEntityTypeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &cxpb.EntityType{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntityType not found: %s", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *entityTypesServer) ListEntityTypes(ctx context.Context, req *cxpb.ListEntityTypesRequest) (*cxpb.ListEntityTypesResponse, error) {
	parent, err := s.parseEntityTypeParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.ParentString() + "/entityTypes/"

	var entityTypes []*cxpb.EntityType
	entityTypeKind := (&cxpb.EntityType{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, entityTypeKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		entityType := obj.(*cxpb.EntityType)
		s.populateDefaults(entityType)
		entityTypes = append(entityTypes, entityType)
		return nil
	}); err != nil {
		return nil, err
	}

	return &cxpb.ListEntityTypesResponse{
		EntityTypes: entityTypes,
	}, nil
}

func RegisterCXEntityTypesHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := cxpb.NewEntityTypesClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// POST /v3/{parent=projects/*/locations/*/agents/*}/entityTypes
	createPath := "/v3/{parent=projects/*/locations/*/agents/*}/entityTypes"
	if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.CreateEntityTypeRequest{}
		protoReq.Parent = pathParams["parent"]

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.EntityType = &cxpb.EntityType{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.EntityType); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}

		if lang := r.URL.Query().Get("languageCode"); lang != "" {
			protoReq.LanguageCode = lang
		}

		obj, err := client.CreateEntityType(ctx, protoReq)
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

	// GET /v3/{name=projects/*/locations/*/agents/*/entityTypes/*}
	getPath := "/v3/{name=projects/*/locations/*/agents/*/entityTypes/*}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.GetEntityTypeRequest{
			Name: pathParams["name"],
		}
		if lang := r.URL.Query().Get("languageCode"); lang != "" {
			protoReq.LanguageCode = lang
		}

		obj, err := client.GetEntityType(ctx, protoReq)
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

	// PATCH /v3/{entity_type.name=projects/*/locations/*/agents/*/entityTypes/*}
	patchPath := "/v3/{entity_type.name=projects/*/locations/*/agents/*/entityTypes/*}"
	if err := mux.HandlePath("PATCH", patchPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.UpdateEntityTypeRequest{}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.EntityType = &cxpb.EntityType{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.EntityType); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.EntityType.Name = pathParams["entity_type.name"]

		if maskStr := r.URL.Query().Get("updateMask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		} else if maskStr := r.URL.Query().Get("update_mask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		}

		if lang := r.URL.Query().Get("languageCode"); lang != "" {
			protoReq.LanguageCode = lang
		}

		obj, err := client.UpdateEntityType(ctx, protoReq)
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

	// DELETE /v3/{name=projects/*/locations/*/agents/*/entityTypes/*}
	deletePath := "/v3/{name=projects/*/locations/*/agents/*/entityTypes/*}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.DeleteEntityTypeRequest{
			Name: pathParams["name"],
		}
		if forceStr := r.URL.Query().Get("force"); forceStr != "" {
			protoReq.Force = strings.EqualFold(forceStr, "true")
		}

		obj, err := client.DeleteEntityType(ctx, protoReq)
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

	// GET /v3/{parent=projects/*/locations/*/agents/*}/entityTypes
	listPath := "/v3/{parent=projects/*/locations/*/agents/*}/entityTypes"
	if err := mux.HandlePath("GET", listPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.ListEntityTypesRequest{
			Parent: pathParams["parent"],
		}
		if lang := r.URL.Query().Get("languageCode"); lang != "" {
			protoReq.LanguageCode = lang
		}

		obj, err := client.ListEntityTypes(ctx, protoReq)
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
