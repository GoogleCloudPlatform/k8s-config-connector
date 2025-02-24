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

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterControlPlaneRootVolumeVolumeTypeEnum converts a ClusterControlPlaneRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum(e containerawspb.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum) *containeraws.ClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerawspb.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := containeraws.ClusterControlPlaneRootVolumeVolumeTypeEnum(n[len("ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterControlPlaneMainVolumeVolumeTypeEnum converts a ClusterControlPlaneMainVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum(e containerawspb.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum) *containeraws.ClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerawspb.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := containeraws.ClusterControlPlaneMainVolumeVolumeTypeEnum(n[len("ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToContainerawsClusterStateEnum(e containerawspb.ContainerawsClusterStateEnum) *containeraws.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerawspb.ContainerawsClusterStateEnum_name[int32(e)]; ok {
		e := containeraws.ClusterStateEnum(n[len("ContainerawsClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterBinaryAuthorizationEvaluationModeEnum converts a ClusterBinaryAuthorizationEvaluationModeEnum enum from its proto representation.
func ProtoToContainerawsClusterBinaryAuthorizationEvaluationModeEnum(e containerawspb.ContainerawsClusterBinaryAuthorizationEvaluationModeEnum) *containeraws.ClusterBinaryAuthorizationEvaluationModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerawspb.ContainerawsClusterBinaryAuthorizationEvaluationModeEnum_name[int32(e)]; ok {
		e := containeraws.ClusterBinaryAuthorizationEvaluationModeEnum(n[len("ContainerawsClusterBinaryAuthorizationEvaluationModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterNetworking converts a ClusterNetworking object from its proto representation.
func ProtoToContainerawsClusterNetworking(p *containerawspb.ContainerawsClusterNetworking) *containeraws.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterNetworking{
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
func ProtoToContainerawsClusterControlPlane(p *containerawspb.ContainerawsClusterControlPlane) *containeraws.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterControlPlane{
		Version:                   dcl.StringOrNil(p.GetVersion()),
		InstanceType:              dcl.StringOrNil(p.GetInstanceType()),
		SshConfig:                 ProtoToContainerawsClusterControlPlaneSshConfig(p.GetSshConfig()),
		ConfigEncryption:          ProtoToContainerawsClusterControlPlaneConfigEncryption(p.GetConfigEncryption()),
		IamInstanceProfile:        dcl.StringOrNil(p.GetIamInstanceProfile()),
		RootVolume:                ProtoToContainerawsClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToContainerawsClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToContainerawsClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToContainerawsClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
		ProxyConfig:               ProtoToContainerawsClusterControlPlaneProxyConfig(p.GetProxyConfig()),
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
func ProtoToContainerawsClusterControlPlaneSshConfig(p *containerawspb.ContainerawsClusterControlPlaneSshConfig) *containeraws.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.GetEc2KeyPair()),
	}
	return obj
}

// ProtoToClusterControlPlaneConfigEncryption converts a ClusterControlPlaneConfigEncryption object from its proto representation.
func ProtoToContainerawsClusterControlPlaneConfigEncryption(p *containerawspb.ContainerawsClusterControlPlaneConfigEncryption) *containeraws.ClusterControlPlaneConfigEncryption {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterControlPlaneConfigEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume object from its proto representation.
func ProtoToContainerawsClusterControlPlaneRootVolume(p *containerawspb.ContainerawsClusterControlPlaneRootVolume) *containeraws.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume object from its proto representation.
func ProtoToContainerawsClusterControlPlaneMainVolume(p *containerawspb.ContainerawsClusterControlPlaneMainVolume) *containeraws.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.GetSizeGib()),
		VolumeType: ProtoToContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.GetIops()),
		Throughput: dcl.Int64OrNil(p.GetThroughput()),
		KmsKeyArn:  dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption object from its proto representation.
func ProtoToContainerawsClusterControlPlaneDatabaseEncryption(p *containerawspb.ContainerawsClusterControlPlaneDatabaseEncryption) *containeraws.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.GetKmsKeyArn()),
	}
	return obj
}

// ProtoToClusterControlPlaneAwsServicesAuthentication converts a ClusterControlPlaneAwsServicesAuthentication object from its proto representation.
func ProtoToContainerawsClusterControlPlaneAwsServicesAuthentication(p *containerawspb.ContainerawsClusterControlPlaneAwsServicesAuthentication) *containeraws.ClusterControlPlaneAwsServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.StringOrNil(p.GetRoleArn()),
		RoleSessionName: dcl.StringOrNil(p.GetRoleSessionName()),
	}
	return obj
}

// ProtoToClusterControlPlaneProxyConfig converts a ClusterControlPlaneProxyConfig object from its proto representation.
func ProtoToContainerawsClusterControlPlaneProxyConfig(p *containerawspb.ContainerawsClusterControlPlaneProxyConfig) *containeraws.ClusterControlPlaneProxyConfig {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterControlPlaneProxyConfig{
		SecretArn:     dcl.StringOrNil(p.GetSecretArn()),
		SecretVersion: dcl.StringOrNil(p.GetSecretVersion()),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization object from its proto representation.
func ProtoToContainerawsClusterAuthorization(p *containerawspb.ContainerawsClusterAuthorization) *containeraws.ClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerawsClusterAuthorizationAdminUsers(r))
	}
	for _, r := range p.GetAdminGroups() {
		obj.AdminGroups = append(obj.AdminGroups, *ProtoToContainerawsClusterAuthorizationAdminGroups(r))
	}
	return obj
}

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers object from its proto representation.
func ProtoToContainerawsClusterAuthorizationAdminUsers(p *containerawspb.ContainerawsClusterAuthorizationAdminUsers) *containeraws.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToClusterAuthorizationAdminGroups converts a ClusterAuthorizationAdminGroups object from its proto representation.
func ProtoToContainerawsClusterAuthorizationAdminGroups(p *containerawspb.ContainerawsClusterAuthorizationAdminGroups) *containeraws.ClusterAuthorizationAdminGroups {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterAuthorizationAdminGroups{
		Group: dcl.StringOrNil(p.GetGroup()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig object from its proto representation.
func ProtoToContainerawsClusterWorkloadIdentityConfig(p *containerawspb.ContainerawsClusterWorkloadIdentityConfig) *containeraws.ClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.GetIssuerUri()),
		WorkloadPool:     dcl.StringOrNil(p.GetWorkloadPool()),
		IdentityProvider: dcl.StringOrNil(p.GetIdentityProvider()),
	}
	return obj
}

// ProtoToClusterFleet converts a ClusterFleet object from its proto representation.
func ProtoToContainerawsClusterFleet(p *containerawspb.ContainerawsClusterFleet) *containeraws.ClusterFleet {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterFleet{
		Project:    dcl.StringOrNil(p.GetProject()),
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToClusterBinaryAuthorization converts a ClusterBinaryAuthorization object from its proto representation.
func ProtoToContainerawsClusterBinaryAuthorization(p *containerawspb.ContainerawsClusterBinaryAuthorization) *containeraws.ClusterBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &containeraws.ClusterBinaryAuthorization{
		EvaluationMode: ProtoToContainerawsClusterBinaryAuthorizationEvaluationModeEnum(p.GetEvaluationMode()),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *containerawspb.ContainerawsCluster) *containeraws.Cluster {
	obj := &containeraws.Cluster{
		Name:                   dcl.StringOrNil(p.GetName()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		Networking:             ProtoToContainerawsClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.GetAwsRegion()),
		ControlPlane:           ProtoToContainerawsClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToContainerawsClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToContainerawsClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.GetEndpoint()),
		Uid:                    dcl.StringOrNil(p.GetUid()),
		Reconciling:            dcl.Bool(p.GetReconciling()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.GetEtag()),
		WorkloadIdentityConfig: ProtoToContainerawsClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
		Fleet:                  ProtoToContainerawsClusterFleet(p.GetFleet()),
		BinaryAuthorization:    ProtoToContainerawsClusterBinaryAuthorization(p.GetBinaryAuthorization()),
	}
	return obj
}

// ClusterControlPlaneRootVolumeVolumeTypeEnumToProto converts a ClusterControlPlaneRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(e *containeraws.ClusterControlPlaneRootVolumeVolumeTypeEnum) containerawspb.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == nil {
		return containerawspb.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := containerawspb.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum_value["ClusterControlPlaneRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return containerawspb.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum(v)
	}
	return containerawspb.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
}

// ClusterControlPlaneMainVolumeVolumeTypeEnumToProto converts a ClusterControlPlaneMainVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(e *containeraws.ClusterControlPlaneMainVolumeVolumeTypeEnum) containerawspb.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == nil {
		return containerawspb.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
	}
	if v, ok := containerawspb.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum_value["ClusterControlPlaneMainVolumeVolumeTypeEnum"+string(*e)]; ok {
		return containerawspb.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum(v)
	}
	return containerawspb.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func ContainerawsClusterStateEnumToProto(e *containeraws.ClusterStateEnum) containerawspb.ContainerawsClusterStateEnum {
	if e == nil {
		return containerawspb.ContainerawsClusterStateEnum(0)
	}
	if v, ok := containerawspb.ContainerawsClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return containerawspb.ContainerawsClusterStateEnum(v)
	}
	return containerawspb.ContainerawsClusterStateEnum(0)
}

// ClusterBinaryAuthorizationEvaluationModeEnumToProto converts a ClusterBinaryAuthorizationEvaluationModeEnum enum to its proto representation.
func ContainerawsClusterBinaryAuthorizationEvaluationModeEnumToProto(e *containeraws.ClusterBinaryAuthorizationEvaluationModeEnum) containerawspb.ContainerawsClusterBinaryAuthorizationEvaluationModeEnum {
	if e == nil {
		return containerawspb.ContainerawsClusterBinaryAuthorizationEvaluationModeEnum(0)
	}
	if v, ok := containerawspb.ContainerawsClusterBinaryAuthorizationEvaluationModeEnum_value["ClusterBinaryAuthorizationEvaluationModeEnum"+string(*e)]; ok {
		return containerawspb.ContainerawsClusterBinaryAuthorizationEvaluationModeEnum(v)
	}
	return containerawspb.ContainerawsClusterBinaryAuthorizationEvaluationModeEnum(0)
}

// ClusterNetworkingToProto converts a ClusterNetworking object to its proto representation.
func ContainerawsClusterNetworkingToProto(o *containeraws.ClusterNetworking) *containerawspb.ContainerawsClusterNetworking {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterNetworking{}
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
func ContainerawsClusterControlPlaneToProto(o *containeraws.ClusterControlPlane) *containerawspb.ContainerawsClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterControlPlane{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetInstanceType(dcl.ValueOrEmptyString(o.InstanceType))
	p.SetSshConfig(ContainerawsClusterControlPlaneSshConfigToProto(o.SshConfig))
	p.SetConfigEncryption(ContainerawsClusterControlPlaneConfigEncryptionToProto(o.ConfigEncryption))
	p.SetIamInstanceProfile(dcl.ValueOrEmptyString(o.IamInstanceProfile))
	p.SetRootVolume(ContainerawsClusterControlPlaneRootVolumeToProto(o.RootVolume))
	p.SetMainVolume(ContainerawsClusterControlPlaneMainVolumeToProto(o.MainVolume))
	p.SetDatabaseEncryption(ContainerawsClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption))
	p.SetAwsServicesAuthentication(ContainerawsClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication))
	p.SetProxyConfig(ContainerawsClusterControlPlaneProxyConfigToProto(o.ProxyConfig))
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
func ContainerawsClusterControlPlaneSshConfigToProto(o *containeraws.ClusterControlPlaneSshConfig) *containerawspb.ContainerawsClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterControlPlaneSshConfig{}
	p.SetEc2KeyPair(dcl.ValueOrEmptyString(o.Ec2KeyPair))
	return p
}

// ClusterControlPlaneConfigEncryptionToProto converts a ClusterControlPlaneConfigEncryption object to its proto representation.
func ContainerawsClusterControlPlaneConfigEncryptionToProto(o *containeraws.ClusterControlPlaneConfigEncryption) *containerawspb.ContainerawsClusterControlPlaneConfigEncryption {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterControlPlaneConfigEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume object to its proto representation.
func ContainerawsClusterControlPlaneRootVolumeToProto(o *containeraws.ClusterControlPlaneRootVolume) *containerawspb.ContainerawsClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterControlPlaneRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume object to its proto representation.
func ContainerawsClusterControlPlaneMainVolumeToProto(o *containeraws.ClusterControlPlaneMainVolume) *containerawspb.ContainerawsClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterControlPlaneMainVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	p.SetVolumeType(ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType))
	p.SetIops(dcl.ValueOrEmptyInt64(o.Iops))
	p.SetThroughput(dcl.ValueOrEmptyInt64(o.Throughput))
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption object to its proto representation.
func ContainerawsClusterControlPlaneDatabaseEncryptionToProto(o *containeraws.ClusterControlPlaneDatabaseEncryption) *containerawspb.ContainerawsClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterControlPlaneDatabaseEncryption{}
	p.SetKmsKeyArn(dcl.ValueOrEmptyString(o.KmsKeyArn))
	return p
}

// ClusterControlPlaneAwsServicesAuthenticationToProto converts a ClusterControlPlaneAwsServicesAuthentication object to its proto representation.
func ContainerawsClusterControlPlaneAwsServicesAuthenticationToProto(o *containeraws.ClusterControlPlaneAwsServicesAuthentication) *containerawspb.ContainerawsClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterControlPlaneAwsServicesAuthentication{}
	p.SetRoleArn(dcl.ValueOrEmptyString(o.RoleArn))
	p.SetRoleSessionName(dcl.ValueOrEmptyString(o.RoleSessionName))
	return p
}

// ClusterControlPlaneProxyConfigToProto converts a ClusterControlPlaneProxyConfig object to its proto representation.
func ContainerawsClusterControlPlaneProxyConfigToProto(o *containeraws.ClusterControlPlaneProxyConfig) *containerawspb.ContainerawsClusterControlPlaneProxyConfig {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterControlPlaneProxyConfig{}
	p.SetSecretArn(dcl.ValueOrEmptyString(o.SecretArn))
	p.SetSecretVersion(dcl.ValueOrEmptyString(o.SecretVersion))
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization object to its proto representation.
func ContainerawsClusterAuthorizationToProto(o *containeraws.ClusterAuthorization) *containerawspb.ContainerawsClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterAuthorization{}
	sAdminUsers := make([]*containerawspb.ContainerawsClusterAuthorizationAdminUsers, len(o.AdminUsers))
	for i, r := range o.AdminUsers {
		sAdminUsers[i] = ContainerawsClusterAuthorizationAdminUsersToProto(&r)
	}
	p.SetAdminUsers(sAdminUsers)
	sAdminGroups := make([]*containerawspb.ContainerawsClusterAuthorizationAdminGroups, len(o.AdminGroups))
	for i, r := range o.AdminGroups {
		sAdminGroups[i] = ContainerawsClusterAuthorizationAdminGroupsToProto(&r)
	}
	p.SetAdminGroups(sAdminGroups)
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers object to its proto representation.
func ContainerawsClusterAuthorizationAdminUsersToProto(o *containeraws.ClusterAuthorizationAdminUsers) *containerawspb.ContainerawsClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterAuthorizationAdminUsers{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ClusterAuthorizationAdminGroupsToProto converts a ClusterAuthorizationAdminGroups object to its proto representation.
func ContainerawsClusterAuthorizationAdminGroupsToProto(o *containeraws.ClusterAuthorizationAdminGroups) *containerawspb.ContainerawsClusterAuthorizationAdminGroups {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterAuthorizationAdminGroups{}
	p.SetGroup(dcl.ValueOrEmptyString(o.Group))
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig object to its proto representation.
func ContainerawsClusterWorkloadIdentityConfigToProto(o *containeraws.ClusterWorkloadIdentityConfig) *containerawspb.ContainerawsClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterWorkloadIdentityConfig{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetWorkloadPool(dcl.ValueOrEmptyString(o.WorkloadPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// ClusterFleetToProto converts a ClusterFleet object to its proto representation.
func ContainerawsClusterFleetToProto(o *containeraws.ClusterFleet) *containerawspb.ContainerawsClusterFleet {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterFleet{}
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// ClusterBinaryAuthorizationToProto converts a ClusterBinaryAuthorization object to its proto representation.
func ContainerawsClusterBinaryAuthorizationToProto(o *containeraws.ClusterBinaryAuthorization) *containerawspb.ContainerawsClusterBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &containerawspb.ContainerawsClusterBinaryAuthorization{}
	p.SetEvaluationMode(ContainerawsClusterBinaryAuthorizationEvaluationModeEnumToProto(o.EvaluationMode))
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *containeraws.Cluster) *containerawspb.ContainerawsCluster {
	p := &containerawspb.ContainerawsCluster{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetNetworking(ContainerawsClusterNetworkingToProto(resource.Networking))
	p.SetAwsRegion(dcl.ValueOrEmptyString(resource.AwsRegion))
	p.SetControlPlane(ContainerawsClusterControlPlaneToProto(resource.ControlPlane))
	p.SetAuthorization(ContainerawsClusterAuthorizationToProto(resource.Authorization))
	p.SetState(ContainerawsClusterStateEnumToProto(resource.State))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkloadIdentityConfig(ContainerawsClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFleet(ContainerawsClusterFleetToProto(resource.Fleet))
	p.SetBinaryAuthorization(ContainerawsClusterBinaryAuthorizationToProto(resource.BinaryAuthorization))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *containeraws.Client, request *containerawspb.ApplyContainerawsClusterRequest) (*containerawspb.ContainerawsCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyContainerawsCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerawsCluster(ctx context.Context, request *containerawspb.ApplyContainerawsClusterRequest) (*containerawspb.ContainerawsCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerawsCluster(ctx context.Context, request *containerawspb.DeleteContainerawsClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerawsCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerawsCluster(ctx context.Context, request *containerawspb.ListContainerawsClusterRequest) (*containerawspb.ListContainerawsClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*containerawspb.ContainerawsCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &containerawspb.ListContainerawsClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*containeraws.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return containeraws.NewClient(conf), nil
}
