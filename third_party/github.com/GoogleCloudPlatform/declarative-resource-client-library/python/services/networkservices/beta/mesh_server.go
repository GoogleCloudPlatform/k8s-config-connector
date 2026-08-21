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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/beta/networkservices_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
)

// MeshServer implements the gRPC interface for Mesh.
type MeshServer struct{}

// ProtoToMesh converts a Mesh resource from its proto representation.
func ProtoToMesh(p *betapb.NetworkservicesBetaMesh) *beta.Mesh {
	obj := &beta.Mesh{
		Name:             dcl.StringOrNil(p.GetName()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		InterceptionPort: dcl.Int64OrNil(p.GetInterceptionPort()),
		Project:          dcl.StringOrNil(p.GetProject()),
		Location:         dcl.StringOrNil(p.GetLocation()),
		SelfLink:         dcl.StringOrNil(p.GetSelfLink()),
	}
	return obj
}

// MeshToProto converts a Mesh resource to its proto representation.
func MeshToProto(resource *beta.Mesh) *betapb.NetworkservicesBetaMesh {
	p := &betapb.NetworkservicesBetaMesh{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInterceptionPort(dcl.ValueOrEmptyInt64(resource.InterceptionPort))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyMesh handles the gRPC request by passing it to the underlying Mesh Apply() method.
func (s *MeshServer) applyMesh(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkservicesBetaMeshRequest) (*betapb.NetworkservicesBetaMesh, error) {
	p := ProtoToMesh(request.GetResource())
	res, err := c.ApplyMesh(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MeshToProto(res)
	return r, nil
}

// applyNetworkservicesBetaMesh handles the gRPC request by passing it to the underlying Mesh Apply() method.
func (s *MeshServer) ApplyNetworkservicesBetaMesh(ctx context.Context, request *betapb.ApplyNetworkservicesBetaMeshRequest) (*betapb.NetworkservicesBetaMesh, error) {
	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMesh(ctx, cl, request)
}

// DeleteMesh handles the gRPC request by passing it to the underlying Mesh Delete() method.
func (s *MeshServer) DeleteNetworkservicesBetaMesh(ctx context.Context, request *betapb.DeleteNetworkservicesBetaMeshRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMesh(ctx, ProtoToMesh(request.GetResource()))

}

// ListNetworkservicesBetaMesh handles the gRPC request by passing it to the underlying MeshList() method.
func (s *MeshServer) ListNetworkservicesBetaMesh(ctx context.Context, request *betapb.ListNetworkservicesBetaMeshRequest) (*betapb.ListNetworkservicesBetaMeshResponse, error) {
	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMesh(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkservicesBetaMesh
	for _, r := range resources.Items {
		rp := MeshToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworkservicesBetaMeshResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMesh(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
