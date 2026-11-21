// Copyright 2026 High Five Labs
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

// Attaching, detaching and
// listing the endpoints of a regional network endpoint group, adapted from
// this package's zonal NetworkEndpointGroupsV1 methods (networkendpointgroupsv1.go)
// for regions. Endpoints are keyed as the zonal ones are (buildEndpointKey).

package mockcompute

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *RegionNetworkEndpointGroupV1) regionalNEG(ctx context.Context, project, region, neg string) (*regionNetworkEndpointGroupName, string, *pb.NetworkEndpointGroup, error) {
	name, err := s.parseRegionNetworkEndpointGroupName("projects/" + project + "/regions/" + region + "/networkEndpointGroups/" + neg)
	if err != nil {
		return nil, "", nil, err
	}
	fqn := name.String()
	obj := &pb.NetworkEndpointGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, "", nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}
	return name, fqn, obj, nil
}

func (s *RegionNetworkEndpointGroupV1) updateRegionalNegSize(ctx context.Context, fqn string, neg *pb.NetworkEndpointGroup) error {
	var count int32
	kind := (&pb.NetworkEndpoint{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: fqn + "/networkEndpoints/"}, func(obj proto.Message) error {
		count++
		return nil
	}); err != nil {
		return err
	}
	neg.Size = &count
	return s.storage.Update(ctx, fqn, neg)
}

func (s *RegionNetworkEndpointGroupV1) AttachNetworkEndpoints(ctx context.Context, req *pb.AttachNetworkEndpointsRegionNetworkEndpointGroupRequest) (*pb.Operation, error) {
	name, fqn, neg, err := s.regionalNEG(ctx, req.GetProject(), req.GetRegion(), req.GetNetworkEndpointGroup())
	if err != nil {
		return nil, err
	}
	for _, endpoint := range req.GetRegionNetworkEndpointGroupsAttachEndpointsRequestResource().GetNetworkEndpoints() {
		key := buildEndpointKey(fqn, endpoint)
		if err := s.storage.Create(ctx, key, endpoint); err != nil {
			if status.Code(err) != codes.AlreadyExists {
				return nil, err
			}
			if err := s.storage.Update(ctx, key, endpoint); err != nil {
				return nil, err
			}
		}
	}
	_ = s.updateRegionalNegSize(ctx, fqn, neg)
	op := &pb.Operation{
		TargetId:      neg.Id,
		TargetLink:    neg.SelfLink,
		OperationType: PtrTo("compute.regionNetworkEndpointGroups.attachNetworkEndpoints"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return neg, nil
	})
}

func (s *RegionNetworkEndpointGroupV1) DetachNetworkEndpoints(ctx context.Context, req *pb.DetachNetworkEndpointsRegionNetworkEndpointGroupRequest) (*pb.Operation, error) {
	name, fqn, neg, err := s.regionalNEG(ctx, req.GetProject(), req.GetRegion(), req.GetNetworkEndpointGroup())
	if err != nil {
		return nil, err
	}
	for _, endpoint := range req.GetRegionNetworkEndpointGroupsDetachEndpointsRequestResource().GetNetworkEndpoints() {
		if err := s.storage.Delete(ctx, buildEndpointKey(fqn, endpoint), &pb.NetworkEndpoint{}); err != nil && status.Code(err) != codes.NotFound {
			return nil, err
		}
	}
	_ = s.updateRegionalNegSize(ctx, fqn, neg)
	op := &pb.Operation{
		TargetId:      neg.Id,
		TargetLink:    neg.SelfLink,
		OperationType: PtrTo("compute.regionNetworkEndpointGroups.detachNetworkEndpoints"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return neg, nil
	})
}

func (s *RegionNetworkEndpointGroupV1) ListNetworkEndpoints(ctx context.Context, req *pb.ListNetworkEndpointsRegionNetworkEndpointGroupsRequest) (*pb.NetworkEndpointGroupsListNetworkEndpoints, error) {
	_, fqn, neg, err := s.regionalNEG(ctx, req.GetProject(), req.GetRegion(), req.GetNetworkEndpointGroup())
	if err != nil {
		return nil, err
	}
	var items []*pb.NetworkEndpointWithHealthStatus
	kind := (&pb.NetworkEndpoint{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: fqn + "/networkEndpoints/"}, func(obj proto.Message) error {
		endpoint := proto.Clone(obj).(*pb.NetworkEndpoint)
		if endpoint.Instance != nil {
			inst := lastComponent(*endpoint.Instance)
			endpoint.Instance = &inst
		}
		items = append(items, &pb.NetworkEndpointWithHealthStatus{NetworkEndpoint: endpoint})
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.NetworkEndpointGroupsListNetworkEndpoints{
		Id:    PtrTo(fmt.Sprintf("%d", ValueOf(neg.Id))),
		Kind:  PtrTo("compute#networkEndpointGroupsListNetworkEndpoints"),
		Items: items,
	}, nil
}
