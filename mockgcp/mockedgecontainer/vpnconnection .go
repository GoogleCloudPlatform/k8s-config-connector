// Copyright 2022 Google LLC
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

package mockedgecontainer

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/edgecontainer/v1"
)

func (s *EdgeContainerV1) GetVpnConnection(ctx context.Context, req *pb.GetVpnConnectionRequest) (*pb.VpnConnection, error) {
	name, err := s.parseVpnConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.VpnConnection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "vpnConnection %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading vpnConnection: %v", err)
		}
	}

	return obj, nil
}

func (s *EdgeContainerV1) CreateVpnConnection(ctx context.Context, req *pb.CreateVpnConnectionRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/vpnConnections/" + req.VpnConnectionId
	name, err := s.parseVpnConnectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.VpnConnection).(*pb.VpnConnection)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating vpnConnection: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *EdgeContainerV1) DeleteVpnConnection(ctx context.Context, req *pb.DeleteVpnConnectionRequest) (*longrunning.Operation, error) {
	name, err := s.parseVpnConnectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.VpnConnection{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "vpnConnection %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting vpnConnection: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
