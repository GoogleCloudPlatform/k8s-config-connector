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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkemulticloud/alpha/gkemulticloud_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/alpha"
)

// Server implements the gRPC interface for AwsCluster.
type AwsClusterServer struct{}

// ProtoToAwsClusterControlPlaneRootVolumeVolumeTypeEnum converts a AwsClusterControlPlaneRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(e alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum) *alpha.AwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.AwsClusterControlPlaneRootVolumeVolumeTypeEnum(n[len("GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterControlPlaneMainVolumeVolumeTypeEnum converts a AwsClusterControlPlaneMainVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(e alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum) *alpha.AwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.AwsClusterControlPlaneMainVolumeVolumeTypeEnum(n[len("GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterStateEnum converts a AwsClusterStateEnum enum from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterStateEnum(e alphapb.GkemulticloudAlphaAwsClusterStateEnum) *alpha.AwsClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkemulticloudAlphaAwsClusterStateEnum_name[int32(e)]; ok {
		e := alpha.AwsClusterStateEnum(n[len("GkemulticloudAlphaAwsClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterNetworking converts a AwsClusterNetworking resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterNetworking(p *alphapb.GkemulticloudAlphaAwsClusterNetworking) *alpha.AwsClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterNetworking{
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
func ProtoToGkemulticloudAlphaAwsClusterControlPlane(p *alphapb.GkemulticloudAlphaAwsClusterControlPlane) *alpha.AwsClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlane{
		Version:                   dcl.StringOrNil(p.Version),
		InstanceType:              dcl.StringOrNil(p.InstanceType),
		SshConfig:                 ProtoToGkemulticloudAlphaAwsClusterControlPlaneSshConfig(p.GetSshConfig()),
		IamInstanceProfile:        dcl.StringOrNil(p.IamInstanceProfile),
		RootVolume:                ProtoToGkemulticloudAlphaAwsClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToGkemulticloudAlphaAwsClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToGkemulticloudAlphaAwsClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToGkemulticloudAlphaAwsClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
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
func ProtoToGkemulticloudAlphaAwsClusterControlPlaneSshConfig(p *alphapb.GkemulticloudAlphaAwsClusterControlPlaneSshConfig) *alpha.AwsClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneRootVolume converts a AwsClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterControlPlaneRootVolume(p *alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolume) *alpha.AwsClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneMainVolume converts a AwsClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterControlPlaneMainVolume(p *alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolume) *alpha.AwsClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneDatabaseEncryption converts a AwsClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterControlPlaneDatabaseEncryption(p *alphapb.GkemulticloudAlphaAwsClusterControlPlaneDatabaseEncryption) *alpha.AwsClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneAwsServicesAuthentication converts a AwsClusterControlPlaneAwsServicesAuthentication resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterControlPlaneAwsServicesAuthentication(p *alphapb.GkemulticloudAlphaAwsClusterControlPlaneAwsServicesAuthentication) *alpha.AwsClusterControlPlaneAwsServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.StringOrNil(p.RoleArn),
		RoleSessionName: dcl.StringOrNil(p.RoleSessionName),
	}
	return obj
}

// ProtoToAwsClusterAuthorization converts a AwsClusterAuthorization resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterAuthorization(p *alphapb.GkemulticloudAlphaAwsClusterAuthorization) *alpha.AwsClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToGkemulticloudAlphaAwsClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToAwsClusterAuthorizationAdminUsers converts a AwsClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterAuthorizationAdminUsers(p *alphapb.GkemulticloudAlphaAwsClusterAuthorizationAdminUsers) *alpha.AwsClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToAwsClusterWorkloadIdentityConfig converts a AwsClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToGkemulticloudAlphaAwsClusterWorkloadIdentityConfig(p *alphapb.GkemulticloudAlphaAwsClusterWorkloadIdentityConfig) *alpha.AwsClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.IssuerUri),
		WorkloadPool:     dcl.StringOrNil(p.WorkloadPool),
		IdentityProvider: dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToAwsCluster converts a AwsCluster resource from its proto representation.
func ProtoToAwsCluster(p *alphapb.GkemulticloudAlphaAwsCluster) *alpha.AwsCluster {
	obj := &alpha.AwsCluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		Networking:             ProtoToGkemulticloudAlphaAwsClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.AwsRegion),
		ControlPlane:           ProtoToGkemulticloudAlphaAwsClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToGkemulticloudAlphaAwsClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToGkemulticloudAlphaAwsClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToGkemulticloudAlphaAwsClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// AwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto converts a AwsClusterControlPlaneRootVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(e *alpha.AwsClusterControlPlaneRootVolumeVolumeTypeEnum) alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum_value["AwsClusterControlPlaneRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(v)
	}
	return alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
}

// AwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto converts a AwsClusterControlPlaneMainVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(e *alpha.AwsClusterControlPlaneMainVolumeVolumeTypeEnum) alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum_value["AwsClusterControlPlaneMainVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(v)
	}
	return alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
}

// AwsClusterStateEnumToProto converts a AwsClusterStateEnum enum to its proto representation.
func GkemulticloudAlphaAwsClusterStateEnumToProto(e *alpha.AwsClusterStateEnum) alphapb.GkemulticloudAlphaAwsClusterStateEnum {
	if e == nil {
		return alphapb.GkemulticloudAlphaAwsClusterStateEnum(0)
	}
	if v, ok := alphapb.GkemulticloudAlphaAwsClusterStateEnum_value["AwsClusterStateEnum"+string(*e)]; ok {
		return alphapb.GkemulticloudAlphaAwsClusterStateEnum(v)
	}
	return alphapb.GkemulticloudAlphaAwsClusterStateEnum(0)
}

// AwsClusterNetworkingToProto converts a AwsClusterNetworking resource to its proto representation.
func GkemulticloudAlphaAwsClusterNetworkingToProto(o *alpha.AwsClusterNetworking) *alphapb.GkemulticloudAlphaAwsClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterNetworking{
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
func GkemulticloudAlphaAwsClusterControlPlaneToProto(o *alpha.AwsClusterControlPlane) *alphapb.GkemulticloudAlphaAwsClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterControlPlane{
		Version:                   dcl.ValueOrEmptyString(o.Version),
		InstanceType:              dcl.ValueOrEmptyString(o.InstanceType),
		SshConfig:                 GkemulticloudAlphaAwsClusterControlPlaneSshConfigToProto(o.SshConfig),
		IamInstanceProfile:        dcl.ValueOrEmptyString(o.IamInstanceProfile),
		RootVolume:                GkemulticloudAlphaAwsClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:                GkemulticloudAlphaAwsClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption:        GkemulticloudAlphaAwsClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
		AwsServicesAuthentication: GkemulticloudAlphaAwsClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication),
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
func GkemulticloudAlphaAwsClusterControlPlaneSshConfigToProto(o *alpha.AwsClusterControlPlaneSshConfig) *alphapb.GkemulticloudAlphaAwsClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// AwsClusterControlPlaneRootVolumeToProto converts a AwsClusterControlPlaneRootVolume resource to its proto representation.
func GkemulticloudAlphaAwsClusterControlPlaneRootVolumeToProto(o *alpha.AwsClusterControlPlaneRootVolume) *alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterControlPlaneRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneMainVolumeToProto converts a AwsClusterControlPlaneMainVolume resource to its proto representation.
func GkemulticloudAlphaAwsClusterControlPlaneMainVolumeToProto(o *alpha.AwsClusterControlPlaneMainVolume) *alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterControlPlaneMainVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneDatabaseEncryptionToProto converts a AwsClusterControlPlaneDatabaseEncryption resource to its proto representation.
func GkemulticloudAlphaAwsClusterControlPlaneDatabaseEncryptionToProto(o *alpha.AwsClusterControlPlaneDatabaseEncryption) *alphapb.GkemulticloudAlphaAwsClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneAwsServicesAuthenticationToProto converts a AwsClusterControlPlaneAwsServicesAuthentication resource to its proto representation.
func GkemulticloudAlphaAwsClusterControlPlaneAwsServicesAuthenticationToProto(o *alpha.AwsClusterControlPlaneAwsServicesAuthentication) *alphapb.GkemulticloudAlphaAwsClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.ValueOrEmptyString(o.RoleArn),
		RoleSessionName: dcl.ValueOrEmptyString(o.RoleSessionName),
	}
	return p
}

// AwsClusterAuthorizationToProto converts a AwsClusterAuthorization resource to its proto representation.
func GkemulticloudAlphaAwsClusterAuthorizationToProto(o *alpha.AwsClusterAuthorization) *alphapb.GkemulticloudAlphaAwsClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, GkemulticloudAlphaAwsClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// AwsClusterAuthorizationAdminUsersToProto converts a AwsClusterAuthorizationAdminUsers resource to its proto representation.
func GkemulticloudAlphaAwsClusterAuthorizationAdminUsersToProto(o *alpha.AwsClusterAuthorizationAdminUsers) *alphapb.GkemulticloudAlphaAwsClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// AwsClusterWorkloadIdentityConfigToProto converts a AwsClusterWorkloadIdentityConfig resource to its proto representation.
func GkemulticloudAlphaAwsClusterWorkloadIdentityConfigToProto(o *alpha.AwsClusterWorkloadIdentityConfig) *alphapb.GkemulticloudAlphaAwsClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAwsClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// AwsClusterToProto converts a AwsCluster resource to its proto representation.
func AwsClusterToProto(resource *alpha.AwsCluster) *alphapb.GkemulticloudAlphaAwsCluster {
	p := &alphapb.GkemulticloudAlphaAwsCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		Networking:             GkemulticloudAlphaAwsClusterNetworkingToProto(resource.Networking),
		AwsRegion:              dcl.ValueOrEmptyString(resource.AwsRegion),
		ControlPlane:           GkemulticloudAlphaAwsClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          GkemulticloudAlphaAwsClusterAuthorizationToProto(resource.Authorization),
		State:                  GkemulticloudAlphaAwsClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: GkemulticloudAlphaAwsClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Apply() method.
func (s *AwsClusterServer) applyAwsCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkemulticloudAlphaAwsClusterRequest) (*alphapb.GkemulticloudAlphaAwsCluster, error) {
	p := ProtoToAwsCluster(request.GetResource())
	res, err := c.ApplyAwsCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AwsClusterToProto(res)
	return r, nil
}

// ApplyAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Apply() method.
func (s *AwsClusterServer) ApplyGkemulticloudAlphaAwsCluster(ctx context.Context, request *alphapb.ApplyGkemulticloudAlphaAwsClusterRequest) (*alphapb.GkemulticloudAlphaAwsCluster, error) {
	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAwsCluster(ctx, cl, request)
}

// DeleteAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Delete() method.
func (s *AwsClusterServer) DeleteGkemulticloudAlphaAwsCluster(ctx context.Context, request *alphapb.DeleteGkemulticloudAlphaAwsClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAwsCluster(ctx, ProtoToAwsCluster(request.GetResource()))

}

// ListGkemulticloudAlphaAwsCluster handles the gRPC request by passing it to the underlying AwsClusterList() method.
func (s *AwsClusterServer) ListGkemulticloudAlphaAwsCluster(ctx context.Context, request *alphapb.ListGkemulticloudAlphaAwsClusterRequest) (*alphapb.ListGkemulticloudAlphaAwsClusterResponse, error) {
	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAwsCluster(ctx, ProtoToAwsCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkemulticloudAlphaAwsCluster
	for _, r := range resources.Items {
		rp := AwsClusterToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListGkemulticloudAlphaAwsClusterResponse{Items: protos}, nil
}

func createConfigAwsCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
