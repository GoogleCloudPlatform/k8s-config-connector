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

package mockgkemulticloud

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkemulticloud/v1"
)

type GKEMulticloudV1 struct {
	*MockService
	pb.UnimplementedAttachedClustersServer
}

func (s *GKEMulticloudV1) GetAttachedCluster(ctx context.Context, req *pb.GetAttachedClusterRequest) (*pb.AttachedCluster, error) {
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

func (s *GKEMulticloudV1) CreateAttachedCluster(ctx context.Context, req *pb.CreateAttachedClusterRequest) (*longrunning.Operation, error) {
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

func (s *GKEMulticloudV1) UpdateAttachedCluster(ctx context.Context, req *pb.UpdateAttachedClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseAttachedClustersName(req.GetAttachedCluster().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AttachedCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "attachedCluster %q not found", fqn)
		}
		return nil, status.Errorf(codes.Internal, "error reading attachedCluster: %v", err)
	}
	// Mask of fields to update. At least one path must be supplied in
	// this field. The elements of the repeated paths field can only include these
	// fields from
	// [AttachedCluster][mockgcp.cloud.gkemulticloud.v1.AttachedCluster]:
	//
	//   - `description`.
	//   - `annotations`.
	//   - `platform_version`.
	//   - `authorization.admin_users`.
	//   - `logging_config.component_config.enable_components`.
	//   - `monitoring_config.managed_prometheus_config.enabled`.

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "must specify updateMask")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetAttachedCluster().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating attachedCluster: %v", err)
	}
	return s.operations.NewLRO(ctx)
}

func (s *GKEMulticloudV1) DeleteAttachedCluster(ctx context.Context, req *pb.DeleteAttachedClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseAttachedClustersName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.AttachedCluster{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "attachedCluster %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting attachedCluster: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
