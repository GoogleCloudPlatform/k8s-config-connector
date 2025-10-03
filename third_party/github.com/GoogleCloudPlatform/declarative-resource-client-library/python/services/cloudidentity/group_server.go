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

// GroupServer implements the gRPC interface for Group.
type GroupServer struct{}

// ProtoToGroupDynamicGroupMetadataQueriesResourceTypeEnum converts a GroupDynamicGroupMetadataQueriesResourceTypeEnum enum from its proto representation.
func ProtoToCloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum(e cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum) *cloudidentity.GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum_name[int32(e)]; ok {
		e := cloudidentity.GroupDynamicGroupMetadataQueriesResourceTypeEnum(n[len("CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupDynamicGroupMetadataStatusStatusEnum converts a GroupDynamicGroupMetadataStatusStatusEnum enum from its proto representation.
func ProtoToCloudidentityGroupDynamicGroupMetadataStatusStatusEnum(e cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatusStatusEnum) *cloudidentity.GroupDynamicGroupMetadataStatusStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatusStatusEnum_name[int32(e)]; ok {
		e := cloudidentity.GroupDynamicGroupMetadataStatusStatusEnum(n[len("CloudidentityGroupDynamicGroupMetadataStatusStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupInitialGroupConfigEnum converts a GroupInitialGroupConfigEnum enum from its proto representation.
func ProtoToCloudidentityGroupInitialGroupConfigEnum(e cloudidentitypb.CloudidentityGroupInitialGroupConfigEnum) *cloudidentity.GroupInitialGroupConfigEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudidentitypb.CloudidentityGroupInitialGroupConfigEnum_name[int32(e)]; ok {
		e := cloudidentity.GroupInitialGroupConfigEnum(n[len("CloudidentityGroupInitialGroupConfigEnum"):])
		return &e
	}
	return nil
}

// ProtoToGroupGroupKey converts a GroupGroupKey object from its proto representation.
func ProtoToCloudidentityGroupGroupKey(p *cloudidentitypb.CloudidentityGroupGroupKey) *cloudidentity.GroupGroupKey {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.GroupGroupKey{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupAdditionalGroupKeys converts a GroupAdditionalGroupKeys object from its proto representation.
func ProtoToCloudidentityGroupAdditionalGroupKeys(p *cloudidentitypb.CloudidentityGroupAdditionalGroupKeys) *cloudidentity.GroupAdditionalGroupKeys {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.GroupAdditionalGroupKeys{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupDirectMemberCountPerType converts a GroupDirectMemberCountPerType object from its proto representation.
func ProtoToCloudidentityGroupDirectMemberCountPerType(p *cloudidentitypb.CloudidentityGroupDirectMemberCountPerType) *cloudidentity.GroupDirectMemberCountPerType {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.GroupDirectMemberCountPerType{
		UserCount:  dcl.Int64OrNil(p.GetUserCount()),
		GroupCount: dcl.Int64OrNil(p.GetGroupCount()),
	}
	return obj
}

// ProtoToGroupDerivedAliases converts a GroupDerivedAliases object from its proto representation.
func ProtoToCloudidentityGroupDerivedAliases(p *cloudidentitypb.CloudidentityGroupDerivedAliases) *cloudidentity.GroupDerivedAliases {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.GroupDerivedAliases{
		Id:        dcl.StringOrNil(p.GetId()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadata converts a GroupDynamicGroupMetadata object from its proto representation.
func ProtoToCloudidentityGroupDynamicGroupMetadata(p *cloudidentitypb.CloudidentityGroupDynamicGroupMetadata) *cloudidentity.GroupDynamicGroupMetadata {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.GroupDynamicGroupMetadata{
		Status: ProtoToCloudidentityGroupDynamicGroupMetadataStatus(p.GetStatus()),
	}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, *ProtoToCloudidentityGroupDynamicGroupMetadataQueries(r))
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadataQueries converts a GroupDynamicGroupMetadataQueries object from its proto representation.
func ProtoToCloudidentityGroupDynamicGroupMetadataQueries(p *cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueries) *cloudidentity.GroupDynamicGroupMetadataQueries {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.GroupDynamicGroupMetadataQueries{
		ResourceType: ProtoToCloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum(p.GetResourceType()),
		Query:        dcl.StringOrNil(p.GetQuery()),
	}
	return obj
}

// ProtoToGroupDynamicGroupMetadataStatus converts a GroupDynamicGroupMetadataStatus object from its proto representation.
func ProtoToCloudidentityGroupDynamicGroupMetadataStatus(p *cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatus) *cloudidentity.GroupDynamicGroupMetadataStatus {
	if p == nil {
		return nil
	}
	obj := &cloudidentity.GroupDynamicGroupMetadataStatus{
		Status:     ProtoToCloudidentityGroupDynamicGroupMetadataStatusStatusEnum(p.GetStatus()),
		StatusTime: dcl.StringOrNil(p.GetStatusTime()),
	}
	return obj
}

// ProtoToGroup converts a Group resource from its proto representation.
func ProtoToGroup(p *cloudidentitypb.CloudidentityGroup) *cloudidentity.Group {
	obj := &cloudidentity.Group{
		Name:                     dcl.StringOrNil(p.GetName()),
		GroupKey:                 ProtoToCloudidentityGroupGroupKey(p.GetGroupKey()),
		Parent:                   dcl.StringOrNil(p.GetParent()),
		DisplayName:              dcl.StringOrNil(p.GetDisplayName()),
		Description:              dcl.StringOrNil(p.GetDescription()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:               dcl.StringOrNil(p.GetUpdateTime()),
		DirectMemberCount:        dcl.Int64OrNil(p.GetDirectMemberCount()),
		DirectMemberCountPerType: ProtoToCloudidentityGroupDirectMemberCountPerType(p.GetDirectMemberCountPerType()),
		DynamicGroupMetadata:     ProtoToCloudidentityGroupDynamicGroupMetadata(p.GetDynamicGroupMetadata()),
		InitialGroupConfig:       ProtoToCloudidentityGroupInitialGroupConfigEnum(p.GetInitialGroupConfig()),
	}
	for _, r := range p.GetAdditionalGroupKeys() {
		obj.AdditionalGroupKeys = append(obj.AdditionalGroupKeys, *ProtoToCloudidentityGroupAdditionalGroupKeys(r))
	}
	for _, r := range p.GetDerivedAliases() {
		obj.DerivedAliases = append(obj.DerivedAliases, *ProtoToCloudidentityGroupDerivedAliases(r))
	}
	return obj
}

// GroupDynamicGroupMetadataQueriesResourceTypeEnumToProto converts a GroupDynamicGroupMetadataQueriesResourceTypeEnum enum to its proto representation.
func CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnumToProto(e *cloudidentity.GroupDynamicGroupMetadataQueriesResourceTypeEnum) cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum {
	if e == nil {
		return cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum(0)
	}
	if v, ok := cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum_value["GroupDynamicGroupMetadataQueriesResourceTypeEnum"+string(*e)]; ok {
		return cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum(v)
	}
	return cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnum(0)
}

// GroupDynamicGroupMetadataStatusStatusEnumToProto converts a GroupDynamicGroupMetadataStatusStatusEnum enum to its proto representation.
func CloudidentityGroupDynamicGroupMetadataStatusStatusEnumToProto(e *cloudidentity.GroupDynamicGroupMetadataStatusStatusEnum) cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatusStatusEnum {
	if e == nil {
		return cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatusStatusEnum(0)
	}
	if v, ok := cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatusStatusEnum_value["GroupDynamicGroupMetadataStatusStatusEnum"+string(*e)]; ok {
		return cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatusStatusEnum(v)
	}
	return cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatusStatusEnum(0)
}

// GroupInitialGroupConfigEnumToProto converts a GroupInitialGroupConfigEnum enum to its proto representation.
func CloudidentityGroupInitialGroupConfigEnumToProto(e *cloudidentity.GroupInitialGroupConfigEnum) cloudidentitypb.CloudidentityGroupInitialGroupConfigEnum {
	if e == nil {
		return cloudidentitypb.CloudidentityGroupInitialGroupConfigEnum(0)
	}
	if v, ok := cloudidentitypb.CloudidentityGroupInitialGroupConfigEnum_value["GroupInitialGroupConfigEnum"+string(*e)]; ok {
		return cloudidentitypb.CloudidentityGroupInitialGroupConfigEnum(v)
	}
	return cloudidentitypb.CloudidentityGroupInitialGroupConfigEnum(0)
}

// GroupGroupKeyToProto converts a GroupGroupKey object to its proto representation.
func CloudidentityGroupGroupKeyToProto(o *cloudidentity.GroupGroupKey) *cloudidentitypb.CloudidentityGroupGroupKey {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityGroupGroupKey{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupAdditionalGroupKeysToProto converts a GroupAdditionalGroupKeys object to its proto representation.
func CloudidentityGroupAdditionalGroupKeysToProto(o *cloudidentity.GroupAdditionalGroupKeys) *cloudidentitypb.CloudidentityGroupAdditionalGroupKeys {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityGroupAdditionalGroupKeys{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupDirectMemberCountPerTypeToProto converts a GroupDirectMemberCountPerType object to its proto representation.
func CloudidentityGroupDirectMemberCountPerTypeToProto(o *cloudidentity.GroupDirectMemberCountPerType) *cloudidentitypb.CloudidentityGroupDirectMemberCountPerType {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityGroupDirectMemberCountPerType{}
	p.SetUserCount(dcl.ValueOrEmptyInt64(o.UserCount))
	p.SetGroupCount(dcl.ValueOrEmptyInt64(o.GroupCount))
	return p
}

// GroupDerivedAliasesToProto converts a GroupDerivedAliases object to its proto representation.
func CloudidentityGroupDerivedAliasesToProto(o *cloudidentity.GroupDerivedAliases) *cloudidentitypb.CloudidentityGroupDerivedAliases {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityGroupDerivedAliases{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	return p
}

// GroupDynamicGroupMetadataToProto converts a GroupDynamicGroupMetadata object to its proto representation.
func CloudidentityGroupDynamicGroupMetadataToProto(o *cloudidentity.GroupDynamicGroupMetadata) *cloudidentitypb.CloudidentityGroupDynamicGroupMetadata {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityGroupDynamicGroupMetadata{}
	p.SetStatus(CloudidentityGroupDynamicGroupMetadataStatusToProto(o.Status))
	sQueries := make([]*cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueries, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = CloudidentityGroupDynamicGroupMetadataQueriesToProto(&r)
	}
	p.SetQueries(sQueries)
	return p
}

// GroupDynamicGroupMetadataQueriesToProto converts a GroupDynamicGroupMetadataQueries object to its proto representation.
func CloudidentityGroupDynamicGroupMetadataQueriesToProto(o *cloudidentity.GroupDynamicGroupMetadataQueries) *cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueries {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityGroupDynamicGroupMetadataQueries{}
	p.SetResourceType(CloudidentityGroupDynamicGroupMetadataQueriesResourceTypeEnumToProto(o.ResourceType))
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	return p
}

// GroupDynamicGroupMetadataStatusToProto converts a GroupDynamicGroupMetadataStatus object to its proto representation.
func CloudidentityGroupDynamicGroupMetadataStatusToProto(o *cloudidentity.GroupDynamicGroupMetadataStatus) *cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatus {
	if o == nil {
		return nil
	}
	p := &cloudidentitypb.CloudidentityGroupDynamicGroupMetadataStatus{}
	p.SetStatus(CloudidentityGroupDynamicGroupMetadataStatusStatusEnumToProto(o.Status))
	p.SetStatusTime(dcl.ValueOrEmptyString(o.StatusTime))
	return p
}

// GroupToProto converts a Group resource to its proto representation.
func GroupToProto(resource *cloudidentity.Group) *cloudidentitypb.CloudidentityGroup {
	p := &cloudidentitypb.CloudidentityGroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetGroupKey(CloudidentityGroupGroupKeyToProto(resource.GroupKey))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDirectMemberCount(dcl.ValueOrEmptyInt64(resource.DirectMemberCount))
	p.SetDirectMemberCountPerType(CloudidentityGroupDirectMemberCountPerTypeToProto(resource.DirectMemberCountPerType))
	p.SetDynamicGroupMetadata(CloudidentityGroupDynamicGroupMetadataToProto(resource.DynamicGroupMetadata))
	p.SetInitialGroupConfig(CloudidentityGroupInitialGroupConfigEnumToProto(resource.InitialGroupConfig))
	sAdditionalGroupKeys := make([]*cloudidentitypb.CloudidentityGroupAdditionalGroupKeys, len(resource.AdditionalGroupKeys))
	for i, r := range resource.AdditionalGroupKeys {
		sAdditionalGroupKeys[i] = CloudidentityGroupAdditionalGroupKeysToProto(&r)
	}
	p.SetAdditionalGroupKeys(sAdditionalGroupKeys)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sDerivedAliases := make([]*cloudidentitypb.CloudidentityGroupDerivedAliases, len(resource.DerivedAliases))
	for i, r := range resource.DerivedAliases {
		sDerivedAliases[i] = CloudidentityGroupDerivedAliasesToProto(&r)
	}
	p.SetDerivedAliases(sDerivedAliases)

	return p
}

// applyGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) applyGroup(ctx context.Context, c *cloudidentity.Client, request *cloudidentitypb.ApplyCloudidentityGroupRequest) (*cloudidentitypb.CloudidentityGroup, error) {
	p := ProtoToGroup(request.GetResource())
	res, err := c.ApplyGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GroupToProto(res)
	return r, nil
}

// applyCloudidentityGroup handles the gRPC request by passing it to the underlying Group Apply() method.
func (s *GroupServer) ApplyCloudidentityGroup(ctx context.Context, request *cloudidentitypb.ApplyCloudidentityGroupRequest) (*cloudidentitypb.CloudidentityGroup, error) {
	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGroup(ctx, cl, request)
}

// DeleteGroup handles the gRPC request by passing it to the underlying Group Delete() method.
func (s *GroupServer) DeleteCloudidentityGroup(ctx context.Context, request *cloudidentitypb.DeleteCloudidentityGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGroup(ctx, ProtoToGroup(request.GetResource()))

}

// ListCloudidentityGroup handles the gRPC request by passing it to the underlying GroupList() method.
func (s *GroupServer) ListCloudidentityGroup(ctx context.Context, request *cloudidentitypb.ListCloudidentityGroupRequest) (*cloudidentitypb.ListCloudidentityGroupResponse, error) {
	cl, err := createConfigGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGroup(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*cloudidentitypb.CloudidentityGroup
	for _, r := range resources.Items {
		rp := GroupToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudidentitypb.ListCloudidentityGroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGroup(ctx context.Context, service_account_file string) (*cloudidentity.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudidentity.NewClient(conf), nil
}
