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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeraws/alpha/containeraws_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/alpha"
)

// NodePoolServer implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolConfigRootVolumeVolumeTypeEnum converts a NodePoolConfigRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum) *alpha.NodePoolConfigRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.NodePoolConfigRootVolumeVolumeTypeEnum(n[len("ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfigTaintsEffectEnum converts a NodePoolConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigTaintsEffectEnum(e alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum) *alpha.NodePoolConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := alpha.NodePoolConfigTaintsEffectEnum(n[len("ContainerawsAlphaNodePoolConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfigInstancePlacementTenancyEnum converts a NodePoolConfigInstancePlacementTenancyEnum enum from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum(e alphapb.ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum) *alpha.NodePoolConfigInstancePlacementTenancyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum_name[int32(e)]; ok {
		e := alpha.NodePoolConfigInstancePlacementTenancyEnum(n[len("ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolStateEnum converts a NodePoolStateEnum enum from its proto representation.
func ProtoToContainerawsAlphaNodePoolStateEnum(e alphapb.ContainerawsAlphaNodePoolStateEnum) *alpha.NodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaNodePoolStateEnum_name[int32(e)]; ok {
		e := alpha.NodePoolStateEnum(n[len("ContainerawsAlphaNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfig(p *alphapb.ContainerawsAlphaNodePoolConfig) *alpha.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfig{
		InstanceType:                 dcl.StringOrNil(p.GetInstanceType()),
		RootVolume:                   ProtoToContainerawsAlphaNodePoolConfigRootVolume(p.GetRootVolume()),
		IamInstanceProfile:           dcl.StringOrNil(p.GetIamInstanceProfile()),
		ConfigEncryption:             ProtoToContainerawsAlphaNodePoolConfigConfigEncryption(p.GetConfigEncryption()),
		SshConfig:                    ProtoToContainerawsAlphaNodePoolConfigSshConfig(p.GetSshConfig()),
		SpotConfig:                   ProtoToContainerawsAlphaNodePoolConfigSpotConfig(p.GetSpotConfig()),
		ProxyConfig:                  ProtoToContainerawsAlphaNodePoolConfigProxyConfig(p.GetProxyConfig()),
		InstancePlacement:            ProtoToContainerawsAlphaNodePoolConfigInstancePlacement(p.GetInstancePlacement()),
		ImageType:                    dcl.StringOrNil(p.GetImageType()),
		AutoscalingMetricsCollection: ProtoToContainerawsAlphaNodePoolConfigAutoscalingMetricsCollection(p.GetAutoscalingMetricsCollection()),
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerawsAlphaNodePoolConfigTaints(r))
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToNodePoolConfigRootVolume converts a NodePoolConfigRootVolume object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigRootVolume(p *alphapb.ContainerawsAlphaNodePoolConfigRootVolume) *alpha.NodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigRootVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToNodePoolConfigTaints converts a NodePoolConfigTaints object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigTaints(p *alphapb.ContainerawsAlphaNodePoolConfigTaints) *alpha.NodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.GetKey()),
		Value:  dcl.StringOrNil(p.GetValue()),
		Effect: ProtoToContainerawsAlphaNodePoolConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToNodePoolConfigConfigEncryption converts a NodePoolConfigConfigEncryption object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigConfigEncryption(p *alphapb.ContainerawsAlphaNodePoolConfigConfigEncryption) *alpha.NodePoolConfigConfigEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigConfigEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToNodePoolConfigSshConfig converts a NodePoolConfigSshConfig object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigSshConfig(p *alphapb.ContainerawsAlphaNodePoolConfigSshConfig) *alpha.NodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.GetEc2KeyPair()),
	}
	return obj
}

// ProtoToNodePoolConfigSpotConfig converts a NodePoolConfigSpotConfig object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigSpotConfig(p *alphapb.ContainerawsAlphaNodePoolConfigSpotConfig) *alpha.NodePoolConfigSpotConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigSpotConfig{}
	for _, r := range p.GetInstanceTypes() {
		obj.InstanceTypes = append(obj.InstanceTypes, r)
	}
	return obj
}

// ProtoToNodePoolConfigProxyConfig converts a NodePoolConfigProxyConfig object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigProxyConfig(p *alphapb.ContainerawsAlphaNodePoolConfigProxyConfig) *alpha.NodePoolConfigProxyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigProxyConfig{
		SecretArn:     dcl.StringOrNil(p.GetSecretArn()),
		SecretVersion: dcl.StringOrNil(p.GetSecretVersion()),
	}
	return obj
}

// ProtoToNodePoolConfigInstancePlacement converts a NodePoolConfigInstancePlacement object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigInstancePlacement(p *alphapb.ContainerawsAlphaNodePoolConfigInstancePlacement) *alpha.NodePoolConfigInstancePlacement {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigInstancePlacement{
		Tenancy: ProtoToContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum(p.GetTenancy()),
	}
	return obj
}

// ProtoToNodePoolConfigAutoscalingMetricsCollection converts a NodePoolConfigAutoscalingMetricsCollection object from its proto representation.
func ProtoToContainerawsAlphaNodePoolConfigAutoscalingMetricsCollection(p *alphapb.ContainerawsAlphaNodePoolConfigAutoscalingMetricsCollection) *alpha.NodePoolConfigAutoscalingMetricsCollection {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigAutoscalingMetricsCollection{
		Granularity: dcl.StringOrNil(p.GetGranularity()),
	}
	for _, r := range p.GetMetrics() {
		obj.Metrics = append(obj.Metrics, r)
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling object from its proto representation.
func ProtoToContainerawsAlphaNodePoolAutoscaling(p *alphapb.ContainerawsAlphaNodePoolAutoscaling) *alpha.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.GetMinNodeCount()),
		MaxNodeCount: dcl.Int64OrNil(p.GetMaxNodeCount()),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint object from its proto representation.
func ProtoToContainerawsAlphaNodePoolMaxPodsConstraint(p *alphapb.ContainerawsAlphaNodePoolMaxPodsConstraint) *alpha.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.GetMaxPodsPerNode()),
	}
	return obj
}

// ProtoToNodePoolManagement converts a NodePoolManagement object from its proto representation.
func ProtoToContainerawsAlphaNodePoolManagement(p *alphapb.ContainerawsAlphaNodePoolManagement) *alpha.NodePoolManagement {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolManagement{
		AutoRepair: dcl.Bool(p.GetAutoRepair()),
	}
	return obj
}

// ProtoToNodePoolUpdateSettings converts a NodePoolUpdateSettings object from its proto representation.
func ProtoToContainerawsAlphaNodePoolUpdateSettings(p *alphapb.ContainerawsAlphaNodePoolUpdateSettings) *alpha.NodePoolUpdateSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolUpdateSettings{
		SurgeSettings: ProtoToContainerawsAlphaNodePoolUpdateSettingsSurgeSettings(p.GetSurgeSettings()),
	}
	return obj
}

// ProtoToNodePoolUpdateSettingsSurgeSettings converts a NodePoolUpdateSettingsSurgeSettings object from its proto representation.
func ProtoToContainerawsAlphaNodePoolUpdateSettingsSurgeSettings(p *alphapb.ContainerawsAlphaNodePoolUpdateSettingsSurgeSettings) *alpha.NodePoolUpdateSettingsSurgeSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolUpdateSettingsSurgeSettings{
		MaxSurge:       dcl.Int64OrNil(p.GetMaxSurge()),
		MaxUnavailable: dcl.Int64OrNil(p.GetMaxUnavailable()),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *alphapb.ContainerawsAlphaNodePool) *alpha.NodePool {
	obj := &alpha.NodePool{
		Name:              dcl.StringOrNil(p.GetName()),
		Version:           dcl.StringOrNil(p.GetVersion()),
		Config:            ProtoToContainerawsAlphaNodePoolConfig(p.GetConfig()),
		Autoscaling:       ProtoToContainerawsAlphaNodePoolAutoscaling(p.GetAutoscaling()),
		SubnetId:          dcl.StringOrNil(p.GetSubnetId()),
		State:             ProtoToContainerawsAlphaNodePoolStateEnum(p.GetState()),
		Uid:               dcl.StringOrNil(p.GetUid()),
		Reconciling:       dcl.Bool(p.GetReconciling()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Etag:              dcl.StringOrNil(p.GetEtag()),
		MaxPodsConstraint: ProtoToContainerawsAlphaNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Management:        ProtoToContainerawsAlphaNodePoolManagement(p.GetManagement()),
		UpdateSettings:    ProtoToContainerawsAlphaNodePoolUpdateSettings(p.GetUpdateSettings()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
		Cluster:           dcl.StringOrNil(p.GetCluster()),
	}
	return obj
}

// NodePoolConfigRootVolumeVolumeTypeEnumToProto converts a NodePoolConfigRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnumToProto(e *alpha.NodePoolConfigRootVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum_value["NodePoolConfigRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnum(0)
}

// NodePoolConfigTaintsEffectEnumToProto converts a NodePoolConfigTaintsEffectEnum enum to its proto representation.
func ContainerawsAlphaNodePoolConfigTaintsEffectEnumToProto(e *alpha.NodePoolConfigTaintsEffectEnum) alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum_value["NodePoolConfigTaintsEffectEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum(v)
	}
	return alphapb.ContainerawsAlphaNodePoolConfigTaintsEffectEnum(0)
}

// NodePoolConfigInstancePlacementTenancyEnumToProto converts a NodePoolConfigInstancePlacementTenancyEnum enum to its proto representation.
func ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnumToProto(e *alpha.NodePoolConfigInstancePlacementTenancyEnum) alphapb.ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum_value["NodePoolConfigInstancePlacementTenancyEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum(v)
	}
	return alphapb.ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnum(0)
}

// NodePoolStateEnumToProto converts a NodePoolStateEnum enum to its proto representation.
func ContainerawsAlphaNodePoolStateEnumToProto(e *alpha.NodePoolStateEnum) alphapb.ContainerawsAlphaNodePoolStateEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaNodePoolStateEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaNodePoolStateEnum_value["NodePoolStateEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaNodePoolStateEnum(v)
	}
	return alphapb.ContainerawsAlphaNodePoolStateEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig object to its proto representation.
func ContainerawsAlphaNodePoolConfigToProto(o *alpha.NodePoolConfig) *alphapb.ContainerawsAlphaNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfig{}
	p.SetInstanceType(dcl.ValueOrEmptyString(o.InstanceType))
	p.SetRootVolume(ContainerawsAlphaNodePoolConfigRootVolumeToProto(o.RootVolume))
	p.SetIamInstanceProfile(dcl.ValueOrEmptyString(o.IamInstanceProfile))
	p.SetConfigEncryption(ContainerawsAlphaNodePoolConfigConfigEncryptionToProto(o.ConfigEncryption))
	p.SetSshConfig(ContainerawsAlphaNodePoolConfigSshConfigToProto(o.SshConfig))
	p.SetSpotConfig(ContainerawsAlphaNodePoolConfigSpotConfigToProto(o.SpotConfig))
	p.SetProxyConfig(ContainerawsAlphaNodePoolConfigProxyConfigToProto(o.ProxyConfig))
	p.SetInstancePlacement(ContainerawsAlphaNodePoolConfigInstancePlacementToProto(o.InstancePlacement))
	p.SetImageType(dcl.ValueOrEmptyString(o.ImageType))
	p.SetAutoscalingMetricsCollection(ContainerawsAlphaNodePoolConfigAutoscalingMetricsCollectionToProto(o.AutoscalingMetricsCollection))
	sTaints := make([]*alphapb.ContainerawsAlphaNodePoolConfigTaints, len(o.Taints))
	for i, r := range o.Taints {
		sTaints[i] = ContainerawsAlphaNodePoolConfigTaintsToProto(&r)
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
func ContainerawsAlphaNodePoolConfigRootVolumeToProto(o *alpha.NodePoolConfigRootVolume) *alphapb.ContainerawsAlphaNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsAlphaNodePoolConfigRootVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// NodePoolConfigTaintsToProto converts a NodePoolConfigTaints object to its proto representation.
func ContainerawsAlphaNodePoolConfigTaintsToProto(o *alpha.NodePoolConfigTaints) *alphapb.ContainerawsAlphaNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigTaints{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetEffect(ContainerawsAlphaNodePoolConfigTaintsEffectEnumToProto(o.Effect))
	return p
}

// NodePoolConfigConfigEncryptionToProto converts a NodePoolConfigConfigEncryption object to its proto representation.
func ContainerawsAlphaNodePoolConfigConfigEncryptionToProto(o *alpha.NodePoolConfigConfigEncryption) *alphapb.ContainerawsAlphaNodePoolConfigConfigEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigConfigEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// NodePoolConfigSshConfigToProto converts a NodePoolConfigSshConfig object to its proto representation.
func ContainerawsAlphaNodePoolConfigSshConfigToProto(o *alpha.NodePoolConfigSshConfig) *alphapb.ContainerawsAlphaNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigSshConfig{}
	p.SetEc2KeyPair(dcl.ValueOrEmptyString(o.Ec2KeyPair))
	return p
}

// NodePoolConfigSpotConfigToProto converts a NodePoolConfigSpotConfig object to its proto representation.
func ContainerawsAlphaNodePoolConfigSpotConfigToProto(o *alpha.NodePoolConfigSpotConfig) *alphapb.ContainerawsAlphaNodePoolConfigSpotConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigSpotConfig{}
	sInstanceTypes := make([]string, len(o.InstanceTypes))
	for i, r := range o.InstanceTypes {
		sInstanceTypes[i] = r
	}
	p.SetInstanceTypes(sInstanceTypes)
	return p
}

// NodePoolConfigProxyConfigToProto converts a NodePoolConfigProxyConfig object to its proto representation.
func ContainerawsAlphaNodePoolConfigProxyConfigToProto(o *alpha.NodePoolConfigProxyConfig) *alphapb.ContainerawsAlphaNodePoolConfigProxyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigProxyConfig{}
	p.SetSecretArn(dcl.ValueOrEmptyString(o.SecretArn))
	p.SetSecretVersion(dcl.ValueOrEmptyString(o.SecretVersion))
	return p
}

// NodePoolConfigInstancePlacementToProto converts a NodePoolConfigInstancePlacement object to its proto representation.
func ContainerawsAlphaNodePoolConfigInstancePlacementToProto(o *alpha.NodePoolConfigInstancePlacement) *alphapb.ContainerawsAlphaNodePoolConfigInstancePlacement {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigInstancePlacement{}
	p.SetTenancy(ContainerawsAlphaNodePoolConfigInstancePlacementTenancyEnumToProto(o.Tenancy))
	return p
}

// NodePoolConfigAutoscalingMetricsCollectionToProto converts a NodePoolConfigAutoscalingMetricsCollection object to its proto representation.
func ContainerawsAlphaNodePoolConfigAutoscalingMetricsCollectionToProto(o *alpha.NodePoolConfigAutoscalingMetricsCollection) *alphapb.ContainerawsAlphaNodePoolConfigAutoscalingMetricsCollection {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolConfigAutoscalingMetricsCollection{}
	p.SetGranularity(dcl.ValueOrEmptyString(o.Granularity))
	sMetrics := make([]string, len(o.Metrics))
	for i, r := range o.Metrics {
		sMetrics[i] = r
	}
	p.SetMetrics(sMetrics)
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling object to its proto representation.
func ContainerawsAlphaNodePoolAutoscalingToProto(o *alpha.NodePoolAutoscaling) *alphapb.ContainerawsAlphaNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolAutoscaling{}
	p.SetMinNodeCount(dcl.ValueOrEmptyInt64(o.MinNodeCount))
	p.SetMaxNodeCount(dcl.ValueOrEmptyInt64(o.MaxNodeCount))
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint object to its proto representation.
func ContainerawsAlphaNodePoolMaxPodsConstraintToProto(o *alpha.NodePoolMaxPodsConstraint) *alphapb.ContainerawsAlphaNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolMaxPodsConstraint{}
	p.SetMaxPodsPerNode(dcl.ValueOrEmptyInt64(o.MaxPodsPerNode))
	return p
}

// NodePoolManagementToProto converts a NodePoolManagement object to its proto representation.
func ContainerawsAlphaNodePoolManagementToProto(o *alpha.NodePoolManagement) *alphapb.ContainerawsAlphaNodePoolManagement {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolManagement{}
	p.SetAutoRepair(dcl.ValueOrEmptyBool(o.AutoRepair))
	return p
}

// NodePoolUpdateSettingsToProto converts a NodePoolUpdateSettings object to its proto representation.
func ContainerawsAlphaNodePoolUpdateSettingsToProto(o *alpha.NodePoolUpdateSettings) *alphapb.ContainerawsAlphaNodePoolUpdateSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolUpdateSettings{}
	p.SetSurgeSettings(ContainerawsAlphaNodePoolUpdateSettingsSurgeSettingsToProto(o.SurgeSettings))
	return p
}

// NodePoolUpdateSettingsSurgeSettingsToProto converts a NodePoolUpdateSettingsSurgeSettings object to its proto representation.
func ContainerawsAlphaNodePoolUpdateSettingsSurgeSettingsToProto(o *alpha.NodePoolUpdateSettingsSurgeSettings) *alphapb.ContainerawsAlphaNodePoolUpdateSettingsSurgeSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaNodePoolUpdateSettingsSurgeSettings{}
	p.SetMaxSurge(dcl.ValueOrEmptyInt64(o.MaxSurge))
	p.SetMaxUnavailable(dcl.ValueOrEmptyInt64(o.MaxUnavailable))
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *alpha.NodePool) *alphapb.ContainerawsAlphaNodePool {
	p := &alphapb.ContainerawsAlphaNodePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyString(resource.Version))
	p.SetConfig(ContainerawsAlphaNodePoolConfigToProto(resource.Config))
	p.SetAutoscaling(ContainerawsAlphaNodePoolAutoscalingToProto(resource.Autoscaling))
	p.SetSubnetId(dcl.ValueOrEmptyString(resource.SubnetId))
	p.SetState(ContainerawsAlphaNodePoolStateEnumToProto(resource.State))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetMaxPodsConstraint(ContainerawsAlphaNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint))
	p.SetManagement(ContainerawsAlphaNodePoolManagementToProto(resource.Management))
	p.SetUpdateSettings(ContainerawsAlphaNodePoolUpdateSettingsToProto(resource.UpdateSettings))
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
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerawsAlphaNodePoolRequest) (*alphapb.ContainerawsAlphaNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// applyContainerawsAlphaNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerawsAlphaNodePool(ctx context.Context, request *alphapb.ApplyContainerawsAlphaNodePoolRequest) (*alphapb.ContainerawsAlphaNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerawsAlphaNodePool(ctx context.Context, request *alphapb.DeleteContainerawsAlphaNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerawsAlphaNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerawsAlphaNodePool(ctx context.Context, request *alphapb.ListContainerawsAlphaNodePoolRequest) (*alphapb.ListContainerawsAlphaNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.GetProject(), request.GetLocation(), request.GetCluster())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerawsAlphaNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListContainerawsAlphaNodePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
