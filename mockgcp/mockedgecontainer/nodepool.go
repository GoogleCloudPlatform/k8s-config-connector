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

func (s *EdgeContainerV1) GetNodePool(ctx context.Context, req *pb.GetNodePoolRequest) (*pb.NodePool, error) {
	name, err := s.parseNodePoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.NodePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "nodePool %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading nodePool: %v", err)
		}
	}

	return obj, nil
}

func (s *EdgeContainerV1) CreateNodePool(ctx context.Context, req *pb.CreateNodePoolRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/nodePools/" + req.NodePoolId
	name, err := s.parseNodePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.NodePool).(*pb.NodePool)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating nodePool: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *EdgeContainerV1) DeleteNodePool(ctx context.Context, req *pb.DeleteNodePoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseNodePoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.NodePool{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "nodePool %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting nodePool: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
