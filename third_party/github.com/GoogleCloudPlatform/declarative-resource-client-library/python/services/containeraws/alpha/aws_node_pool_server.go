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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeraws/alpha/containeraws_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/alpha"
)

// Server implements the gRPC interface for AwsNodePool.
type AwsNodePoolServer struct{}

// ProtoToAwsNodePoolConfigRootVolumeVolumeTypeEnum converts a AwsNodePoolConfigRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum) *alpha.AwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.AwsNodePoolConfigRootVolumeVolumeTypeEnum(n[len("ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolConfigTaintsEffectEnum converts a AwsNodePoolConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum(e alphapb.ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum) *alpha.AwsNodePoolConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := alpha.AwsNodePoolConfigTaintsEffectEnum(n[len("ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolStateEnum converts a AwsNodePoolStateEnum enum from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolStateEnum(e alphapb.ContainerawsAlphaAwsNodePoolStateEnum) *alpha.AwsNodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaAwsNodePoolStateEnum_name[int32(e)]; ok {
		e := alpha.AwsNodePoolStateEnum(n[len("ContainerawsAlphaAwsNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolConfig converts a AwsNodePoolConfig resource from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolConfig(p *alphapb.ContainerawsAlphaAwsNodePoolConfig) *alpha.AwsNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolConfig{
		InstanceType:       dcl.StringOrNil(p.InstanceType),
		RootVolume:         ProtoToContainerawsAlphaAwsNodePoolConfigRootVolume(p.GetRootVolume()),
		IamInstanceProfile: dcl.StringOrNil(p.IamInstanceProfile),
		SshConfig:          ProtoToContainerawsAlphaAwsNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerawsAlphaAwsNodePoolConfigTaints(r))
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToAwsNodePoolConfigRootVolume converts a AwsNodePoolConfigRootVolume resource from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolConfigRootVolume(p *alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolume) *alpha.AwsNodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolConfigRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsNodePoolConfigTaints converts a AwsNodePoolConfigTaints resource from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolConfigTaints(p *alphapb.ContainerawsAlphaAwsNodePoolConfigTaints) *alpha.AwsNodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToAwsNodePoolConfigSshConfig converts a AwsNodePoolConfigSshConfig resource from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolConfigSshConfig(p *alphapb.ContainerawsAlphaAwsNodePoolConfigSshConfig) *alpha.AwsNodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToAwsNodePoolAutoscaling converts a AwsNodePoolAutoscaling resource from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolAutoscaling(p *alphapb.ContainerawsAlphaAwsNodePoolAutoscaling) *alpha.AwsNodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount: dcl.Int64OrNil(p.MaxNodeCount),
	}
	return obj
}

// ProtoToAwsNodePoolMaxPodsConstraint converts a AwsNodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToContainerawsAlphaAwsNodePoolMaxPodsConstraint(p *alphapb.ContainerawsAlphaAwsNodePoolMaxPodsConstraint) *alpha.AwsNodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToAwsNodePool converts a AwsNodePool resource from its proto representation.
func ProtoToAwsNodePool(p *alphapb.ContainerawsAlphaAwsNodePool) *alpha.AwsNodePool {
	obj := &alpha.AwsNodePool{
		Name:              dcl.StringOrNil(p.Name),
		Version:           dcl.StringOrNil(p.Version),
		Config:            ProtoToContainerawsAlphaAwsNodePoolConfig(p.GetConfig()),
		Autoscaling:       ProtoToContainerawsAlphaAwsNodePoolAutoscaling(p.GetAutoscaling()),
		SubnetId:          dcl.StringOrNil(p.SubnetId),
		State:             ProtoToContainerawsAlphaAwsNodePoolStateEnum(p.GetState()),
		Uid:               dcl.StringOrNil(p.Uid),
		Reconciling:       dcl.Bool(p.Reconciling),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Etag:              dcl.StringOrNil(p.Etag),
		MaxPodsConstraint: ProtoToContainerawsAlphaAwsNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
		AwsCluster:        dcl.StringOrNil(p.AwsCluster),
	}
	return obj
}

// AwsNodePoolConfigRootVolumeVolumeTypeEnumToProto converts a AwsNodePoolConfigRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnumToProto(e *alpha.AwsNodePoolConfigRootVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum_value["AwsNodePoolConfigRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(0)
}

// AwsNodePoolConfigTaintsEffectEnumToProto converts a AwsNodePoolConfigTaintsEffectEnum enum to its proto representation.
func ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnumToProto(e *alpha.AwsNodePoolConfigTaintsEffectEnum) alphapb.ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum_value["AwsNodePoolConfigTaintsEffectEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum(v)
	}
	return alphapb.ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnum(0)
}

// AwsNodePoolStateEnumToProto converts a AwsNodePoolStateEnum enum to its proto representation.
func ContainerawsAlphaAwsNodePoolStateEnumToProto(e *alpha.AwsNodePoolStateEnum) alphapb.ContainerawsAlphaAwsNodePoolStateEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaAwsNodePoolStateEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaAwsNodePoolStateEnum_value["AwsNodePoolStateEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaAwsNodePoolStateEnum(v)
	}
	return alphapb.ContainerawsAlphaAwsNodePoolStateEnum(0)
}

// AwsNodePoolConfigToProto converts a AwsNodePoolConfig resource to its proto representation.
func ContainerawsAlphaAwsNodePoolConfigToProto(o *alpha.AwsNodePoolConfig) *alphapb.ContainerawsAlphaAwsNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsNodePoolConfig{
		InstanceType:       dcl.ValueOrEmptyString(o.InstanceType),
		RootVolume:         ContainerawsAlphaAwsNodePoolConfigRootVolumeToProto(o.RootVolume),
		IamInstanceProfile: dcl.ValueOrEmptyString(o.IamInstanceProfile),
		SshConfig:          ContainerawsAlphaAwsNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, ContainerawsAlphaAwsNodePoolConfigTaintsToProto(&r))
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	for _, r := range o.SecurityGroupIds {
		p.SecurityGroupIds = append(p.SecurityGroupIds, r)
	}
	return p
}

// AwsNodePoolConfigRootVolumeToProto converts a AwsNodePoolConfigRootVolume resource to its proto representation.
func ContainerawsAlphaAwsNodePoolConfigRootVolumeToProto(o *alpha.AwsNodePoolConfigRootVolume) *alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsNodePoolConfigRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: ContainerawsAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsNodePoolConfigTaintsToProto converts a AwsNodePoolConfigTaints resource to its proto representation.
func ContainerawsAlphaAwsNodePoolConfigTaintsToProto(o *alpha.AwsNodePoolConfigTaints) *alphapb.ContainerawsAlphaAwsNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsNodePoolConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: ContainerawsAlphaAwsNodePoolConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// AwsNodePoolConfigSshConfigToProto converts a AwsNodePoolConfigSshConfig resource to its proto representation.
func ContainerawsAlphaAwsNodePoolConfigSshConfigToProto(o *alpha.AwsNodePoolConfigSshConfig) *alphapb.ContainerawsAlphaAwsNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// AwsNodePoolAutoscalingToProto converts a AwsNodePoolAutoscaling resource to its proto representation.
func ContainerawsAlphaAwsNodePoolAutoscalingToProto(o *alpha.AwsNodePoolAutoscaling) *alphapb.ContainerawsAlphaAwsNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// AwsNodePoolMaxPodsConstraintToProto converts a AwsNodePoolMaxPodsConstraint resource to its proto representation.
func ContainerawsAlphaAwsNodePoolMaxPodsConstraintToProto(o *alpha.AwsNodePoolMaxPodsConstraint) *alphapb.ContainerawsAlphaAwsNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// AwsNodePoolToProto converts a AwsNodePool resource to its proto representation.
func AwsNodePoolToProto(resource *alpha.AwsNodePool) *alphapb.ContainerawsAlphaAwsNodePool {
	p := &alphapb.ContainerawsAlphaAwsNodePool{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Version:           dcl.ValueOrEmptyString(resource.Version),
		Config:            ContainerawsAlphaAwsNodePoolConfigToProto(resource.Config),
		Autoscaling:       ContainerawsAlphaAwsNodePoolAutoscalingToProto(resource.Autoscaling),
		SubnetId:          dcl.ValueOrEmptyString(resource.SubnetId),
		State:             ContainerawsAlphaAwsNodePoolStateEnumToProto(resource.State),
		Uid:               dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:       dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:        dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:        dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:              dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint: ContainerawsAlphaAwsNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
		AwsCluster:        dcl.ValueOrEmptyString(resource.AwsCluster),
	}

	return p
}

// ApplyAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Apply() method.
func (s *AwsNodePoolServer) applyAwsNodePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerawsAlphaAwsNodePoolRequest) (*alphapb.ContainerawsAlphaAwsNodePool, error) {
	p := ProtoToAwsNodePool(request.GetResource())
	res, err := c.ApplyAwsNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AwsNodePoolToProto(res)
	return r, nil
}

// ApplyAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Apply() method.
func (s *AwsNodePoolServer) ApplyContainerawsAlphaAwsNodePool(ctx context.Context, request *alphapb.ApplyContainerawsAlphaAwsNodePoolRequest) (*alphapb.ContainerawsAlphaAwsNodePool, error) {
	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAwsNodePool(ctx, cl, request)
}

// DeleteAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Delete() method.
func (s *AwsNodePoolServer) DeleteContainerawsAlphaAwsNodePool(ctx context.Context, request *alphapb.DeleteContainerawsAlphaAwsNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAwsNodePool(ctx, ProtoToAwsNodePool(request.GetResource()))

}

// ListContainerawsAlphaAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePoolList() method.
func (s *AwsNodePoolServer) ListContainerawsAlphaAwsNodePool(ctx context.Context, request *alphapb.ListContainerawsAlphaAwsNodePoolRequest) (*alphapb.ListContainerawsAlphaAwsNodePoolResponse, error) {
	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAwsNodePool(ctx, ProtoToAwsNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerawsAlphaAwsNodePool
	for _, r := range resources.Items {
		rp := AwsNodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContainerawsAlphaAwsNodePoolResponse{Items: protos}, nil
}

func createConfigAwsNodePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
