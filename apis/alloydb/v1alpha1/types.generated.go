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

package v1alpha1


// +kcc:proto=google.cloud.alloydb.v1.SupportedDatabaseFlag
type SupportedDatabaseFlag struct {
	// Restriction on STRING type value.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.string_restrictions
	StringRestrictions *SupportedDatabaseFlag_StringRestrictions `json:"stringRestrictions,omitempty"`

	// Restriction on INTEGER type value.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.integer_restrictions
	IntegerRestrictions *SupportedDatabaseFlag_IntegerRestrictions `json:"integerRestrictions,omitempty"`

	// The name of the flag resource, following Google Cloud conventions, e.g.:
	//   * projects/{project}/locations/{location}/flags/{flag}
	//  This field currently has no semantic meaning.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.name
	Name *string `json:"name,omitempty"`

	// The name of the database flag, e.g. "max_allowed_packets".
	//  The is a possibly key for the Instance.database_flags map field.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.flag_name
	FlagName *string `json:"flagName,omitempty"`

	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.value_type
	ValueType *string `json:"valueType,omitempty"`

	// Whether the database flag accepts multiple values. If true,
	//  a comma-separated list of stringified values may be specified.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.accepts_multiple_values
	AcceptsMultipleValues *bool `json:"acceptsMultipleValues,omitempty"`

	// Major database engine versions for which this flag is supported.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.supported_db_versions
	SupportedDbVersions []string `json:"supportedDbVersions,omitempty"`

	// Whether setting or updating this flag on an Instance requires a database
	//  restart. If a flag that requires database restart is set, the backend
	//  will automatically restart the database (making sure to satisfy any
	//  availability SLO's).
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.requires_db_restart
	RequiresDbRestart *bool `json:"requiresDbRestart,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1.SupportedDatabaseFlag.IntegerRestrictions
type SupportedDatabaseFlag_IntegerRestrictions struct {
	// The minimum value that can be specified, if applicable.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.IntegerRestrictions.min_value
	MinValue *int64 `json:"minValue,omitempty"`

	// The maximum value that can be specified, if applicable.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.IntegerRestrictions.max_value
	MaxValue *int64 `json:"maxValue,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1.SupportedDatabaseFlag.StringRestrictions
type SupportedDatabaseFlag_StringRestrictions struct {
	// The list of allowed values, if bounded. This field will be empty
	//  if there is a unbounded number of allowed values.
	// +kcc:proto:field=google.cloud.alloydb.v1.SupportedDatabaseFlag.StringRestrictions.allowed_values
	AllowedValues []string `json:"allowedValues,omitempty"`
}
