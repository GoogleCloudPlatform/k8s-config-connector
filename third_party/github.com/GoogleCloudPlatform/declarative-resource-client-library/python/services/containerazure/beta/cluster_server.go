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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/beta/containerazure_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/beta"
)

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToContainerazureBetaClusterStateEnum(e betapb.ContainerazureBetaClusterStateEnum) *beta.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerazureBetaClusterStateEnum_name[int32(e)]; ok {
		e := beta.ClusterStateEnum(n[len("ContainerazureBetaClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterLoggingConfigComponentConfigEnableComponentsEnum converts a ClusterLoggingConfigComponentConfigEnableComponentsEnum enum from its proto representation.
func ProtoToContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(e betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum) *beta.ClusterLoggingConfigComponentConfigEnableComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum_name[int32(e)]; ok {
		e := beta.ClusterLoggingConfigComponentConfigEnableComponentsEnum(n[len("ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterAzureServicesAuthentication converts a ClusterAzureServicesAuthentication object from its proto representation.
func ProtoToContainerazureBetaClusterAzureServicesAuthentication(p *betapb.ContainerazureBetaClusterAzureServicesAuthentication) *beta.ClusterAzureServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAzureServicesAuthentication{
		TenantId:      dcl.StringOrNil(p.GetTenantId()),
		ApplicationId: dcl.StringOrNil(p.GetApplicationId()),
	}
	return obj
}

// ProtoToClusterNetworking converts a ClusterNetworking object from its proto representation.
func ProtoToContainerazureBetaClusterNetworking(p *betapb.ContainerazureBetaClusterNetworking) *beta.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterNetworking{
		VirtualNetworkId: dcl.StringOrNil(p.GetVirtualNetworkId()),
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
func ProtoToContainerazureBetaClusterControlPlane(p *betapb.ContainerazureBetaClusterControlPlane) *beta.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlane{
		Version:            dcl.StringOrNil(p.GetVersion()),
		SubnetId:           dcl.StringOrNil(p.GetSubnetId()),
		VmSize:             dcl.StringOrNil(p.GetVmSize()),
		SshConfig:          ProtoToContainerazureBetaClusterControlPlaneSshConfig(p.GetSshConfig()),
		RootVolume:         ProtoToContainerazureBetaClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:         ProtoToContainerazureBetaClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption: ProtoToContainerazureBetaClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		ProxyConfig:        ProtoToContainerazureBetaClusterControlPlaneProxyConfig(p.GetProxyConfig()),
	}
	for _, r := range p.GetReplicaPlacements() {
		obj.ReplicaPlacements = append(obj.ReplicaPlacements, *ProtoToContainerazureBetaClusterControlPlaneReplicaPlacements(r))
	}
	return obj
}

// ProtoToClusterControlPlaneSshConfig converts a ClusterControlPlaneSshConfig object from its proto representation.
func ProtoToContainerazureBetaClusterControlPlaneSshConfig(p *betapb.ContainerazureBetaClusterControlPlaneSshConfig) *beta.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.GetAuthorizedKey()),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume object from its proto representation.
func ProtoToContainerazureBetaClusterControlPlaneRootVolume(p *betapb.ContainerazureBetaClusterControlPlaneRootVolume) *beta.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume object from its proto representation.
func ProtoToContainerazureBetaClusterControlPlaneMainVolume(p *betapb.ContainerazureBetaClusterControlPlaneMainVolume) *beta.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption object from its proto representation.
func ProtoToContainerazureBetaClusterControlPlaneDatabaseEncryption(p *betapb.ContainerazureBetaClusterControlPlaneDatabaseEncryption) *beta.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneDatabaseEncryption{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToClusterControlPlaneProxyConfig converts a ClusterControlPlaneProxyConfig object from its proto representation.
func ProtoToContainerazureBetaClusterControlPlaneProxyConfig(p *betapb.ContainerazureBetaClusterControlPlaneProxyConfig) *beta.ClusterControlPlaneProxyConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneProxyConfig{
		ResourceGroupId: dcl.StringOrNil(p.GetResourceGroupId()),
		SecretId:        dcl.StringOrNil(p.GetSecretId()),
	}
	return obj
}

// ProtoToClusterControlPlaneReplicaPlacements converts a ClusterControlPlaneReplicaPlacements object from its proto representation.
func ProtoToContainerazureBetaClusterControlPlaneReplicaPlacements(p *betapb.ContainerazureBetaClusterControlPlaneReplicaPlacements) *beta.ClusterControlPlaneReplicaPlacements {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterControlPlaneReplicaPlacements{
		SubnetId:              dcl.StringOrNil(p.GetSubnetId()),
		AzureAvailabilityZone: dcl.StringOrNil(p.GetAzureAvailabilityZone()),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization object from its proto representation.
func ProtoToContainerazureBetaClusterAuthorization(p *betapb.ContainerazureBetaClusterAuthorization) *beta.ClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerazureBetaClusterAuthorizationAdminUsers(r))
	}
	for _, r := range p.GetAdminGroups() {
		obj.AdminGroups = append(obj.AdminGroups, *ProtoToContainerazureBetaClusterAuthorizationAdminGroups(r))
	}
	return obj
}

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers object from its proto representation.
func ProtoToContainerazureBetaClusterAuthorizationAdminUsers(p *betapb.ContainerazureBetaClusterAuthorizationAdminUsers) *beta.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToClusterAuthorizationAdminGroups converts a ClusterAuthorizationAdminGroups object from its proto representation.
func ProtoToContainerazureBetaClusterAuthorizationAdminGroups(p *betapb.ContainerazureBetaClusterAuthorizationAdminGroups) *beta.ClusterAuthorizationAdminGroups {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterAuthorizationAdminGroups{
		Group: dcl.StringOrNil(p.GetGroup()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig object from its proto representation.
func ProtoToContainerazureBetaClusterWorkloadIdentityConfig(p *betapb.ContainerazureBetaClusterWorkloadIdentityConfig) *beta.ClusterWorkloadIdentityConfig {
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
func ProtoToContainerazureBetaClusterFleet(p *betapb.ContainerazureBetaClusterFleet) *beta.ClusterFleet {
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
func ProtoToContainerazureBetaClusterLoggingConfig(p *betapb.ContainerazureBetaClusterLoggingConfig) *beta.ClusterLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterLoggingConfig{
		ComponentConfig: ProtoToContainerazureBetaClusterLoggingConfigComponentConfig(p.GetComponentConfig()),
	}
	return obj
}

// ProtoToClusterLoggingConfigComponentConfig converts a ClusterLoggingConfigComponentConfig object from its proto representation.
func ProtoToContainerazureBetaClusterLoggingConfigComponentConfig(p *betapb.ContainerazureBetaClusterLoggingConfigComponentConfig) *beta.ClusterLoggingConfigComponentConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterLoggingConfigComponentConfig{}
	for _, r := range p.GetEnableComponents() {
		obj.EnableComponents = append(obj.EnableComponents, *ProtoToContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterMonitoringConfig converts a ClusterMonitoringConfig object from its proto representation.
func ProtoToContainerazureBetaClusterMonitoringConfig(p *betapb.ContainerazureBetaClusterMonitoringConfig) *beta.ClusterMonitoringConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMonitoringConfig{
		ManagedPrometheusConfig: ProtoToContainerazureBetaClusterMonitoringConfigManagedPrometheusConfig(p.GetManagedPrometheusConfig()),
	}
	return obj
}

// ProtoToClusterMonitoringConfigManagedPrometheusConfig converts a ClusterMonitoringConfigManagedPrometheusConfig object from its proto representation.
func ProtoToContainerazureBetaClusterMonitoringConfigManagedPrometheusConfig(p *betapb.ContainerazureBetaClusterMonitoringConfigManagedPrometheusConfig) *beta.ClusterMonitoringConfigManagedPrometheusConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ClusterMonitoringConfigManagedPrometheusConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *betapb.ContainerazureBetaCluster) *beta.Cluster {
	obj := &beta.Cluster{
		Name:                        dcl.StringOrNil(p.GetName()),
		Description:                 dcl.StringOrNil(p.GetDescription()),
		AzureRegion:                 dcl.StringOrNil(p.GetAzureRegion()),
		ResourceGroupId:             dcl.StringOrNil(p.GetResourceGroupId()),
		Client:                      dcl.StringOrNil(p.GetClient()),
		AzureServicesAuthentication: ProtoToContainerazureBetaClusterAzureServicesAuthentication(p.GetAzureServicesAuthentication()),
		Networking:                  ProtoToContainerazureBetaClusterNetworking(p.GetNetworking()),
		ControlPlane:                ProtoToContainerazureBetaClusterControlPlane(p.GetControlPlane()),
		Authorization:               ProtoToContainerazureBetaClusterAuthorization(p.GetAuthorization()),
		State:                       ProtoToContainerazureBetaClusterStateEnum(p.GetState()),
		Endpoint:                    dcl.StringOrNil(p.GetEndpoint()),
		Uid:                         dcl.StringOrNil(p.GetUid()),
		Reconciling:                 dcl.Bool(p.GetReconciling()),
		CreateTime:                  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                  dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                        dcl.StringOrNil(p.GetEtag()),
		WorkloadIdentityConfig:      ProtoToContainerazureBetaClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                     dcl.StringOrNil(p.GetProject()),
		Location:                    dcl.StringOrNil(p.GetLocation()),
		Fleet:                       ProtoToContainerazureBetaClusterFleet(p.GetFleet()),
		LoggingConfig:               ProtoToContainerazureBetaClusterLoggingConfig(p.GetLoggingConfig()),
		MonitoringConfig:            ProtoToContainerazureBetaClusterMonitoringConfig(p.GetMonitoringConfig()),
	}
	return obj
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func ContainerazureBetaClusterStateEnumToProto(e *beta.ClusterStateEnum) betapb.ContainerazureBetaClusterStateEnum {
	if e == nil {
		return betapb.ContainerazureBetaClusterStateEnum(0)
	}
	if v, ok := betapb.ContainerazureBetaClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return betapb.ContainerazureBetaClusterStateEnum(v)
	}
	return betapb.ContainerazureBetaClusterStateEnum(0)
}

// ClusterLoggingConfigComponentConfigEnableComponentsEnumToProto converts a ClusterLoggingConfigComponentConfigEnableComponentsEnum enum to its proto representation.
func ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnumToProto(e *beta.ClusterLoggingConfigComponentConfigEnableComponentsEnum) betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum {
	if e == nil {
		return betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(0)
	}
	if v, ok := betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum_value["ClusterLoggingConfigComponentConfigEnableComponentsEnum"+string(*e)]; ok {
		return betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(v)
	}
	return betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(0)
}

// ClusterAzureServicesAuthenticationToProto converts a ClusterAzureServicesAuthentication object to its proto representation.
func ContainerazureBetaClusterAzureServicesAuthenticationToProto(o *beta.ClusterAzureServicesAuthentication) *betapb.ContainerazureBetaClusterAzureServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterAzureServicesAuthentication{}
	p.SetTenantId(dcl.ValueOrEmptyString(o.TenantId))
	p.SetApplicationId(dcl.ValueOrEmptyString(o.ApplicationId))
	return p
}

// ClusterNetworkingToProto converts a ClusterNetworking object to its proto representation.
func ContainerazureBetaClusterNetworkingToProto(o *beta.ClusterNetworking) *betapb.ContainerazureBetaClusterNetworking {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterNetworking{}
	p.SetVirtualNetworkId(dcl.ValueOrEmptyString(o.VirtualNetworkId))
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
func ContainerazureBetaClusterControlPlaneToProto(o *beta.ClusterControlPlane) *betapb.ContainerazureBetaClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterControlPlane{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetSubnetId(dcl.ValueOrEmptyString(o.SubnetId))
	p.SetVmSize(dcl.ValueOrEmptyString(o.VmSize))
	p.SetSshConfig(ContainerazureBetaClusterControlPlaneSshConfigToProto(o.SshConfig))
	p.SetRootVolume(ContainerazureBetaClusterControlPlaneRootVolumeToProto(o.RootVolume))
	p.SetMainVolume(ContainerazureBetaClusterControlPlaneMainVolumeToProto(o.MainVolume))
	p.SetDatabaseEncryption(ContainerazureBetaClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption))
	p.SetProxyConfig(ContainerazureBetaClusterControlPlaneProxyConfigToProto(o.ProxyConfig))
	mTags := make(map[string]string, len(o.Tags))
	for k, r := range o.Tags {
		mTags[k] = r
	}
	p.SetTags(mTags)
	sReplicaPlacements := make([]*betapb.ContainerazureBetaClusterControlPlaneReplicaPlacements, len(o.ReplicaPlacements))
	for i, r := range o.ReplicaPlacements {
		sReplicaPlacements[i] = ContainerazureBetaClusterControlPlaneReplicaPlacementsToProto(&r)
	}
	p.SetReplicaPlacements(sReplicaPlacements)
	return p
}

// ClusterControlPlaneSshConfigToProto converts a ClusterControlPlaneSshConfig object to its proto representation.
func ContainerazureBetaClusterControlPlaneSshConfigToProto(o *beta.ClusterControlPlaneSshConfig) *betapb.ContainerazureBetaClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterControlPlaneSshConfig{}
	p.SetAuthorizedKey(dcl.ValueOrEmptyString(o.AuthorizedKey))
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume object to its proto representation.
func ContainerazureBetaClusterControlPlaneRootVolumeToProto(o *beta.ClusterControlPlaneRootVolume) *betapb.ContainerazureBetaClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterControlPlaneRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume object to its proto representation.
func ContainerazureBetaClusterControlPlaneMainVolumeToProto(o *beta.ClusterControlPlaneMainVolume) *betapb.ContainerazureBetaClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterControlPlaneMainVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption object to its proto representation.
func ContainerazureBetaClusterControlPlaneDatabaseEncryptionToProto(o *beta.ClusterControlPlaneDatabaseEncryption) *betapb.ContainerazureBetaClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterControlPlaneDatabaseEncryption{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// ClusterControlPlaneProxyConfigToProto converts a ClusterControlPlaneProxyConfig object to its proto representation.
func ContainerazureBetaClusterControlPlaneProxyConfigToProto(o *beta.ClusterControlPlaneProxyConfig) *betapb.ContainerazureBetaClusterControlPlaneProxyConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterControlPlaneProxyConfig{}
	p.SetResourceGroupId(dcl.ValueOrEmptyString(o.ResourceGroupId))
	p.SetSecretId(dcl.ValueOrEmptyString(o.SecretId))
	return p
}

// ClusterControlPlaneReplicaPlacementsToProto converts a ClusterControlPlaneReplicaPlacements object to its proto representation.
func ContainerazureBetaClusterControlPlaneReplicaPlacementsToProto(o *beta.ClusterControlPlaneReplicaPlacements) *betapb.ContainerazureBetaClusterControlPlaneReplicaPlacements {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterControlPlaneReplicaPlacements{}
	p.SetSubnetId(dcl.ValueOrEmptyString(o.SubnetId))
	p.SetAzureAvailabilityZone(dcl.ValueOrEmptyString(o.AzureAvailabilityZone))
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization object to its proto representation.
func ContainerazureBetaClusterAuthorizationToProto(o *beta.ClusterAuthorization) *betapb.ContainerazureBetaClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterAuthorization{}
	sAdminUsers := make([]*betapb.ContainerazureBetaClusterAuthorizationAdminUsers, len(o.AdminUsers))
	for i, r := range o.AdminUsers {
		sAdminUsers[i] = ContainerazureBetaClusterAuthorizationAdminUsersToProto(&r)
	}
	p.SetAdminUsers(sAdminUsers)
	sAdminGroups := make([]*betapb.ContainerazureBetaClusterAuthorizationAdminGroups, len(o.AdminGroups))
	for i, r := range o.AdminGroups {
		sAdminGroups[i] = ContainerazureBetaClusterAuthorizationAdminGroupsToProto(&r)
	}
	p.SetAdminGroups(sAdminGroups)
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers object to its proto representation.
func ContainerazureBetaClusterAuthorizationAdminUsersToProto(o *beta.ClusterAuthorizationAdminUsers) *betapb.ContainerazureBetaClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterAuthorizationAdminUsers{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ClusterAuthorizationAdminGroupsToProto converts a ClusterAuthorizationAdminGroups object to its proto representation.
func ContainerazureBetaClusterAuthorizationAdminGroupsToProto(o *beta.ClusterAuthorizationAdminGroups) *betapb.ContainerazureBetaClusterAuthorizationAdminGroups {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterAuthorizationAdminGroups{}
	p.SetGroup(dcl.ValueOrEmptyString(o.Group))
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig object to its proto representation.
func ContainerazureBetaClusterWorkloadIdentityConfigToProto(o *beta.ClusterWorkloadIdentityConfig) *betapb.ContainerazureBetaClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterWorkloadIdentityConfig{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetWorkloadPool(dcl.ValueOrEmptyString(o.WorkloadPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// ClusterFleetToProto converts a ClusterFleet object to its proto representation.
func ContainerazureBetaClusterFleetToProto(o *beta.ClusterFleet) *betapb.ContainerazureBetaClusterFleet {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterFleet{}
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// ClusterLoggingConfigToProto converts a ClusterLoggingConfig object to its proto representation.
func ContainerazureBetaClusterLoggingConfigToProto(o *beta.ClusterLoggingConfig) *betapb.ContainerazureBetaClusterLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterLoggingConfig{}
	p.SetComponentConfig(ContainerazureBetaClusterLoggingConfigComponentConfigToProto(o.ComponentConfig))
	return p
}

// ClusterLoggingConfigComponentConfigToProto converts a ClusterLoggingConfigComponentConfig object to its proto representation.
func ContainerazureBetaClusterLoggingConfigComponentConfigToProto(o *beta.ClusterLoggingConfigComponentConfig) *betapb.ContainerazureBetaClusterLoggingConfigComponentConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterLoggingConfigComponentConfig{}
	sEnableComponents := make([]betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum, len(o.EnableComponents))
	for i, r := range o.EnableComponents {
		sEnableComponents[i] = betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum(betapb.ContainerazureBetaClusterLoggingConfigComponentConfigEnableComponentsEnum_value[string(r)])
	}
	p.SetEnableComponents(sEnableComponents)
	return p
}

// ClusterMonitoringConfigToProto converts a ClusterMonitoringConfig object to its proto representation.
func ContainerazureBetaClusterMonitoringConfigToProto(o *beta.ClusterMonitoringConfig) *betapb.ContainerazureBetaClusterMonitoringConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterMonitoringConfig{}
	p.SetManagedPrometheusConfig(ContainerazureBetaClusterMonitoringConfigManagedPrometheusConfigToProto(o.ManagedPrometheusConfig))
	return p
}

// ClusterMonitoringConfigManagedPrometheusConfigToProto converts a ClusterMonitoringConfigManagedPrometheusConfig object to its proto representation.
func ContainerazureBetaClusterMonitoringConfigManagedPrometheusConfigToProto(o *beta.ClusterMonitoringConfigManagedPrometheusConfig) *betapb.ContainerazureBetaClusterMonitoringConfigManagedPrometheusConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaClusterMonitoringConfigManagedPrometheusConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *beta.Cluster) *betapb.ContainerazureBetaCluster {
	p := &betapb.ContainerazureBetaCluster{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetAzureRegion(dcl.ValueOrEmptyString(resource.AzureRegion))
	p.SetResourceGroupId(dcl.ValueOrEmptyString(resource.ResourceGroupId))
	p.SetClient(dcl.ValueOrEmptyString(resource.Client))
	p.SetAzureServicesAuthentication(ContainerazureBetaClusterAzureServicesAuthenticationToProto(resource.AzureServicesAuthentication))
	p.SetNetworking(ContainerazureBetaClusterNetworkingToProto(resource.Networking))
	p.SetControlPlane(ContainerazureBetaClusterControlPlaneToProto(resource.ControlPlane))
	p.SetAuthorization(ContainerazureBetaClusterAuthorizationToProto(resource.Authorization))
	p.SetState(ContainerazureBetaClusterStateEnumToProto(resource.State))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkloadIdentityConfig(ContainerazureBetaClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFleet(ContainerazureBetaClusterFleetToProto(resource.Fleet))
	p.SetLoggingConfig(ContainerazureBetaClusterLoggingConfigToProto(resource.LoggingConfig))
	p.SetMonitoringConfig(ContainerazureBetaClusterMonitoringConfigToProto(resource.MonitoringConfig))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *beta.Client, request *betapb.ApplyContainerazureBetaClusterRequest) (*betapb.ContainerazureBetaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyContainerazureBetaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerazureBetaCluster(ctx context.Context, request *betapb.ApplyContainerazureBetaClusterRequest) (*betapb.ContainerazureBetaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerazureBetaCluster(ctx context.Context, request *betapb.DeleteContainerazureBetaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerazureBetaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerazureBetaCluster(ctx context.Context, request *betapb.ListContainerazureBetaClusterRequest) (*betapb.ListContainerazureBetaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContainerazureBetaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListContainerazureBetaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
