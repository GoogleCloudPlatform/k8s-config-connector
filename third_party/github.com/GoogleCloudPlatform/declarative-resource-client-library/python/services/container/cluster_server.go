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
	containerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/container/container_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container"
)

// Server implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum converts a ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum enum from its proto representation.
func ProtoToContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(e containerpb.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum) *container.ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum_name[int32(e)]; ok {
		e := container.ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(n[len("ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigWorkloadMetadataConfigModeEnum converts a ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum enum from its proto representation.
func ProtoToContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(e containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum) *container.ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum_name[int32(e)]; ok {
		e := container.ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(n[len("ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigTaintsEffectEnum converts a ClusterNodePoolsConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerClusterNodePoolsConfigTaintsEffectEnum(e containerpb.ContainerClusterNodePoolsConfigTaintsEffectEnum) *container.ClusterNodePoolsConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodePoolsConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := container.ClusterNodePoolsConfigTaintsEffectEnum(n[len("ContainerClusterNodePoolsConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigSandboxConfigTypeEnum converts a ClusterNodePoolsConfigSandboxConfigTypeEnum enum from its proto representation.
func ProtoToContainerClusterNodePoolsConfigSandboxConfigTypeEnum(e containerpb.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum) *container.ClusterNodePoolsConfigSandboxConfigTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum_name[int32(e)]; ok {
		e := container.ClusterNodePoolsConfigSandboxConfigTypeEnum(n[len("ContainerClusterNodePoolsConfigSandboxConfigTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(e containerpb.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum) *container.ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := container.ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(n[len("ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsStatusEnum converts a ClusterNodePoolsStatusEnum enum from its proto representation.
func ProtoToContainerClusterNodePoolsStatusEnum(e containerpb.ContainerClusterNodePoolsStatusEnum) *container.ClusterNodePoolsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodePoolsStatusEnum_name[int32(e)]; ok {
		e := container.ClusterNodePoolsStatusEnum(n[len("ContainerClusterNodePoolsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConditionsCodeEnum converts a ClusterNodePoolsConditionsCodeEnum enum from its proto representation.
func ProtoToContainerClusterNodePoolsConditionsCodeEnum(e containerpb.ContainerClusterNodePoolsConditionsCodeEnum) *container.ClusterNodePoolsConditionsCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodePoolsConditionsCodeEnum_name[int32(e)]; ok {
		e := container.ClusterNodePoolsConditionsCodeEnum(n[len("ContainerClusterNodePoolsConditionsCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodePoolsConditionsCanonicalCodeEnum converts a ClusterNodePoolsConditionsCanonicalCodeEnum enum from its proto representation.
func ProtoToContainerClusterNodePoolsConditionsCanonicalCodeEnum(e containerpb.ContainerClusterNodePoolsConditionsCanonicalCodeEnum) *container.ClusterNodePoolsConditionsCanonicalCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodePoolsConditionsCanonicalCodeEnum_name[int32(e)]; ok {
		e := container.ClusterNodePoolsConditionsCanonicalCodeEnum(n[len("ContainerClusterNodePoolsConditionsCanonicalCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworkPolicyProviderEnum converts a ClusterNetworkPolicyProviderEnum enum from its proto representation.
func ProtoToContainerClusterNetworkPolicyProviderEnum(e containerpb.ContainerClusterNetworkPolicyProviderEnum) *container.ClusterNetworkPolicyProviderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNetworkPolicyProviderEnum_name[int32(e)]; ok {
		e := container.ClusterNetworkPolicyProviderEnum(n[len("ContainerClusterNetworkPolicyProviderEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworkConfigPrivateIPv6GoogleAccessEnum converts a ClusterNetworkConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum(e containerpb.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum) *container.ClusterNetworkConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := container.ClusterNetworkConfigPrivateIPv6GoogleAccessEnum(n[len("ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterDatabaseEncryptionStateEnum converts a ClusterDatabaseEncryptionStateEnum enum from its proto representation.
func ProtoToContainerClusterDatabaseEncryptionStateEnum(e containerpb.ContainerClusterDatabaseEncryptionStateEnum) *container.ClusterDatabaseEncryptionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterDatabaseEncryptionStateEnum_name[int32(e)]; ok {
		e := container.ClusterDatabaseEncryptionStateEnum(n[len("ContainerClusterDatabaseEncryptionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConditionsCanonicalCodeEnum converts a ClusterConditionsCanonicalCodeEnum enum from its proto representation.
func ProtoToContainerClusterConditionsCanonicalCodeEnum(e containerpb.ContainerClusterConditionsCanonicalCodeEnum) *container.ClusterConditionsCanonicalCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterConditionsCanonicalCodeEnum_name[int32(e)]; ok {
		e := container.ClusterConditionsCanonicalCodeEnum(n[len("ContainerClusterConditionsCanonicalCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigWorkloadMetadataConfigModeEnum converts a ClusterNodeConfigWorkloadMetadataConfigModeEnum enum from its proto representation.
func ProtoToContainerClusterNodeConfigWorkloadMetadataConfigModeEnum(e containerpb.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum) *container.ClusterNodeConfigWorkloadMetadataConfigModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum_name[int32(e)]; ok {
		e := container.ClusterNodeConfigWorkloadMetadataConfigModeEnum(n[len("ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigTaintsEffectEnum converts a ClusterNodeConfigTaintsEffectEnum enum from its proto representation.
func ProtoToContainerClusterNodeConfigTaintsEffectEnum(e containerpb.ContainerClusterNodeConfigTaintsEffectEnum) *container.ClusterNodeConfigTaintsEffectEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodeConfigTaintsEffectEnum_name[int32(e)]; ok {
		e := container.ClusterNodeConfigTaintsEffectEnum(n[len("ContainerClusterNodeConfigTaintsEffectEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigSandboxConfigTypeEnum converts a ClusterNodeConfigSandboxConfigTypeEnum enum from its proto representation.
func ProtoToContainerClusterNodeConfigSandboxConfigTypeEnum(e containerpb.ContainerClusterNodeConfigSandboxConfigTypeEnum) *container.ClusterNodeConfigSandboxConfigTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodeConfigSandboxConfigTypeEnum_name[int32(e)]; ok {
		e := container.ClusterNodeConfigSandboxConfigTypeEnum(n[len("ContainerClusterNodeConfigSandboxConfigTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNodeConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(e containerpb.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum) *container.ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := container.ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(n[len("ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterReleaseChannelChannelEnum converts a ClusterReleaseChannelChannelEnum enum from its proto representation.
func ProtoToContainerClusterReleaseChannelChannelEnum(e containerpb.ContainerClusterReleaseChannelChannelEnum) *container.ClusterReleaseChannelChannelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerClusterReleaseChannelChannelEnum_name[int32(e)]; ok {
		e := container.ClusterReleaseChannelChannelEnum(n[len("ContainerClusterReleaseChannelChannelEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterMasterAuth converts a ClusterMasterAuth resource from its proto representation.
func ProtoToContainerClusterMasterAuth(p *containerpb.ContainerClusterMasterAuth) *container.ClusterMasterAuth {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMasterAuth{
		Username:                dcl.StringOrNil(p.Username),
		Password:                dcl.StringOrNil(p.Password),
		ClientCertificateConfig: ProtoToContainerClusterMasterAuthClientCertificateConfig(p.GetClientCertificateConfig()),
		ClusterCaCertificate:    dcl.StringOrNil(p.ClusterCaCertificate),
		ClientCertificate:       dcl.StringOrNil(p.ClientCertificate),
		ClientKey:               dcl.StringOrNil(p.ClientKey),
	}
	return obj
}

// ProtoToClusterMasterAuthClientCertificateConfig converts a ClusterMasterAuthClientCertificateConfig resource from its proto representation.
func ProtoToContainerClusterMasterAuthClientCertificateConfig(p *containerpb.ContainerClusterMasterAuthClientCertificateConfig) *container.ClusterMasterAuthClientCertificateConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMasterAuthClientCertificateConfig{
		IssueClientCertificate: dcl.Bool(p.IssueClientCertificate),
	}
	return obj
}

// ProtoToClusterAddonsConfig converts a ClusterAddonsConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfig(p *containerpb.ContainerClusterAddonsConfig) *container.ClusterAddonsConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfig{
		HttpLoadBalancing:                ProtoToContainerClusterAddonsConfigHttpLoadBalancing(p.GetHttpLoadBalancing()),
		HorizontalPodAutoscaling:         ProtoToContainerClusterAddonsConfigHorizontalPodAutoscaling(p.GetHorizontalPodAutoscaling()),
		KubernetesDashboard:              ProtoToContainerClusterAddonsConfigKubernetesDashboard(p.GetKubernetesDashboard()),
		NetworkPolicyConfig:              ProtoToContainerClusterAddonsConfigNetworkPolicyConfig(p.GetNetworkPolicyConfig()),
		CloudRunConfig:                   ProtoToContainerClusterAddonsConfigCloudRunConfig(p.GetCloudRunConfig()),
		DnsCacheConfig:                   ProtoToContainerClusterAddonsConfigDnsCacheConfig(p.GetDnsCacheConfig()),
		ConfigConnectorConfig:            ProtoToContainerClusterAddonsConfigConfigConnectorConfig(p.GetConfigConnectorConfig()),
		GcePersistentDiskCsiDriverConfig: ProtoToContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig(p.GetGcePersistentDiskCsiDriverConfig()),
	}
	return obj
}

// ProtoToClusterAddonsConfigHttpLoadBalancing converts a ClusterAddonsConfigHttpLoadBalancing resource from its proto representation.
func ProtoToContainerClusterAddonsConfigHttpLoadBalancing(p *containerpb.ContainerClusterAddonsConfigHttpLoadBalancing) *container.ClusterAddonsConfigHttpLoadBalancing {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigHttpLoadBalancing{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigHorizontalPodAutoscaling converts a ClusterAddonsConfigHorizontalPodAutoscaling resource from its proto representation.
func ProtoToContainerClusterAddonsConfigHorizontalPodAutoscaling(p *containerpb.ContainerClusterAddonsConfigHorizontalPodAutoscaling) *container.ClusterAddonsConfigHorizontalPodAutoscaling {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigHorizontalPodAutoscaling{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigKubernetesDashboard converts a ClusterAddonsConfigKubernetesDashboard resource from its proto representation.
func ProtoToContainerClusterAddonsConfigKubernetesDashboard(p *containerpb.ContainerClusterAddonsConfigKubernetesDashboard) *container.ClusterAddonsConfigKubernetesDashboard {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigKubernetesDashboard{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigNetworkPolicyConfig converts a ClusterAddonsConfigNetworkPolicyConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfigNetworkPolicyConfig(p *containerpb.ContainerClusterAddonsConfigNetworkPolicyConfig) *container.ClusterAddonsConfigNetworkPolicyConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigNetworkPolicyConfig{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigCloudRunConfig converts a ClusterAddonsConfigCloudRunConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfigCloudRunConfig(p *containerpb.ContainerClusterAddonsConfigCloudRunConfig) *container.ClusterAddonsConfigCloudRunConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigCloudRunConfig{
		Disabled:         dcl.Bool(p.Disabled),
		LoadBalancerType: ProtoToContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(p.GetLoadBalancerType()),
	}
	return obj
}

// ProtoToClusterAddonsConfigDnsCacheConfig converts a ClusterAddonsConfigDnsCacheConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfigDnsCacheConfig(p *containerpb.ContainerClusterAddonsConfigDnsCacheConfig) *container.ClusterAddonsConfigDnsCacheConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigDnsCacheConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigConfigConnectorConfig converts a ClusterAddonsConfigConfigConnectorConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfigConfigConnectorConfig(p *containerpb.ContainerClusterAddonsConfigConfigConnectorConfig) *container.ClusterAddonsConfigConfigConnectorConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigConfigConnectorConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAddonsConfigGcePersistentDiskCsiDriverConfig converts a ClusterAddonsConfigGcePersistentDiskCsiDriverConfig resource from its proto representation.
func ProtoToContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig(p *containerpb.ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig) *container.ClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAddonsConfigGcePersistentDiskCsiDriverConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterNodePools converts a ClusterNodePools resource from its proto representation.
func ProtoToContainerClusterNodePools(p *containerpb.ContainerClusterNodePools) *container.ClusterNodePools {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePools{
		Name:              dcl.StringOrNil(p.Name),
		Config:            ProtoToContainerClusterNodePoolsConfig(p.GetConfig()),
		InitialNodeCount:  dcl.Int64OrNil(p.InitialNodeCount),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Version:           dcl.StringOrNil(p.Version),
		Status:            ProtoToContainerClusterNodePoolsStatusEnum(p.GetStatus()),
		StatusMessage:     dcl.StringOrNil(p.StatusMessage),
		Autoscaling:       ProtoToContainerClusterNodePoolsAutoscaling(p.GetAutoscaling()),
		Management:        ProtoToContainerClusterNodePoolsManagement(p.GetManagement()),
		MaxPodsConstraint: ProtoToContainerClusterNodePoolsMaxPodsConstraint(p.GetMaxPodsConstraint()),
		PodIPv4CidrSize:   dcl.Int64OrNil(p.PodIpv4CidrSize),
		UpgradeSettings:   ProtoToContainerClusterNodePoolsUpgradeSettings(p.GetUpgradeSettings()),
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	for _, r := range p.GetInstanceGroupUrls() {
		obj.InstanceGroupUrls = append(obj.InstanceGroupUrls, r)
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToContainerClusterNodePoolsConditions(r))
	}
	return obj
}

// ProtoToClusterNodePoolsConfig converts a ClusterNodePoolsConfig resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfig(p *containerpb.ContainerClusterNodePoolsConfig) *container.ClusterNodePoolsConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfig{
		MachineType:            dcl.StringOrNil(p.MachineType),
		DiskSizeGb:             dcl.Int64OrNil(p.DiskSizeGb),
		ServiceAccount:         dcl.StringOrNil(p.ServiceAccount),
		ImageType:              dcl.StringOrNil(p.ImageType),
		LocalSsdCount:          dcl.Int64OrNil(p.LocalSsdCount),
		Preemptible:            dcl.Bool(p.Preemptible),
		DiskType:               dcl.StringOrNil(p.DiskType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		WorkloadMetadataConfig: ProtoToContainerClusterNodePoolsConfigWorkloadMetadataConfig(p.GetWorkloadMetadataConfig()),
		SandboxConfig:          ProtoToContainerClusterNodePoolsConfigSandboxConfig(p.GetSandboxConfig()),
		NodeGroup:              dcl.StringOrNil(p.NodeGroup),
		ReservationAffinity:    ProtoToContainerClusterNodePoolsConfigReservationAffinity(p.GetReservationAffinity()),
		ShieldedInstanceConfig: ProtoToContainerClusterNodePoolsConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		LinuxNodeConfig:        ProtoToContainerClusterNodePoolsConfigLinuxNodeConfig(p.GetLinuxNodeConfig()),
		KubeletConfig:          ProtoToContainerClusterNodePoolsConfigKubeletConfig(p.GetKubeletConfig()),
		BootDiskKmsKey:         dcl.StringOrNil(p.BootDiskKmsKey),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToContainerClusterNodePoolsConfigAccelerators(r))
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerClusterNodePoolsConfigTaints(r))
	}
	return obj
}

// ProtoToClusterNodePoolsConfigAccelerators converts a ClusterNodePoolsConfigAccelerators resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfigAccelerators(p *containerpb.ContainerClusterNodePoolsConfigAccelerators) *container.ClusterNodePoolsConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfigAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigWorkloadMetadataConfig converts a ClusterNodePoolsConfigWorkloadMetadataConfig resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfigWorkloadMetadataConfig(p *containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfig) *container.ClusterNodePoolsConfigWorkloadMetadataConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfigWorkloadMetadataConfig{
		Mode: ProtoToContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigTaints converts a ClusterNodePoolsConfigTaints resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfigTaints(p *containerpb.ContainerClusterNodePoolsConfigTaints) *container.ClusterNodePoolsConfigTaints {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToContainerClusterNodePoolsConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigSandboxConfig converts a ClusterNodePoolsConfigSandboxConfig resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfigSandboxConfig(p *containerpb.ContainerClusterNodePoolsConfigSandboxConfig) *container.ClusterNodePoolsConfigSandboxConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfigSandboxConfig{
		Type: ProtoToContainerClusterNodePoolsConfigSandboxConfigTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigReservationAffinity converts a ClusterNodePoolsConfigReservationAffinity resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfigReservationAffinity(p *containerpb.ContainerClusterNodePoolsConfigReservationAffinity) *container.ClusterNodePoolsConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfigReservationAffinity{
		ConsumeReservationType: ProtoToContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterNodePoolsConfigShieldedInstanceConfig converts a ClusterNodePoolsConfigShieldedInstanceConfig resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfigShieldedInstanceConfig(p *containerpb.ContainerClusterNodePoolsConfigShieldedInstanceConfig) *container.ClusterNodePoolsConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToClusterNodePoolsConfigLinuxNodeConfig converts a ClusterNodePoolsConfigLinuxNodeConfig resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfigLinuxNodeConfig(p *containerpb.ContainerClusterNodePoolsConfigLinuxNodeConfig) *container.ClusterNodePoolsConfigLinuxNodeConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfigLinuxNodeConfig{}
	return obj
}

// ProtoToClusterNodePoolsConfigKubeletConfig converts a ClusterNodePoolsConfigKubeletConfig resource from its proto representation.
func ProtoToContainerClusterNodePoolsConfigKubeletConfig(p *containerpb.ContainerClusterNodePoolsConfigKubeletConfig) *container.ClusterNodePoolsConfigKubeletConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConfigKubeletConfig{
		CpuManagerPolicy:  dcl.StringOrNil(p.CpuManagerPolicy),
		CpuCfsQuota:       dcl.Bool(p.CpuCfsQuota),
		CpuCfsQuotaPeriod: dcl.StringOrNil(p.CpuCfsQuotaPeriod),
	}
	return obj
}

// ProtoToClusterNodePoolsAutoscaling converts a ClusterNodePoolsAutoscaling resource from its proto representation.
func ProtoToContainerClusterNodePoolsAutoscaling(p *containerpb.ContainerClusterNodePoolsAutoscaling) *container.ClusterNodePoolsAutoscaling {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsAutoscaling{
		Enabled:         dcl.Bool(p.Enabled),
		MinNodeCount:    dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount:    dcl.Int64OrNil(p.MaxNodeCount),
		Autoprovisioned: dcl.Bool(p.Autoprovisioned),
	}
	return obj
}

// ProtoToClusterNodePoolsManagement converts a ClusterNodePoolsManagement resource from its proto representation.
func ProtoToContainerClusterNodePoolsManagement(p *containerpb.ContainerClusterNodePoolsManagement) *container.ClusterNodePoolsManagement {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsManagement{
		AutoUpgrade:    dcl.Bool(p.AutoUpgrade),
		AutoRepair:     dcl.Bool(p.AutoRepair),
		UpgradeOptions: ProtoToContainerClusterNodePoolsManagementUpgradeOptions(p.GetUpgradeOptions()),
	}
	return obj
}

// ProtoToClusterNodePoolsManagementUpgradeOptions converts a ClusterNodePoolsManagementUpgradeOptions resource from its proto representation.
func ProtoToContainerClusterNodePoolsManagementUpgradeOptions(p *containerpb.ContainerClusterNodePoolsManagementUpgradeOptions) *container.ClusterNodePoolsManagementUpgradeOptions {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.StringOrNil(p.AutoUpgradeStartTime),
		Description:          dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToClusterNodePoolsMaxPodsConstraint converts a ClusterNodePoolsMaxPodsConstraint resource from its proto representation.
func ProtoToContainerClusterNodePoolsMaxPodsConstraint(p *containerpb.ContainerClusterNodePoolsMaxPodsConstraint) *container.ClusterNodePoolsMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToClusterNodePoolsConditions converts a ClusterNodePoolsConditions resource from its proto representation.
func ProtoToContainerClusterNodePoolsConditions(p *containerpb.ContainerClusterNodePoolsConditions) *container.ClusterNodePoolsConditions {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsConditions{
		Code:          ProtoToContainerClusterNodePoolsConditionsCodeEnum(p.GetCode()),
		Message:       dcl.StringOrNil(p.Message),
		CanonicalCode: ProtoToContainerClusterNodePoolsConditionsCanonicalCodeEnum(p.GetCanonicalCode()),
	}
	return obj
}

// ProtoToClusterNodePoolsUpgradeSettings converts a ClusterNodePoolsUpgradeSettings resource from its proto representation.
func ProtoToContainerClusterNodePoolsUpgradeSettings(p *containerpb.ContainerClusterNodePoolsUpgradeSettings) *container.ClusterNodePoolsUpgradeSettings {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodePoolsUpgradeSettings{
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
	}
	return obj
}

// ProtoToClusterLegacyAbac converts a ClusterLegacyAbac resource from its proto representation.
func ProtoToContainerClusterLegacyAbac(p *containerpb.ContainerClusterLegacyAbac) *container.ClusterLegacyAbac {
	if p == nil {
		return nil
	}
	obj := &container.ClusterLegacyAbac{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterNetworkPolicy converts a ClusterNetworkPolicy resource from its proto representation.
func ProtoToContainerClusterNetworkPolicy(p *containerpb.ContainerClusterNetworkPolicy) *container.ClusterNetworkPolicy {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNetworkPolicy{
		Provider: ProtoToContainerClusterNetworkPolicyProviderEnum(p.GetProvider()),
		Enabled:  dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterIPAllocationPolicy converts a ClusterIPAllocationPolicy resource from its proto representation.
func ProtoToContainerClusterIPAllocationPolicy(p *containerpb.ContainerClusterIPAllocationPolicy) *container.ClusterIPAllocationPolicy {
	if p == nil {
		return nil
	}
	obj := &container.ClusterIPAllocationPolicy{
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
	}
	return obj
}

// ProtoToClusterMasterAuthorizedNetworksConfig converts a ClusterMasterAuthorizedNetworksConfig resource from its proto representation.
func ProtoToContainerClusterMasterAuthorizedNetworksConfig(p *containerpb.ContainerClusterMasterAuthorizedNetworksConfig) *container.ClusterMasterAuthorizedNetworksConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMasterAuthorizedNetworksConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	for _, r := range p.GetCidrBlocks() {
		obj.CidrBlocks = append(obj.CidrBlocks, *ProtoToContainerClusterMasterAuthorizedNetworksConfigCidrBlocks(r))
	}
	return obj
}

// ProtoToClusterMasterAuthorizedNetworksConfigCidrBlocks converts a ClusterMasterAuthorizedNetworksConfigCidrBlocks resource from its proto representation.
func ProtoToContainerClusterMasterAuthorizedNetworksConfigCidrBlocks(p *containerpb.ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks) *container.ClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMasterAuthorizedNetworksConfigCidrBlocks{
		DisplayName: dcl.StringOrNil(p.DisplayName),
		CidrBlock:   dcl.StringOrNil(p.CidrBlock),
	}
	return obj
}

// ProtoToClusterBinaryAuthorization converts a ClusterBinaryAuthorization resource from its proto representation.
func ProtoToContainerClusterBinaryAuthorization(p *containerpb.ContainerClusterBinaryAuthorization) *container.ClusterBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &container.ClusterBinaryAuthorization{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAutoscaling converts a ClusterAutoscaling resource from its proto representation.
func ProtoToContainerClusterAutoscaling(p *containerpb.ContainerClusterAutoscaling) *container.ClusterAutoscaling {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscaling{
		EnableNodeAutoprovisioning:       dcl.Bool(p.EnableNodeAutoprovisioning),
		AutoprovisioningNodePoolDefaults: ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaults(p.GetAutoprovisioningNodePoolDefaults()),
	}
	for _, r := range p.GetResourceLimits() {
		obj.ResourceLimits = append(obj.ResourceLimits, *ProtoToContainerClusterAutoscalingResourceLimits(r))
	}
	for _, r := range p.GetAutoprovisioningLocations() {
		obj.AutoprovisioningLocations = append(obj.AutoprovisioningLocations, r)
	}
	return obj
}

// ProtoToClusterAutoscalingResourceLimits converts a ClusterAutoscalingResourceLimits resource from its proto representation.
func ProtoToContainerClusterAutoscalingResourceLimits(p *containerpb.ContainerClusterAutoscalingResourceLimits) *container.ClusterAutoscalingResourceLimits {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingResourceLimits{
		ResourceType: dcl.StringOrNil(p.ResourceType),
		Minimum:      dcl.Int64OrNil(p.Minimum),
		Maximum:      dcl.Int64OrNil(p.Maximum),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaults converts a ClusterAutoscalingAutoprovisioningNodePoolDefaults resource from its proto representation.
func ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaults(p *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaults) *container.ClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingAutoprovisioningNodePoolDefaults{
		ServiceAccount:         dcl.StringOrNil(p.ServiceAccount),
		UpgradeSettings:        ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(p.GetUpgradeSettings()),
		Management:             ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(p.GetManagement()),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		DiskSizeGb:             dcl.Int64OrNil(p.DiskSizeGb),
		DiskType:               dcl.StringOrNil(p.DiskType),
		ShieldedInstanceConfig: ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		BootDiskKmsKey:         dcl.StringOrNil(p.BootDiskKmsKey),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings resource from its proto representation.
func ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(p *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement resource from its proto representation.
func ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(p *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{
		AutoUpgrade:    dcl.Bool(p.AutoUpgrade),
		AutoRepair:     dcl.Bool(p.AutoRepair),
		UpgradeOptions: ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(p.GetUpgradeOptions()),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions resource from its proto representation.
func ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions(p *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.StringOrNil(p.AutoUpgradeStartTime),
		Description:          dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig resource from its proto representation.
func ProtoToContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig(p *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToClusterNetworkConfig converts a ClusterNetworkConfig resource from its proto representation.
func ProtoToContainerClusterNetworkConfig(p *containerpb.ContainerClusterNetworkConfig) *container.ClusterNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNetworkConfig{
		Network:                   dcl.StringOrNil(p.Network),
		Subnetwork:                dcl.StringOrNil(p.Subnetwork),
		EnableIntraNodeVisibility: dcl.Bool(p.EnableIntraNodeVisibility),
		DefaultSnatStatus:         ProtoToContainerClusterNetworkConfigDefaultSnatStatus(p.GetDefaultSnatStatus()),
		PrivateIPv6GoogleAccess:   ProtoToContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
	}
	return obj
}

// ProtoToClusterNetworkConfigDefaultSnatStatus converts a ClusterNetworkConfigDefaultSnatStatus resource from its proto representation.
func ProtoToContainerClusterNetworkConfigDefaultSnatStatus(p *containerpb.ContainerClusterNetworkConfigDefaultSnatStatus) *container.ClusterNetworkConfigDefaultSnatStatus {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNetworkConfigDefaultSnatStatus{
		Disabled: dcl.Bool(p.Disabled),
	}
	return obj
}

// ProtoToClusterMaintenancePolicy converts a ClusterMaintenancePolicy resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicy(p *containerpb.ContainerClusterMaintenancePolicy) *container.ClusterMaintenancePolicy {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicy{
		Window:          ProtoToContainerClusterMaintenancePolicyWindow(p.GetWindow()),
		ResourceVersion: dcl.StringOrNil(p.ResourceVersion),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindow converts a ClusterMaintenancePolicyWindow resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicyWindow(p *containerpb.ContainerClusterMaintenancePolicyWindow) *container.ClusterMaintenancePolicyWindow {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicyWindow{
		DailyMaintenanceWindow: ProtoToContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow(p.GetDailyMaintenanceWindow()),
		RecurringWindow:        ProtoToContainerClusterMaintenancePolicyWindowRecurringWindow(p.GetRecurringWindow()),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowDailyMaintenanceWindow converts a ClusterMaintenancePolicyWindowDailyMaintenanceWindow resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow(p *containerpb.ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow) *container.ClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicyWindowDailyMaintenanceWindow{
		StartTime: dcl.StringOrNil(p.GetStartTime()),
		Duration:  dcl.StringOrNil(p.Duration),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowRecurringWindow converts a ClusterMaintenancePolicyWindowRecurringWindow resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicyWindowRecurringWindow(p *containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindow) *container.ClusterMaintenancePolicyWindowRecurringWindow {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicyWindowRecurringWindow{
		Window:     ProtoToContainerClusterMaintenancePolicyWindowRecurringWindowWindow(p.GetWindow()),
		Recurrence: dcl.StringOrNil(p.Recurrence),
	}
	return obj
}

// ProtoToClusterMaintenancePolicyWindowRecurringWindowWindow converts a ClusterMaintenancePolicyWindowRecurringWindowWindow resource from its proto representation.
func ProtoToContainerClusterMaintenancePolicyWindowRecurringWindowWindow(p *containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindowWindow) *container.ClusterMaintenancePolicyWindowRecurringWindowWindow {
	if p == nil {
		return nil
	}
	obj := &container.ClusterMaintenancePolicyWindowRecurringWindowWindow{
		StartTime: dcl.StringOrNil(p.GetStartTime()),
		EndTime:   dcl.StringOrNil(p.GetEndTime()),
	}
	return obj
}

// ProtoToClusterDefaultMaxPodsConstraint converts a ClusterDefaultMaxPodsConstraint resource from its proto representation.
func ProtoToContainerClusterDefaultMaxPodsConstraint(p *containerpb.ContainerClusterDefaultMaxPodsConstraint) *container.ClusterDefaultMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &container.ClusterDefaultMaxPodsConstraint{
		MaxPodsPerNode: dcl.StringOrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfig converts a ClusterResourceUsageExportConfig resource from its proto representation.
func ProtoToContainerClusterResourceUsageExportConfig(p *containerpb.ContainerClusterResourceUsageExportConfig) *container.ClusterResourceUsageExportConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterResourceUsageExportConfig{
		BigqueryDestination:           ProtoToContainerClusterResourceUsageExportConfigBigqueryDestination(p.GetBigqueryDestination()),
		EnableNetworkEgressMonitoring: dcl.Bool(p.EnableNetworkEgressMonitoring),
		ConsumptionMeteringConfig:     ProtoToContainerClusterResourceUsageExportConfigConsumptionMeteringConfig(p.GetConsumptionMeteringConfig()),
		EnableNetworkEgressMetering:   dcl.Bool(p.EnableNetworkEgressMetering),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfigBigqueryDestination converts a ClusterResourceUsageExportConfigBigqueryDestination resource from its proto representation.
func ProtoToContainerClusterResourceUsageExportConfigBigqueryDestination(p *containerpb.ContainerClusterResourceUsageExportConfigBigqueryDestination) *container.ClusterResourceUsageExportConfigBigqueryDestination {
	if p == nil {
		return nil
	}
	obj := &container.ClusterResourceUsageExportConfigBigqueryDestination{
		DatasetId: dcl.StringOrNil(p.DatasetId),
	}
	return obj
}

// ProtoToClusterResourceUsageExportConfigConsumptionMeteringConfig converts a ClusterResourceUsageExportConfigConsumptionMeteringConfig resource from its proto representation.
func ProtoToContainerClusterResourceUsageExportConfigConsumptionMeteringConfig(p *containerpb.ContainerClusterResourceUsageExportConfigConsumptionMeteringConfig) *container.ClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterResourceUsageExportConfigConsumptionMeteringConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterAuthenticatorGroupsConfig converts a ClusterAuthenticatorGroupsConfig resource from its proto representation.
func ProtoToContainerClusterAuthenticatorGroupsConfig(p *containerpb.ContainerClusterAuthenticatorGroupsConfig) *container.ClusterAuthenticatorGroupsConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAuthenticatorGroupsConfig{
		Enabled:       dcl.Bool(p.Enabled),
		SecurityGroup: dcl.StringOrNil(p.SecurityGroup),
	}
	return obj
}

// ProtoToClusterPrivateClusterConfig converts a ClusterPrivateClusterConfig resource from its proto representation.
func ProtoToContainerClusterPrivateClusterConfig(p *containerpb.ContainerClusterPrivateClusterConfig) *container.ClusterPrivateClusterConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterPrivateClusterConfig{
		EnablePrivateNodes:       dcl.Bool(p.EnablePrivateNodes),
		EnablePrivateEndpoint:    dcl.Bool(p.EnablePrivateEndpoint),
		MasterIPv4CidrBlock:      dcl.StringOrNil(p.MasterIpv4CidrBlock),
		PrivateEndpoint:          dcl.StringOrNil(p.PrivateEndpoint),
		PublicEndpoint:           dcl.StringOrNil(p.PublicEndpoint),
		PeeringName:              dcl.StringOrNil(p.PeeringName),
		MasterGlobalAccessConfig: ProtoToContainerClusterPrivateClusterConfigMasterGlobalAccessConfig(p.GetMasterGlobalAccessConfig()),
	}
	return obj
}

// ProtoToClusterPrivateClusterConfigMasterGlobalAccessConfig converts a ClusterPrivateClusterConfigMasterGlobalAccessConfig resource from its proto representation.
func ProtoToContainerClusterPrivateClusterConfigMasterGlobalAccessConfig(p *containerpb.ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig) *container.ClusterPrivateClusterConfigMasterGlobalAccessConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterPrivateClusterConfigMasterGlobalAccessConfig{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterDatabaseEncryption converts a ClusterDatabaseEncryption resource from its proto representation.
func ProtoToContainerClusterDatabaseEncryption(p *containerpb.ContainerClusterDatabaseEncryption) *container.ClusterDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &container.ClusterDatabaseEncryption{
		State:   ProtoToContainerClusterDatabaseEncryptionStateEnum(p.GetState()),
		KeyName: dcl.StringOrNil(p.KeyName),
	}
	return obj
}

// ProtoToClusterVerticalPodAutoscaling converts a ClusterVerticalPodAutoscaling resource from its proto representation.
func ProtoToContainerClusterVerticalPodAutoscaling(p *containerpb.ContainerClusterVerticalPodAutoscaling) *container.ClusterVerticalPodAutoscaling {
	if p == nil {
		return nil
	}
	obj := &container.ClusterVerticalPodAutoscaling{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterShieldedNodes converts a ClusterShieldedNodes resource from its proto representation.
func ProtoToContainerClusterShieldedNodes(p *containerpb.ContainerClusterShieldedNodes) *container.ClusterShieldedNodes {
	if p == nil {
		return nil
	}
	obj := &container.ClusterShieldedNodes{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterConditions converts a ClusterConditions resource from its proto representation.
func ProtoToContainerClusterConditions(p *containerpb.ContainerClusterConditions) *container.ClusterConditions {
	if p == nil {
		return nil
	}
	obj := &container.ClusterConditions{
		Code:          dcl.StringOrNil(p.Code),
		Message:       dcl.StringOrNil(p.Message),
		CanonicalCode: ProtoToContainerClusterConditionsCanonicalCodeEnum(p.GetCanonicalCode()),
	}
	return obj
}

// ProtoToClusterAutopilot converts a ClusterAutopilot resource from its proto representation.
func ProtoToContainerClusterAutopilot(p *containerpb.ContainerClusterAutopilot) *container.ClusterAutopilot {
	if p == nil {
		return nil
	}
	obj := &container.ClusterAutopilot{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToClusterNodeConfig converts a ClusterNodeConfig resource from its proto representation.
func ProtoToContainerClusterNodeConfig(p *containerpb.ContainerClusterNodeConfig) *container.ClusterNodeConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfig{
		MachineType:            dcl.StringOrNil(p.MachineType),
		DiskSizeGb:             dcl.Int64OrNil(p.DiskSizeGb),
		ServiceAccount:         dcl.StringOrNil(p.ServiceAccount),
		ImageType:              dcl.StringOrNil(p.ImageType),
		LocalSsdCount:          dcl.Int64OrNil(p.LocalSsdCount),
		Preemptible:            dcl.Bool(p.Preemptible),
		DiskType:               dcl.StringOrNil(p.DiskType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		WorkloadMetadataConfig: ProtoToContainerClusterNodeConfigWorkloadMetadataConfig(p.GetWorkloadMetadataConfig()),
		SandboxConfig:          ProtoToContainerClusterNodeConfigSandboxConfig(p.GetSandboxConfig()),
		NodeGroup:              dcl.StringOrNil(p.NodeGroup),
		ReservationAffinity:    ProtoToContainerClusterNodeConfigReservationAffinity(p.GetReservationAffinity()),
		ShieldedInstanceConfig: ProtoToContainerClusterNodeConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		LinuxNodeConfig:        ProtoToContainerClusterNodeConfigLinuxNodeConfig(p.GetLinuxNodeConfig()),
		KubeletConfig:          ProtoToContainerClusterNodeConfigKubeletConfig(p.GetKubeletConfig()),
		BootDiskKmsKey:         dcl.StringOrNil(p.BootDiskKmsKey),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToContainerClusterNodeConfigAccelerators(r))
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerClusterNodeConfigTaints(r))
	}
	return obj
}

// ProtoToClusterNodeConfigAccelerators converts a ClusterNodeConfigAccelerators resource from its proto representation.
func ProtoToContainerClusterNodeConfigAccelerators(p *containerpb.ContainerClusterNodeConfigAccelerators) *container.ClusterNodeConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfigAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToClusterNodeConfigWorkloadMetadataConfig converts a ClusterNodeConfigWorkloadMetadataConfig resource from its proto representation.
func ProtoToContainerClusterNodeConfigWorkloadMetadataConfig(p *containerpb.ContainerClusterNodeConfigWorkloadMetadataConfig) *container.ClusterNodeConfigWorkloadMetadataConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfigWorkloadMetadataConfig{
		Mode: ProtoToContainerClusterNodeConfigWorkloadMetadataConfigModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToClusterNodeConfigTaints converts a ClusterNodeConfigTaints resource from its proto representation.
func ProtoToContainerClusterNodeConfigTaints(p *containerpb.ContainerClusterNodeConfigTaints) *container.ClusterNodeConfigTaints {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: ProtoToContainerClusterNodeConfigTaintsEffectEnum(p.GetEffect()),
	}
	return obj
}

// ProtoToClusterNodeConfigSandboxConfig converts a ClusterNodeConfigSandboxConfig resource from its proto representation.
func ProtoToContainerClusterNodeConfigSandboxConfig(p *containerpb.ContainerClusterNodeConfigSandboxConfig) *container.ClusterNodeConfigSandboxConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfigSandboxConfig{
		Type: ProtoToContainerClusterNodeConfigSandboxConfigTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToClusterNodeConfigReservationAffinity converts a ClusterNodeConfigReservationAffinity resource from its proto representation.
func ProtoToContainerClusterNodeConfigReservationAffinity(p *containerpb.ContainerClusterNodeConfigReservationAffinity) *container.ClusterNodeConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfigReservationAffinity{
		ConsumeReservationType: ProtoToContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterNodeConfigShieldedInstanceConfig converts a ClusterNodeConfigShieldedInstanceConfig resource from its proto representation.
func ProtoToContainerClusterNodeConfigShieldedInstanceConfig(p *containerpb.ContainerClusterNodeConfigShieldedInstanceConfig) *container.ClusterNodeConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToClusterNodeConfigLinuxNodeConfig converts a ClusterNodeConfigLinuxNodeConfig resource from its proto representation.
func ProtoToContainerClusterNodeConfigLinuxNodeConfig(p *containerpb.ContainerClusterNodeConfigLinuxNodeConfig) *container.ClusterNodeConfigLinuxNodeConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfigLinuxNodeConfig{}
	return obj
}

// ProtoToClusterNodeConfigKubeletConfig converts a ClusterNodeConfigKubeletConfig resource from its proto representation.
func ProtoToContainerClusterNodeConfigKubeletConfig(p *containerpb.ContainerClusterNodeConfigKubeletConfig) *container.ClusterNodeConfigKubeletConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNodeConfigKubeletConfig{
		CpuManagerPolicy:  dcl.StringOrNil(p.CpuManagerPolicy),
		CpuCfsQuota:       dcl.Bool(p.CpuCfsQuota),
		CpuCfsQuotaPeriod: dcl.StringOrNil(p.CpuCfsQuotaPeriod),
	}
	return obj
}

// ProtoToClusterReleaseChannel converts a ClusterReleaseChannel resource from its proto representation.
func ProtoToContainerClusterReleaseChannel(p *containerpb.ContainerClusterReleaseChannel) *container.ClusterReleaseChannel {
	if p == nil {
		return nil
	}
	obj := &container.ClusterReleaseChannel{
		Channel: ProtoToContainerClusterReleaseChannelChannelEnum(p.GetChannel()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToContainerClusterWorkloadIdentityConfig(p *containerpb.ContainerClusterWorkloadIdentityConfig) *container.ClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterWorkloadIdentityConfig{
		WorkloadPool: dcl.StringOrNil(p.WorkloadPool),
	}
	return obj
}

// ProtoToClusterNotificationConfig converts a ClusterNotificationConfig resource from its proto representation.
func ProtoToContainerClusterNotificationConfig(p *containerpb.ContainerClusterNotificationConfig) *container.ClusterNotificationConfig {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNotificationConfig{
		Pubsub: ProtoToContainerClusterNotificationConfigPubsub(p.GetPubsub()),
	}
	return obj
}

// ProtoToClusterNotificationConfigPubsub converts a ClusterNotificationConfigPubsub resource from its proto representation.
func ProtoToContainerClusterNotificationConfigPubsub(p *containerpb.ContainerClusterNotificationConfigPubsub) *container.ClusterNotificationConfigPubsub {
	if p == nil {
		return nil
	}
	obj := &container.ClusterNotificationConfigPubsub{
		Enabled: dcl.Bool(p.Enabled),
		Topic:   dcl.StringOrNil(p.Topic),
	}
	return obj
}

// ProtoToClusterConfidentialNodes converts a ClusterConfidentialNodes resource from its proto representation.
func ProtoToContainerClusterConfidentialNodes(p *containerpb.ContainerClusterConfidentialNodes) *container.ClusterConfidentialNodes {
	if p == nil {
		return nil
	}
	obj := &container.ClusterConfidentialNodes{
		Enabled: dcl.Bool(p.Enabled),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *containerpb.ContainerCluster) *container.Cluster {
	obj := &container.Cluster{
		Name:                           dcl.StringOrNil(p.Name),
		Description:                    dcl.StringOrNil(p.Description),
		InitialNodeCount:               dcl.Int64OrNil(p.InitialNodeCount),
		MasterAuth:                     ProtoToContainerClusterMasterAuth(p.GetMasterAuth()),
		LoggingService:                 dcl.StringOrNil(p.LoggingService),
		MonitoringService:              dcl.StringOrNil(p.MonitoringService),
		Network:                        dcl.StringOrNil(p.Network),
		ClusterIPv4Cidr:                dcl.StringOrNil(p.ClusterIpv4Cidr),
		AddonsConfig:                   ProtoToContainerClusterAddonsConfig(p.GetAddonsConfig()),
		Subnetwork:                     dcl.StringOrNil(p.Subnetwork),
		EnableKubernetesAlpha:          dcl.Bool(p.EnableKubernetesAlpha),
		LabelFingerprint:               dcl.StringOrNil(p.LabelFingerprint),
		LegacyAbac:                     ProtoToContainerClusterLegacyAbac(p.GetLegacyAbac()),
		NetworkPolicy:                  ProtoToContainerClusterNetworkPolicy(p.GetNetworkPolicy()),
		IPAllocationPolicy:             ProtoToContainerClusterIPAllocationPolicy(p.GetIpAllocationPolicy()),
		MasterAuthorizedNetworksConfig: ProtoToContainerClusterMasterAuthorizedNetworksConfig(p.GetMasterAuthorizedNetworksConfig()),
		BinaryAuthorization:            ProtoToContainerClusterBinaryAuthorization(p.GetBinaryAuthorization()),
		Autoscaling:                    ProtoToContainerClusterAutoscaling(p.GetAutoscaling()),
		NetworkConfig:                  ProtoToContainerClusterNetworkConfig(p.GetNetworkConfig()),
		MaintenancePolicy:              ProtoToContainerClusterMaintenancePolicy(p.GetMaintenancePolicy()),
		DefaultMaxPodsConstraint:       ProtoToContainerClusterDefaultMaxPodsConstraint(p.GetDefaultMaxPodsConstraint()),
		ResourceUsageExportConfig:      ProtoToContainerClusterResourceUsageExportConfig(p.GetResourceUsageExportConfig()),
		AuthenticatorGroupsConfig:      ProtoToContainerClusterAuthenticatorGroupsConfig(p.GetAuthenticatorGroupsConfig()),
		PrivateClusterConfig:           ProtoToContainerClusterPrivateClusterConfig(p.GetPrivateClusterConfig()),
		DatabaseEncryption:             ProtoToContainerClusterDatabaseEncryption(p.GetDatabaseEncryption()),
		VerticalPodAutoscaling:         ProtoToContainerClusterVerticalPodAutoscaling(p.GetVerticalPodAutoscaling()),
		ShieldedNodes:                  ProtoToContainerClusterShieldedNodes(p.GetShieldedNodes()),
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
		Autopilot:                      ProtoToContainerClusterAutopilot(p.GetAutopilot()),
		Project:                        dcl.StringOrNil(p.Project),
		NodeConfig:                     ProtoToContainerClusterNodeConfig(p.GetNodeConfig()),
		ReleaseChannel:                 ProtoToContainerClusterReleaseChannel(p.GetReleaseChannel()),
		WorkloadIdentityConfig:         ProtoToContainerClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		NotificationConfig:             ProtoToContainerClusterNotificationConfig(p.GetNotificationConfig()),
		ConfidentialNodes:              ProtoToContainerClusterConfidentialNodes(p.GetConfidentialNodes()),
		SelfLink:                       dcl.StringOrNil(p.SelfLink),
		Zone:                           dcl.StringOrNil(p.Zone),
		InitialClusterVersion:          dcl.StringOrNil(p.InitialClusterVersion),
		CurrentMasterVersion:           dcl.StringOrNil(p.CurrentMasterVersion),
		CurrentNodeVersion:             dcl.StringOrNil(p.CurrentNodeVersion),
		CurrentNodeCount:               dcl.Int64OrNil(p.CurrentNodeCount),
		Id:                             dcl.StringOrNil(p.Id),
	}
	for _, r := range p.GetNodePools() {
		obj.NodePools = append(obj.NodePools, *ProtoToContainerClusterNodePools(r))
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToContainerClusterConditions(r))
	}
	for _, r := range p.GetInstanceGroupUrls() {
		obj.InstanceGroupUrls = append(obj.InstanceGroupUrls, r)
	}
	return obj
}

// ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumToProto converts a ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum enum to its proto representation.
func ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumToProto(e *container.ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum) containerpb.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum {
	if e == nil {
		return containerpb.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum_value["ClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(v)
	}
	return containerpb.ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnum(0)
}

// ClusterNodePoolsConfigWorkloadMetadataConfigModeEnumToProto converts a ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum enum to its proto representation.
func ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnumToProto(e *container.ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum) containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum {
	if e == nil {
		return containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum_value["ClusterNodePoolsConfigWorkloadMetadataConfigModeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(v)
	}
	return containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnum(0)
}

// ClusterNodePoolsConfigTaintsEffectEnumToProto converts a ClusterNodePoolsConfigTaintsEffectEnum enum to its proto representation.
func ContainerClusterNodePoolsConfigTaintsEffectEnumToProto(e *container.ClusterNodePoolsConfigTaintsEffectEnum) containerpb.ContainerClusterNodePoolsConfigTaintsEffectEnum {
	if e == nil {
		return containerpb.ContainerClusterNodePoolsConfigTaintsEffectEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodePoolsConfigTaintsEffectEnum_value["ClusterNodePoolsConfigTaintsEffectEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodePoolsConfigTaintsEffectEnum(v)
	}
	return containerpb.ContainerClusterNodePoolsConfigTaintsEffectEnum(0)
}

// ClusterNodePoolsConfigSandboxConfigTypeEnumToProto converts a ClusterNodePoolsConfigSandboxConfigTypeEnum enum to its proto representation.
func ContainerClusterNodePoolsConfigSandboxConfigTypeEnumToProto(e *container.ClusterNodePoolsConfigSandboxConfigTypeEnum) containerpb.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum {
	if e == nil {
		return containerpb.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum_value["ClusterNodePoolsConfigSandboxConfigTypeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum(v)
	}
	return containerpb.ContainerClusterNodePoolsConfigSandboxConfigTypeEnum(0)
}

// ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumToProto(e *container.ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum) containerpb.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return containerpb.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return containerpb.ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterNodePoolsStatusEnumToProto converts a ClusterNodePoolsStatusEnum enum to its proto representation.
func ContainerClusterNodePoolsStatusEnumToProto(e *container.ClusterNodePoolsStatusEnum) containerpb.ContainerClusterNodePoolsStatusEnum {
	if e == nil {
		return containerpb.ContainerClusterNodePoolsStatusEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodePoolsStatusEnum_value["ClusterNodePoolsStatusEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodePoolsStatusEnum(v)
	}
	return containerpb.ContainerClusterNodePoolsStatusEnum(0)
}

// ClusterNodePoolsConditionsCodeEnumToProto converts a ClusterNodePoolsConditionsCodeEnum enum to its proto representation.
func ContainerClusterNodePoolsConditionsCodeEnumToProto(e *container.ClusterNodePoolsConditionsCodeEnum) containerpb.ContainerClusterNodePoolsConditionsCodeEnum {
	if e == nil {
		return containerpb.ContainerClusterNodePoolsConditionsCodeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodePoolsConditionsCodeEnum_value["ClusterNodePoolsConditionsCodeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodePoolsConditionsCodeEnum(v)
	}
	return containerpb.ContainerClusterNodePoolsConditionsCodeEnum(0)
}

// ClusterNodePoolsConditionsCanonicalCodeEnumToProto converts a ClusterNodePoolsConditionsCanonicalCodeEnum enum to its proto representation.
func ContainerClusterNodePoolsConditionsCanonicalCodeEnumToProto(e *container.ClusterNodePoolsConditionsCanonicalCodeEnum) containerpb.ContainerClusterNodePoolsConditionsCanonicalCodeEnum {
	if e == nil {
		return containerpb.ContainerClusterNodePoolsConditionsCanonicalCodeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodePoolsConditionsCanonicalCodeEnum_value["ClusterNodePoolsConditionsCanonicalCodeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodePoolsConditionsCanonicalCodeEnum(v)
	}
	return containerpb.ContainerClusterNodePoolsConditionsCanonicalCodeEnum(0)
}

// ClusterNetworkPolicyProviderEnumToProto converts a ClusterNetworkPolicyProviderEnum enum to its proto representation.
func ContainerClusterNetworkPolicyProviderEnumToProto(e *container.ClusterNetworkPolicyProviderEnum) containerpb.ContainerClusterNetworkPolicyProviderEnum {
	if e == nil {
		return containerpb.ContainerClusterNetworkPolicyProviderEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNetworkPolicyProviderEnum_value["ClusterNetworkPolicyProviderEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNetworkPolicyProviderEnum(v)
	}
	return containerpb.ContainerClusterNetworkPolicyProviderEnum(0)
}

// ClusterNetworkConfigPrivateIPv6GoogleAccessEnumToProto converts a ClusterNetworkConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnumToProto(e *container.ClusterNetworkConfigPrivateIPv6GoogleAccessEnum) containerpb.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return containerpb.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum_value["ClusterNetworkConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return containerpb.ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnum(0)
}

// ClusterDatabaseEncryptionStateEnumToProto converts a ClusterDatabaseEncryptionStateEnum enum to its proto representation.
func ContainerClusterDatabaseEncryptionStateEnumToProto(e *container.ClusterDatabaseEncryptionStateEnum) containerpb.ContainerClusterDatabaseEncryptionStateEnum {
	if e == nil {
		return containerpb.ContainerClusterDatabaseEncryptionStateEnum(0)
	}
	if v, ok := containerpb.ContainerClusterDatabaseEncryptionStateEnum_value["ClusterDatabaseEncryptionStateEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterDatabaseEncryptionStateEnum(v)
	}
	return containerpb.ContainerClusterDatabaseEncryptionStateEnum(0)
}

// ClusterConditionsCanonicalCodeEnumToProto converts a ClusterConditionsCanonicalCodeEnum enum to its proto representation.
func ContainerClusterConditionsCanonicalCodeEnumToProto(e *container.ClusterConditionsCanonicalCodeEnum) containerpb.ContainerClusterConditionsCanonicalCodeEnum {
	if e == nil {
		return containerpb.ContainerClusterConditionsCanonicalCodeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterConditionsCanonicalCodeEnum_value["ClusterConditionsCanonicalCodeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterConditionsCanonicalCodeEnum(v)
	}
	return containerpb.ContainerClusterConditionsCanonicalCodeEnum(0)
}

// ClusterNodeConfigWorkloadMetadataConfigModeEnumToProto converts a ClusterNodeConfigWorkloadMetadataConfigModeEnum enum to its proto representation.
func ContainerClusterNodeConfigWorkloadMetadataConfigModeEnumToProto(e *container.ClusterNodeConfigWorkloadMetadataConfigModeEnum) containerpb.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum {
	if e == nil {
		return containerpb.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum_value["ClusterNodeConfigWorkloadMetadataConfigModeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum(v)
	}
	return containerpb.ContainerClusterNodeConfigWorkloadMetadataConfigModeEnum(0)
}

// ClusterNodeConfigTaintsEffectEnumToProto converts a ClusterNodeConfigTaintsEffectEnum enum to its proto representation.
func ContainerClusterNodeConfigTaintsEffectEnumToProto(e *container.ClusterNodeConfigTaintsEffectEnum) containerpb.ContainerClusterNodeConfigTaintsEffectEnum {
	if e == nil {
		return containerpb.ContainerClusterNodeConfigTaintsEffectEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodeConfigTaintsEffectEnum_value["ClusterNodeConfigTaintsEffectEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodeConfigTaintsEffectEnum(v)
	}
	return containerpb.ContainerClusterNodeConfigTaintsEffectEnum(0)
}

// ClusterNodeConfigSandboxConfigTypeEnumToProto converts a ClusterNodeConfigSandboxConfigTypeEnum enum to its proto representation.
func ContainerClusterNodeConfigSandboxConfigTypeEnumToProto(e *container.ClusterNodeConfigSandboxConfigTypeEnum) containerpb.ContainerClusterNodeConfigSandboxConfigTypeEnum {
	if e == nil {
		return containerpb.ContainerClusterNodeConfigSandboxConfigTypeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodeConfigSandboxConfigTypeEnum_value["ClusterNodeConfigSandboxConfigTypeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodeConfigSandboxConfigTypeEnum(v)
	}
	return containerpb.ContainerClusterNodeConfigSandboxConfigTypeEnum(0)
}

// ClusterNodeConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnumToProto(e *container.ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum) containerpb.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return containerpb.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := containerpb.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterNodeConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return containerpb.ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterReleaseChannelChannelEnumToProto converts a ClusterReleaseChannelChannelEnum enum to its proto representation.
func ContainerClusterReleaseChannelChannelEnumToProto(e *container.ClusterReleaseChannelChannelEnum) containerpb.ContainerClusterReleaseChannelChannelEnum {
	if e == nil {
		return containerpb.ContainerClusterReleaseChannelChannelEnum(0)
	}
	if v, ok := containerpb.ContainerClusterReleaseChannelChannelEnum_value["ClusterReleaseChannelChannelEnum"+string(*e)]; ok {
		return containerpb.ContainerClusterReleaseChannelChannelEnum(v)
	}
	return containerpb.ContainerClusterReleaseChannelChannelEnum(0)
}

// ClusterMasterAuthToProto converts a ClusterMasterAuth resource to its proto representation.
func ContainerClusterMasterAuthToProto(o *container.ClusterMasterAuth) *containerpb.ContainerClusterMasterAuth {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMasterAuth{
		Username:                dcl.ValueOrEmptyString(o.Username),
		Password:                dcl.ValueOrEmptyString(o.Password),
		ClientCertificateConfig: ContainerClusterMasterAuthClientCertificateConfigToProto(o.ClientCertificateConfig),
		ClusterCaCertificate:    dcl.ValueOrEmptyString(o.ClusterCaCertificate),
		ClientCertificate:       dcl.ValueOrEmptyString(o.ClientCertificate),
		ClientKey:               dcl.ValueOrEmptyString(o.ClientKey),
	}
	return p
}

// ClusterMasterAuthClientCertificateConfigToProto converts a ClusterMasterAuthClientCertificateConfig resource to its proto representation.
func ContainerClusterMasterAuthClientCertificateConfigToProto(o *container.ClusterMasterAuthClientCertificateConfig) *containerpb.ContainerClusterMasterAuthClientCertificateConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMasterAuthClientCertificateConfig{
		IssueClientCertificate: dcl.ValueOrEmptyBool(o.IssueClientCertificate),
	}
	return p
}

// ClusterAddonsConfigToProto converts a ClusterAddonsConfig resource to its proto representation.
func ContainerClusterAddonsConfigToProto(o *container.ClusterAddonsConfig) *containerpb.ContainerClusterAddonsConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfig{
		HttpLoadBalancing:                ContainerClusterAddonsConfigHttpLoadBalancingToProto(o.HttpLoadBalancing),
		HorizontalPodAutoscaling:         ContainerClusterAddonsConfigHorizontalPodAutoscalingToProto(o.HorizontalPodAutoscaling),
		KubernetesDashboard:              ContainerClusterAddonsConfigKubernetesDashboardToProto(o.KubernetesDashboard),
		NetworkPolicyConfig:              ContainerClusterAddonsConfigNetworkPolicyConfigToProto(o.NetworkPolicyConfig),
		CloudRunConfig:                   ContainerClusterAddonsConfigCloudRunConfigToProto(o.CloudRunConfig),
		DnsCacheConfig:                   ContainerClusterAddonsConfigDnsCacheConfigToProto(o.DnsCacheConfig),
		ConfigConnectorConfig:            ContainerClusterAddonsConfigConfigConnectorConfigToProto(o.ConfigConnectorConfig),
		GcePersistentDiskCsiDriverConfig: ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfigToProto(o.GcePersistentDiskCsiDriverConfig),
	}
	return p
}

// ClusterAddonsConfigHttpLoadBalancingToProto converts a ClusterAddonsConfigHttpLoadBalancing resource to its proto representation.
func ContainerClusterAddonsConfigHttpLoadBalancingToProto(o *container.ClusterAddonsConfigHttpLoadBalancing) *containerpb.ContainerClusterAddonsConfigHttpLoadBalancing {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigHttpLoadBalancing{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigHorizontalPodAutoscalingToProto converts a ClusterAddonsConfigHorizontalPodAutoscaling resource to its proto representation.
func ContainerClusterAddonsConfigHorizontalPodAutoscalingToProto(o *container.ClusterAddonsConfigHorizontalPodAutoscaling) *containerpb.ContainerClusterAddonsConfigHorizontalPodAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigHorizontalPodAutoscaling{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigKubernetesDashboardToProto converts a ClusterAddonsConfigKubernetesDashboard resource to its proto representation.
func ContainerClusterAddonsConfigKubernetesDashboardToProto(o *container.ClusterAddonsConfigKubernetesDashboard) *containerpb.ContainerClusterAddonsConfigKubernetesDashboard {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigKubernetesDashboard{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigNetworkPolicyConfigToProto converts a ClusterAddonsConfigNetworkPolicyConfig resource to its proto representation.
func ContainerClusterAddonsConfigNetworkPolicyConfigToProto(o *container.ClusterAddonsConfigNetworkPolicyConfig) *containerpb.ContainerClusterAddonsConfigNetworkPolicyConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigNetworkPolicyConfig{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterAddonsConfigCloudRunConfigToProto converts a ClusterAddonsConfigCloudRunConfig resource to its proto representation.
func ContainerClusterAddonsConfigCloudRunConfigToProto(o *container.ClusterAddonsConfigCloudRunConfig) *containerpb.ContainerClusterAddonsConfigCloudRunConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigCloudRunConfig{
		Disabled:         dcl.ValueOrEmptyBool(o.Disabled),
		LoadBalancerType: ContainerClusterAddonsConfigCloudRunConfigLoadBalancerTypeEnumToProto(o.LoadBalancerType),
	}
	return p
}

// ClusterAddonsConfigDnsCacheConfigToProto converts a ClusterAddonsConfigDnsCacheConfig resource to its proto representation.
func ContainerClusterAddonsConfigDnsCacheConfigToProto(o *container.ClusterAddonsConfigDnsCacheConfig) *containerpb.ContainerClusterAddonsConfigDnsCacheConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigDnsCacheConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAddonsConfigConfigConnectorConfigToProto converts a ClusterAddonsConfigConfigConnectorConfig resource to its proto representation.
func ContainerClusterAddonsConfigConfigConnectorConfigToProto(o *container.ClusterAddonsConfigConfigConnectorConfig) *containerpb.ContainerClusterAddonsConfigConfigConnectorConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigConfigConnectorConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAddonsConfigGcePersistentDiskCsiDriverConfigToProto converts a ClusterAddonsConfigGcePersistentDiskCsiDriverConfig resource to its proto representation.
func ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfigToProto(o *container.ClusterAddonsConfigGcePersistentDiskCsiDriverConfig) *containerpb.ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterNodePoolsToProto converts a ClusterNodePools resource to its proto representation.
func ContainerClusterNodePoolsToProto(o *container.ClusterNodePools) *containerpb.ContainerClusterNodePools {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePools{
		Name:              dcl.ValueOrEmptyString(o.Name),
		Config:            ContainerClusterNodePoolsConfigToProto(o.Config),
		InitialNodeCount:  dcl.ValueOrEmptyInt64(o.InitialNodeCount),
		SelfLink:          dcl.ValueOrEmptyString(o.SelfLink),
		Version:           dcl.ValueOrEmptyString(o.Version),
		Status:            ContainerClusterNodePoolsStatusEnumToProto(o.Status),
		StatusMessage:     dcl.ValueOrEmptyString(o.StatusMessage),
		Autoscaling:       ContainerClusterNodePoolsAutoscalingToProto(o.Autoscaling),
		Management:        ContainerClusterNodePoolsManagementToProto(o.Management),
		MaxPodsConstraint: ContainerClusterNodePoolsMaxPodsConstraintToProto(o.MaxPodsConstraint),
		PodIpv4CidrSize:   dcl.ValueOrEmptyInt64(o.PodIPv4CidrSize),
		UpgradeSettings:   ContainerClusterNodePoolsUpgradeSettingsToProto(o.UpgradeSettings),
	}
	for _, r := range o.Locations {
		p.Locations = append(p.Locations, r)
	}
	for _, r := range o.InstanceGroupUrls {
		p.InstanceGroupUrls = append(p.InstanceGroupUrls, r)
	}
	for _, r := range o.Conditions {
		p.Conditions = append(p.Conditions, ContainerClusterNodePoolsConditionsToProto(&r))
	}
	return p
}

// ClusterNodePoolsConfigToProto converts a ClusterNodePoolsConfig resource to its proto representation.
func ContainerClusterNodePoolsConfigToProto(o *container.ClusterNodePoolsConfig) *containerpb.ContainerClusterNodePoolsConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfig{
		MachineType:            dcl.ValueOrEmptyString(o.MachineType),
		DiskSizeGb:             dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		ServiceAccount:         dcl.ValueOrEmptyString(o.ServiceAccount),
		ImageType:              dcl.ValueOrEmptyString(o.ImageType),
		LocalSsdCount:          dcl.ValueOrEmptyInt64(o.LocalSsdCount),
		Preemptible:            dcl.ValueOrEmptyBool(o.Preemptible),
		DiskType:               dcl.ValueOrEmptyString(o.DiskType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		WorkloadMetadataConfig: ContainerClusterNodePoolsConfigWorkloadMetadataConfigToProto(o.WorkloadMetadataConfig),
		SandboxConfig:          ContainerClusterNodePoolsConfigSandboxConfigToProto(o.SandboxConfig),
		NodeGroup:              dcl.ValueOrEmptyString(o.NodeGroup),
		ReservationAffinity:    ContainerClusterNodePoolsConfigReservationAffinityToProto(o.ReservationAffinity),
		ShieldedInstanceConfig: ContainerClusterNodePoolsConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
		LinuxNodeConfig:        ContainerClusterNodePoolsConfigLinuxNodeConfigToProto(o.LinuxNodeConfig),
		KubeletConfig:          ContainerClusterNodePoolsConfigKubeletConfigToProto(o.KubeletConfig),
		BootDiskKmsKey:         dcl.ValueOrEmptyString(o.BootDiskKmsKey),
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
		p.Accelerators = append(p.Accelerators, ContainerClusterNodePoolsConfigAcceleratorsToProto(&r))
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, ContainerClusterNodePoolsConfigTaintsToProto(&r))
	}
	return p
}

// ClusterNodePoolsConfigAcceleratorsToProto converts a ClusterNodePoolsConfigAccelerators resource to its proto representation.
func ContainerClusterNodePoolsConfigAcceleratorsToProto(o *container.ClusterNodePoolsConfigAccelerators) *containerpb.ContainerClusterNodePoolsConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfigAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// ClusterNodePoolsConfigWorkloadMetadataConfigToProto converts a ClusterNodePoolsConfigWorkloadMetadataConfig resource to its proto representation.
func ContainerClusterNodePoolsConfigWorkloadMetadataConfigToProto(o *container.ClusterNodePoolsConfigWorkloadMetadataConfig) *containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfigWorkloadMetadataConfig{
		Mode: ContainerClusterNodePoolsConfigWorkloadMetadataConfigModeEnumToProto(o.Mode),
	}
	return p
}

// ClusterNodePoolsConfigTaintsToProto converts a ClusterNodePoolsConfigTaints resource to its proto representation.
func ContainerClusterNodePoolsConfigTaintsToProto(o *container.ClusterNodePoolsConfigTaints) *containerpb.ContainerClusterNodePoolsConfigTaints {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: ContainerClusterNodePoolsConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// ClusterNodePoolsConfigSandboxConfigToProto converts a ClusterNodePoolsConfigSandboxConfig resource to its proto representation.
func ContainerClusterNodePoolsConfigSandboxConfigToProto(o *container.ClusterNodePoolsConfigSandboxConfig) *containerpb.ContainerClusterNodePoolsConfigSandboxConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfigSandboxConfig{
		Type: ContainerClusterNodePoolsConfigSandboxConfigTypeEnumToProto(o.Type),
	}
	return p
}

// ClusterNodePoolsConfigReservationAffinityToProto converts a ClusterNodePoolsConfigReservationAffinity resource to its proto representation.
func ContainerClusterNodePoolsConfigReservationAffinityToProto(o *container.ClusterNodePoolsConfigReservationAffinity) *containerpb.ContainerClusterNodePoolsConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfigReservationAffinity{
		ConsumeReservationType: ContainerClusterNodePoolsConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// ClusterNodePoolsConfigShieldedInstanceConfigToProto converts a ClusterNodePoolsConfigShieldedInstanceConfig resource to its proto representation.
func ContainerClusterNodePoolsConfigShieldedInstanceConfigToProto(o *container.ClusterNodePoolsConfigShieldedInstanceConfig) *containerpb.ContainerClusterNodePoolsConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// ClusterNodePoolsConfigLinuxNodeConfigToProto converts a ClusterNodePoolsConfigLinuxNodeConfig resource to its proto representation.
func ContainerClusterNodePoolsConfigLinuxNodeConfigToProto(o *container.ClusterNodePoolsConfigLinuxNodeConfig) *containerpb.ContainerClusterNodePoolsConfigLinuxNodeConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfigLinuxNodeConfig{}
	p.Sysctls = make(map[string]string)
	for k, r := range o.Sysctls {
		p.Sysctls[k] = r
	}
	return p
}

// ClusterNodePoolsConfigKubeletConfigToProto converts a ClusterNodePoolsConfigKubeletConfig resource to its proto representation.
func ContainerClusterNodePoolsConfigKubeletConfigToProto(o *container.ClusterNodePoolsConfigKubeletConfig) *containerpb.ContainerClusterNodePoolsConfigKubeletConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConfigKubeletConfig{
		CpuManagerPolicy:  dcl.ValueOrEmptyString(o.CpuManagerPolicy),
		CpuCfsQuota:       dcl.ValueOrEmptyBool(o.CpuCfsQuota),
		CpuCfsQuotaPeriod: dcl.ValueOrEmptyString(o.CpuCfsQuotaPeriod),
	}
	return p
}

// ClusterNodePoolsAutoscalingToProto converts a ClusterNodePoolsAutoscaling resource to its proto representation.
func ContainerClusterNodePoolsAutoscalingToProto(o *container.ClusterNodePoolsAutoscaling) *containerpb.ContainerClusterNodePoolsAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsAutoscaling{
		Enabled:         dcl.ValueOrEmptyBool(o.Enabled),
		MinNodeCount:    dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount:    dcl.ValueOrEmptyInt64(o.MaxNodeCount),
		Autoprovisioned: dcl.ValueOrEmptyBool(o.Autoprovisioned),
	}
	return p
}

// ClusterNodePoolsManagementToProto converts a ClusterNodePoolsManagement resource to its proto representation.
func ContainerClusterNodePoolsManagementToProto(o *container.ClusterNodePoolsManagement) *containerpb.ContainerClusterNodePoolsManagement {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsManagement{
		AutoUpgrade:    dcl.ValueOrEmptyBool(o.AutoUpgrade),
		AutoRepair:     dcl.ValueOrEmptyBool(o.AutoRepair),
		UpgradeOptions: ContainerClusterNodePoolsManagementUpgradeOptionsToProto(o.UpgradeOptions),
	}
	return p
}

// ClusterNodePoolsManagementUpgradeOptionsToProto converts a ClusterNodePoolsManagementUpgradeOptions resource to its proto representation.
func ContainerClusterNodePoolsManagementUpgradeOptionsToProto(o *container.ClusterNodePoolsManagementUpgradeOptions) *containerpb.ContainerClusterNodePoolsManagementUpgradeOptions {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.ValueOrEmptyString(o.AutoUpgradeStartTime),
		Description:          dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// ClusterNodePoolsMaxPodsConstraintToProto converts a ClusterNodePoolsMaxPodsConstraint resource to its proto representation.
func ContainerClusterNodePoolsMaxPodsConstraintToProto(o *container.ClusterNodePoolsMaxPodsConstraint) *containerpb.ContainerClusterNodePoolsMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// ClusterNodePoolsConditionsToProto converts a ClusterNodePoolsConditions resource to its proto representation.
func ContainerClusterNodePoolsConditionsToProto(o *container.ClusterNodePoolsConditions) *containerpb.ContainerClusterNodePoolsConditions {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsConditions{
		Code:          ContainerClusterNodePoolsConditionsCodeEnumToProto(o.Code),
		Message:       dcl.ValueOrEmptyString(o.Message),
		CanonicalCode: ContainerClusterNodePoolsConditionsCanonicalCodeEnumToProto(o.CanonicalCode),
	}
	return p
}

// ClusterNodePoolsUpgradeSettingsToProto converts a ClusterNodePoolsUpgradeSettings resource to its proto representation.
func ContainerClusterNodePoolsUpgradeSettingsToProto(o *container.ClusterNodePoolsUpgradeSettings) *containerpb.ContainerClusterNodePoolsUpgradeSettings {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodePoolsUpgradeSettings{
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
	}
	return p
}

// ClusterLegacyAbacToProto converts a ClusterLegacyAbac resource to its proto representation.
func ContainerClusterLegacyAbacToProto(o *container.ClusterLegacyAbac) *containerpb.ContainerClusterLegacyAbac {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterLegacyAbac{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterNetworkPolicyToProto converts a ClusterNetworkPolicy resource to its proto representation.
func ContainerClusterNetworkPolicyToProto(o *container.ClusterNetworkPolicy) *containerpb.ContainerClusterNetworkPolicy {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNetworkPolicy{
		Provider: ContainerClusterNetworkPolicyProviderEnumToProto(o.Provider),
		Enabled:  dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterIPAllocationPolicyToProto converts a ClusterIPAllocationPolicy resource to its proto representation.
func ContainerClusterIPAllocationPolicyToProto(o *container.ClusterIPAllocationPolicy) *containerpb.ContainerClusterIPAllocationPolicy {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterIPAllocationPolicy{
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
	}
	return p
}

// ClusterMasterAuthorizedNetworksConfigToProto converts a ClusterMasterAuthorizedNetworksConfig resource to its proto representation.
func ContainerClusterMasterAuthorizedNetworksConfigToProto(o *container.ClusterMasterAuthorizedNetworksConfig) *containerpb.ContainerClusterMasterAuthorizedNetworksConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMasterAuthorizedNetworksConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	for _, r := range o.CidrBlocks {
		p.CidrBlocks = append(p.CidrBlocks, ContainerClusterMasterAuthorizedNetworksConfigCidrBlocksToProto(&r))
	}
	return p
}

// ClusterMasterAuthorizedNetworksConfigCidrBlocksToProto converts a ClusterMasterAuthorizedNetworksConfigCidrBlocks resource to its proto representation.
func ContainerClusterMasterAuthorizedNetworksConfigCidrBlocksToProto(o *container.ClusterMasterAuthorizedNetworksConfigCidrBlocks) *containerpb.ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks{
		DisplayName: dcl.ValueOrEmptyString(o.DisplayName),
		CidrBlock:   dcl.ValueOrEmptyString(o.CidrBlock),
	}
	return p
}

// ClusterBinaryAuthorizationToProto converts a ClusterBinaryAuthorization resource to its proto representation.
func ContainerClusterBinaryAuthorizationToProto(o *container.ClusterBinaryAuthorization) *containerpb.ContainerClusterBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterBinaryAuthorization{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAutoscalingToProto converts a ClusterAutoscaling resource to its proto representation.
func ContainerClusterAutoscalingToProto(o *container.ClusterAutoscaling) *containerpb.ContainerClusterAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscaling{
		EnableNodeAutoprovisioning:       dcl.ValueOrEmptyBool(o.EnableNodeAutoprovisioning),
		AutoprovisioningNodePoolDefaults: ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto(o.AutoprovisioningNodePoolDefaults),
	}
	for _, r := range o.ResourceLimits {
		p.ResourceLimits = append(p.ResourceLimits, ContainerClusterAutoscalingResourceLimitsToProto(&r))
	}
	for _, r := range o.AutoprovisioningLocations {
		p.AutoprovisioningLocations = append(p.AutoprovisioningLocations, r)
	}
	return p
}

// ClusterAutoscalingResourceLimitsToProto converts a ClusterAutoscalingResourceLimits resource to its proto representation.
func ContainerClusterAutoscalingResourceLimitsToProto(o *container.ClusterAutoscalingResourceLimits) *containerpb.ContainerClusterAutoscalingResourceLimits {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingResourceLimits{
		ResourceType: dcl.ValueOrEmptyString(o.ResourceType),
		Minimum:      dcl.ValueOrEmptyInt64(o.Minimum),
		Maximum:      dcl.ValueOrEmptyInt64(o.Maximum),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaults resource to its proto representation.
func ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsToProto(o *container.ClusterAutoscalingAutoprovisioningNodePoolDefaults) *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaults {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaults{
		ServiceAccount:         dcl.ValueOrEmptyString(o.ServiceAccount),
		UpgradeSettings:        ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto(o.UpgradeSettings),
		Management:             ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto(o.Management),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		DiskSizeGb:             dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		DiskType:               dcl.ValueOrEmptyString(o.DiskType),
		ShieldedInstanceConfig: ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
		BootDiskKmsKey:         dcl.ValueOrEmptyString(o.BootDiskKmsKey),
	}
	for _, r := range o.OAuthScopes {
		p.OauthScopes = append(p.OauthScopes, r)
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings resource to its proto representation.
func ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsToProto(o *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings) *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings{
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement resource to its proto representation.
func ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementToProto(o *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement) *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement{
		AutoUpgrade:    dcl.ValueOrEmptyBool(o.AutoUpgrade),
		AutoRepair:     dcl.ValueOrEmptyBool(o.AutoRepair),
		UpgradeOptions: ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsToProto(o.UpgradeOptions),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions resource to its proto representation.
func ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptionsToProto(o *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions) *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.ValueOrEmptyString(o.AutoUpgradeStartTime),
		Description:          dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigToProto converts a ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig resource to its proto representation.
func ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfigToProto(o *container.ClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig) *containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// ClusterNetworkConfigToProto converts a ClusterNetworkConfig resource to its proto representation.
func ContainerClusterNetworkConfigToProto(o *container.ClusterNetworkConfig) *containerpb.ContainerClusterNetworkConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNetworkConfig{
		Network:                   dcl.ValueOrEmptyString(o.Network),
		Subnetwork:                dcl.ValueOrEmptyString(o.Subnetwork),
		EnableIntraNodeVisibility: dcl.ValueOrEmptyBool(o.EnableIntraNodeVisibility),
		DefaultSnatStatus:         ContainerClusterNetworkConfigDefaultSnatStatusToProto(o.DefaultSnatStatus),
		PrivateIpv6GoogleAccess:   ContainerClusterNetworkConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess),
	}
	return p
}

// ClusterNetworkConfigDefaultSnatStatusToProto converts a ClusterNetworkConfigDefaultSnatStatus resource to its proto representation.
func ContainerClusterNetworkConfigDefaultSnatStatusToProto(o *container.ClusterNetworkConfigDefaultSnatStatus) *containerpb.ContainerClusterNetworkConfigDefaultSnatStatus {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNetworkConfigDefaultSnatStatus{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	return p
}

// ClusterMaintenancePolicyToProto converts a ClusterMaintenancePolicy resource to its proto representation.
func ContainerClusterMaintenancePolicyToProto(o *container.ClusterMaintenancePolicy) *containerpb.ContainerClusterMaintenancePolicy {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicy{
		Window:          ContainerClusterMaintenancePolicyWindowToProto(o.Window),
		ResourceVersion: dcl.ValueOrEmptyString(o.ResourceVersion),
	}
	return p
}

// ClusterMaintenancePolicyWindowToProto converts a ClusterMaintenancePolicyWindow resource to its proto representation.
func ContainerClusterMaintenancePolicyWindowToProto(o *container.ClusterMaintenancePolicyWindow) *containerpb.ContainerClusterMaintenancePolicyWindow {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicyWindow{
		DailyMaintenanceWindow: ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto(o.DailyMaintenanceWindow),
		RecurringWindow:        ContainerClusterMaintenancePolicyWindowRecurringWindowToProto(o.RecurringWindow),
	}
	p.MaintenanceExclusions = make(map[string]string)
	for k, r := range o.MaintenanceExclusions {
		p.MaintenanceExclusions[k] = r
	}
	return p
}

// ClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto converts a ClusterMaintenancePolicyWindowDailyMaintenanceWindow resource to its proto representation.
func ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindowToProto(o *container.ClusterMaintenancePolicyWindowDailyMaintenanceWindow) *containerpb.ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		Duration:  dcl.ValueOrEmptyString(o.Duration),
	}
	return p
}

// ClusterMaintenancePolicyWindowRecurringWindowToProto converts a ClusterMaintenancePolicyWindowRecurringWindow resource to its proto representation.
func ContainerClusterMaintenancePolicyWindowRecurringWindowToProto(o *container.ClusterMaintenancePolicyWindowRecurringWindow) *containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindow {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindow{
		Window:     ContainerClusterMaintenancePolicyWindowRecurringWindowWindowToProto(o.Window),
		Recurrence: dcl.ValueOrEmptyString(o.Recurrence),
	}
	return p
}

// ClusterMaintenancePolicyWindowRecurringWindowWindowToProto converts a ClusterMaintenancePolicyWindowRecurringWindowWindow resource to its proto representation.
func ContainerClusterMaintenancePolicyWindowRecurringWindowWindowToProto(o *container.ClusterMaintenancePolicyWindowRecurringWindowWindow) *containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindowWindow {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterMaintenancePolicyWindowRecurringWindowWindow{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		EndTime:   dcl.ValueOrEmptyString(o.EndTime),
	}
	return p
}

// ClusterDefaultMaxPodsConstraintToProto converts a ClusterDefaultMaxPodsConstraint resource to its proto representation.
func ContainerClusterDefaultMaxPodsConstraintToProto(o *container.ClusterDefaultMaxPodsConstraint) *containerpb.ContainerClusterDefaultMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterDefaultMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyString(o.MaxPodsPerNode),
	}
	return p
}

// ClusterResourceUsageExportConfigToProto converts a ClusterResourceUsageExportConfig resource to its proto representation.
func ContainerClusterResourceUsageExportConfigToProto(o *container.ClusterResourceUsageExportConfig) *containerpb.ContainerClusterResourceUsageExportConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterResourceUsageExportConfig{
		BigqueryDestination:           ContainerClusterResourceUsageExportConfigBigqueryDestinationToProto(o.BigqueryDestination),
		EnableNetworkEgressMonitoring: dcl.ValueOrEmptyBool(o.EnableNetworkEgressMonitoring),
		ConsumptionMeteringConfig:     ContainerClusterResourceUsageExportConfigConsumptionMeteringConfigToProto(o.ConsumptionMeteringConfig),
		EnableNetworkEgressMetering:   dcl.ValueOrEmptyBool(o.EnableNetworkEgressMetering),
	}
	return p
}

// ClusterResourceUsageExportConfigBigqueryDestinationToProto converts a ClusterResourceUsageExportConfigBigqueryDestination resource to its proto representation.
func ContainerClusterResourceUsageExportConfigBigqueryDestinationToProto(o *container.ClusterResourceUsageExportConfigBigqueryDestination) *containerpb.ContainerClusterResourceUsageExportConfigBigqueryDestination {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterResourceUsageExportConfigBigqueryDestination{
		DatasetId: dcl.ValueOrEmptyString(o.DatasetId),
	}
	return p
}

// ClusterResourceUsageExportConfigConsumptionMeteringConfigToProto converts a ClusterResourceUsageExportConfigConsumptionMeteringConfig resource to its proto representation.
func ContainerClusterResourceUsageExportConfigConsumptionMeteringConfigToProto(o *container.ClusterResourceUsageExportConfigConsumptionMeteringConfig) *containerpb.ContainerClusterResourceUsageExportConfigConsumptionMeteringConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterResourceUsageExportConfigConsumptionMeteringConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterAuthenticatorGroupsConfigToProto converts a ClusterAuthenticatorGroupsConfig resource to its proto representation.
func ContainerClusterAuthenticatorGroupsConfigToProto(o *container.ClusterAuthenticatorGroupsConfig) *containerpb.ContainerClusterAuthenticatorGroupsConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAuthenticatorGroupsConfig{
		Enabled:       dcl.ValueOrEmptyBool(o.Enabled),
		SecurityGroup: dcl.ValueOrEmptyString(o.SecurityGroup),
	}
	return p
}

// ClusterPrivateClusterConfigToProto converts a ClusterPrivateClusterConfig resource to its proto representation.
func ContainerClusterPrivateClusterConfigToProto(o *container.ClusterPrivateClusterConfig) *containerpb.ContainerClusterPrivateClusterConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterPrivateClusterConfig{
		EnablePrivateNodes:       dcl.ValueOrEmptyBool(o.EnablePrivateNodes),
		EnablePrivateEndpoint:    dcl.ValueOrEmptyBool(o.EnablePrivateEndpoint),
		MasterIpv4CidrBlock:      dcl.ValueOrEmptyString(o.MasterIPv4CidrBlock),
		PrivateEndpoint:          dcl.ValueOrEmptyString(o.PrivateEndpoint),
		PublicEndpoint:           dcl.ValueOrEmptyString(o.PublicEndpoint),
		PeeringName:              dcl.ValueOrEmptyString(o.PeeringName),
		MasterGlobalAccessConfig: ContainerClusterPrivateClusterConfigMasterGlobalAccessConfigToProto(o.MasterGlobalAccessConfig),
	}
	return p
}

// ClusterPrivateClusterConfigMasterGlobalAccessConfigToProto converts a ClusterPrivateClusterConfigMasterGlobalAccessConfig resource to its proto representation.
func ContainerClusterPrivateClusterConfigMasterGlobalAccessConfigToProto(o *container.ClusterPrivateClusterConfigMasterGlobalAccessConfig) *containerpb.ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterPrivateClusterConfigMasterGlobalAccessConfig{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterDatabaseEncryptionToProto converts a ClusterDatabaseEncryption resource to its proto representation.
func ContainerClusterDatabaseEncryptionToProto(o *container.ClusterDatabaseEncryption) *containerpb.ContainerClusterDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterDatabaseEncryption{
		State:   ContainerClusterDatabaseEncryptionStateEnumToProto(o.State),
		KeyName: dcl.ValueOrEmptyString(o.KeyName),
	}
	return p
}

// ClusterVerticalPodAutoscalingToProto converts a ClusterVerticalPodAutoscaling resource to its proto representation.
func ContainerClusterVerticalPodAutoscalingToProto(o *container.ClusterVerticalPodAutoscaling) *containerpb.ContainerClusterVerticalPodAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterVerticalPodAutoscaling{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterShieldedNodesToProto converts a ClusterShieldedNodes resource to its proto representation.
func ContainerClusterShieldedNodesToProto(o *container.ClusterShieldedNodes) *containerpb.ContainerClusterShieldedNodes {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterShieldedNodes{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterConditionsToProto converts a ClusterConditions resource to its proto representation.
func ContainerClusterConditionsToProto(o *container.ClusterConditions) *containerpb.ContainerClusterConditions {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterConditions{
		Code:          dcl.ValueOrEmptyString(o.Code),
		Message:       dcl.ValueOrEmptyString(o.Message),
		CanonicalCode: ContainerClusterConditionsCanonicalCodeEnumToProto(o.CanonicalCode),
	}
	return p
}

// ClusterAutopilotToProto converts a ClusterAutopilot resource to its proto representation.
func ContainerClusterAutopilotToProto(o *container.ClusterAutopilot) *containerpb.ContainerClusterAutopilot {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterAutopilot{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterNodeConfigToProto converts a ClusterNodeConfig resource to its proto representation.
func ContainerClusterNodeConfigToProto(o *container.ClusterNodeConfig) *containerpb.ContainerClusterNodeConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfig{
		MachineType:            dcl.ValueOrEmptyString(o.MachineType),
		DiskSizeGb:             dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		ServiceAccount:         dcl.ValueOrEmptyString(o.ServiceAccount),
		ImageType:              dcl.ValueOrEmptyString(o.ImageType),
		LocalSsdCount:          dcl.ValueOrEmptyInt64(o.LocalSsdCount),
		Preemptible:            dcl.ValueOrEmptyBool(o.Preemptible),
		DiskType:               dcl.ValueOrEmptyString(o.DiskType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		WorkloadMetadataConfig: ContainerClusterNodeConfigWorkloadMetadataConfigToProto(o.WorkloadMetadataConfig),
		SandboxConfig:          ContainerClusterNodeConfigSandboxConfigToProto(o.SandboxConfig),
		NodeGroup:              dcl.ValueOrEmptyString(o.NodeGroup),
		ReservationAffinity:    ContainerClusterNodeConfigReservationAffinityToProto(o.ReservationAffinity),
		ShieldedInstanceConfig: ContainerClusterNodeConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
		LinuxNodeConfig:        ContainerClusterNodeConfigLinuxNodeConfigToProto(o.LinuxNodeConfig),
		KubeletConfig:          ContainerClusterNodeConfigKubeletConfigToProto(o.KubeletConfig),
		BootDiskKmsKey:         dcl.ValueOrEmptyString(o.BootDiskKmsKey),
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
		p.Accelerators = append(p.Accelerators, ContainerClusterNodeConfigAcceleratorsToProto(&r))
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, ContainerClusterNodeConfigTaintsToProto(&r))
	}
	return p
}

// ClusterNodeConfigAcceleratorsToProto converts a ClusterNodeConfigAccelerators resource to its proto representation.
func ContainerClusterNodeConfigAcceleratorsToProto(o *container.ClusterNodeConfigAccelerators) *containerpb.ContainerClusterNodeConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfigAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// ClusterNodeConfigWorkloadMetadataConfigToProto converts a ClusterNodeConfigWorkloadMetadataConfig resource to its proto representation.
func ContainerClusterNodeConfigWorkloadMetadataConfigToProto(o *container.ClusterNodeConfigWorkloadMetadataConfig) *containerpb.ContainerClusterNodeConfigWorkloadMetadataConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfigWorkloadMetadataConfig{
		Mode: ContainerClusterNodeConfigWorkloadMetadataConfigModeEnumToProto(o.Mode),
	}
	return p
}

// ClusterNodeConfigTaintsToProto converts a ClusterNodeConfigTaints resource to its proto representation.
func ContainerClusterNodeConfigTaintsToProto(o *container.ClusterNodeConfigTaints) *containerpb.ContainerClusterNodeConfigTaints {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: ContainerClusterNodeConfigTaintsEffectEnumToProto(o.Effect),
	}
	return p
}

// ClusterNodeConfigSandboxConfigToProto converts a ClusterNodeConfigSandboxConfig resource to its proto representation.
func ContainerClusterNodeConfigSandboxConfigToProto(o *container.ClusterNodeConfigSandboxConfig) *containerpb.ContainerClusterNodeConfigSandboxConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfigSandboxConfig{
		Type: ContainerClusterNodeConfigSandboxConfigTypeEnumToProto(o.Type),
	}
	return p
}

// ClusterNodeConfigReservationAffinityToProto converts a ClusterNodeConfigReservationAffinity resource to its proto representation.
func ContainerClusterNodeConfigReservationAffinityToProto(o *container.ClusterNodeConfigReservationAffinity) *containerpb.ContainerClusterNodeConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfigReservationAffinity{
		ConsumeReservationType: ContainerClusterNodeConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// ClusterNodeConfigShieldedInstanceConfigToProto converts a ClusterNodeConfigShieldedInstanceConfig resource to its proto representation.
func ContainerClusterNodeConfigShieldedInstanceConfigToProto(o *container.ClusterNodeConfigShieldedInstanceConfig) *containerpb.ContainerClusterNodeConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// ClusterNodeConfigLinuxNodeConfigToProto converts a ClusterNodeConfigLinuxNodeConfig resource to its proto representation.
func ContainerClusterNodeConfigLinuxNodeConfigToProto(o *container.ClusterNodeConfigLinuxNodeConfig) *containerpb.ContainerClusterNodeConfigLinuxNodeConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfigLinuxNodeConfig{}
	p.Sysctls = make(map[string]string)
	for k, r := range o.Sysctls {
		p.Sysctls[k] = r
	}
	return p
}

// ClusterNodeConfigKubeletConfigToProto converts a ClusterNodeConfigKubeletConfig resource to its proto representation.
func ContainerClusterNodeConfigKubeletConfigToProto(o *container.ClusterNodeConfigKubeletConfig) *containerpb.ContainerClusterNodeConfigKubeletConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNodeConfigKubeletConfig{
		CpuManagerPolicy:  dcl.ValueOrEmptyString(o.CpuManagerPolicy),
		CpuCfsQuota:       dcl.ValueOrEmptyBool(o.CpuCfsQuota),
		CpuCfsQuotaPeriod: dcl.ValueOrEmptyString(o.CpuCfsQuotaPeriod),
	}
	return p
}

// ClusterReleaseChannelToProto converts a ClusterReleaseChannel resource to its proto representation.
func ContainerClusterReleaseChannelToProto(o *container.ClusterReleaseChannel) *containerpb.ContainerClusterReleaseChannel {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterReleaseChannel{
		Channel: ContainerClusterReleaseChannelChannelEnumToProto(o.Channel),
	}
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig resource to its proto representation.
func ContainerClusterWorkloadIdentityConfigToProto(o *container.ClusterWorkloadIdentityConfig) *containerpb.ContainerClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterWorkloadIdentityConfig{
		WorkloadPool: dcl.ValueOrEmptyString(o.WorkloadPool),
	}
	return p
}

// ClusterNotificationConfigToProto converts a ClusterNotificationConfig resource to its proto representation.
func ContainerClusterNotificationConfigToProto(o *container.ClusterNotificationConfig) *containerpb.ContainerClusterNotificationConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNotificationConfig{
		Pubsub: ContainerClusterNotificationConfigPubsubToProto(o.Pubsub),
	}
	return p
}

// ClusterNotificationConfigPubsubToProto converts a ClusterNotificationConfigPubsub resource to its proto representation.
func ContainerClusterNotificationConfigPubsubToProto(o *container.ClusterNotificationConfigPubsub) *containerpb.ContainerClusterNotificationConfigPubsub {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterNotificationConfigPubsub{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
		Topic:   dcl.ValueOrEmptyString(o.Topic),
	}
	return p
}

// ClusterConfidentialNodesToProto converts a ClusterConfidentialNodes resource to its proto representation.
func ContainerClusterConfidentialNodesToProto(o *container.ClusterConfidentialNodes) *containerpb.ContainerClusterConfidentialNodes {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerClusterConfidentialNodes{
		Enabled: dcl.ValueOrEmptyBool(o.Enabled),
	}
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *container.Cluster) *containerpb.ContainerCluster {
	p := &containerpb.ContainerCluster{
		Name:                           dcl.ValueOrEmptyString(resource.Name),
		Description:                    dcl.ValueOrEmptyString(resource.Description),
		InitialNodeCount:               dcl.ValueOrEmptyInt64(resource.InitialNodeCount),
		MasterAuth:                     ContainerClusterMasterAuthToProto(resource.MasterAuth),
		LoggingService:                 dcl.ValueOrEmptyString(resource.LoggingService),
		MonitoringService:              dcl.ValueOrEmptyString(resource.MonitoringService),
		Network:                        dcl.ValueOrEmptyString(resource.Network),
		ClusterIpv4Cidr:                dcl.ValueOrEmptyString(resource.ClusterIPv4Cidr),
		AddonsConfig:                   ContainerClusterAddonsConfigToProto(resource.AddonsConfig),
		Subnetwork:                     dcl.ValueOrEmptyString(resource.Subnetwork),
		EnableKubernetesAlpha:          dcl.ValueOrEmptyBool(resource.EnableKubernetesAlpha),
		LabelFingerprint:               dcl.ValueOrEmptyString(resource.LabelFingerprint),
		LegacyAbac:                     ContainerClusterLegacyAbacToProto(resource.LegacyAbac),
		NetworkPolicy:                  ContainerClusterNetworkPolicyToProto(resource.NetworkPolicy),
		IpAllocationPolicy:             ContainerClusterIPAllocationPolicyToProto(resource.IPAllocationPolicy),
		MasterAuthorizedNetworksConfig: ContainerClusterMasterAuthorizedNetworksConfigToProto(resource.MasterAuthorizedNetworksConfig),
		BinaryAuthorization:            ContainerClusterBinaryAuthorizationToProto(resource.BinaryAuthorization),
		Autoscaling:                    ContainerClusterAutoscalingToProto(resource.Autoscaling),
		NetworkConfig:                  ContainerClusterNetworkConfigToProto(resource.NetworkConfig),
		MaintenancePolicy:              ContainerClusterMaintenancePolicyToProto(resource.MaintenancePolicy),
		DefaultMaxPodsConstraint:       ContainerClusterDefaultMaxPodsConstraintToProto(resource.DefaultMaxPodsConstraint),
		ResourceUsageExportConfig:      ContainerClusterResourceUsageExportConfigToProto(resource.ResourceUsageExportConfig),
		AuthenticatorGroupsConfig:      ContainerClusterAuthenticatorGroupsConfigToProto(resource.AuthenticatorGroupsConfig),
		PrivateClusterConfig:           ContainerClusterPrivateClusterConfigToProto(resource.PrivateClusterConfig),
		DatabaseEncryption:             ContainerClusterDatabaseEncryptionToProto(resource.DatabaseEncryption),
		VerticalPodAutoscaling:         ContainerClusterVerticalPodAutoscalingToProto(resource.VerticalPodAutoscaling),
		ShieldedNodes:                  ContainerClusterShieldedNodesToProto(resource.ShieldedNodes),
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
		Autopilot:                      ContainerClusterAutopilotToProto(resource.Autopilot),
		Project:                        dcl.ValueOrEmptyString(resource.Project),
		NodeConfig:                     ContainerClusterNodeConfigToProto(resource.NodeConfig),
		ReleaseChannel:                 ContainerClusterReleaseChannelToProto(resource.ReleaseChannel),
		WorkloadIdentityConfig:         ContainerClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		NotificationConfig:             ContainerClusterNotificationConfigToProto(resource.NotificationConfig),
		ConfidentialNodes:              ContainerClusterConfidentialNodesToProto(resource.ConfidentialNodes),
		SelfLink:                       dcl.ValueOrEmptyString(resource.SelfLink),
		Zone:                           dcl.ValueOrEmptyString(resource.Zone),
		InitialClusterVersion:          dcl.ValueOrEmptyString(resource.InitialClusterVersion),
		CurrentMasterVersion:           dcl.ValueOrEmptyString(resource.CurrentMasterVersion),
		CurrentNodeVersion:             dcl.ValueOrEmptyString(resource.CurrentNodeVersion),
		CurrentNodeCount:               dcl.ValueOrEmptyInt64(resource.CurrentNodeCount),
		Id:                             dcl.ValueOrEmptyString(resource.Id),
	}
	for _, r := range resource.NodePools {
		p.NodePools = append(p.NodePools, ContainerClusterNodePoolsToProto(&r))
	}
	for _, r := range resource.Locations {
		p.Locations = append(p.Locations, r)
	}
	for _, r := range resource.Conditions {
		p.Conditions = append(p.Conditions, ContainerClusterConditionsToProto(&r))
	}
	for _, r := range resource.InstanceGroupUrls {
		p.InstanceGroupUrls = append(p.InstanceGroupUrls, r)
	}

	return p
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *container.Client, request *containerpb.ApplyContainerClusterRequest) (*containerpb.ContainerCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// ApplyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerCluster(ctx context.Context, request *containerpb.ApplyContainerClusterRequest) (*containerpb.ContainerCluster, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerCluster(ctx context.Context, request *containerpb.DeleteContainerClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerCluster(ctx context.Context, request *containerpb.ListContainerClusterRequest) (*containerpb.ListContainerClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*containerpb.ContainerCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	return &containerpb.ListContainerClusterResponse{Items: protos}, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*container.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return container.NewClient(conf), nil
}
