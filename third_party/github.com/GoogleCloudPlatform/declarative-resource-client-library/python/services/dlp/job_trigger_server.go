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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dlppb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dlp/dlp_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp"
)

// JobTriggerServer implements the gRPC interface for JobTrigger.
type JobTriggerServer struct{}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(e dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum) *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(n[len("DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(e dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum) *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(n[len("DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum converts a JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(e dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum) *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(n[len("DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigMinLikelihoodEnum converts a JobTriggerInspectJobInspectConfigMinLikelihoodEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum(e dlppb.DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum) *dlp.JobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobInspectConfigMinLikelihoodEnum(n[len("DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(e dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(n[len("DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(e dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(n[len("DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(e dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(n[len("DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum enum from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(e dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum) *dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(n[len("DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerStatusEnum converts a JobTriggerStatusEnum enum from its proto representation.
func ProtoToDlpJobTriggerStatusEnum(e dlppb.DlpJobTriggerStatusEnum) *dlp.JobTriggerStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpJobTriggerStatusEnum_name[int32(e)]; ok {
		e := dlp.JobTriggerStatusEnum(n[len("DlpJobTriggerStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJob converts a JobTriggerInspectJob object from its proto representation.
func ProtoToDlpJobTriggerInspectJob(p *dlppb.DlpJobTriggerInspectJob) *dlp.JobTriggerInspectJob {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJob{
		StorageConfig:       ProtoToDlpJobTriggerInspectJobStorageConfig(p.GetStorageConfig()),
		InspectConfig:       ProtoToDlpJobTriggerInspectJobInspectConfig(p.GetInspectConfig()),
		InspectTemplateName: dcl.StringOrNil(p.GetInspectTemplateName()),
	}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, *ProtoToDlpJobTriggerInspectJobActions(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfig converts a JobTriggerInspectJobStorageConfig object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfig(p *dlppb.DlpJobTriggerInspectJobStorageConfig) *dlp.JobTriggerInspectJobStorageConfig {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfig{
		DatastoreOptions:    ProtoToDlpJobTriggerInspectJobStorageConfigDatastoreOptions(p.GetDatastoreOptions()),
		CloudStorageOptions: ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptions(p.GetCloudStorageOptions()),
		BigQueryOptions:     ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptions(p.GetBigQueryOptions()),
		HybridOptions:       ProtoToDlpJobTriggerInspectJobStorageConfigHybridOptions(p.GetHybridOptions()),
		TimespanConfig:      ProtoToDlpJobTriggerInspectJobStorageConfigTimespanConfig(p.GetTimespanConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptions converts a JobTriggerInspectJobStorageConfigDatastoreOptions object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigDatastoreOptions(p *dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptions) *dlp.JobTriggerInspectJobStorageConfigDatastoreOptions {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigDatastoreOptions{
		PartitionId: ProtoToDlpJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(p.GetPartitionId()),
		Kind:        ProtoToDlpJobTriggerInspectJobStorageConfigDatastoreOptionsKind(p.GetKind()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId converts a JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(p *dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) *dlp.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{
		ProjectId:   dcl.StringOrNil(p.GetProjectId()),
		NamespaceId: dcl.StringOrNil(p.GetNamespaceId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptionsKind converts a JobTriggerInspectJobStorageConfigDatastoreOptionsKind object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigDatastoreOptionsKind(p *dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptionsKind) *dlp.JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigDatastoreOptionsKind{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptions converts a JobTriggerInspectJobStorageConfigCloudStorageOptions object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptions(p *dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptions) *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptions {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigCloudStorageOptions{
		FileSet:                  ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(p.GetFileSet()),
		BytesLimitPerFile:        dcl.Int64OrNil(p.GetBytesLimitPerFile()),
		BytesLimitPerFilePercent: dcl.Int64OrNil(p.GetBytesLimitPerFilePercent()),
		SampleMethod:             ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(p.GetSampleMethod()),
		FilesLimitPercent:        dcl.Int64OrNil(p.GetFilesLimitPercent()),
	}
	for _, r := range p.GetFileTypes() {
		obj.FileTypes = append(obj.FileTypes, *ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(p *dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{
		Url:          dcl.StringOrNil(p.GetUrl()),
		RegexFileSet: ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(p.GetRegexFileSet()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(p *dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{
		BucketName: dcl.StringOrNil(p.GetBucketName()),
	}
	for _, r := range p.GetIncludeRegex() {
		obj.IncludeRegex = append(obj.IncludeRegex, r)
	}
	for _, r := range p.GetExcludeRegex() {
		obj.ExcludeRegex = append(obj.ExcludeRegex, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptions converts a JobTriggerInspectJobStorageConfigBigQueryOptions object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptions(p *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptions) *dlp.JobTriggerInspectJobStorageConfigBigQueryOptions {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigBigQueryOptions{
		TableReference:   ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(p.GetTableReference()),
		RowsLimit:        dcl.Int64OrNil(p.GetRowsLimit()),
		RowsLimitPercent: dcl.Int64OrNil(p.GetRowsLimitPercent()),
		SampleMethod:     ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(p.GetSampleMethod()),
	}
	for _, r := range p.GetIdentifyingFields() {
		obj.IdentifyingFields = append(obj.IdentifyingFields, *ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(r))
	}
	for _, r := range p.GetExcludedFields() {
		obj.ExcludedFields = append(obj.ExcludedFields, *ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(r))
	}
	for _, r := range p.GetIncludedFields() {
		obj.IncludedFields = append(obj.IncludedFields, *ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference converts a JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(p *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(p *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(p *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(p *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptions converts a JobTriggerInspectJobStorageConfigHybridOptions object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigHybridOptions(p *dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptions) *dlp.JobTriggerInspectJobStorageConfigHybridOptions {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigHybridOptions{
		Description:  dcl.StringOrNil(p.GetDescription()),
		TableOptions: ProtoToDlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(p.GetTableOptions()),
	}
	for _, r := range p.GetRequiredFindingLabelKeys() {
		obj.RequiredFindingLabelKeys = append(obj.RequiredFindingLabelKeys, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptionsTableOptions converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptions object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(p *dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptions) *dlp.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	for _, r := range p.GetIdentifyingFields() {
		obj.IdentifyingFields = append(obj.IdentifyingFields, *ProtoToDlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(p *dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) *dlp.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigTimespanConfig converts a JobTriggerInspectJobStorageConfigTimespanConfig object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigTimespanConfig(p *dlppb.DlpJobTriggerInspectJobStorageConfigTimespanConfig) *dlp.JobTriggerInspectJobStorageConfigTimespanConfig {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigTimespanConfig{
		StartTime:                            dcl.StringOrNil(p.GetStartTime()),
		EndTime:                              dcl.StringOrNil(p.GetEndTime()),
		TimestampField:                       ProtoToDlpJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(p.GetTimestampField()),
		EnableAutoPopulationOfTimespanConfig: dcl.Bool(p.GetEnableAutoPopulationOfTimespanConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigTimespanConfigTimestampField converts a JobTriggerInspectJobStorageConfigTimespanConfigTimestampField object from its proto representation.
func ProtoToDlpJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(p *dlppb.DlpJobTriggerInspectJobStorageConfigTimespanConfigTimestampField) *dlp.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfig converts a JobTriggerInspectJobInspectConfig object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfig(p *dlppb.DlpJobTriggerInspectJobInspectConfig) *dlp.JobTriggerInspectJobInspectConfig {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfig{
		MinLikelihood:    ProtoToDlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum(p.GetMinLikelihood()),
		Limits:           ProtoToDlpJobTriggerInspectJobInspectConfigLimits(p.GetLimits()),
		IncludeQuote:     dcl.Bool(p.GetIncludeQuote()),
		ExcludeInfoTypes: dcl.Bool(p.GetExcludeInfoTypes()),
	}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpJobTriggerInspectJobInspectConfigInfoTypes(r))
	}
	for _, r := range p.GetCustomInfoTypes() {
		obj.CustomInfoTypes = append(obj.CustomInfoTypes, *ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypes(r))
	}
	for _, r := range p.GetRuleSet() {
		obj.RuleSet = append(obj.RuleSet, *ProtoToDlpJobTriggerInspectJobInspectConfigRuleSet(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigInfoTypes converts a JobTriggerInspectJobInspectConfigInfoTypes object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigInfoTypes(p *dlppb.DlpJobTriggerInspectJobInspectConfigInfoTypes) *dlp.JobTriggerInspectJobInspectConfigInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimits converts a JobTriggerInspectJobInspectConfigLimits object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigLimits(p *dlppb.DlpJobTriggerInspectJobInspectConfigLimits) *dlp.JobTriggerInspectJobInspectConfigLimits {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigLimits{
		MaxFindingsPerItem:    dcl.Int64OrNil(p.GetMaxFindingsPerItem()),
		MaxFindingsPerRequest: dcl.Int64OrNil(p.GetMaxFindingsPerRequest()),
	}
	for _, r := range p.GetMaxFindingsPerInfoType() {
		obj.MaxFindingsPerInfoType = append(obj.MaxFindingsPerInfoType, *ProtoToDlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(p *dlppb.DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) *dlp.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{
		InfoType:    ProtoToDlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p.GetInfoType()),
		MaxFindings: dcl.Int64OrNil(p.GetMaxFindings()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p *dlppb.DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *dlp.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypes converts a JobTriggerInspectJobInspectConfigCustomInfoTypes object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypes(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypes) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypes{
		InfoType:      ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(p.GetInfoType()),
		Likelihood:    ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(p.GetLikelihood()),
		Dictionary:    ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(p.GetDictionary()),
		Regex:         ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(p.GetRegex()),
		SurrogateType: ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(p.GetSurrogateType()),
		StoredType:    ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(p.GetStoredType()),
		ExclusionType: ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(p.GetExclusionType()),
	}
	for _, r := range p.GetDetectionRules() {
		obj.DetectionRules = append(obj.DetectionRules, *ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{
		WordList:         ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesRegex converts a JobTriggerInspectJobInspectConfigCustomInfoTypesRegex object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesRegex) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{
		HotwordRule: ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(p.GetHotwordRule()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{
		HotwordRegex:         ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(p *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSet converts a JobTriggerInspectJobInspectConfigRuleSet object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSet(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSet) *dlp.JobTriggerInspectJobInspectConfigRuleSet {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSet{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetInfoTypes(r))
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRules(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetInfoTypes object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetInfoTypes(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetInfoTypes) *dlp.JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetInfoTypes{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRules converts a JobTriggerInspectJobInspectConfigRuleSetRules object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRules(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRules) *dlp.JobTriggerInspectJobInspectConfigRuleSetRules {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRules{
		HotwordRule:   ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(p.GetHotwordRule()),
		ExclusionRule: ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(p.GetExclusionRule()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{
		HotwordRegex:         ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{
		Dictionary:       ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(p.GetDictionary()),
		Regex:            ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(p.GetRegex()),
		ExcludeInfoTypes: ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p.GetExcludeInfoTypes()),
		MatchingType:     ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(p.GetMatchingType()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{
		WordList:         ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object from its proto representation.
func ProtoToDlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(p *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActions converts a JobTriggerInspectJobActions object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActions(p *dlppb.DlpJobTriggerInspectJobActions) *dlp.JobTriggerInspectJobActions {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActions{
		SaveFindings:                      ProtoToDlpJobTriggerInspectJobActionsSaveFindings(p.GetSaveFindings()),
		PubSub:                            ProtoToDlpJobTriggerInspectJobActionsPubSub(p.GetPubSub()),
		PublishSummaryToCscc:              ProtoToDlpJobTriggerInspectJobActionsPublishSummaryToCscc(p.GetPublishSummaryToCscc()),
		PublishFindingsToCloudDataCatalog: ProtoToDlpJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(p.GetPublishFindingsToCloudDataCatalog()),
		JobNotificationEmails:             ProtoToDlpJobTriggerInspectJobActionsJobNotificationEmails(p.GetJobNotificationEmails()),
		PublishToStackdriver:              ProtoToDlpJobTriggerInspectJobActionsPublishToStackdriver(p.GetPublishToStackdriver()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindings converts a JobTriggerInspectJobActionsSaveFindings object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsSaveFindings(p *dlppb.DlpJobTriggerInspectJobActionsSaveFindings) *dlp.JobTriggerInspectJobActionsSaveFindings {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsSaveFindings{
		OutputConfig: ProtoToDlpJobTriggerInspectJobActionsSaveFindingsOutputConfig(p.GetOutputConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfig converts a JobTriggerInspectJobActionsSaveFindingsOutputConfig object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsSaveFindingsOutputConfig(p *dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfig) *dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfig{
		Table:        ProtoToDlpJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(p.GetTable()),
		DlpStorage:   ProtoToDlpJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(p.GetDlpStorage()),
		OutputSchema: ProtoToDlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(p.GetOutputSchema()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigTable converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigTable object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(p *dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigTable) *dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(p *dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) *dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPubSub converts a JobTriggerInspectJobActionsPubSub object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsPubSub(p *dlppb.DlpJobTriggerInspectJobActionsPubSub) *dlp.JobTriggerInspectJobActionsPubSub {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsPubSub{
		Topic: dcl.StringOrNil(p.GetTopic()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishSummaryToCscc converts a JobTriggerInspectJobActionsPublishSummaryToCscc object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsPublishSummaryToCscc(p *dlppb.DlpJobTriggerInspectJobActionsPublishSummaryToCscc) *dlp.JobTriggerInspectJobActionsPublishSummaryToCscc {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsPublishSummaryToCscc{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog converts a JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(p *dlppb.DlpJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) *dlp.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsJobNotificationEmails converts a JobTriggerInspectJobActionsJobNotificationEmails object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsJobNotificationEmails(p *dlppb.DlpJobTriggerInspectJobActionsJobNotificationEmails) *dlp.JobTriggerInspectJobActionsJobNotificationEmails {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsJobNotificationEmails{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishToStackdriver converts a JobTriggerInspectJobActionsPublishToStackdriver object from its proto representation.
func ProtoToDlpJobTriggerInspectJobActionsPublishToStackdriver(p *dlppb.DlpJobTriggerInspectJobActionsPublishToStackdriver) *dlp.JobTriggerInspectJobActionsPublishToStackdriver {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerInspectJobActionsPublishToStackdriver{}
	return obj
}

// ProtoToJobTriggerTriggers converts a JobTriggerTriggers object from its proto representation.
func ProtoToDlpJobTriggerTriggers(p *dlppb.DlpJobTriggerTriggers) *dlp.JobTriggerTriggers {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerTriggers{
		Schedule: ProtoToDlpJobTriggerTriggersSchedule(p.GetSchedule()),
		Manual:   ProtoToDlpJobTriggerTriggersManual(p.GetManual()),
	}
	return obj
}

// ProtoToJobTriggerTriggersSchedule converts a JobTriggerTriggersSchedule object from its proto representation.
func ProtoToDlpJobTriggerTriggersSchedule(p *dlppb.DlpJobTriggerTriggersSchedule) *dlp.JobTriggerTriggersSchedule {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerTriggersSchedule{
		RecurrencePeriodDuration: dcl.StringOrNil(p.GetRecurrencePeriodDuration()),
	}
	return obj
}

// ProtoToJobTriggerTriggersManual converts a JobTriggerTriggersManual object from its proto representation.
func ProtoToDlpJobTriggerTriggersManual(p *dlppb.DlpJobTriggerTriggersManual) *dlp.JobTriggerTriggersManual {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerTriggersManual{}
	return obj
}

// ProtoToJobTriggerErrors converts a JobTriggerErrors object from its proto representation.
func ProtoToDlpJobTriggerErrors(p *dlppb.DlpJobTriggerErrors) *dlp.JobTriggerErrors {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerErrors{
		Details: ProtoToDlpJobTriggerErrorsDetails(p.GetDetails()),
	}
	for _, r := range p.GetTimestamps() {
		obj.Timestamps = append(obj.Timestamps, r)
	}
	return obj
}

// ProtoToJobTriggerErrorsDetails converts a JobTriggerErrorsDetails object from its proto representation.
func ProtoToDlpJobTriggerErrorsDetails(p *dlppb.DlpJobTriggerErrorsDetails) *dlp.JobTriggerErrorsDetails {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerErrorsDetails{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToDlpJobTriggerErrorsDetailsDetails(r))
	}
	return obj
}

// ProtoToJobTriggerErrorsDetailsDetails converts a JobTriggerErrorsDetailsDetails object from its proto representation.
func ProtoToDlpJobTriggerErrorsDetailsDetails(p *dlppb.DlpJobTriggerErrorsDetailsDetails) *dlp.JobTriggerErrorsDetailsDetails {
	if p == nil {
		return nil
	}
	obj := &dlp.JobTriggerErrorsDetailsDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToJobTrigger converts a JobTrigger resource from its proto representation.
func ProtoToJobTrigger(p *dlppb.DlpJobTrigger) *dlp.JobTrigger {
	obj := &dlp.JobTrigger{
		Name:        dcl.StringOrNil(p.GetName()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		InspectJob:  ProtoToDlpJobTriggerInspectJob(p.GetInspectJob()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		LastRunTime: dcl.StringOrNil(p.GetLastRunTime()),
		Status:      ProtoToDlpJobTriggerStatusEnum(p.GetStatus()),
		LocationId:  dcl.StringOrNil(p.GetLocationId()),
		Parent:      dcl.StringOrNil(p.GetParent()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetTriggers() {
		obj.Triggers = append(obj.Triggers, *ProtoToDlpJobTriggerTriggers(r))
	}
	for _, r := range p.GetErrors() {
		obj.Errors = append(obj.Errors, *ProtoToDlpJobTriggerErrors(r))
	}
	return obj
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum enum to its proto representation.
func DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumToProto(e *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum) dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_value["JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(0)
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum enum to its proto representation.
func DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto(e *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum) dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum_value["JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(0)
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum enum to its proto representation.
func DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto(e *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum) dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum_value["JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(0)
}

// JobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigMinLikelihoodEnum enum to its proto representation.
func DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto(e *dlp.JobTriggerInspectJobInspectConfigMinLikelihoodEnum) dlppb.DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum_value["JobTriggerInspectJobInspectConfigMinLikelihoodEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum enum to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto(e *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum) dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum enum to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto(e *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum) dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(0)
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(e *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_value["JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum enum to its proto representation.
func DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto(e *dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum) dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	if e == nil {
		return dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum_value["JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(v)
	}
	return dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(0)
}

// JobTriggerStatusEnumToProto converts a JobTriggerStatusEnum enum to its proto representation.
func DlpJobTriggerStatusEnumToProto(e *dlp.JobTriggerStatusEnum) dlppb.DlpJobTriggerStatusEnum {
	if e == nil {
		return dlppb.DlpJobTriggerStatusEnum(0)
	}
	if v, ok := dlppb.DlpJobTriggerStatusEnum_value["JobTriggerStatusEnum"+string(*e)]; ok {
		return dlppb.DlpJobTriggerStatusEnum(v)
	}
	return dlppb.DlpJobTriggerStatusEnum(0)
}

// JobTriggerInspectJobToProto converts a JobTriggerInspectJob object to its proto representation.
func DlpJobTriggerInspectJobToProto(o *dlp.JobTriggerInspectJob) *dlppb.DlpJobTriggerInspectJob {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJob{}
	p.SetStorageConfig(DlpJobTriggerInspectJobStorageConfigToProto(o.StorageConfig))
	p.SetInspectConfig(DlpJobTriggerInspectJobInspectConfigToProto(o.InspectConfig))
	p.SetInspectTemplateName(dcl.ValueOrEmptyString(o.InspectTemplateName))
	sActions := make([]*dlppb.DlpJobTriggerInspectJobActions, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = DlpJobTriggerInspectJobActionsToProto(&r)
	}
	p.SetActions(sActions)
	return p
}

// JobTriggerInspectJobStorageConfigToProto converts a JobTriggerInspectJobStorageConfig object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigToProto(o *dlp.JobTriggerInspectJobStorageConfig) *dlppb.DlpJobTriggerInspectJobStorageConfig {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfig{}
	p.SetDatastoreOptions(DlpJobTriggerInspectJobStorageConfigDatastoreOptionsToProto(o.DatastoreOptions))
	p.SetCloudStorageOptions(DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsToProto(o.CloudStorageOptions))
	p.SetBigQueryOptions(DlpJobTriggerInspectJobStorageConfigBigQueryOptionsToProto(o.BigQueryOptions))
	p.SetHybridOptions(DlpJobTriggerInspectJobStorageConfigHybridOptionsToProto(o.HybridOptions))
	p.SetTimespanConfig(DlpJobTriggerInspectJobStorageConfigTimespanConfigToProto(o.TimespanConfig))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptions object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigDatastoreOptionsToProto(o *dlp.JobTriggerInspectJobStorageConfigDatastoreOptions) *dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptions {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptions{}
	p.SetPartitionId(DlpJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto(o.PartitionId))
	p.SetKind(DlpJobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto(o.Kind))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto(o *dlp.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) *dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetNamespaceId(dcl.ValueOrEmptyString(o.NamespaceId))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptionsKind object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto(o *dlp.JobTriggerInspectJobStorageConfigDatastoreOptionsKind) *dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptions object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsToProto(o *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptions) *dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptions {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptions{}
	p.SetFileSet(DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto(o.FileSet))
	p.SetBytesLimitPerFile(dcl.ValueOrEmptyInt64(o.BytesLimitPerFile))
	p.SetBytesLimitPerFilePercent(dcl.ValueOrEmptyInt64(o.BytesLimitPerFilePercent))
	p.SetSampleMethod(DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto(o.SampleMethod))
	p.SetFilesLimitPercent(dcl.ValueOrEmptyInt64(o.FilesLimitPercent))
	sFileTypes := make([]dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum, len(o.FileTypes))
	for i, r := range o.FileTypes {
		sFileTypes[i] = dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_value[string(r)])
	}
	p.SetFileTypes(sFileTypes)
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto(o *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) *dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetRegexFileSet(DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto(o.RegexFileSet))
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto(o *dlp.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) *dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
	p.SetBucketName(dcl.ValueOrEmptyString(o.BucketName))
	sIncludeRegex := make([]string, len(o.IncludeRegex))
	for i, r := range o.IncludeRegex {
		sIncludeRegex[i] = r
	}
	p.SetIncludeRegex(sIncludeRegex)
	sExcludeRegex := make([]string, len(o.ExcludeRegex))
	for i, r := range o.ExcludeRegex {
		sExcludeRegex[i] = r
	}
	p.SetExcludeRegex(sExcludeRegex)
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptions object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigBigQueryOptionsToProto(o *dlp.JobTriggerInspectJobStorageConfigBigQueryOptions) *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptions {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptions{}
	p.SetTableReference(DlpJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto(o.TableReference))
	p.SetRowsLimit(dcl.ValueOrEmptyInt64(o.RowsLimit))
	p.SetRowsLimitPercent(dcl.ValueOrEmptyInt64(o.RowsLimitPercent))
	p.SetSampleMethod(DlpJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto(o.SampleMethod))
	sIdentifyingFields := make([]*dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, len(o.IdentifyingFields))
	for i, r := range o.IdentifyingFields {
		sIdentifyingFields[i] = DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto(&r)
	}
	p.SetIdentifyingFields(sIdentifyingFields)
	sExcludedFields := make([]*dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, len(o.ExcludedFields))
	for i, r := range o.ExcludedFields {
		sExcludedFields[i] = DlpJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto(&r)
	}
	p.SetExcludedFields(sExcludedFields)
	sIncludedFields := make([]*dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, len(o.IncludedFields))
	for i, r := range o.IncludedFields {
		sIncludedFields[i] = DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto(&r)
	}
	p.SetIncludedFields(sIncludedFields)
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto(o *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto(o *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto(o *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto(o *dlp.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) *dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigHybridOptionsToProto converts a JobTriggerInspectJobStorageConfigHybridOptions object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigHybridOptionsToProto(o *dlp.JobTriggerInspectJobStorageConfigHybridOptions) *dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptions {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptions{}
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetTableOptions(DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsToProto(o.TableOptions))
	sRequiredFindingLabelKeys := make([]string, len(o.RequiredFindingLabelKeys))
	for i, r := range o.RequiredFindingLabelKeys {
		sRequiredFindingLabelKeys[i] = r
	}
	p.SetRequiredFindingLabelKeys(sRequiredFindingLabelKeys)
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsToProto converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptions object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsToProto(o *dlp.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) *dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	sIdentifyingFields := make([]*dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, len(o.IdentifyingFields))
	for i, r := range o.IdentifyingFields {
		sIdentifyingFields[i] = DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto(&r)
	}
	p.SetIdentifyingFields(sIdentifyingFields)
	return p
}

// JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto(o *dlp.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) *dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigTimespanConfigToProto converts a JobTriggerInspectJobStorageConfigTimespanConfig object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigTimespanConfigToProto(o *dlp.JobTriggerInspectJobStorageConfigTimespanConfig) *dlppb.DlpJobTriggerInspectJobStorageConfigTimespanConfig {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigTimespanConfig{}
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetEndTime(dcl.ValueOrEmptyString(o.EndTime))
	p.SetTimestampField(DlpJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto(o.TimestampField))
	p.SetEnableAutoPopulationOfTimespanConfig(dcl.ValueOrEmptyBool(o.EnableAutoPopulationOfTimespanConfig))
	return p
}

// JobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto converts a JobTriggerInspectJobStorageConfigTimespanConfigTimestampField object to its proto representation.
func DlpJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto(o *dlp.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) *dlppb.DlpJobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobInspectConfigToProto converts a JobTriggerInspectJobInspectConfig object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigToProto(o *dlp.JobTriggerInspectJobInspectConfig) *dlppb.DlpJobTriggerInspectJobInspectConfig {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfig{}
	p.SetMinLikelihood(DlpJobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto(o.MinLikelihood))
	p.SetLimits(DlpJobTriggerInspectJobInspectConfigLimitsToProto(o.Limits))
	p.SetIncludeQuote(dcl.ValueOrEmptyBool(o.IncludeQuote))
	p.SetExcludeInfoTypes(dcl.ValueOrEmptyBool(o.ExcludeInfoTypes))
	sInfoTypes := make([]*dlppb.DlpJobTriggerInspectJobInspectConfigInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpJobTriggerInspectJobInspectConfigInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sCustomInfoTypes := make([]*dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypes, len(o.CustomInfoTypes))
	for i, r := range o.CustomInfoTypes {
		sCustomInfoTypes[i] = DlpJobTriggerInspectJobInspectConfigCustomInfoTypesToProto(&r)
	}
	p.SetCustomInfoTypes(sCustomInfoTypes)
	sRuleSet := make([]*dlppb.DlpJobTriggerInspectJobInspectConfigRuleSet, len(o.RuleSet))
	for i, r := range o.RuleSet {
		sRuleSet[i] = DlpJobTriggerInspectJobInspectConfigRuleSetToProto(&r)
	}
	p.SetRuleSet(sRuleSet)
	return p
}

// JobTriggerInspectJobInspectConfigInfoTypesToProto converts a JobTriggerInspectJobInspectConfigInfoTypes object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigInfoTypesToProto(o *dlp.JobTriggerInspectJobInspectConfigInfoTypes) *dlppb.DlpJobTriggerInspectJobInspectConfigInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobInspectConfigLimitsToProto converts a JobTriggerInspectJobInspectConfigLimits object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigLimitsToProto(o *dlp.JobTriggerInspectJobInspectConfigLimits) *dlppb.DlpJobTriggerInspectJobInspectConfigLimits {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigLimits{}
	p.SetMaxFindingsPerItem(dcl.ValueOrEmptyInt64(o.MaxFindingsPerItem))
	p.SetMaxFindingsPerRequest(dcl.ValueOrEmptyInt64(o.MaxFindingsPerRequest))
	sMaxFindingsPerInfoType := make([]*dlppb.DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, len(o.MaxFindingsPerInfoType))
	for i, r := range o.MaxFindingsPerInfoType {
		sMaxFindingsPerInfoType[i] = DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto(&r)
	}
	p.SetMaxFindingsPerInfoType(sMaxFindingsPerInfoType)
	return p
}

// JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto(o *dlp.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) *dlppb.DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}
	p.SetInfoType(DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o.InfoType))
	p.SetMaxFindings(dcl.ValueOrEmptyInt64(o.MaxFindings))
	return p
}

// JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o *dlp.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *dlppb.DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypes object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypes) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypes{}
	p.SetInfoType(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto(o.InfoType))
	p.SetLikelihood(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto(o.Likelihood))
	p.SetDictionary(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto(o.Regex))
	p.SetSurrogateType(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto(o.SurrogateType))
	p.SetStoredType(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto(o.StoredType))
	p.SetExclusionType(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto(o.ExclusionType))
	sDetectionRules := make([]*dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, len(o.DetectionRules))
	for i, r := range o.DetectionRules {
		sDetectionRules[i] = DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto(&r)
	}
	p.SetDetectionRules(sDetectionRules)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	p.SetWordList(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesRegex object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}
	p.SetHotwordRule(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto(o.HotwordRule))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	p.SetHotwordRegex(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto(o *dlp.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) *dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetToProto converts a JobTriggerInspectJobInspectConfigRuleSet object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSet) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSet {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSet{}
	sInfoTypes := make([]*dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpJobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sRules := make([]*dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = DlpJobTriggerInspectJobInspectConfigRuleSetRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetInfoTypes object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetInfoTypes) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRules object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRules) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRules {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRules{}
	p.SetHotwordRule(DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto(o.HotwordRule))
	p.SetExclusionRule(DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto(o.ExclusionRule))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	p.SetHotwordRegex(DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	p.SetDictionary(DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto(o.Regex))
	p.SetExcludeInfoTypes(DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o.ExcludeInfoTypes))
	p.SetMatchingType(DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(o.MatchingType))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	p.SetWordList(DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	sInfoTypes := make([]*dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object to its proto representation.
func DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(o *dlp.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobActionsToProto converts a JobTriggerInspectJobActions object to its proto representation.
func DlpJobTriggerInspectJobActionsToProto(o *dlp.JobTriggerInspectJobActions) *dlppb.DlpJobTriggerInspectJobActions {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActions{}
	p.SetSaveFindings(DlpJobTriggerInspectJobActionsSaveFindingsToProto(o.SaveFindings))
	p.SetPubSub(DlpJobTriggerInspectJobActionsPubSubToProto(o.PubSub))
	p.SetPublishSummaryToCscc(DlpJobTriggerInspectJobActionsPublishSummaryToCsccToProto(o.PublishSummaryToCscc))
	p.SetPublishFindingsToCloudDataCatalog(DlpJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto(o.PublishFindingsToCloudDataCatalog))
	p.SetJobNotificationEmails(DlpJobTriggerInspectJobActionsJobNotificationEmailsToProto(o.JobNotificationEmails))
	p.SetPublishToStackdriver(DlpJobTriggerInspectJobActionsPublishToStackdriverToProto(o.PublishToStackdriver))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsToProto converts a JobTriggerInspectJobActionsSaveFindings object to its proto representation.
func DlpJobTriggerInspectJobActionsSaveFindingsToProto(o *dlp.JobTriggerInspectJobActionsSaveFindings) *dlppb.DlpJobTriggerInspectJobActionsSaveFindings {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsSaveFindings{}
	p.SetOutputConfig(DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigToProto(o.OutputConfig))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfig object to its proto representation.
func DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigToProto(o *dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfig) *dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	p.SetTable(DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto(o.Table))
	p.SetDlpStorage(DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto(o.DlpStorage))
	p.SetOutputSchema(DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto(o.OutputSchema))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigTable object to its proto representation.
func DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto(o *dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) *dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage object to its proto representation.
func DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto(o *dlp.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) *dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	return p
}

// JobTriggerInspectJobActionsPubSubToProto converts a JobTriggerInspectJobActionsPubSub object to its proto representation.
func DlpJobTriggerInspectJobActionsPubSubToProto(o *dlp.JobTriggerInspectJobActionsPubSub) *dlppb.DlpJobTriggerInspectJobActionsPubSub {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsPubSub{}
	p.SetTopic(dcl.ValueOrEmptyString(o.Topic))
	return p
}

// JobTriggerInspectJobActionsPublishSummaryToCsccToProto converts a JobTriggerInspectJobActionsPublishSummaryToCscc object to its proto representation.
func DlpJobTriggerInspectJobActionsPublishSummaryToCsccToProto(o *dlp.JobTriggerInspectJobActionsPublishSummaryToCscc) *dlppb.DlpJobTriggerInspectJobActionsPublishSummaryToCscc {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsPublishSummaryToCscc{}
	return p
}

// JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto converts a JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog object to its proto representation.
func DlpJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto(o *dlp.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) *dlppb.DlpJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	return p
}

// JobTriggerInspectJobActionsJobNotificationEmailsToProto converts a JobTriggerInspectJobActionsJobNotificationEmails object to its proto representation.
func DlpJobTriggerInspectJobActionsJobNotificationEmailsToProto(o *dlp.JobTriggerInspectJobActionsJobNotificationEmails) *dlppb.DlpJobTriggerInspectJobActionsJobNotificationEmails {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsJobNotificationEmails{}
	return p
}

// JobTriggerInspectJobActionsPublishToStackdriverToProto converts a JobTriggerInspectJobActionsPublishToStackdriver object to its proto representation.
func DlpJobTriggerInspectJobActionsPublishToStackdriverToProto(o *dlp.JobTriggerInspectJobActionsPublishToStackdriver) *dlppb.DlpJobTriggerInspectJobActionsPublishToStackdriver {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerInspectJobActionsPublishToStackdriver{}
	return p
}

// JobTriggerTriggersToProto converts a JobTriggerTriggers object to its proto representation.
func DlpJobTriggerTriggersToProto(o *dlp.JobTriggerTriggers) *dlppb.DlpJobTriggerTriggers {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerTriggers{}
	p.SetSchedule(DlpJobTriggerTriggersScheduleToProto(o.Schedule))
	p.SetManual(DlpJobTriggerTriggersManualToProto(o.Manual))
	return p
}

// JobTriggerTriggersScheduleToProto converts a JobTriggerTriggersSchedule object to its proto representation.
func DlpJobTriggerTriggersScheduleToProto(o *dlp.JobTriggerTriggersSchedule) *dlppb.DlpJobTriggerTriggersSchedule {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerTriggersSchedule{}
	p.SetRecurrencePeriodDuration(dcl.ValueOrEmptyString(o.RecurrencePeriodDuration))
	return p
}

// JobTriggerTriggersManualToProto converts a JobTriggerTriggersManual object to its proto representation.
func DlpJobTriggerTriggersManualToProto(o *dlp.JobTriggerTriggersManual) *dlppb.DlpJobTriggerTriggersManual {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerTriggersManual{}
	return p
}

// JobTriggerErrorsToProto converts a JobTriggerErrors object to its proto representation.
func DlpJobTriggerErrorsToProto(o *dlp.JobTriggerErrors) *dlppb.DlpJobTriggerErrors {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerErrors{}
	p.SetDetails(DlpJobTriggerErrorsDetailsToProto(o.Details))
	sTimestamps := make([]string, len(o.Timestamps))
	for i, r := range o.Timestamps {
		sTimestamps[i] = r
	}
	p.SetTimestamps(sTimestamps)
	return p
}

// JobTriggerErrorsDetailsToProto converts a JobTriggerErrorsDetails object to its proto representation.
func DlpJobTriggerErrorsDetailsToProto(o *dlp.JobTriggerErrorsDetails) *dlppb.DlpJobTriggerErrorsDetails {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerErrorsDetails{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*dlppb.DlpJobTriggerErrorsDetailsDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = DlpJobTriggerErrorsDetailsDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// JobTriggerErrorsDetailsDetailsToProto converts a JobTriggerErrorsDetailsDetails object to its proto representation.
func DlpJobTriggerErrorsDetailsDetailsToProto(o *dlp.JobTriggerErrorsDetailsDetails) *dlppb.DlpJobTriggerErrorsDetailsDetails {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpJobTriggerErrorsDetailsDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// JobTriggerToProto converts a JobTrigger resource to its proto representation.
func JobTriggerToProto(resource *dlp.JobTrigger) *dlppb.DlpJobTrigger {
	p := &dlppb.DlpJobTrigger{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInspectJob(DlpJobTriggerInspectJobToProto(resource.InspectJob))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetLastRunTime(dcl.ValueOrEmptyString(resource.LastRunTime))
	p.SetStatus(DlpJobTriggerStatusEnumToProto(resource.Status))
	p.SetLocationId(dcl.ValueOrEmptyString(resource.LocationId))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sTriggers := make([]*dlppb.DlpJobTriggerTriggers, len(resource.Triggers))
	for i, r := range resource.Triggers {
		sTriggers[i] = DlpJobTriggerTriggersToProto(&r)
	}
	p.SetTriggers(sTriggers)
	sErrors := make([]*dlppb.DlpJobTriggerErrors, len(resource.Errors))
	for i, r := range resource.Errors {
		sErrors[i] = DlpJobTriggerErrorsToProto(&r)
	}
	p.SetErrors(sErrors)

	return p
}

// applyJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Apply() method.
func (s *JobTriggerServer) applyJobTrigger(ctx context.Context, c *dlp.Client, request *dlppb.ApplyDlpJobTriggerRequest) (*dlppb.DlpJobTrigger, error) {
	p := ProtoToJobTrigger(request.GetResource())
	res, err := c.ApplyJobTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobTriggerToProto(res)
	return r, nil
}

// applyDlpJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Apply() method.
func (s *JobTriggerServer) ApplyDlpJobTrigger(ctx context.Context, request *dlppb.ApplyDlpJobTriggerRequest) (*dlppb.DlpJobTrigger, error) {
	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyJobTrigger(ctx, cl, request)
}

// DeleteJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Delete() method.
func (s *JobTriggerServer) DeleteDlpJobTrigger(ctx context.Context, request *dlppb.DeleteDlpJobTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJobTrigger(ctx, ProtoToJobTrigger(request.GetResource()))

}

// ListDlpJobTrigger handles the gRPC request by passing it to the underlying JobTriggerList() method.
func (s *JobTriggerServer) ListDlpJobTrigger(ctx context.Context, request *dlppb.ListDlpJobTriggerRequest) (*dlppb.ListDlpJobTriggerResponse, error) {
	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJobTrigger(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*dlppb.DlpJobTrigger
	for _, r := range resources.Items {
		rp := JobTriggerToProto(r)
		protos = append(protos, rp)
	}
	p := &dlppb.ListDlpJobTriggerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigJobTrigger(ctx context.Context, service_account_file string) (*dlp.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dlp.NewClient(conf), nil
}
