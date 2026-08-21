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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/alpha/gkehub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
)

// MembershipServer implements the gRPC interface for Membership.
type MembershipServer struct{}

// ProtoToMembershipStateCodeEnum converts a MembershipStateCodeEnum enum from its proto representation.
func ProtoToGkehubAlphaMembershipStateCodeEnum(e alphapb.GkehubAlphaMembershipStateCodeEnum) *alpha.MembershipStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaMembershipStateCodeEnum_name[int32(e)]; ok {
		e := alpha.MembershipStateCodeEnum(n[len("GkehubAlphaMembershipStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipInfrastructureTypeEnum converts a MembershipInfrastructureTypeEnum enum from its proto representation.
func ProtoToGkehubAlphaMembershipInfrastructureTypeEnum(e alphapb.GkehubAlphaMembershipInfrastructureTypeEnum) *alpha.MembershipInfrastructureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaMembershipInfrastructureTypeEnum_name[int32(e)]; ok {
		e := alpha.MembershipInfrastructureTypeEnum(n[len("GkehubAlphaMembershipInfrastructureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipEndpoint converts a MembershipEndpoint object from its proto representation.
func ProtoToGkehubAlphaMembershipEndpoint(p *alphapb.GkehubAlphaMembershipEndpoint) *alpha.MembershipEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpoint{
		GkeCluster:         ProtoToGkehubAlphaMembershipEndpointGkeCluster(p.GetGkeCluster()),
		KubernetesMetadata: ProtoToGkehubAlphaMembershipEndpointKubernetesMetadata(p.GetKubernetesMetadata()),
		KubernetesResource: ProtoToGkehubAlphaMembershipEndpointKubernetesResource(p.GetKubernetesResource()),
	}
	return obj
}

// ProtoToMembershipEndpointGkeCluster converts a MembershipEndpointGkeCluster object from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointGkeCluster(p *alphapb.GkehubAlphaMembershipEndpointGkeCluster) *alpha.MembershipEndpointGkeCluster {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointGkeCluster{
		ResourceLink: dcl.StringOrNil(p.GetResourceLink()),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesMetadata converts a MembershipEndpointKubernetesMetadata object from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesMetadata(p *alphapb.GkehubAlphaMembershipEndpointKubernetesMetadata) *alpha.MembershipEndpointKubernetesMetadata {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesMetadata{
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
func ProtoToGkehubAlphaMembershipEndpointKubernetesResource(p *alphapb.GkehubAlphaMembershipEndpointKubernetesResource) *alpha.MembershipEndpointKubernetesResource {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesResource{
		MembershipCrManifest: dcl.StringOrNil(p.GetMembershipCrManifest()),
		ResourceOptions:      ProtoToGkehubAlphaMembershipEndpointKubernetesResourceResourceOptions(p.GetResourceOptions()),
	}
	for _, r := range p.GetMembershipResources() {
		obj.MembershipResources = append(obj.MembershipResources, *ProtoToGkehubAlphaMembershipEndpointKubernetesResourceMembershipResources(r))
	}
	for _, r := range p.GetConnectResources() {
		obj.ConnectResources = append(obj.ConnectResources, *ProtoToGkehubAlphaMembershipEndpointKubernetesResourceConnectResources(r))
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceMembershipResources converts a MembershipEndpointKubernetesResourceMembershipResources object from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesResourceMembershipResources(p *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceMembershipResources) *alpha.MembershipEndpointKubernetesResourceMembershipResources {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesResourceMembershipResources{
		Manifest:      dcl.StringOrNil(p.GetManifest()),
		ClusterScoped: dcl.Bool(p.GetClusterScoped()),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceConnectResources converts a MembershipEndpointKubernetesResourceConnectResources object from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesResourceConnectResources(p *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceConnectResources) *alpha.MembershipEndpointKubernetesResourceConnectResources {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesResourceConnectResources{
		Manifest:      dcl.StringOrNil(p.GetManifest()),
		ClusterScoped: dcl.Bool(p.GetClusterScoped()),
	}
	return obj
}

// ProtoToMembershipEndpointKubernetesResourceResourceOptions converts a MembershipEndpointKubernetesResourceResourceOptions object from its proto representation.
func ProtoToGkehubAlphaMembershipEndpointKubernetesResourceResourceOptions(p *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceResourceOptions) *alpha.MembershipEndpointKubernetesResourceResourceOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipEndpointKubernetesResourceResourceOptions{
		ConnectVersion: dcl.StringOrNil(p.GetConnectVersion()),
		V1Beta1Crd:     dcl.Bool(p.GetV1Beta1Crd()),
	}
	return obj
}

// ProtoToMembershipState converts a MembershipState object from its proto representation.
func ProtoToGkehubAlphaMembershipState(p *alphapb.GkehubAlphaMembershipState) *alpha.MembershipState {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipState{
		Code: ProtoToGkehubAlphaMembershipStateCodeEnum(p.GetCode()),
	}
	return obj
}

// ProtoToMembershipAuthority converts a MembershipAuthority object from its proto representation.
func ProtoToGkehubAlphaMembershipAuthority(p *alphapb.GkehubAlphaMembershipAuthority) *alpha.MembershipAuthority {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipAuthority{
		Issuer:               dcl.StringOrNil(p.GetIssuer()),
		WorkloadIdentityPool: dcl.StringOrNil(p.GetWorkloadIdentityPool()),
		IdentityProvider:     dcl.StringOrNil(p.GetIdentityProvider()),
	}
	return obj
}

// ProtoToMembership converts a Membership resource from its proto representation.
func ProtoToMembership(p *alphapb.GkehubAlphaMembership) *alpha.Membership {
	obj := &alpha.Membership{
		Endpoint:           ProtoToGkehubAlphaMembershipEndpoint(p.GetEndpoint()),
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		State:              ProtoToGkehubAlphaMembershipState(p.GetState()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:         dcl.StringOrNil(p.GetDeleteTime()),
		ExternalId:         dcl.StringOrNil(p.GetExternalId()),
		LastConnectionTime: dcl.StringOrNil(p.GetLastConnectionTime()),
		UniqueId:           dcl.StringOrNil(p.GetUniqueId()),
		Authority:          ProtoToGkehubAlphaMembershipAuthority(p.GetAuthority()),
		InfrastructureType: ProtoToGkehubAlphaMembershipInfrastructureTypeEnum(p.GetInfrastructureType()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// MembershipStateCodeEnumToProto converts a MembershipStateCodeEnum enum to its proto representation.
func GkehubAlphaMembershipStateCodeEnumToProto(e *alpha.MembershipStateCodeEnum) alphapb.GkehubAlphaMembershipStateCodeEnum {
	if e == nil {
		return alphapb.GkehubAlphaMembershipStateCodeEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaMembershipStateCodeEnum_value["MembershipStateCodeEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaMembershipStateCodeEnum(v)
	}
	return alphapb.GkehubAlphaMembershipStateCodeEnum(0)
}

// MembershipInfrastructureTypeEnumToProto converts a MembershipInfrastructureTypeEnum enum to its proto representation.
func GkehubAlphaMembershipInfrastructureTypeEnumToProto(e *alpha.MembershipInfrastructureTypeEnum) alphapb.GkehubAlphaMembershipInfrastructureTypeEnum {
	if e == nil {
		return alphapb.GkehubAlphaMembershipInfrastructureTypeEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaMembershipInfrastructureTypeEnum_value["MembershipInfrastructureTypeEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaMembershipInfrastructureTypeEnum(v)
	}
	return alphapb.GkehubAlphaMembershipInfrastructureTypeEnum(0)
}

// MembershipEndpointToProto converts a MembershipEndpoint object to its proto representation.
func GkehubAlphaMembershipEndpointToProto(o *alpha.MembershipEndpoint) *alphapb.GkehubAlphaMembershipEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpoint{}
	p.SetGkeCluster(GkehubAlphaMembershipEndpointGkeClusterToProto(o.GkeCluster))
	p.SetKubernetesMetadata(GkehubAlphaMembershipEndpointKubernetesMetadataToProto(o.KubernetesMetadata))
	p.SetKubernetesResource(GkehubAlphaMembershipEndpointKubernetesResourceToProto(o.KubernetesResource))
	return p
}

// MembershipEndpointGkeClusterToProto converts a MembershipEndpointGkeCluster object to its proto representation.
func GkehubAlphaMembershipEndpointGkeClusterToProto(o *alpha.MembershipEndpointGkeCluster) *alphapb.GkehubAlphaMembershipEndpointGkeCluster {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointGkeCluster{}
	p.SetResourceLink(dcl.ValueOrEmptyString(o.ResourceLink))
	return p
}

// MembershipEndpointKubernetesMetadataToProto converts a MembershipEndpointKubernetesMetadata object to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesMetadataToProto(o *alpha.MembershipEndpointKubernetesMetadata) *alphapb.GkehubAlphaMembershipEndpointKubernetesMetadata {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesMetadata{}
	p.SetKubernetesApiServerVersion(dcl.ValueOrEmptyString(o.KubernetesApiServerVersion))
	p.SetNodeProviderId(dcl.ValueOrEmptyString(o.NodeProviderId))
	p.SetNodeCount(dcl.ValueOrEmptyInt64(o.NodeCount))
	p.SetVcpuCount(dcl.ValueOrEmptyInt64(o.VcpuCount))
	p.SetMemoryMb(dcl.ValueOrEmptyInt64(o.MemoryMb))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// MembershipEndpointKubernetesResourceToProto converts a MembershipEndpointKubernetesResource object to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesResourceToProto(o *alpha.MembershipEndpointKubernetesResource) *alphapb.GkehubAlphaMembershipEndpointKubernetesResource {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesResource{}
	p.SetMembershipCrManifest(dcl.ValueOrEmptyString(o.MembershipCrManifest))
	p.SetResourceOptions(GkehubAlphaMembershipEndpointKubernetesResourceResourceOptionsToProto(o.ResourceOptions))
	sMembershipResources := make([]*alphapb.GkehubAlphaMembershipEndpointKubernetesResourceMembershipResources, len(o.MembershipResources))
	for i, r := range o.MembershipResources {
		sMembershipResources[i] = GkehubAlphaMembershipEndpointKubernetesResourceMembershipResourcesToProto(&r)
	}
	p.SetMembershipResources(sMembershipResources)
	sConnectResources := make([]*alphapb.GkehubAlphaMembershipEndpointKubernetesResourceConnectResources, len(o.ConnectResources))
	for i, r := range o.ConnectResources {
		sConnectResources[i] = GkehubAlphaMembershipEndpointKubernetesResourceConnectResourcesToProto(&r)
	}
	p.SetConnectResources(sConnectResources)
	return p
}

// MembershipEndpointKubernetesResourceMembershipResourcesToProto converts a MembershipEndpointKubernetesResourceMembershipResources object to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesResourceMembershipResourcesToProto(o *alpha.MembershipEndpointKubernetesResourceMembershipResources) *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceMembershipResources {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesResourceMembershipResources{}
	p.SetManifest(dcl.ValueOrEmptyString(o.Manifest))
	p.SetClusterScoped(dcl.ValueOrEmptyBool(o.ClusterScoped))
	return p
}

// MembershipEndpointKubernetesResourceConnectResourcesToProto converts a MembershipEndpointKubernetesResourceConnectResources object to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesResourceConnectResourcesToProto(o *alpha.MembershipEndpointKubernetesResourceConnectResources) *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceConnectResources {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesResourceConnectResources{}
	p.SetManifest(dcl.ValueOrEmptyString(o.Manifest))
	p.SetClusterScoped(dcl.ValueOrEmptyBool(o.ClusterScoped))
	return p
}

// MembershipEndpointKubernetesResourceResourceOptionsToProto converts a MembershipEndpointKubernetesResourceResourceOptions object to its proto representation.
func GkehubAlphaMembershipEndpointKubernetesResourceResourceOptionsToProto(o *alpha.MembershipEndpointKubernetesResourceResourceOptions) *alphapb.GkehubAlphaMembershipEndpointKubernetesResourceResourceOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipEndpointKubernetesResourceResourceOptions{}
	p.SetConnectVersion(dcl.ValueOrEmptyString(o.ConnectVersion))
	p.SetV1Beta1Crd(dcl.ValueOrEmptyBool(o.V1Beta1Crd))
	return p
}

// MembershipStateToProto converts a MembershipState object to its proto representation.
func GkehubAlphaMembershipStateToProto(o *alpha.MembershipState) *alphapb.GkehubAlphaMembershipState {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipState{}
	p.SetCode(GkehubAlphaMembershipStateCodeEnumToProto(o.Code))
	return p
}

// MembershipAuthorityToProto converts a MembershipAuthority object to its proto representation.
func GkehubAlphaMembershipAuthorityToProto(o *alpha.MembershipAuthority) *alphapb.GkehubAlphaMembershipAuthority {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaMembershipAuthority{}
	p.SetIssuer(dcl.ValueOrEmptyString(o.Issuer))
	p.SetWorkloadIdentityPool(dcl.ValueOrEmptyString(o.WorkloadIdentityPool))
	p.SetIdentityProvider(dcl.ValueOrEmptyString(o.IdentityProvider))
	return p
}

// MembershipToProto converts a Membership resource to its proto representation.
func MembershipToProto(resource *alpha.Membership) *alphapb.GkehubAlphaMembership {
	p := &alphapb.GkehubAlphaMembership{}
	p.SetEndpoint(GkehubAlphaMembershipEndpointToProto(resource.Endpoint))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(GkehubAlphaMembershipStateToProto(resource.State))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetExternalId(dcl.ValueOrEmptyString(resource.ExternalId))
	p.SetLastConnectionTime(dcl.ValueOrEmptyString(resource.LastConnectionTime))
	p.SetUniqueId(dcl.ValueOrEmptyString(resource.UniqueId))
	p.SetAuthority(GkehubAlphaMembershipAuthorityToProto(resource.Authority))
	p.SetInfrastructureType(GkehubAlphaMembershipInfrastructureTypeEnumToProto(resource.InfrastructureType))
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
func (s *MembershipServer) applyMembership(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkehubAlphaMembershipRequest) (*alphapb.GkehubAlphaMembership, error) {
	p := ProtoToMembership(request.GetResource())
	res, err := c.ApplyMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MembershipToProto(res)
	return r, nil
}

// applyGkehubAlphaMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) ApplyGkehubAlphaMembership(ctx context.Context, request *alphapb.ApplyGkehubAlphaMembershipRequest) (*alphapb.GkehubAlphaMembership, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMembership(ctx, cl, request)
}

// DeleteMembership handles the gRPC request by passing it to the underlying Membership Delete() method.
func (s *MembershipServer) DeleteGkehubAlphaMembership(ctx context.Context, request *alphapb.DeleteGkehubAlphaMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMembership(ctx, ProtoToMembership(request.GetResource()))

}

// ListGkehubAlphaMembership handles the gRPC request by passing it to the underlying MembershipList() method.
func (s *MembershipServer) ListGkehubAlphaMembership(ctx context.Context, request *alphapb.ListGkehubAlphaMembershipRequest) (*alphapb.ListGkehubAlphaMembershipResponse, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMembership(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkehubAlphaMembership
	for _, r := range resources.Items {
		rp := MembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListGkehubAlphaMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMembership(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
