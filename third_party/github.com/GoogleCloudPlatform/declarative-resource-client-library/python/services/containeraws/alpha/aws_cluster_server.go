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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containeraws/alpha/containeraws_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/alpha"
)

// Server implements the gRPC interface for AwsCluster.
type AwsClusterServer struct{}

// ProtoToAwsClusterControlPlaneRootVolumeVolumeTypeEnum converts a AwsClusterControlPlaneRootVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum) *alpha.AwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.AwsClusterControlPlaneRootVolumeVolumeTypeEnum(n[len("ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterControlPlaneMainVolumeVolumeTypeEnum converts a AwsClusterControlPlaneMainVolumeVolumeTypeEnum enum from its proto representation.
func ProtoToContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(e alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum) *alpha.AwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum_name[int32(e)]; ok {
		e := alpha.AwsClusterControlPlaneMainVolumeVolumeTypeEnum(n[len("ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterStateEnum converts a AwsClusterStateEnum enum from its proto representation.
func ProtoToContainerawsAlphaAwsClusterStateEnum(e alphapb.ContainerawsAlphaAwsClusterStateEnum) *alpha.AwsClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerawsAlphaAwsClusterStateEnum_name[int32(e)]; ok {
		e := alpha.AwsClusterStateEnum(n[len("ContainerawsAlphaAwsClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAwsClusterNetworking converts a AwsClusterNetworking resource from its proto representation.
func ProtoToContainerawsAlphaAwsClusterNetworking(p *alphapb.ContainerawsAlphaAwsClusterNetworking) *alpha.AwsClusterNetworking {
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
func ProtoToContainerawsAlphaAwsClusterControlPlane(p *alphapb.ContainerawsAlphaAwsClusterControlPlane) *alpha.AwsClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlane{
		Version:                   dcl.StringOrNil(p.Version),
		InstanceType:              dcl.StringOrNil(p.InstanceType),
		SshConfig:                 ProtoToContainerawsAlphaAwsClusterControlPlaneSshConfig(p.GetSshConfig()),
		IamInstanceProfile:        dcl.StringOrNil(p.IamInstanceProfile),
		RootVolume:                ProtoToContainerawsAlphaAwsClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:                ProtoToContainerawsAlphaAwsClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption:        ProtoToContainerawsAlphaAwsClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		AwsServicesAuthentication: ProtoToContainerawsAlphaAwsClusterControlPlaneAwsServicesAuthentication(p.GetAwsServicesAuthentication()),
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
func ProtoToContainerawsAlphaAwsClusterControlPlaneSshConfig(p *alphapb.ContainerawsAlphaAwsClusterControlPlaneSshConfig) *alpha.AwsClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.StringOrNil(p.Ec2KeyPair),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneRootVolume converts a AwsClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToContainerawsAlphaAwsClusterControlPlaneRootVolume(p *alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolume) *alpha.AwsClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneRootVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneMainVolume converts a AwsClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToContainerawsAlphaAwsClusterControlPlaneMainVolume(p *alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolume) *alpha.AwsClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneMainVolume{
		SizeGib:    dcl.Int64OrNil(p.SizeGib),
		VolumeType: ProtoToContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(p.GetVolumeType()),
		Iops:       dcl.Int64OrNil(p.Iops),
		KmsKeyArn:  dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneDatabaseEncryption converts a AwsClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToContainerawsAlphaAwsClusterControlPlaneDatabaseEncryption(p *alphapb.ContainerawsAlphaAwsClusterControlPlaneDatabaseEncryption) *alpha.AwsClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.StringOrNil(p.KmsKeyArn),
	}
	return obj
}

// ProtoToAwsClusterControlPlaneAwsServicesAuthentication converts a AwsClusterControlPlaneAwsServicesAuthentication resource from its proto representation.
func ProtoToContainerawsAlphaAwsClusterControlPlaneAwsServicesAuthentication(p *alphapb.ContainerawsAlphaAwsClusterControlPlaneAwsServicesAuthentication) *alpha.AwsClusterControlPlaneAwsServicesAuthentication {
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
func ProtoToContainerawsAlphaAwsClusterAuthorization(p *alphapb.ContainerawsAlphaAwsClusterAuthorization) *alpha.AwsClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerawsAlphaAwsClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToAwsClusterAuthorizationAdminUsers converts a AwsClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToContainerawsAlphaAwsClusterAuthorizationAdminUsers(p *alphapb.ContainerawsAlphaAwsClusterAuthorizationAdminUsers) *alpha.AwsClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.AwsClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToAwsClusterWorkloadIdentityConfig converts a AwsClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToContainerawsAlphaAwsClusterWorkloadIdentityConfig(p *alphapb.ContainerawsAlphaAwsClusterWorkloadIdentityConfig) *alpha.AwsClusterWorkloadIdentityConfig {
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
func ProtoToAwsCluster(p *alphapb.ContainerawsAlphaAwsCluster) *alpha.AwsCluster {
	obj := &alpha.AwsCluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		Networking:             ProtoToContainerawsAlphaAwsClusterNetworking(p.GetNetworking()),
		AwsRegion:              dcl.StringOrNil(p.AwsRegion),
		ControlPlane:           ProtoToContainerawsAlphaAwsClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToContainerawsAlphaAwsClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToContainerawsAlphaAwsClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToContainerawsAlphaAwsClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// AwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto converts a AwsClusterControlPlaneRootVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(e *alpha.AwsClusterControlPlaneRootVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum_value["AwsClusterControlPlaneRootVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnum(0)
}

// AwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto converts a AwsClusterControlPlaneMainVolumeVolumeTypeEnum enum to its proto representation.
func ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(e *alpha.AwsClusterControlPlaneMainVolumeVolumeTypeEnum) alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum_value["AwsClusterControlPlaneMainVolumeVolumeTypeEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(v)
	}
	return alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnum(0)
}

// AwsClusterStateEnumToProto converts a AwsClusterStateEnum enum to its proto representation.
func ContainerawsAlphaAwsClusterStateEnumToProto(e *alpha.AwsClusterStateEnum) alphapb.ContainerawsAlphaAwsClusterStateEnum {
	if e == nil {
		return alphapb.ContainerawsAlphaAwsClusterStateEnum(0)
	}
	if v, ok := alphapb.ContainerawsAlphaAwsClusterStateEnum_value["AwsClusterStateEnum"+string(*e)]; ok {
		return alphapb.ContainerawsAlphaAwsClusterStateEnum(v)
	}
	return alphapb.ContainerawsAlphaAwsClusterStateEnum(0)
}

// AwsClusterNetworkingToProto converts a AwsClusterNetworking resource to its proto representation.
func ContainerawsAlphaAwsClusterNetworkingToProto(o *alpha.AwsClusterNetworking) *alphapb.ContainerawsAlphaAwsClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterNetworking{
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
func ContainerawsAlphaAwsClusterControlPlaneToProto(o *alpha.AwsClusterControlPlane) *alphapb.ContainerawsAlphaAwsClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterControlPlane{
		Version:                   dcl.ValueOrEmptyString(o.Version),
		InstanceType:              dcl.ValueOrEmptyString(o.InstanceType),
		SshConfig:                 ContainerawsAlphaAwsClusterControlPlaneSshConfigToProto(o.SshConfig),
		IamInstanceProfile:        dcl.ValueOrEmptyString(o.IamInstanceProfile),
		RootVolume:                ContainerawsAlphaAwsClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:                ContainerawsAlphaAwsClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption:        ContainerawsAlphaAwsClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
		AwsServicesAuthentication: ContainerawsAlphaAwsClusterControlPlaneAwsServicesAuthenticationToProto(o.AwsServicesAuthentication),
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
func ContainerawsAlphaAwsClusterControlPlaneSshConfigToProto(o *alpha.AwsClusterControlPlaneSshConfig) *alphapb.ContainerawsAlphaAwsClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.ValueOrEmptyString(o.Ec2KeyPair),
	}
	return p
}

// AwsClusterControlPlaneRootVolumeToProto converts a AwsClusterControlPlaneRootVolume resource to its proto representation.
func ContainerawsAlphaAwsClusterControlPlaneRootVolumeToProto(o *alpha.AwsClusterControlPlaneRootVolume) *alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterControlPlaneRootVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: ContainerawsAlphaAwsClusterControlPlaneRootVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneMainVolumeToProto converts a AwsClusterControlPlaneMainVolume resource to its proto representation.
func ContainerawsAlphaAwsClusterControlPlaneMainVolumeToProto(o *alpha.AwsClusterControlPlaneMainVolume) *alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterControlPlaneMainVolume{
		SizeGib:    dcl.ValueOrEmptyInt64(o.SizeGib),
		VolumeType: ContainerawsAlphaAwsClusterControlPlaneMainVolumeVolumeTypeEnumToProto(o.VolumeType),
		Iops:       dcl.ValueOrEmptyInt64(o.Iops),
		KmsKeyArn:  dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneDatabaseEncryptionToProto converts a AwsClusterControlPlaneDatabaseEncryption resource to its proto representation.
func ContainerawsAlphaAwsClusterControlPlaneDatabaseEncryptionToProto(o *alpha.AwsClusterControlPlaneDatabaseEncryption) *alphapb.ContainerawsAlphaAwsClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.ValueOrEmptyString(o.KmsKeyArn),
	}
	return p
}

// AwsClusterControlPlaneAwsServicesAuthenticationToProto converts a AwsClusterControlPlaneAwsServicesAuthentication resource to its proto representation.
func ContainerawsAlphaAwsClusterControlPlaneAwsServicesAuthenticationToProto(o *alpha.AwsClusterControlPlaneAwsServicesAuthentication) *alphapb.ContainerawsAlphaAwsClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.ValueOrEmptyString(o.RoleArn),
		RoleSessionName: dcl.ValueOrEmptyString(o.RoleSessionName),
	}
	return p
}

// AwsClusterAuthorizationToProto converts a AwsClusterAuthorization resource to its proto representation.
func ContainerawsAlphaAwsClusterAuthorizationToProto(o *alpha.AwsClusterAuthorization) *alphapb.ContainerawsAlphaAwsClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, ContainerawsAlphaAwsClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// AwsClusterAuthorizationAdminUsersToProto converts a AwsClusterAuthorizationAdminUsers resource to its proto representation.
func ContainerawsAlphaAwsClusterAuthorizationAdminUsersToProto(o *alpha.AwsClusterAuthorizationAdminUsers) *alphapb.ContainerawsAlphaAwsClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// AwsClusterWorkloadIdentityConfigToProto converts a AwsClusterWorkloadIdentityConfig resource to its proto representation.
func ContainerawsAlphaAwsClusterWorkloadIdentityConfigToProto(o *alpha.AwsClusterWorkloadIdentityConfig) *alphapb.ContainerawsAlphaAwsClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerawsAlphaAwsClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// AwsClusterToProto converts a AwsCluster resource to its proto representation.
func AwsClusterToProto(resource *alpha.AwsCluster) *alphapb.ContainerawsAlphaAwsCluster {
	p := &alphapb.ContainerawsAlphaAwsCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		Networking:             ContainerawsAlphaAwsClusterNetworkingToProto(resource.Networking),
		AwsRegion:              dcl.ValueOrEmptyString(resource.AwsRegion),
		ControlPlane:           ContainerawsAlphaAwsClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          ContainerawsAlphaAwsClusterAuthorizationToProto(resource.Authorization),
		State:                  ContainerawsAlphaAwsClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: ContainerawsAlphaAwsClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Apply() method.
func (s *AwsClusterServer) applyAwsCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerawsAlphaAwsClusterRequest) (*alphapb.ContainerawsAlphaAwsCluster, error) {
	p := ProtoToAwsCluster(request.GetResource())
	res, err := c.ApplyAwsCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AwsClusterToProto(res)
	return r, nil
}

// ApplyAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Apply() method.
func (s *AwsClusterServer) ApplyContainerawsAlphaAwsCluster(ctx context.Context, request *alphapb.ApplyContainerawsAlphaAwsClusterRequest) (*alphapb.ContainerawsAlphaAwsCluster, error) {
	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAwsCluster(ctx, cl, request)
}

// DeleteAwsCluster handles the gRPC request by passing it to the underlying AwsCluster Delete() method.
func (s *AwsClusterServer) DeleteContainerawsAlphaAwsCluster(ctx context.Context, request *alphapb.DeleteContainerawsAlphaAwsClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAwsCluster(ctx, ProtoToAwsCluster(request.GetResource()))

}

// ListContainerawsAlphaAwsCluster handles the gRPC request by passing it to the underlying AwsClusterList() method.
func (s *AwsClusterServer) ListContainerawsAlphaAwsCluster(ctx context.Context, request *alphapb.ListContainerawsAlphaAwsClusterRequest) (*alphapb.ListContainerawsAlphaAwsClusterResponse, error) {
	cl, err := createConfigAwsCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAwsCluster(ctx, ProtoToAwsCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerawsAlphaAwsCluster
	for _, r := range resources.Items {
		rp := AwsClusterToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContainerawsAlphaAwsClusterResponse{Items: protos}, nil
}

func createConfigAwsCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
