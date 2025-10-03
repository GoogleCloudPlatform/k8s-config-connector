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
	gkemulticloudpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkemulticloud/gkemulticloud_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud"
)

// Server implements the gRPC interface for AzureNodePool.
type AzureNodePoolServer struct{}

// ProtoToAzureNodePoolStateEnum converts a AzureNodePoolStateEnum enum from its proto representation.
func ProtoToGkemulticloudAzureNodePoolStateEnum(e gkemulticloudpb.GkemulticloudAzureNodePoolStateEnum) *gkemulticloud.AzureNodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkemulticloudpb.GkemulticloudAzureNodePoolStateEnum_name[int32(e)]; ok {
		e := gkemulticloud.AzureNodePoolStateEnum(n[len("GkemulticloudAzureNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAzureNodePoolConfig converts a AzureNodePoolConfig resource from its proto representation.
func ProtoToGkemulticloudAzureNodePoolConfig(p *gkemulticloudpb.GkemulticloudAzureNodePoolConfig) *gkemulticloud.AzureNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureNodePoolConfig{
		VmSize:     dcl.StringOrNil(p.VmSize),
		RootVolume: ProtoToGkemulticloudAzureNodePoolConfigRootVolume(p.GetRootVolume()),
		SshConfig:  ProtoToGkemulticloudAzureNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	return obj
}

// ProtoToAzureNodePoolConfigRootVolume converts a AzureNodePoolConfigRootVolume resource from its proto representation.
func ProtoToGkemulticloudAzureNodePoolConfigRootVolume(p *gkemulticloudpb.GkemulticloudAzureNodePoolConfigRootVolume) *gkemulticloud.AzureNodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureNodePoolConfigRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureNodePoolConfigSshConfig converts a AzureNodePoolConfigSshConfig resource from its proto representation.
func ProtoToGkemulticloudAzureNodePoolConfigSshConfig(p *gkemulticloudpb.GkemulticloudAzureNodePoolConfigSshConfig) *gkemulticloud.AzureNodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToAzureNodePoolAutoscaling converts a AzureNodePoolAutoscaling resource from its proto representation.
func ProtoToGkemulticloudAzureNodePoolAutoscaling(p *gkemulticloudpb.GkemulticloudAzureNodePoolAutoscaling) *gkemulticloud.AzureNodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureNodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount: dcl.Int64OrNil(p.MaxNodeCount),
	}
	return obj
}

// ProtoToAzureNodePoolMaxPodsConstraint converts a AzureNodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToGkemulticloudAzureNodePoolMaxPodsConstraint(p *gkemulticloudpb.GkemulticloudAzureNodePoolMaxPodsConstraint) *gkemulticloud.AzureNodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToAzureNodePool converts a AzureNodePool resource from its proto representation.
func ProtoToAzureNodePool(p *gkemulticloudpb.GkemulticloudAzureNodePool) *gkemulticloud.AzureNodePool {
	obj := &gkemulticloud.AzureNodePool{
		Name:                  dcl.StringOrNil(p.Name),
		Version:               dcl.StringOrNil(p.Version),
		Config:                ProtoToGkemulticloudAzureNodePoolConfig(p.GetConfig()),
		SubnetId:              dcl.StringOrNil(p.SubnetId),
		Autoscaling:           ProtoToGkemulticloudAzureNodePoolAutoscaling(p.GetAutoscaling()),
		State:                 ProtoToGkemulticloudAzureNodePoolStateEnum(p.GetState()),
		Uid:                   dcl.StringOrNil(p.Uid),
		Reconciling:           dcl.Bool(p.Reconciling),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                  dcl.StringOrNil(p.Etag),
		MaxPodsConstraint:     ProtoToGkemulticloudAzureNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		AzureAvailabilityZone: dcl.StringOrNil(p.AzureAvailabilityZone),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
		AzureCluster:          dcl.StringOrNil(p.AzureCluster),
	}
	return obj
}

// AzureNodePoolStateEnumToProto converts a AzureNodePoolStateEnum enum to its proto representation.
func GkemulticloudAzureNodePoolStateEnumToProto(e *gkemulticloud.AzureNodePoolStateEnum) gkemulticloudpb.GkemulticloudAzureNodePoolStateEnum {
	if e == nil {
		return gkemulticloudpb.GkemulticloudAzureNodePoolStateEnum(0)
	}
	if v, ok := gkemulticloudpb.GkemulticloudAzureNodePoolStateEnum_value["AzureNodePoolStateEnum"+string(*e)]; ok {
		return gkemulticloudpb.GkemulticloudAzureNodePoolStateEnum(v)
	}
	return gkemulticloudpb.GkemulticloudAzureNodePoolStateEnum(0)
}

// AzureNodePoolConfigToProto converts a AzureNodePoolConfig resource to its proto representation.
func GkemulticloudAzureNodePoolConfigToProto(o *gkemulticloud.AzureNodePoolConfig) *gkemulticloudpb.GkemulticloudAzureNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureNodePoolConfig{
		VmSize:     dcl.ValueOrEmptyString(o.VmSize),
		RootVolume: GkemulticloudAzureNodePoolConfigRootVolumeToProto(o.RootVolume),
		SshConfig:  GkemulticloudAzureNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AzureNodePoolConfigRootVolumeToProto converts a AzureNodePoolConfigRootVolume resource to its proto representation.
func GkemulticloudAzureNodePoolConfigRootVolumeToProto(o *gkemulticloud.AzureNodePoolConfigRootVolume) *gkemulticloudpb.GkemulticloudAzureNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureNodePoolConfigRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureNodePoolConfigSshConfigToProto converts a AzureNodePoolConfigSshConfig resource to its proto representation.
func GkemulticloudAzureNodePoolConfigSshConfigToProto(o *gkemulticloud.AzureNodePoolConfigSshConfig) *gkemulticloudpb.GkemulticloudAzureNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// AzureNodePoolAutoscalingToProto converts a AzureNodePoolAutoscaling resource to its proto representation.
func GkemulticloudAzureNodePoolAutoscalingToProto(o *gkemulticloud.AzureNodePoolAutoscaling) *gkemulticloudpb.GkemulticloudAzureNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// AzureNodePoolMaxPodsConstraintToProto converts a AzureNodePoolMaxPodsConstraint resource to its proto representation.
func GkemulticloudAzureNodePoolMaxPodsConstraintToProto(o *gkemulticloud.AzureNodePoolMaxPodsConstraint) *gkemulticloudpb.GkemulticloudAzureNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// AzureNodePoolToProto converts a AzureNodePool resource to its proto representation.
func AzureNodePoolToProto(resource *gkemulticloud.AzureNodePool) *gkemulticloudpb.GkemulticloudAzureNodePool {
	p := &gkemulticloudpb.GkemulticloudAzureNodePool{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		Version:               dcl.ValueOrEmptyString(resource.Version),
		Config:                GkemulticloudAzureNodePoolConfigToProto(resource.Config),
		SubnetId:              dcl.ValueOrEmptyString(resource.SubnetId),
		Autoscaling:           GkemulticloudAzureNodePoolAutoscalingToProto(resource.Autoscaling),
		State:                 GkemulticloudAzureNodePoolStateEnumToProto(resource.State),
		Uid:                   dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:           dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                  dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint:     GkemulticloudAzureNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		AzureAvailabilityZone: dcl.ValueOrEmptyString(resource.AzureAvailabilityZone),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
		AzureCluster:          dcl.ValueOrEmptyString(resource.AzureCluster),
	}

	return p
}

// ApplyAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Apply() method.
func (s *AzureNodePoolServer) applyAzureNodePool(ctx context.Context, c *gkemulticloud.Client, request *gkemulticloudpb.ApplyGkemulticloudAzureNodePoolRequest) (*gkemulticloudpb.GkemulticloudAzureNodePool, error) {
	p := ProtoToAzureNodePool(request.GetResource())
	res, err := c.ApplyAzureNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureNodePoolToProto(res)
	return r, nil
}

// ApplyAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Apply() method.
func (s *AzureNodePoolServer) ApplyGkemulticloudAzureNodePool(ctx context.Context, request *gkemulticloudpb.ApplyGkemulticloudAzureNodePoolRequest) (*gkemulticloudpb.GkemulticloudAzureNodePool, error) {
	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureNodePool(ctx, cl, request)
}

// DeleteAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Delete() method.
func (s *AzureNodePoolServer) DeleteGkemulticloudAzureNodePool(ctx context.Context, request *gkemulticloudpb.DeleteGkemulticloudAzureNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureNodePool(ctx, ProtoToAzureNodePool(request.GetResource()))

}

// ListGkemulticloudAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePoolList() method.
func (s *AzureNodePoolServer) ListGkemulticloudAzureNodePool(ctx context.Context, request *gkemulticloudpb.ListGkemulticloudAzureNodePoolRequest) (*gkemulticloudpb.ListGkemulticloudAzureNodePoolResponse, error) {
	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureNodePool(ctx, ProtoToAzureNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*gkemulticloudpb.GkemulticloudAzureNodePool
	for _, r := range resources.Items {
		rp := AzureNodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &gkemulticloudpb.ListGkemulticloudAzureNodePoolResponse{Items: protos}, nil
}

func createConfigAzureNodePool(ctx context.Context, service_account_file string) (*gkemulticloud.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return gkemulticloud.NewClient(conf), nil
}
