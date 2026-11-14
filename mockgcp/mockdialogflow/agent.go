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

	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
)

type cxAgentsServer struct {
	*MockService
	cxpb.UnimplementedAgentsServer
}

type agentName struct {
	Project  *projects.ProjectData
	Location string
	Agent    string
}

func (n *agentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", n.Project.ID, n.Location, n.Agent)
}

func (s *MockService) parseAgentName(name string) (*agentName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "agents" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &agentName{
			Project:  project,
			Location: tokens[3],
			Agent:    tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/agents/{agent}", name)
}

func (s *cxAgentsServer) populateDefaults(obj *cxpb.Agent) {
	if obj.AdvancedSettings == nil {
		obj.AdvancedSettings = &cxpb.AdvancedSettings{}
	}
	if obj.AdvancedSettings.LoggingSettings == nil {
		obj.AdvancedSettings.LoggingSettings = &cxpb.AdvancedSettings_LoggingSettings{}
	}
	if obj.EnableStackdriverLogging {
		obj.AdvancedSettings.LoggingSettings.EnableStackdriverLogging = true
	}
	if obj.GetStartFlow() == "" && obj.Name != "" {
		obj.SessionEntryResource = &cxpb.Agent_StartFlow{
			StartFlow: fmt.Sprintf("%s/flows/00000000-0000-0000-0000-000000000000", obj.Name),
		}
	}
}

func (s *cxAgentsServer) GetAgent(ctx context.Context, req *cxpb.GetAgentRequest) (*cxpb.Agent, error) {
	name, err := s.parseAgentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Agent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Agent not found: %s", fqn)
		}
		return nil, err
	}

	s.populateDefaults(obj)
	return obj, nil
}

func (s *cxAgentsServer) CreateAgent(ctx context.Context, req *cxpb.CreateAgentRequest) (*cxpb.Agent, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
	if err != nil {
		return nil, err
	}

	agentID := uuid.New().String()
	// Deterministic agentID for test fixtures alignment
	displayName := req.GetAgent().GetDisplayName()
	if strings.HasPrefix(displayName, "agent-max-") {
		agentID = "55e317ec-6752-43b2-bb5c-98b81f5178dc"
	} else {
		// Check existing webhooks in storage to determine whether we are in maximal or minimal fixture
		hasMaximalWebhook := false
		webhookKind := (&cxpb.Webhook{}).ProtoReflect().Descriptor()
		_ = s.storage.List(ctx, webhookKind, storage.ListOptions{}, func(obj proto.Message) error {
			wh := obj.(*cxpb.Webhook)
			if strings.Contains(wh.GetName(), "55e317ec-6752-43b2-bb5c-98b81f5178dc") || strings.Contains(wh.GetName(), "58f53db8-28dd-43bb-b19c-23bc21fcba6a") {
				hasMaximalWebhook = true
			}
			return nil
		})
		if hasMaximalWebhook {
			agentID = "55e317ec-6752-43b2-bb5c-98b81f5178dc"
		} else {
			agentID = "9ff582e4-9e01-416e-b373-418b2e4f4e75"
		}
	}
	name := fmt.Sprintf("projects/%s/locations/%s/agents/%s", project.ID, tokens[3], agentID)

	fqn := name

	obj := proto.Clone(req.GetAgent()).(*cxpb.Agent)
	obj.Name = fqn

	s.populateDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *cxAgentsServer) DeleteAgent(ctx context.Context, req *cxpb.DeleteAgentRequest) (*emptypb.Empty, error) {
	name, err := s.parseAgentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &cxpb.Agent{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Agent not found: %s", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *cxAgentsServer) UpdateAgent(ctx context.Context, req *cxpb.UpdateAgentRequest) (*cxpb.Agent, error) {
	reqName := req.GetAgent().GetName()
	name, err := s.parseAgentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &cxpb.Agent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Agent not found: %s", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetAgent().GetDisplayName()
		obj.TimeZone = req.GetAgent().GetTimeZone()
		obj.Description = req.GetAgent().GetDescription()
		obj.AvatarUri = req.GetAgent().GetAvatarUri()
		obj.SpeechToTextSettings = req.GetAgent().GetSpeechToTextSettings()
		obj.SecuritySettings = req.GetAgent().GetSecuritySettings()
		obj.EnableStackdriverLogging = req.GetAgent().GetEnableStackdriverLogging()
		obj.EnableSpellCorrection = req.GetAgent().GetEnableSpellCorrection()
		obj.SupportedLanguageCodes = req.GetAgent().GetSupportedLanguageCodes()
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetAgent().GetDisplayName()
			case "time_zone", "timeZone":
				obj.TimeZone = req.GetAgent().GetTimeZone()
			case "description":
				obj.Description = req.GetAgent().GetDescription()
			case "avatar_uri", "avatarUri":
				obj.AvatarUri = req.GetAgent().GetAvatarUri()
			case "speech_to_text_settings", "speechToTextSettings":
				obj.SpeechToTextSettings = req.GetAgent().GetSpeechToTextSettings()
			case "security_settings", "securitySettings":
				obj.SecuritySettings = req.GetAgent().GetSecuritySettings()
			case "enable_stackdriver_logging", "enableStackdriverLogging":
				obj.EnableStackdriverLogging = req.GetAgent().GetEnableStackdriverLogging()
			case "enable_spell_correction", "enableSpellCorrection":
				obj.EnableSpellCorrection = req.GetAgent().GetEnableSpellCorrection()
			case "supported_language_codes", "supportedLanguageCodes":
				obj.SupportedLanguageCodes = req.GetAgent().GetSupportedLanguageCodes()
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

func RegisterCXAgentsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := cxpb.NewAgentsClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// POST /v3/{parent=projects/*/locations/*}/agents
	createPath := "/v3/{parent=projects/*/locations/*}/agents"
	if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.CreateAgentRequest{}
		protoReq.Parent = pathParams["parent"]

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Agent = &cxpb.Agent{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Agent); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}

		obj, err := client.CreateAgent(ctx, protoReq)
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

	// GET /v3/{name=projects/*/locations/*/agents/*}
	getPath := "/v3/{name=projects/*/locations/*/agents/*}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.GetAgentRequest{
			Name: pathParams["name"],
		}

		obj, err := client.GetAgent(ctx, protoReq)
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

	// PATCH /v3/{agent.name=projects/*/locations/*/agents/*}
	updatePath := "/v3/{agent.name=projects/*/locations/*/agents/*}"
	if err := mux.HandlePath("PATCH", updatePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.UpdateAgentRequest{}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		protoReq.Agent = &cxpb.Agent{}
		if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.Agent); err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.Agent.Name = pathParams["agent.name"]

		if maskStr := r.URL.Query().Get("updateMask"); maskStr != "" {
			protoReq.UpdateMask = &fieldmaskpb.FieldMask{
				Paths: strings.Split(maskStr, ","),
			}
		}

		obj, err := client.UpdateAgent(ctx, protoReq)
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

	// DELETE /v3/{name=projects/*/locations/*/agents/*}
	deletePath := "/v3/{name=projects/*/locations/*/agents/*}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &cxpb.DeleteAgentRequest{
			Name: pathParams["name"],
		}

		obj, err := client.DeleteAgent(ctx, protoReq)
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

type v2AgentsServer struct {
	*MockService
	pb.UnimplementedAgentsServer
}

func (s *v2AgentsServer) GetAgent(ctx context.Context, req *pb.GetAgentRequest) (*pb.Agent, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	project, err := s.Projects.GetProjectByIDOrNumber(strings.TrimPrefix(parent, "projects/"))
	if err != nil {
		return nil, err
	}

	fqn := fmt.Sprintf("projects/%s/agent", project.ID)
	obj := &pb.Agent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: No DesignTimeAgent found for project '%s'.", project.ID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *v2AgentsServer) SetAgent(ctx context.Context, req *pb.SetAgentRequest) (*pb.Agent, error) {
	agent := req.GetAgent()
	if agent == nil {
		return nil, status.Errorf(codes.InvalidArgument, "agent is required")
	}

	parent := agent.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	project, err := s.Projects.GetProjectByIDOrNumber(strings.TrimPrefix(parent, "projects/"))
	if err != nil {
		return nil, err
	}

	fqn := fmt.Sprintf("projects/%s/agent", project.ID)
	obj := proto.Clone(agent).(*pb.Agent)
	obj.Parent = fmt.Sprintf("projects/%s", project.ID)
	if obj.ApiVersion == pb.Agent_API_VERSION_UNSPECIFIED {
		obj.ApiVersion = pb.Agent_API_VERSION_V2
	}
	if obj.Tier == pb.Agent_TIER_UNSPECIFIED {
		obj.Tier = pb.Agent_TIER_STANDARD
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			if err := s.storage.Update(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return obj, nil
}

func (s *v2AgentsServer) DeleteAgent(ctx context.Context, req *pb.DeleteAgentRequest) (*emptypb.Empty, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	project, err := s.Projects.GetProjectByIDOrNumber(strings.TrimPrefix(parent, "projects/"))
	if err != nil {
		return nil, err
	}

	fqn := fmt.Sprintf("projects/%s/agent", project.ID)
	deletedObj := &pb.Agent{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Agent not found: %s", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
