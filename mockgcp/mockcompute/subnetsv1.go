// Copyright 2022 Google LLC
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
	"net"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/regions"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type SubnetsV1 struct {
	*MockService
	pb.UnimplementedSubnetworksServer
}

func (s *SubnetsV1) Get(ctx context.Context, req *pb.GetSubnetworkRequest) (*pb.Subnetwork, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Subnetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SubnetsV1) List(ctx context.Context, req *pb.ListSubnetworksRequest) (*pb.SubnetworkList, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), "placeholder")
	if err != nil {
		return nil, err
	}

	findPrefix := strings.TrimSuffix(name.String(), "placeholder")

	response := &pb.SubnetworkList{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#subnetworkList")
	response.SelfLink = PtrTo(buildComputeSelfLink(ctx, strings.TrimSuffix(findPrefix, "/")))

	findKind := (&pb.Subnetwork{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		subnet := obj.(*pb.Subnetwork)
		isMatch, err := matchFilter(req.GetFilter(), subnet)
		if err != nil {
			return err
		}
		if isMatch {
			response.Items = append(response.Items, subnet)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SubnetsV1) AggregatedList(ctx context.Context, req *pb.AggregatedListSubnetworksRequest) (*pb.SubnetworkAggregatedList, error) {
	name, err := s.newSubnetName(req.GetProject(), "-", "-")
	if err != nil {
		return nil, err
	}

	if req.GetFilter() != "" {
		return nil, fmt.Errorf("filter %q not implemented by mockgcp", req.GetFilter())
	}

	findPrefix := strings.TrimSuffix(name.String(), "-/subnetworks/-")

	response := &pb.SubnetworkAggregatedList{}
	response.Id = PtrTo("0123456789")
	response.Kind = PtrTo("compute#subnetworkAggregatedList")
	response.SelfLink = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/aggregated/subnetworks", name.Project.ID)))
	response.Items = make(map[string]*pb.SubnetworksScopedList)

	// Add empty lists for all regions and zones (and global).
	response.Items["global"] = &pb.SubnetworksScopedList{}
	for _, region := range regions.GetAllRegions(ctx) {
		response.Items["regions/"+region.Name] = &pb.SubnetworksScopedList{}

		for _, zone := range region.Zones(ctx) {
			response.Items["zones/"+zone] = &pb.SubnetworksScopedList{}
		}
	}

	// Actually populate the lists
	findKind := (&pb.Subnetwork{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		subnet := obj.(*pb.Subnetwork)

		key := fmt.Sprintf("regions/%s", lastComponent(subnet.GetRegion()))
		scopedList := response.Items[key]
		scopedList.Subnetworks = append(scopedList.Subnetworks, subnet)

		return nil
	}); err != nil {
		return nil, err
	}

	// Add a warning if there are no subnets in a region or zone.
	for name, scopedList := range response.Items {
		if len(scopedList.Subnetworks) != 0 {
			continue
		}

		scopedList.Warning = &pb.Warning{
			Code:    PtrTo("NO_RESULTS_ON_PAGE"),
			Message: PtrTo(fmt.Sprintf("No results for the scope '%s'.", name)),
			Data: []*pb.Data{
				{
					Key:   PtrTo("scope"),
					Value: PtrTo(name),
				},
			},
		}
	}

	return response, nil
}

func (s *SubnetsV1) Insert(ctx context.Context, req *pb.InsertSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), req.GetSubnetworkResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetSubnetworkResource()).(*pb.Subnetwork)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#subnetwork")
	// if obj.EnableFlowLogs == nil {
	// 	obj.EnableFlowLogs = PtrTo(false)
	// }
	if obj.PrivateIpGoogleAccess == nil {
		obj.PrivateIpGoogleAccess = PtrTo(false)
	}
	if obj.PrivateIpv6GoogleAccess == nil {
		obj.PrivateIpv6GoogleAccess = PtrTo("DISABLE_GOOGLE_ACCESS")
	}
	if obj.Purpose == nil {
		obj.Purpose = PtrTo("PRIVATE")
	}
	obj.Region = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))
	if obj.StackType == nil {
		obj.StackType = PtrTo("IPV4_ONLY")
	}
	networkName, err := s.parseNetworkSelfLink(obj.GetNetwork())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "network %q is not valid", obj.GetNetwork())
	}
	obj.Network = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/networks/%s", networkName.Project.ID, networkName.Name)))

	cidrIP, _, err := net.ParseCIDR(obj.GetIpCidrRange())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "ipCidrRange %q is not valid", obj.GetIpCidrRange())
	}
	cidrIP = cidrIP.To4()
	if cidrIP == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ipCidrRange %q is not valid", obj.GetIpCidrRange())
	}
	gatewayAddress := make(net.IP, len(cidrIP))
	copy(gatewayAddress, cidrIP)
	gatewayAddress[3] = 1
	obj.GatewayAddress = PtrTo(gatewayAddress.String())

	// obj.AllowSubnetCidrRoutesOverlap = PtrTo(false)
	obj.Fingerprint = PtrTo(computeFingerprint(obj))
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Add the subnetwork to the list in the network
	{
		networkFQN := networkName.String()
		network := &pb.Network{}
		if err := s.storage.Get(ctx, networkFQN, network); err != nil {
			return nil, err
		}

		network.Subnetworks = append(network.Subnetworks, obj.GetSelfLink())

		if err := s.storage.Update(ctx, networkFQN, network); err != nil {
			return nil, err
		}
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SubnetsV1) Delete(ctx context.Context, req *pb.DeleteSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Subnetwork{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	networkName, err := s.parseNetworkSelfLink(existing.GetNetwork())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "network %q is not valid", existing.GetNetwork())
	}

	deleted := &pb.Subnetwork{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// Remove the subnetwork from the list in the network
	{
		networkFQN := networkName.String()
		network := &pb.Network{}
		if err := s.storage.Get(ctx, networkFQN, network); err != nil {
			return nil, err
		}

		network.Subnetworks, _ = removeFromSlice(network.Subnetworks, deleted.GetSelfLink())

		if err := s.storage.Update(ctx, networkFQN, network); err != nil {
			return nil, err
		}
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *SubnetsV1) SetPrivateIpGoogleAccess(ctx context.Context, req *pb.SetPrivateIpGoogleAccessSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Subnetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.PrivateIpGoogleAccess = PtrTo(req.GetSubnetworksSetPrivateIpGoogleAccessRequestResource().GetPrivateIpGoogleAccess())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setPrivateIpGoogleAccess"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type subnetName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *subnetName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/subnetworks/" + n.Name
}

// parseSubnetName parses a string into a subnetName.
// The expected form is `projects/*/regions/*/subnetworks/*`.
func (s *MockService) parseSubnetName(name string) (*subnetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "subnetworks" {
		return s.newSubnetName(tokens[1], tokens[3], tokens[5])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

// newSubnetName builds a normalized subnetName from the constituent parts.
func (s *MockService) newSubnetName(project string, region string, name string) (*subnetName, error) {
	projectObj, err := s.Projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}

	return &subnetName{
		Project: projectObj,
		Region:  region,
		Name:    name,
	}, nil
}

func removeFromSlice[T comparable](s []T, removeValue T) ([]T, bool) {
	var keep []T
	removed := false
	for _, t := range s {
		if t == removeValue {
			removed = true
			continue
		}
		keep = append(keep, t)
	}
	return keep, removed
}
