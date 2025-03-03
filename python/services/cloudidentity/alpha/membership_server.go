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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudidentity/alpha/cloudidentity_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity/alpha"
)

// MembershipServer implements the gRPC interface for Membership.
type MembershipServer struct{}

// ProtoToMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum enum from its proto representation.
func ProtoToCloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(e alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum) *alpha.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum_name[int32(e)]; ok {
		e := alpha.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(n[len("CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipTypeEnum converts a MembershipTypeEnum enum from its proto representation.
func ProtoToCloudidentityAlphaMembershipTypeEnum(e alphapb.CloudidentityAlphaMembershipTypeEnum) *alpha.MembershipTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudidentityAlphaMembershipTypeEnum_name[int32(e)]; ok {
		e := alpha.MembershipTypeEnum(n[len("CloudidentityAlphaMembershipTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipDeliverySettingEnum converts a MembershipDeliverySettingEnum enum from its proto representation.
func ProtoToCloudidentityAlphaMembershipDeliverySettingEnum(e alphapb.CloudidentityAlphaMembershipDeliverySettingEnum) *alpha.MembershipDeliverySettingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudidentityAlphaMembershipDeliverySettingEnum_name[int32(e)]; ok {
		e := alpha.MembershipDeliverySettingEnum(n[len("CloudidentityAlphaMembershipDeliverySettingEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipPreferredMemberKey converts a MembershipPreferredMemberKey object from its proto representation.
func ProtoToCloudidentityAlphaMembershipPreferredMemberKey(p *alphapb.CloudidentityAlphaMembershipPreferredMemberKey) *alpha.MembershipPreferredMemberKey {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipPreferredMemberKey{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToMembershipRoles converts a MembershipRoles object from its proto representation.
func ProtoToCloudidentityAlphaMembershipRoles(p *alphapb.CloudidentityAlphaMembershipRoles) *alpha.MembershipRoles {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipRoles{
		Name:                   dcl.StringOrNil(p.GetName()),
		ExpiryDetail:           ProtoToCloudidentityAlphaMembershipRolesExpiryDetail(p.GetExpiryDetail()),
		RestrictionEvaluations: ProtoToCloudidentityAlphaMembershipRolesRestrictionEvaluations(p.GetRestrictionEvaluations()),
	}
	return obj
}

// ProtoToMembershipRolesExpiryDetail converts a MembershipRolesExpiryDetail object from its proto representation.
func ProtoToCloudidentityAlphaMembershipRolesExpiryDetail(p *alphapb.CloudidentityAlphaMembershipRolesExpiryDetail) *alpha.MembershipRolesExpiryDetail {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipRolesExpiryDetail{
		ExpireTime: dcl.StringOrNil(p.GetExpireTime()),
	}
	return obj
}

// ProtoToMembershipRolesRestrictionEvaluations converts a MembershipRolesRestrictionEvaluations object from its proto representation.
func ProtoToCloudidentityAlphaMembershipRolesRestrictionEvaluations(p *alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluations) *alpha.MembershipRolesRestrictionEvaluations {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipRolesRestrictionEvaluations{
		MemberRestrictionEvaluation: ProtoToCloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(p.GetMemberRestrictionEvaluation()),
	}
	return obj
}

// ProtoToMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation object from its proto representation.
func ProtoToCloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(p *alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) *alpha.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{
		State: ProtoToCloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(p.GetState()),
	}
	return obj
}

// ProtoToMembershipDisplayName converts a MembershipDisplayName object from its proto representation.
func ProtoToCloudidentityAlphaMembershipDisplayName(p *alphapb.CloudidentityAlphaMembershipDisplayName) *alpha.MembershipDisplayName {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipDisplayName{
		GivenName:  dcl.StringOrNil(p.GetGivenName()),
		FamilyName: dcl.StringOrNil(p.GetFamilyName()),
		FullName:   dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToMembershipMemberKey converts a MembershipMemberKey object from its proto representation.
func ProtoToCloudidentityAlphaMembershipMemberKey(p *alphapb.CloudidentityAlphaMembershipMemberKey) *alpha.MembershipMemberKey {
	if p == nil {
		return nil
	}
	obj := &alpha.MembershipMemberKey{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToMembership converts a Membership resource from its proto representation.
func ProtoToMembership(p *alphapb.CloudidentityAlphaMembership) *alpha.Membership {
	obj := &alpha.Membership{
		Name:               dcl.StringOrNil(p.GetName()),
		PreferredMemberKey: ProtoToCloudidentityAlphaMembershipPreferredMemberKey(p.GetPreferredMemberKey()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Type:               ProtoToCloudidentityAlphaMembershipTypeEnum(p.GetType()),
		DeliverySetting:    ProtoToCloudidentityAlphaMembershipDeliverySettingEnum(p.GetDeliverySetting()),
		DisplayName:        ProtoToCloudidentityAlphaMembershipDisplayName(p.GetDisplayName()),
		MemberKey:          ProtoToCloudidentityAlphaMembershipMemberKey(p.GetMemberKey()),
		Group:              dcl.StringOrNil(p.GetGroup()),
	}
	for _, r := range p.GetRoles() {
		obj.Roles = append(obj.Roles, *ProtoToCloudidentityAlphaMembershipRoles(r))
	}
	return obj
}

// MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum enum to its proto representation.
func CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto(e *alpha.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum) alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	if e == nil {
		return alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(0)
	}
	if v, ok := alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum_value["MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum"+string(*e)]; ok {
		return alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(v)
	}
	return alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(0)
}

// MembershipTypeEnumToProto converts a MembershipTypeEnum enum to its proto representation.
func CloudidentityAlphaMembershipTypeEnumToProto(e *alpha.MembershipTypeEnum) alphapb.CloudidentityAlphaMembershipTypeEnum {
	if e == nil {
		return alphapb.CloudidentityAlphaMembershipTypeEnum(0)
	}
	if v, ok := alphapb.CloudidentityAlphaMembershipTypeEnum_value["MembershipTypeEnum"+string(*e)]; ok {
		return alphapb.CloudidentityAlphaMembershipTypeEnum(v)
	}
	return alphapb.CloudidentityAlphaMembershipTypeEnum(0)
}

// MembershipDeliverySettingEnumToProto converts a MembershipDeliverySettingEnum enum to its proto representation.
func CloudidentityAlphaMembershipDeliverySettingEnumToProto(e *alpha.MembershipDeliverySettingEnum) alphapb.CloudidentityAlphaMembershipDeliverySettingEnum {
	if e == nil {
		return alphapb.CloudidentityAlphaMembershipDeliverySettingEnum(0)
	}
	if v, ok := alphapb.CloudidentityAlphaMembershipDeliverySettingEnum_value["MembershipDeliverySettingEnum"+string(*e)]; ok {
		return alphapb.CloudidentityAlphaMembershipDeliverySettingEnum(v)
	}
	return alphapb.CloudidentityAlphaMembershipDeliverySettingEnum(0)
}

// MembershipPreferredMemberKeyToProto converts a MembershipPreferredMemberKey object to its proto representation.
func CloudidentityAlphaMembershipPreferredMemberKeyToProto(o *alpha.MembershipPreferredMemberKey) *alphapb.CloudidentityAlphaMembershipPreferredMemberKey {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaMembershipPreferredMemberKey{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// MembershipRolesToProto converts a MembershipRoles object to its proto representation.
func CloudidentityAlphaMembershipRolesToProto(o *alpha.MembershipRoles) *alphapb.CloudidentityAlphaMembershipRoles {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaMembershipRoles{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetExpiryDetail(CloudidentityAlphaMembershipRolesExpiryDetailToProto(o.ExpiryDetail))
	p.SetRestrictionEvaluations(CloudidentityAlphaMembershipRolesRestrictionEvaluationsToProto(o.RestrictionEvaluations))
	return p
}

// MembershipRolesExpiryDetailToProto converts a MembershipRolesExpiryDetail object to its proto representation.
func CloudidentityAlphaMembershipRolesExpiryDetailToProto(o *alpha.MembershipRolesExpiryDetail) *alphapb.CloudidentityAlphaMembershipRolesExpiryDetail {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaMembershipRolesExpiryDetail{}
	p.SetExpireTime(dcl.ValueOrEmptyString(o.ExpireTime))
	return p
}

// MembershipRolesRestrictionEvaluationsToProto converts a MembershipRolesRestrictionEvaluations object to its proto representation.
func CloudidentityAlphaMembershipRolesRestrictionEvaluationsToProto(o *alpha.MembershipRolesRestrictionEvaluations) *alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluations {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluations{}
	p.SetMemberRestrictionEvaluation(CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto(o.MemberRestrictionEvaluation))
	return p
}

// MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation object to its proto representation.
func CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto(o *alpha.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) *alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	p.SetState(CloudidentityAlphaMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto(o.State))
	return p
}

// MembershipDisplayNameToProto converts a MembershipDisplayName object to its proto representation.
func CloudidentityAlphaMembershipDisplayNameToProto(o *alpha.MembershipDisplayName) *alphapb.CloudidentityAlphaMembershipDisplayName {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaMembershipDisplayName{}
	p.SetGivenName(dcl.ValueOrEmptyString(o.GivenName))
	p.SetFamilyName(dcl.ValueOrEmptyString(o.FamilyName))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// MembershipMemberKeyToProto converts a MembershipMemberKey object to its proto representation.
func CloudidentityAlphaMembershipMemberKeyToProto(o *alpha.MembershipMemberKey) *alphapb.CloudidentityAlphaMembershipMemberKey {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaMembershipMemberKey{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// MembershipToProto converts a Membership resource to its proto representation.
func MembershipToProto(resource *alpha.Membership) *alphapb.CloudidentityAlphaMembership {
	p := &alphapb.CloudidentityAlphaMembership{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPreferredMemberKey(CloudidentityAlphaMembershipPreferredMemberKeyToProto(resource.PreferredMemberKey))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetType(CloudidentityAlphaMembershipTypeEnumToProto(resource.Type))
	p.SetDeliverySetting(CloudidentityAlphaMembershipDeliverySettingEnumToProto(resource.DeliverySetting))
	p.SetDisplayName(CloudidentityAlphaMembershipDisplayNameToProto(resource.DisplayName))
	p.SetMemberKey(CloudidentityAlphaMembershipMemberKeyToProto(resource.MemberKey))
	p.SetGroup(dcl.ValueOrEmptyString(resource.Group))
	sRoles := make([]*alphapb.CloudidentityAlphaMembershipRoles, len(resource.Roles))
	for i, r := range resource.Roles {
		sRoles[i] = CloudidentityAlphaMembershipRolesToProto(&r)
	}
	p.SetRoles(sRoles)

	return p
}

// applyMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) applyMembership(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudidentityAlphaMembershipRequest) (*alphapb.CloudidentityAlphaMembership, error) {
	p := ProtoToMembership(request.GetResource())
	res, err := c.ApplyMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MembershipToProto(res)
	return r, nil
}

// applyCloudidentityAlphaMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) ApplyCloudidentityAlphaMembership(ctx context.Context, request *alphapb.ApplyCloudidentityAlphaMembershipRequest) (*alphapb.CloudidentityAlphaMembership, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMembership(ctx, cl, request)
}

// DeleteMembership handles the gRPC request by passing it to the underlying Membership Delete() method.
func (s *MembershipServer) DeleteCloudidentityAlphaMembership(ctx context.Context, request *alphapb.DeleteCloudidentityAlphaMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMembership(ctx, ProtoToMembership(request.GetResource()))

}

// ListCloudidentityAlphaMembership handles the gRPC request by passing it to the underlying MembershipList() method.
func (s *MembershipServer) ListCloudidentityAlphaMembership(ctx context.Context, request *alphapb.ListCloudidentityAlphaMembershipRequest) (*alphapb.ListCloudidentityAlphaMembershipResponse, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMembership(ctx, request.GetGroup())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudidentityAlphaMembership
	for _, r := range resources.Items {
		rp := MembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudidentityAlphaMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMembership(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
