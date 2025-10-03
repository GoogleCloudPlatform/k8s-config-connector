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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkemulticloud/beta/gkemulticloud_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/beta"
)

// Server implements the gRPC interface for AzureNodePool.
type AzureNodePoolServer struct{}

// ProtoToAzureNodePoolStateEnum converts a AzureNodePoolStateEnum enum from its proto representation.
func ProtoToGkemulticloudBetaAzureNodePoolStateEnum(e betapb.GkemulticloudBetaAzureNodePoolStateEnum) *beta.AzureNodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkemulticloudBetaAzureNodePoolStateEnum_name[int32(e)]; ok {
		e := beta.AzureNodePoolStateEnum(n[len("GkemulticloudBetaAzureNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAzureNodePoolConfig converts a AzureNodePoolConfig resource from its proto representation.
func ProtoToGkemulticloudBetaAzureNodePoolConfig(p *betapb.GkemulticloudBetaAzureNodePoolConfig) *beta.AzureNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AzureNodePoolConfig{
		VmSize:     dcl.StringOrNil(p.VmSize),
		RootVolume: ProtoToGkemulticloudBetaAzureNodePoolConfigRootVolume(p.GetRootVolume()),
		SshConfig:  ProtoToGkemulticloudBetaAzureNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	return obj
}

// ProtoToAzureNodePoolConfigRootVolume converts a AzureNodePoolConfigRootVolume resource from its proto representation.
func ProtoToGkemulticloudBetaAzureNodePoolConfigRootVolume(p *betapb.GkemulticloudBetaAzureNodePoolConfigRootVolume) *beta.AzureNodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &beta.AzureNodePoolConfigRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureNodePoolConfigSshConfig converts a AzureNodePoolConfigSshConfig resource from its proto representation.
func ProtoToGkemulticloudBetaAzureNodePoolConfigSshConfig(p *betapb.GkemulticloudBetaAzureNodePoolConfigSshConfig) *beta.AzureNodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToAzureNodePoolAutoscaling converts a AzureNodePoolAutoscaling resource from its proto representation.
func ProtoToGkemulticloudBetaAzureNodePoolAutoscaling(p *betapb.GkemulticloudBetaAzureNodePoolAutoscaling) *beta.AzureNodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.AzureNodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount: dcl.Int64OrNil(p.MaxNodeCount),
	}
	return obj
}

// ProtoToAzureNodePoolMaxPodsConstraint converts a AzureNodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToGkemulticloudBetaAzureNodePoolMaxPodsConstraint(p *betapb.GkemulticloudBetaAzureNodePoolMaxPodsConstraint) *beta.AzureNodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &beta.AzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToAzureNodePool converts a AzureNodePool resource from its proto representation.
func ProtoToAzureNodePool(p *betapb.GkemulticloudBetaAzureNodePool) *beta.AzureNodePool {
	obj := &beta.AzureNodePool{
		Name:                  dcl.StringOrNil(p.Name),
		Version:               dcl.StringOrNil(p.Version),
		Config:                ProtoToGkemulticloudBetaAzureNodePoolConfig(p.GetConfig()),
		SubnetId:              dcl.StringOrNil(p.SubnetId),
		Autoscaling:           ProtoToGkemulticloudBetaAzureNodePoolAutoscaling(p.GetAutoscaling()),
		State:                 ProtoToGkemulticloudBetaAzureNodePoolStateEnum(p.GetState()),
		Uid:                   dcl.StringOrNil(p.Uid),
		Reconciling:           dcl.Bool(p.Reconciling),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                  dcl.StringOrNil(p.Etag),
		MaxPodsConstraint:     ProtoToGkemulticloudBetaAzureNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		AzureAvailabilityZone: dcl.StringOrNil(p.AzureAvailabilityZone),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
		AzureCluster:          dcl.StringOrNil(p.AzureCluster),
	}
	return obj
}

// AzureNodePoolStateEnumToProto converts a AzureNodePoolStateEnum enum to its proto representation.
func GkemulticloudBetaAzureNodePoolStateEnumToProto(e *beta.AzureNodePoolStateEnum) betapb.GkemulticloudBetaAzureNodePoolStateEnum {
	if e == nil {
		return betapb.GkemulticloudBetaAzureNodePoolStateEnum(0)
	}
	if v, ok := betapb.GkemulticloudBetaAzureNodePoolStateEnum_value["AzureNodePoolStateEnum"+string(*e)]; ok {
		return betapb.GkemulticloudBetaAzureNodePoolStateEnum(v)
	}
	return betapb.GkemulticloudBetaAzureNodePoolStateEnum(0)
}

// AzureNodePoolConfigToProto converts a AzureNodePoolConfig resource to its proto representation.
func GkemulticloudBetaAzureNodePoolConfigToProto(o *beta.AzureNodePoolConfig) *betapb.GkemulticloudBetaAzureNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureNodePoolConfig{
		VmSize:     dcl.ValueOrEmptyString(o.VmSize),
		RootVolume: GkemulticloudBetaAzureNodePoolConfigRootVolumeToProto(o.RootVolume),
		SshConfig:  GkemulticloudBetaAzureNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AzureNodePoolConfigRootVolumeToProto converts a AzureNodePoolConfigRootVolume resource to its proto representation.
func GkemulticloudBetaAzureNodePoolConfigRootVolumeToProto(o *beta.AzureNodePoolConfigRootVolume) *betapb.GkemulticloudBetaAzureNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureNodePoolConfigRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureNodePoolConfigSshConfigToProto converts a AzureNodePoolConfigSshConfig resource to its proto representation.
func GkemulticloudBetaAzureNodePoolConfigSshConfigToProto(o *beta.AzureNodePoolConfigSshConfig) *betapb.GkemulticloudBetaAzureNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// AzureNodePoolAutoscalingToProto converts a AzureNodePoolAutoscaling resource to its proto representation.
func GkemulticloudBetaAzureNodePoolAutoscalingToProto(o *beta.AzureNodePoolAutoscaling) *betapb.GkemulticloudBetaAzureNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// AzureNodePoolMaxPodsConstraintToProto converts a AzureNodePoolMaxPodsConstraint resource to its proto representation.
func GkemulticloudBetaAzureNodePoolMaxPodsConstraintToProto(o *beta.AzureNodePoolMaxPodsConstraint) *betapb.GkemulticloudBetaAzureNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// AzureNodePoolToProto converts a AzureNodePool resource to its proto representation.
func AzureNodePoolToProto(resource *beta.AzureNodePool) *betapb.GkemulticloudBetaAzureNodePool {
	p := &betapb.GkemulticloudBetaAzureNodePool{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		Version:               dcl.ValueOrEmptyString(resource.Version),
		Config:                GkemulticloudBetaAzureNodePoolConfigToProto(resource.Config),
		SubnetId:              dcl.ValueOrEmptyString(resource.SubnetId),
		Autoscaling:           GkemulticloudBetaAzureNodePoolAutoscalingToProto(resource.Autoscaling),
		State:                 GkemulticloudBetaAzureNodePoolStateEnumToProto(resource.State),
		Uid:                   dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:           dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                  dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint:     GkemulticloudBetaAzureNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		AzureAvailabilityZone: dcl.ValueOrEmptyString(resource.AzureAvailabilityZone),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
		AzureCluster:          dcl.ValueOrEmptyString(resource.AzureCluster),
	}

	return p
}

// ApplyAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Apply() method.
func (s *AzureNodePoolServer) applyAzureNodePool(ctx context.Context, c *beta.Client, request *betapb.ApplyGkemulticloudBetaAzureNodePoolRequest) (*betapb.GkemulticloudBetaAzureNodePool, error) {
	p := ProtoToAzureNodePool(request.GetResource())
	res, err := c.ApplyAzureNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureNodePoolToProto(res)
	return r, nil
}

// ApplyAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Apply() method.
func (s *AzureNodePoolServer) ApplyGkemulticloudBetaAzureNodePool(ctx context.Context, request *betapb.ApplyGkemulticloudBetaAzureNodePoolRequest) (*betapb.GkemulticloudBetaAzureNodePool, error) {
	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureNodePool(ctx, cl, request)
}

// DeleteAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Delete() method.
func (s *AzureNodePoolServer) DeleteGkemulticloudBetaAzureNodePool(ctx context.Context, request *betapb.DeleteGkemulticloudBetaAzureNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureNodePool(ctx, ProtoToAzureNodePool(request.GetResource()))

}

// ListGkemulticloudBetaAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePoolList() method.
func (s *AzureNodePoolServer) ListGkemulticloudBetaAzureNodePool(ctx context.Context, request *betapb.ListGkemulticloudBetaAzureNodePoolRequest) (*betapb.ListGkemulticloudBetaAzureNodePoolResponse, error) {
	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureNodePool(ctx, ProtoToAzureNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkemulticloudBetaAzureNodePool
	for _, r := range resources.Items {
		rp := AzureNodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListGkemulticloudBetaAzureNodePoolResponse{Items: protos}, nil
}

func createConfigAzureNodePool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
