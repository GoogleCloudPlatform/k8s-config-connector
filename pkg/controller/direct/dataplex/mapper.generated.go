// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dataplex

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DataDiscoveryResult_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoveryResult) *krm.DataDiscoveryResult {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoveryResult{}
	// MISSING: BigqueryPublishing
	return out
}
func DataDiscoveryResult_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoveryResult) *pb.DataDiscoveryResult {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoveryResult{}
	// MISSING: BigqueryPublishing
	return out
}
func DataDiscoveryResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoveryResult) *krm.DataDiscoveryResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoveryResultObservedState{}
	out.BigqueryPublishing = DataDiscoveryResult_BigQueryPublishing_FromProto(mapCtx, in.GetBigqueryPublishing())
	return out
}
func DataDiscoveryResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoveryResultObservedState) *pb.DataDiscoveryResult {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoveryResult{}
	out.BigqueryPublishing = DataDiscoveryResult_BigQueryPublishing_ToProto(mapCtx, in.BigqueryPublishing)
	return out
}
func DataDiscoveryResult_BigQueryPublishing_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoveryResult_BigQueryPublishing) *krm.DataDiscoveryResult_BigQueryPublishing {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoveryResult_BigQueryPublishing{}
	// MISSING: Dataset
	return out
}
func DataDiscoveryResult_BigQueryPublishing_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoveryResult_BigQueryPublishing) *pb.DataDiscoveryResult_BigQueryPublishing {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoveryResult_BigQueryPublishing{}
	// MISSING: Dataset
	return out
}
func DataDiscoveryResult_BigQueryPublishingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoveryResult_BigQueryPublishing) *krm.DataDiscoveryResult_BigQueryPublishingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoveryResult_BigQueryPublishingObservedState{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	return out
}
func DataDiscoveryResult_BigQueryPublishingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoveryResult_BigQueryPublishingObservedState) *pb.DataDiscoveryResult_BigQueryPublishing {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoveryResult_BigQueryPublishing{}
	out.Dataset = direct.ValueOf(in.Dataset)
	return out
}
func DataDiscoverySpec_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoverySpec) *krm.DataDiscoverySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoverySpec{}
	out.BigqueryPublishingConfig = DataDiscoverySpec_BigQueryPublishingConfig_FromProto(mapCtx, in.GetBigqueryPublishingConfig())
	out.StorageConfig = DataDiscoverySpec_StorageConfig_FromProto(mapCtx, in.GetStorageConfig())
	return out
}
func DataDiscoverySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoverySpec) *pb.DataDiscoverySpec {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoverySpec{}
	out.BigqueryPublishingConfig = DataDiscoverySpec_BigQueryPublishingConfig_ToProto(mapCtx, in.BigqueryPublishingConfig)
	if oneof := DataDiscoverySpec_StorageConfig_ToProto(mapCtx, in.StorageConfig); oneof != nil {
		out.ResourceConfig = &pb.DataDiscoverySpec_StorageConfig_{StorageConfig: oneof}
	}
	return out
}
func DataDiscoverySpec_BigQueryPublishingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoverySpec_BigQueryPublishingConfig) *krm.DataDiscoverySpec_BigQueryPublishingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoverySpec_BigQueryPublishingConfig{}
	out.TableType = direct.Enum_FromProto(mapCtx, in.GetTableType())
	out.Connection = direct.LazyPtr(in.GetConnection())
	return out
}
func DataDiscoverySpec_BigQueryPublishingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoverySpec_BigQueryPublishingConfig) *pb.DataDiscoverySpec_BigQueryPublishingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoverySpec_BigQueryPublishingConfig{}
	out.TableType = direct.Enum_ToProto[pb.DataDiscoverySpec_BigQueryPublishingConfig_TableType](mapCtx, in.TableType)
	out.Connection = direct.ValueOf(in.Connection)
	return out
}
func DataDiscoverySpec_StorageConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoverySpec_StorageConfig) *krm.DataDiscoverySpec_StorageConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoverySpec_StorageConfig{}
	out.IncludePatterns = in.IncludePatterns
	out.ExcludePatterns = in.ExcludePatterns
	out.CsvOptions = DataDiscoverySpec_StorageConfig_CsvOptions_FromProto(mapCtx, in.GetCsvOptions())
	out.JsonOptions = DataDiscoverySpec_StorageConfig_JsonOptions_FromProto(mapCtx, in.GetJsonOptions())
	return out
}
func DataDiscoverySpec_StorageConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoverySpec_StorageConfig) *pb.DataDiscoverySpec_StorageConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoverySpec_StorageConfig{}
	out.IncludePatterns = in.IncludePatterns
	out.ExcludePatterns = in.ExcludePatterns
	out.CsvOptions = DataDiscoverySpec_StorageConfig_CsvOptions_ToProto(mapCtx, in.CsvOptions)
	out.JsonOptions = DataDiscoverySpec_StorageConfig_JsonOptions_ToProto(mapCtx, in.JsonOptions)
	return out
}
func DataDiscoverySpec_StorageConfig_CsvOptions_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoverySpec_StorageConfig_CsvOptions) *krm.DataDiscoverySpec_StorageConfig_CsvOptions {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoverySpec_StorageConfig_CsvOptions{}
	out.HeaderRows = direct.LazyPtr(in.GetHeaderRows())
	out.Delimiter = direct.LazyPtr(in.GetDelimiter())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.TypeInferenceDisabled = direct.LazyPtr(in.GetTypeInferenceDisabled())
	out.Quote = direct.LazyPtr(in.GetQuote())
	return out
}
func DataDiscoverySpec_StorageConfig_CsvOptions_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoverySpec_StorageConfig_CsvOptions) *pb.DataDiscoverySpec_StorageConfig_CsvOptions {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoverySpec_StorageConfig_CsvOptions{}
	out.HeaderRows = direct.ValueOf(in.HeaderRows)
	out.Delimiter = direct.ValueOf(in.Delimiter)
	out.Encoding = direct.ValueOf(in.Encoding)
	out.TypeInferenceDisabled = direct.ValueOf(in.TypeInferenceDisabled)
	out.Quote = direct.ValueOf(in.Quote)
	return out
}
func DataDiscoverySpec_StorageConfig_JsonOptions_FromProto(mapCtx *direct.MapContext, in *pb.DataDiscoverySpec_StorageConfig_JsonOptions) *krm.DataDiscoverySpec_StorageConfig_JsonOptions {
	if in == nil {
		return nil
	}
	out := &krm.DataDiscoverySpec_StorageConfig_JsonOptions{}
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.TypeInferenceDisabled = direct.LazyPtr(in.GetTypeInferenceDisabled())
	return out
}
func DataDiscoverySpec_StorageConfig_JsonOptions_ToProto(mapCtx *direct.MapContext, in *krm.DataDiscoverySpec_StorageConfig_JsonOptions) *pb.DataDiscoverySpec_StorageConfig_JsonOptions {
	if in == nil {
		return nil
	}
	out := &pb.DataDiscoverySpec_StorageConfig_JsonOptions{}
	out.Encoding = direct.ValueOf(in.Encoding)
	out.TypeInferenceDisabled = direct.ValueOf(in.TypeInferenceDisabled)
	return out
}
func DataProfileResult_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult) *krm.DataProfileResult {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult{}
	out.RowCount = direct.LazyPtr(in.GetRowCount())
	out.Profile = DataProfileResult_Profile_FromProto(mapCtx, in.GetProfile())
	out.ScannedData = ScannedData_FromProto(mapCtx, in.GetScannedData())
	// MISSING: PostScanActionsResult
	return out
}
func DataProfileResult_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult) *pb.DataProfileResult {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult{}
	out.RowCount = direct.ValueOf(in.RowCount)
	out.Profile = DataProfileResult_Profile_ToProto(mapCtx, in.Profile)
	out.ScannedData = ScannedData_ToProto(mapCtx, in.ScannedData)
	// MISSING: PostScanActionsResult
	return out
}
func DataProfileResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult) *krm.DataProfileResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResultObservedState{}
	// MISSING: RowCount
	// MISSING: Profile
	// MISSING: ScannedData
	out.PostScanActionsResult = DataProfileResult_PostScanActionsResult_FromProto(mapCtx, in.GetPostScanActionsResult())
	return out
}
func DataProfileResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResultObservedState) *pb.DataProfileResult {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult{}
	// MISSING: RowCount
	// MISSING: Profile
	// MISSING: ScannedData
	out.PostScanActionsResult = DataProfileResult_PostScanActionsResult_ToProto(mapCtx, in.PostScanActionsResult)
	return out
}
func DataProfileResult_PostScanActionsResult_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_PostScanActionsResult) *krm.DataProfileResult_PostScanActionsResult {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_PostScanActionsResult{}
	// MISSING: BigqueryExportResult
	return out
}
func DataProfileResult_PostScanActionsResult_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_PostScanActionsResult) *pb.DataProfileResult_PostScanActionsResult {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_PostScanActionsResult{}
	// MISSING: BigqueryExportResult
	return out
}
func DataProfileResult_PostScanActionsResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_PostScanActionsResult) *krm.DataProfileResult_PostScanActionsResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_PostScanActionsResultObservedState{}
	out.BigqueryExportResult = DataProfileResult_PostScanActionsResult_BigQueryExportResult_FromProto(mapCtx, in.GetBigqueryExportResult())
	return out
}
func DataProfileResult_PostScanActionsResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_PostScanActionsResultObservedState) *pb.DataProfileResult_PostScanActionsResult {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_PostScanActionsResult{}
	out.BigqueryExportResult = DataProfileResult_PostScanActionsResult_BigQueryExportResult_ToProto(mapCtx, in.BigqueryExportResult)
	return out
}
func DataProfileResult_PostScanActionsResult_BigQueryExportResult_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_PostScanActionsResult_BigQueryExportResult) *krm.DataProfileResult_PostScanActionsResult_BigQueryExportResult {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_PostScanActionsResult_BigQueryExportResult{}
	// MISSING: State
	// MISSING: Message
	return out
}
func DataProfileResult_PostScanActionsResult_BigQueryExportResult_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_PostScanActionsResult_BigQueryExportResult) *pb.DataProfileResult_PostScanActionsResult_BigQueryExportResult {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_PostScanActionsResult_BigQueryExportResult{}
	// MISSING: State
	// MISSING: Message
	return out
}
func DataProfileResult_PostScanActionsResult_BigQueryExportResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_PostScanActionsResult_BigQueryExportResult) *krm.DataProfileResult_PostScanActionsResult_BigQueryExportResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_PostScanActionsResult_BigQueryExportResultObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func DataProfileResult_PostScanActionsResult_BigQueryExportResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_PostScanActionsResult_BigQueryExportResultObservedState) *pb.DataProfileResult_PostScanActionsResult_BigQueryExportResult {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_PostScanActionsResult_BigQueryExportResult{}
	out.State = direct.Enum_ToProto[pb.DataProfileResult_PostScanActionsResult_BigQueryExportResult_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	return out
}
func DataProfileResult_Profile_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_Profile) *krm.DataProfileResult_Profile {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_Profile{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, DataProfileResult_Profile_Field_FromProto)
	return out
}
func DataProfileResult_Profile_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_Profile) *pb.DataProfileResult_Profile {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_Profile{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, DataProfileResult_Profile_Field_ToProto)
	return out
}
func DataProfileResult_Profile_Field_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_Profile_Field) *krm.DataProfileResult_Profile_Field {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_Profile_Field{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Mode = direct.LazyPtr(in.GetMode())
	out.Profile = DataProfileResult_Profile_Field_ProfileInfo_FromProto(mapCtx, in.GetProfile())
	return out
}
func DataProfileResult_Profile_Field_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_Profile_Field) *pb.DataProfileResult_Profile_Field {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_Profile_Field{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.Mode = direct.ValueOf(in.Mode)
	out.Profile = DataProfileResult_Profile_Field_ProfileInfo_ToProto(mapCtx, in.Profile)
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_Profile_Field_ProfileInfo) *krm.DataProfileResult_Profile_Field_ProfileInfo {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_Profile_Field_ProfileInfo{}
	out.NullRatio = direct.LazyPtr(in.GetNullRatio())
	out.DistinctRatio = direct.LazyPtr(in.GetDistinctRatio())
	out.TopNValues = direct.Slice_FromProto(mapCtx, in.TopNValues, DataProfileResult_Profile_Field_ProfileInfo_TopNValue_FromProto)
	out.StringProfile = DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo_FromProto(mapCtx, in.GetStringProfile())
	out.IntegerProfile = DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo_FromProto(mapCtx, in.GetIntegerProfile())
	out.DoubleProfile = DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo_FromProto(mapCtx, in.GetDoubleProfile())
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_Profile_Field_ProfileInfo) *pb.DataProfileResult_Profile_Field_ProfileInfo {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_Profile_Field_ProfileInfo{}
	out.NullRatio = direct.ValueOf(in.NullRatio)
	out.DistinctRatio = direct.ValueOf(in.DistinctRatio)
	out.TopNValues = direct.Slice_ToProto(mapCtx, in.TopNValues, DataProfileResult_Profile_Field_ProfileInfo_TopNValue_ToProto)
	if oneof := DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo_ToProto(mapCtx, in.StringProfile); oneof != nil {
		out.FieldInfo = &pb.DataProfileResult_Profile_Field_ProfileInfo_StringProfile{StringProfile: oneof}
	}
	if oneof := DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo_ToProto(mapCtx, in.IntegerProfile); oneof != nil {
		out.FieldInfo = &pb.DataProfileResult_Profile_Field_ProfileInfo_IntegerProfile{IntegerProfile: oneof}
	}
	if oneof := DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo_ToProto(mapCtx, in.DoubleProfile); oneof != nil {
		out.FieldInfo = &pb.DataProfileResult_Profile_Field_ProfileInfo_DoubleProfile{DoubleProfile: oneof}
	}
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo) *krm.DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo{}
	out.Average = direct.LazyPtr(in.GetAverage())
	out.StandardDeviation = direct.LazyPtr(in.GetStandardDeviation())
	out.Min = direct.LazyPtr(in.GetMin())
	out.Quartiles = in.Quartiles
	out.Max = direct.LazyPtr(in.GetMax())
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo) *pb.DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_Profile_Field_ProfileInfo_DoubleFieldInfo{}
	out.Average = direct.ValueOf(in.Average)
	out.StandardDeviation = direct.ValueOf(in.StandardDeviation)
	out.Min = direct.ValueOf(in.Min)
	out.Quartiles = in.Quartiles
	out.Max = direct.ValueOf(in.Max)
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo) *krm.DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo{}
	out.Average = direct.LazyPtr(in.GetAverage())
	out.StandardDeviation = direct.LazyPtr(in.GetStandardDeviation())
	out.Min = direct.LazyPtr(in.GetMin())
	out.Quartiles = in.Quartiles
	out.Max = direct.LazyPtr(in.GetMax())
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo) *pb.DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_Profile_Field_ProfileInfo_IntegerFieldInfo{}
	out.Average = direct.ValueOf(in.Average)
	out.StandardDeviation = direct.ValueOf(in.StandardDeviation)
	out.Min = direct.ValueOf(in.Min)
	out.Quartiles = in.Quartiles
	out.Max = direct.ValueOf(in.Max)
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo) *krm.DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo{}
	out.MinLength = direct.LazyPtr(in.GetMinLength())
	out.MaxLength = direct.LazyPtr(in.GetMaxLength())
	out.AverageLength = direct.LazyPtr(in.GetAverageLength())
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo) *pb.DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_Profile_Field_ProfileInfo_StringFieldInfo{}
	out.MinLength = direct.ValueOf(in.MinLength)
	out.MaxLength = direct.ValueOf(in.MaxLength)
	out.AverageLength = direct.ValueOf(in.AverageLength)
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_TopNValue_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileResult_Profile_Field_ProfileInfo_TopNValue) *krm.DataProfileResult_Profile_Field_ProfileInfo_TopNValue {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileResult_Profile_Field_ProfileInfo_TopNValue{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.Count = direct.LazyPtr(in.GetCount())
	out.Ratio = direct.LazyPtr(in.GetRatio())
	return out
}
func DataProfileResult_Profile_Field_ProfileInfo_TopNValue_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileResult_Profile_Field_ProfileInfo_TopNValue) *pb.DataProfileResult_Profile_Field_ProfileInfo_TopNValue {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileResult_Profile_Field_ProfileInfo_TopNValue{}
	out.Value = direct.ValueOf(in.Value)
	out.Count = direct.ValueOf(in.Count)
	out.Ratio = direct.ValueOf(in.Ratio)
	return out
}
func DataProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileSpec) *krm.DataProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileSpec{}
	out.SamplingPercent = direct.LazyPtr(in.GetSamplingPercent())
	out.RowFilter = direct.LazyPtr(in.GetRowFilter())
	out.PostScanActions = DataProfileSpec_PostScanActions_FromProto(mapCtx, in.GetPostScanActions())
	out.IncludeFields = DataProfileSpec_SelectedFields_FromProto(mapCtx, in.GetIncludeFields())
	out.ExcludeFields = DataProfileSpec_SelectedFields_FromProto(mapCtx, in.GetExcludeFields())
	return out
}
func DataProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileSpec) *pb.DataProfileSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileSpec{}
	out.SamplingPercent = direct.ValueOf(in.SamplingPercent)
	out.RowFilter = direct.ValueOf(in.RowFilter)
	out.PostScanActions = DataProfileSpec_PostScanActions_ToProto(mapCtx, in.PostScanActions)
	out.IncludeFields = DataProfileSpec_SelectedFields_ToProto(mapCtx, in.IncludeFields)
	out.ExcludeFields = DataProfileSpec_SelectedFields_ToProto(mapCtx, in.ExcludeFields)
	return out
}
func DataProfileSpec_PostScanActions_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileSpec_PostScanActions) *krm.DataProfileSpec_PostScanActions {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileSpec_PostScanActions{}
	out.BigqueryExport = DataProfileSpec_PostScanActions_BigQueryExport_FromProto(mapCtx, in.GetBigqueryExport())
	return out
}
func DataProfileSpec_PostScanActions_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileSpec_PostScanActions) *pb.DataProfileSpec_PostScanActions {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileSpec_PostScanActions{}
	out.BigqueryExport = DataProfileSpec_PostScanActions_BigQueryExport_ToProto(mapCtx, in.BigqueryExport)
	return out
}
func DataProfileSpec_PostScanActions_BigQueryExport_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileSpec_PostScanActions_BigQueryExport) *krm.DataProfileSpec_PostScanActions_BigQueryExport {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileSpec_PostScanActions_BigQueryExport{}
	out.ResultsTable = direct.LazyPtr(in.GetResultsTable())
	return out
}
func DataProfileSpec_PostScanActions_BigQueryExport_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileSpec_PostScanActions_BigQueryExport) *pb.DataProfileSpec_PostScanActions_BigQueryExport {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileSpec_PostScanActions_BigQueryExport{}
	out.ResultsTable = direct.ValueOf(in.ResultsTable)
	return out
}
func DataProfileSpec_SelectedFields_FromProto(mapCtx *direct.MapContext, in *pb.DataProfileSpec_SelectedFields) *krm.DataProfileSpec_SelectedFields {
	if in == nil {
		return nil
	}
	out := &krm.DataProfileSpec_SelectedFields{}
	out.FieldNames = in.FieldNames
	return out
}
func DataProfileSpec_SelectedFields_ToProto(mapCtx *direct.MapContext, in *krm.DataProfileSpec_SelectedFields) *pb.DataProfileSpec_SelectedFields {
	if in == nil {
		return nil
	}
	out := &pb.DataProfileSpec_SelectedFields{}
	out.FieldNames = in.FieldNames
	return out
}
func DataQualityColumnResult_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityColumnResult) *krm.DataQualityColumnResult {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityColumnResult{}
	// MISSING: Column
	// MISSING: Score
	return out
}
func DataQualityColumnResult_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityColumnResult) *pb.DataQualityColumnResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityColumnResult{}
	// MISSING: Column
	// MISSING: Score
	return out
}
func DataQualityColumnResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityColumnResult) *krm.DataQualityColumnResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityColumnResultObservedState{}
	out.Column = direct.LazyPtr(in.GetColumn())
	out.Score = in.Score
	return out
}
func DataQualityColumnResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityColumnResultObservedState) *pb.DataQualityColumnResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityColumnResult{}
	out.Column = direct.ValueOf(in.Column)
	out.Score = in.Score
	return out
}
func DataQualityDimension_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityDimension) *krm.DataQualityDimension {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityDimension{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func DataQualityDimension_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityDimension) *pb.DataQualityDimension {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityDimension{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func DataQualityDimensionResult_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityDimensionResult) *krm.DataQualityDimensionResult {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityDimensionResult{}
	// MISSING: Dimension
	out.Passed = direct.LazyPtr(in.GetPassed())
	// MISSING: Score
	return out
}
func DataQualityDimensionResult_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityDimensionResult) *pb.DataQualityDimensionResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityDimensionResult{}
	// MISSING: Dimension
	out.Passed = direct.ValueOf(in.Passed)
	// MISSING: Score
	return out
}
func DataQualityDimensionResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityDimensionResult) *krm.DataQualityDimensionResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityDimensionResultObservedState{}
	out.Dimension = DataQualityDimension_FromProto(mapCtx, in.GetDimension())
	// MISSING: Passed
	out.Score = in.Score
	return out
}
func DataQualityDimensionResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityDimensionResultObservedState) *pb.DataQualityDimensionResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityDimensionResult{}
	out.Dimension = DataQualityDimension_ToProto(mapCtx, in.Dimension)
	// MISSING: Passed
	out.Score = in.Score
	return out
}
func DataQualityResult_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityResult) *krm.DataQualityResult {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityResult{}
	out.Passed = direct.LazyPtr(in.GetPassed())
	// MISSING: Score
	out.Dimensions = direct.Slice_FromProto(mapCtx, in.Dimensions, DataQualityDimensionResult_FromProto)
	// MISSING: Columns
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, DataQualityRuleResult_FromProto)
	out.RowCount = direct.LazyPtr(in.GetRowCount())
	out.ScannedData = ScannedData_FromProto(mapCtx, in.GetScannedData())
	// MISSING: PostScanActionsResult
	return out
}
func DataQualityResult_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityResult) *pb.DataQualityResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityResult{}
	out.Passed = direct.ValueOf(in.Passed)
	// MISSING: Score
	out.Dimensions = direct.Slice_ToProto(mapCtx, in.Dimensions, DataQualityDimensionResult_ToProto)
	// MISSING: Columns
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, DataQualityRuleResult_ToProto)
	out.RowCount = direct.ValueOf(in.RowCount)
	out.ScannedData = ScannedData_ToProto(mapCtx, in.ScannedData)
	// MISSING: PostScanActionsResult
	return out
}
func DataQualityResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityResult) *krm.DataQualityResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityResultObservedState{}
	// MISSING: Passed
	out.Score = in.Score
	out.Dimensions = direct.Slice_FromProto(mapCtx, in.Dimensions, DataQualityDimensionResultObservedState_FromProto)
	out.Columns = direct.Slice_FromProto(mapCtx, in.Columns, DataQualityColumnResult_FromProto)
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, DataQualityRuleResultObservedState_FromProto)
	// MISSING: RowCount
	// MISSING: ScannedData
	out.PostScanActionsResult = DataQualityResult_PostScanActionsResult_FromProto(mapCtx, in.GetPostScanActionsResult())
	return out
}
func DataQualityResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityResultObservedState) *pb.DataQualityResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityResult{}
	// MISSING: Passed
	out.Score = in.Score
	out.Dimensions = direct.Slice_ToProto(mapCtx, in.Dimensions, DataQualityDimensionResultObservedState_ToProto)
	out.Columns = direct.Slice_ToProto(mapCtx, in.Columns, DataQualityColumnResult_ToProto)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, DataQualityRuleResultObservedState_ToProto)
	// MISSING: RowCount
	// MISSING: ScannedData
	out.PostScanActionsResult = DataQualityResult_PostScanActionsResult_ToProto(mapCtx, in.PostScanActionsResult)
	return out
}
func DataQualityResult_PostScanActionsResult_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityResult_PostScanActionsResult) *krm.DataQualityResult_PostScanActionsResult {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityResult_PostScanActionsResult{}
	// MISSING: BigqueryExportResult
	return out
}
func DataQualityResult_PostScanActionsResult_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityResult_PostScanActionsResult) *pb.DataQualityResult_PostScanActionsResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityResult_PostScanActionsResult{}
	// MISSING: BigqueryExportResult
	return out
}
func DataQualityResult_PostScanActionsResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityResult_PostScanActionsResult) *krm.DataQualityResult_PostScanActionsResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityResult_PostScanActionsResultObservedState{}
	out.BigqueryExportResult = DataQualityResult_PostScanActionsResult_BigQueryExportResult_FromProto(mapCtx, in.GetBigqueryExportResult())
	return out
}
func DataQualityResult_PostScanActionsResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityResult_PostScanActionsResultObservedState) *pb.DataQualityResult_PostScanActionsResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityResult_PostScanActionsResult{}
	out.BigqueryExportResult = DataQualityResult_PostScanActionsResult_BigQueryExportResult_ToProto(mapCtx, in.BigqueryExportResult)
	return out
}
func DataQualityResult_PostScanActionsResult_BigQueryExportResult_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityResult_PostScanActionsResult_BigQueryExportResult) *krm.DataQualityResult_PostScanActionsResult_BigQueryExportResult {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityResult_PostScanActionsResult_BigQueryExportResult{}
	// MISSING: State
	// MISSING: Message
	return out
}
func DataQualityResult_PostScanActionsResult_BigQueryExportResult_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityResult_PostScanActionsResult_BigQueryExportResult) *pb.DataQualityResult_PostScanActionsResult_BigQueryExportResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityResult_PostScanActionsResult_BigQueryExportResult{}
	// MISSING: State
	// MISSING: Message
	return out
}
func DataQualityResult_PostScanActionsResult_BigQueryExportResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityResult_PostScanActionsResult_BigQueryExportResult) *krm.DataQualityResult_PostScanActionsResult_BigQueryExportResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityResult_PostScanActionsResult_BigQueryExportResultObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func DataQualityResult_PostScanActionsResult_BigQueryExportResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityResult_PostScanActionsResult_BigQueryExportResultObservedState) *pb.DataQualityResult_PostScanActionsResult_BigQueryExportResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityResult_PostScanActionsResult_BigQueryExportResult{}
	out.State = direct.Enum_ToProto[pb.DataQualityResult_PostScanActionsResult_BigQueryExportResult_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	return out
}
func DataQualityRule_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule) *krm.DataQualityRule {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule{}
	out.RangeExpectation = DataQualityRule_RangeExpectation_FromProto(mapCtx, in.GetRangeExpectation())
	out.NonNullExpectation = DataQualityRule_NonNullExpectation_FromProto(mapCtx, in.GetNonNullExpectation())
	out.SetExpectation = DataQualityRule_SetExpectation_FromProto(mapCtx, in.GetSetExpectation())
	out.RegexExpectation = DataQualityRule_RegexExpectation_FromProto(mapCtx, in.GetRegexExpectation())
	out.UniquenessExpectation = DataQualityRule_UniquenessExpectation_FromProto(mapCtx, in.GetUniquenessExpectation())
	out.StatisticRangeExpectation = DataQualityRule_StatisticRangeExpectation_FromProto(mapCtx, in.GetStatisticRangeExpectation())
	out.RowConditionExpectation = DataQualityRule_RowConditionExpectation_FromProto(mapCtx, in.GetRowConditionExpectation())
	out.TableConditionExpectation = DataQualityRule_TableConditionExpectation_FromProto(mapCtx, in.GetTableConditionExpectation())
	out.SqlAssertion = DataQualityRule_SqlAssertion_FromProto(mapCtx, in.GetSqlAssertion())
	out.Column = direct.LazyPtr(in.GetColumn())
	out.IgnoreNull = direct.LazyPtr(in.GetIgnoreNull())
	out.Dimension = direct.LazyPtr(in.GetDimension())
	out.Threshold = direct.LazyPtr(in.GetThreshold())
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	return out
}
func DataQualityRule_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule) *pb.DataQualityRule {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule{}
	if oneof := DataQualityRule_RangeExpectation_ToProto(mapCtx, in.RangeExpectation); oneof != nil {
		out.RuleType = &pb.DataQualityRule_RangeExpectation_{RangeExpectation: oneof}
	}
	if oneof := DataQualityRule_NonNullExpectation_ToProto(mapCtx, in.NonNullExpectation); oneof != nil {
		out.RuleType = &pb.DataQualityRule_NonNullExpectation_{NonNullExpectation: oneof}
	}
	if oneof := DataQualityRule_SetExpectation_ToProto(mapCtx, in.SetExpectation); oneof != nil {
		out.RuleType = &pb.DataQualityRule_SetExpectation_{SetExpectation: oneof}
	}
	if oneof := DataQualityRule_RegexExpectation_ToProto(mapCtx, in.RegexExpectation); oneof != nil {
		out.RuleType = &pb.DataQualityRule_RegexExpectation_{RegexExpectation: oneof}
	}
	if oneof := DataQualityRule_UniquenessExpectation_ToProto(mapCtx, in.UniquenessExpectation); oneof != nil {
		out.RuleType = &pb.DataQualityRule_UniquenessExpectation_{UniquenessExpectation: oneof}
	}
	if oneof := DataQualityRule_StatisticRangeExpectation_ToProto(mapCtx, in.StatisticRangeExpectation); oneof != nil {
		out.RuleType = &pb.DataQualityRule_StatisticRangeExpectation_{StatisticRangeExpectation: oneof}
	}
	if oneof := DataQualityRule_RowConditionExpectation_ToProto(mapCtx, in.RowConditionExpectation); oneof != nil {
		out.RuleType = &pb.DataQualityRule_RowConditionExpectation_{RowConditionExpectation: oneof}
	}
	if oneof := DataQualityRule_TableConditionExpectation_ToProto(mapCtx, in.TableConditionExpectation); oneof != nil {
		out.RuleType = &pb.DataQualityRule_TableConditionExpectation_{TableConditionExpectation: oneof}
	}
	if oneof := DataQualityRule_SqlAssertion_ToProto(mapCtx, in.SqlAssertion); oneof != nil {
		out.RuleType = &pb.DataQualityRule_SqlAssertion_{SqlAssertion: oneof}
	}
	out.Column = direct.ValueOf(in.Column)
	out.IgnoreNull = direct.ValueOf(in.IgnoreNull)
	out.Dimension = direct.ValueOf(in.Dimension)
	out.Threshold = direct.ValueOf(in.Threshold)
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Suspended = direct.ValueOf(in.Suspended)
	return out
}
func DataQualityRuleResult_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRuleResult) *krm.DataQualityRuleResult {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRuleResult{}
	out.Rule = DataQualityRule_FromProto(mapCtx, in.GetRule())
	out.Passed = direct.LazyPtr(in.GetPassed())
	out.EvaluatedCount = direct.LazyPtr(in.GetEvaluatedCount())
	out.PassedCount = direct.LazyPtr(in.GetPassedCount())
	out.NullCount = direct.LazyPtr(in.GetNullCount())
	out.PassRatio = direct.LazyPtr(in.GetPassRatio())
	out.FailingRowsQuery = direct.LazyPtr(in.GetFailingRowsQuery())
	// MISSING: AssertionRowCount
	return out
}
func DataQualityRuleResult_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRuleResult) *pb.DataQualityRuleResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRuleResult{}
	out.Rule = DataQualityRule_ToProto(mapCtx, in.Rule)
	out.Passed = direct.ValueOf(in.Passed)
	out.EvaluatedCount = direct.ValueOf(in.EvaluatedCount)
	out.PassedCount = direct.ValueOf(in.PassedCount)
	out.NullCount = direct.ValueOf(in.NullCount)
	out.PassRatio = direct.ValueOf(in.PassRatio)
	out.FailingRowsQuery = direct.ValueOf(in.FailingRowsQuery)
	// MISSING: AssertionRowCount
	return out
}
func DataQualityRuleResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRuleResult) *krm.DataQualityRuleResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRuleResultObservedState{}
	// MISSING: Rule
	// MISSING: Passed
	// MISSING: EvaluatedCount
	// MISSING: PassedCount
	// MISSING: NullCount
	// MISSING: PassRatio
	// MISSING: FailingRowsQuery
	out.AssertionRowCount = direct.LazyPtr(in.GetAssertionRowCount())
	return out
}
func DataQualityRuleResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRuleResultObservedState) *pb.DataQualityRuleResult {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRuleResult{}
	// MISSING: Rule
	// MISSING: Passed
	// MISSING: EvaluatedCount
	// MISSING: PassedCount
	// MISSING: NullCount
	// MISSING: PassRatio
	// MISSING: FailingRowsQuery
	out.AssertionRowCount = direct.ValueOf(in.AssertionRowCount)
	return out
}
func DataQualityRule_NonNullExpectation_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_NonNullExpectation) *krm.DataQualityRule_NonNullExpectation {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_NonNullExpectation{}
	return out
}
func DataQualityRule_NonNullExpectation_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_NonNullExpectation) *pb.DataQualityRule_NonNullExpectation {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_NonNullExpectation{}
	return out
}
func DataQualityRule_RangeExpectation_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_RangeExpectation) *krm.DataQualityRule_RangeExpectation {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_RangeExpectation{}
	out.MinValue = direct.LazyPtr(in.GetMinValue())
	out.MaxValue = direct.LazyPtr(in.GetMaxValue())
	out.StrictMinEnabled = direct.LazyPtr(in.GetStrictMinEnabled())
	out.StrictMaxEnabled = direct.LazyPtr(in.GetStrictMaxEnabled())
	return out
}
func DataQualityRule_RangeExpectation_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_RangeExpectation) *pb.DataQualityRule_RangeExpectation {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_RangeExpectation{}
	out.MinValue = direct.ValueOf(in.MinValue)
	out.MaxValue = direct.ValueOf(in.MaxValue)
	out.StrictMinEnabled = direct.ValueOf(in.StrictMinEnabled)
	out.StrictMaxEnabled = direct.ValueOf(in.StrictMaxEnabled)
	return out
}
func DataQualityRule_RegexExpectation_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_RegexExpectation) *krm.DataQualityRule_RegexExpectation {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_RegexExpectation{}
	out.Regex = direct.LazyPtr(in.GetRegex())
	return out
}
func DataQualityRule_RegexExpectation_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_RegexExpectation) *pb.DataQualityRule_RegexExpectation {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_RegexExpectation{}
	out.Regex = direct.ValueOf(in.Regex)
	return out
}
func DataQualityRule_RowConditionExpectation_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_RowConditionExpectation) *krm.DataQualityRule_RowConditionExpectation {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_RowConditionExpectation{}
	out.SqlExpression = direct.LazyPtr(in.GetSqlExpression())
	return out
}
func DataQualityRule_RowConditionExpectation_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_RowConditionExpectation) *pb.DataQualityRule_RowConditionExpectation {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_RowConditionExpectation{}
	out.SqlExpression = direct.ValueOf(in.SqlExpression)
	return out
}
func DataQualityRule_SetExpectation_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_SetExpectation) *krm.DataQualityRule_SetExpectation {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_SetExpectation{}
	out.Values = in.Values
	return out
}
func DataQualityRule_SetExpectation_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_SetExpectation) *pb.DataQualityRule_SetExpectation {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_SetExpectation{}
	out.Values = in.Values
	return out
}
func DataQualityRule_SqlAssertion_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_SqlAssertion) *krm.DataQualityRule_SqlAssertion {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_SqlAssertion{}
	out.SqlStatement = direct.LazyPtr(in.GetSqlStatement())
	return out
}
func DataQualityRule_SqlAssertion_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_SqlAssertion) *pb.DataQualityRule_SqlAssertion {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_SqlAssertion{}
	out.SqlStatement = direct.ValueOf(in.SqlStatement)
	return out
}
func DataQualityRule_StatisticRangeExpectation_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_StatisticRangeExpectation) *krm.DataQualityRule_StatisticRangeExpectation {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_StatisticRangeExpectation{}
	out.Statistic = direct.Enum_FromProto(mapCtx, in.GetStatistic())
	out.MinValue = direct.LazyPtr(in.GetMinValue())
	out.MaxValue = direct.LazyPtr(in.GetMaxValue())
	out.StrictMinEnabled = direct.LazyPtr(in.GetStrictMinEnabled())
	out.StrictMaxEnabled = direct.LazyPtr(in.GetStrictMaxEnabled())
	return out
}
func DataQualityRule_StatisticRangeExpectation_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_StatisticRangeExpectation) *pb.DataQualityRule_StatisticRangeExpectation {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_StatisticRangeExpectation{}
	out.Statistic = direct.Enum_ToProto[pb.DataQualityRule_StatisticRangeExpectation_ColumnStatistic](mapCtx, in.Statistic)
	out.MinValue = direct.ValueOf(in.MinValue)
	out.MaxValue = direct.ValueOf(in.MaxValue)
	out.StrictMinEnabled = direct.ValueOf(in.StrictMinEnabled)
	out.StrictMaxEnabled = direct.ValueOf(in.StrictMaxEnabled)
	return out
}
func DataQualityRule_TableConditionExpectation_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_TableConditionExpectation) *krm.DataQualityRule_TableConditionExpectation {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_TableConditionExpectation{}
	out.SqlExpression = direct.LazyPtr(in.GetSqlExpression())
	return out
}
func DataQualityRule_TableConditionExpectation_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_TableConditionExpectation) *pb.DataQualityRule_TableConditionExpectation {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_TableConditionExpectation{}
	out.SqlExpression = direct.ValueOf(in.SqlExpression)
	return out
}
func DataQualityRule_UniquenessExpectation_FromProto(mapCtx *direct.MapContext, in *pb.DataQualityRule_UniquenessExpectation) *krm.DataQualityRule_UniquenessExpectation {
	if in == nil {
		return nil
	}
	out := &krm.DataQualityRule_UniquenessExpectation{}
	return out
}
func DataQualityRule_UniquenessExpectation_ToProto(mapCtx *direct.MapContext, in *krm.DataQualityRule_UniquenessExpectation) *pb.DataQualityRule_UniquenessExpectation {
	if in == nil {
		return nil
	}
	out := &pb.DataQualityRule_UniquenessExpectation{}
	return out
}
func DataQualitySpec_FromProto(mapCtx *direct.MapContext, in *pb.DataQualitySpec) *krm.DataQualitySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataQualitySpec{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, DataQualityRule_FromProto)
	out.SamplingPercent = direct.LazyPtr(in.GetSamplingPercent())
	out.RowFilter = direct.LazyPtr(in.GetRowFilter())
	out.PostScanActions = DataQualitySpec_PostScanActions_FromProto(mapCtx, in.GetPostScanActions())
	return out
}
func DataQualitySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataQualitySpec) *pb.DataQualitySpec {
	if in == nil {
		return nil
	}
	out := &pb.DataQualitySpec{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, DataQualityRule_ToProto)
	out.SamplingPercent = direct.ValueOf(in.SamplingPercent)
	out.RowFilter = direct.ValueOf(in.RowFilter)
	out.PostScanActions = DataQualitySpec_PostScanActions_ToProto(mapCtx, in.PostScanActions)
	return out
}
func DataQualitySpec_PostScanActions_FromProto(mapCtx *direct.MapContext, in *pb.DataQualitySpec_PostScanActions) *krm.DataQualitySpec_PostScanActions {
	if in == nil {
		return nil
	}
	out := &krm.DataQualitySpec_PostScanActions{}
	out.BigqueryExport = DataQualitySpec_PostScanActions_BigQueryExport_FromProto(mapCtx, in.GetBigqueryExport())
	out.NotificationReport = DataQualitySpec_PostScanActions_NotificationReport_FromProto(mapCtx, in.GetNotificationReport())
	return out
}
func DataQualitySpec_PostScanActions_ToProto(mapCtx *direct.MapContext, in *krm.DataQualitySpec_PostScanActions) *pb.DataQualitySpec_PostScanActions {
	if in == nil {
		return nil
	}
	out := &pb.DataQualitySpec_PostScanActions{}
	out.BigqueryExport = DataQualitySpec_PostScanActions_BigQueryExport_ToProto(mapCtx, in.BigqueryExport)
	out.NotificationReport = DataQualitySpec_PostScanActions_NotificationReport_ToProto(mapCtx, in.NotificationReport)
	return out
}
func DataQualitySpec_PostScanActions_BigQueryExport_FromProto(mapCtx *direct.MapContext, in *pb.DataQualitySpec_PostScanActions_BigQueryExport) *krm.DataQualitySpec_PostScanActions_BigQueryExport {
	if in == nil {
		return nil
	}
	out := &krm.DataQualitySpec_PostScanActions_BigQueryExport{}
	out.ResultsTable = direct.LazyPtr(in.GetResultsTable())
	return out
}
func DataQualitySpec_PostScanActions_BigQueryExport_ToProto(mapCtx *direct.MapContext, in *krm.DataQualitySpec_PostScanActions_BigQueryExport) *pb.DataQualitySpec_PostScanActions_BigQueryExport {
	if in == nil {
		return nil
	}
	out := &pb.DataQualitySpec_PostScanActions_BigQueryExport{}
	out.ResultsTable = direct.ValueOf(in.ResultsTable)
	return out
}
func DataQualitySpec_PostScanActions_JobEndTrigger_FromProto(mapCtx *direct.MapContext, in *pb.DataQualitySpec_PostScanActions_JobEndTrigger) *krm.DataQualitySpec_PostScanActions_JobEndTrigger {
	if in == nil {
		return nil
	}
	out := &krm.DataQualitySpec_PostScanActions_JobEndTrigger{}
	return out
}
func DataQualitySpec_PostScanActions_JobEndTrigger_ToProto(mapCtx *direct.MapContext, in *krm.DataQualitySpec_PostScanActions_JobEndTrigger) *pb.DataQualitySpec_PostScanActions_JobEndTrigger {
	if in == nil {
		return nil
	}
	out := &pb.DataQualitySpec_PostScanActions_JobEndTrigger{}
	return out
}
func DataQualitySpec_PostScanActions_JobFailureTrigger_FromProto(mapCtx *direct.MapContext, in *pb.DataQualitySpec_PostScanActions_JobFailureTrigger) *krm.DataQualitySpec_PostScanActions_JobFailureTrigger {
	if in == nil {
		return nil
	}
	out := &krm.DataQualitySpec_PostScanActions_JobFailureTrigger{}
	return out
}
func DataQualitySpec_PostScanActions_JobFailureTrigger_ToProto(mapCtx *direct.MapContext, in *krm.DataQualitySpec_PostScanActions_JobFailureTrigger) *pb.DataQualitySpec_PostScanActions_JobFailureTrigger {
	if in == nil {
		return nil
	}
	out := &pb.DataQualitySpec_PostScanActions_JobFailureTrigger{}
	return out
}
func DataQualitySpec_PostScanActions_NotificationReport_FromProto(mapCtx *direct.MapContext, in *pb.DataQualitySpec_PostScanActions_NotificationReport) *krm.DataQualitySpec_PostScanActions_NotificationReport {
	if in == nil {
		return nil
	}
	out := &krm.DataQualitySpec_PostScanActions_NotificationReport{}
	out.Recipients = DataQualitySpec_PostScanActions_Recipients_FromProto(mapCtx, in.GetRecipients())
	out.ScoreThresholdTrigger = DataQualitySpec_PostScanActions_ScoreThresholdTrigger_FromProto(mapCtx, in.GetScoreThresholdTrigger())
	out.JobFailureTrigger = DataQualitySpec_PostScanActions_JobFailureTrigger_FromProto(mapCtx, in.GetJobFailureTrigger())
	out.JobEndTrigger = DataQualitySpec_PostScanActions_JobEndTrigger_FromProto(mapCtx, in.GetJobEndTrigger())
	return out
}
func DataQualitySpec_PostScanActions_NotificationReport_ToProto(mapCtx *direct.MapContext, in *krm.DataQualitySpec_PostScanActions_NotificationReport) *pb.DataQualitySpec_PostScanActions_NotificationReport {
	if in == nil {
		return nil
	}
	out := &pb.DataQualitySpec_PostScanActions_NotificationReport{}
	out.Recipients = DataQualitySpec_PostScanActions_Recipients_ToProto(mapCtx, in.Recipients)
	out.ScoreThresholdTrigger = DataQualitySpec_PostScanActions_ScoreThresholdTrigger_ToProto(mapCtx, in.ScoreThresholdTrigger)
	out.JobFailureTrigger = DataQualitySpec_PostScanActions_JobFailureTrigger_ToProto(mapCtx, in.JobFailureTrigger)
	out.JobEndTrigger = DataQualitySpec_PostScanActions_JobEndTrigger_ToProto(mapCtx, in.JobEndTrigger)
	return out
}
func DataQualitySpec_PostScanActions_Recipients_FromProto(mapCtx *direct.MapContext, in *pb.DataQualitySpec_PostScanActions_Recipients) *krm.DataQualitySpec_PostScanActions_Recipients {
	if in == nil {
		return nil
	}
	out := &krm.DataQualitySpec_PostScanActions_Recipients{}
	out.Emails = in.Emails
	return out
}
func DataQualitySpec_PostScanActions_Recipients_ToProto(mapCtx *direct.MapContext, in *krm.DataQualitySpec_PostScanActions_Recipients) *pb.DataQualitySpec_PostScanActions_Recipients {
	if in == nil {
		return nil
	}
	out := &pb.DataQualitySpec_PostScanActions_Recipients{}
	out.Emails = in.Emails
	return out
}
func DataQualitySpec_PostScanActions_ScoreThresholdTrigger_FromProto(mapCtx *direct.MapContext, in *pb.DataQualitySpec_PostScanActions_ScoreThresholdTrigger) *krm.DataQualitySpec_PostScanActions_ScoreThresholdTrigger {
	if in == nil {
		return nil
	}
	out := &krm.DataQualitySpec_PostScanActions_ScoreThresholdTrigger{}
	out.ScoreThreshold = direct.LazyPtr(in.GetScoreThreshold())
	return out
}
func DataQualitySpec_PostScanActions_ScoreThresholdTrigger_ToProto(mapCtx *direct.MapContext, in *krm.DataQualitySpec_PostScanActions_ScoreThresholdTrigger) *pb.DataQualitySpec_PostScanActions_ScoreThresholdTrigger {
	if in == nil {
		return nil
	}
	out := &pb.DataQualitySpec_PostScanActions_ScoreThresholdTrigger{}
	out.ScoreThreshold = direct.ValueOf(in.ScoreThreshold)
	return out
}
func DataScan_FromProto(mapCtx *direct.MapContext, in *pb.DataScan) *krm.DataScan {
	if in == nil {
		return nil
	}
	out := &krm.DataScan{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Data = DataSource_FromProto(mapCtx, in.GetData())
	out.ExecutionSpec = DataScan_ExecutionSpec_FromProto(mapCtx, in.GetExecutionSpec())
	// MISSING: ExecutionStatus
	// MISSING: Type
	out.DataQualitySpec = DataQualitySpec_FromProto(mapCtx, in.GetDataQualitySpec())
	out.DataProfileSpec = DataProfileSpec_FromProto(mapCtx, in.GetDataProfileSpec())
	out.DataDiscoverySpec = DataDiscoverySpec_FromProto(mapCtx, in.GetDataDiscoverySpec())
	// MISSING: DataQualityResult
	// MISSING: DataProfileResult
	// MISSING: DataDiscoveryResult
	return out
}
func DataScan_ToProto(mapCtx *direct.MapContext, in *krm.DataScan) *pb.DataScan {
	if in == nil {
		return nil
	}
	out := &pb.DataScan{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Data = DataSource_ToProto(mapCtx, in.Data)
	out.ExecutionSpec = DataScan_ExecutionSpec_ToProto(mapCtx, in.ExecutionSpec)
	// MISSING: ExecutionStatus
	// MISSING: Type
	if oneof := DataQualitySpec_ToProto(mapCtx, in.DataQualitySpec); oneof != nil {
		out.Spec = &pb.DataScan_DataQualitySpec{DataQualitySpec: oneof}
	}
	if oneof := DataProfileSpec_ToProto(mapCtx, in.DataProfileSpec); oneof != nil {
		out.Spec = &pb.DataScan_DataProfileSpec{DataProfileSpec: oneof}
	}
	if oneof := DataDiscoverySpec_ToProto(mapCtx, in.DataDiscoverySpec); oneof != nil {
		out.Spec = &pb.DataScan_DataDiscoverySpec{DataDiscoverySpec: oneof}
	}
	// MISSING: DataQualityResult
	// MISSING: DataProfileResult
	// MISSING: DataDiscoveryResult
	return out
}
func DataScanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataScan) *krm.DataScanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataScanObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Data
	// MISSING: ExecutionSpec
	out.ExecutionStatus = DataScan_ExecutionStatus_FromProto(mapCtx, in.GetExecutionStatus())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: DataQualitySpec
	// MISSING: DataProfileSpec
	// MISSING: DataDiscoverySpec
	out.DataQualityResult = DataQualityResult_FromProto(mapCtx, in.GetDataQualityResult())
	out.DataProfileResult = DataProfileResult_FromProto(mapCtx, in.GetDataProfileResult())
	out.DataDiscoveryResult = DataDiscoveryResult_FromProto(mapCtx, in.GetDataDiscoveryResult())
	return out
}
func DataScanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataScanObservedState) *pb.DataScan {
	if in == nil {
		return nil
	}
	out := &pb.DataScan{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Data
	// MISSING: ExecutionSpec
	out.ExecutionStatus = DataScan_ExecutionStatus_ToProto(mapCtx, in.ExecutionStatus)
	out.Type = direct.Enum_ToProto[pb.DataScanType](mapCtx, in.Type)
	// MISSING: DataQualitySpec
	// MISSING: DataProfileSpec
	// MISSING: DataDiscoverySpec
	if oneof := DataQualityResult_ToProto(mapCtx, in.DataQualityResult); oneof != nil {
		out.Result = &pb.DataScan_DataQualityResult{DataQualityResult: oneof}
	}
	if oneof := DataProfileResult_ToProto(mapCtx, in.DataProfileResult); oneof != nil {
		out.Result = &pb.DataScan_DataProfileResult{DataProfileResult: oneof}
	}
	if oneof := DataDiscoveryResult_ToProto(mapCtx, in.DataDiscoveryResult); oneof != nil {
		out.Result = &pb.DataScan_DataDiscoveryResult{DataDiscoveryResult: oneof}
	}
	return out
}
func DataScan_ExecutionSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataScan_ExecutionSpec) *krm.DataScan_ExecutionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataScan_ExecutionSpec{}
	out.Trigger = Trigger_FromProto(mapCtx, in.GetTrigger())
	out.Field = direct.LazyPtr(in.GetField())
	return out
}
func DataScan_ExecutionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataScan_ExecutionSpec) *pb.DataScan_ExecutionSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataScan_ExecutionSpec{}
	out.Trigger = Trigger_ToProto(mapCtx, in.Trigger)
	if oneof := DataScan_ExecutionSpec_Field_ToProto(mapCtx, in.Field); oneof != nil {
		out.Incremental = oneof
	}
	return out
}
func DataScan_ExecutionStatus_FromProto(mapCtx *direct.MapContext, in *pb.DataScan_ExecutionStatus) *krm.DataScan_ExecutionStatus {
	if in == nil {
		return nil
	}
	out := &krm.DataScan_ExecutionStatus{}
	out.LatestJobStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLatestJobStartTime())
	out.LatestJobEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLatestJobEndTime())
	out.LatestJobCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLatestJobCreateTime())
	return out
}
func DataScan_ExecutionStatus_ToProto(mapCtx *direct.MapContext, in *krm.DataScan_ExecutionStatus) *pb.DataScan_ExecutionStatus {
	if in == nil {
		return nil
	}
	out := &pb.DataScan_ExecutionStatus{}
	out.LatestJobStartTime = direct.StringTimestamp_ToProto(mapCtx, in.LatestJobStartTime)
	out.LatestJobEndTime = direct.StringTimestamp_ToProto(mapCtx, in.LatestJobEndTime)
	out.LatestJobCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.LatestJobCreateTime)
	return out
}
func DataSource_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.DataSource {
	if in == nil {
		return nil
	}
	out := &krm.DataSource{}
	out.Entity = direct.LazyPtr(in.GetEntity())
	out.Resource = direct.LazyPtr(in.GetResource())
	return out
}
func DataSource_ToProto(mapCtx *direct.MapContext, in *krm.DataSource) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	if oneof := DataSource_Entity_ToProto(mapCtx, in.Entity); oneof != nil {
		out.Source = oneof
	}
	if oneof := DataSource_Resource_ToProto(mapCtx, in.Resource); oneof != nil {
		out.Source = oneof
	}
	return out
}
func DataplexDataScanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataScan) *krm.DataplexDataScanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataScanObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Data
	// MISSING: ExecutionSpec
	// MISSING: ExecutionStatus
	// MISSING: Type
	// MISSING: DataQualitySpec
	// MISSING: DataProfileSpec
	// MISSING: DataDiscoverySpec
	// MISSING: DataQualityResult
	// MISSING: DataProfileResult
	// MISSING: DataDiscoveryResult
	return out
}
func DataplexDataScanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataScanObservedState) *pb.DataScan {
	if in == nil {
		return nil
	}
	out := &pb.DataScan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Data
	// MISSING: ExecutionSpec
	// MISSING: ExecutionStatus
	// MISSING: Type
	// MISSING: DataQualitySpec
	// MISSING: DataProfileSpec
	// MISSING: DataDiscoverySpec
	// MISSING: DataQualityResult
	// MISSING: DataProfileResult
	// MISSING: DataDiscoveryResult
	return out
}
func DataplexDataScanSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataScan) *krm.DataplexDataScanSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataScanSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Data
	// MISSING: ExecutionSpec
	// MISSING: ExecutionStatus
	// MISSING: Type
	// MISSING: DataQualitySpec
	// MISSING: DataProfileSpec
	// MISSING: DataDiscoverySpec
	// MISSING: DataQualityResult
	// MISSING: DataProfileResult
	// MISSING: DataDiscoveryResult
	return out
}
func DataplexDataScanSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataScanSpec) *pb.DataScan {
	if in == nil {
		return nil
	}
	out := &pb.DataScan{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Data
	// MISSING: ExecutionSpec
	// MISSING: ExecutionStatus
	// MISSING: Type
	// MISSING: DataQualitySpec
	// MISSING: DataProfileSpec
	// MISSING: DataDiscoverySpec
	// MISSING: DataQualityResult
	// MISSING: DataProfileResult
	// MISSING: DataDiscoveryResult
	return out
}
func ScannedData_FromProto(mapCtx *direct.MapContext, in *pb.ScannedData) *krm.ScannedData {
	if in == nil {
		return nil
	}
	out := &krm.ScannedData{}
	out.IncrementalField = ScannedData_IncrementalField_FromProto(mapCtx, in.GetIncrementalField())
	return out
}
func ScannedData_ToProto(mapCtx *direct.MapContext, in *krm.ScannedData) *pb.ScannedData {
	if in == nil {
		return nil
	}
	out := &pb.ScannedData{}
	if oneof := ScannedData_IncrementalField_ToProto(mapCtx, in.IncrementalField); oneof != nil {
		out.DataRange = &pb.ScannedData_IncrementalField_{IncrementalField: oneof}
	}
	return out
}
func ScannedData_IncrementalField_FromProto(mapCtx *direct.MapContext, in *pb.ScannedData_IncrementalField) *krm.ScannedData_IncrementalField {
	if in == nil {
		return nil
	}
	out := &krm.ScannedData_IncrementalField{}
	out.Field = direct.LazyPtr(in.GetField())
	out.Start = direct.LazyPtr(in.GetStart())
	out.End = direct.LazyPtr(in.GetEnd())
	return out
}
func ScannedData_IncrementalField_ToProto(mapCtx *direct.MapContext, in *krm.ScannedData_IncrementalField) *pb.ScannedData_IncrementalField {
	if in == nil {
		return nil
	}
	out := &pb.ScannedData_IncrementalField{}
	out.Field = direct.ValueOf(in.Field)
	out.Start = direct.ValueOf(in.Start)
	out.End = direct.ValueOf(in.End)
	return out
}
func Trigger_FromProto(mapCtx *direct.MapContext, in *pb.Trigger) *krm.Trigger {
	if in == nil {
		return nil
	}
	out := &krm.Trigger{}
	out.OnDemand = Trigger_OnDemand_FromProto(mapCtx, in.GetOnDemand())
	out.Schedule = Trigger_Schedule_FromProto(mapCtx, in.GetSchedule())
	return out
}
func Trigger_ToProto(mapCtx *direct.MapContext, in *krm.Trigger) *pb.Trigger {
	if in == nil {
		return nil
	}
	out := &pb.Trigger{}
	if oneof := Trigger_OnDemand_ToProto(mapCtx, in.OnDemand); oneof != nil {
		out.Mode = &pb.Trigger_OnDemand_{OnDemand: oneof}
	}
	if oneof := Trigger_Schedule_ToProto(mapCtx, in.Schedule); oneof != nil {
		out.Mode = &pb.Trigger_Schedule_{Schedule: oneof}
	}
	return out
}
func Trigger_OnDemand_FromProto(mapCtx *direct.MapContext, in *pb.Trigger_OnDemand) *krm.Trigger_OnDemand {
	if in == nil {
		return nil
	}
	out := &krm.Trigger_OnDemand{}
	return out
}
func Trigger_OnDemand_ToProto(mapCtx *direct.MapContext, in *krm.Trigger_OnDemand) *pb.Trigger_OnDemand {
	if in == nil {
		return nil
	}
	out := &pb.Trigger_OnDemand{}
	return out
}
func Trigger_Schedule_FromProto(mapCtx *direct.MapContext, in *pb.Trigger_Schedule) *krm.Trigger_Schedule {
	if in == nil {
		return nil
	}
	out := &krm.Trigger_Schedule{}
	out.Cron = direct.LazyPtr(in.GetCron())
	return out
}
func Trigger_Schedule_ToProto(mapCtx *direct.MapContext, in *krm.Trigger_Schedule) *pb.Trigger_Schedule {
	if in == nil {
		return nil
	}
	out := &pb.Trigger_Schedule{}
	out.Cron = direct.ValueOf(in.Cron)
	return out
}
