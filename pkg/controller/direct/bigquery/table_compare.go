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
	// The fields from API can be in a different order.
	// Sort by name before comparing.
	sortSchemaFields(desired)
	sortSchemaFields(actual)
	for i := range desired {
		fieldName := desired[i].Name
		fieldPrefix := fmt.Sprintf("%s[%s]", prefix, fieldName)
		if !reflect.DeepEqual(desired[i].Categories, actual[i].Categories) {
			diff.AddField(fieldPrefix+".categories", desired[i].Categories, actual[i].Categories)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].Collation, actual[i].Collation) {
			diff.AddField(fieldPrefix+".collation", desired[i].Collation, actual[i].Collation)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].DefaultValueExpression, actual[i].DefaultValueExpression) {
			diff.AddField(fieldPrefix+".default_value_expression", desired[i].DefaultValueExpression, actual[i].DefaultValueExpression)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].Description, actual[i].Description) {
			diff.AddField(fieldPrefix+".description", desired[i].Description, actual[i].Description)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].ForeignTypeDefinition, actual[i].ForeignTypeDefinition) {
			diff.AddField(fieldPrefix+".foreign_type_definition", desired[i].ForeignTypeDefinition, actual[i].ForeignTypeDefinition)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].MaxLength, actual[i].MaxLength) {
			diff.AddField(fieldPrefix+".max_length", desired[i].MaxLength, actual[i].MaxLength)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].Mode, actual[i].Mode) {
			diff.AddField(fieldPrefix+".mode", desired[i].Mode, actual[i].Mode)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].Name, actual[i].Name) {
			diff.AddField(fieldPrefix+".name", desired[i].Name, actual[i].Name)
			return false, nil
		}
		if !policyTagsEqual(desired[i].PolicyTags, actual[i].PolicyTags) {
			diff.AddField(fieldPrefix+".policy_tags", desired[i].PolicyTags, actual[i].PolicyTags)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].Precision, actual[i].Precision) {
			diff.AddField(fieldPrefix+".precision", desired[i].Precision, actual[i].Precision)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].RangeElementType, actual[i].RangeElementType) {
			diff.AddField(fieldPrefix+".range_element_type", desired[i].RangeElementType, actual[i].RangeElementType)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].RoundingMode, actual[i].RoundingMode) {
			diff.AddField(fieldPrefix+".rounding_mode", desired[i].RoundingMode, actual[i].RoundingMode)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].Scale, actual[i].Scale) {
			diff.AddField(fieldPrefix+".scale", desired[i].Scale, actual[i].Scale)
			return false, nil
		}
		if !reflect.DeepEqual(desired[i].Type, actual[i].Type) {
			diff.AddField(fieldPrefix+".type", desired[i].Type, actual[i].Type)
			return false, nil
		}
		eq, err := tableFieldsSchemaEqual(desired[i].Fields, actual[i].Fields, fieldPrefix+".fields", diff)
		if err != nil {
			return false, err
		}
		if !eq {
			return false, nil
		}
	}
	return true, nil
}

func externalDataConfigurationEqual(a, b *bigquery.ExternalDataConfiguration, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}
	if a == nil || b == nil {
		diff.AddField(prefix, a, b)
		return false, nil
	}
	if !reflect.DeepEqual(a.Autodetect, b.Autodetect) {
		diff.AddField(prefix+".autodetect", a.Autodetect, b.Autodetect)
		return false, nil
	}
	if !reflect.DeepEqual(a.AvroOptions, b.AvroOptions) {
		diff.AddField(prefix+".avro_options", a.AvroOptions, b.AvroOptions)
		return false, nil
	}
	if !reflect.DeepEqual(a.Compression, b.Compression) {
		diff.AddField(prefix+".compression", a.Compression, b.Compression)
		return false, nil
	}
	if !reflect.DeepEqual(a.ConnectionId, b.ConnectionId) {
		diff.AddField(prefix+".connection_id", a.ConnectionId, b.ConnectionId)
		return false, nil
	}
	if !reflect.DeepEqual(a.CsvOptions, b.CsvOptions) {
		diff.AddField(prefix+".csv_options", a.CsvOptions, b.CsvOptions)
		return false, nil
	}
	if !reflect.DeepEqual(a.GoogleSheetsOptions, b.GoogleSheetsOptions) {
		diff.AddField(prefix+".google_sheets_options", a.GoogleSheetsOptions, b.GoogleSheetsOptions)
		return false, nil
	}
	if !reflect.DeepEqual(a.HivePartitioningOptions, b.HivePartitioningOptions) {
		diff.AddField(prefix+".hive_partitioning_options", a.HivePartitioningOptions, b.HivePartitioningOptions)
		return false, nil
	}
	if !reflect.DeepEqual(a.IgnoreUnknownValues, b.IgnoreUnknownValues) {
		diff.AddField(prefix+".ignore_unknown_values", a.IgnoreUnknownValues, b.IgnoreUnknownValues)
		return false, nil
	}
	if !reflect.DeepEqual(a.JsonOptions, b.JsonOptions) {
		diff.AddField(prefix+".json_options", a.JsonOptions, b.JsonOptions)
		return false, nil
	}
	if !reflect.DeepEqual(a.MaxBadRecords, b.MaxBadRecords) {
		diff.AddField(prefix+".max_bad_records", a.MaxBadRecords, b.MaxBadRecords)
		return false, nil
	}
	if !reflect.DeepEqual(a.MetadataCacheMode, b.MetadataCacheMode) {
		diff.AddField(prefix+".metadata_cache_mode", a.MetadataCacheMode, b.MetadataCacheMode)
		return false, nil
	}
	if !reflect.DeepEqual(a.ParquetOptions, b.ParquetOptions) {
		diff.AddField(prefix+".parquet_options", a.ParquetOptions, b.ParquetOptions)
		return false, nil
	}
	if !reflect.DeepEqual(a.SourceFormat, b.SourceFormat) {
		diff.AddField(prefix+".source_format", a.SourceFormat, b.SourceFormat)
		return false, nil
	}
	if !reflect.DeepEqual(a.SourceUris, b.SourceUris) {
		diff.AddField(prefix+".source_uris", a.SourceUris, b.SourceUris)
		return false, nil
	}

	equal, err := tableSchemaEq(a.Schema, b.Schema, prefix+".schema", diff)
	if err != nil {
		return false, err
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
	if !reflect.DeepEqual(a.AllowNonIncrementalDefinition, b.AllowNonIncrementalDefinition) {
		diff.AddField(prefix+".allow_non_incremental_definition", a.AllowNonIncrementalDefinition, b.AllowNonIncrementalDefinition)
		return false
	}
	if !reflect.DeepEqual(a.EnableRefresh, b.EnableRefresh) {
		diff.AddField(prefix+".enable_refresh", a.EnableRefresh, b.EnableRefresh)
		return false
	}
	if !reflect.DeepEqual(a.Query, b.Query) {
		diff.AddField(prefix+".query", a.Query, b.Query)
		return false
	}
	if !reflect.DeepEqual(a.MaxStaleness, b.MaxStaleness) {
		diff.AddField(prefix+".max_staleness", a.MaxStaleness, b.MaxStaleness)
		return false
	}
	if !reflect.DeepEqual(a.RefreshIntervalMs, b.RefreshIntervalMs) {
		diff.AddField(prefix+".refresh_interval_ms", a.RefreshIntervalMs, b.RefreshIntervalMs)
		return false
	}
	return true
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
	if !reflect.DeepEqual(a.Query, b.Query) {
		diff.AddField(prefix+".query", a.Query, b.Query)
		return false
	}
	if !reflect.DeepEqual(a.UseLegacySql, b.UseLegacySql) {
		diff.AddField(prefix+".use_legacy_sql", a.UseLegacySql, b.UseLegacySql)
		return false
	}
	return true
}

func TableEq(a, b *bigquery.Table, diff *structuredreporting.Diff) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}
	if !reflect.DeepEqual(a.Clustering, b.Clustering) {
		diff.AddField("clustering", a.Clustering, b.Clustering)
		return false, nil
	}
	if !reflect.DeepEqual(a.Description, b.Description) {
		diff.AddField("description", a.Description, b.Description)
		return false, nil
	}
	if !reflect.DeepEqual(a.EncryptionConfiguration, b.EncryptionConfiguration) {
		diff.AddField("encryption_configuration", a.EncryptionConfiguration, b.EncryptionConfiguration)
		return false, nil
	}
	if !reflect.DeepEqual(a.ExpirationTime, b.ExpirationTime) {
		diff.AddField("expiration_time", a.ExpirationTime, b.ExpirationTime)
		return false, nil
	}
	if !reflect.DeepEqual(a.FriendlyName, b.FriendlyName) {
		diff.AddField("friendly_name", a.FriendlyName, b.FriendlyName)
		return false, nil
	}
	if !materializedViewEq(a.MaterializedView, b.MaterializedView, "materialized_view", diff) {
		return false, nil
	}
	if !reflect.DeepEqual(a.MaxStaleness, b.MaxStaleness) {
		diff.AddField("max_staleness", a.MaxStaleness, b.MaxStaleness)
		return false, nil
	}
	if !reflect.DeepEqual(a.RangePartitioning, b.RangePartitioning) {
		diff.AddField("range_partitioning", a.RangePartitioning, b.RangePartitioning)
		return false, nil
	}
	if !reflect.DeepEqual(a.RequirePartitionFilter, b.RequirePartitionFilter) {
		diff.AddField("require_partition_filter", a.RequirePartitionFilter, b.RequirePartitionFilter)
		return false, nil
	}
	if !reflect.DeepEqual(a.TableConstraints, b.TableConstraints) {
		diff.AddField("table_constraints", a.TableConstraints, b.TableConstraints)
		return false, nil
	}

	if !viewEq(a.View, b.View, "view", diff) {
		return false, nil
	}
	if !reflect.DeepEqual(a.Labels, b.Labels) {
		diff.AddField("labels", a.Labels, b.Labels)
		return false, nil
	}
	equal, err := tableSchemaEq(a.Schema, b.Schema, "schema", diff)
	if err != nil {
		return false, err
	}
	if !equal {
		return false, nil
	}
	return externalDataConfigurationEqual(a.ExternalDataConfiguration, b.ExternalDataConfiguration, "external_data_configuration", diff)
}
