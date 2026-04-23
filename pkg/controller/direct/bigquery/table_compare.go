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
	aEmpty := (a == nil || len(a.Names) == 0)
	bEmpty := (b == nil || len(b.Names) == 0)
	if aEmpty && bEmpty {
		return true
	}
	if aEmpty || bEmpty {
		return false
	}

	if len(a.Names) != len(b.Names) {
		return false
	}
	sort.Strings(a.Names)
	sort.Strings(b.Names)
	for i := range a.Names {
		if a.Names[i] != b.Names[i] {
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
	startDiffs := len(diff.Fields)
	// The fields from API can be in a different order.
	// Sort by name before comparing.
	sortSchemaFields(desired)
	sortSchemaFields(actual)
	for i := range desired {
		fieldName := desired[i].Name
		fieldPrefix := fmt.Sprintf("%s[%s]", prefix, fieldName)
		if !reflect.DeepEqual(desired[i].Categories, actual[i].Categories) {
			diff.AddField(fieldPrefix+".categories", desired[i].Categories, actual[i].Categories)
		}
		if !reflect.DeepEqual(desired[i].Collation, actual[i].Collation) {
			diff.AddField(fieldPrefix+".collation", desired[i].Collation, actual[i].Collation)
		}
		if !reflect.DeepEqual(desired[i].DefaultValueExpression, actual[i].DefaultValueExpression) {
			diff.AddField(fieldPrefix+".default_value_expression", desired[i].DefaultValueExpression, actual[i].DefaultValueExpression)
		}
		if !reflect.DeepEqual(desired[i].Description, actual[i].Description) {
			diff.AddField(fieldPrefix+".description", desired[i].Description, actual[i].Description)
		}
		if !reflect.DeepEqual(desired[i].ForeignTypeDefinition, actual[i].ForeignTypeDefinition) {
			diff.AddField(fieldPrefix+".foreign_type_definition", desired[i].ForeignTypeDefinition, actual[i].ForeignTypeDefinition)
		}
		if !reflect.DeepEqual(desired[i].MaxLength, actual[i].MaxLength) {
			diff.AddField(fieldPrefix+".max_length", desired[i].MaxLength, actual[i].MaxLength)
		}
		if !reflect.DeepEqual(desired[i].Mode, actual[i].Mode) {
			diff.AddField(fieldPrefix+".mode", desired[i].Mode, actual[i].Mode)
		}
		if !reflect.DeepEqual(desired[i].Name, actual[i].Name) {
			diff.AddField(fieldPrefix+".name", desired[i].Name, actual[i].Name)
		}
		if !policyTagsEqual(desired[i].PolicyTags, actual[i].PolicyTags) {
			diff.AddField(fieldPrefix+".policy_tags", desired[i].PolicyTags, actual[i].PolicyTags)
		}
		if !reflect.DeepEqual(desired[i].Precision, actual[i].Precision) {
			diff.AddField(fieldPrefix+".precision", desired[i].Precision, actual[i].Precision)
		}
		if !reflect.DeepEqual(desired[i].RangeElementType, actual[i].RangeElementType) {
			diff.AddField(fieldPrefix+".range_element_type", desired[i].RangeElementType, actual[i].RangeElementType)
		}
		if !reflect.DeepEqual(desired[i].RoundingMode, actual[i].RoundingMode) {
			diff.AddField(fieldPrefix+".rounding_mode", desired[i].RoundingMode, actual[i].RoundingMode)
		}
		if !reflect.DeepEqual(desired[i].Scale, actual[i].Scale) {
			diff.AddField(fieldPrefix+".scale", desired[i].Scale, actual[i].Scale)
		}
		if !reflect.DeepEqual(desired[i].Type, actual[i].Type) {
			diff.AddField(fieldPrefix+".type", desired[i].Type, actual[i].Type)
		}
		if _, err := tableFieldsSchemaEqual(desired[i].Fields, actual[i].Fields, fieldPrefix+".fields", diff); err != nil {
			return false, err
		}
	}
	return len(diff.Fields) == startDiffs, nil
}

func externalDataConfigurationEqual(a, b *bigquery.ExternalDataConfiguration, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false, nil
	}
	startDiffs := len(diff.Fields)
	if !reflect.DeepEqual(a.Autodetect, b.Autodetect) {
		diff.AddField(prefix+".autodetect", a.Autodetect, b.Autodetect)
	}
	if !reflect.DeepEqual(a.AvroOptions, b.AvroOptions) {
		diff.AddField(prefix+".avro_options", a.AvroOptions, b.AvroOptions)
	}
	if !reflect.DeepEqual(a.Compression, b.Compression) {
		diff.AddField(prefix+".compression", a.Compression, b.Compression)
	}
	if !reflect.DeepEqual(a.ConnectionId, b.ConnectionId) {
		diff.AddField(prefix+".connection_id", a.ConnectionId, b.ConnectionId)
	}
	if !reflect.DeepEqual(a.CsvOptions, b.CsvOptions) {
		diff.AddField(prefix+".csv_options", a.CsvOptions, b.CsvOptions)
	}
	if !reflect.DeepEqual(a.GoogleSheetsOptions, b.GoogleSheetsOptions) {
		diff.AddField(prefix+".google_sheets_options", a.GoogleSheetsOptions, b.GoogleSheetsOptions)
	}
	if !reflect.DeepEqual(a.HivePartitioningOptions, b.HivePartitioningOptions) {
		diff.AddField(prefix+".hive_partitioning_options", a.HivePartitioningOptions, b.HivePartitioningOptions)
	}
	if !reflect.DeepEqual(a.IgnoreUnknownValues, b.IgnoreUnknownValues) {
		diff.AddField(prefix+".ignore_unknown_values", a.IgnoreUnknownValues, b.IgnoreUnknownValues)
	}
	if !reflect.DeepEqual(a.JsonOptions, b.JsonOptions) {
		diff.AddField(prefix+".json_options", a.JsonOptions, b.JsonOptions)
	}
	if !reflect.DeepEqual(a.MaxBadRecords, b.MaxBadRecords) {
		diff.AddField(prefix+".max_bad_records", a.MaxBadRecords, b.MaxBadRecords)
	}
	if !reflect.DeepEqual(a.MetadataCacheMode, b.MetadataCacheMode) {
		diff.AddField(prefix+".metadata_cache_mode", a.MetadataCacheMode, b.MetadataCacheMode)
	}
	if !reflect.DeepEqual(a.ParquetOptions, b.ParquetOptions) {
		diff.AddField(prefix+".parquet_options", a.ParquetOptions, b.ParquetOptions)
	}
	if !reflect.DeepEqual(a.SourceFormat, b.SourceFormat) {
		diff.AddField(prefix+".source_format", a.SourceFormat, b.SourceFormat)
	}
	if !reflect.DeepEqual(a.SourceUris, b.SourceUris) {
		diff.AddField(prefix+".source_uris", a.SourceUris, b.SourceUris)
	}

	if _, err := tableSchemaEq(a.Schema, b.Schema, prefix+".schema", diff); err != nil {
		return false, err
	}
	return len(diff.Fields) == startDiffs, nil
}

func materializedViewEq(a, b *bigquery.MaterializedViewDefinition, prefix string, diff *structuredreporting.Diff) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false
	}
	startDiffs := len(diff.Fields)
	if !reflect.DeepEqual(a.AllowNonIncrementalDefinition, b.AllowNonIncrementalDefinition) {
		diff.AddField(prefix+".allow_non_incremental_definition", a.AllowNonIncrementalDefinition, b.AllowNonIncrementalDefinition)
	}
	if !reflect.DeepEqual(a.EnableRefresh, b.EnableRefresh) {
		diff.AddField(prefix+".enable_refresh", a.EnableRefresh, b.EnableRefresh)
	}
	if !reflect.DeepEqual(a.Query, b.Query) {
		diff.AddField(prefix+".query", a.Query, b.Query)
	}
	if !reflect.DeepEqual(a.MaxStaleness, b.MaxStaleness) {
		diff.AddField(prefix+".max_staleness", a.MaxStaleness, b.MaxStaleness)
	}
	if !reflect.DeepEqual(a.RefreshIntervalMs, b.RefreshIntervalMs) {
		diff.AddField(prefix+".refresh_interval_ms", a.RefreshIntervalMs, b.RefreshIntervalMs)
	}
	return len(diff.Fields) == startDiffs
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
	startDiffs := len(diff.Fields)
	if !reflect.DeepEqual(a.Query, b.Query) {
		diff.AddField(prefix+".query", a.Query, b.Query)
	}
	if !reflect.DeepEqual(a.UseLegacySql, b.UseLegacySql) {
		diff.AddField(prefix+".use_legacy_sql", a.UseLegacySql, b.UseLegacySql)
	}
	return len(diff.Fields) == startDiffs
}

func TableEq(a, b *bigquery.Table, diff *structuredreporting.Diff) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}
	if a == nil || b == nil {
		diff.AddField("table", a, b)
		return false, nil
	}

	startDiffs := len(diff.Fields)

	if !reflect.DeepEqual(a.Clustering, b.Clustering) {
		diff.AddField("clustering", a.Clustering, b.Clustering)
	}
	if !reflect.DeepEqual(a.Description, b.Description) {
		diff.AddField("description", a.Description, b.Description)
	}
	if !reflect.DeepEqual(a.EncryptionConfiguration, b.EncryptionConfiguration) {
		diff.AddField("encryption_configuration", a.EncryptionConfiguration, b.EncryptionConfiguration)
	}
	if !reflect.DeepEqual(a.ExpirationTime, b.ExpirationTime) {
		diff.AddField("expiration_time", a.ExpirationTime, b.ExpirationTime)
	}
	if !reflect.DeepEqual(a.FriendlyName, b.FriendlyName) {
		diff.AddField("friendly_name", a.FriendlyName, b.FriendlyName)
	}
	materializedViewEq(a.MaterializedView, b.MaterializedView, "materialized_view", diff)

	if !reflect.DeepEqual(a.MaxStaleness, b.MaxStaleness) {
		diff.AddField("max_staleness", a.MaxStaleness, b.MaxStaleness)
	}
	if !reflect.DeepEqual(a.RangePartitioning, b.RangePartitioning) {
		diff.AddField("range_partitioning", a.RangePartitioning, b.RangePartitioning)
	}
	if !reflect.DeepEqual(a.RequirePartitionFilter, b.RequirePartitionFilter) {
		diff.AddField("require_partition_filter", a.RequirePartitionFilter, b.RequirePartitionFilter)
	}
	if !reflect.DeepEqual(a.TableConstraints, b.TableConstraints) {
		diff.AddField("table_constraints", a.TableConstraints, b.TableConstraints)
	}

	viewEq(a.View, b.View, "view", diff)

	if !reflect.DeepEqual(a.Labels, b.Labels) {
		diff.AddField("labels", a.Labels, b.Labels)
	}
	if _, err := tableSchemaEq(a.Schema, b.Schema, "schema", diff); err != nil {
		return false, err
	}
	if _, err := externalDataConfigurationEqual(a.ExternalDataConfiguration, b.ExternalDataConfiguration, "external_data_configuration", diff); err != nil {
		return false, err
	}

	return len(diff.Fields) == startDiffs, nil
}
