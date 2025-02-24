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

// AssetServer implements the gRPC interface for Asset.
type AssetServer struct{}

// ProtoToAssetStateEnum converts a AssetStateEnum enum from its proto representation.
func ProtoToDataplexAssetStateEnum(e dataplexpb.DataplexAssetStateEnum) *dataplex.AssetStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexAssetStateEnum_name[int32(e)]; ok {
		e := dataplex.AssetStateEnum(n[len("DataplexAssetStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpecTypeEnum converts a AssetResourceSpecTypeEnum enum from its proto representation.
func ProtoToDataplexAssetResourceSpecTypeEnum(e dataplexpb.DataplexAssetResourceSpecTypeEnum) *dataplex.AssetResourceSpecTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexAssetResourceSpecTypeEnum_name[int32(e)]; ok {
		e := dataplex.AssetResourceSpecTypeEnum(n[len("DataplexAssetResourceSpecTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpecReadAccessModeEnum converts a AssetResourceSpecReadAccessModeEnum enum from its proto representation.
func ProtoToDataplexAssetResourceSpecReadAccessModeEnum(e dataplexpb.DataplexAssetResourceSpecReadAccessModeEnum) *dataplex.AssetResourceSpecReadAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexAssetResourceSpecReadAccessModeEnum_name[int32(e)]; ok {
		e := dataplex.AssetResourceSpecReadAccessModeEnum(n[len("DataplexAssetResourceSpecReadAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceStatusStateEnum converts a AssetResourceStatusStateEnum enum from its proto representation.
func ProtoToDataplexAssetResourceStatusStateEnum(e dataplexpb.DataplexAssetResourceStatusStateEnum) *dataplex.AssetResourceStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexAssetResourceStatusStateEnum_name[int32(e)]; ok {
		e := dataplex.AssetResourceStatusStateEnum(n[len("DataplexAssetResourceStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetSecurityStatusStateEnum converts a AssetSecurityStatusStateEnum enum from its proto representation.
func ProtoToDataplexAssetSecurityStatusStateEnum(e dataplexpb.DataplexAssetSecurityStatusStateEnum) *dataplex.AssetSecurityStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexAssetSecurityStatusStateEnum_name[int32(e)]; ok {
		e := dataplex.AssetSecurityStatusStateEnum(n[len("DataplexAssetSecurityStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetDiscoveryStatusStateEnum converts a AssetDiscoveryStatusStateEnum enum from its proto representation.
func ProtoToDataplexAssetDiscoveryStatusStateEnum(e dataplexpb.DataplexAssetDiscoveryStatusStateEnum) *dataplex.AssetDiscoveryStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataplexpb.DataplexAssetDiscoveryStatusStateEnum_name[int32(e)]; ok {
		e := dataplex.AssetDiscoveryStatusStateEnum(n[len("DataplexAssetDiscoveryStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpec converts a AssetResourceSpec object from its proto representation.
func ProtoToDataplexAssetResourceSpec(p *dataplexpb.DataplexAssetResourceSpec) *dataplex.AssetResourceSpec {
	if p == nil {
		return nil
	}
	obj := &dataplex.AssetResourceSpec{
		Name:           dcl.StringOrNil(p.GetName()),
		Type:           ProtoToDataplexAssetResourceSpecTypeEnum(p.GetType()),
		ReadAccessMode: ProtoToDataplexAssetResourceSpecReadAccessModeEnum(p.GetReadAccessMode()),
	}
	return obj
}

// ProtoToAssetResourceStatus converts a AssetResourceStatus object from its proto representation.
func ProtoToDataplexAssetResourceStatus(p *dataplexpb.DataplexAssetResourceStatus) *dataplex.AssetResourceStatus {
	if p == nil {
		return nil
	}
	obj := &dataplex.AssetResourceStatus{
		State:      ProtoToDataplexAssetResourceStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToAssetSecurityStatus converts a AssetSecurityStatus object from its proto representation.
func ProtoToDataplexAssetSecurityStatus(p *dataplexpb.DataplexAssetSecurityStatus) *dataplex.AssetSecurityStatus {
	if p == nil {
		return nil
	}
	obj := &dataplex.AssetSecurityStatus{
		State:      ProtoToDataplexAssetSecurityStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToAssetDiscoverySpec converts a AssetDiscoverySpec object from its proto representation.
func ProtoToDataplexAssetDiscoverySpec(p *dataplexpb.DataplexAssetDiscoverySpec) *dataplex.AssetDiscoverySpec {
	if p == nil {
		return nil
	}
	obj := &dataplex.AssetDiscoverySpec{
		Enabled:     dcl.Bool(p.GetEnabled()),
		CsvOptions:  ProtoToDataplexAssetDiscoverySpecCsvOptions(p.GetCsvOptions()),
		JsonOptions: ProtoToDataplexAssetDiscoverySpecJsonOptions(p.GetJsonOptions()),
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

// ProtoToAssetDiscoverySpecCsvOptions converts a AssetDiscoverySpecCsvOptions object from its proto representation.
func ProtoToDataplexAssetDiscoverySpecCsvOptions(p *dataplexpb.DataplexAssetDiscoverySpecCsvOptions) *dataplex.AssetDiscoverySpecCsvOptions {
	if p == nil {
		return nil
	}
	obj := &dataplex.AssetDiscoverySpecCsvOptions{
		HeaderRows:           dcl.Int64OrNil(p.GetHeaderRows()),
		Delimiter:            dcl.StringOrNil(p.GetDelimiter()),
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToAssetDiscoverySpecJsonOptions converts a AssetDiscoverySpecJsonOptions object from its proto representation.
func ProtoToDataplexAssetDiscoverySpecJsonOptions(p *dataplexpb.DataplexAssetDiscoverySpecJsonOptions) *dataplex.AssetDiscoverySpecJsonOptions {
	if p == nil {
		return nil
	}
	obj := &dataplex.AssetDiscoverySpecJsonOptions{
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToAssetDiscoveryStatus converts a AssetDiscoveryStatus object from its proto representation.
func ProtoToDataplexAssetDiscoveryStatus(p *dataplexpb.DataplexAssetDiscoveryStatus) *dataplex.AssetDiscoveryStatus {
	if p == nil {
		return nil
	}
	obj := &dataplex.AssetDiscoveryStatus{
		State:           ProtoToDataplexAssetDiscoveryStatusStateEnum(p.GetState()),
		Message:         dcl.StringOrNil(p.GetMessage()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		LastRunTime:     dcl.StringOrNil(p.GetLastRunTime()),
		Stats:           ProtoToDataplexAssetDiscoveryStatusStats(p.GetStats()),
		LastRunDuration: dcl.StringOrNil(p.GetLastRunDuration()),
	}
	return obj
}

// ProtoToAssetDiscoveryStatusStats converts a AssetDiscoveryStatusStats object from its proto representation.
func ProtoToDataplexAssetDiscoveryStatusStats(p *dataplexpb.DataplexAssetDiscoveryStatusStats) *dataplex.AssetDiscoveryStatusStats {
	if p == nil {
		return nil
	}
	obj := &dataplex.AssetDiscoveryStatusStats{
		DataItems: dcl.Int64OrNil(p.GetDataItems()),
		DataSize:  dcl.Int64OrNil(p.GetDataSize()),
		Tables:    dcl.Int64OrNil(p.GetTables()),
		Filesets:  dcl.Int64OrNil(p.GetFilesets()),
	}
	return obj
}

// ProtoToAsset converts a Asset resource from its proto representation.
func ProtoToAsset(p *dataplexpb.DataplexAsset) *dataplex.Asset {
	obj := &dataplex.Asset{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToDataplexAssetStateEnum(p.GetState()),
		ResourceSpec:    ProtoToDataplexAssetResourceSpec(p.GetResourceSpec()),
		ResourceStatus:  ProtoToDataplexAssetResourceStatus(p.GetResourceStatus()),
		SecurityStatus:  ProtoToDataplexAssetSecurityStatus(p.GetSecurityStatus()),
		DiscoverySpec:   ProtoToDataplexAssetDiscoverySpec(p.GetDiscoverySpec()),
		DiscoveryStatus: ProtoToDataplexAssetDiscoveryStatus(p.GetDiscoveryStatus()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
		Lake:            dcl.StringOrNil(p.GetLake()),
		DataplexZone:    dcl.StringOrNil(p.GetDataplexZone()),
	}
	return obj
}

// AssetStateEnumToProto converts a AssetStateEnum enum to its proto representation.
func DataplexAssetStateEnumToProto(e *dataplex.AssetStateEnum) dataplexpb.DataplexAssetStateEnum {
	if e == nil {
		return dataplexpb.DataplexAssetStateEnum(0)
	}
	if v, ok := dataplexpb.DataplexAssetStateEnum_value["AssetStateEnum"+string(*e)]; ok {
		return dataplexpb.DataplexAssetStateEnum(v)
	}
	return dataplexpb.DataplexAssetStateEnum(0)
}

// AssetResourceSpecTypeEnumToProto converts a AssetResourceSpecTypeEnum enum to its proto representation.
func DataplexAssetResourceSpecTypeEnumToProto(e *dataplex.AssetResourceSpecTypeEnum) dataplexpb.DataplexAssetResourceSpecTypeEnum {
	if e == nil {
		return dataplexpb.DataplexAssetResourceSpecTypeEnum(0)
	}
	if v, ok := dataplexpb.DataplexAssetResourceSpecTypeEnum_value["AssetResourceSpecTypeEnum"+string(*e)]; ok {
		return dataplexpb.DataplexAssetResourceSpecTypeEnum(v)
	}
	return dataplexpb.DataplexAssetResourceSpecTypeEnum(0)
}

// AssetResourceSpecReadAccessModeEnumToProto converts a AssetResourceSpecReadAccessModeEnum enum to its proto representation.
func DataplexAssetResourceSpecReadAccessModeEnumToProto(e *dataplex.AssetResourceSpecReadAccessModeEnum) dataplexpb.DataplexAssetResourceSpecReadAccessModeEnum {
	if e == nil {
		return dataplexpb.DataplexAssetResourceSpecReadAccessModeEnum(0)
	}
	if v, ok := dataplexpb.DataplexAssetResourceSpecReadAccessModeEnum_value["AssetResourceSpecReadAccessModeEnum"+string(*e)]; ok {
		return dataplexpb.DataplexAssetResourceSpecReadAccessModeEnum(v)
	}
	return dataplexpb.DataplexAssetResourceSpecReadAccessModeEnum(0)
}

// AssetResourceStatusStateEnumToProto converts a AssetResourceStatusStateEnum enum to its proto representation.
func DataplexAssetResourceStatusStateEnumToProto(e *dataplex.AssetResourceStatusStateEnum) dataplexpb.DataplexAssetResourceStatusStateEnum {
	if e == nil {
		return dataplexpb.DataplexAssetResourceStatusStateEnum(0)
	}
	if v, ok := dataplexpb.DataplexAssetResourceStatusStateEnum_value["AssetResourceStatusStateEnum"+string(*e)]; ok {
		return dataplexpb.DataplexAssetResourceStatusStateEnum(v)
	}
	return dataplexpb.DataplexAssetResourceStatusStateEnum(0)
}

// AssetSecurityStatusStateEnumToProto converts a AssetSecurityStatusStateEnum enum to its proto representation.
func DataplexAssetSecurityStatusStateEnumToProto(e *dataplex.AssetSecurityStatusStateEnum) dataplexpb.DataplexAssetSecurityStatusStateEnum {
	if e == nil {
		return dataplexpb.DataplexAssetSecurityStatusStateEnum(0)
	}
	if v, ok := dataplexpb.DataplexAssetSecurityStatusStateEnum_value["AssetSecurityStatusStateEnum"+string(*e)]; ok {
		return dataplexpb.DataplexAssetSecurityStatusStateEnum(v)
	}
	return dataplexpb.DataplexAssetSecurityStatusStateEnum(0)
}

// AssetDiscoveryStatusStateEnumToProto converts a AssetDiscoveryStatusStateEnum enum to its proto representation.
func DataplexAssetDiscoveryStatusStateEnumToProto(e *dataplex.AssetDiscoveryStatusStateEnum) dataplexpb.DataplexAssetDiscoveryStatusStateEnum {
	if e == nil {
		return dataplexpb.DataplexAssetDiscoveryStatusStateEnum(0)
	}
	if v, ok := dataplexpb.DataplexAssetDiscoveryStatusStateEnum_value["AssetDiscoveryStatusStateEnum"+string(*e)]; ok {
		return dataplexpb.DataplexAssetDiscoveryStatusStateEnum(v)
	}
	return dataplexpb.DataplexAssetDiscoveryStatusStateEnum(0)
}

// AssetResourceSpecToProto converts a AssetResourceSpec object to its proto representation.
func DataplexAssetResourceSpecToProto(o *dataplex.AssetResourceSpec) *dataplexpb.DataplexAssetResourceSpec {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexAssetResourceSpec{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(DataplexAssetResourceSpecTypeEnumToProto(o.Type))
	p.SetReadAccessMode(DataplexAssetResourceSpecReadAccessModeEnumToProto(o.ReadAccessMode))
	return p
}

// AssetResourceStatusToProto converts a AssetResourceStatus object to its proto representation.
func DataplexAssetResourceStatusToProto(o *dataplex.AssetResourceStatus) *dataplexpb.DataplexAssetResourceStatus {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexAssetResourceStatus{}
	p.SetState(DataplexAssetResourceStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// AssetSecurityStatusToProto converts a AssetSecurityStatus object to its proto representation.
func DataplexAssetSecurityStatusToProto(o *dataplex.AssetSecurityStatus) *dataplexpb.DataplexAssetSecurityStatus {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexAssetSecurityStatus{}
	p.SetState(DataplexAssetSecurityStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// AssetDiscoverySpecToProto converts a AssetDiscoverySpec object to its proto representation.
func DataplexAssetDiscoverySpecToProto(o *dataplex.AssetDiscoverySpec) *dataplexpb.DataplexAssetDiscoverySpec {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexAssetDiscoverySpec{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetCsvOptions(DataplexAssetDiscoverySpecCsvOptionsToProto(o.CsvOptions))
	p.SetJsonOptions(DataplexAssetDiscoverySpecJsonOptionsToProto(o.JsonOptions))
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

// AssetDiscoverySpecCsvOptionsToProto converts a AssetDiscoverySpecCsvOptions object to its proto representation.
func DataplexAssetDiscoverySpecCsvOptionsToProto(o *dataplex.AssetDiscoverySpecCsvOptions) *dataplexpb.DataplexAssetDiscoverySpecCsvOptions {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexAssetDiscoverySpecCsvOptions{}
	p.SetHeaderRows(dcl.ValueOrEmptyInt64(o.HeaderRows))
	p.SetDelimiter(dcl.ValueOrEmptyString(o.Delimiter))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// AssetDiscoverySpecJsonOptionsToProto converts a AssetDiscoverySpecJsonOptions object to its proto representation.
func DataplexAssetDiscoverySpecJsonOptionsToProto(o *dataplex.AssetDiscoverySpecJsonOptions) *dataplexpb.DataplexAssetDiscoverySpecJsonOptions {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexAssetDiscoverySpecJsonOptions{}
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// AssetDiscoveryStatusToProto converts a AssetDiscoveryStatus object to its proto representation.
func DataplexAssetDiscoveryStatusToProto(o *dataplex.AssetDiscoveryStatus) *dataplexpb.DataplexAssetDiscoveryStatus {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexAssetDiscoveryStatus{}
	p.SetState(DataplexAssetDiscoveryStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetLastRunTime(dcl.ValueOrEmptyString(o.LastRunTime))
	p.SetStats(DataplexAssetDiscoveryStatusStatsToProto(o.Stats))
	p.SetLastRunDuration(dcl.ValueOrEmptyString(o.LastRunDuration))
	return p
}

// AssetDiscoveryStatusStatsToProto converts a AssetDiscoveryStatusStats object to its proto representation.
func DataplexAssetDiscoveryStatusStatsToProto(o *dataplex.AssetDiscoveryStatusStats) *dataplexpb.DataplexAssetDiscoveryStatusStats {
	if o == nil {
		return nil
	}
	p := &dataplexpb.DataplexAssetDiscoveryStatusStats{}
	p.SetDataItems(dcl.ValueOrEmptyInt64(o.DataItems))
	p.SetDataSize(dcl.ValueOrEmptyInt64(o.DataSize))
	p.SetTables(dcl.ValueOrEmptyInt64(o.Tables))
	p.SetFilesets(dcl.ValueOrEmptyInt64(o.Filesets))
	return p
}

// AssetToProto converts a Asset resource to its proto representation.
func AssetToProto(resource *dataplex.Asset) *dataplexpb.DataplexAsset {
	p := &dataplexpb.DataplexAsset{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexAssetStateEnumToProto(resource.State))
	p.SetResourceSpec(DataplexAssetResourceSpecToProto(resource.ResourceSpec))
	p.SetResourceStatus(DataplexAssetResourceStatusToProto(resource.ResourceStatus))
	p.SetSecurityStatus(DataplexAssetSecurityStatusToProto(resource.SecurityStatus))
	p.SetDiscoverySpec(DataplexAssetDiscoverySpecToProto(resource.DiscoverySpec))
	p.SetDiscoveryStatus(DataplexAssetDiscoveryStatusToProto(resource.DiscoveryStatus))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetLake(dcl.ValueOrEmptyString(resource.Lake))
	p.SetDataplexZone(dcl.ValueOrEmptyString(resource.DataplexZone))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyAsset handles the gRPC request by passing it to the underlying Asset Apply() method.
func (s *AssetServer) applyAsset(ctx context.Context, c *dataplex.Client, request *dataplexpb.ApplyDataplexAssetRequest) (*dataplexpb.DataplexAsset, error) {
	p := ProtoToAsset(request.GetResource())
	res, err := c.ApplyAsset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AssetToProto(res)
	return r, nil
}

// applyDataplexAsset handles the gRPC request by passing it to the underlying Asset Apply() method.
func (s *AssetServer) ApplyDataplexAsset(ctx context.Context, request *dataplexpb.ApplyDataplexAssetRequest) (*dataplexpb.DataplexAsset, error) {
	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAsset(ctx, cl, request)
}

// DeleteAsset handles the gRPC request by passing it to the underlying Asset Delete() method.
func (s *AssetServer) DeleteDataplexAsset(ctx context.Context, request *dataplexpb.DeleteDataplexAssetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAsset(ctx, ProtoToAsset(request.GetResource()))

}

// ListDataplexAsset handles the gRPC request by passing it to the underlying AssetList() method.
func (s *AssetServer) ListDataplexAsset(ctx context.Context, request *dataplexpb.ListDataplexAssetRequest) (*dataplexpb.ListDataplexAssetResponse, error) {
	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAsset(ctx, request.GetProject(), request.GetLocation(), request.GetDataplexZone(), request.GetLake())
	if err != nil {
		return nil, err
	}
	var protos []*dataplexpb.DataplexAsset
	for _, r := range resources.Items {
		rp := AssetToProto(r)
		protos = append(protos, rp)
	}
	p := &dataplexpb.ListDataplexAssetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAsset(ctx context.Context, service_account_file string) (*dataplex.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataplex.NewClient(conf), nil
}
