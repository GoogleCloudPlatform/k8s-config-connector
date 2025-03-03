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
	dataplexpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataplex/dataplex_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataplex"
)

// ZoneServer implements the gRPC interface for Zone.
type ZoneServer struct{}

// ProtoToZoneStateEnum converts a ZoneStateEnum enum from its proto representation.
func ProtoToDataplexZoneStateEnum(e dataplexpb.DataplexZoneStateEnum) *dataplex.ZoneStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexZoneStateEnum_name[int32(e)]; ok {
		e := dataplex.ZoneStateEnum(n[len("DataplexZoneStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneTypeEnum converts a ZoneTypeEnum enum from its proto representation.
func ProtoToDataplexZoneTypeEnum(e dataplexpb.DataplexZoneTypeEnum) *dataplex.ZoneTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexZoneTypeEnum_name[int32(e)]; ok {
		e := dataplex.ZoneTypeEnum(n[len("DataplexZoneTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneResourceSpecLocationTypeEnum converts a ZoneResourceSpecLocationTypeEnum enum from its proto representation.
func ProtoToDataplexZoneResourceSpecLocationTypeEnum(e dataplexpb.DataplexZoneResourceSpecLocationTypeEnum) *dataplex.ZoneResourceSpecLocationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexZoneResourceSpecLocationTypeEnum_name[int32(e)]; ok {
		e := dataplex.ZoneResourceSpecLocationTypeEnum(n[len("DataplexZoneResourceSpecLocationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToZoneDiscoverySpec converts a ZoneDiscoverySpec object from its proto representation.
func ProtoToDataplexZoneDiscoverySpec(p *dataplexpb.DataplexZoneDiscoverySpec) *dataplex.ZoneDiscoverySpec {
	if p == nil {
		return nil
	}
	obj := &dataplex.ZoneDiscoverySpec{
		Enabled:     dcl.Bool(p.GetEnabled()),
		CsvOptions:  ProtoToDataplexZoneDiscoverySpecCsvOptions(p.GetCsvOptions()),
		JsonOptions: ProtoToDataplexZoneDiscoverySpecJsonOptions(p.GetJsonOptions()),
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
func ProtoToDataplexZoneDiscoverySpecCsvOptions(p *dataplexpb.DataplexZoneDiscoverySpecCsvOptions) *dataplex.ZoneDiscoverySpecCsvOptions {
	if p == nil {
		return nil
	}
	obj := &dataplex.ZoneDiscoverySpecCsvOptions{
		HeaderRows:           dcl.Int64OrNil(p.GetHeaderRows()),
		Delimiter:            dcl.StringOrNil(p.GetDelimiter()),
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToZoneDiscoverySpecJsonOptions converts a ZoneDiscoverySpecJsonOptions object from its proto representation.
func ProtoToDataplexZoneDiscoverySpecJsonOptions(p *dataplexpb.DataplexZoneDiscoverySpecJsonOptions) *dataplex.ZoneDiscoverySpecJsonOptions {
	if p == nil {
		return nil
	}
	obj := &dataplex.ZoneDiscoverySpecJsonOptions{
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToZoneResourceSpec converts a ZoneResourceSpec object from its proto representation.
func ProtoToDataplexZoneResourceSpec(p *dataplexpb.DataplexZoneResourceSpec) *dataplex.ZoneResourceSpec {
	if p == nil {
		return nil
	}
	obj := &dataplex.ZoneResourceSpec{
		LocationType: ProtoToDataplexZoneResourceSpecLocationTypeEnum(p.GetLocationType()),
	}
	return obj
}

// ProtoToZoneAssetStatus converts a ZoneAssetStatus object from its proto representation.
func ProtoToDataplexZoneAssetStatus(p *dataplexpb.DataplexZoneAssetStatus) *dataplex.ZoneAssetStatus {
	if p == nil {
		return nil
	}
	obj := &dataplex.ZoneAssetStatus{
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		ActiveAssets:                 dcl.Int64OrNil(p.GetActiveAssets()),
		SecurityPolicyApplyingAssets: dcl.Int64OrNil(p.GetSecurityPolicyApplyingAssets()),
	}
	return obj
}

// ProtoToZone converts a Zone resource from its proto representation.
func ProtoToZone(p *dataplexpb.DataplexZone) *dataplex.Zone {
	obj := &dataplex.Zone{
		Name:          dcl.StringOrNil(p.GetName()),
		DisplayName:   dcl.StringOrNil(p.GetDisplayName()),
		Uid:           dcl.StringOrNil(p.GetUid()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		State:         ProtoToDataplexZoneStateEnum(p.GetState()),
		Type:          ProtoToDataplexZoneTypeEnum(p.GetType()),
		DiscoverySpec: ProtoToDataplexZoneDiscoverySpec(p.GetDiscoverySpec()),
		ResourceSpec:  ProtoToDataplexZoneResourceSpec(p.GetResourceSpec()),
		AssetStatus:   ProtoToDataplexZoneAssetStatus(p.GetAssetStatus()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		Lake:          dcl.StringOrNil(p.GetLake()),
	}
	return obj
}

// ZoneStateEnumToProto converts a ZoneStateEnum enum to its proto representation.
func DataplexZoneStateEnumToProto(e *dataplex.ZoneStateEnum) dataplexpb.DataplexZoneStateEnum {
	if e == nil {
		return dataplexpb.DataplexZoneStateEnum(0)
	}
	if v, ok := dataplexpb.DataplexZoneStateEnum_value["ZoneStateEnum"+string(*e)]; ok {
		return dataplexpb.DataplexZoneStateEnum(v)
	}
	return dataplexpb.DataplexZoneStateEnum(0)
}

// ZoneTypeEnumToProto converts a ZoneTypeEnum enum to its proto representation.
func DataplexZoneTypeEnumToProto(e *dataplex.ZoneTypeEnum) dataplexpb.DataplexZoneTypeEnum {
	if e == nil {
		return dataplexpb.DataplexZoneTypeEnum(0)
	}
	if v, ok := dataplexpb.DataplexZoneTypeEnum_value["ZoneTypeEnum"+string(*e)]; ok {
		return dataplexpb.DataplexZoneTypeEnum(v)
	}
	return dataplexpb.DataplexZoneTypeEnum(0)
}

// ZoneResourceSpecLocationTypeEnumToProto converts a ZoneResourceSpecLocationTypeEnum enum to its proto representation.
func DataplexZoneResourceSpecLocationTypeEnumToProto(e *dataplex.ZoneResourceSpecLocationTypeEnum) dataplexpb.DataplexZoneResourceSpecLocationTypeEnum {
	if e == nil {
		return dataplexpb.DataplexZoneResourceSpecLocationTypeEnum(0)
	}
	if v, ok := dataplexpb.DataplexZoneResourceSpecLocationTypeEnum_value["ZoneResourceSpecLocationTypeEnum"+string(*e)]; ok {
		return dataplexpb.DataplexZoneResourceSpecLocationTypeEnum(v)
	}
	return dataplexpb.DataplexZoneResourceSpecLocationTypeEnum(0)
}

// ZoneDiscoverySpecToProto converts a ZoneDiscoverySpec object to its proto representation.
func DataplexZoneDiscoverySpecToProto(o *dataplex.ZoneDiscoverySpec) *dataplexpb.DataplexZoneDiscoverySpec {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexZoneDiscoverySpec{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetCsvOptions(DataplexZoneDiscoverySpecCsvOptionsToProto(o.CsvOptions))
	p.SetJsonOptions(DataplexZoneDiscoverySpecJsonOptionsToProto(o.JsonOptions))
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
func DataplexZoneDiscoverySpecCsvOptionsToProto(o *dataplex.ZoneDiscoverySpecCsvOptions) *dataplexpb.DataplexZoneDiscoverySpecCsvOptions {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexZoneDiscoverySpecCsvOptions{}
	p.SetHeaderRows(dcl.ValueOrEmptyInt64(o.HeaderRows))
	p.SetDelimiter(dcl.ValueOrEmptyString(o.Delimiter))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// ZoneDiscoverySpecJsonOptionsToProto converts a ZoneDiscoverySpecJsonOptions object to its proto representation.
func DataplexZoneDiscoverySpecJsonOptionsToProto(o *dataplex.ZoneDiscoverySpecJsonOptions) *dataplexpb.DataplexZoneDiscoverySpecJsonOptions {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexZoneDiscoverySpecJsonOptions{}
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// ZoneResourceSpecToProto converts a ZoneResourceSpec object to its proto representation.
func DataplexZoneResourceSpecToProto(o *dataplex.ZoneResourceSpec) *dataplexpb.DataplexZoneResourceSpec {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexZoneResourceSpec{}
	p.SetLocationType(DataplexZoneResourceSpecLocationTypeEnumToProto(o.LocationType))
	return p
}

// ZoneAssetStatusToProto converts a ZoneAssetStatus object to its proto representation.
func DataplexZoneAssetStatusToProto(o *dataplex.ZoneAssetStatus) *dataplexpb.DataplexZoneAssetStatus {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexZoneAssetStatus{}
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetActiveAssets(dcl.ValueOrEmptyInt64(o.ActiveAssets))
	p.SetSecurityPolicyApplyingAssets(dcl.ValueOrEmptyInt64(o.SecurityPolicyApplyingAssets))
	return p
}

// ZoneToProto converts a Zone resource to its proto representation.
func ZoneToProto(resource *dataplex.Zone) *dataplexpb.DataplexZone {
	p := &dataplexpb.DataplexZone{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexZoneStateEnumToProto(resource.State))
	p.SetType(DataplexZoneTypeEnumToProto(resource.Type))
	p.SetDiscoverySpec(DataplexZoneDiscoverySpecToProto(resource.DiscoverySpec))
	p.SetResourceSpec(DataplexZoneResourceSpecToProto(resource.ResourceSpec))
	p.SetAssetStatus(DataplexZoneAssetStatusToProto(resource.AssetStatus))
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
func (s *ZoneServer) applyZone(ctx context.Context, c *dataplex.Client, request *dataplexpb.ApplyDataplexZoneRequest) (*dataplexpb.DataplexZone, error) {
	p := ProtoToZone(request.GetResource())
	res, err := c.ApplyZone(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ZoneToProto(res)
	return r, nil
}

// applyDataplexZone handles the gRPC request by passing it to the underlying Zone Apply() method.
func (s *ZoneServer) ApplyDataplexZone(ctx context.Context, request *dataplexpb.ApplyDataplexZoneRequest) (*dataplexpb.DataplexZone, error) {
	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyZone(ctx, cl, request)
}

// DeleteZone handles the gRPC request by passing it to the underlying Zone Delete() method.
func (s *ZoneServer) DeleteDataplexZone(ctx context.Context, request *dataplexpb.DeleteDataplexZoneRequest) (*emptypb.Empty, error) {

	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteZone(ctx, ProtoToZone(request.GetResource()))

}

// ListDataplexZone handles the gRPC request by passing it to the underlying ZoneList() method.
func (s *ZoneServer) ListDataplexZone(ctx context.Context, request *dataplexpb.ListDataplexZoneRequest) (*dataplexpb.ListDataplexZoneResponse, error) {
	cl, err := createConfigZone(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListZone(ctx, request.GetProject(), request.GetLocation(), request.GetLake())
	if err != nil {
		return nil, err
	}
	var protos []*dataplexpb.DataplexZone
	for _, r := range resources.Items {
		rp := ZoneToProto(r)
		protos = append(protos, rp)
	}
	p := &dataplexpb.ListDataplexZoneResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigZone(ctx context.Context, service_account_file string) (*dataplex.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataplex.NewClient(conf), nil
}
