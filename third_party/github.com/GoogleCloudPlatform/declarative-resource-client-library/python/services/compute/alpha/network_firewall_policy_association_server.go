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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
)

// NetworkFirewallPolicyAssociationServer implements the gRPC interface for NetworkFirewallPolicyAssociation.
type NetworkFirewallPolicyAssociationServer struct{}

// ProtoToNetworkFirewallPolicyAssociation converts a NetworkFirewallPolicyAssociation resource from its proto representation.
func ProtoToNetworkFirewallPolicyAssociation(p *alphapb.ComputeAlphaNetworkFirewallPolicyAssociation) *alpha.NetworkFirewallPolicyAssociation {
	obj := &alpha.NetworkFirewallPolicyAssociation{
		Name:             dcl.StringOrNil(p.GetName()),
		AttachmentTarget: dcl.StringOrNil(p.GetAttachmentTarget()),
		FirewallPolicy:   dcl.StringOrNil(p.GetFirewallPolicy()),
		ShortName:        dcl.StringOrNil(p.GetShortName()),
		Location:         dcl.StringOrNil(p.GetLocation()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// NetworkFirewallPolicyAssociationToProto converts a NetworkFirewallPolicyAssociation resource to its proto representation.
func NetworkFirewallPolicyAssociationToProto(resource *alpha.NetworkFirewallPolicyAssociation) *alphapb.ComputeAlphaNetworkFirewallPolicyAssociation {
	p := &alphapb.ComputeAlphaNetworkFirewallPolicyAssociation{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAttachmentTarget(dcl.ValueOrEmptyString(resource.AttachmentTarget))
	p.SetFirewallPolicy(dcl.ValueOrEmptyString(resource.FirewallPolicy))
	p.SetShortName(dcl.ValueOrEmptyString(resource.ShortName))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyNetworkFirewallPolicyAssociation handles the gRPC request by passing it to the underlying NetworkFirewallPolicyAssociation Apply() method.
func (s *NetworkFirewallPolicyAssociationServer) applyNetworkFirewallPolicyAssociation(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaNetworkFirewallPolicyAssociationRequest) (*alphapb.ComputeAlphaNetworkFirewallPolicyAssociation, error) {
	p := ProtoToNetworkFirewallPolicyAssociation(request.GetResource())
	res, err := c.ApplyNetworkFirewallPolicyAssociation(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkFirewallPolicyAssociationToProto(res)
	return r, nil
}

// applyComputeAlphaNetworkFirewallPolicyAssociation handles the gRPC request by passing it to the underlying NetworkFirewallPolicyAssociation Apply() method.
func (s *NetworkFirewallPolicyAssociationServer) ApplyComputeAlphaNetworkFirewallPolicyAssociation(ctx context.Context, request *alphapb.ApplyComputeAlphaNetworkFirewallPolicyAssociationRequest) (*alphapb.ComputeAlphaNetworkFirewallPolicyAssociation, error) {
	cl, err := createConfigNetworkFirewallPolicyAssociation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNetworkFirewallPolicyAssociation(ctx, cl, request)
}

// DeleteNetworkFirewallPolicyAssociation handles the gRPC request by passing it to the underlying NetworkFirewallPolicyAssociation Delete() method.
func (s *NetworkFirewallPolicyAssociationServer) DeleteComputeAlphaNetworkFirewallPolicyAssociation(ctx context.Context, request *alphapb.DeleteComputeAlphaNetworkFirewallPolicyAssociationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetworkFirewallPolicyAssociation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetworkFirewallPolicyAssociation(ctx, ProtoToNetworkFirewallPolicyAssociation(request.GetResource()))

}

// ListComputeAlphaNetworkFirewallPolicyAssociation handles the gRPC request by passing it to the underlying NetworkFirewallPolicyAssociationList() method.
func (s *NetworkFirewallPolicyAssociationServer) ListComputeAlphaNetworkFirewallPolicyAssociation(ctx context.Context, request *alphapb.ListComputeAlphaNetworkFirewallPolicyAssociationRequest) (*alphapb.ListComputeAlphaNetworkFirewallPolicyAssociationResponse, error) {
	cl, err := createConfigNetworkFirewallPolicyAssociation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetworkFirewallPolicyAssociation(ctx, request.GetProject(), request.GetLocation(), request.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaNetworkFirewallPolicyAssociation
	for _, r := range resources.Items {
		rp := NetworkFirewallPolicyAssociationToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaNetworkFirewallPolicyAssociationResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNetworkFirewallPolicyAssociation(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
