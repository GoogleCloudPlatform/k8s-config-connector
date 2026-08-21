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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/alpha/gkehub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
)

// FleetServer implements the gRPC interface for Fleet.
type FleetServer struct{}

// ProtoToFleet converts a Fleet resource from its proto representation.
func ProtoToFleet(p *alphapb.GkehubAlphaFleet) *alpha.Fleet {
	obj := &alpha.Fleet{
		Name:              dcl.StringOrNil(p.GetName()),
		DisplayName:       dcl.StringOrNil(p.GetDisplayName()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Uid:               dcl.StringOrNil(p.GetUid()),
		ManagedNamespaces: dcl.Bool(p.GetManagedNamespaces()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// FleetToProto converts a Fleet resource to its proto representation.
func FleetToProto(resource *alpha.Fleet) *alphapb.GkehubAlphaFleet {
	p := &alphapb.GkehubAlphaFleet{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetManagedNamespaces(dcl.ValueOrEmptyBool(resource.ManagedNamespaces))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyFleet handles the gRPC request by passing it to the underlying Fleet Apply() method.
func (s *FleetServer) applyFleet(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkehubAlphaFleetRequest) (*alphapb.GkehubAlphaFleet, error) {
	p := ProtoToFleet(request.GetResource())
	res, err := c.ApplyFleet(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FleetToProto(res)
	return r, nil
}

// applyGkehubAlphaFleet handles the gRPC request by passing it to the underlying Fleet Apply() method.
func (s *FleetServer) ApplyGkehubAlphaFleet(ctx context.Context, request *alphapb.ApplyGkehubAlphaFleetRequest) (*alphapb.GkehubAlphaFleet, error) {
	cl, err := createConfigFleet(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFleet(ctx, cl, request)
}

// DeleteFleet handles the gRPC request by passing it to the underlying Fleet Delete() method.
func (s *FleetServer) DeleteGkehubAlphaFleet(ctx context.Context, request *alphapb.DeleteGkehubAlphaFleetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFleet(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFleet(ctx, ProtoToFleet(request.GetResource()))

}

// ListGkehubAlphaFleet is a no-op method because Fleet has no list method.
func (s *FleetServer) ListGkehubAlphaFleet(_ context.Context, _ *alphapb.ListGkehubAlphaFleetRequest) (*alphapb.ListGkehubAlphaFleetResponse, error) {
	return nil, nil
}

func createConfigFleet(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
