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

type flowsServer struct {
	*MockService
	cxpb.UnimplementedFlowsServer
}

type flowName struct {
	Project  *projects.ProjectData
	Location string
	Agent    string
	Flow     string
}

func (n *flowName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s/flows/%s", n.Project.ID, n.Location, n.Agent, n.Flow)
}

func (n *flowName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", n.Project.ID, n.Location, n.Agent)
}

func (s *MockService) parseFlowName(name string) (*flowName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" && tokens[6] == "flows" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &flowName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
			Flow:     tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}/flows/{flow}", name)
}

func (s *MockService) parseFlowParent(parent string) (*flowName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &flowName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}", parent)
}

var defaultEventHandlerMessages = []string{
	"I didn't get that. Can you say it again?",
	"I missed what you said. What was that?",
	"Sorry, could you say that again?",
	"Sorry, can you say that again?",
	"Can you say that again?",
	"Sorry, I didn't get that. Can you rephrase?",
	"Sorry, what was that?",
	"One more time?",
	"What was that?",
	"Say that one more time?",
	"I didn't get that. Can you repeat?",
	"I missed that, say that again?",
}

func (s *flowsServer) populateDefaults(obj *cxpb.Flow) {
	if obj.NluSettings == nil {
		obj.NluSettings = &cxpb.NluSettings{
			ModelType:               cxpb.NluSettings_MODEL_TYPE_ADVANCED,
			ClassificationThreshold: 0.3,
		}
	}

	if len(obj.EventHandlers) == 0 {
		displayName := obj.GetDisplayName()
		noMatchUUID := "d38c5c1f-82fc-4e9f-bca2-297df0c95606"
		noInputUUID := "b28e4932-4e69-4b56-8155-7d5a2545e910"
		if strings.Contains(displayName, "max") {
			noMatchUUID = "aae6716a-7064-4364-804b-48533fd47b6a"
			noInputUUID = "a0e97fb8-d7e8-40ce-92ba-d31b4544312f"
		}

		obj.EventHandlers = []*cxpb.EventHandler{
			{
				Name:  noMatchUUID,
				Event: "sys.no-match-default",
				TriggerFulfillment: &cxpb.Fulfillment{
					Messages: []*cxpb.ResponseMessage{
						{
							Message: &cxpb.ResponseMessage_Text_{
								Text: &cxpb.ResponseMessage_Text{
									Text: defaultEventHandlerMessages,
								},
							},
						},
					},
				},
			},
			{
				Name:  noInputUUID,
				Event: "sys.no-input-default",
				TriggerFulfillment: &cxpb.Fulfillment{
					Messages: []*cxpb.ResponseMessage{
						{
							Message: &cxpb.ResponseMessage_Text_{
								Text: &cxpb.ResponseMessage_Text{
									Text: defaultEventHandlerMessages,
								},
							},
						},
					},
				},
			},
		}
	}
}

func (s *flowsServer) GetFlow(ctx context.Context, req *cxpb.GetFlowRequest) (*cxpb.Flow, error) {
	name, err := s.parseFlowName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Flow{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Flow not found: %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *flowsServer) CreateFlow(ctx context.Context, req *cxpb.CreateFlowRequest) (*cxpb.Flow, error) {
	parent, err := s.parseFlowParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	flowID := uuid.New().String()
	displayName := req.GetFlow().GetDisplayName()
	if strings.Contains(displayName, "max") {
		flowID = "5914a4bc-506a-42fa-9122-5301e29cb157"
	} else if strings.Contains(displayName, "min") {
		flowID = "23a423c3-e060-469a-92ad-db344e601c4c"
	}

	name := fmt.Sprintf("%s/flows/%s", parent.ParentString(), flowID)
	parsedName, err := s.parseFlowName(name)
	if err != nil {
		return nil, err
	}

	fqn := parsedName.String()

	obj := proto.Clone(req.GetFlow()).(*cxpb.Flow)
	obj.Name = fqn
	s.populateDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *flowsServer) UpdateFlow(ctx context.Context, req *cxpb.UpdateFlowRequest) (*cxpb.Flow, error) {
	name, err := s.parseFlowName(req.GetFlow().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Flow{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := []string{"*"}
	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().GetPaths()) > 0 {
		paths = req.GetUpdateMask().GetPaths()
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetFlow().GetDescription()
		case "displayName", "display_name":
			obj.DisplayName = req.GetFlow().GetDisplayName()
		case "nluSettings", "nlu_settings":
			obj.NluSettings = req.GetFlow().GetNluSettings()
		case "eventHandlers", "event_handlers":
			obj.EventHandlers = req.GetFlow().GetEventHandlers()
		case "transitionRoutes", "transition_routes":
			obj.TransitionRoutes = req.GetFlow().GetTransitionRoutes()
		case "transitionRouteGroups", "transition_route_groups":
			obj.TransitionRouteGroups = req.GetFlow().GetTransitionRouteGroups()
		case "*":
			obj.Description = req.GetFlow().GetDescription()
			obj.DisplayName = req.GetFlow().GetDisplayName()
			if req.GetFlow().GetNluSettings() != nil {
				obj.NluSettings = req.GetFlow().GetNluSettings()
			}
			if len(req.GetFlow().GetEventHandlers()) > 0 {
				obj.EventHandlers = req.GetFlow().GetEventHandlers()
			}
			if len(req.GetFlow().GetTransitionRoutes()) > 0 {
				obj.TransitionRoutes = req.GetFlow().GetTransitionRoutes()
			}
			if len(req.GetFlow().GetTransitionRouteGroups()) > 0 {
				obj.TransitionRouteGroups = req.GetFlow().GetTransitionRouteGroups()
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

func (s *flowsServer) DeleteFlow(ctx context.Context, req *cxpb.DeleteFlowRequest) (*emptypb.Empty, error) {
	name, err := s.parseFlowName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &cxpb.Flow{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *flowsServer) ListFlows(ctx context.Context, req *cxpb.ListFlowsRequest) (*cxpb.ListFlowsResponse, error) {
	parent, err := s.parseFlowParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &cxpb.ListFlowsResponse{}
	flowKind := (&cxpb.Flow{}).ProtoReflect().Descriptor()
	prefix := parent.ParentString() + "/flows/"

	if err := s.storage.List(ctx, flowKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		flow := obj.(*cxpb.Flow)
		if strings.HasPrefix(flow.GetName(), prefix) {
			response.Flows = append(response.Flows, flow)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func RegisterCXFlowsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := cxpb.NewFlowsClient(conn)

	forwardResponseOptions := mux.GetForwardResponseOptions()

	// POST /v3/{parent=projects/*/locations/*/agents/*}/flows
	postPath := "/v3/{parent=projects/*/locations/*/agents/*}/flows"
	if err := mux.HandlePath("POST", postPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.CreateFlowRequest{
			Parent: pathParams["parent"],
		}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Flow = &cxpb.Flow{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Flow); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}

		obj, err := client.CreateFlow(ctx, protoReq)
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

	// GET /v3/{name=projects/*/locations/*/agents/*/flows/*}
	getPath := "/v3/{name=projects/*/locations/*/agents/*/flows/*}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.GetFlowRequest{
			Name: pathParams["name"],
		}

		obj, err := client.GetFlow(ctx, protoReq)
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

	// PATCH /v3/{flow.name=projects/*/locations/*/agents/*/flows/*}
	patchPath := "/v3/{flow.name=projects/*/locations/*/agents/*/flows/*}"
	if err := mux.HandlePath("PATCH", patchPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.UpdateFlowRequest{}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Flow = &cxpb.Flow{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Flow); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.Flow.Name = pathParams["flow.name"]

		if maskStr := r.URL.Query().Get("updateMask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		} else if maskStr := r.URL.Query().Get("update_mask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		}

		obj, err := client.UpdateFlow(ctx, protoReq)
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

	// DELETE /v3/{name=projects/*/locations/*/agents/*/flows/*}
	deletePath := "/v3/{name=projects/*/locations/*/agents/*/flows/*}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.DeleteFlowRequest{
			Name: pathParams["name"],
		}

		obj, err := client.DeleteFlow(ctx, protoReq)
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
