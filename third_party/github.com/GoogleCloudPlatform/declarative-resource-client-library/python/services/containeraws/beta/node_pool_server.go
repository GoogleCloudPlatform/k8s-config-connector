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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeraws/beta/containeraws_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/beta"
)

// NodePoolServer implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolConfigRootVolumeVolumeTypeEnum converts a NodePoolConfigRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum(e betapb.ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum) *beta.NodePoolConfigRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := beta.NodePoolConfigRootVolumeVolumeTypeEnum(n[len("ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfigTaintsEffectEnum converts a NodePoolConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigTaintsEffectEnum(e betapb.ContainerawsBetaNodePoolConfigTaintsEffectEnum) *beta.NodePoolConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaNodePoolConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := beta.NodePoolConfigTaintsEffectEnum(n[len("ContainerawsBetaNodePoolConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfigInstancePlacementTenancyEnum converts a NodePoolConfigInstancePlacementTenancyEnum enum from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum(e betapb.ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum) *beta.NodePoolConfigInstancePlacementTenancyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum_name[int32(e)]; ok {
		e := beta.NodePoolConfigInstancePlacementTenancyEnum(n[len("ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolStateEnum converts a NodePoolStateEnum enum from its proto representation.
func ProtoToContainerawsBetaNodePoolStateEnum(e betapb.ContainerawsBetaNodePoolStateEnum) *beta.NodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaNodePoolStateEnum_name[int32(e)]; ok {
		e := beta.NodePoolStateEnum(n[len("ContainerawsBetaNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfig(p *betapb.ContainerawsBetaNodePoolConfig) *beta.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfig{
		InstanceType:                 dcl.StringOrNil(p.GetInstanceType()),
		RootVolume:                   ProtoToContainerawsBetaNodePoolConfigRootVolume(p.GetRootVolume()),
		IamInstanceProfile:           dcl.StringOrNil(p.GetIamInstanceProfile()),
		ConfigEncryption:             ProtoToContainerawsBetaNodePoolConfigConfigEncryption(p.GetConfigEncryption()),
		SshConfig:                    ProtoToContainerawsBetaNodePoolConfigSshConfig(p.GetSshConfig()),
		SpotConfig:                   ProtoToContainerawsBetaNodePoolConfigSpotConfig(p.GetSpotConfig()),
		ProxyConfig:                  ProtoToContainerawsBetaNodePoolConfigProxyConfig(p.GetProxyConfig()),
		InstancePlacement:            ProtoToContainerawsBetaNodePoolConfigInstancePlacement(p.GetInstancePlacement()),
		ImageType:                    dcl.StringOrNil(p.GetImageType()),
		AutoscalingMetricsCollection: ProtoToContainerawsBetaNodePoolConfigAutoscalingMetricsCollection(p.GetAutoscalingMetricsCollection()),
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerawsBetaNodePoolConfigTaints(r))
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToNodePoolConfigRootVolume converts a NodePoolConfigRootVolume object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigRootVolume(p *betapb.ContainerawsBetaNodePoolConfigRootVolume) *beta.NodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigRootVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToNodePoolConfigTaints converts a NodePoolConfigTaints object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigTaints(p *betapb.ContainerawsBetaNodePoolConfigTaints) *beta.NodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.GetKey()),
		Value:  dcl.StringOrNil(p.GetValue()),
		Effect: ProtoToContainerawsBetaNodePoolConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToNodePoolConfigConfigEncryption converts a NodePoolConfigConfigEncryption object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigConfigEncryption(p *betapb.ContainerawsBetaNodePoolConfigConfigEncryption) *beta.NodePoolConfigConfigEncryption {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigConfigEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToNodePoolConfigSshConfig converts a NodePoolConfigSshConfig object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigSshConfig(p *betapb.ContainerawsBetaNodePoolConfigSshConfig) *beta.NodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.GetEc2KeyPair()),
	}
	return obj
}

// ProtoToNodePoolConfigSpotConfig converts a NodePoolConfigSpotConfig object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigSpotConfig(p *betapb.ContainerawsBetaNodePoolConfigSpotConfig) *beta.NodePoolConfigSpotConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigSpotConfig{}
	for _, r := range p.GetInstanceTypes() {
		obj.InstanceTypes = append(obj.InstanceTypes, r)
	}
	return obj
}

// ProtoToNodePoolConfigProxyConfig converts a NodePoolConfigProxyConfig object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigProxyConfig(p *betapb.ContainerawsBetaNodePoolConfigProxyConfig) *beta.NodePoolConfigProxyConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigProxyConfig{
		SecretArn:     dcl.StringOrNil(p.GetSecretArn()),
		SecretVersion: dcl.StringOrNil(p.GetSecretVersion()),
	}
	return obj
}

// ProtoToNodePoolConfigInstancePlacement converts a NodePoolConfigInstancePlacement object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigInstancePlacement(p *betapb.ContainerawsBetaNodePoolConfigInstancePlacement) *beta.NodePoolConfigInstancePlacement {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigInstancePlacement{
		Tenancy: ProtoToContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum(p.GetTenancy()),
	}
	return obj
}

// ProtoToNodePoolConfigAutoscalingMetricsCollection converts a NodePoolConfigAutoscalingMetricsCollection object from its proto representation.
func ProtoToContainerawsBetaNodePoolConfigAutoscalingMetricsCollection(p *betapb.ContainerawsBetaNodePoolConfigAutoscalingMetricsCollection) *beta.NodePoolConfigAutoscalingMetricsCollection {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigAutoscalingMetricsCollection{
		Granularity: dcl.StringOrNil(p.GetGranularity()),
	}
	for _, r := range p.GetMetrics() {
		obj.Metrics = append(obj.Metrics, r)
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling object from its proto representation.
func ProtoToContainerawsBetaNodePoolAutoscaling(p *betapb.ContainerawsBetaNodePoolAutoscaling) *beta.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.GetMinNodeCount()),
		MaxNodeCount: dcl.Int64OrNil(p.GetMaxNodeCount()),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint object from its proto representation.
func ProtoToContainerawsBetaNodePoolMaxPodsConstraint(p *betapb.ContainerawsBetaNodePoolMaxPodsConstraint) *beta.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.GetMaxPodsPerNode()),
	}
	return obj
}

// ProtoToNodePoolManagement converts a NodePoolManagement object from its proto representation.
func ProtoToContainerawsBetaNodePoolManagement(p *betapb.ContainerawsBetaNodePoolManagement) *beta.NodePoolManagement {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolManagement{
		AutoRepair: dcl.Bool(p.GetAutoRepair()),
	}
	return obj
}

// ProtoToNodePoolUpdateSettings converts a NodePoolUpdateSettings object from its proto representation.
func ProtoToContainerawsBetaNodePoolUpdateSettings(p *betapb.ContainerawsBetaNodePoolUpdateSettings) *beta.NodePoolUpdateSettings {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolUpdateSettings{
		SurgeSettings: ProtoToContainerawsBetaNodePoolUpdateSettingsSurgeSettings(p.GetSurgeSettings()),
	}
	return obj
}

// ProtoToNodePoolUpdateSettingsSurgeSettings converts a NodePoolUpdateSettingsSurgeSettings object from its proto representation.
func ProtoToContainerawsBetaNodePoolUpdateSettingsSurgeSettings(p *betapb.ContainerawsBetaNodePoolUpdateSettingsSurgeSettings) *beta.NodePoolUpdateSettingsSurgeSettings {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolUpdateSettingsSurgeSettings{
		MaxSurge:       dcl.Int64OrNil(p.GetMaxSurge()),
		MaxUnavailable: dcl.Int64OrNil(p.GetMaxUnavailable()),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *betapb.ContainerawsBetaNodePool) *beta.NodePool {
	obj := &beta.NodePool{
		Name:              dcl.StringOrNil(p.GetName()),
		Version:           dcl.StringOrNil(p.GetVersion()),
		Config:            ProtoToContainerawsBetaNodePoolConfig(p.GetConfig()),
		Autoscaling:       ProtoToContainerawsBetaNodePoolAutoscaling(p.GetAutoscaling()),
		SubnetId:          dcl.StringOrNil(p.GetSubnetId()),
		State:             ProtoToContainerawsBetaNodePoolStateEnum(p.GetState()),
		Uid:               dcl.StringOrNil(p.GetUid()),
		Reconciling:       dcl.Bool(p.GetReconciling()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Etag:              dcl.StringOrNil(p.GetEtag()),
		MaxPodsConstraint: ProtoToContainerawsBetaNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Management:        ProtoToContainerawsBetaNodePoolManagement(p.GetManagement()),
		UpdateSettings:    ProtoToContainerawsBetaNodePoolUpdateSettings(p.GetUpdateSettings()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
		Cluster:           dcl.StringOrNil(p.GetCluster()),
	}
	return obj
}

// NodePoolConfigRootVolumeVolumeTypeEnumToProto converts a NodePoolConfigRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnumToProto(e *beta.NodePoolConfigRootVolumeVolumeTypeEnum) betapb.ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == nil {
		return betapb.ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum_value["NodePoolConfigRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum(v)
	}
	return betapb.ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnum(0)
}

// NodePoolConfigTaintsEffectEnumToProto converts a NodePoolConfigTaintsEffectEnum enum to its proto representation.
func ContainerawsBetaNodePoolConfigTaintsEffectEnumToProto(e *beta.NodePoolConfigTaintsEffectEnum) betapb.ContainerawsBetaNodePoolConfigTaintsEffectEnum {
	if e == nil {
		return betapb.ContainerawsBetaNodePoolConfigTaintsEffectEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaNodePoolConfigTaintsEffectEnum_value["NodePoolConfigTaintsEffectEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaNodePoolConfigTaintsEffectEnum(v)
	}
	return betapb.ContainerawsBetaNodePoolConfigTaintsEffectEnum(0)
}

// NodePoolConfigInstancePlacementTenancyEnumToProto converts a NodePoolConfigInstancePlacementTenancyEnum enum to its proto representation.
func ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnumToProto(e *beta.NodePoolConfigInstancePlacementTenancyEnum) betapb.ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum {
	if e == nil {
		return betapb.ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum_value["NodePoolConfigInstancePlacementTenancyEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum(v)
	}
	return betapb.ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnum(0)
}

// NodePoolStateEnumToProto converts a NodePoolStateEnum enum to its proto representation.
func ContainerawsBetaNodePoolStateEnumToProto(e *beta.NodePoolStateEnum) betapb.ContainerawsBetaNodePoolStateEnum {
	if e == nil {
		return betapb.ContainerawsBetaNodePoolStateEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaNodePoolStateEnum_value["NodePoolStateEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaNodePoolStateEnum(v)
	}
	return betapb.ContainerawsBetaNodePoolStateEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig object to its proto representation.
func ContainerawsBetaNodePoolConfigToProto(o *beta.NodePoolConfig) *betapb.ContainerawsBetaNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfig{}
	p.SetInstanceType(dcl.ValueOrEmptyString(o.InstanceType))
	p.SetRootVolume(ContainerawsBetaNodePoolConfigRootVolumeToProto(o.RootVolume))
	p.SetIamInstanceProfile(dcl.ValueOrEmptyString(o.IamInstanceProfile))
	p.SetConfigEncryption(ContainerawsBetaNodePoolConfigConfigEncryptionToProto(o.ConfigEncryption))
	p.SetSshConfig(ContainerawsBetaNodePoolConfigSshConfigToProto(o.SshConfig))
	p.SetSpotConfig(ContainerawsBetaNodePoolConfigSpotConfigToProto(o.SpotConfig))
	p.SetProxyConfig(ContainerawsBetaNodePoolConfigProxyConfigToProto(o.ProxyConfig))
	p.SetInstancePlacement(ContainerawsBetaNodePoolConfigInstancePlacementToProto(o.InstancePlacement))
	p.SetImageType(dcl.ValueOrEmptyString(o.ImageType))
	p.SetAutoscalingMetricsCollection(ContainerawsBetaNodePoolConfigAutoscalingMetricsCollectionToProto(o.AutoscalingMetricsCollection))
	sTaints := make([]*betapb.ContainerawsBetaNodePoolConfigTaints, len(o.Taints))
	for i, r := range o.Taints {
		sTaints[i] = ContainerawsBetaNodePoolConfigTaintsToProto(&r)
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
func ContainerawsBetaNodePoolConfigRootVolumeToProto(o *beta.NodePoolConfigRootVolume) *betapb.ContainerawsBetaNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfigRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsBetaNodePoolConfigRootVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// NodePoolConfigTaintsToProto converts a NodePoolConfigTaints object to its proto representation.
func ContainerawsBetaNodePoolConfigTaintsToProto(o *beta.NodePoolConfigTaints) *betapb.ContainerawsBetaNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfigTaints{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetEffect(ContainerawsBetaNodePoolConfigTaintsEffectEnumToProto(o.Effect))
	return p
}

// NodePoolConfigConfigEncryptionToProto converts a NodePoolConfigConfigEncryption object to its proto representation.
func ContainerawsBetaNodePoolConfigConfigEncryptionToProto(o *beta.NodePoolConfigConfigEncryption) *betapb.ContainerawsBetaNodePoolConfigConfigEncryption {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfigConfigEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// NodePoolConfigSshConfigToProto converts a NodePoolConfigSshConfig object to its proto representation.
func ContainerawsBetaNodePoolConfigSshConfigToProto(o *beta.NodePoolConfigSshConfig) *betapb.ContainerawsBetaNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfigSshConfig{}
	p.SetEc2KeyPair(dcl.ValueOrEmptyString(o.Ec2KeyPair))
	return p
}

// NodePoolConfigSpotConfigToProto converts a NodePoolConfigSpotConfig object to its proto representation.
func ContainerawsBetaNodePoolConfigSpotConfigToProto(o *beta.NodePoolConfigSpotConfig) *betapb.ContainerawsBetaNodePoolConfigSpotConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfigSpotConfig{}
	sInstanceTypes := make([]string, len(o.InstanceTypes))
	for i, r := range o.InstanceTypes {
		sInstanceTypes[i] = r
	}
	p.SetInstanceTypes(sInstanceTypes)
	return p
}

// NodePoolConfigProxyConfigToProto converts a NodePoolConfigProxyConfig object to its proto representation.
func ContainerawsBetaNodePoolConfigProxyConfigToProto(o *beta.NodePoolConfigProxyConfig) *betapb.ContainerawsBetaNodePoolConfigProxyConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfigProxyConfig{}
	p.SetSecretArn(dcl.ValueOrEmptyString(o.SecretArn))
	p.SetSecretVersion(dcl.ValueOrEmptyString(o.SecretVersion))
	return p
}

// NodePoolConfigInstancePlacementToProto converts a NodePoolConfigInstancePlacement object to its proto representation.
func ContainerawsBetaNodePoolConfigInstancePlacementToProto(o *beta.NodePoolConfigInstancePlacement) *betapb.ContainerawsBetaNodePoolConfigInstancePlacement {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfigInstancePlacement{}
	p.SetTenancy(ContainerawsBetaNodePoolConfigInstancePlacementTenancyEnumToProto(o.Tenancy))
	return p
}

// NodePoolConfigAutoscalingMetricsCollectionToProto converts a NodePoolConfigAutoscalingMetricsCollection object to its proto representation.
func ContainerawsBetaNodePoolConfigAutoscalingMetricsCollectionToProto(o *beta.NodePoolConfigAutoscalingMetricsCollection) *betapb.ContainerawsBetaNodePoolConfigAutoscalingMetricsCollection {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolConfigAutoscalingMetricsCollection{}
	p.SetGranularity(dcl.ValueOrEmptyString(o.Granularity))
	sMetrics := make([]string, len(o.Metrics))
	for i, r := range o.Metrics {
		sMetrics[i] = r
	}
	p.SetMetrics(sMetrics)
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling object to its proto representation.
func ContainerawsBetaNodePoolAutoscalingToProto(o *beta.NodePoolAutoscaling) *betapb.ContainerawsBetaNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolAutoscaling{}
	p.SetMinNodeCount(dcl.ValueOrEmptyInt64(o.MinNodeCount))
	p.SetMaxNodeCount(dcl.ValueOrEmptyInt64(o.MaxNodeCount))
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint object to its proto representation.
func ContainerawsBetaNodePoolMaxPodsConstraintToProto(o *beta.NodePoolMaxPodsConstraint) *betapb.ContainerawsBetaNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolMaxPodsConstraint{}
	p.SetMaxPodsPerNode(dcl.ValueOrEmptyInt64(o.MaxPodsPerNode))
	return p
}

// NodePoolManagementToProto converts a NodePoolManagement object to its proto representation.
func ContainerawsBetaNodePoolManagementToProto(o *beta.NodePoolManagement) *betapb.ContainerawsBetaNodePoolManagement {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolManagement{}
	p.SetAutoRepair(dcl.ValueOrEmptyBool(o.AutoRepair))
	return p
}

// NodePoolUpdateSettingsToProto converts a NodePoolUpdateSettings object to its proto representation.
func ContainerawsBetaNodePoolUpdateSettingsToProto(o *beta.NodePoolUpdateSettings) *betapb.ContainerawsBetaNodePoolUpdateSettings {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolUpdateSettings{}
	p.SetSurgeSettings(ContainerawsBetaNodePoolUpdateSettingsSurgeSettingsToProto(o.SurgeSettings))
	return p
}

// NodePoolUpdateSettingsSurgeSettingsToProto converts a NodePoolUpdateSettingsSurgeSettings object to its proto representation.
func ContainerawsBetaNodePoolUpdateSettingsSurgeSettingsToProto(o *beta.NodePoolUpdateSettingsSurgeSettings) *betapb.ContainerawsBetaNodePoolUpdateSettingsSurgeSettings {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaNodePoolUpdateSettingsSurgeSettings{}
	p.SetMaxSurge(dcl.ValueOrEmptyInt64(o.MaxSurge))
	p.SetMaxUnavailable(dcl.ValueOrEmptyInt64(o.MaxUnavailable))
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *beta.NodePool) *betapb.ContainerawsBetaNodePool {
	p := &betapb.ContainerawsBetaNodePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyString(resource.Version))
	p.SetConfig(ContainerawsBetaNodePoolConfigToProto(resource.Config))
	p.SetAutoscaling(ContainerawsBetaNodePoolAutoscalingToProto(resource.Autoscaling))
	p.SetSubnetId(dcl.ValueOrEmptyString(resource.SubnetId))
	p.SetState(ContainerawsBetaNodePoolStateEnumToProto(resource.State))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetMaxPodsConstraint(ContainerawsBetaNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint))
	p.SetManagement(ContainerawsBetaNodePoolManagementToProto(resource.Management))
	p.SetUpdateSettings(ContainerawsBetaNodePoolUpdateSettingsToProto(resource.UpdateSettings))
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
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *beta.Client, request *betapb.ApplyContainerawsBetaNodePoolRequest) (*betapb.ContainerawsBetaNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// applyContainerawsBetaNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerawsBetaNodePool(ctx context.Context, request *betapb.ApplyContainerawsBetaNodePoolRequest) (*betapb.ContainerawsBetaNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerawsBetaNodePool(ctx context.Context, request *betapb.DeleteContainerawsBetaNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerawsBetaNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerawsBetaNodePool(ctx context.Context, request *betapb.ListContainerawsBetaNodePoolRequest) (*betapb.ListContainerawsBetaNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.GetProject(), request.GetLocation(), request.GetCluster())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContainerawsBetaNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListContainerawsBetaNodePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
