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

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb"
)

type ClientConnectorServicesServer struct {
	*MockService
	pb.UnimplementedClientConnectorServicesServiceServer
}

func (s *ClientConnectorServicesServer) GetClientConnectorService(ctx context.Context, req *pb.GetClientConnectorServiceRequest) (*pb.ClientConnectorService, error) {
	name, err := s.parseClientConnectorServiceName(req.Name)
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

func (s *ClientConnectorServicesServer) CreateClientConnectorService(ctx context.Context, req *pb.CreateClientConnectorServiceRequest) (*longrunning.Operation, error) {
	parent := req.Parent
	id := req.ClientConnectorServiceId

	nameStr := fmt.Sprintf("%s/clientConnectorServices/%s", parent, id)
	name, err := s.parseClientConnectorServiceName(nameStr)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ClientConnectorService).(*pb.ClientConnectorService)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = pb.ClientConnectorService_RUNNING // Default to running for mock

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Beyondcorp APIs use LRO for creation
	metadata := &pb.ClientConnectorServiceOperationMetadata{
		CreateTime: now,
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, parent, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ClientConnectorServicesServer) UpdateClientConnectorService(ctx context.Context, req *pb.UpdateClientConnectorServiceRequest) (*longrunning.Operation, error) {
	obj := req.ClientConnectorService
	if obj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "client_connector_service is required")
	}

	name, err := s.parseClientConnectorServiceName(obj.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.ClientConnectorService{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// Update the fields based on update mask
	paths := req.UpdateMask.GetPaths()
	if len(paths) == 0 {
		// If empty, update all updateable fields
		paths = []string{"display_name", "ingress", "egress"}
	}

	for _, path := range paths {
		if path == "display_name" || path == "displayName" {
			existing.DisplayName = obj.DisplayName
		} else if path == "ingress" || strings.HasPrefix(path, "ingress.") {
			existing.Ingress = obj.Ingress
		} else if path == "egress" || strings.HasPrefix(path, "egress.") {
			existing.Egress = obj.Egress
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "unsupported field path %q", path)
		}
	}

	now := timestamppb.Now()
	existing.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	metadata := &pb.ClientConnectorServiceOperationMetadata{
		CreateTime: now,
		Target:     fqn,
		Verb:       "update",
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	return s.operations.StartLRO(ctx, parent, metadata, func() (proto.Message, error) {
		return existing, nil
	})
}

func (s *ClientConnectorServicesServer) DeleteClientConnectorService(ctx context.Context, req *pb.DeleteClientConnectorServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseClientConnectorServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.ClientConnectorService{}
	if err := s.storage.Delete(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	metadata := &pb.ClientConnectorServiceOperationMetadata{
		CreateTime: now,
		Target:     fqn,
		Verb:       "delete",
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	return s.operations.StartLRO(ctx, parent, metadata, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}
