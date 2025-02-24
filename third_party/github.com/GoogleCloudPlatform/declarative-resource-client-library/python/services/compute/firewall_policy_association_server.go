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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// FirewallPolicyAssociationServer implements the gRPC interface for FirewallPolicyAssociation.
type FirewallPolicyAssociationServer struct{}

// ProtoToFirewallPolicyAssociation converts a FirewallPolicyAssociation resource from its proto representation.
func ProtoToFirewallPolicyAssociation(p *computepb.ComputeFirewallPolicyAssociation) *compute.FirewallPolicyAssociation {
	obj := &compute.FirewallPolicyAssociation{
		Name:             dcl.StringOrNil(p.GetName()),
		AttachmentTarget: dcl.StringOrNil(p.GetAttachmentTarget()),
		FirewallPolicy:   dcl.StringOrNil(p.GetFirewallPolicy()),
		ShortName:        dcl.StringOrNil(p.GetShortName()),
	}
	return obj
}

// FirewallPolicyAssociationToProto converts a FirewallPolicyAssociation resource to its proto representation.
func FirewallPolicyAssociationToProto(resource *compute.FirewallPolicyAssociation) *computepb.ComputeFirewallPolicyAssociation {
	p := &computepb.ComputeFirewallPolicyAssociation{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAttachmentTarget(dcl.ValueOrEmptyString(resource.AttachmentTarget))
	p.SetFirewallPolicy(dcl.ValueOrEmptyString(resource.FirewallPolicy))
	p.SetShortName(dcl.ValueOrEmptyString(resource.ShortName))

	return p
}

// applyFirewallPolicyAssociation handles the gRPC request by passing it to the underlying FirewallPolicyAssociation Apply() method.
func (s *FirewallPolicyAssociationServer) applyFirewallPolicyAssociation(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeFirewallPolicyAssociationRequest) (*computepb.ComputeFirewallPolicyAssociation, error) {
	p := ProtoToFirewallPolicyAssociation(request.GetResource())
	res, err := c.ApplyFirewallPolicyAssociation(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallPolicyAssociationToProto(res)
	return r, nil
}

// applyComputeFirewallPolicyAssociation handles the gRPC request by passing it to the underlying FirewallPolicyAssociation Apply() method.
func (s *FirewallPolicyAssociationServer) ApplyComputeFirewallPolicyAssociation(ctx context.Context, request *computepb.ApplyComputeFirewallPolicyAssociationRequest) (*computepb.ComputeFirewallPolicyAssociation, error) {
	cl, err := createConfigFirewallPolicyAssociation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFirewallPolicyAssociation(ctx, cl, request)
}

// DeleteFirewallPolicyAssociation handles the gRPC request by passing it to the underlying FirewallPolicyAssociation Delete() method.
func (s *FirewallPolicyAssociationServer) DeleteComputeFirewallPolicyAssociation(ctx context.Context, request *computepb.DeleteComputeFirewallPolicyAssociationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewallPolicyAssociation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewallPolicyAssociation(ctx, ProtoToFirewallPolicyAssociation(request.GetResource()))

}

// ListComputeFirewallPolicyAssociation handles the gRPC request by passing it to the underlying FirewallPolicyAssociationList() method.
func (s *FirewallPolicyAssociationServer) ListComputeFirewallPolicyAssociation(ctx context.Context, request *computepb.ListComputeFirewallPolicyAssociationRequest) (*computepb.ListComputeFirewallPolicyAssociationResponse, error) {
	cl, err := createConfigFirewallPolicyAssociation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewallPolicyAssociation(ctx, request.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeFirewallPolicyAssociation
	for _, r := range resources.Items {
		rp := FirewallPolicyAssociationToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeFirewallPolicyAssociationResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFirewallPolicyAssociation(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
