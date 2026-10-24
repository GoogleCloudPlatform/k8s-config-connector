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

package mockconnectors

import (
	"context"
	"fmt"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/connectors/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type connectorsServer struct {
	*MockService
	pb.UnimplementedConnectorsServer
}

func (s *connectorsServer) GetConnection(ctx context.Context, req *pb.GetConnectionRequest) (*pb.Connection, error) {
	name, err := s.parseConnectionName(req.GetName())
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

	return obj, nil
}

func (s *connectorsServer) CreateConnection(ctx context.Context, req *pb.CreateConnectionRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/connections/" + req.ConnectionId
	name, err := s.parseConnectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Connection).(*pb.Connection)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	s.populateConnectionDefaults(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		Target:                fqn,
		Verb:                  "create",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *connectorsServer) UpdateConnection(ctx context.Context, req *pb.UpdateConnectionRequest) (*longrunning.Operation, error) {
	name, err := s.parseConnectionName(req.GetConnection().GetName())
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

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, req.GetConnection(), paths); err != nil {
		return nil, fmt.Errorf("updating fields by mask: %w", err)
	}

	now := timestamppb.Now()
	obj.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		Target:                fqn,
		Verb:                  "update",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *connectorsServer) DeleteConnection(ctx context.Context, req *pb.DeleteConnectionRequest) (*longrunning.Operation, error) {
	name, err := s.parseConnectionName(req.GetName())
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

	deleted := &pb.Connection{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		Target:                fqn,
		Verb:                  "delete",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *connectorsServer) ListConnections(ctx context.Context, req *pb.ListConnectionsRequest) (*pb.ListConnectionsResponse, error) {
	name, err := s.parseLocationName(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := name.String() + "/connections/"

	var response pb.ListConnectionsResponse
	connectionKind := (&pb.Connection{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, connectionKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		response.Connections = append(response.Connections, obj.(*pb.Connection))
		return nil
	}); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *connectorsServer) populateConnectionDefaults(name *connectionName, obj *pb.Connection) {
	if obj.Status == nil {
		obj.Status = &pb.ConnectionStatus{
			State:       pb.ConnectionStatus_ERROR,
			Description: "Unable to create a connector instance.",
		}
	}

	if obj.NodeConfig == nil {
		obj.NodeConfig = &pb.NodeConfig{
			MinNodeCount: 2,
			MaxNodeCount: 50,
		}
	}

	obj.EnvoyImageLocation = "us.gcr.io/apigee-saas-prod-repo/thirdparty/connectors-envoy/master:1.0.12"
	obj.ImageLocation = "us-docker.pkg.dev/connectors-producer-prod/connectors-platform/connectors-service-gcp-gcs@sha256:65be052f7049a25dfc0e9e137b1d13533212a50c6e61b5d0c6d9e12334783f0b"
	obj.ServiceDirectory = fmt.Sprintf("projects/p47055cd3e3e1c55d-tp/locations/%s/namespaces/cloudrun/services/runtime", name.Location)
}
