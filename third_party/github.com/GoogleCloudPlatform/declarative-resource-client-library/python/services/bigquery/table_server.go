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
	bigquerypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/bigquery_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery"
)

// TableServer implements the gRPC interface for Table.
type TableServer struct{}

// ProtoToTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum converts a TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum enum from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(e bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum) *bigquery.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum_name[int32(e)]; ok {
		e := bigquery.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(n[len("BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum converts a TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum enum from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(e bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum) *bigquery.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum_name[int32(e)]; ok {
		e := bigquery.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(n[len("BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationDecimalTargetTypesEnum converts a TableExternalDataConfigurationDecimalTargetTypesEnum enum from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationDecimalTargetTypesEnum(e bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum) *bigquery.TableExternalDataConfigurationDecimalTargetTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum_name[int32(e)]; ok {
		e := bigquery.TableExternalDataConfigurationDecimalTargetTypesEnum(n[len("BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableExternalDataConfigurationJsonExtensionEnum converts a TableExternalDataConfigurationJsonExtensionEnum enum from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationJsonExtensionEnum(e bigquerypb.BigqueryTableExternalDataConfigurationJsonExtensionEnum) *bigquery.TableExternalDataConfigurationJsonExtensionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigquerypb.BigqueryTableExternalDataConfigurationJsonExtensionEnum_name[int32(e)]; ok {
		e := bigquery.TableExternalDataConfigurationJsonExtensionEnum(n[len("BigqueryTableExternalDataConfigurationJsonExtensionEnum"):])
		return &e
	}
	return nil
}

// ProtoToTableModel converts a TableModel object from its proto representation.
func ProtoToBigqueryTableModel(p *bigquerypb.BigqueryTableModel) *bigquery.TableModel {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableModel{
		ModelOptions: ProtoToBigqueryTableModelModelOptions(p.GetModelOptions()),
	}
	for _, r := range p.GetTrainingRuns() {
		obj.TrainingRuns = append(obj.TrainingRuns, *ProtoToBigqueryTableModelTrainingRuns(r))
	}
	return obj
}

// ProtoToTableModelModelOptions converts a TableModelModelOptions object from its proto representation.
func ProtoToBigqueryTableModelModelOptions(p *bigquerypb.BigqueryTableModelModelOptions) *bigquery.TableModelModelOptions {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableModelModelOptions{
		ModelType: dcl.StringOrNil(p.GetModelType()),
		LossType:  dcl.StringOrNil(p.GetLossType()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, r)
	}
	return obj
}

// ProtoToTableModelTrainingRuns converts a TableModelTrainingRuns object from its proto representation.
func ProtoToBigqueryTableModelTrainingRuns(p *bigquerypb.BigqueryTableModelTrainingRuns) *bigquery.TableModelTrainingRuns {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableModelTrainingRuns{
		State:           dcl.StringOrNil(p.GetState()),
		StartTime:       dcl.StringOrNil(p.GetStartTime()),
		TrainingOptions: ProtoToBigqueryTableModelTrainingRunsTrainingOptions(p.GetTrainingOptions()),
	}
	for _, r := range p.GetIterationResults() {
		obj.IterationResults = append(obj.IterationResults, *ProtoToBigqueryTableModelTrainingRunsIterationResults(r))
	}
	return obj
}

// ProtoToTableModelTrainingRunsTrainingOptions converts a TableModelTrainingRunsTrainingOptions object from its proto representation.
func ProtoToBigqueryTableModelTrainingRunsTrainingOptions(p *bigquerypb.BigqueryTableModelTrainingRunsTrainingOptions) *bigquery.TableModelTrainingRunsTrainingOptions {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableModelTrainingRunsTrainingOptions{
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
func ProtoToBigqueryTableModelTrainingRunsIterationResults(p *bigquerypb.BigqueryTableModelTrainingRunsIterationResults) *bigquery.TableModelTrainingRunsIterationResults {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableModelTrainingRunsIterationResults{
		Index:        dcl.Int64OrNil(p.GetIndex()),
		LearnRate:    dcl.Float64OrNil(p.GetLearnRate()),
		TrainingLoss: dcl.Float64OrNil(p.GetTrainingLoss()),
		EvalLoss:     dcl.Float64OrNil(p.GetEvalLoss()),
		DurationMs:   dcl.StringOrNil(p.GetDurationMs()),
	}
	return obj
}

// ProtoToTableSchema converts a TableSchema object from its proto representation.
func ProtoToBigqueryTableSchema(p *bigquerypb.BigqueryTableSchema) *bigquery.TableSchema {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableSchema{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2Tablefieldschema converts a TableGooglecloudbigqueryv2Tablefieldschema object from its proto representation.
func ProtoToBigqueryTableGooglecloudbigqueryv2Tablefieldschema(p *bigquerypb.BigqueryTableGooglecloudbigqueryv2Tablefieldschema) *bigquery.TableGooglecloudbigqueryv2Tablefieldschema {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableGooglecloudbigqueryv2Tablefieldschema{
		Name:                   dcl.StringOrNil(p.GetName()),
		Type:                   dcl.StringOrNil(p.GetType()),
		Mode:                   dcl.StringOrNil(p.GetMode()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		Categories:             ProtoToBigqueryTableGooglecloudbigqueryv2TablefieldschemaCategories(p.GetCategories()),
		PolicyTags:             ProtoToBigqueryTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(p.GetPolicyTags()),
		MaxLength:              dcl.Int64OrNil(p.GetMaxLength()),
		Precision:              dcl.Int64OrNil(p.GetPrecision()),
		Scale:                  dcl.Int64OrNil(p.GetScale()),
		Collation:              dcl.StringOrNil(p.GetCollation()),
		DefaultValueExpression: dcl.StringOrNil(p.GetDefaultValueExpression()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	for _, r := range p.GetNameAlternative() {
		obj.NameAlternative = append(obj.NameAlternative, r)
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2TablefieldschemaCategories converts a TableGooglecloudbigqueryv2TablefieldschemaCategories object from its proto representation.
func ProtoToBigqueryTableGooglecloudbigqueryv2TablefieldschemaCategories(p *bigquerypb.BigqueryTableGooglecloudbigqueryv2TablefieldschemaCategories) *bigquery.TableGooglecloudbigqueryv2TablefieldschemaCategories {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableGooglecloudbigqueryv2TablefieldschemaCategories{}
	for _, r := range p.GetNames() {
		obj.Names = append(obj.Names, r)
	}
	return obj
}

// ProtoToTableGooglecloudbigqueryv2TablefieldschemaPolicyTags converts a TableGooglecloudbigqueryv2TablefieldschemaPolicyTags object from its proto representation.
func ProtoToBigqueryTableGooglecloudbigqueryv2TablefieldschemaPolicyTags(p *bigquerypb.BigqueryTableGooglecloudbigqueryv2TablefieldschemaPolicyTags) *bigquery.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	for _, r := range p.GetNames() {
		obj.Names = append(obj.Names, r)
	}
	return obj
}

// ProtoToTableTimePartitioning converts a TableTimePartitioning object from its proto representation.
func ProtoToBigqueryTableTimePartitioning(p *bigquerypb.BigqueryTableTimePartitioning) *bigquery.TableTimePartitioning {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableTimePartitioning{
		Type:         dcl.StringOrNil(p.GetType()),
		ExpirationMs: dcl.StringOrNil(p.GetExpirationMs()),
		Field:        dcl.StringOrNil(p.GetField()),
	}
	return obj
}

// ProtoToTableRangePartitioning converts a TableRangePartitioning object from its proto representation.
func ProtoToBigqueryTableRangePartitioning(p *bigquerypb.BigqueryTableRangePartitioning) *bigquery.TableRangePartitioning {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableRangePartitioning{
		Field: dcl.StringOrNil(p.GetField()),
		Range: ProtoToBigqueryTableRangePartitioningRange(p.GetRange()),
	}
	return obj
}

// ProtoToTableRangePartitioningRange converts a TableRangePartitioningRange object from its proto representation.
func ProtoToBigqueryTableRangePartitioningRange(p *bigquerypb.BigqueryTableRangePartitioningRange) *bigquery.TableRangePartitioningRange {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableRangePartitioningRange{
		Start:    dcl.StringOrNil(p.GetStart()),
		End:      dcl.StringOrNil(p.GetEnd()),
		Interval: dcl.StringOrNil(p.GetInterval()),
	}
	return obj
}

// ProtoToTableClustering converts a TableClustering object from its proto representation.
func ProtoToBigqueryTableClustering(p *bigquerypb.BigqueryTableClustering) *bigquery.TableClustering {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableClustering{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToTableView converts a TableView object from its proto representation.
func ProtoToBigqueryTableView(p *bigquerypb.BigqueryTableView) *bigquery.TableView {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableView{
		Query:                  dcl.StringOrNil(p.GetQuery()),
		UseLegacySql:           dcl.Bool(p.GetUseLegacySql()),
		UseExplicitColumnNames: dcl.Bool(p.GetUseExplicitColumnNames()),
	}
	for _, r := range p.GetUserDefinedFunctionResources() {
		obj.UserDefinedFunctionResources = append(obj.UserDefinedFunctionResources, *ProtoToBigqueryTableViewUserDefinedFunctionResources(r))
	}
	return obj
}

// ProtoToTableViewUserDefinedFunctionResources converts a TableViewUserDefinedFunctionResources object from its proto representation.
func ProtoToBigqueryTableViewUserDefinedFunctionResources(p *bigquerypb.BigqueryTableViewUserDefinedFunctionResources) *bigquery.TableViewUserDefinedFunctionResources {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableViewUserDefinedFunctionResources{
		ResourceUri: dcl.StringOrNil(p.GetResourceUri()),
		InlineCode:  dcl.StringOrNil(p.GetInlineCode()),
	}
	for _, r := range p.GetInlineCodeAlternative() {
		obj.InlineCodeAlternative = append(obj.InlineCodeAlternative, r)
	}
	return obj
}

// ProtoToTableMaterializedView converts a TableMaterializedView object from its proto representation.
func ProtoToBigqueryTableMaterializedView(p *bigquerypb.BigqueryTableMaterializedView) *bigquery.TableMaterializedView {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableMaterializedView{
		Query:             dcl.StringOrNil(p.GetQuery()),
		LastRefreshTime:   dcl.Int64OrNil(p.GetLastRefreshTime()),
		EnableRefresh:     dcl.Bool(p.GetEnableRefresh()),
		RefreshIntervalMs: dcl.Int64OrNil(p.GetRefreshIntervalMs()),
	}
	return obj
}

// ProtoToTableExternalDataConfiguration converts a TableExternalDataConfiguration object from its proto representation.
func ProtoToBigqueryTableExternalDataConfiguration(p *bigquerypb.BigqueryTableExternalDataConfiguration) *bigquery.TableExternalDataConfiguration {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfiguration{
		Schema:                  ProtoToBigqueryTableExternalDataConfigurationSchema(p.GetSchema()),
		SourceFormat:            dcl.StringOrNil(p.GetSourceFormat()),
		MaxBadRecords:           dcl.Int64OrNil(p.GetMaxBadRecords()),
		Autodetect:              dcl.Bool(p.GetAutodetect()),
		IgnoreUnknownValues:     dcl.Bool(p.GetIgnoreUnknownValues()),
		Compression:             dcl.StringOrNil(p.GetCompression()),
		CsvOptions:              ProtoToBigqueryTableExternalDataConfigurationCsvOptions(p.GetCsvOptions()),
		BigtableOptions:         ProtoToBigqueryTableExternalDataConfigurationBigtableOptions(p.GetBigtableOptions()),
		GoogleSheetsOptions:     ProtoToBigqueryTableExternalDataConfigurationGoogleSheetsOptions(p.GetGoogleSheetsOptions()),
		HivePartitioningOptions: ProtoToBigqueryTableExternalDataConfigurationHivePartitioningOptions(p.GetHivePartitioningOptions()),
		ConnectionId:            dcl.StringOrNil(p.GetConnectionId()),
		ValueConversionModes:    ProtoToBigqueryTableExternalDataConfigurationValueConversionModes(p.GetValueConversionModes()),
		AvroOptions:             ProtoToBigqueryTableExternalDataConfigurationAvroOptions(p.GetAvroOptions()),
		JsonExtension:           ProtoToBigqueryTableExternalDataConfigurationJsonExtensionEnum(p.GetJsonExtension()),
		ParquetOptions:          ProtoToBigqueryTableExternalDataConfigurationParquetOptions(p.GetParquetOptions()),
	}
	for _, r := range p.GetSourceUris() {
		obj.SourceUris = append(obj.SourceUris, r)
	}
	for _, r := range p.GetMaxBadRecordsAlternative() {
		obj.MaxBadRecordsAlternative = append(obj.MaxBadRecordsAlternative, r)
	}
	for _, r := range p.GetDecimalTargetTypes() {
		obj.DecimalTargetTypes = append(obj.DecimalTargetTypes, *ProtoToBigqueryTableExternalDataConfigurationDecimalTargetTypesEnum(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationSchema converts a TableExternalDataConfigurationSchema object from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationSchema(p *bigquerypb.BigqueryTableExternalDataConfigurationSchema) *bigquery.TableExternalDataConfigurationSchema {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationSchema{}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, *ProtoToBigqueryTableGooglecloudbigqueryv2Tablefieldschema(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationCsvOptions converts a TableExternalDataConfigurationCsvOptions object from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationCsvOptions(p *bigquerypb.BigqueryTableExternalDataConfigurationCsvOptions) *bigquery.TableExternalDataConfigurationCsvOptions {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationCsvOptions{
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
func ProtoToBigqueryTableExternalDataConfigurationBigtableOptions(p *bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptions) *bigquery.TableExternalDataConfigurationBigtableOptions {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationBigtableOptions{
		IgnoreUnspecifiedColumnFamilies: dcl.Bool(p.GetIgnoreUnspecifiedColumnFamilies()),
		ReadRowkeyAsString:              dcl.Bool(p.GetReadRowkeyAsString()),
	}
	for _, r := range p.GetColumnFamilies() {
		obj.ColumnFamilies = append(obj.ColumnFamilies, *ProtoToBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilies(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationBigtableOptionsColumnFamilies converts a TableExternalDataConfigurationBigtableOptionsColumnFamilies object from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilies(p *bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilies) *bigquery.TableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationBigtableOptionsColumnFamilies{
		FamilyId:       dcl.StringOrNil(p.GetFamilyId()),
		Type:           dcl.StringOrNil(p.GetType()),
		Encoding:       dcl.StringOrNil(p.GetEncoding()),
		OnlyReadLatest: dcl.Bool(p.GetOnlyReadLatest()),
	}
	for _, r := range p.GetColumns() {
		obj.Columns = append(obj.Columns, *ProtoToBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(r))
	}
	return obj
}

// ProtoToTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns converts a TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns object from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns(p *bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) *bigquery.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{
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
func ProtoToBigqueryTableExternalDataConfigurationGoogleSheetsOptions(p *bigquerypb.BigqueryTableExternalDataConfigurationGoogleSheetsOptions) *bigquery.TableExternalDataConfigurationGoogleSheetsOptions {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationGoogleSheetsOptions{
		SkipLeadingRows: dcl.StringOrNil(p.GetSkipLeadingRows()),
		Range:           dcl.StringOrNil(p.GetRange()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationHivePartitioningOptions converts a TableExternalDataConfigurationHivePartitioningOptions object from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationHivePartitioningOptions(p *bigquerypb.BigqueryTableExternalDataConfigurationHivePartitioningOptions) *bigquery.TableExternalDataConfigurationHivePartitioningOptions {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationHivePartitioningOptions{
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
func ProtoToBigqueryTableExternalDataConfigurationValueConversionModes(p *bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModes) *bigquery.TableExternalDataConfigurationValueConversionModes {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationValueConversionModes{
		TemporalTypesOutOfRangeConversionMode: ProtoToBigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(p.GetTemporalTypesOutOfRangeConversionMode()),
		NumericTypeOutOfRangeConversionMode:   ProtoToBigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(p.GetNumericTypeOutOfRangeConversionMode()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationAvroOptions converts a TableExternalDataConfigurationAvroOptions object from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationAvroOptions(p *bigquerypb.BigqueryTableExternalDataConfigurationAvroOptions) *bigquery.TableExternalDataConfigurationAvroOptions {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationAvroOptions{
		UseAvroLogicalTypes: dcl.Bool(p.GetUseAvroLogicalTypes()),
	}
	return obj
}

// ProtoToTableExternalDataConfigurationParquetOptions converts a TableExternalDataConfigurationParquetOptions object from its proto representation.
func ProtoToBigqueryTableExternalDataConfigurationParquetOptions(p *bigquerypb.BigqueryTableExternalDataConfigurationParquetOptions) *bigquery.TableExternalDataConfigurationParquetOptions {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableExternalDataConfigurationParquetOptions{
		EnumAsString:        dcl.Bool(p.GetEnumAsString()),
		EnableListInference: dcl.Bool(p.GetEnableListInference()),
	}
	return obj
}

// ProtoToTableStreamingBuffer converts a TableStreamingBuffer object from its proto representation.
func ProtoToBigqueryTableStreamingBuffer(p *bigquerypb.BigqueryTableStreamingBuffer) *bigquery.TableStreamingBuffer {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableStreamingBuffer{
		EstimatedBytes:  dcl.Int64OrNil(p.GetEstimatedBytes()),
		EstimatedRows:   dcl.Int64OrNil(p.GetEstimatedRows()),
		OldestEntryTime: dcl.Int64OrNil(p.GetOldestEntryTime()),
	}
	return obj
}

// ProtoToTableEncryptionConfiguration converts a TableEncryptionConfiguration object from its proto representation.
func ProtoToBigqueryTableEncryptionConfiguration(p *bigquerypb.BigqueryTableEncryptionConfiguration) *bigquery.TableEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToTableSnapshotDefinition converts a TableSnapshotDefinition object from its proto representation.
func ProtoToBigqueryTableSnapshotDefinition(p *bigquerypb.BigqueryTableSnapshotDefinition) *bigquery.TableSnapshotDefinition {
	if p == nil {
		return nil
	}
	obj := &bigquery.TableSnapshotDefinition{
		Table:        dcl.StringOrNil(p.GetTable()),
		Dataset:      dcl.StringOrNil(p.GetDataset()),
		Project:      dcl.StringOrNil(p.GetProject()),
		SnapshotTime: dcl.StringOrNil(p.GetSnapshotTime()),
	}
	return obj
}

// ProtoToTable converts a Table resource from its proto representation.
func ProtoToTable(p *bigquerypb.BigqueryTable) *bigquery.Table {
	obj := &bigquery.Table{
		Etag:                      dcl.StringOrNil(p.GetEtag()),
		Id:                        dcl.StringOrNil(p.GetId()),
		SelfLink:                  dcl.StringOrNil(p.GetSelfLink()),
		Name:                      dcl.StringOrNil(p.GetName()),
		Dataset:                   dcl.StringOrNil(p.GetDataset()),
		Project:                   dcl.StringOrNil(p.GetProject()),
		FriendlyName:              dcl.StringOrNil(p.GetFriendlyName()),
		Description:               dcl.StringOrNil(p.GetDescription()),
		Model:                     ProtoToBigqueryTableModel(p.GetModel()),
		Schema:                    ProtoToBigqueryTableSchema(p.GetSchema()),
		TimePartitioning:          ProtoToBigqueryTableTimePartitioning(p.GetTimePartitioning()),
		RangePartitioning:         ProtoToBigqueryTableRangePartitioning(p.GetRangePartitioning()),
		Clustering:                ProtoToBigqueryTableClustering(p.GetClustering()),
		RequirePartitionFilter:    dcl.Bool(p.GetRequirePartitionFilter()),
		NumBytes:                  dcl.StringOrNil(p.GetNumBytes()),
		NumPhysicalBytes:          dcl.StringOrNil(p.GetNumPhysicalBytes()),
		NumLongTermBytes:          dcl.StringOrNil(p.GetNumLongTermBytes()),
		NumRows:                   dcl.Int64OrNil(p.GetNumRows()),
		CreationTime:              dcl.Int64OrNil(p.GetCreationTime()),
		ExpirationTime:            dcl.Int64OrNil(p.GetExpirationTime()),
		LastModifiedTime:          dcl.Int64OrNil(p.GetLastModifiedTime()),
		Type:                      dcl.StringOrNil(p.GetType()),
		View:                      ProtoToBigqueryTableView(p.GetView()),
		MaterializedView:          ProtoToBigqueryTableMaterializedView(p.GetMaterializedView()),
		ExternalDataConfiguration: ProtoToBigqueryTableExternalDataConfiguration(p.GetExternalDataConfiguration()),
		Location:                  dcl.StringOrNil(p.GetLocation()),
		StreamingBuffer:           ProtoToBigqueryTableStreamingBuffer(p.GetStreamingBuffer()),
		EncryptionConfiguration:   ProtoToBigqueryTableEncryptionConfiguration(p.GetEncryptionConfiguration()),
		SnapshotDefinition:        ProtoToBigqueryTableSnapshotDefinition(p.GetSnapshotDefinition()),
		DefaultCollation:          dcl.StringOrNil(p.GetDefaultCollation()),
	}
	return obj
}

// TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto converts a TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum enum to its proto representation.
func BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto(e *bigquery.TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum) bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum {
	if e == nil {
		return bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(0)
	}
	if v, ok := bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum_value["TableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(v)
	}
	return bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnum(0)
}

// TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto converts a TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum enum to its proto representation.
func BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto(e *bigquery.TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum) bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum {
	if e == nil {
		return bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(0)
	}
	if v, ok := bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum_value["TableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(v)
	}
	return bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnum(0)
}

// TableExternalDataConfigurationDecimalTargetTypesEnumToProto converts a TableExternalDataConfigurationDecimalTargetTypesEnum enum to its proto representation.
func BigqueryTableExternalDataConfigurationDecimalTargetTypesEnumToProto(e *bigquery.TableExternalDataConfigurationDecimalTargetTypesEnum) bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum {
	if e == nil {
		return bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum(0)
	}
	if v, ok := bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum_value["TableExternalDataConfigurationDecimalTargetTypesEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum(v)
	}
	return bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum(0)
}

// TableExternalDataConfigurationJsonExtensionEnumToProto converts a TableExternalDataConfigurationJsonExtensionEnum enum to its proto representation.
func BigqueryTableExternalDataConfigurationJsonExtensionEnumToProto(e *bigquery.TableExternalDataConfigurationJsonExtensionEnum) bigquerypb.BigqueryTableExternalDataConfigurationJsonExtensionEnum {
	if e == nil {
		return bigquerypb.BigqueryTableExternalDataConfigurationJsonExtensionEnum(0)
	}
	if v, ok := bigquerypb.BigqueryTableExternalDataConfigurationJsonExtensionEnum_value["TableExternalDataConfigurationJsonExtensionEnum"+string(*e)]; ok {
		return bigquerypb.BigqueryTableExternalDataConfigurationJsonExtensionEnum(v)
	}
	return bigquerypb.BigqueryTableExternalDataConfigurationJsonExtensionEnum(0)
}

// TableModelToProto converts a TableModel object to its proto representation.
func BigqueryTableModelToProto(o *bigquery.TableModel) *bigquerypb.BigqueryTableModel {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableModel{}
	p.SetModelOptions(BigqueryTableModelModelOptionsToProto(o.ModelOptions))
	sTrainingRuns := make([]*bigquerypb.BigqueryTableModelTrainingRuns, len(o.TrainingRuns))
	for i, r := range o.TrainingRuns {
		sTrainingRuns[i] = BigqueryTableModelTrainingRunsToProto(&r)
	}
	p.SetTrainingRuns(sTrainingRuns)
	return p
}

// TableModelModelOptionsToProto converts a TableModelModelOptions object to its proto representation.
func BigqueryTableModelModelOptionsToProto(o *bigquery.TableModelModelOptions) *bigquerypb.BigqueryTableModelModelOptions {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableModelModelOptions{}
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
func BigqueryTableModelTrainingRunsToProto(o *bigquery.TableModelTrainingRuns) *bigquerypb.BigqueryTableModelTrainingRuns {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableModelTrainingRuns{}
	p.SetState(dcl.ValueOrEmptyString(o.State))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetTrainingOptions(BigqueryTableModelTrainingRunsTrainingOptionsToProto(o.TrainingOptions))
	sIterationResults := make([]*bigquerypb.BigqueryTableModelTrainingRunsIterationResults, len(o.IterationResults))
	for i, r := range o.IterationResults {
		sIterationResults[i] = BigqueryTableModelTrainingRunsIterationResultsToProto(&r)
	}
	p.SetIterationResults(sIterationResults)
	return p
}

// TableModelTrainingRunsTrainingOptionsToProto converts a TableModelTrainingRunsTrainingOptions object to its proto representation.
func BigqueryTableModelTrainingRunsTrainingOptionsToProto(o *bigquery.TableModelTrainingRunsTrainingOptions) *bigquerypb.BigqueryTableModelTrainingRunsTrainingOptions {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableModelTrainingRunsTrainingOptions{}
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
func BigqueryTableModelTrainingRunsIterationResultsToProto(o *bigquery.TableModelTrainingRunsIterationResults) *bigquerypb.BigqueryTableModelTrainingRunsIterationResults {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableModelTrainingRunsIterationResults{}
	p.SetIndex(dcl.ValueOrEmptyInt64(o.Index))
	p.SetLearnRate(dcl.ValueOrEmptyDouble(o.LearnRate))
	p.SetTrainingLoss(dcl.ValueOrEmptyDouble(o.TrainingLoss))
	p.SetEvalLoss(dcl.ValueOrEmptyDouble(o.EvalLoss))
	p.SetDurationMs(dcl.ValueOrEmptyString(o.DurationMs))
	return p
}

// TableSchemaToProto converts a TableSchema object to its proto representation.
func BigqueryTableSchemaToProto(o *bigquery.TableSchema) *bigquerypb.BigqueryTableSchema {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableSchema{}
	sFields := make([]*bigquerypb.BigqueryTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// TableGooglecloudbigqueryv2TablefieldschemaToProto converts a TableGooglecloudbigqueryv2Tablefieldschema object to its proto representation.
func BigqueryTableGooglecloudbigqueryv2TablefieldschemaToProto(o *bigquery.TableGooglecloudbigqueryv2Tablefieldschema) *bigquerypb.BigqueryTableGooglecloudbigqueryv2Tablefieldschema {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableGooglecloudbigqueryv2Tablefieldschema{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetMode(dcl.ValueOrEmptyString(o.Mode))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetCategories(BigqueryTableGooglecloudbigqueryv2TablefieldschemaCategoriesToProto(o.Categories))
	p.SetPolicyTags(BigqueryTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto(o.PolicyTags))
	p.SetMaxLength(dcl.ValueOrEmptyInt64(o.MaxLength))
	p.SetPrecision(dcl.ValueOrEmptyInt64(o.Precision))
	p.SetScale(dcl.ValueOrEmptyInt64(o.Scale))
	p.SetCollation(dcl.ValueOrEmptyString(o.Collation))
	p.SetDefaultValueExpression(dcl.ValueOrEmptyString(o.DefaultValueExpression))
	sFields := make([]*bigquerypb.BigqueryTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
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
func BigqueryTableGooglecloudbigqueryv2TablefieldschemaCategoriesToProto(o *bigquery.TableGooglecloudbigqueryv2TablefieldschemaCategories) *bigquerypb.BigqueryTableGooglecloudbigqueryv2TablefieldschemaCategories {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableGooglecloudbigqueryv2TablefieldschemaCategories{}
	sNames := make([]string, len(o.Names))
	for i, r := range o.Names {
		sNames[i] = r
	}
	p.SetNames(sNames)
	return p
}

// TableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto converts a TableGooglecloudbigqueryv2TablefieldschemaPolicyTags object to its proto representation.
func BigqueryTableGooglecloudbigqueryv2TablefieldschemaPolicyTagsToProto(o *bigquery.TableGooglecloudbigqueryv2TablefieldschemaPolicyTags) *bigquerypb.BigqueryTableGooglecloudbigqueryv2TablefieldschemaPolicyTags {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableGooglecloudbigqueryv2TablefieldschemaPolicyTags{}
	sNames := make([]string, len(o.Names))
	for i, r := range o.Names {
		sNames[i] = r
	}
	p.SetNames(sNames)
	return p
}

// TableTimePartitioningToProto converts a TableTimePartitioning object to its proto representation.
func BigqueryTableTimePartitioningToProto(o *bigquery.TableTimePartitioning) *bigquerypb.BigqueryTableTimePartitioning {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableTimePartitioning{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetExpirationMs(dcl.ValueOrEmptyString(o.ExpirationMs))
	p.SetField(dcl.ValueOrEmptyString(o.Field))
	return p
}

// TableRangePartitioningToProto converts a TableRangePartitioning object to its proto representation.
func BigqueryTableRangePartitioningToProto(o *bigquery.TableRangePartitioning) *bigquerypb.BigqueryTableRangePartitioning {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableRangePartitioning{}
	p.SetField(dcl.ValueOrEmptyString(o.Field))
	p.SetRange(BigqueryTableRangePartitioningRangeToProto(o.Range))
	return p
}

// TableRangePartitioningRangeToProto converts a TableRangePartitioningRange object to its proto representation.
func BigqueryTableRangePartitioningRangeToProto(o *bigquery.TableRangePartitioningRange) *bigquerypb.BigqueryTableRangePartitioningRange {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableRangePartitioningRange{}
	p.SetStart(dcl.ValueOrEmptyString(o.Start))
	p.SetEnd(dcl.ValueOrEmptyString(o.End))
	p.SetInterval(dcl.ValueOrEmptyString(o.Interval))
	return p
}

// TableClusteringToProto converts a TableClustering object to its proto representation.
func BigqueryTableClusteringToProto(o *bigquery.TableClustering) *bigquerypb.BigqueryTableClustering {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableClustering{}
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// TableViewToProto converts a TableView object to its proto representation.
func BigqueryTableViewToProto(o *bigquery.TableView) *bigquerypb.BigqueryTableView {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableView{}
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	p.SetUseLegacySql(dcl.ValueOrEmptyBool(o.UseLegacySql))
	p.SetUseExplicitColumnNames(dcl.ValueOrEmptyBool(o.UseExplicitColumnNames))
	sUserDefinedFunctionResources := make([]*bigquerypb.BigqueryTableViewUserDefinedFunctionResources, len(o.UserDefinedFunctionResources))
	for i, r := range o.UserDefinedFunctionResources {
		sUserDefinedFunctionResources[i] = BigqueryTableViewUserDefinedFunctionResourcesToProto(&r)
	}
	p.SetUserDefinedFunctionResources(sUserDefinedFunctionResources)
	return p
}

// TableViewUserDefinedFunctionResourcesToProto converts a TableViewUserDefinedFunctionResources object to its proto representation.
func BigqueryTableViewUserDefinedFunctionResourcesToProto(o *bigquery.TableViewUserDefinedFunctionResources) *bigquerypb.BigqueryTableViewUserDefinedFunctionResources {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableViewUserDefinedFunctionResources{}
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
func BigqueryTableMaterializedViewToProto(o *bigquery.TableMaterializedView) *bigquerypb.BigqueryTableMaterializedView {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableMaterializedView{}
	p.SetQuery(dcl.ValueOrEmptyString(o.Query))
	p.SetLastRefreshTime(dcl.ValueOrEmptyInt64(o.LastRefreshTime))
	p.SetEnableRefresh(dcl.ValueOrEmptyBool(o.EnableRefresh))
	p.SetRefreshIntervalMs(dcl.ValueOrEmptyInt64(o.RefreshIntervalMs))
	return p
}

// TableExternalDataConfigurationToProto converts a TableExternalDataConfiguration object to its proto representation.
func BigqueryTableExternalDataConfigurationToProto(o *bigquery.TableExternalDataConfiguration) *bigquerypb.BigqueryTableExternalDataConfiguration {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfiguration{}
	p.SetSchema(BigqueryTableExternalDataConfigurationSchemaToProto(o.Schema))
	p.SetSourceFormat(dcl.ValueOrEmptyString(o.SourceFormat))
	p.SetMaxBadRecords(dcl.ValueOrEmptyInt64(o.MaxBadRecords))
	p.SetAutodetect(dcl.ValueOrEmptyBool(o.Autodetect))
	p.SetIgnoreUnknownValues(dcl.ValueOrEmptyBool(o.IgnoreUnknownValues))
	p.SetCompression(dcl.ValueOrEmptyString(o.Compression))
	p.SetCsvOptions(BigqueryTableExternalDataConfigurationCsvOptionsToProto(o.CsvOptions))
	p.SetBigtableOptions(BigqueryTableExternalDataConfigurationBigtableOptionsToProto(o.BigtableOptions))
	p.SetGoogleSheetsOptions(BigqueryTableExternalDataConfigurationGoogleSheetsOptionsToProto(o.GoogleSheetsOptions))
	p.SetHivePartitioningOptions(BigqueryTableExternalDataConfigurationHivePartitioningOptionsToProto(o.HivePartitioningOptions))
	p.SetConnectionId(dcl.ValueOrEmptyString(o.ConnectionId))
	p.SetValueConversionModes(BigqueryTableExternalDataConfigurationValueConversionModesToProto(o.ValueConversionModes))
	p.SetAvroOptions(BigqueryTableExternalDataConfigurationAvroOptionsToProto(o.AvroOptions))
	p.SetJsonExtension(BigqueryTableExternalDataConfigurationJsonExtensionEnumToProto(o.JsonExtension))
	p.SetParquetOptions(BigqueryTableExternalDataConfigurationParquetOptionsToProto(o.ParquetOptions))
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
	sDecimalTargetTypes := make([]bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum, len(o.DecimalTargetTypes))
	for i, r := range o.DecimalTargetTypes {
		sDecimalTargetTypes[i] = bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum(bigquerypb.BigqueryTableExternalDataConfigurationDecimalTargetTypesEnum_value[string(r)])
	}
	p.SetDecimalTargetTypes(sDecimalTargetTypes)
	return p
}

// TableExternalDataConfigurationSchemaToProto converts a TableExternalDataConfigurationSchema object to its proto representation.
func BigqueryTableExternalDataConfigurationSchemaToProto(o *bigquery.TableExternalDataConfigurationSchema) *bigquerypb.BigqueryTableExternalDataConfigurationSchema {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationSchema{}
	sFields := make([]*bigquerypb.BigqueryTableGooglecloudbigqueryv2Tablefieldschema, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = BigqueryTableGooglecloudbigqueryv2TablefieldschemaToProto(&r)
	}
	p.SetFields(sFields)
	return p
}

// TableExternalDataConfigurationCsvOptionsToProto converts a TableExternalDataConfigurationCsvOptions object to its proto representation.
func BigqueryTableExternalDataConfigurationCsvOptionsToProto(o *bigquery.TableExternalDataConfigurationCsvOptions) *bigquerypb.BigqueryTableExternalDataConfigurationCsvOptions {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationCsvOptions{}
	p.SetFieldDelimiter(dcl.ValueOrEmptyString(o.FieldDelimiter))
	p.SetSkipLeadingRows(dcl.ValueOrEmptyString(o.SkipLeadingRows))
	p.SetQuote(dcl.ValueOrEmptyString(o.Quote))
	p.SetAllowQuotedNewlines(dcl.ValueOrEmptyBool(o.AllowQuotedNewlines))
	p.SetAllowJaggedRows(dcl.ValueOrEmptyBool(o.AllowJaggedRows))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	return p
}

// TableExternalDataConfigurationBigtableOptionsToProto converts a TableExternalDataConfigurationBigtableOptions object to its proto representation.
func BigqueryTableExternalDataConfigurationBigtableOptionsToProto(o *bigquery.TableExternalDataConfigurationBigtableOptions) *bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptions {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptions{}
	p.SetIgnoreUnspecifiedColumnFamilies(dcl.ValueOrEmptyBool(o.IgnoreUnspecifiedColumnFamilies))
	p.SetReadRowkeyAsString(dcl.ValueOrEmptyBool(o.ReadRowkeyAsString))
	sColumnFamilies := make([]*bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilies, len(o.ColumnFamilies))
	for i, r := range o.ColumnFamilies {
		sColumnFamilies[i] = BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto(&r)
	}
	p.SetColumnFamilies(sColumnFamilies)
	return p
}

// TableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto converts a TableExternalDataConfigurationBigtableOptionsColumnFamilies object to its proto representation.
func BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesToProto(o *bigquery.TableExternalDataConfigurationBigtableOptionsColumnFamilies) *bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilies {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilies{}
	p.SetFamilyId(dcl.ValueOrEmptyString(o.FamilyId))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetOnlyReadLatest(dcl.ValueOrEmptyBool(o.OnlyReadLatest))
	sColumns := make([]*bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns, len(o.Columns))
	for i, r := range o.Columns {
		sColumns[i] = BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto(&r)
	}
	p.SetColumns(sColumns)
	return p
}

// TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto converts a TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns object to its proto representation.
func BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumnsToProto(o *bigquery.TableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns) *bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamiliesColumns{}
	p.SetQualifierEncoded(dcl.ValueOrEmptyString(o.QualifierEncoded))
	p.SetQualifierString(dcl.ValueOrEmptyString(o.QualifierString))
	p.SetFieldName(dcl.ValueOrEmptyString(o.FieldName))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetEncoding(dcl.ValueOrEmptyString(o.Encoding))
	p.SetOnlyReadLatest(dcl.ValueOrEmptyBool(o.OnlyReadLatest))
	return p
}

// TableExternalDataConfigurationGoogleSheetsOptionsToProto converts a TableExternalDataConfigurationGoogleSheetsOptions object to its proto representation.
func BigqueryTableExternalDataConfigurationGoogleSheetsOptionsToProto(o *bigquery.TableExternalDataConfigurationGoogleSheetsOptions) *bigquerypb.BigqueryTableExternalDataConfigurationGoogleSheetsOptions {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationGoogleSheetsOptions{}
	p.SetSkipLeadingRows(dcl.ValueOrEmptyString(o.SkipLeadingRows))
	p.SetRange(dcl.ValueOrEmptyString(o.Range))
	return p
}

// TableExternalDataConfigurationHivePartitioningOptionsToProto converts a TableExternalDataConfigurationHivePartitioningOptions object to its proto representation.
func BigqueryTableExternalDataConfigurationHivePartitioningOptionsToProto(o *bigquery.TableExternalDataConfigurationHivePartitioningOptions) *bigquerypb.BigqueryTableExternalDataConfigurationHivePartitioningOptions {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationHivePartitioningOptions{}
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
func BigqueryTableExternalDataConfigurationValueConversionModesToProto(o *bigquery.TableExternalDataConfigurationValueConversionModes) *bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModes {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationValueConversionModes{}
	p.SetTemporalTypesOutOfRangeConversionMode(BigqueryTableExternalDataConfigurationValueConversionModesTemporalTypesOutOfRangeConversionModeEnumToProto(o.TemporalTypesOutOfRangeConversionMode))
	p.SetNumericTypeOutOfRangeConversionMode(BigqueryTableExternalDataConfigurationValueConversionModesNumericTypeOutOfRangeConversionModeEnumToProto(o.NumericTypeOutOfRangeConversionMode))
	return p
}

// TableExternalDataConfigurationAvroOptionsToProto converts a TableExternalDataConfigurationAvroOptions object to its proto representation.
func BigqueryTableExternalDataConfigurationAvroOptionsToProto(o *bigquery.TableExternalDataConfigurationAvroOptions) *bigquerypb.BigqueryTableExternalDataConfigurationAvroOptions {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationAvroOptions{}
	p.SetUseAvroLogicalTypes(dcl.ValueOrEmptyBool(o.UseAvroLogicalTypes))
	return p
}

// TableExternalDataConfigurationParquetOptionsToProto converts a TableExternalDataConfigurationParquetOptions object to its proto representation.
func BigqueryTableExternalDataConfigurationParquetOptionsToProto(o *bigquery.TableExternalDataConfigurationParquetOptions) *bigquerypb.BigqueryTableExternalDataConfigurationParquetOptions {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableExternalDataConfigurationParquetOptions{}
	p.SetEnumAsString(dcl.ValueOrEmptyBool(o.EnumAsString))
	p.SetEnableListInference(dcl.ValueOrEmptyBool(o.EnableListInference))
	return p
}

// TableStreamingBufferToProto converts a TableStreamingBuffer object to its proto representation.
func BigqueryTableStreamingBufferToProto(o *bigquery.TableStreamingBuffer) *bigquerypb.BigqueryTableStreamingBuffer {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableStreamingBuffer{}
	p.SetEstimatedBytes(dcl.ValueOrEmptyInt64(o.EstimatedBytes))
	p.SetEstimatedRows(dcl.ValueOrEmptyInt64(o.EstimatedRows))
	p.SetOldestEntryTime(dcl.ValueOrEmptyInt64(o.OldestEntryTime))
	return p
}

// TableEncryptionConfigurationToProto converts a TableEncryptionConfiguration object to its proto representation.
func BigqueryTableEncryptionConfigurationToProto(o *bigquery.TableEncryptionConfiguration) *bigquerypb.BigqueryTableEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableEncryptionConfiguration{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// TableSnapshotDefinitionToProto converts a TableSnapshotDefinition object to its proto representation.
func BigqueryTableSnapshotDefinitionToProto(o *bigquery.TableSnapshotDefinition) *bigquerypb.BigqueryTableSnapshotDefinition {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryTableSnapshotDefinition{}
	p.SetTable(dcl.ValueOrEmptyString(o.Table))
	p.SetDataset(dcl.ValueOrEmptyString(o.Dataset))
	p.SetProject(dcl.ValueOrEmptyString(o.Project))
	p.SetSnapshotTime(dcl.ValueOrEmptyString(o.SnapshotTime))
	return p
}

// TableToProto converts a Table resource to its proto representation.
func TableToProto(resource *bigquery.Table) *bigquerypb.BigqueryTable {
	p := &bigquerypb.BigqueryTable{}
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetId(dcl.ValueOrEmptyString(resource.Id))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetFriendlyName(dcl.ValueOrEmptyString(resource.FriendlyName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetModel(BigqueryTableModelToProto(resource.Model))
	p.SetSchema(BigqueryTableSchemaToProto(resource.Schema))
	p.SetTimePartitioning(BigqueryTableTimePartitioningToProto(resource.TimePartitioning))
	p.SetRangePartitioning(BigqueryTableRangePartitioningToProto(resource.RangePartitioning))
	p.SetClustering(BigqueryTableClusteringToProto(resource.Clustering))
	p.SetRequirePartitionFilter(dcl.ValueOrEmptyBool(resource.RequirePartitionFilter))
	p.SetNumBytes(dcl.ValueOrEmptyString(resource.NumBytes))
	p.SetNumPhysicalBytes(dcl.ValueOrEmptyString(resource.NumPhysicalBytes))
	p.SetNumLongTermBytes(dcl.ValueOrEmptyString(resource.NumLongTermBytes))
	p.SetNumRows(dcl.ValueOrEmptyInt64(resource.NumRows))
	p.SetCreationTime(dcl.ValueOrEmptyInt64(resource.CreationTime))
	p.SetExpirationTime(dcl.ValueOrEmptyInt64(resource.ExpirationTime))
	p.SetLastModifiedTime(dcl.ValueOrEmptyInt64(resource.LastModifiedTime))
	p.SetType(dcl.ValueOrEmptyString(resource.Type))
	p.SetView(BigqueryTableViewToProto(resource.View))
	p.SetMaterializedView(BigqueryTableMaterializedViewToProto(resource.MaterializedView))
	p.SetExternalDataConfiguration(BigqueryTableExternalDataConfigurationToProto(resource.ExternalDataConfiguration))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetStreamingBuffer(BigqueryTableStreamingBufferToProto(resource.StreamingBuffer))
	p.SetEncryptionConfiguration(BigqueryTableEncryptionConfigurationToProto(resource.EncryptionConfiguration))
	p.SetSnapshotDefinition(BigqueryTableSnapshotDefinitionToProto(resource.SnapshotDefinition))
	p.SetDefaultCollation(dcl.ValueOrEmptyString(resource.DefaultCollation))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyTable handles the gRPC request by passing it to the underlying Table Apply() method.
func (s *TableServer) applyTable(ctx context.Context, c *bigquery.Client, request *bigquerypb.ApplyBigqueryTableRequest) (*bigquerypb.BigqueryTable, error) {
	p := ProtoToTable(request.GetResource())
	res, err := c.ApplyTable(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TableToProto(res)
	return r, nil
}

// applyBigqueryTable handles the gRPC request by passing it to the underlying Table Apply() method.
func (s *TableServer) ApplyBigqueryTable(ctx context.Context, request *bigquerypb.ApplyBigqueryTableRequest) (*bigquerypb.BigqueryTable, error) {
	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTable(ctx, cl, request)
}

// DeleteTable handles the gRPC request by passing it to the underlying Table Delete() method.
func (s *TableServer) DeleteBigqueryTable(ctx context.Context, request *bigquerypb.DeleteBigqueryTableRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTable(ctx, ProtoToTable(request.GetResource()))

}

// ListBigqueryTable handles the gRPC request by passing it to the underlying TableList() method.
func (s *TableServer) ListBigqueryTable(ctx context.Context, request *bigquerypb.ListBigqueryTableRequest) (*bigquerypb.ListBigqueryTableResponse, error) {
	cl, err := createConfigTable(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTable(ctx, request.GetProject(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*bigquerypb.BigqueryTable
	for _, r := range resources.Items {
		rp := TableToProto(r)
		protos = append(protos, rp)
	}
	p := &bigquerypb.ListBigqueryTableResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTable(ctx context.Context, service_account_file string) (*bigquery.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigquery.NewClient(conf), nil
}
