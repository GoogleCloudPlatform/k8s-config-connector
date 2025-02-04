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

// +tool:mockgcp-support
// proto.service: google.cloud.managedkafka.v1.ManagedKafka
// proto.message: google.cloud.managedkafka.v1.Cluster

package mockmanagedkafka

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/managedkafka/v1"
)

type managedKafka struct {
	*MockService
	pb.UnimplementedManagedKafkaServer
}

func (s *managedKafka) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/clusters/%s", req.GetParent(), req.GetClusterId())
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Cluster_CREATING
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Method:     "google.cloud.managedkafka.v1.ManagedKafka.CreateCluster",
		Target:     name.String(),
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		obj.State = pb.Cluster_ACTIVE
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *managedKafka) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "cluster %v not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseClusterName(req.GetCluster().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	proto.Merge(obj, req.GetCluster())
	obj.UpdateTime = timestamppb.Now()
	obj.State = pb.Cluster_UPDATING

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Method:     "google.cloud.managedkafka.v1.ManagedKafka.UpdateCluster",
		Target:     name.String(),
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		obj.State = pb.Cluster_ACTIVE
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *managedKafka) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Method:     "google.cloud.managedkafka.v1.ManagedKafka.DeleteCluster",
		Target:     name.String(),
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type clusterName struct {
	Project  *projects.ProjectData
	Location string
	Cluster  string
}

func (n *clusterName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/clusters/%s", n.Project.Number, n.Location, n.Cluster)
}

func (s *MockService) parseClusterName(name string) (*clusterName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &clusterName{
			Project:  project,
			Location: tokens[3],
			Cluster:  tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid cluster name %q", name)
}
