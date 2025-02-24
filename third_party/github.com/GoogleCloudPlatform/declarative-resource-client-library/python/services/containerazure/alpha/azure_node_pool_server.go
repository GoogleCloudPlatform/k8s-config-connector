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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/alpha/containerazure_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/alpha"
)

// Server implements the gRPC interface for AzureNodePool.
type AzureNodePoolServer struct{}

// ProtoToAzureNodePoolStateEnum converts a AzureNodePoolStateEnum enum from its proto representation.
func ProtoToContainerazureAlphaAzureNodePoolStateEnum(e alphapb.ContainerazureAlphaAzureNodePoolStateEnum) *alpha.AzureNodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerazureAlphaAzureNodePoolStateEnum_name[int32(e)]; ok {
		e := alpha.AzureNodePoolStateEnum(n[len("ContainerazureAlphaAzureNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAzureNodePoolConfig converts a AzureNodePoolConfig resource from its proto representation.
func ProtoToContainerazureAlphaAzureNodePoolConfig(p *alphapb.ContainerazureAlphaAzureNodePoolConfig) *alpha.AzureNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolConfig{
		VmSize:     dcl.StringOrNil(p.VmSize),
		RootVolume: ProtoToContainerazureAlphaAzureNodePoolConfigRootVolume(p.GetRootVolume()),
		SshConfig:  ProtoToContainerazureAlphaAzureNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	return obj
}

// ProtoToAzureNodePoolConfigRootVolume converts a AzureNodePoolConfigRootVolume resource from its proto representation.
func ProtoToContainerazureAlphaAzureNodePoolConfigRootVolume(p *alphapb.ContainerazureAlphaAzureNodePoolConfigRootVolume) *alpha.AzureNodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolConfigRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureNodePoolConfigSshConfig converts a AzureNodePoolConfigSshConfig resource from its proto representation.
func ProtoToContainerazureAlphaAzureNodePoolConfigSshConfig(p *alphapb.ContainerazureAlphaAzureNodePoolConfigSshConfig) *alpha.AzureNodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToAzureNodePoolAutoscaling converts a AzureNodePoolAutoscaling resource from its proto representation.
func ProtoToContainerazureAlphaAzureNodePoolAutoscaling(p *alphapb.ContainerazureAlphaAzureNodePoolAutoscaling) *alpha.AzureNodePoolAutoscaling {
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
func ProtoToContainerazureAlphaAzureNodePoolMaxPodsConstraint(p *alphapb.ContainerazureAlphaAzureNodePoolMaxPodsConstraint) *alpha.AzureNodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToAzureNodePool converts a AzureNodePool resource from its proto representation.
func ProtoToAzureNodePool(p *alphapb.ContainerazureAlphaAzureNodePool) *alpha.AzureNodePool {
	obj := &alpha.AzureNodePool{
		Name:                  dcl.StringOrNil(p.Name),
		Version:               dcl.StringOrNil(p.Version),
		Config:                ProtoToContainerazureAlphaAzureNodePoolConfig(p.GetConfig()),
		SubnetId:              dcl.StringOrNil(p.SubnetId),
		Autoscaling:           ProtoToContainerazureAlphaAzureNodePoolAutoscaling(p.GetAutoscaling()),
		State:                 ProtoToContainerazureAlphaAzureNodePoolStateEnum(p.GetState()),
		Uid:                   dcl.StringOrNil(p.Uid),
		Reconciling:           dcl.Bool(p.Reconciling),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                  dcl.StringOrNil(p.Etag),
		MaxPodsConstraint:     ProtoToContainerazureAlphaAzureNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		AzureAvailabilityZone: dcl.StringOrNil(p.AzureAvailabilityZone),
		Project:               dcl.StringOrNil(p.Project),
		Location:              dcl.StringOrNil(p.Location),
		AzureCluster:          dcl.StringOrNil(p.AzureCluster),
	}
	return obj
}

// AzureNodePoolStateEnumToProto converts a AzureNodePoolStateEnum enum to its proto representation.
func ContainerazureAlphaAzureNodePoolStateEnumToProto(e *alpha.AzureNodePoolStateEnum) alphapb.ContainerazureAlphaAzureNodePoolStateEnum {
	if e == nil {
		return alphapb.ContainerazureAlphaAzureNodePoolStateEnum(0)
	}
	if v, ok := alphapb.ContainerazureAlphaAzureNodePoolStateEnum_value["AzureNodePoolStateEnum"+string(*e)]; ok {
		return alphapb.ContainerazureAlphaAzureNodePoolStateEnum(v)
	}
	return alphapb.ContainerazureAlphaAzureNodePoolStateEnum(0)
}

// AzureNodePoolConfigToProto converts a AzureNodePoolConfig resource to its proto representation.
func ContainerazureAlphaAzureNodePoolConfigToProto(o *alpha.AzureNodePoolConfig) *alphapb.ContainerazureAlphaAzureNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureNodePoolConfig{
		VmSize:     dcl.ValueOrEmptyString(o.VmSize),
		RootVolume: ContainerazureAlphaAzureNodePoolConfigRootVolumeToProto(o.RootVolume),
		SshConfig:  ContainerazureAlphaAzureNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AzureNodePoolConfigRootVolumeToProto converts a AzureNodePoolConfigRootVolume resource to its proto representation.
func ContainerazureAlphaAzureNodePoolConfigRootVolumeToProto(o *alpha.AzureNodePoolConfigRootVolume) *alphapb.ContainerazureAlphaAzureNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureNodePoolConfigRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureNodePoolConfigSshConfigToProto converts a AzureNodePoolConfigSshConfig resource to its proto representation.
func ContainerazureAlphaAzureNodePoolConfigSshConfigToProto(o *alpha.AzureNodePoolConfigSshConfig) *alphapb.ContainerazureAlphaAzureNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// AzureNodePoolAutoscalingToProto converts a AzureNodePoolAutoscaling resource to its proto representation.
func ContainerazureAlphaAzureNodePoolAutoscalingToProto(o *alpha.AzureNodePoolAutoscaling) *alphapb.ContainerazureAlphaAzureNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// AzureNodePoolMaxPodsConstraintToProto converts a AzureNodePoolMaxPodsConstraint resource to its proto representation.
func ContainerazureAlphaAzureNodePoolMaxPodsConstraintToProto(o *alpha.AzureNodePoolMaxPodsConstraint) *alphapb.ContainerazureAlphaAzureNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// AzureNodePoolToProto converts a AzureNodePool resource to its proto representation.
func AzureNodePoolToProto(resource *alpha.AzureNodePool) *alphapb.ContainerazureAlphaAzureNodePool {
	p := &alphapb.ContainerazureAlphaAzureNodePool{
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		Version:               dcl.ValueOrEmptyString(resource.Version),
		Config:                ContainerazureAlphaAzureNodePoolConfigToProto(resource.Config),
		SubnetId:              dcl.ValueOrEmptyString(resource.SubnetId),
		Autoscaling:           ContainerazureAlphaAzureNodePoolAutoscalingToProto(resource.Autoscaling),
		State:                 ContainerazureAlphaAzureNodePoolStateEnumToProto(resource.State),
		Uid:                   dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:           dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                  dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint:     ContainerazureAlphaAzureNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		AzureAvailabilityZone: dcl.ValueOrEmptyString(resource.AzureAvailabilityZone),
		Project:               dcl.ValueOrEmptyString(resource.Project),
		Location:              dcl.ValueOrEmptyString(resource.Location),
		AzureCluster:          dcl.ValueOrEmptyString(resource.AzureCluster),
	}

	return p
}

// ApplyAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Apply() method.
func (s *AzureNodePoolServer) applyAzureNodePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerazureAlphaAzureNodePoolRequest) (*alphapb.ContainerazureAlphaAzureNodePool, error) {
	p := ProtoToAzureNodePool(request.GetResource())
	res, err := c.ApplyAzureNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureNodePoolToProto(res)
	return r, nil
}

// ApplyAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Apply() method.
func (s *AzureNodePoolServer) ApplyContainerazureAlphaAzureNodePool(ctx context.Context, request *alphapb.ApplyContainerazureAlphaAzureNodePoolRequest) (*alphapb.ContainerazureAlphaAzureNodePool, error) {
	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureNodePool(ctx, cl, request)
}

// DeleteAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePool Delete() method.
func (s *AzureNodePoolServer) DeleteContainerazureAlphaAzureNodePool(ctx context.Context, request *alphapb.DeleteContainerazureAlphaAzureNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureNodePool(ctx, ProtoToAzureNodePool(request.GetResource()))

}

// ListContainerazureAlphaAzureNodePool handles the gRPC request by passing it to the underlying AzureNodePoolList() method.
func (s *AzureNodePoolServer) ListContainerazureAlphaAzureNodePool(ctx context.Context, request *alphapb.ListContainerazureAlphaAzureNodePoolRequest) (*alphapb.ListContainerazureAlphaAzureNodePoolResponse, error) {
	cl, err := createConfigAzureNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureNodePool(ctx, ProtoToAzureNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerazureAlphaAzureNodePool
	for _, r := range resources.Items {
		rp := AzureNodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContainerazureAlphaAzureNodePoolResponse{Items: protos}, nil
}

func createConfigAzureNodePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
