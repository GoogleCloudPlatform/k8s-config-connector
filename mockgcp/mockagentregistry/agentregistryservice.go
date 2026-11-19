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

package mockagentregistry

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/agentregistry/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type serviceName struct {
	Project   *projects.ProjectData
	Location  string
	ServiceID string
}

func (s *AgentRegistryServer) parseServiceName(name string) (*serviceName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}
	parts := strings.Split(name, "/")
	if len(parts) != 6 || parts[0] != "projects" || parts[2] != "locations" || parts[4] != "services" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid service name %q", name)
	}
	project, err := s.Projects.GetProjectByIDOrNumber(parts[1])
	if err != nil {
		return nil, err
	}
	return &serviceName{
		Project:   project,
		Location:  parts[3],
		ServiceID: parts[5],
	}, nil
}

func (n *serviceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/services/%s", n.Project.ID, n.Location, n.ServiceID)
}

func (s *AgentRegistryServer) GetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.Service, error) {
	name, err := s.parseServiceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AgentRegistryServer) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*longrunningpb.Operation, error) {
	parentParts := strings.Split(req.GetParent(), "/")
	if len(parentParts) != 4 || parentParts[0] != "projects" || parentParts[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q", req.GetParent())
	}
	if req.GetServiceId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "service_id must be provided")
	}
	if req.GetService() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "service must be provided")
	}

	reqName := req.GetParent() + "/services/" + req.GetServiceId()
	name, err := s.parseServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetService()).(*pb.Service)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if obj.GetMcpServerSpec() != nil {
		obj.RegistryResource = fmt.Sprintf("projects/%d/locations/%s/mcpServers/agentregistry-00000000-0000-0000-0000-000000000001", name.Project.Number, name.Location)
	} else if obj.GetAgentSpec() != nil {
		obj.RegistryResource = fmt.Sprintf("projects/%d/locations/%s/agents/agentregistry-00000000-0000-0000-0000-000000000001", name.Project.Number, name.Location)
	} else if obj.GetEndpointSpec() != nil {
		obj.RegistryResource = fmt.Sprintf("projects/%d/locations/%s/endpoints/agentregistry-00000000-0000-0000-0000-000000000001", name.Project.Number, name.Location)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *AgentRegistryServer) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServiceName(req.GetService().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	service := req.GetService()
	paths := []string{}
	if req.GetUpdateMask() != nil {
		paths = req.GetUpdateMask().GetPaths()
	}

	if len(paths) == 0 {
		obj.DisplayName = service.GetDisplayName()
		obj.Description = service.GetDescription()
		obj.Interfaces = service.GetInterfaces()
		if service.Spec != nil {
			obj.Spec = service.Spec
		}
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = service.GetDisplayName()
			case "description":
				obj.Description = service.GetDescription()
			case "interfaces":
				obj.Interfaces = service.GetInterfaces()
			case "agent_spec", "agentSpec":
				if service.GetAgentSpec() != nil {
					obj.Spec = &pb.Service_AgentSpec_{AgentSpec: service.GetAgentSpec()}
				}
			case "mcp_server_spec", "mcpServerSpec":
				if service.GetMcpServerSpec() != nil {
					obj.Spec = &pb.Service_McpServerSpec_{McpServerSpec: service.GetMcpServerSpec()}
				}
			case "endpoint_spec", "endpointSpec":
				if service.GetEndpointSpec() != nil {
					obj.Spec = &pb.Service_EndpointSpec_{EndpointSpec: service.GetEndpointSpec()}
				}
			default:
				return nil, status.Errorf(codes.InvalidArgument, "unsupported update_mask path %q", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *AgentRegistryServer) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServiceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Service{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
