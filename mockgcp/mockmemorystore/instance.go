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
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/memorystore/v1"
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
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Instance_CREATING

	if err := r.populateDefaultsForInstance(name, obj); err != nil {
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

		retObj := proto.Clone(obj).(*pb.Instance)
		retObj.State = pb.Instance_ACTIVE
		r.storage.Update(ctx, fqn, retObj)
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

	if obj.Mode == pb.Instance_MODE_UNSPECIFIED {
		obj.Mode = pb.Instance_CLUSTER
	}
	if len(obj.PscAttachmentDetails) == 0 {
		var types []pb.ConnectionType
		switch obj.GetMode() {
		case pb.Instance_CLUSTER:
			types = append(types, pb.ConnectionType_CONNECTION_TYPE_DISCOVERY)
			types = append(types, pb.ConnectionType_CONNECTION_TYPE_UNSPECIFIED)
		default:
			types = append(types, pb.ConnectionType_CONNECTION_TYPE_PRIMARY)
			types = append(types, pb.ConnectionType_CONNECTION_TYPE_READER)
		}
		obj.PscAttachmentDetails = []*pb.PscAttachmentDetail{
			{
				ServiceAttachment: fmt.Sprintf("projects/tp-%s/regions/%s/serviceAttachments/sa-%s", name.Name, name.Location, name.Name),
				ConnectionType:    types[0],
			},
			{
				ServiceAttachment: fmt.Sprintf("projects/tp-%s/regions/%s/serviceAttachments/sa-%s-2", name.Name, name.Location, name.Name),
				ConnectionType:    types[1],
			},
		}
	}
	if len(obj.Endpoints) > 0 {
		if obj.Endpoints[0] != nil && len(obj.Endpoints[0].Connections) > 0 {
			connections := obj.Endpoints[0].Connections
			if len(connections) == 1 {
				autoConnection := connections[0].GetPscAutoConnection()
				if autoConnection != nil {
					obj.Endpoints[0].Connections = append(obj.Endpoints[0].Connections, &pb.Instance_ConnectionDetail{
						Connection: &pb.Instance_ConnectionDetail_PscAutoConnection{
							PscAutoConnection: proto.Clone(autoConnection).(*pb.PscAutoConnection),
						},
					})
				}
			}
		}
		pscConnectionID := int64(1234567890123456789)
		for _, endpoint := range obj.Endpoints {
			for i, connections := range endpoint.Connections {
				attachmentDetails := obj.PscAttachmentDetails[i%2]
				if connections.GetPscAutoConnection() != nil {
					autoConnection := connections.GetPscAutoConnection()
					network, err := s.parseNetworkName(autoConnection.GetNetwork())
					if err != nil {
						return status.Errorf(codes.InvalidArgument, "unexpected format for network %q", autoConnection.Network)
					}
					autoConnection.IpAddress = fmt.Sprintf("10.128.0.%d", pscConnectionID%256)
					autoConnection.ForwardingRule = fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/regions/%s/forwardingRules/sca-auto-fr-%x", network.Project.ID, name.Location, pscConnectionID)
					autoConnection.PscConnectionId = fmt.Sprintf("%d", pscConnectionID)
					autoConnection.ConnectionType = attachmentDetails.ConnectionType
					autoConnection.ServiceAttachment = attachmentDetails.ServiceAttachment
					if autoConnection.Ports == nil && autoConnection.ConnectionType != pb.ConnectionType_CONNECTION_TYPE_UNSPECIFIED {
						autoConnection.Ports = &pb.PscAutoConnection_Port{
							Port: 6379,
						}
					}
					pscConnectionID++
				}
				if connections.GetPscConnection() != nil {
					userConnection := connections.GetPscConnection()
					network, err := s.parseNetworkName(userConnection.GetNetwork())
					if err != nil {
						return status.Errorf(codes.InvalidArgument, "unexpected format for network %q", userConnection.Network)
					}
					userConnection.ProjectId = network.Project.ID
					userConnection.PscConnectionStatus = pb.PscConnectionStatus_ACTIVE
					userConnection.ConnectionType = attachmentDetails.ConnectionType
					if userConnection.Ports == nil && userConnection.ConnectionType != pb.ConnectionType_CONNECTION_TYPE_UNSPECIFIED {
						userConnection.Ports = &pb.PscConnection_Port{
							Port: 6379,
						}
					}
				}
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
	if obj.NodeConfig == nil {
		obj.NodeConfig = &pb.NodeConfig{}
	}
	nodeCapacity := float64(1)
	switch obj.GetNodeType() {
	case pb.Instance_SHARED_CORE_NANO:
		nodeCapacity = 1.4
	case pb.Instance_STANDARD_SMALL:
		nodeCapacity = 6.5
	case pb.Instance_HIGHMEM_MEDIUM:
		nodeCapacity = 13.0
	case pb.Instance_HIGHMEM_XLARGE:
		nodeCapacity = 58.0
	default:
		return fmt.Errorf("unknown node type %v", obj.GetNodeType())
	}
	obj.NodeConfig.SizeGb = *mocks.PtrTo(float64(nodeCapacity))

	if obj.TransitEncryptionMode == pb.Instance_TRANSIT_ENCRYPTION_MODE_UNSPECIFIED {
		obj.TransitEncryptionMode = pb.Instance_TRANSIT_ENCRYPTION_DISABLED
	}
	if obj.Uid == "" {
		obj.Uid = fmt.Sprintf("instance-%s", name.Name)
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
	if obj.AutomatedBackupConfig == nil {
		obj.AutomatedBackupConfig = &pb.AutomatedBackupConfig{}
	}
	if obj.AutomatedBackupConfig.AutomatedBackupMode == pb.AutomatedBackupConfig_AUTOMATED_BACKUP_MODE_UNSPECIFIED {
		obj.AutomatedBackupConfig.AutomatedBackupMode = pb.AutomatedBackupConfig_DISABLED
	}
	if obj.MaintenancePolicy != nil && obj.MaintenancePolicy.CreateTime == nil {
		obj.MaintenancePolicy.CreateTime = timestamppb.New(time.Now())
		obj.MaintenancePolicy.UpdateTime = obj.MaintenancePolicy.CreateTime
	}
	if crr := obj.CrossInstanceReplicationConfig; crr != nil {
		switch crr.InstanceRole {
		case pb.CrossInstanceReplicationConfig_PRIMARY:
			if len(crr.SecondaryInstances) == 0 {
				return status.Errorf(codes.InvalidArgument, "no secondary instances specified")
			}
			crr.PrimaryInstance = nil
			crr.Membership = &pb.CrossInstanceReplicationConfig_Membership{
				PrimaryInstance: &pb.CrossInstanceReplicationConfig_RemoteInstance{
					Instance: name.String(),
					Uid:      obj.Uid,
				},
				SecondaryInstances: []*pb.CrossInstanceReplicationConfig_RemoteInstance{},
			}
			for _, secondaryInstance := range crr.SecondaryInstances {
				secondaryName, err := s.parseInstanceName(secondaryInstance.Instance)
				if err != nil {
					return err
				}
				secondaryInstance.Uid = fmt.Sprintf("instance-%s", secondaryName.Name)
				crr.Membership.SecondaryInstances = append(crr.Membership.SecondaryInstances, &pb.CrossInstanceReplicationConfig_RemoteInstance{
					Instance: secondaryInstance.Instance,
					Uid:      secondaryInstance.Uid,
				})
			}
		case pb.CrossInstanceReplicationConfig_SECONDARY:
			if crr.PrimaryInstance == nil {
				return status.Errorf(codes.InvalidArgument, "no primary instance specified")
			}
			primaryName, err := s.parseInstanceName(crr.PrimaryInstance.Instance)
			if err != nil {
				return err
			}
			crr.PrimaryInstance.Uid = fmt.Sprintf("instance-%s", primaryName.Name)
			crr.Membership = &pb.CrossInstanceReplicationConfig_Membership{
				PrimaryInstance: &pb.CrossInstanceReplicationConfig_RemoteInstance{
					Instance: crr.PrimaryInstance.Instance,
					Uid:      crr.PrimaryInstance.Uid,
				},
				SecondaryInstances: []*pb.CrossInstanceReplicationConfig_RemoteInstance{
					{
						Instance: name.String(),
						Uid:      obj.Uid,
					},
				},
			}
		case pb.CrossInstanceReplicationConfig_NONE, pb.CrossInstanceReplicationConfig_INSTANCE_ROLE_UNSPECIFIED:
			obj.CrossInstanceReplicationConfig = nil
		default:
			return fmt.Errorf("unknown instance role %v", crr.InstanceRole)
		}
	}
	// PscAutoConnections is not included in the response
	obj.PscAutoConnections = nil
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
		case "labels":
			obj.Labels = req.Instance.Labels
		case "replicaCount":
			obj.ReplicaCount = req.Instance.ReplicaCount
		case "shardCount":
			obj.ShardCount = req.Instance.ShardCount
		case "nodeType":
			obj.NodeType = req.Instance.NodeType
		case "persistenceConfig":
			obj.PersistenceConfig = req.Instance.PersistenceConfig
		case "engineVersion":
			obj.EngineVersion = req.Instance.EngineVersion
		case "engineConfigs":
			obj.EngineConfigs = req.Instance.EngineConfigs
		case "deletionProtectionEnabled":
			obj.DeletionProtectionEnabled = req.Instance.DeletionProtectionEnabled
		case "endpoints":
			obj.Endpoints = req.Instance.Endpoints
		case "maintenancePolicy":
			if req.Instance.MaintenancePolicy != nil {
				obj.MaintenancePolicy = req.Instance.MaintenancePolicy
				obj.MaintenancePolicy.UpdateTime = timestamppb.New(now)
			}
		case "automatedBackupConfig":
			obj.AutomatedBackupConfig = req.Instance.AutomatedBackupConfig
		case "crossInstanceReplicationConfig":
			if req.Instance.CrossInstanceReplicationConfig != nil {
				obj.CrossInstanceReplicationConfig = req.Instance.CrossInstanceReplicationConfig
				obj.CrossInstanceReplicationConfig.UpdateTime = timestamppb.New(now)
			}
		case "gcsSource":
			if gcsSource := req.Instance.GetGcsSource(); gcsSource != nil {
				obj.ImportSources = &pb.Instance_GcsSource{GcsSource: gcsSource}
			}
		case "managedBackupSource":
			if managedBackupSource := req.Instance.GetManagedBackupSource(); managedBackupSource != nil {
				obj.ImportSources = &pb.Instance_ManagedBackupSource_{ManagedBackupSource: managedBackupSource}
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := r.populateDefaultsForInstance(name, obj); err != nil {
		return nil, err
	}

	obj.State = pb.Instance_UPDATING
	obj.UpdateTime = timestamppb.New(time.Now())
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

		retObj := proto.Clone(obj).(*pb.Instance)
		retObj.State = pb.Instance_ACTIVE
		retObj.UpdateTime = timestamppb.New(time.Now())
		r.storage.Update(ctx, fqn, retObj)
		return retObj, nil
	})
}

func (r *instanceServer) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
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

	obj.State = pb.Instance_DELETING
	if err := r.storage.Update(ctx, fqn, obj); err != nil {
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
		deletedObj := &pb.Instance{}
		r.storage.Delete(ctx, fqn, deletedObj)
		return &emptypb.Empty{}, nil
	})
}

func (r *instanceServer) GetBackup(ctx context.Context, req *pb.GetBackupRequest) (*pb.Backup, error) {
	name, err := r.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Backup{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	retObj := proto.Clone(obj).(*pb.Backup)
	return retObj, nil
}

func (r *instanceServer) BackupInstance(ctx context.Context, req *pb.BackupInstanceRequest) (*longrunning.Operation, error) {
	instanceName, err := r.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	instanceFqn := instanceName.String()
	instanceObj := &pb.Instance{}
	if err := r.storage.Get(ctx, instanceFqn, instanceObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", instanceFqn)
		}
		return nil, err
	}

	backupCollectionName := ""
	if instanceObj.BackupCollection != nil {
		backupCollectionName = *instanceObj.BackupCollection
	} else {
		backupCollectionName = fmt.Sprintf("projects/%s/locations/%s/backupCollections/backupCollection-%s", instanceName.Project.ID, instanceName.Location, instanceName.Name)
		instanceObj.BackupCollection = mocks.PtrTo(backupCollectionName)
		if err := r.storage.Update(ctx, instanceFqn, instanceObj); err != nil {
			return nil, err
		}
	}

	backupID := req.GetBackupId()
	if backupID == "" {
		backupID = time.Now().Format("20060102150405")
	}

	reqName := fmt.Sprintf("%s/backups/%s", backupCollectionName, backupID)
	name, err := r.parseBackupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	// Default TTL is 100 years
	ttl := durationpb.New(100 * 365 * 24 * time.Hour)
	if req.GetTtl() != nil {
		ttl = req.GetTtl()
	}

	obj := &pb.Backup{
		BackupType: pb.Backup_ON_DEMAND,
		BackupFiles: []*pb.BackupFile{
			{
				FileName:   fmt.Sprintf("file-%s.rdb", backupID),
				SizeBytes:  141,
				CreateTime: timestamppb.New(now),
			},
		},
		CreateTime:     timestamppb.New(now),
		ExpireTime:     timestamppb.New(now.Add(ttl.AsDuration())),
		EngineVersion:  instanceObj.EngineVersion,
		Instance:       instanceName.String(),
		InstanceUid:    instanceObj.Uid,
		Name:           fqn,
		NodeType:       instanceObj.NodeType,
		ShardCount:     instanceObj.ShardCount,
		State:          pb.Backup_CREATING,
		TotalSizeBytes: 141,
		Uid:            fmt.Sprintf("backup-%s", backupID),
	}
	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := name.OperationPrefix()
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     instanceName.String(),
		Verb:       "backup",
	}

	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		obj.State = pb.Backup_ACTIVE
		r.storage.Update(ctx, fqn, obj)
		return proto.Clone(instanceObj).(*pb.Instance), nil
	})
}

func (r *instanceServer) DeleteBackup(ctx context.Context, req *pb.DeleteBackupRequest) (*longrunning.Operation, error) {
	name, err := r.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Backup{}

	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	obj.State = pb.Backup_DELETING
	if err := r.storage.Update(ctx, fqn, obj); err != nil {
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
		deletedObj := &pb.Backup{}
		r.storage.Delete(ctx, fqn, deletedObj)
		return &emptypb.Empty{}, nil
	})
}

type instanceName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *instanceName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/instances/" + n.Name
}

// parseInstanceName parses a string into an instance name.
// The expected form is `projects/*/locations/*/instances/*`.
func (r *instanceServer) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
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

type backupName struct {
	Project          *projects.ProjectData
	Location         string
	BackupCollection string
	Name             string
}

func (n *backupName) String() string {
	return fmt.Sprintf("%s/backups/%s", n.Parent(), n.Name)
}

func (n *backupName) Parent() string {
	return fmt.Sprintf("%s/backupCollections/%s", n.OperationPrefix(), n.BackupCollection)
}

func (n *backupName) OperationPrefix() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

// parseBackupName parses a string into a backup name.
// The expected form is `projects/*/locations/*/backupCollections/*/backups/*`.
func (r *instanceServer) parseBackupName(name string) (*backupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupCollections" && tokens[6] == "backups" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupName{
			Project:          project,
			Location:         tokens[3],
			BackupCollection: tokens[5],
			Name:             tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
