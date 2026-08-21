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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataplex/alpha/dataplex_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex/alpha"
)

// ZoneServer implements the gRPC interface for Zone.
type ZoneServer struct{}

// ProtoToZoneStateEnum converts a ZoneStateEnum enum from its proto representation.
func ProtoToDataplexAlphaZoneStateEnum(e alphapb.DataplexAlphaZoneStateEnum) *alpha.ZoneStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaZoneStateEnum_name[int32(e)]; ok {
		e := alpha.ZoneStateEnum(n[len("DataplexAlphaZoneStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneTypeEnum converts a ZoneTypeEnum enum from its proto representation.
func ProtoToDataplexAlphaZoneTypeEnum(e alphapb.DataplexAlphaZoneTypeEnum) *alpha.ZoneTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaZoneTypeEnum_name[int32(e)]; ok {
		e := alpha.ZoneTypeEnum(n[len("DataplexAlphaZoneTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneResourceSpecLocationTypeEnum converts a ZoneResourceSpecLocationTypeEnum enum from its proto representation.
func ProtoToDataplexAlphaZoneResourceSpecLocationTypeEnum(e alphapb.DataplexAlphaZoneResourceSpecLocationTypeEnum) *alpha.ZoneResourceSpecLocationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaZoneResourceSpecLocationTypeEnum_name[int32(e)]; ok {
		e := alpha.ZoneResourceSpecLocationTypeEnum(n[len("DataplexAlphaZoneResourceSpecLocationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneDiscoverySpec converts a ZoneDiscoverySpec object from its proto representation.
func ProtoToDataplexAlphaZoneDiscoverySpec(p *alphapb.DataplexAlphaZoneDiscoverySpec) *alpha.ZoneDiscoverySpec {
	if p == nil {
		return nil
	}
	obj := &alpha.ZoneDiscoverySpec{
		Enabled:     dcl.Bool(p.GetEnabled()),
		CsvOptions:  ProtoToDataplexAlphaZoneDiscoverySpecCsvOptions(p.GetCsvOptions()),
		JsonOptions: ProtoToDataplexAlphaZoneDiscoverySpecJsonOptions(p.GetJsonOptions()),
		Schedule:    dcl.StringOrNil(p.GetSchedule()),
	}
	for _, r := range p.GetIncludePatterns() {
		obj.IncludePatterns = append(obj.IncludePatterns, r)
	}
	for _, r := range p.GetExcludePatterns() {
		obj.ExcludePatterns = append(obj.ExcludePatterns, r)
	}
	return obj
}

// ProtoToZoneDiscoverySpecCsvOptions converts a ZoneDiscoverySpecCsvOptions object from its proto representation.
func ProtoToDataplexAlphaZoneDiscoverySpecCsvOptions(p *alphapb.DataplexAlphaZoneDiscoverySpecCsvOptions) *alpha.ZoneDiscoverySpecCsvOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.ZoneDiscoverySpecCsvOptions{
		HeaderRows:           dcl.Int64OrNil(p.GetHeaderRows()),
		Delimiter:            dcl.StringOrNil(p.GetDelimiter()),
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToZoneDiscoverySpecJsonOptions converts a ZoneDiscoverySpecJsonOptions object from its proto representation.
func ProtoToDataplexAlphaZoneDiscoverySpecJsonOptions(p *alphapb.DataplexAlphaZoneDiscoverySpecJsonOptions) *alpha.ZoneDiscoverySpecJsonOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.ZoneDiscoverySpecJsonOptions{
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToZoneResourceSpec converts a ZoneResourceSpec object from its proto representation.
func ProtoToDataplexAlphaZoneResourceSpec(p *alphapb.DataplexAlphaZoneResourceSpec) *alpha.ZoneResourceSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.ZoneResourceSpec{
		LocationType: ProtoToDataplexAlphaZoneResourceSpecLocationTypeEnum(p.GetLocationType()),
	}
	return obj
}

// ProtoToZoneAssetStatus converts a ZoneAssetStatus object from its proto representation.
func ProtoToDataplexAlphaZoneAssetStatus(p *alphapb.DataplexAlphaZoneAssetStatus) *alpha.ZoneAssetStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.ZoneAssetStatus{
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		ActiveAssets:                 dcl.Int64OrNil(p.GetActiveAssets()),
		SecurityPolicyApplyingAssets: dcl.Int64OrNil(p.GetSecurityPolicyApplyingAssets()),
	}
	return obj
}

// ProtoToZone converts a Zone resource from its proto representation.
func ProtoToZone(p *alphapb.DataplexAlphaZone) *alpha.Zone {
	obj := &alpha.Zone{
		Name:          dcl.StringOrNil(p.GetName()),
		DisplayName:   dcl.StringOrNil(p.GetDisplayName()),
		Uid:           dcl.StringOrNil(p.GetUid()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		State:         ProtoToDataplexAlphaZoneStateEnum(p.GetState()),
		Type:          ProtoToDataplexAlphaZoneTypeEnum(p.GetType()),
		DiscoverySpec: ProtoToDataplexAlphaZoneDiscoverySpec(p.GetDiscoverySpec()),
		ResourceSpec:  ProtoToDataplexAlphaZoneResourceSpec(p.GetResourceSpec()),
		AssetStatus:   ProtoToDataplexAlphaZoneAssetStatus(p.GetAssetStatus()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		Lake:          dcl.StringOrNil(p.GetLake()),
	}
	return obj
}

// ZoneStateEnumToProto converts a ZoneStateEnum enum to its proto representation.
func DataplexAlphaZoneStateEnumToProto(e *alpha.ZoneStateEnum) alphapb.DataplexAlphaZoneStateEnum {
	if e == nil {
		return alphapb.DataplexAlphaZoneStateEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaZoneStateEnum_value["ZoneStateEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaZoneStateEnum(v)
	}
	return alphapb.DataplexAlphaZoneStateEnum(0)
}

// ZoneTypeEnumToProto converts a ZoneTypeEnum enum to its proto representation.
func DataplexAlphaZoneTypeEnumToProto(e *alpha.ZoneTypeEnum) alphapb.DataplexAlphaZoneTypeEnum {
	if e == nil {
		return alphapb.DataplexAlphaZoneTypeEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaZoneTypeEnum_value["ZoneTypeEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaZoneTypeEnum(v)
	}
	return alphapb.DataplexAlphaZoneTypeEnum(0)
}

// ZoneResourceSpecLocationTypeEnumToProto converts a ZoneResourceSpecLocationTypeEnum enum to its proto representation.
func DataplexAlphaZoneResourceSpecLocationTypeEnumToProto(e *alpha.ZoneResourceSpecLocationTypeEnum) alphapb.DataplexAlphaZoneResourceSpecLocationTypeEnum {
	if e == nil {
		return alphapb.DataplexAlphaZoneResourceSpecLocationTypeEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaZoneResourceSpecLocationTypeEnum_value["ZoneResourceSpecLocationTypeEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaZoneResourceSpecLocationTypeEnum(v)
	}
	return alphapb.DataplexAlphaZoneResourceSpecLocationTypeEnum(0)
}

// ZoneDiscoverySpecToProto converts a ZoneDiscoverySpec object to its proto representation.
func DataplexAlphaZoneDiscoverySpecToProto(o *alpha.ZoneDiscoverySpec) *alphapb.DataplexAlphaZoneDiscoverySpec {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaZoneDiscoverySpec{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetCsvOptions(DataplexAlphaZoneDiscoverySpecCsvOptionsToProto(o.CsvOptions))
	p.SetJsonOptions(DataplexAlphaZoneDiscoverySpecJsonOptionsToProto(o.JsonOptions))
	p.SetSchedule(dcl.ValueOrEmptyString(o.Schedule))
	sIncludePatterns := make([]string, len(o.IncludePatterns))
	for i, r := range o.IncludePatterns {
		sIncludePatterns[i] = r
	}
	p.SetIncludePatterns(sIncludePatterns)
	sExcludePatterns := make([]string, len(o.ExcludePatterns))
	for i, r := range o.ExcludePatterns {
		sExcludePatterns[i] = r
	}
	p.SetExcludePatterns(sExcludePatterns)
	return p
}

// ZoneDiscoverySpecCsvOptionsToProto converts a ZoneDiscoverySpecCsvOptions object to its proto representation.
func DataplexAlphaZoneDiscoverySpecCsvOptionsToProto(o *alpha.ZoneDiscoverySpecCsvOptions) *alphapb.DataplexAlphaZoneDiscoverySpecCsvOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaZoneDiscoverySpecCsvOptions{}
	p.SetHeaderRows(dcl.ValueOrEmptyInt64(o.HeaderRows))
	p.SetDelimiter(dcl.ValueOrEmptyString(o.Delimiter))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// ZoneDiscoverySpecJsonOptionsToProto converts a ZoneDiscoverySpecJsonOptions object to its proto representation.
func DataplexAlphaZoneDiscoverySpecJsonOptionsToProto(o *alpha.ZoneDiscoverySpecJsonOptions) *alphapb.DataplexAlphaZoneDiscoverySpecJsonOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaZoneDiscoverySpecJsonOptions{}
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// ZoneResourceSpecToProto converts a ZoneResourceSpec object to its proto representation.
func DataplexAlphaZoneResourceSpecToProto(o *alpha.ZoneResourceSpec) *alphapb.DataplexAlphaZoneResourceSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaZoneResourceSpec{}
	p.SetLocationType(DataplexAlphaZoneResourceSpecLocationTypeEnumToProto(o.LocationType))
	return p
}

// ZoneAssetStatusToProto converts a ZoneAssetStatus object to its proto representation.
func DataplexAlphaZoneAssetStatusToProto(o *alpha.ZoneAssetStatus) *alphapb.DataplexAlphaZoneAssetStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaZoneAssetStatus{}
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetActiveAssets(dcl.ValueOrEmptyInt64(o.ActiveAssets))
	p.SetSecurityPolicyApplyingAssets(dcl.ValueOrEmptyInt64(o.SecurityPolicyApplyingAssets))
	return p
}

// ZoneToProto converts a Zone resource to its proto representation.
func ZoneToProto(resource *alpha.Zone) *alphapb.DataplexAlphaZone {
	p := &alphapb.DataplexAlphaZone{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexAlphaZoneStateEnumToProto(resource.State))
	p.SetType(DataplexAlphaZoneTypeEnumToProto(resource.Type))
	p.SetDiscoverySpec(DataplexAlphaZoneDiscoverySpecToProto(resource.DiscoverySpec))
	p.SetResourceSpec(DataplexAlphaZoneResourceSpecToProto(resource.ResourceSpec))
	p.SetAssetStatus(DataplexAlphaZoneAssetStatusToProto(resource.AssetStatus))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetLake(dcl.ValueOrEmptyString(resource.Lake))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyZone handles the gRPC request by passing it to the underlying Zone Apply() method.
func (s *ZoneServer) applyZone(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataplexAlphaZoneRequest) (*alphapb.DataplexAlphaZone, error) {
	p := ProtoToZone(request.GetResource())
	res, err := c.ApplyZone(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ZoneToProto(res)
	return r, nil
}

// applyDataplexAlphaZone handles the gRPC request by passing it to the underlying Zone Apply() method.
func (s *ZoneServer) ApplyDataplexAlphaZone(ctx context.Context, request *alphapb.ApplyDataplexAlphaZoneRequest) (*alphapb.DataplexAlphaZone, error) {
	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyZone(ctx, cl, request)
}

// DeleteZone handles the gRPC request by passing it to the underlying Zone Delete() method.
func (s *ZoneServer) DeleteDataplexAlphaZone(ctx context.Context, request *alphapb.DeleteDataplexAlphaZoneRequest) (*emptypb.Empty, error) {

	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteZone(ctx, ProtoToZone(request.GetResource()))

}

// ListDataplexAlphaZone handles the gRPC request by passing it to the underlying ZoneList() method.
func (s *ZoneServer) ListDataplexAlphaZone(ctx context.Context, request *alphapb.ListDataplexAlphaZoneRequest) (*alphapb.ListDataplexAlphaZoneResponse, error) {
	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListZone(ctx, request.GetProject(), request.GetLocation(), request.GetLake())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataplexAlphaZone
	for _, r := range resources.Items {
		rp := ZoneToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDataplexAlphaZoneResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigZone(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
