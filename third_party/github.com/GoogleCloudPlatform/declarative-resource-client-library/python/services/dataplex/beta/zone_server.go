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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataplex/beta/dataplex_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex/beta"
)

// ZoneServer implements the gRPC interface for Zone.
type ZoneServer struct{}

// ProtoToZoneStateEnum converts a ZoneStateEnum enum from its proto representation.
func ProtoToDataplexBetaZoneStateEnum(e betapb.DataplexBetaZoneStateEnum) *beta.ZoneStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaZoneStateEnum_name[int32(e)]; ok {
		e := beta.ZoneStateEnum(n[len("DataplexBetaZoneStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneTypeEnum converts a ZoneTypeEnum enum from its proto representation.
func ProtoToDataplexBetaZoneTypeEnum(e betapb.DataplexBetaZoneTypeEnum) *beta.ZoneTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaZoneTypeEnum_name[int32(e)]; ok {
		e := beta.ZoneTypeEnum(n[len("DataplexBetaZoneTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneResourceSpecLocationTypeEnum converts a ZoneResourceSpecLocationTypeEnum enum from its proto representation.
func ProtoToDataplexBetaZoneResourceSpecLocationTypeEnum(e betapb.DataplexBetaZoneResourceSpecLocationTypeEnum) *beta.ZoneResourceSpecLocationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaZoneResourceSpecLocationTypeEnum_name[int32(e)]; ok {
		e := beta.ZoneResourceSpecLocationTypeEnum(n[len("DataplexBetaZoneResourceSpecLocationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneDiscoverySpec converts a ZoneDiscoverySpec object from its proto representation.
func ProtoToDataplexBetaZoneDiscoverySpec(p *betapb.DataplexBetaZoneDiscoverySpec) *beta.ZoneDiscoverySpec {
	if p == nil {
		return nil
	}
	obj := &beta.ZoneDiscoverySpec{
		Enabled:     dcl.Bool(p.GetEnabled()),
		CsvOptions:  ProtoToDataplexBetaZoneDiscoverySpecCsvOptions(p.GetCsvOptions()),
		JsonOptions: ProtoToDataplexBetaZoneDiscoverySpecJsonOptions(p.GetJsonOptions()),
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
func ProtoToDataplexBetaZoneDiscoverySpecCsvOptions(p *betapb.DataplexBetaZoneDiscoverySpecCsvOptions) *beta.ZoneDiscoverySpecCsvOptions {
	if p == nil {
		return nil
	}
	obj := &beta.ZoneDiscoverySpecCsvOptions{
		HeaderRows:           dcl.Int64OrNil(p.GetHeaderRows()),
		Delimiter:            dcl.StringOrNil(p.GetDelimiter()),
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToZoneDiscoverySpecJsonOptions converts a ZoneDiscoverySpecJsonOptions object from its proto representation.
func ProtoToDataplexBetaZoneDiscoverySpecJsonOptions(p *betapb.DataplexBetaZoneDiscoverySpecJsonOptions) *beta.ZoneDiscoverySpecJsonOptions {
	if p == nil {
		return nil
	}
	obj := &beta.ZoneDiscoverySpecJsonOptions{
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToZoneResourceSpec converts a ZoneResourceSpec object from its proto representation.
func ProtoToDataplexBetaZoneResourceSpec(p *betapb.DataplexBetaZoneResourceSpec) *beta.ZoneResourceSpec {
	if p == nil {
		return nil
	}
	obj := &beta.ZoneResourceSpec{
		LocationType: ProtoToDataplexBetaZoneResourceSpecLocationTypeEnum(p.GetLocationType()),
	}
	return obj
}

// ProtoToZoneAssetStatus converts a ZoneAssetStatus object from its proto representation.
func ProtoToDataplexBetaZoneAssetStatus(p *betapb.DataplexBetaZoneAssetStatus) *beta.ZoneAssetStatus {
	if p == nil {
		return nil
	}
	obj := &beta.ZoneAssetStatus{
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		ActiveAssets:                 dcl.Int64OrNil(p.GetActiveAssets()),
		SecurityPolicyApplyingAssets: dcl.Int64OrNil(p.GetSecurityPolicyApplyingAssets()),
	}
	return obj
}

// ProtoToZone converts a Zone resource from its proto representation.
func ProtoToZone(p *betapb.DataplexBetaZone) *beta.Zone {
	obj := &beta.Zone{
		Name:          dcl.StringOrNil(p.GetName()),
		DisplayName:   dcl.StringOrNil(p.GetDisplayName()),
		Uid:           dcl.StringOrNil(p.GetUid()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		State:         ProtoToDataplexBetaZoneStateEnum(p.GetState()),
		Type:          ProtoToDataplexBetaZoneTypeEnum(p.GetType()),
		DiscoverySpec: ProtoToDataplexBetaZoneDiscoverySpec(p.GetDiscoverySpec()),
		ResourceSpec:  ProtoToDataplexBetaZoneResourceSpec(p.GetResourceSpec()),
		AssetStatus:   ProtoToDataplexBetaZoneAssetStatus(p.GetAssetStatus()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		Lake:          dcl.StringOrNil(p.GetLake()),
	}
	return obj
}

// ZoneStateEnumToProto converts a ZoneStateEnum enum to its proto representation.
func DataplexBetaZoneStateEnumToProto(e *beta.ZoneStateEnum) betapb.DataplexBetaZoneStateEnum {
	if e == nil {
		return betapb.DataplexBetaZoneStateEnum(0)
	}
	if v, ok := betapb.DataplexBetaZoneStateEnum_value["ZoneStateEnum"+string(*e)]; ok {
		return betapb.DataplexBetaZoneStateEnum(v)
	}
	return betapb.DataplexBetaZoneStateEnum(0)
}

// ZoneTypeEnumToProto converts a ZoneTypeEnum enum to its proto representation.
func DataplexBetaZoneTypeEnumToProto(e *beta.ZoneTypeEnum) betapb.DataplexBetaZoneTypeEnum {
	if e == nil {
		return betapb.DataplexBetaZoneTypeEnum(0)
	}
	if v, ok := betapb.DataplexBetaZoneTypeEnum_value["ZoneTypeEnum"+string(*e)]; ok {
		return betapb.DataplexBetaZoneTypeEnum(v)
	}
	return betapb.DataplexBetaZoneTypeEnum(0)
}

// ZoneResourceSpecLocationTypeEnumToProto converts a ZoneResourceSpecLocationTypeEnum enum to its proto representation.
func DataplexBetaZoneResourceSpecLocationTypeEnumToProto(e *beta.ZoneResourceSpecLocationTypeEnum) betapb.DataplexBetaZoneResourceSpecLocationTypeEnum {
	if e == nil {
		return betapb.DataplexBetaZoneResourceSpecLocationTypeEnum(0)
	}
	if v, ok := betapb.DataplexBetaZoneResourceSpecLocationTypeEnum_value["ZoneResourceSpecLocationTypeEnum"+string(*e)]; ok {
		return betapb.DataplexBetaZoneResourceSpecLocationTypeEnum(v)
	}
	return betapb.DataplexBetaZoneResourceSpecLocationTypeEnum(0)
}

// ZoneDiscoverySpecToProto converts a ZoneDiscoverySpec object to its proto representation.
func DataplexBetaZoneDiscoverySpecToProto(o *beta.ZoneDiscoverySpec) *betapb.DataplexBetaZoneDiscoverySpec {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaZoneDiscoverySpec{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetCsvOptions(DataplexBetaZoneDiscoverySpecCsvOptionsToProto(o.CsvOptions))
	p.SetJsonOptions(DataplexBetaZoneDiscoverySpecJsonOptionsToProto(o.JsonOptions))
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
func DataplexBetaZoneDiscoverySpecCsvOptionsToProto(o *beta.ZoneDiscoverySpecCsvOptions) *betapb.DataplexBetaZoneDiscoverySpecCsvOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaZoneDiscoverySpecCsvOptions{}
	p.SetHeaderRows(dcl.ValueOrEmptyInt64(o.HeaderRows))
	p.SetDelimiter(dcl.ValueOrEmptyString(o.Delimiter))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// ZoneDiscoverySpecJsonOptionsToProto converts a ZoneDiscoverySpecJsonOptions object to its proto representation.
func DataplexBetaZoneDiscoverySpecJsonOptionsToProto(o *beta.ZoneDiscoverySpecJsonOptions) *betapb.DataplexBetaZoneDiscoverySpecJsonOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaZoneDiscoverySpecJsonOptions{}
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// ZoneResourceSpecToProto converts a ZoneResourceSpec object to its proto representation.
func DataplexBetaZoneResourceSpecToProto(o *beta.ZoneResourceSpec) *betapb.DataplexBetaZoneResourceSpec {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaZoneResourceSpec{}
	p.SetLocationType(DataplexBetaZoneResourceSpecLocationTypeEnumToProto(o.LocationType))
	return p
}

// ZoneAssetStatusToProto converts a ZoneAssetStatus object to its proto representation.
func DataplexBetaZoneAssetStatusToProto(o *beta.ZoneAssetStatus) *betapb.DataplexBetaZoneAssetStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaZoneAssetStatus{}
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetActiveAssets(dcl.ValueOrEmptyInt64(o.ActiveAssets))
	p.SetSecurityPolicyApplyingAssets(dcl.ValueOrEmptyInt64(o.SecurityPolicyApplyingAssets))
	return p
}

// ZoneToProto converts a Zone resource to its proto representation.
func ZoneToProto(resource *beta.Zone) *betapb.DataplexBetaZone {
	p := &betapb.DataplexBetaZone{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexBetaZoneStateEnumToProto(resource.State))
	p.SetType(DataplexBetaZoneTypeEnumToProto(resource.Type))
	p.SetDiscoverySpec(DataplexBetaZoneDiscoverySpecToProto(resource.DiscoverySpec))
	p.SetResourceSpec(DataplexBetaZoneResourceSpecToProto(resource.ResourceSpec))
	p.SetAssetStatus(DataplexBetaZoneAssetStatusToProto(resource.AssetStatus))
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
func (s *ZoneServer) applyZone(ctx context.Context, c *beta.Client, request *betapb.ApplyDataplexBetaZoneRequest) (*betapb.DataplexBetaZone, error) {
	p := ProtoToZone(request.GetResource())
	res, err := c.ApplyZone(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ZoneToProto(res)
	return r, nil
}

// applyDataplexBetaZone handles the gRPC request by passing it to the underlying Zone Apply() method.
func (s *ZoneServer) ApplyDataplexBetaZone(ctx context.Context, request *betapb.ApplyDataplexBetaZoneRequest) (*betapb.DataplexBetaZone, error) {
	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyZone(ctx, cl, request)
}

// DeleteZone handles the gRPC request by passing it to the underlying Zone Delete() method.
func (s *ZoneServer) DeleteDataplexBetaZone(ctx context.Context, request *betapb.DeleteDataplexBetaZoneRequest) (*emptypb.Empty, error) {

	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteZone(ctx, ProtoToZone(request.GetResource()))

}

// ListDataplexBetaZone handles the gRPC request by passing it to the underlying ZoneList() method.
func (s *ZoneServer) ListDataplexBetaZone(ctx context.Context, request *betapb.ListDataplexBetaZoneRequest) (*betapb.ListDataplexBetaZoneResponse, error) {
	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListZone(ctx, request.GetProject(), request.GetLocation(), request.GetLake())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataplexBetaZone
	for _, r := range resources.Items {
		rp := ZoneToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDataplexBetaZoneResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigZone(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
