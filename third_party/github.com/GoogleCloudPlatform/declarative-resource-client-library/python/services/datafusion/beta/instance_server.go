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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/datafusion/beta/datafusion_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta"
)

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceTypeEnum converts a InstanceTypeEnum enum from its proto representation.
func ProtoToDatafusionBetaInstanceTypeEnum(e betapb.DatafusionBetaInstanceTypeEnum) *beta.InstanceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DatafusionBetaInstanceTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceTypeEnum(n[len("DatafusionBetaInstanceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToDatafusionBetaInstanceStateEnum(e betapb.DatafusionBetaInstanceStateEnum) *beta.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DatafusionBetaInstanceStateEnum_name[int32(e)]; ok {
		e := beta.InstanceStateEnum(n[len("DatafusionBetaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkConfig converts a InstanceNetworkConfig object from its proto representation.
func ProtoToDatafusionBetaInstanceNetworkConfig(p *betapb.DatafusionBetaInstanceNetworkConfig) *beta.InstanceNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkConfig{
		Network:      dcl.StringOrNil(p.GetNetwork()),
		IPAllocation: dcl.StringOrNil(p.GetIpAllocation()),
	}
	return obj
}

// ProtoToInstanceAvailableVersion converts a InstanceAvailableVersion object from its proto representation.
func ProtoToDatafusionBetaInstanceAvailableVersion(p *betapb.DatafusionBetaInstanceAvailableVersion) *beta.InstanceAvailableVersion {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceAvailableVersion{
		VersionNumber:  dcl.StringOrNil(p.GetVersionNumber()),
		DefaultVersion: dcl.Bool(p.GetDefaultVersion()),
	}
	for _, r := range p.GetAvailableFeatures() {
		obj.AvailableFeatures = append(obj.AvailableFeatures, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.DatafusionBetaInstance) *beta.Instance {
	obj := &beta.Instance{
		Name:                        dcl.StringOrNil(p.GetName()),
		Description:                 dcl.StringOrNil(p.GetDescription()),
		Type:                        ProtoToDatafusionBetaInstanceTypeEnum(p.GetType()),
		EnableStackdriverLogging:    dcl.Bool(p.GetEnableStackdriverLogging()),
		EnableStackdriverMonitoring: dcl.Bool(p.GetEnableStackdriverMonitoring()),
		PrivateInstance:             dcl.Bool(p.GetPrivateInstance()),
		NetworkConfig:               ProtoToDatafusionBetaInstanceNetworkConfig(p.GetNetworkConfig()),
		CreateTime:                  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                  dcl.StringOrNil(p.GetUpdateTime()),
		State:                       ProtoToDatafusionBetaInstanceStateEnum(p.GetState()),
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
		obj.AvailableVersion = append(obj.AvailableVersion, *ProtoToDatafusionBetaInstanceAvailableVersion(r))
	}
	return obj
}

// InstanceTypeEnumToProto converts a InstanceTypeEnum enum to its proto representation.
func DatafusionBetaInstanceTypeEnumToProto(e *beta.InstanceTypeEnum) betapb.DatafusionBetaInstanceTypeEnum {
	if e == nil {
		return betapb.DatafusionBetaInstanceTypeEnum(0)
	}
	if v, ok := betapb.DatafusionBetaInstanceTypeEnum_value["InstanceTypeEnum"+string(*e)]; ok {
		return betapb.DatafusionBetaInstanceTypeEnum(v)
	}
	return betapb.DatafusionBetaInstanceTypeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func DatafusionBetaInstanceStateEnumToProto(e *beta.InstanceStateEnum) betapb.DatafusionBetaInstanceStateEnum {
	if e == nil {
		return betapb.DatafusionBetaInstanceStateEnum(0)
	}
	if v, ok := betapb.DatafusionBetaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return betapb.DatafusionBetaInstanceStateEnum(v)
	}
	return betapb.DatafusionBetaInstanceStateEnum(0)
}

// InstanceNetworkConfigToProto converts a InstanceNetworkConfig object to its proto representation.
func DatafusionBetaInstanceNetworkConfigToProto(o *beta.InstanceNetworkConfig) *betapb.DatafusionBetaInstanceNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DatafusionBetaInstanceNetworkConfig{}
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetIpAllocation(dcl.ValueOrEmptyString(o.IPAllocation))
	return p
}

// InstanceAvailableVersionToProto converts a InstanceAvailableVersion object to its proto representation.
func DatafusionBetaInstanceAvailableVersionToProto(o *beta.InstanceAvailableVersion) *betapb.DatafusionBetaInstanceAvailableVersion {
	if o == nil {
		return nil
	}
	p := &betapb.DatafusionBetaInstanceAvailableVersion{}
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
func InstanceToProto(resource *beta.Instance) *betapb.DatafusionBetaInstance {
	p := &betapb.DatafusionBetaInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetType(DatafusionBetaInstanceTypeEnumToProto(resource.Type))
	p.SetEnableStackdriverLogging(dcl.ValueOrEmptyBool(resource.EnableStackdriverLogging))
	p.SetEnableStackdriverMonitoring(dcl.ValueOrEmptyBool(resource.EnableStackdriverMonitoring))
	p.SetPrivateInstance(dcl.ValueOrEmptyBool(resource.PrivateInstance))
	p.SetNetworkConfig(DatafusionBetaInstanceNetworkConfigToProto(resource.NetworkConfig))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetState(DatafusionBetaInstanceStateEnumToProto(resource.State))
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
	sAvailableVersion := make([]*betapb.DatafusionBetaInstanceAvailableVersion, len(resource.AvailableVersion))
	for i, r := range resource.AvailableVersion {
		sAvailableVersion[i] = DatafusionBetaInstanceAvailableVersionToProto(&r)
	}
	p.SetAvailableVersion(sAvailableVersion)

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplyDatafusionBetaInstanceRequest) (*betapb.DatafusionBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyDatafusionBetaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyDatafusionBetaInstance(ctx context.Context, request *betapb.ApplyDatafusionBetaInstanceRequest) (*betapb.DatafusionBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteDatafusionBetaInstance(ctx context.Context, request *betapb.DeleteDatafusionBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListDatafusionBetaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListDatafusionBetaInstance(ctx context.Context, request *betapb.ListDatafusionBetaInstanceRequest) (*betapb.ListDatafusionBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DatafusionBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDatafusionBetaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
