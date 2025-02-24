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

// AssetServer implements the gRPC interface for Asset.
type AssetServer struct{}

// ProtoToAssetStateEnum converts a AssetStateEnum enum from its proto representation.
func ProtoToDataplexBetaAssetStateEnum(e betapb.DataplexBetaAssetStateEnum) *beta.AssetStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaAssetStateEnum_name[int32(e)]; ok {
		e := beta.AssetStateEnum(n[len("DataplexBetaAssetStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpecTypeEnum converts a AssetResourceSpecTypeEnum enum from its proto representation.
func ProtoToDataplexBetaAssetResourceSpecTypeEnum(e betapb.DataplexBetaAssetResourceSpecTypeEnum) *beta.AssetResourceSpecTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaAssetResourceSpecTypeEnum_name[int32(e)]; ok {
		e := beta.AssetResourceSpecTypeEnum(n[len("DataplexBetaAssetResourceSpecTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpecReadAccessModeEnum converts a AssetResourceSpecReadAccessModeEnum enum from its proto representation.
func ProtoToDataplexBetaAssetResourceSpecReadAccessModeEnum(e betapb.DataplexBetaAssetResourceSpecReadAccessModeEnum) *beta.AssetResourceSpecReadAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaAssetResourceSpecReadAccessModeEnum_name[int32(e)]; ok {
		e := beta.AssetResourceSpecReadAccessModeEnum(n[len("DataplexBetaAssetResourceSpecReadAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceStatusStateEnum converts a AssetResourceStatusStateEnum enum from its proto representation.
func ProtoToDataplexBetaAssetResourceStatusStateEnum(e betapb.DataplexBetaAssetResourceStatusStateEnum) *beta.AssetResourceStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaAssetResourceStatusStateEnum_name[int32(e)]; ok {
		e := beta.AssetResourceStatusStateEnum(n[len("DataplexBetaAssetResourceStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetSecurityStatusStateEnum converts a AssetSecurityStatusStateEnum enum from its proto representation.
func ProtoToDataplexBetaAssetSecurityStatusStateEnum(e betapb.DataplexBetaAssetSecurityStatusStateEnum) *beta.AssetSecurityStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaAssetSecurityStatusStateEnum_name[int32(e)]; ok {
		e := beta.AssetSecurityStatusStateEnum(n[len("DataplexBetaAssetSecurityStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetDiscoveryStatusStateEnum converts a AssetDiscoveryStatusStateEnum enum from its proto representation.
func ProtoToDataplexBetaAssetDiscoveryStatusStateEnum(e betapb.DataplexBetaAssetDiscoveryStatusStateEnum) *beta.AssetDiscoveryStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataplexBetaAssetDiscoveryStatusStateEnum_name[int32(e)]; ok {
		e := beta.AssetDiscoveryStatusStateEnum(n[len("DataplexBetaAssetDiscoveryStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpec converts a AssetResourceSpec object from its proto representation.
func ProtoToDataplexBetaAssetResourceSpec(p *betapb.DataplexBetaAssetResourceSpec) *beta.AssetResourceSpec {
	if p == nil {
		return nil
	}
	obj := &beta.AssetResourceSpec{
		Name:           dcl.StringOrNil(p.GetName()),
		Type:           ProtoToDataplexBetaAssetResourceSpecTypeEnum(p.GetType()),
		ReadAccessMode: ProtoToDataplexBetaAssetResourceSpecReadAccessModeEnum(p.GetReadAccessMode()),
	}
	return obj
}

// ProtoToAssetResourceStatus converts a AssetResourceStatus object from its proto representation.
func ProtoToDataplexBetaAssetResourceStatus(p *betapb.DataplexBetaAssetResourceStatus) *beta.AssetResourceStatus {
	if p == nil {
		return nil
	}
	obj := &beta.AssetResourceStatus{
		State:      ProtoToDataplexBetaAssetResourceStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToAssetSecurityStatus converts a AssetSecurityStatus object from its proto representation.
func ProtoToDataplexBetaAssetSecurityStatus(p *betapb.DataplexBetaAssetSecurityStatus) *beta.AssetSecurityStatus {
	if p == nil {
		return nil
	}
	obj := &beta.AssetSecurityStatus{
		State:      ProtoToDataplexBetaAssetSecurityStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToAssetDiscoverySpec converts a AssetDiscoverySpec object from its proto representation.
func ProtoToDataplexBetaAssetDiscoverySpec(p *betapb.DataplexBetaAssetDiscoverySpec) *beta.AssetDiscoverySpec {
	if p == nil {
		return nil
	}
	obj := &beta.AssetDiscoverySpec{
		Enabled:     dcl.Bool(p.GetEnabled()),
		CsvOptions:  ProtoToDataplexBetaAssetDiscoverySpecCsvOptions(p.GetCsvOptions()),
		JsonOptions: ProtoToDataplexBetaAssetDiscoverySpecJsonOptions(p.GetJsonOptions()),
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
func ProtoToDataplexBetaAssetDiscoverySpecCsvOptions(p *betapb.DataplexBetaAssetDiscoverySpecCsvOptions) *beta.AssetDiscoverySpecCsvOptions {
	if p == nil {
		return nil
	}
	obj := &beta.AssetDiscoverySpecCsvOptions{
		HeaderRows:           dcl.Int64OrNil(p.GetHeaderRows()),
		Delimiter:            dcl.StringOrNil(p.GetDelimiter()),
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToAssetDiscoverySpecJsonOptions converts a AssetDiscoverySpecJsonOptions object from its proto representation.
func ProtoToDataplexBetaAssetDiscoverySpecJsonOptions(p *betapb.DataplexBetaAssetDiscoverySpecJsonOptions) *beta.AssetDiscoverySpecJsonOptions {
	if p == nil {
		return nil
	}
	obj := &beta.AssetDiscoverySpecJsonOptions{
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToAssetDiscoveryStatus converts a AssetDiscoveryStatus object from its proto representation.
func ProtoToDataplexBetaAssetDiscoveryStatus(p *betapb.DataplexBetaAssetDiscoveryStatus) *beta.AssetDiscoveryStatus {
	if p == nil {
		return nil
	}
	obj := &beta.AssetDiscoveryStatus{
		State:           ProtoToDataplexBetaAssetDiscoveryStatusStateEnum(p.GetState()),
		Message:         dcl.StringOrNil(p.GetMessage()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		LastRunTime:     dcl.StringOrNil(p.GetLastRunTime()),
		Stats:           ProtoToDataplexBetaAssetDiscoveryStatusStats(p.GetStats()),
		LastRunDuration: dcl.StringOrNil(p.GetLastRunDuration()),
	}
	return obj
}

// ProtoToAssetDiscoveryStatusStats converts a AssetDiscoveryStatusStats object from its proto representation.
func ProtoToDataplexBetaAssetDiscoveryStatusStats(p *betapb.DataplexBetaAssetDiscoveryStatusStats) *beta.AssetDiscoveryStatusStats {
	if p == nil {
		return nil
	}
	obj := &beta.AssetDiscoveryStatusStats{
		DataItems: dcl.Int64OrNil(p.GetDataItems()),
		DataSize:  dcl.Int64OrNil(p.GetDataSize()),
		Tables:    dcl.Int64OrNil(p.GetTables()),
		Filesets:  dcl.Int64OrNil(p.GetFilesets()),
	}
	return obj
}

// ProtoToAsset converts a Asset resource from its proto representation.
func ProtoToAsset(p *betapb.DataplexBetaAsset) *beta.Asset {
	obj := &beta.Asset{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToDataplexBetaAssetStateEnum(p.GetState()),
		ResourceSpec:    ProtoToDataplexBetaAssetResourceSpec(p.GetResourceSpec()),
		ResourceStatus:  ProtoToDataplexBetaAssetResourceStatus(p.GetResourceStatus()),
		SecurityStatus:  ProtoToDataplexBetaAssetSecurityStatus(p.GetSecurityStatus()),
		DiscoverySpec:   ProtoToDataplexBetaAssetDiscoverySpec(p.GetDiscoverySpec()),
		DiscoveryStatus: ProtoToDataplexBetaAssetDiscoveryStatus(p.GetDiscoveryStatus()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
		Lake:            dcl.StringOrNil(p.GetLake()),
		DataplexZone:    dcl.StringOrNil(p.GetDataplexZone()),
	}
	return obj
}

// AssetStateEnumToProto converts a AssetStateEnum enum to its proto representation.
func DataplexBetaAssetStateEnumToProto(e *beta.AssetStateEnum) betapb.DataplexBetaAssetStateEnum {
	if e == nil {
		return betapb.DataplexBetaAssetStateEnum(0)
	}
	if v, ok := betapb.DataplexBetaAssetStateEnum_value["AssetStateEnum"+string(*e)]; ok {
		return betapb.DataplexBetaAssetStateEnum(v)
	}
	return betapb.DataplexBetaAssetStateEnum(0)
}

// AssetResourceSpecTypeEnumToProto converts a AssetResourceSpecTypeEnum enum to its proto representation.
func DataplexBetaAssetResourceSpecTypeEnumToProto(e *beta.AssetResourceSpecTypeEnum) betapb.DataplexBetaAssetResourceSpecTypeEnum {
	if e == nil {
		return betapb.DataplexBetaAssetResourceSpecTypeEnum(0)
	}
	if v, ok := betapb.DataplexBetaAssetResourceSpecTypeEnum_value["AssetResourceSpecTypeEnum"+string(*e)]; ok {
		return betapb.DataplexBetaAssetResourceSpecTypeEnum(v)
	}
	return betapb.DataplexBetaAssetResourceSpecTypeEnum(0)
}

// AssetResourceSpecReadAccessModeEnumToProto converts a AssetResourceSpecReadAccessModeEnum enum to its proto representation.
func DataplexBetaAssetResourceSpecReadAccessModeEnumToProto(e *beta.AssetResourceSpecReadAccessModeEnum) betapb.DataplexBetaAssetResourceSpecReadAccessModeEnum {
	if e == nil {
		return betapb.DataplexBetaAssetResourceSpecReadAccessModeEnum(0)
	}
	if v, ok := betapb.DataplexBetaAssetResourceSpecReadAccessModeEnum_value["AssetResourceSpecReadAccessModeEnum"+string(*e)]; ok {
		return betapb.DataplexBetaAssetResourceSpecReadAccessModeEnum(v)
	}
	return betapb.DataplexBetaAssetResourceSpecReadAccessModeEnum(0)
}

// AssetResourceStatusStateEnumToProto converts a AssetResourceStatusStateEnum enum to its proto representation.
func DataplexBetaAssetResourceStatusStateEnumToProto(e *beta.AssetResourceStatusStateEnum) betapb.DataplexBetaAssetResourceStatusStateEnum {
	if e == nil {
		return betapb.DataplexBetaAssetResourceStatusStateEnum(0)
	}
	if v, ok := betapb.DataplexBetaAssetResourceStatusStateEnum_value["AssetResourceStatusStateEnum"+string(*e)]; ok {
		return betapb.DataplexBetaAssetResourceStatusStateEnum(v)
	}
	return betapb.DataplexBetaAssetResourceStatusStateEnum(0)
}

// AssetSecurityStatusStateEnumToProto converts a AssetSecurityStatusStateEnum enum to its proto representation.
func DataplexBetaAssetSecurityStatusStateEnumToProto(e *beta.AssetSecurityStatusStateEnum) betapb.DataplexBetaAssetSecurityStatusStateEnum {
	if e == nil {
		return betapb.DataplexBetaAssetSecurityStatusStateEnum(0)
	}
	if v, ok := betapb.DataplexBetaAssetSecurityStatusStateEnum_value["AssetSecurityStatusStateEnum"+string(*e)]; ok {
		return betapb.DataplexBetaAssetSecurityStatusStateEnum(v)
	}
	return betapb.DataplexBetaAssetSecurityStatusStateEnum(0)
}

// AssetDiscoveryStatusStateEnumToProto converts a AssetDiscoveryStatusStateEnum enum to its proto representation.
func DataplexBetaAssetDiscoveryStatusStateEnumToProto(e *beta.AssetDiscoveryStatusStateEnum) betapb.DataplexBetaAssetDiscoveryStatusStateEnum {
	if e == nil {
		return betapb.DataplexBetaAssetDiscoveryStatusStateEnum(0)
	}
	if v, ok := betapb.DataplexBetaAssetDiscoveryStatusStateEnum_value["AssetDiscoveryStatusStateEnum"+string(*e)]; ok {
		return betapb.DataplexBetaAssetDiscoveryStatusStateEnum(v)
	}
	return betapb.DataplexBetaAssetDiscoveryStatusStateEnum(0)
}

// AssetResourceSpecToProto converts a AssetResourceSpec object to its proto representation.
func DataplexBetaAssetResourceSpecToProto(o *beta.AssetResourceSpec) *betapb.DataplexBetaAssetResourceSpec {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaAssetResourceSpec{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(DataplexBetaAssetResourceSpecTypeEnumToProto(o.Type))
	p.SetReadAccessMode(DataplexBetaAssetResourceSpecReadAccessModeEnumToProto(o.ReadAccessMode))
	return p
}

// AssetResourceStatusToProto converts a AssetResourceStatus object to its proto representation.
func DataplexBetaAssetResourceStatusToProto(o *beta.AssetResourceStatus) *betapb.DataplexBetaAssetResourceStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaAssetResourceStatus{}
	p.SetState(DataplexBetaAssetResourceStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// AssetSecurityStatusToProto converts a AssetSecurityStatus object to its proto representation.
func DataplexBetaAssetSecurityStatusToProto(o *beta.AssetSecurityStatus) *betapb.DataplexBetaAssetSecurityStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaAssetSecurityStatus{}
	p.SetState(DataplexBetaAssetSecurityStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// AssetDiscoverySpecToProto converts a AssetDiscoverySpec object to its proto representation.
func DataplexBetaAssetDiscoverySpecToProto(o *beta.AssetDiscoverySpec) *betapb.DataplexBetaAssetDiscoverySpec {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaAssetDiscoverySpec{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetCsvOptions(DataplexBetaAssetDiscoverySpecCsvOptionsToProto(o.CsvOptions))
	p.SetJsonOptions(DataplexBetaAssetDiscoverySpecJsonOptionsToProto(o.JsonOptions))
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
func DataplexBetaAssetDiscoverySpecCsvOptionsToProto(o *beta.AssetDiscoverySpecCsvOptions) *betapb.DataplexBetaAssetDiscoverySpecCsvOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaAssetDiscoverySpecCsvOptions{}
	p.SetHeaderRows(dcl.ValueOrEmptyInt64(o.HeaderRows))
	p.SetDelimiter(dcl.ValueOrEmptyString(o.Delimiter))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// AssetDiscoverySpecJsonOptionsToProto converts a AssetDiscoverySpecJsonOptions object to its proto representation.
func DataplexBetaAssetDiscoverySpecJsonOptionsToProto(o *beta.AssetDiscoverySpecJsonOptions) *betapb.DataplexBetaAssetDiscoverySpecJsonOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaAssetDiscoverySpecJsonOptions{}
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// AssetDiscoveryStatusToProto converts a AssetDiscoveryStatus object to its proto representation.
func DataplexBetaAssetDiscoveryStatusToProto(o *beta.AssetDiscoveryStatus) *betapb.DataplexBetaAssetDiscoveryStatus {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaAssetDiscoveryStatus{}
	p.SetState(DataplexBetaAssetDiscoveryStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetLastRunTime(dcl.ValueOrEmptyString(o.LastRunTime))
	p.SetStats(DataplexBetaAssetDiscoveryStatusStatsToProto(o.Stats))
	p.SetLastRunDuration(dcl.ValueOrEmptyString(o.LastRunDuration))
	return p
}

// AssetDiscoveryStatusStatsToProto converts a AssetDiscoveryStatusStats object to its proto representation.
func DataplexBetaAssetDiscoveryStatusStatsToProto(o *beta.AssetDiscoveryStatusStats) *betapb.DataplexBetaAssetDiscoveryStatusStats {
	if o == nil {
		return nil
	}
	p := &betapb.DataplexBetaAssetDiscoveryStatusStats{}
	p.SetDataItems(dcl.ValueOrEmptyInt64(o.DataItems))
	p.SetDataSize(dcl.ValueOrEmptyInt64(o.DataSize))
	p.SetTables(dcl.ValueOrEmptyInt64(o.Tables))
	p.SetFilesets(dcl.ValueOrEmptyInt64(o.Filesets))
	return p
}

// AssetToProto converts a Asset resource to its proto representation.
func AssetToProto(resource *beta.Asset) *betapb.DataplexBetaAsset {
	p := &betapb.DataplexBetaAsset{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexBetaAssetStateEnumToProto(resource.State))
	p.SetResourceSpec(DataplexBetaAssetResourceSpecToProto(resource.ResourceSpec))
	p.SetResourceStatus(DataplexBetaAssetResourceStatusToProto(resource.ResourceStatus))
	p.SetSecurityStatus(DataplexBetaAssetSecurityStatusToProto(resource.SecurityStatus))
	p.SetDiscoverySpec(DataplexBetaAssetDiscoverySpecToProto(resource.DiscoverySpec))
	p.SetDiscoveryStatus(DataplexBetaAssetDiscoveryStatusToProto(resource.DiscoveryStatus))
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
func (s *AssetServer) applyAsset(ctx context.Context, c *beta.Client, request *betapb.ApplyDataplexBetaAssetRequest) (*betapb.DataplexBetaAsset, error) {
	p := ProtoToAsset(request.GetResource())
	res, err := c.ApplyAsset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AssetToProto(res)
	return r, nil
}

// applyDataplexBetaAsset handles the gRPC request by passing it to the underlying Asset Apply() method.
func (s *AssetServer) ApplyDataplexBetaAsset(ctx context.Context, request *betapb.ApplyDataplexBetaAssetRequest) (*betapb.DataplexBetaAsset, error) {
	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAsset(ctx, cl, request)
}

// DeleteAsset handles the gRPC request by passing it to the underlying Asset Delete() method.
func (s *AssetServer) DeleteDataplexBetaAsset(ctx context.Context, request *betapb.DeleteDataplexBetaAssetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAsset(ctx, ProtoToAsset(request.GetResource()))

}

// ListDataplexBetaAsset handles the gRPC request by passing it to the underlying AssetList() method.
func (s *AssetServer) ListDataplexBetaAsset(ctx context.Context, request *betapb.ListDataplexBetaAssetRequest) (*betapb.ListDataplexBetaAssetResponse, error) {
	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAsset(ctx, request.GetProject(), request.GetLocation(), request.GetDataplexZone(), request.GetLake())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataplexBetaAsset
	for _, r := range resources.Items {
		rp := AssetToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDataplexBetaAssetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAsset(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
