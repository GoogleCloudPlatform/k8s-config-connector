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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *JobTrigger) validate() error {

	if err := dcl.Required(r, "inspectJob"); err != nil {
		return err
	}
	if err := dcl.Required(r, "triggers"); err != nil {
		return err
	}
	if err := dcl.Required(r, "status"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Parent, "Parent"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.InspectJob) {
		if err := r.InspectJob.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJob) validate() error {
	if err := dcl.Required(r, "storageConfig"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.StorageConfig) {
		if err := r.StorageConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.InspectConfig) {
		if err := r.InspectConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfig) validate() error {
	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"DatastoreOptions", "CloudStorageOptions", "BigQueryOptions", "HybridOptions"}, r.DatastoreOptions, r.CloudStorageOptions, r.BigQueryOptions, r.HybridOptions); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"DatastoreOptions", "CloudStorageOptions", "BigQueryOptions", "HybridOptions"}, r.DatastoreOptions, r.CloudStorageOptions, r.BigQueryOptions, r.HybridOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DatastoreOptions) {
		if err := r.DatastoreOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudStorageOptions) {
		if err := r.CloudStorageOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BigQueryOptions) {
		if err := r.BigQueryOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HybridOptions) {
		if err := r.HybridOptions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.TimespanConfig) {
		if err := r.TimespanConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfigDatastoreOptions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.PartitionId) {
		if err := r.PartitionId.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Kind) {
		if err := r.Kind.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) validate() error {
	return nil
}
func (r *JobTriggerInspectJobStorageConfigDatastoreOptionsKind) validate() error {
	return nil
}
func (r *JobTriggerInspectJobStorageConfigCloudStorageOptions) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"BytesLimitPerFile", "BytesLimitPerFilePercent"}, r.BytesLimitPerFile, r.BytesLimitPerFilePercent); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.FileSet) {
		if err := r.FileSet.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) validate() error {
	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"Url", "RegexFileSet"}, r.Url, r.RegexFileSet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.RegexFileSet) {
		if err := r.RegexFileSet.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) validate() error {
	if err := dcl.Required(r, "bucketName"); err != nil {
		return err
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfigBigQueryOptions) validate() error {
	if err := dcl.Required(r, "tableReference"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"RowsLimit", "RowsLimitPercent"}, r.RowsLimit, r.RowsLimitPercent); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"ExcludedFields", "IncludedFields"}, r.ExcludedFields, r.IncludedFields); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.TableReference) {
		if err := r.TableReference.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) validate() error {
	return nil
}
func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) validate() error {
	return nil
}
func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) validate() error {
	return nil
}
func (r *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) validate() error {
	return nil
}
func (r *JobTriggerInspectJobStorageConfigHybridOptions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TableOptions) {
		if err := r.TableOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) validate() error {
	return nil
}
func (r *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) validate() error {
	return nil
}
func (r *JobTriggerInspectJobStorageConfigTimespanConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TimestampField) {
		if err := r.TimestampField.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Limits) {
		if err := r.Limits.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigInfoTypes) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigLimits) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) validate() error {
	if !dcl.IsEmptyValueIndirect(r.InfoType) {
		if err := r.InfoType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypes) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Dictionary", "Regex", "SurrogateType", "StoredType"}, r.Dictionary, r.Regex, r.SurrogateType, r.StoredType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.InfoType) {
		if err := r.InfoType.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Dictionary) {
		if err := r.Dictionary.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Regex) {
		if err := r.Regex.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SurrogateType) {
		if err := r.SurrogateType.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.StoredType) {
		if err := r.StoredType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"WordList", "CloudStoragePath"}, r.WordList, r.CloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.WordList) {
		if err := r.WordList.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudStoragePath) {
		if err := r.CloudStoragePath.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"HotwordRule"}, r.HotwordRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.HotwordRule) {
		if err := r.HotwordRule.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HotwordRegex) {
		if err := r.HotwordRegex.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Proximity) {
		if err := r.Proximity.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LikelihoodAdjustment) {
		if err := r.LikelihoodAdjustment.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FixedLikelihood", "RelativeLikelihood"}, r.FixedLikelihood, r.RelativeLikelihood); err != nil {
		return err
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSet) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetInfoTypes) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRules) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"HotwordRule", "ExclusionRule"}, r.HotwordRule, r.ExclusionRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.HotwordRule) {
		if err := r.HotwordRule.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExclusionRule) {
		if err := r.ExclusionRule.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HotwordRegex) {
		if err := r.HotwordRegex.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Proximity) {
		if err := r.Proximity.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LikelihoodAdjustment) {
		if err := r.LikelihoodAdjustment.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FixedLikelihood", "RelativeLikelihood"}, r.FixedLikelihood, r.RelativeLikelihood); err != nil {
		return err
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Dictionary", "Regex", "ExcludeInfoTypes"}, r.Dictionary, r.Regex, r.ExcludeInfoTypes); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Dictionary) {
		if err := r.Dictionary.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Regex) {
		if err := r.Regex.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExcludeInfoTypes) {
		if err := r.ExcludeInfoTypes.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"WordList", "CloudStoragePath"}, r.WordList, r.CloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.WordList) {
		if err := r.WordList.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudStoragePath) {
		if err := r.CloudStoragePath.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) validate() error {
	return nil
}
func (r *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) validate() error {
	return nil
}
func (r *JobTriggerInspectJobActions) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"SaveFindings", "PubSub", "PublishSummaryToCscc", "PublishFindingsToCloudDataCatalog", "JobNotificationEmails", "PublishToStackdriver"}, r.SaveFindings, r.PubSub, r.PublishSummaryToCscc, r.PublishFindingsToCloudDataCatalog, r.JobNotificationEmails, r.PublishToStackdriver); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SaveFindings) {
		if err := r.SaveFindings.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PubSub) {
		if err := r.PubSub.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PublishSummaryToCscc) {
		if err := r.PublishSummaryToCscc.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PublishFindingsToCloudDataCatalog) {
		if err := r.PublishFindingsToCloudDataCatalog.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.JobNotificationEmails) {
		if err := r.JobNotificationEmails.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PublishToStackdriver) {
		if err := r.PublishToStackdriver.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobActionsSaveFindings) validate() error {
	if !dcl.IsEmptyValueIndirect(r.OutputConfig) {
		if err := r.OutputConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfig) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Table", "DlpStorage"}, r.Table, r.DlpStorage); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Table) {
		if err := r.Table.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DlpStorage) {
		if err := r.DlpStorage.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) validate() error {
	return nil
}
func (r *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) validate() error {
	return nil
}
func (r *JobTriggerInspectJobActionsPubSub) validate() error {
	return nil
}
func (r *JobTriggerInspectJobActionsPublishSummaryToCscc) validate() error {
	return nil
}
func (r *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) validate() error {
	return nil
}
func (r *JobTriggerInspectJobActionsJobNotificationEmails) validate() error {
	return nil
}
func (r *JobTriggerInspectJobActionsPublishToStackdriver) validate() error {
	return nil
}
func (r *JobTriggerTriggers) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Schedule", "Manual"}, r.Schedule, r.Manual); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Schedule) {
		if err := r.Schedule.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Manual) {
		if err := r.Manual.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerTriggersSchedule) validate() error {
	return nil
}
func (r *JobTriggerTriggersManual) validate() error {
	return nil
}
func (r *JobTriggerErrors) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Details) {
		if err := r.Details.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *JobTriggerErrorsDetails) validate() error {
	return nil
}
func (r *JobTriggerErrorsDetailsDetails) validate() error {
	return nil
}
func (r *JobTrigger) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://dlp.googleapis.com/v2/", params)
}

func (r *JobTrigger) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/jobTriggers/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/jobTriggers/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *JobTrigger) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/jobTriggers", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/jobTriggers", nr.basePath(), userBasePath, params), nil

}

func (r *JobTrigger) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/jobTriggers", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/jobTriggers", nr.basePath(), userBasePath, params), nil

}

func (r *JobTrigger) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/jobTriggers/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/jobTriggers/{{name}}", nr.basePath(), userBasePath, params), nil
}

// jobTriggerApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type jobTriggerApiOperation interface {
	do(context.Context, *JobTrigger, *Client) error
}

// newUpdateJobTriggerUpdateJobTriggerRequest creates a request for an
// JobTrigger resource's UpdateJobTrigger update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateJobTriggerUpdateJobTriggerRequest(ctx context.Context, f *JobTrigger, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandJobTriggerInspectJob(c, f.InspectJob, res); err != nil {
		return nil, fmt.Errorf("error expanding InspectJob into inspectJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["inspectJob"] = v
	}
	if v, err := expandJobTriggerTriggersSlice(c, f.Triggers, res); err != nil {
		return nil, fmt.Errorf("error expanding Triggers into triggers: %w", err)
	} else if v != nil {
		req["triggers"] = v
	}
	return req, nil
}

// marshalUpdateJobTriggerUpdateJobTriggerRequest converts the update into
// the final JSON request body.
func marshalUpdateJobTriggerUpdateJobTriggerRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateJobTriggerUpdateJobTriggerOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listJobTriggerRaw(ctx context.Context, r *JobTrigger, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != JobTriggerMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listJobTriggerOperation struct {
	JobTriggers []map[string]interface{} `json:"jobTriggers"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listJobTrigger(ctx context.Context, r *JobTrigger, pageToken string, pageSize int32) ([]*JobTrigger, string, error) {
	b, err := c.listJobTriggerRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listJobTriggerOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*JobTrigger
	for _, v := range m.JobTriggers {
		res, err := unmarshalMapJobTrigger(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Location = r.Location
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllJobTrigger(ctx context.Context, f func(*JobTrigger) bool, resources []*JobTrigger) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteJobTrigger(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteJobTriggerOperation struct{}

func (op *deleteJobTriggerOperation) do(ctx context.Context, r *JobTrigger, c *Client) error {
	r, err := c.GetJobTrigger(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "JobTrigger not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetJobTrigger checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete JobTrigger: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createJobTriggerOperation struct {
	response map[string]interface{}
}

func (op *createJobTriggerOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createJobTriggerOperation) do(ctx context.Context, r *JobTrigger, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	if r.Name != nil {
		// Allowing creation to continue with Name set could result in a JobTrigger with the wrong Name.
		return fmt.Errorf("server-generated parameter Name was specified by user as %v, should be unspecified", dcl.ValueOrEmptyString(r.Name))
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	// Include Name in URL substitution for initial GET request.
	m := op.response
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))

	if _, err := c.GetJobTrigger(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getJobTriggerRaw(ctx context.Context, r *JobTrigger) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) jobTriggerDiffsForRawDesired(ctx context.Context, rawDesired *JobTrigger, opts ...dcl.ApplyOption) (initial, desired *JobTrigger, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *JobTrigger
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*JobTrigger); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected JobTrigger, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeJobTriggerDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetJobTrigger(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a JobTrigger resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve JobTrigger resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that JobTrigger resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeJobTriggerDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for JobTrigger: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for JobTrigger: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractJobTriggerFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeJobTriggerInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for JobTrigger: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeJobTriggerDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for JobTrigger: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffJobTrigger(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeJobTriggerInitialState(rawInitial, rawDesired *JobTrigger) (*JobTrigger, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeJobTriggerDesiredState(rawDesired, rawInitial *JobTrigger, opts ...dcl.ApplyOption) (*JobTrigger, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.InspectJob = canonicalizeJobTriggerInspectJob(rawDesired.InspectJob, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &JobTrigger{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	canonicalDesired.InspectJob = canonicalizeJobTriggerInspectJob(rawDesired.InspectJob, rawInitial.InspectJob, opts...)
	canonicalDesired.Triggers = canonicalizeJobTriggerTriggersSlice(rawDesired.Triggers, rawInitial.Triggers, opts...)
	if dcl.IsZeroValue(rawDesired.Status) || (dcl.IsEmptyValueIndirect(rawDesired.Status) && dcl.IsEmptyValueIndirect(rawInitial.Status)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Status = rawInitial.Status
	} else {
		canonicalDesired.Status = rawDesired.Status
	}
	if dcl.NameToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	return canonicalDesired, nil
}

func canonicalizeJobTriggerNewState(c *Client, rawNew, rawDesired *JobTrigger) (*JobTrigger, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.InspectJob) && dcl.IsEmptyValueIndirect(rawDesired.InspectJob) {
		rawNew.InspectJob = rawDesired.InspectJob
	} else {
		rawNew.InspectJob = canonicalizeNewJobTriggerInspectJob(c, rawDesired.InspectJob, rawNew.InspectJob)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Triggers) && dcl.IsEmptyValueIndirect(rawDesired.Triggers) {
		rawNew.Triggers = rawDesired.Triggers
	} else {
		rawNew.Triggers = canonicalizeNewJobTriggerTriggersSlice(c, rawDesired.Triggers, rawNew.Triggers)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Errors) && dcl.IsEmptyValueIndirect(rawDesired.Errors) {
		rawNew.Errors = rawDesired.Errors
	} else {
		rawNew.Errors = canonicalizeNewJobTriggerErrorsSlice(c, rawDesired.Errors, rawNew.Errors)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastRunTime) && dcl.IsEmptyValueIndirect(rawDesired.LastRunTime) {
		rawNew.LastRunTime = rawDesired.LastRunTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Status) && dcl.IsEmptyValueIndirect(rawDesired.Status) {
		rawNew.Status = rawDesired.Status
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LocationId) && dcl.IsEmptyValueIndirect(rawDesired.LocationId) {
		rawNew.LocationId = rawDesired.LocationId
	} else {
		if dcl.StringCanonicalize(rawDesired.LocationId, rawNew.LocationId) {
			rawNew.LocationId = rawDesired.LocationId
		}
	}

	rawNew.Parent = rawDesired.Parent

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeJobTriggerInspectJob(des, initial *JobTriggerInspectJob, opts ...dcl.ApplyOption) *JobTriggerInspectJob {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJob{}

	cDes.StorageConfig = canonicalizeJobTriggerInspectJobStorageConfig(des.StorageConfig, initial.StorageConfig, opts...)
	cDes.InspectConfig = canonicalizeJobTriggerInspectJobInspectConfig(des.InspectConfig, initial.InspectConfig, opts...)
	if dcl.IsZeroValue(des.InspectTemplateName) || (dcl.IsEmptyValueIndirect(des.InspectTemplateName) && dcl.IsEmptyValueIndirect(initial.InspectTemplateName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.InspectTemplateName = initial.InspectTemplateName
	} else {
		cDes.InspectTemplateName = des.InspectTemplateName
	}
	cDes.Actions = canonicalizeJobTriggerInspectJobActionsSlice(des.Actions, initial.Actions, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobSlice(des, initial []JobTriggerInspectJob, opts ...dcl.ApplyOption) []JobTriggerInspectJob {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJob, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJob(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJob, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJob(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJob(c *Client, des, nw *JobTriggerInspectJob) *JobTriggerInspectJob {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJob while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.StorageConfig = canonicalizeNewJobTriggerInspectJobStorageConfig(c, des.StorageConfig, nw.StorageConfig)
	nw.InspectConfig = canonicalizeNewJobTriggerInspectJobInspectConfig(c, des.InspectConfig, nw.InspectConfig)
	nw.Actions = canonicalizeNewJobTriggerInspectJobActionsSlice(c, des.Actions, nw.Actions)

	return nw
}

func canonicalizeNewJobTriggerInspectJobSet(c *Client, des, nw []JobTriggerInspectJob) []JobTriggerInspectJob {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJob
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJob(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobSlice(c *Client, des, nw []JobTriggerInspectJob) []JobTriggerInspectJob {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJob
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJob(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfig(des, initial *JobTriggerInspectJobStorageConfig, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.DatastoreOptions != nil || (initial != nil && initial.DatastoreOptions != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudStorageOptions, des.BigQueryOptions, des.HybridOptions) {
			des.DatastoreOptions = nil
			if initial != nil {
				initial.DatastoreOptions = nil
			}
		}
	}

	if des.CloudStorageOptions != nil || (initial != nil && initial.CloudStorageOptions != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.DatastoreOptions, des.BigQueryOptions, des.HybridOptions) {
			des.CloudStorageOptions = nil
			if initial != nil {
				initial.CloudStorageOptions = nil
			}
		}
	}

	if des.BigQueryOptions != nil || (initial != nil && initial.BigQueryOptions != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.DatastoreOptions, des.CloudStorageOptions, des.HybridOptions) {
			des.BigQueryOptions = nil
			if initial != nil {
				initial.BigQueryOptions = nil
			}
		}
	}

	if des.HybridOptions != nil || (initial != nil && initial.HybridOptions != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.DatastoreOptions, des.CloudStorageOptions, des.BigQueryOptions) {
			des.HybridOptions = nil
			if initial != nil {
				initial.HybridOptions = nil
			}
		}
	}

	if des.DatastoreOptions != nil || (initial != nil && initial.DatastoreOptions != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudStorageOptions, des.BigQueryOptions, des.HybridOptions) {
			des.DatastoreOptions = nil
			if initial != nil {
				initial.DatastoreOptions = nil
			}
		}
	}

	if des.CloudStorageOptions != nil || (initial != nil && initial.CloudStorageOptions != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.DatastoreOptions, des.BigQueryOptions, des.HybridOptions) {
			des.CloudStorageOptions = nil
			if initial != nil {
				initial.CloudStorageOptions = nil
			}
		}
	}

	if des.BigQueryOptions != nil || (initial != nil && initial.BigQueryOptions != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.DatastoreOptions, des.CloudStorageOptions, des.HybridOptions) {
			des.BigQueryOptions = nil
			if initial != nil {
				initial.BigQueryOptions = nil
			}
		}
	}

	if des.HybridOptions != nil || (initial != nil && initial.HybridOptions != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.DatastoreOptions, des.CloudStorageOptions, des.BigQueryOptions) {
			des.HybridOptions = nil
			if initial != nil {
				initial.HybridOptions = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfig{}

	cDes.DatastoreOptions = canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptions(des.DatastoreOptions, initial.DatastoreOptions, opts...)
	cDes.CloudStorageOptions = canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptions(des.CloudStorageOptions, initial.CloudStorageOptions, opts...)
	cDes.BigQueryOptions = canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptions(des.BigQueryOptions, initial.BigQueryOptions, opts...)
	cDes.HybridOptions = canonicalizeJobTriggerInspectJobStorageConfigHybridOptions(des.HybridOptions, initial.HybridOptions, opts...)
	cDes.TimespanConfig = canonicalizeJobTriggerInspectJobStorageConfigTimespanConfig(des.TimespanConfig, initial.TimespanConfig, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigSlice(des, initial []JobTriggerInspectJobStorageConfig, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfig(c *Client, des, nw *JobTriggerInspectJobStorageConfig) *JobTriggerInspectJobStorageConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.DatastoreOptions = canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptions(c, des.DatastoreOptions, nw.DatastoreOptions)
	nw.CloudStorageOptions = canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptions(c, des.CloudStorageOptions, nw.CloudStorageOptions)
	nw.BigQueryOptions = canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptions(c, des.BigQueryOptions, nw.BigQueryOptions)
	nw.HybridOptions = canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptions(c, des.HybridOptions, nw.HybridOptions)
	nw.TimespanConfig = canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfig(c, des.TimespanConfig, nw.TimespanConfig)

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigSet(c *Client, des, nw []JobTriggerInspectJobStorageConfig) []JobTriggerInspectJobStorageConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfig) []JobTriggerInspectJobStorageConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfig(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptions(des, initial *JobTriggerInspectJobStorageConfigDatastoreOptions, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigDatastoreOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigDatastoreOptions{}

	cDes.PartitionId = canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(des.PartitionId, initial.PartitionId, opts...)
	cDes.Kind = canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsKind(des.Kind, initial.Kind, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsSlice(des, initial []JobTriggerInspectJobStorageConfigDatastoreOptions, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigDatastoreOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigDatastoreOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigDatastoreOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptions(c *Client, des, nw *JobTriggerInspectJobStorageConfigDatastoreOptions) *JobTriggerInspectJobStorageConfigDatastoreOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigDatastoreOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.PartitionId = canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, des.PartitionId, nw.PartitionId)
	nw.Kind = canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, des.Kind, nw.Kind)

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigDatastoreOptions) []JobTriggerInspectJobStorageConfigDatastoreOptions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigDatastoreOptions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigDatastoreOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigDatastoreOptions) []JobTriggerInspectJobStorageConfigDatastoreOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigDatastoreOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptions(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(des, initial *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}

	if dcl.IsZeroValue(des.ProjectId) || (dcl.IsEmptyValueIndirect(des.ProjectId) && dcl.IsEmptyValueIndirect(initial.ProjectId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ProjectId = initial.ProjectId
	} else {
		cDes.ProjectId = des.ProjectId
	}
	if dcl.StringCanonicalize(des.NamespaceId, initial.NamespaceId) || dcl.IsZeroValue(des.NamespaceId) {
		cDes.NamespaceId = initial.NamespaceId
	} else {
		cDes.NamespaceId = des.NamespaceId
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdSlice(des, initial []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c *Client, des, nw *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.NamespaceId, nw.NamespaceId) {
		nw.NamespaceId = des.NamespaceId
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsKind(des, initial *JobTriggerInspectJobStorageConfigDatastoreOptionsKind, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsKindSlice(des, initial []JobTriggerInspectJobStorageConfigDatastoreOptionsKind, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigDatastoreOptionsKind, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsKind(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigDatastoreOptionsKind, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigDatastoreOptionsKind(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c *Client, des, nw *JobTriggerInspectJobStorageConfigDatastoreOptionsKind) *JobTriggerInspectJobStorageConfigDatastoreOptionsKind {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigDatastoreOptionsKind while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsKindSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigDatastoreOptionsKind) []JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigDatastoreOptionsKind
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigDatastoreOptionsKindNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsKindSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigDatastoreOptionsKind) []JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigDatastoreOptionsKind
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptions(des, initial *JobTriggerInspectJobStorageConfigCloudStorageOptions, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigCloudStorageOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.BytesLimitPerFile != nil || (initial != nil && initial.BytesLimitPerFile != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.BytesLimitPerFilePercent) {
			des.BytesLimitPerFile = nil
			if initial != nil {
				initial.BytesLimitPerFile = nil
			}
		}
	}

	if des.BytesLimitPerFilePercent != nil || (initial != nil && initial.BytesLimitPerFilePercent != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.BytesLimitPerFile) {
			des.BytesLimitPerFilePercent = nil
			if initial != nil {
				initial.BytesLimitPerFilePercent = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigCloudStorageOptions{}

	cDes.FileSet = canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(des.FileSet, initial.FileSet, opts...)
	if dcl.IsZeroValue(des.BytesLimitPerFile) || (dcl.IsEmptyValueIndirect(des.BytesLimitPerFile) && dcl.IsEmptyValueIndirect(initial.BytesLimitPerFile)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BytesLimitPerFile = initial.BytesLimitPerFile
	} else {
		cDes.BytesLimitPerFile = des.BytesLimitPerFile
	}
	if dcl.IsZeroValue(des.BytesLimitPerFilePercent) || (dcl.IsEmptyValueIndirect(des.BytesLimitPerFilePercent) && dcl.IsEmptyValueIndirect(initial.BytesLimitPerFilePercent)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BytesLimitPerFilePercent = initial.BytesLimitPerFilePercent
	} else {
		cDes.BytesLimitPerFilePercent = des.BytesLimitPerFilePercent
	}
	if dcl.IsZeroValue(des.FileTypes) || (dcl.IsEmptyValueIndirect(des.FileTypes) && dcl.IsEmptyValueIndirect(initial.FileTypes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.FileTypes = initial.FileTypes
	} else {
		cDes.FileTypes = des.FileTypes
	}
	if dcl.IsZeroValue(des.SampleMethod) || (dcl.IsEmptyValueIndirect(des.SampleMethod) && dcl.IsEmptyValueIndirect(initial.SampleMethod)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SampleMethod = initial.SampleMethod
	} else {
		cDes.SampleMethod = des.SampleMethod
	}
	if dcl.IsZeroValue(des.FilesLimitPercent) || (dcl.IsEmptyValueIndirect(des.FilesLimitPercent) && dcl.IsEmptyValueIndirect(initial.FilesLimitPercent)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.FilesLimitPercent = initial.FilesLimitPercent
	} else {
		cDes.FilesLimitPercent = des.FilesLimitPercent
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsSlice(des, initial []JobTriggerInspectJobStorageConfigCloudStorageOptions, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigCloudStorageOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptions(c *Client, des, nw *JobTriggerInspectJobStorageConfigCloudStorageOptions) *JobTriggerInspectJobStorageConfigCloudStorageOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigCloudStorageOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.FileSet = canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, des.FileSet, nw.FileSet)

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigCloudStorageOptions) []JobTriggerInspectJobStorageConfigCloudStorageOptions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigCloudStorageOptions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigCloudStorageOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigCloudStorageOptions) []JobTriggerInspectJobStorageConfigCloudStorageOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigCloudStorageOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptions(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(des, initial *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Url != nil || (initial != nil && initial.Url != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RegexFileSet) {
			des.Url = nil
			if initial != nil {
				initial.Url = nil
			}
		}
	}

	if des.RegexFileSet != nil || (initial != nil && initial.RegexFileSet != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Url) {
			des.RegexFileSet = nil
			if initial != nil {
				initial.RegexFileSet = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}

	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		cDes.Url = initial.Url
	} else {
		cDes.Url = des.Url
	}
	cDes.RegexFileSet = canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(des.RegexFileSet, initial.RegexFileSet, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetSlice(des, initial []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c *Client, des, nw *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}
	nw.RegexFileSet = canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, des.RegexFileSet, nw.RegexFileSet)

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(des, initial *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}

	if dcl.IsZeroValue(des.BucketName) || (dcl.IsEmptyValueIndirect(des.BucketName) && dcl.IsEmptyValueIndirect(initial.BucketName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BucketName = initial.BucketName
	} else {
		cDes.BucketName = des.BucketName
	}
	if dcl.StringArrayCanonicalize(des.IncludeRegex, initial.IncludeRegex) {
		cDes.IncludeRegex = initial.IncludeRegex
	} else {
		cDes.IncludeRegex = des.IncludeRegex
	}
	if dcl.StringArrayCanonicalize(des.ExcludeRegex, initial.ExcludeRegex) {
		cDes.ExcludeRegex = initial.ExcludeRegex
	} else {
		cDes.ExcludeRegex = des.ExcludeRegex
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetSlice(des, initial []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c *Client, des, nw *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.IncludeRegex, nw.IncludeRegex) {
		nw.IncludeRegex = des.IncludeRegex
	}
	if dcl.StringArrayCanonicalize(des.ExcludeRegex, nw.ExcludeRegex) {
		nw.ExcludeRegex = des.ExcludeRegex
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptions(des, initial *JobTriggerInspectJobStorageConfigBigQueryOptions, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigBigQueryOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.RowsLimit != nil || (initial != nil && initial.RowsLimit != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RowsLimitPercent) {
			des.RowsLimit = nil
			if initial != nil {
				initial.RowsLimit = nil
			}
		}
	}

	if des.RowsLimitPercent != nil || (initial != nil && initial.RowsLimitPercent != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RowsLimit) {
			des.RowsLimitPercent = nil
			if initial != nil {
				initial.RowsLimitPercent = nil
			}
		}
	}

	if des.ExcludedFields != nil || (initial != nil && initial.ExcludedFields != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.IncludedFields) {
			des.ExcludedFields = nil
			if initial != nil {
				initial.ExcludedFields = nil
			}
		}
	}

	if des.IncludedFields != nil || (initial != nil && initial.IncludedFields != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExcludedFields) {
			des.IncludedFields = nil
			if initial != nil {
				initial.IncludedFields = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigBigQueryOptions{}

	cDes.TableReference = canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(des.TableReference, initial.TableReference, opts...)
	cDes.IdentifyingFields = canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice(des.IdentifyingFields, initial.IdentifyingFields, opts...)
	if dcl.IsZeroValue(des.RowsLimit) || (dcl.IsEmptyValueIndirect(des.RowsLimit) && dcl.IsEmptyValueIndirect(initial.RowsLimit)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RowsLimit = initial.RowsLimit
	} else {
		cDes.RowsLimit = des.RowsLimit
	}
	if dcl.IsZeroValue(des.RowsLimitPercent) || (dcl.IsEmptyValueIndirect(des.RowsLimitPercent) && dcl.IsEmptyValueIndirect(initial.RowsLimitPercent)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RowsLimitPercent = initial.RowsLimitPercent
	} else {
		cDes.RowsLimitPercent = des.RowsLimitPercent
	}
	if dcl.IsZeroValue(des.SampleMethod) || (dcl.IsEmptyValueIndirect(des.SampleMethod) && dcl.IsEmptyValueIndirect(initial.SampleMethod)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SampleMethod = initial.SampleMethod
	} else {
		cDes.SampleMethod = des.SampleMethod
	}
	cDes.ExcludedFields = canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice(des.ExcludedFields, initial.ExcludedFields, opts...)
	cDes.IncludedFields = canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice(des.IncludedFields, initial.IncludedFields, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsSlice(des, initial []JobTriggerInspectJobStorageConfigBigQueryOptions, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigBigQueryOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigBigQueryOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptions(c *Client, des, nw *JobTriggerInspectJobStorageConfigBigQueryOptions) *JobTriggerInspectJobStorageConfigBigQueryOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigBigQueryOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.TableReference = canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, des.TableReference, nw.TableReference)
	nw.IdentifyingFields = canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice(c, des.IdentifyingFields, nw.IdentifyingFields)
	nw.ExcludedFields = canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice(c, des.ExcludedFields, nw.ExcludedFields)
	nw.IncludedFields = canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice(c, des.IncludedFields, nw.IncludedFields)

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptions) []JobTriggerInspectJobStorageConfigBigQueryOptions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigBigQueryOptions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigBigQueryOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptions) []JobTriggerInspectJobStorageConfigBigQueryOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigBigQueryOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptions(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(des, initial *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}

	if dcl.IsZeroValue(des.ProjectId) || (dcl.IsEmptyValueIndirect(des.ProjectId) && dcl.IsEmptyValueIndirect(initial.ProjectId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ProjectId = initial.ProjectId
	} else {
		cDes.ProjectId = des.ProjectId
	}
	if dcl.IsZeroValue(des.DatasetId) || (dcl.IsEmptyValueIndirect(des.DatasetId) && dcl.IsEmptyValueIndirect(initial.DatasetId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DatasetId = initial.DatasetId
	} else {
		cDes.DatasetId = des.DatasetId
	}
	if dcl.IsZeroValue(des.TableId) || (dcl.IsEmptyValueIndirect(des.TableId) && dcl.IsEmptyValueIndirect(initial.TableId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TableId = initial.TableId
	} else {
		cDes.TableId = des.TableId
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceSlice(des, initial []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c *Client, des, nw *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(des, initial *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice(des, initial []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c *Client, des, nw *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(des, initial *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice(des, initial []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c *Client, des, nw *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(des, initial *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice(des, initial []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c *Client, des, nw *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigHybridOptions(des, initial *JobTriggerInspectJobStorageConfigHybridOptions, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigHybridOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigHybridOptions{}

	if dcl.StringCanonicalize(des.Description, initial.Description) || dcl.IsZeroValue(des.Description) {
		cDes.Description = initial.Description
	} else {
		cDes.Description = des.Description
	}
	if dcl.StringArrayCanonicalize(des.RequiredFindingLabelKeys, initial.RequiredFindingLabelKeys) {
		cDes.RequiredFindingLabelKeys = initial.RequiredFindingLabelKeys
	} else {
		cDes.RequiredFindingLabelKeys = des.RequiredFindingLabelKeys
	}
	if dcl.IsZeroValue(des.Labels) || (dcl.IsEmptyValueIndirect(des.Labels) && dcl.IsEmptyValueIndirect(initial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Labels = initial.Labels
	} else {
		cDes.Labels = des.Labels
	}
	cDes.TableOptions = canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(des.TableOptions, initial.TableOptions, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsSlice(des, initial []JobTriggerInspectJobStorageConfigHybridOptions, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigHybridOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigHybridOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigHybridOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigHybridOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigHybridOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptions(c *Client, des, nw *JobTriggerInspectJobStorageConfigHybridOptions) *JobTriggerInspectJobStorageConfigHybridOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigHybridOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Description, nw.Description) {
		nw.Description = des.Description
	}
	if dcl.StringArrayCanonicalize(des.RequiredFindingLabelKeys, nw.RequiredFindingLabelKeys) {
		nw.RequiredFindingLabelKeys = des.RequiredFindingLabelKeys
	}
	nw.TableOptions = canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, des.TableOptions, nw.TableOptions)

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigHybridOptions) []JobTriggerInspectJobStorageConfigHybridOptions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigHybridOptions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigHybridOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigHybridOptions) []JobTriggerInspectJobStorageConfigHybridOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigHybridOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptions(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(des, initial *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}

	cDes.IdentifyingFields = canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice(des.IdentifyingFields, initial.IdentifyingFields, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsSlice(des, initial []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigHybridOptionsTableOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigHybridOptionsTableOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c *Client, des, nw *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigHybridOptionsTableOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.IdentifyingFields = canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice(c, des.IdentifyingFields, nw.IdentifyingFields)

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(des, initial *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice(des, initial []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c *Client, des, nw *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigTimespanConfig(des, initial *JobTriggerInspectJobStorageConfigTimespanConfig, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigTimespanConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigTimespanConfig{}

	if dcl.IsZeroValue(des.StartTime) || (dcl.IsEmptyValueIndirect(des.StartTime) && dcl.IsEmptyValueIndirect(initial.StartTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.StartTime = initial.StartTime
	} else {
		cDes.StartTime = des.StartTime
	}
	if dcl.IsZeroValue(des.EndTime) || (dcl.IsEmptyValueIndirect(des.EndTime) && dcl.IsEmptyValueIndirect(initial.EndTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EndTime = initial.EndTime
	} else {
		cDes.EndTime = des.EndTime
	}
	cDes.TimestampField = canonicalizeJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(des.TimestampField, initial.TimestampField, opts...)
	if dcl.BoolCanonicalize(des.EnableAutoPopulationOfTimespanConfig, initial.EnableAutoPopulationOfTimespanConfig) || dcl.IsZeroValue(des.EnableAutoPopulationOfTimespanConfig) {
		cDes.EnableAutoPopulationOfTimespanConfig = initial.EnableAutoPopulationOfTimespanConfig
	} else {
		cDes.EnableAutoPopulationOfTimespanConfig = des.EnableAutoPopulationOfTimespanConfig
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigTimespanConfigSlice(des, initial []JobTriggerInspectJobStorageConfigTimespanConfig, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigTimespanConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigTimespanConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigTimespanConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigTimespanConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigTimespanConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfig(c *Client, des, nw *JobTriggerInspectJobStorageConfigTimespanConfig) *JobTriggerInspectJobStorageConfigTimespanConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigTimespanConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.TimestampField = canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, des.TimestampField, nw.TimestampField)
	if dcl.BoolCanonicalize(des.EnableAutoPopulationOfTimespanConfig, nw.EnableAutoPopulationOfTimespanConfig) {
		nw.EnableAutoPopulationOfTimespanConfig = des.EnableAutoPopulationOfTimespanConfig
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfigSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigTimespanConfig) []JobTriggerInspectJobStorageConfigTimespanConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigTimespanConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigTimespanConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfigSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigTimespanConfig) []JobTriggerInspectJobStorageConfigTimespanConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigTimespanConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfig(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(des, initial *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField, opts ...dcl.ApplyOption) *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldSlice(des, initial []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField, opts ...dcl.ApplyOption) []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobStorageConfigTimespanConfigTimestampField, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobStorageConfigTimespanConfigTimestampField, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c *Client, des, nw *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobStorageConfigTimespanConfigTimestampField while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldSet(c *Client, des, nw []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldSlice(c *Client, des, nw []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfig(des, initial *JobTriggerInspectJobInspectConfig, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfig{}

	cDes.InfoTypes = canonicalizeJobTriggerInspectJobInspectConfigInfoTypesSlice(des.InfoTypes, initial.InfoTypes, opts...)
	if dcl.IsZeroValue(des.MinLikelihood) || (dcl.IsEmptyValueIndirect(des.MinLikelihood) && dcl.IsEmptyValueIndirect(initial.MinLikelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MinLikelihood = initial.MinLikelihood
	} else {
		cDes.MinLikelihood = des.MinLikelihood
	}
	cDes.Limits = canonicalizeJobTriggerInspectJobInspectConfigLimits(des.Limits, initial.Limits, opts...)
	if dcl.BoolCanonicalize(des.IncludeQuote, initial.IncludeQuote) || dcl.IsZeroValue(des.IncludeQuote) {
		cDes.IncludeQuote = initial.IncludeQuote
	} else {
		cDes.IncludeQuote = des.IncludeQuote
	}
	if dcl.BoolCanonicalize(des.ExcludeInfoTypes, initial.ExcludeInfoTypes) || dcl.IsZeroValue(des.ExcludeInfoTypes) {
		cDes.ExcludeInfoTypes = initial.ExcludeInfoTypes
	} else {
		cDes.ExcludeInfoTypes = des.ExcludeInfoTypes
	}
	cDes.CustomInfoTypes = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesSlice(des.CustomInfoTypes, initial.CustomInfoTypes, opts...)
	cDes.RuleSet = canonicalizeJobTriggerInspectJobInspectConfigRuleSetSlice(des.RuleSet, initial.RuleSet, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigSlice(des, initial []JobTriggerInspectJobInspectConfig, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfig(c *Client, des, nw *JobTriggerInspectJobInspectConfig) *JobTriggerInspectJobInspectConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoTypes = canonicalizeNewJobTriggerInspectJobInspectConfigInfoTypesSlice(c, des.InfoTypes, nw.InfoTypes)
	nw.Limits = canonicalizeNewJobTriggerInspectJobInspectConfigLimits(c, des.Limits, nw.Limits)
	if dcl.BoolCanonicalize(des.IncludeQuote, nw.IncludeQuote) {
		nw.IncludeQuote = des.IncludeQuote
	}
	if dcl.BoolCanonicalize(des.ExcludeInfoTypes, nw.ExcludeInfoTypes) {
		nw.ExcludeInfoTypes = des.ExcludeInfoTypes
	}
	nw.CustomInfoTypes = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSlice(c, des.CustomInfoTypes, nw.CustomInfoTypes)
	nw.RuleSet = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetSlice(c, des.RuleSet, nw.RuleSet)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigSet(c *Client, des, nw []JobTriggerInspectJobInspectConfig) []JobTriggerInspectJobInspectConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfig) []JobTriggerInspectJobInspectConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfig(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigInfoTypes(des, initial *JobTriggerInspectJobInspectConfigInfoTypes, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigInfoTypes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigInfoTypesSlice(des, initial []JobTriggerInspectJobInspectConfigInfoTypes, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigInfoTypes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigInfoTypes(c *Client, des, nw *JobTriggerInspectJobInspectConfigInfoTypes) *JobTriggerInspectJobInspectConfigInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigInfoTypesSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigInfoTypes) []JobTriggerInspectJobInspectConfigInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigInfoTypesSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigInfoTypes) []JobTriggerInspectJobInspectConfigInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigLimits(des, initial *JobTriggerInspectJobInspectConfigLimits, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigLimits {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigLimits{}

	if dcl.IsZeroValue(des.MaxFindingsPerItem) || (dcl.IsEmptyValueIndirect(des.MaxFindingsPerItem) && dcl.IsEmptyValueIndirect(initial.MaxFindingsPerItem)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxFindingsPerItem = initial.MaxFindingsPerItem
	} else {
		cDes.MaxFindingsPerItem = des.MaxFindingsPerItem
	}
	if dcl.IsZeroValue(des.MaxFindingsPerRequest) || (dcl.IsEmptyValueIndirect(des.MaxFindingsPerRequest) && dcl.IsEmptyValueIndirect(initial.MaxFindingsPerRequest)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxFindingsPerRequest = initial.MaxFindingsPerRequest
	} else {
		cDes.MaxFindingsPerRequest = des.MaxFindingsPerRequest
	}
	cDes.MaxFindingsPerInfoType = canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice(des.MaxFindingsPerInfoType, initial.MaxFindingsPerInfoType, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigLimitsSlice(des, initial []JobTriggerInspectJobInspectConfigLimits, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigLimits {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigLimits, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigLimits(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigLimits, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigLimits(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimits(c *Client, des, nw *JobTriggerInspectJobInspectConfigLimits) *JobTriggerInspectJobInspectConfigLimits {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigLimits while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MaxFindingsPerInfoType = canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c, des.MaxFindingsPerInfoType, nw.MaxFindingsPerInfoType)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimitsSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigLimits) []JobTriggerInspectJobInspectConfigLimits {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigLimits
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigLimitsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigLimits(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimitsSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigLimits) []JobTriggerInspectJobInspectConfigLimits {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigLimits
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigLimits(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(des, initial *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}

	cDes.InfoType = canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(des.InfoType, initial.InfoType, opts...)
	if dcl.IsZeroValue(des.MaxFindings) || (dcl.IsEmptyValueIndirect(des.MaxFindings) && dcl.IsEmptyValueIndirect(initial.MaxFindings)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxFindings = initial.MaxFindings
	} else {
		cDes.MaxFindings = des.MaxFindings
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice(des, initial []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c *Client, des, nw *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoType = canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, des.InfoType, nw.InfoType)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(des, initial *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice(des, initial []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c *Client, des, nw *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypes(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypes, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Dictionary != nil || (initial != nil && initial.Dictionary != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Regex, des.SurrogateType, des.StoredType) {
			des.Dictionary = nil
			if initial != nil {
				initial.Dictionary = nil
			}
		}
	}

	if des.Regex != nil || (initial != nil && initial.Regex != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.SurrogateType, des.StoredType) {
			des.Regex = nil
			if initial != nil {
				initial.Regex = nil
			}
		}
	}

	if des.SurrogateType != nil || (initial != nil && initial.SurrogateType != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.Regex, des.StoredType) {
			des.SurrogateType = nil
			if initial != nil {
				initial.SurrogateType = nil
			}
		}
	}

	if des.StoredType != nil || (initial != nil && initial.StoredType != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.Regex, des.SurrogateType) {
			des.StoredType = nil
			if initial != nil {
				initial.StoredType = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypes{}

	cDes.InfoType = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(des.InfoType, initial.InfoType, opts...)
	if dcl.IsZeroValue(des.Likelihood) || (dcl.IsEmptyValueIndirect(des.Likelihood) && dcl.IsEmptyValueIndirect(initial.Likelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Likelihood = initial.Likelihood
	} else {
		cDes.Likelihood = des.Likelihood
	}
	cDes.Dictionary = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(des.Dictionary, initial.Dictionary, opts...)
	cDes.Regex = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(des.Regex, initial.Regex, opts...)
	cDes.SurrogateType = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(des.SurrogateType, initial.SurrogateType, opts...)
	cDes.StoredType = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(des.StoredType, initial.StoredType, opts...)
	cDes.DetectionRules = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice(des.DetectionRules, initial.DetectionRules, opts...)
	if dcl.IsZeroValue(des.ExclusionType) || (dcl.IsEmptyValueIndirect(des.ExclusionType) && dcl.IsEmptyValueIndirect(initial.ExclusionType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ExclusionType = initial.ExclusionType
	} else {
		cDes.ExclusionType = des.ExclusionType
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypes, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypes(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypes) *JobTriggerInspectJobInspectConfigCustomInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoType = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, des.InfoType, nw.InfoType)
	nw.Dictionary = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, des.Dictionary, nw.Dictionary)
	nw.Regex = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, des.Regex, nw.Regex)
	nw.SurrogateType = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, des.SurrogateType, nw.SurrogateType)
	nw.StoredType = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, des.StoredType, nw.StoredType)
	nw.DetectionRules = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice(c, des.DetectionRules, nw.DetectionRules)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypes) []JobTriggerInspectJobInspectConfigCustomInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypes) []JobTriggerInspectJobInspectConfigCustomInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.WordList != nil || (initial != nil && initial.WordList != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudStoragePath) {
			des.WordList = nil
			if initial != nil {
				initial.WordList = nil
			}
		}
	}

	if des.CloudStoragePath != nil || (initial != nil && initial.CloudStoragePath != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.WordList) {
			des.CloudStoragePath = nil
			if initial != nil {
				initial.CloudStoragePath = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}

	cDes.WordList = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(des.WordList, initial.WordList, opts...)
	cDes.CloudStoragePath = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(des.CloudStoragePath, initial.CloudStoragePath, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionarySlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.WordList = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, des.WordList, nw.WordList)
	nw.CloudStoragePath = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, des.CloudStoragePath, nw.CloudStoragePath)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionarySet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionarySlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}

	if dcl.StringArrayCanonicalize(des.Words, initial.Words) {
		cDes.Words = initial.Words
	} else {
		cDes.Words = des.Words
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Words, nw.Words) {
		nw.Words = des.Words
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}

	if dcl.StringCanonicalize(des.Pattern, initial.Pattern) || dcl.IsZeroValue(des.Pattern) {
		cDes.Pattern = initial.Pattern
	} else {
		cDes.Pattern = des.Pattern
	}
	if dcl.IsZeroValue(des.GroupIndexes) || (dcl.IsEmptyValueIndirect(des.GroupIndexes) && dcl.IsEmptyValueIndirect(initial.GroupIndexes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GroupIndexes = initial.GroupIndexes
	} else {
		cDes.GroupIndexes = des.GroupIndexes
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesRegexSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesRegex, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesRegex, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesRegex while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Pattern, nw.Pattern) {
		nw.Pattern = des.Pattern
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesRegexSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesRegexNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesRegexSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}

	if dcl.IsZeroValue(des.Name) || (dcl.IsEmptyValueIndirect(des.Name) && dcl.IsEmptyValueIndirect(initial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.HotwordRule != nil || (initial != nil && initial.HotwordRule != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.HotwordRule = nil
			if initial != nil {
				initial.HotwordRule = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}

	cDes.HotwordRule = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(des.HotwordRule, initial.HotwordRule, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.HotwordRule = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, des.HotwordRule, nw.HotwordRule)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}

	cDes.HotwordRegex = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(des.HotwordRegex, initial.HotwordRegex, opts...)
	cDes.Proximity = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(des.Proximity, initial.Proximity, opts...)
	cDes.LikelihoodAdjustment = canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(des.LikelihoodAdjustment, initial.LikelihoodAdjustment, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.HotwordRegex = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, des.HotwordRegex, nw.HotwordRegex)
	nw.Proximity = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, des.Proximity, nw.Proximity)
	nw.LikelihoodAdjustment = canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, des.LikelihoodAdjustment, nw.LikelihoodAdjustment)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}

	if dcl.StringCanonicalize(des.Pattern, initial.Pattern) || dcl.IsZeroValue(des.Pattern) {
		cDes.Pattern = initial.Pattern
	} else {
		cDes.Pattern = des.Pattern
	}
	if dcl.IsZeroValue(des.GroupIndexes) || (dcl.IsEmptyValueIndirect(des.GroupIndexes) && dcl.IsEmptyValueIndirect(initial.GroupIndexes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GroupIndexes = initial.GroupIndexes
	} else {
		cDes.GroupIndexes = des.GroupIndexes
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Pattern, nw.Pattern) {
		nw.Pattern = des.Pattern
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}

	if dcl.IsZeroValue(des.WindowBefore) || (dcl.IsEmptyValueIndirect(des.WindowBefore) && dcl.IsEmptyValueIndirect(initial.WindowBefore)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.WindowBefore = initial.WindowBefore
	} else {
		cDes.WindowBefore = des.WindowBefore
	}
	if dcl.IsZeroValue(des.WindowAfter) || (dcl.IsEmptyValueIndirect(des.WindowAfter) && dcl.IsEmptyValueIndirect(initial.WindowAfter)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.WindowAfter = initial.WindowAfter
	} else {
		cDes.WindowAfter = des.WindowAfter
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximitySlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximitySet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximitySlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(des, initial *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.FixedLikelihood != nil || (initial != nil && initial.FixedLikelihood != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RelativeLikelihood) {
			des.FixedLikelihood = nil
			if initial != nil {
				initial.FixedLikelihood = nil
			}
		}
	}

	if des.RelativeLikelihood != nil || (initial != nil && initial.RelativeLikelihood != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.FixedLikelihood) {
			des.RelativeLikelihood = nil
			if initial != nil {
				initial.RelativeLikelihood = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}

	if dcl.IsZeroValue(des.FixedLikelihood) || (dcl.IsEmptyValueIndirect(des.FixedLikelihood) && dcl.IsEmptyValueIndirect(initial.FixedLikelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.FixedLikelihood = initial.FixedLikelihood
	} else {
		cDes.FixedLikelihood = des.FixedLikelihood
	}
	if dcl.IsZeroValue(des.RelativeLikelihood) || (dcl.IsEmptyValueIndirect(des.RelativeLikelihood) && dcl.IsEmptyValueIndirect(initial.RelativeLikelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RelativeLikelihood = initial.RelativeLikelihood
	} else {
		cDes.RelativeLikelihood = des.RelativeLikelihood
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentSlice(des, initial []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c *Client, des, nw *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSet(des, initial *JobTriggerInspectJobInspectConfigRuleSet, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSet {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSet{}

	cDes.InfoTypes = canonicalizeJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice(des.InfoTypes, initial.InfoTypes, opts...)
	cDes.Rules = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesSlice(des.Rules, initial.Rules, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSet, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSet {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSet, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSet(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSet, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSet(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSet(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSet) *JobTriggerInspectJobInspectConfigRuleSet {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSet while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoTypes = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice(c, des.InfoTypes, nw.InfoTypes)
	nw.Rules = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesSlice(c, des.Rules, nw.Rules)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSet) []JobTriggerInspectJobInspectConfigRuleSet {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSet
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSet(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSet) []JobTriggerInspectJobInspectConfigRuleSet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSet(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetInfoTypes(des, initial *JobTriggerInspectJobInspectConfigRuleSetInfoTypes, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetInfoTypes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetInfoTypes, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetInfoTypes) *JobTriggerInspectJobInspectConfigRuleSetInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetInfoTypesSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetInfoTypes) []JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetInfoTypes) []JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRules(des, initial *JobTriggerInspectJobInspectConfigRuleSetRules, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.HotwordRule != nil || (initial != nil && initial.HotwordRule != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExclusionRule) {
			des.HotwordRule = nil
			if initial != nil {
				initial.HotwordRule = nil
			}
		}
	}

	if des.ExclusionRule != nil || (initial != nil && initial.ExclusionRule != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.HotwordRule) {
			des.ExclusionRule = nil
			if initial != nil {
				initial.ExclusionRule = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRules{}

	cDes.HotwordRule = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(des.HotwordRule, initial.HotwordRule, opts...)
	cDes.ExclusionRule = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(des.ExclusionRule, initial.ExclusionRule, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRules, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRules {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRules(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRules) *JobTriggerInspectJobInspectConfigRuleSetRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.HotwordRule = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, des.HotwordRule, nw.HotwordRule)
	nw.ExclusionRule = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, des.ExclusionRule, nw.ExclusionRule)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRules) []JobTriggerInspectJobInspectConfigRuleSetRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRules) []JobTriggerInspectJobInspectConfigRuleSetRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRules(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}

	cDes.HotwordRegex = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(des.HotwordRegex, initial.HotwordRegex, opts...)
	cDes.Proximity = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(des.Proximity, initial.Proximity, opts...)
	cDes.LikelihoodAdjustment = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(des.LikelihoodAdjustment, initial.LikelihoodAdjustment, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.HotwordRegex = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, des.HotwordRegex, nw.HotwordRegex)
	nw.Proximity = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, des.Proximity, nw.Proximity)
	nw.LikelihoodAdjustment = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, des.LikelihoodAdjustment, nw.LikelihoodAdjustment)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}

	if dcl.StringCanonicalize(des.Pattern, initial.Pattern) || dcl.IsZeroValue(des.Pattern) {
		cDes.Pattern = initial.Pattern
	} else {
		cDes.Pattern = des.Pattern
	}
	if dcl.IsZeroValue(des.GroupIndexes) || (dcl.IsEmptyValueIndirect(des.GroupIndexes) && dcl.IsEmptyValueIndirect(initial.GroupIndexes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GroupIndexes = initial.GroupIndexes
	} else {
		cDes.GroupIndexes = des.GroupIndexes
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Pattern, nw.Pattern) {
		nw.Pattern = des.Pattern
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}

	if dcl.IsZeroValue(des.WindowBefore) || (dcl.IsEmptyValueIndirect(des.WindowBefore) && dcl.IsEmptyValueIndirect(initial.WindowBefore)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.WindowBefore = initial.WindowBefore
	} else {
		cDes.WindowBefore = des.WindowBefore
	}
	if dcl.IsZeroValue(des.WindowAfter) || (dcl.IsEmptyValueIndirect(des.WindowAfter) && dcl.IsEmptyValueIndirect(initial.WindowAfter)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.WindowAfter = initial.WindowAfter
	} else {
		cDes.WindowAfter = des.WindowAfter
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximitySlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximitySet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximitySlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.FixedLikelihood != nil || (initial != nil && initial.FixedLikelihood != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RelativeLikelihood) {
			des.FixedLikelihood = nil
			if initial != nil {
				initial.FixedLikelihood = nil
			}
		}
	}

	if des.RelativeLikelihood != nil || (initial != nil && initial.RelativeLikelihood != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.FixedLikelihood) {
			des.RelativeLikelihood = nil
			if initial != nil {
				initial.RelativeLikelihood = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}

	if dcl.IsZeroValue(des.FixedLikelihood) || (dcl.IsEmptyValueIndirect(des.FixedLikelihood) && dcl.IsEmptyValueIndirect(initial.FixedLikelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.FixedLikelihood = initial.FixedLikelihood
	} else {
		cDes.FixedLikelihood = des.FixedLikelihood
	}
	if dcl.IsZeroValue(des.RelativeLikelihood) || (dcl.IsEmptyValueIndirect(des.RelativeLikelihood) && dcl.IsEmptyValueIndirect(initial.RelativeLikelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RelativeLikelihood = initial.RelativeLikelihood
	} else {
		cDes.RelativeLikelihood = des.RelativeLikelihood
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Dictionary != nil || (initial != nil && initial.Dictionary != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Regex, des.ExcludeInfoTypes) {
			des.Dictionary = nil
			if initial != nil {
				initial.Dictionary = nil
			}
		}
	}

	if des.Regex != nil || (initial != nil && initial.Regex != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.ExcludeInfoTypes) {
			des.Regex = nil
			if initial != nil {
				initial.Regex = nil
			}
		}
	}

	if des.ExcludeInfoTypes != nil || (initial != nil && initial.ExcludeInfoTypes != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.Regex) {
			des.ExcludeInfoTypes = nil
			if initial != nil {
				initial.ExcludeInfoTypes = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}

	cDes.Dictionary = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(des.Dictionary, initial.Dictionary, opts...)
	cDes.Regex = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(des.Regex, initial.Regex, opts...)
	cDes.ExcludeInfoTypes = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(des.ExcludeInfoTypes, initial.ExcludeInfoTypes, opts...)
	if dcl.IsZeroValue(des.MatchingType) || (dcl.IsEmptyValueIndirect(des.MatchingType) && dcl.IsEmptyValueIndirect(initial.MatchingType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MatchingType = initial.MatchingType
	} else {
		cDes.MatchingType = des.MatchingType
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Dictionary = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, des.Dictionary, nw.Dictionary)
	nw.Regex = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, des.Regex, nw.Regex)
	nw.ExcludeInfoTypes = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, des.ExcludeInfoTypes, nw.ExcludeInfoTypes)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.WordList != nil || (initial != nil && initial.WordList != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudStoragePath) {
			des.WordList = nil
			if initial != nil {
				initial.WordList = nil
			}
		}
	}

	if des.CloudStoragePath != nil || (initial != nil && initial.CloudStoragePath != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.WordList) {
			des.CloudStoragePath = nil
			if initial != nil {
				initial.CloudStoragePath = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}

	cDes.WordList = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(des.WordList, initial.WordList, opts...)
	cDes.CloudStoragePath = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(des.CloudStoragePath, initial.CloudStoragePath, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionarySlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.WordList = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, des.WordList, nw.WordList)
	nw.CloudStoragePath = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, des.CloudStoragePath, nw.CloudStoragePath)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionarySet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionarySlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}

	if dcl.StringArrayCanonicalize(des.Words, initial.Words) {
		cDes.Words = initial.Words
	} else {
		cDes.Words = des.Words
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Words, nw.Words) {
		nw.Words = des.Words
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}

	if dcl.StringCanonicalize(des.Pattern, initial.Pattern) || dcl.IsZeroValue(des.Pattern) {
		cDes.Pattern = initial.Pattern
	} else {
		cDes.Pattern = des.Pattern
	}
	if dcl.IsZeroValue(des.GroupIndexes) || (dcl.IsEmptyValueIndirect(des.GroupIndexes) && dcl.IsEmptyValueIndirect(initial.GroupIndexes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GroupIndexes = initial.GroupIndexes
	} else {
		cDes.GroupIndexes = des.GroupIndexes
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Pattern, nw.Pattern) {
		nw.Pattern = des.Pattern
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}

	cDes.InfoTypes = canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(des.InfoTypes, initial.InfoTypes, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoTypes = canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c, des.InfoTypes, nw.InfoTypes)

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(des, initial *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, opts ...dcl.ApplyOption) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(des, initial []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, opts ...dcl.ApplyOption) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c *Client, des, nw *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSet(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c *Client, des, nw []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActions(des, initial *JobTriggerInspectJobActions, opts ...dcl.ApplyOption) *JobTriggerInspectJobActions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.SaveFindings != nil || (initial != nil && initial.SaveFindings != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.PubSub, des.PublishSummaryToCscc, des.PublishFindingsToCloudDataCatalog, des.JobNotificationEmails, des.PublishToStackdriver) {
			des.SaveFindings = nil
			if initial != nil {
				initial.SaveFindings = nil
			}
		}
	}

	if des.PubSub != nil || (initial != nil && initial.PubSub != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.SaveFindings, des.PublishSummaryToCscc, des.PublishFindingsToCloudDataCatalog, des.JobNotificationEmails, des.PublishToStackdriver) {
			des.PubSub = nil
			if initial != nil {
				initial.PubSub = nil
			}
		}
	}

	if des.PublishSummaryToCscc != nil || (initial != nil && initial.PublishSummaryToCscc != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.SaveFindings, des.PubSub, des.PublishFindingsToCloudDataCatalog, des.JobNotificationEmails, des.PublishToStackdriver) {
			des.PublishSummaryToCscc = nil
			if initial != nil {
				initial.PublishSummaryToCscc = nil
			}
		}
	}

	if des.PublishFindingsToCloudDataCatalog != nil || (initial != nil && initial.PublishFindingsToCloudDataCatalog != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.SaveFindings, des.PubSub, des.PublishSummaryToCscc, des.JobNotificationEmails, des.PublishToStackdriver) {
			des.PublishFindingsToCloudDataCatalog = nil
			if initial != nil {
				initial.PublishFindingsToCloudDataCatalog = nil
			}
		}
	}

	if des.JobNotificationEmails != nil || (initial != nil && initial.JobNotificationEmails != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.SaveFindings, des.PubSub, des.PublishSummaryToCscc, des.PublishFindingsToCloudDataCatalog, des.PublishToStackdriver) {
			des.JobNotificationEmails = nil
			if initial != nil {
				initial.JobNotificationEmails = nil
			}
		}
	}

	if des.PublishToStackdriver != nil || (initial != nil && initial.PublishToStackdriver != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.SaveFindings, des.PubSub, des.PublishSummaryToCscc, des.PublishFindingsToCloudDataCatalog, des.JobNotificationEmails) {
			des.PublishToStackdriver = nil
			if initial != nil {
				initial.PublishToStackdriver = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActions{}

	cDes.SaveFindings = canonicalizeJobTriggerInspectJobActionsSaveFindings(des.SaveFindings, initial.SaveFindings, opts...)
	cDes.PubSub = canonicalizeJobTriggerInspectJobActionsPubSub(des.PubSub, initial.PubSub, opts...)
	cDes.PublishSummaryToCscc = canonicalizeJobTriggerInspectJobActionsPublishSummaryToCscc(des.PublishSummaryToCscc, initial.PublishSummaryToCscc, opts...)
	cDes.PublishFindingsToCloudDataCatalog = canonicalizeJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(des.PublishFindingsToCloudDataCatalog, initial.PublishFindingsToCloudDataCatalog, opts...)
	cDes.JobNotificationEmails = canonicalizeJobTriggerInspectJobActionsJobNotificationEmails(des.JobNotificationEmails, initial.JobNotificationEmails, opts...)
	cDes.PublishToStackdriver = canonicalizeJobTriggerInspectJobActionsPublishToStackdriver(des.PublishToStackdriver, initial.PublishToStackdriver, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsSlice(des, initial []JobTriggerInspectJobActions, opts ...dcl.ApplyOption) []JobTriggerInspectJobActions {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActions(c *Client, des, nw *JobTriggerInspectJobActions) *JobTriggerInspectJobActions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.SaveFindings = canonicalizeNewJobTriggerInspectJobActionsSaveFindings(c, des.SaveFindings, nw.SaveFindings)
	nw.PubSub = canonicalizeNewJobTriggerInspectJobActionsPubSub(c, des.PubSub, nw.PubSub)
	nw.PublishSummaryToCscc = canonicalizeNewJobTriggerInspectJobActionsPublishSummaryToCscc(c, des.PublishSummaryToCscc, nw.PublishSummaryToCscc)
	nw.PublishFindingsToCloudDataCatalog = canonicalizeNewJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, des.PublishFindingsToCloudDataCatalog, nw.PublishFindingsToCloudDataCatalog)
	nw.JobNotificationEmails = canonicalizeNewJobTriggerInspectJobActionsJobNotificationEmails(c, des.JobNotificationEmails, nw.JobNotificationEmails)
	nw.PublishToStackdriver = canonicalizeNewJobTriggerInspectJobActionsPublishToStackdriver(c, des.PublishToStackdriver, nw.PublishToStackdriver)

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsSet(c *Client, des, nw []JobTriggerInspectJobActions) []JobTriggerInspectJobActions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsSlice(c *Client, des, nw []JobTriggerInspectJobActions) []JobTriggerInspectJobActions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActions(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsSaveFindings(des, initial *JobTriggerInspectJobActionsSaveFindings, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsSaveFindings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsSaveFindings{}

	cDes.OutputConfig = canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfig(des.OutputConfig, initial.OutputConfig, opts...)

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsSaveFindingsSlice(des, initial []JobTriggerInspectJobActionsSaveFindings, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsSaveFindings {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsSaveFindings, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsSaveFindings(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsSaveFindings, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsSaveFindings(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindings(c *Client, des, nw *JobTriggerInspectJobActionsSaveFindings) *JobTriggerInspectJobActionsSaveFindings {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsSaveFindings while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.OutputConfig = canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, des.OutputConfig, nw.OutputConfig)

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsSet(c *Client, des, nw []JobTriggerInspectJobActionsSaveFindings) []JobTriggerInspectJobActionsSaveFindings {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsSaveFindings
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsSaveFindingsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsSaveFindings(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsSlice(c *Client, des, nw []JobTriggerInspectJobActionsSaveFindings) []JobTriggerInspectJobActionsSaveFindings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsSaveFindings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsSaveFindings(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfig(des, initial *JobTriggerInspectJobActionsSaveFindingsOutputConfig, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Table != nil || (initial != nil && initial.Table != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.DlpStorage) {
			des.Table = nil
			if initial != nil {
				initial.Table = nil
			}
		}
	}

	if des.DlpStorage != nil || (initial != nil && initial.DlpStorage != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Table) {
			des.DlpStorage = nil
			if initial != nil {
				initial.DlpStorage = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsSaveFindingsOutputConfig{}

	cDes.Table = canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(des.Table, initial.Table, opts...)
	cDes.DlpStorage = canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(des.DlpStorage, initial.DlpStorage, opts...)
	if dcl.IsZeroValue(des.OutputSchema) || (dcl.IsEmptyValueIndirect(des.OutputSchema) && dcl.IsEmptyValueIndirect(initial.OutputSchema)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.OutputSchema = initial.OutputSchema
	} else {
		cDes.OutputSchema = des.OutputSchema
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigSlice(des, initial []JobTriggerInspectJobActionsSaveFindingsOutputConfig, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfig(c *Client, des, nw *JobTriggerInspectJobActionsSaveFindingsOutputConfig) *JobTriggerInspectJobActionsSaveFindingsOutputConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsSaveFindingsOutputConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Table = canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, des.Table, nw.Table)
	nw.DlpStorage = canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, des.DlpStorage, nw.DlpStorage)

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigSet(c *Client, des, nw []JobTriggerInspectJobActionsSaveFindingsOutputConfig) []JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsSaveFindingsOutputConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsSaveFindingsOutputConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigSlice(c *Client, des, nw []JobTriggerInspectJobActionsSaveFindingsOutputConfig) []JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsSaveFindingsOutputConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(des, initial *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}

	if dcl.IsZeroValue(des.ProjectId) || (dcl.IsEmptyValueIndirect(des.ProjectId) && dcl.IsEmptyValueIndirect(initial.ProjectId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ProjectId = initial.ProjectId
	} else {
		cDes.ProjectId = des.ProjectId
	}
	if dcl.IsZeroValue(des.DatasetId) || (dcl.IsEmptyValueIndirect(des.DatasetId) && dcl.IsEmptyValueIndirect(initial.DatasetId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DatasetId = initial.DatasetId
	} else {
		cDes.DatasetId = des.DatasetId
	}
	if dcl.IsZeroValue(des.TableId) || (dcl.IsEmptyValueIndirect(des.TableId) && dcl.IsEmptyValueIndirect(initial.TableId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TableId = initial.TableId
	} else {
		cDes.TableId = des.TableId
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigTableSlice(des, initial []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfigTable, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfigTable, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c *Client, des, nw *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsSaveFindingsOutputConfigTable while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigTableSet(c *Client, des, nw []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsSaveFindingsOutputConfigTableNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigTableSlice(c *Client, des, nw []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(des, initial *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageSlice(des, initial []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c *Client, des, nw *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageSet(c *Client, des, nw []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageSlice(c *Client, des, nw []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsPubSub(des, initial *JobTriggerInspectJobActionsPubSub, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsPubSub {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsPubSub{}

	if dcl.IsZeroValue(des.Topic) || (dcl.IsEmptyValueIndirect(des.Topic) && dcl.IsEmptyValueIndirect(initial.Topic)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Topic = initial.Topic
	} else {
		cDes.Topic = des.Topic
	}

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsPubSubSlice(des, initial []JobTriggerInspectJobActionsPubSub, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsPubSub {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsPubSub, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsPubSub(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsPubSub, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsPubSub(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsPubSub(c *Client, des, nw *JobTriggerInspectJobActionsPubSub) *JobTriggerInspectJobActionsPubSub {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsPubSub while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsPubSubSet(c *Client, des, nw []JobTriggerInspectJobActionsPubSub) []JobTriggerInspectJobActionsPubSub {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsPubSub
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsPubSubNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsPubSub(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsPubSubSlice(c *Client, des, nw []JobTriggerInspectJobActionsPubSub) []JobTriggerInspectJobActionsPubSub {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsPubSub
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsPubSub(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsPublishSummaryToCscc(des, initial *JobTriggerInspectJobActionsPublishSummaryToCscc, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsPublishSummaryToCscc {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsPublishSummaryToCscc{}

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsPublishSummaryToCsccSlice(des, initial []JobTriggerInspectJobActionsPublishSummaryToCscc, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsPublishSummaryToCscc {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsPublishSummaryToCscc, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsPublishSummaryToCscc(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsPublishSummaryToCscc, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsPublishSummaryToCscc(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsPublishSummaryToCscc(c *Client, des, nw *JobTriggerInspectJobActionsPublishSummaryToCscc) *JobTriggerInspectJobActionsPublishSummaryToCscc {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsPublishSummaryToCscc while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsPublishSummaryToCsccSet(c *Client, des, nw []JobTriggerInspectJobActionsPublishSummaryToCscc) []JobTriggerInspectJobActionsPublishSummaryToCscc {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsPublishSummaryToCscc
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsPublishSummaryToCsccNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsPublishSummaryToCscc(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsPublishSummaryToCsccSlice(c *Client, des, nw []JobTriggerInspectJobActionsPublishSummaryToCscc) []JobTriggerInspectJobActionsPublishSummaryToCscc {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsPublishSummaryToCscc
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsPublishSummaryToCscc(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(des, initial *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogSlice(des, initial []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c *Client, des, nw *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogSet(c *Client, des, nw []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogSlice(c *Client, des, nw []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsJobNotificationEmails(des, initial *JobTriggerInspectJobActionsJobNotificationEmails, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsJobNotificationEmails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsJobNotificationEmails{}

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsJobNotificationEmailsSlice(des, initial []JobTriggerInspectJobActionsJobNotificationEmails, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsJobNotificationEmails {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsJobNotificationEmails, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsJobNotificationEmails(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsJobNotificationEmails, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsJobNotificationEmails(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsJobNotificationEmails(c *Client, des, nw *JobTriggerInspectJobActionsJobNotificationEmails) *JobTriggerInspectJobActionsJobNotificationEmails {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsJobNotificationEmails while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsJobNotificationEmailsSet(c *Client, des, nw []JobTriggerInspectJobActionsJobNotificationEmails) []JobTriggerInspectJobActionsJobNotificationEmails {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsJobNotificationEmails
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsJobNotificationEmailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsJobNotificationEmails(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsJobNotificationEmailsSlice(c *Client, des, nw []JobTriggerInspectJobActionsJobNotificationEmails) []JobTriggerInspectJobActionsJobNotificationEmails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsJobNotificationEmails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsJobNotificationEmails(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerInspectJobActionsPublishToStackdriver(des, initial *JobTriggerInspectJobActionsPublishToStackdriver, opts ...dcl.ApplyOption) *JobTriggerInspectJobActionsPublishToStackdriver {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &JobTriggerInspectJobActionsPublishToStackdriver{}

	return cDes
}

func canonicalizeJobTriggerInspectJobActionsPublishToStackdriverSlice(des, initial []JobTriggerInspectJobActionsPublishToStackdriver, opts ...dcl.ApplyOption) []JobTriggerInspectJobActionsPublishToStackdriver {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerInspectJobActionsPublishToStackdriver, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerInspectJobActionsPublishToStackdriver(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerInspectJobActionsPublishToStackdriver, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerInspectJobActionsPublishToStackdriver(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerInspectJobActionsPublishToStackdriver(c *Client, des, nw *JobTriggerInspectJobActionsPublishToStackdriver) *JobTriggerInspectJobActionsPublishToStackdriver {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerInspectJobActionsPublishToStackdriver while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerInspectJobActionsPublishToStackdriverSet(c *Client, des, nw []JobTriggerInspectJobActionsPublishToStackdriver) []JobTriggerInspectJobActionsPublishToStackdriver {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerInspectJobActionsPublishToStackdriver
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerInspectJobActionsPublishToStackdriverNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerInspectJobActionsPublishToStackdriver(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerInspectJobActionsPublishToStackdriverSlice(c *Client, des, nw []JobTriggerInspectJobActionsPublishToStackdriver) []JobTriggerInspectJobActionsPublishToStackdriver {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerInspectJobActionsPublishToStackdriver
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerInspectJobActionsPublishToStackdriver(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerTriggers(des, initial *JobTriggerTriggers, opts ...dcl.ApplyOption) *JobTriggerTriggers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Schedule != nil || (initial != nil && initial.Schedule != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Manual) {
			des.Schedule = nil
			if initial != nil {
				initial.Schedule = nil
			}
		}
	}

	if des.Manual != nil || (initial != nil && initial.Manual != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Schedule) {
			des.Manual = nil
			if initial != nil {
				initial.Manual = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerTriggers{}

	cDes.Schedule = canonicalizeJobTriggerTriggersSchedule(des.Schedule, initial.Schedule, opts...)
	cDes.Manual = canonicalizeJobTriggerTriggersManual(des.Manual, initial.Manual, opts...)

	return cDes
}

func canonicalizeJobTriggerTriggersSlice(des, initial []JobTriggerTriggers, opts ...dcl.ApplyOption) []JobTriggerTriggers {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerTriggers, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerTriggers(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerTriggers, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerTriggers(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerTriggers(c *Client, des, nw *JobTriggerTriggers) *JobTriggerTriggers {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerTriggers while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Schedule = canonicalizeNewJobTriggerTriggersSchedule(c, des.Schedule, nw.Schedule)
	nw.Manual = canonicalizeNewJobTriggerTriggersManual(c, des.Manual, nw.Manual)

	return nw
}

func canonicalizeNewJobTriggerTriggersSet(c *Client, des, nw []JobTriggerTriggers) []JobTriggerTriggers {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerTriggers
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerTriggersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerTriggers(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerTriggersSlice(c *Client, des, nw []JobTriggerTriggers) []JobTriggerTriggers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerTriggers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerTriggers(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerTriggersSchedule(des, initial *JobTriggerTriggersSchedule, opts ...dcl.ApplyOption) *JobTriggerTriggersSchedule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerTriggersSchedule{}

	if dcl.StringCanonicalize(des.RecurrencePeriodDuration, initial.RecurrencePeriodDuration) || dcl.IsZeroValue(des.RecurrencePeriodDuration) {
		cDes.RecurrencePeriodDuration = initial.RecurrencePeriodDuration
	} else {
		cDes.RecurrencePeriodDuration = des.RecurrencePeriodDuration
	}

	return cDes
}

func canonicalizeJobTriggerTriggersScheduleSlice(des, initial []JobTriggerTriggersSchedule, opts ...dcl.ApplyOption) []JobTriggerTriggersSchedule {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerTriggersSchedule, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerTriggersSchedule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerTriggersSchedule, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerTriggersSchedule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerTriggersSchedule(c *Client, des, nw *JobTriggerTriggersSchedule) *JobTriggerTriggersSchedule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerTriggersSchedule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.RecurrencePeriodDuration, nw.RecurrencePeriodDuration) {
		nw.RecurrencePeriodDuration = des.RecurrencePeriodDuration
	}

	return nw
}

func canonicalizeNewJobTriggerTriggersScheduleSet(c *Client, des, nw []JobTriggerTriggersSchedule) []JobTriggerTriggersSchedule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerTriggersSchedule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerTriggersScheduleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerTriggersSchedule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerTriggersScheduleSlice(c *Client, des, nw []JobTriggerTriggersSchedule) []JobTriggerTriggersSchedule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerTriggersSchedule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerTriggersSchedule(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerTriggersManual(des, initial *JobTriggerTriggersManual, opts ...dcl.ApplyOption) *JobTriggerTriggersManual {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &JobTriggerTriggersManual{}

	return cDes
}

func canonicalizeJobTriggerTriggersManualSlice(des, initial []JobTriggerTriggersManual, opts ...dcl.ApplyOption) []JobTriggerTriggersManual {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerTriggersManual, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerTriggersManual(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerTriggersManual, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerTriggersManual(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerTriggersManual(c *Client, des, nw *JobTriggerTriggersManual) *JobTriggerTriggersManual {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerTriggersManual while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewJobTriggerTriggersManualSet(c *Client, des, nw []JobTriggerTriggersManual) []JobTriggerTriggersManual {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerTriggersManual
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerTriggersManualNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerTriggersManual(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerTriggersManualSlice(c *Client, des, nw []JobTriggerTriggersManual) []JobTriggerTriggersManual {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerTriggersManual
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerTriggersManual(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerErrors(des, initial *JobTriggerErrors, opts ...dcl.ApplyOption) *JobTriggerErrors {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerErrors{}

	cDes.Details = canonicalizeJobTriggerErrorsDetails(des.Details, initial.Details, opts...)
	if dcl.IsZeroValue(des.Timestamps) || (dcl.IsEmptyValueIndirect(des.Timestamps) && dcl.IsEmptyValueIndirect(initial.Timestamps)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Timestamps = initial.Timestamps
	} else {
		cDes.Timestamps = des.Timestamps
	}

	return cDes
}

func canonicalizeJobTriggerErrorsSlice(des, initial []JobTriggerErrors, opts ...dcl.ApplyOption) []JobTriggerErrors {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerErrors, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerErrors(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerErrors, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerErrors(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerErrors(c *Client, des, nw *JobTriggerErrors) *JobTriggerErrors {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerErrors while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Details = canonicalizeNewJobTriggerErrorsDetails(c, des.Details, nw.Details)

	return nw
}

func canonicalizeNewJobTriggerErrorsSet(c *Client, des, nw []JobTriggerErrors) []JobTriggerErrors {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerErrors
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerErrorsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerErrors(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerErrorsSlice(c *Client, des, nw []JobTriggerErrors) []JobTriggerErrors {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerErrors
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerErrors(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerErrorsDetails(des, initial *JobTriggerErrorsDetails, opts ...dcl.ApplyOption) *JobTriggerErrorsDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerErrorsDetails{}

	if dcl.IsZeroValue(des.Code) || (dcl.IsEmptyValueIndirect(des.Code) && dcl.IsEmptyValueIndirect(initial.Code)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Code = initial.Code
	} else {
		cDes.Code = des.Code
	}
	if dcl.StringCanonicalize(des.Message, initial.Message) || dcl.IsZeroValue(des.Message) {
		cDes.Message = initial.Message
	} else {
		cDes.Message = des.Message
	}
	cDes.Details = canonicalizeJobTriggerErrorsDetailsDetailsSlice(des.Details, initial.Details, opts...)

	return cDes
}

func canonicalizeJobTriggerErrorsDetailsSlice(des, initial []JobTriggerErrorsDetails, opts ...dcl.ApplyOption) []JobTriggerErrorsDetails {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerErrorsDetails, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerErrorsDetails(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerErrorsDetails, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerErrorsDetails(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerErrorsDetails(c *Client, des, nw *JobTriggerErrorsDetails) *JobTriggerErrorsDetails {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerErrorsDetails while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Message, nw.Message) {
		nw.Message = des.Message
	}
	nw.Details = canonicalizeNewJobTriggerErrorsDetailsDetailsSlice(c, des.Details, nw.Details)

	return nw
}

func canonicalizeNewJobTriggerErrorsDetailsSet(c *Client, des, nw []JobTriggerErrorsDetails) []JobTriggerErrorsDetails {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerErrorsDetails
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerErrorsDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerErrorsDetails(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerErrorsDetailsSlice(c *Client, des, nw []JobTriggerErrorsDetails) []JobTriggerErrorsDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerErrorsDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerErrorsDetails(c, &d, &n))
	}

	return items
}

func canonicalizeJobTriggerErrorsDetailsDetails(des, initial *JobTriggerErrorsDetailsDetails, opts ...dcl.ApplyOption) *JobTriggerErrorsDetailsDetails {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &JobTriggerErrorsDetailsDetails{}

	if dcl.StringCanonicalize(des.TypeUrl, initial.TypeUrl) || dcl.IsZeroValue(des.TypeUrl) {
		cDes.TypeUrl = initial.TypeUrl
	} else {
		cDes.TypeUrl = des.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		cDes.Value = initial.Value
	} else {
		cDes.Value = des.Value
	}

	return cDes
}

func canonicalizeJobTriggerErrorsDetailsDetailsSlice(des, initial []JobTriggerErrorsDetailsDetails, opts ...dcl.ApplyOption) []JobTriggerErrorsDetailsDetails {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]JobTriggerErrorsDetailsDetails, 0, len(des))
		for _, d := range des {
			cd := canonicalizeJobTriggerErrorsDetailsDetails(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]JobTriggerErrorsDetailsDetails, 0, len(des))
	for i, d := range des {
		cd := canonicalizeJobTriggerErrorsDetailsDetails(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewJobTriggerErrorsDetailsDetails(c *Client, des, nw *JobTriggerErrorsDetailsDetails) *JobTriggerErrorsDetailsDetails {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for JobTriggerErrorsDetailsDetails while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.TypeUrl, nw.TypeUrl) {
		nw.TypeUrl = des.TypeUrl
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewJobTriggerErrorsDetailsDetailsSet(c *Client, des, nw []JobTriggerErrorsDetailsDetails) []JobTriggerErrorsDetailsDetails {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []JobTriggerErrorsDetailsDetails
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareJobTriggerErrorsDetailsDetailsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewJobTriggerErrorsDetailsDetails(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewJobTriggerErrorsDetailsDetailsSlice(c *Client, des, nw []JobTriggerErrorsDetailsDetails) []JobTriggerErrorsDetailsDetails {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []JobTriggerErrorsDetailsDetails
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewJobTriggerErrorsDetailsDetails(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffJobTrigger(c *Client, desired, actual *JobTrigger, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InspectJob, actual.InspectJob, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobNewStyle, EmptyObject: EmptyJobTriggerInspectJob, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("InspectJob")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Triggers, actual.Triggers, dcl.DiffInfo{ObjectFunction: compareJobTriggerTriggersNewStyle, EmptyObject: EmptyJobTriggerTriggers, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Triggers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Errors, actual.Errors, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareJobTriggerErrorsNewStyle, EmptyObject: EmptyJobTriggerErrors, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Errors")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastRunTime, actual.LastRunTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastRunTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocationId, actual.LocationId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocationId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parent, actual.Parent, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Parent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}
func compareJobTriggerInspectJobNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJob)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJob)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJob or *JobTriggerInspectJob", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJob)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJob)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJob", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StorageConfig, actual.StorageConfig, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfig, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("StorageConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InspectConfig, actual.InspectConfig, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfig, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("InspectConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InspectTemplateName, actual.InspectTemplateName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("InspectTemplateName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Actions, actual.Actions, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsNewStyle, EmptyObject: EmptyJobTriggerInspectJobActions, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Actions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfig)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfig or *JobTriggerInspectJobStorageConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfig)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DatastoreOptions, actual.DatastoreOptions, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigDatastoreOptionsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigDatastoreOptions, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("DatastoreOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudStorageOptions, actual.CloudStorageOptions, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigCloudStorageOptionsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigCloudStorageOptions, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("CloudStorageOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BigQueryOptions, actual.BigQueryOptions, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigBigQueryOptionsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigBigQueryOptions, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("BigQueryOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HybridOptions, actual.HybridOptions, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigHybridOptionsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigHybridOptions, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("HybridOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimespanConfig, actual.TimespanConfig, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigTimespanConfigNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigTimespanConfig, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("TimespanConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigDatastoreOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigDatastoreOptions)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigDatastoreOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigDatastoreOptions or *JobTriggerInspectJobStorageConfigDatastoreOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigDatastoreOptions)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigDatastoreOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigDatastoreOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PartitionId, actual.PartitionId, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("PartitionId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Kind, actual.Kind, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigDatastoreOptionsKindNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsKind, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Kind")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId or *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NamespaceId, actual.NamespaceId, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("NamespaceId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigDatastoreOptionsKindNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigDatastoreOptionsKind)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigDatastoreOptionsKind)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigDatastoreOptionsKind or *JobTriggerInspectJobStorageConfigDatastoreOptionsKind", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigDatastoreOptionsKind)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigDatastoreOptionsKind)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigDatastoreOptionsKind", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigCloudStorageOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigCloudStorageOptions)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigCloudStorageOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigCloudStorageOptions or *JobTriggerInspectJobStorageConfigCloudStorageOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigCloudStorageOptions)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigCloudStorageOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigCloudStorageOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FileSet, actual.FileSet, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("FileSet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BytesLimitPerFile, actual.BytesLimitPerFile, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("BytesLimitPerFile")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BytesLimitPerFilePercent, actual.BytesLimitPerFilePercent, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("BytesLimitPerFilePercent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FileTypes, actual.FileTypes, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("FileTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SampleMethod, actual.SampleMethod, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("SampleMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FilesLimitPercent, actual.FilesLimitPercent, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("FilesLimitPercent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet or *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RegexFileSet, actual.RegexFileSet, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("RegexFileSet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet or *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BucketName, actual.BucketName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("BucketName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludeRegex, actual.IncludeRegex, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("IncludeRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExcludeRegex, actual.ExcludeRegex, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ExcludeRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigBigQueryOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigBigQueryOptions)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigBigQueryOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptions or *JobTriggerInspectJobStorageConfigBigQueryOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigBigQueryOptions)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigBigQueryOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TableReference, actual.TableReference, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("TableReference")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IdentifyingFields, actual.IdentifyingFields, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("IdentifyingFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RowsLimit, actual.RowsLimit, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("RowsLimit")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RowsLimitPercent, actual.RowsLimitPercent, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("RowsLimitPercent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SampleMethod, actual.SampleMethod, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("SampleMethod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExcludedFields, actual.ExcludedFields, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ExcludedFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludedFields, actual.IncludedFields, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("IncludedFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference or *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatasetId, actual.DatasetId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("DatasetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TableId, actual.TableId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("TableId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields or *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields or *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields or *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigHybridOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigHybridOptions)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigHybridOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigHybridOptions or *JobTriggerInspectJobStorageConfigHybridOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigHybridOptions)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigHybridOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigHybridOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequiredFindingLabelKeys, actual.RequiredFindingLabelKeys, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("RequiredFindingLabelKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TableOptions, actual.TableOptions, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptions, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("TableOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigHybridOptionsTableOptions)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigHybridOptionsTableOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigHybridOptionsTableOptions or *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigHybridOptionsTableOptions)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigHybridOptionsTableOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigHybridOptionsTableOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IdentifyingFields, actual.IdentifyingFields, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("IdentifyingFields")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields or *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigTimespanConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigTimespanConfig)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigTimespanConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigTimespanConfig or *JobTriggerInspectJobStorageConfigTimespanConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigTimespanConfig)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigTimespanConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigTimespanConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndTime, actual.EndTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("EndTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimestampField, actual.TimestampField, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldNewStyle, EmptyObject: EmptyJobTriggerInspectJobStorageConfigTimespanConfigTimestampField, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("TimestampField")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableAutoPopulationOfTimespanConfig, actual.EnableAutoPopulationOfTimespanConfig, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("EnableAutoPopulationOfTimespanConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobStorageConfigTimespanConfigTimestampField)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobStorageConfigTimespanConfigTimestampField)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigTimespanConfigTimestampField or *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobStorageConfigTimespanConfigTimestampField)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobStorageConfigTimespanConfigTimestampField)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobStorageConfigTimespanConfigTimestampField", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfig)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfig or *JobTriggerInspectJobInspectConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfig)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoTypes, actual.InfoTypes, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigInfoTypesNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigInfoTypes, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("InfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinLikelihood, actual.MinLikelihood, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("MinLikelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Limits, actual.Limits, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigLimitsNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigLimits, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Limits")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludeQuote, actual.IncludeQuote, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("IncludeQuote")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExcludeInfoTypes, actual.ExcludeInfoTypes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ExcludeInfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomInfoTypes, actual.CustomInfoTypes, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypes, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("CustomInfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RuleSet, actual.RuleSet, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSet, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("RuleSet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigInfoTypes or *JobTriggerInspectJobInspectConfigInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigLimitsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigLimits)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigLimits)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigLimits or *JobTriggerInspectJobInspectConfigLimits", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigLimits)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigLimits)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigLimits", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxFindingsPerItem, actual.MaxFindingsPerItem, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("MaxFindingsPerItem")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxFindingsPerRequest, actual.MaxFindingsPerRequest, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("MaxFindingsPerRequest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxFindingsPerInfoType, actual.MaxFindingsPerInfoType, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("MaxFindingsPerInfoType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType or *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoType, actual.InfoType, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("InfoType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxFindings, actual.MaxFindings, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("MaxFindings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType or *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypes or *JobTriggerInspectJobInspectConfigCustomInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoType, actual.InfoType, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("InfoType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Likelihood, actual.Likelihood, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Likelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dictionary, actual.Dictionary, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Dictionary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Regex, actual.Regex, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesRegexNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesRegex, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Regex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SurrogateType, actual.SurrogateType, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("SurrogateType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StoredType, actual.StoredType, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("StoredType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DetectionRules, actual.DetectionRules, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("DetectionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExclusionType, actual.ExclusionType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ExclusionType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType or *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary or *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WordList, actual.WordList, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("WordList")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudStoragePath, actual.CloudStoragePath, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("CloudStoragePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList or *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Words, actual.Words, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Words")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath or *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesRegexNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesRegex)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesRegex or *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesRegex)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesRegex", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pattern, actual.Pattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Pattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupIndexes, actual.GroupIndexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("GroupIndexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType or *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules or *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HotwordRule, actual.HotwordRule, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("HotwordRule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule or *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HotwordRegex, actual.HotwordRegex, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("HotwordRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Proximity, actual.Proximity, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Proximity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LikelihoodAdjustment, actual.LikelihoodAdjustment, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("LikelihoodAdjustment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex or *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pattern, actual.Pattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Pattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupIndexes, actual.GroupIndexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("GroupIndexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity or *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WindowBefore, actual.WindowBefore, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("WindowBefore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WindowAfter, actual.WindowAfter, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("WindowAfter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment or *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FixedLikelihood, actual.FixedLikelihood, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("FixedLikelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RelativeLikelihood, actual.RelativeLikelihood, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("RelativeLikelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSet)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSet or *JobTriggerInspectJobInspectConfigRuleSet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSet)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoTypes, actual.InfoTypes, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetInfoTypesNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetInfoTypes, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("InfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rules, actual.Rules, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRules, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Rules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetInfoTypes or *JobTriggerInspectJobInspectConfigRuleSetInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRules)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRules or *JobTriggerInspectJobInspectConfigRuleSetRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRules)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HotwordRule, actual.HotwordRule, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("HotwordRule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExclusionRule, actual.ExclusionRule, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ExclusionRule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule or *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HotwordRegex, actual.HotwordRegex, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("HotwordRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Proximity, actual.Proximity, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Proximity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LikelihoodAdjustment, actual.LikelihoodAdjustment, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("LikelihoodAdjustment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex or *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pattern, actual.Pattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Pattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupIndexes, actual.GroupIndexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("GroupIndexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity or *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WindowBefore, actual.WindowBefore, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("WindowBefore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WindowAfter, actual.WindowAfter, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("WindowAfter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment or *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FixedLikelihood, actual.FixedLikelihood, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("FixedLikelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RelativeLikelihood, actual.RelativeLikelihood, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("RelativeLikelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule or *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Dictionary, actual.Dictionary, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Dictionary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Regex, actual.Regex, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Regex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExcludeInfoTypes, actual.ExcludeInfoTypes, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ExcludeInfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MatchingType, actual.MatchingType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("MatchingType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary or *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WordList, actual.WordList, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("WordList")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudStoragePath, actual.CloudStoragePath, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("CloudStoragePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList or *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Words, actual.Words, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Words")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath or *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex or *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pattern, actual.Pattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Pattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupIndexes, actual.GroupIndexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("GroupIndexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes or *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoTypes, actual.InfoTypes, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesNewStyle, EmptyObject: EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("InfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes or *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobActionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobActions)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobActions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActions or *JobTriggerInspectJobActions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobActions)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobActions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SaveFindings, actual.SaveFindings, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsSaveFindingsNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsSaveFindings, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("SaveFindings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PubSub, actual.PubSub, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsPubSubNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsPubSub, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("PubSub")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublishSummaryToCscc, actual.PublishSummaryToCscc, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsPublishSummaryToCsccNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsPublishSummaryToCscc, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("PublishSummaryToCscc")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublishFindingsToCloudDataCatalog, actual.PublishFindingsToCloudDataCatalog, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("PublishFindingsToCloudDataCatalog")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.JobNotificationEmails, actual.JobNotificationEmails, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsJobNotificationEmailsNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsJobNotificationEmails, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("JobNotificationEmails")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublishToStackdriver, actual.PublishToStackdriver, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsPublishToStackdriverNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsPublishToStackdriver, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("PublishToStackdriver")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobActionsSaveFindingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobActionsSaveFindings)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobActionsSaveFindings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActionsSaveFindings or *JobTriggerInspectJobActionsSaveFindings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobActionsSaveFindings)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobActionsSaveFindings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActionsSaveFindings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.OutputConfig, actual.OutputConfig, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsSaveFindingsOutputConfigNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfig, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("OutputConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobActionsSaveFindingsOutputConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobActionsSaveFindingsOutputConfig)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobActionsSaveFindingsOutputConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActionsSaveFindingsOutputConfig or *JobTriggerInspectJobActionsSaveFindingsOutputConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobActionsSaveFindingsOutputConfig)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobActionsSaveFindingsOutputConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActionsSaveFindingsOutputConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Table, actual.Table, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsSaveFindingsOutputConfigTableNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigTable, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Table")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DlpStorage, actual.DlpStorage, dcl.DiffInfo{ObjectFunction: compareJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageNewStyle, EmptyObject: EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("DlpStorage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OutputSchema, actual.OutputSchema, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("OutputSchema")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobActionsSaveFindingsOutputConfigTableNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobActionsSaveFindingsOutputConfigTable)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobActionsSaveFindingsOutputConfigTable)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActionsSaveFindingsOutputConfigTable or *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobActionsSaveFindingsOutputConfigTable)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobActionsSaveFindingsOutputConfigTable)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActionsSaveFindingsOutputConfigTable", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatasetId, actual.DatasetId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("DatasetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TableId, actual.TableId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("TableId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareJobTriggerInspectJobActionsPubSubNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerInspectJobActionsPubSub)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerInspectJobActionsPubSub)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActionsPubSub or *JobTriggerInspectJobActionsPubSub", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerInspectJobActionsPubSub)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerInspectJobActionsPubSub)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerInspectJobActionsPubSub", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Topic, actual.Topic, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Topic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerInspectJobActionsPublishSummaryToCsccNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareJobTriggerInspectJobActionsJobNotificationEmailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareJobTriggerInspectJobActionsPublishToStackdriverNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareJobTriggerTriggersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerTriggers)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerTriggers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerTriggers or *JobTriggerTriggers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerTriggers)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerTriggers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerTriggers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Schedule, actual.Schedule, dcl.DiffInfo{ObjectFunction: compareJobTriggerTriggersScheduleNewStyle, EmptyObject: EmptyJobTriggerTriggersSchedule, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Schedule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Manual, actual.Manual, dcl.DiffInfo{ObjectFunction: compareJobTriggerTriggersManualNewStyle, EmptyObject: EmptyJobTriggerTriggersManual, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Manual")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerTriggersScheduleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerTriggersSchedule)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerTriggersSchedule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerTriggersSchedule or *JobTriggerTriggersSchedule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerTriggersSchedule)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerTriggersSchedule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerTriggersSchedule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RecurrencePeriodDuration, actual.RecurrencePeriodDuration, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("RecurrencePeriodDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerTriggersManualNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareJobTriggerErrorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerErrors)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerErrors)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerErrors or *JobTriggerErrors", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerErrors)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerErrors)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerErrors", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Details, actual.Details, dcl.DiffInfo{ObjectFunction: compareJobTriggerErrorsDetailsNewStyle, EmptyObject: EmptyJobTriggerErrorsDetails, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Details")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Timestamps, actual.Timestamps, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Timestamps")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerErrorsDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerErrorsDetails)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerErrorsDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerErrorsDetails or *JobTriggerErrorsDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerErrorsDetails)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerErrorsDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerErrorsDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Message, actual.Message, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Message")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Details, actual.Details, dcl.DiffInfo{ObjectFunction: compareJobTriggerErrorsDetailsDetailsNewStyle, EmptyObject: EmptyJobTriggerErrorsDetailsDetails, OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Details")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareJobTriggerErrorsDetailsDetailsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*JobTriggerErrorsDetailsDetails)
	if !ok {
		desiredNotPointer, ok := d.(JobTriggerErrorsDetailsDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerErrorsDetailsDetails or *JobTriggerErrorsDetailsDetails", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*JobTriggerErrorsDetailsDetails)
	if !ok {
		actualNotPointer, ok := a.(JobTriggerErrorsDetailsDetails)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a JobTriggerErrorsDetailsDetails", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TypeUrl, actual.TypeUrl, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("TypeUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateJobTriggerUpdateJobTriggerOperation")}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *JobTrigger) urlNormalized() *JobTrigger {
	normalized := dcl.Copy(*r).(JobTrigger)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.LocationId = dcl.SelfLinkToName(r.LocationId)
	normalized.Parent = r.Parent
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *JobTrigger) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateJobTrigger" {
		fields := map[string]interface{}{
			"location": dcl.ValueOrEmptyString(nr.Location),
			"parent":   dcl.ValueOrEmptyString(nr.Parent),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		if dcl.IsRegion(nr.Location) {
			return dcl.URL("{{parent}}/locations/{{location}}/jobTriggers/{{name}}", nr.basePath(), userBasePath, fields), nil
		}

		return dcl.URL("{{parent}}/jobTriggers/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the JobTrigger resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *JobTrigger) marshal(c *Client) ([]byte, error) {
	m, err := expandJobTrigger(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JobTrigger: %w", err)
	}
	m = encodeJobTriggerCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalJobTrigger decodes JSON responses into the JobTrigger resource schema.
func unmarshalJobTrigger(b []byte, c *Client, res *JobTrigger) (*JobTrigger, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapJobTrigger(m, c, res)
}

func unmarshalMapJobTrigger(m map[string]interface{}, c *Client, res *JobTrigger) (*JobTrigger, error) {

	flattened := flattenJobTrigger(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandJobTrigger expands JobTrigger into a JSON request object.
func expandJobTrigger(c *Client, f *JobTrigger) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := expandJobTriggerInspectJob(c, f.InspectJob, res); err != nil {
		return nil, fmt.Errorf("error expanding InspectJob into inspectJob: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["inspectJob"] = v
	}
	if v, err := expandJobTriggerTriggersSlice(c, f.Triggers, res); err != nil {
		return nil, fmt.Errorf("error expanding Triggers into triggers: %w", err)
	} else if v != nil {
		m["triggers"] = v
	}
	if v := f.Status; dcl.ValueShouldBeSent(v) {
		m["status"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenJobTrigger flattens JobTrigger from a JSON request object into the
// JobTrigger type.
func flattenJobTrigger(c *Client, i interface{}, res *JobTrigger) *JobTrigger {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &JobTrigger{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.InspectJob = flattenJobTriggerInspectJob(c, m["inspectJob"], res)
	resultRes.Triggers = flattenJobTriggerTriggersSlice(c, m["triggers"], res)
	resultRes.Errors = flattenJobTriggerErrorsSlice(c, m["errors"], res)
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.LastRunTime = dcl.FlattenString(m["lastRunTime"])
	resultRes.Status = flattenJobTriggerStatusEnum(m["status"])
	resultRes.LocationId = dcl.FlattenString(m["locationId"])
	resultRes.Parent = dcl.FlattenString(m["parent"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandJobTriggerInspectJobMap expands the contents of JobTriggerInspectJob into a JSON
// request object.
func expandJobTriggerInspectJobMap(c *Client, f map[string]JobTriggerInspectJob, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJob(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobSlice expands the contents of JobTriggerInspectJob into a JSON
// request object.
func expandJobTriggerInspectJobSlice(c *Client, f []JobTriggerInspectJob, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJob(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobMap flattens the contents of JobTriggerInspectJob from a JSON
// response object.
func flattenJobTriggerInspectJobMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJob {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJob{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJob{}
	}

	items := make(map[string]JobTriggerInspectJob)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJob(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobSlice flattens the contents of JobTriggerInspectJob from a JSON
// response object.
func flattenJobTriggerInspectJobSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJob {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJob{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJob{}
	}

	items := make([]JobTriggerInspectJob, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJob(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJob expands an instance of JobTriggerInspectJob into a JSON
// request object.
func expandJobTriggerInspectJob(c *Client, f *JobTriggerInspectJob, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobStorageConfig(c, f.StorageConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding StorageConfig into storageConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["storageConfig"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfig(c, f.InspectConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding InspectConfig into inspectConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["inspectConfig"] = v
	}
	if v := f.InspectTemplateName; !dcl.IsEmptyValueIndirect(v) {
		m["inspectTemplateName"] = v
	}
	if v, err := expandJobTriggerInspectJobActionsSlice(c, f.Actions, res); err != nil {
		return nil, fmt.Errorf("error expanding Actions into actions: %w", err)
	} else if v != nil {
		m["actions"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJob flattens an instance of JobTriggerInspectJob from a JSON
// response object.
func flattenJobTriggerInspectJob(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJob {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJob{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJob
	}
	r.StorageConfig = flattenJobTriggerInspectJobStorageConfig(c, m["storageConfig"], res)
	r.InspectConfig = flattenJobTriggerInspectJobInspectConfig(c, m["inspectConfig"], res)
	r.InspectTemplateName = dcl.FlattenString(m["inspectTemplateName"])
	r.Actions = flattenJobTriggerInspectJobActionsSlice(c, m["actions"], res)

	return r
}

// expandJobTriggerInspectJobStorageConfigMap expands the contents of JobTriggerInspectJobStorageConfig into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigMap(c *Client, f map[string]JobTriggerInspectJobStorageConfig, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigSlice expands the contents of JobTriggerInspectJobStorageConfig into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigSlice(c *Client, f []JobTriggerInspectJobStorageConfig, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigMap flattens the contents of JobTriggerInspectJobStorageConfig from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfig{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfig{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfig)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigSlice flattens the contents of JobTriggerInspectJobStorageConfig from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfig{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfig{}
	}

	items := make([]JobTriggerInspectJobStorageConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfig expands an instance of JobTriggerInspectJobStorageConfig into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfig(c *Client, f *JobTriggerInspectJobStorageConfig, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobStorageConfigDatastoreOptions(c, f.DatastoreOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding DatastoreOptions into datastoreOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["datastoreOptions"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptions(c, f.CloudStorageOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudStorageOptions into cloudStorageOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudStorageOptions"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigBigQueryOptions(c, f.BigQueryOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding BigQueryOptions into bigQueryOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bigQueryOptions"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigHybridOptions(c, f.HybridOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding HybridOptions into hybridOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hybridOptions"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigTimespanConfig(c, f.TimespanConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding TimespanConfig into timespanConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timespanConfig"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfig flattens an instance of JobTriggerInspectJobStorageConfig from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfig(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfig
	}
	r.DatastoreOptions = flattenJobTriggerInspectJobStorageConfigDatastoreOptions(c, m["datastoreOptions"], res)
	r.CloudStorageOptions = flattenJobTriggerInspectJobStorageConfigCloudStorageOptions(c, m["cloudStorageOptions"], res)
	r.BigQueryOptions = flattenJobTriggerInspectJobStorageConfigBigQueryOptions(c, m["bigQueryOptions"], res)
	r.HybridOptions = flattenJobTriggerInspectJobStorageConfigHybridOptions(c, m["hybridOptions"], res)
	r.TimespanConfig = flattenJobTriggerInspectJobStorageConfigTimespanConfig(c, m["timespanConfig"], res)

	return r
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptionsMap expands the contents of JobTriggerInspectJobStorageConfigDatastoreOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptionsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigDatastoreOptions, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigDatastoreOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptionsSlice expands the contents of JobTriggerInspectJobStorageConfigDatastoreOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptionsSlice(c *Client, f []JobTriggerInspectJobStorageConfigDatastoreOptions, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigDatastoreOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptionsMap flattens the contents of JobTriggerInspectJobStorageConfigDatastoreOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptionsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigDatastoreOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigDatastoreOptions{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigDatastoreOptions{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigDatastoreOptions)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigDatastoreOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptionsSlice flattens the contents of JobTriggerInspectJobStorageConfigDatastoreOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptionsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigDatastoreOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigDatastoreOptions{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigDatastoreOptions{}
	}

	items := make([]JobTriggerInspectJobStorageConfigDatastoreOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigDatastoreOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptions expands an instance of JobTriggerInspectJobStorageConfigDatastoreOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptions(c *Client, f *JobTriggerInspectJobStorageConfigDatastoreOptions, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, f.PartitionId, res); err != nil {
		return nil, fmt.Errorf("error expanding PartitionId into partitionId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["partitionId"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, f.Kind, res); err != nil {
		return nil, fmt.Errorf("error expanding Kind into kind: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kind"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptions flattens an instance of JobTriggerInspectJobStorageConfigDatastoreOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptions(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigDatastoreOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigDatastoreOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigDatastoreOptions
	}
	r.PartitionId = flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, m["partitionId"], res)
	r.Kind = flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, m["kind"], res)

	return r
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdMap expands the contents of JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdSlice expands the contents of JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdSlice(c *Client, f []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdMap flattens the contents of JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdSlice flattens the contents of JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	}

	items := make([]JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId expands an instance of JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c *Client, f *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.NamespaceId; !dcl.IsEmptyValueIndirect(v) {
		m["namespaceId"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId flattens an instance of JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.NamespaceId = dcl.FlattenString(m["namespaceId"])

	return r
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptionsKindMap expands the contents of JobTriggerInspectJobStorageConfigDatastoreOptionsKind into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptionsKindMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsKind, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptionsKindSlice expands the contents of JobTriggerInspectJobStorageConfigDatastoreOptionsKind into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptionsKindSlice(c *Client, f []JobTriggerInspectJobStorageConfigDatastoreOptionsKind, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKindMap flattens the contents of JobTriggerInspectJobStorageConfigDatastoreOptionsKind from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKindMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigDatastoreOptionsKind)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKindSlice flattens the contents of JobTriggerInspectJobStorageConfigDatastoreOptionsKind from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKindSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	}

	items := make([]JobTriggerInspectJobStorageConfigDatastoreOptionsKind, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigDatastoreOptionsKind expands an instance of JobTriggerInspectJobStorageConfigDatastoreOptionsKind into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c *Client, f *JobTriggerInspectJobStorageConfigDatastoreOptionsKind, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKind flattens an instance of JobTriggerInspectJobStorageConfigDatastoreOptionsKind from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigDatastoreOptionsKind(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigDatastoreOptionsKind {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsKind
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptionsMap expands the contents of JobTriggerInspectJobStorageConfigCloudStorageOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptionsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigCloudStorageOptions, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptionsSlice expands the contents of JobTriggerInspectJobStorageConfigCloudStorageOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptionsSlice(c *Client, f []JobTriggerInspectJobStorageConfigCloudStorageOptions, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsMap flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigCloudStorageOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptions{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptions{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigCloudStorageOptions)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigCloudStorageOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSlice flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigCloudStorageOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptions{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptions{}
	}

	items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigCloudStorageOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptions expands an instance of JobTriggerInspectJobStorageConfigCloudStorageOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptions(c *Client, f *JobTriggerInspectJobStorageConfigCloudStorageOptions, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, f.FileSet, res); err != nil {
		return nil, fmt.Errorf("error expanding FileSet into fileSet: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["fileSet"] = v
	}
	if v := f.BytesLimitPerFile; !dcl.IsEmptyValueIndirect(v) {
		m["bytesLimitPerFile"] = v
	}
	if v := f.BytesLimitPerFilePercent; !dcl.IsEmptyValueIndirect(v) {
		m["bytesLimitPerFilePercent"] = v
	}
	if v := f.FileTypes; v != nil {
		m["fileTypes"] = v
	}
	if v := f.SampleMethod; !dcl.IsEmptyValueIndirect(v) {
		m["sampleMethod"] = v
	}
	if v := f.FilesLimitPercent; !dcl.IsEmptyValueIndirect(v) {
		m["filesLimitPercent"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptions flattens an instance of JobTriggerInspectJobStorageConfigCloudStorageOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptions(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigCloudStorageOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigCloudStorageOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigCloudStorageOptions
	}
	r.FileSet = flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, m["fileSet"], res)
	r.BytesLimitPerFile = dcl.FlattenInteger(m["bytesLimitPerFile"])
	r.BytesLimitPerFilePercent = dcl.FlattenInteger(m["bytesLimitPerFilePercent"])
	r.FileTypes = flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumSlice(c, m["fileTypes"], res)
	r.SampleMethod = flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(m["sampleMethod"])
	r.FilesLimitPercent = dcl.FlattenInteger(m["filesLimitPercent"])

	return r
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetMap expands the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetSlice expands the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetSlice(c *Client, f []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetMap flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetSlice flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	}

	items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet expands an instance of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c *Client, f *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, f.RegexFileSet, res); err != nil {
		return nil, fmt.Errorf("error expanding RegexFileSet into regexFileSet: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["regexFileSet"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet flattens an instance of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet
	}
	r.Url = dcl.FlattenString(m["url"])
	r.RegexFileSet = flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, m["regexFileSet"], res)

	return r
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetMap expands the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetSlice expands the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetSlice(c *Client, f []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetMap flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetSlice flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
	}

	items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet expands an instance of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c *Client, f *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.BucketName; !dcl.IsEmptyValueIndirect(v) {
		m["bucketName"] = v
	}
	if v := f.IncludeRegex; v != nil {
		m["includeRegex"] = v
	}
	if v := f.ExcludeRegex; v != nil {
		m["excludeRegex"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet flattens an instance of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet
	}
	r.BucketName = dcl.FlattenString(m["bucketName"])
	r.IncludeRegex = dcl.FlattenStringSlice(m["includeRegex"])
	r.ExcludeRegex = dcl.FlattenStringSlice(m["excludeRegex"])

	return r
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsMap expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigBigQueryOptions, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsSlice expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsSlice(c *Client, f []JobTriggerInspectJobStorageConfigBigQueryOptions, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsMap flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigBigQueryOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptions{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptions{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigBigQueryOptions)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigBigQueryOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSlice flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigBigQueryOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigBigQueryOptions{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigBigQueryOptions{}
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigBigQueryOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptions expands an instance of JobTriggerInspectJobStorageConfigBigQueryOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptions(c *Client, f *JobTriggerInspectJobStorageConfigBigQueryOptions, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, f.TableReference, res); err != nil {
		return nil, fmt.Errorf("error expanding TableReference into tableReference: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tableReference"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice(c, f.IdentifyingFields, res); err != nil {
		return nil, fmt.Errorf("error expanding IdentifyingFields into identifyingFields: %w", err)
	} else if v != nil {
		m["identifyingFields"] = v
	}
	if v := f.RowsLimit; !dcl.IsEmptyValueIndirect(v) {
		m["rowsLimit"] = v
	}
	if v := f.RowsLimitPercent; !dcl.IsEmptyValueIndirect(v) {
		m["rowsLimitPercent"] = v
	}
	if v := f.SampleMethod; !dcl.IsEmptyValueIndirect(v) {
		m["sampleMethod"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice(c, f.ExcludedFields, res); err != nil {
		return nil, fmt.Errorf("error expanding ExcludedFields into excludedFields: %w", err)
	} else if v != nil {
		m["excludedFields"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice(c, f.IncludedFields, res); err != nil {
		return nil, fmt.Errorf("error expanding IncludedFields into includedFields: %w", err)
	} else if v != nil {
		m["includedFields"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptions flattens an instance of JobTriggerInspectJobStorageConfigBigQueryOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptions(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigBigQueryOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigBigQueryOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigBigQueryOptions
	}
	r.TableReference = flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, m["tableReference"], res)
	r.IdentifyingFields = flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice(c, m["identifyingFields"], res)
	r.RowsLimit = dcl.FlattenInteger(m["rowsLimit"])
	r.RowsLimitPercent = dcl.FlattenInteger(m["rowsLimitPercent"])
	r.SampleMethod = flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(m["sampleMethod"])
	r.ExcludedFields = flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice(c, m["excludedFields"], res)
	r.IncludedFields = flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice(c, m["includedFields"], res)

	return r
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceMap expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceSlice expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceSlice(c *Client, f []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceMap flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceSlice flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference expands an instance of JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c *Client, f *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.DatasetId); err != nil {
		return nil, fmt.Errorf("error expanding DatasetId into datasetId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["datasetId"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.TableId); err != nil {
		return nil, fmt.Errorf("error expanding TableId into tableId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tableId"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference flattens an instance of JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.DatasetId = dcl.FlattenString(m["datasetId"])
	r.TableId = dcl.FlattenString(m["tableId"])

	return r
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsMap expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice(c *Client, f []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsMap flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields expands an instance of JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c *Client, f *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields flattens an instance of JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsMap expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice(c *Client, f []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsMap flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields expands an instance of JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c *Client, f *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields flattens an instance of JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsMap expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice expands the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice(c *Client, f []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsMap flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields expands an instance of JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c *Client, f *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields flattens an instance of JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandJobTriggerInspectJobStorageConfigHybridOptionsMap expands the contents of JobTriggerInspectJobStorageConfigHybridOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptionsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigHybridOptions, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigHybridOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigHybridOptionsSlice expands the contents of JobTriggerInspectJobStorageConfigHybridOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptionsSlice(c *Client, f []JobTriggerInspectJobStorageConfigHybridOptions, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigHybridOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigHybridOptionsMap flattens the contents of JobTriggerInspectJobStorageConfigHybridOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptionsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigHybridOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigHybridOptions{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigHybridOptions{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigHybridOptions)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigHybridOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigHybridOptionsSlice flattens the contents of JobTriggerInspectJobStorageConfigHybridOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptionsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigHybridOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigHybridOptions{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigHybridOptions{}
	}

	items := make([]JobTriggerInspectJobStorageConfigHybridOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigHybridOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigHybridOptions expands an instance of JobTriggerInspectJobStorageConfigHybridOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptions(c *Client, f *JobTriggerInspectJobStorageConfigHybridOptions, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		m["description"] = v
	}
	if v := f.RequiredFindingLabelKeys; v != nil {
		m["requiredFindingLabelKeys"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, f.TableOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding TableOptions into tableOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tableOptions"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigHybridOptions flattens an instance of JobTriggerInspectJobStorageConfigHybridOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptions(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigHybridOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigHybridOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigHybridOptions
	}
	r.Description = dcl.FlattenString(m["description"])
	r.RequiredFindingLabelKeys = dcl.FlattenStringSlice(m["requiredFindingLabelKeys"])
	r.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	r.TableOptions = flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, m["tableOptions"], res)

	return r
}

// expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsMap expands the contents of JobTriggerInspectJobStorageConfigHybridOptionsTableOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptions, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsSlice expands the contents of JobTriggerInspectJobStorageConfigHybridOptionsTableOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsSlice(c *Client, f []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsMap flattens the contents of JobTriggerInspectJobStorageConfigHybridOptionsTableOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptions)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsSlice flattens the contents of JobTriggerInspectJobStorageConfigHybridOptionsTableOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	}

	items := make([]JobTriggerInspectJobStorageConfigHybridOptionsTableOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptions expands an instance of JobTriggerInspectJobStorageConfigHybridOptionsTableOptions into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c *Client, f *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice(c, f.IdentifyingFields, res); err != nil {
		return nil, fmt.Errorf("error expanding IdentifyingFields into identifyingFields: %w", err)
	} else if v != nil {
		m["identifyingFields"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptions flattens an instance of JobTriggerInspectJobStorageConfigHybridOptionsTableOptions from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptions(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptions
	}
	r.IdentifyingFields = flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice(c, m["identifyingFields"], res)

	return r
}

// expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsMap expands the contents of JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice expands the contents of JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice(c *Client, f []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsMap flattens the contents of JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice flattens the contents of JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}
	}

	items := make([]JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields expands an instance of JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c *Client, f *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields flattens an instance of JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandJobTriggerInspectJobStorageConfigTimespanConfigMap expands the contents of JobTriggerInspectJobStorageConfigTimespanConfig into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigTimespanConfigMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigTimespanConfig, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigTimespanConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigTimespanConfigSlice expands the contents of JobTriggerInspectJobStorageConfigTimespanConfig into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigTimespanConfigSlice(c *Client, f []JobTriggerInspectJobStorageConfigTimespanConfig, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigTimespanConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigTimespanConfigMap flattens the contents of JobTriggerInspectJobStorageConfigTimespanConfig from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigTimespanConfigMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigTimespanConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigTimespanConfig{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigTimespanConfig{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigTimespanConfig)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigTimespanConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigTimespanConfigSlice flattens the contents of JobTriggerInspectJobStorageConfigTimespanConfig from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigTimespanConfigSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigTimespanConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigTimespanConfig{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigTimespanConfig{}
	}

	items := make([]JobTriggerInspectJobStorageConfigTimespanConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigTimespanConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigTimespanConfig expands an instance of JobTriggerInspectJobStorageConfigTimespanConfig into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigTimespanConfig(c *Client, f *JobTriggerInspectJobStorageConfigTimespanConfig, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.EndTime; !dcl.IsEmptyValueIndirect(v) {
		m["endTime"] = v
	}
	if v, err := expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, f.TimestampField, res); err != nil {
		return nil, fmt.Errorf("error expanding TimestampField into timestampField: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["timestampField"] = v
	}
	if v := f.EnableAutoPopulationOfTimespanConfig; !dcl.IsEmptyValueIndirect(v) {
		m["enableAutoPopulationOfTimespanConfig"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigTimespanConfig flattens an instance of JobTriggerInspectJobStorageConfigTimespanConfig from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigTimespanConfig(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigTimespanConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigTimespanConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigTimespanConfig
	}
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.EndTime = dcl.FlattenString(m["endTime"])
	r.TimestampField = flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, m["timestampField"], res)
	r.EnableAutoPopulationOfTimespanConfig = dcl.FlattenBool(m["enableAutoPopulationOfTimespanConfig"])

	return r
}

// expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldMap expands the contents of JobTriggerInspectJobStorageConfigTimespanConfigTimestampField into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldMap(c *Client, f map[string]JobTriggerInspectJobStorageConfigTimespanConfigTimestampField, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldSlice expands the contents of JobTriggerInspectJobStorageConfigTimespanConfigTimestampField into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldSlice(c *Client, f []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldMap flattens the contents of JobTriggerInspectJobStorageConfigTimespanConfigTimestampField from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigTimespanConfigTimestampField)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldSlice flattens the contents of JobTriggerInspectJobStorageConfigTimespanConfigTimestampField from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	}

	items := make([]JobTriggerInspectJobStorageConfigTimespanConfigTimestampField, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampField expands an instance of JobTriggerInspectJobStorageConfigTimespanConfigTimestampField into a JSON
// request object.
func expandJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c *Client, f *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampField flattens an instance of JobTriggerInspectJobStorageConfigTimespanConfigTimestampField from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigTimespanConfigTimestampField(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobStorageConfigTimespanConfigTimestampField
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandJobTriggerInspectJobInspectConfigMap expands the contents of JobTriggerInspectJobInspectConfig into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigMap(c *Client, f map[string]JobTriggerInspectJobInspectConfig, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigSlice expands the contents of JobTriggerInspectJobInspectConfig into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigSlice(c *Client, f []JobTriggerInspectJobInspectConfig, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigMap flattens the contents of JobTriggerInspectJobInspectConfig from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfig{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfig{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfig)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigSlice flattens the contents of JobTriggerInspectJobInspectConfig from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfig{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfig{}
	}

	items := make([]JobTriggerInspectJobInspectConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfig expands an instance of JobTriggerInspectJobInspectConfig into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfig(c *Client, f *JobTriggerInspectJobInspectConfig, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigInfoTypesSlice(c, f.InfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoTypes into infoTypes: %w", err)
	} else if v != nil {
		m["infoTypes"] = v
	}
	if v := f.MinLikelihood; !dcl.IsEmptyValueIndirect(v) {
		m["minLikelihood"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigLimits(c, f.Limits, res); err != nil {
		return nil, fmt.Errorf("error expanding Limits into limits: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["limits"] = v
	}
	if v := f.IncludeQuote; !dcl.IsEmptyValueIndirect(v) {
		m["includeQuote"] = v
	}
	if v := f.ExcludeInfoTypes; !dcl.IsEmptyValueIndirect(v) {
		m["excludeInfoTypes"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesSlice(c, f.CustomInfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding CustomInfoTypes into customInfoTypes: %w", err)
	} else if v != nil {
		m["customInfoTypes"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetSlice(c, f.RuleSet, res); err != nil {
		return nil, fmt.Errorf("error expanding RuleSet into ruleSet: %w", err)
	} else if v != nil {
		m["ruleSet"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfig flattens an instance of JobTriggerInspectJobInspectConfig from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfig(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfig
	}
	r.InfoTypes = flattenJobTriggerInspectJobInspectConfigInfoTypesSlice(c, m["infoTypes"], res)
	r.MinLikelihood = flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnum(m["minLikelihood"])
	r.Limits = flattenJobTriggerInspectJobInspectConfigLimits(c, m["limits"], res)
	r.IncludeQuote = dcl.FlattenBool(m["includeQuote"])
	r.ExcludeInfoTypes = dcl.FlattenBool(m["excludeInfoTypes"])
	r.CustomInfoTypes = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSlice(c, m["customInfoTypes"], res)
	r.RuleSet = flattenJobTriggerInspectJobInspectConfigRuleSetSlice(c, m["ruleSet"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigInfoTypesMap expands the contents of JobTriggerInspectJobInspectConfigInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigInfoTypesMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigInfoTypesSlice expands the contents of JobTriggerInspectJobInspectConfigInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigInfoTypesSlice(c *Client, f []JobTriggerInspectJobInspectConfigInfoTypes, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigInfoTypesMap flattens the contents of JobTriggerInspectJobInspectConfigInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigInfoTypesMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigInfoTypes{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigInfoTypes)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigInfoTypesSlice flattens the contents of JobTriggerInspectJobInspectConfigInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigInfoTypesSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigInfoTypes{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigInfoTypes{}
	}

	items := make([]JobTriggerInspectJobInspectConfigInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigInfoTypes expands an instance of JobTriggerInspectJobInspectConfigInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigInfoTypes(c *Client, f *JobTriggerInspectJobInspectConfigInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigInfoTypes flattens an instance of JobTriggerInspectJobInspectConfigInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigInfoTypes(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigInfoTypes
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandJobTriggerInspectJobInspectConfigLimitsMap expands the contents of JobTriggerInspectJobInspectConfigLimits into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimitsMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigLimits, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigLimits(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigLimitsSlice expands the contents of JobTriggerInspectJobInspectConfigLimits into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimitsSlice(c *Client, f []JobTriggerInspectJobInspectConfigLimits, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigLimits(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigLimitsMap flattens the contents of JobTriggerInspectJobInspectConfigLimits from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimitsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigLimits {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigLimits{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigLimits{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigLimits)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigLimits(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigLimitsSlice flattens the contents of JobTriggerInspectJobInspectConfigLimits from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimitsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigLimits {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigLimits{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigLimits{}
	}

	items := make([]JobTriggerInspectJobInspectConfigLimits, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigLimits(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigLimits expands an instance of JobTriggerInspectJobInspectConfigLimits into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimits(c *Client, f *JobTriggerInspectJobInspectConfigLimits, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxFindingsPerItem; !dcl.IsEmptyValueIndirect(v) {
		m["maxFindingsPerItem"] = v
	}
	if v := f.MaxFindingsPerRequest; !dcl.IsEmptyValueIndirect(v) {
		m["maxFindingsPerRequest"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c, f.MaxFindingsPerInfoType, res); err != nil {
		return nil, fmt.Errorf("error expanding MaxFindingsPerInfoType into maxFindingsPerInfoType: %w", err)
	} else if v != nil {
		m["maxFindingsPerInfoType"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigLimits flattens an instance of JobTriggerInspectJobInspectConfigLimits from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimits(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigLimits {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigLimits{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigLimits
	}
	r.MaxFindingsPerItem = dcl.FlattenInteger(m["maxFindingsPerItem"])
	r.MaxFindingsPerRequest = dcl.FlattenInteger(m["maxFindingsPerRequest"])
	r.MaxFindingsPerInfoType = flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c, m["maxFindingsPerInfoType"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeMap expands the contents of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice expands the contents of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c *Client, f []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeMap flattens the contents of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice flattens the contents of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}
	}

	items := make([]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType expands an instance of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c *Client, f *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, f.InfoType, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoType into infoType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["infoType"] = v
	}
	if v := f.MaxFindings; !dcl.IsEmptyValueIndirect(v) {
		m["maxFindings"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType flattens an instance of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType
	}
	r.InfoType = flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, m["infoType"], res)
	r.MaxFindings = dcl.FlattenInteger(m["maxFindings"])

	return r
}

// expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeMap expands the contents of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice expands the contents of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice(c *Client, f []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeMap flattens the contents of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice flattens the contents of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}

	items := make([]JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType expands an instance of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c *Client, f *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType flattens an instance of JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypes, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypes{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypes)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypes{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypes{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypes expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypes(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, f.InfoType, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoType into infoType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["infoType"] = v
	}
	if v := f.Likelihood; !dcl.IsEmptyValueIndirect(v) {
		m["likelihood"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, f.Dictionary, res); err != nil {
		return nil, fmt.Errorf("error expanding Dictionary into dictionary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dictionary"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, f.Regex, res); err != nil {
		return nil, fmt.Errorf("error expanding Regex into regex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["regex"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, f.SurrogateType, res); err != nil {
		return nil, fmt.Errorf("error expanding SurrogateType into surrogateType: %w", err)
	} else if v != nil {
		m["surrogateType"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, f.StoredType, res); err != nil {
		return nil, fmt.Errorf("error expanding StoredType into storedType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["storedType"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice(c, f.DetectionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding DetectionRules into detectionRules: %w", err)
	} else if v != nil {
		m["detectionRules"] = v
	}
	if v := f.ExclusionType; !dcl.IsEmptyValueIndirect(v) {
		m["exclusionType"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypes flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypes(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypes
	}
	r.InfoType = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, m["infoType"], res)
	r.Likelihood = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(m["likelihood"])
	r.Dictionary = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, m["dictionary"], res)
	r.Regex = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, m["regex"], res)
	r.SurrogateType = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, m["surrogateType"], res)
	r.StoredType = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, m["storedType"], res)
	r.DetectionRules = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice(c, m["detectionRules"], res)
	r.ExclusionType = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(m["exclusionType"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionarySlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionarySlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionarySlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionarySlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, f.WordList, res); err != nil {
		return nil, fmt.Errorf("error expanding WordList into wordList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["wordList"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, f.CloudStoragePath, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudStoragePath into cloudStoragePath: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudStoragePath"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary
	}
	r.WordList = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, m["wordList"], res)
	r.CloudStoragePath = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, m["cloudStoragePath"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Words; v != nil {
		m["words"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList
	}
	r.Words = dcl.FlattenStringSlice(m["words"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	}
	r.Path = dcl.FlattenString(m["path"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegexMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegexMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesRegex, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegexSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegexSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegexMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegexMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesRegex)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegexSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegexSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegex expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Pattern; !dcl.IsEmptyValueIndirect(v) {
		m["pattern"] = v
	}
	if v := f.GroupIndexes; v != nil {
		m["groupIndexes"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegex flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesRegex(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesRegex
	}
	r.Pattern = dcl.FlattenString(m["pattern"])
	r.GroupIndexes = dcl.FlattenIntSlice(m["groupIndexes"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType
	}

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType
	}
	r.Name = dcl.FlattenString(m["name"])
	r.CreateTime = dcl.FlattenString(m["createTime"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, f.HotwordRule, res); err != nil {
		return nil, fmt.Errorf("error expanding HotwordRule into hotwordRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hotwordRule"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules
	}
	r.HotwordRule = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, m["hotwordRule"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, f.HotwordRegex, res); err != nil {
		return nil, fmt.Errorf("error expanding HotwordRegex into hotwordRegex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hotwordRegex"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, f.Proximity, res); err != nil {
		return nil, fmt.Errorf("error expanding Proximity into proximity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["proximity"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, f.LikelihoodAdjustment, res); err != nil {
		return nil, fmt.Errorf("error expanding LikelihoodAdjustment into likelihoodAdjustment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["likelihoodAdjustment"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule
	}
	r.HotwordRegex = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, m["hotwordRegex"], res)
	r.Proximity = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, m["proximity"], res)
	r.LikelihoodAdjustment = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, m["likelihoodAdjustment"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Pattern; !dcl.IsEmptyValueIndirect(v) {
		m["pattern"] = v
	}
	if v := f.GroupIndexes; v != nil {
		m["groupIndexes"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex
	}
	r.Pattern = dcl.FlattenString(m["pattern"])
	r.GroupIndexes = dcl.FlattenIntSlice(m["groupIndexes"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximitySlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximitySlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximitySlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximitySlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.WindowBefore; !dcl.IsEmptyValueIndirect(v) {
		m["windowBefore"] = v
	}
	if v := f.WindowAfter; !dcl.IsEmptyValueIndirect(v) {
		m["windowAfter"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity
	}
	r.WindowBefore = dcl.FlattenInteger(m["windowBefore"])
	r.WindowAfter = dcl.FlattenInteger(m["windowAfter"])

	return r
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentMap expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentSlice expands the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, f []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment expands an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c *Client, f *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FixedLikelihood; !dcl.IsEmptyValueIndirect(v) {
		m["fixedLikelihood"] = v
	}
	if v := f.RelativeLikelihood; !dcl.IsEmptyValueIndirect(v) {
		m["relativeLikelihood"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment flattens an instance of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment
	}
	r.FixedLikelihood = flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(m["fixedLikelihood"])
	r.RelativeLikelihood = dcl.FlattenInteger(m["relativeLikelihood"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetMap expands the contents of JobTriggerInspectJobInspectConfigRuleSet into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSet, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSet(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSet into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSet, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSet(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSet from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSet{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSet{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSet)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSet(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSet from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSet {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSet{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSet{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSet(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSet expands an instance of JobTriggerInspectJobInspectConfigRuleSet into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSet(c *Client, f *JobTriggerInspectJobInspectConfigRuleSet, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice(c, f.InfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoTypes into infoTypes: %w", err)
	} else if v != nil {
		m["infoTypes"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesSlice(c, f.Rules, res); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if v != nil {
		m["rules"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSet flattens an instance of JobTriggerInspectJobInspectConfigRuleSet from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSet(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSet{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSet
	}
	r.InfoTypes = flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice(c, m["infoTypes"], res)
	r.Rules = flattenJobTriggerInspectJobInspectConfigRuleSetRulesSlice(c, m["rules"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetInfoTypesMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetInfoTypesMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetInfoTypes, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypesMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypesMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetInfoTypes{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetInfoTypes)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypesSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetInfoTypes{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetInfoTypes{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetInfoTypes expands an instance of JobTriggerInspectJobInspectConfigRuleSetInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypes flattens an instance of JobTriggerInspectJobInspectConfigRuleSetInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetInfoTypes(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetInfoTypes
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRules into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRules, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRules into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRules, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRules from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRules{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRules{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRules)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRules from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRules {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRules{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRules{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRules expands an instance of JobTriggerInspectJobInspectConfigRuleSetRules into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRules(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRules, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, f.HotwordRule, res); err != nil {
		return nil, fmt.Errorf("error expanding HotwordRule into hotwordRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hotwordRule"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, f.ExclusionRule, res); err != nil {
		return nil, fmt.Errorf("error expanding ExclusionRule into exclusionRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exclusionRule"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRules flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRules from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRules(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRules
	}
	r.HotwordRule = flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, m["hotwordRule"], res)
	r.ExclusionRule = flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, m["exclusionRule"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, f.HotwordRegex, res); err != nil {
		return nil, fmt.Errorf("error expanding HotwordRegex into hotwordRegex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hotwordRegex"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, f.Proximity, res); err != nil {
		return nil, fmt.Errorf("error expanding Proximity into proximity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["proximity"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, f.LikelihoodAdjustment, res); err != nil {
		return nil, fmt.Errorf("error expanding LikelihoodAdjustment into likelihoodAdjustment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["likelihoodAdjustment"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	}
	r.HotwordRegex = flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, m["hotwordRegex"], res)
	r.Proximity = flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, m["proximity"], res)
	r.LikelihoodAdjustment = flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, m["likelihoodAdjustment"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Pattern; !dcl.IsEmptyValueIndirect(v) {
		m["pattern"] = v
	}
	if v := f.GroupIndexes; v != nil {
		m["groupIndexes"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	}
	r.Pattern = dcl.FlattenString(m["pattern"])
	r.GroupIndexes = dcl.FlattenIntSlice(m["groupIndexes"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximitySlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximitySlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximitySlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximitySlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.WindowBefore; !dcl.IsEmptyValueIndirect(v) {
		m["windowBefore"] = v
	}
	if v := f.WindowAfter; !dcl.IsEmptyValueIndirect(v) {
		m["windowAfter"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity
	}
	r.WindowBefore = dcl.FlattenInteger(m["windowBefore"])
	r.WindowAfter = dcl.FlattenInteger(m["windowAfter"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FixedLikelihood; !dcl.IsEmptyValueIndirect(v) {
		m["fixedLikelihood"] = v
	}
	if v := f.RelativeLikelihood; !dcl.IsEmptyValueIndirect(v) {
		m["relativeLikelihood"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	}
	r.FixedLikelihood = flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(m["fixedLikelihood"])
	r.RelativeLikelihood = dcl.FlattenInteger(m["relativeLikelihood"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, f.Dictionary, res); err != nil {
		return nil, fmt.Errorf("error expanding Dictionary into dictionary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dictionary"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, f.Regex, res); err != nil {
		return nil, fmt.Errorf("error expanding Regex into regex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["regex"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, f.ExcludeInfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding ExcludeInfoTypes into excludeInfoTypes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["excludeInfoTypes"] = v
	}
	if v := f.MatchingType; !dcl.IsEmptyValueIndirect(v) {
		m["matchingType"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	}
	r.Dictionary = flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, m["dictionary"], res)
	r.Regex = flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, m["regex"], res)
	r.ExcludeInfoTypes = flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, m["excludeInfoTypes"], res)
	r.MatchingType = flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(m["matchingType"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionarySlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionarySlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionarySlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionarySlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, f.WordList, res); err != nil {
		return nil, fmt.Errorf("error expanding WordList into wordList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["wordList"] = v
	}
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, f.CloudStoragePath, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudStoragePath into cloudStoragePath: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudStoragePath"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	}
	r.WordList = flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, m["wordList"], res)
	r.CloudStoragePath = flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, m["cloudStoragePath"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Words; v != nil {
		m["words"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	}
	r.Words = dcl.FlattenStringSlice(m["words"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	}
	r.Path = dcl.FlattenString(m["path"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Pattern; !dcl.IsEmptyValueIndirect(v) {
		m["pattern"] = v
	}
	if v := f.GroupIndexes; v != nil {
		m["groupIndexes"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
	}
	r.Pattern = dcl.FlattenString(m["pattern"])
	r.GroupIndexes = dcl.FlattenIntSlice(m["groupIndexes"])

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c, f.InfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoTypes into infoTypes: %w", err)
	} else if v != nil {
		m["infoTypes"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	}
	r.InfoTypes = flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c, m["infoTypes"], res)

	return r
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesMap expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesMap(c *Client, f map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice expands the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c *Client, f []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes expands an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes into a JSON
// request object.
func expandJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c *Client, f *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes flattens an instance of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandJobTriggerInspectJobActionsMap expands the contents of JobTriggerInspectJobActions into a JSON
// request object.
func expandJobTriggerInspectJobActionsMap(c *Client, f map[string]JobTriggerInspectJobActions, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsSlice expands the contents of JobTriggerInspectJobActions into a JSON
// request object.
func expandJobTriggerInspectJobActionsSlice(c *Client, f []JobTriggerInspectJobActions, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsMap flattens the contents of JobTriggerInspectJobActions from a JSON
// response object.
func flattenJobTriggerInspectJobActionsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActions{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActions{}
	}

	items := make(map[string]JobTriggerInspectJobActions)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsSlice flattens the contents of JobTriggerInspectJobActions from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActions {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActions{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActions{}
	}

	items := make([]JobTriggerInspectJobActions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActions expands an instance of JobTriggerInspectJobActions into a JSON
// request object.
func expandJobTriggerInspectJobActions(c *Client, f *JobTriggerInspectJobActions, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobActionsSaveFindings(c, f.SaveFindings, res); err != nil {
		return nil, fmt.Errorf("error expanding SaveFindings into saveFindings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["saveFindings"] = v
	}
	if v, err := expandJobTriggerInspectJobActionsPubSub(c, f.PubSub, res); err != nil {
		return nil, fmt.Errorf("error expanding PubSub into pubSub: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pubSub"] = v
	}
	if v, err := expandJobTriggerInspectJobActionsPublishSummaryToCscc(c, f.PublishSummaryToCscc, res); err != nil {
		return nil, fmt.Errorf("error expanding PublishSummaryToCscc into publishSummaryToCscc: %w", err)
	} else if v != nil {
		m["publishSummaryToCscc"] = v
	}
	if v, err := expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, f.PublishFindingsToCloudDataCatalog, res); err != nil {
		return nil, fmt.Errorf("error expanding PublishFindingsToCloudDataCatalog into publishFindingsToCloudDataCatalog: %w", err)
	} else if v != nil {
		m["publishFindingsToCloudDataCatalog"] = v
	}
	if v, err := expandJobTriggerInspectJobActionsJobNotificationEmails(c, f.JobNotificationEmails, res); err != nil {
		return nil, fmt.Errorf("error expanding JobNotificationEmails into jobNotificationEmails: %w", err)
	} else if v != nil {
		m["jobNotificationEmails"] = v
	}
	if v, err := expandJobTriggerInspectJobActionsPublishToStackdriver(c, f.PublishToStackdriver, res); err != nil {
		return nil, fmt.Errorf("error expanding PublishToStackdriver into publishToStackdriver: %w", err)
	} else if v != nil {
		m["publishToStackdriver"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobActions flattens an instance of JobTriggerInspectJobActions from a JSON
// response object.
func flattenJobTriggerInspectJobActions(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActions
	}
	r.SaveFindings = flattenJobTriggerInspectJobActionsSaveFindings(c, m["saveFindings"], res)
	r.PubSub = flattenJobTriggerInspectJobActionsPubSub(c, m["pubSub"], res)
	r.PublishSummaryToCscc = flattenJobTriggerInspectJobActionsPublishSummaryToCscc(c, m["publishSummaryToCscc"], res)
	r.PublishFindingsToCloudDataCatalog = flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, m["publishFindingsToCloudDataCatalog"], res)
	r.JobNotificationEmails = flattenJobTriggerInspectJobActionsJobNotificationEmails(c, m["jobNotificationEmails"], res)
	r.PublishToStackdriver = flattenJobTriggerInspectJobActionsPublishToStackdriver(c, m["publishToStackdriver"], res)

	return r
}

// expandJobTriggerInspectJobActionsSaveFindingsMap expands the contents of JobTriggerInspectJobActionsSaveFindings into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsMap(c *Client, f map[string]JobTriggerInspectJobActionsSaveFindings, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsSaveFindings(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsSaveFindingsSlice expands the contents of JobTriggerInspectJobActionsSaveFindings into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsSlice(c *Client, f []JobTriggerInspectJobActionsSaveFindings, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsSaveFindings(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsSaveFindingsMap flattens the contents of JobTriggerInspectJobActionsSaveFindings from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsSaveFindings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsSaveFindings{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsSaveFindings{}
	}

	items := make(map[string]JobTriggerInspectJobActionsSaveFindings)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsSaveFindings(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsSaveFindingsSlice flattens the contents of JobTriggerInspectJobActionsSaveFindings from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsSaveFindings {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsSaveFindings{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsSaveFindings{}
	}

	items := make([]JobTriggerInspectJobActionsSaveFindings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsSaveFindings(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsSaveFindings expands an instance of JobTriggerInspectJobActionsSaveFindings into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindings(c *Client, f *JobTriggerInspectJobActionsSaveFindings, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, f.OutputConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding OutputConfig into outputConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["outputConfig"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobActionsSaveFindings flattens an instance of JobTriggerInspectJobActionsSaveFindings from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindings(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsSaveFindings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsSaveFindings{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsSaveFindings
	}
	r.OutputConfig = flattenJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, m["outputConfig"], res)

	return r
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfigMap expands the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfig into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfigMap(c *Client, f map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfig, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfigSlice expands the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfig into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfigSlice(c *Client, f []JobTriggerInspectJobActionsSaveFindingsOutputConfig, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigMap flattens the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfig from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	}

	items := make(map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfig)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigSlice flattens the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfig from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	}

	items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsSaveFindingsOutputConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfig expands an instance of JobTriggerInspectJobActionsSaveFindingsOutputConfig into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfig(c *Client, f *JobTriggerInspectJobActionsSaveFindingsOutputConfig, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, f.Table, res); err != nil {
		return nil, fmt.Errorf("error expanding Table into table: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["table"] = v
	}
	if v, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, f.DlpStorage, res); err != nil {
		return nil, fmt.Errorf("error expanding DlpStorage into dlpStorage: %w", err)
	} else if v != nil {
		m["dlpStorage"] = v
	}
	if v := f.OutputSchema; !dcl.IsEmptyValueIndirect(v) {
		m["outputSchema"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfig flattens an instance of JobTriggerInspectJobActionsSaveFindingsOutputConfig from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfig(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsSaveFindingsOutputConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsSaveFindingsOutputConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfig
	}
	r.Table = flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, m["table"], res)
	r.DlpStorage = flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, m["dlpStorage"], res)
	r.OutputSchema = flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(m["outputSchema"])

	return r
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTableMap expands the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigTable into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTableMap(c *Client, f map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigTable, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTableSlice expands the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigTable into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTableSlice(c *Client, f []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTableMap flattens the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigTable from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTableMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	}

	items := make(map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigTable)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTableSlice flattens the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigTable from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTableSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	}

	items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfigTable, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTable expands an instance of JobTriggerInspectJobActionsSaveFindingsOutputConfigTable into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c *Client, f *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.SelfLinkToNameExpander(f.ProjectId); err != nil {
		return nil, fmt.Errorf("error expanding ProjectId into projectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.DatasetId); err != nil {
		return nil, fmt.Errorf("error expanding DatasetId into datasetId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["datasetId"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.TableId); err != nil {
		return nil, fmt.Errorf("error expanding TableId into tableId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tableId"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTable flattens an instance of JobTriggerInspectJobActionsSaveFindingsOutputConfigTable from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigTable(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigTable
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.DatasetId = dcl.FlattenString(m["datasetId"])
	r.TableId = dcl.FlattenString(m["tableId"])

	return r
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageMap expands the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageMap(c *Client, f map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageSlice expands the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageSlice(c *Client, f []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageMap flattens the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	}

	items := make(map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageSlice flattens the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	}

	items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage expands an instance of JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage into a JSON
// request object.
func expandJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c *Client, f *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage flattens an instance of JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage
	}

	return r
}

// expandJobTriggerInspectJobActionsPubSubMap expands the contents of JobTriggerInspectJobActionsPubSub into a JSON
// request object.
func expandJobTriggerInspectJobActionsPubSubMap(c *Client, f map[string]JobTriggerInspectJobActionsPubSub, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsPubSub(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsPubSubSlice expands the contents of JobTriggerInspectJobActionsPubSub into a JSON
// request object.
func expandJobTriggerInspectJobActionsPubSubSlice(c *Client, f []JobTriggerInspectJobActionsPubSub, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsPubSub(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsPubSubMap flattens the contents of JobTriggerInspectJobActionsPubSub from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPubSubMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsPubSub {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsPubSub{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsPubSub{}
	}

	items := make(map[string]JobTriggerInspectJobActionsPubSub)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsPubSub(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsPubSubSlice flattens the contents of JobTriggerInspectJobActionsPubSub from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPubSubSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsPubSub {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsPubSub{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsPubSub{}
	}

	items := make([]JobTriggerInspectJobActionsPubSub, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsPubSub(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsPubSub expands an instance of JobTriggerInspectJobActionsPubSub into a JSON
// request object.
func expandJobTriggerInspectJobActionsPubSub(c *Client, f *JobTriggerInspectJobActionsPubSub, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Topic; !dcl.IsEmptyValueIndirect(v) {
		m["topic"] = v
	}

	return m, nil
}

// flattenJobTriggerInspectJobActionsPubSub flattens an instance of JobTriggerInspectJobActionsPubSub from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPubSub(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsPubSub {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsPubSub{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsPubSub
	}
	r.Topic = dcl.FlattenString(m["topic"])

	return r
}

// expandJobTriggerInspectJobActionsPublishSummaryToCsccMap expands the contents of JobTriggerInspectJobActionsPublishSummaryToCscc into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishSummaryToCsccMap(c *Client, f map[string]JobTriggerInspectJobActionsPublishSummaryToCscc, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsPublishSummaryToCscc(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsPublishSummaryToCsccSlice expands the contents of JobTriggerInspectJobActionsPublishSummaryToCscc into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishSummaryToCsccSlice(c *Client, f []JobTriggerInspectJobActionsPublishSummaryToCscc, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsPublishSummaryToCscc(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsPublishSummaryToCsccMap flattens the contents of JobTriggerInspectJobActionsPublishSummaryToCscc from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishSummaryToCsccMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsPublishSummaryToCscc {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsPublishSummaryToCscc{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsPublishSummaryToCscc{}
	}

	items := make(map[string]JobTriggerInspectJobActionsPublishSummaryToCscc)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsPublishSummaryToCscc(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsPublishSummaryToCsccSlice flattens the contents of JobTriggerInspectJobActionsPublishSummaryToCscc from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishSummaryToCsccSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsPublishSummaryToCscc {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsPublishSummaryToCscc{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsPublishSummaryToCscc{}
	}

	items := make([]JobTriggerInspectJobActionsPublishSummaryToCscc, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsPublishSummaryToCscc(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsPublishSummaryToCscc expands an instance of JobTriggerInspectJobActionsPublishSummaryToCscc into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishSummaryToCscc(c *Client, f *JobTriggerInspectJobActionsPublishSummaryToCscc, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenJobTriggerInspectJobActionsPublishSummaryToCscc flattens an instance of JobTriggerInspectJobActionsPublishSummaryToCscc from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishSummaryToCscc(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsPublishSummaryToCscc {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsPublishSummaryToCscc{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsPublishSummaryToCscc
	}

	return r
}

// expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogMap expands the contents of JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogMap(c *Client, f map[string]JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogSlice expands the contents of JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogSlice(c *Client, f []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogMap flattens the contents of JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	}

	items := make(map[string]JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogSlice flattens the contents of JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	}

	items := make([]JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog expands an instance of JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c *Client, f *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog flattens an instance of JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	}

	return r
}

// expandJobTriggerInspectJobActionsJobNotificationEmailsMap expands the contents of JobTriggerInspectJobActionsJobNotificationEmails into a JSON
// request object.
func expandJobTriggerInspectJobActionsJobNotificationEmailsMap(c *Client, f map[string]JobTriggerInspectJobActionsJobNotificationEmails, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsJobNotificationEmails(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsJobNotificationEmailsSlice expands the contents of JobTriggerInspectJobActionsJobNotificationEmails into a JSON
// request object.
func expandJobTriggerInspectJobActionsJobNotificationEmailsSlice(c *Client, f []JobTriggerInspectJobActionsJobNotificationEmails, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsJobNotificationEmails(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsJobNotificationEmailsMap flattens the contents of JobTriggerInspectJobActionsJobNotificationEmails from a JSON
// response object.
func flattenJobTriggerInspectJobActionsJobNotificationEmailsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsJobNotificationEmails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsJobNotificationEmails{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsJobNotificationEmails{}
	}

	items := make(map[string]JobTriggerInspectJobActionsJobNotificationEmails)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsJobNotificationEmails(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsJobNotificationEmailsSlice flattens the contents of JobTriggerInspectJobActionsJobNotificationEmails from a JSON
// response object.
func flattenJobTriggerInspectJobActionsJobNotificationEmailsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsJobNotificationEmails {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsJobNotificationEmails{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsJobNotificationEmails{}
	}

	items := make([]JobTriggerInspectJobActionsJobNotificationEmails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsJobNotificationEmails(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsJobNotificationEmails expands an instance of JobTriggerInspectJobActionsJobNotificationEmails into a JSON
// request object.
func expandJobTriggerInspectJobActionsJobNotificationEmails(c *Client, f *JobTriggerInspectJobActionsJobNotificationEmails, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenJobTriggerInspectJobActionsJobNotificationEmails flattens an instance of JobTriggerInspectJobActionsJobNotificationEmails from a JSON
// response object.
func flattenJobTriggerInspectJobActionsJobNotificationEmails(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsJobNotificationEmails {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsJobNotificationEmails{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsJobNotificationEmails
	}

	return r
}

// expandJobTriggerInspectJobActionsPublishToStackdriverMap expands the contents of JobTriggerInspectJobActionsPublishToStackdriver into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishToStackdriverMap(c *Client, f map[string]JobTriggerInspectJobActionsPublishToStackdriver, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerInspectJobActionsPublishToStackdriver(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerInspectJobActionsPublishToStackdriverSlice expands the contents of JobTriggerInspectJobActionsPublishToStackdriver into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishToStackdriverSlice(c *Client, f []JobTriggerInspectJobActionsPublishToStackdriver, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerInspectJobActionsPublishToStackdriver(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerInspectJobActionsPublishToStackdriverMap flattens the contents of JobTriggerInspectJobActionsPublishToStackdriver from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishToStackdriverMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsPublishToStackdriver {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsPublishToStackdriver{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsPublishToStackdriver{}
	}

	items := make(map[string]JobTriggerInspectJobActionsPublishToStackdriver)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsPublishToStackdriver(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerInspectJobActionsPublishToStackdriverSlice flattens the contents of JobTriggerInspectJobActionsPublishToStackdriver from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishToStackdriverSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsPublishToStackdriver {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsPublishToStackdriver{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsPublishToStackdriver{}
	}

	items := make([]JobTriggerInspectJobActionsPublishToStackdriver, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsPublishToStackdriver(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerInspectJobActionsPublishToStackdriver expands an instance of JobTriggerInspectJobActionsPublishToStackdriver into a JSON
// request object.
func expandJobTriggerInspectJobActionsPublishToStackdriver(c *Client, f *JobTriggerInspectJobActionsPublishToStackdriver, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenJobTriggerInspectJobActionsPublishToStackdriver flattens an instance of JobTriggerInspectJobActionsPublishToStackdriver from a JSON
// response object.
func flattenJobTriggerInspectJobActionsPublishToStackdriver(c *Client, i interface{}, res *JobTrigger) *JobTriggerInspectJobActionsPublishToStackdriver {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerInspectJobActionsPublishToStackdriver{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerInspectJobActionsPublishToStackdriver
	}

	return r
}

// expandJobTriggerTriggersMap expands the contents of JobTriggerTriggers into a JSON
// request object.
func expandJobTriggerTriggersMap(c *Client, f map[string]JobTriggerTriggers, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerTriggers(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerTriggersSlice expands the contents of JobTriggerTriggers into a JSON
// request object.
func expandJobTriggerTriggersSlice(c *Client, f []JobTriggerTriggers, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerTriggers(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerTriggersMap flattens the contents of JobTriggerTriggers from a JSON
// response object.
func flattenJobTriggerTriggersMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerTriggers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerTriggers{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerTriggers{}
	}

	items := make(map[string]JobTriggerTriggers)
	for k, item := range a {
		items[k] = *flattenJobTriggerTriggers(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerTriggersSlice flattens the contents of JobTriggerTriggers from a JSON
// response object.
func flattenJobTriggerTriggersSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerTriggers {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerTriggers{}
	}

	if len(a) == 0 {
		return []JobTriggerTriggers{}
	}

	items := make([]JobTriggerTriggers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerTriggers(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerTriggers expands an instance of JobTriggerTriggers into a JSON
// request object.
func expandJobTriggerTriggers(c *Client, f *JobTriggerTriggers, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerTriggersSchedule(c, f.Schedule, res); err != nil {
		return nil, fmt.Errorf("error expanding Schedule into schedule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["schedule"] = v
	}
	if v, err := expandJobTriggerTriggersManual(c, f.Manual, res); err != nil {
		return nil, fmt.Errorf("error expanding Manual into manual: %w", err)
	} else if v != nil {
		m["manual"] = v
	}

	return m, nil
}

// flattenJobTriggerTriggers flattens an instance of JobTriggerTriggers from a JSON
// response object.
func flattenJobTriggerTriggers(c *Client, i interface{}, res *JobTrigger) *JobTriggerTriggers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerTriggers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerTriggers
	}
	r.Schedule = flattenJobTriggerTriggersSchedule(c, m["schedule"], res)
	r.Manual = flattenJobTriggerTriggersManual(c, m["manual"], res)

	return r
}

// expandJobTriggerTriggersScheduleMap expands the contents of JobTriggerTriggersSchedule into a JSON
// request object.
func expandJobTriggerTriggersScheduleMap(c *Client, f map[string]JobTriggerTriggersSchedule, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerTriggersSchedule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerTriggersScheduleSlice expands the contents of JobTriggerTriggersSchedule into a JSON
// request object.
func expandJobTriggerTriggersScheduleSlice(c *Client, f []JobTriggerTriggersSchedule, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerTriggersSchedule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerTriggersScheduleMap flattens the contents of JobTriggerTriggersSchedule from a JSON
// response object.
func flattenJobTriggerTriggersScheduleMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerTriggersSchedule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerTriggersSchedule{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerTriggersSchedule{}
	}

	items := make(map[string]JobTriggerTriggersSchedule)
	for k, item := range a {
		items[k] = *flattenJobTriggerTriggersSchedule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerTriggersScheduleSlice flattens the contents of JobTriggerTriggersSchedule from a JSON
// response object.
func flattenJobTriggerTriggersScheduleSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerTriggersSchedule {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerTriggersSchedule{}
	}

	if len(a) == 0 {
		return []JobTriggerTriggersSchedule{}
	}

	items := make([]JobTriggerTriggersSchedule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerTriggersSchedule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerTriggersSchedule expands an instance of JobTriggerTriggersSchedule into a JSON
// request object.
func expandJobTriggerTriggersSchedule(c *Client, f *JobTriggerTriggersSchedule, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.RecurrencePeriodDuration; !dcl.IsEmptyValueIndirect(v) {
		m["recurrencePeriodDuration"] = v
	}

	return m, nil
}

// flattenJobTriggerTriggersSchedule flattens an instance of JobTriggerTriggersSchedule from a JSON
// response object.
func flattenJobTriggerTriggersSchedule(c *Client, i interface{}, res *JobTrigger) *JobTriggerTriggersSchedule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerTriggersSchedule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerTriggersSchedule
	}
	r.RecurrencePeriodDuration = dcl.FlattenString(m["recurrencePeriodDuration"])

	return r
}

// expandJobTriggerTriggersManualMap expands the contents of JobTriggerTriggersManual into a JSON
// request object.
func expandJobTriggerTriggersManualMap(c *Client, f map[string]JobTriggerTriggersManual, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerTriggersManual(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerTriggersManualSlice expands the contents of JobTriggerTriggersManual into a JSON
// request object.
func expandJobTriggerTriggersManualSlice(c *Client, f []JobTriggerTriggersManual, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerTriggersManual(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerTriggersManualMap flattens the contents of JobTriggerTriggersManual from a JSON
// response object.
func flattenJobTriggerTriggersManualMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerTriggersManual {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerTriggersManual{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerTriggersManual{}
	}

	items := make(map[string]JobTriggerTriggersManual)
	for k, item := range a {
		items[k] = *flattenJobTriggerTriggersManual(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerTriggersManualSlice flattens the contents of JobTriggerTriggersManual from a JSON
// response object.
func flattenJobTriggerTriggersManualSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerTriggersManual {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerTriggersManual{}
	}

	if len(a) == 0 {
		return []JobTriggerTriggersManual{}
	}

	items := make([]JobTriggerTriggersManual, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerTriggersManual(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerTriggersManual expands an instance of JobTriggerTriggersManual into a JSON
// request object.
func expandJobTriggerTriggersManual(c *Client, f *JobTriggerTriggersManual, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenJobTriggerTriggersManual flattens an instance of JobTriggerTriggersManual from a JSON
// response object.
func flattenJobTriggerTriggersManual(c *Client, i interface{}, res *JobTrigger) *JobTriggerTriggersManual {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerTriggersManual{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerTriggersManual
	}

	return r
}

// expandJobTriggerErrorsMap expands the contents of JobTriggerErrors into a JSON
// request object.
func expandJobTriggerErrorsMap(c *Client, f map[string]JobTriggerErrors, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerErrors(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerErrorsSlice expands the contents of JobTriggerErrors into a JSON
// request object.
func expandJobTriggerErrorsSlice(c *Client, f []JobTriggerErrors, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerErrors(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerErrorsMap flattens the contents of JobTriggerErrors from a JSON
// response object.
func flattenJobTriggerErrorsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerErrors {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerErrors{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerErrors{}
	}

	items := make(map[string]JobTriggerErrors)
	for k, item := range a {
		items[k] = *flattenJobTriggerErrors(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerErrorsSlice flattens the contents of JobTriggerErrors from a JSON
// response object.
func flattenJobTriggerErrorsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerErrors {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerErrors{}
	}

	if len(a) == 0 {
		return []JobTriggerErrors{}
	}

	items := make([]JobTriggerErrors, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerErrors(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerErrors expands an instance of JobTriggerErrors into a JSON
// request object.
func expandJobTriggerErrors(c *Client, f *JobTriggerErrors, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandJobTriggerErrorsDetails(c, f.Details, res); err != nil {
		return nil, fmt.Errorf("error expanding Details into details: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["details"] = v
	}
	if v := f.Timestamps; v != nil {
		m["timestamps"] = v
	}

	return m, nil
}

// flattenJobTriggerErrors flattens an instance of JobTriggerErrors from a JSON
// response object.
func flattenJobTriggerErrors(c *Client, i interface{}, res *JobTrigger) *JobTriggerErrors {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerErrors{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerErrors
	}
	r.Details = flattenJobTriggerErrorsDetails(c, m["details"], res)
	r.Timestamps = dcl.FlattenStringSlice(m["timestamps"])

	return r
}

// expandJobTriggerErrorsDetailsMap expands the contents of JobTriggerErrorsDetails into a JSON
// request object.
func expandJobTriggerErrorsDetailsMap(c *Client, f map[string]JobTriggerErrorsDetails, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerErrorsDetails(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerErrorsDetailsSlice expands the contents of JobTriggerErrorsDetails into a JSON
// request object.
func expandJobTriggerErrorsDetailsSlice(c *Client, f []JobTriggerErrorsDetails, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerErrorsDetails(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerErrorsDetailsMap flattens the contents of JobTriggerErrorsDetails from a JSON
// response object.
func flattenJobTriggerErrorsDetailsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerErrorsDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerErrorsDetails{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerErrorsDetails{}
	}

	items := make(map[string]JobTriggerErrorsDetails)
	for k, item := range a {
		items[k] = *flattenJobTriggerErrorsDetails(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerErrorsDetailsSlice flattens the contents of JobTriggerErrorsDetails from a JSON
// response object.
func flattenJobTriggerErrorsDetailsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerErrorsDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerErrorsDetails{}
	}

	if len(a) == 0 {
		return []JobTriggerErrorsDetails{}
	}

	items := make([]JobTriggerErrorsDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerErrorsDetails(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerErrorsDetails expands an instance of JobTriggerErrorsDetails into a JSON
// request object.
func expandJobTriggerErrorsDetails(c *Client, f *JobTriggerErrorsDetails, res *JobTrigger) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Code; !dcl.IsEmptyValueIndirect(v) {
		m["code"] = v
	}
	if v := f.Message; !dcl.IsEmptyValueIndirect(v) {
		m["message"] = v
	}
	if v, err := expandJobTriggerErrorsDetailsDetailsSlice(c, f.Details, res); err != nil {
		return nil, fmt.Errorf("error expanding Details into details: %w", err)
	} else if v != nil {
		m["details"] = v
	}

	return m, nil
}

// flattenJobTriggerErrorsDetails flattens an instance of JobTriggerErrorsDetails from a JSON
// response object.
func flattenJobTriggerErrorsDetails(c *Client, i interface{}, res *JobTrigger) *JobTriggerErrorsDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerErrorsDetails{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerErrorsDetails
	}
	r.Code = dcl.FlattenInteger(m["code"])
	r.Message = dcl.FlattenString(m["message"])
	r.Details = flattenJobTriggerErrorsDetailsDetailsSlice(c, m["details"], res)

	return r
}

// expandJobTriggerErrorsDetailsDetailsMap expands the contents of JobTriggerErrorsDetailsDetails into a JSON
// request object.
func expandJobTriggerErrorsDetailsDetailsMap(c *Client, f map[string]JobTriggerErrorsDetailsDetails, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandJobTriggerErrorsDetailsDetails(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandJobTriggerErrorsDetailsDetailsSlice expands the contents of JobTriggerErrorsDetailsDetails into a JSON
// request object.
func expandJobTriggerErrorsDetailsDetailsSlice(c *Client, f []JobTriggerErrorsDetailsDetails, res *JobTrigger) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandJobTriggerErrorsDetailsDetails(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenJobTriggerErrorsDetailsDetailsMap flattens the contents of JobTriggerErrorsDetailsDetails from a JSON
// response object.
func flattenJobTriggerErrorsDetailsDetailsMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerErrorsDetailsDetails {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerErrorsDetailsDetails{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerErrorsDetailsDetails{}
	}

	items := make(map[string]JobTriggerErrorsDetailsDetails)
	for k, item := range a {
		items[k] = *flattenJobTriggerErrorsDetailsDetails(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenJobTriggerErrorsDetailsDetailsSlice flattens the contents of JobTriggerErrorsDetailsDetails from a JSON
// response object.
func flattenJobTriggerErrorsDetailsDetailsSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerErrorsDetailsDetails {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerErrorsDetailsDetails{}
	}

	if len(a) == 0 {
		return []JobTriggerErrorsDetailsDetails{}
	}

	items := make([]JobTriggerErrorsDetailsDetails, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerErrorsDetailsDetails(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandJobTriggerErrorsDetailsDetails expands an instance of JobTriggerErrorsDetailsDetails into a JSON
// request object.
func expandJobTriggerErrorsDetailsDetails(c *Client, f *JobTriggerErrorsDetailsDetails, res *JobTrigger) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.TypeUrl; !dcl.IsEmptyValueIndirect(v) {
		m["typeUrl"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenJobTriggerErrorsDetailsDetails flattens an instance of JobTriggerErrorsDetailsDetails from a JSON
// response object.
func flattenJobTriggerErrorsDetailsDetails(c *Client, i interface{}, res *JobTrigger) *JobTriggerErrorsDetailsDetails {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &JobTriggerErrorsDetailsDetails{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyJobTriggerErrorsDetailsDetails
	}
	r.TypeUrl = dcl.FlattenString(m["typeUrl"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumMap flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumSlice flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum{}
	}

	items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum with the same value as that string.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(i interface{}) *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnumRef(s)
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumMap flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumSlice flattens the contents of JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum{}
	}

	items := make([]JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum with the same value as that string.
func flattenJobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum(i interface{}) *JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumRef(s)
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumMap flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum{}
	}

	items := make(map[string]JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumSlice flattens the contents of JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum{}
	}

	items := make([]JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum with the same value as that string.
func flattenJobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum(i interface{}) *JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumRef(s)
}

// flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnumMap flattens the contents of JobTriggerInspectJobInspectConfigMinLikelihoodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigMinLikelihoodEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigMinLikelihoodEnum{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigMinLikelihoodEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnumSlice flattens the contents of JobTriggerInspectJobInspectConfigMinLikelihoodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigMinLikelihoodEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigMinLikelihoodEnum{}
	}

	items := make([]JobTriggerInspectJobInspectConfigMinLikelihoodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobInspectConfigMinLikelihoodEnum with the same value as that string.
func flattenJobTriggerInspectJobInspectConfigMinLikelihoodEnum(i interface{}) *JobTriggerInspectJobInspectConfigMinLikelihoodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobInspectConfigMinLikelihoodEnumRef(s)
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum with the same value as that string.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum(i interface{}) *JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumRef(s)
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum with the same value as that string.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(i interface{}) *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s)
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumMap flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumSlice flattens the contents of JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum{}
	}

	items := make([]JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum with the same value as that string.
func flattenJobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum(i interface{}) *JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumRef(s)
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum with the same value as that string.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(i interface{}) *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s)
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumMap flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum{}
	}

	items := make(map[string]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumSlice flattens the contents of JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum from a JSON
// response object.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum{}
	}

	items := make([]JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum with the same value as that string.
func flattenJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(i interface{}) *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumRef(s)
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumMap flattens the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum{}
	}

	items := make(map[string]JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumSlice flattens the contents of JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum from a JSON
// response object.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum{}
	}

	items := make([]JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum with the same value as that string.
func flattenJobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum(i interface{}) *JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumRef(s)
}

// flattenJobTriggerStatusEnumMap flattens the contents of JobTriggerStatusEnum from a JSON
// response object.
func flattenJobTriggerStatusEnumMap(c *Client, i interface{}, res *JobTrigger) map[string]JobTriggerStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]JobTriggerStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]JobTriggerStatusEnum{}
	}

	items := make(map[string]JobTriggerStatusEnum)
	for k, item := range a {
		items[k] = *flattenJobTriggerStatusEnum(item.(interface{}))
	}

	return items
}

// flattenJobTriggerStatusEnumSlice flattens the contents of JobTriggerStatusEnum from a JSON
// response object.
func flattenJobTriggerStatusEnumSlice(c *Client, i interface{}, res *JobTrigger) []JobTriggerStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []JobTriggerStatusEnum{}
	}

	if len(a) == 0 {
		return []JobTriggerStatusEnum{}
	}

	items := make([]JobTriggerStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenJobTriggerStatusEnum(item.(interface{})))
	}

	return items
}

// flattenJobTriggerStatusEnum asserts that an interface is a string, and returns a
// pointer to a *JobTriggerStatusEnum with the same value as that string.
func flattenJobTriggerStatusEnum(i interface{}) *JobTriggerStatusEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return JobTriggerStatusEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *JobTrigger) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalJobTrigger(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Parent == nil && ncr.Parent == nil {
			c.Config.Logger.Info("Both Parent fields null - considering equal.")
		} else if nr.Parent == nil || ncr.Parent == nil {
			c.Config.Logger.Info("Only one Parent field is null - considering unequal.")
			return false
		} else if *nr.Parent != *ncr.Parent {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type jobTriggerDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         jobTriggerApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToJobTriggerDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]jobTriggerDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []jobTriggerDiff
	// For each operation name, create a jobTriggerDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := jobTriggerDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToJobTriggerApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToJobTriggerApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (jobTriggerApiOperation, error) {
	switch opName {

	case "updateJobTriggerUpdateJobTriggerOperation":
		return &updateJobTriggerUpdateJobTriggerOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractJobTriggerFields(r *JobTrigger) error {
	vInspectJob := r.InspectJob
	if vInspectJob == nil {
		// note: explicitly not the empty object.
		vInspectJob = &JobTriggerInspectJob{}
	}
	if err := extractJobTriggerInspectJobFields(r, vInspectJob); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInspectJob) {
		r.InspectJob = vInspectJob
	}
	return nil
}
func extractJobTriggerInspectJobFields(r *JobTrigger, o *JobTriggerInspectJob) error {
	vStorageConfig := o.StorageConfig
	if vStorageConfig == nil {
		// note: explicitly not the empty object.
		vStorageConfig = &JobTriggerInspectJobStorageConfig{}
	}
	if err := extractJobTriggerInspectJobStorageConfigFields(r, vStorageConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStorageConfig) {
		o.StorageConfig = vStorageConfig
	}
	vInspectConfig := o.InspectConfig
	if vInspectConfig == nil {
		// note: explicitly not the empty object.
		vInspectConfig = &JobTriggerInspectJobInspectConfig{}
	}
	if err := extractJobTriggerInspectJobInspectConfigFields(r, vInspectConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInspectConfig) {
		o.InspectConfig = vInspectConfig
	}
	return nil
}
func extractJobTriggerInspectJobStorageConfigFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfig) error {
	vDatastoreOptions := o.DatastoreOptions
	if vDatastoreOptions == nil {
		// note: explicitly not the empty object.
		vDatastoreOptions = &JobTriggerInspectJobStorageConfigDatastoreOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigDatastoreOptionsFields(r, vDatastoreOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDatastoreOptions) {
		o.DatastoreOptions = vDatastoreOptions
	}
	vCloudStorageOptions := o.CloudStorageOptions
	if vCloudStorageOptions == nil {
		// note: explicitly not the empty object.
		vCloudStorageOptions = &JobTriggerInspectJobStorageConfigCloudStorageOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFields(r, vCloudStorageOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStorageOptions) {
		o.CloudStorageOptions = vCloudStorageOptions
	}
	vBigQueryOptions := o.BigQueryOptions
	if vBigQueryOptions == nil {
		// note: explicitly not the empty object.
		vBigQueryOptions = &JobTriggerInspectJobStorageConfigBigQueryOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigBigQueryOptionsFields(r, vBigQueryOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBigQueryOptions) {
		o.BigQueryOptions = vBigQueryOptions
	}
	vHybridOptions := o.HybridOptions
	if vHybridOptions == nil {
		// note: explicitly not the empty object.
		vHybridOptions = &JobTriggerInspectJobStorageConfigHybridOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigHybridOptionsFields(r, vHybridOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHybridOptions) {
		o.HybridOptions = vHybridOptions
	}
	vTimespanConfig := o.TimespanConfig
	if vTimespanConfig == nil {
		// note: explicitly not the empty object.
		vTimespanConfig = &JobTriggerInspectJobStorageConfigTimespanConfig{}
	}
	if err := extractJobTriggerInspectJobStorageConfigTimespanConfigFields(r, vTimespanConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTimespanConfig) {
		o.TimespanConfig = vTimespanConfig
	}
	return nil
}
func extractJobTriggerInspectJobStorageConfigDatastoreOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigDatastoreOptions) error {
	vPartitionId := o.PartitionId
	if vPartitionId == nil {
		// note: explicitly not the empty object.
		vPartitionId = &JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	}
	if err := extractJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdFields(r, vPartitionId); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPartitionId) {
		o.PartitionId = vPartitionId
	}
	vKind := o.Kind
	if vKind == nil {
		// note: explicitly not the empty object.
		vKind = &JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	}
	if err := extractJobTriggerInspectJobStorageConfigDatastoreOptionsKindFields(r, vKind); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKind) {
		o.Kind = vKind
	}
	return nil
}
func extractJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigDatastoreOptionsKindFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigDatastoreOptionsKind) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigCloudStorageOptions) error {
	vFileSet := o.FileSet
	if vFileSet == nil {
		// note: explicitly not the empty object.
		vFileSet = &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	}
	if err := extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetFields(r, vFileSet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFileSet) {
		o.FileSet = vFileSet
	}
	return nil
}
func extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) error {
	vRegexFileSet := o.RegexFileSet
	if vRegexFileSet == nil {
		// note: explicitly not the empty object.
		vRegexFileSet = &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
	}
	if err := extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetFields(r, vRegexFileSet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegexFileSet) {
		o.RegexFileSet = vRegexFileSet
	}
	return nil
}
func extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigBigQueryOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptions) error {
	vTableReference := o.TableReference
	if vTableReference == nil {
		// note: explicitly not the empty object.
		vTableReference = &JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	}
	if err := extractJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceFields(r, vTableReference); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTableReference) {
		o.TableReference = vTableReference
	}
	return nil
}
func extractJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigHybridOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigHybridOptions) error {
	vTableOptions := o.TableOptions
	if vTableOptions == nil {
		// note: explicitly not the empty object.
		vTableOptions = &JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsFields(r, vTableOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTableOptions) {
		o.TableOptions = vTableOptions
	}
	return nil
}
func extractJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) error {
	return nil
}
func extractJobTriggerInspectJobStorageConfigTimespanConfigFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigTimespanConfig) error {
	vTimestampField := o.TimestampField
	if vTimestampField == nil {
		// note: explicitly not the empty object.
		vTimestampField = &JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	}
	if err := extractJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldFields(r, vTimestampField); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTimestampField) {
		o.TimestampField = vTimestampField
	}
	return nil
}
func extractJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfig) error {
	vLimits := o.Limits
	if vLimits == nil {
		// note: explicitly not the empty object.
		vLimits = &JobTriggerInspectJobInspectConfigLimits{}
	}
	if err := extractJobTriggerInspectJobInspectConfigLimitsFields(r, vLimits); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLimits) {
		o.Limits = vLimits
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigInfoTypes) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigLimitsFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigLimits) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) error {
	vInfoType := o.InfoType
	if vInfoType == nil {
		// note: explicitly not the empty object.
		vInfoType = &JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}
	if err := extractJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeFields(r, vInfoType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInfoType) {
		o.InfoType = vInfoType
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypes) error {
	vInfoType := o.InfoType
	if vInfoType == nil {
		// note: explicitly not the empty object.
		vInfoType = &JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeFields(r, vInfoType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInfoType) {
		o.InfoType = vInfoType
	}
	vDictionary := o.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		o.Dictionary = vDictionary
	}
	vRegex := o.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		o.Regex = vRegex
	}
	vSurrogateType := o.SurrogateType
	if vSurrogateType == nil {
		// note: explicitly not the empty object.
		vSurrogateType = &JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeFields(r, vSurrogateType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSurrogateType) {
		o.SurrogateType = vSurrogateType
	}
	vStoredType := o.StoredType
	if vStoredType == nil {
		// note: explicitly not the empty object.
		vStoredType = &JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeFields(r, vStoredType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStoredType) {
		o.StoredType = vStoredType
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesRegexFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) error {
	vHotwordRule := o.HotwordRule
	if vHotwordRule == nil {
		// note: explicitly not the empty object.
		vHotwordRule = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleFields(r, vHotwordRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRule) {
		o.HotwordRule = vHotwordRule
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) error {
	vHotwordRegex := o.HotwordRegex
	if vHotwordRegex == nil {
		// note: explicitly not the empty object.
		vHotwordRegex = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexFields(r, vHotwordRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRegex) {
		o.HotwordRegex = vHotwordRegex
	}
	vProximity := o.Proximity
	if vProximity == nil {
		// note: explicitly not the empty object.
		vProximity = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityFields(r, vProximity); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vProximity) {
		o.Proximity = vProximity
	}
	vLikelihoodAdjustment := o.LikelihoodAdjustment
	if vLikelihoodAdjustment == nil {
		// note: explicitly not the empty object.
		vLikelihoodAdjustment = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFields(r, vLikelihoodAdjustment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLikelihoodAdjustment) {
		o.LikelihoodAdjustment = vLikelihoodAdjustment
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSet) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetInfoTypes) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRules) error {
	vHotwordRule := o.HotwordRule
	if vHotwordRule == nil {
		// note: explicitly not the empty object.
		vHotwordRule = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleFields(r, vHotwordRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRule) {
		o.HotwordRule = vHotwordRule
	}
	vExclusionRule := o.ExclusionRule
	if vExclusionRule == nil {
		// note: explicitly not the empty object.
		vExclusionRule = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleFields(r, vExclusionRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExclusionRule) {
		o.ExclusionRule = vExclusionRule
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) error {
	vHotwordRegex := o.HotwordRegex
	if vHotwordRegex == nil {
		// note: explicitly not the empty object.
		vHotwordRegex = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexFields(r, vHotwordRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRegex) {
		o.HotwordRegex = vHotwordRegex
	}
	vProximity := o.Proximity
	if vProximity == nil {
		// note: explicitly not the empty object.
		vProximity = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityFields(r, vProximity); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vProximity) {
		o.Proximity = vProximity
	}
	vLikelihoodAdjustment := o.LikelihoodAdjustment
	if vLikelihoodAdjustment == nil {
		// note: explicitly not the empty object.
		vLikelihoodAdjustment = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFields(r, vLikelihoodAdjustment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLikelihoodAdjustment) {
		o.LikelihoodAdjustment = vLikelihoodAdjustment
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) error {
	vDictionary := o.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		o.Dictionary = vDictionary
	}
	vRegex := o.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		o.Regex = vRegex
	}
	vExcludeInfoTypes := o.ExcludeInfoTypes
	if vExcludeInfoTypes == nil {
		// note: explicitly not the empty object.
		vExcludeInfoTypes = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesFields(r, vExcludeInfoTypes); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExcludeInfoTypes) {
		o.ExcludeInfoTypes = vExcludeInfoTypes
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) error {
	return nil
}
func extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) error {
	return nil
}
func extractJobTriggerInspectJobActionsFields(r *JobTrigger, o *JobTriggerInspectJobActions) error {
	vSaveFindings := o.SaveFindings
	if vSaveFindings == nil {
		// note: explicitly not the empty object.
		vSaveFindings = &JobTriggerInspectJobActionsSaveFindings{}
	}
	if err := extractJobTriggerInspectJobActionsSaveFindingsFields(r, vSaveFindings); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSaveFindings) {
		o.SaveFindings = vSaveFindings
	}
	vPubSub := o.PubSub
	if vPubSub == nil {
		// note: explicitly not the empty object.
		vPubSub = &JobTriggerInspectJobActionsPubSub{}
	}
	if err := extractJobTriggerInspectJobActionsPubSubFields(r, vPubSub); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPubSub) {
		o.PubSub = vPubSub
	}
	vPublishSummaryToCscc := o.PublishSummaryToCscc
	if vPublishSummaryToCscc == nil {
		// note: explicitly not the empty object.
		vPublishSummaryToCscc = &JobTriggerInspectJobActionsPublishSummaryToCscc{}
	}
	if err := extractJobTriggerInspectJobActionsPublishSummaryToCsccFields(r, vPublishSummaryToCscc); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPublishSummaryToCscc) {
		o.PublishSummaryToCscc = vPublishSummaryToCscc
	}
	vPublishFindingsToCloudDataCatalog := o.PublishFindingsToCloudDataCatalog
	if vPublishFindingsToCloudDataCatalog == nil {
		// note: explicitly not the empty object.
		vPublishFindingsToCloudDataCatalog = &JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	}
	if err := extractJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogFields(r, vPublishFindingsToCloudDataCatalog); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPublishFindingsToCloudDataCatalog) {
		o.PublishFindingsToCloudDataCatalog = vPublishFindingsToCloudDataCatalog
	}
	vJobNotificationEmails := o.JobNotificationEmails
	if vJobNotificationEmails == nil {
		// note: explicitly not the empty object.
		vJobNotificationEmails = &JobTriggerInspectJobActionsJobNotificationEmails{}
	}
	if err := extractJobTriggerInspectJobActionsJobNotificationEmailsFields(r, vJobNotificationEmails); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vJobNotificationEmails) {
		o.JobNotificationEmails = vJobNotificationEmails
	}
	vPublishToStackdriver := o.PublishToStackdriver
	if vPublishToStackdriver == nil {
		// note: explicitly not the empty object.
		vPublishToStackdriver = &JobTriggerInspectJobActionsPublishToStackdriver{}
	}
	if err := extractJobTriggerInspectJobActionsPublishToStackdriverFields(r, vPublishToStackdriver); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPublishToStackdriver) {
		o.PublishToStackdriver = vPublishToStackdriver
	}
	return nil
}
func extractJobTriggerInspectJobActionsSaveFindingsFields(r *JobTrigger, o *JobTriggerInspectJobActionsSaveFindings) error {
	vOutputConfig := o.OutputConfig
	if vOutputConfig == nil {
		// note: explicitly not the empty object.
		vOutputConfig = &JobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	}
	if err := extractJobTriggerInspectJobActionsSaveFindingsOutputConfigFields(r, vOutputConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOutputConfig) {
		o.OutputConfig = vOutputConfig
	}
	return nil
}
func extractJobTriggerInspectJobActionsSaveFindingsOutputConfigFields(r *JobTrigger, o *JobTriggerInspectJobActionsSaveFindingsOutputConfig) error {
	vTable := o.Table
	if vTable == nil {
		// note: explicitly not the empty object.
		vTable = &JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	}
	if err := extractJobTriggerInspectJobActionsSaveFindingsOutputConfigTableFields(r, vTable); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTable) {
		o.Table = vTable
	}
	vDlpStorage := o.DlpStorage
	if vDlpStorage == nil {
		// note: explicitly not the empty object.
		vDlpStorage = &JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	}
	if err := extractJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageFields(r, vDlpStorage); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDlpStorage) {
		o.DlpStorage = vDlpStorage
	}
	return nil
}
func extractJobTriggerInspectJobActionsSaveFindingsOutputConfigTableFields(r *JobTrigger, o *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) error {
	return nil
}
func extractJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageFields(r *JobTrigger, o *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) error {
	return nil
}
func extractJobTriggerInspectJobActionsPubSubFields(r *JobTrigger, o *JobTriggerInspectJobActionsPubSub) error {
	return nil
}
func extractJobTriggerInspectJobActionsPublishSummaryToCsccFields(r *JobTrigger, o *JobTriggerInspectJobActionsPublishSummaryToCscc) error {
	return nil
}
func extractJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogFields(r *JobTrigger, o *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) error {
	return nil
}
func extractJobTriggerInspectJobActionsJobNotificationEmailsFields(r *JobTrigger, o *JobTriggerInspectJobActionsJobNotificationEmails) error {
	return nil
}
func extractJobTriggerInspectJobActionsPublishToStackdriverFields(r *JobTrigger, o *JobTriggerInspectJobActionsPublishToStackdriver) error {
	return nil
}
func extractJobTriggerTriggersFields(r *JobTrigger, o *JobTriggerTriggers) error {
	vSchedule := o.Schedule
	if vSchedule == nil {
		// note: explicitly not the empty object.
		vSchedule = &JobTriggerTriggersSchedule{}
	}
	if err := extractJobTriggerTriggersScheduleFields(r, vSchedule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSchedule) {
		o.Schedule = vSchedule
	}
	vManual := o.Manual
	if vManual == nil {
		// note: explicitly not the empty object.
		vManual = &JobTriggerTriggersManual{}
	}
	if err := extractJobTriggerTriggersManualFields(r, vManual); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vManual) {
		o.Manual = vManual
	}
	return nil
}
func extractJobTriggerTriggersScheduleFields(r *JobTrigger, o *JobTriggerTriggersSchedule) error {
	return nil
}
func extractJobTriggerTriggersManualFields(r *JobTrigger, o *JobTriggerTriggersManual) error {
	return nil
}
func extractJobTriggerErrorsFields(r *JobTrigger, o *JobTriggerErrors) error {
	vDetails := o.Details
	if vDetails == nil {
		// note: explicitly not the empty object.
		vDetails = &JobTriggerErrorsDetails{}
	}
	if err := extractJobTriggerErrorsDetailsFields(r, vDetails); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDetails) {
		o.Details = vDetails
	}
	return nil
}
func extractJobTriggerErrorsDetailsFields(r *JobTrigger, o *JobTriggerErrorsDetails) error {
	return nil
}
func extractJobTriggerErrorsDetailsDetailsFields(r *JobTrigger, o *JobTriggerErrorsDetailsDetails) error {
	return nil
}

func postReadExtractJobTriggerFields(r *JobTrigger) error {
	vInspectJob := r.InspectJob
	if vInspectJob == nil {
		// note: explicitly not the empty object.
		vInspectJob = &JobTriggerInspectJob{}
	}
	if err := postReadExtractJobTriggerInspectJobFields(r, vInspectJob); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInspectJob) {
		r.InspectJob = vInspectJob
	}
	return nil
}
func postReadExtractJobTriggerInspectJobFields(r *JobTrigger, o *JobTriggerInspectJob) error {
	vStorageConfig := o.StorageConfig
	if vStorageConfig == nil {
		// note: explicitly not the empty object.
		vStorageConfig = &JobTriggerInspectJobStorageConfig{}
	}
	if err := extractJobTriggerInspectJobStorageConfigFields(r, vStorageConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStorageConfig) {
		o.StorageConfig = vStorageConfig
	}
	vInspectConfig := o.InspectConfig
	if vInspectConfig == nil {
		// note: explicitly not the empty object.
		vInspectConfig = &JobTriggerInspectJobInspectConfig{}
	}
	if err := extractJobTriggerInspectJobInspectConfigFields(r, vInspectConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInspectConfig) {
		o.InspectConfig = vInspectConfig
	}
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfig) error {
	vDatastoreOptions := o.DatastoreOptions
	if vDatastoreOptions == nil {
		// note: explicitly not the empty object.
		vDatastoreOptions = &JobTriggerInspectJobStorageConfigDatastoreOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigDatastoreOptionsFields(r, vDatastoreOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDatastoreOptions) {
		o.DatastoreOptions = vDatastoreOptions
	}
	vCloudStorageOptions := o.CloudStorageOptions
	if vCloudStorageOptions == nil {
		// note: explicitly not the empty object.
		vCloudStorageOptions = &JobTriggerInspectJobStorageConfigCloudStorageOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFields(r, vCloudStorageOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStorageOptions) {
		o.CloudStorageOptions = vCloudStorageOptions
	}
	vBigQueryOptions := o.BigQueryOptions
	if vBigQueryOptions == nil {
		// note: explicitly not the empty object.
		vBigQueryOptions = &JobTriggerInspectJobStorageConfigBigQueryOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigBigQueryOptionsFields(r, vBigQueryOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBigQueryOptions) {
		o.BigQueryOptions = vBigQueryOptions
	}
	vHybridOptions := o.HybridOptions
	if vHybridOptions == nil {
		// note: explicitly not the empty object.
		vHybridOptions = &JobTriggerInspectJobStorageConfigHybridOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigHybridOptionsFields(r, vHybridOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHybridOptions) {
		o.HybridOptions = vHybridOptions
	}
	vTimespanConfig := o.TimespanConfig
	if vTimespanConfig == nil {
		// note: explicitly not the empty object.
		vTimespanConfig = &JobTriggerInspectJobStorageConfigTimespanConfig{}
	}
	if err := extractJobTriggerInspectJobStorageConfigTimespanConfigFields(r, vTimespanConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTimespanConfig) {
		o.TimespanConfig = vTimespanConfig
	}
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigDatastoreOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigDatastoreOptions) error {
	vPartitionId := o.PartitionId
	if vPartitionId == nil {
		// note: explicitly not the empty object.
		vPartitionId = &JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
	}
	if err := extractJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdFields(r, vPartitionId); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPartitionId) {
		o.PartitionId = vPartitionId
	}
	vKind := o.Kind
	if vKind == nil {
		// note: explicitly not the empty object.
		vKind = &JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
	}
	if err := extractJobTriggerInspectJobStorageConfigDatastoreOptionsKindFields(r, vKind); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKind) {
		o.Kind = vKind
	}
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigDatastoreOptionsKindFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigDatastoreOptionsKind) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigCloudStorageOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigCloudStorageOptions) error {
	vFileSet := o.FileSet
	if vFileSet == nil {
		// note: explicitly not the empty object.
		vFileSet = &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
	}
	if err := extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetFields(r, vFileSet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vFileSet) {
		o.FileSet = vFileSet
	}
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) error {
	vRegexFileSet := o.RegexFileSet
	if vRegexFileSet == nil {
		// note: explicitly not the empty object.
		vRegexFileSet = &JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
	}
	if err := extractJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetFields(r, vRegexFileSet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegexFileSet) {
		o.RegexFileSet = vRegexFileSet
	}
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigBigQueryOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptions) error {
	vTableReference := o.TableReference
	if vTableReference == nil {
		// note: explicitly not the empty object.
		vTableReference = &JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
	}
	if err := extractJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceFields(r, vTableReference); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTableReference) {
		o.TableReference = vTableReference
	}
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigHybridOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigHybridOptions) error {
	vTableOptions := o.TableOptions
	if vTableOptions == nil {
		// note: explicitly not the empty object.
		vTableOptions = &JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
	}
	if err := extractJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsFields(r, vTableOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTableOptions) {
		o.TableOptions = vTableOptions
	}
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigHybridOptionsTableOptions) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields) error {
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigTimespanConfigFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigTimespanConfig) error {
	vTimestampField := o.TimestampField
	if vTimestampField == nil {
		// note: explicitly not the empty object.
		vTimestampField = &JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
	}
	if err := extractJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldFields(r, vTimestampField); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTimestampField) {
		o.TimestampField = vTimestampField
	}
	return nil
}
func postReadExtractJobTriggerInspectJobStorageConfigTimespanConfigTimestampFieldFields(r *JobTrigger, o *JobTriggerInspectJobStorageConfigTimespanConfigTimestampField) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfig) error {
	vLimits := o.Limits
	if vLimits == nil {
		// note: explicitly not the empty object.
		vLimits = &JobTriggerInspectJobInspectConfigLimits{}
	}
	if err := extractJobTriggerInspectJobInspectConfigLimitsFields(r, vLimits); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLimits) {
		o.Limits = vLimits
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigInfoTypes) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigLimitsFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigLimits) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType) error {
	vInfoType := o.InfoType
	if vInfoType == nil {
		// note: explicitly not the empty object.
		vInfoType = &JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}
	if err := extractJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeFields(r, vInfoType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInfoType) {
		o.InfoType = vInfoType
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypes) error {
	vInfoType := o.InfoType
	if vInfoType == nil {
		// note: explicitly not the empty object.
		vInfoType = &JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeFields(r, vInfoType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInfoType) {
		o.InfoType = vInfoType
	}
	vDictionary := o.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		o.Dictionary = vDictionary
	}
	vRegex := o.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		o.Regex = vRegex
	}
	vSurrogateType := o.SurrogateType
	if vSurrogateType == nil {
		// note: explicitly not the empty object.
		vSurrogateType = &JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeFields(r, vSurrogateType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSurrogateType) {
		o.SurrogateType = vSurrogateType
	}
	vStoredType := o.StoredType
	if vStoredType == nil {
		// note: explicitly not the empty object.
		vStoredType = &JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeFields(r, vStoredType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStoredType) {
		o.StoredType = vStoredType
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesInfoTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordListFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePathFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesRegexFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesRegex) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesStoredTypeFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules) error {
	vHotwordRule := o.HotwordRule
	if vHotwordRule == nil {
		// note: explicitly not the empty object.
		vHotwordRule = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleFields(r, vHotwordRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRule) {
		o.HotwordRule = vHotwordRule
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule) error {
	vHotwordRegex := o.HotwordRegex
	if vHotwordRegex == nil {
		// note: explicitly not the empty object.
		vHotwordRegex = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexFields(r, vHotwordRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRegex) {
		o.HotwordRegex = vHotwordRegex
	}
	vProximity := o.Proximity
	if vProximity == nil {
		// note: explicitly not the empty object.
		vProximity = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityFields(r, vProximity); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vProximity) {
		o.Proximity = vProximity
	}
	vLikelihoodAdjustment := o.LikelihoodAdjustment
	if vLikelihoodAdjustment == nil {
		// note: explicitly not the empty object.
		vLikelihoodAdjustment = &JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
	}
	if err := extractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFields(r, vLikelihoodAdjustment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLikelihoodAdjustment) {
		o.LikelihoodAdjustment = vLikelihoodAdjustment
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegexFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximityFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSet) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetInfoTypes) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRules) error {
	vHotwordRule := o.HotwordRule
	if vHotwordRule == nil {
		// note: explicitly not the empty object.
		vHotwordRule = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleFields(r, vHotwordRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRule) {
		o.HotwordRule = vHotwordRule
	}
	vExclusionRule := o.ExclusionRule
	if vExclusionRule == nil {
		// note: explicitly not the empty object.
		vExclusionRule = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleFields(r, vExclusionRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExclusionRule) {
		o.ExclusionRule = vExclusionRule
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) error {
	vHotwordRegex := o.HotwordRegex
	if vHotwordRegex == nil {
		// note: explicitly not the empty object.
		vHotwordRegex = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexFields(r, vHotwordRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRegex) {
		o.HotwordRegex = vHotwordRegex
	}
	vProximity := o.Proximity
	if vProximity == nil {
		// note: explicitly not the empty object.
		vProximity = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityFields(r, vProximity); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vProximity) {
		o.Proximity = vProximity
	}
	vLikelihoodAdjustment := o.LikelihoodAdjustment
	if vLikelihoodAdjustment == nil {
		// note: explicitly not the empty object.
		vLikelihoodAdjustment = &JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFields(r, vLikelihoodAdjustment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLikelihoodAdjustment) {
		o.LikelihoodAdjustment = vLikelihoodAdjustment
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegexFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximityFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) error {
	vDictionary := o.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		o.Dictionary = vDictionary
	}
	vRegex := o.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		o.Regex = vRegex
	}
	vExcludeInfoTypes := o.ExcludeInfoTypes
	if vExcludeInfoTypes == nil {
		// note: explicitly not the empty object.
		vExcludeInfoTypes = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesFields(r, vExcludeInfoTypes); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExcludeInfoTypes) {
		o.ExcludeInfoTypes = vExcludeInfoTypes
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}
	if err := extractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) error {
	return nil
}
func postReadExtractJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesFields(r *JobTrigger, o *JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) error {
	return nil
}
func postReadExtractJobTriggerInspectJobActionsFields(r *JobTrigger, o *JobTriggerInspectJobActions) error {
	vSaveFindings := o.SaveFindings
	if vSaveFindings == nil {
		// note: explicitly not the empty object.
		vSaveFindings = &JobTriggerInspectJobActionsSaveFindings{}
	}
	if err := extractJobTriggerInspectJobActionsSaveFindingsFields(r, vSaveFindings); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSaveFindings) {
		o.SaveFindings = vSaveFindings
	}
	vPubSub := o.PubSub
	if vPubSub == nil {
		// note: explicitly not the empty object.
		vPubSub = &JobTriggerInspectJobActionsPubSub{}
	}
	if err := extractJobTriggerInspectJobActionsPubSubFields(r, vPubSub); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPubSub) {
		o.PubSub = vPubSub
	}
	vPublishSummaryToCscc := o.PublishSummaryToCscc
	if vPublishSummaryToCscc == nil {
		// note: explicitly not the empty object.
		vPublishSummaryToCscc = &JobTriggerInspectJobActionsPublishSummaryToCscc{}
	}
	if err := extractJobTriggerInspectJobActionsPublishSummaryToCsccFields(r, vPublishSummaryToCscc); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPublishSummaryToCscc) {
		o.PublishSummaryToCscc = vPublishSummaryToCscc
	}
	vPublishFindingsToCloudDataCatalog := o.PublishFindingsToCloudDataCatalog
	if vPublishFindingsToCloudDataCatalog == nil {
		// note: explicitly not the empty object.
		vPublishFindingsToCloudDataCatalog = &JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
	}
	if err := extractJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogFields(r, vPublishFindingsToCloudDataCatalog); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPublishFindingsToCloudDataCatalog) {
		o.PublishFindingsToCloudDataCatalog = vPublishFindingsToCloudDataCatalog
	}
	vJobNotificationEmails := o.JobNotificationEmails
	if vJobNotificationEmails == nil {
		// note: explicitly not the empty object.
		vJobNotificationEmails = &JobTriggerInspectJobActionsJobNotificationEmails{}
	}
	if err := extractJobTriggerInspectJobActionsJobNotificationEmailsFields(r, vJobNotificationEmails); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vJobNotificationEmails) {
		o.JobNotificationEmails = vJobNotificationEmails
	}
	vPublishToStackdriver := o.PublishToStackdriver
	if vPublishToStackdriver == nil {
		// note: explicitly not the empty object.
		vPublishToStackdriver = &JobTriggerInspectJobActionsPublishToStackdriver{}
	}
	if err := extractJobTriggerInspectJobActionsPublishToStackdriverFields(r, vPublishToStackdriver); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPublishToStackdriver) {
		o.PublishToStackdriver = vPublishToStackdriver
	}
	return nil
}
func postReadExtractJobTriggerInspectJobActionsSaveFindingsFields(r *JobTrigger, o *JobTriggerInspectJobActionsSaveFindings) error {
	vOutputConfig := o.OutputConfig
	if vOutputConfig == nil {
		// note: explicitly not the empty object.
		vOutputConfig = &JobTriggerInspectJobActionsSaveFindingsOutputConfig{}
	}
	if err := extractJobTriggerInspectJobActionsSaveFindingsOutputConfigFields(r, vOutputConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOutputConfig) {
		o.OutputConfig = vOutputConfig
	}
	return nil
}
func postReadExtractJobTriggerInspectJobActionsSaveFindingsOutputConfigFields(r *JobTrigger, o *JobTriggerInspectJobActionsSaveFindingsOutputConfig) error {
	vTable := o.Table
	if vTable == nil {
		// note: explicitly not the empty object.
		vTable = &JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
	}
	if err := extractJobTriggerInspectJobActionsSaveFindingsOutputConfigTableFields(r, vTable); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTable) {
		o.Table = vTable
	}
	vDlpStorage := o.DlpStorage
	if vDlpStorage == nil {
		// note: explicitly not the empty object.
		vDlpStorage = &JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
	}
	if err := extractJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageFields(r, vDlpStorage); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDlpStorage) {
		o.DlpStorage = vDlpStorage
	}
	return nil
}
func postReadExtractJobTriggerInspectJobActionsSaveFindingsOutputConfigTableFields(r *JobTrigger, o *JobTriggerInspectJobActionsSaveFindingsOutputConfigTable) error {
	return nil
}
func postReadExtractJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorageFields(r *JobTrigger, o *JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage) error {
	return nil
}
func postReadExtractJobTriggerInspectJobActionsPubSubFields(r *JobTrigger, o *JobTriggerInspectJobActionsPubSub) error {
	return nil
}
func postReadExtractJobTriggerInspectJobActionsPublishSummaryToCsccFields(r *JobTrigger, o *JobTriggerInspectJobActionsPublishSummaryToCscc) error {
	return nil
}
func postReadExtractJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogFields(r *JobTrigger, o *JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) error {
	return nil
}
func postReadExtractJobTriggerInspectJobActionsJobNotificationEmailsFields(r *JobTrigger, o *JobTriggerInspectJobActionsJobNotificationEmails) error {
	return nil
}
func postReadExtractJobTriggerInspectJobActionsPublishToStackdriverFields(r *JobTrigger, o *JobTriggerInspectJobActionsPublishToStackdriver) error {
	return nil
}
func postReadExtractJobTriggerTriggersFields(r *JobTrigger, o *JobTriggerTriggers) error {
	vSchedule := o.Schedule
	if vSchedule == nil {
		// note: explicitly not the empty object.
		vSchedule = &JobTriggerTriggersSchedule{}
	}
	if err := extractJobTriggerTriggersScheduleFields(r, vSchedule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSchedule) {
		o.Schedule = vSchedule
	}
	vManual := o.Manual
	if vManual == nil {
		// note: explicitly not the empty object.
		vManual = &JobTriggerTriggersManual{}
	}
	if err := extractJobTriggerTriggersManualFields(r, vManual); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vManual) {
		o.Manual = vManual
	}
	return nil
}
func postReadExtractJobTriggerTriggersScheduleFields(r *JobTrigger, o *JobTriggerTriggersSchedule) error {
	return nil
}
func postReadExtractJobTriggerTriggersManualFields(r *JobTrigger, o *JobTriggerTriggersManual) error {
	return nil
}
func postReadExtractJobTriggerErrorsFields(r *JobTrigger, o *JobTriggerErrors) error {
	vDetails := o.Details
	if vDetails == nil {
		// note: explicitly not the empty object.
		vDetails = &JobTriggerErrorsDetails{}
	}
	if err := extractJobTriggerErrorsDetailsFields(r, vDetails); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDetails) {
		o.Details = vDetails
	}
	return nil
}
func postReadExtractJobTriggerErrorsDetailsFields(r *JobTrigger, o *JobTriggerErrorsDetails) error {
	return nil
}
func postReadExtractJobTriggerErrorsDetailsDetailsFields(r *JobTrigger, o *JobTriggerErrorsDetailsDetails) error {
	return nil
}
