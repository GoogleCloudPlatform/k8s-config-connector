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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dlp/beta/dlp_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/beta"
)

// JobTriggerServer implements the gRPC interface for JobTrigger.
type JobTriggerServer struct{}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(e betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum) *beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(n[len("DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(e betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum) *beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(n[len("DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum converts a JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(e betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum) *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(n[len("DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigMinLikelihoodEnum converts a JobTriggerInspectJobInspectConfigMinLikelihoodEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(e betapb.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum) *beta.JobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobInspectConfigMinLikelihoodEnum(n[len("DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(e betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(n[len("DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum converts a JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(e betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(n[len("DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(e betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(n[len("DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(e betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum) *beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum_name[int32(e)]; ok {
		e := beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(n[len("DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerStatusEnum converts a JobTriggerStatusEnum enum from its proto representation.
func ProtoToDlpBetaJobTriggerStatusEnum(e betapb.DlpBetaJobTriggerStatusEnum) *beta.JobTriggerStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaJobTriggerStatusEnum_name[int32(e)]; ok {
		e := beta.JobTriggerStatusEnum(n[len("DlpBetaJobTriggerStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTriggerInspectJob converts a JobTriggerInspectJob object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJob(p *betapb.DlpBetaJobTriggerInspectJob) *beta.JobTriggerInspectJob {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJob{
		StorageConfig:       ProtoToDlpBetaJobTriggerInspectJobStorageConfig(p.GetStorageConfig()),
		InspectConfig:       ProtoToDlpBetaJobTriggerInspectJobInspectConfig(p.GetInspectConfig()),
		InspectTemplateName: dcl.StringOrNil(p.GetInspectTemplateName()),
	}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, *ProtoToDlpBetaJobTriggerInspectJobActions(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfig converts a JobTriggerInspectJobStorageConfig object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfig(p *betapb.DlpBetaJobTriggerInspectJobStorageConfig) *beta.JobTriggerInspectJobStorageConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfig{
		DatastoreOptions:    ProtoToDlpBetaJobTriggerInspectJobStorageConfigDatastoreOptions(p.GetDatastoreOptions()),
		CloudStorageOptions: ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptions(p.GetCloudStorageOptions()),
		BigQueryOptions:     ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptions(p.GetBigQueryOptions()),
		HybridOptions:       ProtoToDlpBetaJobTriggerInspectJobStorageConfigHybridOptions(p.GetHybridOptions()),
		TimespanConfig:      ProtoToDlpBetaJobTriggerInspectJobStorageConfigTimespanConfig(p.GetTimespanConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptions converts a JobTriggerInspectJobStorageConfigDatastoreOptions object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigDatastoreOptions(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptions) *beta.JobTriggerInspectJobStorageConfigDatastoreOptions {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigDatastoreOptions{
		PartitionId: ProtoToDlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(p.GetPartitionId()),
		Kind:        ProtoToDlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsKind(p.GetKind()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId converts a JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) *beta.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{
		ProjectId:   dcl.StringOrNil(p.GetProjectId()),
		NamespaceId: dcl.StringOrNil(p.GetNamespaceId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigDatastoreOptionsKind converts a JobTriggerInspectJobStorageConfigDatastoreOptionsKind object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsKind(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsKind) *beta.JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigDatastoreOptionsKind{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptions converts a JobTriggerInspectJobStorageConfigCloudStorageOptions object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptions(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptions) *beta.JobTriggerInspectJobStorageConfigCloudStorageOptions {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigCloudStorageOptions{
		FileSet:                  ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(p.GetFileSet()),
		BytesLimitPerFile:        dcl.Int64OrNil(p.GetBytesLimitPerFile()),
		BytesLimitPerFilePercent: dcl.Int64OrNil(p.GetBytesLimitPerFilePercent()),
		SampleMethod:             ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(p.GetSampleMethod()),
		FilesLimitPercent:        dcl.Int64OrNil(p.GetFilesLimitPercent()),
	}
	for _, r := range p.GetFileTypes() {
		obj.FileTypes = append(obj.FileTypes, *ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) *beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{
		Url:          dcl.StringOrNil(p.GetUrl()),
		RegexFileSet: ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(p.GetRegexFileSet()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) *beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{
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
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptions(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptions) *beta.JobTriggerInspectJobStorageConfigBigQueryOptions {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigBigQueryOptions{
		TableReference:   ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(p.GetTableReference()),
		RowsLimit:        dcl.Int64OrNil(p.GetRowsLimit()),
		RowsLimitPercent: dcl.Int64OrNil(p.GetRowsLimitPercent()),
		SampleMethod:     ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(p.GetSampleMethod()),
	}
	for _, r := range p.GetIdentifyingFields() {
		obj.IdentifyingFields = append(obj.IdentifyingFields, *ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(r))
	}
	for _, r := range p.GetExcludedFields() {
		obj.ExcludedFields = append(obj.ExcludedFields, *ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(r))
	}
	for _, r := range p.GetIncludedFields() {
		obj.IncludedFields = append(obj.IncludedFields, *ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference converts a JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptions converts a JobTriggerInspectJobStorageConfigHybridOptions object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigHybridOptions(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptions) *beta.JobTriggerInspectJobStorageConfigHybridOptions {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigHybridOptions{
		Description:  dcl.StringOrNil(p.GetDescription()),
		TableOptions: ProtoToDlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(p.GetTableOptions()),
	}
	for _, r := range p.GetRequiredFindingLabelKeys() {
		obj.RequiredFindingLabelKeys = append(obj.RequiredFindingLabelKeys, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptionsTableOptions converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptions object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions) *beta.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	for _, r := range p.GetIdentifyingFields() {
		obj.IdentifyingFields = append(obj.IdentifyingFields, *ProtoToDlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) *beta.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigTimespanConfig converts a JobTriggerInspectJobStorageConfigTimespanConfig object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigTimespanConfig(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigTimespanConfig) *beta.JobTriggerInspectJobStorageConfigTimespanConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigTimespanConfig{
		StartTime:                            dcl.StringOrNil(p.GetStartTime()),
		EndTime:                              dcl.StringOrNil(p.GetEndTime()),
		TimestampField:                       ProtoToDlpBetaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(p.GetTimestampField()),
		EnableAutoPopulationOfTimespanConfig: dcl.Bool(p.GetEnableAutoPopulationOfTimespanConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobStorageConfigTimespanConfigTimestampField converts a JobTriggerInspectJobStorageConfigTimespanConfigTimestampField object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(p *betapb.DlpBetaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField) *beta.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfig converts a JobTriggerInspectJobInspectConfig object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfig(p *betapb.DlpBetaJobTriggerInspectJobInspectConfig) *beta.JobTriggerInspectJobInspectConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfig{
		MinLikelihood:    ProtoToDlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(p.GetMinLikelihood()),
		Limits:           ProtoToDlpBetaJobTriggerInspectJobInspectConfigLimits(p.GetLimits()),
		IncludeQuote:     dcl.Bool(p.GetIncludeQuote()),
		ExcludeInfoTypes: dcl.Bool(p.GetExcludeInfoTypes()),
	}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpBetaJobTriggerInspectJobInspectConfigInfoTypes(r))
	}
	for _, r := range p.GetCustomInfoTypes() {
		obj.CustomInfoTypes = append(obj.CustomInfoTypes, *ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypes(r))
	}
	for _, r := range p.GetRuleSet() {
		obj.RuleSet = append(obj.RuleSet, *ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSet(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigInfoTypes converts a JobTriggerInspectJobInspectConfigInfoTypes object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigInfoTypes(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigInfoTypes) *beta.JobTriggerInspectJobInspectConfigInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimits converts a JobTriggerInspectJobInspectConfigLimits object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigLimits(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigLimits) *beta.JobTriggerInspectJobInspectConfigLimits {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigLimits{
		MaxFindingsPerItem:    dcl.Int64OrNil(p.GetMaxFindingsPerItem()),
		MaxFindingsPerRequest: dcl.Int64OrNil(p.GetMaxFindingsPerRequest()),
	}
	for _, r := range p.GetMaxFindingsPerInfoType() {
		obj.MaxFindingsPerInfoType = append(obj.MaxFindingsPerInfoType, *ProtoToDlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) *beta.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{
		InfoType:    ProtoToDlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p.GetInfoType()),
		MaxFindings: dcl.Int64OrNil(p.GetMaxFindings()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *beta.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypes converts a JobTriggerInspectJobInspectConfigCustomInfoTypes object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypes(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypes) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypes{
		InfoType:      ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(p.GetInfoType()),
		Likelihood:    ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(p.GetLikelihood()),
		Dictionary:    ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(p.GetDictionary()),
		Regex:         ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(p.GetRegex()),
		SurrogateType: ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(p.GetSurrogateType()),
		StoredType:    ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(p.GetStoredType()),
		ExclusionType: ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(p.GetExclusionType()),
	}
	for _, r := range p.GetDetectionRules() {
		obj.DetectionRules = append(obj.DetectionRules, *ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{
		WordList:         ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesRegex converts a JobTriggerInspectJobInspectConfigCustomInfoTypesRegex object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType converts a JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{
		HotwordRule: ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(p.GetHotwordRule()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{
		HotwordRegex:         ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSet converts a JobTriggerInspectJobInspectConfigRuleSet object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSet(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSet) *beta.JobTriggerInspectJobInspectConfigRuleSet {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSet{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypes(r))
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRules(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetInfoTypes object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypes(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypes) *beta.JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetInfoTypes{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRules converts a JobTriggerInspectJobInspectConfigRuleSetRules object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRules(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRules) *beta.JobTriggerInspectJobInspectConfigRuleSetRules {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRules{
		HotwordRule:   ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(p.GetHotwordRule()),
		ExclusionRule: ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(p.GetExclusionRule()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{
		HotwordRegex:         ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{
		Dictionary:       ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(p.GetDictionary()),
		Regex:            ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(p.GetRegex()),
		ExcludeInfoTypes: ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p.GetExcludeInfoTypes()),
		MatchingType:     ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(p.GetMatchingType()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{
		WordList:         ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(r))
	}
	return obj
}

// ProtoToJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(p *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{
		Name:    dcl.StringOrNil(p.GetName()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActions converts a JobTriggerInspectJobActions object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActions(p *betapb.DlpBetaJobTriggerInspectJobActions) *beta.JobTriggerInspectJobActions {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActions{
		SaveFindings:                      ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindings(p.GetSaveFindings()),
		PubSub:                            ProtoToDlpBetaJobTriggerInspectJobActionsPubSub(p.GetPubSub()),
		PublishSummaryToCscc:              ProtoToDlpBetaJobTriggerInspectJobActionsPublishSummaryToCscc(p.GetPublishSummaryToCscc()),
		PublishFindingsToCloudDataCatalog: ProtoToDlpBetaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(p.GetPublishFindingsToCloudDataCatalog()),
		JobNotificationEmails:             ProtoToDlpBetaJobTriggerInspectJobActionsJobNotificationEmails(p.GetJobNotificationEmails()),
		PublishToStackdriver:              ProtoToDlpBetaJobTriggerInspectJobActionsPublishToStackdriver(p.GetPublishToStackdriver()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindings converts a JobTriggerInspectJobActionsSaveFindings object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindings(p *betapb.DlpBetaJobTriggerInspectJobActionsSaveFindings) *beta.JobTriggerInspectJobActionsSaveFindings {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsSaveFindings{
		OutputConfig: ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfig(p.GetOutputConfig()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfig converts a JobTriggerInspectJobActionsSaveFindingsOutputConfig object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfig(p *betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfig) *beta.JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsSaveFindingsOutputConfig{
		Table:        ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(p.GetTable()),
		DlpStorage:   ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(p.GetDlpStorage()),
		OutputSchema: ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(p.GetOutputSchema()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigTable converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigTable object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(p *betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable) *beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(p *betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) *beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPubSub converts a JobTriggerInspectJobActionsPubSub object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsPubSub(p *betapb.DlpBetaJobTriggerInspectJobActionsPubSub) *beta.JobTriggerInspectJobActionsPubSub {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsPubSub{
		Topic: dcl.StringOrNil(p.GetTopic()),
	}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishSummaryToCscc converts a JobTriggerInspectJobActionsPublishSummaryToCscc object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsPublishSummaryToCscc(p *betapb.DlpBetaJobTriggerInspectJobActionsPublishSummaryToCscc) *beta.JobTriggerInspectJobActionsPublishSummaryToCscc {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsPublishSummaryToCscc{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog converts a JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(p *betapb.DlpBetaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) *beta.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsJobNotificationEmails converts a JobTriggerInspectJobActionsJobNotificationEmails object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsJobNotificationEmails(p *betapb.DlpBetaJobTriggerInspectJobActionsJobNotificationEmails) *beta.JobTriggerInspectJobActionsJobNotificationEmails {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsJobNotificationEmails{}
	return obj
}

// ProtoToJobTriggerInspectJobActionsPublishToStackdriver converts a JobTriggerInspectJobActionsPublishToStackdriver object from its proto representation.
func ProtoToDlpBetaJobTriggerInspectJobActionsPublishToStackdriver(p *betapb.DlpBetaJobTriggerInspectJobActionsPublishToStackdriver) *beta.JobTriggerInspectJobActionsPublishToStackdriver {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerInspectJobActionsPublishToStackdriver{}
	return obj
}

// ProtoToJobTriggerTriggers converts a JobTriggerTriggers object from its proto representation.
func ProtoToDlpBetaJobTriggerTriggers(p *betapb.DlpBetaJobTriggerTriggers) *beta.JobTriggerTriggers {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerTriggers{
		Schedule: ProtoToDlpBetaJobTriggerTriggersSchedule(p.GetSchedule()),
		Manual:   ProtoToDlpBetaJobTriggerTriggersManual(p.GetManual()),
	}
	return obj
}

// ProtoToJobTriggerTriggersSchedule converts a JobTriggerTriggersSchedule object from its proto representation.
func ProtoToDlpBetaJobTriggerTriggersSchedule(p *betapb.DlpBetaJobTriggerTriggersSchedule) *beta.JobTriggerTriggersSchedule {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerTriggersSchedule{
		RecurrencePeriodDuration: dcl.StringOrNil(p.GetRecurrencePeriodDuration()),
	}
	return obj
}

// ProtoToJobTriggerTriggersManual converts a JobTriggerTriggersManual object from its proto representation.
func ProtoToDlpBetaJobTriggerTriggersManual(p *betapb.DlpBetaJobTriggerTriggersManual) *beta.JobTriggerTriggersManual {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerTriggersManual{}
	return obj
}

// ProtoToJobTriggerErrors converts a JobTriggerErrors object from its proto representation.
func ProtoToDlpBetaJobTriggerErrors(p *betapb.DlpBetaJobTriggerErrors) *beta.JobTriggerErrors {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerErrors{
		Details: ProtoToDlpBetaJobTriggerErrorsDetails(p.GetDetails()),
	}
	for _, r := range p.GetTimestamps() {
		obj.Timestamps = append(obj.Timestamps, r)
	}
	return obj
}

// ProtoToJobTriggerErrorsDetails converts a JobTriggerErrorsDetails object from its proto representation.
func ProtoToDlpBetaJobTriggerErrorsDetails(p *betapb.DlpBetaJobTriggerErrorsDetails) *beta.JobTriggerErrorsDetails {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerErrorsDetails{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToDlpBetaJobTriggerErrorsDetailsDetails(r))
	}
	return obj
}

// ProtoToJobTriggerErrorsDetailsDetails converts a JobTriggerErrorsDetailsDetails object from its proto representation.
func ProtoToDlpBetaJobTriggerErrorsDetailsDetails(p *betapb.DlpBetaJobTriggerErrorsDetailsDetails) *beta.JobTriggerErrorsDetailsDetails {
	if p == nil {
		return nil
	}
	obj := &beta.JobTriggerErrorsDetailsDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToJobTrigger converts a JobTrigger resource from its proto representation.
func ProtoToJobTrigger(p *betapb.DlpBetaJobTrigger) *beta.JobTrigger {
	obj := &beta.JobTrigger{
		Name:        dcl.StringOrNil(p.GetName()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		InspectJob:  ProtoToDlpBetaJobTriggerInspectJob(p.GetInspectJob()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		LastRunTime: dcl.StringOrNil(p.GetLastRunTime()),
		Status:      ProtoToDlpBetaJobTriggerStatusEnum(p.GetStatus()),
		LocationId:  dcl.StringOrNil(p.GetLocationId()),
		Parent:      dcl.StringOrNil(p.GetParent()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetTriggers() {
		obj.Triggers = append(obj.Triggers, *ProtoToDlpBetaJobTriggerTriggers(r))
	}
	for _, r := range p.GetErrors() {
		obj.Errors = append(obj.Errors, *ProtoToDlpBetaJobTriggerErrors(r))
	}
	return obj
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumToProto(e *beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum) betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_value["JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(0)
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto(e *beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum) betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum_value["JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(0)
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto(e *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum) betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum_value["JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(0)
}

// JobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigMinLikelihoodEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto(e *beta.JobTriggerInspectJobInspectConfigMinLikelihoodEnum) betapb.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum_value["JobTriggerInspectJobInspectConfigMinLikelihoodEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto(e *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum) betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto(e *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum) betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum_value["JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(0)
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(e *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_value["JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum enum to its proto representation.
func DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto(e *beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum) betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum_value["JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(v)
	}
	return betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(0)
}

// JobTriggerStatusEnumToProto converts a JobTriggerStatusEnum enum to its proto representation.
func DlpBetaJobTriggerStatusEnumToProto(e *beta.JobTriggerStatusEnum) betapb.DlpBetaJobTriggerStatusEnum {
	if e == nil {
		return betapb.DlpBetaJobTriggerStatusEnum(0)
	}
	if v, ok := betapb.DlpBetaJobTriggerStatusEnum_value["JobTriggerStatusEnum"+string(*e)]; ok {
		return betapb.DlpBetaJobTriggerStatusEnum(v)
	}
	return betapb.DlpBetaJobTriggerStatusEnum(0)
}

// JobTriggerInspectJobToProto converts a JobTriggerInspectJob object to its proto representation.
func DlpBetaJobTriggerInspectJobToProto(o *beta.JobTriggerInspectJob) *betapb.DlpBetaJobTriggerInspectJob {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJob{}
	p.SetStorageConfig(DlpBetaJobTriggerInspectJobStorageConfigToProto(o.StorageConfig))
	p.SetInspectConfig(DlpBetaJobTriggerInspectJobInspectConfigToProto(o.InspectConfig))
	p.SetInspectTemplateName(dcl.ValueOrEmptyString(o.InspectTemplateName))
	sActions := make([]*betapb.DlpBetaJobTriggerInspectJobActions, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = DlpBetaJobTriggerInspectJobActionsToProto(&r)
	}
	p.SetActions(sActions)
	return p
}

// JobTriggerInspectJobStorageConfigToProto converts a JobTriggerInspectJobStorageConfig object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigToProto(o *beta.JobTriggerInspectJobStorageConfig) *betapb.DlpBetaJobTriggerInspectJobStorageConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfig{}
	p.SetDatastoreOptions(DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsToProto(o.DatastoreOptions))
	p.SetCloudStorageOptions(DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsToProto(o.CloudStorageOptions))
	p.SetBigQueryOptions(DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsToProto(o.BigQueryOptions))
	p.SetHybridOptions(DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsToProto(o.HybridOptions))
	p.SetTimespanConfig(DlpBetaJobTriggerInspectJobStorageConfigTimespanConfigToProto(o.TimespanConfig))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptions object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsToProto(o *beta.JobTriggerInspectJobStorageConfigDatastoreOptions) *betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptions{}
	p.SetPartitionId(DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto(o.PartitionId))
	p.SetKind(DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto(o.Kind))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdToProto(o *beta.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) *betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetNamespaceId(dcl.ValueOrEmptyString(o.NamespaceId))
	return p
}

// JobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto converts a JobTriggerInspectJobStorageConfigDatastoreOptionsKind object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsKindToProto(o *beta.JobTriggerInspectJobStorageConfigDatastoreOptionsKind) *betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptions object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsToProto(o *beta.JobTriggerInspectJobStorageConfigCloudStorageOptions) *betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptions{}
	p.SetFileSet(DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto(o.FileSet))
	p.SetBytesLimitPerFile(dcl.ValueOrEmptyInt64(o.BytesLimitPerFile))
	p.SetBytesLimitPerFilePercent(dcl.ValueOrEmptyInt64(o.BytesLimitPerFilePercent))
	p.SetSampleMethod(DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumToProto(o.SampleMethod))
	p.SetFilesLimitPercent(dcl.ValueOrEmptyInt64(o.FilesLimitPercent))
	sFileTypes := make([]betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum, len(o.FileTypes))
	for i, r := range o.FileTypes {
		sFileTypes[i] = betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum_value[string(r)])
	}
	p.SetFileTypes(sFileTypes)
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetToProto(o *beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) *betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetRegexFileSet(DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto(o.RegexFileSet))
	return p
}

// JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto converts a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetToProto(o *beta.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) *betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
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
func DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsToProto(o *beta.JobTriggerInspectJobStorageConfigBigQueryOptions) *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptions{}
	p.SetTableReference(DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto(o.TableReference))
	p.SetRowsLimit(dcl.ValueOrEmptyInt64(o.RowsLimit))
	p.SetRowsLimitPercent(dcl.ValueOrEmptyInt64(o.RowsLimitPercent))
	p.SetSampleMethod(DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumToProto(o.SampleMethod))
	sIdentifyingFields := make([]*betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, len(o.IdentifyingFields))
	for i, r := range o.IdentifyingFields {
		sIdentifyingFields[i] = DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto(&r)
	}
	p.SetIdentifyingFields(sIdentifyingFields)
	sExcludedFields := make([]*betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, len(o.ExcludedFields))
	for i, r := range o.ExcludedFields {
		sExcludedFields[i] = DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto(&r)
	}
	p.SetExcludedFields(sExcludedFields)
	sIncludedFields := make([]*betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, len(o.IncludedFields))
	for i, r := range o.IncludedFields {
		sIncludedFields[i] = DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto(&r)
	}
	p.SetIncludedFields(sIncludedFields)
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceToProto(o *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsToProto(o *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsToProto(o *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto converts a JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsToProto(o *beta.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) *betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigHybridOptionsToProto converts a JobTriggerInspectJobStorageConfigHybridOptions object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsToProto(o *beta.JobTriggerInspectJobStorageConfigHybridOptions) *betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptions{}
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetTableOptions(DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsToProto(o.TableOptions))
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
func DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsToProto(o *beta.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) *betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	sIdentifyingFields := make([]*betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, len(o.IdentifyingFields))
	for i, r := range o.IdentifyingFields {
		sIdentifyingFields[i] = DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto(&r)
	}
	p.SetIdentifyingFields(sIdentifyingFields)
	return p
}

// JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto converts a JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsToProto(o *beta.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) *betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobStorageConfigTimespanConfigToProto converts a JobTriggerInspectJobStorageConfigTimespanConfig object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigTimespanConfigToProto(o *beta.JobTriggerInspectJobStorageConfigTimespanConfig) *betapb.DlpBetaJobTriggerInspectJobStorageConfigTimespanConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigTimespanConfig{}
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetEndTime(dcl.ValueOrEmptyString(o.EndTime))
	p.SetTimestampField(DlpBetaJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto(o.TimestampField))
	p.SetEnableAutoPopulationOfTimespanConfig(dcl.ValueOrEmptyBool(o.EnableAutoPopulationOfTimespanConfig))
	return p
}

// JobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto converts a JobTriggerInspectJobStorageConfigTimespanConfigTimestampField object to its proto representation.
func DlpBetaJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldToProto(o *beta.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) *betapb.DlpBetaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobInspectConfigToProto converts a JobTriggerInspectJobInspectConfig object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigToProto(o *beta.JobTriggerInspectJobInspectConfig) *betapb.DlpBetaJobTriggerInspectJobInspectConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfig{}
	p.SetMinLikelihood(DlpBetaJobTriggerInspectJobInspectConfigMinLikelihoodEnumToProto(o.MinLikelihood))
	p.SetLimits(DlpBetaJobTriggerInspectJobInspectConfigLimitsToProto(o.Limits))
	p.SetIncludeQuote(dcl.ValueOrEmptyBool(o.IncludeQuote))
	p.SetExcludeInfoTypes(dcl.ValueOrEmptyBool(o.ExcludeInfoTypes))
	sInfoTypes := make([]*betapb.DlpBetaJobTriggerInspectJobInspectConfigInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpBetaJobTriggerInspectJobInspectConfigInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sCustomInfoTypes := make([]*betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypes, len(o.CustomInfoTypes))
	for i, r := range o.CustomInfoTypes {
		sCustomInfoTypes[i] = DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesToProto(&r)
	}
	p.SetCustomInfoTypes(sCustomInfoTypes)
	sRuleSet := make([]*betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSet, len(o.RuleSet))
	for i, r := range o.RuleSet {
		sRuleSet[i] = DlpBetaJobTriggerInspectJobInspectConfigRuleSetToProto(&r)
	}
	p.SetRuleSet(sRuleSet)
	return p
}

// JobTriggerInspectJobInspectConfigInfoTypesToProto converts a JobTriggerInspectJobInspectConfigInfoTypes object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigInfoTypesToProto(o *beta.JobTriggerInspectJobInspectConfigInfoTypes) *betapb.DlpBetaJobTriggerInspectJobInspectConfigInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// JobTriggerInspectJobInspectConfigLimitsToProto converts a JobTriggerInspectJobInspectConfigLimits object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigLimitsToProto(o *beta.JobTriggerInspectJobInspectConfigLimits) *betapb.DlpBetaJobTriggerInspectJobInspectConfigLimits {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigLimits{}
	p.SetMaxFindingsPerItem(dcl.ValueOrEmptyInt64(o.MaxFindingsPerItem))
	p.SetMaxFindingsPerRequest(dcl.ValueOrEmptyInt64(o.MaxFindingsPerRequest))
	sMaxFindingsPerInfoType := make([]*betapb.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, len(o.MaxFindingsPerInfoType))
	for i, r := range o.MaxFindingsPerInfoType {
		sMaxFindingsPerInfoType[i] = DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto(&r)
	}
	p.SetMaxFindingsPerInfoType(sMaxFindingsPerInfoType)
	return p
}

// JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeToProto(o *beta.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) *betapb.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}
	p.SetInfoType(DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o.InfoType))
	p.SetMaxFindings(dcl.ValueOrEmptyInt64(o.MaxFindings))
	return p
}

// JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto converts a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o *beta.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *betapb.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypes object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypes) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypes{}
	p.SetInfoType(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto(o.InfoType))
	p.SetLikelihood(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumToProto(o.Likelihood))
	p.SetDictionary(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto(o.Regex))
	p.SetSurrogateType(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto(o.SurrogateType))
	p.SetStoredType(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto(o.StoredType))
	p.SetExclusionType(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumToProto(o.ExclusionType))
	sDetectionRules := make([]*betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, len(o.DetectionRules))
	for i, r := range o.DetectionRules {
		sDetectionRules[i] = DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto(&r)
	}
	p.SetDetectionRules(sDetectionRules)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	p.SetWordList(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesRegex object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesRegexToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}
	p.SetHotwordRule(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto(o.HotwordRule))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	p.SetHotwordRegex(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto converts a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentToProto(o *beta.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) *betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpBetaJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetToProto converts a JobTriggerInspectJobInspectConfigRuleSet object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSet) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSet {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSet{}
	sInfoTypes := make([]*betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sRules := make([]*betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetInfoTypes object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypesToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetInfoTypes) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRules object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRules) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRules {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRules{}
	p.SetHotwordRule(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto(o.HotwordRule))
	p.SetExclusionRule(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto(o.ExclusionRule))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	p.SetHotwordRegex(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	p.SetDictionary(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto(o.Regex))
	p.SetExcludeInfoTypes(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o.ExcludeInfoTypes))
	p.SetMatchingType(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(o.MatchingType))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	p.SetWordList(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	sInfoTypes := make([]*betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	return p
}

// JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto converts a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object to its proto representation.
func DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(o *beta.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTriggerInspectJobActionsToProto converts a JobTriggerInspectJobActions object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsToProto(o *beta.JobTriggerInspectJobActions) *betapb.DlpBetaJobTriggerInspectJobActions {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActions{}
	p.SetSaveFindings(DlpBetaJobTriggerInspectJobActionsSaveFindingsToProto(o.SaveFindings))
	p.SetPubSub(DlpBetaJobTriggerInspectJobActionsPubSubToProto(o.PubSub))
	p.SetPublishSummaryToCscc(DlpBetaJobTriggerInspectJobActionsPublishSummaryToCsccToProto(o.PublishSummaryToCscc))
	p.SetPublishFindingsToCloudDataCatalog(DlpBetaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto(o.PublishFindingsToCloudDataCatalog))
	p.SetJobNotificationEmails(DlpBetaJobTriggerInspectJobActionsJobNotificationEmailsToProto(o.JobNotificationEmails))
	p.SetPublishToStackdriver(DlpBetaJobTriggerInspectJobActionsPublishToStackdriverToProto(o.PublishToStackdriver))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsToProto converts a JobTriggerInspectJobActionsSaveFindings object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsSaveFindingsToProto(o *beta.JobTriggerInspectJobActionsSaveFindings) *betapb.DlpBetaJobTriggerInspectJobActionsSaveFindings {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsSaveFindings{}
	p.SetOutputConfig(DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigToProto(o.OutputConfig))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfig object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigToProto(o *beta.JobTriggerInspectJobActionsSaveFindingsOutputConfig) *betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	p.SetTable(DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto(o.Table))
	p.SetDlpStorage(DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto(o.DlpStorage))
	p.SetOutputSchema(DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumToProto(o.OutputSchema))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigTable object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigTableToProto(o *beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) *betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto converts a JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageToProto(o *beta.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) *betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	return p
}

// JobTriggerInspectJobActionsPubSubToProto converts a JobTriggerInspectJobActionsPubSub object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsPubSubToProto(o *beta.JobTriggerInspectJobActionsPubSub) *betapb.DlpBetaJobTriggerInspectJobActionsPubSub {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsPubSub{}
	p.SetTopic(dcl.ValueOrEmptyString(o.Topic))
	return p
}

// JobTriggerInspectJobActionsPublishSummaryToCsccToProto converts a JobTriggerInspectJobActionsPublishSummaryToCscc object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsPublishSummaryToCsccToProto(o *beta.JobTriggerInspectJobActionsPublishSummaryToCscc) *betapb.DlpBetaJobTriggerInspectJobActionsPublishSummaryToCscc {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsPublishSummaryToCscc{}
	return p
}

// JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto converts a JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogToProto(o *beta.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) *betapb.DlpBetaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	return p
}

// JobTriggerInspectJobActionsJobNotificationEmailsToProto converts a JobTriggerInspectJobActionsJobNotificationEmails object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsJobNotificationEmailsToProto(o *beta.JobTriggerInspectJobActionsJobNotificationEmails) *betapb.DlpBetaJobTriggerInspectJobActionsJobNotificationEmails {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsJobNotificationEmails{}
	return p
}

// JobTriggerInspectJobActionsPublishToStackdriverToProto converts a JobTriggerInspectJobActionsPublishToStackdriver object to its proto representation.
func DlpBetaJobTriggerInspectJobActionsPublishToStackdriverToProto(o *beta.JobTriggerInspectJobActionsPublishToStackdriver) *betapb.DlpBetaJobTriggerInspectJobActionsPublishToStackdriver {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerInspectJobActionsPublishToStackdriver{}
	return p
}

// JobTriggerTriggersToProto converts a JobTriggerTriggers object to its proto representation.
func DlpBetaJobTriggerTriggersToProto(o *beta.JobTriggerTriggers) *betapb.DlpBetaJobTriggerTriggers {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerTriggers{}
	p.SetSchedule(DlpBetaJobTriggerTriggersScheduleToProto(o.Schedule))
	p.SetManual(DlpBetaJobTriggerTriggersManualToProto(o.Manual))
	return p
}

// JobTriggerTriggersScheduleToProto converts a JobTriggerTriggersSchedule object to its proto representation.
func DlpBetaJobTriggerTriggersScheduleToProto(o *beta.JobTriggerTriggersSchedule) *betapb.DlpBetaJobTriggerTriggersSchedule {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerTriggersSchedule{}
	p.SetRecurrencePeriodDuration(dcl.ValueOrEmptyString(o.RecurrencePeriodDuration))
	return p
}

// JobTriggerTriggersManualToProto converts a JobTriggerTriggersManual object to its proto representation.
func DlpBetaJobTriggerTriggersManualToProto(o *beta.JobTriggerTriggersManual) *betapb.DlpBetaJobTriggerTriggersManual {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerTriggersManual{}
	return p
}

// JobTriggerErrorsToProto converts a JobTriggerErrors object to its proto representation.
func DlpBetaJobTriggerErrorsToProto(o *beta.JobTriggerErrors) *betapb.DlpBetaJobTriggerErrors {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerErrors{}
	p.SetDetails(DlpBetaJobTriggerErrorsDetailsToProto(o.Details))
	sTimestamps := make([]string, len(o.Timestamps))
	for i, r := range o.Timestamps {
		sTimestamps[i] = r
	}
	p.SetTimestamps(sTimestamps)
	return p
}

// JobTriggerErrorsDetailsToProto converts a JobTriggerErrorsDetails object to its proto representation.
func DlpBetaJobTriggerErrorsDetailsToProto(o *beta.JobTriggerErrorsDetails) *betapb.DlpBetaJobTriggerErrorsDetails {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerErrorsDetails{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*betapb.DlpBetaJobTriggerErrorsDetailsDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = DlpBetaJobTriggerErrorsDetailsDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// JobTriggerErrorsDetailsDetailsToProto converts a JobTriggerErrorsDetailsDetails object to its proto representation.
func DlpBetaJobTriggerErrorsDetailsDetailsToProto(o *beta.JobTriggerErrorsDetailsDetails) *betapb.DlpBetaJobTriggerErrorsDetailsDetails {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaJobTriggerErrorsDetailsDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// JobTriggerToProto converts a JobTrigger resource to its proto representation.
func JobTriggerToProto(resource *beta.JobTrigger) *betapb.DlpBetaJobTrigger {
	p := &betapb.DlpBetaJobTrigger{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInspectJob(DlpBetaJobTriggerInspectJobToProto(resource.InspectJob))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetLastRunTime(dcl.ValueOrEmptyString(resource.LastRunTime))
	p.SetStatus(DlpBetaJobTriggerStatusEnumToProto(resource.Status))
	p.SetLocationId(dcl.ValueOrEmptyString(resource.LocationId))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sTriggers := make([]*betapb.DlpBetaJobTriggerTriggers, len(resource.Triggers))
	for i, r := range resource.Triggers {
		sTriggers[i] = DlpBetaJobTriggerTriggersToProto(&r)
	}
	p.SetTriggers(sTriggers)
	sErrors := make([]*betapb.DlpBetaJobTriggerErrors, len(resource.Errors))
	for i, r := range resource.Errors {
		sErrors[i] = DlpBetaJobTriggerErrorsToProto(&r)
	}
	p.SetErrors(sErrors)

	return p
}

// applyJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Apply() method.
func (s *JobTriggerServer) applyJobTrigger(ctx context.Context, c *beta.Client, request *betapb.ApplyDlpBetaJobTriggerRequest) (*betapb.DlpBetaJobTrigger, error) {
	p := ProtoToJobTrigger(request.GetResource())
	res, err := c.ApplyJobTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobTriggerToProto(res)
	return r, nil
}

// applyDlpBetaJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Apply() method.
func (s *JobTriggerServer) ApplyDlpBetaJobTrigger(ctx context.Context, request *betapb.ApplyDlpBetaJobTriggerRequest) (*betapb.DlpBetaJobTrigger, error) {
	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyJobTrigger(ctx, cl, request)
}

// DeleteJobTrigger handles the gRPC request by passing it to the underlying JobTrigger Delete() method.
func (s *JobTriggerServer) DeleteDlpBetaJobTrigger(ctx context.Context, request *betapb.DeleteDlpBetaJobTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJobTrigger(ctx, ProtoToJobTrigger(request.GetResource()))

}

// ListDlpBetaJobTrigger handles the gRPC request by passing it to the underlying JobTriggerList() method.
func (s *JobTriggerServer) ListDlpBetaJobTrigger(ctx context.Context, request *betapb.ListDlpBetaJobTriggerRequest) (*betapb.ListDlpBetaJobTriggerResponse, error) {
	cl, err := createConfigJobTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJobTrigger(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DlpBetaJobTrigger
	for _, r := range resources.Items {
		rp := JobTriggerToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDlpBetaJobTriggerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigJobTrigger(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
