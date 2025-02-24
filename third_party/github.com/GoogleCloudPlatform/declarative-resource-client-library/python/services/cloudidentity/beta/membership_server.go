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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudidentity/beta/cloudidentity_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity/beta"
)

// MembershipServer implements the gRPC interface for Membership.
type MembershipServer struct{}

// ProtoToMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum enum from its proto representation.
func ProtoToCloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(e betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum) *beta.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum_name[int32(e)]; ok {
		e := beta.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(n[len("CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipTypeEnum converts a MembershipTypeEnum enum from its proto representation.
func ProtoToCloudidentityBetaMembershipTypeEnum(e betapb.CloudidentityBetaMembershipTypeEnum) *beta.MembershipTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudidentityBetaMembershipTypeEnum_name[int32(e)]; ok {
		e := beta.MembershipTypeEnum(n[len("CloudidentityBetaMembershipTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipDeliverySettingEnum converts a MembershipDeliverySettingEnum enum from its proto representation.
func ProtoToCloudidentityBetaMembershipDeliverySettingEnum(e betapb.CloudidentityBetaMembershipDeliverySettingEnum) *beta.MembershipDeliverySettingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudidentityBetaMembershipDeliverySettingEnum_name[int32(e)]; ok {
		e := beta.MembershipDeliverySettingEnum(n[len("CloudidentityBetaMembershipDeliverySettingEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipPreferredMemberKey converts a MembershipPreferredMemberKey object from its proto representation.
func ProtoToCloudidentityBetaMembershipPreferredMemberKey(p *betapb.CloudidentityBetaMembershipPreferredMemberKey) *beta.MembershipPreferredMemberKey {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipPreferredMemberKey{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToMembershipRoles converts a MembershipRoles object from its proto representation.
func ProtoToCloudidentityBetaMembershipRoles(p *betapb.CloudidentityBetaMembershipRoles) *beta.MembershipRoles {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipRoles{
		Name:                   dcl.StringOrNil(p.GetName()),
		ExpiryDetail:           ProtoToCloudidentityBetaMembershipRolesExpiryDetail(p.GetExpiryDetail()),
		RestrictionEvaluations: ProtoToCloudidentityBetaMembershipRolesRestrictionEvaluations(p.GetRestrictionEvaluations()),
	}
	return obj
}

// ProtoToMembershipRolesExpiryDetail converts a MembershipRolesExpiryDetail object from its proto representation.
func ProtoToCloudidentityBetaMembershipRolesExpiryDetail(p *betapb.CloudidentityBetaMembershipRolesExpiryDetail) *beta.MembershipRolesExpiryDetail {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipRolesExpiryDetail{
		ExpireTime: dcl.StringOrNil(p.GetExpireTime()),
	}
	return obj
}

// ProtoToMembershipRolesRestrictionEvaluations converts a MembershipRolesRestrictionEvaluations object from its proto representation.
func ProtoToCloudidentityBetaMembershipRolesRestrictionEvaluations(p *betapb.CloudidentityBetaMembershipRolesRestrictionEvaluations) *beta.MembershipRolesRestrictionEvaluations {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipRolesRestrictionEvaluations{
		MemberRestrictionEvaluation: ProtoToCloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(p.GetMemberRestrictionEvaluation()),
	}
	return obj
}

// ProtoToMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation object from its proto representation.
func ProtoToCloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(p *betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) *beta.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{
		State: ProtoToCloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(p.GetState()),
	}
	return obj
}

// ProtoToMembershipDisplayName converts a MembershipDisplayName object from its proto representation.
func ProtoToCloudidentityBetaMembershipDisplayName(p *betapb.CloudidentityBetaMembershipDisplayName) *beta.MembershipDisplayName {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipDisplayName{
		GivenName:  dcl.StringOrNil(p.GetGivenName()),
		FamilyName: dcl.StringOrNil(p.GetFamilyName()),
		FullName:   dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToMembershipMemberKey converts a MembershipMemberKey object from its proto representation.
func ProtoToCloudidentityBetaMembershipMemberKey(p *betapb.CloudidentityBetaMembershipMemberKey) *beta.MembershipMemberKey {
	if p == nil {
		return nil
	}
	obj := &beta.MembershipMemberKey{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToMembership converts a Membership resource from its proto representation.
func ProtoToMembership(p *betapb.CloudidentityBetaMembership) *beta.Membership {
	obj := &beta.Membership{
		Name:               dcl.StringOrNil(p.GetName()),
		PreferredMemberKey: ProtoToCloudidentityBetaMembershipPreferredMemberKey(p.GetPreferredMemberKey()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Type:               ProtoToCloudidentityBetaMembershipTypeEnum(p.GetType()),
		DeliverySetting:    ProtoToCloudidentityBetaMembershipDeliverySettingEnum(p.GetDeliverySetting()),
		DisplayName:        ProtoToCloudidentityBetaMembershipDisplayName(p.GetDisplayName()),
		MemberKey:          ProtoToCloudidentityBetaMembershipMemberKey(p.GetMemberKey()),
		Group:              dcl.StringOrNil(p.GetGroup()),
	}
	for _, r := range p.GetRoles() {
		obj.Roles = append(obj.Roles, *ProtoToCloudidentityBetaMembershipRoles(r))
	}
	return obj
}

// MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum enum to its proto representation.
func CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto(e *beta.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum) betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	if e == nil {
		return betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(0)
	}
	if v, ok := betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum_value["MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum"+string(*e)]; ok {
		return betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(v)
	}
	return betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(0)
}

// MembershipTypeEnumToProto converts a MembershipTypeEnum enum to its proto representation.
func CloudidentityBetaMembershipTypeEnumToProto(e *beta.MembershipTypeEnum) betapb.CloudidentityBetaMembershipTypeEnum {
	if e == nil {
		return betapb.CloudidentityBetaMembershipTypeEnum(0)
	}
	if v, ok := betapb.CloudidentityBetaMembershipTypeEnum_value["MembershipTypeEnum"+string(*e)]; ok {
		return betapb.CloudidentityBetaMembershipTypeEnum(v)
	}
	return betapb.CloudidentityBetaMembershipTypeEnum(0)
}

// MembershipDeliverySettingEnumToProto converts a MembershipDeliverySettingEnum enum to its proto representation.
func CloudidentityBetaMembershipDeliverySettingEnumToProto(e *beta.MembershipDeliverySettingEnum) betapb.CloudidentityBetaMembershipDeliverySettingEnum {
	if e == nil {
		return betapb.CloudidentityBetaMembershipDeliverySettingEnum(0)
	}
	if v, ok := betapb.CloudidentityBetaMembershipDeliverySettingEnum_value["MembershipDeliverySettingEnum"+string(*e)]; ok {
		return betapb.CloudidentityBetaMembershipDeliverySettingEnum(v)
	}
	return betapb.CloudidentityBetaMembershipDeliverySettingEnum(0)
}

// MembershipPreferredMemberKeyToProto converts a MembershipPreferredMemberKey object to its proto representation.
func CloudidentityBetaMembershipPreferredMemberKeyToProto(o *beta.MembershipPreferredMemberKey) *betapb.CloudidentityBetaMembershipPreferredMemberKey {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaMembershipPreferredMemberKey{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// MembershipRolesToProto converts a MembershipRoles object to its proto representation.
func CloudidentityBetaMembershipRolesToProto(o *beta.MembershipRoles) *betapb.CloudidentityBetaMembershipRoles {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaMembershipRoles{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetExpiryDetail(CloudidentityBetaMembershipRolesExpiryDetailToProto(o.ExpiryDetail))
	p.SetRestrictionEvaluations(CloudidentityBetaMembershipRolesRestrictionEvaluationsToProto(o.RestrictionEvaluations))
	return p
}

// MembershipRolesExpiryDetailToProto converts a MembershipRolesExpiryDetail object to its proto representation.
func CloudidentityBetaMembershipRolesExpiryDetailToProto(o *beta.MembershipRolesExpiryDetail) *betapb.CloudidentityBetaMembershipRolesExpiryDetail {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaMembershipRolesExpiryDetail{}
	p.SetExpireTime(dcl.ValueOrEmptyString(o.ExpireTime))
	return p
}

// MembershipRolesRestrictionEvaluationsToProto converts a MembershipRolesRestrictionEvaluations object to its proto representation.
func CloudidentityBetaMembershipRolesRestrictionEvaluationsToProto(o *beta.MembershipRolesRestrictionEvaluations) *betapb.CloudidentityBetaMembershipRolesRestrictionEvaluations {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaMembershipRolesRestrictionEvaluations{}
	p.SetMemberRestrictionEvaluation(CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto(o.MemberRestrictionEvaluation))
	return p
}

// MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation object to its proto representation.
func CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto(o *beta.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) *betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	p.SetState(CloudidentityBetaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto(o.State))
	return p
}

// MembershipDisplayNameToProto converts a MembershipDisplayName object to its proto representation.
func CloudidentityBetaMembershipDisplayNameToProto(o *beta.MembershipDisplayName) *betapb.CloudidentityBetaMembershipDisplayName {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaMembershipDisplayName{}
	p.SetGivenName(dcl.ValueOrEmptyString(o.GivenName))
	p.SetFamilyName(dcl.ValueOrEmptyString(o.FamilyName))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// MembershipMemberKeyToProto converts a MembershipMemberKey object to its proto representation.
func CloudidentityBetaMembershipMemberKeyToProto(o *beta.MembershipMemberKey) *betapb.CloudidentityBetaMembershipMemberKey {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaMembershipMemberKey{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// MembershipToProto converts a Membership resource to its proto representation.
func MembershipToProto(resource *beta.Membership) *betapb.CloudidentityBetaMembership {
	p := &betapb.CloudidentityBetaMembership{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPreferredMemberKey(CloudidentityBetaMembershipPreferredMemberKeyToProto(resource.PreferredMemberKey))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetType(CloudidentityBetaMembershipTypeEnumToProto(resource.Type))
	p.SetDeliverySetting(CloudidentityBetaMembershipDeliverySettingEnumToProto(resource.DeliverySetting))
	p.SetDisplayName(CloudidentityBetaMembershipDisplayNameToProto(resource.DisplayName))
	p.SetMemberKey(CloudidentityBetaMembershipMemberKeyToProto(resource.MemberKey))
	p.SetGroup(dcl.ValueOrEmptyString(resource.Group))
	sRoles := make([]*betapb.CloudidentityBetaMembershipRoles, len(resource.Roles))
	for i, r := range resource.Roles {
		sRoles[i] = CloudidentityBetaMembershipRolesToProto(&r)
	}
	p.SetRoles(sRoles)

	return p
}

// applyMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) applyMembership(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudidentityBetaMembershipRequest) (*betapb.CloudidentityBetaMembership, error) {
	p := ProtoToMembership(request.GetResource())
	res, err := c.ApplyMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MembershipToProto(res)
	return r, nil
}

// applyCloudidentityBetaMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) ApplyCloudidentityBetaMembership(ctx context.Context, request *betapb.ApplyCloudidentityBetaMembershipRequest) (*betapb.CloudidentityBetaMembership, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMembership(ctx, cl, request)
}

// DeleteMembership handles the gRPC request by passing it to the underlying Membership Delete() method.
func (s *MembershipServer) DeleteCloudidentityBetaMembership(ctx context.Context, request *betapb.DeleteCloudidentityBetaMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMembership(ctx, ProtoToMembership(request.GetResource()))

}

// ListCloudidentityBetaMembership handles the gRPC request by passing it to the underlying MembershipList() method.
func (s *MembershipServer) ListCloudidentityBetaMembership(ctx context.Context, request *betapb.ListCloudidentityBetaMembershipRequest) (*betapb.ListCloudidentityBetaMembershipResponse, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMembership(ctx, request.GetGroup())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudidentityBetaMembership
	for _, r := range resources.Items {
		rp := MembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudidentityBetaMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMembership(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
