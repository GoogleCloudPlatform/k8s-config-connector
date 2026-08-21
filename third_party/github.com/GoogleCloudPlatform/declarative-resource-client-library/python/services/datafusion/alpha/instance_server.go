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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/datafusion/alpha/datafusion_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/alpha"
)

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceTypeEnum converts a InstanceTypeEnum enum from its proto representation.
func ProtoToDatafusionAlphaInstanceTypeEnum(e alphapb.DatafusionAlphaInstanceTypeEnum) *alpha.InstanceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DatafusionAlphaInstanceTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceTypeEnum(n[len("DatafusionAlphaInstanceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToDatafusionAlphaInstanceStateEnum(e alphapb.DatafusionAlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DatafusionAlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("DatafusionAlphaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkConfig converts a InstanceNetworkConfig object from its proto representation.
func ProtoToDatafusionAlphaInstanceNetworkConfig(p *alphapb.DatafusionAlphaInstanceNetworkConfig) *alpha.InstanceNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNetworkConfig{
		Network:      dcl.StringOrNil(p.GetNetwork()),
		IPAllocation: dcl.StringOrNil(p.GetIpAllocation()),
	}
	return obj
}

// ProtoToInstanceAvailableVersion converts a InstanceAvailableVersion object from its proto representation.
func ProtoToDatafusionAlphaInstanceAvailableVersion(p *alphapb.DatafusionAlphaInstanceAvailableVersion) *alpha.InstanceAvailableVersion {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceAvailableVersion{
		VersionNumber:  dcl.StringOrNil(p.GetVersionNumber()),
		DefaultVersion: dcl.Bool(p.GetDefaultVersion()),
	}
	for _, r := range p.GetAvailableFeatures() {
		obj.AvailableFeatures = append(obj.AvailableFeatures, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.DatafusionAlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:                        dcl.StringOrNil(p.GetName()),
		Description:                 dcl.StringOrNil(p.GetDescription()),
		Type:                        ProtoToDatafusionAlphaInstanceTypeEnum(p.GetType()),
		EnableStackdriverLogging:    dcl.Bool(p.GetEnableStackdriverLogging()),
		EnableStackdriverMonitoring: dcl.Bool(p.GetEnableStackdriverMonitoring()),
		PrivateInstance:             dcl.Bool(p.GetPrivateInstance()),
		NetworkConfig:               ProtoToDatafusionAlphaInstanceNetworkConfig(p.GetNetworkConfig()),
		CreateTime:                  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                  dcl.StringOrNil(p.GetUpdateTime()),
		State:                       ProtoToDatafusionAlphaInstanceStateEnum(p.GetState()),
		StateMessage:                dcl.StringOrNil(p.GetStateMessage()),
		ServiceEndpoint:             dcl.StringOrNil(p.GetServiceEndpoint()),
		Zone:                        dcl.StringOrNil(p.GetZone()),
		Version:                     dcl.StringOrNil(p.GetVersion()),
		DisplayName:                 dcl.StringOrNil(p.GetDisplayName()),
		ApiEndpoint:                 dcl.StringOrNil(p.GetApiEndpoint()),
		GcsBucket:                   dcl.StringOrNil(p.GetGcsBucket()),
		P4ServiceAccount:            dcl.StringOrNil(p.GetP4ServiceAccount()),
		TenantProjectId:             dcl.StringOrNil(p.GetTenantProjectId()),
		DataprocServiceAccount:      dcl.StringOrNil(p.GetDataprocServiceAccount()),
		Project:                     dcl.StringOrNil(p.GetProject()),
		Location:                    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetAvailableVersion() {
		obj.AvailableVersion = append(obj.AvailableVersion, *ProtoToDatafusionAlphaInstanceAvailableVersion(r))
	}
	return obj
}

// InstanceTypeEnumToProto converts a InstanceTypeEnum enum to its proto representation.
func DatafusionAlphaInstanceTypeEnumToProto(e *alpha.InstanceTypeEnum) alphapb.DatafusionAlphaInstanceTypeEnum {
	if e == nil {
		return alphapb.DatafusionAlphaInstanceTypeEnum(0)
	}
	if v, ok := alphapb.DatafusionAlphaInstanceTypeEnum_value["InstanceTypeEnum"+string(*e)]; ok {
		return alphapb.DatafusionAlphaInstanceTypeEnum(v)
	}
	return alphapb.DatafusionAlphaInstanceTypeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func DatafusionAlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.DatafusionAlphaInstanceStateEnum {
	if e == nil {
		return alphapb.DatafusionAlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.DatafusionAlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.DatafusionAlphaInstanceStateEnum(v)
	}
	return alphapb.DatafusionAlphaInstanceStateEnum(0)
}

// InstanceNetworkConfigToProto converts a InstanceNetworkConfig object to its proto representation.
func DatafusionAlphaInstanceNetworkConfigToProto(o *alpha.InstanceNetworkConfig) *alphapb.DatafusionAlphaInstanceNetworkConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DatafusionAlphaInstanceNetworkConfig{}
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetIpAllocation(dcl.ValueOrEmptyString(o.IPAllocation))
	return p
}

// InstanceAvailableVersionToProto converts a InstanceAvailableVersion object to its proto representation.
func DatafusionAlphaInstanceAvailableVersionToProto(o *alpha.InstanceAvailableVersion) *alphapb.DatafusionAlphaInstanceAvailableVersion {
	if o == nil {
		return nil
	}
	p := &alphapb.DatafusionAlphaInstanceAvailableVersion{}
	p.SetVersionNumber(dcl.ValueOrEmptyString(o.VersionNumber))
	p.SetDefaultVersion(dcl.ValueOrEmptyBool(o.DefaultVersion))
	sAvailableFeatures := make([]string, len(o.AvailableFeatures))
	for i, r := range o.AvailableFeatures {
		sAvailableFeatures[i] = r
	}
	p.SetAvailableFeatures(sAvailableFeatures)
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *alpha.Instance) *alphapb.DatafusionAlphaInstance {
	p := &alphapb.DatafusionAlphaInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetType(DatafusionAlphaInstanceTypeEnumToProto(resource.Type))
	p.SetEnableStackdriverLogging(dcl.ValueOrEmptyBool(resource.EnableStackdriverLogging))
	p.SetEnableStackdriverMonitoring(dcl.ValueOrEmptyBool(resource.EnableStackdriverMonitoring))
	p.SetPrivateInstance(dcl.ValueOrEmptyBool(resource.PrivateInstance))
	p.SetNetworkConfig(DatafusionAlphaInstanceNetworkConfigToProto(resource.NetworkConfig))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetState(DatafusionAlphaInstanceStateEnumToProto(resource.State))
	p.SetStateMessage(dcl.ValueOrEmptyString(resource.StateMessage))
	p.SetServiceEndpoint(dcl.ValueOrEmptyString(resource.ServiceEndpoint))
	p.SetZone(dcl.ValueOrEmptyString(resource.Zone))
	p.SetVersion(dcl.ValueOrEmptyString(resource.Version))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetApiEndpoint(dcl.ValueOrEmptyString(resource.ApiEndpoint))
	p.SetGcsBucket(dcl.ValueOrEmptyString(resource.GcsBucket))
	p.SetP4ServiceAccount(dcl.ValueOrEmptyString(resource.P4ServiceAccount))
	p.SetTenantProjectId(dcl.ValueOrEmptyString(resource.TenantProjectId))
	p.SetDataprocServiceAccount(dcl.ValueOrEmptyString(resource.DataprocServiceAccount))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	mOptions := make(map[string]string, len(resource.Options))
	for k, r := range resource.Options {
		mOptions[k] = r
	}
	p.SetOptions(mOptions)
	sAvailableVersion := make([]*alphapb.DatafusionAlphaInstanceAvailableVersion, len(resource.AvailableVersion))
	for i, r := range resource.AvailableVersion {
		sAvailableVersion[i] = DatafusionAlphaInstanceAvailableVersionToProto(&r)
	}
	p.SetAvailableVersion(sAvailableVersion)

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDatafusionAlphaInstanceRequest) (*alphapb.DatafusionAlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyDatafusionAlphaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyDatafusionAlphaInstance(ctx context.Context, request *alphapb.ApplyDatafusionAlphaInstanceRequest) (*alphapb.DatafusionAlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteDatafusionAlphaInstance(ctx context.Context, request *alphapb.DeleteDatafusionAlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListDatafusionAlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListDatafusionAlphaInstance(ctx context.Context, request *alphapb.ListDatafusionAlphaInstanceRequest) (*alphapb.ListDatafusionAlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DatafusionAlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDatafusionAlphaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
