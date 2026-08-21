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
	appenginepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/appengine/appengine_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/appengine"
)

// Server implements the gRPC interface for FirewallRule.
type FirewallRuleServer struct{}

// ProtoToFirewallRuleActionEnum converts a FirewallRuleActionEnum enum from its proto representation.
func ProtoToAppengineFirewallRuleActionEnum(e appenginepb.AppengineFirewallRuleActionEnum) *appengine.FirewallRuleActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineFirewallRuleActionEnum_name[int32(e)]; ok {
		e := appengine.FirewallRuleActionEnum(n[len("AppengineFirewallRuleActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToFirewallRule converts a FirewallRule resource from its proto representation.
func ProtoToFirewallRule(p *appenginepb.AppengineFirewallRule) *appengine.FirewallRule {
	obj := &appengine.FirewallRule{
		Action:      ProtoToAppengineFirewallRuleActionEnum(p.GetAction()),
		Description: dcl.StringOrNil(p.Description),
		Priority:    dcl.Int64OrNil(p.Priority),
		SourceRange: dcl.StringOrNil(p.SourceRange),
		App:         dcl.StringOrNil(p.App),
	}
	return obj
}

// FirewallRuleActionEnumToProto converts a FirewallRuleActionEnum enum to its proto representation.
func AppengineFirewallRuleActionEnumToProto(e *appengine.FirewallRuleActionEnum) appenginepb.AppengineFirewallRuleActionEnum {
	if e == nil {
		return appenginepb.AppengineFirewallRuleActionEnum(0)
	}
	if v, ok := appenginepb.AppengineFirewallRuleActionEnum_value["FirewallRuleActionEnum"+string(*e)]; ok {
		return appenginepb.AppengineFirewallRuleActionEnum(v)
	}
	return appenginepb.AppengineFirewallRuleActionEnum(0)
}

// FirewallRuleToProto converts a FirewallRule resource to its proto representation.
func FirewallRuleToProto(resource *appengine.FirewallRule) *appenginepb.AppengineFirewallRule {
	p := &appenginepb.AppengineFirewallRule{
		Action:      AppengineFirewallRuleActionEnumToProto(resource.Action),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Priority:    dcl.ValueOrEmptyInt64(resource.Priority),
		SourceRange: dcl.ValueOrEmptyString(resource.SourceRange),
		App:         dcl.ValueOrEmptyString(resource.App),
	}

	return p
}

// ApplyFirewallRule handles the gRPC request by passing it to the underlying FirewallRule Apply() method.
func (s *FirewallRuleServer) applyFirewallRule(ctx context.Context, c *appengine.Client, request *appenginepb.ApplyAppengineFirewallRuleRequest) (*appenginepb.AppengineFirewallRule, error) {
	p := ProtoToFirewallRule(request.GetResource())
	res, err := c.ApplyFirewallRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallRuleToProto(res)
	return r, nil
}

// ApplyFirewallRule handles the gRPC request by passing it to the underlying FirewallRule Apply() method.
func (s *FirewallRuleServer) ApplyAppengineFirewallRule(ctx context.Context, request *appenginepb.ApplyAppengineFirewallRuleRequest) (*appenginepb.AppengineFirewallRule, error) {
	cl, err := createConfigFirewallRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFirewallRule(ctx, cl, request)
}

// DeleteFirewallRule handles the gRPC request by passing it to the underlying FirewallRule Delete() method.
func (s *FirewallRuleServer) DeleteAppengineFirewallRule(ctx context.Context, request *appenginepb.DeleteAppengineFirewallRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewallRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewallRule(ctx, ProtoToFirewallRule(request.GetResource()))

}

// ListAppengineFirewallRule handles the gRPC request by passing it to the underlying FirewallRuleList() method.
func (s *FirewallRuleServer) ListAppengineFirewallRule(ctx context.Context, request *appenginepb.ListAppengineFirewallRuleRequest) (*appenginepb.ListAppengineFirewallRuleResponse, error) {
	cl, err := createConfigFirewallRule(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewallRule(ctx, request.App)
	if err != nil {
		return nil, err
	}
	var protos []*appenginepb.AppengineFirewallRule
	for _, r := range resources.Items {
		rp := FirewallRuleToProto(r)
		protos = append(protos, rp)
	}
	return &appenginepb.ListAppengineFirewallRuleResponse{Items: protos}, nil
}

func createConfigFirewallRule(ctx context.Context, service_account_file string) (*appengine.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return appengine.NewClient(conf), nil
}
