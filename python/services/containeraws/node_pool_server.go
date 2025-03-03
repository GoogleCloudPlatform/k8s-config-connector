// Copyright 2024 Google LLC. All Rights Reserved.
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
	containerawspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeraws/containeraws_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws"
)

// NodePoolServer implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolConfigRootVolumeVolumeTypeEnum converts a NodePoolConfigRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsNodePoolConfigRootVolumeVolumeTypeEnum(e containerawspb.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum) *containeraws.NodePoolConfigRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerawspb.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := containeraws.NodePoolConfigRootVolumeVolumeTypeEnum(n[len("ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfigTaintsEffectEnum converts a NodePoolConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerawsNodePoolConfigTaintsEffectEnum(e containerawspb.ContainerawsNodePoolConfigTaintsEffectEnum) *containeraws.NodePoolConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerawspb.ContainerawsNodePoolConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := containeraws.NodePoolConfigTaintsEffectEnum(n[len("ContainerawsNodePoolConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolStateEnum converts a NodePoolStateEnum enum from its proto representation.
func ProtoToContainerawsNodePoolStateEnum(e containerawspb.ContainerawsNodePoolStateEnum) *containeraws.NodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerawspb.ContainerawsNodePoolStateEnum_name[int32(e)]; ok {
		e := containeraws.NodePoolStateEnum(n[len("ContainerawsNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig object from its proto representation.
func ProtoToContainerawsNodePoolConfig(p *containerawspb.ContainerawsNodePoolConfig) *containeraws.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolConfig{
		InstanceType:                 dcl.StringOrNil(p.GetInstanceType()),
		RootVolume:                   ProtoToContainerawsNodePoolConfigRootVolume(p.GetRootVolume()),
		IamInstanceProfile:           dcl.StringOrNil(p.GetIamInstanceProfile()),
		ConfigEncryption:             ProtoToContainerawsNodePoolConfigConfigEncryption(p.GetConfigEncryption()),
		SshConfig:                    ProtoToContainerawsNodePoolConfigSshConfig(p.GetSshConfig()),
		ProxyConfig:                  ProtoToContainerawsNodePoolConfigProxyConfig(p.GetProxyConfig()),
		AutoscalingMetricsCollection: ProtoToContainerawsNodePoolConfigAutoscalingMetricsCollection(p.GetAutoscalingMetricsCollection()),
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerawsNodePoolConfigTaints(r))
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToNodePoolConfigRootVolume converts a NodePoolConfigRootVolume object from its proto representation.
func ProtoToContainerawsNodePoolConfigRootVolume(p *containerawspb.ContainerawsNodePoolConfigRootVolume) *containeraws.NodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolConfigRootVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsNodePoolConfigRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToNodePoolConfigTaints converts a NodePoolConfigTaints object from its proto representation.
func ProtoToContainerawsNodePoolConfigTaints(p *containerawspb.ContainerawsNodePoolConfigTaints) *containeraws.NodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.GetKey()),
		Value:  dcl.StringOrNil(p.GetValue()),
		Effect: ProtoToContainerawsNodePoolConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToNodePoolConfigConfigEncryption converts a NodePoolConfigConfigEncryption object from its proto representation.
func ProtoToContainerawsNodePoolConfigConfigEncryption(p *containerawspb.ContainerawsNodePoolConfigConfigEncryption) *containeraws.NodePoolConfigConfigEncryption {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolConfigConfigEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToNodePoolConfigSshConfig converts a NodePoolConfigSshConfig object from its proto representation.
func ProtoToContainerawsNodePoolConfigSshConfig(p *containerawspb.ContainerawsNodePoolConfigSshConfig) *containeraws.NodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolConfigSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.GetEc2KeyPair()),
	}
	return obj
}

// ProtoToNodePoolConfigProxyConfig converts a NodePoolConfigProxyConfig object from its proto representation.
func ProtoToContainerawsNodePoolConfigProxyConfig(p *containerawspb.ContainerawsNodePoolConfigProxyConfig) *containeraws.NodePoolConfigProxyConfig {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolConfigProxyConfig{
		SecretArn:     dcl.StringOrNil(p.GetSecretArn()),
		SecretVersion: dcl.StringOrNil(p.GetSecretVersion()),
	}
	return obj
}

// ProtoToNodePoolConfigAutoscalingMetricsCollection converts a NodePoolConfigAutoscalingMetricsCollection object from its proto representation.
func ProtoToContainerawsNodePoolConfigAutoscalingMetricsCollection(p *containerawspb.ContainerawsNodePoolConfigAutoscalingMetricsCollection) *containeraws.NodePoolConfigAutoscalingMetricsCollection {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolConfigAutoscalingMetricsCollection{
		Granularity: dcl.StringOrNil(p.GetGranularity()),
	}
	for _, r := range p.GetMetrics() {
		obj.Metrics = append(obj.Metrics, r)
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling object from its proto representation.
func ProtoToContainerawsNodePoolAutoscaling(p *containerawspb.ContainerawsNodePoolAutoscaling) *containeraws.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.GetMinNodeCount()),
		MaxNodeCount: dcl.Int64OrNil(p.GetMaxNodeCount()),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint object from its proto representation.
func ProtoToContainerawsNodePoolMaxPodsConstraint(p *containerawspb.ContainerawsNodePoolMaxPodsConstraint) *containeraws.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.GetMaxPodsPerNode()),
	}
	return obj
}

// ProtoToNodePoolManagement converts a NodePoolManagement object from its proto representation.
func ProtoToContainerawsNodePoolManagement(p *containerawspb.ContainerawsNodePoolManagement) *containeraws.NodePoolManagement {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolManagement{
		AutoRepair: dcl.Bool(p.GetAutoRepair()),
	}
	return obj
}

// ProtoToNodePoolUpdateSettings converts a NodePoolUpdateSettings object from its proto representation.
func ProtoToContainerawsNodePoolUpdateSettings(p *containerawspb.ContainerawsNodePoolUpdateSettings) *containeraws.NodePoolUpdateSettings {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolUpdateSettings{
		SurgeSettings: ProtoToContainerawsNodePoolUpdateSettingsSurgeSettings(p.GetSurgeSettings()),
	}
	return obj
}

// ProtoToNodePoolUpdateSettingsSurgeSettings converts a NodePoolUpdateSettingsSurgeSettings object from its proto representation.
func ProtoToContainerawsNodePoolUpdateSettingsSurgeSettings(p *containerawspb.ContainerawsNodePoolUpdateSettingsSurgeSettings) *containeraws.NodePoolUpdateSettingsSurgeSettings {
	if p == nil {
		return nil
	}
	obj := &containeraws.NodePoolUpdateSettingsSurgeSettings{
		MaxSurge:       dcl.Int64OrNil(p.GetMaxSurge()),
		MaxUnavailable: dcl.Int64OrNil(p.GetMaxUnavailable()),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *containerawspb.ContainerawsNodePool) *containeraws.NodePool {
	obj := &containeraws.NodePool{
		Name:              dcl.StringOrNil(p.GetName()),
		Version:           dcl.StringOrNil(p.GetVersion()),
		Config:            ProtoToContainerawsNodePoolConfig(p.GetConfig()),
		Autoscaling:       ProtoToContainerawsNodePoolAutoscaling(p.GetAutoscaling()),
		SubnetId:          dcl.StringOrNil(p.GetSubnetId()),
		State:             ProtoToContainerawsNodePoolStateEnum(p.GetState()),
		Uid:               dcl.StringOrNil(p.GetUid()),
		Reconciling:       dcl.Bool(p.GetReconciling()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Etag:              dcl.StringOrNil(p.GetEtag()),
		MaxPodsConstraint: ProtoToContainerawsNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Management:        ProtoToContainerawsNodePoolManagement(p.GetManagement()),
		UpdateSettings:    ProtoToContainerawsNodePoolUpdateSettings(p.GetUpdateSettings()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
		Cluster:           dcl.StringOrNil(p.GetCluster()),
	}
	return obj
}

// NodePoolConfigRootVolumeVolumeTypeEnumToProto converts a NodePoolConfigRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsNodePoolConfigRootVolumeVolumeTypeEnumToProto(e *containeraws.NodePoolConfigRootVolumeVolumeTypeEnum) containerawspb.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == nil {
		return containerawspb.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := containerawspb.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum_value["NodePoolConfigRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return containerawspb.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum(v)
	}
	return containerawspb.ContainerawsNodePoolConfigRootVolumeVolumeTypeEnum(0)
}

// NodePoolConfigTaintsEffectEnumToProto converts a NodePoolConfigTaintsEffectEnum enum to its proto representation.
func ContainerawsNodePoolConfigTaintsEffectEnumToProto(e *containeraws.NodePoolConfigTaintsEffectEnum) containerawspb.ContainerawsNodePoolConfigTaintsEffectEnum {
	if e == nil {
		return containerawspb.ContainerawsNodePoolConfigTaintsEffectEnum(0)
	}
	if v, ok := containerawspb.ContainerawsNodePoolConfigTaintsEffectEnum_value["NodePoolConfigTaintsEffectEnum"+string(*e)]; ok {
		return containerawspb.ContainerawsNodePoolConfigTaintsEffectEnum(v)
	}
	return containerawspb.ContainerawsNodePoolConfigTaintsEffectEnum(0)
}

// NodePoolStateEnumToProto converts a NodePoolStateEnum enum to its proto representation.
func ContainerawsNodePoolStateEnumToProto(e *containeraws.NodePoolStateEnum) containerawspb.ContainerawsNodePoolStateEnum {
	if e == nil {
		return containerawspb.ContainerawsNodePoolStateEnum(0)
	}
	if v, ok := containerawspb.ContainerawsNodePoolStateEnum_value["NodePoolStateEnum"+string(*e)]; ok {
		return containerawspb.ContainerawsNodePoolStateEnum(v)
	}
	return containerawspb.ContainerawsNodePoolStateEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig object to its proto representation.
func ContainerawsNodePoolConfigToProto(o *containeraws.NodePoolConfig) *containerawspb.ContainerawsNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolConfig{}
	p.SetInstanceType(dcl.ValueOrEmptyString(o.InstanceType))
	p.SetRootVolume(ContainerawsNodePoolConfigRootVolumeToProto(o.RootVolume))
	p.SetIamInstanceProfile(dcl.ValueOrEmptyString(o.IamInstanceProfile))
	p.SetConfigEncryption(ContainerawsNodePoolConfigConfigEncryptionToProto(o.ConfigEncryption))
	p.SetSshConfig(ContainerawsNodePoolConfigSshConfigToProto(o.SshConfig))
	p.SetProxyConfig(ContainerawsNodePoolConfigProxyConfigToProto(o.ProxyConfig))
	p.SetAutoscalingMetricsCollection(ContainerawsNodePoolConfigAutoscalingMetricsCollectionToProto(o.AutoscalingMetricsCollection))
	sTaints := make([]*containerawspb.ContainerawsNodePoolConfigTaints, len(o.Taints))
	for i, r := range o.Taints {
		sTaints[i] = ContainerawsNodePoolConfigTaintsToProto(&r)
	}
	p.SetTaints(sTaints)
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	mTags := make(map[string]string, len(o.Tags))
	for k, r := range o.Tags {
		mTags[k] = r
	}
	p.SetTags(mTags)
	sSecurityGroupIds := make([]string, len(o.SecurityGroupIds))
	for i, r := range o.SecurityGroupIds {
		sSecurityGroupIds[i] = r
	}
	p.SetSecurityGroupIds(sSecurityGroupIds)
	return p
}

// NodePoolConfigRootVolumeToProto converts a NodePoolConfigRootVolume object to its proto representation.
func ContainerawsNodePoolConfigRootVolumeToProto(o *containeraws.NodePoolConfigRootVolume) *containerawspb.ContainerawsNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolConfigRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsNodePoolConfigRootVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// NodePoolConfigTaintsToProto converts a NodePoolConfigTaints object to its proto representation.
func ContainerawsNodePoolConfigTaintsToProto(o *containeraws.NodePoolConfigTaints) *containerawspb.ContainerawsNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolConfigTaints{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetEffect(ContainerawsNodePoolConfigTaintsEffectEnumToProto(o.Effect))
	return p
}

// NodePoolConfigConfigEncryptionToProto converts a NodePoolConfigConfigEncryption object to its proto representation.
func ContainerawsNodePoolConfigConfigEncryptionToProto(o *containeraws.NodePoolConfigConfigEncryption) *containerawspb.ContainerawsNodePoolConfigConfigEncryption {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolConfigConfigEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// NodePoolConfigSshConfigToProto converts a NodePoolConfigSshConfig object to its proto representation.
func ContainerawsNodePoolConfigSshConfigToProto(o *containeraws.NodePoolConfigSshConfig) *containerawspb.ContainerawsNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolConfigSshConfig{}
	p.SetEc2KeyPair(dcl.ValueOrEmptyString(o.Ec2KeyPair))
	return p
}

// NodePoolConfigProxyConfigToProto converts a NodePoolConfigProxyConfig object to its proto representation.
func ContainerawsNodePoolConfigProxyConfigToProto(o *containeraws.NodePoolConfigProxyConfig) *containerawspb.ContainerawsNodePoolConfigProxyConfig {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolConfigProxyConfig{}
	p.SetSecretArn(dcl.ValueOrEmptyString(o.SecretArn))
	p.SetSecretVersion(dcl.ValueOrEmptyString(o.SecretVersion))
	return p
}

// NodePoolConfigAutoscalingMetricsCollectionToProto converts a NodePoolConfigAutoscalingMetricsCollection object to its proto representation.
func ContainerawsNodePoolConfigAutoscalingMetricsCollectionToProto(o *containeraws.NodePoolConfigAutoscalingMetricsCollection) *containerawspb.ContainerawsNodePoolConfigAutoscalingMetricsCollection {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolConfigAutoscalingMetricsCollection{}
	p.SetGranularity(dcl.ValueOrEmptyString(o.Granularity))
	sMetrics := make([]string, len(o.Metrics))
	for i, r := range o.Metrics {
		sMetrics[i] = r
	}
	p.SetMetrics(sMetrics)
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling object to its proto representation.
func ContainerawsNodePoolAutoscalingToProto(o *containeraws.NodePoolAutoscaling) *containerawspb.ContainerawsNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolAutoscaling{}
	p.SetMinNodeCount(dcl.ValueOrEmptyInt64(o.MinNodeCount))
	p.SetMaxNodeCount(dcl.ValueOrEmptyInt64(o.MaxNodeCount))
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint object to its proto representation.
func ContainerawsNodePoolMaxPodsConstraintToProto(o *containeraws.NodePoolMaxPodsConstraint) *containerawspb.ContainerawsNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolMaxPodsConstraint{}
	p.SetMaxPodsPerNode(dcl.ValueOrEmptyInt64(o.MaxPodsPerNode))
	return p
}

// NodePoolManagementToProto converts a NodePoolManagement object to its proto representation.
func ContainerawsNodePoolManagementToProto(o *containeraws.NodePoolManagement) *containerawspb.ContainerawsNodePoolManagement {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolManagement{}
	p.SetAutoRepair(dcl.ValueOrEmptyBool(o.AutoRepair))
	return p
}

// NodePoolUpdateSettingsToProto converts a NodePoolUpdateSettings object to its proto representation.
func ContainerawsNodePoolUpdateSettingsToProto(o *containeraws.NodePoolUpdateSettings) *containerawspb.ContainerawsNodePoolUpdateSettings {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolUpdateSettings{}
	p.SetSurgeSettings(ContainerawsNodePoolUpdateSettingsSurgeSettingsToProto(o.SurgeSettings))
	return p
}

// NodePoolUpdateSettingsSurgeSettingsToProto converts a NodePoolUpdateSettingsSurgeSettings object to its proto representation.
func ContainerawsNodePoolUpdateSettingsSurgeSettingsToProto(o *containeraws.NodePoolUpdateSettingsSurgeSettings) *containerawspb.ContainerawsNodePoolUpdateSettingsSurgeSettings {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsNodePoolUpdateSettingsSurgeSettings{}
	p.SetMaxSurge(dcl.ValueOrEmptyInt64(o.MaxSurge))
	p.SetMaxUnavailable(dcl.ValueOrEmptyInt64(o.MaxUnavailable))
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *containeraws.NodePool) *containerawspb.ContainerawsNodePool {
	p := &containerawspb.ContainerawsNodePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyString(resource.Version))
	p.SetConfig(ContainerawsNodePoolConfigToProto(resource.Config))
	p.SetAutoscaling(ContainerawsNodePoolAutoscalingToProto(resource.Autoscaling))
	p.SetSubnetId(dcl.ValueOrEmptyString(resource.SubnetId))
	p.SetState(ContainerawsNodePoolStateEnumToProto(resource.State))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetMaxPodsConstraint(ContainerawsNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint))
	p.SetManagement(ContainerawsNodePoolManagementToProto(resource.Management))
	p.SetUpdateSettings(ContainerawsNodePoolUpdateSettingsToProto(resource.UpdateSettings))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetCluster(dcl.ValueOrEmptyString(resource.Cluster))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *containeraws.Client, request *containerawspb.ApplyContainerawsNodePoolRequest) (*containerawspb.ContainerawsNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// applyContainerawsNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerawsNodePool(ctx context.Context, request *containerawspb.ApplyContainerawsNodePoolRequest) (*containerawspb.ContainerawsNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerawsNodePool(ctx context.Context, request *containerawspb.DeleteContainerawsNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerawsNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerawsNodePool(ctx context.Context, request *containerawspb.ListContainerawsNodePoolRequest) (*containerawspb.ListContainerawsNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.GetProject(), request.GetLocation(), request.GetCluster())
	if err != nil {
		return nil, err
	}
	var protos []*containerawspb.ContainerawsNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &containerawspb.ListContainerawsNodePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*containeraws.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return containeraws.NewClient(conf), nil
}
