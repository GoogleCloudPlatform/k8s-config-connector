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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkemulticloud/alpha/gkemulticloud_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/alpha"
)

// Server implements the gRPC interface for AzureNodePool.
type AzureNodePoolServer struct{}

// ProtoToAzureNodePoolStateEnum converts a AzureNodePoolStateEnum enum from its proto representation.
func ProtoToGkemulticloudAlphaAzureNodePoolStateEnum(e alphapb.GkemulticloudAlphaAzureNodePoolStateEnum) *alpha.AzureNodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkemulticloudAlphaAzureNodePoolStateEnum_name[int32(e)]; ok {
		e := alpha.AzureNodePoolStateEnum(n[len("GkemulticloudAlphaAzureNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAzureNodePoolConfig converts a AzureNodePoolConfig resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureNodePoolConfig(p *alphapb.GkemulticloudAlphaAzureNodePoolConfig) *alpha.AzureNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolConfig{
		VmSize:     dcl.StringOrNil(p.VmSize),
		RootVolume: ProtoToGkemulticloudAlphaAzureNodePoolConfigRootVolume(p.GetRootVolume()),
		SshConfig:  ProtoToGkemulticloudAlphaAzureNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	return obj
}

// ProtoToAzureNodePoolConfigRootVolume converts a AzureNodePoolConfigRootVolume resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureNodePoolConfigRootVolume(p *alphapb.GkemulticloudAlphaAzureNodePoolConfigRootVolume) *alpha.AzureNodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolConfigRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureNodePoolConfigSshConfig converts a AzureNodePoolConfigSshConfig resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureNodePoolConfigSshConfig(p *alphapb.GkemulticloudAlphaAzureNodePoolConfigSshConfig) *alpha.AzureNodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToAzureNodePoolAutoscaling converts a AzureNodePoolAutoscaling resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureNodePoolAutoscaling(p *alphapb.GkemulticloudAlphaAzureNodePoolAutoscaling) *alpha.AzureNodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount: dcl.Int64OrNil(p.MaxNodeCount),
	}
	return obj
}

// ProtoToAzureNodePoolMaxPodsConstraint converts a AzureNodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureNodePoolMaxPodsConstraint(p *alphapb.GkemulticloudAlphaAzureNodePoolMaxPodsConstraint) *alpha.AzureNodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToAzureNodePool converts a AzureNodePool resource from its proto representation.
func ProtoToAzureNodePool(p *alphapb.GkemulticloudAlphaAzureNodePool) *alpha.AzureNodePool {
	obj := &alpha.AzureNodePool{
		Name:                  dcl.StringOrNil(p.Name),
		Version:               dcl.StringOrNil(p.Version),
		Config:                ProtoToGkemulticloudAlphaAzureNodePoolConfig(p.GetConfig()),
		SubnetId:              dcl.StringOrNil(p.SubnetId),
		Autoscaling:           ProtoToGkemulticloudAlphaAzureNodePoolAutoscaling(p.GetAutoscaling()),
		State:                 ProtoToGkemulticloudAlphaAzureNodePoolStateEnum(p.GetState()),
		Uid:                   dcl.StringOrNil(p.Uid),
		Reconciling:           dcl.Bool(p.Reconciling),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                  dcl.StringOrNil(p.Etag),
		MaxPodsConstraint:     ProtoToGkemulticloudAlphaAzureNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		AzureAvailabilityZone: dcl.StringOrNil(p.AzureAvailabilityZone),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
		AzureCluster:          dcl.StringOrNil(p.AzureCluster),
	}
	return obj
}

// AzureNodePoolStateEnumToProto converts a AzureNodePoolStateEnum enum to its proto representation.
func GkemulticloudAlphaAzureNodePoolStateEnumToProto(e *alpha.AzureNodePoolStateEnum) alphapb.GkemulticloudAlphaAzureNodePoolStateEnum {
	if e == nil {
		return alphapb.GkemulticloudAlphaAzureNodePoolStateEnum(0)
	}
	if v, ok := alphapb.GkemulticloudAlphaAzureNodePoolStateEnum_value["AzureNodePoolStateEnum"+string(*e)]; ok {
		return alphapb.GkemulticloudAlphaAzureNodePoolStateEnum(v)
	}
	return alphapb.GkemulticloudAlphaAzureNodePoolStateEnum(0)
}

// AzureNodePoolConfigToProto converts a AzureNodePoolConfig resource to its proto representation.
func GkemulticloudAlphaAzureNodePoolConfigToProto(o *alpha.AzureNodePoolConfig) *alphapb.GkemulticloudAlphaAzureNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureNodePoolConfig{
		VmSize:     dcl.ValueOrEmptyString(o.VmSize),
		RootVolume: GkemulticloudAlphaAzureNodePoolConfigRootVolumeToProto(o.RootVolume),
		SshConfig:  GkemulticloudAlphaAzureNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AzureNodePoolConfigRootVolumeToProto converts a AzureNodePoolConfigRootVolume resource to its proto representation.
func GkemulticloudAlphaAzureNodePoolConfigRootVolumeToProto(o *alpha.AzureNodePoolConfigRootVolume) *alphapb.GkemulticloudAlphaAzureNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureNodePoolConfigRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureNodePoolConfigSshConfigToProto converts a AzureNodePoolConfigSshConfig resource to its proto representation.
func GkemulticloudAlphaAzureNodePoolConfigSshConfigToProto(o *alpha.AzureNodePoolConfigSshConfig) *alphapb.GkemulticloudAlphaAzureNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// AzureNodePoolAutoscalingToProto converts a AzureNodePoolAutoscaling resource to its proto representation.
func GkemulticloudAlphaAzureNodePoolAutoscalingToProto(o *alpha.AzureNodePoolAutoscaling) *alphapb.GkemulticloudAlphaAzureNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// AzureNodePoolMaxPodsConstraintToProto converts a AzureNodePoolMaxPodsConstraint resource to its proto representation.
func GkemulticloudAlphaAzureNodePoolMaxPodsConstraintToProto(o *alpha.AzureNodePoolMaxPodsConstraint) *alphapb.GkemulticloudAlphaAzureNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// AzureNodePoolToProto converts a AzureNodePool resource to its proto representation.
func AzureNodePoolToProto(resource *alpha.AzureNodePool) *alphapb.GkemulticloudAlphaAzureNodePool {
	p := &alphapb.GkemulticloudAlphaAzureNodePool{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		Version:               dcl.ValueOrEmptyString(resource.Version),
		Config:                GkemulticloudAlphaAzureNodePoolConfigToProto(resource.Config),
		SubnetId:              dcl.ValueOrEmptyString(resource.SubnetId),
		Autoscaling:           GkemulticloudAlphaAzureNodePoolAutoscalingToProto(resource.Autoscaling),
		State:                 GkemulticloudAlphaAzureNodePoolStateEnumToProto(resource.State),
		Uid:                   dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:           dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                  dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint:     GkemulticloudAlphaAzureNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		AzureAvailabilityZone: dcl.ValueOrEmptyString(resource.AzureAvailabilityZone),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
		AzureCluster:          dcl.ValueOrEmptyString(resource.AzureCluster),
	}

	return p
}

// ApplyAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Apply() method.
func (s *AzureNodePoolServer) applyAzureNodePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkemulticloudAlphaAzureNodePoolRequest) (*alphapb.GkemulticloudAlphaAzureNodePool, error) {
	p := ProtoToAzureNodePool(request.GetResource())
	res, err := c.ApplyAzureNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureNodePoolToProto(res)
	return r, nil
}

// ApplyAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Apply() method.
func (s *AzureNodePoolServer) ApplyGkemulticloudAlphaAzureNodePool(ctx context.Context, request *alphapb.ApplyGkemulticloudAlphaAzureNodePoolRequest) (*alphapb.GkemulticloudAlphaAzureNodePool, error) {
	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureNodePool(ctx, cl, request)
}

// DeleteAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Delete() method.
func (s *AzureNodePoolServer) DeleteGkemulticloudAlphaAzureNodePool(ctx context.Context, request *alphapb.DeleteGkemulticloudAlphaAzureNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureNodePool(ctx, ProtoToAzureNodePool(request.GetResource()))

}

// ListGkemulticloudAlphaAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePoolList() method.
func (s *AzureNodePoolServer) ListGkemulticloudAlphaAzureNodePool(ctx context.Context, request *alphapb.ListGkemulticloudAlphaAzureNodePoolRequest) (*alphapb.ListGkemulticloudAlphaAzureNodePoolResponse, error) {
	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureNodePool(ctx, ProtoToAzureNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkemulticloudAlphaAzureNodePool
	for _, r := range resources.Items {
		rp := AzureNodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListGkemulticloudAlphaAzureNodePoolResponse{Items: protos}, nil
}

func createConfigAzureNodePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
