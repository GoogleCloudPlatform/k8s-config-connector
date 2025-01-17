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

package mockmemorystore

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/memorystore/v1beta"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocks"
)

type instanceServer struct {
	*MockService
	pb.UnimplementedMemorystoreServer
}

func (r *instanceServer) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := r.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	retObj := proto.Clone(obj).(*pb.Instance)
	// pscConfigs is not included in the response
	retObj.PscAutoConnections = nil
	return retObj, nil
}

func (r *instanceServer) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/instances/%s", req.GetParent(), req.GetInstanceId())
	name, err := r.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetInstance()).(*pb.Instance)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	obj.State = pb.Instance_CREATING

	if err := r.populateDefaultsForInstance(name, obj); err != nil {
		return nil, err
	}

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1beta",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		obj.State = pb.Instance_ACTIVE

		if err := r.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		retObj := proto.Clone(obj).(*pb.Instance)
		// pscConfigs is not included in the response
		retObj.PscAutoConnections = nil
		return retObj, nil
	})
}

func (s *instanceServer) populateDefaultsForInstance(name *instanceName, obj *pb.Instance) error {
	if obj.AuthorizationMode == pb.Instance_AUTHORIZATION_MODE_UNSPECIFIED {
		obj.AuthorizationMode = pb.Instance_AUTH_DISABLED
	}

	if obj.DeletionProtectionEnabled == nil {
		obj.DeletionProtectionEnabled = mocks.PtrTo(false)
	}

	if obj.NodeType == pb.Instance_NODE_TYPE_UNSPECIFIED {
		obj.NodeType = pb.Instance_HIGHMEM_MEDIUM
	}

	if obj.DiscoveryEndpoints == nil {
		for _, pscConfig := range obj.PscAutoConnections {
			discoveryEndpoint := &pb.DiscoveryEndpoint{
				// The assigned addresses are (seemingly) not deterministic
				Address: fmt.Sprintf("10.128.0.%d", rand.IntN(100)),
				Port:    6379,
				Network: pscConfig.Network,
			}
			obj.DiscoveryEndpoints = append(obj.DiscoveryEndpoints, discoveryEndpoint)
		}
	}

	if obj.PscAutoConnections == nil {
		pscConnectionID := time.Now().UnixNano()

		for _, pscConfig := range obj.PscAutoConnections {
			for i := 0; i < 2; i++ {
				network, err := s.parseNetworkName(pscConfig.Network)
				if err != nil {
					return status.Errorf(codes.InvalidArgument, "unexpected format for network %q", pscConfig.Network)
				}
				pscConnectionID++
				forwardingRuleID := fmt.Sprintf("ssc-auto-fr-%x", pscConnectionID)
				pscConnection := &pb.PscAutoConnection{
					// The assigned addresses are (seemingly) not deterministic
					IpAddress:       fmt.Sprintf("10.128.0.%d", rand.IntN(100)),
					ForwardingRule:  fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/forwardingRules/%s", network.Project.ID, name.Location, forwardingRuleID),
					Network:         pscConfig.Network,
					ProjectId:       network.Project.ID,
					PscConnectionId: fmt.Sprintf("%d", pscConnectionID),
				}
				obj.PscAutoConnections = append(obj.PscAutoConnections, pscConnection)
			}
		}
	}

	if obj.PersistenceConfig == nil {
		obj.PersistenceConfig = &pb.PersistenceConfig{}
	}
	if obj.PersistenceConfig.Mode == pb.PersistenceConfig_PERSISTENCE_MODE_UNSPECIFIED {
		obj.PersistenceConfig.Mode = pb.PersistenceConfig_DISABLED
	}

	if obj.ReplicaCount == nil {
		obj.ReplicaCount = mocks.PtrTo[int32](0)
	}

	// nodeCapacity := float64(1)
	// switch obj.GetNodeType() {
	// case pb.Instance_SHARED_CORE_NANO:
	// 	nodeCapacity = 1.4
	// case pb.Instance_STANDARD_SMALL:
	// 	nodeCapacity = 6.5
	// case pb.Instance_HIGHMEM_MEDIUM:
	// 	nodeCapacity = 13.0
	// case pb.Instance_HIGHMEM_XLARGE:
	// 	nodeCapacity = 58.0
	// default:
	// 	return fmt.Errorf("unknown node type %v", obj.GetNodeType())
	// }
	// obj.M = mocks.PtrTo(float64(nodeCapacity * float64(obj.GetShardCount())))
	// obj.SizeGb = mocks.PtrTo(int32(obj.GetPreciseSizeGb()))

	if obj.TransitEncryptionMode == pb.Instance_TRANSIT_ENCRYPTION_MODE_UNSPECIFIED {
		obj.TransitEncryptionMode = pb.Instance_TRANSIT_ENCRYPTION_DISABLED
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
	if obj.EngineVersion == "" {
		obj.EngineVersion = "VALKEY_7_2"
	}
	return nil
}

func (r *instanceServer) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.GetInstance().GetName()

	name, err := r.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Instance{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. Mask of fields to update. At least one path must be supplied in
	// this field. The elements of the repeated paths field may only include these
	// fields from [instance][mockgcp.cloud.redis.instance.v1.instance]:
	//
	//   - `size_gb`
	//   - `replica_count`
	paths := req.GetUpdateMask().GetPaths()

	if req.Instance == nil {
		return nil, status.Errorf(codes.InvalidArgument, "instance is required")
	}

	if len(paths) != 1 {
		return nil, status.Errorf(codes.InvalidArgument, "exactly 1 update_mask field must be specified per update request")
	}

	if obj.GetPersistenceConfig().GetAofConfig() != nil && obj.GetPersistenceConfig().GetRdbConfig() != nil {
		return nil, status.Errorf(codes.InvalidArgument, "unable to update RDB and AOF config at the same time")
	}

	for _, path := range paths {
		switch path {
		case "replicaCount":
			obj.ReplicaCount = req.Instance.ReplicaCount
		case "shardCount":
			obj.ShardCount = req.Instance.ShardCount
		case "deletionProtectionEnabled":
			obj.DeletionProtectionEnabled = req.Instance.DeletionProtectionEnabled
		case "persistenceConfig":
			obj.PersistenceConfig = req.Instance.PersistenceConfig
		case "engineConfigs":
			obj.EngineConfigs = req.Instance.EngineConfigs

		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := r.populateDefaultsForInstance(name, obj); err != nil {
		return nil, err
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1beta",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		retObj := proto.Clone(obj).(*pb.Instance)
		// pscConfigs is not included in the response
		retObj.PscAutoConnections = nil
		return retObj, nil
	})
}

func (r *instanceServer) DeleteCluster(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := r.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Instance{}

	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if obj.GetDeletionProtectionEnabled() {
		return nil, status.Errorf(codes.FailedPrecondition, "The instance is deletion protected. Please disable deletion protection to delete the instance. To disable, update DeleteProtectionEnabled to false via the Update API")
	}

	deletedObj := &pb.Instance{}
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

type instanceName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *instanceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clusters/" + n.Name
}

// parseClusterName parses a string into an clusterName.
// The expected form is `projects/*/locations/*/clusters/*`.
func (r *instanceServer) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &instanceName{
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
func (s *instanceServer) parseNetworkName(name string) (*networkName, error) {
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
