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
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/redis/apiv1/redispb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type redisServer struct {
	*MockService
	pb.UnimplementedCloudRedisServer
}

func (r *redisServer) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
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

	return obj, nil
}

func (r *redisServer) GetInstanceAuthString(ctx context.Context, req *pb.GetInstanceAuthStringRequest) (*pb.InstanceAuthString, error) {
	name, err := r.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	response := &pb.InstanceAuthString{
		// We use a hard-coded auth string in our mock, until we discover something better is needed
		AuthString: "secret-squirrel",
	}
	return response, nil
}

func (r *redisServer) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/instances/%s", req.GetParent(), req.GetInstanceId())
	name, err := r.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.GetInstance())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	r.populateDefaultsForInstance(name, obj)

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:      "v1",
		CancelRequested: false,
		CreateTime:      timestamppb.New(now),
		Target:          fqn,
		Verb:            "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		obj.State = pb.Instance_READY

		if err := r.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (r *redisServer) populateDefaultsForInstance(name *instanceName, obj *pb.Instance) {
	zone := name.Location + "-a"
	obj.CurrentLocationId = zone
	obj.LocationId = zone

	obj.Host = "10.20.30.40"
	obj.ReservedIpRange = "10.20.30.0/24"

	obj.PersistenceIamIdentity = fmt.Sprintf("serviceAccount:service-%d@cloud-redis.iam.gserviceaccount.com", name.Project.Number)

	obj.Port = 6379

	if obj.RedisVersion == "" {
		obj.RedisVersion = "REDIS_7_0"
	}

	obj.State = pb.Instance_CREATING

	obj.Nodes = []*pb.NodeInfo{
		{
			Id:   "node-0",
			Zone: zone,
		},
	}

	// alternativeLocationId will be present in the instance response if and only if the instance is created as STANDARD_HA tier
	if obj.Tier == pb.Instance_STANDARD_HA && obj.AlternativeLocationId == "" {
		obj.AlternativeLocationId = zone
		obj.Nodes = append(obj.Nodes, &pb.NodeInfo{
			Id:   "node-1",
			Zone: zone,
		})
	}

	// The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2.
	// If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1.
	// The valid value for basic tier is 0 and the default is also 0.
	if obj.ReplicaCount == 0 && obj.Tier == pb.Instance_STANDARD_HA {
		obj.ReplicaCount = 1
		if obj.ReadReplicasMode == pb.Instance_READ_REPLICAS_ENABLED {
			obj.ReplicaCount = 2
		}
	}

	if obj.AuthorizedNetwork == "" {
		obj.AuthorizedNetwork = "projects/" + name.Project.ID + "/global/networks/default"
	}

	if obj.ConnectMode == pb.Instance_CONNECT_MODE_UNSPECIFIED {
		obj.ConnectMode = pb.Instance_DIRECT_PEERING
	}

	if obj.TransitEncryptionMode == pb.Instance_TRANSIT_ENCRYPTION_MODE_UNSPECIFIED {
		obj.TransitEncryptionMode = pb.Instance_DISABLED
	}

	if obj.PersistenceConfig == nil {
		obj.PersistenceConfig = &pb.PersistenceConfig{}
	}

	if obj.PersistenceConfig.PersistenceMode == pb.PersistenceConfig_PERSISTENCE_MODE_UNSPECIFIED {
		obj.PersistenceConfig.PersistenceMode = pb.PersistenceConfig_DISABLED
	}

	if obj.ReadReplicasMode == pb.Instance_READ_REPLICAS_MODE_UNSPECIFIED {
		obj.ReadReplicasMode = pb.Instance_READ_REPLICAS_DISABLED
	}
}

func (r *redisServer) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunning.Operation, error) {
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
	updated := proto.CloneOf(obj)

	// Required. Mask of fields to update. At least one path must be supplied in
	// this field. The elements of the repeated paths field may only include these
	// fields from Instance:
	//
	//  *   `displayName`
	//  *   `labels`
	//  *   `memorySizeGb`
	//  *   `redisConfig`
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "displayName", "display_name":
			updated.DisplayName = req.GetInstance().GetDisplayName()
		case "labels":
			updated.Labels = req.GetInstance().GetLabels()
		case "memorySizeGb", "memory_size_gb":
			updated.MemorySizeGb = req.GetInstance().GetMemorySizeGb()
		case "redisConfig", "redisConfigs", "redis_configs":
			updated.RedisConfigs = req.GetInstance().GetRedisConfigs()
		case "readReplicasMode", "read_replicas_mode":
			updated.ReadReplicasMode = req.GetInstance().GetReadReplicasMode()
			// SecondaryIpRange can only be set during Update call when enabling readReplicasMode, or it will be ignored.
			// See https://b.corp.google.com/issues/374126107#comment6
			updated.SecondaryIpRange = req.GetInstance().GetSecondaryIpRange()
			if updated.ReadReplicasMode == pb.Instance_READ_REPLICAS_ENABLED {
				if updated.SecondaryIpRange == "auto" {
					updated.SecondaryIpRange = "10.46.102.96/28"
				} else if updated.SecondaryIpRange != "" && !strings.Contains(updated.SecondaryIpRange, "/") {
					updated.SecondaryIpRange = "10.87.192.0/28"
				}
				if updated.SecondaryIpRange == "" {
					return nil, status.Errorf(codes.InvalidArgument, "Secondary IP Range is required when enabling read replicas on an existing instance.")
				}
			}
		case "secondaryIpRange", "secondary_ip_range":
			newVal := req.GetInstance().GetSecondaryIpRange()
			// SecondaryIpRange cannot be updated on instances that already have readReplicasMode enabled.
			if obj.ReadReplicasMode == pb.Instance_READ_REPLICAS_ENABLED {
				if newVal != obj.SecondaryIpRange {
					return nil, status.Errorf(codes.InvalidArgument, "Secondary IP Range can not be updated on instances that use read replicas")
				}
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if updated.SecondaryIpRange != "" {
		updated.ReadEndpoint = "10.57.59.77"
		updated.ReadEndpointPort = 6379
	}

	if err := r.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:      "v1",
		CancelRequested: false,
		CreateTime:      timestamppb.New(now),
		Target:          fqn,
		Verb:            "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return updated, nil
	})
}

func (r *redisServer) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := r.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	oldObj := &pb.Instance{}
	if err := r.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:      "v1",
		CancelRequested: false,
		CreateTime:      timestamppb.New(now),
		Target:          fqn,
		Verb:            "delete",
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
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/instances/" + n.Name
}

// parseInstanceName parses a string into an instanceName.
// The expected form is `projects/*/locations/*/instances/*`.
func (r *redisServer) parseInstanceName(name string) (*instanceName, error) {
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
