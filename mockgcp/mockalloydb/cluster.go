// Copyright 2024 Google LLC
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

package mockalloydb

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/alloydb/v1beta"
)

type AlloyDBAdminV1 struct {
	*MockService
	pb.UnimplementedAlloyDBAdminServer
}

func (s *AlloyDBAdminV1) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AlloyDBAdminV1) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/clusters/" + req.ClusterId
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Cluster).(*pb.Cluster)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AlloyDBAdminV1) CreateSecondaryCluster(ctx context.Context, req *pb.CreateSecondaryClusterRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/clusters/" + req.ClusterId
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Cluster).(*pb.Cluster)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AlloyDBAdminV1) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunning.Operation, error) {
	reqName := req.GetCluster().GetName()

	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.Cluster.GetDisplayName()
		case "automatedBackupPolicy":
			obj.AutomatedBackupPolicy = req.Cluster.GetAutomatedBackupPolicy()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AlloyDBAdminV1) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
