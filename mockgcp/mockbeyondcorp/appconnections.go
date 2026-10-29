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
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type appConnectionsServer struct {
	*MockService
	pb.UnimplementedAppConnectionsServiceServer
}

func (s *appConnectionsServer) GetAppConnection(ctx context.Context, req *pb.GetAppConnectionRequest) (*pb.AppConnection, error) {
	name, err := s.parseAppConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AppConnection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *appConnectionsServer) CreateAppConnection(ctx context.Context, req *pb.CreateAppConnectionRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/appConnections/%s", req.GetParent(), req.GetAppConnectionId())
	name, err := s.parseAppConnectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetAppConnection())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = uuid.NewString()
	obj.State = pb.AppConnection_CREATING

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.AppConnectionOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		obj.State = pb.AppConnection_CREATED
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *appConnectionsServer) UpdateAppConnection(ctx context.Context, req *pb.UpdateAppConnectionRequest) (*longrunning.Operation, error) {
	name, err := s.parseAppConnectionName(req.GetAppConnection().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AppConnection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetAppConnection().GetDisplayName()
		case "labels":
			obj.Labels = req.GetAppConnection().GetLabels()
		case "application_endpoint", "applicationEndpoint":
			obj.ApplicationEndpoint = req.GetAppConnection().GetApplicationEndpoint()
		case "connectors":
			obj.Connectors = req.GetAppConnection().GetConnectors()
		case "gateway":
			obj.Gateway = req.GetAppConnection().GetGateway()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.AppConnectionOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *appConnectionsServer) DeleteAppConnection(ctx context.Context, req *pb.DeleteAppConnectionRequest) (*longrunning.Operation, error) {
	name, err := s.parseAppConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.AppConnection{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.AppConnectionOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *appConnectionsServer) ListAppConnections(ctx context.Context, req *pb.ListAppConnectionsRequest) (*pb.ListAppConnectionsResponse, error) {
	name, err := s.parseAppConnectionName(req.GetParent() + "/appConnections/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListAppConnectionsResponse{}
	findPrefix := strings.TrimSuffix(name.String(), "dummy")
	listKind := (&pb.AppConnection{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		conn := obj.(*pb.AppConnection)
		if strings.HasPrefix(conn.GetName(), findPrefix) {
			response.AppConnections = append(response.AppConnections, conn)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
