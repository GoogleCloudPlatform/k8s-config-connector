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

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterControlPlaneRootVolumeVolumeTypeEnum converts a ClusterControlPlaneRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum(e betapb.ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum) *beta.ClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterControlPlaneRootVolumeVolumeTypeEnum(n[len("ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterControlPlaneMainVolumeVolumeTypeEnum converts a ClusterControlPlaneMainVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum(e betapb.ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum) *beta.ClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := beta.ClusterControlPlaneMainVolumeVolumeTypeEnum(n[len("ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterControlPlaneInstancePlacementTenancyEnum converts a ClusterControlPlaneInstancePlacementTenancyEnum enum from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum(e betapb.ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum) *beta.ClusterControlPlaneInstancePlacementTenancyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum_name[int32(e)]; ok {
		e := beta.ClusterControlPlaneInstancePlacementTenancyEnum(n[len("ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToContainerawsBetaClusterStateEnum(e betapb.ContainerawsBetaClusterStateEnum) *beta.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaClusterStateEnum_name[int32(e)]; ok {
		e := beta.ClusterStateEnum(n[len("ContainerawsBetaClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterLoggingConfigComponentConfigEnableComponentsEnum converts a ClusterLoggingConfigComponentConfigEnableComponentsEnum enum from its proto representation.
func ProtoToContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(e betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum) *beta.ClusterLoggingConfigComponentConfigEnableComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum_name[int32(e)]; ok {
		e := beta.ClusterLoggingConfigComponentConfigEnableComponentsEnum(n[len("ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterBinaryAuthorizationEvaluationModeEnum converts a ClusterBinaryAuthorizationEvaluationModeEnum enum from its proto representation.
func ProtoToContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum(e betapb.ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum) *beta.ClusterBinaryAuthorizationEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum_name[int32(e)]; ok {
		e := beta.ClusterBinaryAuthorizationEvaluationModeEnum(n[len("ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworking converts a ClusterNetworking object from its proto representation.
func ProtoToContainerawsBetaClusterNetworking(p *betapb.ContainerawsBetaClusterNetworking) *beta.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNetworking{
		VPCId:                      dcl.StringOrNil(p.GetVpcId()),
		PerNodePoolSgRulesDisabled: dcl.Bool(p.GetPerNodePoolSgRulesDisabled()),
	}
	for _, r := range p.GetPodAddressCidrBlocks() {
		obj.PodAddressCidrBlocks = append(obj.PodAddressCidrBlocks, r)
	}
	for _, r := range p.GetServiceAddressCidrBlocks() {
		obj.ServiceAddressCidrBlocks = append(obj.ServiceAddressCidrBlocks, r)
	}
	return obj
}

// ProtoToClusterControlPlane converts a ClusterControlPlane object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlane(p *betapb.ContainerawsBetaClusterControlPlane) *beta.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlane{
		Version:                   dcl.StringOrNil(p.GetVersion()),
		InstanceType:              dcl.StringOrNil(p.GetInstanceType()),
		SshConfig:                 ProtoToContainerawsBetaClusterControlPlaneSshConfig(p.GetSshConfig()),
		ConfigEncryption:          ProtoToContainerawsBetaClusterControlPlaneConfigEncryption(p.GetConfigEncryption()),
		IamInstanceProfile:        dcl.StringOrNil(p.GetIamInstanceProfile()),
		RootVolume:                ProtoToContainerawsBetaClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToContainerawsBetaClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToContainerawsBetaClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToContainerawsBetaClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
		ProxyConfig:               ProtoToContainerawsBetaClusterControlPlaneProxyConfig(p.GetProxyConfig()),
		InstancePlacement:         ProtoToContainerawsBetaClusterControlPlaneInstancePlacement(p.GetInstancePlacement()),
	}
	for _, r := range p.GetSubnetIds() {
		obj.SubnetIds = append(obj.SubnetIds, r)
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToClusterControlPlaneSshConfig converts a ClusterControlPlaneSshConfig object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneSshConfig(p *betapb.ContainerawsBetaClusterControlPlaneSshConfig) *beta.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.GetEc2KeyPair()),
	}
	return obj
}

// ProtoToClusterControlPlaneConfigEncryption converts a ClusterControlPlaneConfigEncryption object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneConfigEncryption(p *betapb.ContainerawsBetaClusterControlPlaneConfigEncryption) *beta.ClusterControlPlaneConfigEncryption {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneConfigEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneRootVolume(p *betapb.ContainerawsBetaClusterControlPlaneRootVolume) *beta.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneMainVolume(p *betapb.ContainerawsBetaClusterControlPlaneMainVolume) *beta.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneDatabaseEncryption(p *betapb.ContainerawsBetaClusterControlPlaneDatabaseEncryption) *beta.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneAwsServicesAuthentication converts a ClusterControlPlaneAwsServicesAuthentication object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneAwsServicesAuthentication(p *betapb.ContainerawsBetaClusterControlPlaneAwsServicesAuthentication) *beta.ClusterControlPlaneAwsServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.StringOrNil(p.GetRoleArn()),
		RoleSessionName: dcl.StringOrNil(p.GetRoleSessionName()),
	}
	return obj
}

// ProtoToClusterControlPlaneProxyConfig converts a ClusterControlPlaneProxyConfig object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneProxyConfig(p *betapb.ContainerawsBetaClusterControlPlaneProxyConfig) *beta.ClusterControlPlaneProxyConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneProxyConfig{
		SecretArn:     dcl.StringOrNil(p.GetSecretArn()),
		SecretVersion: dcl.StringOrNil(p.GetSecretVersion()),
	}
	return obj
}

// ProtoToClusterControlPlaneInstancePlacement converts a ClusterControlPlaneInstancePlacement object from its proto representation.
func ProtoToContainerawsBetaClusterControlPlaneInstancePlacement(p *betapb.ContainerawsBetaClusterControlPlaneInstancePlacement) *beta.ClusterControlPlaneInstancePlacement {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneInstancePlacement{
		Tenancy: ProtoToContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum(p.GetTenancy()),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization object from its proto representation.
func ProtoToContainerawsBetaClusterAuthorization(p *betapb.ContainerawsBetaClusterAuthorization) *beta.ClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerawsBetaClusterAuthorizationAdminUsers(r))
	}
	for _, r := range p.GetAdminGroups() {
		obj.AdminGroups = append(obj.AdminGroups, *ProtoToContainerawsBetaClusterAuthorizationAdminGroups(r))
	}
	return obj
}

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers object from its proto representation.
func ProtoToContainerawsBetaClusterAuthorizationAdminUsers(p *betapb.ContainerawsBetaClusterAuthorizationAdminUsers) *beta.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToClusterAuthorizationAdminGroups converts a ClusterAuthorizationAdminGroups object from its proto representation.
func ProtoToContainerawsBetaClusterAuthorizationAdminGroups(p *betapb.ContainerawsBetaClusterAuthorizationAdminGroups) *beta.ClusterAuthorizationAdminGroups {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAuthorizationAdminGroups{
		Group: dcl.StringOrNil(p.GetGroup()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig object from its proto representation.
func ProtoToContainerawsBetaClusterWorkloadIdentityConfig(p *betapb.ContainerawsBetaClusterWorkloadIdentityConfig) *beta.ClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.GetIssuerUri()),
		WorkloadPool:     dcl.StringOrNil(p.GetWorkloadPool()),
		IdentityProvider: dcl.StringOrNil(p.GetIdentityProvider()),
	}
	return obj
}

// ProtoToClusterFleet converts a ClusterFleet object from its proto representation.
func ProtoToContainerawsBetaClusterFleet(p *betapb.ContainerawsBetaClusterFleet) *beta.ClusterFleet {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterFleet{
		Project:    dcl.StringOrNil(p.GetProject()),
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToClusterLoggingConfig converts a ClusterLoggingConfig object from its proto representation.
func ProtoToContainerawsBetaClusterLoggingConfig(p *betapb.ContainerawsBetaClusterLoggingConfig) *beta.ClusterLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterLoggingConfig{
		ComponentConfig: ProtoToContainerawsBetaClusterLoggingConfigComponentConfig(p.GetComponentConfig()),
	}
	return obj
}

// ProtoToClusterLoggingConfigComponentConfig converts a ClusterLoggingConfigComponentConfig object from its proto representation.
func ProtoToContainerawsBetaClusterLoggingConfigComponentConfig(p *betapb.ContainerawsBetaClusterLoggingConfigComponentConfig) *beta.ClusterLoggingConfigComponentConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterLoggingConfigComponentConfig{}
	for _, r := range p.GetEnableComponents() {
		obj.EnableComponents = append(obj.EnableComponents, *ProtoToContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterMonitoringConfig converts a ClusterMonitoringConfig object from its proto representation.
func ProtoToContainerawsBetaClusterMonitoringConfig(p *betapb.ContainerawsBetaClusterMonitoringConfig) *beta.ClusterMonitoringConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMonitoringConfig{
		ManagedPrometheusConfig: ProtoToContainerawsBetaClusterMonitoringConfigManagedPrometheusConfig(p.GetManagedPrometheusConfig()),
	}
	return obj
}

// ProtoToClusterMonitoringConfigManagedPrometheusConfig converts a ClusterMonitoringConfigManagedPrometheusConfig object from its proto representation.
func ProtoToContainerawsBetaClusterMonitoringConfigManagedPrometheusConfig(p *betapb.ContainerawsBetaClusterMonitoringConfigManagedPrometheusConfig) *beta.ClusterMonitoringConfigManagedPrometheusConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMonitoringConfigManagedPrometheusConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToClusterBinaryAuthorization converts a ClusterBinaryAuthorization object from its proto representation.
func ProtoToContainerawsBetaClusterBinaryAuthorization(p *betapb.ContainerawsBetaClusterBinaryAuthorization) *beta.ClusterBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterBinaryAuthorization{
		EvaluationMode: ProtoToContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum(p.GetEvaluationMode()),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *betapb.ContainerawsBetaCluster) *beta.Cluster {
	obj := &beta.Cluster{
		Name:                   dcl.StringOrNil(p.GetName()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		Networking:             ProtoToContainerawsBetaClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.GetAwsRegion()),
		ControlPlane:           ProtoToContainerawsBetaClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToContainerawsBetaClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToContainerawsBetaClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.GetEndpoint()),
		Uid:                    dcl.StringOrNil(p.GetUid()),
		Reconciling:            dcl.Bool(p.GetReconciling()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.GetEtag()),
		WorkloadIdentityConfig: ProtoToContainerawsBetaClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
		Fleet:                  ProtoToContainerawsBetaClusterFleet(p.GetFleet()),
		LoggingConfig:          ProtoToContainerawsBetaClusterLoggingConfig(p.GetLoggingConfig()),
		MonitoringConfig:       ProtoToContainerawsBetaClusterMonitoringConfig(p.GetMonitoringConfig()),
		BinaryAuthorization:    ProtoToContainerawsBetaClusterBinaryAuthorization(p.GetBinaryAuthorization()),
	}
	return obj
}

// ClusterControlPlaneRootVolumeVolumeTypeEnumToProto converts a ClusterControlPlaneRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnumToProto(e *beta.ClusterControlPlaneRootVolumeVolumeTypeEnum) betapb.ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == nil {
		return betapb.ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum_value["ClusterControlPlaneRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum(v)
	}
	return betapb.ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnum(0)
}

// ClusterControlPlaneMainVolumeVolumeTypeEnumToProto converts a ClusterControlPlaneMainVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnumToProto(e *beta.ClusterControlPlaneMainVolumeVolumeTypeEnum) betapb.ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == nil {
		return betapb.ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum_value["ClusterControlPlaneMainVolumeVolumeTypeEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum(v)
	}
	return betapb.ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnum(0)
}

// ClusterControlPlaneInstancePlacementTenancyEnumToProto converts a ClusterControlPlaneInstancePlacementTenancyEnum enum to its proto representation.
func ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnumToProto(e *beta.ClusterControlPlaneInstancePlacementTenancyEnum) betapb.ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum {
	if e == nil {
		return betapb.ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum_value["ClusterControlPlaneInstancePlacementTenancyEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum(v)
	}
	return betapb.ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnum(0)
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func ContainerawsBetaClusterStateEnumToProto(e *beta.ClusterStateEnum) betapb.ContainerawsBetaClusterStateEnum {
	if e == nil {
		return betapb.ContainerawsBetaClusterStateEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaClusterStateEnum(v)
	}
	return betapb.ContainerawsBetaClusterStateEnum(0)
}

// ClusterLoggingConfigComponentConfigEnableComponentsEnumToProto converts a ClusterLoggingConfigComponentConfigEnableComponentsEnum enum to its proto representation.
func ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnumToProto(e *beta.ClusterLoggingConfigComponentConfigEnableComponentsEnum) betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum {
	if e == nil {
		return betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum_value["ClusterLoggingConfigComponentConfigEnableComponentsEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(v)
	}
	return betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(0)
}

// ClusterBinaryAuthorizationEvaluationModeEnumToProto converts a ClusterBinaryAuthorizationEvaluationModeEnum enum to its proto representation.
func ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnumToProto(e *beta.ClusterBinaryAuthorizationEvaluationModeEnum) betapb.ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum {
	if e == nil {
		return betapb.ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum(0)
	}
	if v, ok := betapb.ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum_value["ClusterBinaryAuthorizationEvaluationModeEnum"+string(*e)]; ok {
		return betapb.ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum(v)
	}
	return betapb.ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnum(0)
}

// ClusterNetworkingToProto converts a ClusterNetworking object to its proto representation.
func ContainerawsBetaClusterNetworkingToProto(o *beta.ClusterNetworking) *betapb.ContainerawsBetaClusterNetworking {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterNetworking{}
	p.SetVpcId(dcl.ValueOrEmptyString(o.VPCId))
	p.SetPerNodePoolSgRulesDisabled(dcl.ValueOrEmptyBool(o.PerNodePoolSgRulesDisabled))
	sPodAddressCidrBlocks := make([]string, len(o.PodAddressCidrBlocks))
	for i, r := range o.PodAddressCidrBlocks {
		sPodAddressCidrBlocks[i] = r
	}
	p.SetPodAddressCidrBlocks(sPodAddressCidrBlocks)
	sServiceAddressCidrBlocks := make([]string, len(o.ServiceAddressCidrBlocks))
	for i, r := range o.ServiceAddressCidrBlocks {
		sServiceAddressCidrBlocks[i] = r
	}
	p.SetServiceAddressCidrBlocks(sServiceAddressCidrBlocks)
	return p
}

// ClusterControlPlaneToProto converts a ClusterControlPlane object to its proto representation.
func ContainerawsBetaClusterControlPlaneToProto(o *beta.ClusterControlPlane) *betapb.ContainerawsBetaClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlane{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetInstanceType(dcl.ValueOrEmptyString(o.InstanceType))
	p.SetSshConfig(ContainerawsBetaClusterControlPlaneSshConfigToProto(o.SshConfig))
	p.SetConfigEncryption(ContainerawsBetaClusterControlPlaneConfigEncryptionToProto(o.ConfigEncryption))
	p.SetIamInstanceProfile(dcl.ValueOrEmptyString(o.IamInstanceProfile))
	p.SetRootVolume(ContainerawsBetaClusterControlPlaneRootVolumeToProto(o.RootVolume))
	p.SetMainVolume(ContainerawsBetaClusterControlPlaneMainVolumeToProto(o.MainVolume))
	p.SetDatabaseEncryption(ContainerawsBetaClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption))
	p.SetAwsServicesAuthentication(ContainerawsBetaClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication))
	p.SetProxyConfig(ContainerawsBetaClusterControlPlaneProxyConfigToProto(o.ProxyConfig))
	p.SetInstancePlacement(ContainerawsBetaClusterControlPlaneInstancePlacementToProto(o.InstancePlacement))
	sSubnetIds := make([]string, len(o.SubnetIds))
	for i, r := range o.SubnetIds {
		sSubnetIds[i] = r
	}
	p.SetSubnetIds(sSubnetIds)
	sSecurityGroupIds := make([]string, len(o.SecurityGroupIds))
	for i, r := range o.SecurityGroupIds {
		sSecurityGroupIds[i] = r
	}
	p.SetSecurityGroupIds(sSecurityGroupIds)
	mTags := make(map[string]string, len(o.Tags))
	for k, r := range o.Tags {
		mTags[k] = r
	}
	p.SetTags(mTags)
	return p
}

// ClusterControlPlaneSshConfigToProto converts a ClusterControlPlaneSshConfig object to its proto representation.
func ContainerawsBetaClusterControlPlaneSshConfigToProto(o *beta.ClusterControlPlaneSshConfig) *betapb.ContainerawsBetaClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlaneSshConfig{}
	p.SetEc2KeyPair(dcl.ValueOrEmptyString(o.Ec2KeyPair))
	return p
}

// ClusterControlPlaneConfigEncryptionToProto converts a ClusterControlPlaneConfigEncryption object to its proto representation.
func ContainerawsBetaClusterControlPlaneConfigEncryptionToProto(o *beta.ClusterControlPlaneConfigEncryption) *betapb.ContainerawsBetaClusterControlPlaneConfigEncryption {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlaneConfigEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume object to its proto representation.
func ContainerawsBetaClusterControlPlaneRootVolumeToProto(o *beta.ClusterControlPlaneRootVolume) *betapb.ContainerawsBetaClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlaneRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsBetaClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume object to its proto representation.
func ContainerawsBetaClusterControlPlaneMainVolumeToProto(o *beta.ClusterControlPlaneMainVolume) *betapb.ContainerawsBetaClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlaneMainVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsBetaClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption object to its proto representation.
func ContainerawsBetaClusterControlPlaneDatabaseEncryptionToProto(o *beta.ClusterControlPlaneDatabaseEncryption) *betapb.ContainerawsBetaClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlaneDatabaseEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneAwsServicesAuthenticationToProto converts a ClusterControlPlaneAwsServicesAuthentication object to its proto representation.
func ContainerawsBetaClusterControlPlaneAwsServicesAuthenticationToProto(o *beta.ClusterControlPlaneAwsServicesAuthentication) *betapb.ContainerawsBetaClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlaneAwsServicesAuthentication{}
	p.SetRoleArn(dcl.ValueOrEmptyString(o.RoleArn))
	p.SetRoleSessionName(dcl.ValueOrEmptyString(o.RoleSessionName))
	return p
}

// ClusterControlPlaneProxyConfigToProto converts a ClusterControlPlaneProxyConfig object to its proto representation.
func ContainerawsBetaClusterControlPlaneProxyConfigToProto(o *beta.ClusterControlPlaneProxyConfig) *betapb.ContainerawsBetaClusterControlPlaneProxyConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlaneProxyConfig{}
	p.SetSecretArn(dcl.ValueOrEmptyString(o.SecretArn))
	p.SetSecretVersion(dcl.ValueOrEmptyString(o.SecretVersion))
	return p
}

// ClusterControlPlaneInstancePlacementToProto converts a ClusterControlPlaneInstancePlacement object to its proto representation.
func ContainerawsBetaClusterControlPlaneInstancePlacementToProto(o *beta.ClusterControlPlaneInstancePlacement) *betapb.ContainerawsBetaClusterControlPlaneInstancePlacement {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterControlPlaneInstancePlacement{}
	p.SetTenancy(ContainerawsBetaClusterControlPlaneInstancePlacementTenancyEnumToProto(o.Tenancy))
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization object to its proto representation.
func ContainerawsBetaClusterAuthorizationToProto(o *beta.ClusterAuthorization) *betapb.ContainerawsBetaClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterAuthorization{}
	sAdminUsers := make([]*betapb.ContainerawsBetaClusterAuthorizationAdminUsers, len(o.AdminUsers))
	for i, r := range o.AdminUsers {
		sAdminUsers[i] = ContainerawsBetaClusterAuthorizationAdminUsersToProto(&r)
	}
	p.SetAdminUsers(sAdminUsers)
	sAdminGroups := make([]*betapb.ContainerawsBetaClusterAuthorizationAdminGroups, len(o.AdminGroups))
	for i, r := range o.AdminGroups {
		sAdminGroups[i] = ContainerawsBetaClusterAuthorizationAdminGroupsToProto(&r)
	}
	p.SetAdminGroups(sAdminGroups)
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers object to its proto representation.
func ContainerawsBetaClusterAuthorizationAdminUsersToProto(o *beta.ClusterAuthorizationAdminUsers) *betapb.ContainerawsBetaClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterAuthorizationAdminUsers{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ClusterAuthorizationAdminGroupsToProto converts a ClusterAuthorizationAdminGroups object to its proto representation.
func ContainerawsBetaClusterAuthorizationAdminGroupsToProto(o *beta.ClusterAuthorizationAdminGroups) *betapb.ContainerawsBetaClusterAuthorizationAdminGroups {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterAuthorizationAdminGroups{}
	p.SetGroup(dcl.ValueOrEmptyString(o.Group))
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig object to its proto representation.
func ContainerawsBetaClusterWorkloadIdentityConfigToProto(o *beta.ClusterWorkloadIdentityConfig) *betapb.ContainerawsBetaClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterWorkloadIdentityConfig{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetWorkloadPool(dcl.ValueOrEmptyString(o.WorkloadPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// ClusterFleetToProto converts a ClusterFleet object to its proto representation.
func ContainerawsBetaClusterFleetToProto(o *beta.ClusterFleet) *betapb.ContainerawsBetaClusterFleet {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterFleet{}
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// ClusterLoggingConfigToProto converts a ClusterLoggingConfig object to its proto representation.
func ContainerawsBetaClusterLoggingConfigToProto(o *beta.ClusterLoggingConfig) *betapb.ContainerawsBetaClusterLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterLoggingConfig{}
	p.SetComponentConfig(ContainerawsBetaClusterLoggingConfigComponentConfigToProto(o.ComponentConfig))
	return p
}

// ClusterLoggingConfigComponentConfigToProto converts a ClusterLoggingConfigComponentConfig object to its proto representation.
func ContainerawsBetaClusterLoggingConfigComponentConfigToProto(o *beta.ClusterLoggingConfigComponentConfig) *betapb.ContainerawsBetaClusterLoggingConfigComponentConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterLoggingConfigComponentConfig{}
	sEnableComponents := make([]betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum, len(o.EnableComponents))
	for i, r := range o.EnableComponents {
		sEnableComponents[i] = betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(betapb.ContainerawsBetaClusterLoggingConfigComponentConfigEnableComponentsEnum_value[string(r)])
	}
	p.SetEnableComponents(sEnableComponents)
	return p
}

// ClusterMonitoringConfigToProto converts a ClusterMonitoringConfig object to its proto representation.
func ContainerawsBetaClusterMonitoringConfigToProto(o *beta.ClusterMonitoringConfig) *betapb.ContainerawsBetaClusterMonitoringConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterMonitoringConfig{}
	p.SetManagedPrometheusConfig(ContainerawsBetaClusterMonitoringConfigManagedPrometheusConfigToProto(o.ManagedPrometheusConfig))
	return p
}

// ClusterMonitoringConfigManagedPrometheusConfigToProto converts a ClusterMonitoringConfigManagedPrometheusConfig object to its proto representation.
func ContainerawsBetaClusterMonitoringConfigManagedPrometheusConfigToProto(o *beta.ClusterMonitoringConfigManagedPrometheusConfig) *betapb.ContainerawsBetaClusterMonitoringConfigManagedPrometheusConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterMonitoringConfigManagedPrometheusConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ClusterBinaryAuthorizationToProto converts a ClusterBinaryAuthorization object to its proto representation.
func ContainerawsBetaClusterBinaryAuthorizationToProto(o *beta.ClusterBinaryAuthorization) *betapb.ContainerawsBetaClusterBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerawsBetaClusterBinaryAuthorization{}
	p.SetEvaluationMode(ContainerawsBetaClusterBinaryAuthorizationEvaluationModeEnumToProto(o.EvaluationMode))
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *beta.Cluster) *betapb.ContainerawsBetaCluster {
	p := &betapb.ContainerawsBetaCluster{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetNetworking(ContainerawsBetaClusterNetworkingToProto(resource.Networking))
	p.SetAwsRegion(dcl.ValueOrEmptyString(resource.AwsRegion))
	p.SetControlPlane(ContainerawsBetaClusterControlPlaneToProto(resource.ControlPlane))
	p.SetAuthorization(ContainerawsBetaClusterAuthorizationToProto(resource.Authorization))
	p.SetState(ContainerawsBetaClusterStateEnumToProto(resource.State))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkloadIdentityConfig(ContainerawsBetaClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFleet(ContainerawsBetaClusterFleetToProto(resource.Fleet))
	p.SetLoggingConfig(ContainerawsBetaClusterLoggingConfigToProto(resource.LoggingConfig))
	p.SetMonitoringConfig(ContainerawsBetaClusterMonitoringConfigToProto(resource.MonitoringConfig))
	p.SetBinaryAuthorization(ContainerawsBetaClusterBinaryAuthorizationToProto(resource.BinaryAuthorization))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *beta.Client, request *betapb.ApplyContainerawsBetaClusterRequest) (*betapb.ContainerawsBetaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyContainerawsBetaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerawsBetaCluster(ctx context.Context, request *betapb.ApplyContainerawsBetaClusterRequest) (*betapb.ContainerawsBetaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerawsBetaCluster(ctx context.Context, request *betapb.DeleteContainerawsBetaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerawsBetaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerawsBetaCluster(ctx context.Context, request *betapb.ListContainerawsBetaClusterRequest) (*betapb.ListContainerawsBetaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContainerawsBetaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListContainerawsBetaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
