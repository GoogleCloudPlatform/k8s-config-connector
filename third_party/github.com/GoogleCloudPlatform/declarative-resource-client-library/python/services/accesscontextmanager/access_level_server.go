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
	accesscontextmanagerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/accesscontextmanager/accesscontextmanager_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/accesscontextmanager"
)

// Server implements the gRPC interface for AccessLevel.
type AccessLevelServer struct{}

// ProtoToAccessLevelBasicCombiningFunctionEnum converts a AccessLevelBasicCombiningFunctionEnum enum from its proto representation.
func ProtoToAccesscontextmanagerAccessLevelBasicCombiningFunctionEnum(e accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum) *accesscontextmanager.AccessLevelBasicCombiningFunctionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum_name[int32(e)]; ok {
		e := accesscontextmanager.AccessLevelBasicCombiningFunctionEnum(n[len("AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum"):])
		return &e
	}
	return nil
}

// ProtoToAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum converts a AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum enum from its proto representation.
func ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(e accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum) *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum_name[int32(e)]; ok {
		e := accesscontextmanager.AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(n[len("AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum"):])
		return &e
	}
	return nil
}

// ProtoToAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum converts a AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum enum from its proto representation.
func ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(e accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum) *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum_name[int32(e)]; ok {
		e := accesscontextmanager.AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(n[len("AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum"):])
		return &e
	}
	return nil
}

// ProtoToAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum converts a AccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum enum from its proto representation.
func ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum(e accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum) *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum_name[int32(e)]; ok {
		e := accesscontextmanager.AccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum(n[len("AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAccessLevelBasic converts a AccessLevelBasic resource from its proto representation.
func ProtoToAccesscontextmanagerAccessLevelBasic(p *accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasic) *accesscontextmanager.AccessLevelBasic {
	if p == nil {
		return nil
	}
	obj := &accesscontextmanager.AccessLevelBasic{
		CombiningFunction: ProtoToAccesscontextmanagerAccessLevelBasicCombiningFunctionEnum(p.GetCombiningFunction()),
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToAccesscontextmanagerAccessLevelBasicConditions(r))
	}
	return obj
}

// ProtoToAccessLevelBasicConditions converts a AccessLevelBasicConditions resource from its proto representation.
func ProtoToAccesscontextmanagerAccessLevelBasicConditions(p *accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditions) *accesscontextmanager.AccessLevelBasicConditions {
	if p == nil {
		return nil
	}
	obj := &accesscontextmanager.AccessLevelBasicConditions{
		Negate:       dcl.Bool(p.Negate),
		DevicePolicy: ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicy(p.GetDevicePolicy()),
	}
	for _, r := range p.GetRegions() {
		obj.Regions = append(obj.Regions, r)
	}
	for _, r := range p.GetIpSubnetworks() {
		obj.IPSubnetworks = append(obj.IPSubnetworks, r)
	}
	for _, r := range p.GetRequiredAccessLevels() {
		obj.RequiredAccessLevels = append(obj.RequiredAccessLevels, r)
	}
	for _, r := range p.GetMembers() {
		obj.Members = append(obj.Members, r)
	}
	return obj
}

// ProtoToAccessLevelBasicConditionsDevicePolicy converts a AccessLevelBasicConditionsDevicePolicy resource from its proto representation.
func ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicy(p *accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicy) *accesscontextmanager.AccessLevelBasicConditionsDevicePolicy {
	if p == nil {
		return nil
	}
	obj := &accesscontextmanager.AccessLevelBasicConditionsDevicePolicy{
		RequireScreenlock:    dcl.Bool(p.RequireScreenlock),
		RequireAdminApproval: dcl.Bool(p.RequireAdminApproval),
		RequireCorpOwned:     dcl.Bool(p.RequireCorpOwned),
	}
	for _, r := range p.GetAllowedEncryptionStatuses() {
		obj.AllowedEncryptionStatuses = append(obj.AllowedEncryptionStatuses, *ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(r))
	}
	for _, r := range p.GetAllowedDeviceManagementLevels() {
		obj.AllowedDeviceManagementLevels = append(obj.AllowedDeviceManagementLevels, *ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(r))
	}
	for _, r := range p.GetOsConstraints() {
		obj.OSConstraints = append(obj.OSConstraints, *ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraints(r))
	}
	return obj
}

// ProtoToAccessLevelBasicConditionsDevicePolicyOSConstraints converts a AccessLevelBasicConditionsDevicePolicyOSConstraints resource from its proto representation.
func ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraints(p *accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraints) *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyOSConstraints {
	if p == nil {
		return nil
	}
	obj := &accesscontextmanager.AccessLevelBasicConditionsDevicePolicyOSConstraints{
		MinimumVersion:          dcl.StringOrNil(p.MinimumVersion),
		OSType:                  ProtoToAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum(p.GetOsType()),
		RequireVerifiedChromeOS: dcl.Bool(p.RequireVerifiedChromeOs),
	}
	return obj
}

// ProtoToAccessLevel converts a AccessLevel resource from its proto representation.
func ProtoToAccessLevel(p *accesscontextmanagerpb.AccesscontextmanagerAccessLevel) *accesscontextmanager.AccessLevel {
	obj := &accesscontextmanager.AccessLevel{
		Title:       dcl.StringOrNil(p.Title),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.Description),
		Basic:       ProtoToAccesscontextmanagerAccessLevelBasic(p.GetBasic()),
		Name:        dcl.StringOrNil(p.Name),
		Policy:      dcl.StringOrNil(p.Policy),
	}
	return obj
}

// AccessLevelBasicCombiningFunctionEnumToProto converts a AccessLevelBasicCombiningFunctionEnum enum to its proto representation.
func AccesscontextmanagerAccessLevelBasicCombiningFunctionEnumToProto(e *accesscontextmanager.AccessLevelBasicCombiningFunctionEnum) accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum {
	if e == nil {
		return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum(0)
	}
	if v, ok := accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum_value["AccessLevelBasicCombiningFunctionEnum"+string(*e)]; ok {
		return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum(v)
	}
	return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicCombiningFunctionEnum(0)
}

// AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumToProto converts a AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum enum to its proto representation.
func AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnumToProto(e *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum) accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum {
	if e == nil {
		return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(0)
	}
	if v, ok := accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum_value["AccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum"+string(*e)]; ok {
		return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(v)
	}
	return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(0)
}

// AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumToProto converts a AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum enum to its proto representation.
func AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnumToProto(e *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum) accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum {
	if e == nil {
		return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(0)
	}
	if v, ok := accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum_value["AccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum"+string(*e)]; ok {
		return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(v)
	}
	return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(0)
}

// AccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnumToProto converts a AccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum enum to its proto representation.
func AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnumToProto(e *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum) accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum {
	if e == nil {
		return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum(0)
	}
	if v, ok := accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum_value["AccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum"+string(*e)]; ok {
		return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum(v)
	}
	return accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnum(0)
}

// AccessLevelBasicToProto converts a AccessLevelBasic resource to its proto representation.
func AccesscontextmanagerAccessLevelBasicToProto(o *accesscontextmanager.AccessLevelBasic) *accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasic {
	if o == nil {
		return nil
	}
	p := &accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasic{
		CombiningFunction: AccesscontextmanagerAccessLevelBasicCombiningFunctionEnumToProto(o.CombiningFunction),
	}
	for _, r := range o.Conditions {
		p.Conditions = append(p.Conditions, AccesscontextmanagerAccessLevelBasicConditionsToProto(&r))
	}
	return p
}

// AccessLevelBasicConditionsToProto converts a AccessLevelBasicConditions resource to its proto representation.
func AccesscontextmanagerAccessLevelBasicConditionsToProto(o *accesscontextmanager.AccessLevelBasicConditions) *accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditions {
	if o == nil {
		return nil
	}
	p := &accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditions{
		Negate:       dcl.ValueOrEmptyBool(o.Negate),
		DevicePolicy: AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyToProto(o.DevicePolicy),
	}
	for _, r := range o.Regions {
		p.Regions = append(p.Regions, r)
	}
	for _, r := range o.IPSubnetworks {
		p.IpSubnetworks = append(p.IpSubnetworks, r)
	}
	for _, r := range o.RequiredAccessLevels {
		p.RequiredAccessLevels = append(p.RequiredAccessLevels, r)
	}
	for _, r := range o.Members {
		p.Members = append(p.Members, r)
	}
	return p
}

// AccessLevelBasicConditionsDevicePolicyToProto converts a AccessLevelBasicConditionsDevicePolicy resource to its proto representation.
func AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyToProto(o *accesscontextmanager.AccessLevelBasicConditionsDevicePolicy) *accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicy {
	if o == nil {
		return nil
	}
	p := &accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicy{
		RequireScreenlock:    dcl.ValueOrEmptyBool(o.RequireScreenlock),
		RequireAdminApproval: dcl.ValueOrEmptyBool(o.RequireAdminApproval),
		RequireCorpOwned:     dcl.ValueOrEmptyBool(o.RequireCorpOwned),
	}
	for _, r := range o.AllowedEncryptionStatuses {
		p.AllowedEncryptionStatuses = append(p.AllowedEncryptionStatuses, accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum(accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatusesEnum_value[string(r)]))
	}
	for _, r := range o.AllowedDeviceManagementLevels {
		p.AllowedDeviceManagementLevels = append(p.AllowedDeviceManagementLevels, accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum(accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevelsEnum_value[string(r)]))
	}
	for _, r := range o.OSConstraints {
		p.OsConstraints = append(p.OsConstraints, AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsToProto(&r))
	}
	return p
}

// AccessLevelBasicConditionsDevicePolicyOSConstraintsToProto converts a AccessLevelBasicConditionsDevicePolicyOSConstraints resource to its proto representation.
func AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsToProto(o *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyOSConstraints) *accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraints {
	if o == nil {
		return nil
	}
	p := &accesscontextmanagerpb.AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraints{
		MinimumVersion:          dcl.ValueOrEmptyString(o.MinimumVersion),
		OsType:                  AccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOSConstraintsOSTypeEnumToProto(o.OSType),
		RequireVerifiedChromeOs: dcl.ValueOrEmptyBool(o.RequireVerifiedChromeOS),
	}
	return p
}

// AccessLevelToProto converts a AccessLevel resource to its proto representation.
func AccessLevelToProto(resource *accesscontextmanager.AccessLevel) *accesscontextmanagerpb.AccesscontextmanagerAccessLevel {
	p := &accesscontextmanagerpb.AccesscontextmanagerAccessLevel{
		Title:       dcl.ValueOrEmptyString(resource.Title),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Basic:       AccesscontextmanagerAccessLevelBasicToProto(resource.Basic),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Policy:      dcl.ValueOrEmptyString(resource.Policy),
	}

	return p
}

// ApplyAccessLevel handles the gRPC request by passing it to the underlying AccessLevel Apply() method.
func (s *AccessLevelServer) applyAccessLevel(ctx context.Context, c *accesscontextmanager.Client, request *accesscontextmanagerpb.ApplyAccesscontextmanagerAccessLevelRequest) (*accesscontextmanagerpb.AccesscontextmanagerAccessLevel, error) {
	p := ProtoToAccessLevel(request.GetResource())
	res, err := c.ApplyAccessLevel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AccessLevelToProto(res)
	return r, nil
}

// ApplyAccessLevel handles the gRPC request by passing it to the underlying AccessLevel Apply() method.
func (s *AccessLevelServer) ApplyAccesscontextmanagerAccessLevel(ctx context.Context, request *accesscontextmanagerpb.ApplyAccesscontextmanagerAccessLevelRequest) (*accesscontextmanagerpb.AccesscontextmanagerAccessLevel, error) {
	cl, err := createConfigAccessLevel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAccessLevel(ctx, cl, request)
}

// DeleteAccessLevel handles the gRPC request by passing it to the underlying AccessLevel Delete() method.
func (s *AccessLevelServer) DeleteAccesscontextmanagerAccessLevel(ctx context.Context, request *accesscontextmanagerpb.DeleteAccesscontextmanagerAccessLevelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAccessLevel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAccessLevel(ctx, ProtoToAccessLevel(request.GetResource()))

}

// ListAccesscontextmanagerAccessLevel handles the gRPC request by passing it to the underlying AccessLevelList() method.
func (s *AccessLevelServer) ListAccesscontextmanagerAccessLevel(ctx context.Context, request *accesscontextmanagerpb.ListAccesscontextmanagerAccessLevelRequest) (*accesscontextmanagerpb.ListAccesscontextmanagerAccessLevelResponse, error) {
	cl, err := createConfigAccessLevel(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAccessLevel(ctx, request.Policy)
	if err != nil {
		return nil, err
	}
	var protos []*accesscontextmanagerpb.AccesscontextmanagerAccessLevel
	for _, r := range resources.Items {
		rp := AccessLevelToProto(r)
		protos = append(protos, rp)
	}
	return &accesscontextmanagerpb.ListAccesscontextmanagerAccessLevelResponse{Items: protos}, nil
}

func createConfigAccessLevel(ctx context.Context, service_account_file string) (*accesscontextmanager.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return accesscontextmanager.NewClient(conf), nil
}
