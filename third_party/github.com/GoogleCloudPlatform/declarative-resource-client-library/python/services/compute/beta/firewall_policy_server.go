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

// FirewallPolicyServer implements the gRPC interface for FirewallPolicy.
type FirewallPolicyServer struct{}

// ProtoToFirewallPolicy converts a FirewallPolicy resource from its proto representation.
func ProtoToFirewallPolicy(p *betapb.ComputeBetaFirewallPolicy) *beta.FirewallPolicy {
	obj := &beta.FirewallPolicy{
		Name:              dcl.StringOrNil(p.GetName()),
		Id:                dcl.StringOrNil(p.GetId()),
		CreationTimestamp: dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		Fingerprint:       dcl.StringOrNil(p.GetFingerprint()),
		SelfLink:          dcl.StringOrNil(p.GetSelfLink()),
		SelfLinkWithId:    dcl.StringOrNil(p.GetSelfLinkWithId()),
		RuleTupleCount:    dcl.Int64OrNil(p.GetRuleTupleCount()),
		ShortName:         dcl.StringOrNil(p.GetShortName()),
		Parent:            dcl.StringOrNil(p.GetParent()),
	}
	return obj
}

// FirewallPolicyToProto converts a FirewallPolicy resource to its proto representation.
func FirewallPolicyToProto(resource *beta.FirewallPolicy) *betapb.ComputeBetaFirewallPolicy {
	p := &betapb.ComputeBetaFirewallPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetId(dcl.ValueOrEmptyString(resource.Id))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetSelfLinkWithId(dcl.ValueOrEmptyString(resource.SelfLinkWithId))
	p.SetRuleTupleCount(dcl.ValueOrEmptyInt64(resource.RuleTupleCount))
	p.SetShortName(dcl.ValueOrEmptyString(resource.ShortName))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))

	return p
}

// applyFirewallPolicy handles the gRPC request by passing it to the underlying FirewallPolicy Apply() method.
func (s *FirewallPolicyServer) applyFirewallPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaFirewallPolicyRequest) (*betapb.ComputeBetaFirewallPolicy, error) {
	p := ProtoToFirewallPolicy(request.GetResource())
	res, err := c.ApplyFirewallPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallPolicyToProto(res)
	return r, nil
}

// applyComputeBetaFirewallPolicy handles the gRPC request by passing it to the underlying FirewallPolicy Apply() method.
func (s *FirewallPolicyServer) ApplyComputeBetaFirewallPolicy(ctx context.Context, request *betapb.ApplyComputeBetaFirewallPolicyRequest) (*betapb.ComputeBetaFirewallPolicy, error) {
	cl, err := createConfigFirewallPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFirewallPolicy(ctx, cl, request)
}

// DeleteFirewallPolicy handles the gRPC request by passing it to the underlying FirewallPolicy Delete() method.
func (s *FirewallPolicyServer) DeleteComputeBetaFirewallPolicy(ctx context.Context, request *betapb.DeleteComputeBetaFirewallPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewallPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewallPolicy(ctx, ProtoToFirewallPolicy(request.GetResource()))

}

// ListComputeBetaFirewallPolicy handles the gRPC request by passing it to the underlying FirewallPolicyList() method.
func (s *FirewallPolicyServer) ListComputeBetaFirewallPolicy(ctx context.Context, request *betapb.ListComputeBetaFirewallPolicyRequest) (*betapb.ListComputeBetaFirewallPolicyResponse, error) {
	cl, err := createConfigFirewallPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewallPolicy(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaFirewallPolicy
	for _, r := range resources.Items {
		rp := FirewallPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaFirewallPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFirewallPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
