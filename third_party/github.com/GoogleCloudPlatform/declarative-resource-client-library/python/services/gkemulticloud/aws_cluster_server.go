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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	gkemulticloudpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkemulticloud/gkemulticloud_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud"
)

// Server implements the gRPC interface for AwsCluster.
type AwsClusterServer struct{}

// ProtoToAwsClusterControlPlaneRootVolumeVolumeTypeEnum converts a AwsClusterControlPlaneRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum(e gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum) *gkemulticloud.AwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := gkemulticloud.AwsClusterControlPlaneRootVolumeVolumeTypeEnum(n[len("GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterControlPlaneMainVolumeVolumeTypeEnum converts a AwsClusterControlPlaneMainVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum(e gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum) *gkemulticloud.AwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := gkemulticloud.AwsClusterControlPlaneMainVolumeVolumeTypeEnum(n[len("GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterStateEnum converts a AwsClusterStateEnum enum from its proto representation.
func ProtoToGkemulticloudAwsClusterStateEnum(e gkemulticloudpb.GkemulticloudAwsClusterStateEnum) *gkemulticloud.AwsClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkemulticloudpb.GkemulticloudAwsClusterStateEnum_name[int32(e)]; ok {
		e := gkemulticloud.AwsClusterStateEnum(n[len("GkemulticloudAwsClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterNetworking converts a AwsClusterNetworking resource from its proto representation.
func ProtoToGkemulticloudAwsClusterNetworking(p *gkemulticloudpb.GkemulticloudAwsClusterNetworking) *gkemulticloud.AwsClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterNetworking{
		VPCId: dcl.StringOrNil(p.VpcId),
	}
	for _, r := range p.GetPodAddressCidrBlocks() {
		obj.PodAddressCidrBlocks = append(obj.PodAddressCidrBlocks, r)
	}
	for _, r := range p.GetServiceAddressCidrBlocks() {
		obj.ServiceAddressCidrBlocks = append(obj.ServiceAddressCidrBlocks, r)
	}
	for _, r := range p.GetServiceLoadBalancerSubnetIds() {
		obj.ServiceLoadBalancerSubnetIds = append(obj.ServiceLoadBalancerSubnetIds, r)
	}
	return obj
}

// ProtoToAwsClusterControlPlane converts a AwsClusterControlPlane resource from its proto representation.
func ProtoToGkemulticloudAwsClusterControlPlane(p *gkemulticloudpb.GkemulticloudAwsClusterControlPlane) *gkemulticloud.AwsClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterControlPlane{
		Version:                   dcl.StringOrNil(p.Version),
		InstanceType:              dcl.StringOrNil(p.InstanceType),
		SshConfig:                 ProtoToGkemulticloudAwsClusterControlPlaneSshConfig(p.GetSshConfig()),
		IamInstanceProfile:        dcl.StringOrNil(p.IamInstanceProfile),
		RootVolume:                ProtoToGkemulticloudAwsClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToGkemulticloudAwsClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToGkemulticloudAwsClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToGkemulticloudAwsClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
	}
	for _, r := range p.GetSubnetIds() {
		obj.SubnetIds = append(obj.SubnetIds, r)
	}
	for _, r := range p.GetSecurityGroupIds() {
		obj.SecurityGroupIds = append(obj.SecurityGroupIds, r)
	}
	return obj
}

// ProtoToAwsClusterControlPlaneSshConfig converts a AwsClusterControlPlaneSshConfig resource from its proto representation.
func ProtoToGkemulticloudAwsClusterControlPlaneSshConfig(p *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneSshConfig) *gkemulticloud.AwsClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneRootVolume converts a AwsClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToGkemulticloudAwsClusterControlPlaneRootVolume(p *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolume) *gkemulticloud.AwsClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneMainVolume converts a AwsClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToGkemulticloudAwsClusterControlPlaneMainVolume(p *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolume) *gkemulticloud.AwsClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneDatabaseEncryption converts a AwsClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToGkemulticloudAwsClusterControlPlaneDatabaseEncryption(p *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneDatabaseEncryption) *gkemulticloud.AwsClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneAwsServicesAuthentication converts a AwsClusterControlPlaneAwsServicesAuthentication resource from its proto representation.
func ProtoToGkemulticloudAwsClusterControlPlaneAwsServicesAuthentication(p *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneAwsServicesAuthentication) *gkemulticloud.AwsClusterControlPlaneAwsServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.StringOrNil(p.RoleArn),
		RoleSessionName: dcl.StringOrNil(p.RoleSessionName),
	}
	return obj
}

// ProtoToAwsClusterAuthorization converts a AwsClusterAuthorization resource from its proto representation.
func ProtoToGkemulticloudAwsClusterAuthorization(p *gkemulticloudpb.GkemulticloudAwsClusterAuthorization) *gkemulticloud.AwsClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToGkemulticloudAwsClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToAwsClusterAuthorizationAdminUsers converts a AwsClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToGkemulticloudAwsClusterAuthorizationAdminUsers(p *gkemulticloudpb.GkemulticloudAwsClusterAuthorizationAdminUsers) *gkemulticloud.AwsClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToAwsClusterWorkloadIdentityConfig converts a AwsClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToGkemulticloudAwsClusterWorkloadIdentityConfig(p *gkemulticloudpb.GkemulticloudAwsClusterWorkloadIdentityConfig) *gkemulticloud.AwsClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AwsClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.IssuerUri),
		WorkloadPool:     dcl.StringOrNil(p.WorkloadPool),
		IdentityProvider: dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToAwsCluster converts a AwsCluster resource from its proto representation.
func ProtoToAwsCluster(p *gkemulticloudpb.GkemulticloudAwsCluster) *gkemulticloud.AwsCluster {
	obj := &gkemulticloud.AwsCluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		Networking:             ProtoToGkemulticloudAwsClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.AwsRegion),
		ControlPlane:           ProtoToGkemulticloudAwsClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToGkemulticloudAwsClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToGkemulticloudAwsClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToGkemulticloudAwsClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// AwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto converts a AwsClusterControlPlaneRootVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(e *gkemulticloud.AwsClusterControlPlaneRootVolumeVolumeTypeEnum) gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == nil {
		return gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum_value["AwsClusterControlPlaneRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum(v)
	}
	return gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
}

// AwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto converts a AwsClusterControlPlaneMainVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(e *gkemulticloud.AwsClusterControlPlaneMainVolumeVolumeTypeEnum) gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == nil {
		return gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
	}
	if v, ok := gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum_value["AwsClusterControlPlaneMainVolumeVolumeTypeEnum"+string(*e)]; ok {
		return gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum(v)
	}
	return gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
}

// AwsClusterStateEnumToProto converts a AwsClusterStateEnum enum to its proto representation.
func GkemulticloudAwsClusterStateEnumToProto(e *gkemulticloud.AwsClusterStateEnum) gkemulticloudpb.GkemulticloudAwsClusterStateEnum {
	if e == nil {
		return gkemulticloudpb.GkemulticloudAwsClusterStateEnum(0)
	}
	if v, ok := gkemulticloudpb.GkemulticloudAwsClusterStateEnum_value["AwsClusterStateEnum"+string(*e)]; ok {
		return gkemulticloudpb.GkemulticloudAwsClusterStateEnum(v)
	}
	return gkemulticloudpb.GkemulticloudAwsClusterStateEnum(0)
}

// AwsClusterNetworkingToProto converts a AwsClusterNetworking resource to its proto representation.
func GkemulticloudAwsClusterNetworkingToProto(o *gkemulticloud.AwsClusterNetworking) *gkemulticloudpb.GkemulticloudAwsClusterNetworking {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterNetworking{
		VpcId: dcl.ValueOrEmptyString(o.VPCId),
	}
	for _, r := range o.PodAddressCidrBlocks {
		p.PodAddressCidrBlocks = append(p.PodAddressCidrBlocks, r)
	}
	for _, r := range o.ServiceAddressCidrBlocks {
		p.ServiceAddressCidrBlocks = append(p.ServiceAddressCidrBlocks, r)
	}
	for _, r := range o.ServiceLoadBalancerSubnetIds {
		p.ServiceLoadBalancerSubnetIds = append(p.ServiceLoadBalancerSubnetIds, r)
	}
	return p
}

// AwsClusterControlPlaneToProto converts a AwsClusterControlPlane resource to its proto representation.
func GkemulticloudAwsClusterControlPlaneToProto(o *gkemulticloud.AwsClusterControlPlane) *gkemulticloudpb.GkemulticloudAwsClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterControlPlane{
		Version:                   dcl.ValueOrEmptyString(o.Version),
		InstanceType:              dcl.ValueOrEmptyString(o.InstanceType),
		SshConfig:                 GkemulticloudAwsClusterControlPlaneSshConfigToProto(o.SshConfig),
		IamInstanceProfile:        dcl.ValueOrEmptyString(o.IamInstanceProfile),
		RootVolume:                GkemulticloudAwsClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:                GkemulticloudAwsClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption:        GkemulticloudAwsClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
		AwsServicesAuthentication: GkemulticloudAwsClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication),
	}
	for _, r := range o.SubnetIds {
		p.SubnetIds = append(p.SubnetIds, r)
	}
	for _, r := range o.SecurityGroupIds {
		p.SecurityGroupIds = append(p.SecurityGroupIds, r)
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AwsClusterControlPlaneSshConfigToProto converts a AwsClusterControlPlaneSshConfig resource to its proto representation.
func GkemulticloudAwsClusterControlPlaneSshConfigToProto(o *gkemulticloud.AwsClusterControlPlaneSshConfig) *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// AwsClusterControlPlaneRootVolumeToProto converts a AwsClusterControlPlaneRootVolume resource to its proto representation.
func GkemulticloudAwsClusterControlPlaneRootVolumeToProto(o *gkemulticloud.AwsClusterControlPlaneRootVolume) *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterControlPlaneRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudAwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneMainVolumeToProto converts a AwsClusterControlPlaneMainVolume resource to its proto representation.
func GkemulticloudAwsClusterControlPlaneMainVolumeToProto(o *gkemulticloud.AwsClusterControlPlaneMainVolume) *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterControlPlaneMainVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudAwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneDatabaseEncryptionToProto converts a AwsClusterControlPlaneDatabaseEncryption resource to its proto representation.
func GkemulticloudAwsClusterControlPlaneDatabaseEncryptionToProto(o *gkemulticloud.AwsClusterControlPlaneDatabaseEncryption) *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneAwsServicesAuthenticationToProto converts a AwsClusterControlPlaneAwsServicesAuthentication resource to its proto representation.
func GkemulticloudAwsClusterControlPlaneAwsServicesAuthenticationToProto(o *gkemulticloud.AwsClusterControlPlaneAwsServicesAuthentication) *gkemulticloudpb.GkemulticloudAwsClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.ValueOrEmptyString(o.RoleArn),
		RoleSessionName: dcl.ValueOrEmptyString(o.RoleSessionName),
	}
	return p
}

// AwsClusterAuthorizationToProto converts a AwsClusterAuthorization resource to its proto representation.
func GkemulticloudAwsClusterAuthorizationToProto(o *gkemulticloud.AwsClusterAuthorization) *gkemulticloudpb.GkemulticloudAwsClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, GkemulticloudAwsClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// AwsClusterAuthorizationAdminUsersToProto converts a AwsClusterAuthorizationAdminUsers resource to its proto representation.
func GkemulticloudAwsClusterAuthorizationAdminUsersToProto(o *gkemulticloud.AwsClusterAuthorizationAdminUsers) *gkemulticloudpb.GkemulticloudAwsClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// AwsClusterWorkloadIdentityConfigToProto converts a AwsClusterWorkloadIdentityConfig resource to its proto representation.
func GkemulticloudAwsClusterWorkloadIdentityConfigToProto(o *gkemulticloud.AwsClusterWorkloadIdentityConfig) *gkemulticloudpb.GkemulticloudAwsClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAwsClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// AwsClusterToProto converts a AwsCluster resource to its proto representation.
func AwsClusterToProto(resource *gkemulticloud.AwsCluster) *gkemulticloudpb.GkemulticloudAwsCluster {
	p := &gkemulticloudpb.GkemulticloudAwsCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		Networking:             GkemulticloudAwsClusterNetworkingToProto(resource.Networking),
		AwsRegion:              dcl.ValueOrEmptyString(resource.AwsRegion),
		ControlPlane:           GkemulticloudAwsClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          GkemulticloudAwsClusterAuthorizationToProto(resource.Authorization),
		State:                  GkemulticloudAwsClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: GkemulticloudAwsClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Apply() method.
func (s *AwsClusterServer) applyAwsCluster(ctx context.Context, c *gkemulticloud.Client, request *gkemulticloudpb.ApplyGkemulticloudAwsClusterRequest) (*gkemulticloudpb.GkemulticloudAwsCluster, error) {
	p := ProtoToAwsCluster(request.GetResource())
	res, err := c.ApplyAwsCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AwsClusterToProto(res)
	return r, nil
}

// ApplyAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Apply() method.
func (s *AwsClusterServer) ApplyGkemulticloudAwsCluster(ctx context.Context, request *gkemulticloudpb.ApplyGkemulticloudAwsClusterRequest) (*gkemulticloudpb.GkemulticloudAwsCluster, error) {
	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAwsCluster(ctx, cl, request)
}

// DeleteAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Delete() method.
func (s *AwsClusterServer) DeleteGkemulticloudAwsCluster(ctx context.Context, request *gkemulticloudpb.DeleteGkemulticloudAwsClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAwsCluster(ctx, ProtoToAwsCluster(request.GetResource()))

}

// ListGkemulticloudAwsCluster handles the gRPC request by passing it to the underlying AwsClusterList() method.
func (s *AwsClusterServer) ListGkemulticloudAwsCluster(ctx context.Context, request *gkemulticloudpb.ListGkemulticloudAwsClusterRequest) (*gkemulticloudpb.ListGkemulticloudAwsClusterResponse, error) {
	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAwsCluster(ctx, ProtoToAwsCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*gkemulticloudpb.GkemulticloudAwsCluster
	for _, r := range resources.Items {
		rp := AwsClusterToProto(r)
		protos = append(protos, rp)
	}
	return &gkemulticloudpb.ListGkemulticloudAwsClusterResponse{Items: protos}, nil
}

func createConfigAwsCluster(ctx context.Context, service_account_file string) (*gkemulticloud.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return gkemulticloud.NewClient(conf), nil
}
