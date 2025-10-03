// Copyright 2024 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// NetworkFirewallPolicyServer implements the gRPC interface for NetworkFirewallPolicy.
type NetworkFirewallPolicyServer struct{}

// ProtoToNetworkFirewallPolicy converts a NetworkFirewallPolicy resource from its proto representation.
func ProtoToNetworkFirewallPolicy(p *betapb.ComputeBetaNetworkFirewallPolicy) *beta.NetworkFirewallPolicy {
	obj := &beta.NetworkFirewallPolicy{
		Location:          dcl.StringOrNil(p.GetLocation()),
		CreationTimestamp: dcl.StringOrNil(p.GetCreationTimestamp()),
		Name:              dcl.StringOrNil(p.GetName()),
		Id:                dcl.StringOrNil(p.GetId()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		Fingerprint:       dcl.StringOrNil(p.GetFingerprint()),
		SelfLink:          dcl.StringOrNil(p.GetSelfLink()),
		SelfLinkWithId:    dcl.StringOrNil(p.GetSelfLinkWithId()),
		RuleTupleCount:    dcl.Int64OrNil(p.GetRuleTupleCount()),
		Region:            dcl.StringOrNil(p.GetRegion()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// NetworkFirewallPolicyToProto converts a NetworkFirewallPolicy resource to its proto representation.
func NetworkFirewallPolicyToProto(resource *beta.NetworkFirewallPolicy) *betapb.ComputeBetaNetworkFirewallPolicy {
	p := &betapb.ComputeBetaNetworkFirewallPolicy{}
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetId(dcl.ValueOrEmptyString(resource.Id))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetSelfLinkWithId(dcl.ValueOrEmptyString(resource.SelfLinkWithId))
	p.SetRuleTupleCount(dcl.ValueOrEmptyInt64(resource.RuleTupleCount))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyNetworkFirewallPolicy handles the gRPC request by passing it to the underlying NetworkFirewallPolicy Apply() method.
func (s *NetworkFirewallPolicyServer) applyNetworkFirewallPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaNetworkFirewallPolicyRequest) (*betapb.ComputeBetaNetworkFirewallPolicy, error) {
	p := ProtoToNetworkFirewallPolicy(request.GetResource())
	res, err := c.ApplyNetworkFirewallPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkFirewallPolicyToProto(res)
	return r, nil
}

// applyComputeBetaNetworkFirewallPolicy handles the gRPC request by passing it to the underlying NetworkFirewallPolicy Apply() method.
func (s *NetworkFirewallPolicyServer) ApplyComputeBetaNetworkFirewallPolicy(ctx context.Context, request *betapb.ApplyComputeBetaNetworkFirewallPolicyRequest) (*betapb.ComputeBetaNetworkFirewallPolicy, error) {
	cl, err := createConfigNetworkFirewallPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNetworkFirewallPolicy(ctx, cl, request)
}

// DeleteNetworkFirewallPolicy handles the gRPC request by passing it to the underlying NetworkFirewallPolicy Delete() method.
func (s *NetworkFirewallPolicyServer) DeleteComputeBetaNetworkFirewallPolicy(ctx context.Context, request *betapb.DeleteComputeBetaNetworkFirewallPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetworkFirewallPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetworkFirewallPolicy(ctx, ProtoToNetworkFirewallPolicy(request.GetResource()))

}

// ListComputeBetaNetworkFirewallPolicy handles the gRPC request by passing it to the underlying NetworkFirewallPolicyList() method.
func (s *NetworkFirewallPolicyServer) ListComputeBetaNetworkFirewallPolicy(ctx context.Context, request *betapb.ListComputeBetaNetworkFirewallPolicyRequest) (*betapb.ListComputeBetaNetworkFirewallPolicyResponse, error) {
	cl, err := createConfigNetworkFirewallPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetworkFirewallPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaNetworkFirewallPolicy
	for _, r := range resources.Items {
		rp := NetworkFirewallPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaNetworkFirewallPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNetworkFirewallPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
