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

// Server implements the gRPC interface for AwsNodePool.
type AwsNodePoolServer struct{}

// ProtoToAwsNodePoolConfigRootVolumeVolumeTypeEnum converts a AwsNodePoolConfigRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum(e betapb.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum) *beta.AwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := beta.AwsNodePoolConfigRootVolumeVolumeTypeEnum(n[len("GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolConfigTaintsEffectEnum converts a AwsNodePoolConfigTaintsEffectEnum enum from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum(e betapb.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum) *beta.AwsNodePoolConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := beta.AwsNodePoolConfigTaintsEffectEnum(n[len("GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolStateEnum converts a AwsNodePoolStateEnum enum from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolStateEnum(e betapb.GkemulticloudBetaAwsNodePoolStateEnum) *beta.AwsNodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkemulticloudBetaAwsNodePoolStateEnum_name[int32(e)]; ok {
		e := beta.AwsNodePoolStateEnum(n[len("GkemulticloudBetaAwsNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolConfig converts a AwsNodePoolConfig resource from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolConfig(p *betapb.GkemulticloudBetaAwsNodePoolConfig) *beta.AwsNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AwsNodePoolConfig{
		InstanceType:       dcl.StringOrNil(p.InstanceType),
		RootVolume:         ProtoToGkemulticloudBetaAwsNodePoolConfigRootVolume(p.GetRootVolume()),
		IamInstanceProfile: dcl.StringOrNil(p.IamInstanceProfile),
		SshConfig:          ProtoToGkemulticloudBetaAwsNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToGkemulticloudBetaAwsNodePoolConfigTaints(r))
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToAwsNodePoolConfigRootVolume converts a AwsNodePoolConfigRootVolume resource from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolConfigRootVolume(p *betapb.GkemulticloudBetaAwsNodePoolConfigRootVolume) *beta.AwsNodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &beta.AwsNodePoolConfigRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsNodePoolConfigTaints converts a AwsNodePoolConfigTaints resource from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolConfigTaints(p *betapb.GkemulticloudBetaAwsNodePoolConfigTaints) *beta.AwsNodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &beta.AwsNodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToGkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToAwsNodePoolConfigSshConfig converts a AwsNodePoolConfigSshConfig resource from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolConfigSshConfig(p *betapb.GkemulticloudBetaAwsNodePoolConfigSshConfig) *beta.AwsNodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToAwsNodePoolAutoscaling converts a AwsNodePoolAutoscaling resource from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolAutoscaling(p *betapb.GkemulticloudBetaAwsNodePoolAutoscaling) *beta.AwsNodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.AwsNodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount: dcl.Int64OrNil(p.MaxNodeCount),
	}
	return obj
}

// ProtoToAwsNodePoolMaxPodsConstraint converts a AwsNodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToGkemulticloudBetaAwsNodePoolMaxPodsConstraint(p *betapb.GkemulticloudBetaAwsNodePoolMaxPodsConstraint) *beta.AwsNodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &beta.AwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToAwsNodePool converts a AwsNodePool resource from its proto representation.
func ProtoToAwsNodePool(p *betapb.GkemulticloudBetaAwsNodePool) *beta.AwsNodePool {
	obj := &beta.AwsNodePool{
		Name:              dcl.StringOrNil(p.Name),
		Version:           dcl.StringOrNil(p.Version),
		Config:            ProtoToGkemulticloudBetaAwsNodePoolConfig(p.GetConfig()),
		Autoscaling:       ProtoToGkemulticloudBetaAwsNodePoolAutoscaling(p.GetAutoscaling()),
		SubnetId:          dcl.StringOrNil(p.SubnetId),
		State:             ProtoToGkemulticloudBetaAwsNodePoolStateEnum(p.GetState()),
		Uid:               dcl.StringOrNil(p.Uid),
		Reconciling:       dcl.Bool(p.Reconciling),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Etag:              dcl.StringOrNil(p.Etag),
		MaxPodsConstraint: ProtoToGkemulticloudBetaAwsNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
		AwsCluster:        dcl.StringOrNil(p.AwsCluster),
	}
	return obj
}

// AwsNodePoolConfigRootVolumeVolumeTypeEnumToProto converts a AwsNodePoolConfigRootVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnumToProto(e *beta.AwsNodePoolConfigRootVolumeVolumeTypeEnum) betapb.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == nil {
		return betapb.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := betapb.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum_value["AwsNodePoolConfigRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return betapb.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum(v)
	}
	return betapb.GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnum(0)
}

// AwsNodePoolConfigTaintsEffectEnumToProto converts a AwsNodePoolConfigTaintsEffectEnum enum to its proto representation.
func GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnumToProto(e *beta.AwsNodePoolConfigTaintsEffectEnum) betapb.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum {
	if e == nil {
		return betapb.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum(0)
	}
	if v, ok := betapb.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum_value["AwsNodePoolConfigTaintsEffectEnum"+string(*e)]; ok {
		return betapb.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum(v)
	}
	return betapb.GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnum(0)
}

// AwsNodePoolStateEnumToProto converts a AwsNodePoolStateEnum enum to its proto representation.
func GkemulticloudBetaAwsNodePoolStateEnumToProto(e *beta.AwsNodePoolStateEnum) betapb.GkemulticloudBetaAwsNodePoolStateEnum {
	if e == nil {
		return betapb.GkemulticloudBetaAwsNodePoolStateEnum(0)
	}
	if v, ok := betapb.GkemulticloudBetaAwsNodePoolStateEnum_value["AwsNodePoolStateEnum"+string(*e)]; ok {
		return betapb.GkemulticloudBetaAwsNodePoolStateEnum(v)
	}
	return betapb.GkemulticloudBetaAwsNodePoolStateEnum(0)
}

// AwsNodePoolConfigToProto converts a AwsNodePoolConfig resource to its proto representation.
func GkemulticloudBetaAwsNodePoolConfigToProto(o *beta.AwsNodePoolConfig) *betapb.GkemulticloudBetaAwsNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsNodePoolConfig{
		InstanceType:       dcl.ValueOrEmptyString(o.InstanceType),
		RootVolume:         GkemulticloudBetaAwsNodePoolConfigRootVolumeToProto(o.RootVolume),
		IamInstanceProfile: dcl.ValueOrEmptyString(o.IamInstanceProfile),
		SshConfig:          GkemulticloudBetaAwsNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, GkemulticloudBetaAwsNodePoolConfigTaintsToProto(&r))
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
func GkemulticloudBetaAwsNodePoolConfigRootVolumeToProto(o *beta.AwsNodePoolConfigRootVolume) *betapb.GkemulticloudBetaAwsNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsNodePoolConfigRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudBetaAwsNodePoolConfigRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsNodePoolConfigTaintsToProto converts a AwsNodePoolConfigTaints resource to its proto representation.
func GkemulticloudBetaAwsNodePoolConfigTaintsToProto(o *beta.AwsNodePoolConfigTaints) *betapb.GkemulticloudBetaAwsNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsNodePoolConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: GkemulticloudBetaAwsNodePoolConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// AwsNodePoolConfigSshConfigToProto converts a AwsNodePoolConfigSshConfig resource to its proto representation.
func GkemulticloudBetaAwsNodePoolConfigSshConfigToProto(o *beta.AwsNodePoolConfigSshConfig) *betapb.GkemulticloudBetaAwsNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// AwsNodePoolAutoscalingToProto converts a AwsNodePoolAutoscaling resource to its proto representation.
func GkemulticloudBetaAwsNodePoolAutoscalingToProto(o *beta.AwsNodePoolAutoscaling) *betapb.GkemulticloudBetaAwsNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// AwsNodePoolMaxPodsConstraintToProto converts a AwsNodePoolMaxPodsConstraint resource to its proto representation.
func GkemulticloudBetaAwsNodePoolMaxPodsConstraintToProto(o *beta.AwsNodePoolMaxPodsConstraint) *betapb.GkemulticloudBetaAwsNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// AwsNodePoolToProto converts a AwsNodePool resource to its proto representation.
func AwsNodePoolToProto(resource *beta.AwsNodePool) *betapb.GkemulticloudBetaAwsNodePool {
	p := &betapb.GkemulticloudBetaAwsNodePool{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Version:           dcl.ValueOrEmptyString(resource.Version),
		Config:            GkemulticloudBetaAwsNodePoolConfigToProto(resource.Config),
		Autoscaling:       GkemulticloudBetaAwsNodePoolAutoscalingToProto(resource.Autoscaling),
		SubnetId:          dcl.ValueOrEmptyString(resource.SubnetId),
		State:             GkemulticloudBetaAwsNodePoolStateEnumToProto(resource.State),
		Uid:               dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:       dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:        dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:        dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:              dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint: GkemulticloudBetaAwsNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
		AwsCluster:        dcl.ValueOrEmptyString(resource.AwsCluster),
	}

	return p
}

// ApplyAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Apply() method.
func (s *AwsNodePoolServer) applyAwsNodePool(ctx context.Context, c *beta.Client, request *betapb.ApplyGkemulticloudBetaAwsNodePoolRequest) (*betapb.GkemulticloudBetaAwsNodePool, error) {
	p := ProtoToAwsNodePool(request.GetResource())
	res, err := c.ApplyAwsNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AwsNodePoolToProto(res)
	return r, nil
}

// ApplyAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Apply() method.
func (s *AwsNodePoolServer) ApplyGkemulticloudBetaAwsNodePool(ctx context.Context, request *betapb.ApplyGkemulticloudBetaAwsNodePoolRequest) (*betapb.GkemulticloudBetaAwsNodePool, error) {
	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAwsNodePool(ctx, cl, request)
}

// DeleteAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Delete() method.
func (s *AwsNodePoolServer) DeleteGkemulticloudBetaAwsNodePool(ctx context.Context, request *betapb.DeleteGkemulticloudBetaAwsNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAwsNodePool(ctx, ProtoToAwsNodePool(request.GetResource()))

}

// ListGkemulticloudBetaAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePoolList() method.
func (s *AwsNodePoolServer) ListGkemulticloudBetaAwsNodePool(ctx context.Context, request *betapb.ListGkemulticloudBetaAwsNodePoolRequest) (*betapb.ListGkemulticloudBetaAwsNodePoolResponse, error) {
	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAwsNodePool(ctx, ProtoToAwsNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkemulticloudBetaAwsNodePool
	for _, r := range resources.Items {
		rp := AwsNodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListGkemulticloudBetaAwsNodePoolResponse{Items: protos}, nil
}

func createConfigAwsNodePool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
