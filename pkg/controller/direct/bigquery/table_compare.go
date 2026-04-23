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
	startDiffs := len(diff.Fields)
	for i := range actual {
		fieldName := actual[i].Name
		fieldPrefix := fmt.Sprintf("%s[%s]", prefix, fieldName)
		if !reflect.DeepEqual(actual[i].Categories, desired[i].Categories) {
			diff.AddField(fieldPrefix+".categories", actual[i].Categories, desired[i].Categories)
		}
		if !reflect.DeepEqual(actual[i].Collation, desired[i].Collation) {
			diff.AddField(fieldPrefix+".collation", actual[i].Collation, desired[i].Collation)
		}
		if !reflect.DeepEqual(actual[i].DefaultValueExpression, desired[i].DefaultValueExpression) {
			diff.AddField(fieldPrefix+".default_value_expression", actual[i].DefaultValueExpression, desired[i].DefaultValueExpression)
		}
		if !reflect.DeepEqual(actual[i].Description, desired[i].Description) {
			diff.AddField(fieldPrefix+".description", actual[i].Description, desired[i].Description)
		}
		if !reflect.DeepEqual(actual[i].ForeignTypeDefinition, desired[i].ForeignTypeDefinition) {
			diff.AddField(fieldPrefix+".foreign_type_definition", actual[i].ForeignTypeDefinition, desired[i].ForeignTypeDefinition)
		}
		if !reflect.DeepEqual(actual[i].MaxLength, desired[i].MaxLength) {
			diff.AddField(fieldPrefix+".max_length", actual[i].MaxLength, desired[i].MaxLength)
		}
		if !reflect.DeepEqual(actual[i].Mode, desired[i].Mode) {
			diff.AddField(fieldPrefix+".mode", actual[i].Mode, desired[i].Mode)
		}
		if !reflect.DeepEqual(actual[i].Name, desired[i].Name) {
			diff.AddField(fieldPrefix+".name", actual[i].Name, desired[i].Name)
		}
		if !policyTagsEqual(actual[i].PolicyTags, desired[i].PolicyTags) {
			diff.AddField(fieldPrefix+".policy_tags", actual[i].PolicyTags, desired[i].PolicyTags)
		}
		if !reflect.DeepEqual(actual[i].Precision, desired[i].Precision) {
			diff.AddField(fieldPrefix+".precision", actual[i].Precision, desired[i].Precision)
		}
		if !reflect.DeepEqual(actual[i].RangeElementType, desired[i].RangeElementType) {
			diff.AddField(fieldPrefix+".range_element_type", actual[i].RangeElementType, desired[i].RangeElementType)
		}
		if !reflect.DeepEqual(actual[i].RoundingMode, desired[i].RoundingMode) {
			diff.AddField(fieldPrefix+".rounding_mode", actual[i].RoundingMode, desired[i].RoundingMode)
		}
		if !reflect.DeepEqual(actual[i].Scale, desired[i].Scale) {
			diff.AddField(fieldPrefix+".scale", actual[i].Scale, desired[i].Scale)
		}
		if !reflect.DeepEqual(actual[i].Type, desired[i].Type) {
			diff.AddField(fieldPrefix+".type", actual[i].Type, desired[i].Type)
		}
		if _, err := tableFieldsSchemaEqual(actual[i].Fields, desired[i].Fields, fieldPrefix+".fields", diff); err != nil {
			return false, err
		}
	}
	return len(diff.Fields) == startDiffs, nil
}

func externalDataConfigurationEqual(actual, desired *bigquery.ExternalDataConfiguration, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if actual == nil && desired == nil {
		return true, nil
	}
	if actual == nil || desired == nil {
		diff.AddField(prefix, actual, desired)
		return false, nil
	}
	startDiffs := len(diff.Fields)
	if !reflect.DeepEqual(actual.Autodetect, desired.Autodetect) {
		diff.AddField(prefix+".autodetect", actual.Autodetect, desired.Autodetect)
	}
	if !reflect.DeepEqual(actual.AvroOptions, desired.AvroOptions) {
		diff.AddField(prefix+".avro_options", actual.AvroOptions, desired.AvroOptions)
	}
	if !reflect.DeepEqual(actual.Compression, desired.Compression) {
		diff.AddField(prefix+".compression", actual.Compression, desired.Compression)
	}
	if !reflect.DeepEqual(actual.ConnectionId, desired.ConnectionId) {
		diff.AddField(prefix+".connection_id", actual.ConnectionId, desired.ConnectionId)
	}
	if !reflect.DeepEqual(actual.CsvOptions, desired.CsvOptions) {
		diff.AddField(prefix+".csv_options", actual.CsvOptions, desired.CsvOptions)
	}
	if !reflect.DeepEqual(actual.GoogleSheetsOptions, desired.GoogleSheetsOptions) {
		diff.AddField(prefix+".google_sheets_options", actual.GoogleSheetsOptions, desired.GoogleSheetsOptions)
	}
	if !reflect.DeepEqual(actual.HivePartitioningOptions, desired.HivePartitioningOptions) {
		diff.AddField(prefix+".hive_partitioning_options", actual.HivePartitioningOptions, desired.HivePartitioningOptions)
	}
	if !reflect.DeepEqual(actual.IgnoreUnknownValues, desired.IgnoreUnknownValues) {
		diff.AddField(prefix+".ignore_unknown_values", actual.IgnoreUnknownValues, desired.IgnoreUnknownValues)
	}
	if !reflect.DeepEqual(actual.JsonOptions, desired.JsonOptions) {
		diff.AddField(prefix+".json_options", actual.JsonOptions, desired.JsonOptions)
	}
	if !reflect.DeepEqual(actual.MaxBadRecords, desired.MaxBadRecords) {
		diff.AddField(prefix+".max_bad_records", actual.MaxBadRecords, desired.MaxBadRecords)
	}
	if !reflect.DeepEqual(actual.MetadataCacheMode, desired.MetadataCacheMode) {
		diff.AddField(prefix+".metadata_cache_mode", actual.MetadataCacheMode, desired.MetadataCacheMode)
	}
	if !reflect.DeepEqual(actual.ParquetOptions, desired.ParquetOptions) {
		diff.AddField(prefix+".parquet_options", actual.ParquetOptions, desired.ParquetOptions)
	}
	if !reflect.DeepEqual(actual.SourceFormat, desired.SourceFormat) {
		diff.AddField(prefix+".source_format", actual.SourceFormat, desired.SourceFormat)
	}
	if !reflect.DeepEqual(actual.SourceUris, desired.SourceUris) {
		diff.AddField(prefix+".source_uris", actual.SourceUris, desired.SourceUris)
	}

	if _, err := tableSchemaEq(actual.Schema, desired.Schema, prefix+".schema", diff); err != nil {
		return false, err
	}
	return len(diff.Fields) == startDiffs, nil
}

func materializedViewEq(actual, desired *bigquery.MaterializedViewDefinition, prefix string, diff *structuredreporting.Diff) bool {
	if actual == nil && desired == nil {
		return true
	}
	if actual == nil || desired == nil {
		diff.AddField(prefix, actual, desired)
		return false
	}
	startDiffs := len(diff.Fields)
	if !reflect.DeepEqual(actual.AllowNonIncrementalDefinition, desired.AllowNonIncrementalDefinition) {
		diff.AddField(prefix+".allow_non_incremental_definition", actual.AllowNonIncrementalDefinition, desired.AllowNonIncrementalDefinition)
	}
	if !reflect.DeepEqual(actual.EnableRefresh, desired.EnableRefresh) {
		diff.AddField(prefix+".enable_refresh", actual.EnableRefresh, desired.EnableRefresh)
	}
	if !reflect.DeepEqual(actual.Query, desired.Query) {
		diff.AddField(prefix+".query", actual.Query, desired.Query)
	}
	if !reflect.DeepEqual(actual.MaxStaleness, desired.MaxStaleness) {
		diff.AddField(prefix+".max_staleness", actual.MaxStaleness, desired.MaxStaleness)
	}
	if !reflect.DeepEqual(actual.RefreshIntervalMs, desired.RefreshIntervalMs) {
		diff.AddField(prefix+".refresh_interval_ms", actual.RefreshIntervalMs, desired.RefreshIntervalMs)
	}
	return len(diff.Fields) == startDiffs
}

func tableSchemaEq(actual, desired *bigquery.TableSchema, prefix string, diff *structuredreporting.Diff) (bool, error) {
	if actual == nil && desired == nil {
		return true, nil
	}
	if actual == nil || desired == nil {
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
	startDiffs := len(diff.Fields)
	if !reflect.DeepEqual(actual.Query, desired.Query) {
		diff.AddField(prefix+".query", actual.Query, desired.Query)
	}
	if !reflect.DeepEqual(actual.UseLegacySql, desired.UseLegacySql) {
		diff.AddField(prefix+".use_legacy_sql", actual.UseLegacySql, desired.UseLegacySql)
	}
	return len(diff.Fields) == startDiffs
}

func TableEq(actual, desired *bigquery.Table, diff *structuredreporting.Diff) (bool, error) {
	if actual == nil && desired == nil {
		return true, nil
	}
	if actual == nil || desired == nil {
		diff.AddField("table", actual, desired)
		return false, nil
	}

	startDiffs := len(diff.Fields)

	if !reflect.DeepEqual(actual.Clustering, desired.Clustering) {
		diff.AddField("clustering", actual.Clustering, desired.Clustering)
	}
	if !reflect.DeepEqual(actual.Description, desired.Description) {
		diff.AddField("description", actual.Description, desired.Description)
	}
	if !reflect.DeepEqual(actual.EncryptionConfiguration, desired.EncryptionConfiguration) {
		diff.AddField("encryption_configuration", actual.EncryptionConfiguration, desired.EncryptionConfiguration)
	}
	if !reflect.DeepEqual(actual.ExpirationTime, desired.ExpirationTime) {
		diff.AddField("expiration_time", actual.ExpirationTime, desired.ExpirationTime)
	}
	if !reflect.DeepEqual(actual.FriendlyName, desired.FriendlyName) {
		diff.AddField("friendly_name", actual.FriendlyName, desired.FriendlyName)
	}
	materializedViewEq(actual.MaterializedView, desired.MaterializedView, "materialized_view", diff)

	if !reflect.DeepEqual(actual.MaxStaleness, desired.MaxStaleness) {
		diff.AddField("max_staleness", actual.MaxStaleness, desired.MaxStaleness)
	}
	if !reflect.DeepEqual(actual.RangePartitioning, desired.RangePartitioning) {
		diff.AddField("range_partitioning", actual.RangePartitioning, desired.RangePartitioning)
	}
	if !reflect.DeepEqual(actual.RequirePartitionFilter, desired.RequirePartitionFilter) {
		diff.AddField("require_partition_filter", actual.RequirePartitionFilter, desired.RequirePartitionFilter)
	}
	if !reflect.DeepEqual(actual.TableConstraints, desired.TableConstraints) {
		diff.AddField("table_constraints", actual.TableConstraints, desired.TableConstraints)
	}
	if !reflect.DeepEqual(actual.TimePartitioning, desired.TimePartitioning) {
		diff.AddField("time_partitioning", actual.TimePartitioning, desired.TimePartitioning)
	}

	viewEq(actual.View, desired.View, "view", diff)

	if !reflect.DeepEqual(actual.Labels, desired.Labels) {
		diff.AddField("labels", actual.Labels, desired.Labels)
	}
	if _, err := tableSchemaEq(actual.Schema, desired.Schema, "schema", diff); err != nil {
		return false, err
	}
	if _, err := externalDataConfigurationEqual(actual.ExternalDataConfiguration, desired.ExternalDataConfiguration, "external_data_configuration", diff); err != nil {
		return false, err
	}

	return len(diff.Fields) == startDiffs, nil
}
