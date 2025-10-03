// Copyright 2024 Google LLC. All Rights Reserved.
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
package alpha

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type JobTrigger struct {
	Name        *string               `json:"name"`
	DisplayName *string               `json:"displayName"`
	Description *string               `json:"description"`
	InspectJob  *JobTriggerInspectJob `json:"inspectJob"`
	Triggers    []JobTriggerTriggers  `json:"triggers"`
	Errors      []JobTriggerErrors    `json:"errors"`
	CreateTime  *string               `json:"createTime"`
	UpdateTime  *string               `json:"updateTime"`
	LastRunTime *string               `json:"lastRunTime"`
	Status      *JobTriggerStatusEnum `json:"status"`
	LocationId  *string               `json:"locationId"`
	Parent      *string               `json:"parent"`
	Location    *string               `json:"location"`
}

func (r *JobTrigger) String() string {
	return dcl.SprintResource(r)
}

// The enum JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum.
type JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum string

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumRef returns a *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumRef(s string) *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	v := JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(s)
	return &v
}

func (v JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"FILE_TYPE_UNSPECIFIED", "BINARY_FILE", "TEXT_FILE", "IMAGE", "WORD", "PDF", "AVRO", "CSV", "TSV"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum.
type JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum string

// JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumRef returns a *JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumRef(s string) *JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	v := JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(s)
	return &v
}

func (v JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SAMPLE_METHOD_UNSPECIFIED", "TOP", "RANDOM_START"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum.
type JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum string

// JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumRef returns a *JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumRef(s string) *JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	v := JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(s)
	return &v
}

func (v JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SAMPLE_METHOD_UNSPECIFIED", "TOP", "RANDOM_START"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobInspectConfigMinLikelihoodEnum.
type JobTriggerInspectJobInspectConfigMinLikelihoodEnum string

// JobTriggerInspectJobInspectConfigMinLikelihoodEnumRef returns a *JobTriggerInspectJobInspectConfigMinLikelihoodEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobInspectConfigMinLikelihoodEnumRef(s string) *JobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	v := JobTriggerInspectJobInspectConfigMinLikelihoodEnum(s)
	return &v
}

func (v JobTriggerInspectJobInspectConfigMinLikelihoodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"LIKELIHOOD_UNSPECIFIED", "VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobInspectConfigMinLikelihoodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum.
type JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum string

// JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumRef returns a *JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumRef(s string) *JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	v := JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(s)
	return &v
}

func (v JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"LIKELIHOOD_UNSPECIFIED", "VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.
type JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum string

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef returns a *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s string) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	v := JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(s)
	return &v
}

func (v JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"LIKELIHOOD_UNSPECIFIED", "VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum.
type JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum string

// JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumRef returns a *JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumRef(s string) *JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	v := JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(s)
	return &v
}

func (v JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EXCLUSION_TYPE_UNSPECIFIED", "EXCLUSION_TYPE_EXCLUDE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.
type JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum string

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef returns a *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s string) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	v := JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(s)
	return &v
}

func (v JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"LIKELIHOOD_UNSPECIFIED", "VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.
type JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum string

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumRef returns a *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumRef(s string) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	v := JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(s)
	return &v
}

func (v JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MATCHING_TYPE_UNSPECIFIED", "MATCHING_TYPE_FULL_MATCH", "MATCHING_TYPE_PARTIAL_MATCH", "MATCHING_TYPE_INVERSE_MATCH"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum.
type JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum string

// JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumRef returns a *JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumRef(s string) *JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	v := JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(s)
	return &v
}

func (v JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"OUTPUT_SCHEMA_UNSPECIFIED", "BASIC_COLUMNS", "GCS_COLUMNS", "DATASTORE_COLUMNS", "BIG_QUERY_COLUMNS", "ALL_COLUMNS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTriggerStatusEnum.
type JobTriggerStatusEnum string

// JobTriggerStatusEnumRef returns a *JobTriggerStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTriggerStatusEnumRef(s string) *JobTriggerStatusEnum {
	v := JobTriggerStatusEnum(s)
	return &v
}

func (v JobTriggerStatusEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATUS_UNSPECIFIED", "HEALTHY", "PAUSED", "CANCELLED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTriggerStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type JobTriggerInspectJob struct {
	empty               bool                               `json:"-"`
	StorageConfig       *JobTriggerInspectJobStorageConfig `json:"storageConfig"`
	InspectConfig       *JobTriggerInspectJobInspectConfig `json:"inspectConfig"`
	InspectTemplateName *string                            `json:"inspectTemplateName"`
	Actions             []JobTriggerInspectJobActions      `json:"actions"`
}

type jsonJobTriggerInspectJob JobTriggerInspectJob

func (r *JobTriggerInspectJob) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJob
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJob
	} else {

		r.StorageConfig = res.StorageConfig

		r.InspectConfig = res.InspectConfig

		r.InspectTemplateName = res.InspectTemplateName

		r.Actions = res.Actions

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJob is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJob *JobTriggerInspectJob = &JobTriggerInspectJob{empty: true}

func (r *JobTriggerInspectJob) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJob) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJob) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfig struct {
	empty               bool                                                  `json:"-"`
	DatastoreOptions    *JobTriggerInspectJobStorageConfigDatastoreOptions    `json:"datastoreOptions"`
	CloudStorageOptions *JobTriggerInspectJobStorageConfigCloudStorageOptions `json:"cloudStorageOptions"`
	BigQueryOptions     *JobTriggerInspectJobStorageConfigBigQueryOptions     `json:"bigQueryOptions"`
	HybridOptions       *JobTriggerInspectJobStorageConfigHybridOptions       `json:"hybridOptions"`
	TimespanConfig      *JobTriggerInspectJobStorageConfigTimespanConfig      `json:"timespanConfig"`
}

type jsonJobTriggerInspectJobStorageConfig JobTriggerInspectJobStorageConfig

func (r *JobTriggerInspectJobStorageConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfig
	} else {

		r.DatastoreOptions = res.DatastoreOptions

		r.CloudStorageOptions = res.CloudStorageOptions

		r.BigQueryOptions = res.BigQueryOptions

		r.HybridOptions = res.HybridOptions

		r.TimespanConfig = res.TimespanConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfig *JobTriggerInspectJobStorageConfig = &JobTriggerInspectJobStorageConfig{empty: true}

func (r *JobTriggerInspectJobStorageConfig) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigDatastoreOptions struct {
	empty       bool                                                          `json:"-"`
	PartitionId *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId `json:"partitionId"`
	Kind        *JobTriggerInspectJobStorageConfigDatastoreOptionsKind        `json:"kind"`
}

type jsonJobTriggerInspectJobStorageConfigDatastoreOptions JobTriggerInspectJobStorageConfigDatastoreOptions

func (r *JobTriggerInspectJobStorageConfigDatastoreOptions) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigDatastoreOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigDatastoreOptions
	} else {

		r.PartitionId = res.PartitionId

		r.Kind = res.Kind

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigDatastoreOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigDatastoreOptions *JobTriggerInspectJobStorageConfigDatastoreOptions = &JobTriggerInspectJobStorageConfigDatastoreOptions{empty: true}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptions) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId struct {
	empty       bool    `json:"-"`
	ProjectId   *string `json:"projectId"`
	NamespaceId *string `json:"namespaceId"`
}

type jsonJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId

func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId
	} else {

		r.ProjectId = res.ProjectId

		r.NamespaceId = res.NamespaceId

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId = &JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{empty: true}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigDatastoreOptionsKind struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonJobTriggerInspectJobStorageConfigDatastoreOptionsKind JobTriggerInspectJobStorageConfigDatastoreOptionsKind

func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsKind) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigDatastoreOptionsKind
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsKind
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigDatastoreOptionsKind is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsKind *JobTriggerInspectJobStorageConfigDatastoreOptionsKind = &JobTriggerInspectJobStorageConfigDatastoreOptionsKind{empty: true}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsKind) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsKind) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsKind) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigCloudStorageOptions struct {
	empty                    bool                                                                  `json:"-"`
	FileSet                  *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet          `json:"fileSet"`
	BytesLimitPerFile        *int64                                                                `json:"bytesLimitPerFile"`
	BytesLimitPerFilePercent *int64                                                                `json:"bytesLimitPerFilePercent"`
	FileTypes                []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum   `json:"fileTypes"`
	SampleMethod             *JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum `json:"sampleMethod"`
	FilesLimitPercent        *int64                                                                `json:"filesLimitPercent"`
}

type jsonJobTriggerInspectJobStorageConfigCloudStorageOptions JobTriggerInspectJobStorageConfigCloudStorageOptions

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptions) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigCloudStorageOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigCloudStorageOptions
	} else {

		r.FileSet = res.FileSet

		r.BytesLimitPerFile = res.BytesLimitPerFile

		r.BytesLimitPerFilePercent = res.BytesLimitPerFilePercent

		r.FileTypes = res.FileTypes

		r.SampleMethod = res.SampleMethod

		r.FilesLimitPercent = res.FilesLimitPercent

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigCloudStorageOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigCloudStorageOptions *JobTriggerInspectJobStorageConfigCloudStorageOptions = &JobTriggerInspectJobStorageConfigCloudStorageOptions{empty: true}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptions) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet struct {
	empty        bool                                                                     `json:"-"`
	Url          *string                                                                  `json:"url"`
	RegexFileSet *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet `json:"regexFileSet"`
}

type jsonJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet
	} else {

		r.Url = res.Url

		r.RegexFileSet = res.RegexFileSet

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet = &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{empty: true}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet struct {
	empty        bool     `json:"-"`
	BucketName   *string  `json:"bucketName"`
	IncludeRegex []string `json:"includeRegex"`
	ExcludeRegex []string `json:"excludeRegex"`
}

type jsonJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet
	} else {

		r.BucketName = res.BucketName

		r.IncludeRegex = res.IncludeRegex

		r.ExcludeRegex = res.ExcludeRegex

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet = &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{empty: true}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigBigQueryOptions struct {
	empty             bool                                                                `json:"-"`
	TableReference    *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference     `json:"tableReference"`
	IdentifyingFields []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields `json:"identifyingFields"`
	RowsLimit         *int64                                                              `json:"rowsLimit"`
	RowsLimitPercent  *int64                                                              `json:"rowsLimitPercent"`
	SampleMethod      *JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum   `json:"sampleMethod"`
	ExcludedFields    []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields    `json:"excludedFields"`
	IncludedFields    []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields    `json:"includedFields"`
}

type jsonJobTriggerInspectJobStorageConfigBigQueryOptions JobTriggerInspectJobStorageConfigBigQueryOptions

func (r *JobTriggerInspectJobStorageConfigBigQueryOptions) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigBigQueryOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigBigQueryOptions
	} else {

		r.TableReference = res.TableReference

		r.IdentifyingFields = res.IdentifyingFields

		r.RowsLimit = res.RowsLimit

		r.RowsLimitPercent = res.RowsLimitPercent

		r.SampleMethod = res.SampleMethod

		r.ExcludedFields = res.ExcludedFields

		r.IncludedFields = res.IncludedFields

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigBigQueryOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigBigQueryOptions *JobTriggerInspectJobStorageConfigBigQueryOptions = &JobTriggerInspectJobStorageConfigBigQueryOptions{empty: true}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptions) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference struct {
	empty     bool    `json:"-"`
	ProjectId *string `json:"projectId"`
	DatasetId *string `json:"datasetId"`
	TableId   *string `json:"tableId"`
}

type jsonJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
	} else {

		r.ProjectId = res.ProjectId

		r.DatasetId = res.DatasetId

		r.TableId = res.TableId

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference = &JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{empty: true}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields = &JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{empty: true}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields = &JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{empty: true}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields = &JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{empty: true}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigHybridOptions struct {
	empty                    bool                                                        `json:"-"`
	Description              *string                                                     `json:"description"`
	RequiredFindingLabelKeys []string                                                    `json:"requiredFindingLabelKeys"`
	Labels                   map[string]string                                           `json:"labels"`
	TableOptions             *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions `json:"tableOptions"`
}

type jsonJobTriggerInspectJobStorageConfigHybridOptions JobTriggerInspectJobStorageConfigHybridOptions

func (r *JobTriggerInspectJobStorageConfigHybridOptions) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigHybridOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigHybridOptions
	} else {

		r.Description = res.Description

		r.RequiredFindingLabelKeys = res.RequiredFindingLabelKeys

		r.Labels = res.Labels

		r.TableOptions = res.TableOptions

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigHybridOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigHybridOptions *JobTriggerInspectJobStorageConfigHybridOptions = &JobTriggerInspectJobStorageConfigHybridOptions{empty: true}

func (r *JobTriggerInspectJobStorageConfigHybridOptions) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigHybridOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigHybridOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigHybridOptionsTableOptions struct {
	empty             bool                                                                          `json:"-"`
	IdentifyingFields []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields `json:"identifyingFields"`
}

type jsonJobTriggerInspectJobStorageConfigHybridOptionsTableOptions JobTriggerInspectJobStorageConfigHybridOptionsTableOptions

func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigHybridOptionsTableOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptions
	} else {

		r.IdentifyingFields = res.IdentifyingFields

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigHybridOptionsTableOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptions *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions = &JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{empty: true}

func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields

func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields = &JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{empty: true}

func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigTimespanConfig struct {
	empty                                bool                                                           `json:"-"`
	StartTime                            *string                                                        `json:"startTime"`
	EndTime                              *string                                                        `json:"endTime"`
	TimestampField                       *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField `json:"timestampField"`
	EnableAutoPopulationOfTimespanConfig *bool                                                          `json:"enableAutoPopulationOfTimespanConfig"`
}

type jsonJobTriggerInspectJobStorageConfigTimespanConfig JobTriggerInspectJobStorageConfigTimespanConfig

func (r *JobTriggerInspectJobStorageConfigTimespanConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigTimespanConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigTimespanConfig
	} else {

		r.StartTime = res.StartTime

		r.EndTime = res.EndTime

		r.TimestampField = res.TimestampField

		r.EnableAutoPopulationOfTimespanConfig = res.EnableAutoPopulationOfTimespanConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigTimespanConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigTimespanConfig *JobTriggerInspectJobStorageConfigTimespanConfig = &JobTriggerInspectJobStorageConfigTimespanConfig{empty: true}

func (r *JobTriggerInspectJobStorageConfigTimespanConfig) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigTimespanConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigTimespanConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobStorageConfigTimespanConfigTimestampField struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonJobTriggerInspectJobStorageConfigTimespanConfigTimestampField JobTriggerInspectJobStorageConfigTimespanConfigTimestampField

func (r *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobStorageConfigTimespanConfigTimestampField
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobStorageConfigTimespanConfigTimestampField
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobStorageConfigTimespanConfigTimestampField is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobStorageConfigTimespanConfigTimestampField *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField = &JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{empty: true}

func (r *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfig struct {
	empty            bool                                                `json:"-"`
	InfoTypes        []JobTriggerInspectJobInspectConfigInfoTypes        `json:"infoTypes"`
	MinLikelihood    *JobTriggerInspectJobInspectConfigMinLikelihoodEnum `json:"minLikelihood"`
	Limits           *JobTriggerInspectJobInspectConfigLimits            `json:"limits"`
	IncludeQuote     *bool                                               `json:"includeQuote"`
	ExcludeInfoTypes *bool                                               `json:"excludeInfoTypes"`
	CustomInfoTypes  []JobTriggerInspectJobInspectConfigCustomInfoTypes  `json:"customInfoTypes"`
	RuleSet          []JobTriggerInspectJobInspectConfigRuleSet          `json:"ruleSet"`
}

type jsonJobTriggerInspectJobInspectConfig JobTriggerInspectJobInspectConfig

func (r *JobTriggerInspectJobInspectConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfig
	} else {

		r.InfoTypes = res.InfoTypes

		r.MinLikelihood = res.MinLikelihood

		r.Limits = res.Limits

		r.IncludeQuote = res.IncludeQuote

		r.ExcludeInfoTypes = res.ExcludeInfoTypes

		r.CustomInfoTypes = res.CustomInfoTypes

		r.RuleSet = res.RuleSet

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfig *JobTriggerInspectJobInspectConfig = &JobTriggerInspectJobInspectConfig{empty: true}

func (r *JobTriggerInspectJobInspectConfig) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigInfoTypes struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonJobTriggerInspectJobInspectConfigInfoTypes JobTriggerInspectJobInspectConfigInfoTypes

func (r *JobTriggerInspectJobInspectConfigInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigInfoTypes
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigInfoTypes *JobTriggerInspectJobInspectConfigInfoTypes = &JobTriggerInspectJobInspectConfigInfoTypes{empty: true}

func (r *JobTriggerInspectJobInspectConfigInfoTypes) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigLimits struct {
	empty                  bool                                                            `json:"-"`
	MaxFindingsPerItem     *int64                                                          `json:"maxFindingsPerItem"`
	MaxFindingsPerRequest  *int64                                                          `json:"maxFindingsPerRequest"`
	MaxFindingsPerInfoType []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType `json:"maxFindingsPerInfoType"`
}

type jsonJobTriggerInspectJobInspectConfigLimits JobTriggerInspectJobInspectConfigLimits

func (r *JobTriggerInspectJobInspectConfigLimits) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigLimits
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigLimits
	} else {

		r.MaxFindingsPerItem = res.MaxFindingsPerItem

		r.MaxFindingsPerRequest = res.MaxFindingsPerRequest

		r.MaxFindingsPerInfoType = res.MaxFindingsPerInfoType

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigLimits is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigLimits *JobTriggerInspectJobInspectConfigLimits = &JobTriggerInspectJobInspectConfigLimits{empty: true}

func (r *JobTriggerInspectJobInspectConfigLimits) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigLimits) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigLimits) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType struct {
	empty       bool                                                                   `json:"-"`
	InfoType    *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `json:"infoType"`
	MaxFindings *int64                                                                 `json:"maxFindings"`
}

type jsonJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType

func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType
	} else {

		r.InfoType = res.InfoType

		r.MaxFindings = res.MaxFindings

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType = &JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{empty: true}

func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType struct {
	empty   bool    `json:"-"`
	Name    *string `json:"name"`
	Version *string `json:"version"`
}

type jsonJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType

func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	} else {

		r.Name = res.Name

		r.Version = res.Version

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType = &JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{empty: true}

func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypes struct {
	empty          bool                                                               `json:"-"`
	InfoType       *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType          `json:"infoType"`
	Likelihood     *JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum    `json:"likelihood"`
	Dictionary     *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary        `json:"dictionary"`
	Regex          *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex             `json:"regex"`
	SurrogateType  *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType     `json:"surrogateType"`
	StoredType     *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType        `json:"storedType"`
	DetectionRules []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules   `json:"detectionRules"`
	ExclusionType  *JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum `json:"exclusionType"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypes JobTriggerInspectJobInspectConfigCustomInfoTypes

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypes
	} else {

		r.InfoType = res.InfoType

		r.Likelihood = res.Likelihood

		r.Dictionary = res.Dictionary

		r.Regex = res.Regex

		r.SurrogateType = res.SurrogateType

		r.StoredType = res.StoredType

		r.DetectionRules = res.DetectionRules

		r.ExclusionType = res.ExclusionType

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypes *JobTriggerInspectJobInspectConfigCustomInfoTypes = &JobTriggerInspectJobInspectConfigCustomInfoTypes{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypes) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType struct {
	empty   bool    `json:"-"`
	Name    *string `json:"name"`
	Version *string `json:"version"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType
	} else {

		r.Name = res.Name

		r.Version = res.Version

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType = &JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary struct {
	empty            bool                                                                        `json:"-"`
	WordList         *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList         `json:"wordList"`
	CloudStoragePath *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath `json:"cloudStoragePath"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary
	} else {

		r.WordList = res.WordList

		r.CloudStoragePath = res.CloudStoragePath

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList struct {
	empty bool     `json:"-"`
	Words []string `json:"words"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList
	} else {

		r.Words = res.Words

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath struct {
	empty bool    `json:"-"`
	Path  *string `json:"path"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	} else {

		r.Path = res.Path

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesRegex struct {
	empty        bool    `json:"-"`
	Pattern      *string `json:"pattern"`
	GroupIndexes []int64 `json:"groupIndexes"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesRegex JobTriggerInspectJobInspectConfigCustomInfoTypesRegex

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesRegex
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesRegex
	} else {

		r.Pattern = res.Pattern

		r.GroupIndexes = res.GroupIndexes

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesRegex is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesRegex *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex = &JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType struct {
	empty bool `json:"-"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType = &JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType struct {
	empty      bool    `json:"-"`
	Name       *string `json:"name"`
	CreateTime *string `json:"createTime"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType
	} else {

		r.Name = res.Name

		r.CreateTime = res.CreateTime

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType = &JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules struct {
	empty       bool                                                                       `json:"-"`
	HotwordRule *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule `json:"hotwordRule"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules
	} else {

		r.HotwordRule = res.HotwordRule

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule struct {
	empty                bool                                                                                           `json:"-"`
	HotwordRegex         *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex         `json:"hotwordRegex"`
	Proximity            *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity            `json:"proximity"`
	LikelihoodAdjustment *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment `json:"likelihoodAdjustment"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule
	} else {

		r.HotwordRegex = res.HotwordRegex

		r.Proximity = res.Proximity

		r.LikelihoodAdjustment = res.LikelihoodAdjustment

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex struct {
	empty        bool    `json:"-"`
	Pattern      *string `json:"pattern"`
	GroupIndexes []int64 `json:"groupIndexes"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex
	} else {

		r.Pattern = res.Pattern

		r.GroupIndexes = res.GroupIndexes

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity struct {
	empty        bool   `json:"-"`
	WindowBefore *int64 `json:"windowBefore"`
	WindowAfter  *int64 `json:"windowAfter"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
	} else {

		r.WindowBefore = res.WindowBefore

		r.WindowAfter = res.WindowAfter

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment struct {
	empty              bool                                                                                                              `json:"-"`
	FixedLikelihood    *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum `json:"fixedLikelihood"`
	RelativeLikelihood *int64                                                                                                            `json:"relativeLikelihood"`
}

type jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment
	} else {

		r.FixedLikelihood = res.FixedLikelihood

		r.RelativeLikelihood = res.RelativeLikelihood

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{empty: true}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSet struct {
	empty     bool                                                `json:"-"`
	InfoTypes []JobTriggerInspectJobInspectConfigRuleSetInfoTypes `json:"infoTypes"`
	Rules     []JobTriggerInspectJobInspectConfigRuleSetRules     `json:"rules"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSet JobTriggerInspectJobInspectConfigRuleSet

func (r *JobTriggerInspectJobInspectConfigRuleSet) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSet
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSet
	} else {

		r.InfoTypes = res.InfoTypes

		r.Rules = res.Rules

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSet is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSet *JobTriggerInspectJobInspectConfigRuleSet = &JobTriggerInspectJobInspectConfigRuleSet{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSet) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSet) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSet) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetInfoTypes struct {
	empty   bool    `json:"-"`
	Name    *string `json:"name"`
	Version *string `json:"version"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetInfoTypes JobTriggerInspectJobInspectConfigRuleSetInfoTypes

func (r *JobTriggerInspectJobInspectConfigRuleSetInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetInfoTypes
	} else {

		r.Name = res.Name

		r.Version = res.Version

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetInfoTypes *JobTriggerInspectJobInspectConfigRuleSetInfoTypes = &JobTriggerInspectJobInspectConfigRuleSetInfoTypes{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetInfoTypes) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRules struct {
	empty         bool                                                        `json:"-"`
	HotwordRule   *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule   `json:"hotwordRule"`
	ExclusionRule *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule `json:"exclusionRule"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRules JobTriggerInspectJobInspectConfigRuleSetRules

func (r *JobTriggerInspectJobInspectConfigRuleSetRules) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRules
	} else {

		r.HotwordRule = res.HotwordRule

		r.ExclusionRule = res.ExclusionRule

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRules *JobTriggerInspectJobInspectConfigRuleSetRules = &JobTriggerInspectJobInspectConfigRuleSetRules{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRules) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRules) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule struct {
	empty                bool                                                                          `json:"-"`
	HotwordRegex         *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex         `json:"hotwordRegex"`
	Proximity            *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity            `json:"proximity"`
	LikelihoodAdjustment *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment `json:"likelihoodAdjustment"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	} else {

		r.HotwordRegex = res.HotwordRegex

		r.Proximity = res.Proximity

		r.LikelihoodAdjustment = res.LikelihoodAdjustment

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex struct {
	empty        bool    `json:"-"`
	Pattern      *string `json:"pattern"`
	GroupIndexes []int64 `json:"groupIndexes"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	} else {

		r.Pattern = res.Pattern

		r.GroupIndexes = res.GroupIndexes

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity struct {
	empty        bool   `json:"-"`
	WindowBefore *int64 `json:"windowBefore"`
	WindowAfter  *int64 `json:"windowAfter"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity
	} else {

		r.WindowBefore = res.WindowBefore

		r.WindowAfter = res.WindowAfter

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment struct {
	empty              bool                                                                                             `json:"-"`
	FixedLikelihood    *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum `json:"fixedLikelihood"`
	RelativeLikelihood *int64                                                                                           `json:"relativeLikelihood"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	} else {

		r.FixedLikelihood = res.FixedLikelihood

		r.RelativeLikelihood = res.RelativeLikelihood

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule struct {
	empty            bool                                                                        `json:"-"`
	Dictionary       *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary       `json:"dictionary"`
	Regex            *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex            `json:"regex"`
	ExcludeInfoTypes *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes `json:"excludeInfoTypes"`
	MatchingType     *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum `json:"matchingType"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	} else {

		r.Dictionary = res.Dictionary

		r.Regex = res.Regex

		r.ExcludeInfoTypes = res.ExcludeInfoTypes

		r.MatchingType = res.MatchingType

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary struct {
	empty            bool                                                                                  `json:"-"`
	WordList         *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList         `json:"wordList"`
	CloudStoragePath *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath `json:"cloudStoragePath"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	} else {

		r.WordList = res.WordList

		r.CloudStoragePath = res.CloudStoragePath

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList struct {
	empty bool     `json:"-"`
	Words []string `json:"words"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	} else {

		r.Words = res.Words

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath struct {
	empty bool    `json:"-"`
	Path  *string `json:"path"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	} else {

		r.Path = res.Path

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex struct {
	empty        bool    `json:"-"`
	Pattern      *string `json:"pattern"`
	GroupIndexes []int64 `json:"groupIndexes"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
	} else {

		r.Pattern = res.Pattern

		r.GroupIndexes = res.GroupIndexes

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes struct {
	empty     bool                                                                                  `json:"-"`
	InfoTypes []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes `json:"infoTypes"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	} else {

		r.InfoTypes = res.InfoTypes

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes struct {
	empty   bool    `json:"-"`
	Name    *string `json:"name"`
	Version *string `json:"version"`
}

type jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	} else {

		r.Name = res.Name

		r.Version = res.Version

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{empty: true}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActions struct {
	empty                             bool                                                          `json:"-"`
	SaveFindings                      *JobTriggerInspectJobActionsSaveFindings                      `json:"saveFindings"`
	PubSub                            *JobTriggerInspectJobActionsPubSub                            `json:"pubSub"`
	PublishSummaryToCscc              *JobTriggerInspectJobActionsPublishSummaryToCscc              `json:"publishSummaryToCscc"`
	PublishFindingsToCloudDataCatalog *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog `json:"publishFindingsToCloudDataCatalog"`
	JobNotificationEmails             *JobTriggerInspectJobActionsJobNotificationEmails             `json:"jobNotificationEmails"`
	PublishToStackdriver              *JobTriggerInspectJobActionsPublishToStackdriver              `json:"publishToStackdriver"`
}

type jsonJobTriggerInspectJobActions JobTriggerInspectJobActions

func (r *JobTriggerInspectJobActions) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActions
	} else {

		r.SaveFindings = res.SaveFindings

		r.PubSub = res.PubSub

		r.PublishSummaryToCscc = res.PublishSummaryToCscc

		r.PublishFindingsToCloudDataCatalog = res.PublishFindingsToCloudDataCatalog

		r.JobNotificationEmails = res.JobNotificationEmails

		r.PublishToStackdriver = res.PublishToStackdriver

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActions *JobTriggerInspectJobActions = &JobTriggerInspectJobActions{empty: true}

func (r *JobTriggerInspectJobActions) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActions) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsSaveFindings struct {
	empty        bool                                                 `json:"-"`
	OutputConfig *JobTriggerInspectJobActionsSaveFindingsOutputConfig `json:"outputConfig"`
}

type jsonJobTriggerInspectJobActionsSaveFindings JobTriggerInspectJobActionsSaveFindings

func (r *JobTriggerInspectJobActionsSaveFindings) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsSaveFindings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsSaveFindings
	} else {

		r.OutputConfig = res.OutputConfig

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsSaveFindings is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsSaveFindings *JobTriggerInspectJobActionsSaveFindings = &JobTriggerInspectJobActionsSaveFindings{empty: true}

func (r *JobTriggerInspectJobActionsSaveFindings) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsSaveFindings) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsSaveFindings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsSaveFindingsOutputConfig struct {
	empty        bool                                                                 `json:"-"`
	Table        *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable            `json:"table"`
	DlpStorage   *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage       `json:"dlpStorage"`
	OutputSchema *JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum `json:"outputSchema"`
}

type jsonJobTriggerInspectJobActionsSaveFindingsOutputConfig JobTriggerInspectJobActionsSaveFindingsOutputConfig

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfig) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsSaveFindingsOutputConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfig
	} else {

		r.Table = res.Table

		r.DlpStorage = res.DlpStorage

		r.OutputSchema = res.OutputSchema

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsSaveFindingsOutputConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfig *JobTriggerInspectJobActionsSaveFindingsOutputConfig = &JobTriggerInspectJobActionsSaveFindingsOutputConfig{empty: true}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfig) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsSaveFindingsOutputConfigTable struct {
	empty     bool    `json:"-"`
	ProjectId *string `json:"projectId"`
	DatasetId *string `json:"datasetId"`
	TableId   *string `json:"tableId"`
}

type jsonJobTriggerInspectJobActionsSaveFindingsOutputConfigTable JobTriggerInspectJobActionsSaveFindingsOutputConfigTable

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsSaveFindingsOutputConfigTable
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigTable
	} else {

		r.ProjectId = res.ProjectId

		r.DatasetId = res.DatasetId

		r.TableId = res.TableId

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsSaveFindingsOutputConfigTable is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigTable *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable = &JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{empty: true}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage struct {
	empty bool `json:"-"`
}

type jsonJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage = &JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{empty: true}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsPubSub struct {
	empty bool    `json:"-"`
	Topic *string `json:"topic"`
}

type jsonJobTriggerInspectJobActionsPubSub JobTriggerInspectJobActionsPubSub

func (r *JobTriggerInspectJobActionsPubSub) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsPubSub
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsPubSub
	} else {

		r.Topic = res.Topic

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsPubSub is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsPubSub *JobTriggerInspectJobActionsPubSub = &JobTriggerInspectJobActionsPubSub{empty: true}

func (r *JobTriggerInspectJobActionsPubSub) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsPubSub) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsPubSub) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsPublishSummaryToCscc struct {
	empty bool `json:"-"`
}

type jsonJobTriggerInspectJobActionsPublishSummaryToCscc JobTriggerInspectJobActionsPublishSummaryToCscc

func (r *JobTriggerInspectJobActionsPublishSummaryToCscc) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsPublishSummaryToCscc
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsPublishSummaryToCscc
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsPublishSummaryToCscc is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsPublishSummaryToCscc *JobTriggerInspectJobActionsPublishSummaryToCscc = &JobTriggerInspectJobActionsPublishSummaryToCscc{empty: true}

func (r *JobTriggerInspectJobActionsPublishSummaryToCscc) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsPublishSummaryToCscc) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsPublishSummaryToCscc) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog struct {
	empty bool `json:"-"`
}

type jsonJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog

func (r *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog = &JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{empty: true}

func (r *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsJobNotificationEmails struct {
	empty bool `json:"-"`
}

type jsonJobTriggerInspectJobActionsJobNotificationEmails JobTriggerInspectJobActionsJobNotificationEmails

func (r *JobTriggerInspectJobActionsJobNotificationEmails) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsJobNotificationEmails
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsJobNotificationEmails
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsJobNotificationEmails is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsJobNotificationEmails *JobTriggerInspectJobActionsJobNotificationEmails = &JobTriggerInspectJobActionsJobNotificationEmails{empty: true}

func (r *JobTriggerInspectJobActionsJobNotificationEmails) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsJobNotificationEmails) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsJobNotificationEmails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerInspectJobActionsPublishToStackdriver struct {
	empty bool `json:"-"`
}

type jsonJobTriggerInspectJobActionsPublishToStackdriver JobTriggerInspectJobActionsPublishToStackdriver

func (r *JobTriggerInspectJobActionsPublishToStackdriver) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerInspectJobActionsPublishToStackdriver
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerInspectJobActionsPublishToStackdriver
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerInspectJobActionsPublishToStackdriver is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerInspectJobActionsPublishToStackdriver *JobTriggerInspectJobActionsPublishToStackdriver = &JobTriggerInspectJobActionsPublishToStackdriver{empty: true}

func (r *JobTriggerInspectJobActionsPublishToStackdriver) Empty() bool {
	return r.empty
}

func (r *JobTriggerInspectJobActionsPublishToStackdriver) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerInspectJobActionsPublishToStackdriver) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerTriggers struct {
	empty    bool                        `json:"-"`
	Schedule *JobTriggerTriggersSchedule `json:"schedule"`
	Manual   *JobTriggerTriggersManual   `json:"manual"`
}

type jsonJobTriggerTriggers JobTriggerTriggers

func (r *JobTriggerTriggers) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerTriggers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerTriggers
	} else {

		r.Schedule = res.Schedule

		r.Manual = res.Manual

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerTriggers is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerTriggers *JobTriggerTriggers = &JobTriggerTriggers{empty: true}

func (r *JobTriggerTriggers) Empty() bool {
	return r.empty
}

func (r *JobTriggerTriggers) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerTriggers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerTriggersSchedule struct {
	empty                    bool    `json:"-"`
	RecurrencePeriodDuration *string `json:"recurrencePeriodDuration"`
}

type jsonJobTriggerTriggersSchedule JobTriggerTriggersSchedule

func (r *JobTriggerTriggersSchedule) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerTriggersSchedule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerTriggersSchedule
	} else {

		r.RecurrencePeriodDuration = res.RecurrencePeriodDuration

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerTriggersSchedule is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerTriggersSchedule *JobTriggerTriggersSchedule = &JobTriggerTriggersSchedule{empty: true}

func (r *JobTriggerTriggersSchedule) Empty() bool {
	return r.empty
}

func (r *JobTriggerTriggersSchedule) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerTriggersSchedule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerTriggersManual struct {
	empty bool `json:"-"`
}

type jsonJobTriggerTriggersManual JobTriggerTriggersManual

func (r *JobTriggerTriggersManual) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerTriggersManual
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerTriggersManual
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerTriggersManual is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerTriggersManual *JobTriggerTriggersManual = &JobTriggerTriggersManual{empty: true}

func (r *JobTriggerTriggersManual) Empty() bool {
	return r.empty
}

func (r *JobTriggerTriggersManual) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerTriggersManual) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerErrors struct {
	empty      bool                     `json:"-"`
	Details    *JobTriggerErrorsDetails `json:"details"`
	Timestamps []string                 `json:"timestamps"`
}

type jsonJobTriggerErrors JobTriggerErrors

func (r *JobTriggerErrors) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerErrors
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerErrors
	} else {

		r.Details = res.Details

		r.Timestamps = res.Timestamps

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerErrors is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerErrors *JobTriggerErrors = &JobTriggerErrors{empty: true}

func (r *JobTriggerErrors) Empty() bool {
	return r.empty
}

func (r *JobTriggerErrors) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerErrors) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerErrorsDetails struct {
	empty   bool                             `json:"-"`
	Code    *int64                           `json:"code"`
	Message *string                          `json:"message"`
	Details []JobTriggerErrorsDetailsDetails `json:"details"`
}

type jsonJobTriggerErrorsDetails JobTriggerErrorsDetails

func (r *JobTriggerErrorsDetails) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerErrorsDetails
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerErrorsDetails
	} else {

		r.Code = res.Code

		r.Message = res.Message

		r.Details = res.Details

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerErrorsDetails is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerErrorsDetails *JobTriggerErrorsDetails = &JobTriggerErrorsDetails{empty: true}

func (r *JobTriggerErrorsDetails) Empty() bool {
	return r.empty
}

func (r *JobTriggerErrorsDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerErrorsDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTriggerErrorsDetailsDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

type jsonJobTriggerErrorsDetailsDetails JobTriggerErrorsDetailsDetails

func (r *JobTriggerErrorsDetailsDetails) UnmarshalJSON(data []byte) error {
	var res jsonJobTriggerErrorsDetailsDetails
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTriggerErrorsDetailsDetails
	} else {

		r.TypeUrl = res.TypeUrl

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this JobTriggerErrorsDetailsDetails is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTriggerErrorsDetailsDetails *JobTriggerErrorsDetailsDetails = &JobTriggerErrorsDetailsDetails{empty: true}

func (r *JobTriggerErrorsDetailsDetails) Empty() bool {
	return r.empty
}

func (r *JobTriggerErrorsDetailsDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTriggerErrorsDetailsDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *JobTrigger) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dlp",
		Type:    "JobTrigger",
		Version: "alpha",
	}
}

func (r *JobTrigger) ID() (string, error) {
	if err := extractJobTriggerFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":          dcl.ValueOrEmptyString(nr.Name),
		"display_name":  dcl.ValueOrEmptyString(nr.DisplayName),
		"description":   dcl.ValueOrEmptyString(nr.Description),
		"inspect_job":   dcl.ValueOrEmptyString(nr.InspectJob),
		"triggers":      dcl.ValueOrEmptyString(nr.Triggers),
		"errors":        dcl.ValueOrEmptyString(nr.Errors),
		"create_time":   dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":   dcl.ValueOrEmptyString(nr.UpdateTime),
		"last_run_time": dcl.ValueOrEmptyString(nr.LastRunTime),
		"status":        dcl.ValueOrEmptyString(nr.Status),
		"location_id":   dcl.ValueOrEmptyString(nr.LocationId),
		"parent":        dcl.ValueOrEmptyString(nr.Parent),
		"location":      dcl.ValueOrEmptyString(nr.Location),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.Nprintf("{{parent}}/locations/{{location}}/jobTriggers/{{name}}", params), nil
	}

	return dcl.Nprintf("{{parent}}/jobTriggers/{{name}}", params), nil
}

const JobTriggerMaxPage = -1

type JobTriggerList struct {
	Items []*JobTrigger

	nextToken string

	pageSize int32

	resource *JobTrigger
}

func (l *JobTriggerList) HasNext() bool {
	return l.nextToken != ""
}

func (l *JobTriggerList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listJobTrigger(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListJobTrigger(ctx context.Context, location, parent string) (*JobTriggerList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListJobTriggerWithMaxResults(ctx, location, parent, JobTriggerMaxPage)

}

func (c *Client) ListJobTriggerWithMaxResults(ctx context.Context, location, parent string, pageSize int32) (*JobTriggerList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &JobTrigger{
		Location: &location,
		Parent:   &parent,
	}
	items, token, err := c.listJobTrigger(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &JobTriggerList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetJobTrigger(ctx context.Context, r *JobTrigger) (*JobTrigger, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractJobTriggerFields(r)

	b, err := c.getJobTriggerRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalJobTrigger(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Location = r.Location
	result.Parent = r.Parent
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeJobTriggerNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractJobTriggerFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteJobTrigger(ctx context.Context, r *JobTrigger) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("JobTrigger resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting JobTrigger...")
	deleteOp := deleteJobTriggerOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllJobTrigger deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllJobTrigger(ctx context.Context, location, parent string, filter func(*JobTrigger) bool) error {
	listObj, err := c.ListJobTrigger(ctx, location, parent)
	if err != nil {
		return err
	}

	err = c.deleteAllJobTrigger(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllJobTrigger(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyJobTrigger(ctx context.Context, rawDesired *JobTrigger, opts ...dcl.ApplyOption) (*JobTrigger, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *JobTrigger
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyJobTriggerHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyJobTriggerHelper(c *Client, ctx context.Context, rawDesired *JobTrigger, opts ...dcl.ApplyOption) (*JobTrigger, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyJobTrigger...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractJobTriggerFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.jobTriggerDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToJobTriggerDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []jobTriggerApiOperation
	if create {
		ops = append(ops, &createJobTriggerOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyJobTriggerDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyJobTriggerDiff(c *Client, ctx context.Context, desired *JobTrigger, rawDesired *JobTrigger, ops []jobTriggerApiOperation, opts ...dcl.ApplyOption) (*JobTrigger, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetJobTrigger(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createJobTriggerOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapJobTrigger(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeJobTriggerNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeJobTriggerNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeJobTriggerDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractJobTriggerFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractJobTriggerFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffJobTrigger(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
