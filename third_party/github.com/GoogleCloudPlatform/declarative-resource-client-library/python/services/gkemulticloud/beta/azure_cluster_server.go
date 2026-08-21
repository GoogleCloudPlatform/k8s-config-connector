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

// Server implements the gRPC interface for AzureCluster.
type AzureClusterServer struct{}

// ProtoToAzureClusterStateEnum converts a AzureClusterStateEnum enum from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterStateEnum(e betapb.GkemulticloudBetaAzureClusterStateEnum) *beta.AzureClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkemulticloudBetaAzureClusterStateEnum_name[int32(e)]; ok {
		e := beta.AzureClusterStateEnum(n[len("GkemulticloudBetaAzureClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAzureClusterNetworking converts a AzureClusterNetworking resource from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterNetworking(p *betapb.GkemulticloudBetaAzureClusterNetworking) *beta.AzureClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterNetworking{
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
func ProtoToGkemulticloudBetaAzureClusterControlPlane(p *betapb.GkemulticloudBetaAzureClusterControlPlane) *beta.AzureClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterControlPlane{
		Version:            dcl.StringOrNil(p.Version),
		SubnetId:           dcl.StringOrNil(p.SubnetId),
		VmSize:             dcl.StringOrNil(p.VmSize),
		SshConfig:          ProtoToGkemulticloudBetaAzureClusterControlPlaneSshConfig(p.GetSshConfig()),
		RootVolume:         ProtoToGkemulticloudBetaAzureClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:         ProtoToGkemulticloudBetaAzureClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption: ProtoToGkemulticloudBetaAzureClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneSshConfig converts a AzureClusterControlPlaneSshConfig resource from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterControlPlaneSshConfig(p *betapb.GkemulticloudBetaAzureClusterControlPlaneSshConfig) *beta.AzureClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneRootVolume converts a AzureClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterControlPlaneRootVolume(p *betapb.GkemulticloudBetaAzureClusterControlPlaneRootVolume) *beta.AzureClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneMainVolume converts a AzureClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterControlPlaneMainVolume(p *betapb.GkemulticloudBetaAzureClusterControlPlaneMainVolume) *beta.AzureClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneDatabaseEncryption converts a AzureClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterControlPlaneDatabaseEncryption(p *betapb.GkemulticloudBetaAzureClusterControlPlaneDatabaseEncryption) *beta.AzureClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.StringOrNil(p.ResourceGroupId),
		KmsKeyIdentifier: dcl.StringOrNil(p.KmsKeyIdentifier),
	}
	return obj
}

// ProtoToAzureClusterAuthorization converts a AzureClusterAuthorization resource from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterAuthorization(p *betapb.GkemulticloudBetaAzureClusterAuthorization) *beta.AzureClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToGkemulticloudBetaAzureClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToAzureClusterAuthorizationAdminUsers converts a AzureClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterAuthorizationAdminUsers(p *betapb.GkemulticloudBetaAzureClusterAuthorizationAdminUsers) *beta.AzureClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToAzureClusterWorkloadIdentityConfig converts a AzureClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToGkemulticloudBetaAzureClusterWorkloadIdentityConfig(p *betapb.GkemulticloudBetaAzureClusterWorkloadIdentityConfig) *beta.AzureClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.AzureClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.IssuerUri),
		WorkloadPool:     dcl.StringOrNil(p.WorkloadPool),
		IdentityProvider: dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToAzureCluster converts a AzureCluster resource from its proto representation.
func ProtoToAzureCluster(p *betapb.GkemulticloudBetaAzureCluster) *beta.AzureCluster {
	obj := &beta.AzureCluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		AzureRegion:            dcl.StringOrNil(p.AzureRegion),
		ResourceGroupId:        dcl.StringOrNil(p.ResourceGroupId),
		AzureClient:            dcl.StringOrNil(p.AzureClient),
		Networking:             ProtoToGkemulticloudBetaAzureClusterNetworking(p.GetNetworking()),
		ControlPlane:           ProtoToGkemulticloudBetaAzureClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToGkemulticloudBetaAzureClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToGkemulticloudBetaAzureClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToGkemulticloudBetaAzureClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// AzureClusterStateEnumToProto converts a AzureClusterStateEnum enum to its proto representation.
func GkemulticloudBetaAzureClusterStateEnumToProto(e *beta.AzureClusterStateEnum) betapb.GkemulticloudBetaAzureClusterStateEnum {
	if e == nil {
		return betapb.GkemulticloudBetaAzureClusterStateEnum(0)
	}
	if v, ok := betapb.GkemulticloudBetaAzureClusterStateEnum_value["AzureClusterStateEnum"+string(*e)]; ok {
		return betapb.GkemulticloudBetaAzureClusterStateEnum(v)
	}
	return betapb.GkemulticloudBetaAzureClusterStateEnum(0)
}

// AzureClusterNetworkingToProto converts a AzureClusterNetworking resource to its proto representation.
func GkemulticloudBetaAzureClusterNetworkingToProto(o *beta.AzureClusterNetworking) *betapb.GkemulticloudBetaAzureClusterNetworking {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterNetworking{
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
func GkemulticloudBetaAzureClusterControlPlaneToProto(o *beta.AzureClusterControlPlane) *betapb.GkemulticloudBetaAzureClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterControlPlane{
		Version:            dcl.ValueOrEmptyString(o.Version),
		SubnetId:           dcl.ValueOrEmptyString(o.SubnetId),
		VmSize:             dcl.ValueOrEmptyString(o.VmSize),
		SshConfig:          GkemulticloudBetaAzureClusterControlPlaneSshConfigToProto(o.SshConfig),
		RootVolume:         GkemulticloudBetaAzureClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:         GkemulticloudBetaAzureClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption: GkemulticloudBetaAzureClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AzureClusterControlPlaneSshConfigToProto converts a AzureClusterControlPlaneSshConfig resource to its proto representation.
func GkemulticloudBetaAzureClusterControlPlaneSshConfigToProto(o *beta.AzureClusterControlPlaneSshConfig) *betapb.GkemulticloudBetaAzureClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// AzureClusterControlPlaneRootVolumeToProto converts a AzureClusterControlPlaneRootVolume resource to its proto representation.
func GkemulticloudBetaAzureClusterControlPlaneRootVolumeToProto(o *beta.AzureClusterControlPlaneRootVolume) *betapb.GkemulticloudBetaAzureClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterControlPlaneRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureClusterControlPlaneMainVolumeToProto converts a AzureClusterControlPlaneMainVolume resource to its proto representation.
func GkemulticloudBetaAzureClusterControlPlaneMainVolumeToProto(o *beta.AzureClusterControlPlaneMainVolume) *betapb.GkemulticloudBetaAzureClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterControlPlaneMainVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureClusterControlPlaneDatabaseEncryptionToProto converts a AzureClusterControlPlaneDatabaseEncryption resource to its proto representation.
func GkemulticloudBetaAzureClusterControlPlaneDatabaseEncryptionToProto(o *beta.AzureClusterControlPlaneDatabaseEncryption) *betapb.GkemulticloudBetaAzureClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.ValueOrEmptyString(o.ResourceGroupId),
		KmsKeyIdentifier: dcl.ValueOrEmptyString(o.KmsKeyIdentifier),
	}
	return p
}

// AzureClusterAuthorizationToProto converts a AzureClusterAuthorization resource to its proto representation.
func GkemulticloudBetaAzureClusterAuthorizationToProto(o *beta.AzureClusterAuthorization) *betapb.GkemulticloudBetaAzureClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, GkemulticloudBetaAzureClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// AzureClusterAuthorizationAdminUsersToProto converts a AzureClusterAuthorizationAdminUsers resource to its proto representation.
func GkemulticloudBetaAzureClusterAuthorizationAdminUsersToProto(o *beta.AzureClusterAuthorizationAdminUsers) *betapb.GkemulticloudBetaAzureClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// AzureClusterWorkloadIdentityConfigToProto converts a AzureClusterWorkloadIdentityConfig resource to its proto representation.
func GkemulticloudBetaAzureClusterWorkloadIdentityConfigToProto(o *beta.AzureClusterWorkloadIdentityConfig) *betapb.GkemulticloudBetaAzureClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkemulticloudBetaAzureClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// AzureClusterToProto converts a AzureCluster resource to its proto representation.
func AzureClusterToProto(resource *beta.AzureCluster) *betapb.GkemulticloudBetaAzureCluster {
	p := &betapb.GkemulticloudBetaAzureCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		AzureRegion:            dcl.ValueOrEmptyString(resource.AzureRegion),
		ResourceGroupId:        dcl.ValueOrEmptyString(resource.ResourceGroupId),
		AzureClient:            dcl.ValueOrEmptyString(resource.AzureClient),
		Networking:             GkemulticloudBetaAzureClusterNetworkingToProto(resource.Networking),
		ControlPlane:           GkemulticloudBetaAzureClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          GkemulticloudBetaAzureClusterAuthorizationToProto(resource.Authorization),
		State:                  GkemulticloudBetaAzureClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: GkemulticloudBetaAzureClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Apply() method.
func (s *AzureClusterServer) applyAzureCluster(ctx context.Context, c *beta.Client, request *betapb.ApplyGkemulticloudBetaAzureClusterRequest) (*betapb.GkemulticloudBetaAzureCluster, error) {
	p := ProtoToAzureCluster(request.GetResource())
	res, err := c.ApplyAzureCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureClusterToProto(res)
	return r, nil
}

// ApplyAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Apply() method.
func (s *AzureClusterServer) ApplyGkemulticloudBetaAzureCluster(ctx context.Context, request *betapb.ApplyGkemulticloudBetaAzureClusterRequest) (*betapb.GkemulticloudBetaAzureCluster, error) {
	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureCluster(ctx, cl, request)
}

// DeleteAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Delete() method.
func (s *AzureClusterServer) DeleteGkemulticloudBetaAzureCluster(ctx context.Context, request *betapb.DeleteGkemulticloudBetaAzureClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureCluster(ctx, ProtoToAzureCluster(request.GetResource()))

}

// ListGkemulticloudBetaAzureCluster handles the gRPC request by passing it to the underlying AzureClusterList() method.
func (s *AzureClusterServer) ListGkemulticloudBetaAzureCluster(ctx context.Context, request *betapb.ListGkemulticloudBetaAzureClusterRequest) (*betapb.ListGkemulticloudBetaAzureClusterResponse, error) {
	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureCluster(ctx, ProtoToAzureCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkemulticloudBetaAzureCluster
	for _, r := range resources.Items {
		rp := AzureClusterToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListGkemulticloudBetaAzureClusterResponse{Items: protos}, nil
}

func createConfigAzureCluster(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
