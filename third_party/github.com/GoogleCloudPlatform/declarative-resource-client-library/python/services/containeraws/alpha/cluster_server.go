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

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterControlPlaneRootVolumeVolumeTypeEnum converts a ClusterControlPlaneRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum) *alpha.ClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.ClusterControlPlaneRootVolumeVolumeTypeEnum(n[len("ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterControlPlaneMainVolumeVolumeTypeEnum converts a ClusterControlPlaneMainVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum) *alpha.ClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.ClusterControlPlaneMainVolumeVolumeTypeEnum(n[len("ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterControlPlaneInstancePlacementTenancyEnum converts a ClusterControlPlaneInstancePlacementTenancyEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum(e alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum) *alpha.ClusterControlPlaneInstancePlacementTenancyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum_name[int32(e)]; ok {
		e := alpha.ClusterControlPlaneInstancePlacementTenancyEnum(n[len("ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterStateEnum(e alphapb.ContainerawsAlphaClusterStateEnum) *alpha.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStateEnum(n[len("ContainerawsAlphaClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterLoggingConfigComponentConfigEnableComponentsEnum converts a ClusterLoggingConfigComponentConfigEnableComponentsEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(e alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum) *alpha.ClusterLoggingConfigComponentConfigEnableComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum_name[int32(e)]; ok {
		e := alpha.ClusterLoggingConfigComponentConfigEnableComponentsEnum(n[len("ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterBinaryAuthorizationEvaluationModeEnum converts a ClusterBinaryAuthorizationEvaluationModeEnum enum from its proto representation.
func ProtoToContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum(e alphapb.ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum) *alpha.ClusterBinaryAuthorizationEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum_name[int32(e)]; ok {
		e := alpha.ClusterBinaryAuthorizationEvaluationModeEnum(n[len("ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworking converts a ClusterNetworking object from its proto representation.
func ProtoToContainerawsAlphaClusterNetworking(p *alphapb.ContainerawsAlphaClusterNetworking) *alpha.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterNetworking{
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
func ProtoToContainerawsAlphaClusterControlPlane(p *alphapb.ContainerawsAlphaClusterControlPlane) *alpha.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlane{
		Version:                   dcl.StringOrNil(p.GetVersion()),
		InstanceType:              dcl.StringOrNil(p.GetInstanceType()),
		SshConfig:                 ProtoToContainerawsAlphaClusterControlPlaneSshConfig(p.GetSshConfig()),
		ConfigEncryption:          ProtoToContainerawsAlphaClusterControlPlaneConfigEncryption(p.GetConfigEncryption()),
		IamInstanceProfile:        dcl.StringOrNil(p.GetIamInstanceProfile()),
		RootVolume:                ProtoToContainerawsAlphaClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToContainerawsAlphaClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToContainerawsAlphaClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToContainerawsAlphaClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
		ProxyConfig:               ProtoToContainerawsAlphaClusterControlPlaneProxyConfig(p.GetProxyConfig()),
		InstancePlacement:         ProtoToContainerawsAlphaClusterControlPlaneInstancePlacement(p.GetInstancePlacement()),
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
func ProtoToContainerawsAlphaClusterControlPlaneSshConfig(p *alphapb.ContainerawsAlphaClusterControlPlaneSshConfig) *alpha.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.GetEc2KeyPair()),
	}
	return obj
}

// ProtoToClusterControlPlaneConfigEncryption converts a ClusterControlPlaneConfigEncryption object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneConfigEncryption(p *alphapb.ContainerawsAlphaClusterControlPlaneConfigEncryption) *alpha.ClusterControlPlaneConfigEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneConfigEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneRootVolume(p *alphapb.ContainerawsAlphaClusterControlPlaneRootVolume) *alpha.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneMainVolume(p *alphapb.ContainerawsAlphaClusterControlPlaneMainVolume) *alpha.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneDatabaseEncryption(p *alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption) *alpha.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneAwsServicesAuthentication converts a ClusterControlPlaneAwsServicesAuthentication object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneAwsServicesAuthentication(p *alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication) *alpha.ClusterControlPlaneAwsServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.StringOrNil(p.GetRoleArn()),
		RoleSessionName: dcl.StringOrNil(p.GetRoleSessionName()),
	}
	return obj
}

// ProtoToClusterControlPlaneProxyConfig converts a ClusterControlPlaneProxyConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneProxyConfig(p *alphapb.ContainerawsAlphaClusterControlPlaneProxyConfig) *alpha.ClusterControlPlaneProxyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneProxyConfig{
		SecretArn:     dcl.StringOrNil(p.GetSecretArn()),
		SecretVersion: dcl.StringOrNil(p.GetSecretVersion()),
	}
	return obj
}

// ProtoToClusterControlPlaneInstancePlacement converts a ClusterControlPlaneInstancePlacement object from its proto representation.
func ProtoToContainerawsAlphaClusterControlPlaneInstancePlacement(p *alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacement) *alpha.ClusterControlPlaneInstancePlacement {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneInstancePlacement{
		Tenancy: ProtoToContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum(p.GetTenancy()),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization object from its proto representation.
func ProtoToContainerawsAlphaClusterAuthorization(p *alphapb.ContainerawsAlphaClusterAuthorization) *alpha.ClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerawsAlphaClusterAuthorizationAdminUsers(r))
	}
	for _, r := range p.GetAdminGroups() {
		obj.AdminGroups = append(obj.AdminGroups, *ProtoToContainerawsAlphaClusterAuthorizationAdminGroups(r))
	}
	return obj
}

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers object from its proto representation.
func ProtoToContainerawsAlphaClusterAuthorizationAdminUsers(p *alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers) *alpha.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToClusterAuthorizationAdminGroups converts a ClusterAuthorizationAdminGroups object from its proto representation.
func ProtoToContainerawsAlphaClusterAuthorizationAdminGroups(p *alphapb.ContainerawsAlphaClusterAuthorizationAdminGroups) *alpha.ClusterAuthorizationAdminGroups {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorizationAdminGroups{
		Group: dcl.StringOrNil(p.GetGroup()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterWorkloadIdentityConfig(p *alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig) *alpha.ClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.GetIssuerUri()),
		WorkloadPool:     dcl.StringOrNil(p.GetWorkloadPool()),
		IdentityProvider: dcl.StringOrNil(p.GetIdentityProvider()),
	}
	return obj
}

// ProtoToClusterFleet converts a ClusterFleet object from its proto representation.
func ProtoToContainerawsAlphaClusterFleet(p *alphapb.ContainerawsAlphaClusterFleet) *alpha.ClusterFleet {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterFleet{
		Project:    dcl.StringOrNil(p.GetProject()),
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToClusterLoggingConfig converts a ClusterLoggingConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterLoggingConfig(p *alphapb.ContainerawsAlphaClusterLoggingConfig) *alpha.ClusterLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterLoggingConfig{
		ComponentConfig: ProtoToContainerawsAlphaClusterLoggingConfigComponentConfig(p.GetComponentConfig()),
	}
	return obj
}

// ProtoToClusterLoggingConfigComponentConfig converts a ClusterLoggingConfigComponentConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterLoggingConfigComponentConfig(p *alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfig) *alpha.ClusterLoggingConfigComponentConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterLoggingConfigComponentConfig{}
	for _, r := range p.GetEnableComponents() {
		obj.EnableComponents = append(obj.EnableComponents, *ProtoToContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterMonitoringConfig converts a ClusterMonitoringConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterMonitoringConfig(p *alphapb.ContainerawsAlphaClusterMonitoringConfig) *alpha.ClusterMonitoringConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterMonitoringConfig{
		ManagedPrometheusConfig: ProtoToContainerawsAlphaClusterMonitoringConfigManagedPrometheusConfig(p.GetManagedPrometheusConfig()),
	}
	return obj
}

// ProtoToClusterMonitoringConfigManagedPrometheusConfig converts a ClusterMonitoringConfigManagedPrometheusConfig object from its proto representation.
func ProtoToContainerawsAlphaClusterMonitoringConfigManagedPrometheusConfig(p *alphapb.ContainerawsAlphaClusterMonitoringConfigManagedPrometheusConfig) *alpha.ClusterMonitoringConfigManagedPrometheusConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterMonitoringConfigManagedPrometheusConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToClusterBinaryAuthorization converts a ClusterBinaryAuthorization object from its proto representation.
func ProtoToContainerawsAlphaClusterBinaryAuthorization(p *alphapb.ContainerawsAlphaClusterBinaryAuthorization) *alpha.ClusterBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterBinaryAuthorization{
		EvaluationMode: ProtoToContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum(p.GetEvaluationMode()),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *alphapb.ContainerawsAlphaCluster) *alpha.Cluster {
	obj := &alpha.Cluster{
		Name:                   dcl.StringOrNil(p.GetName()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		Networking:             ProtoToContainerawsAlphaClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.GetAwsRegion()),
		ControlPlane:           ProtoToContainerawsAlphaClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToContainerawsAlphaClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToContainerawsAlphaClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.GetEndpoint()),
		Uid:                    dcl.StringOrNil(p.GetUid()),
		Reconciling:            dcl.Bool(p.GetReconciling()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.GetEtag()),
		WorkloadIdentityConfig: ProtoToContainerawsAlphaClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
		Fleet:                  ProtoToContainerawsAlphaClusterFleet(p.GetFleet()),
		LoggingConfig:          ProtoToContainerawsAlphaClusterLoggingConfig(p.GetLoggingConfig()),
		MonitoringConfig:       ProtoToContainerawsAlphaClusterMonitoringConfig(p.GetMonitoringConfig()),
		BinaryAuthorization:    ProtoToContainerawsAlphaClusterBinaryAuthorization(p.GetBinaryAuthorization()),
	}
	return obj
}

// ClusterControlPlaneRootVolumeVolumeTypeEnumToProto converts a ClusterControlPlaneRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnumToProto(e *alpha.ClusterControlPlaneRootVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum_value["ClusterControlPlaneRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnum(0)
}

// ClusterControlPlaneMainVolumeVolumeTypeEnumToProto converts a ClusterControlPlaneMainVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnumToProto(e *alpha.ClusterControlPlaneMainVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum_value["ClusterControlPlaneMainVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnum(0)
}

// ClusterControlPlaneInstancePlacementTenancyEnumToProto converts a ClusterControlPlaneInstancePlacementTenancyEnum enum to its proto representation.
func ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnumToProto(e *alpha.ClusterControlPlaneInstancePlacementTenancyEnum) alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum_value["ClusterControlPlaneInstancePlacementTenancyEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnum(0)
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func ContainerawsAlphaClusterStateEnumToProto(e *alpha.ClusterStateEnum) alphapb.ContainerawsAlphaClusterStateEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterStateEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterStateEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterStateEnum(0)
}

// ClusterLoggingConfigComponentConfigEnableComponentsEnumToProto converts a ClusterLoggingConfigComponentConfigEnableComponentsEnum enum to its proto representation.
func ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnumToProto(e *alpha.ClusterLoggingConfigComponentConfigEnableComponentsEnum) alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum_value["ClusterLoggingConfigComponentConfigEnableComponentsEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(0)
}

// ClusterBinaryAuthorizationEvaluationModeEnumToProto converts a ClusterBinaryAuthorizationEvaluationModeEnum enum to its proto representation.
func ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnumToProto(e *alpha.ClusterBinaryAuthorizationEvaluationModeEnum) alphapb.ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum_value["ClusterBinaryAuthorizationEvaluationModeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum(v)
	}
	return alphapb.ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnum(0)
}

// ClusterNetworkingToProto converts a ClusterNetworking object to its proto representation.
func ContainerawsAlphaClusterNetworkingToProto(o *alpha.ClusterNetworking) *alphapb.ContainerawsAlphaClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterNetworking{}
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
func ContainerawsAlphaClusterControlPlaneToProto(o *alpha.ClusterControlPlane) *alphapb.ContainerawsAlphaClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlane{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetInstanceType(dcl.ValueOrEmptyString(o.InstanceType))
	p.SetSshConfig(ContainerawsAlphaClusterControlPlaneSshConfigToProto(o.SshConfig))
	p.SetConfigEncryption(ContainerawsAlphaClusterControlPlaneConfigEncryptionToProto(o.ConfigEncryption))
	p.SetIamInstanceProfile(dcl.ValueOrEmptyString(o.IamInstanceProfile))
	p.SetRootVolume(ContainerawsAlphaClusterControlPlaneRootVolumeToProto(o.RootVolume))
	p.SetMainVolume(ContainerawsAlphaClusterControlPlaneMainVolumeToProto(o.MainVolume))
	p.SetDatabaseEncryption(ContainerawsAlphaClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption))
	p.SetAwsServicesAuthentication(ContainerawsAlphaClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication))
	p.SetProxyConfig(ContainerawsAlphaClusterControlPlaneProxyConfigToProto(o.ProxyConfig))
	p.SetInstancePlacement(ContainerawsAlphaClusterControlPlaneInstancePlacementToProto(o.InstancePlacement))
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
func ContainerawsAlphaClusterControlPlaneSshConfigToProto(o *alpha.ClusterControlPlaneSshConfig) *alphapb.ContainerawsAlphaClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneSshConfig{}
	p.SetEc2KeyPair(dcl.ValueOrEmptyString(o.Ec2KeyPair))
	return p
}

// ClusterControlPlaneConfigEncryptionToProto converts a ClusterControlPlaneConfigEncryption object to its proto representation.
func ContainerawsAlphaClusterControlPlaneConfigEncryptionToProto(o *alpha.ClusterControlPlaneConfigEncryption) *alphapb.ContainerawsAlphaClusterControlPlaneConfigEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneConfigEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume object to its proto representation.
func ContainerawsAlphaClusterControlPlaneRootVolumeToProto(o *alpha.ClusterControlPlaneRootVolume) *alphapb.ContainerawsAlphaClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsAlphaClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume object to its proto representation.
func ContainerawsAlphaClusterControlPlaneMainVolumeToProto(o *alpha.ClusterControlPlaneMainVolume) *alphapb.ContainerawsAlphaClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneMainVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsAlphaClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption object to its proto representation.
func ContainerawsAlphaClusterControlPlaneDatabaseEncryptionToProto(o *alpha.ClusterControlPlaneDatabaseEncryption) *alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneDatabaseEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneAwsServicesAuthenticationToProto converts a ClusterControlPlaneAwsServicesAuthentication object to its proto representation.
func ContainerawsAlphaClusterControlPlaneAwsServicesAuthenticationToProto(o *alpha.ClusterControlPlaneAwsServicesAuthentication) *alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneAwsServicesAuthentication{}
	p.SetRoleArn(dcl.ValueOrEmptyString(o.RoleArn))
	p.SetRoleSessionName(dcl.ValueOrEmptyString(o.RoleSessionName))
	return p
}

// ClusterControlPlaneProxyConfigToProto converts a ClusterControlPlaneProxyConfig object to its proto representation.
func ContainerawsAlphaClusterControlPlaneProxyConfigToProto(o *alpha.ClusterControlPlaneProxyConfig) *alphapb.ContainerawsAlphaClusterControlPlaneProxyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneProxyConfig{}
	p.SetSecretArn(dcl.ValueOrEmptyString(o.SecretArn))
	p.SetSecretVersion(dcl.ValueOrEmptyString(o.SecretVersion))
	return p
}

// ClusterControlPlaneInstancePlacementToProto converts a ClusterControlPlaneInstancePlacement object to its proto representation.
func ContainerawsAlphaClusterControlPlaneInstancePlacementToProto(o *alpha.ClusterControlPlaneInstancePlacement) *alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacement {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterControlPlaneInstancePlacement{}
	p.SetTenancy(ContainerawsAlphaClusterControlPlaneInstancePlacementTenancyEnumToProto(o.Tenancy))
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization object to its proto representation.
func ContainerawsAlphaClusterAuthorizationToProto(o *alpha.ClusterAuthorization) *alphapb.ContainerawsAlphaClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterAuthorization{}
	sAdminUsers := make([]*alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers, len(o.AdminUsers))
	for i, r := range o.AdminUsers {
		sAdminUsers[i] = ContainerawsAlphaClusterAuthorizationAdminUsersToProto(&r)
	}
	p.SetAdminUsers(sAdminUsers)
	sAdminGroups := make([]*alphapb.ContainerawsAlphaClusterAuthorizationAdminGroups, len(o.AdminGroups))
	for i, r := range o.AdminGroups {
		sAdminGroups[i] = ContainerawsAlphaClusterAuthorizationAdminGroupsToProto(&r)
	}
	p.SetAdminGroups(sAdminGroups)
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers object to its proto representation.
func ContainerawsAlphaClusterAuthorizationAdminUsersToProto(o *alpha.ClusterAuthorizationAdminUsers) *alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterAuthorizationAdminUsers{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ClusterAuthorizationAdminGroupsToProto converts a ClusterAuthorizationAdminGroups object to its proto representation.
func ContainerawsAlphaClusterAuthorizationAdminGroupsToProto(o *alpha.ClusterAuthorizationAdminGroups) *alphapb.ContainerawsAlphaClusterAuthorizationAdminGroups {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterAuthorizationAdminGroups{}
	p.SetGroup(dcl.ValueOrEmptyString(o.Group))
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig object to its proto representation.
func ContainerawsAlphaClusterWorkloadIdentityConfigToProto(o *alpha.ClusterWorkloadIdentityConfig) *alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterWorkloadIdentityConfig{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetWorkloadPool(dcl.ValueOrEmptyString(o.WorkloadPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// ClusterFleetToProto converts a ClusterFleet object to its proto representation.
func ContainerawsAlphaClusterFleetToProto(o *alpha.ClusterFleet) *alphapb.ContainerawsAlphaClusterFleet {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterFleet{}
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// ClusterLoggingConfigToProto converts a ClusterLoggingConfig object to its proto representation.
func ContainerawsAlphaClusterLoggingConfigToProto(o *alpha.ClusterLoggingConfig) *alphapb.ContainerawsAlphaClusterLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterLoggingConfig{}
	p.SetComponentConfig(ContainerawsAlphaClusterLoggingConfigComponentConfigToProto(o.ComponentConfig))
	return p
}

// ClusterLoggingConfigComponentConfigToProto converts a ClusterLoggingConfigComponentConfig object to its proto representation.
func ContainerawsAlphaClusterLoggingConfigComponentConfigToProto(o *alpha.ClusterLoggingConfigComponentConfig) *alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfig{}
	sEnableComponents := make([]alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum, len(o.EnableComponents))
	for i, r := range o.EnableComponents {
		sEnableComponents[i] = alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(alphapb.ContainerawsAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum_value[string(r)])
	}
	p.SetEnableComponents(sEnableComponents)
	return p
}

// ClusterMonitoringConfigToProto converts a ClusterMonitoringConfig object to its proto representation.
func ContainerawsAlphaClusterMonitoringConfigToProto(o *alpha.ClusterMonitoringConfig) *alphapb.ContainerawsAlphaClusterMonitoringConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterMonitoringConfig{}
	p.SetManagedPrometheusConfig(ContainerawsAlphaClusterMonitoringConfigManagedPrometheusConfigToProto(o.ManagedPrometheusConfig))
	return p
}

// ClusterMonitoringConfigManagedPrometheusConfigToProto converts a ClusterMonitoringConfigManagedPrometheusConfig object to its proto representation.
func ContainerawsAlphaClusterMonitoringConfigManagedPrometheusConfigToProto(o *alpha.ClusterMonitoringConfigManagedPrometheusConfig) *alphapb.ContainerawsAlphaClusterMonitoringConfigManagedPrometheusConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterMonitoringConfigManagedPrometheusConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ClusterBinaryAuthorizationToProto converts a ClusterBinaryAuthorization object to its proto representation.
func ContainerawsAlphaClusterBinaryAuthorizationToProto(o *alpha.ClusterBinaryAuthorization) *alphapb.ContainerawsAlphaClusterBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaClusterBinaryAuthorization{}
	p.SetEvaluationMode(ContainerawsAlphaClusterBinaryAuthorizationEvaluationModeEnumToProto(o.EvaluationMode))
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *alpha.Cluster) *alphapb.ContainerawsAlphaCluster {
	p := &alphapb.ContainerawsAlphaCluster{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetNetworking(ContainerawsAlphaClusterNetworkingToProto(resource.Networking))
	p.SetAwsRegion(dcl.ValueOrEmptyString(resource.AwsRegion))
	p.SetControlPlane(ContainerawsAlphaClusterControlPlaneToProto(resource.ControlPlane))
	p.SetAuthorization(ContainerawsAlphaClusterAuthorizationToProto(resource.Authorization))
	p.SetState(ContainerawsAlphaClusterStateEnumToProto(resource.State))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkloadIdentityConfig(ContainerawsAlphaClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFleet(ContainerawsAlphaClusterFleetToProto(resource.Fleet))
	p.SetLoggingConfig(ContainerawsAlphaClusterLoggingConfigToProto(resource.LoggingConfig))
	p.SetMonitoringConfig(ContainerawsAlphaClusterMonitoringConfigToProto(resource.MonitoringConfig))
	p.SetBinaryAuthorization(ContainerawsAlphaClusterBinaryAuthorizationToProto(resource.BinaryAuthorization))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerawsAlphaClusterRequest) (*alphapb.ContainerawsAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyContainerawsAlphaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerawsAlphaCluster(ctx context.Context, request *alphapb.ApplyContainerawsAlphaClusterRequest) (*alphapb.ContainerawsAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerawsAlphaCluster(ctx context.Context, request *alphapb.DeleteContainerawsAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerawsAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerawsAlphaCluster(ctx context.Context, request *alphapb.ListContainerawsAlphaClusterRequest) (*alphapb.ListContainerawsAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerawsAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListContainerawsAlphaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
