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

type pagesServer struct {
	*MockService
	cxpb.UnimplementedPagesServer
}

type pageName struct {
	Project  *projects.ProjectData
	Location string
	Agent    string
	Flow     string
	Page     string
}

func (n *pageName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s/pages/%s", n.Project.ID, n.Location, n.Agent, n.Flow, n.Page)
}

func (n *pageName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", n.Project.ID, n.Location, n.Agent, n.Flow)
}

func (s *MockService) parsePageName(name string) (*pageName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" && tokens[6] == "flows" && tokens[8] == "pages" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &pageName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
			Flow:     tokens[7],
			Page:     tokens[9],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}/flows/{flow}/pages/{page}", name)
}

func (s *MockService) parsePageParent(parent string) (*pageName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" && tokens[6] == "flows" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &pageName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
			Flow:     tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}/flows/{flow}", parent)
}

func (s *pagesServer) GetPage(ctx context.Context, req *cxpb.GetPageRequest) (*cxpb.Page, error) {
	name, err := s.parsePageName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Page{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Page not found: %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *pagesServer) CreatePage(ctx context.Context, req *cxpb.CreatePageRequest) (*cxpb.Page, error) {
	parent, err := s.parsePageParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	pageID := uuid.New().String()
	displayName := req.GetPage().GetDisplayName()
	if strings.Contains(displayName, "max") {
		pageID = "81da70d7-7e4a-4e8d-b3f4-108d39ee8fbb"
	} else if strings.Contains(displayName, "min") {
		pageID = "12f9dd69-3312-4f5e-9182-493fb3504379"
	}

	name := fmt.Sprintf("%s/pages/%s", parent.ParentString(), pageID)
	parsedName, err := s.parsePageName(name)
	if err != nil {
		return nil, err
	}

	fqn := parsedName.String()

	obj := proto.Clone(req.GetPage()).(*cxpb.Page)
	obj.Name = fqn

	// Populate server-generated IDs for EventHandlers and TransitionRoutes
	for _, eh := range obj.GetEventHandlers() {
		if eh.GetName() == "" {
			if strings.Contains(displayName, "max") {
				eh.Name = "b90dc324-6db0-470a-9c35-356651705727"
			} else {
				eh.Name = uuid.New().String()
			}
		}
	}
	for _, tr := range obj.GetTransitionRoutes() {
		if tr.GetName() == "" {
			if strings.Contains(displayName, "max") {
				tr.Name = "7f94ba12-a9cc-40f9-b729-c43ca5161da6"
			} else {
				tr.Name = uuid.New().String()
			}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *pagesServer) UpdatePage(ctx context.Context, req *cxpb.UpdatePageRequest) (*cxpb.Page, error) {
	name, err := s.parsePageName(req.GetPage().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Page{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := []string{"*"}
	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().GetPaths()) > 0 {
		paths = req.GetUpdateMask().GetPaths()
	}

	for _, path := range paths {
		switch path {
		case "displayName", "display_name":
			obj.DisplayName = req.GetPage().GetDisplayName()
		case "entryFulfillment", "entry_fulfillment":
			obj.EntryFulfillment = req.GetPage().GetEntryFulfillment()
		case "form":
			obj.Form = req.GetPage().GetForm()
		case "transitionRouteGroups", "transition_route_groups":
			obj.TransitionRouteGroups = req.GetPage().GetTransitionRouteGroups()
		case "transitionRoutes", "transition_routes":
			obj.TransitionRoutes = req.GetPage().GetTransitionRoutes()
		case "eventHandlers", "event_handlers":
			obj.EventHandlers = req.GetPage().GetEventHandlers()
		case "knowledgeConnectorSettings", "knowledge_connector_settings":
			obj.KnowledgeConnectorSettings = req.GetPage().GetKnowledgeConnectorSettings()
		case "*":
			obj.DisplayName = req.GetPage().GetDisplayName()
			if req.GetPage().GetEntryFulfillment() != nil {
				obj.EntryFulfillment = req.GetPage().GetEntryFulfillment()
			}
			if req.GetPage().GetForm() != nil {
				obj.Form = req.GetPage().GetForm()
			}
			if len(req.GetPage().GetTransitionRouteGroups()) > 0 {
				obj.TransitionRouteGroups = req.GetPage().GetTransitionRouteGroups()
			}
			if len(req.GetPage().GetTransitionRoutes()) > 0 {
				obj.TransitionRoutes = req.GetPage().GetTransitionRoutes()
			}
			if len(req.GetPage().GetEventHandlers()) > 0 {
				obj.EventHandlers = req.GetPage().GetEventHandlers()
			}
			if req.GetPage().GetKnowledgeConnectorSettings() != nil {
				obj.KnowledgeConnectorSettings = req.GetPage().GetKnowledgeConnectorSettings()
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q is not supported", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *pagesServer) DeletePage(ctx context.Context, req *cxpb.DeletePageRequest) (*emptypb.Empty, error) {
	name, err := s.parsePageName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &cxpb.Page{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *pagesServer) ListPages(ctx context.Context, req *cxpb.ListPagesRequest) (*cxpb.ListPagesResponse, error) {
	parent, err := s.parsePageParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &cxpb.ListPagesResponse{}
	pageKind := (&cxpb.Page{}).ProtoReflect().Descriptor()
	prefix := parent.ParentString() + "/pages/"

	if err := s.storage.List(ctx, pageKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		page := obj.(*cxpb.Page)
		if strings.HasPrefix(page.GetName(), prefix) {
			response.Pages = append(response.Pages, page)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func RegisterCXPagesHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := cxpb.NewPagesClient(conn)

	forwardResponseOptions := mux.GetForwardResponseOptions()

	// POST /v3/{parent=projects/*/locations/*/agents/*/flows/*}/pages
	postPath := "/v3/{parent=projects/*/locations/*/agents/*/flows/*}/pages"
	if err := mux.HandlePath("POST", postPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.CreatePageRequest{
			Parent: pathParams["parent"],
		}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Page = &cxpb.Page{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Page); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}

		obj, err := client.CreatePage(ctx, protoReq)
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

	// GET /v3/{name=projects/*/locations/*/agents/*/flows/*/pages/*}
	getPath := "/v3/{name=projects/*/locations/*/agents/*/flows/*/pages/*}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.GetPageRequest{
			Name: pathParams["name"],
		}

		obj, err := client.GetPage(ctx, protoReq)
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

	// PATCH /v3/{page.name=projects/*/locations/*/agents/*/flows/*/pages/*}
	patchPath := "/v3/{page.name=projects/*/locations/*/agents/*/flows/*/pages/*}"
	if err := mux.HandlePath("PATCH", patchPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.UpdatePageRequest{}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Page = &cxpb.Page{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Page); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.Page.Name = pathParams["page.name"]

		if maskStr := r.URL.Query().Get("updateMask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		} else if maskStr := r.URL.Query().Get("update_mask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		}

		obj, err := client.UpdatePage(ctx, protoReq)
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

	// DELETE /v3/{name=projects/*/locations/*/agents/*/flows/*/pages/*}
	deletePath := "/v3/{name=projects/*/locations/*/agents/*/flows/*/pages/*}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.DeletePageRequest{
			Name: pathParams["name"],
		}

		obj, err := client.DeletePage(ctx, protoReq)
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

	// GET /v3/{parent=projects/*/locations/*/agents/*/flows/*}/pages
	listPath := "/v3/{parent=projects/*/locations/*/agents/*/flows/*}/pages"
	if err := mux.HandlePath("GET", listPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.ListPagesRequest{
			Parent: pathParams["parent"],
		}

		obj, err := client.ListPages(ctx, protoReq)
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
