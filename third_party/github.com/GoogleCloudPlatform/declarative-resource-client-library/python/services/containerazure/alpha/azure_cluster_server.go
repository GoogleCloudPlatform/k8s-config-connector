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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/alpha/containerazure_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/alpha"
)

// Server implements the gRPC interface for AzureCluster.
type AzureClusterServer struct{}

// ProtoToAzureClusterStateEnum converts a AzureClusterStateEnum enum from its proto representation.
func ProtoToContainerazureAlphaAzureClusterStateEnum(e alphapb.ContainerazureAlphaAzureClusterStateEnum) *alpha.AzureClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerazureAlphaAzureClusterStateEnum_name[int32(e)]; ok {
		e := alpha.AzureClusterStateEnum(n[len("ContainerazureAlphaAzureClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAzureClusterNetworking converts a AzureClusterNetworking resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterNetworking(p *alphapb.ContainerazureAlphaAzureClusterNetworking) *alpha.AzureClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterNetworking{
		VirtualNetworkId: dcl.StringOrNil(p.VirtualNetworkId),
	}
	for _, r := range p.GetPodAddressCidrBlocks() {
		obj.PodAddressCidrBlocks = append(obj.PodAddressCidrBlocks, r)
	}
	for _, r := range p.GetServiceAddressCidrBlocks() {
		obj.ServiceAddressCidrBlocks = append(obj.ServiceAddressCidrBlocks, r)
	}
	return obj
}

// ProtoToAzureClusterControlPlane converts a AzureClusterControlPlane resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterControlPlane(p *alphapb.ContainerazureAlphaAzureClusterControlPlane) *alpha.AzureClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlane{
		Version:            dcl.StringOrNil(p.Version),
		SubnetId:           dcl.StringOrNil(p.SubnetId),
		VmSize:             dcl.StringOrNil(p.VmSize),
		SshConfig:          ProtoToContainerazureAlphaAzureClusterControlPlaneSshConfig(p.GetSshConfig()),
		RootVolume:         ProtoToContainerazureAlphaAzureClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:         ProtoToContainerazureAlphaAzureClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption: ProtoToContainerazureAlphaAzureClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneSshConfig converts a AzureClusterControlPlaneSshConfig resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterControlPlaneSshConfig(p *alphapb.ContainerazureAlphaAzureClusterControlPlaneSshConfig) *alpha.AzureClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneRootVolume converts a AzureClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterControlPlaneRootVolume(p *alphapb.ContainerazureAlphaAzureClusterControlPlaneRootVolume) *alpha.AzureClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneMainVolume converts a AzureClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterControlPlaneMainVolume(p *alphapb.ContainerazureAlphaAzureClusterControlPlaneMainVolume) *alpha.AzureClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneDatabaseEncryption converts a AzureClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterControlPlaneDatabaseEncryption(p *alphapb.ContainerazureAlphaAzureClusterControlPlaneDatabaseEncryption) *alpha.AzureClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.StringOrNil(p.ResourceGroupId),
		KmsKeyIdentifier: dcl.StringOrNil(p.KmsKeyIdentifier),
	}
	return obj
}

// ProtoToAzureClusterAuthorization converts a AzureClusterAuthorization resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterAuthorization(p *alphapb.ContainerazureAlphaAzureClusterAuthorization) *alpha.AzureClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerazureAlphaAzureClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToAzureClusterAuthorizationAdminUsers converts a AzureClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterAuthorizationAdminUsers(p *alphapb.ContainerazureAlphaAzureClusterAuthorizationAdminUsers) *alpha.AzureClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToAzureClusterWorkloadIdentityConfig converts a AzureClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToContainerazureAlphaAzureClusterWorkloadIdentityConfig(p *alphapb.ContainerazureAlphaAzureClusterWorkloadIdentityConfig) *alpha.AzureClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.IssuerUri),
		WorkloadPool:     dcl.StringOrNil(p.WorkloadPool),
		IdentityProvider: dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToAzureCluster converts a AzureCluster resource from its proto representation.
func ProtoToAzureCluster(p *alphapb.ContainerazureAlphaAzureCluster) *alpha.AzureCluster {
	obj := &alpha.AzureCluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		AzureRegion:            dcl.StringOrNil(p.AzureRegion),
		ResourceGroupId:        dcl.StringOrNil(p.ResourceGroupId),
		AzureClient:            dcl.StringOrNil(p.AzureClient),
		Networking:             ProtoToContainerazureAlphaAzureClusterNetworking(p.GetNetworking()),
		ControlPlane:           ProtoToContainerazureAlphaAzureClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToContainerazureAlphaAzureClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToContainerazureAlphaAzureClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToContainerazureAlphaAzureClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// AzureClusterStateEnumToProto converts a AzureClusterStateEnum enum to its proto representation.
func ContainerazureAlphaAzureClusterStateEnumToProto(e *alpha.AzureClusterStateEnum) alphapb.ContainerazureAlphaAzureClusterStateEnum {
	if e == nil {
		return alphapb.ContainerazureAlphaAzureClusterStateEnum(0)
	}
	if v, ok := alphapb.ContainerazureAlphaAzureClusterStateEnum_value["AzureClusterStateEnum"+string(*e)]; ok {
		return alphapb.ContainerazureAlphaAzureClusterStateEnum(v)
	}
	return alphapb.ContainerazureAlphaAzureClusterStateEnum(0)
}

// AzureClusterNetworkingToProto converts a AzureClusterNetworking resource to its proto representation.
func ContainerazureAlphaAzureClusterNetworkingToProto(o *alpha.AzureClusterNetworking) *alphapb.ContainerazureAlphaAzureClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterNetworking{
		VirtualNetworkId: dcl.ValueOrEmptyString(o.VirtualNetworkId),
	}
	for _, r := range o.PodAddressCidrBlocks {
		p.PodAddressCidrBlocks = append(p.PodAddressCidrBlocks, r)
	}
	for _, r := range o.ServiceAddressCidrBlocks {
		p.ServiceAddressCidrBlocks = append(p.ServiceAddressCidrBlocks, r)
	}
	return p
}

// AzureClusterControlPlaneToProto converts a AzureClusterControlPlane resource to its proto representation.
func ContainerazureAlphaAzureClusterControlPlaneToProto(o *alpha.AzureClusterControlPlane) *alphapb.ContainerazureAlphaAzureClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterControlPlane{
		Version:            dcl.ValueOrEmptyString(o.Version),
		SubnetId:           dcl.ValueOrEmptyString(o.SubnetId),
		VmSize:             dcl.ValueOrEmptyString(o.VmSize),
		SshConfig:          ContainerazureAlphaAzureClusterControlPlaneSshConfigToProto(o.SshConfig),
		RootVolume:         ContainerazureAlphaAzureClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:         ContainerazureAlphaAzureClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption: ContainerazureAlphaAzureClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AzureClusterControlPlaneSshConfigToProto converts a AzureClusterControlPlaneSshConfig resource to its proto representation.
func ContainerazureAlphaAzureClusterControlPlaneSshConfigToProto(o *alpha.AzureClusterControlPlaneSshConfig) *alphapb.ContainerazureAlphaAzureClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// AzureClusterControlPlaneRootVolumeToProto converts a AzureClusterControlPlaneRootVolume resource to its proto representation.
func ContainerazureAlphaAzureClusterControlPlaneRootVolumeToProto(o *alpha.AzureClusterControlPlaneRootVolume) *alphapb.ContainerazureAlphaAzureClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterControlPlaneRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureClusterControlPlaneMainVolumeToProto converts a AzureClusterControlPlaneMainVolume resource to its proto representation.
func ContainerazureAlphaAzureClusterControlPlaneMainVolumeToProto(o *alpha.AzureClusterControlPlaneMainVolume) *alphapb.ContainerazureAlphaAzureClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterControlPlaneMainVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureClusterControlPlaneDatabaseEncryptionToProto converts a AzureClusterControlPlaneDatabaseEncryption resource to its proto representation.
func ContainerazureAlphaAzureClusterControlPlaneDatabaseEncryptionToProto(o *alpha.AzureClusterControlPlaneDatabaseEncryption) *alphapb.ContainerazureAlphaAzureClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.ValueOrEmptyString(o.ResourceGroupId),
		KmsKeyIdentifier: dcl.ValueOrEmptyString(o.KmsKeyIdentifier),
	}
	return p
}

// AzureClusterAuthorizationToProto converts a AzureClusterAuthorization resource to its proto representation.
func ContainerazureAlphaAzureClusterAuthorizationToProto(o *alpha.AzureClusterAuthorization) *alphapb.ContainerazureAlphaAzureClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, ContainerazureAlphaAzureClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// AzureClusterAuthorizationAdminUsersToProto converts a AzureClusterAuthorizationAdminUsers resource to its proto representation.
func ContainerazureAlphaAzureClusterAuthorizationAdminUsersToProto(o *alpha.AzureClusterAuthorizationAdminUsers) *alphapb.ContainerazureAlphaAzureClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// AzureClusterWorkloadIdentityConfigToProto converts a AzureClusterWorkloadIdentityConfig resource to its proto representation.
func ContainerazureAlphaAzureClusterWorkloadIdentityConfigToProto(o *alpha.AzureClusterWorkloadIdentityConfig) *alphapb.ContainerazureAlphaAzureClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaAzureClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// AzureClusterToProto converts a AzureCluster resource to its proto representation.
func AzureClusterToProto(resource *alpha.AzureCluster) *alphapb.ContainerazureAlphaAzureCluster {
	p := &alphapb.ContainerazureAlphaAzureCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		AzureRegion:            dcl.ValueOrEmptyString(resource.AzureRegion),
		ResourceGroupId:        dcl.ValueOrEmptyString(resource.ResourceGroupId),
		AzureClient:            dcl.ValueOrEmptyString(resource.AzureClient),
		Networking:             ContainerazureAlphaAzureClusterNetworkingToProto(resource.Networking),
		ControlPlane:           ContainerazureAlphaAzureClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          ContainerazureAlphaAzureClusterAuthorizationToProto(resource.Authorization),
		State:                  ContainerazureAlphaAzureClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: ContainerazureAlphaAzureClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Apply() method.
func (s *AzureClusterServer) applyAzureCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerazureAlphaAzureClusterRequest) (*alphapb.ContainerazureAlphaAzureCluster, error) {
	p := ProtoToAzureCluster(request.GetResource())
	res, err := c.ApplyAzureCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureClusterToProto(res)
	return r, nil
}

// ApplyAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Apply() method.
func (s *AzureClusterServer) ApplyContainerazureAlphaAzureCluster(ctx context.Context, request *alphapb.ApplyContainerazureAlphaAzureClusterRequest) (*alphapb.ContainerazureAlphaAzureCluster, error) {
	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureCluster(ctx, cl, request)
}

// DeleteAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Delete() method.
func (s *AzureClusterServer) DeleteContainerazureAlphaAzureCluster(ctx context.Context, request *alphapb.DeleteContainerazureAlphaAzureClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureCluster(ctx, ProtoToAzureCluster(request.GetResource()))

}

// ListContainerazureAlphaAzureCluster handles the gRPC request by passing it to the underlying AzureClusterList() method.
func (s *AzureClusterServer) ListContainerazureAlphaAzureCluster(ctx context.Context, request *alphapb.ListContainerazureAlphaAzureClusterRequest) (*alphapb.ListContainerazureAlphaAzureClusterResponse, error) {
	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureCluster(ctx, ProtoToAzureCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerazureAlphaAzureCluster
	for _, r := range resources.Items {
		rp := AzureClusterToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListContainerazureAlphaAzureClusterResponse{Items: protos}, nil
}

func createConfigAzureCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
