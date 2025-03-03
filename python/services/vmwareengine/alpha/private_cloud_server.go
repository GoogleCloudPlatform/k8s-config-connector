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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vmwareengine/alpha/vmwareengine_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmwareengine/alpha"
)

// Server implements the gRPC interface for PrivateCloud.
type PrivateCloudServer struct{}

// ProtoToPrivateCloudStateEnum converts a PrivateCloudStateEnum enum from its proto representation.
func ProtoToVmwareengineAlphaPrivateCloudStateEnum(e alphapb.VmwareengineAlphaPrivateCloudStateEnum) *alpha.PrivateCloudStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareengineAlphaPrivateCloudStateEnum_name[int32(e)]; ok {
		e := alpha.PrivateCloudStateEnum(n[len("VmwareengineAlphaPrivateCloudStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToPrivateCloudNetworkConfig converts a PrivateCloudNetworkConfig object from its proto representation.
func ProtoToVmwareengineAlphaPrivateCloudNetworkConfig(p *alphapb.VmwareengineAlphaPrivateCloudNetworkConfig) *alpha.PrivateCloudNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudNetworkConfig{
		Network:        dcl.StringOrNil(p.GetNetwork()),
		ServiceNetwork: dcl.StringOrNil(p.GetServiceNetwork()),
		ManagementCidr: dcl.StringOrNil(p.GetManagementCidr()),
	}
	return obj
}

// ProtoToPrivateCloudManagementCluster converts a PrivateCloudManagementCluster object from its proto representation.
func ProtoToVmwareengineAlphaPrivateCloudManagementCluster(p *alphapb.VmwareengineAlphaPrivateCloudManagementCluster) *alpha.PrivateCloudManagementCluster {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudManagementCluster{
		ClusterId:  dcl.StringOrNil(p.GetClusterId()),
		NodeTypeId: dcl.StringOrNil(p.GetNodeTypeId()),
		NodeCount:  dcl.Int64OrNil(p.GetNodeCount()),
	}
	return obj
}

// ProtoToPrivateCloudConditions converts a PrivateCloudConditions object from its proto representation.
func ProtoToVmwareengineAlphaPrivateCloudConditions(p *alphapb.VmwareengineAlphaPrivateCloudConditions) *alpha.PrivateCloudConditions {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudConditions{
		Code:    dcl.StringOrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	return obj
}

// ProtoToPrivateCloudHcx converts a PrivateCloudHcx object from its proto representation.
func ProtoToVmwareengineAlphaPrivateCloudHcx(p *alphapb.VmwareengineAlphaPrivateCloudHcx) *alpha.PrivateCloudHcx {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudHcx{
		Fdqn:       dcl.StringOrNil(p.GetFdqn()),
		InternalIP: dcl.StringOrNil(p.GetInternalIp()),
		ExternalIP: dcl.StringOrNil(p.GetExternalIp()),
		Version:    dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToPrivateCloudNsx converts a PrivateCloudNsx object from its proto representation.
func ProtoToVmwareengineAlphaPrivateCloudNsx(p *alphapb.VmwareengineAlphaPrivateCloudNsx) *alpha.PrivateCloudNsx {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudNsx{
		Fdqn:       dcl.StringOrNil(p.GetFdqn()),
		InternalIP: dcl.StringOrNil(p.GetInternalIp()),
		ExternalIP: dcl.StringOrNil(p.GetExternalIp()),
		Version:    dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToPrivateCloudVcenter converts a PrivateCloudVcenter object from its proto representation.
func ProtoToVmwareengineAlphaPrivateCloudVcenter(p *alphapb.VmwareengineAlphaPrivateCloudVcenter) *alpha.PrivateCloudVcenter {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudVcenter{
		Fdqn:       dcl.StringOrNil(p.GetFdqn()),
		InternalIP: dcl.StringOrNil(p.GetInternalIp()),
		ExternalIP: dcl.StringOrNil(p.GetExternalIp()),
		Version:    dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToPrivateCloud converts a PrivateCloud resource from its proto representation.
func ProtoToPrivateCloud(p *alphapb.VmwareengineAlphaPrivateCloud) *alpha.PrivateCloud {
	obj := &alpha.PrivateCloud{
		Name:              dcl.StringOrNil(p.GetName()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:        dcl.StringOrNil(p.GetDeleteTime()),
		ExpireTime:        dcl.StringOrNil(p.GetExpireTime()),
		State:             ProtoToVmwareengineAlphaPrivateCloudStateEnum(p.GetState()),
		NetworkConfig:     ProtoToVmwareengineAlphaPrivateCloudNetworkConfig(p.GetNetworkConfig()),
		ManagementCluster: ProtoToVmwareengineAlphaPrivateCloudManagementCluster(p.GetManagementCluster()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		Hcx:               ProtoToVmwareengineAlphaPrivateCloudHcx(p.GetHcx()),
		Nsx:               ProtoToVmwareengineAlphaPrivateCloudNsx(p.GetNsx()),
		Vcenter:           ProtoToVmwareengineAlphaPrivateCloudVcenter(p.GetVcenter()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToVmwareengineAlphaPrivateCloudConditions(r))
	}
	return obj
}

// PrivateCloudStateEnumToProto converts a PrivateCloudStateEnum enum to its proto representation.
func VmwareengineAlphaPrivateCloudStateEnumToProto(e *alpha.PrivateCloudStateEnum) alphapb.VmwareengineAlphaPrivateCloudStateEnum {
	if e == nil {
		return alphapb.VmwareengineAlphaPrivateCloudStateEnum(0)
	}
	if v, ok := alphapb.VmwareengineAlphaPrivateCloudStateEnum_value["PrivateCloudStateEnum"+string(*e)]; ok {
		return alphapb.VmwareengineAlphaPrivateCloudStateEnum(v)
	}
	return alphapb.VmwareengineAlphaPrivateCloudStateEnum(0)
}

// PrivateCloudNetworkConfigToProto converts a PrivateCloudNetworkConfig object to its proto representation.
func VmwareengineAlphaPrivateCloudNetworkConfigToProto(o *alpha.PrivateCloudNetworkConfig) *alphapb.VmwareengineAlphaPrivateCloudNetworkConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareengineAlphaPrivateCloudNetworkConfig{}
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetServiceNetwork(dcl.ValueOrEmptyString(o.ServiceNetwork))
	p.SetManagementCidr(dcl.ValueOrEmptyString(o.ManagementCidr))
	return p
}

// PrivateCloudManagementClusterToProto converts a PrivateCloudManagementCluster object to its proto representation.
func VmwareengineAlphaPrivateCloudManagementClusterToProto(o *alpha.PrivateCloudManagementCluster) *alphapb.VmwareengineAlphaPrivateCloudManagementCluster {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareengineAlphaPrivateCloudManagementCluster{}
	p.SetClusterId(dcl.ValueOrEmptyString(o.ClusterId))
	p.SetNodeTypeId(dcl.ValueOrEmptyString(o.NodeTypeId))
	p.SetNodeCount(dcl.ValueOrEmptyInt64(o.NodeCount))
	return p
}

// PrivateCloudConditionsToProto converts a PrivateCloudConditions object to its proto representation.
func VmwareengineAlphaPrivateCloudConditionsToProto(o *alpha.PrivateCloudConditions) *alphapb.VmwareengineAlphaPrivateCloudConditions {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareengineAlphaPrivateCloudConditions{}
	p.SetCode(dcl.ValueOrEmptyString(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	return p
}

// PrivateCloudHcxToProto converts a PrivateCloudHcx object to its proto representation.
func VmwareengineAlphaPrivateCloudHcxToProto(o *alpha.PrivateCloudHcx) *alphapb.VmwareengineAlphaPrivateCloudHcx {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareengineAlphaPrivateCloudHcx{}
	p.SetFdqn(dcl.ValueOrEmptyString(o.Fdqn))
	p.SetInternalIp(dcl.ValueOrEmptyString(o.InternalIP))
	p.SetExternalIp(dcl.ValueOrEmptyString(o.ExternalIP))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// PrivateCloudNsxToProto converts a PrivateCloudNsx object to its proto representation.
func VmwareengineAlphaPrivateCloudNsxToProto(o *alpha.PrivateCloudNsx) *alphapb.VmwareengineAlphaPrivateCloudNsx {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareengineAlphaPrivateCloudNsx{}
	p.SetFdqn(dcl.ValueOrEmptyString(o.Fdqn))
	p.SetInternalIp(dcl.ValueOrEmptyString(o.InternalIP))
	p.SetExternalIp(dcl.ValueOrEmptyString(o.ExternalIP))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// PrivateCloudVcenterToProto converts a PrivateCloudVcenter object to its proto representation.
func VmwareengineAlphaPrivateCloudVcenterToProto(o *alpha.PrivateCloudVcenter) *alphapb.VmwareengineAlphaPrivateCloudVcenter {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareengineAlphaPrivateCloudVcenter{}
	p.SetFdqn(dcl.ValueOrEmptyString(o.Fdqn))
	p.SetInternalIp(dcl.ValueOrEmptyString(o.InternalIP))
	p.SetExternalIp(dcl.ValueOrEmptyString(o.ExternalIP))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// PrivateCloudToProto converts a PrivateCloud resource to its proto representation.
func PrivateCloudToProto(resource *alpha.PrivateCloud) *alphapb.VmwareengineAlphaPrivateCloud {
	p := &alphapb.VmwareengineAlphaPrivateCloud{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetExpireTime(dcl.ValueOrEmptyString(resource.ExpireTime))
	p.SetState(VmwareengineAlphaPrivateCloudStateEnumToProto(resource.State))
	p.SetNetworkConfig(VmwareengineAlphaPrivateCloudNetworkConfigToProto(resource.NetworkConfig))
	p.SetManagementCluster(VmwareengineAlphaPrivateCloudManagementClusterToProto(resource.ManagementCluster))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetHcx(VmwareengineAlphaPrivateCloudHcxToProto(resource.Hcx))
	p.SetNsx(VmwareengineAlphaPrivateCloudNsxToProto(resource.Nsx))
	p.SetVcenter(VmwareengineAlphaPrivateCloudVcenterToProto(resource.Vcenter))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sConditions := make([]*alphapb.VmwareengineAlphaPrivateCloudConditions, len(resource.Conditions))
	for i, r := range resource.Conditions {
		sConditions[i] = VmwareengineAlphaPrivateCloudConditionsToProto(&r)
	}
	p.SetConditions(sConditions)

	return p
}

// applyPrivateCloud handles the gRPC request by passing it to the underlying PrivateCloud Apply() method.
func (s *PrivateCloudServer) applyPrivateCloud(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVmwareengineAlphaPrivateCloudRequest) (*alphapb.VmwareengineAlphaPrivateCloud, error) {
	p := ProtoToPrivateCloud(request.GetResource())
	res, err := c.ApplyPrivateCloud(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PrivateCloudToProto(res)
	return r, nil
}

// applyVmwareengineAlphaPrivateCloud handles the gRPC request by passing it to the underlying PrivateCloud Apply() method.
func (s *PrivateCloudServer) ApplyVmwareengineAlphaPrivateCloud(ctx context.Context, request *alphapb.ApplyVmwareengineAlphaPrivateCloudRequest) (*alphapb.VmwareengineAlphaPrivateCloud, error) {
	cl, err := createConfigPrivateCloud(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPrivateCloud(ctx, cl, request)
}

// DeletePrivateCloud handles the gRPC request by passing it to the underlying PrivateCloud Delete() method.
func (s *PrivateCloudServer) DeleteVmwareengineAlphaPrivateCloud(ctx context.Context, request *alphapb.DeleteVmwareengineAlphaPrivateCloudRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPrivateCloud(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePrivateCloud(ctx, ProtoToPrivateCloud(request.GetResource()))

}

// ListVmwareengineAlphaPrivateCloud handles the gRPC request by passing it to the underlying PrivateCloudList() method.
func (s *PrivateCloudServer) ListVmwareengineAlphaPrivateCloud(ctx context.Context, request *alphapb.ListVmwareengineAlphaPrivateCloudRequest) (*alphapb.ListVmwareengineAlphaPrivateCloudResponse, error) {
	cl, err := createConfigPrivateCloud(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPrivateCloud(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VmwareengineAlphaPrivateCloud
	for _, r := range resources.Items {
		rp := PrivateCloudToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVmwareengineAlphaPrivateCloudResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPrivateCloud(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
