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

// Server implements the gRPC interface for AwsNodePool.
type AwsNodePoolServer struct{}

// ProtoToAwsNodePoolConfigRootVolumeVolumeTypeEnum converts a AwsNodePoolConfigRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum(e gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum) *gkemulticloud.AwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := gkemulticloud.AwsNodePoolConfigRootVolumeVolumeTypeEnum(n[len("GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolConfigTaintsEffectEnum converts a AwsNodePoolConfigTaintsEffectEnum enum from its proto representation.
func ProtoToGkemulticloudAwsNodePoolConfigTaintsEffectEnum(e gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaintsEffectEnum) *gkemulticloud.AwsNodePoolConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := gkemulticloud.AwsNodePoolConfigTaintsEffectEnum(n[len("GkemulticloudAwsNodePoolConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolStateEnum converts a AwsNodePoolStateEnum enum from its proto representation.
func ProtoToGkemulticloudAwsNodePoolStateEnum(e gkemulticloudpb.GkemulticloudAwsNodePoolStateEnum) *gkemulticloud.AwsNodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkemulticloudpb.GkemulticloudAwsNodePoolStateEnum_name[int32(e)]; ok {
		e := gkemulticloud.AwsNodePoolStateEnum(n[len("GkemulticloudAwsNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsNodePoolConfig converts a AwsNodePoolConfig resource from its proto representation.
func ProtoToGkemulticloudAwsNodePoolConfig(p *gkemulticloudpb.GkemulticloudAwsNodePoolConfig) *gkemulticloud.AwsNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsNodePoolConfig{
		InstanceType:       dcl.StringOrNil(p.InstanceType),
		RootVolume:         ProtoToGkemulticloudAwsNodePoolConfigRootVolume(p.GetRootVolume()),
		IamInstanceProfile: dcl.StringOrNil(p.IamInstanceProfile),
		SshConfig:          ProtoToGkemulticloudAwsNodePoolConfigSshConfig(p.GetSshConfig()),
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToGkemulticloudAwsNodePoolConfigTaints(r))
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToAwsNodePoolConfigRootVolume converts a AwsNodePoolConfigRootVolume resource from its proto representation.
func ProtoToGkemulticloudAwsNodePoolConfigRootVolume(p *gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolume) *gkemulticloud.AwsNodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsNodePoolConfigRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsNodePoolConfigTaints converts a AwsNodePoolConfigTaints resource from its proto representation.
func ProtoToGkemulticloudAwsNodePoolConfigTaints(p *gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaints) *gkemulticloud.AwsNodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsNodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToGkemulticloudAwsNodePoolConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToAwsNodePoolConfigSshConfig converts a AwsNodePoolConfigSshConfig resource from its proto representation.
func ProtoToGkemulticloudAwsNodePoolConfigSshConfig(p *gkemulticloudpb.GkemulticloudAwsNodePoolConfigSshConfig) *gkemulticloud.AwsNodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToAwsNodePoolAutoscaling converts a AwsNodePoolAutoscaling resource from its proto representation.
func ProtoToGkemulticloudAwsNodePoolAutoscaling(p *gkemulticloudpb.GkemulticloudAwsNodePoolAutoscaling) *gkemulticloud.AwsNodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsNodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount: dcl.Int64OrNil(p.MaxNodeCount),
	}
	return obj
}

// ProtoToAwsNodePoolMaxPodsConstraint converts a AwsNodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToGkemulticloudAwsNodePoolMaxPodsConstraint(p *gkemulticloudpb.GkemulticloudAwsNodePoolMaxPodsConstraint) *gkemulticloud.AwsNodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToAwsNodePool converts a AwsNodePool resource from its proto representation.
func ProtoToAwsNodePool(p *gkemulticloudpb.GkemulticloudAwsNodePool) *gkemulticloud.AwsNodePool {
	obj := &gkemulticloud.AwsNodePool{
		Name:              dcl.StringOrNil(p.Name),
		Version:           dcl.StringOrNil(p.Version),
		Config:            ProtoToGkemulticloudAwsNodePoolConfig(p.GetConfig()),
		Autoscaling:       ProtoToGkemulticloudAwsNodePoolAutoscaling(p.GetAutoscaling()),
		SubnetId:          dcl.StringOrNil(p.SubnetId),
		State:             ProtoToGkemulticloudAwsNodePoolStateEnum(p.GetState()),
		Uid:               dcl.StringOrNil(p.Uid),
		Reconciling:       dcl.Bool(p.Reconciling),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Etag:              dcl.StringOrNil(p.Etag),
		MaxPodsConstraint: ProtoToGkemulticloudAwsNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
		AwsCluster:        dcl.StringOrNil(p.AwsCluster),
	}
	return obj
}

// AwsNodePoolConfigRootVolumeVolumeTypeEnumToProto converts a AwsNodePoolConfigRootVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnumToProto(e *gkemulticloud.AwsNodePoolConfigRootVolumeVolumeTypeEnum) gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == nil {
		return gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum_value["AwsNodePoolConfigRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum(v)
	}
	return gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnum(0)
}

// AwsNodePoolConfigTaintsEffectEnumToProto converts a AwsNodePoolConfigTaintsEffectEnum enum to its proto representation.
func GkemulticloudAwsNodePoolConfigTaintsEffectEnumToProto(e *gkemulticloud.AwsNodePoolConfigTaintsEffectEnum) gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaintsEffectEnum {
	if e == nil {
		return gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaintsEffectEnum(0)
	}
	if v, ok := gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaintsEffectEnum_value["AwsNodePoolConfigTaintsEffectEnum"+string(*e)]; ok {
		return gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaintsEffectEnum(v)
	}
	return gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaintsEffectEnum(0)
}

// AwsNodePoolStateEnumToProto converts a AwsNodePoolStateEnum enum to its proto representation.
func GkemulticloudAwsNodePoolStateEnumToProto(e *gkemulticloud.AwsNodePoolStateEnum) gkemulticloudpb.GkemulticloudAwsNodePoolStateEnum {
	if e == nil {
		return gkemulticloudpb.GkemulticloudAwsNodePoolStateEnum(0)
	}
	if v, ok := gkemulticloudpb.GkemulticloudAwsNodePoolStateEnum_value["AwsNodePoolStateEnum"+string(*e)]; ok {
		return gkemulticloudpb.GkemulticloudAwsNodePoolStateEnum(v)
	}
	return gkemulticloudpb.GkemulticloudAwsNodePoolStateEnum(0)
}

// AwsNodePoolConfigToProto converts a AwsNodePoolConfig resource to its proto representation.
func GkemulticloudAwsNodePoolConfigToProto(o *gkemulticloud.AwsNodePoolConfig) *gkemulticloudpb.GkemulticloudAwsNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsNodePoolConfig{
		InstanceType:       dcl.ValueOrEmptyString(o.InstanceType),
		RootVolume:         GkemulticloudAwsNodePoolConfigRootVolumeToProto(o.RootVolume),
		IamInstanceProfile: dcl.ValueOrEmptyString(o.IamInstanceProfile),
		SshConfig:          GkemulticloudAwsNodePoolConfigSshConfigToProto(o.SshConfig),
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, GkemulticloudAwsNodePoolConfigTaintsToProto(&r))
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
func GkemulticloudAwsNodePoolConfigRootVolumeToProto(o *gkemulticloud.AwsNodePoolConfigRootVolume) *gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsNodePoolConfigRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudAwsNodePoolConfigRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsNodePoolConfigTaintsToProto converts a AwsNodePoolConfigTaints resource to its proto representation.
func GkemulticloudAwsNodePoolConfigTaintsToProto(o *gkemulticloud.AwsNodePoolConfigTaints) *gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsNodePoolConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: GkemulticloudAwsNodePoolConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// AwsNodePoolConfigSshConfigToProto converts a AwsNodePoolConfigSshConfig resource to its proto representation.
func GkemulticloudAwsNodePoolConfigSshConfigToProto(o *gkemulticloud.AwsNodePoolConfigSshConfig) *gkemulticloudpb.GkemulticloudAwsNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// AwsNodePoolAutoscalingToProto converts a AwsNodePoolAutoscaling resource to its proto representation.
func GkemulticloudAwsNodePoolAutoscalingToProto(o *gkemulticloud.AwsNodePoolAutoscaling) *gkemulticloudpb.GkemulticloudAwsNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsNodePoolAutoscaling{
		MinNodeCount: dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount: dcl.ValueOrEmptyInt64(o.MaxNodeCount),
	}
	return p
}

// AwsNodePoolMaxPodsConstraintToProto converts a AwsNodePoolMaxPodsConstraint resource to its proto representation.
func GkemulticloudAwsNodePoolMaxPodsConstraintToProto(o *gkemulticloud.AwsNodePoolMaxPodsConstraint) *gkemulticloudpb.GkemulticloudAwsNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// AwsNodePoolToProto converts a AwsNodePool resource to its proto representation.
func AwsNodePoolToProto(resource *gkemulticloud.AwsNodePool) *gkemulticloudpb.GkemulticloudAwsNodePool {
	p := &gkemulticloudpb.GkemulticloudAwsNodePool{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Version:           dcl.ValueOrEmptyString(resource.Version),
		Config:            GkemulticloudAwsNodePoolConfigToProto(resource.Config),
		Autoscaling:       GkemulticloudAwsNodePoolAutoscalingToProto(resource.Autoscaling),
		SubnetId:          dcl.ValueOrEmptyString(resource.SubnetId),
		State:             GkemulticloudAwsNodePoolStateEnumToProto(resource.State),
		Uid:               dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:       dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:        dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:        dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:              dcl.ValueOrEmptyString(resource.Etag),
		MaxPodsConstraint: GkemulticloudAwsNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
		AwsCluster:        dcl.ValueOrEmptyString(resource.AwsCluster),
	}

	return p
}

// ApplyAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Apply() method.
func (s *AwsNodePoolServer) applyAwsNodePool(ctx context.Context, c *gkemulticloud.Client, request *gkemulticloudpb.ApplyGkemulticloudAwsNodePoolRequest) (*gkemulticloudpb.GkemulticloudAwsNodePool, error) {
	p := ProtoToAwsNodePool(request.GetResource())
	res, err := c.ApplyAwsNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AwsNodePoolToProto(res)
	return r, nil
}

// ApplyAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Apply() method.
func (s *AwsNodePoolServer) ApplyGkemulticloudAwsNodePool(ctx context.Context, request *gkemulticloudpb.ApplyGkemulticloudAwsNodePoolRequest) (*gkemulticloudpb.GkemulticloudAwsNodePool, error) {
	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAwsNodePool(ctx, cl, request)
}

// DeleteAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePool Delete() method.
func (s *AwsNodePoolServer) DeleteGkemulticloudAwsNodePool(ctx context.Context, request *gkemulticloudpb.DeleteGkemulticloudAwsNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAwsNodePool(ctx, ProtoToAwsNodePool(request.GetResource()))

}

// ListGkemulticloudAwsNodePool handles the gRPC request by passing it to the underlying AwsNodePoolList() method.
func (s *AwsNodePoolServer) ListGkemulticloudAwsNodePool(ctx context.Context, request *gkemulticloudpb.ListGkemulticloudAwsNodePoolRequest) (*gkemulticloudpb.ListGkemulticloudAwsNodePoolResponse, error) {
	cl, err := createConfigAwsNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAwsNodePool(ctx, ProtoToAwsNodePool(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*gkemulticloudpb.GkemulticloudAwsNodePool
	for _, r := range resources.Items {
		rp := AwsNodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &gkemulticloudpb.ListGkemulticloudAwsNodePoolResponse{Items: protos}, nil
}

func createConfigAwsNodePool(ctx context.Context, service_account_file string) (*gkemulticloud.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return gkemulticloud.NewClient(conf), nil
}
