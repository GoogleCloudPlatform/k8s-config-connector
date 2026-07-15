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

	"github.com/google/uuid"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
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

	retObj := proto.CloneOf(obj)
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

	obj := proto.CloneOf(req.GetCluster())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	obj.State = pb.Cluster_CREATING

	if err := r.populateDefaultsForCluster(ctx, name, obj); err != nil {
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

		retObj := proto.CloneOf(obj)
		// pscConfigs is not included in the response
		retObj.PscConfigs = nil
		return retObj, nil
	})
}

func (r *clusterServer) syncReplication(ctx context.Context, obj *pb.Cluster) error {
	config := obj.GetCrossClusterReplicationConfig()
	if config == nil {
		return nil
	}

	if config.ClusterRole == pb.CrossClusterReplicationConfig_PRIMARY {
		membership := &pb.CrossClusterReplicationConfig_Membership{
			PrimaryCluster: &pb.CrossClusterReplicationConfig_RemoteCluster{
				Cluster: obj.Name,
				Uid:     obj.Uid,
			},
		}
		var secondaryClusters []*pb.CrossClusterReplicationConfig_RemoteCluster
		for _, sc := range config.SecondaryClusters {
			secondary := &pb.Cluster{}
			if err := r.storage.Get(ctx, sc.GetCluster(), secondary); err == nil {
				secondaryClusters = append(secondaryClusters, &pb.CrossClusterReplicationConfig_RemoteCluster{
					Cluster: secondary.Name,
					Uid:     secondary.Uid,
				})
			} else {
				secondaryClusters = append(secondaryClusters, sc)
			}
		}
		membership.SecondaryClusters = secondaryClusters
		config.Membership = membership
		config.SecondaryClusters = secondaryClusters
		config.PrimaryCluster = nil

		// Update secondaries
		for _, sc := range membership.SecondaryClusters {
			secondary := &pb.Cluster{}
			if err := r.storage.Get(ctx, sc.GetCluster(), secondary); err == nil {
				if secondary.CrossClusterReplicationConfig == nil {
					secondary.CrossClusterReplicationConfig = &pb.CrossClusterReplicationConfig{}
				}
				secondary.CrossClusterReplicationConfig.ClusterRole = pb.CrossClusterReplicationConfig_SECONDARY
				secondary.CrossClusterReplicationConfig.PrimaryCluster = membership.PrimaryCluster
				secondary.CrossClusterReplicationConfig.SecondaryClusters = nil
				secondary.CrossClusterReplicationConfig.Membership = membership
				secondary.CrossClusterReplicationConfig.UpdateTime = timestamppb.Now()
				if err := r.storage.Update(ctx, secondary.Name, secondary); err != nil {
					return err
				}
			}
		}
		// Save obj with updated membership
		if err := r.storage.Update(ctx, obj.Name, obj); err != nil {
			return err
		}
	} else if config.ClusterRole == pb.CrossClusterReplicationConfig_SECONDARY {
		pc := config.GetPrimaryCluster()
		if pc.GetCluster() == "" {
			return nil
		}
		primaryName := pc.GetCluster()
		primary := &pb.Cluster{}
		if err := r.storage.Get(ctx, primaryName, primary); err == nil {
			// Ensure primary knows about this secondary
			found := false
			if primary.CrossClusterReplicationConfig == nil {
				primary.CrossClusterReplicationConfig = &pb.CrossClusterReplicationConfig{
					ClusterRole: pb.CrossClusterReplicationConfig_PRIMARY,
				}
			}
			for _, sc := range primary.CrossClusterReplicationConfig.SecondaryClusters {
				if sc.GetCluster() == obj.Name {
					found = true
					sc.Uid = obj.Uid
					break
				}
			}
			if !found {
				primary.CrossClusterReplicationConfig.SecondaryClusters = append(primary.CrossClusterReplicationConfig.SecondaryClusters, &pb.CrossClusterReplicationConfig_RemoteCluster{
					Cluster: obj.Name,
					Uid:     obj.Uid,
				})
			}
			// Sync from primary's perspective
			return r.syncReplication(ctx, primary)
		}
	}

	// Reload obj to ensure it has the latest state (including membership)
	latest := &pb.Cluster{}
	if err := r.storage.Get(ctx, obj.Name, latest); err != nil {
		return err
	}
	proto.Reset(obj)
	proto.Merge(obj, latest)

	return nil
}

func (s *clusterServer) populateDefaultsForCluster(ctx context.Context, name *clusterName, obj *pb.Cluster) error {
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

	if obj.PscServiceAttachments == nil {
		suffix := "abcdef0123456789"
		obj.PscServiceAttachments = []*pb.PscServiceAttachment{
			{
				ServiceAttachment: fmt.Sprintf("projects/%d/regions/%s/serviceAttachments/gcp-memorystore-auto-%s-psc-sa", name.Project.Number, name.Location, suffix),
				ConnectionType:    pb.ConnectionType_CONNECTION_TYPE_DISCOVERY,
			},
			{
				ServiceAttachment: fmt.Sprintf("projects/%d/regions/%s/serviceAttachments/gcp-memorystore-auto-%s-psc-sa-2", name.Project.Number, name.Location, suffix),
			},
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

	if obj.CrossClusterReplicationConfig != nil {
		config := obj.CrossClusterReplicationConfig
		if config.UpdateTime == nil {
			config.UpdateTime = timestamppb.Now()
		}
		membership := &pb.CrossClusterReplicationConfig_Membership{}
		switch config.ClusterRole {
		case pb.CrossClusterReplicationConfig_PRIMARY:
			membership.PrimaryCluster = &pb.CrossClusterReplicationConfig_RemoteCluster{
				Cluster: obj.Name,
				Uid:     obj.Uid,
			}
			for _, secondary := range config.SecondaryClusters {
				secondaryCluster := &pb.CrossClusterReplicationConfig_RemoteCluster{
					Cluster: secondary.Cluster,
					Uid:     uuid.NewString(),
				}
				membership.SecondaryClusters = append(membership.SecondaryClusters, secondaryCluster)
			}
			config.PrimaryCluster = nil
			config.SecondaryClusters = membership.SecondaryClusters
		case pb.CrossClusterReplicationConfig_SECONDARY:
			if config.PrimaryCluster != nil {
				primaryCluster := &pb.CrossClusterReplicationConfig_RemoteCluster{
					Cluster: config.PrimaryCluster.GetCluster(),
					Uid:     uuid.NewString(),
				}
				membership.PrimaryCluster = primaryCluster
			}
			membership.SecondaryClusters = append(membership.SecondaryClusters, &pb.CrossClusterReplicationConfig_RemoteCluster{
				Cluster: obj.Name,
				Uid:     obj.Uid,
			})
			// This field is only set for a primary cluster.
			config.PrimaryCluster = membership.PrimaryCluster
			config.SecondaryClusters = nil
		}
		config.Membership = membership
	}
	if obj.AutomatedBackupConfig != nil {
		if obj.AutomatedBackupConfig.AutomatedBackupMode == pb.AutomatedBackupConfig_AUTOMATED_BACKUP_MODE_UNSPECIFIED {
			obj.AutomatedBackupConfig.AutomatedBackupMode = pb.AutomatedBackupConfig_DISABLED
		}
		if obj.AutomatedBackupConfig.AutomatedBackupMode == pb.AutomatedBackupConfig_DISABLED {
			obj.AutomatedBackupConfig.Schedule = nil
			obj.AutomatedBackupConfig.Retention = nil
		}
	} else {
		obj.AutomatedBackupConfig = &pb.AutomatedBackupConfig{AutomatedBackupMode: pb.AutomatedBackupConfig_DISABLED}
	}

	if obj.GetKmsKey() != "" {
		if obj.EncryptionInfo == nil {
			obj.EncryptionInfo = &pb.EncryptionInfo{
				EncryptionType:     pb.EncryptionInfo_CUSTOMER_MANAGED_ENCRYPTION,
				KmsKeyVersions:     []string{obj.GetKmsKey() + "/cryptoKeyVersions/1"},
				KmsKeyPrimaryState: pb.EncryptionInfo_ENABLED,
				LastUpdateTime:     timestamppb.Now(),
			}
		}
	} else {
		if obj.EncryptionInfo == nil {
			obj.EncryptionInfo = &pb.EncryptionInfo{
				EncryptionType: pb.EncryptionInfo_GOOGLE_DEFAULT_ENCRYPTION,
			}
		}
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
		case "automatedBackupConfig":
			obj.AutomatedBackupConfig = req.Cluster.AutomatedBackupConfig
		case "maintenancePolicy":
			obj.MaintenancePolicy = req.Cluster.MaintenancePolicy
		case "crossClusterReplicationConfig":
			obj.CrossClusterReplicationConfig = req.Cluster.CrossClusterReplicationConfig
		case "clusterEndpoints", "cluster_endpoints":
			obj.ClusterEndpoints = req.Cluster.ClusterEndpoints

		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := r.populateDefaultsForCluster(ctx, name, obj); err != nil {
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

		if err := r.syncReplication(ctx, obj); err != nil {
			return nil, err
		}

		retObj := proto.CloneOf(obj)
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

	if obj.GetCrossClusterReplicationConfig().GetClusterRole() == pb.CrossClusterReplicationConfig_PRIMARY {
		secondaryClusters := obj.CrossClusterReplicationConfig.SecondaryClusters
		if len(secondaryClusters) > 0 {
			return nil, status.Errorf(codes.FailedPrecondition, "Primary cluster %q cannot be deleted because it still has secondary clusters: %v", obj.GetName(), secondaryClusters)
		}
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
