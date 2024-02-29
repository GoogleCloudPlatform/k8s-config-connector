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
)

func (s *EdgenetworkV1) GetSubnet(ctx context.Context, req *pb.GetSubnetRequest) (*pb.Subnet, error) {
	name, err := s.parseSubnetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Subnet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *EdgenetworkV1) CreateSubnet(ctx context.Context, req *pb.CreateSubnetRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/subnets/" + req.SubnetId
	name, err := s.parseSubnetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Subnet).(*pb.Subnet)
	obj.Name = fqn

	// Network reference must exist
	networkName := obj.GetNetwork()
	if _, err := s.GetNetwork(ctx, &pb.GetNetworkRequest{Name: networkName}); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "network not found: %v", err)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *EdgenetworkV1) DeleteSubnet(ctx context.Context, req *pb.DeleteSubnetRequest) (*longrunning.Operation, error) {
	name, err := s.parseSubnetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Subnet{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
