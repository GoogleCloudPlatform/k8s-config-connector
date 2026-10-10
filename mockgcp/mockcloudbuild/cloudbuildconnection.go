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

package mockcloudbuild

import (
	"context"
	"fmt"
	"strings"

	pbv2 "cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ pbv2.RepositoryManagerServer = &CloudBuildConnectionServer{}

type CloudBuildConnectionServer struct {
	*MockService
	pbv2.UnimplementedRepositoryManagerServer
}

func (s *CloudBuildConnectionServer) GetConnection(ctx context.Context, req *pbv2.GetConnectionRequest) (*pbv2.Connection, error) {
	name, err := s.parseConnectionName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pbv2.Connection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *CloudBuildConnectionServer) CreateConnection(ctx context.Context, req *pbv2.CreateConnectionRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/connections/" + req.GetConnectionId()
	name, err := s.parseConnectionName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.GetConnection()).(*pbv2.Connection)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now

	s.populateDefaultsForConnection(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pbv2.OperationMetadata{
		ApiVersion: "v2",
		CreateTime: now,
		Target:     fqn,
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, name.Parent(), metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *CloudBuildConnectionServer) UpdateConnection(ctx context.Context, req *pbv2.UpdateConnectionRequest) (*longrunningpb.Operation, error) {
	reqConnection := req.GetConnection()
	name, err := s.parseConnectionName(reqConnection.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := timestamppb.Now()

	obj := &pbv2.Connection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(obj, reqConnection, paths); err != nil {
		return nil, err
	}

	obj.UpdateTime = now
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pbv2.OperationMetadata{
		ApiVersion: "v2",
		CreateTime: now,
		Target:     fqn,
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, name.Parent(), metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *CloudBuildConnectionServer) DeleteConnection(ctx context.Context, req *pbv2.DeleteConnectionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseConnectionName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := timestamppb.Now()

	oldObj := &pbv2.Connection{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	metadata := &pbv2.OperationMetadata{
		ApiVersion: "v2",
		CreateTime: now,
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, name.Parent(), metadata, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *CloudBuildConnectionServer) ListConnections(ctx context.Context, req *pbv2.ListConnectionsRequest) (*pbv2.ListConnectionsResponse, error) {
	parent := req.GetParent()

	response := &pbv2.ListConnectionsResponse{}
	connectionKind := (&pbv2.Connection{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, connectionKind, storage.ListOptions{
		Prefix: parent + "/connections",
	}, func(obj proto.Message) error {
		connection := obj.(*pbv2.Connection)
		response.Connections = append(response.Connections, connection)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *CloudBuildConnectionServer) populateDefaultsForConnection(obj *pbv2.Connection, name *connectionName) {
	obj.Etag = fields.ComputeWeakEtag(obj)
	obj.Reconciling = false

	if obj.GetGithubConfig() != nil && obj.GetInstallationState() == nil {
		actionURI := fmt.Sprintf("https://accounts.google.com/AccountChooser?continue=https%%3A%%2F%%2Fconsole.cloud.google.com%%2Fm%%2Fgcb%%2Fgithub%%2Flocations%%2F%s%%2Foauth_v2%%3Fconnection_name%%3Dprojects%%252F%d%%252Flocations%%252F%s%%252Fconnections%%252F%s",
			name.Location, name.Project.Number, name.Location, name.ConnectionID)
		obj.InstallationState = &pbv2.InstallationState{
			Stage:     pbv2.InstallationState_PENDING_USER_OAUTH,
			Message:   "Please log in to https://github.com using a robot account and then follow this link to authorize Cloud Build to access that account. After authorization, your GitHub authorization token will be stored in Cloud Secret Manager.",
			ActionUri: actionURI,
		}
	}
}

type connectionName struct {
	Project      *projects.ProjectData
	Location     string
	ConnectionID string
}

func (n *connectionName) String() string {
	return n.Parent() + "/connections/" + n.ConnectionID
}

func (n *connectionName) Parent() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
}

func (s *MockService) parseConnectionName(name string) (*connectionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connections" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		cn := &connectionName{
			Project:      project,
			Location:     tokens[3],
			ConnectionID: tokens[5],
		}
		return cn, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
