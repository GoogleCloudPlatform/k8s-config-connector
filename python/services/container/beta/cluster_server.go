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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/container/beta/container_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container/beta"
)

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum converts a ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum enum from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(e betapb.ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum) *beta.ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(n[len("ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterAddonsConfigIstioConfigAuthEnum converts a ClusterAddonsConfigIstioConfigAuthEnum enum from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigIstioConfigAuthEnum(e betapb.ContainerBetaClusterAddonsConfigIstioConfigAuthEnum) *beta.ClusterAddonsConfigIstioConfigAuthEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterAddonsConfigIstioConfigAuthEnum_name[int32(e)]; ok {
		e := beta.ClusterAddonsConfigIstioConfigAuthEnum(n[len("ContainerBetaClusterAddonsConfigIstioConfigAuthEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigWorkloadMetadataConfigModeEnum converts a ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(e betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum) *beta.ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum_name[int32(e)]; ok {
		e := beta.ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(n[len("ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum converts a ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum(e betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum) *beta.ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum_name[int32(e)]; ok {
		e := beta.ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum(n[len("ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigTaintsEffectEnum converts a ClusterNodePoolsConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigTaintsEffectEnum(e betapb.ContainerBetaClusterNodePoolsConfigTaintsEffectEnum) *beta.ClusterNodePoolsConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodePoolsConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := beta.ClusterNodePoolsConfigTaintsEffectEnum(n[len("ContainerBetaClusterNodePoolsConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigSandboxConfigTypeEnum converts a ClusterNodePoolsConfigSandboxConfigTypeEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum(e betapb.ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum) *beta.ClusterNodePoolsConfigSandboxConfigTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterNodePoolsConfigSandboxConfigTypeEnum(n[len("ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(e betapb.ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum) *beta.ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(n[len("ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsStatusEnum converts a ClusterNodePoolsStatusEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodePoolsStatusEnum(e betapb.ContainerBetaClusterNodePoolsStatusEnum) *beta.ClusterNodePoolsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodePoolsStatusEnum_name[int32(e)]; ok {
		e := beta.ClusterNodePoolsStatusEnum(n[len("ContainerBetaClusterNodePoolsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConditionsCodeEnum converts a ClusterNodePoolsConditionsCodeEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConditionsCodeEnum(e betapb.ContainerBetaClusterNodePoolsConditionsCodeEnum) *beta.ClusterNodePoolsConditionsCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodePoolsConditionsCodeEnum_name[int32(e)]; ok {
		e := beta.ClusterNodePoolsConditionsCodeEnum(n[len("ContainerBetaClusterNodePoolsConditionsCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConditionsCanonicalCodeEnum converts a ClusterNodePoolsConditionsCanonicalCodeEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum(e betapb.ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum) *beta.ClusterNodePoolsConditionsCanonicalCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum_name[int32(e)]; ok {
		e := beta.ClusterNodePoolsConditionsCanonicalCodeEnum(n[len("ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworkPolicyProviderEnum converts a ClusterNetworkPolicyProviderEnum enum from its proto representation.
func ProtoToContainerBetaClusterNetworkPolicyProviderEnum(e betapb.ContainerBetaClusterNetworkPolicyProviderEnum) *beta.ClusterNetworkPolicyProviderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNetworkPolicyProviderEnum_name[int32(e)]; ok {
		e := beta.ClusterNetworkPolicyProviderEnum(n[len("ContainerBetaClusterNetworkPolicyProviderEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterAutoscalingAutoscalingProfileEnum converts a ClusterAutoscalingAutoscalingProfileEnum enum from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoscalingProfileEnum(e betapb.ContainerBetaClusterAutoscalingAutoscalingProfileEnum) *beta.ClusterAutoscalingAutoscalingProfileEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterAutoscalingAutoscalingProfileEnum_name[int32(e)]; ok {
		e := beta.ClusterAutoscalingAutoscalingProfileEnum(n[len("ContainerBetaClusterAutoscalingAutoscalingProfileEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworkConfigPrivateIPv6GoogleAccessEnum converts a ClusterNetworkConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum(e betapb.ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum) *beta.ClusterNetworkConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := beta.ClusterNetworkConfigPrivateIPv6GoogleAccessEnum(n[len("ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworkConfigDatapathProviderEnum converts a ClusterNetworkConfigDatapathProviderEnum enum from its proto representation.
func ProtoToContainerBetaClusterNetworkConfigDatapathProviderEnum(e betapb.ContainerBetaClusterNetworkConfigDatapathProviderEnum) *beta.ClusterNetworkConfigDatapathProviderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNetworkConfigDatapathProviderEnum_name[int32(e)]; ok {
		e := beta.ClusterNetworkConfigDatapathProviderEnum(n[len("ContainerBetaClusterNetworkConfigDatapathProviderEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterDatabaseEncryptionStateEnum converts a ClusterDatabaseEncryptionStateEnum enum from its proto representation.
func ProtoToContainerBetaClusterDatabaseEncryptionStateEnum(e betapb.ContainerBetaClusterDatabaseEncryptionStateEnum) *beta.ClusterDatabaseEncryptionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterDatabaseEncryptionStateEnum_name[int32(e)]; ok {
		e := beta.ClusterDatabaseEncryptionStateEnum(n[len("ContainerBetaClusterDatabaseEncryptionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConditionsCanonicalCodeEnum converts a ClusterConditionsCanonicalCodeEnum enum from its proto representation.
func ProtoToContainerBetaClusterConditionsCanonicalCodeEnum(e betapb.ContainerBetaClusterConditionsCanonicalCodeEnum) *beta.ClusterConditionsCanonicalCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterConditionsCanonicalCodeEnum_name[int32(e)]; ok {
		e := beta.ClusterConditionsCanonicalCodeEnum(n[len("ContainerBetaClusterConditionsCanonicalCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigWorkloadMetadataConfigModeEnum converts a ClusterNodeConfigWorkloadMetadataConfigModeEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum(e betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum) *beta.ClusterNodeConfigWorkloadMetadataConfigModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum_name[int32(e)]; ok {
		e := beta.ClusterNodeConfigWorkloadMetadataConfigModeEnum(n[len("ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum converts a ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum(e betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum) *beta.ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum_name[int32(e)]; ok {
		e := beta.ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum(n[len("ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigTaintsEffectEnum converts a ClusterNodeConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodeConfigTaintsEffectEnum(e betapb.ContainerBetaClusterNodeConfigTaintsEffectEnum) *beta.ClusterNodeConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodeConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := beta.ClusterNodeConfigTaintsEffectEnum(n[len("ContainerBetaClusterNodeConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigSandboxConfigTypeEnum converts a ClusterNodeConfigSandboxConfigTypeEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodeConfigSandboxConfigTypeEnum(e betapb.ContainerBetaClusterNodeConfigSandboxConfigTypeEnum) *beta.ClusterNodeConfigSandboxConfigTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodeConfigSandboxConfigTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterNodeConfigSandboxConfigTypeEnum(n[len("ContainerBetaClusterNodeConfigSandboxConfigTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(e betapb.ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum) *beta.ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(n[len("ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterReleaseChannelChannelEnum converts a ClusterReleaseChannelChannelEnum enum from its proto representation.
func ProtoToContainerBetaClusterReleaseChannelChannelEnum(e betapb.ContainerBetaClusterReleaseChannelChannelEnum) *beta.ClusterReleaseChannelChannelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterReleaseChannelChannelEnum_name[int32(e)]; ok {
		e := beta.ClusterReleaseChannelChannelEnum(n[len("ContainerBetaClusterReleaseChannelChannelEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterClusterTelemetryTypeEnum converts a ClusterClusterTelemetryTypeEnum enum from its proto representation.
func ProtoToContainerBetaClusterClusterTelemetryTypeEnum(e betapb.ContainerBetaClusterClusterTelemetryTypeEnum) *beta.ClusterClusterTelemetryTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaClusterClusterTelemetryTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterClusterTelemetryTypeEnum(n[len("ContainerBetaClusterClusterTelemetryTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterMasterAuth converts a ClusterMasterAuth resource from its proto representation.
func ProtoToContainerBetaClusterMasterAuth(p *betapb.ContainerBetaClusterMasterAuth) *beta.ClusterMasterAuth {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMasterAuth{
		Username:                dcl.StringOrNil(p.Username),
		Password:                dcl.StringOrNil(p.Password),
		ClientCertificateConfig: ProtoToContainerBetaClusterMasterAuthClientCertificateConfig(p.GetClientCertificateConfig()),
		ClusterCaCertificate:    dcl.StringOrNil(p.ClusterCaCertificate),
		ClientCertificate:       dcl.StringOrNil(p.ClientCertificate),
		ClientKey:               dcl.StringOrNil(p.ClientKey),
	}
	return obj
}

// ProtoToClusterMasterAuthClientCertificateConfig converts a ClusterMasterAuthClientCertificateConfig resource from its proto representation.
func ProtoToContainerBetaClusterMasterAuthClientCertificateConfig(p *betapb.ContainerBetaClusterMasterAuthClientCertificateConfig) *beta.ClusterMasterAuthClientCertificateConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMasterAuthClientCertificateConfig{
		IssueClientCertificate: dcl.Bool(p.IssueClientCertificate),
	}
	return obj
}

// ProtoToClusterAddonsConfig converts a ClusterAddonsConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfig(p *betapb.ContainerBetaClusterAddonsConfig) *beta.ClusterAddonsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfig{
		HttpLoadBalancing:                ProtoToContainerBetaClusterAddonsConfigHttpLoadBalancing(p.GetHttpLoadBalancing()),
		HorizontalPodAutoscaling:         ProtoToContainerBetaClusterAddonsConfigHorizontalPodAutoscaling(p.GetHorizontalPodAutoscaling()),
		KubernetesDashboard:              ProtoToContainerBetaClusterAddonsConfigKubernetesDashboard(p.GetKubernetesDashboard()),
		NetworkPolicyConfig:              ProtoToContainerBetaClusterAddonsConfigNetworkPolicyConfig(p.GetNetworkPolicyConfig()),
		CloudRunConfig:                   ProtoToContainerBetaClusterAddonsConfigCloudRunConfig(p.GetCloudRunConfig()),
		DnsCacheConfig:                   ProtoToContainerBetaClusterAddonsConfigDnsCacheConfig(p.GetDnsCacheConfig()),
		ConfigConnectorConfig:            ProtoToContainerBetaClusterAddonsConfigConfigConnectorConfig(p.GetConfigConnectorConfig()),
		GcePersistentDiskCsiDriverConfig: ProtoToContainerBetaClusterAddonsConfigGcePersistentDiskCsiDriverConfig(p.GetGcePersistentDiskCsiDriverConfig()),
		IstioConfig:                      ProtoToContainerBetaClusterAddonsConfigIstioConfig(p.GetIstioConfig()),
		KalmConfig:                       ProtoToContainerBetaClusterAddonsConfigKalmConfig(p.GetKalmConfig()),
	}
	return obj
}

// ProtoToClusterAddonsConfigHttpLoadBalancing converts a ClusterAddonsConfigHttpLoadBalancing resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigHttpLoadBalancing(p *betapb.ContainerBetaClusterAddonsConfigHttpLoadBalancing) *beta.ClusterAddonsConfigHttpLoadBalancing {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigHttpLoadBalancing{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigHorizontalPodAutoscaling converts a ClusterAddonsConfigHorizontalPodAutoscaling resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigHorizontalPodAutoscaling(p *betapb.ContainerBetaClusterAddonsConfigHorizontalPodAutoscaling) *beta.ClusterAddonsConfigHorizontalPodAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigHorizontalPodAutoscaling{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigKubernetesDashboard converts a ClusterAddonsConfigKubernetesDashboard resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigKubernetesDashboard(p *betapb.ContainerBetaClusterAddonsConfigKubernetesDashboard) *beta.ClusterAddonsConfigKubernetesDashboard {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigKubernetesDashboard{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigNetworkPolicyConfig converts a ClusterAddonsConfigNetworkPolicyConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigNetworkPolicyConfig(p *betapb.ContainerBetaClusterAddonsConfigNetworkPolicyConfig) *beta.ClusterAddonsConfigNetworkPolicyConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigNetworkPolicyConfig{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigCloudRunConfig converts a ClusterAddonsConfigCloudRunConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigCloudRunConfig(p *betapb.ContainerBetaClusterAddonsConfigCloudRunConfig) *beta.ClusterAddonsConfigCloudRunConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigCloudRunConfig{
		Disabled:         dcl.Bool(p.Disabled),
		LoadBalancerType: ProtoToContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(p.GetLoadBalancerType()),
	}
	return obj
}

// ProtoToClusterAddonsConfigDnsCacheConfig converts a ClusterAddonsConfigDnsCacheConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigDnsCacheConfig(p *betapb.ContainerBetaClusterAddonsConfigDnsCacheConfig) *beta.ClusterAddonsConfigDnsCacheConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigDnsCacheConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigConfigConnectorConfig converts a ClusterAddonsConfigConfigConnectorConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigConfigConnectorConfig(p *betapb.ContainerBetaClusterAddonsConfigConfigConnectorConfig) *beta.ClusterAddonsConfigConfigConnectorConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigConfigConnectorConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigGcePersistentDiskCsiDriverConfig converts a ClusterAddonsConfigGcePersistentDiskCsiDriverConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigGcePersistentDiskCsiDriverConfig(p *betapb.ContainerBetaClusterAddonsConfigGcePersistentDiskCsiDriverConfig) *beta.ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigGcePersistentDiskCsiDriverConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigIstioConfig converts a ClusterAddonsConfigIstioConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigIstioConfig(p *betapb.ContainerBetaClusterAddonsConfigIstioConfig) *beta.ClusterAddonsConfigIstioConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigIstioConfig{
		Disabled: dcl.Bool(p.Disabled),
		Auth:     ProtoToContainerBetaClusterAddonsConfigIstioConfigAuthEnum(p.GetAuth()),
	}
	return obj
}

// ProtoToClusterAddonsConfigKalmConfig converts a ClusterAddonsConfigKalmConfig resource from its proto representation.
func ProtoToContainerBetaClusterAddonsConfigKalmConfig(p *betapb.ContainerBetaClusterAddonsConfigKalmConfig) *beta.ClusterAddonsConfigKalmConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAddonsConfigKalmConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterNodePools converts a ClusterNodePools resource from its proto representation.
func ProtoToContainerBetaClusterNodePools(p *betapb.ContainerBetaClusterNodePools) *beta.ClusterNodePools {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePools{
		Name:              dcl.StringOrNil(p.Name),
		Config:            ProtoToContainerBetaClusterNodePoolsConfig(p.GetConfig()),
		InitialNodeCount:  dcl.Int64OrNil(p.InitialNodeCount),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Version:           dcl.StringOrNil(p.Version),
		Status:            ProtoToContainerBetaClusterNodePoolsStatusEnum(p.GetStatus()),
		StatusMessage:     dcl.StringOrNil(p.StatusMessage),
		Autoscaling:       ProtoToContainerBetaClusterNodePoolsAutoscaling(p.GetAutoscaling()),
		Management:        ProtoToContainerBetaClusterNodePoolsManagement(p.GetManagement()),
		MaxPodsConstraint: ProtoToContainerBetaClusterNodePoolsMaxPodsConstraint(p.GetMaxPodsConstraint()),
		PodIPv4CidrSize:   dcl.Int64OrNil(p.PodIpv4CidrSize),
		UpgradeSettings:   ProtoToContainerBetaClusterNodePoolsUpgradeSettings(p.GetUpgradeSettings()),
		NetworkConfig:     ProtoToContainerBetaClusterNodePoolsNetworkConfig(p.GetNetworkConfig()),
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	for _, r := range p.GetInstanceGroupUrls() {
		obj.InstanceGroupUrls = append(obj.InstanceGroupUrls, r)
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToContainerBetaClusterNodePoolsConditions(r))
	}
	return obj
}

// ProtoToClusterNodePoolsConfig converts a ClusterNodePoolsConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfig(p *betapb.ContainerBetaClusterNodePoolsConfig) *beta.ClusterNodePoolsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfig{
		MachineType:            dcl.StringOrNil(p.MachineType),
		DiskSizeGb:             dcl.Int64OrNil(p.DiskSizeGb),
		ServiceAccount:         dcl.StringOrNil(p.ServiceAccount),
		ImageType:              dcl.StringOrNil(p.ImageType),
		LocalSsdCount:          dcl.Int64OrNil(p.LocalSsdCount),
		Preemptible:            dcl.Bool(p.Preemptible),
		DiskType:               dcl.StringOrNil(p.DiskType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		WorkloadMetadataConfig: ProtoToContainerBetaClusterNodePoolsConfigWorkloadMetadataConfig(p.GetWorkloadMetadataConfig()),
		SandboxConfig:          ProtoToContainerBetaClusterNodePoolsConfigSandboxConfig(p.GetSandboxConfig()),
		NodeGroup:              dcl.StringOrNil(p.NodeGroup),
		ReservationAffinity:    ProtoToContainerBetaClusterNodePoolsConfigReservationAffinity(p.GetReservationAffinity()),
		ShieldedInstanceConfig: ProtoToContainerBetaClusterNodePoolsConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		LinuxNodeConfig:        ProtoToContainerBetaClusterNodePoolsConfigLinuxNodeConfig(p.GetLinuxNodeConfig()),
		KubeletConfig:          ProtoToContainerBetaClusterNodePoolsConfigKubeletConfig(p.GetKubeletConfig()),
		BootDiskKmsKey:         dcl.StringOrNil(p.BootDiskKmsKey),
		EphemeralStorageConfig: ProtoToContainerBetaClusterNodePoolsConfigEphemeralStorageConfig(p.GetEphemeralStorageConfig()),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToContainerBetaClusterNodePoolsConfigAccelerators(r))
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerBetaClusterNodePoolsConfigTaints(r))
	}
	return obj
}

// ProtoToClusterNodePoolsConfigAccelerators converts a ClusterNodePoolsConfigAccelerators resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigAccelerators(p *betapb.ContainerBetaClusterNodePoolsConfigAccelerators) *beta.ClusterNodePoolsConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigWorkloadMetadataConfig converts a ClusterNodePoolsConfigWorkloadMetadataConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigWorkloadMetadataConfig(p *betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfig) *beta.ClusterNodePoolsConfigWorkloadMetadataConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigWorkloadMetadataConfig{
		Mode:         ProtoToContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(p.GetMode()),
		NodeMetadata: ProtoToContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum(p.GetNodeMetadata()),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigTaints converts a ClusterNodePoolsConfigTaints resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigTaints(p *betapb.ContainerBetaClusterNodePoolsConfigTaints) *beta.ClusterNodePoolsConfigTaints {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToContainerBetaClusterNodePoolsConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigSandboxConfig converts a ClusterNodePoolsConfigSandboxConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigSandboxConfig(p *betapb.ContainerBetaClusterNodePoolsConfigSandboxConfig) *beta.ClusterNodePoolsConfigSandboxConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigSandboxConfig{
		Type:        ProtoToContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum(p.GetType()),
		SandboxType: dcl.StringOrNil(p.SandboxType),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigReservationAffinity converts a ClusterNodePoolsConfigReservationAffinity resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigReservationAffinity(p *betapb.ContainerBetaClusterNodePoolsConfigReservationAffinity) *beta.ClusterNodePoolsConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigReservationAffinity{
		ConsumeReservationType: ProtoToContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterNodePoolsConfigShieldedInstanceConfig converts a ClusterNodePoolsConfigShieldedInstanceConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigShieldedInstanceConfig(p *betapb.ContainerBetaClusterNodePoolsConfigShieldedInstanceConfig) *beta.ClusterNodePoolsConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigLinuxNodeConfig converts a ClusterNodePoolsConfigLinuxNodeConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigLinuxNodeConfig(p *betapb.ContainerBetaClusterNodePoolsConfigLinuxNodeConfig) *beta.ClusterNodePoolsConfigLinuxNodeConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigLinuxNodeConfig{}
	return obj
}

// ProtoToClusterNodePoolsConfigKubeletConfig converts a ClusterNodePoolsConfigKubeletConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigKubeletConfig(p *betapb.ContainerBetaClusterNodePoolsConfigKubeletConfig) *beta.ClusterNodePoolsConfigKubeletConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigKubeletConfig{
		CpuManagerPolicy:  dcl.StringOrNil(p.CpuManagerPolicy),
		CpuCfsQuota:       dcl.Bool(p.CpuCfsQuota),
		CpuCfsQuotaPeriod: dcl.StringOrNil(p.CpuCfsQuotaPeriod),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigEphemeralStorageConfig converts a ClusterNodePoolsConfigEphemeralStorageConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConfigEphemeralStorageConfig(p *betapb.ContainerBetaClusterNodePoolsConfigEphemeralStorageConfig) *beta.ClusterNodePoolsConfigEphemeralStorageConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConfigEphemeralStorageConfig{
		LocalSsdCount: dcl.Int64OrNil(p.LocalSsdCount),
	}
	return obj
}

// ProtoToClusterNodePoolsAutoscaling converts a ClusterNodePoolsAutoscaling resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsAutoscaling(p *betapb.ContainerBetaClusterNodePoolsAutoscaling) *beta.ClusterNodePoolsAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsAutoscaling{
		Enabled:         dcl.Bool(p.Enabled),
		MinNodeCount:    dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount:    dcl.Int64OrNil(p.MaxNodeCount),
		Autoprovisioned: dcl.Bool(p.Autoprovisioned),
	}
	return obj
}

// ProtoToClusterNodePoolsManagement converts a ClusterNodePoolsManagement resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsManagement(p *betapb.ContainerBetaClusterNodePoolsManagement) *beta.ClusterNodePoolsManagement {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsManagement{
		AutoUpgrade:    dcl.Bool(p.AutoUpgrade),
		AutoRepair:     dcl.Bool(p.AutoRepair),
		UpgradeOptions: ProtoToContainerBetaClusterNodePoolsManagementUpgradeOptions(p.GetUpgradeOptions()),
	}
	return obj
}

// ProtoToClusterNodePoolsManagementUpgradeOptions converts a ClusterNodePoolsManagementUpgradeOptions resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsManagementUpgradeOptions(p *betapb.ContainerBetaClusterNodePoolsManagementUpgradeOptions) *beta.ClusterNodePoolsManagementUpgradeOptions {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.StringOrNil(p.AutoUpgradeStartTime),
		Description:          dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToClusterNodePoolsMaxPodsConstraint converts a ClusterNodePoolsMaxPodsConstraint resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsMaxPodsConstraint(p *betapb.ContainerBetaClusterNodePoolsMaxPodsConstraint) *beta.ClusterNodePoolsMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToClusterNodePoolsConditions converts a ClusterNodePoolsConditions resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsConditions(p *betapb.ContainerBetaClusterNodePoolsConditions) *beta.ClusterNodePoolsConditions {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsConditions{
		Code:          ProtoToContainerBetaClusterNodePoolsConditionsCodeEnum(p.GetCode()),
		Message:       dcl.StringOrNil(p.Message),
		CanonicalCode: ProtoToContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum(p.GetCanonicalCode()),
	}
	return obj
}

// ProtoToClusterNodePoolsUpgradeSettings converts a ClusterNodePoolsUpgradeSettings resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsUpgradeSettings(p *betapb.ContainerBetaClusterNodePoolsUpgradeSettings) *beta.ClusterNodePoolsUpgradeSettings {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsUpgradeSettings{
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
	}
	return obj
}

// ProtoToClusterNodePoolsNetworkConfig converts a ClusterNodePoolsNetworkConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodePoolsNetworkConfig(p *betapb.ContainerBetaClusterNodePoolsNetworkConfig) *beta.ClusterNodePoolsNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodePoolsNetworkConfig{
		CreatePodRange:   dcl.Bool(p.CreatePodRange),
		PodRange:         dcl.StringOrNil(p.PodRange),
		PodIPv4CidrBlock: dcl.StringOrNil(p.PodIpv4CidrBlock),
	}
	return obj
}

// ProtoToClusterLegacyAbac converts a ClusterLegacyAbac resource from its proto representation.
func ProtoToContainerBetaClusterLegacyAbac(p *betapb.ContainerBetaClusterLegacyAbac) *beta.ClusterLegacyAbac {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterLegacyAbac{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterNetworkPolicy converts a ClusterNetworkPolicy resource from its proto representation.
func ProtoToContainerBetaClusterNetworkPolicy(p *betapb.ContainerBetaClusterNetworkPolicy) *beta.ClusterNetworkPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNetworkPolicy{
		Provider: ProtoToContainerBetaClusterNetworkPolicyProviderEnum(p.GetProvider()),
		Enabled:  dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterIPAllocationPolicy converts a ClusterIPAllocationPolicy resource from its proto representation.
func ProtoToContainerBetaClusterIPAllocationPolicy(p *betapb.ContainerBetaClusterIPAllocationPolicy) *beta.ClusterIPAllocationPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterIPAllocationPolicy{
		UseIPAliases:               dcl.Bool(p.UseIpAliases),
		CreateSubnetwork:           dcl.Bool(p.CreateSubnetwork),
		SubnetworkName:             dcl.StringOrNil(p.SubnetworkName),
		ClusterSecondaryRangeName:  dcl.StringOrNil(p.ClusterSecondaryRangeName),
		ServicesSecondaryRangeName: dcl.StringOrNil(p.ServicesSecondaryRangeName),
		ClusterIPv4CidrBlock:       dcl.StringOrNil(p.ClusterIpv4CidrBlock),
		NodeIPv4CidrBlock:          dcl.StringOrNil(p.NodeIpv4CidrBlock),
		ServicesIPv4CidrBlock:      dcl.StringOrNil(p.ServicesIpv4CidrBlock),
		TPUIPv4CidrBlock:           dcl.StringOrNil(p.TpuIpv4CidrBlock),
		ClusterIPv4Cidr:            dcl.StringOrNil(p.ClusterIpv4Cidr),
		NodeIPv4Cidr:               dcl.StringOrNil(p.NodeIpv4Cidr),
		ServicesIPv4Cidr:           dcl.StringOrNil(p.ServicesIpv4Cidr),
		UseRoutes:                  dcl.Bool(p.UseRoutes),
		AllowRouteOverlap:          dcl.Bool(p.AllowRouteOverlap),
	}
	return obj
}

// ProtoToClusterMasterAuthorizedNetworksConfig converts a ClusterMasterAuthorizedNetworksConfig resource from its proto representation.
func ProtoToContainerBetaClusterMasterAuthorizedNetworksConfig(p *betapb.ContainerBetaClusterMasterAuthorizedNetworksConfig) *beta.ClusterMasterAuthorizedNetworksConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMasterAuthorizedNetworksConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	for _, r := range p.GetCidrBlocks() {
		obj.CidrBlocks = append(obj.CidrBlocks, *ProtoToContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks(r))
	}
	return obj
}

// ProtoToClusterMasterAuthorizedNetworksConfigCidrBlocks converts a ClusterMasterAuthorizedNetworksConfigCidrBlocks resource from its proto representation.
func ProtoToContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks(p *betapb.ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks) *beta.ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMasterAuthorizedNetworksConfigCidrBlocks{
		DisplayName: dcl.StringOrNil(p.DisplayName),
		CidrBlock:   dcl.StringOrNil(p.CidrBlock),
	}
	return obj
}

// ProtoToClusterBinaryAuthorization converts a ClusterBinaryAuthorization resource from its proto representation.
func ProtoToContainerBetaClusterBinaryAuthorization(p *betapb.ContainerBetaClusterBinaryAuthorization) *beta.ClusterBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterBinaryAuthorization{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAutoscaling converts a ClusterAutoscaling resource from its proto representation.
func ProtoToContainerBetaClusterAutoscaling(p *betapb.ContainerBetaClusterAutoscaling) *beta.ClusterAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscaling{
		EnableNodeAutoprovisioning:       dcl.Bool(p.EnableNodeAutoprovisioning),
		AutoprovisioningNodePoolDefaults: ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults(p.GetAutoprovisioningNodePoolDefaults()),
		AutoscalingProfile:               ProtoToContainerBetaClusterAutoscalingAutoscalingProfileEnum(p.GetAutoscalingProfile()),
	}
	for _, r := range p.GetResourceLimits() {
		obj.ResourceLimits = append(obj.ResourceLimits, *ProtoToContainerBetaClusterAutoscalingResourceLimits(r))
	}
	for _, r := range p.GetAutoprovisioningLocations() {
		obj.AutoprovisioningLocations = append(obj.AutoprovisioningLocations, r)
	}
	return obj
}

// ProtoToClusterAutoscalingResourceLimits converts a ClusterAutoscalingResourceLimits resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingResourceLimits(p *betapb.ContainerBetaClusterAutoscalingResourceLimits) *beta.ClusterAutoscalingResourceLimits {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingResourceLimits{
		ResourceType: dcl.StringOrNil(p.ResourceType),
		Minimum:      dcl.Int64OrNil(p.Minimum),
		Maximum:      dcl.Int64OrNil(p.Maximum),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaults converts a ClusterAutoscalingAutoprovisioningNodePoolDefaults resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults(p *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults) *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingAutoprovisioningNodePoolDefaults{
		ServiceAccount:         dcl.StringOrNil(p.ServiceAccount),
		UpgradeSettings:        ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(p.GetUpgradeSettings()),
		Management:             ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(p.GetManagement()),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		DiskSizeGb:             dcl.Int64OrNil(p.DiskSizeGb),
		DiskType:               dcl.StringOrNil(p.DiskType),
		ShieldedInstanceConfig: ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		BootDiskKmsKey:         dcl.StringOrNil(p.BootDiskKmsKey),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(p *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(p *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{
		AutoUpgrade:    dcl.Bool(p.AutoUpgrade),
		AutoRepair:     dcl.Bool(p.AutoRepair),
		UpgradeOptions: ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(p.GetUpgradeOptions()),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(p *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.StringOrNil(p.AutoUpgradeStartTime),
		Description:          dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig resource from its proto representation.
func ProtoToContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(p *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToClusterNetworkConfig converts a ClusterNetworkConfig resource from its proto representation.
func ProtoToContainerBetaClusterNetworkConfig(p *betapb.ContainerBetaClusterNetworkConfig) *beta.ClusterNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNetworkConfig{
		Network:                   dcl.StringOrNil(p.Network),
		Subnetwork:                dcl.StringOrNil(p.Subnetwork),
		EnableIntraNodeVisibility: dcl.Bool(p.EnableIntraNodeVisibility),
		DefaultSnatStatus:         ProtoToContainerBetaClusterNetworkConfigDefaultSnatStatus(p.GetDefaultSnatStatus()),
		PrivateIPv6GoogleAccess:   ProtoToContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		DatapathProvider:          ProtoToContainerBetaClusterNetworkConfigDatapathProviderEnum(p.GetDatapathProvider()),
	}
	return obj
}

// ProtoToClusterNetworkConfigDefaultSnatStatus converts a ClusterNetworkConfigDefaultSnatStatus resource from its proto representation.
func ProtoToContainerBetaClusterNetworkConfigDefaultSnatStatus(p *betapb.ContainerBetaClusterNetworkConfigDefaultSnatStatus) *beta.ClusterNetworkConfigDefaultSnatStatus {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNetworkConfigDefaultSnatStatus{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterMaintenancePolicy converts a ClusterMaintenancePolicy resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicy(p *betapb.ContainerBetaClusterMaintenancePolicy) *beta.ClusterMaintenancePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicy{
		Window:          ProtoToContainerBetaClusterMaintenancePolicyWindow(p.GetWindow()),
		ResourceVersion: dcl.StringOrNil(p.ResourceVersion),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindow converts a ClusterMaintenancePolicyWindow resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicyWindow(p *betapb.ContainerBetaClusterMaintenancePolicyWindow) *beta.ClusterMaintenancePolicyWindow {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicyWindow{
		DailyMaintenanceWindow: ProtoToContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow(p.GetDailyMaintenanceWindow()),
		RecurringWindow:        ProtoToContainerBetaClusterMaintenancePolicyWindowRecurringWindow(p.GetRecurringWindow()),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowDailyMaintenanceWindow converts a ClusterMaintenancePolicyWindowDailyMaintenanceWindow resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow(p *betapb.ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow) *beta.ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicyWindowDailyMaintenanceWindow{
		StartTime: dcl.StringOrNil(p.GetStartTime()),
		Duration:  dcl.StringOrNil(p.Duration),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowRecurringWindow converts a ClusterMaintenancePolicyWindowRecurringWindow resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicyWindowRecurringWindow(p *betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindow) *beta.ClusterMaintenancePolicyWindowRecurringWindow {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicyWindowRecurringWindow{
		Window:     ProtoToContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow(p.GetWindow()),
		Recurrence: dcl.StringOrNil(p.Recurrence),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowRecurringWindowWindow converts a ClusterMaintenancePolicyWindowRecurringWindowWindow resource from its proto representation.
func ProtoToContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow(p *betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow) *beta.ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaintenancePolicyWindowRecurringWindowWindow{
		StartTime: dcl.StringOrNil(p.GetStartTime()),
		EndTime:   dcl.StringOrNil(p.GetEndTime()),
	}
	return obj
}

// ProtoToClusterDefaultMaxPodsConstraint converts a ClusterDefaultMaxPodsConstraint resource from its proto representation.
func ProtoToContainerBetaClusterDefaultMaxPodsConstraint(p *betapb.ContainerBetaClusterDefaultMaxPodsConstraint) *beta.ClusterDefaultMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterDefaultMaxPodsConstraint{
		MaxPodsPerNode: dcl.StringOrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfig converts a ClusterResourceUsageExportConfig resource from its proto representation.
func ProtoToContainerBetaClusterResourceUsageExportConfig(p *betapb.ContainerBetaClusterResourceUsageExportConfig) *beta.ClusterResourceUsageExportConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterResourceUsageExportConfig{
		BigqueryDestination:           ProtoToContainerBetaClusterResourceUsageExportConfigBigqueryDestination(p.GetBigqueryDestination()),
		EnableNetworkEgressMonitoring: dcl.Bool(p.EnableNetworkEgressMonitoring),
		ConsumptionMeteringConfig:     ProtoToContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig(p.GetConsumptionMeteringConfig()),
		EnableNetworkEgressMetering:   dcl.Bool(p.EnableNetworkEgressMetering),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfigBigqueryDestination converts a ClusterResourceUsageExportConfigBigqueryDestination resource from its proto representation.
func ProtoToContainerBetaClusterResourceUsageExportConfigBigqueryDestination(p *betapb.ContainerBetaClusterResourceUsageExportConfigBigqueryDestination) *beta.ClusterResourceUsageExportConfigBigqueryDestination {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterResourceUsageExportConfigBigqueryDestination{
		DatasetId: dcl.StringOrNil(p.DatasetId),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfigConsumptionMeteringConfig converts a ClusterResourceUsageExportConfigConsumptionMeteringConfig resource from its proto representation.
func ProtoToContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig(p *betapb.ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig) *beta.ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterResourceUsageExportConfigConsumptionMeteringConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAuthenticatorGroupsConfig converts a ClusterAuthenticatorGroupsConfig resource from its proto representation.
func ProtoToContainerBetaClusterAuthenticatorGroupsConfig(p *betapb.ContainerBetaClusterAuthenticatorGroupsConfig) *beta.ClusterAuthenticatorGroupsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAuthenticatorGroupsConfig{
		Enabled:       dcl.Bool(p.Enabled),
		SecurityGroup: dcl.StringOrNil(p.SecurityGroup),
	}
	return obj
}

// ProtoToClusterPrivateClusterConfig converts a ClusterPrivateClusterConfig resource from its proto representation.
func ProtoToContainerBetaClusterPrivateClusterConfig(p *betapb.ContainerBetaClusterPrivateClusterConfig) *beta.ClusterPrivateClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterPrivateClusterConfig{
		EnablePrivateNodes:       dcl.Bool(p.EnablePrivateNodes),
		EnablePrivateEndpoint:    dcl.Bool(p.EnablePrivateEndpoint),
		MasterIPv4CidrBlock:      dcl.StringOrNil(p.MasterIpv4CidrBlock),
		PrivateEndpoint:          dcl.StringOrNil(p.PrivateEndpoint),
		PublicEndpoint:           dcl.StringOrNil(p.PublicEndpoint),
		PeeringName:              dcl.StringOrNil(p.PeeringName),
		MasterGlobalAccessConfig: ProtoToContainerBetaClusterPrivateClusterConfigMasterGlobalAccessConfig(p.GetMasterGlobalAccessConfig()),
	}
	return obj
}

// ProtoToClusterPrivateClusterConfigMasterGlobalAccessConfig converts a ClusterPrivateClusterConfigMasterGlobalAccessConfig resource from its proto representation.
func ProtoToContainerBetaClusterPrivateClusterConfigMasterGlobalAccessConfig(p *betapb.ContainerBetaClusterPrivateClusterConfigMasterGlobalAccessConfig) *beta.ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterPrivateClusterConfigMasterGlobalAccessConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterDatabaseEncryption converts a ClusterDatabaseEncryption resource from its proto representation.
func ProtoToContainerBetaClusterDatabaseEncryption(p *betapb.ContainerBetaClusterDatabaseEncryption) *beta.ClusterDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterDatabaseEncryption{
		State:   ProtoToContainerBetaClusterDatabaseEncryptionStateEnum(p.GetState()),
		KeyName: dcl.StringOrNil(p.KeyName),
	}
	return obj
}

// ProtoToClusterVerticalPodAutoscaling converts a ClusterVerticalPodAutoscaling resource from its proto representation.
func ProtoToContainerBetaClusterVerticalPodAutoscaling(p *betapb.ContainerBetaClusterVerticalPodAutoscaling) *beta.ClusterVerticalPodAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterVerticalPodAutoscaling{
		Enabled:                    dcl.Bool(p.Enabled),
		EnableExperimentalFeatures: dcl.Bool(p.EnableExperimentalFeatures),
	}
	return obj
}

// ProtoToClusterShieldedNodes converts a ClusterShieldedNodes resource from its proto representation.
func ProtoToContainerBetaClusterShieldedNodes(p *betapb.ContainerBetaClusterShieldedNodes) *beta.ClusterShieldedNodes {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterShieldedNodes{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterConditions converts a ClusterConditions resource from its proto representation.
func ProtoToContainerBetaClusterConditions(p *betapb.ContainerBetaClusterConditions) *beta.ClusterConditions {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConditions{
		Code:          dcl.StringOrNil(p.Code),
		Message:       dcl.StringOrNil(p.Message),
		CanonicalCode: ProtoToContainerBetaClusterConditionsCanonicalCodeEnum(p.GetCanonicalCode()),
	}
	return obj
}

// ProtoToClusterAutopilot converts a ClusterAutopilot resource from its proto representation.
func ProtoToContainerBetaClusterAutopilot(p *betapb.ContainerBetaClusterAutopilot) *beta.ClusterAutopilot {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAutopilot{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterNodeConfig converts a ClusterNodeConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfig(p *betapb.ContainerBetaClusterNodeConfig) *beta.ClusterNodeConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfig{
		MachineType:            dcl.StringOrNil(p.MachineType),
		DiskSizeGb:             dcl.Int64OrNil(p.DiskSizeGb),
		ServiceAccount:         dcl.StringOrNil(p.ServiceAccount),
		ImageType:              dcl.StringOrNil(p.ImageType),
		LocalSsdCount:          dcl.Int64OrNil(p.LocalSsdCount),
		Preemptible:            dcl.Bool(p.Preemptible),
		DiskType:               dcl.StringOrNil(p.DiskType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		WorkloadMetadataConfig: ProtoToContainerBetaClusterNodeConfigWorkloadMetadataConfig(p.GetWorkloadMetadataConfig()),
		SandboxConfig:          ProtoToContainerBetaClusterNodeConfigSandboxConfig(p.GetSandboxConfig()),
		NodeGroup:              dcl.StringOrNil(p.NodeGroup),
		ReservationAffinity:    ProtoToContainerBetaClusterNodeConfigReservationAffinity(p.GetReservationAffinity()),
		ShieldedInstanceConfig: ProtoToContainerBetaClusterNodeConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		LinuxNodeConfig:        ProtoToContainerBetaClusterNodeConfigLinuxNodeConfig(p.GetLinuxNodeConfig()),
		KubeletConfig:          ProtoToContainerBetaClusterNodeConfigKubeletConfig(p.GetKubeletConfig()),
		BootDiskKmsKey:         dcl.StringOrNil(p.BootDiskKmsKey),
		EphemeralStorageConfig: ProtoToContainerBetaClusterNodeConfigEphemeralStorageConfig(p.GetEphemeralStorageConfig()),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToContainerBetaClusterNodeConfigAccelerators(r))
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerBetaClusterNodeConfigTaints(r))
	}
	return obj
}

// ProtoToClusterNodeConfigAccelerators converts a ClusterNodeConfigAccelerators resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigAccelerators(p *betapb.ContainerBetaClusterNodeConfigAccelerators) *beta.ClusterNodeConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToClusterNodeConfigWorkloadMetadataConfig converts a ClusterNodeConfigWorkloadMetadataConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigWorkloadMetadataConfig(p *betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfig) *beta.ClusterNodeConfigWorkloadMetadataConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigWorkloadMetadataConfig{
		Mode:         ProtoToContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum(p.GetMode()),
		NodeMetadata: ProtoToContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum(p.GetNodeMetadata()),
	}
	return obj
}

// ProtoToClusterNodeConfigTaints converts a ClusterNodeConfigTaints resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigTaints(p *betapb.ContainerBetaClusterNodeConfigTaints) *beta.ClusterNodeConfigTaints {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToContainerBetaClusterNodeConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToClusterNodeConfigSandboxConfig converts a ClusterNodeConfigSandboxConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigSandboxConfig(p *betapb.ContainerBetaClusterNodeConfigSandboxConfig) *beta.ClusterNodeConfigSandboxConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigSandboxConfig{
		Type:        ProtoToContainerBetaClusterNodeConfigSandboxConfigTypeEnum(p.GetType()),
		SandboxType: dcl.StringOrNil(p.SandboxType),
	}
	return obj
}

// ProtoToClusterNodeConfigReservationAffinity converts a ClusterNodeConfigReservationAffinity resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigReservationAffinity(p *betapb.ContainerBetaClusterNodeConfigReservationAffinity) *beta.ClusterNodeConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigReservationAffinity{
		ConsumeReservationType: ProtoToContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterNodeConfigShieldedInstanceConfig converts a ClusterNodeConfigShieldedInstanceConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigShieldedInstanceConfig(p *betapb.ContainerBetaClusterNodeConfigShieldedInstanceConfig) *beta.ClusterNodeConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToClusterNodeConfigLinuxNodeConfig converts a ClusterNodeConfigLinuxNodeConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigLinuxNodeConfig(p *betapb.ContainerBetaClusterNodeConfigLinuxNodeConfig) *beta.ClusterNodeConfigLinuxNodeConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigLinuxNodeConfig{}
	return obj
}

// ProtoToClusterNodeConfigKubeletConfig converts a ClusterNodeConfigKubeletConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigKubeletConfig(p *betapb.ContainerBetaClusterNodeConfigKubeletConfig) *beta.ClusterNodeConfigKubeletConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigKubeletConfig{
		CpuManagerPolicy:  dcl.StringOrNil(p.CpuManagerPolicy),
		CpuCfsQuota:       dcl.Bool(p.CpuCfsQuota),
		CpuCfsQuotaPeriod: dcl.StringOrNil(p.CpuCfsQuotaPeriod),
	}
	return obj
}

// ProtoToClusterNodeConfigEphemeralStorageConfig converts a ClusterNodeConfigEphemeralStorageConfig resource from its proto representation.
func ProtoToContainerBetaClusterNodeConfigEphemeralStorageConfig(p *betapb.ContainerBetaClusterNodeConfigEphemeralStorageConfig) *beta.ClusterNodeConfigEphemeralStorageConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNodeConfigEphemeralStorageConfig{
		LocalSsdCount: dcl.Int64OrNil(p.LocalSsdCount),
	}
	return obj
}

// ProtoToClusterReleaseChannel converts a ClusterReleaseChannel resource from its proto representation.
func ProtoToContainerBetaClusterReleaseChannel(p *betapb.ContainerBetaClusterReleaseChannel) *beta.ClusterReleaseChannel {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterReleaseChannel{
		Channel: ProtoToContainerBetaClusterReleaseChannelChannelEnum(p.GetChannel()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToContainerBetaClusterWorkloadIdentityConfig(p *betapb.ContainerBetaClusterWorkloadIdentityConfig) *beta.ClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterWorkloadIdentityConfig{
		WorkloadPool:      dcl.StringOrNil(p.WorkloadPool),
		IdentityNamespace: dcl.StringOrNil(p.IdentityNamespace),
		IdentityProvider:  dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToClusterNotificationConfig converts a ClusterNotificationConfig resource from its proto representation.
func ProtoToContainerBetaClusterNotificationConfig(p *betapb.ContainerBetaClusterNotificationConfig) *beta.ClusterNotificationConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNotificationConfig{
		Pubsub: ProtoToContainerBetaClusterNotificationConfigPubsub(p.GetPubsub()),
	}
	return obj
}

// ProtoToClusterNotificationConfigPubsub converts a ClusterNotificationConfigPubsub resource from its proto representation.
func ProtoToContainerBetaClusterNotificationConfigPubsub(p *betapb.ContainerBetaClusterNotificationConfigPubsub) *beta.ClusterNotificationConfigPubsub {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNotificationConfigPubsub{
		Enabled: dcl.Bool(p.Enabled),
		Topic:   dcl.StringOrNil(p.Topic),
	}
	return obj
}

// ProtoToClusterConfidentialNodes converts a ClusterConfidentialNodes resource from its proto representation.
func ProtoToContainerBetaClusterConfidentialNodes(p *betapb.ContainerBetaClusterConfidentialNodes) *beta.ClusterConfidentialNodes {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterConfidentialNodes{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterPodSecurityPolicyConfig converts a ClusterPodSecurityPolicyConfig resource from its proto representation.
func ProtoToContainerBetaClusterPodSecurityPolicyConfig(p *betapb.ContainerBetaClusterPodSecurityPolicyConfig) *beta.ClusterPodSecurityPolicyConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterPodSecurityPolicyConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterClusterTelemetry converts a ClusterClusterTelemetry resource from its proto representation.
func ProtoToContainerBetaClusterClusterTelemetry(p *betapb.ContainerBetaClusterClusterTelemetry) *beta.ClusterClusterTelemetry {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterClusterTelemetry{
		Type: ProtoToContainerBetaClusterClusterTelemetryTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToClusterTPUConfig converts a ClusterTPUConfig resource from its proto representation.
func ProtoToContainerBetaClusterTPUConfig(p *betapb.ContainerBetaClusterTPUConfig) *beta.ClusterTPUConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterTPUConfig{
		Enabled:              dcl.Bool(p.Enabled),
		UseServiceNetworking: dcl.Bool(p.UseServiceNetworking),
		IPv4CidrBlock:        dcl.StringOrNil(p.Ipv4CidrBlock),
	}
	return obj
}

// ProtoToClusterMaster converts a ClusterMaster resource from its proto representation.
func ProtoToContainerBetaClusterMaster(p *betapb.ContainerBetaClusterMaster) *beta.ClusterMaster {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMaster{}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *betapb.ContainerBetaCluster) *beta.Cluster {
	obj := &beta.Cluster{
		Name:                           dcl.StringOrNil(p.Name),
		Description:                    dcl.StringOrNil(p.Description),
		InitialNodeCount:               dcl.Int64OrNil(p.InitialNodeCount),
		MasterAuth:                     ProtoToContainerBetaClusterMasterAuth(p.GetMasterAuth()),
		LoggingService:                 dcl.StringOrNil(p.LoggingService),
		MonitoringService:              dcl.StringOrNil(p.MonitoringService),
		Network:                        dcl.StringOrNil(p.Network),
		ClusterIPv4Cidr:                dcl.StringOrNil(p.ClusterIpv4Cidr),
		AddonsConfig:                   ProtoToContainerBetaClusterAddonsConfig(p.GetAddonsConfig()),
		Subnetwork:                     dcl.StringOrNil(p.Subnetwork),
		EnableKubernetesAlpha:          dcl.Bool(p.EnableKubernetesAlpha),
		LabelFingerprint:               dcl.StringOrNil(p.LabelFingerprint),
		LegacyAbac:                     ProtoToContainerBetaClusterLegacyAbac(p.GetLegacyAbac()),
		NetworkPolicy:                  ProtoToContainerBetaClusterNetworkPolicy(p.GetNetworkPolicy()),
		IPAllocationPolicy:             ProtoToContainerBetaClusterIPAllocationPolicy(p.GetIpAllocationPolicy()),
		MasterAuthorizedNetworksConfig: ProtoToContainerBetaClusterMasterAuthorizedNetworksConfig(p.GetMasterAuthorizedNetworksConfig()),
		BinaryAuthorization:            ProtoToContainerBetaClusterBinaryAuthorization(p.GetBinaryAuthorization()),
		Autoscaling:                    ProtoToContainerBetaClusterAutoscaling(p.GetAutoscaling()),
		NetworkConfig:                  ProtoToContainerBetaClusterNetworkConfig(p.GetNetworkConfig()),
		MaintenancePolicy:              ProtoToContainerBetaClusterMaintenancePolicy(p.GetMaintenancePolicy()),
		DefaultMaxPodsConstraint:       ProtoToContainerBetaClusterDefaultMaxPodsConstraint(p.GetDefaultMaxPodsConstraint()),
		ResourceUsageExportConfig:      ProtoToContainerBetaClusterResourceUsageExportConfig(p.GetResourceUsageExportConfig()),
		AuthenticatorGroupsConfig:      ProtoToContainerBetaClusterAuthenticatorGroupsConfig(p.GetAuthenticatorGroupsConfig()),
		PrivateClusterConfig:           ProtoToContainerBetaClusterPrivateClusterConfig(p.GetPrivateClusterConfig()),
		DatabaseEncryption:             ProtoToContainerBetaClusterDatabaseEncryption(p.GetDatabaseEncryption()),
		VerticalPodAutoscaling:         ProtoToContainerBetaClusterVerticalPodAutoscaling(p.GetVerticalPodAutoscaling()),
		ShieldedNodes:                  ProtoToContainerBetaClusterShieldedNodes(p.GetShieldedNodes()),
		Endpoint:                       dcl.StringOrNil(p.Endpoint),
		MasterVersion:                  dcl.StringOrNil(p.MasterVersion),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		Status:                         dcl.StringOrNil(p.Status),
		StatusMessage:                  dcl.StringOrNil(p.StatusMessage),
		NodeIPv4CidrSize:               dcl.Int64OrNil(p.NodeIpv4CidrSize),
		ServicesIPv4Cidr:               dcl.StringOrNil(p.ServicesIpv4Cidr),
		ExpireTime:                     dcl.StringOrNil(p.GetExpireTime()),
		Location:                       dcl.StringOrNil(p.Location),
		EnableTPU:                      dcl.Bool(p.EnableTpu),
		TPUIPv4CidrBlock:               dcl.StringOrNil(p.TpuIpv4CidrBlock),
		Autopilot:                      ProtoToContainerBetaClusterAutopilot(p.GetAutopilot()),
		Project:                        dcl.StringOrNil(p.Project),
		NodeConfig:                     ProtoToContainerBetaClusterNodeConfig(p.GetNodeConfig()),
		ReleaseChannel:                 ProtoToContainerBetaClusterReleaseChannel(p.GetReleaseChannel()),
		WorkloadIdentityConfig:         ProtoToContainerBetaClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		NotificationConfig:             ProtoToContainerBetaClusterNotificationConfig(p.GetNotificationConfig()),
		ConfidentialNodes:              ProtoToContainerBetaClusterConfidentialNodes(p.GetConfidentialNodes()),
		SelfLink:                       dcl.StringOrNil(p.SelfLink),
		Zone:                           dcl.StringOrNil(p.Zone),
		InitialClusterVersion:          dcl.StringOrNil(p.InitialClusterVersion),
		CurrentMasterVersion:           dcl.StringOrNil(p.CurrentMasterVersion),
		CurrentNodeVersion:             dcl.StringOrNil(p.CurrentNodeVersion),
		CurrentNodeCount:               dcl.Int64OrNil(p.CurrentNodeCount),
		Id:                             dcl.StringOrNil(p.Id),
		PodSecurityPolicyConfig:        ProtoToContainerBetaClusterPodSecurityPolicyConfig(p.GetPodSecurityPolicyConfig()),
		PrivateCluster:                 dcl.Bool(p.PrivateCluster),
		MasterIPv4CidrBlock:            dcl.StringOrNil(p.MasterIpv4CidrBlock),
		ClusterTelemetry:               ProtoToContainerBetaClusterClusterTelemetry(p.GetClusterTelemetry()),
		TPUConfig:                      ProtoToContainerBetaClusterTPUConfig(p.GetTpuConfig()),
		Master:                         ProtoToContainerBetaClusterMaster(p.GetMaster()),
	}
	for _, r := range p.GetNodePools() {
		obj.NodePools = append(obj.NodePools, *ProtoToContainerBetaClusterNodePools(r))
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToContainerBetaClusterConditions(r))
	}
	for _, r := range p.GetInstanceGroupUrls() {
		obj.InstanceGroupUrls = append(obj.InstanceGroupUrls, r)
	}
	return obj
}

// ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumToProto converts a ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum enum to its proto representation.
func ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumToProto(e *beta.ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum) betapb.ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum_value["ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(v)
	}
	return betapb.ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(0)
}

// ClusterAddonsConfigIstioConfigAuthEnumToProto converts a ClusterAddonsConfigIstioConfigAuthEnum enum to its proto representation.
func ContainerBetaClusterAddonsConfigIstioConfigAuthEnumToProto(e *beta.ClusterAddonsConfigIstioConfigAuthEnum) betapb.ContainerBetaClusterAddonsConfigIstioConfigAuthEnum {
	if e == nil {
		return betapb.ContainerBetaClusterAddonsConfigIstioConfigAuthEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterAddonsConfigIstioConfigAuthEnum_value["ClusterAddonsConfigIstioConfigAuthEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterAddonsConfigIstioConfigAuthEnum(v)
	}
	return betapb.ContainerBetaClusterAddonsConfigIstioConfigAuthEnum(0)
}

// ClusterNodePoolsConfigWorkloadMetadataConfigModeEnumToProto converts a ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum enum to its proto representation.
func ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnumToProto(e *beta.ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum) betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum_value["ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(v)
	}
	return betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(0)
}

// ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnumToProto converts a ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum enum to its proto representation.
func ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnumToProto(e *beta.ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum) betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum_value["ClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum(v)
	}
	return betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnum(0)
}

// ClusterNodePoolsConfigTaintsEffectEnumToProto converts a ClusterNodePoolsConfigTaintsEffectEnum enum to its proto representation.
func ContainerBetaClusterNodePoolsConfigTaintsEffectEnumToProto(e *beta.ClusterNodePoolsConfigTaintsEffectEnum) betapb.ContainerBetaClusterNodePoolsConfigTaintsEffectEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodePoolsConfigTaintsEffectEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodePoolsConfigTaintsEffectEnum_value["ClusterNodePoolsConfigTaintsEffectEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodePoolsConfigTaintsEffectEnum(v)
	}
	return betapb.ContainerBetaClusterNodePoolsConfigTaintsEffectEnum(0)
}

// ClusterNodePoolsConfigSandboxConfigTypeEnumToProto converts a ClusterNodePoolsConfigSandboxConfigTypeEnum enum to its proto representation.
func ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnumToProto(e *beta.ClusterNodePoolsConfigSandboxConfigTypeEnum) betapb.ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum_value["ClusterNodePoolsConfigSandboxConfigTypeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum(v)
	}
	return betapb.ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnum(0)
}

// ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumToProto(e *beta.ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum) betapb.ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return betapb.ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterNodePoolsStatusEnumToProto converts a ClusterNodePoolsStatusEnum enum to its proto representation.
func ContainerBetaClusterNodePoolsStatusEnumToProto(e *beta.ClusterNodePoolsStatusEnum) betapb.ContainerBetaClusterNodePoolsStatusEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodePoolsStatusEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodePoolsStatusEnum_value["ClusterNodePoolsStatusEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodePoolsStatusEnum(v)
	}
	return betapb.ContainerBetaClusterNodePoolsStatusEnum(0)
}

// ClusterNodePoolsConditionsCodeEnumToProto converts a ClusterNodePoolsConditionsCodeEnum enum to its proto representation.
func ContainerBetaClusterNodePoolsConditionsCodeEnumToProto(e *beta.ClusterNodePoolsConditionsCodeEnum) betapb.ContainerBetaClusterNodePoolsConditionsCodeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodePoolsConditionsCodeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodePoolsConditionsCodeEnum_value["ClusterNodePoolsConditionsCodeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodePoolsConditionsCodeEnum(v)
	}
	return betapb.ContainerBetaClusterNodePoolsConditionsCodeEnum(0)
}

// ClusterNodePoolsConditionsCanonicalCodeEnumToProto converts a ClusterNodePoolsConditionsCanonicalCodeEnum enum to its proto representation.
func ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnumToProto(e *beta.ClusterNodePoolsConditionsCanonicalCodeEnum) betapb.ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum_value["ClusterNodePoolsConditionsCanonicalCodeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum(v)
	}
	return betapb.ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnum(0)
}

// ClusterNetworkPolicyProviderEnumToProto converts a ClusterNetworkPolicyProviderEnum enum to its proto representation.
func ContainerBetaClusterNetworkPolicyProviderEnumToProto(e *beta.ClusterNetworkPolicyProviderEnum) betapb.ContainerBetaClusterNetworkPolicyProviderEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNetworkPolicyProviderEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNetworkPolicyProviderEnum_value["ClusterNetworkPolicyProviderEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNetworkPolicyProviderEnum(v)
	}
	return betapb.ContainerBetaClusterNetworkPolicyProviderEnum(0)
}

// ClusterAutoscalingAutoscalingProfileEnumToProto converts a ClusterAutoscalingAutoscalingProfileEnum enum to its proto representation.
func ContainerBetaClusterAutoscalingAutoscalingProfileEnumToProto(e *beta.ClusterAutoscalingAutoscalingProfileEnum) betapb.ContainerBetaClusterAutoscalingAutoscalingProfileEnum {
	if e == nil {
		return betapb.ContainerBetaClusterAutoscalingAutoscalingProfileEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterAutoscalingAutoscalingProfileEnum_value["ClusterAutoscalingAutoscalingProfileEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterAutoscalingAutoscalingProfileEnum(v)
	}
	return betapb.ContainerBetaClusterAutoscalingAutoscalingProfileEnum(0)
}

// ClusterNetworkConfigPrivateIPv6GoogleAccessEnumToProto converts a ClusterNetworkConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnumToProto(e *beta.ClusterNetworkConfigPrivateIPv6GoogleAccessEnum) betapb.ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum_value["ClusterNetworkConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return betapb.ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnum(0)
}

// ClusterNetworkConfigDatapathProviderEnumToProto converts a ClusterNetworkConfigDatapathProviderEnum enum to its proto representation.
func ContainerBetaClusterNetworkConfigDatapathProviderEnumToProto(e *beta.ClusterNetworkConfigDatapathProviderEnum) betapb.ContainerBetaClusterNetworkConfigDatapathProviderEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNetworkConfigDatapathProviderEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNetworkConfigDatapathProviderEnum_value["ClusterNetworkConfigDatapathProviderEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNetworkConfigDatapathProviderEnum(v)
	}
	return betapb.ContainerBetaClusterNetworkConfigDatapathProviderEnum(0)
}

// ClusterDatabaseEncryptionStateEnumToProto converts a ClusterDatabaseEncryptionStateEnum enum to its proto representation.
func ContainerBetaClusterDatabaseEncryptionStateEnumToProto(e *beta.ClusterDatabaseEncryptionStateEnum) betapb.ContainerBetaClusterDatabaseEncryptionStateEnum {
	if e == nil {
		return betapb.ContainerBetaClusterDatabaseEncryptionStateEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterDatabaseEncryptionStateEnum_value["ClusterDatabaseEncryptionStateEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterDatabaseEncryptionStateEnum(v)
	}
	return betapb.ContainerBetaClusterDatabaseEncryptionStateEnum(0)
}

// ClusterConditionsCanonicalCodeEnumToProto converts a ClusterConditionsCanonicalCodeEnum enum to its proto representation.
func ContainerBetaClusterConditionsCanonicalCodeEnumToProto(e *beta.ClusterConditionsCanonicalCodeEnum) betapb.ContainerBetaClusterConditionsCanonicalCodeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterConditionsCanonicalCodeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterConditionsCanonicalCodeEnum_value["ClusterConditionsCanonicalCodeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterConditionsCanonicalCodeEnum(v)
	}
	return betapb.ContainerBetaClusterConditionsCanonicalCodeEnum(0)
}

// ClusterNodeConfigWorkloadMetadataConfigModeEnumToProto converts a ClusterNodeConfigWorkloadMetadataConfigModeEnum enum to its proto representation.
func ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnumToProto(e *beta.ClusterNodeConfigWorkloadMetadataConfigModeEnum) betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum_value["ClusterNodeConfigWorkloadMetadataConfigModeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum(v)
	}
	return betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnum(0)
}

// ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnumToProto converts a ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum enum to its proto representation.
func ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnumToProto(e *beta.ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum) betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum_value["ClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum(v)
	}
	return betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnum(0)
}

// ClusterNodeConfigTaintsEffectEnumToProto converts a ClusterNodeConfigTaintsEffectEnum enum to its proto representation.
func ContainerBetaClusterNodeConfigTaintsEffectEnumToProto(e *beta.ClusterNodeConfigTaintsEffectEnum) betapb.ContainerBetaClusterNodeConfigTaintsEffectEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodeConfigTaintsEffectEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodeConfigTaintsEffectEnum_value["ClusterNodeConfigTaintsEffectEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodeConfigTaintsEffectEnum(v)
	}
	return betapb.ContainerBetaClusterNodeConfigTaintsEffectEnum(0)
}

// ClusterNodeConfigSandboxConfigTypeEnumToProto converts a ClusterNodeConfigSandboxConfigTypeEnum enum to its proto representation.
func ContainerBetaClusterNodeConfigSandboxConfigTypeEnumToProto(e *beta.ClusterNodeConfigSandboxConfigTypeEnum) betapb.ContainerBetaClusterNodeConfigSandboxConfigTypeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodeConfigSandboxConfigTypeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodeConfigSandboxConfigTypeEnum_value["ClusterNodeConfigSandboxConfigTypeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodeConfigSandboxConfigTypeEnum(v)
	}
	return betapb.ContainerBetaClusterNodeConfigSandboxConfigTypeEnum(0)
}

// ClusterNodeConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnumToProto(e *beta.ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum) betapb.ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return betapb.ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterReleaseChannelChannelEnumToProto converts a ClusterReleaseChannelChannelEnum enum to its proto representation.
func ContainerBetaClusterReleaseChannelChannelEnumToProto(e *beta.ClusterReleaseChannelChannelEnum) betapb.ContainerBetaClusterReleaseChannelChannelEnum {
	if e == nil {
		return betapb.ContainerBetaClusterReleaseChannelChannelEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterReleaseChannelChannelEnum_value["ClusterReleaseChannelChannelEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterReleaseChannelChannelEnum(v)
	}
	return betapb.ContainerBetaClusterReleaseChannelChannelEnum(0)
}

// ClusterClusterTelemetryTypeEnumToProto converts a ClusterClusterTelemetryTypeEnum enum to its proto representation.
func ContainerBetaClusterClusterTelemetryTypeEnumToProto(e *beta.ClusterClusterTelemetryTypeEnum) betapb.ContainerBetaClusterClusterTelemetryTypeEnum {
	if e == nil {
		return betapb.ContainerBetaClusterClusterTelemetryTypeEnum(0)
	}
	if v, ok := betapb.ContainerBetaClusterClusterTelemetryTypeEnum_value["ClusterClusterTelemetryTypeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaClusterClusterTelemetryTypeEnum(v)
	}
	return betapb.ContainerBetaClusterClusterTelemetryTypeEnum(0)
}

// ClusterMasterAuthToProto converts a ClusterMasterAuth resource to its proto representation.
func ContainerBetaClusterMasterAuthToProto(o *beta.ClusterMasterAuth) *betapb.ContainerBetaClusterMasterAuth {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMasterAuth{
		Username:                dcl.ValueOrEmptyString(o.Username),
		Password:                dcl.ValueOrEmptyString(o.Password),
		ClientCertificateConfig: ContainerBetaClusterMasterAuthClientCertificateConfigToProto(o.ClientCertificateConfig),
		ClusterCaCertificate:    dcl.ValueOrEmptyString(o.ClusterCaCertificate),
		ClientCertificate:       dcl.ValueOrEmptyString(o.ClientCertificate),
		ClientKey:               dcl.ValueOrEmptyString(o.ClientKey),
	}
	return p
}

// ClusterMasterAuthClientCertificateConfigToProto converts a ClusterMasterAuthClientCertificateConfig resource to its proto representation.
func ContainerBetaClusterMasterAuthClientCertificateConfigToProto(o *beta.ClusterMasterAuthClientCertificateConfig) *betapb.ContainerBetaClusterMasterAuthClientCertificateConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMasterAuthClientCertificateConfig{
		IssueClientCertificate: dcl.ValueOrEmptyBool(o.IssueClientCertificate),
	}
	return p
}

// ClusterAddonsConfigToProto converts a ClusterAddonsConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigToProto(o *beta.ClusterAddonsConfig) *betapb.ContainerBetaClusterAddonsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfig{
		HttpLoadBalancing:                ContainerBetaClusterAddonsConfigHttpLoadBalancingToProto(o.HttpLoadBalancing),
		HorizontalPodAutoscaling:         ContainerBetaClusterAddonsConfigHorizontalPodAutoscalingToProto(o.HorizontalPodAutoscaling),
		KubernetesDashboard:              ContainerBetaClusterAddonsConfigKubernetesDashboardToProto(o.KubernetesDashboard),
		NetworkPolicyConfig:              ContainerBetaClusterAddonsConfigNetworkPolicyConfigToProto(o.NetworkPolicyConfig),
		CloudRunConfig:                   ContainerBetaClusterAddonsConfigCloudRunConfigToProto(o.CloudRunConfig),
		DnsCacheConfig:                   ContainerBetaClusterAddonsConfigDnsCacheConfigToProto(o.DnsCacheConfig),
		ConfigConnectorConfig:            ContainerBetaClusterAddonsConfigConfigConnectorConfigToProto(o.ConfigConnectorConfig),
		GcePersistentDiskCsiDriverConfig: ContainerBetaClusterAddonsConfigGcePersistentDiskCsiDriverConfigToProto(o.GcePersistentDiskCsiDriverConfig),
		IstioConfig:                      ContainerBetaClusterAddonsConfigIstioConfigToProto(o.IstioConfig),
		KalmConfig:                       ContainerBetaClusterAddonsConfigKalmConfigToProto(o.KalmConfig),
	}
	return p
}

// ClusterAddonsConfigHttpLoadBalancingToProto converts a ClusterAddonsConfigHttpLoadBalancing resource to its proto representation.
func ContainerBetaClusterAddonsConfigHttpLoadBalancingToProto(o *beta.ClusterAddonsConfigHttpLoadBalancing) *betapb.ContainerBetaClusterAddonsConfigHttpLoadBalancing {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigHttpLoadBalancing{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigHorizontalPodAutoscalingToProto converts a ClusterAddonsConfigHorizontalPodAutoscaling resource to its proto representation.
func ContainerBetaClusterAddonsConfigHorizontalPodAutoscalingToProto(o *beta.ClusterAddonsConfigHorizontalPodAutoscaling) *betapb.ContainerBetaClusterAddonsConfigHorizontalPodAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigHorizontalPodAutoscaling{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigKubernetesDashboardToProto converts a ClusterAddonsConfigKubernetesDashboard resource to its proto representation.
func ContainerBetaClusterAddonsConfigKubernetesDashboardToProto(o *beta.ClusterAddonsConfigKubernetesDashboard) *betapb.ContainerBetaClusterAddonsConfigKubernetesDashboard {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigKubernetesDashboard{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigNetworkPolicyConfigToProto converts a ClusterAddonsConfigNetworkPolicyConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigNetworkPolicyConfigToProto(o *beta.ClusterAddonsConfigNetworkPolicyConfig) *betapb.ContainerBetaClusterAddonsConfigNetworkPolicyConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigNetworkPolicyConfig{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigCloudRunConfigToProto converts a ClusterAddonsConfigCloudRunConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigCloudRunConfigToProto(o *beta.ClusterAddonsConfigCloudRunConfig) *betapb.ContainerBetaClusterAddonsConfigCloudRunConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigCloudRunConfig{
		Disabled:         dcl.ValueOrEmptyBool(o.Disabled),
		LoadBalancerType: ContainerBetaClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumToProto(o.LoadBalancerType),
	}
	return p
}

// ClusterAddonsConfigDnsCacheConfigToProto converts a ClusterAddonsConfigDnsCacheConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigDnsCacheConfigToProto(o *beta.ClusterAddonsConfigDnsCacheConfig) *betapb.ContainerBetaClusterAddonsConfigDnsCacheConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigDnsCacheConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAddonsConfigConfigConnectorConfigToProto converts a ClusterAddonsConfigConfigConnectorConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigConfigConnectorConfigToProto(o *beta.ClusterAddonsConfigConfigConnectorConfig) *betapb.ContainerBetaClusterAddonsConfigConfigConnectorConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigConfigConnectorConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAddonsConfigGcePersistentDiskCsiDriverConfigToProto converts a ClusterAddonsConfigGcePersistentDiskCsiDriverConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigGcePersistentDiskCsiDriverConfigToProto(o *beta.ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) *betapb.ContainerBetaClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigGcePersistentDiskCsiDriverConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAddonsConfigIstioConfigToProto converts a ClusterAddonsConfigIstioConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigIstioConfigToProto(o *beta.ClusterAddonsConfigIstioConfig) *betapb.ContainerBetaClusterAddonsConfigIstioConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigIstioConfig{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
		Auth:     ContainerBetaClusterAddonsConfigIstioConfigAuthEnumToProto(o.Auth),
	}
	return p
}

// ClusterAddonsConfigKalmConfigToProto converts a ClusterAddonsConfigKalmConfig resource to its proto representation.
func ContainerBetaClusterAddonsConfigKalmConfigToProto(o *beta.ClusterAddonsConfigKalmConfig) *betapb.ContainerBetaClusterAddonsConfigKalmConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAddonsConfigKalmConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterNodePoolsToProto converts a ClusterNodePools resource to its proto representation.
func ContainerBetaClusterNodePoolsToProto(o *beta.ClusterNodePools) *betapb.ContainerBetaClusterNodePools {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePools{
		Name:              dcl.ValueOrEmptyString(o.Name),
		Config:            ContainerBetaClusterNodePoolsConfigToProto(o.Config),
		InitialNodeCount:  dcl.ValueOrEmptyInt64(o.InitialNodeCount),
		SelfLink:          dcl.ValueOrEmptyString(o.SelfLink),
		Version:           dcl.ValueOrEmptyString(o.Version),
		Status:            ContainerBetaClusterNodePoolsStatusEnumToProto(o.Status),
		StatusMessage:     dcl.ValueOrEmptyString(o.StatusMessage),
		Autoscaling:       ContainerBetaClusterNodePoolsAutoscalingToProto(o.Autoscaling),
		Management:        ContainerBetaClusterNodePoolsManagementToProto(o.Management),
		MaxPodsConstraint: ContainerBetaClusterNodePoolsMaxPodsConstraintToProto(o.MaxPodsConstraint),
		PodIpv4CidrSize:   dcl.ValueOrEmptyInt64(o.PodIPv4CidrSize),
		UpgradeSettings:   ContainerBetaClusterNodePoolsUpgradeSettingsToProto(o.UpgradeSettings),
		NetworkConfig:     ContainerBetaClusterNodePoolsNetworkConfigToProto(o.NetworkConfig),
	}
	for _, r := range o.Locations {
		p.Locations = append(p.Locations, r)
	}
	for _, r := range o.InstanceGroupUrls {
		p.InstanceGroupUrls = append(p.InstanceGroupUrls, r)
	}
	for _, r := range o.Conditions {
		p.Conditions = append(p.Conditions, ContainerBetaClusterNodePoolsConditionsToProto(&r))
	}
	return p
}

// ClusterNodePoolsConfigToProto converts a ClusterNodePoolsConfig resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigToProto(o *beta.ClusterNodePoolsConfig) *betapb.ContainerBetaClusterNodePoolsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfig{
		MachineType:            dcl.ValueOrEmptyString(o.MachineType),
		DiskSizeGb:             dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		ServiceAccount:         dcl.ValueOrEmptyString(o.ServiceAccount),
		ImageType:              dcl.ValueOrEmptyString(o.ImageType),
		LocalSsdCount:          dcl.ValueOrEmptyInt64(o.LocalSsdCount),
		Preemptible:            dcl.ValueOrEmptyBool(o.Preemptible),
		DiskType:               dcl.ValueOrEmptyString(o.DiskType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		WorkloadMetadataConfig: ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigToProto(o.WorkloadMetadataConfig),
		SandboxConfig:          ContainerBetaClusterNodePoolsConfigSandboxConfigToProto(o.SandboxConfig),
		NodeGroup:              dcl.ValueOrEmptyString(o.NodeGroup),
		ReservationAffinity:    ContainerBetaClusterNodePoolsConfigReservationAffinityToProto(o.ReservationAffinity),
		ShieldedInstanceConfig: ContainerBetaClusterNodePoolsConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
		LinuxNodeConfig:        ContainerBetaClusterNodePoolsConfigLinuxNodeConfigToProto(o.LinuxNodeConfig),
		KubeletConfig:          ContainerBetaClusterNodePoolsConfigKubeletConfigToProto(o.KubeletConfig),
		BootDiskKmsKey:         dcl.ValueOrEmptyString(o.BootDiskKmsKey),
		EphemeralStorageConfig: ContainerBetaClusterNodePoolsConfigEphemeralStorageConfigToProto(o.EphemeralStorageConfig),
	}
	for _, r := range o.OAuthScopes {
		p.OauthScopes = append(p.OauthScopes, r)
	}
	p.Metadata = make(map[string]string)
	for k, r := range o.Metadata {
		p.Metadata[k] = r
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	for _, r := range o.Accelerators {
		p.Accelerators = append(p.Accelerators, ContainerBetaClusterNodePoolsConfigAcceleratorsToProto(&r))
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, ContainerBetaClusterNodePoolsConfigTaintsToProto(&r))
	}
	return p
}

// ClusterNodePoolsConfigAcceleratorsToProto converts a ClusterNodePoolsConfigAccelerators resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigAcceleratorsToProto(o *beta.ClusterNodePoolsConfigAccelerators) *betapb.ContainerBetaClusterNodePoolsConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// ClusterNodePoolsConfigWorkloadMetadataConfigToProto converts a ClusterNodePoolsConfigWorkloadMetadataConfig resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigToProto(o *beta.ClusterNodePoolsConfigWorkloadMetadataConfig) *betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfig{
		Mode:         ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigModeEnumToProto(o.Mode),
		NodeMetadata: ContainerBetaClusterNodePoolsConfigWorkloadMetadataConfigNodeMetadataEnumToProto(o.NodeMetadata),
	}
	return p
}

// ClusterNodePoolsConfigTaintsToProto converts a ClusterNodePoolsConfigTaints resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigTaintsToProto(o *beta.ClusterNodePoolsConfigTaints) *betapb.ContainerBetaClusterNodePoolsConfigTaints {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: ContainerBetaClusterNodePoolsConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// ClusterNodePoolsConfigSandboxConfigToProto converts a ClusterNodePoolsConfigSandboxConfig resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigSandboxConfigToProto(o *beta.ClusterNodePoolsConfigSandboxConfig) *betapb.ContainerBetaClusterNodePoolsConfigSandboxConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigSandboxConfig{
		Type:        ContainerBetaClusterNodePoolsConfigSandboxConfigTypeEnumToProto(o.Type),
		SandboxType: dcl.ValueOrEmptyString(o.SandboxType),
	}
	return p
}

// ClusterNodePoolsConfigReservationAffinityToProto converts a ClusterNodePoolsConfigReservationAffinity resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigReservationAffinityToProto(o *beta.ClusterNodePoolsConfigReservationAffinity) *betapb.ContainerBetaClusterNodePoolsConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigReservationAffinity{
		ConsumeReservationType: ContainerBetaClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// ClusterNodePoolsConfigShieldedInstanceConfigToProto converts a ClusterNodePoolsConfigShieldedInstanceConfig resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigShieldedInstanceConfigToProto(o *beta.ClusterNodePoolsConfigShieldedInstanceConfig) *betapb.ContainerBetaClusterNodePoolsConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// ClusterNodePoolsConfigLinuxNodeConfigToProto converts a ClusterNodePoolsConfigLinuxNodeConfig resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigLinuxNodeConfigToProto(o *beta.ClusterNodePoolsConfigLinuxNodeConfig) *betapb.ContainerBetaClusterNodePoolsConfigLinuxNodeConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigLinuxNodeConfig{}
	p.Sysctls = make(map[string]string)
	for k, r := range o.Sysctls {
		p.Sysctls[k] = r
	}
	return p
}

// ClusterNodePoolsConfigKubeletConfigToProto converts a ClusterNodePoolsConfigKubeletConfig resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigKubeletConfigToProto(o *beta.ClusterNodePoolsConfigKubeletConfig) *betapb.ContainerBetaClusterNodePoolsConfigKubeletConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigKubeletConfig{
		CpuManagerPolicy:  dcl.ValueOrEmptyString(o.CpuManagerPolicy),
		CpuCfsQuota:       dcl.ValueOrEmptyBool(o.CpuCfsQuota),
		CpuCfsQuotaPeriod: dcl.ValueOrEmptyString(o.CpuCfsQuotaPeriod),
	}
	return p
}

// ClusterNodePoolsConfigEphemeralStorageConfigToProto converts a ClusterNodePoolsConfigEphemeralStorageConfig resource to its proto representation.
func ContainerBetaClusterNodePoolsConfigEphemeralStorageConfigToProto(o *beta.ClusterNodePoolsConfigEphemeralStorageConfig) *betapb.ContainerBetaClusterNodePoolsConfigEphemeralStorageConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConfigEphemeralStorageConfig{
		LocalSsdCount: dcl.ValueOrEmptyInt64(o.LocalSsdCount),
	}
	return p
}

// ClusterNodePoolsAutoscalingToProto converts a ClusterNodePoolsAutoscaling resource to its proto representation.
func ContainerBetaClusterNodePoolsAutoscalingToProto(o *beta.ClusterNodePoolsAutoscaling) *betapb.ContainerBetaClusterNodePoolsAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsAutoscaling{
		Enabled:         dcl.ValueOrEmptyBool(o.Enabled),
		MinNodeCount:    dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount:    dcl.ValueOrEmptyInt64(o.MaxNodeCount),
		Autoprovisioned: dcl.ValueOrEmptyBool(o.Autoprovisioned),
	}
	return p
}

// ClusterNodePoolsManagementToProto converts a ClusterNodePoolsManagement resource to its proto representation.
func ContainerBetaClusterNodePoolsManagementToProto(o *beta.ClusterNodePoolsManagement) *betapb.ContainerBetaClusterNodePoolsManagement {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsManagement{
		AutoUpgrade:    dcl.ValueOrEmptyBool(o.AutoUpgrade),
		AutoRepair:     dcl.ValueOrEmptyBool(o.AutoRepair),
		UpgradeOptions: ContainerBetaClusterNodePoolsManagementUpgradeOptionsToProto(o.UpgradeOptions),
	}
	return p
}

// ClusterNodePoolsManagementUpgradeOptionsToProto converts a ClusterNodePoolsManagementUpgradeOptions resource to its proto representation.
func ContainerBetaClusterNodePoolsManagementUpgradeOptionsToProto(o *beta.ClusterNodePoolsManagementUpgradeOptions) *betapb.ContainerBetaClusterNodePoolsManagementUpgradeOptions {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.ValueOrEmptyString(o.AutoUpgradeStartTime),
		Description:          dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// ClusterNodePoolsMaxPodsConstraintToProto converts a ClusterNodePoolsMaxPodsConstraint resource to its proto representation.
func ContainerBetaClusterNodePoolsMaxPodsConstraintToProto(o *beta.ClusterNodePoolsMaxPodsConstraint) *betapb.ContainerBetaClusterNodePoolsMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// ClusterNodePoolsConditionsToProto converts a ClusterNodePoolsConditions resource to its proto representation.
func ContainerBetaClusterNodePoolsConditionsToProto(o *beta.ClusterNodePoolsConditions) *betapb.ContainerBetaClusterNodePoolsConditions {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsConditions{
		Code:          ContainerBetaClusterNodePoolsConditionsCodeEnumToProto(o.Code),
		Message:       dcl.ValueOrEmptyString(o.Message),
		CanonicalCode: ContainerBetaClusterNodePoolsConditionsCanonicalCodeEnumToProto(o.CanonicalCode),
	}
	return p
}

// ClusterNodePoolsUpgradeSettingsToProto converts a ClusterNodePoolsUpgradeSettings resource to its proto representation.
func ContainerBetaClusterNodePoolsUpgradeSettingsToProto(o *beta.ClusterNodePoolsUpgradeSettings) *betapb.ContainerBetaClusterNodePoolsUpgradeSettings {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsUpgradeSettings{
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
	}
	return p
}

// ClusterNodePoolsNetworkConfigToProto converts a ClusterNodePoolsNetworkConfig resource to its proto representation.
func ContainerBetaClusterNodePoolsNetworkConfigToProto(o *beta.ClusterNodePoolsNetworkConfig) *betapb.ContainerBetaClusterNodePoolsNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodePoolsNetworkConfig{
		CreatePodRange:   dcl.ValueOrEmptyBool(o.CreatePodRange),
		PodRange:         dcl.ValueOrEmptyString(o.PodRange),
		PodIpv4CidrBlock: dcl.ValueOrEmptyString(o.PodIPv4CidrBlock),
	}
	return p
}

// ClusterLegacyAbacToProto converts a ClusterLegacyAbac resource to its proto representation.
func ContainerBetaClusterLegacyAbacToProto(o *beta.ClusterLegacyAbac) *betapb.ContainerBetaClusterLegacyAbac {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterLegacyAbac{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterNetworkPolicyToProto converts a ClusterNetworkPolicy resource to its proto representation.
func ContainerBetaClusterNetworkPolicyToProto(o *beta.ClusterNetworkPolicy) *betapb.ContainerBetaClusterNetworkPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNetworkPolicy{
		Provider: ContainerBetaClusterNetworkPolicyProviderEnumToProto(o.Provider),
		Enabled:  dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterIPAllocationPolicyToProto converts a ClusterIPAllocationPolicy resource to its proto representation.
func ContainerBetaClusterIPAllocationPolicyToProto(o *beta.ClusterIPAllocationPolicy) *betapb.ContainerBetaClusterIPAllocationPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterIPAllocationPolicy{
		UseIpAliases:               dcl.ValueOrEmptyBool(o.UseIPAliases),
		CreateSubnetwork:           dcl.ValueOrEmptyBool(o.CreateSubnetwork),
		SubnetworkName:             dcl.ValueOrEmptyString(o.SubnetworkName),
		ClusterSecondaryRangeName:  dcl.ValueOrEmptyString(o.ClusterSecondaryRangeName),
		ServicesSecondaryRangeName: dcl.ValueOrEmptyString(o.ServicesSecondaryRangeName),
		ClusterIpv4CidrBlock:       dcl.ValueOrEmptyString(o.ClusterIPv4CidrBlock),
		NodeIpv4CidrBlock:          dcl.ValueOrEmptyString(o.NodeIPv4CidrBlock),
		ServicesIpv4CidrBlock:      dcl.ValueOrEmptyString(o.ServicesIPv4CidrBlock),
		TpuIpv4CidrBlock:           dcl.ValueOrEmptyString(o.TPUIPv4CidrBlock),
		ClusterIpv4Cidr:            dcl.ValueOrEmptyString(o.ClusterIPv4Cidr),
		NodeIpv4Cidr:               dcl.ValueOrEmptyString(o.NodeIPv4Cidr),
		ServicesIpv4Cidr:           dcl.ValueOrEmptyString(o.ServicesIPv4Cidr),
		UseRoutes:                  dcl.ValueOrEmptyBool(o.UseRoutes),
		AllowRouteOverlap:          dcl.ValueOrEmptyBool(o.AllowRouteOverlap),
	}
	return p
}

// ClusterMasterAuthorizedNetworksConfigToProto converts a ClusterMasterAuthorizedNetworksConfig resource to its proto representation.
func ContainerBetaClusterMasterAuthorizedNetworksConfigToProto(o *beta.ClusterMasterAuthorizedNetworksConfig) *betapb.ContainerBetaClusterMasterAuthorizedNetworksConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMasterAuthorizedNetworksConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	for _, r := range o.CidrBlocks {
		p.CidrBlocks = append(p.CidrBlocks, ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocksToProto(&r))
	}
	return p
}

// ClusterMasterAuthorizedNetworksConfigCidrBlocksToProto converts a ClusterMasterAuthorizedNetworksConfigCidrBlocks resource to its proto representation.
func ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocksToProto(o *beta.ClusterMasterAuthorizedNetworksConfigCidrBlocks) *betapb.ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMasterAuthorizedNetworksConfigCidrBlocks{
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		CidrBlock:   dcl.ValueOrEmptyString(o.CidrBlock),
	}
	return p
}

// ClusterBinaryAuthorizationToProto converts a ClusterBinaryAuthorization resource to its proto representation.
func ContainerBetaClusterBinaryAuthorizationToProto(o *beta.ClusterBinaryAuthorization) *betapb.ContainerBetaClusterBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterBinaryAuthorization{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAutoscalingToProto converts a ClusterAutoscaling resource to its proto representation.
func ContainerBetaClusterAutoscalingToProto(o *beta.ClusterAutoscaling) *betapb.ContainerBetaClusterAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscaling{
		EnableNodeAutoprovisioning:       dcl.ValueOrEmptyBool(o.EnableNodeAutoprovisioning),
		AutoprovisioningNodePoolDefaults: ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto(o.AutoprovisioningNodePoolDefaults),
		AutoscalingProfile:               ContainerBetaClusterAutoscalingAutoscalingProfileEnumToProto(o.AutoscalingProfile),
	}
	for _, r := range o.ResourceLimits {
		p.ResourceLimits = append(p.ResourceLimits, ContainerBetaClusterAutoscalingResourceLimitsToProto(&r))
	}
	for _, r := range o.AutoprovisioningLocations {
		p.AutoprovisioningLocations = append(p.AutoprovisioningLocations, r)
	}
	return p
}

// ClusterAutoscalingResourceLimitsToProto converts a ClusterAutoscalingResourceLimits resource to its proto representation.
func ContainerBetaClusterAutoscalingResourceLimitsToProto(o *beta.ClusterAutoscalingResourceLimits) *betapb.ContainerBetaClusterAutoscalingResourceLimits {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingResourceLimits{
		ResourceType: dcl.ValueOrEmptyString(o.ResourceType),
		Minimum:      dcl.ValueOrEmptyInt64(o.Minimum),
		Maximum:      dcl.ValueOrEmptyInt64(o.Maximum),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaults resource to its proto representation.
func ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto(o *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaults) *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaults{
		ServiceAccount:         dcl.ValueOrEmptyString(o.ServiceAccount),
		UpgradeSettings:        ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto(o.UpgradeSettings),
		Management:             ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto(o.Management),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		DiskSizeGb:             dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		DiskType:               dcl.ValueOrEmptyString(o.DiskType),
		ShieldedInstanceConfig: ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
		BootDiskKmsKey:         dcl.ValueOrEmptyString(o.BootDiskKmsKey),
	}
	for _, r := range o.OAuthScopes {
		p.OauthScopes = append(p.OauthScopes, r)
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings resource to its proto representation.
func ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto(o *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement resource to its proto representation.
func ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto(o *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{
		AutoUpgrade:    dcl.ValueOrEmptyBool(o.AutoUpgrade),
		AutoRepair:     dcl.ValueOrEmptyBool(o.AutoRepair),
		UpgradeOptions: ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsToProto(o.UpgradeOptions),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions resource to its proto representation.
func ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsToProto(o *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.ValueOrEmptyString(o.AutoUpgradeStartTime),
		Description:          dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig resource to its proto representation.
func ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigToProto(o *beta.ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) *betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// ClusterNetworkConfigToProto converts a ClusterNetworkConfig resource to its proto representation.
func ContainerBetaClusterNetworkConfigToProto(o *beta.ClusterNetworkConfig) *betapb.ContainerBetaClusterNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNetworkConfig{
		Network:                   dcl.ValueOrEmptyString(o.Network),
		Subnetwork:                dcl.ValueOrEmptyString(o.Subnetwork),
		EnableIntraNodeVisibility: dcl.ValueOrEmptyBool(o.EnableIntraNodeVisibility),
		DefaultSnatStatus:         ContainerBetaClusterNetworkConfigDefaultSnatStatusToProto(o.DefaultSnatStatus),
		PrivateIpv6GoogleAccess:   ContainerBetaClusterNetworkConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess),
		DatapathProvider:          ContainerBetaClusterNetworkConfigDatapathProviderEnumToProto(o.DatapathProvider),
	}
	return p
}

// ClusterNetworkConfigDefaultSnatStatusToProto converts a ClusterNetworkConfigDefaultSnatStatus resource to its proto representation.
func ContainerBetaClusterNetworkConfigDefaultSnatStatusToProto(o *beta.ClusterNetworkConfigDefaultSnatStatus) *betapb.ContainerBetaClusterNetworkConfigDefaultSnatStatus {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNetworkConfigDefaultSnatStatus{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterMaintenancePolicyToProto converts a ClusterMaintenancePolicy resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyToProto(o *beta.ClusterMaintenancePolicy) *betapb.ContainerBetaClusterMaintenancePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicy{
		Window:          ContainerBetaClusterMaintenancePolicyWindowToProto(o.Window),
		ResourceVersion: dcl.ValueOrEmptyString(o.ResourceVersion),
	}
	return p
}

// ClusterMaintenancePolicyWindowToProto converts a ClusterMaintenancePolicyWindow resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyWindowToProto(o *beta.ClusterMaintenancePolicyWindow) *betapb.ContainerBetaClusterMaintenancePolicyWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicyWindow{
		DailyMaintenanceWindow: ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto(o.DailyMaintenanceWindow),
		RecurringWindow:        ContainerBetaClusterMaintenancePolicyWindowRecurringWindowToProto(o.RecurringWindow),
	}
	p.MaintenanceExclusions = make(map[string]string)
	for k, r := range o.MaintenanceExclusions {
		p.MaintenanceExclusions[k] = r
	}
	return p
}

// ClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto converts a ClusterMaintenancePolicyWindowDailyMaintenanceWindow resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto(o *beta.ClusterMaintenancePolicyWindowDailyMaintenanceWindow) *betapb.ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicyWindowDailyMaintenanceWindow{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		Duration:  dcl.ValueOrEmptyString(o.Duration),
	}
	return p
}

// ClusterMaintenancePolicyWindowRecurringWindowToProto converts a ClusterMaintenancePolicyWindowRecurringWindow resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyWindowRecurringWindowToProto(o *beta.ClusterMaintenancePolicyWindowRecurringWindow) *betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindow{
		Window:     ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindowToProto(o.Window),
		Recurrence: dcl.ValueOrEmptyString(o.Recurrence),
	}
	return p
}

// ClusterMaintenancePolicyWindowRecurringWindowWindowToProto converts a ClusterMaintenancePolicyWindowRecurringWindowWindow resource to its proto representation.
func ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindowToProto(o *beta.ClusterMaintenancePolicyWindowRecurringWindowWindow) *betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaintenancePolicyWindowRecurringWindowWindow{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		EndTime:   dcl.ValueOrEmptyString(o.EndTime),
	}
	return p
}

// ClusterDefaultMaxPodsConstraintToProto converts a ClusterDefaultMaxPodsConstraint resource to its proto representation.
func ContainerBetaClusterDefaultMaxPodsConstraintToProto(o *beta.ClusterDefaultMaxPodsConstraint) *betapb.ContainerBetaClusterDefaultMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterDefaultMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyString(o.MaxPodsPerNode),
	}
	return p
}

// ClusterResourceUsageExportConfigToProto converts a ClusterResourceUsageExportConfig resource to its proto representation.
func ContainerBetaClusterResourceUsageExportConfigToProto(o *beta.ClusterResourceUsageExportConfig) *betapb.ContainerBetaClusterResourceUsageExportConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterResourceUsageExportConfig{
		BigqueryDestination:           ContainerBetaClusterResourceUsageExportConfigBigqueryDestinationToProto(o.BigqueryDestination),
		EnableNetworkEgressMonitoring: dcl.ValueOrEmptyBool(o.EnableNetworkEgressMonitoring),
		ConsumptionMeteringConfig:     ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfigToProto(o.ConsumptionMeteringConfig),
		EnableNetworkEgressMetering:   dcl.ValueOrEmptyBool(o.EnableNetworkEgressMetering),
	}
	return p
}

// ClusterResourceUsageExportConfigBigqueryDestinationToProto converts a ClusterResourceUsageExportConfigBigqueryDestination resource to its proto representation.
func ContainerBetaClusterResourceUsageExportConfigBigqueryDestinationToProto(o *beta.ClusterResourceUsageExportConfigBigqueryDestination) *betapb.ContainerBetaClusterResourceUsageExportConfigBigqueryDestination {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterResourceUsageExportConfigBigqueryDestination{
		DatasetId: dcl.ValueOrEmptyString(o.DatasetId),
	}
	return p
}

// ClusterResourceUsageExportConfigConsumptionMeteringConfigToProto converts a ClusterResourceUsageExportConfigConsumptionMeteringConfig resource to its proto representation.
func ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfigToProto(o *beta.ClusterResourceUsageExportConfigConsumptionMeteringConfig) *betapb.ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterResourceUsageExportConfigConsumptionMeteringConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAuthenticatorGroupsConfigToProto converts a ClusterAuthenticatorGroupsConfig resource to its proto representation.
func ContainerBetaClusterAuthenticatorGroupsConfigToProto(o *beta.ClusterAuthenticatorGroupsConfig) *betapb.ContainerBetaClusterAuthenticatorGroupsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAuthenticatorGroupsConfig{
		Enabled:       dcl.ValueOrEmptyBool(o.Enabled),
		SecurityGroup: dcl.ValueOrEmptyString(o.SecurityGroup),
	}
	return p
}

// ClusterPrivateClusterConfigToProto converts a ClusterPrivateClusterConfig resource to its proto representation.
func ContainerBetaClusterPrivateClusterConfigToProto(o *beta.ClusterPrivateClusterConfig) *betapb.ContainerBetaClusterPrivateClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterPrivateClusterConfig{
		EnablePrivateNodes:       dcl.ValueOrEmptyBool(o.EnablePrivateNodes),
		EnablePrivateEndpoint:    dcl.ValueOrEmptyBool(o.EnablePrivateEndpoint),
		MasterIpv4CidrBlock:      dcl.ValueOrEmptyString(o.MasterIPv4CidrBlock),
		PrivateEndpoint:          dcl.ValueOrEmptyString(o.PrivateEndpoint),
		PublicEndpoint:           dcl.ValueOrEmptyString(o.PublicEndpoint),
		PeeringName:              dcl.ValueOrEmptyString(o.PeeringName),
		MasterGlobalAccessConfig: ContainerBetaClusterPrivateClusterConfigMasterGlobalAccessConfigToProto(o.MasterGlobalAccessConfig),
	}
	return p
}

// ClusterPrivateClusterConfigMasterGlobalAccessConfigToProto converts a ClusterPrivateClusterConfigMasterGlobalAccessConfig resource to its proto representation.
func ContainerBetaClusterPrivateClusterConfigMasterGlobalAccessConfigToProto(o *beta.ClusterPrivateClusterConfigMasterGlobalAccessConfig) *betapb.ContainerBetaClusterPrivateClusterConfigMasterGlobalAccessConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterPrivateClusterConfigMasterGlobalAccessConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterDatabaseEncryptionToProto converts a ClusterDatabaseEncryption resource to its proto representation.
func ContainerBetaClusterDatabaseEncryptionToProto(o *beta.ClusterDatabaseEncryption) *betapb.ContainerBetaClusterDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterDatabaseEncryption{
		State:   ContainerBetaClusterDatabaseEncryptionStateEnumToProto(o.State),
		KeyName: dcl.ValueOrEmptyString(o.KeyName),
	}
	return p
}

// ClusterVerticalPodAutoscalingToProto converts a ClusterVerticalPodAutoscaling resource to its proto representation.
func ContainerBetaClusterVerticalPodAutoscalingToProto(o *beta.ClusterVerticalPodAutoscaling) *betapb.ContainerBetaClusterVerticalPodAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterVerticalPodAutoscaling{
		Enabled:                    dcl.ValueOrEmptyBool(o.Enabled),
		EnableExperimentalFeatures: dcl.ValueOrEmptyBool(o.EnableExperimentalFeatures),
	}
	return p
}

// ClusterShieldedNodesToProto converts a ClusterShieldedNodes resource to its proto representation.
func ContainerBetaClusterShieldedNodesToProto(o *beta.ClusterShieldedNodes) *betapb.ContainerBetaClusterShieldedNodes {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterShieldedNodes{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterConditionsToProto converts a ClusterConditions resource to its proto representation.
func ContainerBetaClusterConditionsToProto(o *beta.ClusterConditions) *betapb.ContainerBetaClusterConditions {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterConditions{
		Code:          dcl.ValueOrEmptyString(o.Code),
		Message:       dcl.ValueOrEmptyString(o.Message),
		CanonicalCode: ContainerBetaClusterConditionsCanonicalCodeEnumToProto(o.CanonicalCode),
	}
	return p
}

// ClusterAutopilotToProto converts a ClusterAutopilot resource to its proto representation.
func ContainerBetaClusterAutopilotToProto(o *beta.ClusterAutopilot) *betapb.ContainerBetaClusterAutopilot {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterAutopilot{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterNodeConfigToProto converts a ClusterNodeConfig resource to its proto representation.
func ContainerBetaClusterNodeConfigToProto(o *beta.ClusterNodeConfig) *betapb.ContainerBetaClusterNodeConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfig{
		MachineType:            dcl.ValueOrEmptyString(o.MachineType),
		DiskSizeGb:             dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		ServiceAccount:         dcl.ValueOrEmptyString(o.ServiceAccount),
		ImageType:              dcl.ValueOrEmptyString(o.ImageType),
		LocalSsdCount:          dcl.ValueOrEmptyInt64(o.LocalSsdCount),
		Preemptible:            dcl.ValueOrEmptyBool(o.Preemptible),
		DiskType:               dcl.ValueOrEmptyString(o.DiskType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		WorkloadMetadataConfig: ContainerBetaClusterNodeConfigWorkloadMetadataConfigToProto(o.WorkloadMetadataConfig),
		SandboxConfig:          ContainerBetaClusterNodeConfigSandboxConfigToProto(o.SandboxConfig),
		NodeGroup:              dcl.ValueOrEmptyString(o.NodeGroup),
		ReservationAffinity:    ContainerBetaClusterNodeConfigReservationAffinityToProto(o.ReservationAffinity),
		ShieldedInstanceConfig: ContainerBetaClusterNodeConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
		LinuxNodeConfig:        ContainerBetaClusterNodeConfigLinuxNodeConfigToProto(o.LinuxNodeConfig),
		KubeletConfig:          ContainerBetaClusterNodeConfigKubeletConfigToProto(o.KubeletConfig),
		BootDiskKmsKey:         dcl.ValueOrEmptyString(o.BootDiskKmsKey),
		EphemeralStorageConfig: ContainerBetaClusterNodeConfigEphemeralStorageConfigToProto(o.EphemeralStorageConfig),
	}
	for _, r := range o.OAuthScopes {
		p.OauthScopes = append(p.OauthScopes, r)
	}
	p.Metadata = make(map[string]string)
	for k, r := range o.Metadata {
		p.Metadata[k] = r
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	for _, r := range o.Accelerators {
		p.Accelerators = append(p.Accelerators, ContainerBetaClusterNodeConfigAcceleratorsToProto(&r))
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, ContainerBetaClusterNodeConfigTaintsToProto(&r))
	}
	return p
}

// ClusterNodeConfigAcceleratorsToProto converts a ClusterNodeConfigAccelerators resource to its proto representation.
func ContainerBetaClusterNodeConfigAcceleratorsToProto(o *beta.ClusterNodeConfigAccelerators) *betapb.ContainerBetaClusterNodeConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// ClusterNodeConfigWorkloadMetadataConfigToProto converts a ClusterNodeConfigWorkloadMetadataConfig resource to its proto representation.
func ContainerBetaClusterNodeConfigWorkloadMetadataConfigToProto(o *beta.ClusterNodeConfigWorkloadMetadataConfig) *betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigWorkloadMetadataConfig{
		Mode:         ContainerBetaClusterNodeConfigWorkloadMetadataConfigModeEnumToProto(o.Mode),
		NodeMetadata: ContainerBetaClusterNodeConfigWorkloadMetadataConfigNodeMetadataEnumToProto(o.NodeMetadata),
	}
	return p
}

// ClusterNodeConfigTaintsToProto converts a ClusterNodeConfigTaints resource to its proto representation.
func ContainerBetaClusterNodeConfigTaintsToProto(o *beta.ClusterNodeConfigTaints) *betapb.ContainerBetaClusterNodeConfigTaints {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: ContainerBetaClusterNodeConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// ClusterNodeConfigSandboxConfigToProto converts a ClusterNodeConfigSandboxConfig resource to its proto representation.
func ContainerBetaClusterNodeConfigSandboxConfigToProto(o *beta.ClusterNodeConfigSandboxConfig) *betapb.ContainerBetaClusterNodeConfigSandboxConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigSandboxConfig{
		Type:        ContainerBetaClusterNodeConfigSandboxConfigTypeEnumToProto(o.Type),
		SandboxType: dcl.ValueOrEmptyString(o.SandboxType),
	}
	return p
}

// ClusterNodeConfigReservationAffinityToProto converts a ClusterNodeConfigReservationAffinity resource to its proto representation.
func ContainerBetaClusterNodeConfigReservationAffinityToProto(o *beta.ClusterNodeConfigReservationAffinity) *betapb.ContainerBetaClusterNodeConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigReservationAffinity{
		ConsumeReservationType: ContainerBetaClusterNodeConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// ClusterNodeConfigShieldedInstanceConfigToProto converts a ClusterNodeConfigShieldedInstanceConfig resource to its proto representation.
func ContainerBetaClusterNodeConfigShieldedInstanceConfigToProto(o *beta.ClusterNodeConfigShieldedInstanceConfig) *betapb.ContainerBetaClusterNodeConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// ClusterNodeConfigLinuxNodeConfigToProto converts a ClusterNodeConfigLinuxNodeConfig resource to its proto representation.
func ContainerBetaClusterNodeConfigLinuxNodeConfigToProto(o *beta.ClusterNodeConfigLinuxNodeConfig) *betapb.ContainerBetaClusterNodeConfigLinuxNodeConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigLinuxNodeConfig{}
	p.Sysctls = make(map[string]string)
	for k, r := range o.Sysctls {
		p.Sysctls[k] = r
	}
	return p
}

// ClusterNodeConfigKubeletConfigToProto converts a ClusterNodeConfigKubeletConfig resource to its proto representation.
func ContainerBetaClusterNodeConfigKubeletConfigToProto(o *beta.ClusterNodeConfigKubeletConfig) *betapb.ContainerBetaClusterNodeConfigKubeletConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigKubeletConfig{
		CpuManagerPolicy:  dcl.ValueOrEmptyString(o.CpuManagerPolicy),
		CpuCfsQuota:       dcl.ValueOrEmptyBool(o.CpuCfsQuota),
		CpuCfsQuotaPeriod: dcl.ValueOrEmptyString(o.CpuCfsQuotaPeriod),
	}
	return p
}

// ClusterNodeConfigEphemeralStorageConfigToProto converts a ClusterNodeConfigEphemeralStorageConfig resource to its proto representation.
func ContainerBetaClusterNodeConfigEphemeralStorageConfigToProto(o *beta.ClusterNodeConfigEphemeralStorageConfig) *betapb.ContainerBetaClusterNodeConfigEphemeralStorageConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNodeConfigEphemeralStorageConfig{
		LocalSsdCount: dcl.ValueOrEmptyInt64(o.LocalSsdCount),
	}
	return p
}

// ClusterReleaseChannelToProto converts a ClusterReleaseChannel resource to its proto representation.
func ContainerBetaClusterReleaseChannelToProto(o *beta.ClusterReleaseChannel) *betapb.ContainerBetaClusterReleaseChannel {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterReleaseChannel{
		Channel: ContainerBetaClusterReleaseChannelChannelEnumToProto(o.Channel),
	}
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig resource to its proto representation.
func ContainerBetaClusterWorkloadIdentityConfigToProto(o *beta.ClusterWorkloadIdentityConfig) *betapb.ContainerBetaClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterWorkloadIdentityConfig{
		WorkloadPool:      dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityNamespace: dcl.ValueOrEmptyString(o.IdentityNamespace),
		IdentityProvider:  dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// ClusterNotificationConfigToProto converts a ClusterNotificationConfig resource to its proto representation.
func ContainerBetaClusterNotificationConfigToProto(o *beta.ClusterNotificationConfig) *betapb.ContainerBetaClusterNotificationConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNotificationConfig{
		Pubsub: ContainerBetaClusterNotificationConfigPubsubToProto(o.Pubsub),
	}
	return p
}

// ClusterNotificationConfigPubsubToProto converts a ClusterNotificationConfigPubsub resource to its proto representation.
func ContainerBetaClusterNotificationConfigPubsubToProto(o *beta.ClusterNotificationConfigPubsub) *betapb.ContainerBetaClusterNotificationConfigPubsub {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterNotificationConfigPubsub{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
		Topic:   dcl.ValueOrEmptyString(o.Topic),
	}
	return p
}

// ClusterConfidentialNodesToProto converts a ClusterConfidentialNodes resource to its proto representation.
func ContainerBetaClusterConfidentialNodesToProto(o *beta.ClusterConfidentialNodes) *betapb.ContainerBetaClusterConfidentialNodes {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterConfidentialNodes{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterPodSecurityPolicyConfigToProto converts a ClusterPodSecurityPolicyConfig resource to its proto representation.
func ContainerBetaClusterPodSecurityPolicyConfigToProto(o *beta.ClusterPodSecurityPolicyConfig) *betapb.ContainerBetaClusterPodSecurityPolicyConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterPodSecurityPolicyConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterClusterTelemetryToProto converts a ClusterClusterTelemetry resource to its proto representation.
func ContainerBetaClusterClusterTelemetryToProto(o *beta.ClusterClusterTelemetry) *betapb.ContainerBetaClusterClusterTelemetry {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterClusterTelemetry{
		Type: ContainerBetaClusterClusterTelemetryTypeEnumToProto(o.Type),
	}
	return p
}

// ClusterTPUConfigToProto converts a ClusterTPUConfig resource to its proto representation.
func ContainerBetaClusterTPUConfigToProto(o *beta.ClusterTPUConfig) *betapb.ContainerBetaClusterTPUConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterTPUConfig{
		Enabled:              dcl.ValueOrEmptyBool(o.Enabled),
		UseServiceNetworking: dcl.ValueOrEmptyBool(o.UseServiceNetworking),
		Ipv4CidrBlock:        dcl.ValueOrEmptyString(o.IPv4CidrBlock),
	}
	return p
}

// ClusterMasterToProto converts a ClusterMaster resource to its proto representation.
func ContainerBetaClusterMasterToProto(o *beta.ClusterMaster) *betapb.ContainerBetaClusterMaster {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaClusterMaster{}
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *beta.Cluster) *betapb.ContainerBetaCluster {
	p := &betapb.ContainerBetaCluster{
		Name:                           dcl.ValueOrEmptyString(resource.Name),
		Description:                    dcl.ValueOrEmptyString(resource.Description),
		InitialNodeCount:               dcl.ValueOrEmptyInt64(resource.InitialNodeCount),
		MasterAuth:                     ContainerBetaClusterMasterAuthToProto(resource.MasterAuth),
		LoggingService:                 dcl.ValueOrEmptyString(resource.LoggingService),
		MonitoringService:              dcl.ValueOrEmptyString(resource.MonitoringService),
		Network:                        dcl.ValueOrEmptyString(resource.Network),
		ClusterIpv4Cidr:                dcl.ValueOrEmptyString(resource.ClusterIPv4Cidr),
		AddonsConfig:                   ContainerBetaClusterAddonsConfigToProto(resource.AddonsConfig),
		Subnetwork:                     dcl.ValueOrEmptyString(resource.Subnetwork),
		EnableKubernetesAlpha:          dcl.ValueOrEmptyBool(resource.EnableKubernetesAlpha),
		LabelFingerprint:               dcl.ValueOrEmptyString(resource.LabelFingerprint),
		LegacyAbac:                     ContainerBetaClusterLegacyAbacToProto(resource.LegacyAbac),
		NetworkPolicy:                  ContainerBetaClusterNetworkPolicyToProto(resource.NetworkPolicy),
		IpAllocationPolicy:             ContainerBetaClusterIPAllocationPolicyToProto(resource.IPAllocationPolicy),
		MasterAuthorizedNetworksConfig: ContainerBetaClusterMasterAuthorizedNetworksConfigToProto(resource.MasterAuthorizedNetworksConfig),
		BinaryAuthorization:            ContainerBetaClusterBinaryAuthorizationToProto(resource.BinaryAuthorization),
		Autoscaling:                    ContainerBetaClusterAutoscalingToProto(resource.Autoscaling),
		NetworkConfig:                  ContainerBetaClusterNetworkConfigToProto(resource.NetworkConfig),
		MaintenancePolicy:              ContainerBetaClusterMaintenancePolicyToProto(resource.MaintenancePolicy),
		DefaultMaxPodsConstraint:       ContainerBetaClusterDefaultMaxPodsConstraintToProto(resource.DefaultMaxPodsConstraint),
		ResourceUsageExportConfig:      ContainerBetaClusterResourceUsageExportConfigToProto(resource.ResourceUsageExportConfig),
		AuthenticatorGroupsConfig:      ContainerBetaClusterAuthenticatorGroupsConfigToProto(resource.AuthenticatorGroupsConfig),
		PrivateClusterConfig:           ContainerBetaClusterPrivateClusterConfigToProto(resource.PrivateClusterConfig),
		DatabaseEncryption:             ContainerBetaClusterDatabaseEncryptionToProto(resource.DatabaseEncryption),
		VerticalPodAutoscaling:         ContainerBetaClusterVerticalPodAutoscalingToProto(resource.VerticalPodAutoscaling),
		ShieldedNodes:                  ContainerBetaClusterShieldedNodesToProto(resource.ShieldedNodes),
		Endpoint:                       dcl.ValueOrEmptyString(resource.Endpoint),
		MasterVersion:                  dcl.ValueOrEmptyString(resource.MasterVersion),
		CreateTime:                     dcl.ValueOrEmptyString(resource.CreateTime),
		Status:                         dcl.ValueOrEmptyString(resource.Status),
		StatusMessage:                  dcl.ValueOrEmptyString(resource.StatusMessage),
		NodeIpv4CidrSize:               dcl.ValueOrEmptyInt64(resource.NodeIPv4CidrSize),
		ServicesIpv4Cidr:               dcl.ValueOrEmptyString(resource.ServicesIPv4Cidr),
		ExpireTime:                     dcl.ValueOrEmptyString(resource.ExpireTime),
		Location:                       dcl.ValueOrEmptyString(resource.Location),
		EnableTpu:                      dcl.ValueOrEmptyBool(resource.EnableTPU),
		TpuIpv4CidrBlock:               dcl.ValueOrEmptyString(resource.TPUIPv4CidrBlock),
		Autopilot:                      ContainerBetaClusterAutopilotToProto(resource.Autopilot),
		Project:                        dcl.ValueOrEmptyString(resource.Project),
		NodeConfig:                     ContainerBetaClusterNodeConfigToProto(resource.NodeConfig),
		ReleaseChannel:                 ContainerBetaClusterReleaseChannelToProto(resource.ReleaseChannel),
		WorkloadIdentityConfig:         ContainerBetaClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		NotificationConfig:             ContainerBetaClusterNotificationConfigToProto(resource.NotificationConfig),
		ConfidentialNodes:              ContainerBetaClusterConfidentialNodesToProto(resource.ConfidentialNodes),
		SelfLink:                       dcl.ValueOrEmptyString(resource.SelfLink),
		Zone:                           dcl.ValueOrEmptyString(resource.Zone),
		InitialClusterVersion:          dcl.ValueOrEmptyString(resource.InitialClusterVersion),
		CurrentMasterVersion:           dcl.ValueOrEmptyString(resource.CurrentMasterVersion),
		CurrentNodeVersion:             dcl.ValueOrEmptyString(resource.CurrentNodeVersion),
		CurrentNodeCount:               dcl.ValueOrEmptyInt64(resource.CurrentNodeCount),
		Id:                             dcl.ValueOrEmptyString(resource.Id),
		PodSecurityPolicyConfig:        ContainerBetaClusterPodSecurityPolicyConfigToProto(resource.PodSecurityPolicyConfig),
		PrivateCluster:                 dcl.ValueOrEmptyBool(resource.PrivateCluster),
		MasterIpv4CidrBlock:            dcl.ValueOrEmptyString(resource.MasterIPv4CidrBlock),
		ClusterTelemetry:               ContainerBetaClusterClusterTelemetryToProto(resource.ClusterTelemetry),
		TpuConfig:                      ContainerBetaClusterTPUConfigToProto(resource.TPUConfig),
		Master:                         ContainerBetaClusterMasterToProto(resource.Master),
	}
	for _, r := range resource.NodePools {
		p.NodePools = append(p.NodePools, ContainerBetaClusterNodePoolsToProto(&r))
	}
	for _, r := range resource.Locations {
		p.Locations = append(p.Locations, r)
	}
	for _, r := range resource.Conditions {
		p.Conditions = append(p.Conditions, ContainerBetaClusterConditionsToProto(&r))
	}
	for _, r := range resource.InstanceGroupUrls {
		p.InstanceGroupUrls = append(p.InstanceGroupUrls, r)
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *beta.Client, request *betapb.ApplyContainerBetaClusterRequest) (*betapb.ContainerBetaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerBetaCluster(ctx context.Context, request *betapb.ApplyContainerBetaClusterRequest) (*betapb.ContainerBetaCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerBetaCluster(ctx context.Context, request *betapb.DeleteContainerBetaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerBetaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerBetaCluster(ctx context.Context, request *betapb.ListContainerBetaClusterRequest) (*betapb.ListContainerBetaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContainerBetaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListContainerBetaClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
