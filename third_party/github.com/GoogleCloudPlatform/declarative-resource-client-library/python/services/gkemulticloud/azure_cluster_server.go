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

// Server implements the gRPC interface for AzureCluster.
type AzureClusterServer struct{}

// ProtoToAzureClusterStateEnum converts a AzureClusterStateEnum enum from its proto representation.
func ProtoToGkemulticloudAzureClusterStateEnum(e gkemulticloudpb.GkemulticloudAzureClusterStateEnum) *gkemulticloud.AzureClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkemulticloudpb.GkemulticloudAzureClusterStateEnum_name[int32(e)]; ok {
		e := gkemulticloud.AzureClusterStateEnum(n[len("GkemulticloudAzureClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAzureClusterNetworking converts a AzureClusterNetworking resource from its proto representation.
func ProtoToGkemulticloudAzureClusterNetworking(p *gkemulticloudpb.GkemulticloudAzureClusterNetworking) *gkemulticloud.AzureClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterNetworking{
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
func ProtoToGkemulticloudAzureClusterControlPlane(p *gkemulticloudpb.GkemulticloudAzureClusterControlPlane) *gkemulticloud.AzureClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterControlPlane{
		Version:            dcl.StringOrNil(p.Version),
		SubnetId:           dcl.StringOrNil(p.SubnetId),
		VmSize:             dcl.StringOrNil(p.VmSize),
		SshConfig:          ProtoToGkemulticloudAzureClusterControlPlaneSshConfig(p.GetSshConfig()),
		RootVolume:         ProtoToGkemulticloudAzureClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:         ProtoToGkemulticloudAzureClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption: ProtoToGkemulticloudAzureClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneSshConfig converts a AzureClusterControlPlaneSshConfig resource from its proto representation.
func ProtoToGkemulticloudAzureClusterControlPlaneSshConfig(p *gkemulticloudpb.GkemulticloudAzureClusterControlPlaneSshConfig) *gkemulticloud.AzureClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.AuthorizedKey),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneRootVolume converts a AzureClusterControlPlaneRootVolume resource from its proto representation.
func ProtoToGkemulticloudAzureClusterControlPlaneRootVolume(p *gkemulticloudpb.GkemulticloudAzureClusterControlPlaneRootVolume) *gkemulticloud.AzureClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneMainVolume converts a AzureClusterControlPlaneMainVolume resource from its proto representation.
func ProtoToGkemulticloudAzureClusterControlPlaneMainVolume(p *gkemulticloudpb.GkemulticloudAzureClusterControlPlaneMainVolume) *gkemulticloud.AzureClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64OrNil(p.SizeGib),
	}
	return obj
}

// ProtoToAzureClusterControlPlaneDatabaseEncryption converts a AzureClusterControlPlaneDatabaseEncryption resource from its proto representation.
func ProtoToGkemulticloudAzureClusterControlPlaneDatabaseEncryption(p *gkemulticloudpb.GkemulticloudAzureClusterControlPlaneDatabaseEncryption) *gkemulticloud.AzureClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.StringOrNil(p.ResourceGroupId),
		KmsKeyIdentifier: dcl.StringOrNil(p.KmsKeyIdentifier),
	}
	return obj
}

// ProtoToAzureClusterAuthorization converts a AzureClusterAuthorization resource from its proto representation.
func ProtoToGkemulticloudAzureClusterAuthorization(p *gkemulticloudpb.GkemulticloudAzureClusterAuthorization) *gkemulticloud.AzureClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToGkemulticloudAzureClusterAuthorizationAdminUsers(r))
	}
	return obj
}

// ProtoToAzureClusterAuthorizationAdminUsers converts a AzureClusterAuthorizationAdminUsers resource from its proto representation.
func ProtoToGkemulticloudAzureClusterAuthorizationAdminUsers(p *gkemulticloudpb.GkemulticloudAzureClusterAuthorizationAdminUsers) *gkemulticloud.AzureClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.Username),
	}
	return obj
}

// ProtoToAzureClusterWorkloadIdentityConfig converts a AzureClusterWorkloadIdentityConfig resource from its proto representation.
func ProtoToGkemulticloudAzureClusterWorkloadIdentityConfig(p *gkemulticloudpb.GkemulticloudAzureClusterWorkloadIdentityConfig) *gkemulticloud.AzureClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &gkemulticloud.AzureClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.IssuerUri),
		WorkloadPool:     dcl.StringOrNil(p.WorkloadPool),
		IdentityProvider: dcl.StringOrNil(p.IdentityProvider),
	}
	return obj
}

// ProtoToAzureCluster converts a AzureCluster resource from its proto representation.
func ProtoToAzureCluster(p *gkemulticloudpb.GkemulticloudAzureCluster) *gkemulticloud.AzureCluster {
	obj := &gkemulticloud.AzureCluster{
		Name:                   dcl.StringOrNil(p.Name),
		Description:            dcl.StringOrNil(p.Description),
		AzureRegion:            dcl.StringOrNil(p.AzureRegion),
		ResourceGroupId:        dcl.StringOrNil(p.ResourceGroupId),
		AzureClient:            dcl.StringOrNil(p.AzureClient),
		Networking:             ProtoToGkemulticloudAzureClusterNetworking(p.GetNetworking()),
		ControlPlane:           ProtoToGkemulticloudAzureClusterControlPlane(p.GetControlPlane()),
		Authorization:          ProtoToGkemulticloudAzureClusterAuthorization(p.GetAuthorization()),
		State:                  ProtoToGkemulticloudAzureClusterStateEnum(p.GetState()),
		Endpoint:               dcl.StringOrNil(p.Endpoint),
		Uid:                    dcl.StringOrNil(p.Uid),
		Reconciling:            dcl.Bool(p.Reconciling),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                   dcl.StringOrNil(p.Etag),
		WorkloadIdentityConfig: ProtoToGkemulticloudAzureClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                dcl.StringOrNil(p.Project),
		Location:               dcl.StringOrNil(p.Location),
	}
	return obj
}

// AzureClusterStateEnumToProto converts a AzureClusterStateEnum enum to its proto representation.
func GkemulticloudAzureClusterStateEnumToProto(e *gkemulticloud.AzureClusterStateEnum) gkemulticloudpb.GkemulticloudAzureClusterStateEnum {
	if e == nil {
		return gkemulticloudpb.GkemulticloudAzureClusterStateEnum(0)
	}
	if v, ok := gkemulticloudpb.GkemulticloudAzureClusterStateEnum_value["AzureClusterStateEnum"+string(*e)]; ok {
		return gkemulticloudpb.GkemulticloudAzureClusterStateEnum(v)
	}
	return gkemulticloudpb.GkemulticloudAzureClusterStateEnum(0)
}

// AzureClusterNetworkingToProto converts a AzureClusterNetworking resource to its proto representation.
func GkemulticloudAzureClusterNetworkingToProto(o *gkemulticloud.AzureClusterNetworking) *gkemulticloudpb.GkemulticloudAzureClusterNetworking {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterNetworking{
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
func GkemulticloudAzureClusterControlPlaneToProto(o *gkemulticloud.AzureClusterControlPlane) *gkemulticloudpb.GkemulticloudAzureClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterControlPlane{
		Version:            dcl.ValueOrEmptyString(o.Version),
		SubnetId:           dcl.ValueOrEmptyString(o.SubnetId),
		VmSize:             dcl.ValueOrEmptyString(o.VmSize),
		SshConfig:          GkemulticloudAzureClusterControlPlaneSshConfigToProto(o.SshConfig),
		RootVolume:         GkemulticloudAzureClusterControlPlaneRootVolumeToProto(o.RootVolume),
		MainVolume:         GkemulticloudAzureClusterControlPlaneMainVolumeToProto(o.MainVolume),
		DatabaseEncryption: GkemulticloudAzureClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption),
	}
	p.Tags = make(map[string]string)
	for k, r := range o.Tags {
		p.Tags[k] = r
	}
	return p
}

// AzureClusterControlPlaneSshConfigToProto converts a AzureClusterControlPlaneSshConfig resource to its proto representation.
func GkemulticloudAzureClusterControlPlaneSshConfigToProto(o *gkemulticloud.AzureClusterControlPlaneSshConfig) *gkemulticloudpb.GkemulticloudAzureClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.ValueOrEmptyString(o.AuthorizedKey),
	}
	return p
}

// AzureClusterControlPlaneRootVolumeToProto converts a AzureClusterControlPlaneRootVolume resource to its proto representation.
func GkemulticloudAzureClusterControlPlaneRootVolumeToProto(o *gkemulticloud.AzureClusterControlPlaneRootVolume) *gkemulticloudpb.GkemulticloudAzureClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterControlPlaneRootVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureClusterControlPlaneMainVolumeToProto converts a AzureClusterControlPlaneMainVolume resource to its proto representation.
func GkemulticloudAzureClusterControlPlaneMainVolumeToProto(o *gkemulticloud.AzureClusterControlPlaneMainVolume) *gkemulticloudpb.GkemulticloudAzureClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterControlPlaneMainVolume{
		SizeGib: dcl.ValueOrEmptyInt64(o.SizeGib),
	}
	return p
}

// AzureClusterControlPlaneDatabaseEncryptionToProto converts a AzureClusterControlPlaneDatabaseEncryption resource to its proto representation.
func GkemulticloudAzureClusterControlPlaneDatabaseEncryptionToProto(o *gkemulticloud.AzureClusterControlPlaneDatabaseEncryption) *gkemulticloudpb.GkemulticloudAzureClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterControlPlaneDatabaseEncryption{
		ResourceGroupId:  dcl.ValueOrEmptyString(o.ResourceGroupId),
		KmsKeyIdentifier: dcl.ValueOrEmptyString(o.KmsKeyIdentifier),
	}
	return p
}

// AzureClusterAuthorizationToProto converts a AzureClusterAuthorization resource to its proto representation.
func GkemulticloudAzureClusterAuthorizationToProto(o *gkemulticloud.AzureClusterAuthorization) *gkemulticloudpb.GkemulticloudAzureClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterAuthorization{}
	for _, r := range o.AdminUsers {
		p.AdminUsers = append(p.AdminUsers, GkemulticloudAzureClusterAuthorizationAdminUsersToProto(&r))
	}
	return p
}

// AzureClusterAuthorizationAdminUsersToProto converts a AzureClusterAuthorizationAdminUsers resource to its proto representation.
func GkemulticloudAzureClusterAuthorizationAdminUsersToProto(o *gkemulticloud.AzureClusterAuthorizationAdminUsers) *gkemulticloudpb.GkemulticloudAzureClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterAuthorizationAdminUsers{
		Username: dcl.ValueOrEmptyString(o.Username),
	}
	return p
}

// AzureClusterWorkloadIdentityConfigToProto converts a AzureClusterWorkloadIdentityConfig resource to its proto representation.
func GkemulticloudAzureClusterWorkloadIdentityConfigToProto(o *gkemulticloud.AzureClusterWorkloadIdentityConfig) *gkemulticloudpb.GkemulticloudAzureClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &gkemulticloudpb.GkemulticloudAzureClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.ValueOrEmptyString(o.IssuerUri),
		WorkloadPool:     dcl.ValueOrEmptyString(o.WorkloadPool),
		IdentityProvider: dcl.ValueOrEmptyString(o.IdentityProvider),
	}
	return p
}

// AzureClusterToProto converts a AzureCluster resource to its proto representation.
func AzureClusterToProto(resource *gkemulticloud.AzureCluster) *gkemulticloudpb.GkemulticloudAzureCluster {
	p := &gkemulticloudpb.GkemulticloudAzureCluster{
		Name:                   dcl.ValueOrEmptyString(resource.Name),
		Description:            dcl.ValueOrEmptyString(resource.Description),
		AzureRegion:            dcl.ValueOrEmptyString(resource.AzureRegion),
		ResourceGroupId:        dcl.ValueOrEmptyString(resource.ResourceGroupId),
		AzureClient:            dcl.ValueOrEmptyString(resource.AzureClient),
		Networking:             GkemulticloudAzureClusterNetworkingToProto(resource.Networking),
		ControlPlane:           GkemulticloudAzureClusterControlPlaneToProto(resource.ControlPlane),
		Authorization:          GkemulticloudAzureClusterAuthorizationToProto(resource.Authorization),
		State:                  GkemulticloudAzureClusterStateEnumToProto(resource.State),
		Endpoint:               dcl.ValueOrEmptyString(resource.Endpoint),
		Uid:                    dcl.ValueOrEmptyString(resource.Uid),
		Reconciling:            dcl.ValueOrEmptyBool(resource.Reconciling),
		CreateTime:             dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:             dcl.ValueOrEmptyString(resource.UpdateTime),
		Etag:                   dcl.ValueOrEmptyString(resource.Etag),
		WorkloadIdentityConfig: GkemulticloudAzureClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig),
		Project:                dcl.ValueOrEmptyString(resource.Project),
		Location:               dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Apply() method.
func (s *AzureClusterServer) applyAzureCluster(ctx context.Context, c *gkemulticloud.Client, request *gkemulticloudpb.ApplyGkemulticloudAzureClusterRequest) (*gkemulticloudpb.GkemulticloudAzureCluster, error) {
	p := ProtoToAzureCluster(request.GetResource())
	res, err := c.ApplyAzureCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureClusterToProto(res)
	return r, nil
}

// ApplyAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Apply() method.
func (s *AzureClusterServer) ApplyGkemulticloudAzureCluster(ctx context.Context, request *gkemulticloudpb.ApplyGkemulticloudAzureClusterRequest) (*gkemulticloudpb.GkemulticloudAzureCluster, error) {
	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureCluster(ctx, cl, request)
}

// DeleteAzureCluster handles the gRPC request by passing it to the underlying AzureCluster Delete() method.
func (s *AzureClusterServer) DeleteGkemulticloudAzureCluster(ctx context.Context, request *gkemulticloudpb.DeleteGkemulticloudAzureClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureCluster(ctx, ProtoToAzureCluster(request.GetResource()))

}

// ListGkemulticloudAzureCluster handles the gRPC request by passing it to the underlying AzureClusterList() method.
func (s *AzureClusterServer) ListGkemulticloudAzureCluster(ctx context.Context, request *gkemulticloudpb.ListGkemulticloudAzureClusterRequest) (*gkemulticloudpb.ListGkemulticloudAzureClusterResponse, error) {
	cl, err := createConfigAzureCluster(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureCluster(ctx, ProtoToAzureCluster(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*gkemulticloudpb.GkemulticloudAzureCluster
	for _, r := range resources.Items {
		rp := AzureClusterToProto(r)
		protos = append(protos, rp)
	}
	return &gkemulticloudpb.ListGkemulticloudAzureClusterResponse{Items: protos}, nil
}

func createConfigAzureCluster(ctx context.Context, service_account_file string) (*gkemulticloud.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return gkemulticloud.NewClient(conf), nil
}
