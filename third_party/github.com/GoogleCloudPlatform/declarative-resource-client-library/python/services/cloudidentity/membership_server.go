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
	cloudidentitypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudidentity/cloudidentity_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity"
)

// MembershipServer implements the gRPC interface for Membership.
type MembershipServer struct{}

// ProtoToMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum enum from its proto representation.
func ProtoToCloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(e cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum) *cloudidentity.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum_name[int32(e)]; ok {
		e := cloudidentity.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(n[len("CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipTypeEnum converts a MembershipTypeEnum enum from its proto representation.
func ProtoToCloudidentityMembershipTypeEnum(e cloudidentitypb.CloudidentityMembershipTypeEnum) *cloudidentity.MembershipTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudidentitypb.CloudidentityMembershipTypeEnum_name[int32(e)]; ok {
		e := cloudidentity.MembershipTypeEnum(n[len("CloudidentityMembershipTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipDeliverySettingEnum converts a MembershipDeliverySettingEnum enum from its proto representation.
func ProtoToCloudidentityMembershipDeliverySettingEnum(e cloudidentitypb.CloudidentityMembershipDeliverySettingEnum) *cloudidentity.MembershipDeliverySettingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudidentitypb.CloudidentityMembershipDeliverySettingEnum_name[int32(e)]; ok {
		e := cloudidentity.MembershipDeliverySettingEnum(n[len("CloudidentityMembershipDeliverySettingEnum"):])
		return &e
	}
	return nil
}

// ProtoToMembershipPreferredMemberKey converts a MembershipPreferredMemberKey object from its proto representation.
func ProtoToCloudidentityMembershipPreferredMemberKey(p *cloudidentitypb.CloudidentityMembershipPreferredMemberKey) *cloudidentity.MembershipPreferredMemberKey {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.MembershipPreferredMemberKey{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToMembershipRoles converts a MembershipRoles object from its proto representation.
func ProtoToCloudidentityMembershipRoles(p *cloudidentitypb.CloudidentityMembershipRoles) *cloudidentity.MembershipRoles {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.MembershipRoles{
		Name:                   dcl.StringOrNil(p.GetName()),
		ExpiryDetail:           ProtoToCloudidentityMembershipRolesExpiryDetail(p.GetExpiryDetail()),
		RestrictionEvaluations: ProtoToCloudidentityMembershipRolesRestrictionEvaluations(p.GetRestrictionEvaluations()),
	}
	return obj
}

// ProtoToMembershipRolesExpiryDetail converts a MembershipRolesExpiryDetail object from its proto representation.
func ProtoToCloudidentityMembershipRolesExpiryDetail(p *cloudidentitypb.CloudidentityMembershipRolesExpiryDetail) *cloudidentity.MembershipRolesExpiryDetail {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.MembershipRolesExpiryDetail{
		ExpireTime: dcl.StringOrNil(p.GetExpireTime()),
	}
	return obj
}

// ProtoToMembershipRolesRestrictionEvaluations converts a MembershipRolesRestrictionEvaluations object from its proto representation.
func ProtoToCloudidentityMembershipRolesRestrictionEvaluations(p *cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluations) *cloudidentity.MembershipRolesRestrictionEvaluations {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.MembershipRolesRestrictionEvaluations{
		MemberRestrictionEvaluation: ProtoToCloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(p.GetMemberRestrictionEvaluation()),
	}
	return obj
}

// ProtoToMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation object from its proto representation.
func ProtoToCloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(p *cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) *cloudidentity.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{
		State: ProtoToCloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(p.GetState()),
	}
	return obj
}

// ProtoToMembershipDisplayName converts a MembershipDisplayName object from its proto representation.
func ProtoToCloudidentityMembershipDisplayName(p *cloudidentitypb.CloudidentityMembershipDisplayName) *cloudidentity.MembershipDisplayName {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.MembershipDisplayName{
		GivenName:  dcl.StringOrNil(p.GetGivenName()),
		FamilyName: dcl.StringOrNil(p.GetFamilyName()),
		FullName:   dcl.StringOrNil(p.GetFullName()),
	}
	return obj
}

// ProtoToMembership converts a Membership resource from its proto representation.
func ProtoToMembership(p *cloudidentitypb.CloudidentityMembership) *cloudidentity.Membership {
	obj := &cloudidentity.Membership{
		Name:               dcl.StringOrNil(p.GetName()),
		PreferredMemberKey: ProtoToCloudidentityMembershipPreferredMemberKey(p.GetPreferredMemberKey()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		Type:               ProtoToCloudidentityMembershipTypeEnum(p.GetType()),
		DeliverySetting:    ProtoToCloudidentityMembershipDeliverySettingEnum(p.GetDeliverySetting()),
		DisplayName:        ProtoToCloudidentityMembershipDisplayName(p.GetDisplayName()),
		Group:              dcl.StringOrNil(p.GetGroup()),
	}
	for _, r := range p.GetRoles() {
		obj.Roles = append(obj.Roles, *ProtoToCloudidentityMembershipRoles(r))
	}
	return obj
}

// MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum enum to its proto representation.
func CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto(e *cloudidentity.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum) cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	if e == nil {
		return cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(0)
	}
	if v, ok := cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum_value["MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum"+string(*e)]; ok {
		return cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(v)
	}
	return cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(0)
}

// MembershipTypeEnumToProto converts a MembershipTypeEnum enum to its proto representation.
func CloudidentityMembershipTypeEnumToProto(e *cloudidentity.MembershipTypeEnum) cloudidentitypb.CloudidentityMembershipTypeEnum {
	if e == nil {
		return cloudidentitypb.CloudidentityMembershipTypeEnum(0)
	}
	if v, ok := cloudidentitypb.CloudidentityMembershipTypeEnum_value["MembershipTypeEnum"+string(*e)]; ok {
		return cloudidentitypb.CloudidentityMembershipTypeEnum(v)
	}
	return cloudidentitypb.CloudidentityMembershipTypeEnum(0)
}

// MembershipDeliverySettingEnumToProto converts a MembershipDeliverySettingEnum enum to its proto representation.
func CloudidentityMembershipDeliverySettingEnumToProto(e *cloudidentity.MembershipDeliverySettingEnum) cloudidentitypb.CloudidentityMembershipDeliverySettingEnum {
	if e == nil {
		return cloudidentitypb.CloudidentityMembershipDeliverySettingEnum(0)
	}
	if v, ok := cloudidentitypb.CloudidentityMembershipDeliverySettingEnum_value["MembershipDeliverySettingEnum"+string(*e)]; ok {
		return cloudidentitypb.CloudidentityMembershipDeliverySettingEnum(v)
	}
	return cloudidentitypb.CloudidentityMembershipDeliverySettingEnum(0)
}

// MembershipPreferredMemberKeyToProto converts a MembershipPreferredMemberKey object to its proto representation.
func CloudidentityMembershipPreferredMemberKeyToProto(o *cloudidentity.MembershipPreferredMemberKey) *cloudidentitypb.CloudidentityMembershipPreferredMemberKey {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityMembershipPreferredMemberKey{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// MembershipRolesToProto converts a MembershipRoles object to its proto representation.
func CloudidentityMembershipRolesToProto(o *cloudidentity.MembershipRoles) *cloudidentitypb.CloudidentityMembershipRoles {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityMembershipRoles{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetExpiryDetail(CloudidentityMembershipRolesExpiryDetailToProto(o.ExpiryDetail))
	p.SetRestrictionEvaluations(CloudidentityMembershipRolesRestrictionEvaluationsToProto(o.RestrictionEvaluations))
	return p
}

// MembershipRolesExpiryDetailToProto converts a MembershipRolesExpiryDetail object to its proto representation.
func CloudidentityMembershipRolesExpiryDetailToProto(o *cloudidentity.MembershipRolesExpiryDetail) *cloudidentitypb.CloudidentityMembershipRolesExpiryDetail {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityMembershipRolesExpiryDetail{}
	p.SetExpireTime(dcl.ValueOrEmptyString(o.ExpireTime))
	return p
}

// MembershipRolesRestrictionEvaluationsToProto converts a MembershipRolesRestrictionEvaluations object to its proto representation.
func CloudidentityMembershipRolesRestrictionEvaluationsToProto(o *cloudidentity.MembershipRolesRestrictionEvaluations) *cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluations {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluations{}
	p.SetMemberRestrictionEvaluation(CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto(o.MemberRestrictionEvaluation))
	return p
}

// MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto converts a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation object to its proto representation.
func CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationToProto(o *cloudidentity.MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) *cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	p.SetState(CloudidentityMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumToProto(o.State))
	return p
}

// MembershipDisplayNameToProto converts a MembershipDisplayName object to its proto representation.
func CloudidentityMembershipDisplayNameToProto(o *cloudidentity.MembershipDisplayName) *cloudidentitypb.CloudidentityMembershipDisplayName {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityMembershipDisplayName{}
	p.SetGivenName(dcl.ValueOrEmptyString(o.GivenName))
	p.SetFamilyName(dcl.ValueOrEmptyString(o.FamilyName))
	p.SetFullName(dcl.ValueOrEmptyString(o.FullName))
	return p
}

// MembershipToProto converts a Membership resource to its proto representation.
func MembershipToProto(resource *cloudidentity.Membership) *cloudidentitypb.CloudidentityMembership {
	p := &cloudidentitypb.CloudidentityMembership{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetPreferredMemberKey(CloudidentityMembershipPreferredMemberKeyToProto(resource.PreferredMemberKey))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetType(CloudidentityMembershipTypeEnumToProto(resource.Type))
	p.SetDeliverySetting(CloudidentityMembershipDeliverySettingEnumToProto(resource.DeliverySetting))
	p.SetDisplayName(CloudidentityMembershipDisplayNameToProto(resource.DisplayName))
	p.SetGroup(dcl.ValueOrEmptyString(resource.Group))
	sRoles := make([]*cloudidentitypb.CloudidentityMembershipRoles, len(resource.Roles))
	for i, r := range resource.Roles {
		sRoles[i] = CloudidentityMembershipRolesToProto(&r)
	}
	p.SetRoles(sRoles)

	return p
}

// applyMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) applyMembership(ctx context.Context, c *cloudidentity.Client, request *cloudidentitypb.ApplyCloudidentityMembershipRequest) (*cloudidentitypb.CloudidentityMembership, error) {
	p := ProtoToMembership(request.GetResource())
	res, err := c.ApplyMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MembershipToProto(res)
	return r, nil
}

// applyCloudidentityMembership handles the gRPC request by passing it to the underlying Membership Apply() method.
func (s *MembershipServer) ApplyCloudidentityMembership(ctx context.Context, request *cloudidentitypb.ApplyCloudidentityMembershipRequest) (*cloudidentitypb.CloudidentityMembership, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMembership(ctx, cl, request)
}

// DeleteMembership handles the gRPC request by passing it to the underlying Membership Delete() method.
func (s *MembershipServer) DeleteCloudidentityMembership(ctx context.Context, request *cloudidentitypb.DeleteCloudidentityMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMembership(ctx, ProtoToMembership(request.GetResource()))

}

// ListCloudidentityMembership handles the gRPC request by passing it to the underlying MembershipList() method.
func (s *MembershipServer) ListCloudidentityMembership(ctx context.Context, request *cloudidentitypb.ListCloudidentityMembershipRequest) (*cloudidentitypb.ListCloudidentityMembershipResponse, error) {
	cl, err := createConfigMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMembership(ctx, request.GetGroup())
	if err != nil {
		return nil, err
	}
	var protos []*cloudidentitypb.CloudidentityMembership
	for _, r := range resources.Items {
		rp := MembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudidentitypb.ListCloudidentityMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMembership(ctx context.Context, service_account_file string) (*cloudidentity.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudidentity.NewClient(conf), nil
}
