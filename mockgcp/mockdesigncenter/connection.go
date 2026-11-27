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

package mockdesigncenter

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *DesignCenterServer) GetConnection(ctx context.Context, req *pb.GetConnectionRequest) (*pb.Connection, error) {
	name, err := s.parseConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Connection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.populateConnection(obj), nil
}

func (s *DesignCenterServer) CreateConnection(ctx context.Context, req *pb.CreateConnectionRequest) (*pb.Connection, error) {
	parent, err := s.parseComponentName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Reject if parent component does not exist
	parentComponent := &pb.Component{}
	if err := s.storage.Get(ctx, parent.String(), parentComponent); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	if req.ConnectionId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "connectionId is required")
	}

	if req.Connection == nil {
		return nil, status.Errorf(codes.InvalidArgument, "connection is required")
	}

	connectionID := req.ConnectionId
	fqn := fmt.Sprintf("%s/connections/%s", parent.String(), connectionID)

	name, err := s.parseConnectionName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.Connection).(*pb.Connection)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj = s.populateConnection(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DesignCenterServer) UpdateConnection(ctx context.Context, req *pb.UpdateConnectionRequest) (*pb.Connection, error) {
	if req.Connection == nil {
		return nil, status.Errorf(codes.InvalidArgument, "connection is required")
	}

	name, err := s.parseConnectionName(req.Connection.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Connection{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Connection)
	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(updated, req.Connection, paths); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DesignCenterServer) DeleteConnection(ctx context.Context, req *pb.DeleteConnectionRequest) (*emptypb.Empty, error) {
	name, err := s.parseConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Connection{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DesignCenterServer) ListConnections(ctx context.Context, req *pb.ListConnectionsRequest) (*pb.ListConnectionsResponse, error) {
	parent, err := s.parseComponentName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/connections/"
	response := &pb.ListConnectionsResponse{}

	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	connectionKind := (&pb.Connection{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, connectionKind, findOpts, func(item proto.Message) error {
		obj := item.(*pb.Connection)
		response.Connections = append(response.Connections, s.populateConnection(obj))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DesignCenterServer) populateConnection(obj *pb.Connection) *pb.Connection {
	if len(obj.SourceComponentParameters) == 0 && obj.DestinationComponentUri != "" {
		// Real GCP automatically derives default sourceComponentParameters when connecting
		// to destination component if not explicitly provided.
		tokens := strings.Split(obj.DestinationComponentUri, "/")
		destCompID := tokens[len(tokens)-1]
		envSecretKey := strings.ReplaceAll(destCompID, "-", "_") + "_SECRET"
		envSecretVal := fmt.Sprintf("module.%s.env_vars.SECRET", destCompID)

		obj.SourceComponentParameters = []*pb.Parameter{
			{
				Key: "containers[0].env_secret_vars",
				Value: structpb.NewStructValue(&structpb.Struct{
					Fields: map[string]*structpb.Value{
						envSecretKey: structpb.NewStringValue(envSecretVal),
					},
				}),
			},
			{
				Key: "service_account_project_roles",
				Value: structpb.NewListValue(&structpb.ListValue{
					Values: []*structpb.Value{
						structpb.NewStringValue("roles/secretmanager.secretAccessor"),
					},
				}),
			},
		}
	}
	return obj
}
