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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkemulticloud/beta/gkemulticloud_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/beta"
)

// Server implements the gRPC interface for AwsCluster.
type AwsClusterServer struct{}

// ProtoToAwsClusterControlPlaneRootVolumeVolumeTypeEnum converts a AwsClusterControlPlaneRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(e betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum) *beta.AwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := beta.AwsClusterControlPlaneRootVolumeVolumeTypeEnum(n[len("GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterControlPlaneMainVolumeVolumeTypeEnum converts a AwsClusterControlPlaneMainVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(e betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum) *beta.AwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := beta.AwsClusterControlPlaneMainVolumeVolumeTypeEnum(n[len("GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterStateEnum converts a AwsClusterStateEnum enum from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterStateEnum(e betapb.GkemulticloudBetaAwsClusterStateEnum) *beta.AwsClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkemulticloudBetaAwsClusterStateEnum_name[int32(e)]; ok {
		e := beta.AwsClusterStateEnum(n[len("GkemulticloudBetaAwsClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterNetworking converts a AwsClusterNetworking resource from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterNetworking(p *betapb.GkemulticloudBetaAwsClusterNetworking) *beta.AwsClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterNetworking{
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
func ProtoToGkemulticloudBetaAwsClusterControlPlane(p *betapb.GkemulticloudBetaAwsClusterControlPlane) *beta.AwsClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterControlPlane{
		Version:                   dcl.StringOrNil(p.Version),
		InstanceType:              dcl.StringOrNil(p.InstanceType),
		SshConfig:                 ProtoToGkemulticloudBetaAwsClusterControlPlaneSshConfig(p.GetSshConfig()),
		IamInstanceProfile:        dcl.StringOrNil(p.IamInstanceProfile),
		RootVolume:                ProtoToGkemulticloudBetaAwsClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToGkemulticloudBetaAwsClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToGkemulticloudBetaAwsClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToGkemulticloudBetaAwsClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
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
func ProtoToGkemulticloudBetaAwsClusterControlPlaneSshConfig(p *betapb.GkemulticloudBetaAwsClusterControlPlaneSshConfig) *beta.AwsClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneRootVolume converts a AwsClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterControlPlaneRootVolume(p *betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolume) *beta.AwsClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneMainVolume converts a AwsClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterControlPlaneMainVolume(p *betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolume) *beta.AwsClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToGkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneDatabaseEncryption converts a AwsClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterControlPlaneDatabaseEncryption(p *betapb.GkemulticloudBetaAwsClusterControlPlaneDatabaseEncryption) *beta.AwsClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneAwsServicesAuthentication converts a AwsClusterControlPlaneAwsServicesAuthentication resource from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterControlPlaneAwsServicesAuthentication(p *betapb.GkemulticloudBetaAwsClusterControlPlaneAwsServicesAuthentication) *beta.AwsClusterControlPlaneAwsServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.StringOrNil(p.RoleArn),
		RoleSessionName: dcl.StringOrNil(p.RoleSessionName),
	}
	return obj
}

// ProtoToAwsClusterAuthorization converts a AwsClusterAuthorization resource from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterAuthorization(p *betapb.GkemulticloudBetaAwsClusterAuthorization) *beta.AwsClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToGkemulticloudBetaAwsClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToAwsClusterAuthorizationAdminUsers converts a AwsClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterAuthorizationAdminUsers(p *betapb.GkemulticloudBetaAwsClusterAuthorizationAdminUsers) *beta.AwsClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToAwsClusterWorkloadIdentityConfig converts a AwsClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToGkemulticloudBetaAwsClusterWorkloadIdentityConfig(p *betapb.GkemulticloudBetaAwsClusterWorkloadIdentityConfig) *beta.AwsClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AwsClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.IssuerUri),
		WorkloadPool:     dcl.StringOrNil(p.WorkloadPool),
		IdentityProvider: dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToAwsCluster converts a AwsCluster resource from its proto representation.
func ProtoToAwsCluster(p *betapb.GkemulticloudBetaAwsCluster) *beta.AwsCluster {
	obj := &beta.AwsCluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		Networking:             ProtoToGkemulticloudBetaAwsClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.AwsRegion),
		ControlPlane:           ProtoToGkemulticloudBetaAwsClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToGkemulticloudBetaAwsClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToGkemulticloudBetaAwsClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToGkemulticloudBetaAwsClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// AwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto converts a AwsClusterControlPlaneRootVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(e *beta.AwsClusterControlPlaneRootVolumeVolumeTypeEnum) betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == nil {
		return betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum_value["AwsClusterControlPlaneRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(v)
	}
	return betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
}

// AwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto converts a AwsClusterControlPlaneMainVolumeVolumeTypeEnum enum to its proto representation.
func GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(e *beta.AwsClusterControlPlaneMainVolumeVolumeTypeEnum) betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == nil {
		return betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
	}
	if v, ok := betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum_value["AwsClusterControlPlaneMainVolumeVolumeTypeEnum"+string(*e)]; ok {
		return betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(v)
	}
	return betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
}

// AwsClusterStateEnumToProto converts a AwsClusterStateEnum enum to its proto representation.
func GkemulticloudBetaAwsClusterStateEnumToProto(e *beta.AwsClusterStateEnum) betapb.GkemulticloudBetaAwsClusterStateEnum {
	if e == nil {
		return betapb.GkemulticloudBetaAwsClusterStateEnum(0)
	}
	if v, ok := betapb.GkemulticloudBetaAwsClusterStateEnum_value["AwsClusterStateEnum"+string(*e)]; ok {
		return betapb.GkemulticloudBetaAwsClusterStateEnum(v)
	}
	return betapb.GkemulticloudBetaAwsClusterStateEnum(0)
}

// AwsClusterNetworkingToProto converts a AwsClusterNetworking resource to its proto representation.
func GkemulticloudBetaAwsClusterNetworkingToProto(o *beta.AwsClusterNetworking) *betapb.GkemulticloudBetaAwsClusterNetworking {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterNetworking{
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
func GkemulticloudBetaAwsClusterControlPlaneToProto(o *beta.AwsClusterControlPlane) *betapb.GkemulticloudBetaAwsClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterControlPlane{
		Version:                   dcl.ValueOrEmptyString(o.Version),
		InstanceType:              dcl.ValueOrEmptyString(o.InstanceType),
		SshConfig:                 GkemulticloudBetaAwsClusterControlPlaneSshConfigToProto(o.SshConfig),
		IamInstanceProfile:        dcl.ValueOrEmptyString(o.IamInstanceProfile),
		RootVolume:                GkemulticloudBetaAwsClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:                GkemulticloudBetaAwsClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption:        GkemulticloudBetaAwsClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
		AwsServicesAuthentication: GkemulticloudBetaAwsClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication),
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
func GkemulticloudBetaAwsClusterControlPlaneSshConfigToProto(o *beta.AwsClusterControlPlaneSshConfig) *betapb.GkemulticloudBetaAwsClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// AwsClusterControlPlaneRootVolumeToProto converts a AwsClusterControlPlaneRootVolume resource to its proto representation.
func GkemulticloudBetaAwsClusterControlPlaneRootVolumeToProto(o *beta.AwsClusterControlPlaneRootVolume) *betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterControlPlaneRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudBetaAwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneMainVolumeToProto converts a AwsClusterControlPlaneMainVolume resource to its proto representation.
func GkemulticloudBetaAwsClusterControlPlaneMainVolumeToProto(o *beta.AwsClusterControlPlaneMainVolume) *betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterControlPlaneMainVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: GkemulticloudBetaAwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneDatabaseEncryptionToProto converts a AwsClusterControlPlaneDatabaseEncryption resource to its proto representation.
func GkemulticloudBetaAwsClusterControlPlaneDatabaseEncryptionToProto(o *beta.AwsClusterControlPlaneDatabaseEncryption) *betapb.GkemulticloudBetaAwsClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneAwsServicesAuthenticationToProto converts a AwsClusterControlPlaneAwsServicesAuthentication resource to its proto representation.
func GkemulticloudBetaAwsClusterControlPlaneAwsServicesAuthenticationToProto(o *beta.AwsClusterControlPlaneAwsServicesAuthentication) *betapb.GkemulticloudBetaAwsClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.ValueOrEmptyString(o.RoleArn),
		RoleSessionName: dcl.ValueOrEmptyString(o.RoleSessionName),
	}
	return p
}

// AwsClusterAuthorizationToProto converts a AwsClusterAuthorization resource to its proto representation.
func GkemulticloudBetaAwsClusterAuthorizationToProto(o *beta.AwsClusterAuthorization) *betapb.GkemulticloudBetaAwsClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, GkemulticloudBetaAwsClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// AwsClusterAuthorizationAdminUsersToProto converts a AwsClusterAuthorizationAdminUsers resource to its proto representation.
func GkemulticloudBetaAwsClusterAuthorizationAdminUsersToProto(o *beta.AwsClusterAuthorizationAdminUsers) *betapb.GkemulticloudBetaAwsClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// AwsClusterWorkloadIdentityConfigToProto converts a AwsClusterWorkloadIdentityConfig resource to its proto representation.
func GkemulticloudBetaAwsClusterWorkloadIdentityConfigToProto(o *beta.AwsClusterWorkloadIdentityConfig) *betapb.GkemulticloudBetaAwsClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAwsClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// AwsClusterToProto converts a AwsCluster resource to its proto representation.
func AwsClusterToProto(resource *beta.AwsCluster) *betapb.GkemulticloudBetaAwsCluster {
	p := &betapb.GkemulticloudBetaAwsCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		Networking:             GkemulticloudBetaAwsClusterNetworkingToProto(resource.Networking),
		AwsRegion:              dcl.ValueOrEmptyString(resource.AwsRegion),
		ControlPlane:           GkemulticloudBetaAwsClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          GkemulticloudBetaAwsClusterAuthorizationToProto(resource.Authorization),
		State:                  GkemulticloudBetaAwsClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: GkemulticloudBetaAwsClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Apply() method.
func (s *AwsClusterServer) applyAwsCluster(ctx context.Context, c *beta.Client, request *betapb.ApplyGkemulticloudBetaAwsClusterRequest) (*betapb.GkemulticloudBetaAwsCluster, error) {
	p := ProtoToAwsCluster(request.GetResource())
	res, err := c.ApplyAwsCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AwsClusterToProto(res)
	return r, nil
}

// ApplyAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Apply() method.
func (s *AwsClusterServer) ApplyGkemulticloudBetaAwsCluster(ctx context.Context, request *betapb.ApplyGkemulticloudBetaAwsClusterRequest) (*betapb.GkemulticloudBetaAwsCluster, error) {
	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAwsCluster(ctx, cl, request)
}

// DeleteAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Delete() method.
func (s *AwsClusterServer) DeleteGkemulticloudBetaAwsCluster(ctx context.Context, request *betapb.DeleteGkemulticloudBetaAwsClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAwsCluster(ctx, ProtoToAwsCluster(request.GetResource()))

}

// ListGkemulticloudBetaAwsCluster handles the gRPC request by passing it to the underlying AwsClusterList() method.
func (s *AwsClusterServer) ListGkemulticloudBetaAwsCluster(ctx context.Context, request *betapb.ListGkemulticloudBetaAwsClusterRequest) (*betapb.ListGkemulticloudBetaAwsClusterResponse, error) {
	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAwsCluster(ctx, ProtoToAwsCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkemulticloudBetaAwsCluster
	for _, r := range resources.Items {
		rp := AwsClusterToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListGkemulticloudBetaAwsClusterResponse{Items: protos}, nil
}

func createConfigAwsCluster(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
