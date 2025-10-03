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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/configcontroller/alpha/configcontroller_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/configcontroller/alpha"
)

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToConfigcontrollerAlphaInstanceStateEnum(e alphapb.ConfigcontrollerAlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ConfigcontrollerAlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("ConfigcontrollerAlphaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceBundlesConfig converts a InstanceBundlesConfig object from its proto representation.
func ProtoToConfigcontrollerAlphaInstanceBundlesConfig(p *alphapb.ConfigcontrollerAlphaInstanceBundlesConfig) *alpha.InstanceBundlesConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceBundlesConfig{
		ConfigControllerConfig: ProtoToConfigcontrollerAlphaInstanceBundlesConfigConfigControllerConfig(p.GetConfigControllerConfig()),
	}
	return obj
}

// ProtoToInstanceBundlesConfigConfigControllerConfig converts a InstanceBundlesConfigConfigControllerConfig object from its proto representation.
func ProtoToConfigcontrollerAlphaInstanceBundlesConfigConfigControllerConfig(p *alphapb.ConfigcontrollerAlphaInstanceBundlesConfigConfigControllerConfig) *alpha.InstanceBundlesConfigConfigControllerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceBundlesConfigConfigControllerConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToInstanceManagementConfig converts a InstanceManagementConfig object from its proto representation.
func ProtoToConfigcontrollerAlphaInstanceManagementConfig(p *alphapb.ConfigcontrollerAlphaInstanceManagementConfig) *alpha.InstanceManagementConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceManagementConfig{
		StandardManagementConfig: ProtoToConfigcontrollerAlphaInstanceManagementConfigStandardManagementConfig(p.GetStandardManagementConfig()),
		FullManagementConfig:     ProtoToConfigcontrollerAlphaInstanceManagementConfigFullManagementConfig(p.GetFullManagementConfig()),
	}
	return obj
}

// ProtoToInstanceManagementConfigStandardManagementConfig converts a InstanceManagementConfigStandardManagementConfig object from its proto representation.
func ProtoToConfigcontrollerAlphaInstanceManagementConfigStandardManagementConfig(p *alphapb.ConfigcontrollerAlphaInstanceManagementConfigStandardManagementConfig) *alpha.InstanceManagementConfigStandardManagementConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceManagementConfigStandardManagementConfig{
		Network:             dcl.StringOrNil(p.GetNetwork()),
		MasterIPv4CidrBlock: dcl.StringOrNil(p.GetMasterIpv4CidrBlock()),
		ManBlock:            dcl.StringOrNil(p.GetManBlock()),
		ClusterCidrBlock:    dcl.StringOrNil(p.GetClusterCidrBlock()),
		ServicesCidrBlock:   dcl.StringOrNil(p.GetServicesCidrBlock()),
		ClusterNamedRange:   dcl.StringOrNil(p.GetClusterNamedRange()),
		ServicesNamedRange:  dcl.StringOrNil(p.GetServicesNamedRange()),
	}
	return obj
}

// ProtoToInstanceManagementConfigFullManagementConfig converts a InstanceManagementConfigFullManagementConfig object from its proto representation.
func ProtoToConfigcontrollerAlphaInstanceManagementConfigFullManagementConfig(p *alphapb.ConfigcontrollerAlphaInstanceManagementConfigFullManagementConfig) *alpha.InstanceManagementConfigFullManagementConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceManagementConfigFullManagementConfig{
		Network:             dcl.StringOrNil(p.GetNetwork()),
		MasterIPv4CidrBlock: dcl.StringOrNil(p.GetMasterIpv4CidrBlock()),
		ManBlock:            dcl.StringOrNil(p.GetManBlock()),
		ClusterCidrBlock:    dcl.StringOrNil(p.GetClusterCidrBlock()),
		ServicesCidrBlock:   dcl.StringOrNil(p.GetServicesCidrBlock()),
		ClusterNamedRange:   dcl.StringOrNil(p.GetClusterNamedRange()),
		ServicesNamedRange:  dcl.StringOrNil(p.GetServicesNamedRange()),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.ConfigcontrollerAlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:               dcl.StringOrNil(p.GetName()),
		BundlesConfig:      ProtoToConfigcontrollerAlphaInstanceBundlesConfig(p.GetBundlesConfig()),
		UsePrivateEndpoint: dcl.Bool(p.GetUsePrivateEndpoint()),
		GkeResourceLink:    dcl.StringOrNil(p.GetGkeResourceLink()),
		State:              ProtoToConfigcontrollerAlphaInstanceStateEnum(p.GetState()),
		ManagementConfig:   ProtoToConfigcontrollerAlphaInstanceManagementConfig(p.GetManagementConfig()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func ConfigcontrollerAlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.ConfigcontrollerAlphaInstanceStateEnum {
	if e == nil {
		return alphapb.ConfigcontrollerAlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.ConfigcontrollerAlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.ConfigcontrollerAlphaInstanceStateEnum(v)
	}
	return alphapb.ConfigcontrollerAlphaInstanceStateEnum(0)
}

// InstanceBundlesConfigToProto converts a InstanceBundlesConfig object to its proto representation.
func ConfigcontrollerAlphaInstanceBundlesConfigToProto(o *alpha.InstanceBundlesConfig) *alphapb.ConfigcontrollerAlphaInstanceBundlesConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ConfigcontrollerAlphaInstanceBundlesConfig{}
	p.SetConfigControllerConfig(ConfigcontrollerAlphaInstanceBundlesConfigConfigControllerConfigToProto(o.ConfigControllerConfig))
	return p
}

// InstanceBundlesConfigConfigControllerConfigToProto converts a InstanceBundlesConfigConfigControllerConfig object to its proto representation.
func ConfigcontrollerAlphaInstanceBundlesConfigConfigControllerConfigToProto(o *alpha.InstanceBundlesConfigConfigControllerConfig) *alphapb.ConfigcontrollerAlphaInstanceBundlesConfigConfigControllerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ConfigcontrollerAlphaInstanceBundlesConfigConfigControllerConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// InstanceManagementConfigToProto converts a InstanceManagementConfig object to its proto representation.
func ConfigcontrollerAlphaInstanceManagementConfigToProto(o *alpha.InstanceManagementConfig) *alphapb.ConfigcontrollerAlphaInstanceManagementConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ConfigcontrollerAlphaInstanceManagementConfig{}
	p.SetStandardManagementConfig(ConfigcontrollerAlphaInstanceManagementConfigStandardManagementConfigToProto(o.StandardManagementConfig))
	p.SetFullManagementConfig(ConfigcontrollerAlphaInstanceManagementConfigFullManagementConfigToProto(o.FullManagementConfig))
	return p
}

// InstanceManagementConfigStandardManagementConfigToProto converts a InstanceManagementConfigStandardManagementConfig object to its proto representation.
func ConfigcontrollerAlphaInstanceManagementConfigStandardManagementConfigToProto(o *alpha.InstanceManagementConfigStandardManagementConfig) *alphapb.ConfigcontrollerAlphaInstanceManagementConfigStandardManagementConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ConfigcontrollerAlphaInstanceManagementConfigStandardManagementConfig{}
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetMasterIpv4CidrBlock(dcl.ValueOrEmptyString(o.MasterIPv4CidrBlock))
	p.SetManBlock(dcl.ValueOrEmptyString(o.ManBlock))
	p.SetClusterCidrBlock(dcl.ValueOrEmptyString(o.ClusterCidrBlock))
	p.SetServicesCidrBlock(dcl.ValueOrEmptyString(o.ServicesCidrBlock))
	p.SetClusterNamedRange(dcl.ValueOrEmptyString(o.ClusterNamedRange))
	p.SetServicesNamedRange(dcl.ValueOrEmptyString(o.ServicesNamedRange))
	return p
}

// InstanceManagementConfigFullManagementConfigToProto converts a InstanceManagementConfigFullManagementConfig object to its proto representation.
func ConfigcontrollerAlphaInstanceManagementConfigFullManagementConfigToProto(o *alpha.InstanceManagementConfigFullManagementConfig) *alphapb.ConfigcontrollerAlphaInstanceManagementConfigFullManagementConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ConfigcontrollerAlphaInstanceManagementConfigFullManagementConfig{}
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetMasterIpv4CidrBlock(dcl.ValueOrEmptyString(o.MasterIPv4CidrBlock))
	p.SetManBlock(dcl.ValueOrEmptyString(o.ManBlock))
	p.SetClusterCidrBlock(dcl.ValueOrEmptyString(o.ClusterCidrBlock))
	p.SetServicesCidrBlock(dcl.ValueOrEmptyString(o.ServicesCidrBlock))
	p.SetClusterNamedRange(dcl.ValueOrEmptyString(o.ClusterNamedRange))
	p.SetServicesNamedRange(dcl.ValueOrEmptyString(o.ServicesNamedRange))
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *alpha.Instance) *alphapb.ConfigcontrollerAlphaInstance {
	p := &alphapb.ConfigcontrollerAlphaInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetBundlesConfig(ConfigcontrollerAlphaInstanceBundlesConfigToProto(resource.BundlesConfig))
	p.SetUsePrivateEndpoint(dcl.ValueOrEmptyBool(resource.UsePrivateEndpoint))
	p.SetGkeResourceLink(dcl.ValueOrEmptyString(resource.GkeResourceLink))
	p.SetState(ConfigcontrollerAlphaInstanceStateEnumToProto(resource.State))
	p.SetManagementConfig(ConfigcontrollerAlphaInstanceManagementConfigToProto(resource.ManagementConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyConfigcontrollerAlphaInstanceRequest) (*alphapb.ConfigcontrollerAlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyConfigcontrollerAlphaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyConfigcontrollerAlphaInstance(ctx context.Context, request *alphapb.ApplyConfigcontrollerAlphaInstanceRequest) (*alphapb.ConfigcontrollerAlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteConfigcontrollerAlphaInstance(ctx context.Context, request *alphapb.DeleteConfigcontrollerAlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListConfigcontrollerAlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListConfigcontrollerAlphaInstance(ctx context.Context, request *alphapb.ListConfigcontrollerAlphaInstanceRequest) (*alphapb.ListConfigcontrollerAlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ConfigcontrollerAlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListConfigcontrollerAlphaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
