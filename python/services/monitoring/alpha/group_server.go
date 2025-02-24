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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/alpha/monitoring_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
)

// GroupServer implements the gRPC interface for Group.
type GroupServer struct{}

// ProtoToGroup converts a Group resource from its proto representation.
func ProtoToGroup(p *alphapb.MonitoringAlphaGroup) *alpha.Group {
	obj := &alpha.Group{
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Filter:      dcl.StringOrNil(p.GetFilter()),
		IsCluster:   dcl.Bool(p.GetIsCluster()),
		Name:        dcl.StringOrNil(p.GetName()),
		ParentName:  dcl.StringOrNil(p.GetParentName()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// GroupToProto converts a Group resource to its proto representation.
func GroupToProto(resource *alpha.Group) *alphapb.MonitoringAlphaGroup {
	p := &alphapb.MonitoringAlphaGroup{}
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetFilter(dcl.ValueOrEmptyString(resource.Filter))
	p.SetIsCluster(dcl.ValueOrEmptyBool(resource.IsCluster))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetParentName(dcl.ValueOrEmptyString(resource.ParentName))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) applyGroup(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaGroupRequest) (*alphapb.MonitoringAlphaGroup, error) {
	p := ProtoToGroup(request.GetResource())
	res, err := c.ApplyGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GroupToProto(res)
	return r, nil
}

// applyMonitoringAlphaGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) ApplyMonitoringAlphaGroup(ctx context.Context, request *alphapb.ApplyMonitoringAlphaGroupRequest) (*alphapb.MonitoringAlphaGroup, error) {
	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGroup(ctx, cl, request)
}

// DeleteGroup handles the gRPC request by passing it to the underlying Group Delete() method.
func (s *GroupServer) DeleteMonitoringAlphaGroup(ctx context.Context, request *alphapb.DeleteMonitoringAlphaGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGroup(ctx, ProtoToGroup(request.GetResource()))

}

// ListMonitoringAlphaGroup handles the gRPC request by passing it to the underlying GroupList() method.
func (s *GroupServer) ListMonitoringAlphaGroup(ctx context.Context, request *alphapb.ListMonitoringAlphaGroupRequest) (*alphapb.ListMonitoringAlphaGroupResponse, error) {
	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGroup(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaGroup
	for _, r := range resources.Items {
		rp := GroupToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListMonitoringAlphaGroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGroup(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
