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
// proto.service: google.cloud.managedkafka.v1.ManagedKafkaConnect
// proto.message: google.cloud.managedkafka.v1.ConnectCluster

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

	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type managedKafkaConnect struct {
	*MockService
	pb.UnimplementedManagedKafkaConnectServer
}

func (s *managedKafkaConnect) CreateConnectCluster(ctx context.Context, req *pb.CreateConnectClusterRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/connectClusters/%s", req.GetParent(), req.GetConnectClusterId())
	name, err := s.parseConnectClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetConnectCluster())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.ConnectCluster_CREATING
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		obj.State = pb.ConnectCluster_ACTIVE
		obj.UpdateTime = timestamppb.New(now)
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *managedKafkaConnect) GetConnectCluster(ctx context.Context, req *pb.GetConnectClusterRequest) (*pb.ConnectCluster, error) {
	name, err := s.parseConnectClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConnectCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *managedKafkaConnect) ListConnectClusters(ctx context.Context, req *pb.ListConnectClustersRequest) (*pb.ListConnectClustersResponse, error) {
	name, err := s.parseConnectClusterName(req.GetParent() + "/connectClusters/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListConnectClustersResponse{}

	findPrefix := strings.TrimSuffix(name.String(), "dummy")

	listKind := (&pb.ConnectCluster{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		cluster := obj.(*pb.ConnectCluster)
		if strings.HasPrefix(cluster.GetName(), findPrefix) {
			response.ConnectClusters = append(response.ConnectClusters, cluster)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *managedKafkaConnect) UpdateConnectCluster(ctx context.Context, req *pb.UpdateConnectClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseConnectClusterName(req.GetConnectCluster().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.ConnectCluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetConnectCluster().GetLabels()
		case "capacityConfig.memoryBytes":
			if obj.CapacityConfig == nil {
				obj.CapacityConfig = &pb.CapacityConfig{}
			}
			obj.CapacityConfig.MemoryBytes = req.GetConnectCluster().GetCapacityConfig().GetMemoryBytes()
		case "capacityConfig.vcpuCount":
			if obj.CapacityConfig == nil {
				obj.CapacityConfig = &pb.CapacityConfig{}
			}
			obj.CapacityConfig.VcpuCount = req.GetConnectCluster().GetCapacityConfig().GetVcpuCount()
		case "gcpConfig.accessConfig.networkConfigs":
			// In ManagedKafka, platformConfig is used, so GetGcpConfig returns it
			if obj.GetGcpConfig() == nil {
				obj.PlatformConfig = &pb.ConnectCluster_GcpConfig{GcpConfig: &pb.ConnectGcpConfig{}}
			}
			obj.GetGcpConfig().AccessConfig = req.GetConnectCluster().GetGcpConfig().GetAccessConfig()
		case "name":
			obj.Name = req.GetConnectCluster().GetName()
		case "config":
			obj.Config = req.GetConnectCluster().GetConfig()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		obj.State = pb.ConnectCluster_ACTIVE
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *managedKafkaConnect) DeleteConnectCluster(ctx context.Context, req *pb.DeleteConnectClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseConnectClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ConnectCluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
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
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type connectClusterName struct {
	Project        *projects.ProjectData
	Location       string
	ConnectCluster string
}

func (n *connectClusterName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/connectClusters/%s", n.Project.ID, n.Location, n.ConnectCluster)
}

func (s *MockService) parseConnectClusterName(name string) (*connectClusterName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connectClusters" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &connectClusterName{
			Project:        project,
			Location:       tokens[3],
			ConnectCluster: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid connect cluster name %q", name)
}
