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

// AssetServer implements the gRPC interface for Asset.
type AssetServer struct{}

// ProtoToAssetStateEnum converts a AssetStateEnum enum from its proto representation.
func ProtoToDataplexAlphaAssetStateEnum(e alphapb.DataplexAlphaAssetStateEnum) *alpha.AssetStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaAssetStateEnum_name[int32(e)]; ok {
		e := alpha.AssetStateEnum(n[len("DataplexAlphaAssetStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpecTypeEnum converts a AssetResourceSpecTypeEnum enum from its proto representation.
func ProtoToDataplexAlphaAssetResourceSpecTypeEnum(e alphapb.DataplexAlphaAssetResourceSpecTypeEnum) *alpha.AssetResourceSpecTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaAssetResourceSpecTypeEnum_name[int32(e)]; ok {
		e := alpha.AssetResourceSpecTypeEnum(n[len("DataplexAlphaAssetResourceSpecTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpecReadAccessModeEnum converts a AssetResourceSpecReadAccessModeEnum enum from its proto representation.
func ProtoToDataplexAlphaAssetResourceSpecReadAccessModeEnum(e alphapb.DataplexAlphaAssetResourceSpecReadAccessModeEnum) *alpha.AssetResourceSpecReadAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaAssetResourceSpecReadAccessModeEnum_name[int32(e)]; ok {
		e := alpha.AssetResourceSpecReadAccessModeEnum(n[len("DataplexAlphaAssetResourceSpecReadAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceStatusStateEnum converts a AssetResourceStatusStateEnum enum from its proto representation.
func ProtoToDataplexAlphaAssetResourceStatusStateEnum(e alphapb.DataplexAlphaAssetResourceStatusStateEnum) *alpha.AssetResourceStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaAssetResourceStatusStateEnum_name[int32(e)]; ok {
		e := alpha.AssetResourceStatusStateEnum(n[len("DataplexAlphaAssetResourceStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetSecurityStatusStateEnum converts a AssetSecurityStatusStateEnum enum from its proto representation.
func ProtoToDataplexAlphaAssetSecurityStatusStateEnum(e alphapb.DataplexAlphaAssetSecurityStatusStateEnum) *alpha.AssetSecurityStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaAssetSecurityStatusStateEnum_name[int32(e)]; ok {
		e := alpha.AssetSecurityStatusStateEnum(n[len("DataplexAlphaAssetSecurityStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetDiscoveryStatusStateEnum converts a AssetDiscoveryStatusStateEnum enum from its proto representation.
func ProtoToDataplexAlphaAssetDiscoveryStatusStateEnum(e alphapb.DataplexAlphaAssetDiscoveryStatusStateEnum) *alpha.AssetDiscoveryStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataplexAlphaAssetDiscoveryStatusStateEnum_name[int32(e)]; ok {
		e := alpha.AssetDiscoveryStatusStateEnum(n[len("DataplexAlphaAssetDiscoveryStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssetResourceSpec converts a AssetResourceSpec object from its proto representation.
func ProtoToDataplexAlphaAssetResourceSpec(p *alphapb.DataplexAlphaAssetResourceSpec) *alpha.AssetResourceSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.AssetResourceSpec{
		Name:           dcl.StringOrNil(p.GetName()),
		Type:           ProtoToDataplexAlphaAssetResourceSpecTypeEnum(p.GetType()),
		ReadAccessMode: ProtoToDataplexAlphaAssetResourceSpecReadAccessModeEnum(p.GetReadAccessMode()),
	}
	return obj
}

// ProtoToAssetResourceStatus converts a AssetResourceStatus object from its proto representation.
func ProtoToDataplexAlphaAssetResourceStatus(p *alphapb.DataplexAlphaAssetResourceStatus) *alpha.AssetResourceStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.AssetResourceStatus{
		State:      ProtoToDataplexAlphaAssetResourceStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToAssetSecurityStatus converts a AssetSecurityStatus object from its proto representation.
func ProtoToDataplexAlphaAssetSecurityStatus(p *alphapb.DataplexAlphaAssetSecurityStatus) *alpha.AssetSecurityStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.AssetSecurityStatus{
		State:      ProtoToDataplexAlphaAssetSecurityStatusStateEnum(p.GetState()),
		Message:    dcl.StringOrNil(p.GetMessage()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToAssetDiscoverySpec converts a AssetDiscoverySpec object from its proto representation.
func ProtoToDataplexAlphaAssetDiscoverySpec(p *alphapb.DataplexAlphaAssetDiscoverySpec) *alpha.AssetDiscoverySpec {
	if p == nil {
		return nil
	}
	obj := &alpha.AssetDiscoverySpec{
		Enabled:     dcl.Bool(p.GetEnabled()),
		CsvOptions:  ProtoToDataplexAlphaAssetDiscoverySpecCsvOptions(p.GetCsvOptions()),
		JsonOptions: ProtoToDataplexAlphaAssetDiscoverySpecJsonOptions(p.GetJsonOptions()),
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
func ProtoToDataplexAlphaAssetDiscoverySpecCsvOptions(p *alphapb.DataplexAlphaAssetDiscoverySpecCsvOptions) *alpha.AssetDiscoverySpecCsvOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.AssetDiscoverySpecCsvOptions{
		HeaderRows:           dcl.Int64OrNil(p.GetHeaderRows()),
		Delimiter:            dcl.StringOrNil(p.GetDelimiter()),
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToAssetDiscoverySpecJsonOptions converts a AssetDiscoverySpecJsonOptions object from its proto representation.
func ProtoToDataplexAlphaAssetDiscoverySpecJsonOptions(p *alphapb.DataplexAlphaAssetDiscoverySpecJsonOptions) *alpha.AssetDiscoverySpecJsonOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.AssetDiscoverySpecJsonOptions{
		Encoding:             dcl.StringOrNil(p.GetEncoding()),
		DisableTypeInference: dcl.Bool(p.GetDisableTypeInference()),
	}
	return obj
}

// ProtoToAssetDiscoveryStatus converts a AssetDiscoveryStatus object from its proto representation.
func ProtoToDataplexAlphaAssetDiscoveryStatus(p *alphapb.DataplexAlphaAssetDiscoveryStatus) *alpha.AssetDiscoveryStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.AssetDiscoveryStatus{
		State:           ProtoToDataplexAlphaAssetDiscoveryStatusStateEnum(p.GetState()),
		Message:         dcl.StringOrNil(p.GetMessage()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		LastRunTime:     dcl.StringOrNil(p.GetLastRunTime()),
		Stats:           ProtoToDataplexAlphaAssetDiscoveryStatusStats(p.GetStats()),
		LastRunDuration: dcl.StringOrNil(p.GetLastRunDuration()),
	}
	return obj
}

// ProtoToAssetDiscoveryStatusStats converts a AssetDiscoveryStatusStats object from its proto representation.
func ProtoToDataplexAlphaAssetDiscoveryStatusStats(p *alphapb.DataplexAlphaAssetDiscoveryStatusStats) *alpha.AssetDiscoveryStatusStats {
	if p == nil {
		return nil
	}
	obj := &alpha.AssetDiscoveryStatusStats{
		DataItems: dcl.Int64OrNil(p.GetDataItems()),
		DataSize:  dcl.Int64OrNil(p.GetDataSize()),
		Tables:    dcl.Int64OrNil(p.GetTables()),
		Filesets:  dcl.Int64OrNil(p.GetFilesets()),
	}
	return obj
}

// ProtoToAsset converts a Asset resource from its proto representation.
func ProtoToAsset(p *alphapb.DataplexAlphaAsset) *alpha.Asset {
	obj := &alpha.Asset{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Uid:             dcl.StringOrNil(p.GetUid()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToDataplexAlphaAssetStateEnum(p.GetState()),
		ResourceSpec:    ProtoToDataplexAlphaAssetResourceSpec(p.GetResourceSpec()),
		ResourceStatus:  ProtoToDataplexAlphaAssetResourceStatus(p.GetResourceStatus()),
		SecurityStatus:  ProtoToDataplexAlphaAssetSecurityStatus(p.GetSecurityStatus()),
		DiscoverySpec:   ProtoToDataplexAlphaAssetDiscoverySpec(p.GetDiscoverySpec()),
		DiscoveryStatus: ProtoToDataplexAlphaAssetDiscoveryStatus(p.GetDiscoveryStatus()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
		Lake:            dcl.StringOrNil(p.GetLake()),
		DataplexZone:    dcl.StringOrNil(p.GetDataplexZone()),
	}
	return obj
}

// AssetStateEnumToProto converts a AssetStateEnum enum to its proto representation.
func DataplexAlphaAssetStateEnumToProto(e *alpha.AssetStateEnum) alphapb.DataplexAlphaAssetStateEnum {
	if e == nil {
		return alphapb.DataplexAlphaAssetStateEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaAssetStateEnum_value["AssetStateEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaAssetStateEnum(v)
	}
	return alphapb.DataplexAlphaAssetStateEnum(0)
}

// AssetResourceSpecTypeEnumToProto converts a AssetResourceSpecTypeEnum enum to its proto representation.
func DataplexAlphaAssetResourceSpecTypeEnumToProto(e *alpha.AssetResourceSpecTypeEnum) alphapb.DataplexAlphaAssetResourceSpecTypeEnum {
	if e == nil {
		return alphapb.DataplexAlphaAssetResourceSpecTypeEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaAssetResourceSpecTypeEnum_value["AssetResourceSpecTypeEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaAssetResourceSpecTypeEnum(v)
	}
	return alphapb.DataplexAlphaAssetResourceSpecTypeEnum(0)
}

// AssetResourceSpecReadAccessModeEnumToProto converts a AssetResourceSpecReadAccessModeEnum enum to its proto representation.
func DataplexAlphaAssetResourceSpecReadAccessModeEnumToProto(e *alpha.AssetResourceSpecReadAccessModeEnum) alphapb.DataplexAlphaAssetResourceSpecReadAccessModeEnum {
	if e == nil {
		return alphapb.DataplexAlphaAssetResourceSpecReadAccessModeEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaAssetResourceSpecReadAccessModeEnum_value["AssetResourceSpecReadAccessModeEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaAssetResourceSpecReadAccessModeEnum(v)
	}
	return alphapb.DataplexAlphaAssetResourceSpecReadAccessModeEnum(0)
}

// AssetResourceStatusStateEnumToProto converts a AssetResourceStatusStateEnum enum to its proto representation.
func DataplexAlphaAssetResourceStatusStateEnumToProto(e *alpha.AssetResourceStatusStateEnum) alphapb.DataplexAlphaAssetResourceStatusStateEnum {
	if e == nil {
		return alphapb.DataplexAlphaAssetResourceStatusStateEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaAssetResourceStatusStateEnum_value["AssetResourceStatusStateEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaAssetResourceStatusStateEnum(v)
	}
	return alphapb.DataplexAlphaAssetResourceStatusStateEnum(0)
}

// AssetSecurityStatusStateEnumToProto converts a AssetSecurityStatusStateEnum enum to its proto representation.
func DataplexAlphaAssetSecurityStatusStateEnumToProto(e *alpha.AssetSecurityStatusStateEnum) alphapb.DataplexAlphaAssetSecurityStatusStateEnum {
	if e == nil {
		return alphapb.DataplexAlphaAssetSecurityStatusStateEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaAssetSecurityStatusStateEnum_value["AssetSecurityStatusStateEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaAssetSecurityStatusStateEnum(v)
	}
	return alphapb.DataplexAlphaAssetSecurityStatusStateEnum(0)
}

// AssetDiscoveryStatusStateEnumToProto converts a AssetDiscoveryStatusStateEnum enum to its proto representation.
func DataplexAlphaAssetDiscoveryStatusStateEnumToProto(e *alpha.AssetDiscoveryStatusStateEnum) alphapb.DataplexAlphaAssetDiscoveryStatusStateEnum {
	if e == nil {
		return alphapb.DataplexAlphaAssetDiscoveryStatusStateEnum(0)
	}
	if v, ok := alphapb.DataplexAlphaAssetDiscoveryStatusStateEnum_value["AssetDiscoveryStatusStateEnum"+string(*e)]; ok {
		return alphapb.DataplexAlphaAssetDiscoveryStatusStateEnum(v)
	}
	return alphapb.DataplexAlphaAssetDiscoveryStatusStateEnum(0)
}

// AssetResourceSpecToProto converts a AssetResourceSpec object to its proto representation.
func DataplexAlphaAssetResourceSpecToProto(o *alpha.AssetResourceSpec) *alphapb.DataplexAlphaAssetResourceSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaAssetResourceSpec{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(DataplexAlphaAssetResourceSpecTypeEnumToProto(o.Type))
	p.SetReadAccessMode(DataplexAlphaAssetResourceSpecReadAccessModeEnumToProto(o.ReadAccessMode))
	return p
}

// AssetResourceStatusToProto converts a AssetResourceStatus object to its proto representation.
func DataplexAlphaAssetResourceStatusToProto(o *alpha.AssetResourceStatus) *alphapb.DataplexAlphaAssetResourceStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaAssetResourceStatus{}
	p.SetState(DataplexAlphaAssetResourceStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// AssetSecurityStatusToProto converts a AssetSecurityStatus object to its proto representation.
func DataplexAlphaAssetSecurityStatusToProto(o *alpha.AssetSecurityStatus) *alphapb.DataplexAlphaAssetSecurityStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaAssetSecurityStatus{}
	p.SetState(DataplexAlphaAssetSecurityStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// AssetDiscoverySpecToProto converts a AssetDiscoverySpec object to its proto representation.
func DataplexAlphaAssetDiscoverySpecToProto(o *alpha.AssetDiscoverySpec) *alphapb.DataplexAlphaAssetDiscoverySpec {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaAssetDiscoverySpec{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetCsvOptions(DataplexAlphaAssetDiscoverySpecCsvOptionsToProto(o.CsvOptions))
	p.SetJsonOptions(DataplexAlphaAssetDiscoverySpecJsonOptionsToProto(o.JsonOptions))
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
func DataplexAlphaAssetDiscoverySpecCsvOptionsToProto(o *alpha.AssetDiscoverySpecCsvOptions) *alphapb.DataplexAlphaAssetDiscoverySpecCsvOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaAssetDiscoverySpecCsvOptions{}
	p.SetHeaderRows(dcl.ValueOrEmptyInt64(o.HeaderRows))
	p.SetDelimiter(dcl.ValueOrEmptyString(o.Delimiter))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// AssetDiscoverySpecJsonOptionsToProto converts a AssetDiscoverySpecJsonOptions object to its proto representation.
func DataplexAlphaAssetDiscoverySpecJsonOptionsToProto(o *alpha.AssetDiscoverySpecJsonOptions) *alphapb.DataplexAlphaAssetDiscoverySpecJsonOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaAssetDiscoverySpecJsonOptions{}
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetDisableTypeInference(dcl.ValueOrEmptyBool(o.DisableTypeInference))
	return p
}

// AssetDiscoveryStatusToProto converts a AssetDiscoveryStatus object to its proto representation.
func DataplexAlphaAssetDiscoveryStatusToProto(o *alpha.AssetDiscoveryStatus) *alphapb.DataplexAlphaAssetDiscoveryStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaAssetDiscoveryStatus{}
	p.SetState(DataplexAlphaAssetDiscoveryStatusStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetLastRunTime(dcl.ValueOrEmptyString(o.LastRunTime))
	p.SetStats(DataplexAlphaAssetDiscoveryStatusStatsToProto(o.Stats))
	p.SetLastRunDuration(dcl.ValueOrEmptyString(o.LastRunDuration))
	return p
}

// AssetDiscoveryStatusStatsToProto converts a AssetDiscoveryStatusStats object to its proto representation.
func DataplexAlphaAssetDiscoveryStatusStatsToProto(o *alpha.AssetDiscoveryStatusStats) *alphapb.DataplexAlphaAssetDiscoveryStatusStats {
	if o == nil {
		return nil
	}
	p := &alphapb.DataplexAlphaAssetDiscoveryStatusStats{}
	p.SetDataItems(dcl.ValueOrEmptyInt64(o.DataItems))
	p.SetDataSize(dcl.ValueOrEmptyInt64(o.DataSize))
	p.SetTables(dcl.ValueOrEmptyInt64(o.Tables))
	p.SetFilesets(dcl.ValueOrEmptyInt64(o.Filesets))
	return p
}

// AssetToProto converts a Asset resource to its proto representation.
func AssetToProto(resource *alpha.Asset) *alphapb.DataplexAlphaAsset {
	p := &alphapb.DataplexAlphaAsset{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(DataplexAlphaAssetStateEnumToProto(resource.State))
	p.SetResourceSpec(DataplexAlphaAssetResourceSpecToProto(resource.ResourceSpec))
	p.SetResourceStatus(DataplexAlphaAssetResourceStatusToProto(resource.ResourceStatus))
	p.SetSecurityStatus(DataplexAlphaAssetSecurityStatusToProto(resource.SecurityStatus))
	p.SetDiscoverySpec(DataplexAlphaAssetDiscoverySpecToProto(resource.DiscoverySpec))
	p.SetDiscoveryStatus(DataplexAlphaAssetDiscoveryStatusToProto(resource.DiscoveryStatus))
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
func (s *AssetServer) applyAsset(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataplexAlphaAssetRequest) (*alphapb.DataplexAlphaAsset, error) {
	p := ProtoToAsset(request.GetResource())
	res, err := c.ApplyAsset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AssetToProto(res)
	return r, nil
}

// applyDataplexAlphaAsset handles the gRPC request by passing it to the underlying Asset Apply() method.
func (s *AssetServer) ApplyDataplexAlphaAsset(ctx context.Context, request *alphapb.ApplyDataplexAlphaAssetRequest) (*alphapb.DataplexAlphaAsset, error) {
	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAsset(ctx, cl, request)
}

// DeleteAsset handles the gRPC request by passing it to the underlying Asset Delete() method.
func (s *AssetServer) DeleteDataplexAlphaAsset(ctx context.Context, request *alphapb.DeleteDataplexAlphaAssetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAsset(ctx, ProtoToAsset(request.GetResource()))

}

// ListDataplexAlphaAsset handles the gRPC request by passing it to the underlying AssetList() method.
func (s *AssetServer) ListDataplexAlphaAsset(ctx context.Context, request *alphapb.ListDataplexAlphaAssetRequest) (*alphapb.ListDataplexAlphaAssetResponse, error) {
	cl, err := createConfigAsset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAsset(ctx, request.GetProject(), request.GetLocation(), request.GetDataplexZone(), request.GetLake())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataplexAlphaAsset
	for _, r := range resources.Items {
		rp := AssetToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDataplexAlphaAssetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAsset(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
