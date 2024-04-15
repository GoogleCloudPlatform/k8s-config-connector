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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	commonpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/common"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/redis/v1beta1"
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

	obj := proto.Clone(req.GetInstance()).(*pb.Instance)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	zone := name.Location + "-a"
	obj.CurrentLocationId = zone
	obj.LocationId = zone

	obj.Nodes = []*pb.NodeInfo{
		{
			Id:   "node-0",
			Zone: zone,
		},
	}

	obj.Host = "10.20.30.40"
	obj.ReservedIpRange = "10.20.30.0/24"

	obj.PersistenceIamIdentity = fmt.Sprintf("serviceAccount:service-%d@cloud-redis.iam.gserviceaccount.com", name.Project.Number)

	obj.Port = 6379

	if obj.RedisVersion == "" {
		obj.RedisVersion = "REDIS_7_0"
	}

	obj.State = pb.Instance_CREATING

	r.populateDefaultsForInstance(name, obj)

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &commonpb.OperationMetadata{
		ApiVersion:      "v1beta1",
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
	if obj.AuthorizedNetwork == "" {
		obj.AuthorizedNetwork = "projects/" + name.Project.ID + "/global/networks/default"
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
		case "displayName":
			obj.DisplayName = req.GetInstance().GetDisplayName()
		case "labels":
			obj.Labels = req.GetInstance().GetLabels()
		case "memorySizeGb":
			obj.MemorySizeGb = req.GetInstance().GetMemorySizeGb()
		case "redisConfig":
			obj.RedisConfigs = req.GetInstance().GetRedisConfigs()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &commonpb.OperationMetadata{
		ApiVersion:      "v1beta1",
		CancelRequested: false,
		CreateTime:      timestamppb.New(now),
		Target:          fqn,
		Verb:            "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
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

	metadata := &commonpb.OperationMetadata{
		ApiVersion:      "v1beta1",
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
