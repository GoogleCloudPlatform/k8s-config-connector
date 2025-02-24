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

// Server implements the gRPC interface for AzureCluster.
type AzureClusterServer struct{}

// ProtoToAzureClusterStateEnum converts a AzureClusterStateEnum enum from its proto representation.
func ProtoToGkemulticloudAlphaAzureClusterStateEnum(e alphapb.GkemulticloudAlphaAzureClusterStateEnum) *alpha.AzureClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkemulticloudAlphaAzureClusterStateEnum_name[int32(e)]; ok {
		e := alpha.AzureClusterStateEnum(n[len("GkemulticloudAlphaAzureClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAzureClusterNetworking converts a AzureClusterNetworking resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureClusterNetworking(p *alphapb.GkemulticloudAlphaAzureClusterNetworking) *alpha.AzureClusterNetworking {
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
func ProtoToGkemulticloudAlphaAzureClusterControlPlane(p *alphapb.GkemulticloudAlphaAzureClusterControlPlane) *alpha.AzureClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlane{
		Version:            dcl.StringOrNil(p.Version),
		SubnetId:           dcl.StringOrNil(p.SubnetId),
		VmSize:             dcl.StringOrNil(p.VmSize),
		SshConfig:          ProtoToGkemulticloudAlphaAzureClusterControlPlaneSshConfig(p.GetSshConfig()),
		RootVolume:         ProtoToGkemulticloudAlphaAzureClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:         ProtoToGkemulticloudAlphaAzureClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption: ProtoToGkemulticloudAlphaAzureClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneSshConfig converts a AzureClusterControlPlaneSshConfig resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureClusterControlPlaneSshConfig(p *alphapb.GkemulticloudAlphaAzureClusterControlPlaneSshConfig) *alpha.AzureClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneRootVolume converts a AzureClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureClusterControlPlaneRootVolume(p *alphapb.GkemulticloudAlphaAzureClusterControlPlaneRootVolume) *alpha.AzureClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneMainVolume converts a AzureClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureClusterControlPlaneMainVolume(p *alphapb.GkemulticloudAlphaAzureClusterControlPlaneMainVolume) *alpha.AzureClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneDatabaseEncryption converts a AzureClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureClusterControlPlaneDatabaseEncryption(p *alphapb.GkemulticloudAlphaAzureClusterControlPlaneDatabaseEncryption) *alpha.AzureClusterControlPlaneDatabaseEncryption {
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
func ProtoToGkemulticloudAlphaAzureClusterAuthorization(p *alphapb.GkemulticloudAlphaAzureClusterAuthorization) *alpha.AzureClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToGkemulticloudAlphaAzureClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToAzureClusterAuthorizationAdminUsers converts a AzureClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureClusterAuthorizationAdminUsers(p *alphapb.GkemulticloudAlphaAzureClusterAuthorizationAdminUsers) *alpha.AzureClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &alpha.AzureClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToAzureClusterWorkloadIdentityConfig converts a AzureClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToGkemulticloudAlphaAzureClusterWorkloadIdentityConfig(p *alphapb.GkemulticloudAlphaAzureClusterWorkloadIdentityConfig) *alpha.AzureClusterWorkloadIdentityConfig {
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
func ProtoToAzureCluster(p *alphapb.GkemulticloudAlphaAzureCluster) *alpha.AzureCluster {
	obj := &alpha.AzureCluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		AzureRegion:            dcl.StringOrNil(p.AzureRegion),
		ResourceGroupId:        dcl.StringOrNil(p.ResourceGroupId),
		AzureClient:            dcl.StringOrNil(p.AzureClient),
		Networking:             ProtoToGkemulticloudAlphaAzureClusterNetworking(p.GetNetworking()),
		ControlPlane:           ProtoToGkemulticloudAlphaAzureClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToGkemulticloudAlphaAzureClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToGkemulticloudAlphaAzureClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToGkemulticloudAlphaAzureClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// AzureClusterStateEnumToProto converts a AzureClusterStateEnum enum to its proto representation.
func GkemulticloudAlphaAzureClusterStateEnumToProto(e *alpha.AzureClusterStateEnum) alphapb.GkemulticloudAlphaAzureClusterStateEnum {
	if e == nil {
		return alphapb.GkemulticloudAlphaAzureClusterStateEnum(0)
	}
	if v, ok := alphapb.GkemulticloudAlphaAzureClusterStateEnum_value["AzureClusterStateEnum"+string(*e)]; ok {
		return alphapb.GkemulticloudAlphaAzureClusterStateEnum(v)
	}
	return alphapb.GkemulticloudAlphaAzureClusterStateEnum(0)
}

// AzureClusterNetworkingToProto converts a AzureClusterNetworking resource to its proto representation.
func GkemulticloudAlphaAzureClusterNetworkingToProto(o *alpha.AzureClusterNetworking) *alphapb.GkemulticloudAlphaAzureClusterNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterNetworking{
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
func GkemulticloudAlphaAzureClusterControlPlaneToProto(o *alpha.AzureClusterControlPlane) *alphapb.GkemulticloudAlphaAzureClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterControlPlane{
		Version:            dcl.ValueOrEmptyString(o.Version),
		SubnetId:           dcl.ValueOrEmptyString(o.SubnetId),
		VmSize:             dcl.ValueOrEmptyString(o.VmSize),
		SshConfig:          GkemulticloudAlphaAzureClusterControlPlaneSshConfigToProto(o.SshConfig),
		RootVolume:         GkemulticloudAlphaAzureClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:         GkemulticloudAlphaAzureClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption: GkemulticloudAlphaAzureClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AzureClusterControlPlaneSshConfigToProto converts a AzureClusterControlPlaneSshConfig resource to its proto representation.
func GkemulticloudAlphaAzureClusterControlPlaneSshConfigToProto(o *alpha.AzureClusterControlPlaneSshConfig) *alphapb.GkemulticloudAlphaAzureClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// AzureClusterControlPlaneRootVolumeToProto converts a AzureClusterControlPlaneRootVolume resource to its proto representation.
func GkemulticloudAlphaAzureClusterControlPlaneRootVolumeToProto(o *alpha.AzureClusterControlPlaneRootVolume) *alphapb.GkemulticloudAlphaAzureClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterControlPlaneRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureClusterControlPlaneMainVolumeToProto converts a AzureClusterControlPlaneMainVolume resource to its proto representation.
func GkemulticloudAlphaAzureClusterControlPlaneMainVolumeToProto(o *alpha.AzureClusterControlPlaneMainVolume) *alphapb.GkemulticloudAlphaAzureClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterControlPlaneMainVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureClusterControlPlaneDatabaseEncryptionToProto converts a AzureClusterControlPlaneDatabaseEncryption resource to its proto representation.
func GkemulticloudAlphaAzureClusterControlPlaneDatabaseEncryptionToProto(o *alpha.AzureClusterControlPlaneDatabaseEncryption) *alphapb.GkemulticloudAlphaAzureClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.ValueOrEmptyString(o.ResourceGroupId),
		KmsKeyIdentifier: dcl.ValueOrEmptyString(o.KmsKeyIdentifier),
	}
	return p
}

// AzureClusterAuthorizationToProto converts a AzureClusterAuthorization resource to its proto representation.
func GkemulticloudAlphaAzureClusterAuthorizationToProto(o *alpha.AzureClusterAuthorization) *alphapb.GkemulticloudAlphaAzureClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, GkemulticloudAlphaAzureClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// AzureClusterAuthorizationAdminUsersToProto converts a AzureClusterAuthorizationAdminUsers resource to its proto representation.
func GkemulticloudAlphaAzureClusterAuthorizationAdminUsersToProto(o *alpha.AzureClusterAuthorizationAdminUsers) *alphapb.GkemulticloudAlphaAzureClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// AzureClusterWorkloadIdentityConfigToProto converts a AzureClusterWorkloadIdentityConfig resource to its proto representation.
func GkemulticloudAlphaAzureClusterWorkloadIdentityConfigToProto(o *alpha.AzureClusterWorkloadIdentityConfig) *alphapb.GkemulticloudAlphaAzureClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkemulticloudAlphaAzureClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// AzureClusterToProto converts a AzureCluster resource to its proto representation.
func AzureClusterToProto(resource *alpha.AzureCluster) *alphapb.GkemulticloudAlphaAzureCluster {
	p := &alphapb.GkemulticloudAlphaAzureCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		AzureRegion:            dcl.ValueOrEmptyString(resource.AzureRegion),
		ResourceGroupId:        dcl.ValueOrEmptyString(resource.ResourceGroupId),
		AzureClient:            dcl.ValueOrEmptyString(resource.AzureClient),
		Networking:             GkemulticloudAlphaAzureClusterNetworkingToProto(resource.Networking),
		ControlPlane:           GkemulticloudAlphaAzureClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          GkemulticloudAlphaAzureClusterAuthorizationToProto(resource.Authorization),
		State:                  GkemulticloudAlphaAzureClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: GkemulticloudAlphaAzureClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Apply() method.
func (s *AzureClusterServer) applyAzureCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkemulticloudAlphaAzureClusterRequest) (*alphapb.GkemulticloudAlphaAzureCluster, error) {
	p := ProtoToAzureCluster(request.GetResource())
	res, err := c.ApplyAzureCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureClusterToProto(res)
	return r, nil
}

// ApplyAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Apply() method.
func (s *AzureClusterServer) ApplyGkemulticloudAlphaAzureCluster(ctx context.Context, request *alphapb.ApplyGkemulticloudAlphaAzureClusterRequest) (*alphapb.GkemulticloudAlphaAzureCluster, error) {
	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureCluster(ctx, cl, request)
}

// DeleteAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Delete() method.
func (s *AzureClusterServer) DeleteGkemulticloudAlphaAzureCluster(ctx context.Context, request *alphapb.DeleteGkemulticloudAlphaAzureClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureCluster(ctx, ProtoToAzureCluster(request.GetResource()))

}

// ListGkemulticloudAlphaAzureCluster handles the gRPC request by passing it to the underlying AzureClusterList() method.
func (s *AzureClusterServer) ListGkemulticloudAlphaAzureCluster(ctx context.Context, request *alphapb.ListGkemulticloudAlphaAzureClusterRequest) (*alphapb.ListGkemulticloudAlphaAzureClusterResponse, error) {
	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureCluster(ctx, ProtoToAzureCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkemulticloudAlphaAzureCluster
	for _, r := range resources.Items {
		rp := AzureClusterToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListGkemulticloudAlphaAzureClusterResponse{Items: protos}, nil
}

func createConfigAzureCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
