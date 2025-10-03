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
	containerazurepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/containerazure_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure"
)

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToContainerazureClusterStateEnum(e containerazurepb.ContainerazureClusterStateEnum) *containerazure.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerazurepb.ContainerazureClusterStateEnum_name[int32(e)]; ok {
		e := containerazure.ClusterStateEnum(n[len("ContainerazureClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterAzureServicesAuthentication converts a ClusterAzureServicesAuthentication object from its proto representation.
func ProtoToContainerazureClusterAzureServicesAuthentication(p *containerazurepb.ContainerazureClusterAzureServicesAuthentication) *containerazure.ClusterAzureServicesAuthentication {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterAzureServicesAuthentication{
		TenantId:      dcl.StringOrNil(p.GetTenantId()),
		ApplicationId: dcl.StringOrNil(p.GetApplicationId()),
	}
	return obj
}

// ProtoToClusterNetworking converts a ClusterNetworking object from its proto representation.
func ProtoToContainerazureClusterNetworking(p *containerazurepb.ContainerazureClusterNetworking) *containerazure.ClusterNetworking {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterNetworking{
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
func ProtoToContainerazureClusterControlPlane(p *containerazurepb.ContainerazureClusterControlPlane) *containerazure.ClusterControlPlane {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterControlPlane{
		Version:            dcl.StringOrNil(p.GetVersion()),
		SubnetId:           dcl.StringOrNil(p.GetSubnetId()),
		VmSize:             dcl.StringOrNil(p.GetVmSize()),
		SshConfig:          ProtoToContainerazureClusterControlPlaneSshConfig(p.GetSshConfig()),
		RootVolume:         ProtoToContainerazureClusterControlPlaneRootVolume(p.GetRootVolume()),
		MainVolume:         ProtoToContainerazureClusterControlPlaneMainVolume(p.GetMainVolume()),
		DatabaseEncryption: ProtoToContainerazureClusterControlPlaneDatabaseEncryption(p.GetDatabaseEncryption()),
		ProxyConfig:        ProtoToContainerazureClusterControlPlaneProxyConfig(p.GetProxyConfig()),
	}
	for _, r := range p.GetReplicaPlacements() {
		obj.ReplicaPlacements = append(obj.ReplicaPlacements, *ProtoToContainerazureClusterControlPlaneReplicaPlacements(r))
	}
	return obj
}

// ProtoToClusterControlPlaneSshConfig converts a ClusterControlPlaneSshConfig object from its proto representation.
func ProtoToContainerazureClusterControlPlaneSshConfig(p *containerazurepb.ContainerazureClusterControlPlaneSshConfig) *containerazure.ClusterControlPlaneSshConfig {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.GetAuthorizedKey()),
	}
	return obj
}

// ProtoToClusterControlPlaneRootVolume converts a ClusterControlPlaneRootVolume object from its proto representation.
func ProtoToContainerazureClusterControlPlaneRootVolume(p *containerazurepb.ContainerazureClusterControlPlaneRootVolume) *containerazure.ClusterControlPlaneRootVolume {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToClusterControlPlaneMainVolume converts a ClusterControlPlaneMainVolume object from its proto representation.
func ProtoToContainerazureClusterControlPlaneMainVolume(p *containerazurepb.ContainerazureClusterControlPlaneMainVolume) *containerazure.ClusterControlPlaneMainVolume {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToClusterControlPlaneDatabaseEncryption converts a ClusterControlPlaneDatabaseEncryption object from its proto representation.
func ProtoToContainerazureClusterControlPlaneDatabaseEncryption(p *containerazurepb.ContainerazureClusterControlPlaneDatabaseEncryption) *containerazure.ClusterControlPlaneDatabaseEncryption {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterControlPlaneDatabaseEncryption{
		KeyId: dcl.StringOrNil(p.GetKeyId()),
	}
	return obj
}

// ProtoToClusterControlPlaneProxyConfig converts a ClusterControlPlaneProxyConfig object from its proto representation.
func ProtoToContainerazureClusterControlPlaneProxyConfig(p *containerazurepb.ContainerazureClusterControlPlaneProxyConfig) *containerazure.ClusterControlPlaneProxyConfig {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterControlPlaneProxyConfig{
		ResourceGroupId: dcl.StringOrNil(p.GetResourceGroupId()),
		SecretId:        dcl.StringOrNil(p.GetSecretId()),
	}
	return obj
}

// ProtoToClusterControlPlaneReplicaPlacements converts a ClusterControlPlaneReplicaPlacements object from its proto representation.
func ProtoToContainerazureClusterControlPlaneReplicaPlacements(p *containerazurepb.ContainerazureClusterControlPlaneReplicaPlacements) *containerazure.ClusterControlPlaneReplicaPlacements {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterControlPlaneReplicaPlacements{
		SubnetId:              dcl.StringOrNil(p.GetSubnetId()),
		AzureAvailabilityZone: dcl.StringOrNil(p.GetAzureAvailabilityZone()),
	}
	return obj
}

// ProtoToClusterAuthorization converts a ClusterAuthorization object from its proto representation.
func ProtoToContainerazureClusterAuthorization(p *containerazurepb.ContainerazureClusterAuthorization) *containerazure.ClusterAuthorization {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterAuthorization{}
	for _, r := range p.GetAdminUsers() {
		obj.AdminUsers = append(obj.AdminUsers, *ProtoToContainerazureClusterAuthorizationAdminUsers(r))
	}
	for _, r := range p.GetAdminGroups() {
		obj.AdminGroups = append(obj.AdminGroups, *ProtoToContainerazureClusterAuthorizationAdminGroups(r))
	}
	return obj
}

// ProtoToClusterAuthorizationAdminUsers converts a ClusterAuthorizationAdminUsers object from its proto representation.
func ProtoToContainerazureClusterAuthorizationAdminUsers(p *containerazurepb.ContainerazureClusterAuthorizationAdminUsers) *containerazure.ClusterAuthorizationAdminUsers {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterAuthorizationAdminUsers{
		Username: dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToClusterAuthorizationAdminGroups converts a ClusterAuthorizationAdminGroups object from its proto representation.
func ProtoToContainerazureClusterAuthorizationAdminGroups(p *containerazurepb.ContainerazureClusterAuthorizationAdminGroups) *containerazure.ClusterAuthorizationAdminGroups {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterAuthorizationAdminGroups{
		Group: dcl.StringOrNil(p.GetGroup()),
	}
	return obj
}

// ProtoToClusterWorkloadIdentityConfig converts a ClusterWorkloadIdentityConfig object from its proto representation.
func ProtoToContainerazureClusterWorkloadIdentityConfig(p *containerazurepb.ContainerazureClusterWorkloadIdentityConfig) *containerazure.ClusterWorkloadIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterWorkloadIdentityConfig{
		IssuerUri:        dcl.StringOrNil(p.GetIssuerUri()),
		WorkloadPool:     dcl.StringOrNil(p.GetWorkloadPool()),
		IdentityProvider: dcl.StringOrNil(p.GetIdentityProvider()),
	}
	return obj
}

// ProtoToClusterFleet converts a ClusterFleet object from its proto representation.
func ProtoToContainerazureClusterFleet(p *containerazurepb.ContainerazureClusterFleet) *containerazure.ClusterFleet {
	if p == nil {
		return nil
	}
	obj := &containerazure.ClusterFleet{
		Project:    dcl.StringOrNil(p.GetProject()),
		Membership: dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *containerazurepb.ContainerazureCluster) *containerazure.Cluster {
	obj := &containerazure.Cluster{
		Name:                        dcl.StringOrNil(p.GetName()),
		Description:                 dcl.StringOrNil(p.GetDescription()),
		AzureRegion:                 dcl.StringOrNil(p.GetAzureRegion()),
		ResourceGroupId:             dcl.StringOrNil(p.GetResourceGroupId()),
		Client:                      dcl.StringOrNil(p.GetClient()),
		AzureServicesAuthentication: ProtoToContainerazureClusterAzureServicesAuthentication(p.GetAzureServicesAuthentication()),
		Networking:                  ProtoToContainerazureClusterNetworking(p.GetNetworking()),
		ControlPlane:                ProtoToContainerazureClusterControlPlane(p.GetControlPlane()),
		Authorization:               ProtoToContainerazureClusterAuthorization(p.GetAuthorization()),
		State:                       ProtoToContainerazureClusterStateEnum(p.GetState()),
		Endpoint:                    dcl.StringOrNil(p.GetEndpoint()),
		Uid:                         dcl.StringOrNil(p.GetUid()),
		Reconciling:                 dcl.Bool(p.GetReconciling()),
		CreateTime:                  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                  dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                        dcl.StringOrNil(p.GetEtag()),
		WorkloadIdentityConfig:      ProtoToContainerazureClusterWorkloadIdentityConfig(p.GetWorkloadIdentityConfig()),
		Project:                     dcl.StringOrNil(p.GetProject()),
		Location:                    dcl.StringOrNil(p.GetLocation()),
		Fleet:                       ProtoToContainerazureClusterFleet(p.GetFleet()),
	}
	return obj
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func ContainerazureClusterStateEnumToProto(e *containerazure.ClusterStateEnum) containerazurepb.ContainerazureClusterStateEnum {
	if e == nil {
		return containerazurepb.ContainerazureClusterStateEnum(0)
	}
	if v, ok := containerazurepb.ContainerazureClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return containerazurepb.ContainerazureClusterStateEnum(v)
	}
	return containerazurepb.ContainerazureClusterStateEnum(0)
}

// ClusterAzureServicesAuthenticationToProto converts a ClusterAzureServicesAuthentication object to its proto representation.
func ContainerazureClusterAzureServicesAuthenticationToProto(o *containerazure.ClusterAzureServicesAuthentication) *containerazurepb.ContainerazureClusterAzureServicesAuthentication {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterAzureServicesAuthentication{}
	p.SetTenantId(dcl.ValueOrEmptyString(o.TenantId))
	p.SetApplicationId(dcl.ValueOrEmptyString(o.ApplicationId))
	return p
}

// ClusterNetworkingToProto converts a ClusterNetworking object to its proto representation.
func ContainerazureClusterNetworkingToProto(o *containerazure.ClusterNetworking) *containerazurepb.ContainerazureClusterNetworking {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterNetworking{}
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
func ContainerazureClusterControlPlaneToProto(o *containerazure.ClusterControlPlane) *containerazurepb.ContainerazureClusterControlPlane {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterControlPlane{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetSubnetId(dcl.ValueOrEmptyString(o.SubnetId))
	p.SetVmSize(dcl.ValueOrEmptyString(o.VmSize))
	p.SetSshConfig(ContainerazureClusterControlPlaneSshConfigToProto(o.SshConfig))
	p.SetRootVolume(ContainerazureClusterControlPlaneRootVolumeToProto(o.RootVolume))
	p.SetMainVolume(ContainerazureClusterControlPlaneMainVolumeToProto(o.MainVolume))
	p.SetDatabaseEncryption(ContainerazureClusterControlPlaneDatabaseEncryptionToProto(o.DatabaseEncryption))
	p.SetProxyConfig(ContainerazureClusterControlPlaneProxyConfigToProto(o.ProxyConfig))
	mTags := make(map[string]string, len(o.Tags))
	for k, r := range o.Tags {
		mTags[k] = r
	}
	p.SetTags(mTags)
	sReplicaPlacements := make([]*containerazurepb.ContainerazureClusterControlPlaneReplicaPlacements, len(o.ReplicaPlacements))
	for i, r := range o.ReplicaPlacements {
		sReplicaPlacements[i] = ContainerazureClusterControlPlaneReplicaPlacementsToProto(&r)
	}
	p.SetReplicaPlacements(sReplicaPlacements)
	return p
}

// ClusterControlPlaneSshConfigToProto converts a ClusterControlPlaneSshConfig object to its proto representation.
func ContainerazureClusterControlPlaneSshConfigToProto(o *containerazure.ClusterControlPlaneSshConfig) *containerazurepb.ContainerazureClusterControlPlaneSshConfig {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterControlPlaneSshConfig{}
	p.SetAuthorizedKey(dcl.ValueOrEmptyString(o.AuthorizedKey))
	return p
}

// ClusterControlPlaneRootVolumeToProto converts a ClusterControlPlaneRootVolume object to its proto representation.
func ContainerazureClusterControlPlaneRootVolumeToProto(o *containerazure.ClusterControlPlaneRootVolume) *containerazurepb.ContainerazureClusterControlPlaneRootVolume {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterControlPlaneRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// ClusterControlPlaneMainVolumeToProto converts a ClusterControlPlaneMainVolume object to its proto representation.
func ContainerazureClusterControlPlaneMainVolumeToProto(o *containerazure.ClusterControlPlaneMainVolume) *containerazurepb.ContainerazureClusterControlPlaneMainVolume {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterControlPlaneMainVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// ClusterControlPlaneDatabaseEncryptionToProto converts a ClusterControlPlaneDatabaseEncryption object to its proto representation.
func ContainerazureClusterControlPlaneDatabaseEncryptionToProto(o *containerazure.ClusterControlPlaneDatabaseEncryption) *containerazurepb.ContainerazureClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterControlPlaneDatabaseEncryption{}
	p.SetKeyId(dcl.ValueOrEmptyString(o.KeyId))
	return p
}

// ClusterControlPlaneProxyConfigToProto converts a ClusterControlPlaneProxyConfig object to its proto representation.
func ContainerazureClusterControlPlaneProxyConfigToProto(o *containerazure.ClusterControlPlaneProxyConfig) *containerazurepb.ContainerazureClusterControlPlaneProxyConfig {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterControlPlaneProxyConfig{}
	p.SetResourceGroupId(dcl.ValueOrEmptyString(o.ResourceGroupId))
	p.SetSecretId(dcl.ValueOrEmptyString(o.SecretId))
	return p
}

// ClusterControlPlaneReplicaPlacementsToProto converts a ClusterControlPlaneReplicaPlacements object to its proto representation.
func ContainerazureClusterControlPlaneReplicaPlacementsToProto(o *containerazure.ClusterControlPlaneReplicaPlacements) *containerazurepb.ContainerazureClusterControlPlaneReplicaPlacements {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterControlPlaneReplicaPlacements{}
	p.SetSubnetId(dcl.ValueOrEmptyString(o.SubnetId))
	p.SetAzureAvailabilityZone(dcl.ValueOrEmptyString(o.AzureAvailabilityZone))
	return p
}

// ClusterAuthorizationToProto converts a ClusterAuthorization object to its proto representation.
func ContainerazureClusterAuthorizationToProto(o *containerazure.ClusterAuthorization) *containerazurepb.ContainerazureClusterAuthorization {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterAuthorization{}
	sAdminUsers := make([]*containerazurepb.ContainerazureClusterAuthorizationAdminUsers, len(o.AdminUsers))
	for i, r := range o.AdminUsers {
		sAdminUsers[i] = ContainerazureClusterAuthorizationAdminUsersToProto(&r)
	}
	p.SetAdminUsers(sAdminUsers)
	sAdminGroups := make([]*containerazurepb.ContainerazureClusterAuthorizationAdminGroups, len(o.AdminGroups))
	for i, r := range o.AdminGroups {
		sAdminGroups[i] = ContainerazureClusterAuthorizationAdminGroupsToProto(&r)
	}
	p.SetAdminGroups(sAdminGroups)
	return p
}

// ClusterAuthorizationAdminUsersToProto converts a ClusterAuthorizationAdminUsers object to its proto representation.
func ContainerazureClusterAuthorizationAdminUsersToProto(o *containerazure.ClusterAuthorizationAdminUsers) *containerazurepb.ContainerazureClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterAuthorizationAdminUsers{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ClusterAuthorizationAdminGroupsToProto converts a ClusterAuthorizationAdminGroups object to its proto representation.
func ContainerazureClusterAuthorizationAdminGroupsToProto(o *containerazure.ClusterAuthorizationAdminGroups) *containerazurepb.ContainerazureClusterAuthorizationAdminGroups {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterAuthorizationAdminGroups{}
	p.SetGroup(dcl.ValueOrEmptyString(o.Group))
	return p
}

// ClusterWorkloadIdentityConfigToProto converts a ClusterWorkloadIdentityConfig object to its proto representation.
func ContainerazureClusterWorkloadIdentityConfigToProto(o *containerazure.ClusterWorkloadIdentityConfig) *containerazurepb.ContainerazureClusterWorkloadIdentityConfig {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterWorkloadIdentityConfig{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetWorkloadPool(dcl.ValueOrEmptyString(o.WorkloadPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// ClusterFleetToProto converts a ClusterFleet object to its proto representation.
func ContainerazureClusterFleetToProto(o *containerazure.ClusterFleet) *containerazurepb.ContainerazureClusterFleet {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureClusterFleet{}
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetMembership(dcl.ValueOrEmptyString(o.Membership))
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *containerazure.Cluster) *containerazurepb.ContainerazureCluster {
	p := &containerazurepb.ContainerazureCluster{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetAzureRegion(dcl.ValueOrEmptyString(resource.AzureRegion))
	p.SetResourceGroupId(dcl.ValueOrEmptyString(resource.ResourceGroupId))
	p.SetClient(dcl.ValueOrEmptyString(resource.Client))
	p.SetAzureServicesAuthentication(ContainerazureClusterAzureServicesAuthenticationToProto(resource.AzureServicesAuthentication))
	p.SetNetworking(ContainerazureClusterNetworkingToProto(resource.Networking))
	p.SetControlPlane(ContainerazureClusterControlPlaneToProto(resource.ControlPlane))
	p.SetAuthorization(ContainerazureClusterAuthorizationToProto(resource.Authorization))
	p.SetState(ContainerazureClusterStateEnumToProto(resource.State))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetWorkloadIdentityConfig(ContainerazureClusterWorkloadIdentityConfigToProto(resource.WorkloadIdentityConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFleet(ContainerazureClusterFleetToProto(resource.Fleet))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *containerazure.Client, request *containerazurepb.ApplyContainerazureClusterRequest) (*containerazurepb.ContainerazureCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyContainerazureCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyContainerazureCluster(ctx context.Context, request *containerazurepb.ApplyContainerazureClusterRequest) (*containerazurepb.ContainerazureCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteContainerazureCluster(ctx context.Context, request *containerazurepb.DeleteContainerazureClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListContainerazureCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListContainerazureCluster(ctx context.Context, request *containerazurepb.ListContainerazureClusterRequest) (*containerazurepb.ListContainerazureClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*containerazurepb.ContainerazureCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &containerazurepb.ListContainerazureClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*containerazure.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return containerazure.NewClient(conf), nil
}
