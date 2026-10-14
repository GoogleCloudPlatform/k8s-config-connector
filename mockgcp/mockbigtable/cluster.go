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

package mockbigtable

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	"cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func (s *instanceAdminServer) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
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

func (s *instanceAdminServer) ListClusters(ctx context.Context, req *pb.ListClustersRequest) (*pb.ListClustersResponse, error) {
	instanceName, err := s.parseInstanceName(req.GetParent())
	if err != nil {
		return nil, err
	}

	clusters, err := s.listClustersForInstance(ctx, instanceName)
	if err != nil {
		return nil, err
	}

	response := &pb.ListClustersResponse{}
	response.Clusters = clusters

	return response, nil
}

func (s *MockService) listClustersForInstance(ctx context.Context, instanceName *instanceName) ([]*pb.Cluster, error) {
	if instanceName.InstanceName == "-" {
		return nil, fmt.Errorf("mock does not implement ListClusters for wildcard instances")
	}

	var response []*pb.Cluster

	findKind := (&pb.Cluster{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: instanceName.String() + "/clusters/",
	}, func(obj proto.Message) error {
		cluster := obj.(*pb.Cluster)
		response = append(response, cluster)
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Name < response[j].Name
	})

	return response, nil
}

func (s *instanceAdminServer) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunning.Operation, error) {
	reqName := req.GetParent() + "/clusters/" + req.GetClusterId()
	clusterName, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	// instance, err := s.getInstance(ctx, clusterName.instanceName)
	// if err != nil {
	// 	return nil, err
	// }

	now := time.Now()

	clusterFQN := clusterName.String()

	obj := proto.CloneOf(req.Cluster)
	obj.Name = clusterFQN

	lroMetadata := &pb.CreateClusterMetadata{
		RequestTime: timestamppb.New(now),
	}

	// Populate LRO originalRequest, though not all fields are included
	lroMetadata.OriginalRequest = proto.CloneOf(req)
	if lroMetadata.OriginalRequest.Cluster.DefaultStorageType == pb.StorageType_STORAGE_TYPE_UNSPECIFIED {
		lroMetadata.OriginalRequest.Cluster.DefaultStorageType = pb.StorageType_SSD
	}

	lroPrefix := fmt.Sprintf("operations/%s/locations/%s", clusterName.String(), lastComponent(req.GetCluster().GetLocation()))

	if obj.ServeNodes != 0 && obj.GetClusterConfig().GetClusterAutoscalingConfig() != nil {
		return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
			return nil, status.Errorf(codes.Aborted, "Operation successfully rolled back: Both manual scaling (serve_nodes) and autoscaling (cluster_autoscaling_config) enabled. Exactly one must be set for CreateInstance/CreateCluster")
		})
	}

	if err := s.populateDefaultsForCluster(obj); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, clusterFQN, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.FinishTime = timestamppb.New(time.Now())

		returnObj := proto.CloneOf(obj)
		returnObj.NodeScalingFactor = adminpb.Cluster_NODE_SCALING_FACTOR_UNSPECIFIED // For some reason, not populated here
		return returnObj, nil
	})
}

func (s *instanceAdminServer) PartialUpdateCluster(ctx context.Context, req *pb.PartialUpdateClusterRequest) (*longrunning.Operation, error) {
	clusterName := req.GetCluster().GetName()

	name, err := s.parseClusterName(clusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()

	updateMask := req.GetUpdateMask()

	for _, path := range updateMask.GetPaths() {
		switch path {
		case "serve_nodes":
			obj.ServeNodes = req.GetCluster().GetServeNodes()

		case "cluster_config.cluster_autoscaling_config":
			if req.Cluster.GetClusterConfig().GetClusterAutoscalingConfig() == nil {
				if cc := obj.GetClusterConfig(); cc != nil {
					cc.ClusterAutoscalingConfig = nil
				}
			} else {
				if obj.Config == nil {
					obj.Config = &pb.Cluster_ClusterConfig_{
						ClusterConfig: &pb.Cluster_ClusterConfig{},
					}
				}
				clusterConfig := obj.Config.(*pb.Cluster_ClusterConfig_)
				if clusterConfig.ClusterConfig == nil {
					clusterConfig.ClusterConfig = &pb.Cluster_ClusterConfig{}
				}
				clusterConfig.ClusterConfig.ClusterAutoscalingConfig = req.Cluster.GetClusterConfig().GetClusterAutoscalingConfig()
			}

		default:
			return nil, fmt.Errorf("mock does not implement update of %q", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	zone := lastComponent(obj.GetLocation())
	lroPrefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)
	metadata := &pb.PartialUpdateClusterMetadata{
		RequestTime:     timestamppb.New(now),
		OriginalRequest: req,
	}
	return s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		metadata.FinishTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *instanceAdminServer) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*emptypb.Empty, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MockService) populateDefaultsForCluster(obj *pb.Cluster) error {
	obj.State = pb.Cluster_READY

	if obj.GetClusterConfig() != nil {
		autoscaling := obj.GetClusterConfig().ClusterAutoscalingConfig
		if autoscaling != nil {
			if autoscaling.AutoscalingTargets == nil {
				autoscaling.AutoscalingTargets = &pb.AutoscalingTargets{}
			}
			if autoscaling.AutoscalingTargets.CpuUtilizationPercent == 0 {
				autoscaling.AutoscalingTargets.CpuUtilizationPercent = 70
			}
			if autoscaling.AutoscalingTargets.StorageUtilizationGibPerNode == 0 {
				autoscaling.AutoscalingTargets.StorageUtilizationGibPerNode = 2560
			}
		}
	}

	if obj.ServeNodes == 0 {
		autoscaling := obj.GetClusterConfig().ClusterAutoscalingConfig
		if autoscaling != nil {
			obj.ServeNodes = autoscaling.AutoscalingLimits.GetMinServeNodes()
		}
		if obj.ServeNodes == 0 {
			obj.ServeNodes = 2
		}
	}

	if obj.DefaultStorageType == pb.StorageType_STORAGE_TYPE_UNSPECIFIED {
		obj.DefaultStorageType = pb.StorageType_SSD
	}

	return nil
}

type clusterName struct {
	instanceName
	ClusterName string
}

func (n *clusterName) String() string {
	return n.instanceName.String() + "/clusters/" + n.ClusterName
}

// parseClusterName parses a string into a clusterName.
// The expected form is projects/<projectID>/locations/global/clusters/<clusterName>
func (s *MockService) parseClusterName(name string) (*clusterName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[4] == "clusters" {
		instanceName, err := s.parseInstanceName(strings.Join(tokens[0:4], "/"))
		if err != nil {
			return nil, err
		}

		name := &clusterName{
			instanceName: *instanceName,
			ClusterName:  tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type memoryLayerName struct {
	cluster *clusterName
}

func (n *memoryLayerName) String() string {
	return n.cluster.String() + "/memoryLayer"
}

func (s *MockService) parseMemoryLayerName(name string) (*memoryLayerName, error) {
	if !strings.HasSuffix(name, "/memoryLayer") {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid for memoryLayer", name)
	}
	clusterPart := strings.TrimSuffix(name, "/memoryLayer")
	clusterN, err := s.parseClusterName(clusterPart)
	if err != nil {
		return nil, err
	}
	return &memoryLayerName{cluster: clusterN}, nil
}

func (s *instanceAdminServer) GetMemoryLayer(ctx context.Context, req *adminpb.GetMemoryLayerRequest) (*adminpb.MemoryLayer, error) {
	name, err := s.parseMemoryLayerName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &adminpb.MemoryLayer{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return &adminpb.MemoryLayer{
				Name:  req.GetName(),
				State: adminpb.MemoryLayer_DISABLED,
			}, nil
		}
		return nil, err
	}

	return obj, nil
}

func (s *instanceAdminServer) UpdateMemoryLayer(ctx context.Context, req *adminpb.UpdateMemoryLayerRequest) (*longrunning.Operation, error) {
	ml := req.GetMemoryLayer()
	if ml == nil {
		return nil, status.Errorf(codes.InvalidArgument, "memory_layer is required")
	}

	name, err := s.parseMemoryLayerName(ml.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Fetch the cluster to get location for LRO prefix.
	cluster := &pb.Cluster{}
	if err := s.storage.Get(ctx, name.cluster.String(), cluster); err != nil {
		return nil, err
	}

	existing := &adminpb.MemoryLayer{}
	exists := true
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}
		exists = false
		existing = &adminpb.MemoryLayer{
			Name:  ml.GetName(),
			State: adminpb.MemoryLayer_DISABLED,
		}
	}

	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.GetPaths() {
		switch path {
		case "memory_config":
			existing.MemoryConfig = ml.MemoryConfig
			if ml.MemoryConfig != nil {
				existing.State = adminpb.MemoryLayer_READY
			} else {
				existing.State = adminpb.MemoryLayer_DISABLED
			}
		default:
			return nil, fmt.Errorf("mock does not implement update of MemoryLayer field %q", path)
		}
	}

	if exists {
		if err := s.storage.Update(ctx, fqn, existing); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, fqn, existing); err != nil {
			return nil, err
		}
	}

	zone := lastComponent(cluster.GetLocation())
	lroPrefix := fmt.Sprintf("operations/%s/locations/%s", name.cluster.String(), zone)
	metadata := &adminpb.UpdateMemoryLayerMetadata{}
	return s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		return existing, nil
	})
}
