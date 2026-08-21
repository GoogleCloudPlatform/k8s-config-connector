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

package mockmemcache

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/memcache/apiv1beta2/memcachepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type MemcacheServer struct {
	*MockService
	pb.UnimplementedCloudMemcacheServer
}

func (s *MemcacheServer) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Instance %s not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MemcacheServer) ListInstances(ctx context.Context, req *pb.ListInstancesRequest) (*pb.ListInstancesResponse, error) {
	// Simplified implementation
	var resources []*pb.Instance
	kind := (&pb.Instance{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		resources = append(resources, obj.(*pb.Instance))
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.ListInstancesResponse{
		Resources: resources,
	}, nil
}

func (s *MemcacheServer) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Parent + "/instances/" + req.InstanceId)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.CloneOf(req.Resource)
	obj.Name = fqn

	t := timestamppb.New(time.Now())
	obj.CreateTime = t
	obj.UpdateTime = t
	obj.State = pb.Instance_CREATING

	if obj.MemcacheVersion == pb.MemcacheVersion_MEMCACHE_VERSION_UNSPECIFIED {
		obj.MemcacheVersion = pb.MemcacheVersion_MEMCACHE_1_5
	}
	obj.MemcacheFullVersion = "memcached-1.5.16"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Memcache creation is async, but we'll just make it ready for now in the mock
	obj.State = pb.Instance_READY
	s.populateNodes(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: t,
		Target:     fqn,
		Verb:       "create",
	}
	return s.operations.DoneLRO(ctx, fqn, metadata, obj)
}

func (s *MemcacheServer) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Resource.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Simplified update logic
	paths := req.UpdateMask.GetPaths()
	if len(paths) == 0 {
		// If no mask, update everything? GCP usually requires a mask.
	}

	for _, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.Resource.DisplayName
		case "parameters":
			obj.Parameters = req.Resource.Parameters
		case "nodeCount":
			obj.NodeCount = req.Resource.NodeCount
			s.populateNodes(obj)
		case "labels":
			obj.Labels = req.Resource.Labels
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "update",
	}
	return s.operations.DoneLRO(ctx, fqn, metadata, obj)
}

func (s *MemcacheServer) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	oldObj := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.DoneLRO(ctx, fqn, metadata, nil)
}

func (s *MemcacheServer) populateNodes(obj *pb.Instance) {
	obj.MemcacheNodes = nil
	for i := int32(0); i < obj.NodeCount; i++ {
		nodeID := fmt.Sprintf("node-%d", i)
		zone := ""
		if len(obj.Zones) > 0 {
			zone = obj.Zones[int(i)%len(obj.Zones)]
		}
		obj.MemcacheNodes = append(obj.MemcacheNodes, &pb.Instance_Node{
			NodeId: nodeID,
			Zone:   zone,
			State:  pb.Instance_Node_READY,
			Host:   fmt.Sprintf("10.0.0.%d", i+1),
			Port:   11211,
		})
	}
}

type instanceName struct {
	Project  *projects.ProjectData
	Location string
	Instance string
}

func (n *instanceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", n.Project.ID, n.Location, n.Instance)
}

func (s *MockService) parseInstanceName(name string) (*instanceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &instanceName{
			Project:  project,
			Location: tokens[3],
			Instance: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid instance name %q", name)
}
