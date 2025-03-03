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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/beta/bigquery_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/beta"
)

// TableServer implements the gRPC interface for Table.
type TableServer struct{}

// ProtoToTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum converts a TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum enum from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(e betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum) *beta.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum_name[int32(e)]; ok {
		e := beta.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(n[len("BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum converts a TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum enum from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(e betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum) *beta.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum_name[int32(e)]; ok {
		e := beta.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(n[len("BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationDecimalTargetTypesEnum converts a TableExternalDataConfigurationDecimalTargetTypesEnum enum from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum(e betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum) *beta.TableExternalDataConfigurationDecimalTargetTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum_name[int32(e)]; ok {
		e := beta.TableExternalDataConfigurationDecimalTargetTypesEnum(n[len("BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationJsonExtensionEnum converts a TableExternalDataConfigurationJsonExtensionEnum enum from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationJsonExtensionEnum(e betapb.BigqueryBetaTableExternalDataConfigurationJsonExtensionEnum) *beta.TableExternalDataConfigurationJsonExtensionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryBetaTableExternalDataConfigurationJsonExtensionEnum_name[int32(e)]; ok {
		e := beta.TableExternalDataConfigurationJsonExtensionEnum(n[len("BigqueryBetaTableExternalDataConfigurationJsonExtensionEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableModel converts a TableModel object from its proto representation.
func ProtoToBigqueryBetaTableModel(p *betapb.BigqueryBetaTableModel) *beta.TableModel {
	if p == nil {
		return nil
	}
	obj := &beta.TableModel{
		ModelOptions: ProtoToBigqueryBetaTableModelModelOptions(p.GetModelOptions()),
	}
	for _, r := range p.GetTrainingRuns() {
		obj.TrainingRuns = append(obj.TrainingRuns, *ProtoToBigqueryBetaTableModelTrainingRuns(r))
	}
	return obj
}

// ProtoToTableModelModelOptions converts a TableModelModelOptions object from its proto representation.
func ProtoToBigqueryBetaTableModelModelOptions(p *betapb.BigqueryBetaTableModelModelOptions) *beta.TableModelModelOptions {
	if p == nil {
		return nil
	}
	obj := &beta.TableModelModelOptions{
		ModelType: dcl.StringOrNil(p.GetModelType()),
		LossType:  dcl.StringOrNil(p.GetLossType()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, r)
	}
	return obj
}

// ProtoToTableModelTrainingRuns converts a TableModelTrainingRuns object from its proto representation.
func ProtoToBigqueryBetaTableModelTrainingRuns(p *betapb.BigqueryBetaTableModelTrainingRuns) *beta.TableModelTrainingRuns {
	if p == nil {
		return nil
	}
	obj := &beta.TableModelTrainingRuns{
		State:           dcl.StringOrNil(p.GetState()),
		StartTime:       dcl.StringOrNil(p.GetStartTime()),
		TrainingOptions: ProtoToBigqueryBetaTableModelTrainingRunsTrainingOptions(p.GetTrainingOptions()),
	}
	for _, r := range p.GetIterationResults() {
		obj.IterationResults = append(obj.IterationResults, *ProtoToBigqueryBetaTableModelTrainingRunsIterationResults(r))
	}
	return obj
}

// ProtoToTableModelTrainingRunsTrainingOptions converts a TableModelTrainingRunsTrainingOptions object from its proto representation.
func ProtoToBigqueryBetaTableModelTrainingRunsTrainingOptions(p *betapb.BigqueryBetaTableModelTrainingRunsTrainingOptions) *beta.TableModelTrainingRunsTrainingOptions {
	if p == nil {
		return nil
	}
	obj := &beta.TableModelTrainingRunsTrainingOptions{
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
func ProtoToBigqueryBetaTableModelTrainingRunsIterationResults(p *betapb.BigqueryBetaTableModelTrainingRunsIterationResults) *beta.TableModelTrainingRunsIterationResults {
	if p == nil {
		return nil
	}
	obj := &beta.TableModelTrainingRunsIterationResults{
		Index:        dcl.Int64OrNil(p.GetIndex()),
		LearnRate:    dcl.Float64OrNil(p.GetLearnRate()),
		TrainingLoss: dcl.Float64OrNil(p.GetTrainingLoss()),
		EvalLoss:     dcl.Float64OrNil(p.GetEvalLoss()),
		DurationMs:   dcl.StringOrNil(p.GetDurationMs()),
	}
	return obj
}

// ProtoToTableSchema converts a TableSchema object from its proto representation.
func ProtoToBigqueryBetaTableSchema(p *betapb.BigqueryBetaTableSchema) *beta.TableSchema {
	if p == nil {
		return nil
	}
	obj := &beta.TableSchema{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2Tablefieldschema converts a TableGooglecloudbigqueryv2Tablefieldschema object from its proto representation.
func ProtoToBigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema(p *betapb.BigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema) *beta.TableGooglecloudbigqueryv2Tablefieldschema {
	if p == nil {
		return nil
	}
	obj := &beta.TableGooglecloudbigqueryv2Tablefieldschema{
		Name:                   dcl.StringOrNil(p.GetName()),
		Type:                   dcl.StringOrNil(p.GetType()),
		Mode:                   dcl.StringOrNil(p.GetMode()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		Categories:             ProtoToBigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaCategories(p.GetCategories()),
		PolicyTags:             ProtoToBigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(p.GetPolicyTags()),
		MaxLength:              dcl.Int64OrNil(p.GetMaxLength()),
		Precision:              dcl.Int64OrNil(p.GetPrecision()),
		Scale:                  dcl.Int64OrNil(p.GetScale()),
		Collation:              dcl.StringOrNil(p.GetCollation()),
		DefaultValueExpression: dcl.StringOrNil(p.GetDefaultValueExpression()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	for _, r := range p.GetNameAlternative() {
		obj.NameAlternative = append(obj.NameAlternative, r)
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2TablefieldschemaCategories converts a TableGooglecloudbigqueryv2TablefieldschemaCategories object from its proto representation.
func ProtoToBigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaCategories(p *betapb.BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaCategories) *beta.TableGooglecloudbigqueryv2TablefieldschemaCategories {
	if p == nil {
		return nil
	}
	obj := &beta.TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	for _, r := range p.GetNames() {
		obj.Names = append(obj.Names, r)
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2TablefieldschemaPolicyTags converts a TableGooglecloudbigqueryv2TablefieldschemaPolicyTags object from its proto representation.
func ProtoToBigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(p *betapb.BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags) *beta.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if p == nil {
		return nil
	}
	obj := &beta.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	for _, r := range p.GetNames() {
		obj.Names = append(obj.Names, r)
	}
	return obj
}

// ProtoToTableTimePartitioning converts a TableTimePartitioning object from its proto representation.
func ProtoToBigqueryBetaTableTimePartitioning(p *betapb.BigqueryBetaTableTimePartitioning) *beta.TableTimePartitioning {
	if p == nil {
		return nil
	}
	obj := &beta.TableTimePartitioning{
		Type:         dcl.StringOrNil(p.GetType()),
		ExpirationMs: dcl.StringOrNil(p.GetExpirationMs()),
		Field:        dcl.StringOrNil(p.GetField()),
	}
	return obj
}

// ProtoToTableRangePartitioning converts a TableRangePartitioning object from its proto representation.
func ProtoToBigqueryBetaTableRangePartitioning(p *betapb.BigqueryBetaTableRangePartitioning) *beta.TableRangePartitioning {
	if p == nil {
		return nil
	}
	obj := &beta.TableRangePartitioning{
		Field: dcl.StringOrNil(p.GetField()),
		Range: ProtoToBigqueryBetaTableRangePartitioningRange(p.GetRange()),
	}
	return obj
}

// ProtoToTableRangePartitioningRange converts a TableRangePartitioningRange object from its proto representation.
func ProtoToBigqueryBetaTableRangePartitioningRange(p *betapb.BigqueryBetaTableRangePartitioningRange) *beta.TableRangePartitioningRange {
	if p == nil {
		return nil
	}
	obj := &beta.TableRangePartitioningRange{
		Start:    dcl.StringOrNil(p.GetStart()),
		End:      dcl.StringOrNil(p.GetEnd()),
		Interval: dcl.StringOrNil(p.GetInterval()),
	}
	return obj
}

// ProtoToTableClustering converts a TableClustering object from its proto representation.
func ProtoToBigqueryBetaTableClustering(p *betapb.BigqueryBetaTableClustering) *beta.TableClustering {
	if p == nil {
		return nil
	}
	obj := &beta.TableClustering{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToTableView converts a TableView object from its proto representation.
func ProtoToBigqueryBetaTableView(p *betapb.BigqueryBetaTableView) *beta.TableView {
	if p == nil {
		return nil
	}
	obj := &beta.TableView{
		Query:                  dcl.StringOrNil(p.GetQuery()),
		UseLegacySql:           dcl.Bool(p.GetUseLegacySql()),
		UseExplicitColumnNames: dcl.Bool(p.GetUseExplicitColumnNames()),
	}
	for _, r := range p.GetUserDefinedFunctionResources() {
		obj.UserDefinedFunctionResources = append(obj.UserDefinedFunctionResources, *ProtoToBigqueryBetaTableViewUserDefinedFunctionResources(r))
	}
	return obj
}

// ProtoToTableViewUserDefinedFunctionResources converts a TableViewUserDefinedFunctionResources object from its proto representation.
func ProtoToBigqueryBetaTableViewUserDefinedFunctionResources(p *betapb.BigqueryBetaTableViewUserDefinedFunctionResources) *beta.TableViewUserDefinedFunctionResources {
	if p == nil {
		return nil
	}
	obj := &beta.TableViewUserDefinedFunctionResources{
		ResourceUri: dcl.StringOrNil(p.GetResourceUri()),
		InlineCode:  dcl.StringOrNil(p.GetInlineCode()),
	}
	for _, r := range p.GetInlineCodeAlternative() {
		obj.InlineCodeAlternative = append(obj.InlineCodeAlternative, r)
	}
	return obj
}

// ProtoToTableMaterializedView converts a TableMaterializedView object from its proto representation.
func ProtoToBigqueryBetaTableMaterializedView(p *betapb.BigqueryBetaTableMaterializedView) *beta.TableMaterializedView {
	if p == nil {
		return nil
	}
	obj := &beta.TableMaterializedView{
		Query:             dcl.StringOrNil(p.GetQuery()),
		LastRefreshTime:   dcl.Int64OrNil(p.GetLastRefreshTime()),
		EnableRefresh:     dcl.Bool(p.GetEnableRefresh()),
		RefreshIntervalMs: dcl.Int64OrNil(p.GetRefreshIntervalMs()),
	}
	return obj
}

// ProtoToTableExternalDataConfiguration converts a TableExternalDataConfiguration object from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfiguration(p *betapb.BigqueryBetaTableExternalDataConfiguration) *beta.TableExternalDataConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfiguration{
		Schema:                  ProtoToBigqueryBetaTableExternalDataConfigurationSchema(p.GetSchema()),
		SourceFormat:            dcl.StringOrNil(p.GetSourceFormat()),
		MaxBadRecords:           dcl.Int64OrNil(p.GetMaxBadRecords()),
		Autodetect:              dcl.Bool(p.GetAutodetect()),
		IgnoreUnknownValues:     dcl.Bool(p.GetIgnoreUnknownValues()),
		Compression:             dcl.StringOrNil(p.GetCompression()),
		CsvOptions:              ProtoToBigqueryBetaTableExternalDataConfigurationCsvOptions(p.GetCsvOptions()),
		BigtableOptions:         ProtoToBigqueryBetaTableExternalDataConfigurationBigtableOptions(p.GetBigtableOptions()),
		GoogleSheetsOptions:     ProtoToBigqueryBetaTableExternalDataConfigurationGoogleSheetsOptions(p.GetGoogleSheetsOptions()),
		HivePartitioningOptions: ProtoToBigqueryBetaTableExternalDataConfigurationHivePartitioningOptions(p.GetHivePartitioningOptions()),
		ConnectionId:            dcl.StringOrNil(p.GetConnectionId()),
		ValueConversionModes:    ProtoToBigqueryBetaTableExternalDataConfigurationValueConversionModes(p.GetValueConversionModes()),
		AvroOptions:             ProtoToBigqueryBetaTableExternalDataConfigurationAvroOptions(p.GetAvroOptions()),
		JsonExtension:           ProtoToBigqueryBetaTableExternalDataConfigurationJsonExtensionEnum(p.GetJsonExtension()),
		ParquetOptions:          ProtoToBigqueryBetaTableExternalDataConfigurationParquetOptions(p.GetParquetOptions()),
	}
	for _, r := range p.GetSourceUris() {
		obj.SourceUris = append(obj.SourceUris, r)
	}
	for _, r := range p.GetMaxBadRecordsAlternative() {
		obj.MaxBadRecordsAlternative = append(obj.MaxBadRecordsAlternative, r)
	}
	for _, r := range p.GetDecimalTargetTypes() {
		obj.DecimalTargetTypes = append(obj.DecimalTargetTypes, *ProtoToBigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationSchema converts a TableExternalDataConfigurationSchema object from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationSchema(p *betapb.BigqueryBetaTableExternalDataConfigurationSchema) *beta.TableExternalDataConfigurationSchema {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationSchema{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationCsvOptions converts a TableExternalDataConfigurationCsvOptions object from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationCsvOptions(p *betapb.BigqueryBetaTableExternalDataConfigurationCsvOptions) *beta.TableExternalDataConfigurationCsvOptions {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationCsvOptions{
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
func ProtoToBigqueryBetaTableExternalDataConfigurationBigtableOptions(p *betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptions) *beta.TableExternalDataConfigurationBigtableOptions {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationBigtableOptions{
		IgnoreUnspecifiedColumnFamilies: dcl.Bool(p.GetIgnoreUnspecifiedColumnFamilies()),
		ReadRowkeyAsString:              dcl.Bool(p.GetReadRowkeyAsString()),
	}
	for _, r := range p.GetColumnFamilies() {
		obj.ColumnFamilies = append(obj.ColumnFamilies, *ProtoToBigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamilies(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationBigtableOptionsColumnFamilies converts a TableExternalDataConfigurationBigtableOptionsColumnFamilies object from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamilies(p *betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamilies) *beta.TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationBigtableOptionsColumnFamilies{
		FamilyId:       dcl.StringOrNil(p.GetFamilyId()),
		Type:           dcl.StringOrNil(p.GetType()),
		Encoding:       dcl.StringOrNil(p.GetEncoding()),
		OnlyReadLatest: dcl.Bool(p.GetOnlyReadLatest()),
	}
	for _, r := range p.GetColumns() {
		obj.Columns = append(obj.Columns, *ProtoToBigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns converts a TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns object from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(p *betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) *beta.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{
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
func ProtoToBigqueryBetaTableExternalDataConfigurationGoogleSheetsOptions(p *betapb.BigqueryBetaTableExternalDataConfigurationGoogleSheetsOptions) *beta.TableExternalDataConfigurationGoogleSheetsOptions {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationGoogleSheetsOptions{
		SkipLeadingRows: dcl.StringOrNil(p.GetSkipLeadingRows()),
		Range:           dcl.StringOrNil(p.GetRange()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationHivePartitioningOptions converts a TableExternalDataConfigurationHivePartitioningOptions object from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationHivePartitioningOptions(p *betapb.BigqueryBetaTableExternalDataConfigurationHivePartitioningOptions) *beta.TableExternalDataConfigurationHivePartitioningOptions {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationHivePartitioningOptions{
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
func ProtoToBigqueryBetaTableExternalDataConfigurationValueConversionModes(p *betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModes) *beta.TableExternalDataConfigurationValueConversionModes {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationValueConversionModes{
		TemporalTypesOutOfRangeConversionMode: ProtoToBigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(p.GetTemporalTypesOutOfRangeConversionMode()),
		NumericTypeOutOfRangeConversionMode:   ProtoToBigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(p.GetNumericTypeOutOfRangeConversionMode()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationAvroOptions converts a TableExternalDataConfigurationAvroOptions object from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationAvroOptions(p *betapb.BigqueryBetaTableExternalDataConfigurationAvroOptions) *beta.TableExternalDataConfigurationAvroOptions {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationAvroOptions{
		UseAvroLogicalTypes: dcl.Bool(p.GetUseAvroLogicalTypes()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationParquetOptions converts a TableExternalDataConfigurationParquetOptions object from its proto representation.
func ProtoToBigqueryBetaTableExternalDataConfigurationParquetOptions(p *betapb.BigqueryBetaTableExternalDataConfigurationParquetOptions) *beta.TableExternalDataConfigurationParquetOptions {
	if p == nil {
		return nil
	}
	obj := &beta.TableExternalDataConfigurationParquetOptions{
		EnumAsString:        dcl.Bool(p.GetEnumAsString()),
		EnableListInference: dcl.Bool(p.GetEnableListInference()),
	}
	return obj
}

// ProtoToTableStreamingBuffer converts a TableStreamingBuffer object from its proto representation.
func ProtoToBigqueryBetaTableStreamingBuffer(p *betapb.BigqueryBetaTableStreamingBuffer) *beta.TableStreamingBuffer {
	if p == nil {
		return nil
	}
	obj := &beta.TableStreamingBuffer{
		EstimatedBytes:  dcl.Int64OrNil(p.GetEstimatedBytes()),
		EstimatedRows:   dcl.Int64OrNil(p.GetEstimatedRows()),
		OldestEntryTime: dcl.Int64OrNil(p.GetOldestEntryTime()),
	}
	return obj
}

// ProtoToTableEncryptionConfiguration converts a TableEncryptionConfiguration object from its proto representation.
func ProtoToBigqueryBetaTableEncryptionConfiguration(p *betapb.BigqueryBetaTableEncryptionConfiguration) *beta.TableEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.TableEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToTableSnapshotDefinition converts a TableSnapshotDefinition object from its proto representation.
func ProtoToBigqueryBetaTableSnapshotDefinition(p *betapb.BigqueryBetaTableSnapshotDefinition) *beta.TableSnapshotDefinition {
	if p == nil {
		return nil
	}
	obj := &beta.TableSnapshotDefinition{
		Table:        dcl.StringOrNil(p.GetTable()),
		Dataset:      dcl.StringOrNil(p.GetDataset()),
		Project:      dcl.StringOrNil(p.GetProject()),
		SnapshotTime: dcl.StringOrNil(p.GetSnapshotTime()),
	}
	return obj
}

// ProtoToTable converts a Table resource from its proto representation.
func ProtoToTable(p *betapb.BigqueryBetaTable) *beta.Table {
	obj := &beta.Table{
		Etag:                      dcl.StringOrNil(p.GetEtag()),
		Id:                        dcl.StringOrNil(p.GetId()),
		SelfLink:                  dcl.StringOrNil(p.GetSelfLink()),
		Name:                      dcl.StringOrNil(p.GetName()),
		Dataset:                   dcl.StringOrNil(p.GetDataset()),
		Project:                   dcl.StringOrNil(p.GetProject()),
		FriendlyName:              dcl.StringOrNil(p.GetFriendlyName()),
		Description:               dcl.StringOrNil(p.GetDescription()),
		Model:                     ProtoToBigqueryBetaTableModel(p.GetModel()),
		Schema:                    ProtoToBigqueryBetaTableSchema(p.GetSchema()),
		TimePartitioning:          ProtoToBigqueryBetaTableTimePartitioning(p.GetTimePartitioning()),
		RangePartitioning:         ProtoToBigqueryBetaTableRangePartitioning(p.GetRangePartitioning()),
		Clustering:                ProtoToBigqueryBetaTableClustering(p.GetClustering()),
		RequirePartitionFilter:    dcl.Bool(p.GetRequirePartitionFilter()),
		NumBytes:                  dcl.StringOrNil(p.GetNumBytes()),
		NumPhysicalBytes:          dcl.StringOrNil(p.GetNumPhysicalBytes()),
		NumLongTermBytes:          dcl.StringOrNil(p.GetNumLongTermBytes()),
		NumRows:                   dcl.Int64OrNil(p.GetNumRows()),
		CreationTime:              dcl.Int64OrNil(p.GetCreationTime()),
		ExpirationTime:            dcl.Int64OrNil(p.GetExpirationTime()),
		LastModifiedTime:          dcl.Int64OrNil(p.GetLastModifiedTime()),
		Type:                      dcl.StringOrNil(p.GetType()),
		View:                      ProtoToBigqueryBetaTableView(p.GetView()),
		MaterializedView:          ProtoToBigqueryBetaTableMaterializedView(p.GetMaterializedView()),
		ExternalDataConfiguration: ProtoToBigqueryBetaTableExternalDataConfiguration(p.GetExternalDataConfiguration()),
		Location:                  dcl.StringOrNil(p.GetLocation()),
		StreamingBuffer:           ProtoToBigqueryBetaTableStreamingBuffer(p.GetStreamingBuffer()),
		EncryptionConfiguration:   ProtoToBigqueryBetaTableEncryptionConfiguration(p.GetEncryptionConfiguration()),
		SnapshotDefinition:        ProtoToBigqueryBetaTableSnapshotDefinition(p.GetSnapshotDefinition()),
		DefaultCollation:          dcl.StringOrNil(p.GetDefaultCollation()),
	}
	return obj
}

// TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto converts a TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum enum to its proto representation.
func BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto(e *beta.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum) betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	if e == nil {
		return betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(0)
	}
	if v, ok := betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum_value["TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(v)
	}
	return betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(0)
}

// TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto converts a TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum enum to its proto representation.
func BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto(e *beta.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum) betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	if e == nil {
		return betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(0)
	}
	if v, ok := betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum_value["TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(v)
	}
	return betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(0)
}

// TableExternalDataConfigurationDecimalTargetTypesEnumToProto converts a TableExternalDataConfigurationDecimalTargetTypesEnum enum to its proto representation.
func BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnumToProto(e *beta.TableExternalDataConfigurationDecimalTargetTypesEnum) betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum {
	if e == nil {
		return betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum(0)
	}
	if v, ok := betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum_value["TableExternalDataConfigurationDecimalTargetTypesEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum(v)
	}
	return betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum(0)
}

// TableExternalDataConfigurationJsonExtensionEnumToProto converts a TableExternalDataConfigurationJsonExtensionEnum enum to its proto representation.
func BigqueryBetaTableExternalDataConfigurationJsonExtensionEnumToProto(e *beta.TableExternalDataConfigurationJsonExtensionEnum) betapb.BigqueryBetaTableExternalDataConfigurationJsonExtensionEnum {
	if e == nil {
		return betapb.BigqueryBetaTableExternalDataConfigurationJsonExtensionEnum(0)
	}
	if v, ok := betapb.BigqueryBetaTableExternalDataConfigurationJsonExtensionEnum_value["TableExternalDataConfigurationJsonExtensionEnum"+string(*e)]; ok {
		return betapb.BigqueryBetaTableExternalDataConfigurationJsonExtensionEnum(v)
	}
	return betapb.BigqueryBetaTableExternalDataConfigurationJsonExtensionEnum(0)
}

// TableModelToProto converts a TableModel object to its proto representation.
func BigqueryBetaTableModelToProto(o *beta.TableModel) *betapb.BigqueryBetaTableModel {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableModel{}
	p.SetModelOptions(BigqueryBetaTableModelModelOptionsToProto(o.ModelOptions))
	sTrainingRuns := make([]*betapb.BigqueryBetaTableModelTrainingRuns, len(o.TrainingRuns))
	for i, r := range o.TrainingRuns {
		sTrainingRuns[i] = BigqueryBetaTableModelTrainingRunsToProto(&r)
	}
	p.SetTrainingRuns(sTrainingRuns)
	return p
}

// TableModelModelOptionsToProto converts a TableModelModelOptions object to its proto representation.
func BigqueryBetaTableModelModelOptionsToProto(o *beta.TableModelModelOptions) *betapb.BigqueryBetaTableModelModelOptions {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableModelModelOptions{}
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
func BigqueryBetaTableModelTrainingRunsToProto(o *beta.TableModelTrainingRuns) *betapb.BigqueryBetaTableModelTrainingRuns {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableModelTrainingRuns{}
	p.SetState(dcl.ValueOrEmptyString(o.State))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetTrainingOptions(BigqueryBetaTableModelTrainingRunsTrainingOptionsToProto(o.TrainingOptions))
	sIterationResults := make([]*betapb.BigqueryBetaTableModelTrainingRunsIterationResults, len(o.IterationResults))
	for i, r := range o.IterationResults {
		sIterationResults[i] = BigqueryBetaTableModelTrainingRunsIterationResultsToProto(&r)
	}
	p.SetIterationResults(sIterationResults)
	return p
}

// TableModelTrainingRunsTrainingOptionsToProto converts a TableModelTrainingRunsTrainingOptions object to its proto representation.
func BigqueryBetaTableModelTrainingRunsTrainingOptionsToProto(o *beta.TableModelTrainingRunsTrainingOptions) *betapb.BigqueryBetaTableModelTrainingRunsTrainingOptions {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableModelTrainingRunsTrainingOptions{}
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
func BigqueryBetaTableModelTrainingRunsIterationResultsToProto(o *beta.TableModelTrainingRunsIterationResults) *betapb.BigqueryBetaTableModelTrainingRunsIterationResults {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableModelTrainingRunsIterationResults{}
	p.SetIndex(dcl.ValueOrEmptyInt64(o.Index))
	p.SetLearnRate(dcl.ValueOrEmptyDouble(o.LearnRate))
	p.SetTrainingLoss(dcl.ValueOrEmptyDouble(o.TrainingLoss))
	p.SetEvalLoss(dcl.ValueOrEmptyDouble(o.EvalLoss))
	p.SetDurationMs(dcl.ValueOrEmptyString(o.DurationMs))
	return p
}

// TableSchemaToProto converts a TableSchema object to its proto representation.
func BigqueryBetaTableSchemaToProto(o *beta.TableSchema) *betapb.BigqueryBetaTableSchema {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableSchema{}
	sFields := make([]*betapb.BigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// TableGooglecloudbigqueryv2TablefieldschemaToProto converts a TableGooglecloudbigqueryv2Tablefieldschema object to its proto representation.
func BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaToProto(o *beta.TableGooglecloudbigqueryv2Tablefieldschema) *betapb.BigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetMode(dcl.ValueOrEmptyString(o.Mode))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetCategories(BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaCategoriesToProto(o.Categories))
	p.SetPolicyTags(BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto(o.PolicyTags))
	p.SetMaxLength(dcl.ValueOrEmptyInt64(o.MaxLength))
	p.SetPrecision(dcl.ValueOrEmptyInt64(o.Precision))
	p.SetScale(dcl.ValueOrEmptyInt64(o.Scale))
	p.SetCollation(dcl.ValueOrEmptyString(o.Collation))
	p.SetDefaultValueExpression(dcl.ValueOrEmptyString(o.DefaultValueExpression))
	sFields := make([]*betapb.BigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
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
func BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaCategoriesToProto(o *beta.TableGooglecloudbigqueryv2TablefieldschemaCategories) *betapb.BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaCategories {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaCategories{}
	sNames := make([]string, len(o.Names))
	for i, r := range o.Names {
		sNames[i] = r
	}
	p.SetNames(sNames)
	return p
}

// TableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto converts a TableGooglecloudbigqueryv2TablefieldschemaPolicyTags object to its proto representation.
func BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto(o *beta.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) *betapb.BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	sNames := make([]string, len(o.Names))
	for i, r := range o.Names {
		sNames[i] = r
	}
	p.SetNames(sNames)
	return p
}

// TableTimePartitioningToProto converts a TableTimePartitioning object to its proto representation.
func BigqueryBetaTableTimePartitioningToProto(o *beta.TableTimePartitioning) *betapb.BigqueryBetaTableTimePartitioning {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableTimePartitioning{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetExpirationMs(dcl.ValueOrEmptyString(o.ExpirationMs))
	p.SetField(dcl.ValueOrEmptyString(o.Field))
	return p
}

// TableRangePartitioningToProto converts a TableRangePartitioning object to its proto representation.
func BigqueryBetaTableRangePartitioningToProto(o *beta.TableRangePartitioning) *betapb.BigqueryBetaTableRangePartitioning {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableRangePartitioning{}
	p.SetField(dcl.ValueOrEmptyString(o.Field))
	p.SetRange(BigqueryBetaTableRangePartitioningRangeToProto(o.Range))
	return p
}

// TableRangePartitioningRangeToProto converts a TableRangePartitioningRange object to its proto representation.
func BigqueryBetaTableRangePartitioningRangeToProto(o *beta.TableRangePartitioningRange) *betapb.BigqueryBetaTableRangePartitioningRange {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableRangePartitioningRange{}
	p.SetStart(dcl.ValueOrEmptyString(o.Start))
	p.SetEnd(dcl.ValueOrEmptyString(o.End))
	p.SetInterval(dcl.ValueOrEmptyString(o.Interval))
	return p
}

// TableClusteringToProto converts a TableClustering object to its proto representation.
func BigqueryBetaTableClusteringToProto(o *beta.TableClustering) *betapb.BigqueryBetaTableClustering {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableClustering{}
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// TableViewToProto converts a TableView object to its proto representation.
func BigqueryBetaTableViewToProto(o *beta.TableView) *betapb.BigqueryBetaTableView {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableView{}
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	p.SetUseLegacySql(dcl.ValueOrEmptyBool(o.UseLegacySql))
	p.SetUseExplicitColumnNames(dcl.ValueOrEmptyBool(o.UseExplicitColumnNames))
	sUserDefinedFunctionResources := make([]*betapb.BigqueryBetaTableViewUserDefinedFunctionResources, len(o.UserDefinedFunctionResources))
	for i, r := range o.UserDefinedFunctionResources {
		sUserDefinedFunctionResources[i] = BigqueryBetaTableViewUserDefinedFunctionResourcesToProto(&r)
	}
	p.SetUserDefinedFunctionResources(sUserDefinedFunctionResources)
	return p
}

// TableViewUserDefinedFunctionResourcesToProto converts a TableViewUserDefinedFunctionResources object to its proto representation.
func BigqueryBetaTableViewUserDefinedFunctionResourcesToProto(o *beta.TableViewUserDefinedFunctionResources) *betapb.BigqueryBetaTableViewUserDefinedFunctionResources {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableViewUserDefinedFunctionResources{}
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
func BigqueryBetaTableMaterializedViewToProto(o *beta.TableMaterializedView) *betapb.BigqueryBetaTableMaterializedView {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableMaterializedView{}
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	p.SetLastRefreshTime(dcl.ValueOrEmptyInt64(o.LastRefreshTime))
	p.SetEnableRefresh(dcl.ValueOrEmptyBool(o.EnableRefresh))
	p.SetRefreshIntervalMs(dcl.ValueOrEmptyInt64(o.RefreshIntervalMs))
	return p
}

// TableExternalDataConfigurationToProto converts a TableExternalDataConfiguration object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationToProto(o *beta.TableExternalDataConfiguration) *betapb.BigqueryBetaTableExternalDataConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfiguration{}
	p.SetSchema(BigqueryBetaTableExternalDataConfigurationSchemaToProto(o.Schema))
	p.SetSourceFormat(dcl.ValueOrEmptyString(o.SourceFormat))
	p.SetMaxBadRecords(dcl.ValueOrEmptyInt64(o.MaxBadRecords))
	p.SetAutodetect(dcl.ValueOrEmptyBool(o.Autodetect))
	p.SetIgnoreUnknownValues(dcl.ValueOrEmptyBool(o.IgnoreUnknownValues))
	p.SetCompression(dcl.ValueOrEmptyString(o.Compression))
	p.SetCsvOptions(BigqueryBetaTableExternalDataConfigurationCsvOptionsToProto(o.CsvOptions))
	p.SetBigtableOptions(BigqueryBetaTableExternalDataConfigurationBigtableOptionsToProto(o.BigtableOptions))
	p.SetGoogleSheetsOptions(BigqueryBetaTableExternalDataConfigurationGoogleSheetsOptionsToProto(o.GoogleSheetsOptions))
	p.SetHivePartitioningOptions(BigqueryBetaTableExternalDataConfigurationHivePartitioningOptionsToProto(o.HivePartitioningOptions))
	p.SetConnectionId(dcl.ValueOrEmptyString(o.ConnectionId))
	p.SetValueConversionModes(BigqueryBetaTableExternalDataConfigurationValueConversionModesToProto(o.ValueConversionModes))
	p.SetAvroOptions(BigqueryBetaTableExternalDataConfigurationAvroOptionsToProto(o.AvroOptions))
	p.SetJsonExtension(BigqueryBetaTableExternalDataConfigurationJsonExtensionEnumToProto(o.JsonExtension))
	p.SetParquetOptions(BigqueryBetaTableExternalDataConfigurationParquetOptionsToProto(o.ParquetOptions))
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
	sDecimalTargetTypes := make([]betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum, len(o.DecimalTargetTypes))
	for i, r := range o.DecimalTargetTypes {
		sDecimalTargetTypes[i] = betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum(betapb.BigqueryBetaTableExternalDataConfigurationDecimalTargetTypesEnum_value[string(r)])
	}
	p.SetDecimalTargetTypes(sDecimalTargetTypes)
	return p
}

// TableExternalDataConfigurationSchemaToProto converts a TableExternalDataConfigurationSchema object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationSchemaToProto(o *beta.TableExternalDataConfigurationSchema) *betapb.BigqueryBetaTableExternalDataConfigurationSchema {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationSchema{}
	sFields := make([]*betapb.BigqueryBetaTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryBetaTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// TableExternalDataConfigurationCsvOptionsToProto converts a TableExternalDataConfigurationCsvOptions object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationCsvOptionsToProto(o *beta.TableExternalDataConfigurationCsvOptions) *betapb.BigqueryBetaTableExternalDataConfigurationCsvOptions {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationCsvOptions{}
	p.SetFieldDelimiter(dcl.ValueOrEmptyString(o.FieldDelimiter))
	p.SetSkipLeadingRows(dcl.ValueOrEmptyString(o.SkipLeadingRows))
	p.SetQuote(dcl.ValueOrEmptyString(o.Quote))
	p.SetAllowQuotedNewlines(dcl.ValueOrEmptyBool(o.AllowQuotedNewlines))
	p.SetAllowJaggedRows(dcl.ValueOrEmptyBool(o.AllowJaggedRows))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	return p
}

// TableExternalDataConfigurationBigtableOptionsToProto converts a TableExternalDataConfigurationBigtableOptions object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationBigtableOptionsToProto(o *beta.TableExternalDataConfigurationBigtableOptions) *betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptions {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptions{}
	p.SetIgnoreUnspecifiedColumnFamilies(dcl.ValueOrEmptyBool(o.IgnoreUnspecifiedColumnFamilies))
	p.SetReadRowkeyAsString(dcl.ValueOrEmptyBool(o.ReadRowkeyAsString))
	sColumnFamilies := make([]*betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamilies, len(o.ColumnFamilies))
	for i, r := range o.ColumnFamilies {
		sColumnFamilies[i] = BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto(&r)
	}
	p.SetColumnFamilies(sColumnFamilies)
	return p
}

// TableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto converts a TableExternalDataConfigurationBigtableOptionsColumnFamilies object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto(o *beta.TableExternalDataConfigurationBigtableOptionsColumnFamilies) *betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamilies{}
	p.SetFamilyId(dcl.ValueOrEmptyString(o.FamilyId))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetOnlyReadLatest(dcl.ValueOrEmptyBool(o.OnlyReadLatest))
	sColumns := make([]*betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, len(o.Columns))
	for i, r := range o.Columns {
		sColumns[i] = BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto(&r)
	}
	p.SetColumns(sColumns)
	return p
}

// TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto converts a TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto(o *beta.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) *betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}
	p.SetQualifierEncoded(dcl.ValueOrEmptyString(o.QualifierEncoded))
	p.SetQualifierString(dcl.ValueOrEmptyString(o.QualifierString))
	p.SetFieldName(dcl.ValueOrEmptyString(o.FieldName))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetOnlyReadLatest(dcl.ValueOrEmptyBool(o.OnlyReadLatest))
	return p
}

// TableExternalDataConfigurationGoogleSheetsOptionsToProto converts a TableExternalDataConfigurationGoogleSheetsOptions object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationGoogleSheetsOptionsToProto(o *beta.TableExternalDataConfigurationGoogleSheetsOptions) *betapb.BigqueryBetaTableExternalDataConfigurationGoogleSheetsOptions {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationGoogleSheetsOptions{}
	p.SetSkipLeadingRows(dcl.ValueOrEmptyString(o.SkipLeadingRows))
	p.SetRange(dcl.ValueOrEmptyString(o.Range))
	return p
}

// TableExternalDataConfigurationHivePartitioningOptionsToProto converts a TableExternalDataConfigurationHivePartitioningOptions object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationHivePartitioningOptionsToProto(o *beta.TableExternalDataConfigurationHivePartitioningOptions) *betapb.BigqueryBetaTableExternalDataConfigurationHivePartitioningOptions {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationHivePartitioningOptions{}
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
func BigqueryBetaTableExternalDataConfigurationValueConversionModesToProto(o *beta.TableExternalDataConfigurationValueConversionModes) *betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModes {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationValueConversionModes{}
	p.SetTemporalTypesOutOfRangeConversionMode(BigqueryBetaTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto(o.TemporalTypesOutOfRangeConversionMode))
	p.SetNumericTypeOutOfRangeConversionMode(BigqueryBetaTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto(o.NumericTypeOutOfRangeConversionMode))
	return p
}

// TableExternalDataConfigurationAvroOptionsToProto converts a TableExternalDataConfigurationAvroOptions object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationAvroOptionsToProto(o *beta.TableExternalDataConfigurationAvroOptions) *betapb.BigqueryBetaTableExternalDataConfigurationAvroOptions {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationAvroOptions{}
	p.SetUseAvroLogicalTypes(dcl.ValueOrEmptyBool(o.UseAvroLogicalTypes))
	return p
}

// TableExternalDataConfigurationParquetOptionsToProto converts a TableExternalDataConfigurationParquetOptions object to its proto representation.
func BigqueryBetaTableExternalDataConfigurationParquetOptionsToProto(o *beta.TableExternalDataConfigurationParquetOptions) *betapb.BigqueryBetaTableExternalDataConfigurationParquetOptions {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableExternalDataConfigurationParquetOptions{}
	p.SetEnumAsString(dcl.ValueOrEmptyBool(o.EnumAsString))
	p.SetEnableListInference(dcl.ValueOrEmptyBool(o.EnableListInference))
	return p
}

// TableStreamingBufferToProto converts a TableStreamingBuffer object to its proto representation.
func BigqueryBetaTableStreamingBufferToProto(o *beta.TableStreamingBuffer) *betapb.BigqueryBetaTableStreamingBuffer {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableStreamingBuffer{}
	p.SetEstimatedBytes(dcl.ValueOrEmptyInt64(o.EstimatedBytes))
	p.SetEstimatedRows(dcl.ValueOrEmptyInt64(o.EstimatedRows))
	p.SetOldestEntryTime(dcl.ValueOrEmptyInt64(o.OldestEntryTime))
	return p
}

// TableEncryptionConfigurationToProto converts a TableEncryptionConfiguration object to its proto representation.
func BigqueryBetaTableEncryptionConfigurationToProto(o *beta.TableEncryptionConfiguration) *betapb.BigqueryBetaTableEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableEncryptionConfiguration{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// TableSnapshotDefinitionToProto converts a TableSnapshotDefinition object to its proto representation.
func BigqueryBetaTableSnapshotDefinitionToProto(o *beta.TableSnapshotDefinition) *betapb.BigqueryBetaTableSnapshotDefinition {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaTableSnapshotDefinition{}
	p.SetTable(dcl.ValueOrEmptyString(o.Table))
	p.SetDataset(dcl.ValueOrEmptyString(o.Dataset))
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetSnapshotTime(dcl.ValueOrEmptyString(o.SnapshotTime))
	return p
}

// TableToProto converts a Table resource to its proto representation.
func TableToProto(resource *beta.Table) *betapb.BigqueryBetaTable {
	p := &betapb.BigqueryBetaTable{}
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetId(dcl.ValueOrEmptyString(resource.Id))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetFriendlyName(dcl.ValueOrEmptyString(resource.FriendlyName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetModel(BigqueryBetaTableModelToProto(resource.Model))
	p.SetSchema(BigqueryBetaTableSchemaToProto(resource.Schema))
	p.SetTimePartitioning(BigqueryBetaTableTimePartitioningToProto(resource.TimePartitioning))
	p.SetRangePartitioning(BigqueryBetaTableRangePartitioningToProto(resource.RangePartitioning))
	p.SetClustering(BigqueryBetaTableClusteringToProto(resource.Clustering))
	p.SetRequirePartitionFilter(dcl.ValueOrEmptyBool(resource.RequirePartitionFilter))
	p.SetNumBytes(dcl.ValueOrEmptyString(resource.NumBytes))
	p.SetNumPhysicalBytes(dcl.ValueOrEmptyString(resource.NumPhysicalBytes))
	p.SetNumLongTermBytes(dcl.ValueOrEmptyString(resource.NumLongTermBytes))
	p.SetNumRows(dcl.ValueOrEmptyInt64(resource.NumRows))
	p.SetCreationTime(dcl.ValueOrEmptyInt64(resource.CreationTime))
	p.SetExpirationTime(dcl.ValueOrEmptyInt64(resource.ExpirationTime))
	p.SetLastModifiedTime(dcl.ValueOrEmptyInt64(resource.LastModifiedTime))
	p.SetType(dcl.ValueOrEmptyString(resource.Type))
	p.SetView(BigqueryBetaTableViewToProto(resource.View))
	p.SetMaterializedView(BigqueryBetaTableMaterializedViewToProto(resource.MaterializedView))
	p.SetExternalDataConfiguration(BigqueryBetaTableExternalDataConfigurationToProto(resource.ExternalDataConfiguration))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetStreamingBuffer(BigqueryBetaTableStreamingBufferToProto(resource.StreamingBuffer))
	p.SetEncryptionConfiguration(BigqueryBetaTableEncryptionConfigurationToProto(resource.EncryptionConfiguration))
	p.SetSnapshotDefinition(BigqueryBetaTableSnapshotDefinitionToProto(resource.SnapshotDefinition))
	p.SetDefaultCollation(dcl.ValueOrEmptyString(resource.DefaultCollation))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyTable handles the gRPC request by passing it to the underlying Table Apply() method.
func (s *TableServer) applyTable(ctx context.Context, c *beta.Client, request *betapb.ApplyBigqueryBetaTableRequest) (*betapb.BigqueryBetaTable, error) {
	p := ProtoToTable(request.GetResource())
	res, err := c.ApplyTable(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TableToProto(res)
	return r, nil
}

// applyBigqueryBetaTable handles the gRPC request by passing it to the underlying Table Apply() method.
func (s *TableServer) ApplyBigqueryBetaTable(ctx context.Context, request *betapb.ApplyBigqueryBetaTableRequest) (*betapb.BigqueryBetaTable, error) {
	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTable(ctx, cl, request)
}

// DeleteTable handles the gRPC request by passing it to the underlying Table Delete() method.
func (s *TableServer) DeleteBigqueryBetaTable(ctx context.Context, request *betapb.DeleteBigqueryBetaTableRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTable(ctx, ProtoToTable(request.GetResource()))

}

// ListBigqueryBetaTable handles the gRPC request by passing it to the underlying TableList() method.
func (s *TableServer) ListBigqueryBetaTable(ctx context.Context, request *betapb.ListBigqueryBetaTableRequest) (*betapb.ListBigqueryBetaTableResponse, error) {
	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTable(ctx, request.GetProject(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.BigqueryBetaTable
	for _, r := range resources.Items {
		rp := TableToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListBigqueryBetaTableResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTable(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
