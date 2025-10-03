// Copyright 2023 Google LLC. All Rights Reserved.
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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vmware/alpha/vmware_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmware/alpha"
)

// VmwareEngineNetworkServer implements the gRPC interface for VmwareEngineNetwork.
type VmwareEngineNetworkServer struct{}

// ProtoToVmwareEngineNetworkVPCNetworksTypeEnum converts a VmwareEngineNetworkVPCNetworksTypeEnum enum from its proto representation.
func ProtoToVmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum(e alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum) *alpha.VmwareEngineNetworkVPCNetworksTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum_name[int32(e)]; ok {
		e := alpha.VmwareEngineNetworkVPCNetworksTypeEnum(n[len("VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToVmwareEngineNetworkStateEnum converts a VmwareEngineNetworkStateEnum enum from its proto representation.
func ProtoToVmwareAlphaVmwareEngineNetworkStateEnum(e alphapb.VmwareAlphaVmwareEngineNetworkStateEnum) *alpha.VmwareEngineNetworkStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaVmwareEngineNetworkStateEnum_name[int32(e)]; ok {
		e := alpha.VmwareEngineNetworkStateEnum(n[len("VmwareAlphaVmwareEngineNetworkStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToVmwareEngineNetworkTypeEnum converts a VmwareEngineNetworkTypeEnum enum from its proto representation.
func ProtoToVmwareAlphaVmwareEngineNetworkTypeEnum(e alphapb.VmwareAlphaVmwareEngineNetworkTypeEnum) *alpha.VmwareEngineNetworkTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaVmwareEngineNetworkTypeEnum_name[int32(e)]; ok {
		e := alpha.VmwareEngineNetworkTypeEnum(n[len("VmwareAlphaVmwareEngineNetworkTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToVmwareEngineNetworkVPCNetworks converts a VmwareEngineNetworkVPCNetworks object from its proto representation.
func ProtoToVmwareAlphaVmwareEngineNetworkVPCNetworks(p *alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworks) *alpha.VmwareEngineNetworkVPCNetworks {
	if p == nil {
		return nil
	}
	obj := &alpha.VmwareEngineNetworkVPCNetworks{
		Type:    ProtoToVmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum(p.GetType()),
		Network: dcl.StringOrNil(p.GetNetwork()),
	}
	return obj
}

// ProtoToVmwareEngineNetwork converts a VmwareEngineNetwork resource from its proto representation.
func ProtoToVmwareEngineNetwork(p *alphapb.VmwareAlphaVmwareEngineNetwork) *alpha.VmwareEngineNetwork {
	obj := &alpha.VmwareEngineNetwork{
		Name:        dcl.StringOrNil(p.GetName()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.GetDescription()),
		State:       ProtoToVmwareAlphaVmwareEngineNetworkStateEnum(p.GetState()),
		Type:        ProtoToVmwareAlphaVmwareEngineNetworkTypeEnum(p.GetType()),
		Uid:         dcl.StringOrNil(p.GetUid()),
		Etag:        dcl.StringOrNil(p.GetEtag()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetVpcNetworks() {
		obj.VPCNetworks = append(obj.VPCNetworks, *ProtoToVmwareAlphaVmwareEngineNetworkVPCNetworks(r))
	}
	return obj
}

// VmwareEngineNetworkVPCNetworksTypeEnumToProto converts a VmwareEngineNetworkVPCNetworksTypeEnum enum to its proto representation.
func VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnumToProto(e *alpha.VmwareEngineNetworkVPCNetworksTypeEnum) alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum {
	if e == nil {
		return alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum_value["VmwareEngineNetworkVPCNetworksTypeEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum(v)
	}
	return alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnum(0)
}

// VmwareEngineNetworkStateEnumToProto converts a VmwareEngineNetworkStateEnum enum to its proto representation.
func VmwareAlphaVmwareEngineNetworkStateEnumToProto(e *alpha.VmwareEngineNetworkStateEnum) alphapb.VmwareAlphaVmwareEngineNetworkStateEnum {
	if e == nil {
		return alphapb.VmwareAlphaVmwareEngineNetworkStateEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaVmwareEngineNetworkStateEnum_value["VmwareEngineNetworkStateEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaVmwareEngineNetworkStateEnum(v)
	}
	return alphapb.VmwareAlphaVmwareEngineNetworkStateEnum(0)
}

// VmwareEngineNetworkTypeEnumToProto converts a VmwareEngineNetworkTypeEnum enum to its proto representation.
func VmwareAlphaVmwareEngineNetworkTypeEnumToProto(e *alpha.VmwareEngineNetworkTypeEnum) alphapb.VmwareAlphaVmwareEngineNetworkTypeEnum {
	if e == nil {
		return alphapb.VmwareAlphaVmwareEngineNetworkTypeEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaVmwareEngineNetworkTypeEnum_value["VmwareEngineNetworkTypeEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaVmwareEngineNetworkTypeEnum(v)
	}
	return alphapb.VmwareAlphaVmwareEngineNetworkTypeEnum(0)
}

// VmwareEngineNetworkVPCNetworksToProto converts a VmwareEngineNetworkVPCNetworks object to its proto representation.
func VmwareAlphaVmwareEngineNetworkVPCNetworksToProto(o *alpha.VmwareEngineNetworkVPCNetworks) *alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworks {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworks{}
	p.SetType(VmwareAlphaVmwareEngineNetworkVPCNetworksTypeEnumToProto(o.Type))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	return p
}

// VmwareEngineNetworkToProto converts a VmwareEngineNetwork resource to its proto representation.
func VmwareEngineNetworkToProto(resource *alpha.VmwareEngineNetwork) *alphapb.VmwareAlphaVmwareEngineNetwork {
	p := &alphapb.VmwareAlphaVmwareEngineNetwork{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(VmwareAlphaVmwareEngineNetworkStateEnumToProto(resource.State))
	p.SetType(VmwareAlphaVmwareEngineNetworkTypeEnumToProto(resource.Type))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sVPCNetworks := make([]*alphapb.VmwareAlphaVmwareEngineNetworkVPCNetworks, len(resource.VPCNetworks))
	for i, r := range resource.VPCNetworks {
		sVPCNetworks[i] = VmwareAlphaVmwareEngineNetworkVPCNetworksToProto(&r)
	}
	p.SetVpcNetworks(sVPCNetworks)

	return p
}

// applyVmwareEngineNetwork handles the gRPC request by passing it to the underlying VmwareEngineNetwork Apply() method.
func (s *VmwareEngineNetworkServer) applyVmwareEngineNetwork(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVmwareAlphaVmwareEngineNetworkRequest) (*alphapb.VmwareAlphaVmwareEngineNetwork, error) {
	p := ProtoToVmwareEngineNetwork(request.GetResource())
	res, err := c.ApplyVmwareEngineNetwork(ctx, p)
	if err != nil {
		return nil, err
	}
	r := VmwareEngineNetworkToProto(res)
	return r, nil
}

// applyVmwareAlphaVmwareEngineNetwork handles the gRPC request by passing it to the underlying VmwareEngineNetwork Apply() method.
func (s *VmwareEngineNetworkServer) ApplyVmwareAlphaVmwareEngineNetwork(ctx context.Context, request *alphapb.ApplyVmwareAlphaVmwareEngineNetworkRequest) (*alphapb.VmwareAlphaVmwareEngineNetwork, error) {
	cl, err := createConfigVmwareEngineNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyVmwareEngineNetwork(ctx, cl, request)
}

// DeleteVmwareEngineNetwork handles the gRPC request by passing it to the underlying VmwareEngineNetwork Delete() method.
func (s *VmwareEngineNetworkServer) DeleteVmwareAlphaVmwareEngineNetwork(ctx context.Context, request *alphapb.DeleteVmwareAlphaVmwareEngineNetworkRequest) (*emptypb.Empty, error) {

	cl, err := createConfigVmwareEngineNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteVmwareEngineNetwork(ctx, ProtoToVmwareEngineNetwork(request.GetResource()))

}

// ListVmwareAlphaVmwareEngineNetwork handles the gRPC request by passing it to the underlying VmwareEngineNetworkList() method.
func (s *VmwareEngineNetworkServer) ListVmwareAlphaVmwareEngineNetwork(ctx context.Context, request *alphapb.ListVmwareAlphaVmwareEngineNetworkRequest) (*alphapb.ListVmwareAlphaVmwareEngineNetworkResponse, error) {
	cl, err := createConfigVmwareEngineNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListVmwareEngineNetwork(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VmwareAlphaVmwareEngineNetwork
	for _, r := range resources.Items {
		rp := VmwareEngineNetworkToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVmwareAlphaVmwareEngineNetworkResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigVmwareEngineNetwork(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
