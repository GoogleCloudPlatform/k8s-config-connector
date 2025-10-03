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

// GroupServer implements the gRPC interface for Group.
type GroupServer struct{}

// ProtoToGroupDynamicGroupMetadataQueriesResourceTypeEnum converts a GroupDynamicGroupMetadataQueriesResourceTypeEnum enum from its proto representation.
func ProtoToCloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum(e alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum) *alpha.GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum_name[int32(e)]; ok {
		e := alpha.GroupDynamicGroupMetadataQueriesResourceTypeEnum(n[len("CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupDynamicGroupMetadataStatusStatusEnum converts a GroupDynamicGroupMetadataStatusStatusEnum enum from its proto representation.
func ProtoToCloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum(e alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum) *alpha.GroupDynamicGroupMetadataStatusStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum_name[int32(e)]; ok {
		e := alpha.GroupDynamicGroupMetadataStatusStatusEnum(n[len("CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupInitialGroupConfigEnum converts a GroupInitialGroupConfigEnum enum from its proto representation.
func ProtoToCloudidentityAlphaGroupInitialGroupConfigEnum(e alphapb.CloudidentityAlphaGroupInitialGroupConfigEnum) *alpha.GroupInitialGroupConfigEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudidentityAlphaGroupInitialGroupConfigEnum_name[int32(e)]; ok {
		e := alpha.GroupInitialGroupConfigEnum(n[len("CloudidentityAlphaGroupInitialGroupConfigEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupGroupKey converts a GroupGroupKey object from its proto representation.
func ProtoToCloudidentityAlphaGroupGroupKey(p *alphapb.CloudidentityAlphaGroupGroupKey) *alpha.GroupGroupKey {
	if p == nil {
		return nil
	}
	obj := &alpha.GroupGroupKey{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupAdditionalGroupKeys converts a GroupAdditionalGroupKeys object from its proto representation.
func ProtoToCloudidentityAlphaGroupAdditionalGroupKeys(p *alphapb.CloudidentityAlphaGroupAdditionalGroupKeys) *alpha.GroupAdditionalGroupKeys {
	if p == nil {
		return nil
	}
	obj := &alpha.GroupAdditionalGroupKeys{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupDirectMemberCountPerType converts a GroupDirectMemberCountPerType object from its proto representation.
func ProtoToCloudidentityAlphaGroupDirectMemberCountPerType(p *alphapb.CloudidentityAlphaGroupDirectMemberCountPerType) *alpha.GroupDirectMemberCountPerType {
	if p == nil {
		return nil
	}
	obj := &alpha.GroupDirectMemberCountPerType{
		UserCount:  dcl.Int64OrNil(p.GetUserCount()),
		GroupCount: dcl.Int64OrNil(p.GetGroupCount()),
	}
	return obj
}

// ProtoToGroupDerivedAliases converts a GroupDerivedAliases object from its proto representation.
func ProtoToCloudidentityAlphaGroupDerivedAliases(p *alphapb.CloudidentityAlphaGroupDerivedAliases) *alpha.GroupDerivedAliases {
	if p == nil {
		return nil
	}
	obj := &alpha.GroupDerivedAliases{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadata converts a GroupDynamicGroupMetadata object from its proto representation.
func ProtoToCloudidentityAlphaGroupDynamicGroupMetadata(p *alphapb.CloudidentityAlphaGroupDynamicGroupMetadata) *alpha.GroupDynamicGroupMetadata {
	if p == nil {
		return nil
	}
	obj := &alpha.GroupDynamicGroupMetadata{
		Status: ProtoToCloudidentityAlphaGroupDynamicGroupMetadataStatus(p.GetStatus()),
	}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, *ProtoToCloudidentityAlphaGroupDynamicGroupMetadataQueries(r))
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadataQueries converts a GroupDynamicGroupMetadataQueries object from its proto representation.
func ProtoToCloudidentityAlphaGroupDynamicGroupMetadataQueries(p *alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueries) *alpha.GroupDynamicGroupMetadataQueries {
	if p == nil {
		return nil
	}
	obj := &alpha.GroupDynamicGroupMetadataQueries{
		ResourceType: ProtoToCloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum(p.GetResourceType()),
		Query:        dcl.StringOrNil(p.GetQuery()),
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadataStatus converts a GroupDynamicGroupMetadataStatus object from its proto representation.
func ProtoToCloudidentityAlphaGroupDynamicGroupMetadataStatus(p *alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatus) *alpha.GroupDynamicGroupMetadataStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.GroupDynamicGroupMetadataStatus{
		Status:     ProtoToCloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum(p.GetStatus()),
		StatusTime: dcl.StringOrNil(p.GetStatusTime()),
	}
	return obj
}

// ProtoToGroupPosixGroups converts a GroupPosixGroups object from its proto representation.
func ProtoToCloudidentityAlphaGroupPosixGroups(p *alphapb.CloudidentityAlphaGroupPosixGroups) *alpha.GroupPosixGroups {
	if p == nil {
		return nil
	}
	obj := &alpha.GroupPosixGroups{
		Name:     dcl.StringOrNil(p.GetName()),
		Gid:      dcl.StringOrNil(p.GetGid()),
		SystemId: dcl.StringOrNil(p.GetSystemId()),
	}
	return obj
}

// ProtoToGroup converts a Group resource from its proto representation.
func ProtoToGroup(p *alphapb.CloudidentityAlphaGroup) *alpha.Group {
	obj := &alpha.Group{
		Name:                     dcl.StringOrNil(p.GetName()),
		GroupKey:                 ProtoToCloudidentityAlphaGroupGroupKey(p.GetGroupKey()),
		Parent:                   dcl.StringOrNil(p.GetParent()),
		DisplayName:              dcl.StringOrNil(p.GetDisplayName()),
		Description:              dcl.StringOrNil(p.GetDescription()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:               dcl.StringOrNil(p.GetUpdateTime()),
		DirectMemberCount:        dcl.Int64OrNil(p.GetDirectMemberCount()),
		DirectMemberCountPerType: ProtoToCloudidentityAlphaGroupDirectMemberCountPerType(p.GetDirectMemberCountPerType()),
		DynamicGroupMetadata:     ProtoToCloudidentityAlphaGroupDynamicGroupMetadata(p.GetDynamicGroupMetadata()),
		InitialGroupConfig:       ProtoToCloudidentityAlphaGroupInitialGroupConfigEnum(p.GetInitialGroupConfig()),
	}
	for _, r := range p.GetAdditionalGroupKeys() {
		obj.AdditionalGroupKeys = append(obj.AdditionalGroupKeys, *ProtoToCloudidentityAlphaGroupAdditionalGroupKeys(r))
	}
	for _, r := range p.GetDerivedAliases() {
		obj.DerivedAliases = append(obj.DerivedAliases, *ProtoToCloudidentityAlphaGroupDerivedAliases(r))
	}
	for _, r := range p.GetPosixGroups() {
		obj.PosixGroups = append(obj.PosixGroups, *ProtoToCloudidentityAlphaGroupPosixGroups(r))
	}
	return obj
}

// GroupDynamicGroupMetadataQueriesResourceTypeEnumToProto converts a GroupDynamicGroupMetadataQueriesResourceTypeEnum enum to its proto representation.
func CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnumToProto(e *alpha.GroupDynamicGroupMetadataQueriesResourceTypeEnum) alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum {
	if e == nil {
		return alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum(0)
	}
	if v, ok := alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum_value["GroupDynamicGroupMetadataQueriesResourceTypeEnum"+string(*e)]; ok {
		return alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum(v)
	}
	return alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnum(0)
}

// GroupDynamicGroupMetadataStatusStatusEnumToProto converts a GroupDynamicGroupMetadataStatusStatusEnum enum to its proto representation.
func CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnumToProto(e *alpha.GroupDynamicGroupMetadataStatusStatusEnum) alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum {
	if e == nil {
		return alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum(0)
	}
	if v, ok := alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum_value["GroupDynamicGroupMetadataStatusStatusEnum"+string(*e)]; ok {
		return alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum(v)
	}
	return alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnum(0)
}

// GroupInitialGroupConfigEnumToProto converts a GroupInitialGroupConfigEnum enum to its proto representation.
func CloudidentityAlphaGroupInitialGroupConfigEnumToProto(e *alpha.GroupInitialGroupConfigEnum) alphapb.CloudidentityAlphaGroupInitialGroupConfigEnum {
	if e == nil {
		return alphapb.CloudidentityAlphaGroupInitialGroupConfigEnum(0)
	}
	if v, ok := alphapb.CloudidentityAlphaGroupInitialGroupConfigEnum_value["GroupInitialGroupConfigEnum"+string(*e)]; ok {
		return alphapb.CloudidentityAlphaGroupInitialGroupConfigEnum(v)
	}
	return alphapb.CloudidentityAlphaGroupInitialGroupConfigEnum(0)
}

// GroupGroupKeyToProto converts a GroupGroupKey object to its proto representation.
func CloudidentityAlphaGroupGroupKeyToProto(o *alpha.GroupGroupKey) *alphapb.CloudidentityAlphaGroupGroupKey {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaGroupGroupKey{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupAdditionalGroupKeysToProto converts a GroupAdditionalGroupKeys object to its proto representation.
func CloudidentityAlphaGroupAdditionalGroupKeysToProto(o *alpha.GroupAdditionalGroupKeys) *alphapb.CloudidentityAlphaGroupAdditionalGroupKeys {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaGroupAdditionalGroupKeys{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupDirectMemberCountPerTypeToProto converts a GroupDirectMemberCountPerType object to its proto representation.
func CloudidentityAlphaGroupDirectMemberCountPerTypeToProto(o *alpha.GroupDirectMemberCountPerType) *alphapb.CloudidentityAlphaGroupDirectMemberCountPerType {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaGroupDirectMemberCountPerType{}
	p.SetUserCount(dcl.ValueOrEmptyInt64(o.UserCount))
	p.SetGroupCount(dcl.ValueOrEmptyInt64(o.GroupCount))
	return p
}

// GroupDerivedAliasesToProto converts a GroupDerivedAliases object to its proto representation.
func CloudidentityAlphaGroupDerivedAliasesToProto(o *alpha.GroupDerivedAliases) *alphapb.CloudidentityAlphaGroupDerivedAliases {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaGroupDerivedAliases{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupDynamicGroupMetadataToProto converts a GroupDynamicGroupMetadata object to its proto representation.
func CloudidentityAlphaGroupDynamicGroupMetadataToProto(o *alpha.GroupDynamicGroupMetadata) *alphapb.CloudidentityAlphaGroupDynamicGroupMetadata {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaGroupDynamicGroupMetadata{}
	p.SetStatus(CloudidentityAlphaGroupDynamicGroupMetadataStatusToProto(o.Status))
	sQueries := make([]*alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueries, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = CloudidentityAlphaGroupDynamicGroupMetadataQueriesToProto(&r)
	}
	p.SetQueries(sQueries)
	return p
}

// GroupDynamicGroupMetadataQueriesToProto converts a GroupDynamicGroupMetadataQueries object to its proto representation.
func CloudidentityAlphaGroupDynamicGroupMetadataQueriesToProto(o *alpha.GroupDynamicGroupMetadataQueries) *alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueries {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaGroupDynamicGroupMetadataQueries{}
	p.SetResourceType(CloudidentityAlphaGroupDynamicGroupMetadataQueriesResourceTypeEnumToProto(o.ResourceType))
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	return p
}

// GroupDynamicGroupMetadataStatusToProto converts a GroupDynamicGroupMetadataStatus object to its proto representation.
func CloudidentityAlphaGroupDynamicGroupMetadataStatusToProto(o *alpha.GroupDynamicGroupMetadataStatus) *alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaGroupDynamicGroupMetadataStatus{}
	p.SetStatus(CloudidentityAlphaGroupDynamicGroupMetadataStatusStatusEnumToProto(o.Status))
	p.SetStatusTime(dcl.ValueOrEmptyString(o.StatusTime))
	return p
}

// GroupPosixGroupsToProto converts a GroupPosixGroups object to its proto representation.
func CloudidentityAlphaGroupPosixGroupsToProto(o *alpha.GroupPosixGroups) *alphapb.CloudidentityAlphaGroupPosixGroups {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudidentityAlphaGroupPosixGroups{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetGid(dcl.ValueOrEmptyString(o.Gid))
	p.SetSystemId(dcl.ValueOrEmptyString(o.SystemId))
	return p
}

// GroupToProto converts a Group resource to its proto representation.
func GroupToProto(resource *alpha.Group) *alphapb.CloudidentityAlphaGroup {
	p := &alphapb.CloudidentityAlphaGroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetGroupKey(CloudidentityAlphaGroupGroupKeyToProto(resource.GroupKey))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDirectMemberCount(dcl.ValueOrEmptyInt64(resource.DirectMemberCount))
	p.SetDirectMemberCountPerType(CloudidentityAlphaGroupDirectMemberCountPerTypeToProto(resource.DirectMemberCountPerType))
	p.SetDynamicGroupMetadata(CloudidentityAlphaGroupDynamicGroupMetadataToProto(resource.DynamicGroupMetadata))
	p.SetInitialGroupConfig(CloudidentityAlphaGroupInitialGroupConfigEnumToProto(resource.InitialGroupConfig))
	sAdditionalGroupKeys := make([]*alphapb.CloudidentityAlphaGroupAdditionalGroupKeys, len(resource.AdditionalGroupKeys))
	for i, r := range resource.AdditionalGroupKeys {
		sAdditionalGroupKeys[i] = CloudidentityAlphaGroupAdditionalGroupKeysToProto(&r)
	}
	p.SetAdditionalGroupKeys(sAdditionalGroupKeys)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sDerivedAliases := make([]*alphapb.CloudidentityAlphaGroupDerivedAliases, len(resource.DerivedAliases))
	for i, r := range resource.DerivedAliases {
		sDerivedAliases[i] = CloudidentityAlphaGroupDerivedAliasesToProto(&r)
	}
	p.SetDerivedAliases(sDerivedAliases)
	sPosixGroups := make([]*alphapb.CloudidentityAlphaGroupPosixGroups, len(resource.PosixGroups))
	for i, r := range resource.PosixGroups {
		sPosixGroups[i] = CloudidentityAlphaGroupPosixGroupsToProto(&r)
	}
	p.SetPosixGroups(sPosixGroups)

	return p
}

// applyGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) applyGroup(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudidentityAlphaGroupRequest) (*alphapb.CloudidentityAlphaGroup, error) {
	p := ProtoToGroup(request.GetResource())
	res, err := c.ApplyGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GroupToProto(res)
	return r, nil
}

// applyCloudidentityAlphaGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) ApplyCloudidentityAlphaGroup(ctx context.Context, request *alphapb.ApplyCloudidentityAlphaGroupRequest) (*alphapb.CloudidentityAlphaGroup, error) {
	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGroup(ctx, cl, request)
}

// DeleteGroup handles the gRPC request by passing it to the underlying Group Delete() method.
func (s *GroupServer) DeleteCloudidentityAlphaGroup(ctx context.Context, request *alphapb.DeleteCloudidentityAlphaGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGroup(ctx, ProtoToGroup(request.GetResource()))

}

// ListCloudidentityAlphaGroup handles the gRPC request by passing it to the underlying GroupList() method.
func (s *GroupServer) ListCloudidentityAlphaGroup(ctx context.Context, request *alphapb.ListCloudidentityAlphaGroupRequest) (*alphapb.ListCloudidentityAlphaGroupResponse, error) {
	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGroup(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudidentityAlphaGroup
	for _, r := range resources.Items {
		rp := GroupToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudidentityAlphaGroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGroup(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
