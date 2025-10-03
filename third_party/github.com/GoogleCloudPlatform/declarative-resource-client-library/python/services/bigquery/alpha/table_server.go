// Copyright 2022 Google LLC. All Rights Reserved.
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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/alpha/bigquery_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/alpha"
)

// TableServer implements the gRPC interface for Table.
type TableServer struct{}

// ProtoToTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum converts a TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum enum from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(e alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum) *alpha.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum_name[int32(e)]; ok {
		e := alpha.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(n[len("BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum converts a TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum enum from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(e alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum) *alpha.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum_name[int32(e)]; ok {
		e := alpha.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(n[len("BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationDecimalTargetTypesEnum converts a TableExternalDataConfigurationDecimalTargetTypesEnum enum from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum(e alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum) *alpha.TableExternalDataConfigurationDecimalTargetTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum_name[int32(e)]; ok {
		e := alpha.TableExternalDataConfigurationDecimalTargetTypesEnum(n[len("BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationJsonExtensionEnum converts a TableExternalDataConfigurationJsonExtensionEnum enum from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum(e alphapb.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum) *alpha.TableExternalDataConfigurationJsonExtensionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum_name[int32(e)]; ok {
		e := alpha.TableExternalDataConfigurationJsonExtensionEnum(n[len("BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableModel converts a TableModel object from its proto representation.
func ProtoToBigqueryAlphaTableModel(p *alphapb.BigqueryAlphaTableModel) *alpha.TableModel {
	if p == nil {
		return nil
	}
	obj := &alpha.TableModel{
		ModelOptions: ProtoToBigqueryAlphaTableModelModelOptions(p.GetModelOptions()),
	}
	for _, r := range p.GetTrainingRuns() {
		obj.TrainingRuns = append(obj.TrainingRuns, *ProtoToBigqueryAlphaTableModelTrainingRuns(r))
	}
	return obj
}

// ProtoToTableModelModelOptions converts a TableModelModelOptions object from its proto representation.
func ProtoToBigqueryAlphaTableModelModelOptions(p *alphapb.BigqueryAlphaTableModelModelOptions) *alpha.TableModelModelOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.TableModelModelOptions{
		ModelType: dcl.StringOrNil(p.GetModelType()),
		LossType:  dcl.StringOrNil(p.GetLossType()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, r)
	}
	return obj
}

// ProtoToTableModelTrainingRuns converts a TableModelTrainingRuns object from its proto representation.
func ProtoToBigqueryAlphaTableModelTrainingRuns(p *alphapb.BigqueryAlphaTableModelTrainingRuns) *alpha.TableModelTrainingRuns {
	if p == nil {
		return nil
	}
	obj := &alpha.TableModelTrainingRuns{
		State:           dcl.StringOrNil(p.GetState()),
		StartTime:       dcl.StringOrNil(p.GetStartTime()),
		TrainingOptions: ProtoToBigqueryAlphaTableModelTrainingRunsTrainingOptions(p.GetTrainingOptions()),
	}
	for _, r := range p.GetIterationResults() {
		obj.IterationResults = append(obj.IterationResults, *ProtoToBigqueryAlphaTableModelTrainingRunsIterationResults(r))
	}
	return obj
}

// ProtoToTableModelTrainingRunsTrainingOptions converts a TableModelTrainingRunsTrainingOptions object from its proto representation.
func ProtoToBigqueryAlphaTableModelTrainingRunsTrainingOptions(p *alphapb.BigqueryAlphaTableModelTrainingRunsTrainingOptions) *alpha.TableModelTrainingRunsTrainingOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.TableModelTrainingRunsTrainingOptions{
		MaxIteration:            dcl.Int64OrNil(p.GetMaxIteration()),
		LearnRate:               dcl.Float64OrNil(p.GetLearnRate()),
		L1Reg:                   dcl.Float64OrNil(p.GetL1Reg()),
		L2Reg:                   dcl.Float64OrNil(p.GetL2Reg()),
		MinRelProgress:          dcl.Float64OrNil(p.GetMinRelProgress()),
		WarmStart:               dcl.Bool(p.GetWarmStart()),
		EarlyStop:               dcl.Bool(p.GetEarlyStop()),
		LearnRateStrategy:       dcl.StringOrNil(p.GetLearnRateStrategy()),
		LineSearchInitLearnRate: dcl.Float64OrNil(p.GetLineSearchInitLearnRate()),
	}
	return obj
}

// ProtoToTableModelTrainingRunsIterationResults converts a TableModelTrainingRunsIterationResults object from its proto representation.
func ProtoToBigqueryAlphaTableModelTrainingRunsIterationResults(p *alphapb.BigqueryAlphaTableModelTrainingRunsIterationResults) *alpha.TableModelTrainingRunsIterationResults {
	if p == nil {
		return nil
	}
	obj := &alpha.TableModelTrainingRunsIterationResults{
		Index:        dcl.Int64OrNil(p.GetIndex()),
		LearnRate:    dcl.Float64OrNil(p.GetLearnRate()),
		TrainingLoss: dcl.Float64OrNil(p.GetTrainingLoss()),
		EvalLoss:     dcl.Float64OrNil(p.GetEvalLoss()),
		DurationMs:   dcl.StringOrNil(p.GetDurationMs()),
	}
	return obj
}

// ProtoToTableSchema converts a TableSchema object from its proto representation.
func ProtoToBigqueryAlphaTableSchema(p *alphapb.BigqueryAlphaTableSchema) *alpha.TableSchema {
	if p == nil {
		return nil
	}
	obj := &alpha.TableSchema{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2Tablefieldschema converts a TableGooglecloudbigqueryv2Tablefieldschema object from its proto representation.
func ProtoToBigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema(p *alphapb.BigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema) *alpha.TableGooglecloudbigqueryv2Tablefieldschema {
	if p == nil {
		return nil
	}
	obj := &alpha.TableGooglecloudbigqueryv2Tablefieldschema{
		Name:                   dcl.StringOrNil(p.GetName()),
		Type:                   dcl.StringOrNil(p.GetType()),
		Mode:                   dcl.StringOrNil(p.GetMode()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		Categories:             ProtoToBigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaCategories(p.GetCategories()),
		PolicyTags:             ProtoToBigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(p.GetPolicyTags()),
		MaxLength:              dcl.Int64OrNil(p.GetMaxLength()),
		Precision:              dcl.Int64OrNil(p.GetPrecision()),
		Scale:                  dcl.Int64OrNil(p.GetScale()),
		Collation:              dcl.StringOrNil(p.GetCollation()),
		DefaultValueExpression: dcl.StringOrNil(p.GetDefaultValueExpression()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	for _, r := range p.GetNameAlternative() {
		obj.NameAlternative = append(obj.NameAlternative, r)
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2TablefieldschemaCategories converts a TableGooglecloudbigqueryv2TablefieldschemaCategories object from its proto representation.
func ProtoToBigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaCategories(p *alphapb.BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaCategories) *alpha.TableGooglecloudbigqueryv2TablefieldschemaCategories {
	if p == nil {
		return nil
	}
	obj := &alpha.TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	for _, r := range p.GetNames() {
		obj.Names = append(obj.Names, r)
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2TablefieldschemaPolicyTags converts a TableGooglecloudbigqueryv2TablefieldschemaPolicyTags object from its proto representation.
func ProtoToBigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(p *alphapb.BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags) *alpha.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if p == nil {
		return nil
	}
	obj := &alpha.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	for _, r := range p.GetNames() {
		obj.Names = append(obj.Names, r)
	}
	return obj
}

// ProtoToTableTimePartitioning converts a TableTimePartitioning object from its proto representation.
func ProtoToBigqueryAlphaTableTimePartitioning(p *alphapb.BigqueryAlphaTableTimePartitioning) *alpha.TableTimePartitioning {
	if p == nil {
		return nil
	}
	obj := &alpha.TableTimePartitioning{
		Type:         dcl.StringOrNil(p.GetType()),
		ExpirationMs: dcl.StringOrNil(p.GetExpirationMs()),
		Field:        dcl.StringOrNil(p.GetField()),
	}
	return obj
}

// ProtoToTableRangePartitioning converts a TableRangePartitioning object from its proto representation.
func ProtoToBigqueryAlphaTableRangePartitioning(p *alphapb.BigqueryAlphaTableRangePartitioning) *alpha.TableRangePartitioning {
	if p == nil {
		return nil
	}
	obj := &alpha.TableRangePartitioning{
		Field: dcl.StringOrNil(p.GetField()),
		Range: ProtoToBigqueryAlphaTableRangePartitioningRange(p.GetRange()),
	}
	return obj
}

// ProtoToTableRangePartitioningRange converts a TableRangePartitioningRange object from its proto representation.
func ProtoToBigqueryAlphaTableRangePartitioningRange(p *alphapb.BigqueryAlphaTableRangePartitioningRange) *alpha.TableRangePartitioningRange {
	if p == nil {
		return nil
	}
	obj := &alpha.TableRangePartitioningRange{
		Start:    dcl.StringOrNil(p.GetStart()),
		End:      dcl.StringOrNil(p.GetEnd()),
		Interval: dcl.StringOrNil(p.GetInterval()),
	}
	return obj
}

// ProtoToTableClustering converts a TableClustering object from its proto representation.
func ProtoToBigqueryAlphaTableClustering(p *alphapb.BigqueryAlphaTableClustering) *alpha.TableClustering {
	if p == nil {
		return nil
	}
	obj := &alpha.TableClustering{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToTableView converts a TableView object from its proto representation.
func ProtoToBigqueryAlphaTableView(p *alphapb.BigqueryAlphaTableView) *alpha.TableView {
	if p == nil {
		return nil
	}
	obj := &alpha.TableView{
		Query:                  dcl.StringOrNil(p.GetQuery()),
		UseLegacySql:           dcl.Bool(p.GetUseLegacySql()),
		UseExplicitColumnNames: dcl.Bool(p.GetUseExplicitColumnNames()),
	}
	for _, r := range p.GetUserDefinedFunctionResources() {
		obj.UserDefinedFunctionResources = append(obj.UserDefinedFunctionResources, *ProtoToBigqueryAlphaTableViewUserDefinedFunctionResources(r))
	}
	return obj
}

// ProtoToTableViewUserDefinedFunctionResources converts a TableViewUserDefinedFunctionResources object from its proto representation.
func ProtoToBigqueryAlphaTableViewUserDefinedFunctionResources(p *alphapb.BigqueryAlphaTableViewUserDefinedFunctionResources) *alpha.TableViewUserDefinedFunctionResources {
	if p == nil {
		return nil
	}
	obj := &alpha.TableViewUserDefinedFunctionResources{
		ResourceUri: dcl.StringOrNil(p.GetResourceUri()),
		InlineCode:  dcl.StringOrNil(p.GetInlineCode()),
	}
	for _, r := range p.GetInlineCodeAlternative() {
		obj.InlineCodeAlternative = append(obj.InlineCodeAlternative, r)
	}
	return obj
}

// ProtoToTableMaterializedView converts a TableMaterializedView object from its proto representation.
func ProtoToBigqueryAlphaTableMaterializedView(p *alphapb.BigqueryAlphaTableMaterializedView) *alpha.TableMaterializedView {
	if p == nil {
		return nil
	}
	obj := &alpha.TableMaterializedView{
		Query:             dcl.StringOrNil(p.GetQuery()),
		LastRefreshTime:   dcl.Int64OrNil(p.GetLastRefreshTime()),
		EnableRefresh:     dcl.Bool(p.GetEnableRefresh()),
		RefreshIntervalMs: dcl.Int64OrNil(p.GetRefreshIntervalMs()),
	}
	return obj
}

// ProtoToTableExternalDataConfiguration converts a TableExternalDataConfiguration object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfiguration(p *alphapb.BigqueryAlphaTableExternalDataConfiguration) *alpha.TableExternalDataConfiguration {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfiguration{
		Schema:                  ProtoToBigqueryAlphaTableExternalDataConfigurationSchema(p.GetSchema()),
		SourceFormat:            dcl.StringOrNil(p.GetSourceFormat()),
		MaxBadRecords:           dcl.Int64OrNil(p.GetMaxBadRecords()),
		Autodetect:              dcl.Bool(p.GetAutodetect()),
		IgnoreUnknownValues:     dcl.Bool(p.GetIgnoreUnknownValues()),
		Compression:             dcl.StringOrNil(p.GetCompression()),
		CsvOptions:              ProtoToBigqueryAlphaTableExternalDataConfigurationCsvOptions(p.GetCsvOptions()),
		BigtableOptions:         ProtoToBigqueryAlphaTableExternalDataConfigurationBigtableOptions(p.GetBigtableOptions()),
		GoogleSheetsOptions:     ProtoToBigqueryAlphaTableExternalDataConfigurationGoogleSheetsOptions(p.GetGoogleSheetsOptions()),
		HivePartitioningOptions: ProtoToBigqueryAlphaTableExternalDataConfigurationHivePartitioningOptions(p.GetHivePartitioningOptions()),
		ConnectionId:            dcl.StringOrNil(p.GetConnectionId()),
		ValueConversionModes:    ProtoToBigqueryAlphaTableExternalDataConfigurationValueConversionModes(p.GetValueConversionModes()),
		AvroOptions:             ProtoToBigqueryAlphaTableExternalDataConfigurationAvroOptions(p.GetAvroOptions()),
		JsonExtension:           ProtoToBigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum(p.GetJsonExtension()),
		ParquetOptions:          ProtoToBigqueryAlphaTableExternalDataConfigurationParquetOptions(p.GetParquetOptions()),
	}
	for _, r := range p.GetSourceUris() {
		obj.SourceUris = append(obj.SourceUris, r)
	}
	for _, r := range p.GetMaxBadRecordsAlternative() {
		obj.MaxBadRecordsAlternative = append(obj.MaxBadRecordsAlternative, r)
	}
	for _, r := range p.GetDecimalTargetTypes() {
		obj.DecimalTargetTypes = append(obj.DecimalTargetTypes, *ProtoToBigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationSchema converts a TableExternalDataConfigurationSchema object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationSchema(p *alphapb.BigqueryAlphaTableExternalDataConfigurationSchema) *alpha.TableExternalDataConfigurationSchema {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationSchema{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationCsvOptions converts a TableExternalDataConfigurationCsvOptions object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationCsvOptions(p *alphapb.BigqueryAlphaTableExternalDataConfigurationCsvOptions) *alpha.TableExternalDataConfigurationCsvOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationCsvOptions{
		FieldDelimiter:      dcl.StringOrNil(p.GetFieldDelimiter()),
		SkipLeadingRows:     dcl.StringOrNil(p.GetSkipLeadingRows()),
		Quote:               dcl.StringOrNil(p.GetQuote()),
		AllowQuotedNewlines: dcl.Bool(p.GetAllowQuotedNewlines()),
		AllowJaggedRows:     dcl.Bool(p.GetAllowJaggedRows()),
		Encoding:            dcl.StringOrNil(p.GetEncoding()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationBigtableOptions converts a TableExternalDataConfigurationBigtableOptions object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationBigtableOptions(p *alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptions) *alpha.TableExternalDataConfigurationBigtableOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationBigtableOptions{
		IgnoreUnspecifiedColumnFamilies: dcl.Bool(p.GetIgnoreUnspecifiedColumnFamilies()),
		ReadRowkeyAsString:              dcl.Bool(p.GetReadRowkeyAsString()),
	}
	for _, r := range p.GetColumnFamilies() {
		obj.ColumnFamilies = append(obj.ColumnFamilies, *ProtoToBigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamilies(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationBigtableOptionsColumnFamilies converts a TableExternalDataConfigurationBigtableOptionsColumnFamilies object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamilies(p *alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamilies) *alpha.TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationBigtableOptionsColumnFamilies{
		FamilyId:       dcl.StringOrNil(p.GetFamilyId()),
		Type:           dcl.StringOrNil(p.GetType()),
		Encoding:       dcl.StringOrNil(p.GetEncoding()),
		OnlyReadLatest: dcl.Bool(p.GetOnlyReadLatest()),
	}
	for _, r := range p.GetColumns() {
		obj.Columns = append(obj.Columns, *ProtoToBigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns converts a TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(p *alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) *alpha.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{
		QualifierEncoded: dcl.StringOrNil(p.GetQualifierEncoded()),
		QualifierString:  dcl.StringOrNil(p.GetQualifierString()),
		FieldName:        dcl.StringOrNil(p.GetFieldName()),
		Type:             dcl.StringOrNil(p.GetType()),
		Encoding:         dcl.StringOrNil(p.GetEncoding()),
		OnlyReadLatest:   dcl.Bool(p.GetOnlyReadLatest()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationGoogleSheetsOptions converts a TableExternalDataConfigurationGoogleSheetsOptions object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationGoogleSheetsOptions(p *alphapb.BigqueryAlphaTableExternalDataConfigurationGoogleSheetsOptions) *alpha.TableExternalDataConfigurationGoogleSheetsOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationGoogleSheetsOptions{
		SkipLeadingRows: dcl.StringOrNil(p.GetSkipLeadingRows()),
		Range:           dcl.StringOrNil(p.GetRange()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationHivePartitioningOptions converts a TableExternalDataConfigurationHivePartitioningOptions object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationHivePartitioningOptions(p *alphapb.BigqueryAlphaTableExternalDataConfigurationHivePartitioningOptions) *alpha.TableExternalDataConfigurationHivePartitioningOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationHivePartitioningOptions{
		Mode:                   dcl.StringOrNil(p.GetMode()),
		SourceUriPrefix:        dcl.StringOrNil(p.GetSourceUriPrefix()),
		RequirePartitionFilter: dcl.Bool(p.GetRequirePartitionFilter()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToTableExternalDataConfigurationValueConversionModes converts a TableExternalDataConfigurationValueConversionModes object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationValueConversionModes(p *alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModes) *alpha.TableExternalDataConfigurationValueConversionModes {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationValueConversionModes{
		TemporalTypesOutOfRangeConversionMode: ProtoToBigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(p.GetTemporalTypesOutOfRangeConversionMode()),
		NumericTypeOutOfRangeConversionMode:   ProtoToBigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(p.GetNumericTypeOutOfRangeConversionMode()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationAvroOptions converts a TableExternalDataConfigurationAvroOptions object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationAvroOptions(p *alphapb.BigqueryAlphaTableExternalDataConfigurationAvroOptions) *alpha.TableExternalDataConfigurationAvroOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationAvroOptions{
		UseAvroLogicalTypes: dcl.Bool(p.GetUseAvroLogicalTypes()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationParquetOptions converts a TableExternalDataConfigurationParquetOptions object from its proto representation.
func ProtoToBigqueryAlphaTableExternalDataConfigurationParquetOptions(p *alphapb.BigqueryAlphaTableExternalDataConfigurationParquetOptions) *alpha.TableExternalDataConfigurationParquetOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.TableExternalDataConfigurationParquetOptions{
		EnumAsString:        dcl.Bool(p.GetEnumAsString()),
		EnableListInference: dcl.Bool(p.GetEnableListInference()),
	}
	return obj
}

// ProtoToTableStreamingBuffer converts a TableStreamingBuffer object from its proto representation.
func ProtoToBigqueryAlphaTableStreamingBuffer(p *alphapb.BigqueryAlphaTableStreamingBuffer) *alpha.TableStreamingBuffer {
	if p == nil {
		return nil
	}
	obj := &alpha.TableStreamingBuffer{
		EstimatedBytes:  dcl.Int64OrNil(p.GetEstimatedBytes()),
		EstimatedRows:   dcl.Int64OrNil(p.GetEstimatedRows()),
		OldestEntryTime: dcl.Int64OrNil(p.GetOldestEntryTime()),
	}
	return obj
}

// ProtoToTableEncryptionConfiguration converts a TableEncryptionConfiguration object from its proto representation.
func ProtoToBigqueryAlphaTableEncryptionConfiguration(p *alphapb.BigqueryAlphaTableEncryptionConfiguration) *alpha.TableEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &alpha.TableEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToTableSnapshotDefinition converts a TableSnapshotDefinition object from its proto representation.
func ProtoToBigqueryAlphaTableSnapshotDefinition(p *alphapb.BigqueryAlphaTableSnapshotDefinition) *alpha.TableSnapshotDefinition {
	if p == nil {
		return nil
	}
	obj := &alpha.TableSnapshotDefinition{
		Table:        dcl.StringOrNil(p.GetTable()),
		Dataset:      dcl.StringOrNil(p.GetDataset()),
		Project:      dcl.StringOrNil(p.GetProject()),
		SnapshotTime: dcl.StringOrNil(p.GetSnapshotTime()),
	}
	return obj
}

// ProtoToTable converts a Table resource from its proto representation.
func ProtoToTable(p *alphapb.BigqueryAlphaTable) *alpha.Table {
	obj := &alpha.Table{
		Etag:                      dcl.StringOrNil(p.GetEtag()),
		Id:                        dcl.StringOrNil(p.GetId()),
		SelfLink:                  dcl.StringOrNil(p.GetSelfLink()),
		Name:                      dcl.StringOrNil(p.GetName()),
		Dataset:                   dcl.StringOrNil(p.GetDataset()),
		Project:                   dcl.StringOrNil(p.GetProject()),
		FriendlyName:              dcl.StringOrNil(p.GetFriendlyName()),
		Description:               dcl.StringOrNil(p.GetDescription()),
		Model:                     ProtoToBigqueryAlphaTableModel(p.GetModel()),
		Schema:                    ProtoToBigqueryAlphaTableSchema(p.GetSchema()),
		TimePartitioning:          ProtoToBigqueryAlphaTableTimePartitioning(p.GetTimePartitioning()),
		RangePartitioning:         ProtoToBigqueryAlphaTableRangePartitioning(p.GetRangePartitioning()),
		Clustering:                ProtoToBigqueryAlphaTableClustering(p.GetClustering()),
		RequirePartitionFilter:    dcl.Bool(p.GetRequirePartitionFilter()),
		NumBytes:                  dcl.StringOrNil(p.GetNumBytes()),
		NumPhysicalBytes:          dcl.StringOrNil(p.GetNumPhysicalBytes()),
		NumLongTermBytes:          dcl.StringOrNil(p.GetNumLongTermBytes()),
		NumRows:                   dcl.Int64OrNil(p.GetNumRows()),
		CreationTime:              dcl.Int64OrNil(p.GetCreationTime()),
		ExpirationTime:            dcl.Int64OrNil(p.GetExpirationTime()),
		LastModifiedTime:          dcl.Int64OrNil(p.GetLastModifiedTime()),
		Type:                      dcl.StringOrNil(p.GetType()),
		View:                      ProtoToBigqueryAlphaTableView(p.GetView()),
		MaterializedView:          ProtoToBigqueryAlphaTableMaterializedView(p.GetMaterializedView()),
		ExternalDataConfiguration: ProtoToBigqueryAlphaTableExternalDataConfiguration(p.GetExternalDataConfiguration()),
		Location:                  dcl.StringOrNil(p.GetLocation()),
		StreamingBuffer:           ProtoToBigqueryAlphaTableStreamingBuffer(p.GetStreamingBuffer()),
		EncryptionConfiguration:   ProtoToBigqueryAlphaTableEncryptionConfiguration(p.GetEncryptionConfiguration()),
		SnapshotDefinition:        ProtoToBigqueryAlphaTableSnapshotDefinition(p.GetSnapshotDefinition()),
		DefaultCollation:          dcl.StringOrNil(p.GetDefaultCollation()),
	}
	return obj
}

// TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto converts a TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum enum to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto(e *alpha.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum) alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	if e == nil {
		return alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum_value["TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(v)
	}
	return alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(0)
}

// TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto converts a TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum enum to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto(e *alpha.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum) alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	if e == nil {
		return alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum_value["TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(v)
	}
	return alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(0)
}

// TableExternalDataConfigurationDecimalTargetTypesEnumToProto converts a TableExternalDataConfigurationDecimalTargetTypesEnum enum to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnumToProto(e *alpha.TableExternalDataConfigurationDecimalTargetTypesEnum) alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum {
	if e == nil {
		return alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum_value["TableExternalDataConfigurationDecimalTargetTypesEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum(v)
	}
	return alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum(0)
}

// TableExternalDataConfigurationJsonExtensionEnumToProto converts a TableExternalDataConfigurationJsonExtensionEnum enum to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnumToProto(e *alpha.TableExternalDataConfigurationJsonExtensionEnum) alphapb.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum {
	if e == nil {
		return alphapb.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum(0)
	}
	if v, ok := alphapb.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum_value["TableExternalDataConfigurationJsonExtensionEnum"+string(*e)]; ok {
		return alphapb.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum(v)
	}
	return alphapb.BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnum(0)
}

// TableModelToProto converts a TableModel object to its proto representation.
func BigqueryAlphaTableModelToProto(o *alpha.TableModel) *alphapb.BigqueryAlphaTableModel {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableModel{}
	p.SetModelOptions(BigqueryAlphaTableModelModelOptionsToProto(o.ModelOptions))
	sTrainingRuns := make([]*alphapb.BigqueryAlphaTableModelTrainingRuns, len(o.TrainingRuns))
	for i, r := range o.TrainingRuns {
		sTrainingRuns[i] = BigqueryAlphaTableModelTrainingRunsToProto(&r)
	}
	p.SetTrainingRuns(sTrainingRuns)
	return p
}

// TableModelModelOptionsToProto converts a TableModelModelOptions object to its proto representation.
func BigqueryAlphaTableModelModelOptionsToProto(o *alpha.TableModelModelOptions) *alphapb.BigqueryAlphaTableModelModelOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableModelModelOptions{}
	p.SetModelType(dcl.ValueOrEmptyString(o.ModelType))
	p.SetLossType(dcl.ValueOrEmptyString(o.LossType))
	sLabels := make([]string, len(o.Labels))
	for i, r := range o.Labels {
		sLabels[i] = r
	}
	p.SetLabels(sLabels)
	return p
}

// TableModelTrainingRunsToProto converts a TableModelTrainingRuns object to its proto representation.
func BigqueryAlphaTableModelTrainingRunsToProto(o *alpha.TableModelTrainingRuns) *alphapb.BigqueryAlphaTableModelTrainingRuns {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableModelTrainingRuns{}
	p.SetState(dcl.ValueOrEmptyString(o.State))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetTrainingOptions(BigqueryAlphaTableModelTrainingRunsTrainingOptionsToProto(o.TrainingOptions))
	sIterationResults := make([]*alphapb.BigqueryAlphaTableModelTrainingRunsIterationResults, len(o.IterationResults))
	for i, r := range o.IterationResults {
		sIterationResults[i] = BigqueryAlphaTableModelTrainingRunsIterationResultsToProto(&r)
	}
	p.SetIterationResults(sIterationResults)
	return p
}

// TableModelTrainingRunsTrainingOptionsToProto converts a TableModelTrainingRunsTrainingOptions object to its proto representation.
func BigqueryAlphaTableModelTrainingRunsTrainingOptionsToProto(o *alpha.TableModelTrainingRunsTrainingOptions) *alphapb.BigqueryAlphaTableModelTrainingRunsTrainingOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableModelTrainingRunsTrainingOptions{}
	p.SetMaxIteration(dcl.ValueOrEmptyInt64(o.MaxIteration))
	p.SetLearnRate(dcl.ValueOrEmptyDouble(o.LearnRate))
	p.SetL1Reg(dcl.ValueOrEmptyDouble(o.L1Reg))
	p.SetL2Reg(dcl.ValueOrEmptyDouble(o.L2Reg))
	p.SetMinRelProgress(dcl.ValueOrEmptyDouble(o.MinRelProgress))
	p.SetWarmStart(dcl.ValueOrEmptyBool(o.WarmStart))
	p.SetEarlyStop(dcl.ValueOrEmptyBool(o.EarlyStop))
	p.SetLearnRateStrategy(dcl.ValueOrEmptyString(o.LearnRateStrategy))
	p.SetLineSearchInitLearnRate(dcl.ValueOrEmptyDouble(o.LineSearchInitLearnRate))
	return p
}

// TableModelTrainingRunsIterationResultsToProto converts a TableModelTrainingRunsIterationResults object to its proto representation.
func BigqueryAlphaTableModelTrainingRunsIterationResultsToProto(o *alpha.TableModelTrainingRunsIterationResults) *alphapb.BigqueryAlphaTableModelTrainingRunsIterationResults {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableModelTrainingRunsIterationResults{}
	p.SetIndex(dcl.ValueOrEmptyInt64(o.Index))
	p.SetLearnRate(dcl.ValueOrEmptyDouble(o.LearnRate))
	p.SetTrainingLoss(dcl.ValueOrEmptyDouble(o.TrainingLoss))
	p.SetEvalLoss(dcl.ValueOrEmptyDouble(o.EvalLoss))
	p.SetDurationMs(dcl.ValueOrEmptyString(o.DurationMs))
	return p
}

// TableSchemaToProto converts a TableSchema object to its proto representation.
func BigqueryAlphaTableSchemaToProto(o *alpha.TableSchema) *alphapb.BigqueryAlphaTableSchema {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableSchema{}
	sFields := make([]*alphapb.BigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// TableGooglecloudbigqueryv2TablefieldschemaToProto converts a TableGooglecloudbigqueryv2Tablefieldschema object to its proto representation.
func BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaToProto(o *alpha.TableGooglecloudbigqueryv2Tablefieldschema) *alphapb.BigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetMode(dcl.ValueOrEmptyString(o.Mode))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetCategories(BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaCategoriesToProto(o.Categories))
	p.SetPolicyTags(BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto(o.PolicyTags))
	p.SetMaxLength(dcl.ValueOrEmptyInt64(o.MaxLength))
	p.SetPrecision(dcl.ValueOrEmptyInt64(o.Precision))
	p.SetScale(dcl.ValueOrEmptyInt64(o.Scale))
	p.SetCollation(dcl.ValueOrEmptyString(o.Collation))
	p.SetDefaultValueExpression(dcl.ValueOrEmptyString(o.DefaultValueExpression))
	sFields := make([]*alphapb.BigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
	}
	p.SetFields(sFields)
	sNameAlternative := make([]string, len(o.NameAlternative))
	for i, r := range o.NameAlternative {
		sNameAlternative[i] = r
	}
	p.SetNameAlternative(sNameAlternative)
	return p
}

// TableGooglecloudbigqueryv2TablefieldschemaCategoriesToProto converts a TableGooglecloudbigqueryv2TablefieldschemaCategories object to its proto representation.
func BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaCategoriesToProto(o *alpha.TableGooglecloudbigqueryv2TablefieldschemaCategories) *alphapb.BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaCategories {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaCategories{}
	sNames := make([]string, len(o.Names))
	for i, r := range o.Names {
		sNames[i] = r
	}
	p.SetNames(sNames)
	return p
}

// TableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto converts a TableGooglecloudbigqueryv2TablefieldschemaPolicyTags object to its proto representation.
func BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto(o *alpha.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) *alphapb.BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	sNames := make([]string, len(o.Names))
	for i, r := range o.Names {
		sNames[i] = r
	}
	p.SetNames(sNames)
	return p
}

// TableTimePartitioningToProto converts a TableTimePartitioning object to its proto representation.
func BigqueryAlphaTableTimePartitioningToProto(o *alpha.TableTimePartitioning) *alphapb.BigqueryAlphaTableTimePartitioning {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableTimePartitioning{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetExpirationMs(dcl.ValueOrEmptyString(o.ExpirationMs))
	p.SetField(dcl.ValueOrEmptyString(o.Field))
	return p
}

// TableRangePartitioningToProto converts a TableRangePartitioning object to its proto representation.
func BigqueryAlphaTableRangePartitioningToProto(o *alpha.TableRangePartitioning) *alphapb.BigqueryAlphaTableRangePartitioning {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableRangePartitioning{}
	p.SetField(dcl.ValueOrEmptyString(o.Field))
	p.SetRange(BigqueryAlphaTableRangePartitioningRangeToProto(o.Range))
	return p
}

// TableRangePartitioningRangeToProto converts a TableRangePartitioningRange object to its proto representation.
func BigqueryAlphaTableRangePartitioningRangeToProto(o *alpha.TableRangePartitioningRange) *alphapb.BigqueryAlphaTableRangePartitioningRange {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableRangePartitioningRange{}
	p.SetStart(dcl.ValueOrEmptyString(o.Start))
	p.SetEnd(dcl.ValueOrEmptyString(o.End))
	p.SetInterval(dcl.ValueOrEmptyString(o.Interval))
	return p
}

// TableClusteringToProto converts a TableClustering object to its proto representation.
func BigqueryAlphaTableClusteringToProto(o *alpha.TableClustering) *alphapb.BigqueryAlphaTableClustering {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableClustering{}
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// TableViewToProto converts a TableView object to its proto representation.
func BigqueryAlphaTableViewToProto(o *alpha.TableView) *alphapb.BigqueryAlphaTableView {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableView{}
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	p.SetUseLegacySql(dcl.ValueOrEmptyBool(o.UseLegacySql))
	p.SetUseExplicitColumnNames(dcl.ValueOrEmptyBool(o.UseExplicitColumnNames))
	sUserDefinedFunctionResources := make([]*alphapb.BigqueryAlphaTableViewUserDefinedFunctionResources, len(o.UserDefinedFunctionResources))
	for i, r := range o.UserDefinedFunctionResources {
		sUserDefinedFunctionResources[i] = BigqueryAlphaTableViewUserDefinedFunctionResourcesToProto(&r)
	}
	p.SetUserDefinedFunctionResources(sUserDefinedFunctionResources)
	return p
}

// TableViewUserDefinedFunctionResourcesToProto converts a TableViewUserDefinedFunctionResources object to its proto representation.
func BigqueryAlphaTableViewUserDefinedFunctionResourcesToProto(o *alpha.TableViewUserDefinedFunctionResources) *alphapb.BigqueryAlphaTableViewUserDefinedFunctionResources {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableViewUserDefinedFunctionResources{}
	p.SetResourceUri(dcl.ValueOrEmptyString(o.ResourceUri))
	p.SetInlineCode(dcl.ValueOrEmptyString(o.InlineCode))
	sInlineCodeAlternative := make([]string, len(o.InlineCodeAlternative))
	for i, r := range o.InlineCodeAlternative {
		sInlineCodeAlternative[i] = r
	}
	p.SetInlineCodeAlternative(sInlineCodeAlternative)
	return p
}

// TableMaterializedViewToProto converts a TableMaterializedView object to its proto representation.
func BigqueryAlphaTableMaterializedViewToProto(o *alpha.TableMaterializedView) *alphapb.BigqueryAlphaTableMaterializedView {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableMaterializedView{}
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	p.SetLastRefreshTime(dcl.ValueOrEmptyInt64(o.LastRefreshTime))
	p.SetEnableRefresh(dcl.ValueOrEmptyBool(o.EnableRefresh))
	p.SetRefreshIntervalMs(dcl.ValueOrEmptyInt64(o.RefreshIntervalMs))
	return p
}

// TableExternalDataConfigurationToProto converts a TableExternalDataConfiguration object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationToProto(o *alpha.TableExternalDataConfiguration) *alphapb.BigqueryAlphaTableExternalDataConfiguration {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfiguration{}
	p.SetSchema(BigqueryAlphaTableExternalDataConfigurationSchemaToProto(o.Schema))
	p.SetSourceFormat(dcl.ValueOrEmptyString(o.SourceFormat))
	p.SetMaxBadRecords(dcl.ValueOrEmptyInt64(o.MaxBadRecords))
	p.SetAutodetect(dcl.ValueOrEmptyBool(o.Autodetect))
	p.SetIgnoreUnknownValues(dcl.ValueOrEmptyBool(o.IgnoreUnknownValues))
	p.SetCompression(dcl.ValueOrEmptyString(o.Compression))
	p.SetCsvOptions(BigqueryAlphaTableExternalDataConfigurationCsvOptionsToProto(o.CsvOptions))
	p.SetBigtableOptions(BigqueryAlphaTableExternalDataConfigurationBigtableOptionsToProto(o.BigtableOptions))
	p.SetGoogleSheetsOptions(BigqueryAlphaTableExternalDataConfigurationGoogleSheetsOptionsToProto(o.GoogleSheetsOptions))
	p.SetHivePartitioningOptions(BigqueryAlphaTableExternalDataConfigurationHivePartitioningOptionsToProto(o.HivePartitioningOptions))
	p.SetConnectionId(dcl.ValueOrEmptyString(o.ConnectionId))
	p.SetValueConversionModes(BigqueryAlphaTableExternalDataConfigurationValueConversionModesToProto(o.ValueConversionModes))
	p.SetAvroOptions(BigqueryAlphaTableExternalDataConfigurationAvroOptionsToProto(o.AvroOptions))
	p.SetJsonExtension(BigqueryAlphaTableExternalDataConfigurationJsonExtensionEnumToProto(o.JsonExtension))
	p.SetParquetOptions(BigqueryAlphaTableExternalDataConfigurationParquetOptionsToProto(o.ParquetOptions))
	sSourceUris := make([]string, len(o.SourceUris))
	for i, r := range o.SourceUris {
		sSourceUris[i] = r
	}
	p.SetSourceUris(sSourceUris)
	sMaxBadRecordsAlternative := make([]int64, len(o.MaxBadRecordsAlternative))
	for i, r := range o.MaxBadRecordsAlternative {
		sMaxBadRecordsAlternative[i] = r
	}
	p.SetMaxBadRecordsAlternative(sMaxBadRecordsAlternative)
	sDecimalTargetTypes := make([]alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum, len(o.DecimalTargetTypes))
	for i, r := range o.DecimalTargetTypes {
		sDecimalTargetTypes[i] = alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum(alphapb.BigqueryAlphaTableExternalDataConfigurationDecimalTargetTypesEnum_value[string(r)])
	}
	p.SetDecimalTargetTypes(sDecimalTargetTypes)
	return p
}

// TableExternalDataConfigurationSchemaToProto converts a TableExternalDataConfigurationSchema object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationSchemaToProto(o *alpha.TableExternalDataConfigurationSchema) *alphapb.BigqueryAlphaTableExternalDataConfigurationSchema {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationSchema{}
	sFields := make([]*alphapb.BigqueryAlphaTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryAlphaTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// TableExternalDataConfigurationCsvOptionsToProto converts a TableExternalDataConfigurationCsvOptions object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationCsvOptionsToProto(o *alpha.TableExternalDataConfigurationCsvOptions) *alphapb.BigqueryAlphaTableExternalDataConfigurationCsvOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationCsvOptions{}
	p.SetFieldDelimiter(dcl.ValueOrEmptyString(o.FieldDelimiter))
	p.SetSkipLeadingRows(dcl.ValueOrEmptyString(o.SkipLeadingRows))
	p.SetQuote(dcl.ValueOrEmptyString(o.Quote))
	p.SetAllowQuotedNewlines(dcl.ValueOrEmptyBool(o.AllowQuotedNewlines))
	p.SetAllowJaggedRows(dcl.ValueOrEmptyBool(o.AllowJaggedRows))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	return p
}

// TableExternalDataConfigurationBigtableOptionsToProto converts a TableExternalDataConfigurationBigtableOptions object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationBigtableOptionsToProto(o *alpha.TableExternalDataConfigurationBigtableOptions) *alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptions{}
	p.SetIgnoreUnspecifiedColumnFamilies(dcl.ValueOrEmptyBool(o.IgnoreUnspecifiedColumnFamilies))
	p.SetReadRowkeyAsString(dcl.ValueOrEmptyBool(o.ReadRowkeyAsString))
	sColumnFamilies := make([]*alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamilies, len(o.ColumnFamilies))
	for i, r := range o.ColumnFamilies {
		sColumnFamilies[i] = BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto(&r)
	}
	p.SetColumnFamilies(sColumnFamilies)
	return p
}

// TableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto converts a TableExternalDataConfigurationBigtableOptionsColumnFamilies object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto(o *alpha.TableExternalDataConfigurationBigtableOptionsColumnFamilies) *alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamilies{}
	p.SetFamilyId(dcl.ValueOrEmptyString(o.FamilyId))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetOnlyReadLatest(dcl.ValueOrEmptyBool(o.OnlyReadLatest))
	sColumns := make([]*alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, len(o.Columns))
	for i, r := range o.Columns {
		sColumns[i] = BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto(&r)
	}
	p.SetColumns(sColumns)
	return p
}

// TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto converts a TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto(o *alpha.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) *alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}
	p.SetQualifierEncoded(dcl.ValueOrEmptyString(o.QualifierEncoded))
	p.SetQualifierString(dcl.ValueOrEmptyString(o.QualifierString))
	p.SetFieldName(dcl.ValueOrEmptyString(o.FieldName))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetOnlyReadLatest(dcl.ValueOrEmptyBool(o.OnlyReadLatest))
	return p
}

// TableExternalDataConfigurationGoogleSheetsOptionsToProto converts a TableExternalDataConfigurationGoogleSheetsOptions object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationGoogleSheetsOptionsToProto(o *alpha.TableExternalDataConfigurationGoogleSheetsOptions) *alphapb.BigqueryAlphaTableExternalDataConfigurationGoogleSheetsOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationGoogleSheetsOptions{}
	p.SetSkipLeadingRows(dcl.ValueOrEmptyString(o.SkipLeadingRows))
	p.SetRange(dcl.ValueOrEmptyString(o.Range))
	return p
}

// TableExternalDataConfigurationHivePartitioningOptionsToProto converts a TableExternalDataConfigurationHivePartitioningOptions object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationHivePartitioningOptionsToProto(o *alpha.TableExternalDataConfigurationHivePartitioningOptions) *alphapb.BigqueryAlphaTableExternalDataConfigurationHivePartitioningOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationHivePartitioningOptions{}
	p.SetMode(dcl.ValueOrEmptyString(o.Mode))
	p.SetSourceUriPrefix(dcl.ValueOrEmptyString(o.SourceUriPrefix))
	p.SetRequirePartitionFilter(dcl.ValueOrEmptyBool(o.RequirePartitionFilter))
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// TableExternalDataConfigurationValueConversionModesToProto converts a TableExternalDataConfigurationValueConversionModes object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationValueConversionModesToProto(o *alpha.TableExternalDataConfigurationValueConversionModes) *alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModes {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationValueConversionModes{}
	p.SetTemporalTypesOutOfRangeConversionMode(BigqueryAlphaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto(o.TemporalTypesOutOfRangeConversionMode))
	p.SetNumericTypeOutOfRangeConversionMode(BigqueryAlphaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto(o.NumericTypeOutOfRangeConversionMode))
	return p
}

// TableExternalDataConfigurationAvroOptionsToProto converts a TableExternalDataConfigurationAvroOptions object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationAvroOptionsToProto(o *alpha.TableExternalDataConfigurationAvroOptions) *alphapb.BigqueryAlphaTableExternalDataConfigurationAvroOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationAvroOptions{}
	p.SetUseAvroLogicalTypes(dcl.ValueOrEmptyBool(o.UseAvroLogicalTypes))
	return p
}

// TableExternalDataConfigurationParquetOptionsToProto converts a TableExternalDataConfigurationParquetOptions object to its proto representation.
func BigqueryAlphaTableExternalDataConfigurationParquetOptionsToProto(o *alpha.TableExternalDataConfigurationParquetOptions) *alphapb.BigqueryAlphaTableExternalDataConfigurationParquetOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableExternalDataConfigurationParquetOptions{}
	p.SetEnumAsString(dcl.ValueOrEmptyBool(o.EnumAsString))
	p.SetEnableListInference(dcl.ValueOrEmptyBool(o.EnableListInference))
	return p
}

// TableStreamingBufferToProto converts a TableStreamingBuffer object to its proto representation.
func BigqueryAlphaTableStreamingBufferToProto(o *alpha.TableStreamingBuffer) *alphapb.BigqueryAlphaTableStreamingBuffer {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableStreamingBuffer{}
	p.SetEstimatedBytes(dcl.ValueOrEmptyInt64(o.EstimatedBytes))
	p.SetEstimatedRows(dcl.ValueOrEmptyInt64(o.EstimatedRows))
	p.SetOldestEntryTime(dcl.ValueOrEmptyInt64(o.OldestEntryTime))
	return p
}

// TableEncryptionConfigurationToProto converts a TableEncryptionConfiguration object to its proto representation.
func BigqueryAlphaTableEncryptionConfigurationToProto(o *alpha.TableEncryptionConfiguration) *alphapb.BigqueryAlphaTableEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableEncryptionConfiguration{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// TableSnapshotDefinitionToProto converts a TableSnapshotDefinition object to its proto representation.
func BigqueryAlphaTableSnapshotDefinitionToProto(o *alpha.TableSnapshotDefinition) *alphapb.BigqueryAlphaTableSnapshotDefinition {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaTableSnapshotDefinition{}
	p.SetTable(dcl.ValueOrEmptyString(o.Table))
	p.SetDataset(dcl.ValueOrEmptyString(o.Dataset))
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetSnapshotTime(dcl.ValueOrEmptyString(o.SnapshotTime))
	return p
}

// TableToProto converts a Table resource to its proto representation.
func TableToProto(resource *alpha.Table) *alphapb.BigqueryAlphaTable {
	p := &alphapb.BigqueryAlphaTable{}
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetId(dcl.ValueOrEmptyString(resource.Id))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetFriendlyName(dcl.ValueOrEmptyString(resource.FriendlyName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetModel(BigqueryAlphaTableModelToProto(resource.Model))
	p.SetSchema(BigqueryAlphaTableSchemaToProto(resource.Schema))
	p.SetTimePartitioning(BigqueryAlphaTableTimePartitioningToProto(resource.TimePartitioning))
	p.SetRangePartitioning(BigqueryAlphaTableRangePartitioningToProto(resource.RangePartitioning))
	p.SetClustering(BigqueryAlphaTableClusteringToProto(resource.Clustering))
	p.SetRequirePartitionFilter(dcl.ValueOrEmptyBool(resource.RequirePartitionFilter))
	p.SetNumBytes(dcl.ValueOrEmptyString(resource.NumBytes))
	p.SetNumPhysicalBytes(dcl.ValueOrEmptyString(resource.NumPhysicalBytes))
	p.SetNumLongTermBytes(dcl.ValueOrEmptyString(resource.NumLongTermBytes))
	p.SetNumRows(dcl.ValueOrEmptyInt64(resource.NumRows))
	p.SetCreationTime(dcl.ValueOrEmptyInt64(resource.CreationTime))
	p.SetExpirationTime(dcl.ValueOrEmptyInt64(resource.ExpirationTime))
	p.SetLastModifiedTime(dcl.ValueOrEmptyInt64(resource.LastModifiedTime))
	p.SetType(dcl.ValueOrEmptyString(resource.Type))
	p.SetView(BigqueryAlphaTableViewToProto(resource.View))
	p.SetMaterializedView(BigqueryAlphaTableMaterializedViewToProto(resource.MaterializedView))
	p.SetExternalDataConfiguration(BigqueryAlphaTableExternalDataConfigurationToProto(resource.ExternalDataConfiguration))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetStreamingBuffer(BigqueryAlphaTableStreamingBufferToProto(resource.StreamingBuffer))
	p.SetEncryptionConfiguration(BigqueryAlphaTableEncryptionConfigurationToProto(resource.EncryptionConfiguration))
	p.SetSnapshotDefinition(BigqueryAlphaTableSnapshotDefinitionToProto(resource.SnapshotDefinition))
	p.SetDefaultCollation(dcl.ValueOrEmptyString(resource.DefaultCollation))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyTable handles the gRPC request by passing it to the underlying Table Apply() method.
func (s *TableServer) applyTable(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBigqueryAlphaTableRequest) (*alphapb.BigqueryAlphaTable, error) {
	p := ProtoToTable(request.GetResource())
	res, err := c.ApplyTable(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TableToProto(res)
	return r, nil
}

// applyBigqueryAlphaTable handles the gRPC request by passing it to the underlying Table Apply() method.
func (s *TableServer) ApplyBigqueryAlphaTable(ctx context.Context, request *alphapb.ApplyBigqueryAlphaTableRequest) (*alphapb.BigqueryAlphaTable, error) {
	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTable(ctx, cl, request)
}

// DeleteTable handles the gRPC request by passing it to the underlying Table Delete() method.
func (s *TableServer) DeleteBigqueryAlphaTable(ctx context.Context, request *alphapb.DeleteBigqueryAlphaTableRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTable(ctx, ProtoToTable(request.GetResource()))

}

// ListBigqueryAlphaTable handles the gRPC request by passing it to the underlying TableList() method.
func (s *TableServer) ListBigqueryAlphaTable(ctx context.Context, request *alphapb.ListBigqueryAlphaTableRequest) (*alphapb.ListBigqueryAlphaTableResponse, error) {
	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTable(ctx, request.GetProject(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.BigqueryAlphaTable
	for _, r := range resources.Items {
		rp := TableToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListBigqueryAlphaTableResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTable(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
