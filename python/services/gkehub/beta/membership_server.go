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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/beta/gkehub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

// MembershipServer implements the gRPC interface for Membership.
type MembershipServer struct{}

// ProtoToMembershipStateCodeEnum converts a MembershipStateCodeEnum enum from its proto representation.
func ProtoToGkehubBetaMembershipStateCodeEnum(e betapb.GkehubBetaMembershipStateCodeEnum) *beta.MembershipStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaMembershipStateCodeEnum_name[int32(e)]; ok {
		e := beta.MembershipStateCodeEnum(n[len("GkehubBetaMembershipStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipInfrastructureTypeEnum converts a MembershipInfrastructureTypeEnum enum from its proto representation.
func ProtoToGkehubBetaMembershipInfrastructureTypeEnum(e betapb.GkehubBetaMembershipInfrastructureTypeEnum) *beta.MembershipInfrastructureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaMembershipInfrastructureTypeEnum_name[int32(e)]; ok {
		e := beta.MembershipInfrastructureTypeEnum(n[len("GkehubBetaMembershipInfrastructureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipEndpoint converts a MembershipEndpoint object from its proto representation.
func ProtoToGkehubBetaMembershipEndpoint(p *betapb.GkehubBetaMembershipEndpoint) *beta.MembershipEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpoint{
		GkeCluster:         ProtoToGkehubBetaMembershipEndpointGkeCluster(p.GetGkeCluster()),
		KubernetesMetadata: ProtoToGkehubBetaMembershipEndpointKubernetesMetadata(p.GetKubernetesMetadata()),
		KubernetesResource: ProtoToGkehubBetaMembershipEndpointKubernetesResource(p.GetKubernetesResource()),
	}
	return obj
}

// ProtoToMembershipEndpointGkeCluster converts a MembershipEndpointGkeCluster object from its proto representation.
func ProtoToGkehubBetaMembershipEndpointGkeCluster(p *betapb.GkehubBetaMembershipEndpointGkeCluster) *beta.MembershipEndpointGkeCluster {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointGkeCluster{
		ResourceLink: dcl.StringOrNil(p.GetResourceLink()),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesMetadata converts a MembershipEndpointKubernetesMetadata object from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesMetadata(p *betapb.GkehubBetaMembershipEndpointKubernetesMetadata) *beta.MembershipEndpointKubernetesMetadata {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesMetadata{
		KubernetesApiServerVersion: dcl.StringOrNil(p.GetKubernetesApiServerVersion()),
		NodeProviderId:             dcl.StringOrNil(p.GetNodeProviderId()),
		NodeCount:                  dcl.Int64OrNil(p.GetNodeCount()),
		VcpuCount:                  dcl.Int64OrNil(p.GetVcpuCount()),
		MemoryMb:                   dcl.Int64OrNil(p.GetMemoryMb()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResource converts a MembershipEndpointKubernetesResource object from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesResource(p *betapb.GkehubBetaMembershipEndpointKubernetesResource) *beta.MembershipEndpointKubernetesResource {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesResource{
		MembershipCrManifest: dcl.StringOrNil(p.GetMembershipCrManifest()),
		ResourceOptions:      ProtoToGkehubBetaMembershipEndpointKubernetesResourceResourceOptions(p.GetResourceOptions()),
	}
	for _, r := range p.GetMembershipResources() {
		obj.MembershipResources = append(obj.MembershipResources, *ProtoToGkehubBetaMembershipEndpointKubernetesResourceMembershipResources(r))
	}
	for _, r := range p.GetConnectResources() {
		obj.ConnectResources = append(obj.ConnectResources, *ProtoToGkehubBetaMembershipEndpointKubernetesResourceConnectResources(r))
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceMembershipResources converts a MembershipEndpointKubernetesResourceMembershipResources object from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesResourceMembershipResources(p *betapb.GkehubBetaMembershipEndpointKubernetesResourceMembershipResources) *beta.MembershipEndpointKubernetesResourceMembershipResources {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesResourceMembershipResources{
		Manifest:      dcl.StringOrNil(p.GetManifest()),
		ClusterScoped: dcl.Bool(p.GetClusterScoped()),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceConnectResources converts a MembershipEndpointKubernetesResourceConnectResources object from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesResourceConnectResources(p *betapb.GkehubBetaMembershipEndpointKubernetesResourceConnectResources) *beta.MembershipEndpointKubernetesResourceConnectResources {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesResourceConnectResources{
		Manifest:      dcl.StringOrNil(p.GetManifest()),
		ClusterScoped: dcl.Bool(p.GetClusterScoped()),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceResourceOptions converts a MembershipEndpointKubernetesResourceResourceOptions object from its proto representation.
func ProtoToGkehubBetaMembershipEndpointKubernetesResourceResourceOptions(p *betapb.GkehubBetaMembershipEndpointKubernetesResourceResourceOptions) *beta.MembershipEndpointKubernetesResourceResourceOptions {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipEndpointKubernetesResourceResourceOptions{
		ConnectVersion: dcl.StringOrNil(p.GetConnectVersion()),
		V1Beta1Crd:     dcl.Bool(p.GetV1Beta1Crd()),
	}
	return obj
}

// ProtoToMembershipState converts a MembershipState object from its proto representation.
func ProtoToGkehubBetaMembershipState(p *betapb.GkehubBetaMembershipState) *beta.MembershipState {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipState{
		Code: ProtoToGkehubBetaMembershipStateCodeEnum(p.GetCode()),
	}
	return obj
}

// ProtoToMembershipAuthority converts a MembershipAuthority object from its proto representation.
func ProtoToGkehubBetaMembershipAuthority(p *betapb.GkehubBetaMembershipAuthority) *beta.MembershipAuthority {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipAuthority{
		Issuer:               dcl.StringOrNil(p.GetIssuer()),
		WorkloadIdentityPool: dcl.StringOrNil(p.GetWorkloadIdentityPool()),
		IdentityProvider:     dcl.StringOrNil(p.GetIdentityProvider()),
	}
	return obj
}

// ProtoToMembership converts a Membership resource from its proto representation.
func ProtoToMembership(p *betapb.GkehubBetaMembership) *beta.Membership {
	obj := &beta.Membership{
		Endpoint:           ProtoToGkehubBetaMembershipEndpoint(p.GetEndpoint()),
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		State:              ProtoToGkehubBetaMembershipState(p.GetState()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:         dcl.StringOrNil(p.GetDeleteTime()),
		ExternalId:         dcl.StringOrNil(p.GetExternalId()),
		LastConnectionTime: dcl.StringOrNil(p.GetLastConnectionTime()),
		UniqueId:           dcl.StringOrNil(p.GetUniqueId()),
		Authority:          ProtoToGkehubBetaMembershipAuthority(p.GetAuthority()),
		InfrastructureType: ProtoToGkehubBetaMembershipInfrastructureTypeEnum(p.GetInfrastructureType()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// MembershipStateCodeEnumToProto converts a MembershipStateCodeEnum enum to its proto representation.
func GkehubBetaMembershipStateCodeEnumToProto(e *beta.MembershipStateCodeEnum) betapb.GkehubBetaMembershipStateCodeEnum {
	if e == nil {
		return betapb.GkehubBetaMembershipStateCodeEnum(0)
	}
	if v, ok := betapb.GkehubBetaMembershipStateCodeEnum_value["MembershipStateCodeEnum"+string(*e)]; ok {
		return betapb.GkehubBetaMembershipStateCodeEnum(v)
	}
	return betapb.GkehubBetaMembershipStateCodeEnum(0)
}

// MembershipInfrastructureTypeEnumToProto converts a MembershipInfrastructureTypeEnum enum to its proto representation.
func GkehubBetaMembershipInfrastructureTypeEnumToProto(e *beta.MembershipInfrastructureTypeEnum) betapb.GkehubBetaMembershipInfrastructureTypeEnum {
	if e == nil {
		return betapb.GkehubBetaMembershipInfrastructureTypeEnum(0)
	}
	if v, ok := betapb.GkehubBetaMembershipInfrastructureTypeEnum_value["MembershipInfrastructureTypeEnum"+string(*e)]; ok {
		return betapb.GkehubBetaMembershipInfrastructureTypeEnum(v)
	}
	return betapb.GkehubBetaMembershipInfrastructureTypeEnum(0)
}

// MembershipEndpointToProto converts a MembershipEndpoint object to its proto representation.
func GkehubBetaMembershipEndpointToProto(o *beta.MembershipEndpoint) *betapb.GkehubBetaMembershipEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpoint{}
	p.SetGkeCluster(GkehubBetaMembershipEndpointGkeClusterToProto(o.GkeCluster))
	p.SetKubernetesMetadata(GkehubBetaMembershipEndpointKubernetesMetadataToProto(o.KubernetesMetadata))
	p.SetKubernetesResource(GkehubBetaMembershipEndpointKubernetesResourceToProto(o.KubernetesResource))
	return p
}

// MembershipEndpointGkeClusterToProto converts a MembershipEndpointGkeCluster object to its proto representation.
func GkehubBetaMembershipEndpointGkeClusterToProto(o *beta.MembershipEndpointGkeCluster) *betapb.GkehubBetaMembershipEndpointGkeCluster {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointGkeCluster{}
	p.SetResourceLink(dcl.ValueOrEmptyString(o.ResourceLink))
	return p
}

// MembershipEndpointKubernetesMetadataToProto converts a MembershipEndpointKubernetesMetadata object to its proto representation.
func GkehubBetaMembershipEndpointKubernetesMetadataToProto(o *beta.MembershipEndpointKubernetesMetadata) *betapb.GkehubBetaMembershipEndpointKubernetesMetadata {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesMetadata{}
	p.SetKubernetesApiServerVersion(dcl.ValueOrEmptyString(o.KubernetesApiServerVersion))
	p.SetNodeProviderId(dcl.ValueOrEmptyString(o.NodeProviderId))
	p.SetNodeCount(dcl.ValueOrEmptyInt64(o.NodeCount))
	p.SetVcpuCount(dcl.ValueOrEmptyInt64(o.VcpuCount))
	p.SetMemoryMb(dcl.ValueOrEmptyInt64(o.MemoryMb))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// MembershipEndpointKubernetesResourceToProto converts a MembershipEndpointKubernetesResource object to its proto representation.
func GkehubBetaMembershipEndpointKubernetesResourceToProto(o *beta.MembershipEndpointKubernetesResource) *betapb.GkehubBetaMembershipEndpointKubernetesResource {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesResource{}
	p.SetMembershipCrManifest(dcl.ValueOrEmptyString(o.MembershipCrManifest))
	p.SetResourceOptions(GkehubBetaMembershipEndpointKubernetesResourceResourceOptionsToProto(o.ResourceOptions))
	sMembershipResources := make([]*betapb.GkehubBetaMembershipEndpointKubernetesResourceMembershipResources, len(o.MembershipResources))
	for i, r := range o.MembershipResources {
		sMembershipResources[i] = GkehubBetaMembershipEndpointKubernetesResourceMembershipResourcesToProto(&r)
	}
	p.SetMembershipResources(sMembershipResources)
	sConnectResources := make([]*betapb.GkehubBetaMembershipEndpointKubernetesResourceConnectResources, len(o.ConnectResources))
	for i, r := range o.ConnectResources {
		sConnectResources[i] = GkehubBetaMembershipEndpointKubernetesResourceConnectResourcesToProto(&r)
	}
	p.SetConnectResources(sConnectResources)
	return p
}

// MembershipEndpointKubernetesResourceMembershipResourcesToProto converts a MembershipEndpointKubernetesResourceMembershipResources object to its proto representation.
func GkehubBetaMembershipEndpointKubernetesResourceMembershipResourcesToProto(o *beta.MembershipEndpointKubernetesResourceMembershipResources) *betapb.GkehubBetaMembershipEndpointKubernetesResourceMembershipResources {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesResourceMembershipResources{}
	p.SetManifest(dcl.ValueOrEmptyString(o.Manifest))
	p.SetClusterScoped(dcl.ValueOrEmptyBool(o.ClusterScoped))
	return p
}

// MembershipEndpointKubernetesResourceConnectResourcesToProto converts a MembershipEndpointKubernetesResourceConnectResources object to its proto representation.
func GkehubBetaMembershipEndpointKubernetesResourceConnectResourcesToProto(o *beta.MembershipEndpointKubernetesResourceConnectResources) *betapb.GkehubBetaMembershipEndpointKubernetesResourceConnectResources {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesResourceConnectResources{}
	p.SetManifest(dcl.ValueOrEmptyString(o.Manifest))
	p.SetClusterScoped(dcl.ValueOrEmptyBool(o.ClusterScoped))
	return p
}

// MembershipEndpointKubernetesResourceResourceOptionsToProto converts a MembershipEndpointKubernetesResourceResourceOptions object to its proto representation.
func GkehubBetaMembershipEndpointKubernetesResourceResourceOptionsToProto(o *beta.MembershipEndpointKubernetesResourceResourceOptions) *betapb.GkehubBetaMembershipEndpointKubernetesResourceResourceOptions {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipEndpointKubernetesResourceResourceOptions{}
	p.SetConnectVersion(dcl.ValueOrEmptyString(o.ConnectVersion))
	p.SetV1Beta1Crd(dcl.ValueOrEmptyBool(o.V1Beta1Crd))
	return p
}

// MembershipStateToProto converts a MembershipState object to its proto representation.
func GkehubBetaMembershipStateToProto(o *beta.MembershipState) *betapb.GkehubBetaMembershipState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipState{}
	p.SetCode(GkehubBetaMembershipStateCodeEnumToProto(o.Code))
	return p
}

// MembershipAuthorityToProto converts a MembershipAuthority object to its proto representation.
func GkehubBetaMembershipAuthorityToProto(o *beta.MembershipAuthority) *betapb.GkehubBetaMembershipAuthority {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaMembershipAuthority{}
	p.SetIssuer(dcl.ValueOrEmptyString(o.Issuer))
	p.SetWorkloadIdentityPool(dcl.ValueOrEmptyString(o.WorkloadIdentityPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// MembershipToProto converts a Membership resource to its proto representation.
func MembershipToProto(resource *beta.Membership) *betapb.GkehubBetaMembership {
	p := &betapb.GkehubBetaMembership{}
	p.SetEndpoint(GkehubBetaMembershipEndpointToProto(resource.Endpoint))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(GkehubBetaMembershipStateToProto(resource.State))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetExternalId(dcl.ValueOrEmptyString(resource.ExternalId))
	p.SetLastConnectionTime(dcl.ValueOrEmptyString(resource.LastConnectionTime))
	p.SetUniqueId(dcl.ValueOrEmptyString(resource.UniqueId))
	p.SetAuthority(GkehubBetaMembershipAuthorityToProto(resource.Authority))
	p.SetInfrastructureType(GkehubBetaMembershipInfrastructureTypeEnumToProto(resource.InfrastructureType))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) applyMembership(ctx context.Context, c *beta.Client, request *betapb.ApplyGkehubBetaMembershipRequest) (*betapb.GkehubBetaMembership, error) {
	p := ProtoToMembership(request.GetResource())
	res, err := c.ApplyMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MembershipToProto(res)
	return r, nil
}

// applyGkehubBetaMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) ApplyGkehubBetaMembership(ctx context.Context, request *betapb.ApplyGkehubBetaMembershipRequest) (*betapb.GkehubBetaMembership, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMembership(ctx, cl, request)
}

// DeleteMembership handles the gRPC request by passing it to the underlying Membership Delete() method.
func (s *MembershipServer) DeleteGkehubBetaMembership(ctx context.Context, request *betapb.DeleteGkehubBetaMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMembership(ctx, ProtoToMembership(request.GetResource()))

}

// ListGkehubBetaMembership handles the gRPC request by passing it to the underlying MembershipList() method.
func (s *MembershipServer) ListGkehubBetaMembership(ctx context.Context, request *betapb.ListGkehubBetaMembershipRequest) (*betapb.ListGkehubBetaMembershipResponse, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMembership(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkehubBetaMembership
	for _, r := range resources.Items {
		rp := MembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListGkehubBetaMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMembership(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
