// Copyright 2024 Google LLC
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

package bigquery

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	bigquery "google.golang.org/api/bigquery/v2"
)

func checkFieldValid(fields []*bigquery.TableFieldSchema) error {
	for _, field := range fields {
		if field == nil {
			return fmt.Errorf("schema contains nil field")
		}
		if field.Name == "" {
			return fmt.Errorf("field contains empty field name")
		}
	}
	return nil
}

func policyTagsEqual(a, b *bigquery.TableFieldSchemaPolicyTags) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Names == nil && b.Names == nil {
		return true
	}
	// If one of a.Names or b.Names is nil.
	if a.Names == nil || b.Names == nil {
		// Suppress nil string and emptry string different.
		if len(a.Names) == len(b.Names) {
			return true
		}
		return false
	}
	if len(a.Names) != len(b.Names) {
		return false
	}
	aNames := make([]string, len(a.Names))
	copy(aNames, a.Names)
	bNames := make([]string, len(b.Names))
	copy(bNames, b.Names)
	sort.Strings(aNames)
	sort.Strings(bNames)
	for i := range aNames {
		if aNames[i] != bNames[i] {
			return false
		}
	}
	return true
}

// Sort the fields in place by name.
func sortSchemaFields(fields []*bigquery.TableFieldSchema) {
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name < fields[j].Name // For ascending order
	})
}

func tableFieldsSchemaEqual(desired, actual []*bigquery.TableFieldSchema, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if len(desired) == 0 && len(actual) == 0 {
		return true, nil
	}
	if err := checkFieldValid(desired); err != nil {
		return false, err
	}
	if err := checkFieldValid(actual); err != nil {
		return false, err

	}
	if len(desired) != len(actual) {
		diff.AddField(prefix, desired, actual)
		return false, nil
	}

	desiredCopy := make([]*bigquery.TableFieldSchema, len(desired))
	copy(desiredCopy, desired)
	desired = desiredCopy

	actualCopy := make([]*bigquery.TableFieldSchema, len(actual))
	copy(actualCopy, actual)
	actual = actualCopy

	// The fields from API can be in a different order.
	// Sort by name before comparing.
	sortSchemaFields(desired)
	sortSchemaFields(actual)
	equal := true
	for i := range desired {
		fieldName := desired[i].Name
		fieldPrefix := fmt.Sprintf("%s[%s]", prefix, fieldName)
		if !reflect.DeepEqual(desired[i].Categories, actual[i].Categories) {
			diff.AddField(fieldPrefix+".categories", desired[i].Categories, actual[i].Categories)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].Collation, actual[i].Collation) {
			diff.AddField(fieldPrefix+".collation", desired[i].Collation, actual[i].Collation)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].DefaultValueExpression, actual[i].DefaultValueExpression) {
			diff.AddField(fieldPrefix+".default_value_expression", desired[i].DefaultValueExpression, actual[i].DefaultValueExpression)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].Description, actual[i].Description) {
			diff.AddField(fieldPrefix+".description", desired[i].Description, actual[i].Description)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].ForeignTypeDefinition, actual[i].ForeignTypeDefinition) {
			diff.AddField(fieldPrefix+".foreign_type_definition", desired[i].ForeignTypeDefinition, actual[i].ForeignTypeDefinition)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].MaxLength, actual[i].MaxLength) {
			diff.AddField(fieldPrefix+".max_length", desired[i].MaxLength, actual[i].MaxLength)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].Mode, actual[i].Mode) {
			diff.AddField(fieldPrefix+".mode", desired[i].Mode, actual[i].Mode)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].Name, actual[i].Name) {
			diff.AddField(fieldPrefix+".name", desired[i].Name, actual[i].Name)
			equal = false
		}
		if !policyTagsEqual(desired[i].PolicyTags, actual[i].PolicyTags) {
			diff.AddField(fieldPrefix+".policy_tags", desired[i].PolicyTags, actual[i].PolicyTags)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].Precision, actual[i].Precision) {
			diff.AddField(fieldPrefix+".precision", desired[i].Precision, actual[i].Precision)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].RangeElementType, actual[i].RangeElementType) {
			diff.AddField(fieldPrefix+".range_element_type", desired[i].RangeElementType, actual[i].RangeElementType)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].RoundingMode, actual[i].RoundingMode) {
			diff.AddField(fieldPrefix+".rounding_mode", desired[i].RoundingMode, actual[i].RoundingMode)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].Scale, actual[i].Scale) {
			diff.AddField(fieldPrefix+".scale", desired[i].Scale, actual[i].Scale)
			equal = false
		}
		if !reflect.DeepEqual(desired[i].Type, actual[i].Type) {
			diff.AddField(fieldPrefix+".type", desired[i].Type, actual[i].Type)
			equal = false
		}
		fieldsEqual, err := tableFieldsSchemaEqual(desired[i].Fields, actual[i].Fields, fieldPrefix+".fields", diff)
		if err != nil {
			return false, err
		}
		if !fieldsEqual {
			equal = false
		}
	}
	return equal, nil
}

func avroOptionsEq(a, b *bigquery.AvroOptions) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.UseAvroLogicalTypes == b.UseAvroLogicalTypes
}

func clusteringEq(a, b *bigquery.Clustering) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return reflect.DeepEqual(a.Fields, b.Fields)
}

func csvOptionsEq(a, b *bigquery.CsvOptions) bool {
	if a == nil && b == nil {
		return true
	}
	// Defaults from mapper
	aFieldDelimiter := ","
	if a != nil && a.FieldDelimiter != "" {
		aFieldDelimiter = a.FieldDelimiter
	}
	bFieldDelimiter := ","
	if b != nil && b.FieldDelimiter != "" {
		bFieldDelimiter = b.FieldDelimiter
	}
	if aFieldDelimiter != bFieldDelimiter {
		return false
	}

	aSkipLeadingRows := int64(0)
	if a != nil {
		aSkipLeadingRows = a.SkipLeadingRows
	}
	bSkipLeadingRows := int64(0)
	if b != nil {
		bSkipLeadingRows = b.SkipLeadingRows
	}
	if aSkipLeadingRows != bSkipLeadingRows {
		return false
	}

	aQuote := ""
	if a != nil && a.Quote != nil {
		aQuote = *a.Quote
	}
	bQuote := ""
	if b != nil && b.Quote != nil {
		bQuote = *b.Quote
	}
	if aQuote != bQuote {
		return false
	}

	aAllowQuotedNewlines := false
	if a != nil {
		aAllowQuotedNewlines = a.AllowQuotedNewlines
	}
	bAllowQuotedNewlines := false
	if b != nil {
		bAllowQuotedNewlines = b.AllowQuotedNewlines
	}
	if aAllowQuotedNewlines != bAllowQuotedNewlines {
		return false
	}

	aAllowJaggedRows := false
	if a != nil {
		aAllowJaggedRows = a.AllowJaggedRows
	}
	bAllowJaggedRows := false
	if b != nil {
		bAllowJaggedRows = b.AllowJaggedRows
	}
	if aAllowJaggedRows != bAllowJaggedRows {
		return false
	}

	aEncoding := "UTF8"
	if a != nil && a.Encoding != "" {
		aEncoding = a.Encoding
	}
	bEncoding := "UTF8"
	if b != nil && b.Encoding != "" {
		bEncoding = b.Encoding
	}
	if aEncoding != bEncoding {
		return false
	}

	return true
}

func encryptionConfigurationEq(a, b *bigquery.EncryptionConfiguration) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b.KmsKeyName == ""
	}
	if b == nil {
		return a.KmsKeyName == ""
	}
	return a.KmsKeyName == b.KmsKeyName
}

func googleSheetsOptionsEq(a, b *bigquery.GoogleSheetsOptions) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Range == b.Range && a.SkipLeadingRows == b.SkipLeadingRows
}

func hivePartitioningOptionsEq(a, b *bigquery.HivePartitioningOptions) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Mode == b.Mode && a.RequirePartitionFilter == b.RequirePartitionFilter
}

func jsonOptionsEq(a, b *bigquery.JsonOptions) bool {
	if a == nil && b == nil {
		return true
	}
	aEncoding := "UTF8"
	if a != nil && a.Encoding != "" {
		aEncoding = a.Encoding
	}
	bEncoding := "UTF8"
	if b != nil && b.Encoding != "" {
		bEncoding = b.Encoding
	}
	return aEncoding == bEncoding
}

func parquetOptionsEq(a, b *bigquery.ParquetOptions) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.EnumAsString == b.EnumAsString && a.EnableListInference == b.EnableListInference
}

func rangePartitioningEq(a, b *bigquery.RangePartitioning) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Field != b.Field {
		return false
	}
	if a.Range == nil && b.Range == nil {
		return true
	}
	if a.Range == nil || b.Range == nil {
		return false
	}
	return a.Range.End == b.Range.End && a.Range.Interval == b.Range.Interval && a.Range.Start == b.Range.Start
}

func externalDataConfigurationEqual(a, b *bigquery.ExternalDataConfiguration, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false, nil
	}

	equal := true
	if a.Autodetect != b.Autodetect {
		diff.AddField(prefix+".autodetect", a.Autodetect, b.Autodetect)
		equal = false
	}
	if !avroOptionsEq(a.AvroOptions, b.AvroOptions) {
		diff.AddField(prefix+".avro_options", a.AvroOptions, b.AvroOptions)
		equal = false
	}

	aCompression := "NONE"
	if a.Compression != "" {
		aCompression = a.Compression
	}
	bCompression := "NONE"
	if b.Compression != "" {
		bCompression = b.Compression
	}
	if aCompression != bCompression {
		diff.AddField(prefix+".compression", a.Compression, b.Compression)
		equal = false
	}

	if a.ConnectionId != b.ConnectionId {
		diff.AddField(prefix+".connection_id", a.ConnectionId, b.ConnectionId)
		equal = false
	}
	if !csvOptionsEq(a.CsvOptions, b.CsvOptions) {
		diff.AddField(prefix+".csv_options", a.CsvOptions, b.CsvOptions)
		equal = false
	}
	if !googleSheetsOptionsEq(a.GoogleSheetsOptions, b.GoogleSheetsOptions) {
		diff.AddField(prefix+".google_sheets_options", a.GoogleSheetsOptions, b.GoogleSheetsOptions)
		equal = false
	}
	if !hivePartitioningOptionsEq(a.HivePartitioningOptions, b.HivePartitioningOptions) {
		diff.AddField(prefix+".hive_partitioning_options", a.HivePartitioningOptions, b.HivePartitioningOptions)
		equal = false
	}
	if a.IgnoreUnknownValues != b.IgnoreUnknownValues {
		diff.AddField(prefix+".ignore_unknown_values", a.IgnoreUnknownValues, b.IgnoreUnknownValues)
		equal = false
	}
	if !jsonOptionsEq(a.JsonOptions, b.JsonOptions) {
		diff.AddField(prefix+".json_options", a.JsonOptions, b.JsonOptions)
		equal = false
	}
	if a.MaxBadRecords != b.MaxBadRecords {
		diff.AddField(prefix+".max_bad_records", a.MaxBadRecords, b.MaxBadRecords)
		equal = false
	}
	if a.MetadataCacheMode != b.MetadataCacheMode {
		diff.AddField(prefix+".metadata_cache_mode", a.MetadataCacheMode, b.MetadataCacheMode)
		equal = false
	}
	if !parquetOptionsEq(a.ParquetOptions, b.ParquetOptions) {
		diff.AddField(prefix+".parquet_options", a.ParquetOptions, b.ParquetOptions)
		equal = false
	}
	if a.SourceFormat != b.SourceFormat {
		diff.AddField(prefix+".source_format", a.SourceFormat, b.SourceFormat)
		equal = false
	}
	if !reflect.DeepEqual(a.SourceUris, b.SourceUris) {
		diff.AddField(prefix+".source_uris", a.SourceUris, b.SourceUris)
		equal = false
	}

	schemaEqual, err := tableSchemaEq(a.Schema, b.Schema, prefix+".schema", diff)
	if err != nil {
		return false, err
	}
	if !schemaEqual {
		equal = false
	}
	return equal, nil
}

func materializedViewEq(a, b *bigquery.MaterializedViewDefinition, prefix string, diff *structuredreporting.Diff) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false
	}
	equal := true
	if a.AllowNonIncrementalDefinition != b.AllowNonIncrementalDefinition {
		diff.AddField(prefix+".allow_non_incremental_definition", a.AllowNonIncrementalDefinition, b.AllowNonIncrementalDefinition)
		equal = false
	}

	if a.EnableRefresh != b.EnableRefresh {
		diff.AddField(prefix+".enable_refresh", a.EnableRefresh, b.EnableRefresh)
		equal = false
	}
	if a.Query != b.Query {
		diff.AddField(prefix+".query", a.Query, b.Query)
		equal = false
	}
	if a.MaxStaleness != b.MaxStaleness {
		diff.AddField(prefix+".max_staleness", a.MaxStaleness, b.MaxStaleness)
		equal = false
	}

	aRefreshIntervalMs := int64(1800000)
	if a.RefreshIntervalMs != 0 {
		aRefreshIntervalMs = a.RefreshIntervalMs
	}
	bRefreshIntervalMs := int64(1800000)
	if b.RefreshIntervalMs != 0 {
		bRefreshIntervalMs = b.RefreshIntervalMs
	}
	if aRefreshIntervalMs != bRefreshIntervalMs {
		diff.AddField(prefix+".refresh_interval_ms", a.RefreshIntervalMs, b.RefreshIntervalMs)
		equal = false
	}
	return equal
}

func tableSchemaEq(a, b *bigquery.TableSchema, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false, nil
	}
	return tableFieldsSchemaEqual(a.Fields, b.Fields, prefix+".fields", diff)
}

func viewEq(a, b *bigquery.ViewDefinition, prefix string, diff *structuredreporting.Diff) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false
	}
	equal := true
	if a.Query != b.Query {
		diff.AddField(prefix+".query", a.Query, b.Query)
		equal = false
	}
	if a.UseLegacySql != b.UseLegacySql {
		diff.AddField(prefix+".use_legacy_sql", a.UseLegacySql, b.UseLegacySql)
		equal = false
	}
	return equal
}

func labelsEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || v != bv {
			return false
		}
	}
	return true
}

func tableConstraintsEq(a, b *bigquery.TableConstraints, prefix string, diff *structuredreporting.Diff) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false
	}
	equal := true
	if !reflect.DeepEqual(a.PrimaryKey, b.PrimaryKey) {
		diff.AddField(prefix+".primary_key", a.PrimaryKey, b.PrimaryKey)
		equal = false
	}

	if len(a.ForeignKeys) != len(b.ForeignKeys) {
		diff.AddField(prefix+".foreign_keys", a.ForeignKeys, b.ForeignKeys)
		return false
	}

	aForeignKeys := make([]*bigquery.TableConstraintsForeignKeys, len(a.ForeignKeys))
	copy(aForeignKeys, a.ForeignKeys)
	sort.Slice(aForeignKeys, func(i, j int) bool {
		return aForeignKeys[i].Name < aForeignKeys[j].Name
	})

	bForeignKeys := make([]*bigquery.TableConstraintsForeignKeys, len(b.ForeignKeys))
	copy(bForeignKeys, b.ForeignKeys)
	sort.Slice(bForeignKeys, func(i, j int) bool {
		return bForeignKeys[i].Name < bForeignKeys[j].Name
	})

	for i := range aForeignKeys {
		if !reflect.DeepEqual(aForeignKeys[i], bForeignKeys[i]) {
			diff.AddField(fmt.Sprintf("%s.foreign_keys[%d]", prefix, i), aForeignKeys[i], bForeignKeys[i])
			equal = false
		}
	}

	return equal
}

func timePartitioningEq(a, b *bigquery.TimePartitioning, prefix string, diff *structuredreporting.Diff) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false
	}
	equal := true
	if a.ExpirationMs != b.ExpirationMs {
		diff.AddField(prefix+".expiration_ms", a.ExpirationMs, b.ExpirationMs)
		equal = false
	}
	if a.Field != b.Field {
		diff.AddField(prefix+".field", a.Field, b.Field)
		equal = false
	}
	if a.RequirePartitionFilter != b.RequirePartitionFilter {
		diff.AddField(prefix+".require_partition_filter", a.RequirePartitionFilter, b.RequirePartitionFilter)
		equal = false
	}
	if a.Type != b.Type {
		diff.AddField(prefix+".type", a.Type, b.Type)
		equal = false
	}
	return equal
}

func TableEq(a, b *bigquery.Table, diff *structuredreporting.Diff) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}
	if a == nil || b == nil {
		diff.AddField("table", a, b)
		return false, nil
	}

	equal := true
	if !clusteringEq(a.Clustering, b.Clustering) {
		diff.AddField("clustering", a.Clustering, b.Clustering)
		equal = false
	}
	if a.Description != b.Description {
		diff.AddField("description", a.Description, b.Description)
		equal = false
	}
	if !encryptionConfigurationEq(a.EncryptionConfiguration, b.EncryptionConfiguration) {
		diff.AddField("encryption_configuration", a.EncryptionConfiguration, b.EncryptionConfiguration)
		equal = false
	}
	if a.ExpirationTime != b.ExpirationTime {
		diff.AddField("expiration_time", a.ExpirationTime, b.ExpirationTime)
		equal = false
	}
	if a.FriendlyName != b.FriendlyName {
		diff.AddField("friendly_name", a.FriendlyName, b.FriendlyName)
		equal = false
	}
	if !materializedViewEq(a.MaterializedView, b.MaterializedView, "materialized_view", diff) {
		equal = false
	}
	if a.MaxStaleness != b.MaxStaleness {
		diff.AddField("max_staleness", a.MaxStaleness, b.MaxStaleness)
		equal = false
	}
	if !rangePartitioningEq(a.RangePartitioning, b.RangePartitioning) {
		diff.AddField("range_partitioning", a.RangePartitioning, b.RangePartitioning)
		equal = false
	}
	if a.RequirePartitionFilter != b.RequirePartitionFilter {
		diff.AddField("require_partition_filter", a.RequirePartitionFilter, b.RequirePartitionFilter)
		equal = false
	}
	if !tableConstraintsEq(a.TableConstraints, b.TableConstraints, "table_constraints", diff) {
		equal = false
	}
	if !timePartitioningEq(a.TimePartitioning, b.TimePartitioning, "time_partitioning", diff) {
		equal = false
	}

	if !viewEq(a.View, b.View, "view", diff) {
		equal = false
	}
	if !labelsEqual(a.Labels, b.Labels) {
		diff.AddField("labels", a.Labels, b.Labels)
		equal = false
	}
	schemaEqual, err := tableSchemaEq(a.Schema, b.Schema, "schema", diff)
	if err != nil {
		return false, err
	}
	if !schemaEqual {
		equal = false
	}
	externalDataEqual, err := externalDataConfigurationEqual(a.ExternalDataConfiguration, b.ExternalDataConfiguration, "external_data_configuration", diff)
	if err != nil {
		return false, err
	}
	if !externalDataEqual {
		equal = false
	}
	return equal, nil
}

