// Copyright 2023 Google LLC
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

package mockcontainerattached

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkemulticloud/v1"
)

type ContainerAttachedV1 struct {
	*MockService
	pb.UnimplementedAttachedClustersServer
}

func (s *ContainerAttachedV1) GetAttachedCluster(ctx context.Context, req *pb.GetAttachedClusterRequest) (*pb.AttachedCluster, error) {
	name, err := s.parseAttachedClustersName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AttachedCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "attachedCluster %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading attachedCluster: %v", err)
		}
	}

	return obj, nil
}

func (s *ContainerAttachedV1) CreateAttachedCluster(ctx context.Context, req *pb.CreateAttachedClusterRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/attachedClusters/" + req.AttachedClusterId
	name, err := s.parseAttachedClustersName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.AttachedCluster).(*pb.AttachedCluster)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating attachedCluster: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *ContainerAttachedV1) DeleteAttachedCluster(ctx context.Context, req *pb.DeleteAttachedClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseAttachedClustersName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	attachedClusterKind := (&pb.AttachedCluster{}).ProtoReflect().Descriptor()
	if err := s.storage.Delete(ctx, attachedClusterKind, fqn); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "attachedCluster %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting attachedCluster: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
