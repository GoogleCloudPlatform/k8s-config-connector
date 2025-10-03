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
	networkservicespb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/networkservices_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices"
)

// MeshServer implements the gRPC interface for Mesh.
type MeshServer struct{}

// ProtoToMesh converts a Mesh resource from its proto representation.
func ProtoToMesh(p *networkservicespb.NetworkservicesMesh) *networkservices.Mesh {
	obj := &networkservices.Mesh{
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
func MeshToProto(resource *networkservices.Mesh) *networkservicespb.NetworkservicesMesh {
	p := &networkservicespb.NetworkservicesMesh{}
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
func (s *MeshServer) applyMesh(ctx context.Context, c *networkservices.Client, request *networkservicespb.ApplyNetworkservicesMeshRequest) (*networkservicespb.NetworkservicesMesh, error) {
	p := ProtoToMesh(request.GetResource())
	res, err := c.ApplyMesh(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MeshToProto(res)
	return r, nil
}

// applyNetworkservicesMesh handles the gRPC request by passing it to the underlying Mesh Apply() method.
func (s *MeshServer) ApplyNetworkservicesMesh(ctx context.Context, request *networkservicespb.ApplyNetworkservicesMeshRequest) (*networkservicespb.NetworkservicesMesh, error) {
	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMesh(ctx, cl, request)
}

// DeleteMesh handles the gRPC request by passing it to the underlying Mesh Delete() method.
func (s *MeshServer) DeleteNetworkservicesMesh(ctx context.Context, request *networkservicespb.DeleteNetworkservicesMeshRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMesh(ctx, ProtoToMesh(request.GetResource()))

}

// ListNetworkservicesMesh handles the gRPC request by passing it to the underlying MeshList() method.
func (s *MeshServer) ListNetworkservicesMesh(ctx context.Context, request *networkservicespb.ListNetworkservicesMeshRequest) (*networkservicespb.ListNetworkservicesMeshResponse, error) {
	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMesh(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*networkservicespb.NetworkservicesMesh
	for _, r := range resources.Items {
		rp := MeshToProto(r)
		protos = append(protos, rp)
	}
	p := &networkservicespb.ListNetworkservicesMeshResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMesh(ctx context.Context, service_account_file string) (*networkservices.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return networkservices.NewClient(conf), nil
}
