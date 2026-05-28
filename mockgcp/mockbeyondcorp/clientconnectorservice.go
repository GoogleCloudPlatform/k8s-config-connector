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

package mockbeyondcorp

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/beyondcorp/clientconnectorservices/v1"
)

type ClientConnectorServicesV1 struct {
	*MockService
	pb.UnimplementedClientConnectorServicesServiceServer
}

type clientConnectorServiceName struct {
	Project  *projects.ProjectData
	Location string
	ID       string
}

func (n *clientConnectorServiceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/clientConnectorServices/%s", n.Project.ID, n.Location, n.ID)
}

func (s *MockService) parseName(name string) (*clientConnectorServiceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clientConnectorServices" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &clientConnectorServiceName{
			Project:  project,
			Location: tokens[3],
			ID:       tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}

func (s *ClientConnectorServicesV1) GetClientConnectorService(ctx context.Context, req *pb.GetClientConnectorServiceRequest) (*pb.ClientConnectorService, error) {
	name, err := s.parseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ClientConnectorService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ClientConnectorServicesV1) CreateClientConnectorService(ctx context.Context, req *pb.CreateClientConnectorServiceRequest) (*longrunning.Operation, error) {
	id := req.ClientConnectorServiceId
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "client_connector_service_id is required")
	}

	parent, err := s.parseName(req.Parent + "/clientConnectorServices/" + id)
	if err != nil {
		return nil, err
	}

	fqn := parent.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.ClientConnectorService).(*pb.ClientConnectorService)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = pb.ClientConnectorService_RUNNING

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, fqn, nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ClientConnectorServicesV1) UpdateClientConnectorService(ctx context.Context, req *pb.UpdateClientConnectorServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseName(req.ClientConnectorService.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ClientConnectorService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	obj.UpdateTime = now

	updateMask := req.UpdateMask
	if updateMask == nil {
		updateMask = &fieldmaskpb.FieldMask{Paths: []string{"*"}}
	}

	isSpecified := func(path string) bool {
		for _, p := range updateMask.Paths {
			if p == path || p == "*" || strings.HasPrefix(p, path+".") {
				return true
			}
		}
		return false
	}

	if isSpecified("display_name") {
		obj.DisplayName = req.ClientConnectorService.DisplayName
	}
	if isSpecified("ingress") {
		obj.Ingress = req.ClientConnectorService.Ingress
	}
	if isSpecified("egress") {
		obj.Egress = req.ClientConnectorService.Egress
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, fqn, nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ClientConnectorServicesV1) DeleteClientConnectorService(ctx context.Context, req *pb.DeleteClientConnectorServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.ClientConnectorService{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, fqn, nil, func() (proto.Message, error) {
		return deleted, nil
	})
}
