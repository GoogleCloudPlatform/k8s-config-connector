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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/alpha/containerazure_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/alpha"
)

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToContainerazureAlphaClusterStateEnum(e alphapb.ContainerazureAlphaClusterStateEnum) *alpha.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerazureAlphaClusterStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStateEnum(n[len("ContainerazureAlphaClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterLoggingConfigComponentConfigEnableComponentsEnum converts a ClusterLoggingConfigComponentConfigEnableComponentsEnum enum from its proto representation.
func ProtoToContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(e alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum) *alpha.ClusterLoggingConfigComponentConfigEnableComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum_name[int32(e)]; ok {
		e := alpha.ClusterLoggingConfigComponentConfigEnableComponentsEnum(n[len("ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterAzureServicesAuthentication converts a ClusterAzureServicesAuthentication object from its proto representation.
func ProtoToContainerazureAlphaClusterAzureServicesAuthentication(p *alphapb.ContainerazureAlphaClusterAzureServicesAuthentication) *alpha.ClusterAzureServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAzureServicesAuthentication{
		TenantId:      dcl.StringOrNil(p.GetTenantId()),
		ApplicationId: dcl.StringOrNil(p.GetApplicationId()),
	}
	return obj
}

// ProtoToClusterNetworking converts a ClusterNetworking object from its proto representation.
func ProtoToContainerazureAlphaClusterNetworking(p *alphapb.ContainerazureAlphaClusterNetworking) *alpha.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterNetworking{
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
func ProtoToContainerazureAlphaClusterControlPlane(p *alphapb.ContainerazureAlphaClusterControlPlane) *alpha.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlane{
		Version:            dcl.StringOrNil(p.GetVersion()),
		SubnetId:           dcl.StringOrNil(p.GetSubnetId()),
		VmSize:             dcl.StringOrNil(p.GetVmSize()),
		SshConfig:          ProtoToContainerazureAlphaClusterControlPlaneSshConfig(p.GetSshConfig()),
		RootVolume:         ProtoToContainerazureAlphaClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:         ProtoToContainerazureAlphaClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption: ProtoToContainerazureAlphaClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		ProxyConfig:        ProtoToContainerazureAlphaClusterControlPlaneProxyConfig(p.GetProxyConfig()),
	}
	for _, r := range p.GetReplicaPlacements() {
		obj.ReplicaPlacements = append(obj.ReplicaPlacements, *ProtoToContainerazureAlphaClusterControlPlaneReplicaPlacements(r))
	}
	return obj
}

// ProtoToClusterControlPlaneSshConfig converts a ClusterControlPlaneSshConfig object from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneSshConfig(p *alphapb.ContainerazureAlphaClusterControlPlaneSshConfig) *alpha.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.GetAuthorizedKey()),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume object from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneRootVolume(p *alphapb.ContainerazureAlphaClusterControlPlaneRootVolume) *alpha.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume object from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneMainVolume(p *alphapb.ContainerazureAlphaClusterControlPlaneMainVolume) *alpha.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption object from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneDatabaseEncryption(p *alphapb.ContainerazureAlphaClusterControlPlaneDatabaseEncryption) *alpha.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneDatabaseEncryption{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToClusterControlPlaneProxyConfig converts a ClusterControlPlaneProxyConfig object from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneProxyConfig(p *alphapb.ContainerazureAlphaClusterControlPlaneProxyConfig) *alpha.ClusterControlPlaneProxyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneProxyConfig{
		ResourceGroupId: dcl.StringOrNil(p.GetResourceGroupId()),
		SecretId:        dcl.StringOrNil(p.GetSecretId()),
	}
	return obj
}

// ProtoToClusterControlPlaneReplicaPlacements converts a ClusterControlPlaneReplicaPlacements object from its proto representation.
func ProtoToContainerazureAlphaClusterControlPlaneReplicaPlacements(p *alphapb.ContainerazureAlphaClusterControlPlaneReplicaPlacements) *alpha.ClusterControlPlaneReplicaPlacements {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterControlPlaneReplicaPlacements{
		SubnetId:              dcl.StringOrNil(p.GetSubnetId()),
		AzureAvailabilityZone: dcl.StringOrNil(p.GetAzureAvailabilityZone()),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization object from its proto representation.
func ProtoToContainerazureAlphaClusterAuthorization(p *alphapb.ContainerazureAlphaClusterAuthorization) *alpha.ClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerazureAlphaClusterAuthorizationAdminUsers(r))
	}
	for _, r := range p.GetAdminGroups() {
		obj.AdminGroups = append(obj.AdminGroups, *ProtoToContainerazureAlphaClusterAuthorizationAdminGroups(r))
	}
	return obj
}

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers object from its proto representation.
func ProtoToContainerazureAlphaClusterAuthorizationAdminUsers(p *alphapb.ContainerazureAlphaClusterAuthorizationAdminUsers) *alpha.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToClusterAuthorizationAdminGroups converts a ClusterAuthorizationAdminGroups object from its proto representation.
func ProtoToContainerazureAlphaClusterAuthorizationAdminGroups(p *alphapb.ContainerazureAlphaClusterAuthorizationAdminGroups) *alpha.ClusterAuthorizationAdminGroups {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterAuthorizationAdminGroups{
		Group: dcl.StringOrNil(p.GetGroup()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig object from its proto representation.
func ProtoToContainerazureAlphaClusterWorkloadIdentityConfig(p *alphapb.ContainerazureAlphaClusterWorkloadIdentityConfig) *alpha.ClusterWorkloadIdentityConfig {
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
func ProtoToContainerazureAlphaClusterFleet(p *alphapb.ContainerazureAlphaClusterFleet) *alpha.ClusterFleet {
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
func ProtoToContainerazureAlphaClusterLoggingConfig(p *alphapb.ContainerazureAlphaClusterLoggingConfig) *alpha.ClusterLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterLoggingConfig{
		ComponentConfig: ProtoToContainerazureAlphaClusterLoggingConfigComponentConfig(p.GetComponentConfig()),
	}
	return obj
}

// ProtoToClusterLoggingConfigComponentConfig converts a ClusterLoggingConfigComponentConfig object from its proto representation.
func ProtoToContainerazureAlphaClusterLoggingConfigComponentConfig(p *alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfig) *alpha.ClusterLoggingConfigComponentConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterLoggingConfigComponentConfig{}
	for _, r := range p.GetEnableComponents() {
		obj.EnableComponents = append(obj.EnableComponents, *ProtoToContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterMonitoringConfig converts a ClusterMonitoringConfig object from its proto representation.
func ProtoToContainerazureAlphaClusterMonitoringConfig(p *alphapb.ContainerazureAlphaClusterMonitoringConfig) *alpha.ClusterMonitoringConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterMonitoringConfig{
		ManagedPrometheusConfig: ProtoToContainerazureAlphaClusterMonitoringConfigManagedPrometheusConfig(p.GetManagedPrometheusConfig()),
	}
	return obj
}

// ProtoToClusterMonitoringConfigManagedPrometheusConfig converts a ClusterMonitoringConfigManagedPrometheusConfig object from its proto representation.
func ProtoToContainerazureAlphaClusterMonitoringConfigManagedPrometheusConfig(p *alphapb.ContainerazureAlphaClusterMonitoringConfigManagedPrometheusConfig) *alpha.ClusterMonitoringConfigManagedPrometheusConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterMonitoringConfigManagedPrometheusConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *alphapb.ContainerazureAlphaCluster) *alpha.Cluster {
	obj := &alpha.Cluster{
		Name:                        dcl.StringOrNil(p.GetName()),
		Description:                 dcl.StringOrNil(p.GetDescription()),
		AzureRegion:                 dcl.StringOrNil(p.GetAzureRegion()),
		ResourceGroupId:             dcl.StringOrNil(p.GetResourceGroupId()),
		Client:                      dcl.StringOrNil(p.GetClient()),
		AzureServicesAuthentication: ProtoToContainerazureAlphaClusterAzureServicesAuthentication(p.GetAzureServicesAuthentication()),
		Networking:                  ProtoToContainerazureAlphaClusterNetworking(p.GetNetworking()),
		ControlPlane:                ProtoToContainerazureAlphaClusterControlPlane(p.GetControlPlane()),
		Authorization:               ProtoToContainerazureAlphaClusterAuthorization(p.GetAuthorization()),
		State:                       ProtoToContainerazureAlphaClusterStateEnum(p.GetState()),
		Endpoint:                    dcl.StringOrNil(p.GetEndpoint()),
		Uid:                         dcl.StringOrNil(p.GetUid()),
		Reconciling:                 dcl.Bool(p.GetReconciling()),
		CreateTime:                  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                  dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                        dcl.StringOrNil(p.GetEtag()),
		WorkloadIdentityConfig:      ProtoToContainerazureAlphaClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                     dcl.StringOrNil(p.GetProject()),
		Location:                    dcl.StringOrNil(p.GetLocation()),
		Fleet:                       ProtoToContainerazureAlphaClusterFleet(p.GetFleet()),
		LoggingConfig:               ProtoToContainerazureAlphaClusterLoggingConfig(p.GetLoggingConfig()),
		MonitoringConfig:            ProtoToContainerazureAlphaClusterMonitoringConfig(p.GetMonitoringConfig()),
	}
	return obj
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func ContainerazureAlphaClusterStateEnumToProto(e *alpha.ClusterStateEnum) alphapb.ContainerazureAlphaClusterStateEnum {
	if e == nil {
		return alphapb.ContainerazureAlphaClusterStateEnum(0)
	}
	if v, ok := alphapb.ContainerazureAlphaClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return alphapb.ContainerazureAlphaClusterStateEnum(v)
	}
	return alphapb.ContainerazureAlphaClusterStateEnum(0)
}

// ClusterLoggingConfigComponentConfigEnableComponentsEnumToProto converts a ClusterLoggingConfigComponentConfigEnableComponentsEnum enum to its proto representation.
func ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnumToProto(e *alpha.ClusterLoggingConfigComponentConfigEnableComponentsEnum) alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum {
	if e == nil {
		return alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(0)
	}
	if v, ok := alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum_value["ClusterLoggingConfigComponentConfigEnableComponentsEnum"+string(*e)]; ok {
		return alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(v)
	}
	return alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(0)
}

// ClusterAzureServicesAuthenticationToProto converts a ClusterAzureServicesAuthentication object to its proto representation.
func ContainerazureAlphaClusterAzureServicesAuthenticationToProto(o *alpha.ClusterAzureServicesAuthentication) *alphapb.ContainerazureAlphaClusterAzureServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterAzureServicesAuthentication{}
	p.SetTenantId(dcl.ValueOrEmptyString(o.TenantId))
	p.SetApplicationId(dcl.ValueOrEmptyString(o.ApplicationId))
	return p
}

// ClusterNetworkingToProto converts a ClusterNetworking object to its proto representation.
func ContainerazureAlphaClusterNetworkingToProto(o *alpha.ClusterNetworking) *alphapb.ContainerazureAlphaClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterNetworking{}
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
func ContainerazureAlphaClusterControlPlaneToProto(o *alpha.ClusterControlPlane) *alphapb.ContainerazureAlphaClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlane{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetSubnetId(dcl.ValueOrEmptyString(o.SubnetId))
	p.SetVmSize(dcl.ValueOrEmptyString(o.VmSize))
	p.SetSshConfig(ContainerazureAlphaClusterControlPlaneSshConfigToProto(o.SshConfig))
	p.SetRootVolume(ContainerazureAlphaClusterControlPlaneRootVolumeToProto(o.RootVolume))
	p.SetMainVolume(ContainerazureAlphaClusterControlPlaneMainVolumeToProto(o.MainVolume))
	p.SetDatabaseEncryption(ContainerazureAlphaClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption))
	p.SetProxyConfig(ContainerazureAlphaClusterControlPlaneProxyConfigToProto(o.ProxyConfig))
	mTags := make(map[string]string, len(o.Tags))
	for k, r := range o.Tags {
		mTags[k] = r
	}
	p.SetTags(mTags)
	sReplicaPlacements := make([]*alphapb.ContainerazureAlphaClusterControlPlaneReplicaPlacements, len(o.ReplicaPlacements))
	for i, r := range o.ReplicaPlacements {
		sReplicaPlacements[i] = ContainerazureAlphaClusterControlPlaneReplicaPlacementsToProto(&r)
	}
	p.SetReplicaPlacements(sReplicaPlacements)
	return p
}

// ClusterControlPlaneSshConfigToProto converts a ClusterControlPlaneSshConfig object to its proto representation.
func ContainerazureAlphaClusterControlPlaneSshConfigToProto(o *alpha.ClusterControlPlaneSshConfig) *alphapb.ContainerazureAlphaClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneSshConfig{}
	p.SetAuthorizedKey(dcl.ValueOrEmptyString(o.AuthorizedKey))
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume object to its proto representation.
func ContainerazureAlphaClusterControlPlaneRootVolumeToProto(o *alpha.ClusterControlPlaneRootVolume) *alphapb.ContainerazureAlphaClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume object to its proto representation.
func ContainerazureAlphaClusterControlPlaneMainVolumeToProto(o *alpha.ClusterControlPlaneMainVolume) *alphapb.ContainerazureAlphaClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneMainVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption object to its proto representation.
func ContainerazureAlphaClusterControlPlaneDatabaseEncryptionToProto(o *alpha.ClusterControlPlaneDatabaseEncryption) *alphapb.ContainerazureAlphaClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneDatabaseEncryption{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// ClusterControlPlaneProxyConfigToProto converts a ClusterControlPlaneProxyConfig object to its proto representation.
func ContainerazureAlphaClusterControlPlaneProxyConfigToProto(o *alpha.ClusterControlPlaneProxyConfig) *alphapb.ContainerazureAlphaClusterControlPlaneProxyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneProxyConfig{}
	p.SetResourceGroupId(dcl.ValueOrEmptyString(o.ResourceGroupId))
	p.SetSecretId(dcl.ValueOrEmptyString(o.SecretId))
	return p
}

// ClusterControlPlaneReplicaPlacementsToProto converts a ClusterControlPlaneReplicaPlacements object to its proto representation.
func ContainerazureAlphaClusterControlPlaneReplicaPlacementsToProto(o *alpha.ClusterControlPlaneReplicaPlacements) *alphapb.ContainerazureAlphaClusterControlPlaneReplicaPlacements {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterControlPlaneReplicaPlacements{}
	p.SetSubnetId(dcl.ValueOrEmptyString(o.SubnetId))
	p.SetAzureAvailabilityZone(dcl.ValueOrEmptyString(o.AzureAvailabilityZone))
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization object to its proto representation.
func ContainerazureAlphaClusterAuthorizationToProto(o *alpha.ClusterAuthorization) *alphapb.ContainerazureAlphaClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterAuthorization{}
	sAdminUsers := make([]*alphapb.ContainerazureAlphaClusterAuthorizationAdminUsers, len(o.AdminUsers))
	for i, r := range o.AdminUsers {
		sAdminUsers[i] = ContainerazureAlphaClusterAuthorizationAdminUsersToProto(&r)
	}
	p.SetAdminUsers(sAdminUsers)
	sAdminGroups := make([]*alphapb.ContainerazureAlphaClusterAuthorizationAdminGroups, len(o.AdminGroups))
	for i, r := range o.AdminGroups {
		sAdminGroups[i] = ContainerazureAlphaClusterAuthorizationAdminGroupsToProto(&r)
	}
	p.SetAdminGroups(sAdminGroups)
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers object to its proto representation.
func ContainerazureAlphaClusterAuthorizationAdminUsersToProto(o *alpha.ClusterAuthorizationAdminUsers) *alphapb.ContainerazureAlphaClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterAuthorizationAdminUsers{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ClusterAuthorizationAdminGroupsToProto converts a ClusterAuthorizationAdminGroups object to its proto representation.
func ContainerazureAlphaClusterAuthorizationAdminGroupsToProto(o *alpha.ClusterAuthorizationAdminGroups) *alphapb.ContainerazureAlphaClusterAuthorizationAdminGroups {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterAuthorizationAdminGroups{}
	p.SetGroup(dcl.ValueOrEmptyString(o.Group))
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig object to its proto representation.
func ContainerazureAlphaClusterWorkloadIdentityConfigToProto(o *alpha.ClusterWorkloadIdentityConfig) *alphapb.ContainerazureAlphaClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterWorkloadIdentityConfig{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetWorkloadPool(dcl.ValueOrEmptyString(o.WorkloadPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// ClusterFleetToProto converts a ClusterFleet object to its proto representation.
func ContainerazureAlphaClusterFleetToProto(o *alpha.ClusterFleet) *alphapb.ContainerazureAlphaClusterFleet {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterFleet{}
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// ClusterLoggingConfigToProto converts a ClusterLoggingConfig object to its proto representation.
func ContainerazureAlphaClusterLoggingConfigToProto(o *alpha.ClusterLoggingConfig) *alphapb.ContainerazureAlphaClusterLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterLoggingConfig{}
	p.SetComponentConfig(ContainerazureAlphaClusterLoggingConfigComponentConfigToProto(o.ComponentConfig))
	return p
}

// ClusterLoggingConfigComponentConfigToProto converts a ClusterLoggingConfigComponentConfig object to its proto representation.
func ContainerazureAlphaClusterLoggingConfigComponentConfigToProto(o *alpha.ClusterLoggingConfigComponentConfig) *alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfig{}
	sEnableComponents := make([]alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum, len(o.EnableComponents))
	for i, r := range o.EnableComponents {
		sEnableComponents[i] = alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum(alphapb.ContainerazureAlphaClusterLoggingConfigComponentConfigEnableComponentsEnum_value[string(r)])
	}
	p.SetEnableComponents(sEnableComponents)
	return p
}

// ClusterMonitoringConfigToProto converts a ClusterMonitoringConfig object to its proto representation.
func ContainerazureAlphaClusterMonitoringConfigToProto(o *alpha.ClusterMonitoringConfig) *alphapb.ContainerazureAlphaClusterMonitoringConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterMonitoringConfig{}
	p.SetManagedPrometheusConfig(ContainerazureAlphaClusterMonitoringConfigManagedPrometheusConfigToProto(o.ManagedPrometheusConfig))
	return p
}

// ClusterMonitoringConfigManagedPrometheusConfigToProto converts a ClusterMonitoringConfigManagedPrometheusConfig object to its proto representation.
func ContainerazureAlphaClusterMonitoringConfigManagedPrometheusConfigToProto(o *alpha.ClusterMonitoringConfigManagedPrometheusConfig) *alphapb.ContainerazureAlphaClusterMonitoringConfigManagedPrometheusConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaClusterMonitoringConfigManagedPrometheusConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *alpha.Cluster) *alphapb.ContainerazureAlphaCluster {
	p := &alphapb.ContainerazureAlphaCluster{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetAzureRegion(dcl.ValueOrEmptyString(resource.AzureRegion))
	p.SetResourceGroupId(dcl.ValueOrEmptyString(resource.ResourceGroupId))
	p.SetClient(dcl.ValueOrEmptyString(resource.Client))
	p.SetAzureServicesAuthentication(ContainerazureAlphaClusterAzureServicesAuthenticationToProto(resource.AzureServicesAuthentication))
	p.SetNetworking(ContainerazureAlphaClusterNetworkingToProto(resource.Networking))
	p.SetControlPlane(ContainerazureAlphaClusterControlPlaneToProto(resource.ControlPlane))
	p.SetAuthorization(ContainerazureAlphaClusterAuthorizationToProto(resource.Authorization))
	p.SetState(ContainerazureAlphaClusterStateEnumToProto(resource.State))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkloadIdentityConfig(ContainerazureAlphaClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFleet(ContainerazureAlphaClusterFleetToProto(resource.Fleet))
	p.SetLoggingConfig(ContainerazureAlphaClusterLoggingConfigToProto(resource.LoggingConfig))
	p.SetMonitoringConfig(ContainerazureAlphaClusterMonitoringConfigToProto(resource.MonitoringConfig))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerazureAlphaClusterRequest) (*alphapb.ContainerazureAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyContainerazureAlphaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerazureAlphaCluster(ctx context.Context, request *alphapb.ApplyContainerazureAlphaClusterRequest) (*alphapb.ContainerazureAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerazureAlphaCluster(ctx context.Context, request *alphapb.DeleteContainerazureAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerazureAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerazureAlphaCluster(ctx context.Context, request *alphapb.ListContainerazureAlphaClusterRequest) (*alphapb.ListContainerazureAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerazureAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListContainerazureAlphaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
