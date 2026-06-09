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

func policyTagsEqual(actual, desired *bigquery.TableFieldSchemaPolicyTags) bool {
	actualEmpty := (actual == nil || len(actual.Names) == 0)
	desiredEmpty := (desired == nil || len(desired.Names) == 0)
	if actualEmpty && desiredEmpty {
		return true
	}
	if actualEmpty || desiredEmpty {
		return false
	}

	if len(actual.Names) != len(desired.Names) {
		return false
	}
	sort.Strings(actual.Names)
	sort.Strings(desired.Names)
	for i := range actual.Names {
		if actual.Names[i] != desired.Names[i] {
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

func tableFieldsSchemaEqual(actual, desired []*bigquery.TableFieldSchema, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if len(actual) == 0 && len(desired) == 0 {
		return true, nil
	}
	if err := checkFieldValid(actual); err != nil {
		return false, err
	}
	if err := checkFieldValid(desired); err != nil {
		return false, err
	}
	if len(actual) != len(desired) {
		diff.AddField(prefix, actual, desired)
		return false, nil
	}
	// The fields from API can be in a different order.
	// Sort by name before comparing.
	sortSchemaFields(actual)
	sortSchemaFields(desired)
	for i := range actual {
		fieldName := actual[i].Name
		fieldPrefix := fmt.Sprintf("%s[%s]", prefix, fieldName)
		if !reflect.DeepEqual(actual[i].Categories, desired[i].Categories) {
			diff.AddField(fieldPrefix+".categories", actual[i].Categories, desired[i].Categories)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].Collation, desired[i].Collation) {
			diff.AddField(fieldPrefix+".collation", actual[i].Collation, desired[i].Collation)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].DefaultValueExpression, desired[i].DefaultValueExpression) {
			diff.AddField(fieldPrefix+".default_value_expression", actual[i].DefaultValueExpression, desired[i].DefaultValueExpression)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].Description, desired[i].Description) {
			diff.AddField(fieldPrefix+".description", actual[i].Description, desired[i].Description)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].ForeignTypeDefinition, desired[i].ForeignTypeDefinition) {
			diff.AddField(fieldPrefix+".foreign_type_definition", actual[i].ForeignTypeDefinition, desired[i].ForeignTypeDefinition)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].MaxLength, desired[i].MaxLength) {
			diff.AddField(fieldPrefix+".max_length", actual[i].MaxLength, desired[i].MaxLength)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].Mode, desired[i].Mode) {
			diff.AddField(fieldPrefix+".mode", actual[i].Mode, desired[i].Mode)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].Name, desired[i].Name) {
			diff.AddField(fieldPrefix+".name", actual[i].Name, desired[i].Name)
			return false, nil
		}
		if !policyTagsEqual(actual[i].PolicyTags, desired[i].PolicyTags) {
			diff.AddField(fieldPrefix+".policy_tags", actual[i].PolicyTags, desired[i].PolicyTags)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].Precision, desired[i].Precision) {
			diff.AddField(fieldPrefix+".precision", actual[i].Precision, desired[i].Precision)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].RangeElementType, desired[i].RangeElementType) {
			diff.AddField(fieldPrefix+".range_element_type", actual[i].RangeElementType, desired[i].RangeElementType)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].RoundingMode, desired[i].RoundingMode) {
			diff.AddField(fieldPrefix+".rounding_mode", actual[i].RoundingMode, desired[i].RoundingMode)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].Scale, desired[i].Scale) {
			diff.AddField(fieldPrefix+".scale", actual[i].Scale, desired[i].Scale)
			return false, nil
		}
		if !reflect.DeepEqual(actual[i].Type, desired[i].Type) {
			diff.AddField(fieldPrefix+".type", actual[i].Type, desired[i].Type)
			return false, nil
		}
		eq, err := tableFieldsSchemaEqual(actual[i].Fields, desired[i].Fields, fieldPrefix+".fields", diff)
		if err != nil {
			return false, err
		}
		if !eq {
			return false, nil
		}
	}
	return true, nil
}

func externalDataConfigurationEqual(actual, desired *bigquery.ExternalDataConfiguration, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if actual == nil && desired == nil {
		return true, nil
	}
	if actual == nil || desired == nil {
		diff.AddField(prefix, actual, desired)
		return false, nil
	}
	if !reflect.DeepEqual(actual.Autodetect, desired.Autodetect) {
		diff.AddField(prefix+".autodetect", actual.Autodetect, desired.Autodetect)
		return false, nil
	}
	if !reflect.DeepEqual(actual.AvroOptions, desired.AvroOptions) {
		diff.AddField(prefix+".avro_options", actual.AvroOptions, desired.AvroOptions)
		return false, nil
	}
	if !reflect.DeepEqual(actual.Compression, desired.Compression) {
		diff.AddField(prefix+".compression", actual.Compression, desired.Compression)
		return false, nil
	}
	if !reflect.DeepEqual(actual.ConnectionId, desired.ConnectionId) {
		diff.AddField(prefix+".connection_id", actual.ConnectionId, desired.ConnectionId)
		return false, nil
	}
	if !reflect.DeepEqual(actual.CsvOptions, desired.CsvOptions) {
		diff.AddField(prefix+".csv_options", actual.CsvOptions, desired.CsvOptions)
		return false, nil
	}
	if !reflect.DeepEqual(actual.GoogleSheetsOptions, desired.GoogleSheetsOptions) {
		diff.AddField(prefix+".google_sheets_options", actual.GoogleSheetsOptions, desired.GoogleSheetsOptions)
		return false, nil
	}
	if !reflect.DeepEqual(actual.HivePartitioningOptions, desired.HivePartitioningOptions) {
		diff.AddField(prefix+".hive_partitioning_options", actual.HivePartitioningOptions, desired.HivePartitioningOptions)
		return false, nil
	}
	if !reflect.DeepEqual(actual.IgnoreUnknownValues, desired.IgnoreUnknownValues) {
		diff.AddField(prefix+".ignore_unknown_values", actual.IgnoreUnknownValues, desired.IgnoreUnknownValues)
		return false, nil
	}
	if !reflect.DeepEqual(actual.JsonOptions, desired.JsonOptions) {
		diff.AddField(prefix+".json_options", actual.JsonOptions, desired.JsonOptions)
		return false, nil
	}
	if !reflect.DeepEqual(actual.MaxBadRecords, desired.MaxBadRecords) {
		diff.AddField(prefix+".max_bad_records", actual.MaxBadRecords, desired.MaxBadRecords)
		return false, nil
	}
	if !reflect.DeepEqual(actual.MetadataCacheMode, desired.MetadataCacheMode) {
		diff.AddField(prefix+".metadata_cache_mode", actual.MetadataCacheMode, desired.MetadataCacheMode)
		return false, nil
	}
	if !reflect.DeepEqual(actual.ParquetOptions, desired.ParquetOptions) {
		diff.AddField(prefix+".parquet_options", actual.ParquetOptions, desired.ParquetOptions)
		return false, nil
	}
	if !reflect.DeepEqual(actual.SourceFormat, desired.SourceFormat) {
		diff.AddField(prefix+".source_format", actual.SourceFormat, desired.SourceFormat)
		return false, nil
	}
	if !reflect.DeepEqual(actual.SourceUris, desired.SourceUris) {
		diff.AddField(prefix+".source_uris", actual.SourceUris, desired.SourceUris)
		return false, nil
	}

	equal, err := tableSchemaEq(actual.Schema, desired.Schema, prefix+".schema", diff)
	if err != nil {
		return false, err
	}
	return equal, nil
}

func materializedViewEq(actual, desired *bigquery.MaterializedViewDefinition, prefix string, diff *structuredreporting.Diff) bool {
	if actual == nil && desired == nil {
		return true
	}
	if actual == nil || desired == nil {
		diff.AddField(prefix, actual, desired)
		return false
	}
	if !reflect.DeepEqual(actual.AllowNonIncrementalDefinition, desired.AllowNonIncrementalDefinition) {
		diff.AddField(prefix+".allow_non_incremental_definition", actual.AllowNonIncrementalDefinition, desired.AllowNonIncrementalDefinition)
		return false
	}
	if !reflect.DeepEqual(actual.EnableRefresh, desired.EnableRefresh) {
		diff.AddField(prefix+".enable_refresh", actual.EnableRefresh, desired.EnableRefresh)
		return false
	}
	if !reflect.DeepEqual(actual.Query, desired.Query) {
		diff.AddField(prefix+".query", actual.Query, desired.Query)
		return false
	}
	if !reflect.DeepEqual(actual.MaxStaleness, desired.MaxStaleness) {
		diff.AddField(prefix+".max_staleness", actual.MaxStaleness, desired.MaxStaleness)
		return false
	}
	if !reflect.DeepEqual(actual.RefreshIntervalMs, desired.RefreshIntervalMs) {
		diff.AddField(prefix+".refresh_interval_ms", actual.RefreshIntervalMs, desired.RefreshIntervalMs)
		return false
	}
	return true
}

func tableSchemaEq(actual, desired *bigquery.TableSchema, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if desired == nil {
		// If the desired schema is not specified in the KRM spec, we do not enforce it.
		// This is common for Views, Materialized Views, and External Tables where the schema is derived or autodetected.
		return true, nil
	}
	if actual == nil {
		// Desired schema is specified, but actual is nil. This is a diff.
		diff.AddField(prefix, actual, desired)
		return false, nil
	}
	return tableFieldsSchemaEqual(actual.Fields, desired.Fields, prefix+".fields", diff)
}

func viewEq(actual, desired *bigquery.ViewDefinition, prefix string, diff *structuredreporting.Diff) bool {
	if actual == nil && desired == nil {
		return true
	}
	if actual == nil || desired == nil {
		diff.AddField(prefix, actual, desired)
		return false
	}
	if !reflect.DeepEqual(actual.Query, desired.Query) {
		diff.AddField(prefix+".query", actual.Query, desired.Query)
		return false
	}
	if !reflect.DeepEqual(actual.UseLegacySql, desired.UseLegacySql) {
		diff.AddField(prefix+".use_legacy_sql", actual.UseLegacySql, desired.UseLegacySql)
		return false
	}
	return true
}

func TableEq(actual, desired *bigquery.Table, diff *structuredreporting.Diff) (bool, error) {
	if actual == nil && desired == nil {
		return true, nil
	}
	if !reflect.DeepEqual(actual.Clustering, desired.Clustering) {
		diff.AddField("clustering", actual.Clustering, desired.Clustering)
		return false, nil
	}
	if !reflect.DeepEqual(actual.Description, desired.Description) {
		diff.AddField("description", actual.Description, desired.Description)
		return false, nil
	}
	if desired.EncryptionConfiguration != nil && desired.EncryptionConfiguration.KmsKeyName != "" {
		if !reflect.DeepEqual(actual.EncryptionConfiguration, desired.EncryptionConfiguration) {
			diff.AddField("encryption_configuration", actual.EncryptionConfiguration, desired.EncryptionConfiguration)
			return false, nil
		}
	}
	if !reflect.DeepEqual(actual.ExpirationTime, desired.ExpirationTime) {
		diff.AddField("expiration_time", actual.ExpirationTime, desired.ExpirationTime)
		return false, nil
	}
	if !reflect.DeepEqual(actual.FriendlyName, desired.FriendlyName) {
		diff.AddField("friendly_name", actual.FriendlyName, desired.FriendlyName)
		return false, nil
	}
	if !materializedViewEq(actual.MaterializedView, desired.MaterializedView, "materialized_view", diff) {
		return false, nil
	}
	if !reflect.DeepEqual(actual.MaxStaleness, desired.MaxStaleness) {
		diff.AddField("max_staleness", actual.MaxStaleness, desired.MaxStaleness)
		return false, nil
	}
	if !reflect.DeepEqual(actual.RangePartitioning, desired.RangePartitioning) {
		diff.AddField("range_partitioning", actual.RangePartitioning, desired.RangePartitioning)
		return false, nil
	}
	if !reflect.DeepEqual(actual.RequirePartitionFilter, desired.RequirePartitionFilter) {
		diff.AddField("require_partition_filter", actual.RequirePartitionFilter, desired.RequirePartitionFilter)
		return false, nil
	}
	if !reflect.DeepEqual(actual.TableConstraints, desired.TableConstraints) {
		diff.AddField("table_constraints", actual.TableConstraints, desired.TableConstraints)
		return false, nil
	}

	if !viewEq(actual.View, desired.View, "view", diff) {
		return false, nil
	}
	if !reflect.DeepEqual(actual.Labels, desired.Labels) {
		diff.AddField("labels", actual.Labels, desired.Labels)
		return false, nil
	}
	equal, err := tableSchemaEq(actual.Schema, desired.Schema, "schema", diff)
	if err != nil {
		return false, err
	}
	if !equal {
		return false, nil
	}
	return externalDataConfigurationEqual(actual.ExternalDataConfiguration, desired.ExternalDataConfiguration, "external_data_configuration", diff)
}
