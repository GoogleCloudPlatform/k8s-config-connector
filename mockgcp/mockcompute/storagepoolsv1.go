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

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type StoragePoolsV1 struct {
	*MockService
	pb.UnimplementedStoragePoolsServer
}

type zonalStoragePoolName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalStoragePoolName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/storagePools/" + n.Name
}

func (s *MockService) parseZonalStoragePoolName(name string) (*zonalStoragePoolName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "storagePools" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &zonalStoragePoolName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *StoragePoolsV1) Get(ctx context.Context, req *pb.GetStoragePoolRequest) (*pb.StoragePool, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/storagePools/%s", req.GetProject(), req.GetZone(), req.GetStoragePool())
	name, err := s.parseZonalStoragePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.StoragePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *StoragePoolsV1) Insert(ctx context.Context, req *pb.InsertStoragePoolRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/storagePools/%s", req.GetProject(), req.GetZone(), req.GetStoragePoolResource().GetName())
	name, err := s.parseZonalStoragePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetStoragePoolResource()).(*pb.StoragePool)
	selfLink := BuildComputeSelfLink(ctx, fqn)
	obj.SelfLink = PtrTo(selfLink)
	obj.SelfLinkWithId = PtrTo(selfLink)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#storagePool")
	obj.State = PtrTo("READY")
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))

	if obj.CapacityProvisioningType == nil {
		obj.CapacityProvisioningType = PtrTo("ADVANCED")
	}
	if obj.PerformanceProvisioningType == nil {
		obj.PerformanceProvisioningType = PtrTo("ADVANCED")
	}
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.StoragePoolType != nil {
		poolType := *obj.StoragePoolType
		if !strings.HasPrefix(poolType, "https://") {
			poolType = BuildComputeSelfLink(ctx, poolType)
		}
		obj.StoragePoolType = PtrTo(poolType)
	}

	var maxDiskCapacity int64
	if obj.PoolProvisionedCapacityGb != nil {
		maxDiskCapacity = *obj.PoolProvisionedCapacityGb * 10
	}
	obj.Status = &pb.StoragePoolResourceStatus{
		DiskCount:                         PtrTo(int64(0)),
		MaxTotalProvisionedDiskCapacityGb: PtrTo(maxDiskCapacity),
		PoolUsedCapacityBytes:             PtrTo(int64(0)),
		PoolUserWrittenBytes:              PtrTo(int64(0)),
		TotalProvisionedDiskCapacityGb:    PtrTo(int64(0)),
		TotalProvisionedDiskIops:          PtrTo(int64(0)),
		TotalProvisionedDiskThroughput:    PtrTo(int64(0)),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      &id,
		TargetLink:    PtrTo(selfLink),
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *StoragePoolsV1) Update(ctx context.Context, req *pb.UpdateStoragePoolRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/storagePools/%s", req.GetProject(), req.GetZone(), req.GetStoragePool())
	name, err := s.parseZonalStoragePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.StoragePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	reqResource := req.GetStoragePoolResource()
	if req.GetUpdateMask() != "" {
		for _, path := range strings.Split(req.GetUpdateMask(), ",") {
			path = strings.TrimSpace(path)
			switch path {
			case "pool_provisioned_capacity_gb", "poolProvisionedCapacityGb":
				obj.PoolProvisionedCapacityGb = reqResource.PoolProvisionedCapacityGb
			case "pool_provisioned_iops", "poolProvisionedIops":
				obj.PoolProvisionedIops = reqResource.PoolProvisionedIops
			case "pool_provisioned_throughput", "poolProvisionedThroughput":
				obj.PoolProvisionedThroughput = reqResource.PoolProvisionedThroughput
			case "description":
				obj.Description = reqResource.Description
			case "labels":
				obj.Labels = reqResource.Labels
			}
		}
	} else if reqResource != nil {
		proto.Merge(obj, reqResource)
	}

	if obj.PoolProvisionedCapacityGb != nil && obj.Status != nil {
		obj.Status.MaxTotalProvisionedDiskCapacityGb = PtrTo(*obj.PoolProvisionedCapacityGb * 10)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *StoragePoolsV1) Delete(ctx context.Context, req *pb.DeleteStoragePoolRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/storagePools/%s", req.GetProject(), req.GetZone(), req.GetStoragePool())
	name, err := s.parseZonalStoragePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.StoragePool{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *StoragePoolsV1) List(ctx context.Context, req *pb.ListStoragePoolsRequest) (*pb.StoragePoolList, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	zone := req.GetZone()
	prefix := fmt.Sprintf("projects/%s/zones/%s/storagePools/", project.ID, zone)

	response := &pb.StoragePoolList{
		Id:   PtrTo(fmt.Sprintf("projects/%s/zones/%s/storagePools", project.ID, zone)),
		Kind: PtrTo("compute#storagePoolList"),
	}

	findOptions := storage.ListOptions{
		Prefix: prefix,
	}

	if err := s.storage.List(ctx, (&pb.StoragePool{}).ProtoReflect().Descriptor(), findOptions, func(obj proto.Message) error {
		item := obj.(*pb.StoragePool)
		response.Items = append(response.Items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
