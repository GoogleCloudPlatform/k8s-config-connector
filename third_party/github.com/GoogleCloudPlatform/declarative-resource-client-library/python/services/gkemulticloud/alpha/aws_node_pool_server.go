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

// Server implements the gRPC interface for AwsNodePool.
type AwsNodePoolServer struct{}

// ProtoToAwsNodePoolConfigRootVolumeVolumeTypeEnum converts a AwsNodePoolConfigRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(e alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum) *alpha.AwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.AwsNodePoolConfigRootVolumeVolumeTypeEnum(n[len("GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolConfigTaintsEffectEnum converts a AwsNodePoolConfigTaintsEffectEnum enum from its proto representation.
func ProtoToGkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum(e alphapb.GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum) *alpha.AwsNodePoolConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := alpha.AwsNodePoolConfigTaintsEffectEnum(n[len("GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolStateEnum converts a AwsNodePoolStateEnum enum from its proto representation.
func ProtoToGkemulticloudAlphaAwsNodePoolStateEnum(e alphapb.GkemulticloudAlphaAwsNodePoolStateEnum) *alpha.AwsNodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkemulticloudAlphaAwsNodePoolStateEnum_name[int32(e)]; ok {
		e := alpha.AwsNodePoolStateEnum(n[len("GkemulticloudAlphaAwsNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolConfig converts a AwsNodePoolConfig resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsNodePoolConfig(p *alphapb.GkemulticloudAlphaAwsNodePoolConfig) *alpha.AwsNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolConfig{
		InstanceType:       dcl.StringOrNil(p.InstanceType),
		RootVolume:         ProtoToGkemulticloudAlphaAwsNodePoolConfigRootVolume(p.GetRootVolume()),
		IamInstanceProfile: dcl.StringOrNil(p.IamInstanceProfile),
		SshConfig:          ProtoToGkemulticloudAlphaAwsNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToGkemulticloudAlphaAwsNodePoolConfigTaints(r))
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToAwsNodePoolConfigRootVolume converts a AwsNodePoolConfigRootVolume resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsNodePoolConfigRootVolume(p *alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolume) *alpha.AwsNodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolConfigRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsNodePoolConfigTaints converts a AwsNodePoolConfigTaints resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsNodePoolConfigTaints(p *alphapb.GkemulticloudAlphaAwsNodePoolConfigTaints) *alpha.AwsNodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToGkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToAwsNodePoolConfigSshConfig converts a AwsNodePoolConfigSshConfig resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsNodePoolConfigSshConfig(p *alphapb.GkemulticloudAlphaAwsNodePoolConfigSshConfig) *alpha.AwsNodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToAwsNodePoolAutoscaling converts a AwsNodePoolAutoscaling resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsNodePoolAutoscaling(p *alphapb.GkemulticloudAlphaAwsNodePoolAutoscaling) *alpha.AwsNodePoolAutoscaling {
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
func ProtoToGkemulticloudAlphaAwsNodePoolMaxPodsConstraint(p *alphapb.GkemulticloudAlphaAwsNodePoolMaxPodsConstraint) *alpha.AwsNodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToAwsNodePool converts a AwsNodePool resource from its proto representation.
func ProtoToAwsNodePool(p *alphapb.GkemulticloudAlphaAwsNodePool) *alpha.AwsNodePool {
	obj := &alpha.AwsNodePool{
		Name:              dcl.StringOrNil(p.Name),
		Version:           dcl.StringOrNil(p.Version),
		Config:            ProtoToGkemulticloudAlphaAwsNodePoolConfig(p.GetConfig()),
		Autoscaling:       ProtoToGkemulticloudAlphaAwsNodePoolAutoscaling(p.GetAutoscaling()),
		SubnetId:          dcl.StringOrNil(p.SubnetId),
		State:             ProtoToGkemulticloudAlphaAwsNodePoolStateEnum(p.GetState()),
		Uid:               dcl.StringOrNil(p.Uid),
		Reconciling:       dcl.Bool(p.Reconciling),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Etag:              dcl.StringOrNil(p.Etag),
		MaxPodsConstraint: ProtoToGkemulticloudAlphaAwsNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
		AwsCluster:        dcl.StringOrNil(p.AwsCluster),
	}
	return obj
}

// AwsNodePoolConfigRootVolumeVolumeTypeEnumToProto converts a AwsNodePoolConfigRootVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnumToProto(e *alpha.AwsNodePoolConfigRootVolumeVolumeTypeEnum) alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum_value["AwsNodePoolConfigRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(v)
	}
	return alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnum(0)
}

// AwsNodePoolConfigTaintsEffectEnumToProto converts a AwsNodePoolConfigTaintsEffectEnum enum to its proto representation.
func GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnumToProto(e *alpha.AwsNodePoolConfigTaintsEffectEnum) alphapb.GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum {
	if e == nil {
		return alphapb.GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum(0)
	}
	if v, ok := alphapb.GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum_value["AwsNodePoolConfigTaintsEffectEnum"+string(*e)]; ok {
		return alphapb.GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum(v)
	}
	return alphapb.GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnum(0)
}

// AwsNodePoolStateEnumToProto converts a AwsNodePoolStateEnum enum to its proto representation.
func GkemulticloudAlphaAwsNodePoolStateEnumToProto(e *alpha.AwsNodePoolStateEnum) alphapb.GkemulticloudAlphaAwsNodePoolStateEnum {
	if e == nil {
		return alphapb.GkemulticloudAlphaAwsNodePoolStateEnum(0)
	}
	if v, ok := alphapb.GkemulticloudAlphaAwsNodePoolStateEnum_value["AwsNodePoolStateEnum"+string(*e)]; ok {
		return alphapb.GkemulticloudAlphaAwsNodePoolStateEnum(v)
	}
	return alphapb.GkemulticloudAlphaAwsNodePoolStateEnum(0)
}

// AwsNodePoolConfigToProto converts a AwsNodePoolConfig resource to its proto representation.
func GkemulticloudAlphaAwsNodePoolConfigToProto(o *alpha.AwsNodePoolConfig) *alphapb.GkemulticloudAlphaAwsNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsNodePoolConfig{
		InstanceType:       dcl.ValueOrEmptyString(o.InstanceType),
		RootVolume:         GkemulticloudAlphaAwsNodePoolConfigRootVolumeToProto(o.RootVolume),
		IamInstanceProfile: dcl.ValueOrEmptyString(o.IamInstanceProfile),
		SshConfig:          GkemulticloudAlphaAwsNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, GkemulticloudAlphaAwsNodePoolConfigTaintsToProto(&r))
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
func GkemulticloudAlphaAwsNodePoolConfigRootVolumeToProto(o *alpha.AwsNodePoolConfigRootVolume) *alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsNodePoolConfigRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudAlphaAwsNodePoolConfigRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsNodePoolConfigTaintsToProto converts a AwsNodePoolConfigTaints resource to its proto representation.
func GkemulticloudAlphaAwsNodePoolConfigTaintsToProto(o *alpha.AwsNodePoolConfigTaints) *alphapb.GkemulticloudAlphaAwsNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsNodePoolConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: GkemulticloudAlphaAwsNodePoolConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// AwsNodePoolConfigSshConfigToProto converts a AwsNodePoolConfigSshConfig resource to its proto representation.
func GkemulticloudAlphaAwsNodePoolConfigSshConfigToProto(o *alpha.AwsNodePoolConfigSshConfig) *alphapb.GkemulticloudAlphaAwsNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// AwsNodePoolAutoscalingToProto converts a AwsNodePoolAutoscaling resource to its proto representation.
func GkemulticloudAlphaAwsNodePoolAutoscalingToProto(o *alpha.AwsNodePoolAutoscaling) *alphapb.GkemulticloudAlphaAwsNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// AwsNodePoolMaxPodsConstraintToProto converts a AwsNodePoolMaxPodsConstraint resource to its proto representation.
func GkemulticloudAlphaAwsNodePoolMaxPodsConstraintToProto(o *alpha.AwsNodePoolMaxPodsConstraint) *alphapb.GkemulticloudAlphaAwsNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// AwsNodePoolToProto converts a AwsNodePool resource to its proto representation.
func AwsNodePoolToProto(resource *alpha.AwsNodePool) *alphapb.GkemulticloudAlphaAwsNodePool {
	p := &alphapb.GkemulticloudAlphaAwsNodePool{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Version:           dcl.ValueOrEmptyString(resource.Version),
		Config:            GkemulticloudAlphaAwsNodePoolConfigToProto(resource.Config),
		Autoscaling:       GkemulticloudAlphaAwsNodePoolAutoscalingToProto(resource.Autoscaling),
		SubnetId:          dcl.ValueOrEmptyString(resource.SubnetId),
		State:             GkemulticloudAlphaAwsNodePoolStateEnumToProto(resource.State),
		Uid:               dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:       dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:        dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:        dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:              dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint: GkemulticloudAlphaAwsNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
		AwsCluster:        dcl.ValueOrEmptyString(resource.AwsCluster),
	}

	return p
}

// ApplyAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Apply() method.
func (s *AwsNodePoolServer) applyAwsNodePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkemulticloudAlphaAwsNodePoolRequest) (*alphapb.GkemulticloudAlphaAwsNodePool, error) {
	p := ProtoToAwsNodePool(request.GetResource())
	res, err := c.ApplyAwsNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AwsNodePoolToProto(res)
	return r, nil
}

// ApplyAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Apply() method.
func (s *AwsNodePoolServer) ApplyGkemulticloudAlphaAwsNodePool(ctx context.Context, request *alphapb.ApplyGkemulticloudAlphaAwsNodePoolRequest) (*alphapb.GkemulticloudAlphaAwsNodePool, error) {
	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAwsNodePool(ctx, cl, request)
}

// DeleteAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Delete() method.
func (s *AwsNodePoolServer) DeleteGkemulticloudAlphaAwsNodePool(ctx context.Context, request *alphapb.DeleteGkemulticloudAlphaAwsNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAwsNodePool(ctx, ProtoToAwsNodePool(request.GetResource()))

}

// ListGkemulticloudAlphaAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePoolList() method.
func (s *AwsNodePoolServer) ListGkemulticloudAlphaAwsNodePool(ctx context.Context, request *alphapb.ListGkemulticloudAlphaAwsNodePoolRequest) (*alphapb.ListGkemulticloudAlphaAwsNodePoolResponse, error) {
	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAwsNodePool(ctx, ProtoToAwsNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkemulticloudAlphaAwsNodePool
	for _, r := range resources.Items {
		rp := AwsNodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListGkemulticloudAlphaAwsNodePoolResponse{Items: protos}, nil
}

func createConfigAwsNodePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
