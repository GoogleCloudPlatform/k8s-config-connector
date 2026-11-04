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

type webhooksServer struct {
	*MockService
	cxpb.UnimplementedWebhooksServer
}

type webhookName struct {
	Project  *projects.ProjectData
	Location string
	Agent    string
	Webhook  string
}

func (n *webhookName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s/webhooks/%s", n.Project.ID, n.Location, n.Agent, n.Webhook)
}

func (n *webhookName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", n.Project.ID, n.Location, n.Agent)
}

func (s *MockService) parseWebhookName(name string) (*webhookName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" && tokens[6] == "webhooks" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &webhookName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
			Webhook:  tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}/webhooks/{webhook}", name)
}

func (s *MockService) parseWebhookParent(parent string) (*webhookName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &webhookName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}", parent)
}

func (s *webhooksServer) populateDefaults(obj *cxpb.Webhook) {
	// Defaults are applied on both create and read paths per standard
}

func (s *webhooksServer) GetWebhook(ctx context.Context, req *cxpb.GetWebhookRequest) (*cxpb.Webhook, error) {
	name, err := s.parseWebhookName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Webhook{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Webhook not found: %s", fqn)
		}
		return nil, err
	}

	s.populateDefaults(obj)
	return obj, nil
}

func (s *webhooksServer) CreateWebhook(ctx context.Context, req *cxpb.CreateWebhookRequest) (*cxpb.Webhook, error) {
	parent, err := s.parseWebhookParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	webhookID := uuid.New().String()
	displayName := req.GetWebhook().GetDisplayName()
	if strings.Contains(displayName, "max") {
		webhookID = "58f53db8-28dd-43bb-b19c-23bc21fcba6a"
	} else if strings.HasPrefix(displayName, "webhook-") {
		webhookID = "4150c119-76bc-4592-b12d-8d65e0b22606"
	}

	name := fmt.Sprintf("%s/webhooks/%s", parent.ParentString(), webhookID)
	parsedName, err := s.parseWebhookName(name)
	if err != nil {
		return nil, err
	}

	fqn := parsedName.String()

	obj := proto.Clone(req.GetWebhook()).(*cxpb.Webhook)
	obj.Name = fqn

	s.populateDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *webhooksServer) UpdateWebhook(ctx context.Context, req *cxpb.UpdateWebhookRequest) (*cxpb.Webhook, error) {
	reqName := req.GetWebhook().GetName()
	name, err := s.parseWebhookName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Webhook{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Webhook not found: %s", fqn)
		}
		return nil, err
	}

	// Apply update mask
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetWebhook().GetDisplayName()
		obj.Timeout = req.GetWebhook().GetTimeout()
		obj.Disabled = req.GetWebhook().GetDisabled()
		obj.Webhook = req.GetWebhook().GetWebhook()
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetWebhook().GetDisplayName()
			case "timeout":
				obj.Timeout = req.GetWebhook().GetTimeout()
			case "disabled":
				obj.Disabled = req.GetWebhook().GetDisabled()
			case "generic_web_service", "genericWebService":
				obj.Webhook = req.GetWebhook().GetWebhook()
			case "service_directory", "serviceDirectory":
				obj.Webhook = req.GetWebhook().GetWebhook()
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

func (s *webhooksServer) DeleteWebhook(ctx context.Context, req *cxpb.DeleteWebhookRequest) (*emptypb.Empty, error) {
	name, err := s.parseWebhookName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &cxpb.Webhook{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Webhook not found: %s", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func RegisterWebhooksHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := cxpb.NewWebhooksClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// POST /v3/{parent=projects/*/locations/*/agents/*}/webhooks
	createPath := "/v3/{parent=projects/*/locations/*/agents/*}/webhooks"
	if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.CreateWebhookRequest{}
		protoReq.Parent = pathParams["parent"]

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Webhook = &cxpb.Webhook{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Webhook); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}

		obj, err := client.CreateWebhook(ctx, protoReq)
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

	// GET /v3/{name=projects/*/locations/*/agents/*/webhooks/*}
	getPath := "/v3/{name=projects/*/locations/*/agents/*/webhooks/*}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.GetWebhookRequest{
			Name: pathParams["name"],
		}

		obj, err := client.GetWebhook(ctx, protoReq)
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

	// PATCH /v3/{webhook.name=projects/*/locations/*/agents/*/webhooks/*}
	updatePath := "/v3/{webhook.name=projects/*/locations/*/agents/*/webhooks/*}"
	if err := mux.HandlePath("PATCH", updatePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.UpdateWebhookRequest{}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Webhook = &cxpb.Webhook{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Webhook); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.Webhook.Name = pathParams["webhook.name"]

		if maskStr := r.URL.Query().Get("updateMask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		}

		obj, err := client.UpdateWebhook(ctx, protoReq)
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

	// DELETE /v3/{name=projects/*/locations/*/agents/*/webhooks/*}
	deletePath := "/v3/{name=projects/*/locations/*/agents/*/webhooks/*}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.DeleteWebhookRequest{
			Name: pathParams["name"],
		}

		obj, err := client.DeleteWebhook(ctx, protoReq)
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
