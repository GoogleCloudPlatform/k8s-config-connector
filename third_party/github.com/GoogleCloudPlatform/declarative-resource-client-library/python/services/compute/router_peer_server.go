// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for RouterPeer.
type RouterPeerServer struct{}

// ProtoToRouterPeerAdvertisedIPRanges converts a RouterPeerAdvertisedIPRanges resource from its proto representation.
func ProtoToComputeRouterPeerAdvertisedIPRanges(p *computepb.ComputeRouterPeerAdvertisedIPRanges) *compute.RouterPeerAdvertisedIPRanges {
	if p == nil {
		return nil
	}
	obj := &compute.RouterPeerAdvertisedIPRanges{
		Range:       dcl.StringOrNil(p.Range),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToRouterPeer converts a RouterPeer resource from its proto representation.
func ProtoToRouterPeer(p *computepb.ComputeRouterPeer) *compute.RouterPeer {
	obj := &compute.RouterPeer{
		CreationTimestamp:       dcl.StringOrNil(p.GetCreationTimestamp()),
		Router:                  dcl.StringOrNil(p.Router),
		Name:                    dcl.StringOrNil(p.Name),
		InterfaceName:           dcl.StringOrNil(p.InterfaceName),
		IPAddress:               dcl.StringOrNil(p.IpAddress),
		PeerIPAddress:           dcl.StringOrNil(p.PeerIpAddress),
		PeerAsn:                 dcl.Int64OrNil(p.PeerAsn),
		AdvertisedRoutePriority: dcl.Int64OrNil(p.AdvertisedRoutePriority),
		AdvertiseMode:           dcl.StringOrNil(p.AdvertiseMode),
		ManagementType:          dcl.StringOrNil(p.ManagementType),
		Region:                  dcl.StringOrNil(p.Region),
		Project:                 dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetAdvertisedGroups() {
		obj.AdvertisedGroups = append(obj.AdvertisedGroups, r)
	}
	for _, r := range p.GetAdvertisedIpRanges() {
		obj.AdvertisedIPRanges = append(obj.AdvertisedIPRanges, *ProtoToComputeRouterPeerAdvertisedIPRanges(r))
	}
	return obj
}

// RouterPeerAdvertisedIPRangesToProto converts a RouterPeerAdvertisedIPRanges resource to its proto representation.
func ComputeRouterPeerAdvertisedIPRangesToProto(o *compute.RouterPeerAdvertisedIPRanges) *computepb.ComputeRouterPeerAdvertisedIPRanges {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterPeerAdvertisedIPRanges{
		Range:       dcl.ValueOrEmptyString(o.Range),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// RouterPeerToProto converts a RouterPeer resource to its proto representation.
func RouterPeerToProto(resource *compute.RouterPeer) *computepb.ComputeRouterPeer {
	p := &computepb.ComputeRouterPeer{
		CreationTimestamp:       dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Router:                  dcl.ValueOrEmptyString(resource.Router),
		Name:                    dcl.ValueOrEmptyString(resource.Name),
		InterfaceName:           dcl.ValueOrEmptyString(resource.InterfaceName),
		IpAddress:               dcl.ValueOrEmptyString(resource.IPAddress),
		PeerIpAddress:           dcl.ValueOrEmptyString(resource.PeerIPAddress),
		PeerAsn:                 dcl.ValueOrEmptyInt64(resource.PeerAsn),
		AdvertisedRoutePriority: dcl.ValueOrEmptyInt64(resource.AdvertisedRoutePriority),
		AdvertiseMode:           dcl.ValueOrEmptyString(resource.AdvertiseMode),
		ManagementType:          dcl.ValueOrEmptyString(resource.ManagementType),
		Region:                  dcl.ValueOrEmptyString(resource.Region),
		Project:                 dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.AdvertisedGroups {
		p.AdvertisedGroups = append(p.AdvertisedGroups, r)
	}
	for _, r := range resource.AdvertisedIPRanges {
		p.AdvertisedIpRanges = append(p.AdvertisedIpRanges, ComputeRouterPeerAdvertisedIPRangesToProto(&r))
	}

	return p
}

// ApplyRouterPeer handles the gRPC request by passing it to the underlying RouterPeer Apply() method.
func (s *RouterPeerServer) applyRouterPeer(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeRouterPeerRequest) (*computepb.ComputeRouterPeer, error) {
	p := ProtoToRouterPeer(request.GetResource())
	res, err := c.ApplyRouterPeer(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RouterPeerToProto(res)
	return r, nil
}

// ApplyRouterPeer handles the gRPC request by passing it to the underlying RouterPeer Apply() method.
func (s *RouterPeerServer) ApplyComputeRouterPeer(ctx context.Context, request *computepb.ApplyComputeRouterPeerRequest) (*computepb.ComputeRouterPeer, error) {
	cl, err := createConfigRouterPeer(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRouterPeer(ctx, cl, request)
}

// DeleteRouterPeer handles the gRPC request by passing it to the underlying RouterPeer Delete() method.
func (s *RouterPeerServer) DeleteComputeRouterPeer(ctx context.Context, request *computepb.DeleteComputeRouterPeerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRouterPeer(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRouterPeer(ctx, ProtoToRouterPeer(request.GetResource()))

}

// ListComputeRouterPeer handles the gRPC request by passing it to the underlying RouterPeerList() method.
func (s *RouterPeerServer) ListComputeRouterPeer(ctx context.Context, request *computepb.ListComputeRouterPeerRequest) (*computepb.ListComputeRouterPeerResponse, error) {
	cl, err := createConfigRouterPeer(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRouterPeer(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeRouterPeer
	for _, r := range resources.Items {
		rp := RouterPeerToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeRouterPeerResponse{Items: protos}, nil
}

func createConfigRouterPeer(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
