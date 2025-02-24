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

// GroupServer implements the gRPC interface for Group.
type GroupServer struct{}

// ProtoToGroupDynamicGroupMetadataQueriesResourceTypeEnum converts a GroupDynamicGroupMetadataQueriesResourceTypeEnum enum from its proto representation.
func ProtoToCloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum(e betapb.CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum) *beta.GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum_name[int32(e)]; ok {
		e := beta.GroupDynamicGroupMetadataQueriesResourceTypeEnum(n[len("CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupDynamicGroupMetadataStatusStatusEnum converts a GroupDynamicGroupMetadataStatusStatusEnum enum from its proto representation.
func ProtoToCloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum(e betapb.CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum) *beta.GroupDynamicGroupMetadataStatusStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum_name[int32(e)]; ok {
		e := beta.GroupDynamicGroupMetadataStatusStatusEnum(n[len("CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupInitialGroupConfigEnum converts a GroupInitialGroupConfigEnum enum from its proto representation.
func ProtoToCloudidentityBetaGroupInitialGroupConfigEnum(e betapb.CloudidentityBetaGroupInitialGroupConfigEnum) *beta.GroupInitialGroupConfigEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudidentityBetaGroupInitialGroupConfigEnum_name[int32(e)]; ok {
		e := beta.GroupInitialGroupConfigEnum(n[len("CloudidentityBetaGroupInitialGroupConfigEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupGroupKey converts a GroupGroupKey object from its proto representation.
func ProtoToCloudidentityBetaGroupGroupKey(p *betapb.CloudidentityBetaGroupGroupKey) *beta.GroupGroupKey {
	if p == nil {
		return nil
	}
	obj := &beta.GroupGroupKey{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupAdditionalGroupKeys converts a GroupAdditionalGroupKeys object from its proto representation.
func ProtoToCloudidentityBetaGroupAdditionalGroupKeys(p *betapb.CloudidentityBetaGroupAdditionalGroupKeys) *beta.GroupAdditionalGroupKeys {
	if p == nil {
		return nil
	}
	obj := &beta.GroupAdditionalGroupKeys{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupDirectMemberCountPerType converts a GroupDirectMemberCountPerType object from its proto representation.
func ProtoToCloudidentityBetaGroupDirectMemberCountPerType(p *betapb.CloudidentityBetaGroupDirectMemberCountPerType) *beta.GroupDirectMemberCountPerType {
	if p == nil {
		return nil
	}
	obj := &beta.GroupDirectMemberCountPerType{
		UserCount:  dcl.Int64OrNil(p.GetUserCount()),
		GroupCount: dcl.Int64OrNil(p.GetGroupCount()),
	}
	return obj
}

// ProtoToGroupDerivedAliases converts a GroupDerivedAliases object from its proto representation.
func ProtoToCloudidentityBetaGroupDerivedAliases(p *betapb.CloudidentityBetaGroupDerivedAliases) *beta.GroupDerivedAliases {
	if p == nil {
		return nil
	}
	obj := &beta.GroupDerivedAliases{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadata converts a GroupDynamicGroupMetadata object from its proto representation.
func ProtoToCloudidentityBetaGroupDynamicGroupMetadata(p *betapb.CloudidentityBetaGroupDynamicGroupMetadata) *beta.GroupDynamicGroupMetadata {
	if p == nil {
		return nil
	}
	obj := &beta.GroupDynamicGroupMetadata{
		Status: ProtoToCloudidentityBetaGroupDynamicGroupMetadataStatus(p.GetStatus()),
	}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, *ProtoToCloudidentityBetaGroupDynamicGroupMetadataQueries(r))
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadataQueries converts a GroupDynamicGroupMetadataQueries object from its proto representation.
func ProtoToCloudidentityBetaGroupDynamicGroupMetadataQueries(p *betapb.CloudidentityBetaGroupDynamicGroupMetadataQueries) *beta.GroupDynamicGroupMetadataQueries {
	if p == nil {
		return nil
	}
	obj := &beta.GroupDynamicGroupMetadataQueries{
		ResourceType: ProtoToCloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum(p.GetResourceType()),
		Query:        dcl.StringOrNil(p.GetQuery()),
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadataStatus converts a GroupDynamicGroupMetadataStatus object from its proto representation.
func ProtoToCloudidentityBetaGroupDynamicGroupMetadataStatus(p *betapb.CloudidentityBetaGroupDynamicGroupMetadataStatus) *beta.GroupDynamicGroupMetadataStatus {
	if p == nil {
		return nil
	}
	obj := &beta.GroupDynamicGroupMetadataStatus{
		Status:     ProtoToCloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum(p.GetStatus()),
		StatusTime: dcl.StringOrNil(p.GetStatusTime()),
	}
	return obj
}

// ProtoToGroupPosixGroups converts a GroupPosixGroups object from its proto representation.
func ProtoToCloudidentityBetaGroupPosixGroups(p *betapb.CloudidentityBetaGroupPosixGroups) *beta.GroupPosixGroups {
	if p == nil {
		return nil
	}
	obj := &beta.GroupPosixGroups{
		Name:     dcl.StringOrNil(p.GetName()),
		Gid:      dcl.StringOrNil(p.GetGid()),
		SystemId: dcl.StringOrNil(p.GetSystemId()),
	}
	return obj
}

// ProtoToGroup converts a Group resource from its proto representation.
func ProtoToGroup(p *betapb.CloudidentityBetaGroup) *beta.Group {
	obj := &beta.Group{
		Name:                     dcl.StringOrNil(p.GetName()),
		GroupKey:                 ProtoToCloudidentityBetaGroupGroupKey(p.GetGroupKey()),
		Parent:                   dcl.StringOrNil(p.GetParent()),
		DisplayName:              dcl.StringOrNil(p.GetDisplayName()),
		Description:              dcl.StringOrNil(p.GetDescription()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:               dcl.StringOrNil(p.GetUpdateTime()),
		DirectMemberCount:        dcl.Int64OrNil(p.GetDirectMemberCount()),
		DirectMemberCountPerType: ProtoToCloudidentityBetaGroupDirectMemberCountPerType(p.GetDirectMemberCountPerType()),
		DynamicGroupMetadata:     ProtoToCloudidentityBetaGroupDynamicGroupMetadata(p.GetDynamicGroupMetadata()),
		InitialGroupConfig:       ProtoToCloudidentityBetaGroupInitialGroupConfigEnum(p.GetInitialGroupConfig()),
	}
	for _, r := range p.GetAdditionalGroupKeys() {
		obj.AdditionalGroupKeys = append(obj.AdditionalGroupKeys, *ProtoToCloudidentityBetaGroupAdditionalGroupKeys(r))
	}
	for _, r := range p.GetDerivedAliases() {
		obj.DerivedAliases = append(obj.DerivedAliases, *ProtoToCloudidentityBetaGroupDerivedAliases(r))
	}
	for _, r := range p.GetPosixGroups() {
		obj.PosixGroups = append(obj.PosixGroups, *ProtoToCloudidentityBetaGroupPosixGroups(r))
	}
	return obj
}

// GroupDynamicGroupMetadataQueriesResourceTypeEnumToProto converts a GroupDynamicGroupMetadataQueriesResourceTypeEnum enum to its proto representation.
func CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnumToProto(e *beta.GroupDynamicGroupMetadataQueriesResourceTypeEnum) betapb.CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum {
	if e == nil {
		return betapb.CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum(0)
	}
	if v, ok := betapb.CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum_value["GroupDynamicGroupMetadataQueriesResourceTypeEnum"+string(*e)]; ok {
		return betapb.CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum(v)
	}
	return betapb.CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnum(0)
}

// GroupDynamicGroupMetadataStatusStatusEnumToProto converts a GroupDynamicGroupMetadataStatusStatusEnum enum to its proto representation.
func CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnumToProto(e *beta.GroupDynamicGroupMetadataStatusStatusEnum) betapb.CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum {
	if e == nil {
		return betapb.CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum(0)
	}
	if v, ok := betapb.CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum_value["GroupDynamicGroupMetadataStatusStatusEnum"+string(*e)]; ok {
		return betapb.CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum(v)
	}
	return betapb.CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnum(0)
}

// GroupInitialGroupConfigEnumToProto converts a GroupInitialGroupConfigEnum enum to its proto representation.
func CloudidentityBetaGroupInitialGroupConfigEnumToProto(e *beta.GroupInitialGroupConfigEnum) betapb.CloudidentityBetaGroupInitialGroupConfigEnum {
	if e == nil {
		return betapb.CloudidentityBetaGroupInitialGroupConfigEnum(0)
	}
	if v, ok := betapb.CloudidentityBetaGroupInitialGroupConfigEnum_value["GroupInitialGroupConfigEnum"+string(*e)]; ok {
		return betapb.CloudidentityBetaGroupInitialGroupConfigEnum(v)
	}
	return betapb.CloudidentityBetaGroupInitialGroupConfigEnum(0)
}

// GroupGroupKeyToProto converts a GroupGroupKey object to its proto representation.
func CloudidentityBetaGroupGroupKeyToProto(o *beta.GroupGroupKey) *betapb.CloudidentityBetaGroupGroupKey {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaGroupGroupKey{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupAdditionalGroupKeysToProto converts a GroupAdditionalGroupKeys object to its proto representation.
func CloudidentityBetaGroupAdditionalGroupKeysToProto(o *beta.GroupAdditionalGroupKeys) *betapb.CloudidentityBetaGroupAdditionalGroupKeys {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaGroupAdditionalGroupKeys{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupDirectMemberCountPerTypeToProto converts a GroupDirectMemberCountPerType object to its proto representation.
func CloudidentityBetaGroupDirectMemberCountPerTypeToProto(o *beta.GroupDirectMemberCountPerType) *betapb.CloudidentityBetaGroupDirectMemberCountPerType {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaGroupDirectMemberCountPerType{}
	p.SetUserCount(dcl.ValueOrEmptyInt64(o.UserCount))
	p.SetGroupCount(dcl.ValueOrEmptyInt64(o.GroupCount))
	return p
}

// GroupDerivedAliasesToProto converts a GroupDerivedAliases object to its proto representation.
func CloudidentityBetaGroupDerivedAliasesToProto(o *beta.GroupDerivedAliases) *betapb.CloudidentityBetaGroupDerivedAliases {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaGroupDerivedAliases{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupDynamicGroupMetadataToProto converts a GroupDynamicGroupMetadata object to its proto representation.
func CloudidentityBetaGroupDynamicGroupMetadataToProto(o *beta.GroupDynamicGroupMetadata) *betapb.CloudidentityBetaGroupDynamicGroupMetadata {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaGroupDynamicGroupMetadata{}
	p.SetStatus(CloudidentityBetaGroupDynamicGroupMetadataStatusToProto(o.Status))
	sQueries := make([]*betapb.CloudidentityBetaGroupDynamicGroupMetadataQueries, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = CloudidentityBetaGroupDynamicGroupMetadataQueriesToProto(&r)
	}
	p.SetQueries(sQueries)
	return p
}

// GroupDynamicGroupMetadataQueriesToProto converts a GroupDynamicGroupMetadataQueries object to its proto representation.
func CloudidentityBetaGroupDynamicGroupMetadataQueriesToProto(o *beta.GroupDynamicGroupMetadataQueries) *betapb.CloudidentityBetaGroupDynamicGroupMetadataQueries {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaGroupDynamicGroupMetadataQueries{}
	p.SetResourceType(CloudidentityBetaGroupDynamicGroupMetadataQueriesResourceTypeEnumToProto(o.ResourceType))
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	return p
}

// GroupDynamicGroupMetadataStatusToProto converts a GroupDynamicGroupMetadataStatus object to its proto representation.
func CloudidentityBetaGroupDynamicGroupMetadataStatusToProto(o *beta.GroupDynamicGroupMetadataStatus) *betapb.CloudidentityBetaGroupDynamicGroupMetadataStatus {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaGroupDynamicGroupMetadataStatus{}
	p.SetStatus(CloudidentityBetaGroupDynamicGroupMetadataStatusStatusEnumToProto(o.Status))
	p.SetStatusTime(dcl.ValueOrEmptyString(o.StatusTime))
	return p
}

// GroupPosixGroupsToProto converts a GroupPosixGroups object to its proto representation.
func CloudidentityBetaGroupPosixGroupsToProto(o *beta.GroupPosixGroups) *betapb.CloudidentityBetaGroupPosixGroups {
	if o == nil {
		return nil
	}
	p := &betapb.CloudidentityBetaGroupPosixGroups{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetGid(dcl.ValueOrEmptyString(o.Gid))
	p.SetSystemId(dcl.ValueOrEmptyString(o.SystemId))
	return p
}

// GroupToProto converts a Group resource to its proto representation.
func GroupToProto(resource *beta.Group) *betapb.CloudidentityBetaGroup {
	p := &betapb.CloudidentityBetaGroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetGroupKey(CloudidentityBetaGroupGroupKeyToProto(resource.GroupKey))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDirectMemberCount(dcl.ValueOrEmptyInt64(resource.DirectMemberCount))
	p.SetDirectMemberCountPerType(CloudidentityBetaGroupDirectMemberCountPerTypeToProto(resource.DirectMemberCountPerType))
	p.SetDynamicGroupMetadata(CloudidentityBetaGroupDynamicGroupMetadataToProto(resource.DynamicGroupMetadata))
	p.SetInitialGroupConfig(CloudidentityBetaGroupInitialGroupConfigEnumToProto(resource.InitialGroupConfig))
	sAdditionalGroupKeys := make([]*betapb.CloudidentityBetaGroupAdditionalGroupKeys, len(resource.AdditionalGroupKeys))
	for i, r := range resource.AdditionalGroupKeys {
		sAdditionalGroupKeys[i] = CloudidentityBetaGroupAdditionalGroupKeysToProto(&r)
	}
	p.SetAdditionalGroupKeys(sAdditionalGroupKeys)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sDerivedAliases := make([]*betapb.CloudidentityBetaGroupDerivedAliases, len(resource.DerivedAliases))
	for i, r := range resource.DerivedAliases {
		sDerivedAliases[i] = CloudidentityBetaGroupDerivedAliasesToProto(&r)
	}
	p.SetDerivedAliases(sDerivedAliases)
	sPosixGroups := make([]*betapb.CloudidentityBetaGroupPosixGroups, len(resource.PosixGroups))
	for i, r := range resource.PosixGroups {
		sPosixGroups[i] = CloudidentityBetaGroupPosixGroupsToProto(&r)
	}
	p.SetPosixGroups(sPosixGroups)

	return p
}

// applyGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) applyGroup(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudidentityBetaGroupRequest) (*betapb.CloudidentityBetaGroup, error) {
	p := ProtoToGroup(request.GetResource())
	res, err := c.ApplyGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GroupToProto(res)
	return r, nil
}

// applyCloudidentityBetaGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) ApplyCloudidentityBetaGroup(ctx context.Context, request *betapb.ApplyCloudidentityBetaGroupRequest) (*betapb.CloudidentityBetaGroup, error) {
	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGroup(ctx, cl, request)
}

// DeleteGroup handles the gRPC request by passing it to the underlying Group Delete() method.
func (s *GroupServer) DeleteCloudidentityBetaGroup(ctx context.Context, request *betapb.DeleteCloudidentityBetaGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGroup(ctx, ProtoToGroup(request.GetResource()))

}

// ListCloudidentityBetaGroup handles the gRPC request by passing it to the underlying GroupList() method.
func (s *GroupServer) ListCloudidentityBetaGroup(ctx context.Context, request *betapb.ListCloudidentityBetaGroupRequest) (*betapb.ListCloudidentityBetaGroupResponse, error) {
	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGroup(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudidentityBetaGroup
	for _, r := range resources.Items {
		rp := GroupToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudidentityBetaGroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGroup(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
