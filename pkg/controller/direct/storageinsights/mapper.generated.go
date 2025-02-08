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

package storageinsights

import (
	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storageinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func CSVOptions_FromProto(mapCtx *direct.MapContext, in *pb.CSVOptions) *krm.CSVOptions {
	if in == nil {
		return nil
	}
	out := &krm.CSVOptions{}
	out.RecordSeparator = direct.LazyPtr(in.GetRecordSeparator())
	out.Delimiter = direct.LazyPtr(in.GetDelimiter())
	out.HeaderRequired = direct.LazyPtr(in.GetHeaderRequired())
	return out
}
func CSVOptions_ToProto(mapCtx *direct.MapContext, in *krm.CSVOptions) *pb.CSVOptions {
	if in == nil {
		return nil
	}
	out := &pb.CSVOptions{}
	out.RecordSeparator = direct.ValueOf(in.RecordSeparator)
	out.Delimiter = direct.ValueOf(in.Delimiter)
	out.HeaderRequired = direct.ValueOf(in.HeaderRequired)
	return out
}
func CloudStorageDestinationOptions_FromProto(mapCtx *direct.MapContext, in *pb.CloudStorageDestinationOptions) *krm.CloudStorageDestinationOptions {
	if in == nil {
		return nil
	}
	out := &krm.CloudStorageDestinationOptions{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.DestinationPath = direct.LazyPtr(in.GetDestinationPath())
	return out
}
func CloudStorageDestinationOptions_ToProto(mapCtx *direct.MapContext, in *krm.CloudStorageDestinationOptions) *pb.CloudStorageDestinationOptions {
	if in == nil {
		return nil
	}
	out := &pb.CloudStorageDestinationOptions{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.DestinationPath = direct.ValueOf(in.DestinationPath)
	return out
}
func CloudStorageFilters_FromProto(mapCtx *direct.MapContext, in *pb.CloudStorageFilters) *krm.CloudStorageFilters {
	if in == nil {
		return nil
	}
	out := &krm.CloudStorageFilters{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	return out
}
func CloudStorageFilters_ToProto(mapCtx *direct.MapContext, in *krm.CloudStorageFilters) *pb.CloudStorageFilters {
	if in == nil {
		return nil
	}
	out := &pb.CloudStorageFilters{}
	out.Bucket = direct.ValueOf(in.Bucket)
	return out
}
func FrequencyOptions_FromProto(mapCtx *direct.MapContext, in *pb.FrequencyOptions) *krm.FrequencyOptions {
	if in == nil {
		return nil
	}
	out := &krm.FrequencyOptions{}
	out.Frequency = direct.Enum_FromProto(mapCtx, in.GetFrequency())
	out.StartDate = Date_FromProto(mapCtx, in.GetStartDate())
	out.EndDate = Date_FromProto(mapCtx, in.GetEndDate())
	return out
}
func FrequencyOptions_ToProto(mapCtx *direct.MapContext, in *krm.FrequencyOptions) *pb.FrequencyOptions {
	if in == nil {
		return nil
	}
	out := &pb.FrequencyOptions{}
	out.Frequency = direct.Enum_ToProto[pb.FrequencyOptions_Frequency](mapCtx, in.Frequency)
	out.StartDate = Date_ToProto(mapCtx, in.StartDate)
	out.EndDate = Date_ToProto(mapCtx, in.EndDate)
	return out
}
func ObjectMetadataReportOptions_FromProto(mapCtx *direct.MapContext, in *pb.ObjectMetadataReportOptions) *krm.ObjectMetadataReportOptions {
	if in == nil {
		return nil
	}
	out := &krm.ObjectMetadataReportOptions{}
	out.MetadataFields = in.MetadataFields
	out.StorageFilters = CloudStorageFilters_FromProto(mapCtx, in.GetStorageFilters())
	out.StorageDestinationOptions = CloudStorageDestinationOptions_FromProto(mapCtx, in.GetStorageDestinationOptions())
	return out
}
func ObjectMetadataReportOptions_ToProto(mapCtx *direct.MapContext, in *krm.ObjectMetadataReportOptions) *pb.ObjectMetadataReportOptions {
	if in == nil {
		return nil
	}
	out := &pb.ObjectMetadataReportOptions{}
	out.MetadataFields = in.MetadataFields
	if oneof := CloudStorageFilters_ToProto(mapCtx, in.StorageFilters); oneof != nil {
		out.Filter = &pb.ObjectMetadataReportOptions_StorageFilters{StorageFilters: oneof}
	}
	if oneof := CloudStorageDestinationOptions_ToProto(mapCtx, in.StorageDestinationOptions); oneof != nil {
		out.DestinationOptions = &pb.ObjectMetadataReportOptions_StorageDestinationOptions{StorageDestinationOptions: oneof}
	}
	return out
}
func ParquetOptions_FromProto(mapCtx *direct.MapContext, in *pb.ParquetOptions) *krm.ParquetOptions {
	if in == nil {
		return nil
	}
	out := &krm.ParquetOptions{}
	return out
}
func ParquetOptions_ToProto(mapCtx *direct.MapContext, in *krm.ParquetOptions) *pb.ParquetOptions {
	if in == nil {
		return nil
	}
	out := &pb.ParquetOptions{}
	return out
}
func ReportConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig) *krm.ReportConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReportConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.FrequencyOptions = FrequencyOptions_FromProto(mapCtx, in.GetFrequencyOptions())
	out.CsvOptions = CSVOptions_FromProto(mapCtx, in.GetCsvOptions())
	out.ParquetOptions = ParquetOptions_FromProto(mapCtx, in.GetParquetOptions())
	out.ObjectMetadataReportOptions = ObjectMetadataReportOptions_FromProto(mapCtx, in.GetObjectMetadataReportOptions())
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func ReportConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReportConfig) *pb.ReportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.FrequencyOptions = FrequencyOptions_ToProto(mapCtx, in.FrequencyOptions)
	if oneof := CSVOptions_ToProto(mapCtx, in.CsvOptions); oneof != nil {
		out.ReportFormat = &pb.ReportConfig_CsvOptions{CsvOptions: oneof}
	}
	if oneof := ParquetOptions_ToProto(mapCtx, in.ParquetOptions); oneof != nil {
		out.ReportFormat = &pb.ReportConfig_ParquetOptions{ParquetOptions: oneof}
	}
	if oneof := ObjectMetadataReportOptions_ToProto(mapCtx, in.ObjectMetadataReportOptions); oneof != nil {
		out.ReportKind = &pb.ReportConfig_ObjectMetadataReportOptions{ObjectMetadataReportOptions: oneof}
	}
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func ReportConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig) *krm.ReportConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReportConfigObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: FrequencyOptions
	// MISSING: CsvOptions
	// MISSING: ParquetOptions
	// MISSING: ObjectMetadataReportOptions
	// MISSING: Labels
	// MISSING: DisplayName
	return out
}
func ReportConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReportConfigObservedState) *pb.ReportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: FrequencyOptions
	// MISSING: CsvOptions
	// MISSING: ParquetOptions
	// MISSING: ObjectMetadataReportOptions
	// MISSING: Labels
	// MISSING: DisplayName
	return out
}
func StorageinsightsReportConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig) *krm.StorageinsightsReportConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageinsightsReportConfigObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: FrequencyOptions
	// MISSING: CsvOptions
	// MISSING: ParquetOptions
	// MISSING: ObjectMetadataReportOptions
	// MISSING: Labels
	// MISSING: DisplayName
	return out
}
func StorageinsightsReportConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageinsightsReportConfigObservedState) *pb.ReportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: FrequencyOptions
	// MISSING: CsvOptions
	// MISSING: ParquetOptions
	// MISSING: ObjectMetadataReportOptions
	// MISSING: Labels
	// MISSING: DisplayName
	return out
}
func StorageinsightsReportConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig) *krm.StorageinsightsReportConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageinsightsReportConfigSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: FrequencyOptions
	// MISSING: CsvOptions
	// MISSING: ParquetOptions
	// MISSING: ObjectMetadataReportOptions
	// MISSING: Labels
	// MISSING: DisplayName
	return out
}
func StorageinsightsReportConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageinsightsReportConfigSpec) *pb.ReportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: FrequencyOptions
	// MISSING: CsvOptions
	// MISSING: ParquetOptions
	// MISSING: ObjectMetadataReportOptions
	// MISSING: Labels
	// MISSING: DisplayName
	return out
}
