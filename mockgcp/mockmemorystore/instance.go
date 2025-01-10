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

package mockredis

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/redis/cluster/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocks"
)

type clusterServer struct {
	*MockService
	pb.UnimplementedCloudRedisClusterServer
}

func (r *clusterServer) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := r.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	retObj := proto.Clone(obj).(*pb.Cluster)
	// pscConfigs is not included in the response
	retObj.PscConfigs = nil
	return retObj, nil
}

func (r *clusterServer) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/clusters/%s", req.GetParent(), req.GetClusterId())
	name, err := r.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	obj.State = pb.Cluster_CREATING

	if err := r.populateDefaultsForCluster(name, obj); err != nil {
		return nil, err
	}

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		obj.State = pb.Cluster_ACTIVE

		if err := r.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		retObj := proto.Clone(obj).(*pb.Cluster)
		// pscConfigs is not included in the response
		retObj.PscConfigs = nil
		return retObj, nil
	})
}

func (s *clusterServer) populateDefaultsForCluster(name *clusterName, obj *pb.Cluster) error {
	if obj.AuthorizationMode == pb.AuthorizationMode_AUTH_MODE_UNSPECIFIED {
		obj.AuthorizationMode = pb.AuthorizationMode_AUTH_MODE_DISABLED
	}

	if obj.DeletionProtectionEnabled == nil {
		obj.DeletionProtectionEnabled = mocks.PtrTo(false)
	}

	if obj.NodeType == pb.NodeType_NODE_TYPE_UNSPECIFIED {
		obj.NodeType = pb.NodeType_REDIS_HIGHMEM_MEDIUM
	}

	if obj.DiscoveryEndpoints == nil {
		for _, pscConfig := range obj.PscConfigs {
			discoveryEndpoint := &pb.DiscoveryEndpoint{
				// The assigned addresses are (seemingly) not deterministic
				Address: fmt.Sprintf("10.128.0.%d", rand.IntN(100)),
				Port:    6379,
				PscConfig: &pb.PscConfig{
					Network: pscConfig.Network,
				},
			}
			obj.DiscoveryEndpoints = append(obj.DiscoveryEndpoints, discoveryEndpoint)
		}
	}

	if obj.PscConnections == nil {
		pscConnectionID := time.Now().UnixNano()

		for _, pscConfig := range obj.PscConfigs {
			for i := 0; i < 2; i++ {
				network, err := s.parseNetworkName(pscConfig.Network)
				if err != nil {
					return status.Errorf(codes.InvalidArgument, "unexpected format for network %q", pscConfig.Network)
				}
				pscConnectionID++
				forwardingRuleID := fmt.Sprintf("ssc-auto-fr-%x", pscConnectionID)
				pscConnection := &pb.PscConnection{
					// The assigned addresses are (seemingly) not deterministic
					Address:         fmt.Sprintf("10.128.0.%d", rand.IntN(100)),
					ForwardingRule:  fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/forwardingRules/%s", network.Project.ID, name.Location, forwardingRuleID),
					Network:         pscConfig.Network,
					ProjectId:       network.Project.ID,
					PscConnectionId: fmt.Sprintf("%d", pscConnectionID),
				}
				obj.PscConnections = append(obj.PscConnections, pscConnection)
			}
		}
	}

	if obj.PersistenceConfig == nil {
		obj.PersistenceConfig = &pb.ClusterPersistenceConfig{}
	}
	if obj.PersistenceConfig.Mode == pb.ClusterPersistenceConfig_PERSISTENCE_MODE_UNSPECIFIED {
		obj.PersistenceConfig.Mode = pb.ClusterPersistenceConfig_DISABLED
	}

	if obj.ReplicaCount == nil {
		obj.ReplicaCount = mocks.PtrTo[int32](0)
	}

	nodeCapacity := float64(1)
	switch obj.GetNodeType() {
	case pb.NodeType_REDIS_SHARED_CORE_NANO:
		nodeCapacity = 1.4
	case pb.NodeType_REDIS_STANDARD_SMALL:
		nodeCapacity = 6.5
	case pb.NodeType_REDIS_HIGHMEM_MEDIUM:
		nodeCapacity = 13.0
	case pb.NodeType_REDIS_HIGHMEM_XLARGE:
		nodeCapacity = 58.0
	default:
		return fmt.Errorf("unknown node type %v", obj.GetNodeType())
	}
	obj.PreciseSizeGb = mocks.PtrTo(float64(nodeCapacity * float64(obj.GetShardCount())))
	obj.SizeGb = mocks.PtrTo(int32(obj.GetPreciseSizeGb()))

	if obj.TransitEncryptionMode == pb.TransitEncryptionMode_TRANSIT_ENCRYPTION_MODE_UNSPECIFIED {
		obj.TransitEncryptionMode = pb.TransitEncryptionMode_TRANSIT_ENCRYPTION_MODE_DISABLED
	}
	if obj.Uid == "" {
		obj.Uid = fmt.Sprintf("%x", time.Now().UnixNano())
	}
	if obj.ZoneDistributionConfig == nil {
		obj.ZoneDistributionConfig = &pb.ZoneDistributionConfig{}
	}
	if obj.ZoneDistributionConfig.Mode == pb.ZoneDistributionConfig_ZONE_DISTRIBUTION_MODE_UNSPECIFIED {
		obj.ZoneDistributionConfig.Mode = pb.ZoneDistributionConfig_MULTI_ZONE
	}
	return nil
}

func (r *clusterServer) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunning.Operation, error) {
	reqName := req.GetCluster().GetName()

	name, err := r.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Cluster{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. Mask of fields to update. At least one path must be supplied in
	// this field. The elements of the repeated paths field may only include these
	// fields from [Cluster][mockgcp.cloud.redis.cluster.v1.Cluster]:
	//
	//   - `size_gb`
	//   - `replica_count`
	paths := req.GetUpdateMask().GetPaths()

	if req.Cluster == nil {
		return nil, status.Errorf(codes.InvalidArgument, "cluster is required")
	}

	if len(paths) != 1 {
		return nil, status.Errorf(codes.InvalidArgument, "exactly 1 update_mask field must be specified per update request")
	}

	if obj.GetPersistenceConfig().GetAofConfig() != nil && obj.GetPersistenceConfig().GetRdbConfig() != nil {
		return nil, status.Errorf(codes.InvalidArgument, "unable to update RDB and AOF config at the same time")
	}

	for _, path := range paths {
		switch path {
		case "sizeGb":
			obj.SizeGb = req.Cluster.SizeGb
		case "replicaCount":
			obj.ReplicaCount = req.Cluster.ReplicaCount
		case "shardCount":
			obj.ShardCount = req.Cluster.ShardCount
		case "deletionProtectionEnabled":
			obj.DeletionProtectionEnabled = req.Cluster.DeletionProtectionEnabled
		case "persistenceConfig":
			obj.PersistenceConfig = req.Cluster.PersistenceConfig
		case "redisConfigs":
			obj.RedisConfigs = req.Cluster.RedisConfigs

		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := r.populateDefaultsForCluster(name, obj); err != nil {
		return nil, err
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		retObj := proto.Clone(obj).(*pb.Cluster)
		// pscConfigs is not included in the response
		retObj.PscConfigs = nil
		return retObj, nil
	})
}

func (r *clusterServer) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunning.Operation, error) {
	name, err := r.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Cluster{}

	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if obj.GetDeletionProtectionEnabled() {
		return nil, status.Errorf(codes.FailedPrecondition, "The cluster is deletion protected. Please disable deletion protection to delete the cluster. To disable, update DeleteProtectionEnabled to false via the Update API")
	}

	deletedObj := &pb.Cluster{}
	if err := r.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type clusterName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *clusterName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clusters/" + n.Name
}

// parseClusterName parses a string into an clusterName.
// The expected form is `projects/*/locations/*/clusters/*`.
func (r *clusterServer) parseClusterName(name string) (*clusterName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &clusterName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type networkName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *networkName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/networks/" + n.Name
}

// parseNetworkName parses a string into a networkName.
// The expected form is `projects/*/global/networks/*`.
func (s *clusterServer) parseNetworkName(name string) (*networkName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
		projectObj, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &networkName{
			Project: projectObj,
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
