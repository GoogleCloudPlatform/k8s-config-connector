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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dlp/alpha/dlp_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/alpha"
)

// JobTriggerServer implements the gRPC interface for JobTrigger.
type JobTriggerServer struct{}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(e alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum) *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(n[len("DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(e alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum) *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(n[len("DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum converts a JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(e alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum) *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(n[len("DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigMinLikelihoodEnum converts a JobTriggerInspectJobInspectConfigMinLikelihoodEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(e alphapb.DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum) *alpha.JobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobInspectConfigMinLikelihoodEnum(n[len("DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(e alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(n[len("DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(e alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(n[len("DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(e alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(n[len("DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(e alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum) *alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(n[len("DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerStatusEnum converts a JobTriggerStatusEnum enum from its proto representation.
func ProtoToDlpAlphaJobTriggerStatusEnum(e alphapb.DlpAlphaJobTriggerStatusEnum) *alpha.JobTriggerStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaJobTriggerStatusEnum_name[int32(e)]; ok {
		e := alpha.JobTriggerStatusEnum(n[len("DlpAlphaJobTriggerStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJob converts a JobTriggerInspectJob object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJob(p *alphapb.DlpAlphaJobTriggerInspectJob) *alpha.JobTriggerInspectJob {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJob{
		StorageConfig:       ProtoToDlpAlphaJobTriggerInspectJobStorageConfig(p.GetStorageConfig()),
		InspectConfig:       ProtoToDlpAlphaJobTriggerInspectJobInspectConfig(p.GetInspectConfig()),
		InspectTemplateName: dcl.StringOrNil(p.GetInspectTemplateName()),
	}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, *ProtoToDlpAlphaJobTriggerInspectJobActions(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfig converts a JobTriggerInspectJobStorageConfig object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfig(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfig) *alpha.JobTriggerInspectJobStorageConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfig{
		DatastoreOptions:    ProtoToDlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptions(p.GetDatastoreOptions()),
		CloudStorageOptions: ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptions(p.GetCloudStorageOptions()),
		BigQueryOptions:     ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptions(p.GetBigQueryOptions()),
		HybridOptions:       ProtoToDlpAlphaJobTriggerInspectJobStorageConfigHybridOptions(p.GetHybridOptions()),
		TimespanConfig:      ProtoToDlpAlphaJobTriggerInspectJobStorageConfigTimespanConfig(p.GetTimespanConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptions converts a JobTriggerInspectJobStorageConfigDatastoreOptions object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptions(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptions) *alpha.JobTriggerInspectJobStorageConfigDatastoreOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigDatastoreOptions{
		PartitionId: ProtoToDlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(p.GetPartitionId()),
		Kind:        ProtoToDlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsKind(p.GetKind()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId converts a JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) *alpha.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{
		ProjectId:   dcl.StringOrNil(p.GetProjectId()),
		NamespaceId: dcl.StringOrNil(p.GetNamespaceId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptionsKind converts a JobTriggerInspectJobStorageConfigDatastoreOptionsKind object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsKind(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsKind) *alpha.JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigDatastoreOptionsKind{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptions converts a JobTriggerInspectJobStorageConfigCloudStorageOptions object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptions(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptions) *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigCloudStorageOptions{
		FileSet:                  ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(p.GetFileSet()),
		BytesLimitPerFile:        dcl.Int64OrNil(p.GetBytesLimitPerFile()),
		BytesLimitPerFilePercent: dcl.Int64OrNil(p.GetBytesLimitPerFilePercent()),
		SampleMethod:             ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(p.GetSampleMethod()),
		FilesLimitPercent:        dcl.Int64OrNil(p.GetFilesLimitPercent()),
	}
	for _, r := range p.GetFileTypes() {
		obj.FileTypes = append(obj.FileTypes, *ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{
		Url:          dcl.StringOrNil(p.GetUrl()),
		RegexFileSet: ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(p.GetRegexFileSet()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{
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
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptions(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptions) *alpha.JobTriggerInspectJobStorageConfigBigQueryOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigBigQueryOptions{
		TableReference:   ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(p.GetTableReference()),
		RowsLimit:        dcl.Int64OrNil(p.GetRowsLimit()),
		RowsLimitPercent: dcl.Int64OrNil(p.GetRowsLimitPercent()),
		SampleMethod:     ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(p.GetSampleMethod()),
	}
	for _, r := range p.GetIdentifyingFields() {
		obj.IdentifyingFields = append(obj.IdentifyingFields, *ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(r))
	}
	for _, r := range p.GetExcludedFields() {
		obj.ExcludedFields = append(obj.ExcludedFields, *ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(r))
	}
	for _, r := range p.GetIncludedFields() {
		obj.IncludedFields = append(obj.IncludedFields, *ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference converts a JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptions converts a JobTriggerInspectJobStorageConfigHybridOptions object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigHybridOptions(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptions) *alpha.JobTriggerInspectJobStorageConfigHybridOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigHybridOptions{
		Description:  dcl.StringOrNil(p.GetDescription()),
		TableOptions: ProtoToDlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(p.GetTableOptions()),
	}
	for _, r := range p.GetRequiredFindingLabelKeys() {
		obj.RequiredFindingLabelKeys = append(obj.RequiredFindingLabelKeys, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptionsTableOptions converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptions object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions) *alpha.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	for _, r := range p.GetIdentifyingFields() {
		obj.IdentifyingFields = append(obj.IdentifyingFields, *ProtoToDlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) *alpha.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigTimespanConfig converts a JobTriggerInspectJobStorageConfigTimespanConfig object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigTimespanConfig(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfig) *alpha.JobTriggerInspectJobStorageConfigTimespanConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigTimespanConfig{
		StartTime:                            dcl.StringOrNil(p.GetStartTime()),
		EndTime:                              dcl.StringOrNil(p.GetEndTime()),
		TimestampField:                       ProtoToDlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(p.GetTimestampField()),
		EnableAutoPopulationOfTimespanConfig: dcl.Bool(p.GetEnableAutoPopulationOfTimespanConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigTimespanConfigTimestampField converts a JobTriggerInspectJobStorageConfigTimespanConfigTimestampField object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(p *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField) *alpha.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfig converts a JobTriggerInspectJobInspectConfig object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfig(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfig) *alpha.JobTriggerInspectJobInspectConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfig{
		MinLikelihood:    ProtoToDlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(p.GetMinLikelihood()),
		Limits:           ProtoToDlpAlphaJobTriggerInspectJobInspectConfigLimits(p.GetLimits()),
		IncludeQuote:     dcl.Bool(p.GetIncludeQuote()),
		ExcludeInfoTypes: dcl.Bool(p.GetExcludeInfoTypes()),
	}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpAlphaJobTriggerInspectJobInspectConfigInfoTypes(r))
	}
	for _, r := range p.GetCustomInfoTypes() {
		obj.CustomInfoTypes = append(obj.CustomInfoTypes, *ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypes(r))
	}
	for _, r := range p.GetRuleSet() {
		obj.RuleSet = append(obj.RuleSet, *ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSet(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigInfoTypes converts a JobTriggerInspectJobInspectConfigInfoTypes object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigInfoTypes(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigInfoTypes) *alpha.JobTriggerInspectJobInspectConfigInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimits converts a JobTriggerInspectJobInspectConfigLimits object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigLimits(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimits) *alpha.JobTriggerInspectJobInspectConfigLimits {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigLimits{
		MaxFindingsPerItem:    dcl.Int64OrNil(p.GetMaxFindingsPerItem()),
		MaxFindingsPerRequest: dcl.Int64OrNil(p.GetMaxFindingsPerRequest()),
	}
	for _, r := range p.GetMaxFindingsPerInfoType() {
		obj.MaxFindingsPerInfoType = append(obj.MaxFindingsPerInfoType, *ProtoToDlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) *alpha.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{
		InfoType:    ProtoToDlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p.GetInfoType()),
		MaxFindings: dcl.Int64OrNil(p.GetMaxFindings()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *alpha.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypes converts a JobTriggerInspectJobInspectConfigCustomInfoTypes object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypes(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypes) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypes{
		InfoType:      ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(p.GetInfoType()),
		Likelihood:    ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(p.GetLikelihood()),
		Dictionary:    ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(p.GetDictionary()),
		Regex:         ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(p.GetRegex()),
		SurrogateType: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(p.GetSurrogateType()),
		StoredType:    ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(p.GetStoredType()),
		ExclusionType: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(p.GetExclusionType()),
	}
	for _, r := range p.GetDetectionRules() {
		obj.DetectionRules = append(obj.DetectionRules, *ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{
		WordList:         ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesRegex converts a JobTriggerInspectJobInspectConfigCustomInfoTypesRegex object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{
		HotwordRule: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(p.GetHotwordRule()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{
		HotwordRegex:         ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSet converts a JobTriggerInspectJobInspectConfigRuleSet object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSet(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSet) *alpha.JobTriggerInspectJobInspectConfigRuleSet {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSet{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetInfoTypes(r))
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRules(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetInfoTypes object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetInfoTypes(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetInfoTypes) *alpha.JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetInfoTypes{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRules converts a JobTriggerInspectJobInspectConfigRuleSetRules object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRules(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRules) *alpha.JobTriggerInspectJobInspectConfigRuleSetRules {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRules{
		HotwordRule:   ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(p.GetHotwordRule()),
		ExclusionRule: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(p.GetExclusionRule()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{
		HotwordRegex:         ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{
		Dictionary:       ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(p.GetDictionary()),
		Regex:            ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(p.GetRegex()),
		ExcludeInfoTypes: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p.GetExcludeInfoTypes()),
		MatchingType:     ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(p.GetMatchingType()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{
		WordList:         ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(p *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActions converts a JobTriggerInspectJobActions object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActions(p *alphapb.DlpAlphaJobTriggerInspectJobActions) *alpha.JobTriggerInspectJobActions {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActions{
		SaveFindings:                      ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindings(p.GetSaveFindings()),
		PubSub:                            ProtoToDlpAlphaJobTriggerInspectJobActionsPubSub(p.GetPubSub()),
		PublishSummaryToCscc:              ProtoToDlpAlphaJobTriggerInspectJobActionsPublishSummaryToCscc(p.GetPublishSummaryToCscc()),
		PublishFindingsToCloudDataCatalog: ProtoToDlpAlphaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(p.GetPublishFindingsToCloudDataCatalog()),
		JobNotificationEmails:             ProtoToDlpAlphaJobTriggerInspectJobActionsJobNotificationEmails(p.GetJobNotificationEmails()),
		PublishToStackdriver:              ProtoToDlpAlphaJobTriggerInspectJobActionsPublishToStackdriver(p.GetPublishToStackdriver()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindings converts a JobTriggerInspectJobActionsSaveFindings object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindings(p *alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindings) *alpha.JobTriggerInspectJobActionsSaveFindings {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsSaveFindings{
		OutputConfig: ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfig(p.GetOutputConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfig converts a JobTriggerInspectJobActionsSaveFindingsOutputConfig object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfig(p *alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfig) *alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfig{
		Table:        ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(p.GetTable()),
		DlpStorage:   ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(p.GetDlpStorage()),
		OutputSchema: ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(p.GetOutputSchema()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigTable converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigTable object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(p *alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable) *alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(p *alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) *alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPubSub converts a JobTriggerInspectJobActionsPubSub object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsPubSub(p *alphapb.DlpAlphaJobTriggerInspectJobActionsPubSub) *alpha.JobTriggerInspectJobActionsPubSub {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsPubSub{
		Topic: dcl.StringOrNil(p.GetTopic()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishSummaryToCscc converts a JobTriggerInspectJobActionsPublishSummaryToCscc object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsPublishSummaryToCscc(p *alphapb.DlpAlphaJobTriggerInspectJobActionsPublishSummaryToCscc) *alpha.JobTriggerInspectJobActionsPublishSummaryToCscc {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsPublishSummaryToCscc{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog converts a JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(p *alphapb.DlpAlphaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) *alpha.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsJobNotificationEmails converts a JobTriggerInspectJobActionsJobNotificationEmails object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsJobNotificationEmails(p *alphapb.DlpAlphaJobTriggerInspectJobActionsJobNotificationEmails) *alpha.JobTriggerInspectJobActionsJobNotificationEmails {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsJobNotificationEmails{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishToStackdriver converts a JobTriggerInspectJobActionsPublishToStackdriver object from its proto representation.
func ProtoToDlpAlphaJobTriggerInspectJobActionsPublishToStackdriver(p *alphapb.DlpAlphaJobTriggerInspectJobActionsPublishToStackdriver) *alpha.JobTriggerInspectJobActionsPublishToStackdriver {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerInspectJobActionsPublishToStackdriver{}
	return obj
}

// ProtoToJobTriggerTriggers converts a JobTriggerTriggers object from its proto representation.
func ProtoToDlpAlphaJobTriggerTriggers(p *alphapb.DlpAlphaJobTriggerTriggers) *alpha.JobTriggerTriggers {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerTriggers{
		Schedule: ProtoToDlpAlphaJobTriggerTriggersSchedule(p.GetSchedule()),
		Manual:   ProtoToDlpAlphaJobTriggerTriggersManual(p.GetManual()),
	}
	return obj
}

// ProtoToJobTriggerTriggersSchedule converts a JobTriggerTriggersSchedule object from its proto representation.
func ProtoToDlpAlphaJobTriggerTriggersSchedule(p *alphapb.DlpAlphaJobTriggerTriggersSchedule) *alpha.JobTriggerTriggersSchedule {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerTriggersSchedule{
		RecurrencePeriodDuration: dcl.StringOrNil(p.GetRecurrencePeriodDuration()),
	}
	return obj
}

// ProtoToJobTriggerTriggersManual converts a JobTriggerTriggersManual object from its proto representation.
func ProtoToDlpAlphaJobTriggerTriggersManual(p *alphapb.DlpAlphaJobTriggerTriggersManual) *alpha.JobTriggerTriggersManual {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerTriggersManual{}
	return obj
}

// ProtoToJobTriggerErrors converts a JobTriggerErrors object from its proto representation.
func ProtoToDlpAlphaJobTriggerErrors(p *alphapb.DlpAlphaJobTriggerErrors) *alpha.JobTriggerErrors {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerErrors{
		Details: ProtoToDlpAlphaJobTriggerErrorsDetails(p.GetDetails()),
	}
	for _, r := range p.GetTimestamps() {
		obj.Timestamps = append(obj.Timestamps, r)
	}
	return obj
}

// ProtoToJobTriggerErrorsDetails converts a JobTriggerErrorsDetails object from its proto representation.
func ProtoToDlpAlphaJobTriggerErrorsDetails(p *alphapb.DlpAlphaJobTriggerErrorsDetails) *alpha.JobTriggerErrorsDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerErrorsDetails{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToDlpAlphaJobTriggerErrorsDetailsDetails(r))
	}
	return obj
}

// ProtoToJobTriggerErrorsDetailsDetails converts a JobTriggerErrorsDetailsDetails object from its proto representation.
func ProtoToDlpAlphaJobTriggerErrorsDetailsDetails(p *alphapb.DlpAlphaJobTriggerErrorsDetailsDetails) *alpha.JobTriggerErrorsDetailsDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTriggerErrorsDetailsDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToJobTrigger converts a JobTrigger resource from its proto representation.
func ProtoToJobTrigger(p *alphapb.DlpAlphaJobTrigger) *alpha.JobTrigger {
	obj := &alpha.JobTrigger{
		Name:        dcl.StringOrNil(p.GetName()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		InspectJob:  ProtoToDlpAlphaJobTriggerInspectJob(p.GetInspectJob()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		LastRunTime: dcl.StringOrNil(p.GetLastRunTime()),
		Status:      ProtoToDlpAlphaJobTriggerStatusEnum(p.GetStatus()),
		LocationId:  dcl.StringOrNil(p.GetLocationId()),
		Parent:      dcl.StringOrNil(p.GetParent()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetTriggers() {
		obj.Triggers = append(obj.Triggers, *ProtoToDlpAlphaJobTriggerTriggers(r))
	}
	for _, r := range p.GetErrors() {
		obj.Errors = append(obj.Errors, *ProtoToDlpAlphaJobTriggerErrors(r))
	}
	return obj
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumToProto(e *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum) alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_value["JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(0)
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto(e *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum) alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum_value["JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(0)
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto(e *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum) alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum_value["JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(0)
}

// JobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigMinLikelihoodEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto(e *alpha.JobTriggerInspectJobInspectConfigMinLikelihoodEnum) alphapb.DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum_value["JobTriggerInspectJobInspectConfigMinLikelihoodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto(e *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum) alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto(e *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum) alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(0)
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(e *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_value["JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum enum to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto(e *alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum) alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum_value["JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(0)
}

// JobTriggerStatusEnumToProto converts a JobTriggerStatusEnum enum to its proto representation.
func DlpAlphaJobTriggerStatusEnumToProto(e *alpha.JobTriggerStatusEnum) alphapb.DlpAlphaJobTriggerStatusEnum {
	if e == nil {
		return alphapb.DlpAlphaJobTriggerStatusEnum(0)
	}
	if v, ok := alphapb.DlpAlphaJobTriggerStatusEnum_value["JobTriggerStatusEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaJobTriggerStatusEnum(v)
	}
	return alphapb.DlpAlphaJobTriggerStatusEnum(0)
}

// JobTriggerInspectJobToProto converts a JobTriggerInspectJob object to its proto representation.
func DlpAlphaJobTriggerInspectJobToProto(o *alpha.JobTriggerInspectJob) *alphapb.DlpAlphaJobTriggerInspectJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJob{}
	p.SetStorageConfig(DlpAlphaJobTriggerInspectJobStorageConfigToProto(o.StorageConfig))
	p.SetInspectConfig(DlpAlphaJobTriggerInspectJobInspectConfigToProto(o.InspectConfig))
	p.SetInspectTemplateName(dcl.ValueOrEmptyString(o.InspectTemplateName))
	sActions := make([]*alphapb.DlpAlphaJobTriggerInspectJobActions, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = DlpAlphaJobTriggerInspectJobActionsToProto(&r)
	}
	p.SetActions(sActions)
	return p
}

// JobTriggerInspectJobStorageConfigToProto converts a JobTriggerInspectJobStorageConfig object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigToProto(o *alpha.JobTriggerInspectJobStorageConfig) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfig{}
	p.SetDatastoreOptions(DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsToProto(o.DatastoreOptions))
	p.SetCloudStorageOptions(DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsToProto(o.CloudStorageOptions))
	p.SetBigQueryOptions(DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsToProto(o.BigQueryOptions))
	p.SetHybridOptions(DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsToProto(o.HybridOptions))
	p.SetTimespanConfig(DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigToProto(o.TimespanConfig))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptions object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsToProto(o *alpha.JobTriggerInspectJobStorageConfigDatastoreOptions) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptions{}
	p.SetPartitionId(DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto(o.PartitionId))
	p.SetKind(DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto(o.Kind))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto(o *alpha.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetNamespaceId(dcl.ValueOrEmptyString(o.NamespaceId))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptionsKind object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto(o *alpha.JobTriggerInspectJobStorageConfigDatastoreOptionsKind) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptions object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsToProto(o *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptions) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptions{}
	p.SetFileSet(DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto(o.FileSet))
	p.SetBytesLimitPerFile(dcl.ValueOrEmptyInt64(o.BytesLimitPerFile))
	p.SetBytesLimitPerFilePercent(dcl.ValueOrEmptyInt64(o.BytesLimitPerFilePercent))
	p.SetSampleMethod(DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto(o.SampleMethod))
	p.SetFilesLimitPercent(dcl.ValueOrEmptyInt64(o.FilesLimitPercent))
	sFileTypes := make([]alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum, len(o.FileTypes))
	for i, r := range o.FileTypes {
		sFileTypes[i] = alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_value[string(r)])
	}
	p.SetFileTypes(sFileTypes)
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto(o *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetRegexFileSet(DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto(o.RegexFileSet))
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto(o *alpha.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
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
func DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsToProto(o *alpha.JobTriggerInspectJobStorageConfigBigQueryOptions) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptions{}
	p.SetTableReference(DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto(o.TableReference))
	p.SetRowsLimit(dcl.ValueOrEmptyInt64(o.RowsLimit))
	p.SetRowsLimitPercent(dcl.ValueOrEmptyInt64(o.RowsLimitPercent))
	p.SetSampleMethod(DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto(o.SampleMethod))
	sIdentifyingFields := make([]*alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, len(o.IdentifyingFields))
	for i, r := range o.IdentifyingFields {
		sIdentifyingFields[i] = DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto(&r)
	}
	p.SetIdentifyingFields(sIdentifyingFields)
	sExcludedFields := make([]*alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, len(o.ExcludedFields))
	for i, r := range o.ExcludedFields {
		sExcludedFields[i] = DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto(&r)
	}
	p.SetExcludedFields(sExcludedFields)
	sIncludedFields := make([]*alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, len(o.IncludedFields))
	for i, r := range o.IncludedFields {
		sIncludedFields[i] = DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto(&r)
	}
	p.SetIncludedFields(sIncludedFields)
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto(o *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto(o *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto(o *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto(o *alpha.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigHybridOptionsToProto converts a JobTriggerInspectJobStorageConfigHybridOptions object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsToProto(o *alpha.JobTriggerInspectJobStorageConfigHybridOptions) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptions{}
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetTableOptions(DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsToProto(o.TableOptions))
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
func DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsToProto(o *alpha.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	sIdentifyingFields := make([]*alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, len(o.IdentifyingFields))
	for i, r := range o.IdentifyingFields {
		sIdentifyingFields[i] = DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto(&r)
	}
	p.SetIdentifyingFields(sIdentifyingFields)
	return p
}

// JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto(o *alpha.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigTimespanConfigToProto converts a JobTriggerInspectJobStorageConfigTimespanConfig object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigToProto(o *alpha.JobTriggerInspectJobStorageConfigTimespanConfig) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfig{}
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetEndTime(dcl.ValueOrEmptyString(o.EndTime))
	p.SetTimestampField(DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto(o.TimestampField))
	p.SetEnableAutoPopulationOfTimespanConfig(dcl.ValueOrEmptyBool(o.EnableAutoPopulationOfTimespanConfig))
	return p
}

// JobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto converts a JobTriggerInspectJobStorageConfigTimespanConfigTimestampField object to its proto representation.
func DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto(o *alpha.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) *alphapb.DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobInspectConfigToProto converts a JobTriggerInspectJobInspectConfig object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigToProto(o *alpha.JobTriggerInspectJobInspectConfig) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfig{}
	p.SetMinLikelihood(DlpAlphaJobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto(o.MinLikelihood))
	p.SetLimits(DlpAlphaJobTriggerInspectJobInspectConfigLimitsToProto(o.Limits))
	p.SetIncludeQuote(dcl.ValueOrEmptyBool(o.IncludeQuote))
	p.SetExcludeInfoTypes(dcl.ValueOrEmptyBool(o.ExcludeInfoTypes))
	sInfoTypes := make([]*alphapb.DlpAlphaJobTriggerInspectJobInspectConfigInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpAlphaJobTriggerInspectJobInspectConfigInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sCustomInfoTypes := make([]*alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypes, len(o.CustomInfoTypes))
	for i, r := range o.CustomInfoTypes {
		sCustomInfoTypes[i] = DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesToProto(&r)
	}
	p.SetCustomInfoTypes(sCustomInfoTypes)
	sRuleSet := make([]*alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSet, len(o.RuleSet))
	for i, r := range o.RuleSet {
		sRuleSet[i] = DlpAlphaJobTriggerInspectJobInspectConfigRuleSetToProto(&r)
	}
	p.SetRuleSet(sRuleSet)
	return p
}

// JobTriggerInspectJobInspectConfigInfoTypesToProto converts a JobTriggerInspectJobInspectConfigInfoTypes object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigInfoTypesToProto(o *alpha.JobTriggerInspectJobInspectConfigInfoTypes) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobInspectConfigLimitsToProto converts a JobTriggerInspectJobInspectConfigLimits object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigLimitsToProto(o *alpha.JobTriggerInspectJobInspectConfigLimits) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimits {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimits{}
	p.SetMaxFindingsPerItem(dcl.ValueOrEmptyInt64(o.MaxFindingsPerItem))
	p.SetMaxFindingsPerRequest(dcl.ValueOrEmptyInt64(o.MaxFindingsPerRequest))
	sMaxFindingsPerInfoType := make([]*alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, len(o.MaxFindingsPerInfoType))
	for i, r := range o.MaxFindingsPerInfoType {
		sMaxFindingsPerInfoType[i] = DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto(&r)
	}
	p.SetMaxFindingsPerInfoType(sMaxFindingsPerInfoType)
	return p
}

// JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto(o *alpha.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}
	p.SetInfoType(DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o.InfoType))
	p.SetMaxFindings(dcl.ValueOrEmptyInt64(o.MaxFindings))
	return p
}

// JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o *alpha.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypes object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypes) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypes{}
	p.SetInfoType(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto(o.InfoType))
	p.SetLikelihood(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto(o.Likelihood))
	p.SetDictionary(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto(o.Regex))
	p.SetSurrogateType(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto(o.SurrogateType))
	p.SetStoredType(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto(o.StoredType))
	p.SetExclusionType(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto(o.ExclusionType))
	sDetectionRules := make([]*alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, len(o.DetectionRules))
	for i, r := range o.DetectionRules {
		sDetectionRules[i] = DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto(&r)
	}
	p.SetDetectionRules(sDetectionRules)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	p.SetWordList(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesRegex object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}
	p.SetHotwordRule(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto(o.HotwordRule))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	p.SetHotwordRegex(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto(o *alpha.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpAlphaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetToProto converts a JobTriggerInspectJobInspectConfigRuleSet object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSet) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSet {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSet{}
	sInfoTypes := make([]*alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpAlphaJobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sRules := make([]*alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetInfoTypes object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetInfoTypes) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRules object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRules) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRules {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRules{}
	p.SetHotwordRule(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto(o.HotwordRule))
	p.SetExclusionRule(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto(o.ExclusionRule))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	p.SetHotwordRegex(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	p.SetDictionary(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto(o.Regex))
	p.SetExcludeInfoTypes(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o.ExcludeInfoTypes))
	p.SetMatchingType(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(o.MatchingType))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	p.SetWordList(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	sInfoTypes := make([]*alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object to its proto representation.
func DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(o *alpha.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobActionsToProto converts a JobTriggerInspectJobActions object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsToProto(o *alpha.JobTriggerInspectJobActions) *alphapb.DlpAlphaJobTriggerInspectJobActions {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActions{}
	p.SetSaveFindings(DlpAlphaJobTriggerInspectJobActionsSaveFindingsToProto(o.SaveFindings))
	p.SetPubSub(DlpAlphaJobTriggerInspectJobActionsPubSubToProto(o.PubSub))
	p.SetPublishSummaryToCscc(DlpAlphaJobTriggerInspectJobActionsPublishSummaryToCsccToProto(o.PublishSummaryToCscc))
	p.SetPublishFindingsToCloudDataCatalog(DlpAlphaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto(o.PublishFindingsToCloudDataCatalog))
	p.SetJobNotificationEmails(DlpAlphaJobTriggerInspectJobActionsJobNotificationEmailsToProto(o.JobNotificationEmails))
	p.SetPublishToStackdriver(DlpAlphaJobTriggerInspectJobActionsPublishToStackdriverToProto(o.PublishToStackdriver))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsToProto converts a JobTriggerInspectJobActionsSaveFindings object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsSaveFindingsToProto(o *alpha.JobTriggerInspectJobActionsSaveFindings) *alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindings {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindings{}
	p.SetOutputConfig(DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigToProto(o.OutputConfig))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfig object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigToProto(o *alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfig) *alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	p.SetTable(DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto(o.Table))
	p.SetDlpStorage(DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto(o.DlpStorage))
	p.SetOutputSchema(DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto(o.OutputSchema))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigTable object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto(o *alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) *alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto(o *alpha.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) *alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	return p
}

// JobTriggerInspectJobActionsPubSubToProto converts a JobTriggerInspectJobActionsPubSub object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsPubSubToProto(o *alpha.JobTriggerInspectJobActionsPubSub) *alphapb.DlpAlphaJobTriggerInspectJobActionsPubSub {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsPubSub{}
	p.SetTopic(dcl.ValueOrEmptyString(o.Topic))
	return p
}

// JobTriggerInspectJobActionsPublishSummaryToCsccToProto converts a JobTriggerInspectJobActionsPublishSummaryToCscc object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsPublishSummaryToCsccToProto(o *alpha.JobTriggerInspectJobActionsPublishSummaryToCscc) *alphapb.DlpAlphaJobTriggerInspectJobActionsPublishSummaryToCscc {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsPublishSummaryToCscc{}
	return p
}

// JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto converts a JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto(o *alpha.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) *alphapb.DlpAlphaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	return p
}

// JobTriggerInspectJobActionsJobNotificationEmailsToProto converts a JobTriggerInspectJobActionsJobNotificationEmails object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsJobNotificationEmailsToProto(o *alpha.JobTriggerInspectJobActionsJobNotificationEmails) *alphapb.DlpAlphaJobTriggerInspectJobActionsJobNotificationEmails {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsJobNotificationEmails{}
	return p
}

// JobTriggerInspectJobActionsPublishToStackdriverToProto converts a JobTriggerInspectJobActionsPublishToStackdriver object to its proto representation.
func DlpAlphaJobTriggerInspectJobActionsPublishToStackdriverToProto(o *alpha.JobTriggerInspectJobActionsPublishToStackdriver) *alphapb.DlpAlphaJobTriggerInspectJobActionsPublishToStackdriver {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerInspectJobActionsPublishToStackdriver{}
	return p
}

// JobTriggerTriggersToProto converts a JobTriggerTriggers object to its proto representation.
func DlpAlphaJobTriggerTriggersToProto(o *alpha.JobTriggerTriggers) *alphapb.DlpAlphaJobTriggerTriggers {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerTriggers{}
	p.SetSchedule(DlpAlphaJobTriggerTriggersScheduleToProto(o.Schedule))
	p.SetManual(DlpAlphaJobTriggerTriggersManualToProto(o.Manual))
	return p
}

// JobTriggerTriggersScheduleToProto converts a JobTriggerTriggersSchedule object to its proto representation.
func DlpAlphaJobTriggerTriggersScheduleToProto(o *alpha.JobTriggerTriggersSchedule) *alphapb.DlpAlphaJobTriggerTriggersSchedule {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerTriggersSchedule{}
	p.SetRecurrencePeriodDuration(dcl.ValueOrEmptyString(o.RecurrencePeriodDuration))
	return p
}

// JobTriggerTriggersManualToProto converts a JobTriggerTriggersManual object to its proto representation.
func DlpAlphaJobTriggerTriggersManualToProto(o *alpha.JobTriggerTriggersManual) *alphapb.DlpAlphaJobTriggerTriggersManual {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerTriggersManual{}
	return p
}

// JobTriggerErrorsToProto converts a JobTriggerErrors object to its proto representation.
func DlpAlphaJobTriggerErrorsToProto(o *alpha.JobTriggerErrors) *alphapb.DlpAlphaJobTriggerErrors {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerErrors{}
	p.SetDetails(DlpAlphaJobTriggerErrorsDetailsToProto(o.Details))
	sTimestamps := make([]string, len(o.Timestamps))
	for i, r := range o.Timestamps {
		sTimestamps[i] = r
	}
	p.SetTimestamps(sTimestamps)
	return p
}

// JobTriggerErrorsDetailsToProto converts a JobTriggerErrorsDetails object to its proto representation.
func DlpAlphaJobTriggerErrorsDetailsToProto(o *alpha.JobTriggerErrorsDetails) *alphapb.DlpAlphaJobTriggerErrorsDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerErrorsDetails{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*alphapb.DlpAlphaJobTriggerErrorsDetailsDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = DlpAlphaJobTriggerErrorsDetailsDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// JobTriggerErrorsDetailsDetailsToProto converts a JobTriggerErrorsDetailsDetails object to its proto representation.
func DlpAlphaJobTriggerErrorsDetailsDetailsToProto(o *alpha.JobTriggerErrorsDetailsDetails) *alphapb.DlpAlphaJobTriggerErrorsDetailsDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaJobTriggerErrorsDetailsDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// JobTriggerToProto converts a JobTrigger resource to its proto representation.
func JobTriggerToProto(resource *alpha.JobTrigger) *alphapb.DlpAlphaJobTrigger {
	p := &alphapb.DlpAlphaJobTrigger{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInspectJob(DlpAlphaJobTriggerInspectJobToProto(resource.InspectJob))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetLastRunTime(dcl.ValueOrEmptyString(resource.LastRunTime))
	p.SetStatus(DlpAlphaJobTriggerStatusEnumToProto(resource.Status))
	p.SetLocationId(dcl.ValueOrEmptyString(resource.LocationId))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sTriggers := make([]*alphapb.DlpAlphaJobTriggerTriggers, len(resource.Triggers))
	for i, r := range resource.Triggers {
		sTriggers[i] = DlpAlphaJobTriggerTriggersToProto(&r)
	}
	p.SetTriggers(sTriggers)
	sErrors := make([]*alphapb.DlpAlphaJobTriggerErrors, len(resource.Errors))
	for i, r := range resource.Errors {
		sErrors[i] = DlpAlphaJobTriggerErrorsToProto(&r)
	}
	p.SetErrors(sErrors)

	return p
}

// applyJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Apply() method.
func (s *JobTriggerServer) applyJobTrigger(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDlpAlphaJobTriggerRequest) (*alphapb.DlpAlphaJobTrigger, error) {
	p := ProtoToJobTrigger(request.GetResource())
	res, err := c.ApplyJobTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobTriggerToProto(res)
	return r, nil
}

// applyDlpAlphaJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Apply() method.
func (s *JobTriggerServer) ApplyDlpAlphaJobTrigger(ctx context.Context, request *alphapb.ApplyDlpAlphaJobTriggerRequest) (*alphapb.DlpAlphaJobTrigger, error) {
	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyJobTrigger(ctx, cl, request)
}

// DeleteJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Delete() method.
func (s *JobTriggerServer) DeleteDlpAlphaJobTrigger(ctx context.Context, request *alphapb.DeleteDlpAlphaJobTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJobTrigger(ctx, ProtoToJobTrigger(request.GetResource()))

}

// ListDlpAlphaJobTrigger handles the gRPC request by passing it to the underlying JobTriggerList() method.
func (s *JobTriggerServer) ListDlpAlphaJobTrigger(ctx context.Context, request *alphapb.ListDlpAlphaJobTriggerRequest) (*alphapb.ListDlpAlphaJobTriggerResponse, error) {
	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJobTrigger(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DlpAlphaJobTrigger
	for _, r := range resources.Items {
		rp := JobTriggerToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDlpAlphaJobTriggerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigJobTrigger(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
