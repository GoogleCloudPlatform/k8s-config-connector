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

// +tool:mockgcp-support
// proto.service: google.cloud.hypercomputecluster.v1.HypercomputeCluster
// proto.message: google.cloud.hypercomputecluster.v1.Cluster

package mockhypercomputecluster

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/hypercomputecluster/apiv1/hypercomputeclusterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *HypercomputeClusterServer) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.parseClusterName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "cluster %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *HypercomputeClusterServer) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/clusters/%s", req.GetParent(), req.GetClusterId())
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Reconciling = false

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *HypercomputeClusterServer) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseClusterName(req.GetCluster().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	actual := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, actual); err != nil {
		return nil, err
	}

	desired := req.GetCluster()
	actual.Description = desired.Description
	actual.Labels = desired.Labels
	actual.NetworkResources = desired.NetworkResources
	actual.StorageResources = desired.StorageResources
	actual.ComputeResources = desired.ComputeResources
	actual.Orchestrator = desired.Orchestrator
	actual.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, actual); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return actual, nil
	})
}

func (s *HypercomputeClusterServer) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseClusterName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type clusterName struct {
	Project  *projects.ProjectData
	Location string
	Cluster  string
}

func (n *clusterName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clusters/" + n.Cluster
}

func (n *clusterName) ParentString() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
}

func (s *MockService) parseClusterName(name string) (*clusterName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &clusterName{
			Project:  project,
			Location: tokens[3],
			Cluster:  tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
