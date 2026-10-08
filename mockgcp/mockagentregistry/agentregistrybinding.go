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

type AgentRegistryServer struct {
	*MockService
	pb.UnimplementedAgentRegistryServer
}

type bindingName struct {
	Project   *projects.ProjectData
	Location  string
	BindingID string
}

func (s *AgentRegistryServer) parseBindingName(name string) (*bindingName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}
	parts := strings.Split(name, "/")
	if len(parts) != 6 || parts[0] != "projects" || parts[2] != "locations" || parts[4] != "bindings" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid name format %q", name)
	}
	project, err := s.Projects.GetProjectByIDOrNumber(parts[1])
	if err != nil {
		return nil, err
	}
	return &bindingName{
		Project:   project,
		Location:  parts[3],
		BindingID: parts[5],
	}, nil
}

func (n *bindingName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/bindings/%s", n.Project.ID, n.Location, n.BindingID)
}

func (s *AgentRegistryServer) GetBinding(ctx context.Context, req *pb.GetBindingRequest) (*pb.Binding, error) {
	name, err := s.parseBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Binding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AgentRegistryServer) CreateBinding(ctx context.Context, req *pb.CreateBindingRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/bindings/" + req.BindingId
	name, err := s.parseBindingName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Binding).(*pb.Binding)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

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

func (s *AgentRegistryServer) UpdateBinding(ctx context.Context, req *pb.UpdateBindingRequest) (*longrunningpb.Operation, error) {
	binding := req.GetBinding()
	name, err := s.parseBindingName(binding.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Binding{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	paths := updateMask.GetPaths()
	if len(paths) == 0 {
		paths = []string{"display_name", "description", "auth_provider_binding", "source", "target"}
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = binding.GetDisplayName()
		case "description":
			obj.Description = binding.GetDescription()
		case "auth_provider_binding", "authProviderBinding":
			obj.Binding = binding.Binding
		case "source":
			obj.Source = binding.GetSource()
		case "target":
			obj.Target = binding.GetTarget()
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

func (s *AgentRegistryServer) DeleteBinding(ctx context.Context, req *pb.DeleteBindingRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBindingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Binding{}
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
