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

// Server implements the gRPC interface for KrmApiHost.
type KrmApiHostServer struct{}

// ProtoToKrmApiHostStateEnum converts a KrmApiHostStateEnum enum from its proto representation.
func ProtoToKrmapihostingAlphaKrmApiHostStateEnum(e alphapb.KrmapihostingAlphaKrmApiHostStateEnum) *alpha.KrmApiHostStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.KrmapihostingAlphaKrmApiHostStateEnum_name[int32(e)]; ok {
		e := alpha.KrmApiHostStateEnum(n[len("KrmapihostingAlphaKrmApiHostStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToKrmApiHostBundlesConfig converts a KrmApiHostBundlesConfig resource from its proto representation.
func ProtoToKrmapihostingAlphaKrmApiHostBundlesConfig(p *alphapb.KrmapihostingAlphaKrmApiHostBundlesConfig) *alpha.KrmApiHostBundlesConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.KrmApiHostBundlesConfig{
		ConfigControllerConfig: ProtoToKrmapihostingAlphaKrmApiHostBundlesConfigConfigControllerConfig(p.GetConfigControllerConfig()),
	}
	return obj
}

// ProtoToKrmApiHostBundlesConfigConfigControllerConfig converts a KrmApiHostBundlesConfigConfigControllerConfig resource from its proto representation.
func ProtoToKrmapihostingAlphaKrmApiHostBundlesConfigConfigControllerConfig(p *alphapb.KrmapihostingAlphaKrmApiHostBundlesConfigConfigControllerConfig) *alpha.KrmApiHostBundlesConfigConfigControllerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.KrmApiHostBundlesConfigConfigControllerConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToKrmApiHostManagementConfig converts a KrmApiHostManagementConfig resource from its proto representation.
func ProtoToKrmapihostingAlphaKrmApiHostManagementConfig(p *alphapb.KrmapihostingAlphaKrmApiHostManagementConfig) *alpha.KrmApiHostManagementConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.KrmApiHostManagementConfig{
		StandardManagementConfig: ProtoToKrmapihostingAlphaKrmApiHostManagementConfigStandardManagementConfig(p.GetStandardManagementConfig()),
	}
	return obj
}

// ProtoToKrmApiHostManagementConfigStandardManagementConfig converts a KrmApiHostManagementConfigStandardManagementConfig resource from its proto representation.
func ProtoToKrmapihostingAlphaKrmApiHostManagementConfigStandardManagementConfig(p *alphapb.KrmapihostingAlphaKrmApiHostManagementConfigStandardManagementConfig) *alpha.KrmApiHostManagementConfigStandardManagementConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.KrmApiHostManagementConfigStandardManagementConfig{
		Network:             dcl.StringOrNil(p.Network),
		MasterIPv4CidrBlock: dcl.StringOrNil(p.MasterIpv4CidrBlock),
		ManBlock:            dcl.StringOrNil(p.ManBlock),
		ClusterCidrBlock:    dcl.StringOrNil(p.ClusterCidrBlock),
		ServicesCidrBlock:   dcl.StringOrNil(p.ServicesCidrBlock),
		ClusterNamedRange:   dcl.StringOrNil(p.ClusterNamedRange),
		ServicesNamedRange:  dcl.StringOrNil(p.ServicesNamedRange),
	}
	return obj
}

// ProtoToKrmApiHost converts a KrmApiHost resource from its proto representation.
func ProtoToKrmApiHost(p *alphapb.KrmapihostingAlphaKrmApiHost) *alpha.KrmApiHost {
	obj := &alpha.KrmApiHost{
		Name:               dcl.StringOrNil(p.Name),
		BundlesConfig:      ProtoToKrmapihostingAlphaKrmApiHostBundlesConfig(p.GetBundlesConfig()),
		UsePrivateEndpoint: dcl.Bool(p.UsePrivateEndpoint),
		GkeResourceLink:    dcl.StringOrNil(p.GkeResourceLink),
		State:              ProtoToKrmapihostingAlphaKrmApiHostStateEnum(p.GetState()),
		ManagementConfig:   ProtoToKrmapihostingAlphaKrmApiHostManagementConfig(p.GetManagementConfig()),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	return obj
}

// KrmApiHostStateEnumToProto converts a KrmApiHostStateEnum enum to its proto representation.
func KrmapihostingAlphaKrmApiHostStateEnumToProto(e *alpha.KrmApiHostStateEnum) alphapb.KrmapihostingAlphaKrmApiHostStateEnum {
	if e == nil {
		return alphapb.KrmapihostingAlphaKrmApiHostStateEnum(0)
	}
	if v, ok := alphapb.KrmapihostingAlphaKrmApiHostStateEnum_value["KrmApiHostStateEnum"+string(*e)]; ok {
		return alphapb.KrmapihostingAlphaKrmApiHostStateEnum(v)
	}
	return alphapb.KrmapihostingAlphaKrmApiHostStateEnum(0)
}

// KrmApiHostBundlesConfigToProto converts a KrmApiHostBundlesConfig resource to its proto representation.
func KrmapihostingAlphaKrmApiHostBundlesConfigToProto(o *alpha.KrmApiHostBundlesConfig) *alphapb.KrmapihostingAlphaKrmApiHostBundlesConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.KrmapihostingAlphaKrmApiHostBundlesConfig{
		ConfigControllerConfig: KrmapihostingAlphaKrmApiHostBundlesConfigConfigControllerConfigToProto(o.ConfigControllerConfig),
	}
	return p
}

// KrmApiHostBundlesConfigConfigControllerConfigToProto converts a KrmApiHostBundlesConfigConfigControllerConfig resource to its proto representation.
func KrmapihostingAlphaKrmApiHostBundlesConfigConfigControllerConfigToProto(o *alpha.KrmApiHostBundlesConfigConfigControllerConfig) *alphapb.KrmapihostingAlphaKrmApiHostBundlesConfigConfigControllerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.KrmapihostingAlphaKrmApiHostBundlesConfigConfigControllerConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// KrmApiHostManagementConfigToProto converts a KrmApiHostManagementConfig resource to its proto representation.
func KrmapihostingAlphaKrmApiHostManagementConfigToProto(o *alpha.KrmApiHostManagementConfig) *alphapb.KrmapihostingAlphaKrmApiHostManagementConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.KrmapihostingAlphaKrmApiHostManagementConfig{
		StandardManagementConfig: KrmapihostingAlphaKrmApiHostManagementConfigStandardManagementConfigToProto(o.StandardManagementConfig),
	}
	return p
}

// KrmApiHostManagementConfigStandardManagementConfigToProto converts a KrmApiHostManagementConfigStandardManagementConfig resource to its proto representation.
func KrmapihostingAlphaKrmApiHostManagementConfigStandardManagementConfigToProto(o *alpha.KrmApiHostManagementConfigStandardManagementConfig) *alphapb.KrmapihostingAlphaKrmApiHostManagementConfigStandardManagementConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.KrmapihostingAlphaKrmApiHostManagementConfigStandardManagementConfig{
		Network:             dcl.ValueOrEmptyString(o.Network),
		MasterIpv4CidrBlock: dcl.ValueOrEmptyString(o.MasterIPv4CidrBlock),
		ManBlock:            dcl.ValueOrEmptyString(o.ManBlock),
		ClusterCidrBlock:    dcl.ValueOrEmptyString(o.ClusterCidrBlock),
		ServicesCidrBlock:   dcl.ValueOrEmptyString(o.ServicesCidrBlock),
		ClusterNamedRange:   dcl.ValueOrEmptyString(o.ClusterNamedRange),
		ServicesNamedRange:  dcl.ValueOrEmptyString(o.ServicesNamedRange),
	}
	return p
}

// KrmApiHostToProto converts a KrmApiHost resource to its proto representation.
func KrmApiHostToProto(resource *alpha.KrmApiHost) *alphapb.KrmapihostingAlphaKrmApiHost {
	p := &alphapb.KrmapihostingAlphaKrmApiHost{
		Name:               dcl.ValueOrEmptyString(resource.Name),
		BundlesConfig:      KrmapihostingAlphaKrmApiHostBundlesConfigToProto(resource.BundlesConfig),
		UsePrivateEndpoint: dcl.ValueOrEmptyBool(resource.UsePrivateEndpoint),
		GkeResourceLink:    dcl.ValueOrEmptyString(resource.GkeResourceLink),
		State:              KrmapihostingAlphaKrmApiHostStateEnumToProto(resource.State),
		ManagementConfig:   KrmapihostingAlphaKrmApiHostManagementConfigToProto(resource.ManagementConfig),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyKrmApiHost handles the gRPC request by passing it to the underlying KrmApiHost Apply() method.
func (s *KrmApiHostServer) applyKrmApiHost(ctx context.Context, c *alpha.Client, request *alphapb.ApplyKrmapihostingAlphaKrmApiHostRequest) (*alphapb.KrmapihostingAlphaKrmApiHost, error) {
	p := ProtoToKrmApiHost(request.GetResource())
	res, err := c.ApplyKrmApiHost(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KrmApiHostToProto(res)
	return r, nil
}

// ApplyKrmApiHost handles the gRPC request by passing it to the underlying KrmApiHost Apply() method.
func (s *KrmApiHostServer) ApplyKrmapihostingAlphaKrmApiHost(ctx context.Context, request *alphapb.ApplyKrmapihostingAlphaKrmApiHostRequest) (*alphapb.KrmapihostingAlphaKrmApiHost, error) {
	cl, err := createConfigKrmApiHost(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyKrmApiHost(ctx, cl, request)
}

// DeleteKrmApiHost handles the gRPC request by passing it to the underlying KrmApiHost Delete() method.
func (s *KrmApiHostServer) DeleteKrmapihostingAlphaKrmApiHost(ctx context.Context, request *alphapb.DeleteKrmapihostingAlphaKrmApiHostRequest) (*emptypb.Empty, error) {

	cl, err := createConfigKrmApiHost(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteKrmApiHost(ctx, ProtoToKrmApiHost(request.GetResource()))

}

// ListKrmapihostingAlphaKrmApiHost handles the gRPC request by passing it to the underlying KrmApiHostList() method.
func (s *KrmApiHostServer) ListKrmapihostingAlphaKrmApiHost(ctx context.Context, request *alphapb.ListKrmapihostingAlphaKrmApiHostRequest) (*alphapb.ListKrmapihostingAlphaKrmApiHostResponse, error) {
	cl, err := createConfigKrmApiHost(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKrmApiHost(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.KrmapihostingAlphaKrmApiHost
	for _, r := range resources.Items {
		rp := KrmApiHostToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListKrmapihostingAlphaKrmApiHostResponse{Items: protos}, nil
}

func createConfigKrmApiHost(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
