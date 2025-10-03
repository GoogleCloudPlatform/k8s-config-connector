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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/krmapihosting/alpha/krmapihosting_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/krmapihosting/alpha"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToKrmapihostingAlphaInstanceStateEnum(e alphapb.KrmapihostingAlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.KrmapihostingAlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("KrmapihostingAlphaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceBundlesConfig converts a InstanceBundlesConfig object from its proto representation.
func ProtoToKrmapihostingAlphaInstanceBundlesConfig(p *alphapb.KrmapihostingAlphaInstanceBundlesConfig) *alpha.InstanceBundlesConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceBundlesConfig{
		ConfigControllerConfig: ProtoToKrmapihostingAlphaInstanceBundlesConfigConfigControllerConfig(p.GetConfigControllerConfig()),
	}
	return obj
}

// ProtoToInstanceBundlesConfigConfigControllerConfig converts a InstanceBundlesConfigConfigControllerConfig object from its proto representation.
func ProtoToKrmapihostingAlphaInstanceBundlesConfigConfigControllerConfig(p *alphapb.KrmapihostingAlphaInstanceBundlesConfigConfigControllerConfig) *alpha.InstanceBundlesConfigConfigControllerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceBundlesConfigConfigControllerConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToInstanceManagementConfig converts a InstanceManagementConfig object from its proto representation.
func ProtoToKrmapihostingAlphaInstanceManagementConfig(p *alphapb.KrmapihostingAlphaInstanceManagementConfig) *alpha.InstanceManagementConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceManagementConfig{
		StandardManagementConfig: ProtoToKrmapihostingAlphaInstanceManagementConfigStandardManagementConfig(p.GetStandardManagementConfig()),
	}
	return obj
}

// ProtoToInstanceManagementConfigStandardManagementConfig converts a InstanceManagementConfigStandardManagementConfig object from its proto representation.
func ProtoToKrmapihostingAlphaInstanceManagementConfigStandardManagementConfig(p *alphapb.KrmapihostingAlphaInstanceManagementConfigStandardManagementConfig) *alpha.InstanceManagementConfigStandardManagementConfig {
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

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.KrmapihostingAlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:               dcl.StringOrNil(p.GetName()),
		BundlesConfig:      ProtoToKrmapihostingAlphaInstanceBundlesConfig(p.GetBundlesConfig()),
		UsePrivateEndpoint: dcl.Bool(p.GetUsePrivateEndpoint()),
		GkeResourceLink:    dcl.StringOrNil(p.GetGkeResourceLink()),
		State:              ProtoToKrmapihostingAlphaInstanceStateEnum(p.GetState()),
		ManagementConfig:   ProtoToKrmapihostingAlphaInstanceManagementConfig(p.GetManagementConfig()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func KrmapihostingAlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.KrmapihostingAlphaInstanceStateEnum {
	if e == nil {
		return alphapb.KrmapihostingAlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.KrmapihostingAlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.KrmapihostingAlphaInstanceStateEnum(v)
	}
	return alphapb.KrmapihostingAlphaInstanceStateEnum(0)
}

// InstanceBundlesConfigToProto converts a InstanceBundlesConfig object to its proto representation.
func KrmapihostingAlphaInstanceBundlesConfigToProto(o *alpha.InstanceBundlesConfig) *alphapb.KrmapihostingAlphaInstanceBundlesConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.KrmapihostingAlphaInstanceBundlesConfig{}
	p.SetConfigControllerConfig(KrmapihostingAlphaInstanceBundlesConfigConfigControllerConfigToProto(o.ConfigControllerConfig))
	return p
}

// InstanceBundlesConfigConfigControllerConfigToProto converts a InstanceBundlesConfigConfigControllerConfig object to its proto representation.
func KrmapihostingAlphaInstanceBundlesConfigConfigControllerConfigToProto(o *alpha.InstanceBundlesConfigConfigControllerConfig) *alphapb.KrmapihostingAlphaInstanceBundlesConfigConfigControllerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.KrmapihostingAlphaInstanceBundlesConfigConfigControllerConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// InstanceManagementConfigToProto converts a InstanceManagementConfig object to its proto representation.
func KrmapihostingAlphaInstanceManagementConfigToProto(o *alpha.InstanceManagementConfig) *alphapb.KrmapihostingAlphaInstanceManagementConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.KrmapihostingAlphaInstanceManagementConfig{}
	p.SetStandardManagementConfig(KrmapihostingAlphaInstanceManagementConfigStandardManagementConfigToProto(o.StandardManagementConfig))
	return p
}

// InstanceManagementConfigStandardManagementConfigToProto converts a InstanceManagementConfigStandardManagementConfig object to its proto representation.
func KrmapihostingAlphaInstanceManagementConfigStandardManagementConfigToProto(o *alpha.InstanceManagementConfigStandardManagementConfig) *alphapb.KrmapihostingAlphaInstanceManagementConfigStandardManagementConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.KrmapihostingAlphaInstanceManagementConfigStandardManagementConfig{}
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
func InstanceToProto(resource *alpha.Instance) *alphapb.KrmapihostingAlphaInstance {
	p := &alphapb.KrmapihostingAlphaInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetBundlesConfig(KrmapihostingAlphaInstanceBundlesConfigToProto(resource.BundlesConfig))
	p.SetUsePrivateEndpoint(dcl.ValueOrEmptyBool(resource.UsePrivateEndpoint))
	p.SetGkeResourceLink(dcl.ValueOrEmptyString(resource.GkeResourceLink))
	p.SetState(KrmapihostingAlphaInstanceStateEnumToProto(resource.State))
	p.SetManagementConfig(KrmapihostingAlphaInstanceManagementConfigToProto(resource.ManagementConfig))
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
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyKrmapihostingAlphaInstanceRequest) (*alphapb.KrmapihostingAlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyKrmapihostingAlphaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyKrmapihostingAlphaInstance(ctx context.Context, request *alphapb.ApplyKrmapihostingAlphaInstanceRequest) (*alphapb.KrmapihostingAlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteKrmapihostingAlphaInstance(ctx context.Context, request *alphapb.DeleteKrmapihostingAlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListKrmapihostingAlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListKrmapihostingAlphaInstance(ctx context.Context, request *alphapb.ListKrmapihostingAlphaInstanceRequest) (*alphapb.ListKrmapihostingAlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.KrmapihostingAlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListKrmapihostingAlphaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
