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

package mockedgenetwork

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/edgenetwork/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *EdgenetworkV1) GetNetwork(ctx context.Context, req *pb.GetNetworkRequest) (*pb.Network, error) {
	name, err := s.parseNetworkName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Network{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *EdgenetworkV1) CreateNetwork(ctx context.Context, req *pb.CreateNetworkRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/networks/" + req.NetworkId
	name, err := s.parseNetworkName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Network).(*pb.Network)
	obj.Name = fqn

	if obj.GetMtu() != 9000 && obj.GetMtu() != 1500 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid network MTU: %v", obj.GetMtu())
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *EdgenetworkV1) DeleteNetwork(ctx context.Context, req *pb.DeleteNetworkRequest) (*longrunning.Operation, error) {
	name, err := s.parseNetworkName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Network must not have any subnets depending on it
	subnetKind := (&pb.Subnet{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, subnetKind, storage.ListOptions{}, func(obj proto.Message) error {
		subnet := obj.(*pb.Subnet)
		if subnet.GetNetwork() == fqn {
			return status.Errorf(codes.FailedPrecondition,
				"cannot delete network with attached subnet: %v", subnet.GetName())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	deletedObj := &pb.Network{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
