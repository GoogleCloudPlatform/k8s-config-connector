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
package dlp

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type JobTrigger struct{}

func JobTriggerToUnstructured(r *dclService.JobTrigger) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dlp",
			Version: "alpha",
			Type:    "JobTrigger",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	var rErrors []interface{}
	for _, rErrorsVal := range r.Errors {
		rErrorsObject := make(map[string]interface{})
		if rErrorsVal.Details != nil && rErrorsVal.Details != dclService.EmptyJobTriggerErrorsDetails {
			rErrorsValDetails := make(map[string]interface{})
			if rErrorsVal.Details.Code != nil {
				rErrorsValDetails["code"] = *rErrorsVal.Details.Code
			}
			var rErrorsValDetailsDetails []interface{}
			for _, rErrorsValDetailsDetailsVal := range rErrorsVal.Details.Details {
				rErrorsValDetailsDetailsObject := make(map[string]interface{})
				if rErrorsValDetailsDetailsVal.TypeUrl != nil {
					rErrorsValDetailsDetailsObject["typeUrl"] = *rErrorsValDetailsDetailsVal.TypeUrl
				}
				if rErrorsValDetailsDetailsVal.Value != nil {
					rErrorsValDetailsDetailsObject["value"] = *rErrorsValDetailsDetailsVal.Value
				}
				rErrorsValDetailsDetails = append(rErrorsValDetailsDetails, rErrorsValDetailsDetailsObject)
			}
			rErrorsValDetails["details"] = rErrorsValDetailsDetails
			if rErrorsVal.Details.Message != nil {
				rErrorsValDetails["message"] = *rErrorsVal.Details.Message
			}
			rErrorsObject["details"] = rErrorsValDetails
		}
		var rErrorsValTimestamps []interface{}
		for _, rErrorsValTimestampsVal := range rErrorsVal.Timestamps {
			rErrorsValTimestamps = append(rErrorsValTimestamps, rErrorsValTimestampsVal)
		}
		rErrorsObject["timestamps"] = rErrorsValTimestamps
		rErrors = append(rErrors, rErrorsObject)
	}
	u.Object["errors"] = rErrors
	if r.InspectJob != nil && r.InspectJob != dclService.EmptyJobTriggerInspectJob {
		rInspectJob := make(map[string]interface{})
		var rInspectJobActions []interface{}
		for _, rInspectJobActionsVal := range r.InspectJob.Actions {
			rInspectJobActionsObject := make(map[string]interface{})
			if rInspectJobActionsVal.JobNotificationEmails != nil && rInspectJobActionsVal.JobNotificationEmails != dclService.EmptyJobTriggerInspectJobActionsJobNotificationEmails {
				rInspectJobActionsValJobNotificationEmails := make(map[string]interface{})
				rInspectJobActionsObject["jobNotificationEmails"] = rInspectJobActionsValJobNotificationEmails
			}
			if rInspectJobActionsVal.PubSub != nil && rInspectJobActionsVal.PubSub != dclService.EmptyJobTriggerInspectJobActionsPubSub {
				rInspectJobActionsValPubSub := make(map[string]interface{})
				if rInspectJobActionsVal.PubSub.Topic != nil {
					rInspectJobActionsValPubSub["topic"] = *rInspectJobActionsVal.PubSub.Topic
				}
				rInspectJobActionsObject["pubSub"] = rInspectJobActionsValPubSub
			}
			if rInspectJobActionsVal.PublishFindingsToCloudDataCatalog != nil && rInspectJobActionsVal.PublishFindingsToCloudDataCatalog != dclService.EmptyJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
				rInspectJobActionsValPublishFindingsToCloudDataCatalog := make(map[string]interface{})
				rInspectJobActionsObject["publishFindingsToCloudDataCatalog"] = rInspectJobActionsValPublishFindingsToCloudDataCatalog
			}
			if rInspectJobActionsVal.PublishSummaryToCscc != nil && rInspectJobActionsVal.PublishSummaryToCscc != dclService.EmptyJobTriggerInspectJobActionsPublishSummaryToCscc {
				rInspectJobActionsValPublishSummaryToCscc := make(map[string]interface{})
				rInspectJobActionsObject["publishSummaryToCscc"] = rInspectJobActionsValPublishSummaryToCscc
			}
			if rInspectJobActionsVal.PublishToStackdriver != nil && rInspectJobActionsVal.PublishToStackdriver != dclService.EmptyJobTriggerInspectJobActionsPublishToStackdriver {
				rInspectJobActionsValPublishToStackdriver := make(map[string]interface{})
				rInspectJobActionsObject["publishToStackdriver"] = rInspectJobActionsValPublishToStackdriver
			}
			if rInspectJobActionsVal.SaveFindings != nil && rInspectJobActionsVal.SaveFindings != dclService.EmptyJobTriggerInspectJobActionsSaveFindings {
				rInspectJobActionsValSaveFindings := make(map[string]interface{})
				if rInspectJobActionsVal.SaveFindings.OutputConfig != nil && rInspectJobActionsVal.SaveFindings.OutputConfig != dclService.EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfig {
					rInspectJobActionsValSaveFindingsOutputConfig := make(map[string]interface{})
					if rInspectJobActionsVal.SaveFindings.OutputConfig.DlpStorage != nil && rInspectJobActionsVal.SaveFindings.OutputConfig.DlpStorage != dclService.EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage {
						rInspectJobActionsValSaveFindingsOutputConfigDlpStorage := make(map[string]interface{})
						rInspectJobActionsValSaveFindingsOutputConfig["dlpStorage"] = rInspectJobActionsValSaveFindingsOutputConfigDlpStorage
					}
					if rInspectJobActionsVal.SaveFindings.OutputConfig.OutputSchema != nil {
						rInspectJobActionsValSaveFindingsOutputConfig["outputSchema"] = string(*rInspectJobActionsVal.SaveFindings.OutputConfig.OutputSchema)
					}
					if rInspectJobActionsVal.SaveFindings.OutputConfig.Table != nil && rInspectJobActionsVal.SaveFindings.OutputConfig.Table != dclService.EmptyJobTriggerInspectJobActionsSaveFindingsOutputConfigTable {
						rInspectJobActionsValSaveFindingsOutputConfigTable := make(map[string]interface{})
						if rInspectJobActionsVal.SaveFindings.OutputConfig.Table.DatasetId != nil {
							rInspectJobActionsValSaveFindingsOutputConfigTable["datasetId"] = *rInspectJobActionsVal.SaveFindings.OutputConfig.Table.DatasetId
						}
						if rInspectJobActionsVal.SaveFindings.OutputConfig.Table.ProjectId != nil {
							rInspectJobActionsValSaveFindingsOutputConfigTable["projectId"] = *rInspectJobActionsVal.SaveFindings.OutputConfig.Table.ProjectId
						}
						if rInspectJobActionsVal.SaveFindings.OutputConfig.Table.TableId != nil {
							rInspectJobActionsValSaveFindingsOutputConfigTable["tableId"] = *rInspectJobActionsVal.SaveFindings.OutputConfig.Table.TableId
						}
						rInspectJobActionsValSaveFindingsOutputConfig["table"] = rInspectJobActionsValSaveFindingsOutputConfigTable
					}
					rInspectJobActionsValSaveFindings["outputConfig"] = rInspectJobActionsValSaveFindingsOutputConfig
				}
				rInspectJobActionsObject["saveFindings"] = rInspectJobActionsValSaveFindings
			}
			rInspectJobActions = append(rInspectJobActions, rInspectJobActionsObject)
		}
		rInspectJob["actions"] = rInspectJobActions
		if r.InspectJob.InspectConfig != nil && r.InspectJob.InspectConfig != dclService.EmptyJobTriggerInspectJobInspectConfig {
			rInspectJobInspectConfig := make(map[string]interface{})
			var rInspectJobInspectConfigCustomInfoTypes []interface{}
			for _, rInspectJobInspectConfigCustomInfoTypesVal := range r.InspectJob.InspectConfig.CustomInfoTypes {
				rInspectJobInspectConfigCustomInfoTypesObject := make(map[string]interface{})
				var rInspectJobInspectConfigCustomInfoTypesValDetectionRules []interface{}
				for _, rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal := range rInspectJobInspectConfigCustomInfoTypesVal.DetectionRules {
					rInspectJobInspectConfigCustomInfoTypesValDetectionRulesObject := make(map[string]interface{})
					if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule != nil && rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule {
						rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRule := make(map[string]interface{})
						if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.HotwordRegex != nil && rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.HotwordRegex != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex {
							rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegex := make(map[string]interface{})
							var rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegexGroupIndexes []interface{}
							for _, rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegexGroupIndexesVal := range rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.HotwordRegex.GroupIndexes {
								rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegexGroupIndexes = append(rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegexGroupIndexes, rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegexGroupIndexesVal)
							}
							rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegex["groupIndexes"] = rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegexGroupIndexes
							if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.HotwordRegex.Pattern != nil {
								rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegex["pattern"] = *rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.HotwordRegex.Pattern
							}
							rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRule["hotwordRegex"] = rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleHotwordRegex
						}
						if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.LikelihoodAdjustment != nil && rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.LikelihoodAdjustment != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment {
							rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleLikelihoodAdjustment := make(map[string]interface{})
							if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.LikelihoodAdjustment.FixedLikelihood != nil {
								rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleLikelihoodAdjustment["fixedLikelihood"] = string(*rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.LikelihoodAdjustment.FixedLikelihood)
							}
							if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.LikelihoodAdjustment.RelativeLikelihood != nil {
								rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleLikelihoodAdjustment["relativeLikelihood"] = *rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.LikelihoodAdjustment.RelativeLikelihood
							}
							rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRule["likelihoodAdjustment"] = rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleLikelihoodAdjustment
						}
						if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.Proximity != nil && rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.Proximity != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity {
							rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleProximity := make(map[string]interface{})
							if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.Proximity.WindowAfter != nil {
								rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleProximity["windowAfter"] = *rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.Proximity.WindowAfter
							}
							if rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.Proximity.WindowBefore != nil {
								rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleProximity["windowBefore"] = *rInspectJobInspectConfigCustomInfoTypesValDetectionRulesVal.HotwordRule.Proximity.WindowBefore
							}
							rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRule["proximity"] = rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRuleProximity
						}
						rInspectJobInspectConfigCustomInfoTypesValDetectionRulesObject["hotwordRule"] = rInspectJobInspectConfigCustomInfoTypesValDetectionRulesValHotwordRule
					}
					rInspectJobInspectConfigCustomInfoTypesValDetectionRules = append(rInspectJobInspectConfigCustomInfoTypesValDetectionRules, rInspectJobInspectConfigCustomInfoTypesValDetectionRulesObject)
				}
				rInspectJobInspectConfigCustomInfoTypesObject["detectionRules"] = rInspectJobInspectConfigCustomInfoTypesValDetectionRules
				if rInspectJobInspectConfigCustomInfoTypesVal.Dictionary != nil && rInspectJobInspectConfigCustomInfoTypesVal.Dictionary != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary {
					rInspectJobInspectConfigCustomInfoTypesValDictionary := make(map[string]interface{})
					if rInspectJobInspectConfigCustomInfoTypesVal.Dictionary.CloudStoragePath != nil && rInspectJobInspectConfigCustomInfoTypesVal.Dictionary.CloudStoragePath != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
						rInspectJobInspectConfigCustomInfoTypesValDictionaryCloudStoragePath := make(map[string]interface{})
						if rInspectJobInspectConfigCustomInfoTypesVal.Dictionary.CloudStoragePath.Path != nil {
							rInspectJobInspectConfigCustomInfoTypesValDictionaryCloudStoragePath["path"] = *rInspectJobInspectConfigCustomInfoTypesVal.Dictionary.CloudStoragePath.Path
						}
						rInspectJobInspectConfigCustomInfoTypesValDictionary["cloudStoragePath"] = rInspectJobInspectConfigCustomInfoTypesValDictionaryCloudStoragePath
					}
					if rInspectJobInspectConfigCustomInfoTypesVal.Dictionary.WordList != nil && rInspectJobInspectConfigCustomInfoTypesVal.Dictionary.WordList != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList {
						rInspectJobInspectConfigCustomInfoTypesValDictionaryWordList := make(map[string]interface{})
						var rInspectJobInspectConfigCustomInfoTypesValDictionaryWordListWords []interface{}
						for _, rInspectJobInspectConfigCustomInfoTypesValDictionaryWordListWordsVal := range rInspectJobInspectConfigCustomInfoTypesVal.Dictionary.WordList.Words {
							rInspectJobInspectConfigCustomInfoTypesValDictionaryWordListWords = append(rInspectJobInspectConfigCustomInfoTypesValDictionaryWordListWords, rInspectJobInspectConfigCustomInfoTypesValDictionaryWordListWordsVal)
						}
						rInspectJobInspectConfigCustomInfoTypesValDictionaryWordList["words"] = rInspectJobInspectConfigCustomInfoTypesValDictionaryWordListWords
						rInspectJobInspectConfigCustomInfoTypesValDictionary["wordList"] = rInspectJobInspectConfigCustomInfoTypesValDictionaryWordList
					}
					rInspectJobInspectConfigCustomInfoTypesObject["dictionary"] = rInspectJobInspectConfigCustomInfoTypesValDictionary
				}
				if rInspectJobInspectConfigCustomInfoTypesVal.ExclusionType != nil {
					rInspectJobInspectConfigCustomInfoTypesObject["exclusionType"] = string(*rInspectJobInspectConfigCustomInfoTypesVal.ExclusionType)
				}
				if rInspectJobInspectConfigCustomInfoTypesVal.InfoType != nil && rInspectJobInspectConfigCustomInfoTypesVal.InfoType != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType {
					rInspectJobInspectConfigCustomInfoTypesValInfoType := make(map[string]interface{})
					if rInspectJobInspectConfigCustomInfoTypesVal.InfoType.Name != nil {
						rInspectJobInspectConfigCustomInfoTypesValInfoType["name"] = *rInspectJobInspectConfigCustomInfoTypesVal.InfoType.Name
					}
					if rInspectJobInspectConfigCustomInfoTypesVal.InfoType.Version != nil {
						rInspectJobInspectConfigCustomInfoTypesValInfoType["version"] = *rInspectJobInspectConfigCustomInfoTypesVal.InfoType.Version
					}
					rInspectJobInspectConfigCustomInfoTypesObject["infoType"] = rInspectJobInspectConfigCustomInfoTypesValInfoType
				}
				if rInspectJobInspectConfigCustomInfoTypesVal.Likelihood != nil {
					rInspectJobInspectConfigCustomInfoTypesObject["likelihood"] = string(*rInspectJobInspectConfigCustomInfoTypesVal.Likelihood)
				}
				if rInspectJobInspectConfigCustomInfoTypesVal.Regex != nil && rInspectJobInspectConfigCustomInfoTypesVal.Regex != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesRegex {
					rInspectJobInspectConfigCustomInfoTypesValRegex := make(map[string]interface{})
					var rInspectJobInspectConfigCustomInfoTypesValRegexGroupIndexes []interface{}
					for _, rInspectJobInspectConfigCustomInfoTypesValRegexGroupIndexesVal := range rInspectJobInspectConfigCustomInfoTypesVal.Regex.GroupIndexes {
						rInspectJobInspectConfigCustomInfoTypesValRegexGroupIndexes = append(rInspectJobInspectConfigCustomInfoTypesValRegexGroupIndexes, rInspectJobInspectConfigCustomInfoTypesValRegexGroupIndexesVal)
					}
					rInspectJobInspectConfigCustomInfoTypesValRegex["groupIndexes"] = rInspectJobInspectConfigCustomInfoTypesValRegexGroupIndexes
					if rInspectJobInspectConfigCustomInfoTypesVal.Regex.Pattern != nil {
						rInspectJobInspectConfigCustomInfoTypesValRegex["pattern"] = *rInspectJobInspectConfigCustomInfoTypesVal.Regex.Pattern
					}
					rInspectJobInspectConfigCustomInfoTypesObject["regex"] = rInspectJobInspectConfigCustomInfoTypesValRegex
				}
				if rInspectJobInspectConfigCustomInfoTypesVal.StoredType != nil && rInspectJobInspectConfigCustomInfoTypesVal.StoredType != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType {
					rInspectJobInspectConfigCustomInfoTypesValStoredType := make(map[string]interface{})
					if rInspectJobInspectConfigCustomInfoTypesVal.StoredType.CreateTime != nil {
						rInspectJobInspectConfigCustomInfoTypesValStoredType["createTime"] = *rInspectJobInspectConfigCustomInfoTypesVal.StoredType.CreateTime
					}
					if rInspectJobInspectConfigCustomInfoTypesVal.StoredType.Name != nil {
						rInspectJobInspectConfigCustomInfoTypesValStoredType["name"] = *rInspectJobInspectConfigCustomInfoTypesVal.StoredType.Name
					}
					rInspectJobInspectConfigCustomInfoTypesObject["storedType"] = rInspectJobInspectConfigCustomInfoTypesValStoredType
				}
				if rInspectJobInspectConfigCustomInfoTypesVal.SurrogateType != nil && rInspectJobInspectConfigCustomInfoTypesVal.SurrogateType != dclService.EmptyJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType {
					rInspectJobInspectConfigCustomInfoTypesValSurrogateType := make(map[string]interface{})
					rInspectJobInspectConfigCustomInfoTypesObject["surrogateType"] = rInspectJobInspectConfigCustomInfoTypesValSurrogateType
				}
				rInspectJobInspectConfigCustomInfoTypes = append(rInspectJobInspectConfigCustomInfoTypes, rInspectJobInspectConfigCustomInfoTypesObject)
			}
			rInspectJobInspectConfig["customInfoTypes"] = rInspectJobInspectConfigCustomInfoTypes
			if r.InspectJob.InspectConfig.ExcludeInfoTypes != nil {
				rInspectJobInspectConfig["excludeInfoTypes"] = *r.InspectJob.InspectConfig.ExcludeInfoTypes
			}
			if r.InspectJob.InspectConfig.IncludeQuote != nil {
				rInspectJobInspectConfig["includeQuote"] = *r.InspectJob.InspectConfig.IncludeQuote
			}
			var rInspectJobInspectConfigInfoTypes []interface{}
			for _, rInspectJobInspectConfigInfoTypesVal := range r.InspectJob.InspectConfig.InfoTypes {
				rInspectJobInspectConfigInfoTypesObject := make(map[string]interface{})
				if rInspectJobInspectConfigInfoTypesVal.Name != nil {
					rInspectJobInspectConfigInfoTypesObject["name"] = *rInspectJobInspectConfigInfoTypesVal.Name
				}
				rInspectJobInspectConfigInfoTypes = append(rInspectJobInspectConfigInfoTypes, rInspectJobInspectConfigInfoTypesObject)
			}
			rInspectJobInspectConfig["infoTypes"] = rInspectJobInspectConfigInfoTypes
			if r.InspectJob.InspectConfig.Limits != nil && r.InspectJob.InspectConfig.Limits != dclService.EmptyJobTriggerInspectJobInspectConfigLimits {
				rInspectJobInspectConfigLimits := make(map[string]interface{})
				var rInspectJobInspectConfigLimitsMaxFindingsPerInfoType []interface{}
				for _, rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal := range r.InspectJob.InspectConfig.Limits.MaxFindingsPerInfoType {
					rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeObject := make(map[string]interface{})
					if rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType != nil && rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType != dclService.EmptyJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
						rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeValInfoType := make(map[string]interface{})
						if rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType.Name != nil {
							rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeValInfoType["name"] = *rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType.Name
						}
						if rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType.Version != nil {
							rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeValInfoType["version"] = *rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType.Version
						}
						rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeObject["infoType"] = rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeValInfoType
					}
					if rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal.MaxFindings != nil {
						rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeObject["maxFindings"] = *rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeVal.MaxFindings
					}
					rInspectJobInspectConfigLimitsMaxFindingsPerInfoType = append(rInspectJobInspectConfigLimitsMaxFindingsPerInfoType, rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeObject)
				}
				rInspectJobInspectConfigLimits["maxFindingsPerInfoType"] = rInspectJobInspectConfigLimitsMaxFindingsPerInfoType
				if r.InspectJob.InspectConfig.Limits.MaxFindingsPerItem != nil {
					rInspectJobInspectConfigLimits["maxFindingsPerItem"] = *r.InspectJob.InspectConfig.Limits.MaxFindingsPerItem
				}
				if r.InspectJob.InspectConfig.Limits.MaxFindingsPerRequest != nil {
					rInspectJobInspectConfigLimits["maxFindingsPerRequest"] = *r.InspectJob.InspectConfig.Limits.MaxFindingsPerRequest
				}
				rInspectJobInspectConfig["limits"] = rInspectJobInspectConfigLimits
			}
			if r.InspectJob.InspectConfig.MinLikelihood != nil {
				rInspectJobInspectConfig["minLikelihood"] = string(*r.InspectJob.InspectConfig.MinLikelihood)
			}
			var rInspectJobInspectConfigRuleSet []interface{}
			for _, rInspectJobInspectConfigRuleSetVal := range r.InspectJob.InspectConfig.RuleSet {
				rInspectJobInspectConfigRuleSetObject := make(map[string]interface{})
				var rInspectJobInspectConfigRuleSetValInfoTypes []interface{}
				for _, rInspectJobInspectConfigRuleSetValInfoTypesVal := range rInspectJobInspectConfigRuleSetVal.InfoTypes {
					rInspectJobInspectConfigRuleSetValInfoTypesObject := make(map[string]interface{})
					if rInspectJobInspectConfigRuleSetValInfoTypesVal.Name != nil {
						rInspectJobInspectConfigRuleSetValInfoTypesObject["name"] = *rInspectJobInspectConfigRuleSetValInfoTypesVal.Name
					}
					if rInspectJobInspectConfigRuleSetValInfoTypesVal.Version != nil {
						rInspectJobInspectConfigRuleSetValInfoTypesObject["version"] = *rInspectJobInspectConfigRuleSetValInfoTypesVal.Version
					}
					rInspectJobInspectConfigRuleSetValInfoTypes = append(rInspectJobInspectConfigRuleSetValInfoTypes, rInspectJobInspectConfigRuleSetValInfoTypesObject)
				}
				rInspectJobInspectConfigRuleSetObject["infoTypes"] = rInspectJobInspectConfigRuleSetValInfoTypes
				var rInspectJobInspectConfigRuleSetValRules []interface{}
				for _, rInspectJobInspectConfigRuleSetValRulesVal := range rInspectJobInspectConfigRuleSetVal.Rules {
					rInspectJobInspectConfigRuleSetValRulesObject := make(map[string]interface{})
					if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule != nil && rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
						rInspectJobInspectConfigRuleSetValRulesValExclusionRule := make(map[string]interface{})
						if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary != nil && rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
							rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionary := make(map[string]interface{})
							if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.CloudStoragePath != nil && rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.CloudStoragePath != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryCloudStoragePath := make(map[string]interface{})
								if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.CloudStoragePath.Path != nil {
									rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryCloudStoragePath["path"] = *rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.CloudStoragePath.Path
								}
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionary["cloudStoragePath"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryCloudStoragePath
							}
							if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.WordList != nil && rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.WordList != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordList := make(map[string]interface{})
								var rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWords []interface{}
								for _, rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWordsVal := range rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.WordList.Words {
									rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWords = append(rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWords, rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWordsVal)
								}
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordList["words"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWords
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionary["wordList"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordList
							}
							rInspectJobInspectConfigRuleSetValRulesValExclusionRule["dictionary"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRuleDictionary
						}
						if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.ExcludeInfoTypes != nil && rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.ExcludeInfoTypes != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
							rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypes := make(map[string]interface{})
							var rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypes []interface{}
							for _, rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesVal := range rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.ExcludeInfoTypes.InfoTypes {
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesObject := make(map[string]interface{})
								if rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesVal.Name != nil {
									rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesObject["name"] = *rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesVal.Name
								}
								if rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesVal.Version != nil {
									rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesObject["version"] = *rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesVal.Version
								}
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypes = append(rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypes, rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesObject)
							}
							rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypes["infoTypes"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypes
							rInspectJobInspectConfigRuleSetValRulesValExclusionRule["excludeInfoTypes"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypes
						}
						if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.MatchingType != nil {
							rInspectJobInspectConfigRuleSetValRulesValExclusionRule["matchingType"] = string(*rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.MatchingType)
						}
						if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Regex != nil && rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Regex != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
							rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegex := make(map[string]interface{})
							var rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexes []interface{}
							for _, rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexesVal := range rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Regex.GroupIndexes {
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexes = append(rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexes, rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexesVal)
							}
							rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegex["groupIndexes"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexes
							if rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Regex.Pattern != nil {
								rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegex["pattern"] = *rInspectJobInspectConfigRuleSetValRulesVal.ExclusionRule.Regex.Pattern
							}
							rInspectJobInspectConfigRuleSetValRulesValExclusionRule["regex"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRuleRegex
						}
						rInspectJobInspectConfigRuleSetValRulesObject["exclusionRule"] = rInspectJobInspectConfigRuleSetValRulesValExclusionRule
					}
					if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule != nil && rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
						rInspectJobInspectConfigRuleSetValRulesValHotwordRule := make(map[string]interface{})
						if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex != nil && rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
							rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegex := make(map[string]interface{})
							var rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexes []interface{}
							for _, rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexesVal := range rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex.GroupIndexes {
								rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexes = append(rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexes, rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexesVal)
							}
							rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegex["groupIndexes"] = rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexes
							if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex.Pattern != nil {
								rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegex["pattern"] = *rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex.Pattern
							}
							rInspectJobInspectConfigRuleSetValRulesValHotwordRule["hotwordRegex"] = rInspectJobInspectConfigRuleSetValRulesValHotwordRuleHotwordRegex
						}
						if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment != nil && rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
							rInspectJobInspectConfigRuleSetValRulesValHotwordRuleLikelihoodAdjustment := make(map[string]interface{})
							if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment.FixedLikelihood != nil {
								rInspectJobInspectConfigRuleSetValRulesValHotwordRuleLikelihoodAdjustment["fixedLikelihood"] = string(*rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment.FixedLikelihood)
							}
							if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment.RelativeLikelihood != nil {
								rInspectJobInspectConfigRuleSetValRulesValHotwordRuleLikelihoodAdjustment["relativeLikelihood"] = *rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment.RelativeLikelihood
							}
							rInspectJobInspectConfigRuleSetValRulesValHotwordRule["likelihoodAdjustment"] = rInspectJobInspectConfigRuleSetValRulesValHotwordRuleLikelihoodAdjustment
						}
						if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.Proximity != nil && rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.Proximity != dclService.EmptyJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity {
							rInspectJobInspectConfigRuleSetValRulesValHotwordRuleProximity := make(map[string]interface{})
							if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.Proximity.WindowAfter != nil {
								rInspectJobInspectConfigRuleSetValRulesValHotwordRuleProximity["windowAfter"] = *rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.Proximity.WindowAfter
							}
							if rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.Proximity.WindowBefore != nil {
								rInspectJobInspectConfigRuleSetValRulesValHotwordRuleProximity["windowBefore"] = *rInspectJobInspectConfigRuleSetValRulesVal.HotwordRule.Proximity.WindowBefore
							}
							rInspectJobInspectConfigRuleSetValRulesValHotwordRule["proximity"] = rInspectJobInspectConfigRuleSetValRulesValHotwordRuleProximity
						}
						rInspectJobInspectConfigRuleSetValRulesObject["hotwordRule"] = rInspectJobInspectConfigRuleSetValRulesValHotwordRule
					}
					rInspectJobInspectConfigRuleSetValRules = append(rInspectJobInspectConfigRuleSetValRules, rInspectJobInspectConfigRuleSetValRulesObject)
				}
				rInspectJobInspectConfigRuleSetObject["rules"] = rInspectJobInspectConfigRuleSetValRules
				rInspectJobInspectConfigRuleSet = append(rInspectJobInspectConfigRuleSet, rInspectJobInspectConfigRuleSetObject)
			}
			rInspectJobInspectConfig["ruleSet"] = rInspectJobInspectConfigRuleSet
			rInspectJob["inspectConfig"] = rInspectJobInspectConfig
		}
		if r.InspectJob.InspectTemplateName != nil {
			rInspectJob["inspectTemplateName"] = *r.InspectJob.InspectTemplateName
		}
		if r.InspectJob.StorageConfig != nil && r.InspectJob.StorageConfig != dclService.EmptyJobTriggerInspectJobStorageConfig {
			rInspectJobStorageConfig := make(map[string]interface{})
			if r.InspectJob.StorageConfig.BigQueryOptions != nil && r.InspectJob.StorageConfig.BigQueryOptions != dclService.EmptyJobTriggerInspectJobStorageConfigBigQueryOptions {
				rInspectJobStorageConfigBigQueryOptions := make(map[string]interface{})
				var rInspectJobStorageConfigBigQueryOptionsExcludedFields []interface{}
				for _, rInspectJobStorageConfigBigQueryOptionsExcludedFieldsVal := range r.InspectJob.StorageConfig.BigQueryOptions.ExcludedFields {
					rInspectJobStorageConfigBigQueryOptionsExcludedFieldsObject := make(map[string]interface{})
					if rInspectJobStorageConfigBigQueryOptionsExcludedFieldsVal.Name != nil {
						rInspectJobStorageConfigBigQueryOptionsExcludedFieldsObject["name"] = *rInspectJobStorageConfigBigQueryOptionsExcludedFieldsVal.Name
					}
					rInspectJobStorageConfigBigQueryOptionsExcludedFields = append(rInspectJobStorageConfigBigQueryOptionsExcludedFields, rInspectJobStorageConfigBigQueryOptionsExcludedFieldsObject)
				}
				rInspectJobStorageConfigBigQueryOptions["excludedFields"] = rInspectJobStorageConfigBigQueryOptionsExcludedFields
				var rInspectJobStorageConfigBigQueryOptionsIdentifyingFields []interface{}
				for _, rInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsVal := range r.InspectJob.StorageConfig.BigQueryOptions.IdentifyingFields {
					rInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsObject := make(map[string]interface{})
					if rInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsVal.Name != nil {
						rInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsObject["name"] = *rInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsVal.Name
					}
					rInspectJobStorageConfigBigQueryOptionsIdentifyingFields = append(rInspectJobStorageConfigBigQueryOptionsIdentifyingFields, rInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsObject)
				}
				rInspectJobStorageConfigBigQueryOptions["identifyingFields"] = rInspectJobStorageConfigBigQueryOptionsIdentifyingFields
				var rInspectJobStorageConfigBigQueryOptionsIncludedFields []interface{}
				for _, rInspectJobStorageConfigBigQueryOptionsIncludedFieldsVal := range r.InspectJob.StorageConfig.BigQueryOptions.IncludedFields {
					rInspectJobStorageConfigBigQueryOptionsIncludedFieldsObject := make(map[string]interface{})
					if rInspectJobStorageConfigBigQueryOptionsIncludedFieldsVal.Name != nil {
						rInspectJobStorageConfigBigQueryOptionsIncludedFieldsObject["name"] = *rInspectJobStorageConfigBigQueryOptionsIncludedFieldsVal.Name
					}
					rInspectJobStorageConfigBigQueryOptionsIncludedFields = append(rInspectJobStorageConfigBigQueryOptionsIncludedFields, rInspectJobStorageConfigBigQueryOptionsIncludedFieldsObject)
				}
				rInspectJobStorageConfigBigQueryOptions["includedFields"] = rInspectJobStorageConfigBigQueryOptionsIncludedFields
				if r.InspectJob.StorageConfig.BigQueryOptions.RowsLimit != nil {
					rInspectJobStorageConfigBigQueryOptions["rowsLimit"] = *r.InspectJob.StorageConfig.BigQueryOptions.RowsLimit
				}
				if r.InspectJob.StorageConfig.BigQueryOptions.RowsLimitPercent != nil {
					rInspectJobStorageConfigBigQueryOptions["rowsLimitPercent"] = *r.InspectJob.StorageConfig.BigQueryOptions.RowsLimitPercent
				}
				if r.InspectJob.StorageConfig.BigQueryOptions.SampleMethod != nil {
					rInspectJobStorageConfigBigQueryOptions["sampleMethod"] = string(*r.InspectJob.StorageConfig.BigQueryOptions.SampleMethod)
				}
				if r.InspectJob.StorageConfig.BigQueryOptions.TableReference != nil && r.InspectJob.StorageConfig.BigQueryOptions.TableReference != dclService.EmptyJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
					rInspectJobStorageConfigBigQueryOptionsTableReference := make(map[string]interface{})
					if r.InspectJob.StorageConfig.BigQueryOptions.TableReference.DatasetId != nil {
						rInspectJobStorageConfigBigQueryOptionsTableReference["datasetId"] = *r.InspectJob.StorageConfig.BigQueryOptions.TableReference.DatasetId
					}
					if r.InspectJob.StorageConfig.BigQueryOptions.TableReference.ProjectId != nil {
						rInspectJobStorageConfigBigQueryOptionsTableReference["projectId"] = *r.InspectJob.StorageConfig.BigQueryOptions.TableReference.ProjectId
					}
					if r.InspectJob.StorageConfig.BigQueryOptions.TableReference.TableId != nil {
						rInspectJobStorageConfigBigQueryOptionsTableReference["tableId"] = *r.InspectJob.StorageConfig.BigQueryOptions.TableReference.TableId
					}
					rInspectJobStorageConfigBigQueryOptions["tableReference"] = rInspectJobStorageConfigBigQueryOptionsTableReference
				}
				rInspectJobStorageConfig["bigQueryOptions"] = rInspectJobStorageConfigBigQueryOptions
			}
			if r.InspectJob.StorageConfig.CloudStorageOptions != nil && r.InspectJob.StorageConfig.CloudStorageOptions != dclService.EmptyJobTriggerInspectJobStorageConfigCloudStorageOptions {
				rInspectJobStorageConfigCloudStorageOptions := make(map[string]interface{})
				if r.InspectJob.StorageConfig.CloudStorageOptions.BytesLimitPerFile != nil {
					rInspectJobStorageConfigCloudStorageOptions["bytesLimitPerFile"] = *r.InspectJob.StorageConfig.CloudStorageOptions.BytesLimitPerFile
				}
				if r.InspectJob.StorageConfig.CloudStorageOptions.BytesLimitPerFilePercent != nil {
					rInspectJobStorageConfigCloudStorageOptions["bytesLimitPerFilePercent"] = *r.InspectJob.StorageConfig.CloudStorageOptions.BytesLimitPerFilePercent
				}
				if r.InspectJob.StorageConfig.CloudStorageOptions.FileSet != nil && r.InspectJob.StorageConfig.CloudStorageOptions.FileSet != dclService.EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
					rInspectJobStorageConfigCloudStorageOptionsFileSet := make(map[string]interface{})
					if r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet != nil && r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet != dclService.EmptyJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet {
						rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet := make(map[string]interface{})
						if r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.BucketName != nil {
							rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["bucketName"] = *r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.BucketName
						}
						var rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetExcludeRegex []interface{}
						for _, rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetExcludeRegexVal := range r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.ExcludeRegex {
							rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetExcludeRegex = append(rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetExcludeRegex, rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetExcludeRegexVal)
						}
						rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["excludeRegex"] = rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetExcludeRegex
						var rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetIncludeRegex []interface{}
						for _, rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetIncludeRegexVal := range r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.IncludeRegex {
							rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetIncludeRegex = append(rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetIncludeRegex, rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetIncludeRegexVal)
						}
						rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["includeRegex"] = rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSetIncludeRegex
						rInspectJobStorageConfigCloudStorageOptionsFileSet["regexFileSet"] = rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet
					}
					if r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.Url != nil {
						rInspectJobStorageConfigCloudStorageOptionsFileSet["url"] = *r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.Url
					}
					rInspectJobStorageConfigCloudStorageOptions["fileSet"] = rInspectJobStorageConfigCloudStorageOptionsFileSet
				}
				var rInspectJobStorageConfigCloudStorageOptionsFileTypes []interface{}
				for _, rInspectJobStorageConfigCloudStorageOptionsFileTypesVal := range r.InspectJob.StorageConfig.CloudStorageOptions.FileTypes {
					rInspectJobStorageConfigCloudStorageOptionsFileTypes = append(rInspectJobStorageConfigCloudStorageOptionsFileTypes, string(rInspectJobStorageConfigCloudStorageOptionsFileTypesVal))
				}
				rInspectJobStorageConfigCloudStorageOptions["fileTypes"] = rInspectJobStorageConfigCloudStorageOptionsFileTypes
				if r.InspectJob.StorageConfig.CloudStorageOptions.FilesLimitPercent != nil {
					rInspectJobStorageConfigCloudStorageOptions["filesLimitPercent"] = *r.InspectJob.StorageConfig.CloudStorageOptions.FilesLimitPercent
				}
				if r.InspectJob.StorageConfig.CloudStorageOptions.SampleMethod != nil {
					rInspectJobStorageConfigCloudStorageOptions["sampleMethod"] = string(*r.InspectJob.StorageConfig.CloudStorageOptions.SampleMethod)
				}
				rInspectJobStorageConfig["cloudStorageOptions"] = rInspectJobStorageConfigCloudStorageOptions
			}
			if r.InspectJob.StorageConfig.DatastoreOptions != nil && r.InspectJob.StorageConfig.DatastoreOptions != dclService.EmptyJobTriggerInspectJobStorageConfigDatastoreOptions {
				rInspectJobStorageConfigDatastoreOptions := make(map[string]interface{})
				if r.InspectJob.StorageConfig.DatastoreOptions.Kind != nil && r.InspectJob.StorageConfig.DatastoreOptions.Kind != dclService.EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsKind {
					rInspectJobStorageConfigDatastoreOptionsKind := make(map[string]interface{})
					if r.InspectJob.StorageConfig.DatastoreOptions.Kind.Name != nil {
						rInspectJobStorageConfigDatastoreOptionsKind["name"] = *r.InspectJob.StorageConfig.DatastoreOptions.Kind.Name
					}
					rInspectJobStorageConfigDatastoreOptions["kind"] = rInspectJobStorageConfigDatastoreOptionsKind
				}
				if r.InspectJob.StorageConfig.DatastoreOptions.PartitionId != nil && r.InspectJob.StorageConfig.DatastoreOptions.PartitionId != dclService.EmptyJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId {
					rInspectJobStorageConfigDatastoreOptionsPartitionId := make(map[string]interface{})
					if r.InspectJob.StorageConfig.DatastoreOptions.PartitionId.NamespaceId != nil {
						rInspectJobStorageConfigDatastoreOptionsPartitionId["namespaceId"] = *r.InspectJob.StorageConfig.DatastoreOptions.PartitionId.NamespaceId
					}
					if r.InspectJob.StorageConfig.DatastoreOptions.PartitionId.ProjectId != nil {
						rInspectJobStorageConfigDatastoreOptionsPartitionId["projectId"] = *r.InspectJob.StorageConfig.DatastoreOptions.PartitionId.ProjectId
					}
					rInspectJobStorageConfigDatastoreOptions["partitionId"] = rInspectJobStorageConfigDatastoreOptionsPartitionId
				}
				rInspectJobStorageConfig["datastoreOptions"] = rInspectJobStorageConfigDatastoreOptions
			}
			if r.InspectJob.StorageConfig.HybridOptions != nil && r.InspectJob.StorageConfig.HybridOptions != dclService.EmptyJobTriggerInspectJobStorageConfigHybridOptions {
				rInspectJobStorageConfigHybridOptions := make(map[string]interface{})
				if r.InspectJob.StorageConfig.HybridOptions.Description != nil {
					rInspectJobStorageConfigHybridOptions["description"] = *r.InspectJob.StorageConfig.HybridOptions.Description
				}
				if r.InspectJob.StorageConfig.HybridOptions.Labels != nil {
					rInspectJobStorageConfigHybridOptionsLabels := make(map[string]interface{})
					for k, v := range r.InspectJob.StorageConfig.HybridOptions.Labels {
						rInspectJobStorageConfigHybridOptionsLabels[k] = v
					}
					rInspectJobStorageConfigHybridOptions["labels"] = rInspectJobStorageConfigHybridOptionsLabels
				}
				var rInspectJobStorageConfigHybridOptionsRequiredFindingLabelKeys []interface{}
				for _, rInspectJobStorageConfigHybridOptionsRequiredFindingLabelKeysVal := range r.InspectJob.StorageConfig.HybridOptions.RequiredFindingLabelKeys {
					rInspectJobStorageConfigHybridOptionsRequiredFindingLabelKeys = append(rInspectJobStorageConfigHybridOptionsRequiredFindingLabelKeys, rInspectJobStorageConfigHybridOptionsRequiredFindingLabelKeysVal)
				}
				rInspectJobStorageConfigHybridOptions["requiredFindingLabelKeys"] = rInspectJobStorageConfigHybridOptionsRequiredFindingLabelKeys
				if r.InspectJob.StorageConfig.HybridOptions.TableOptions != nil && r.InspectJob.StorageConfig.HybridOptions.TableOptions != dclService.EmptyJobTriggerInspectJobStorageConfigHybridOptionsTableOptions {
					rInspectJobStorageConfigHybridOptionsTableOptions := make(map[string]interface{})
					var rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields []interface{}
					for _, rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsVal := range r.InspectJob.StorageConfig.HybridOptions.TableOptions.IdentifyingFields {
						rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsObject := make(map[string]interface{})
						if rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsVal.Name != nil {
							rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsObject["name"] = *rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsVal.Name
						}
						rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields = append(rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields, rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFieldsObject)
					}
					rInspectJobStorageConfigHybridOptionsTableOptions["identifyingFields"] = rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields
					rInspectJobStorageConfigHybridOptions["tableOptions"] = rInspectJobStorageConfigHybridOptionsTableOptions
				}
				rInspectJobStorageConfig["hybridOptions"] = rInspectJobStorageConfigHybridOptions
			}
			if r.InspectJob.StorageConfig.TimespanConfig != nil && r.InspectJob.StorageConfig.TimespanConfig != dclService.EmptyJobTriggerInspectJobStorageConfigTimespanConfig {
				rInspectJobStorageConfigTimespanConfig := make(map[string]interface{})
				if r.InspectJob.StorageConfig.TimespanConfig.EnableAutoPopulationOfTimespanConfig != nil {
					rInspectJobStorageConfigTimespanConfig["enableAutoPopulationOfTimespanConfig"] = *r.InspectJob.StorageConfig.TimespanConfig.EnableAutoPopulationOfTimespanConfig
				}
				if r.InspectJob.StorageConfig.TimespanConfig.EndTime != nil {
					rInspectJobStorageConfigTimespanConfig["endTime"] = *r.InspectJob.StorageConfig.TimespanConfig.EndTime
				}
				if r.InspectJob.StorageConfig.TimespanConfig.StartTime != nil {
					rInspectJobStorageConfigTimespanConfig["startTime"] = *r.InspectJob.StorageConfig.TimespanConfig.StartTime
				}
				if r.InspectJob.StorageConfig.TimespanConfig.TimestampField != nil && r.InspectJob.StorageConfig.TimespanConfig.TimestampField != dclService.EmptyJobTriggerInspectJobStorageConfigTimespanConfigTimestampField {
					rInspectJobStorageConfigTimespanConfigTimestampField := make(map[string]interface{})
					if r.InspectJob.StorageConfig.TimespanConfig.TimestampField.Name != nil {
						rInspectJobStorageConfigTimespanConfigTimestampField["name"] = *r.InspectJob.StorageConfig.TimespanConfig.TimestampField.Name
					}
					rInspectJobStorageConfigTimespanConfig["timestampField"] = rInspectJobStorageConfigTimespanConfigTimestampField
				}
				rInspectJobStorageConfig["timespanConfig"] = rInspectJobStorageConfigTimespanConfig
			}
			rInspectJob["storageConfig"] = rInspectJobStorageConfig
		}
		u.Object["inspectJob"] = rInspectJob
	}
	if r.LastRunTime != nil {
		u.Object["lastRunTime"] = *r.LastRunTime
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.LocationId != nil {
		u.Object["locationId"] = *r.LocationId
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.Status != nil {
		u.Object["status"] = string(*r.Status)
	}
	var rTriggers []interface{}
	for _, rTriggersVal := range r.Triggers {
		rTriggersObject := make(map[string]interface{})
		if rTriggersVal.Manual != nil && rTriggersVal.Manual != dclService.EmptyJobTriggerTriggersManual {
			rTriggersValManual := make(map[string]interface{})
			rTriggersObject["manual"] = rTriggersValManual
		}
		if rTriggersVal.Schedule != nil && rTriggersVal.Schedule != dclService.EmptyJobTriggerTriggersSchedule {
			rTriggersValSchedule := make(map[string]interface{})
			if rTriggersVal.Schedule.RecurrencePeriodDuration != nil {
				rTriggersValSchedule["recurrencePeriodDuration"] = *rTriggersVal.Schedule.RecurrencePeriodDuration
			}
			rTriggersObject["schedule"] = rTriggersValSchedule
		}
		rTriggers = append(rTriggers, rTriggersObject)
	}
	u.Object["triggers"] = rTriggers
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToJobTrigger(u *unstructured.Resource) (*dclService.JobTrigger, error) {
	r := &dclService.JobTrigger{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["errors"]; ok {
		if s, ok := u.Object["errors"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rErrors dclService.JobTriggerErrors
					if _, ok := objval["details"]; ok {
						if rErrorsDetails, ok := objval["details"].(map[string]interface{}); ok {
							rErrors.Details = &dclService.JobTriggerErrorsDetails{}
							if _, ok := rErrorsDetails["code"]; ok {
								if i, ok := rErrorsDetails["code"].(int64); ok {
									rErrors.Details.Code = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rErrors.Details.Code: expected int64")
								}
							}
							if _, ok := rErrorsDetails["details"]; ok {
								if s, ok := rErrorsDetails["details"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rErrorsDetailsDetails dclService.JobTriggerErrorsDetailsDetails
											if _, ok := objval["typeUrl"]; ok {
												if s, ok := objval["typeUrl"].(string); ok {
													rErrorsDetailsDetails.TypeUrl = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rErrorsDetailsDetails.TypeUrl: expected string")
												}
											}
											if _, ok := objval["value"]; ok {
												if s, ok := objval["value"].(string); ok {
													rErrorsDetailsDetails.Value = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rErrorsDetailsDetails.Value: expected string")
												}
											}
											rErrors.Details.Details = append(rErrors.Details.Details, rErrorsDetailsDetails)
										}
									}
								} else {
									return nil, fmt.Errorf("rErrors.Details.Details: expected []interface{}")
								}
							}
							if _, ok := rErrorsDetails["message"]; ok {
								if s, ok := rErrorsDetails["message"].(string); ok {
									rErrors.Details.Message = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rErrors.Details.Message: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rErrors.Details: expected map[string]interface{}")
						}
					}
					if _, ok := objval["timestamps"]; ok {
						if s, ok := objval["timestamps"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rErrors.Timestamps = append(rErrors.Timestamps, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rErrors.Timestamps: expected []interface{}")
						}
					}
					r.Errors = append(r.Errors, rErrors)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Errors: expected []interface{}")
		}
	}
	if _, ok := u.Object["inspectJob"]; ok {
		if rInspectJob, ok := u.Object["inspectJob"].(map[string]interface{}); ok {
			r.InspectJob = &dclService.JobTriggerInspectJob{}
			if _, ok := rInspectJob["actions"]; ok {
				if s, ok := rInspectJob["actions"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rInspectJobActions dclService.JobTriggerInspectJobActions
							if _, ok := objval["jobNotificationEmails"]; ok {
								if _, ok := objval["jobNotificationEmails"].(map[string]interface{}); ok {
									rInspectJobActions.JobNotificationEmails = &dclService.JobTriggerInspectJobActionsJobNotificationEmails{}
								} else {
									return nil, fmt.Errorf("rInspectJobActions.JobNotificationEmails: expected map[string]interface{}")
								}
							}
							if _, ok := objval["pubSub"]; ok {
								if rInspectJobActionsPubSub, ok := objval["pubSub"].(map[string]interface{}); ok {
									rInspectJobActions.PubSub = &dclService.JobTriggerInspectJobActionsPubSub{}
									if _, ok := rInspectJobActionsPubSub["topic"]; ok {
										if s, ok := rInspectJobActionsPubSub["topic"].(string); ok {
											rInspectJobActions.PubSub.Topic = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rInspectJobActions.PubSub.Topic: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rInspectJobActions.PubSub: expected map[string]interface{}")
								}
							}
							if _, ok := objval["publishFindingsToCloudDataCatalog"]; ok {
								if _, ok := objval["publishFindingsToCloudDataCatalog"].(map[string]interface{}); ok {
									rInspectJobActions.PublishFindingsToCloudDataCatalog = &dclService.JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog{}
								} else {
									return nil, fmt.Errorf("rInspectJobActions.PublishFindingsToCloudDataCatalog: expected map[string]interface{}")
								}
							}
							if _, ok := objval["publishSummaryToCscc"]; ok {
								if _, ok := objval["publishSummaryToCscc"].(map[string]interface{}); ok {
									rInspectJobActions.PublishSummaryToCscc = &dclService.JobTriggerInspectJobActionsPublishSummaryToCscc{}
								} else {
									return nil, fmt.Errorf("rInspectJobActions.PublishSummaryToCscc: expected map[string]interface{}")
								}
							}
							if _, ok := objval["publishToStackdriver"]; ok {
								if _, ok := objval["publishToStackdriver"].(map[string]interface{}); ok {
									rInspectJobActions.PublishToStackdriver = &dclService.JobTriggerInspectJobActionsPublishToStackdriver{}
								} else {
									return nil, fmt.Errorf("rInspectJobActions.PublishToStackdriver: expected map[string]interface{}")
								}
							}
							if _, ok := objval["saveFindings"]; ok {
								if rInspectJobActionsSaveFindings, ok := objval["saveFindings"].(map[string]interface{}); ok {
									rInspectJobActions.SaveFindings = &dclService.JobTriggerInspectJobActionsSaveFindings{}
									if _, ok := rInspectJobActionsSaveFindings["outputConfig"]; ok {
										if rInspectJobActionsSaveFindingsOutputConfig, ok := rInspectJobActionsSaveFindings["outputConfig"].(map[string]interface{}); ok {
											rInspectJobActions.SaveFindings.OutputConfig = &dclService.JobTriggerInspectJobActionsSaveFindingsOutputConfig{}
											if _, ok := rInspectJobActionsSaveFindingsOutputConfig["dlpStorage"]; ok {
												if _, ok := rInspectJobActionsSaveFindingsOutputConfig["dlpStorage"].(map[string]interface{}); ok {
													rInspectJobActions.SaveFindings.OutputConfig.DlpStorage = &dclService.JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage{}
												} else {
													return nil, fmt.Errorf("rInspectJobActions.SaveFindings.OutputConfig.DlpStorage: expected map[string]interface{}")
												}
											}
											if _, ok := rInspectJobActionsSaveFindingsOutputConfig["outputSchema"]; ok {
												if s, ok := rInspectJobActionsSaveFindingsOutputConfig["outputSchema"].(string); ok {
													rInspectJobActions.SaveFindings.OutputConfig.OutputSchema = dclService.JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnumRef(s)
												} else {
													return nil, fmt.Errorf("rInspectJobActions.SaveFindings.OutputConfig.OutputSchema: expected string")
												}
											}
											if _, ok := rInspectJobActionsSaveFindingsOutputConfig["table"]; ok {
												if rInspectJobActionsSaveFindingsOutputConfigTable, ok := rInspectJobActionsSaveFindingsOutputConfig["table"].(map[string]interface{}); ok {
													rInspectJobActions.SaveFindings.OutputConfig.Table = &dclService.JobTriggerInspectJobActionsSaveFindingsOutputConfigTable{}
													if _, ok := rInspectJobActionsSaveFindingsOutputConfigTable["datasetId"]; ok {
														if s, ok := rInspectJobActionsSaveFindingsOutputConfigTable["datasetId"].(string); ok {
															rInspectJobActions.SaveFindings.OutputConfig.Table.DatasetId = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobActions.SaveFindings.OutputConfig.Table.DatasetId: expected string")
														}
													}
													if _, ok := rInspectJobActionsSaveFindingsOutputConfigTable["projectId"]; ok {
														if s, ok := rInspectJobActionsSaveFindingsOutputConfigTable["projectId"].(string); ok {
															rInspectJobActions.SaveFindings.OutputConfig.Table.ProjectId = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobActions.SaveFindings.OutputConfig.Table.ProjectId: expected string")
														}
													}
													if _, ok := rInspectJobActionsSaveFindingsOutputConfigTable["tableId"]; ok {
														if s, ok := rInspectJobActionsSaveFindingsOutputConfigTable["tableId"].(string); ok {
															rInspectJobActions.SaveFindings.OutputConfig.Table.TableId = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobActions.SaveFindings.OutputConfig.Table.TableId: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rInspectJobActions.SaveFindings.OutputConfig.Table: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectJobActions.SaveFindings.OutputConfig: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rInspectJobActions.SaveFindings: expected map[string]interface{}")
								}
							}
							r.InspectJob.Actions = append(r.InspectJob.Actions, rInspectJobActions)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InspectJob.Actions: expected []interface{}")
				}
			}
			if _, ok := rInspectJob["inspectConfig"]; ok {
				if rInspectJobInspectConfig, ok := rInspectJob["inspectConfig"].(map[string]interface{}); ok {
					r.InspectJob.InspectConfig = &dclService.JobTriggerInspectJobInspectConfig{}
					if _, ok := rInspectJobInspectConfig["customInfoTypes"]; ok {
						if s, ok := rInspectJobInspectConfig["customInfoTypes"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rInspectJobInspectConfigCustomInfoTypes dclService.JobTriggerInspectJobInspectConfigCustomInfoTypes
									if _, ok := objval["detectionRules"]; ok {
										if s, ok := objval["detectionRules"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rInspectJobInspectConfigCustomInfoTypesDetectionRules dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules
													if _, ok := objval["hotwordRule"]; ok {
														if rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule, ok := objval["hotwordRule"].(map[string]interface{}); ok {
															rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule{}
															if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule["hotwordRegex"]; ok {
																if rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule["hotwordRegex"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.HotwordRegex = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex{}
																	if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex["groupIndexes"]; ok {
																		if s, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex["groupIndexes"].([]interface{}); ok {
																			for _, ss := range s {
																				if intval, ok := ss.(int64); ok {
																					rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.HotwordRegex.GroupIndexes = append(rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.HotwordRegex.GroupIndexes, intval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.HotwordRegex.GroupIndexes: expected []interface{}")
																		}
																	}
																	if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex["pattern"]; ok {
																		if s, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex["pattern"].(string); ok {
																			rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.HotwordRegex.Pattern = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.HotwordRegex.Pattern: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.HotwordRegex: expected map[string]interface{}")
																}
															}
															if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule["likelihoodAdjustment"]; ok {
																if rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule["likelihoodAdjustment"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.LikelihoodAdjustment = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment{}
																	if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment["fixedLikelihood"]; ok {
																		if s, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment["fixedLikelihood"].(string); ok {
																			rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.LikelihoodAdjustment.FixedLikelihood = dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.LikelihoodAdjustment.FixedLikelihood: expected string")
																		}
																	}
																	if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment["relativeLikelihood"]; ok {
																		if i, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment["relativeLikelihood"].(int64); ok {
																			rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.LikelihoodAdjustment.RelativeLikelihood = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.LikelihoodAdjustment.RelativeLikelihood: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.LikelihoodAdjustment: expected map[string]interface{}")
																}
															}
															if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule["proximity"]; ok {
																if rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule["proximity"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.Proximity = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity{}
																	if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity["windowAfter"]; ok {
																		if i, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity["windowAfter"].(int64); ok {
																			rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.Proximity.WindowAfter = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.Proximity.WindowAfter: expected int64")
																		}
																	}
																	if _, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity["windowBefore"]; ok {
																		if i, ok := rInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity["windowBefore"].(int64); ok {
																			rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.Proximity.WindowBefore = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.Proximity.WindowBefore: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule.Proximity: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypesDetectionRules.HotwordRule: expected map[string]interface{}")
														}
													}
													rInspectJobInspectConfigCustomInfoTypes.DetectionRules = append(rInspectJobInspectConfigCustomInfoTypes.DetectionRules, rInspectJobInspectConfigCustomInfoTypesDetectionRules)
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.DetectionRules: expected []interface{}")
										}
									}
									if _, ok := objval["dictionary"]; ok {
										if rInspectJobInspectConfigCustomInfoTypesDictionary, ok := objval["dictionary"].(map[string]interface{}); ok {
											rInspectJobInspectConfigCustomInfoTypes.Dictionary = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary{}
											if _, ok := rInspectJobInspectConfigCustomInfoTypesDictionary["cloudStoragePath"]; ok {
												if rInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath, ok := rInspectJobInspectConfigCustomInfoTypesDictionary["cloudStoragePath"].(map[string]interface{}); ok {
													rInspectJobInspectConfigCustomInfoTypes.Dictionary.CloudStoragePath = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
													if _, ok := rInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath["path"]; ok {
														if s, ok := rInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath["path"].(string); ok {
															rInspectJobInspectConfigCustomInfoTypes.Dictionary.CloudStoragePath.Path = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Dictionary.CloudStoragePath.Path: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Dictionary.CloudStoragePath: expected map[string]interface{}")
												}
											}
											if _, ok := rInspectJobInspectConfigCustomInfoTypesDictionary["wordList"]; ok {
												if rInspectJobInspectConfigCustomInfoTypesDictionaryWordList, ok := rInspectJobInspectConfigCustomInfoTypesDictionary["wordList"].(map[string]interface{}); ok {
													rInspectJobInspectConfigCustomInfoTypes.Dictionary.WordList = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList{}
													if _, ok := rInspectJobInspectConfigCustomInfoTypesDictionaryWordList["words"]; ok {
														if s, ok := rInspectJobInspectConfigCustomInfoTypesDictionaryWordList["words"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rInspectJobInspectConfigCustomInfoTypes.Dictionary.WordList.Words = append(rInspectJobInspectConfigCustomInfoTypes.Dictionary.WordList.Words, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Dictionary.WordList.Words: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Dictionary.WordList: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Dictionary: expected map[string]interface{}")
										}
									}
									if _, ok := objval["exclusionType"]; ok {
										if s, ok := objval["exclusionType"].(string); ok {
											rInspectJobInspectConfigCustomInfoTypes.ExclusionType = dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.ExclusionType: expected string")
										}
									}
									if _, ok := objval["infoType"]; ok {
										if rInspectJobInspectConfigCustomInfoTypesInfoType, ok := objval["infoType"].(map[string]interface{}); ok {
											rInspectJobInspectConfigCustomInfoTypes.InfoType = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType{}
											if _, ok := rInspectJobInspectConfigCustomInfoTypesInfoType["name"]; ok {
												if s, ok := rInspectJobInspectConfigCustomInfoTypesInfoType["name"].(string); ok {
													rInspectJobInspectConfigCustomInfoTypes.InfoType.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.InfoType.Name: expected string")
												}
											}
											if _, ok := rInspectJobInspectConfigCustomInfoTypesInfoType["version"]; ok {
												if s, ok := rInspectJobInspectConfigCustomInfoTypesInfoType["version"].(string); ok {
													rInspectJobInspectConfigCustomInfoTypes.InfoType.Version = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.InfoType.Version: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.InfoType: expected map[string]interface{}")
										}
									}
									if _, ok := objval["likelihood"]; ok {
										if s, ok := objval["likelihood"].(string); ok {
											rInspectJobInspectConfigCustomInfoTypes.Likelihood = dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnumRef(s)
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Likelihood: expected string")
										}
									}
									if _, ok := objval["regex"]; ok {
										if rInspectJobInspectConfigCustomInfoTypesRegex, ok := objval["regex"].(map[string]interface{}); ok {
											rInspectJobInspectConfigCustomInfoTypes.Regex = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesRegex{}
											if _, ok := rInspectJobInspectConfigCustomInfoTypesRegex["groupIndexes"]; ok {
												if s, ok := rInspectJobInspectConfigCustomInfoTypesRegex["groupIndexes"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rInspectJobInspectConfigCustomInfoTypes.Regex.GroupIndexes = append(rInspectJobInspectConfigCustomInfoTypes.Regex.GroupIndexes, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Regex.GroupIndexes: expected []interface{}")
												}
											}
											if _, ok := rInspectJobInspectConfigCustomInfoTypesRegex["pattern"]; ok {
												if s, ok := rInspectJobInspectConfigCustomInfoTypesRegex["pattern"].(string); ok {
													rInspectJobInspectConfigCustomInfoTypes.Regex.Pattern = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Regex.Pattern: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.Regex: expected map[string]interface{}")
										}
									}
									if _, ok := objval["storedType"]; ok {
										if rInspectJobInspectConfigCustomInfoTypesStoredType, ok := objval["storedType"].(map[string]interface{}); ok {
											rInspectJobInspectConfigCustomInfoTypes.StoredType = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType{}
											if _, ok := rInspectJobInspectConfigCustomInfoTypesStoredType["createTime"]; ok {
												if s, ok := rInspectJobInspectConfigCustomInfoTypesStoredType["createTime"].(string); ok {
													rInspectJobInspectConfigCustomInfoTypes.StoredType.CreateTime = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.StoredType.CreateTime: expected string")
												}
											}
											if _, ok := rInspectJobInspectConfigCustomInfoTypesStoredType["name"]; ok {
												if s, ok := rInspectJobInspectConfigCustomInfoTypesStoredType["name"].(string); ok {
													rInspectJobInspectConfigCustomInfoTypes.StoredType.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.StoredType.Name: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.StoredType: expected map[string]interface{}")
										}
									}
									if _, ok := objval["surrogateType"]; ok {
										if _, ok := objval["surrogateType"].(map[string]interface{}); ok {
											rInspectJobInspectConfigCustomInfoTypes.SurrogateType = &dclService.JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType{}
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigCustomInfoTypes.SurrogateType: expected map[string]interface{}")
										}
									}
									r.InspectJob.InspectConfig.CustomInfoTypes = append(r.InspectJob.InspectConfig.CustomInfoTypes, rInspectJobInspectConfigCustomInfoTypes)
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.InspectConfig.CustomInfoTypes: expected []interface{}")
						}
					}
					if _, ok := rInspectJobInspectConfig["excludeInfoTypes"]; ok {
						if b, ok := rInspectJobInspectConfig["excludeInfoTypes"].(bool); ok {
							r.InspectJob.InspectConfig.ExcludeInfoTypes = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.InspectJob.InspectConfig.ExcludeInfoTypes: expected bool")
						}
					}
					if _, ok := rInspectJobInspectConfig["includeQuote"]; ok {
						if b, ok := rInspectJobInspectConfig["includeQuote"].(bool); ok {
							r.InspectJob.InspectConfig.IncludeQuote = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.InspectJob.InspectConfig.IncludeQuote: expected bool")
						}
					}
					if _, ok := rInspectJobInspectConfig["infoTypes"]; ok {
						if s, ok := rInspectJobInspectConfig["infoTypes"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rInspectJobInspectConfigInfoTypes dclService.JobTriggerInspectJobInspectConfigInfoTypes
									if _, ok := objval["name"]; ok {
										if s, ok := objval["name"].(string); ok {
											rInspectJobInspectConfigInfoTypes.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigInfoTypes.Name: expected string")
										}
									}
									r.InspectJob.InspectConfig.InfoTypes = append(r.InspectJob.InspectConfig.InfoTypes, rInspectJobInspectConfigInfoTypes)
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.InspectConfig.InfoTypes: expected []interface{}")
						}
					}
					if _, ok := rInspectJobInspectConfig["limits"]; ok {
						if rInspectJobInspectConfigLimits, ok := rInspectJobInspectConfig["limits"].(map[string]interface{}); ok {
							r.InspectJob.InspectConfig.Limits = &dclService.JobTriggerInspectJobInspectConfigLimits{}
							if _, ok := rInspectJobInspectConfigLimits["maxFindingsPerInfoType"]; ok {
								if s, ok := rInspectJobInspectConfigLimits["maxFindingsPerInfoType"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rInspectJobInspectConfigLimitsMaxFindingsPerInfoType dclService.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType
											if _, ok := objval["infoType"]; ok {
												if rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, ok := objval["infoType"].(map[string]interface{}); ok {
													rInspectJobInspectConfigLimitsMaxFindingsPerInfoType.InfoType = &dclService.JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
													if _, ok := rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType["name"]; ok {
														if s, ok := rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType["name"].(string); ok {
															rInspectJobInspectConfigLimitsMaxFindingsPerInfoType.InfoType.Name = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigLimitsMaxFindingsPerInfoType.InfoType.Name: expected string")
														}
													}
													if _, ok := rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType["version"]; ok {
														if s, ok := rInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType["version"].(string); ok {
															rInspectJobInspectConfigLimitsMaxFindingsPerInfoType.InfoType.Version = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigLimitsMaxFindingsPerInfoType.InfoType.Version: expected string")
														}
													}
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigLimitsMaxFindingsPerInfoType.InfoType: expected map[string]interface{}")
												}
											}
											if _, ok := objval["maxFindings"]; ok {
												if i, ok := objval["maxFindings"].(int64); ok {
													rInspectJobInspectConfigLimitsMaxFindingsPerInfoType.MaxFindings = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rInspectJobInspectConfigLimitsMaxFindingsPerInfoType.MaxFindings: expected int64")
												}
											}
											r.InspectJob.InspectConfig.Limits.MaxFindingsPerInfoType = append(r.InspectJob.InspectConfig.Limits.MaxFindingsPerInfoType, rInspectJobInspectConfigLimitsMaxFindingsPerInfoType)
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.InspectConfig.Limits.MaxFindingsPerInfoType: expected []interface{}")
								}
							}
							if _, ok := rInspectJobInspectConfigLimits["maxFindingsPerItem"]; ok {
								if i, ok := rInspectJobInspectConfigLimits["maxFindingsPerItem"].(int64); ok {
									r.InspectJob.InspectConfig.Limits.MaxFindingsPerItem = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.InspectJob.InspectConfig.Limits.MaxFindingsPerItem: expected int64")
								}
							}
							if _, ok := rInspectJobInspectConfigLimits["maxFindingsPerRequest"]; ok {
								if i, ok := rInspectJobInspectConfigLimits["maxFindingsPerRequest"].(int64); ok {
									r.InspectJob.InspectConfig.Limits.MaxFindingsPerRequest = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.InspectJob.InspectConfig.Limits.MaxFindingsPerRequest: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.InspectConfig.Limits: expected map[string]interface{}")
						}
					}
					if _, ok := rInspectJobInspectConfig["minLikelihood"]; ok {
						if s, ok := rInspectJobInspectConfig["minLikelihood"].(string); ok {
							r.InspectJob.InspectConfig.MinLikelihood = dclService.JobTriggerInspectJobInspectConfigMinLikelihoodEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.InspectJob.InspectConfig.MinLikelihood: expected string")
						}
					}
					if _, ok := rInspectJobInspectConfig["ruleSet"]; ok {
						if s, ok := rInspectJobInspectConfig["ruleSet"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rInspectJobInspectConfigRuleSet dclService.JobTriggerInspectJobInspectConfigRuleSet
									if _, ok := objval["infoTypes"]; ok {
										if s, ok := objval["infoTypes"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rInspectJobInspectConfigRuleSetInfoTypes dclService.JobTriggerInspectJobInspectConfigRuleSetInfoTypes
													if _, ok := objval["name"]; ok {
														if s, ok := objval["name"].(string); ok {
															rInspectJobInspectConfigRuleSetInfoTypes.Name = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetInfoTypes.Name: expected string")
														}
													}
													if _, ok := objval["version"]; ok {
														if s, ok := objval["version"].(string); ok {
															rInspectJobInspectConfigRuleSetInfoTypes.Version = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetInfoTypes.Version: expected string")
														}
													}
													rInspectJobInspectConfigRuleSet.InfoTypes = append(rInspectJobInspectConfigRuleSet.InfoTypes, rInspectJobInspectConfigRuleSetInfoTypes)
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigRuleSet.InfoTypes: expected []interface{}")
										}
									}
									if _, ok := objval["rules"]; ok {
										if s, ok := objval["rules"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rInspectJobInspectConfigRuleSetRules dclService.JobTriggerInspectJobInspectConfigRuleSetRules
													if _, ok := objval["exclusionRule"]; ok {
														if rInspectJobInspectConfigRuleSetRulesExclusionRule, ok := objval["exclusionRule"].(map[string]interface{}); ok {
															rInspectJobInspectConfigRuleSetRules.ExclusionRule = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule{}
															if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRule["dictionary"]; ok {
																if rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary, ok := rInspectJobInspectConfigRuleSetRulesExclusionRule["dictionary"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary{}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary["cloudStoragePath"]; ok {
																		if rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary["cloudStoragePath"].(map[string]interface{}); ok {
																			rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.CloudStoragePath = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
																			if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath["path"]; ok {
																				if s, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath["path"].(string); ok {
																					rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.CloudStoragePath.Path = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.CloudStoragePath.Path: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.CloudStoragePath: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary["wordList"]; ok {
																		if rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary["wordList"].(map[string]interface{}); ok {
																			rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
																			if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList["words"]; ok {
																				if s, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList["words"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList.Words = append(rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList.Words, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList.Words: expected []interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.Dictionary: expected map[string]interface{}")
																}
															}
															if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRule["excludeInfoTypes"]; ok {
																if rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, ok := rInspectJobInspectConfigRuleSetRulesExclusionRule["excludeInfoTypes"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes["infoTypes"]; ok {
																		if s, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes["infoTypes"].([]interface{}); ok {
																			for _, o := range s {
																				if objval, ok := o.(map[string]interface{}); ok {
																					var rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes dclService.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
																					if _, ok := objval["name"]; ok {
																						if s, ok := objval["name"].(string); ok {
																							rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.Name = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.Name: expected string")
																						}
																					}
																					if _, ok := objval["version"]; ok {
																						if s, ok := objval["version"].(string); ok {
																							rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.Version = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.Version: expected string")
																						}
																					}
																					rInspectJobInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes.InfoTypes = append(rInspectJobInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes.InfoTypes, rInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes.InfoTypes: expected []interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes: expected map[string]interface{}")
																}
															}
															if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRule["matchingType"]; ok {
																if s, ok := rInspectJobInspectConfigRuleSetRulesExclusionRule["matchingType"].(string); ok {
																	rInspectJobInspectConfigRuleSetRules.ExclusionRule.MatchingType = dclService.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.MatchingType: expected string")
																}
															}
															if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRule["regex"]; ok {
																if rInspectJobInspectConfigRuleSetRulesExclusionRuleRegex, ok := rInspectJobInspectConfigRuleSetRulesExclusionRule["regex"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigRuleSetRules.ExclusionRule.Regex = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex{}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleRegex["groupIndexes"]; ok {
																		if s, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleRegex["groupIndexes"].([]interface{}); ok {
																			for _, ss := range s {
																				if intval, ok := ss.(int64); ok {
																					rInspectJobInspectConfigRuleSetRules.ExclusionRule.Regex.GroupIndexes = append(rInspectJobInspectConfigRuleSetRules.ExclusionRule.Regex.GroupIndexes, intval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.Regex.GroupIndexes: expected []interface{}")
																		}
																	}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleRegex["pattern"]; ok {
																		if s, ok := rInspectJobInspectConfigRuleSetRulesExclusionRuleRegex["pattern"].(string); ok {
																			rInspectJobInspectConfigRuleSetRules.ExclusionRule.Regex.Pattern = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.Regex.Pattern: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule.Regex: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.ExclusionRule: expected map[string]interface{}")
														}
													}
													if _, ok := objval["hotwordRule"]; ok {
														if rInspectJobInspectConfigRuleSetRulesHotwordRule, ok := objval["hotwordRule"].(map[string]interface{}); ok {
															rInspectJobInspectConfigRuleSetRules.HotwordRule = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule{}
															if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRule["hotwordRegex"]; ok {
																if rInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex, ok := rInspectJobInspectConfigRuleSetRulesHotwordRule["hotwordRegex"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigRuleSetRules.HotwordRule.HotwordRegex = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex["groupIndexes"]; ok {
																		if s, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex["groupIndexes"].([]interface{}); ok {
																			for _, ss := range s {
																				if intval, ok := ss.(int64); ok {
																					rInspectJobInspectConfigRuleSetRules.HotwordRule.HotwordRegex.GroupIndexes = append(rInspectJobInspectConfigRuleSetRules.HotwordRule.HotwordRegex.GroupIndexes, intval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.HotwordRegex.GroupIndexes: expected []interface{}")
																		}
																	}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex["pattern"]; ok {
																		if s, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex["pattern"].(string); ok {
																			rInspectJobInspectConfigRuleSetRules.HotwordRule.HotwordRegex.Pattern = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.HotwordRegex.Pattern: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.HotwordRegex: expected map[string]interface{}")
																}
															}
															if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRule["likelihoodAdjustment"]; ok {
																if rInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, ok := rInspectJobInspectConfigRuleSetRulesHotwordRule["likelihoodAdjustment"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment["fixedLikelihood"]; ok {
																		if s, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment["fixedLikelihood"].(string); ok {
																			rInspectJobInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment.FixedLikelihood = dclService.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment.FixedLikelihood: expected string")
																		}
																	}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment["relativeLikelihood"]; ok {
																		if i, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment["relativeLikelihood"].(int64); ok {
																			rInspectJobInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment.RelativeLikelihood = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment.RelativeLikelihood: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment: expected map[string]interface{}")
																}
															}
															if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRule["proximity"]; ok {
																if rInspectJobInspectConfigRuleSetRulesHotwordRuleProximity, ok := rInspectJobInspectConfigRuleSetRulesHotwordRule["proximity"].(map[string]interface{}); ok {
																	rInspectJobInspectConfigRuleSetRules.HotwordRule.Proximity = &dclService.JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity{}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleProximity["windowAfter"]; ok {
																		if i, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleProximity["windowAfter"].(int64); ok {
																			rInspectJobInspectConfigRuleSetRules.HotwordRule.Proximity.WindowAfter = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.Proximity.WindowAfter: expected int64")
																		}
																	}
																	if _, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleProximity["windowBefore"]; ok {
																		if i, ok := rInspectJobInspectConfigRuleSetRulesHotwordRuleProximity["windowBefore"].(int64); ok {
																			rInspectJobInspectConfigRuleSetRules.HotwordRule.Proximity.WindowBefore = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.Proximity.WindowBefore: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule.Proximity: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectJobInspectConfigRuleSetRules.HotwordRule: expected map[string]interface{}")
														}
													}
													rInspectJobInspectConfigRuleSet.Rules = append(rInspectJobInspectConfigRuleSet.Rules, rInspectJobInspectConfigRuleSetRules)
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectJobInspectConfigRuleSet.Rules: expected []interface{}")
										}
									}
									r.InspectJob.InspectConfig.RuleSet = append(r.InspectJob.InspectConfig.RuleSet, rInspectJobInspectConfigRuleSet)
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.InspectConfig.RuleSet: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.InspectJob.InspectConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rInspectJob["inspectTemplateName"]; ok {
				if s, ok := rInspectJob["inspectTemplateName"].(string); ok {
					r.InspectJob.InspectTemplateName = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.InspectJob.InspectTemplateName: expected string")
				}
			}
			if _, ok := rInspectJob["storageConfig"]; ok {
				if rInspectJobStorageConfig, ok := rInspectJob["storageConfig"].(map[string]interface{}); ok {
					r.InspectJob.StorageConfig = &dclService.JobTriggerInspectJobStorageConfig{}
					if _, ok := rInspectJobStorageConfig["bigQueryOptions"]; ok {
						if rInspectJobStorageConfigBigQueryOptions, ok := rInspectJobStorageConfig["bigQueryOptions"].(map[string]interface{}); ok {
							r.InspectJob.StorageConfig.BigQueryOptions = &dclService.JobTriggerInspectJobStorageConfigBigQueryOptions{}
							if _, ok := rInspectJobStorageConfigBigQueryOptions["excludedFields"]; ok {
								if s, ok := rInspectJobStorageConfigBigQueryOptions["excludedFields"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rInspectJobStorageConfigBigQueryOptionsExcludedFields dclService.JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields
											if _, ok := objval["name"]; ok {
												if s, ok := objval["name"].(string); ok {
													rInspectJobStorageConfigBigQueryOptionsExcludedFields.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectJobStorageConfigBigQueryOptionsExcludedFields.Name: expected string")
												}
											}
											r.InspectJob.StorageConfig.BigQueryOptions.ExcludedFields = append(r.InspectJob.StorageConfig.BigQueryOptions.ExcludedFields, rInspectJobStorageConfigBigQueryOptionsExcludedFields)
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.ExcludedFields: expected []interface{}")
								}
							}
							if _, ok := rInspectJobStorageConfigBigQueryOptions["identifyingFields"]; ok {
								if s, ok := rInspectJobStorageConfigBigQueryOptions["identifyingFields"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rInspectJobStorageConfigBigQueryOptionsIdentifyingFields dclService.JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields
											if _, ok := objval["name"]; ok {
												if s, ok := objval["name"].(string); ok {
													rInspectJobStorageConfigBigQueryOptionsIdentifyingFields.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectJobStorageConfigBigQueryOptionsIdentifyingFields.Name: expected string")
												}
											}
											r.InspectJob.StorageConfig.BigQueryOptions.IdentifyingFields = append(r.InspectJob.StorageConfig.BigQueryOptions.IdentifyingFields, rInspectJobStorageConfigBigQueryOptionsIdentifyingFields)
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.IdentifyingFields: expected []interface{}")
								}
							}
							if _, ok := rInspectJobStorageConfigBigQueryOptions["includedFields"]; ok {
								if s, ok := rInspectJobStorageConfigBigQueryOptions["includedFields"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rInspectJobStorageConfigBigQueryOptionsIncludedFields dclService.JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields
											if _, ok := objval["name"]; ok {
												if s, ok := objval["name"].(string); ok {
													rInspectJobStorageConfigBigQueryOptionsIncludedFields.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectJobStorageConfigBigQueryOptionsIncludedFields.Name: expected string")
												}
											}
											r.InspectJob.StorageConfig.BigQueryOptions.IncludedFields = append(r.InspectJob.StorageConfig.BigQueryOptions.IncludedFields, rInspectJobStorageConfigBigQueryOptionsIncludedFields)
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.IncludedFields: expected []interface{}")
								}
							}
							if _, ok := rInspectJobStorageConfigBigQueryOptions["rowsLimit"]; ok {
								if i, ok := rInspectJobStorageConfigBigQueryOptions["rowsLimit"].(int64); ok {
									r.InspectJob.StorageConfig.BigQueryOptions.RowsLimit = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.RowsLimit: expected int64")
								}
							}
							if _, ok := rInspectJobStorageConfigBigQueryOptions["rowsLimitPercent"]; ok {
								if i, ok := rInspectJobStorageConfigBigQueryOptions["rowsLimitPercent"].(int64); ok {
									r.InspectJob.StorageConfig.BigQueryOptions.RowsLimitPercent = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.RowsLimitPercent: expected int64")
								}
							}
							if _, ok := rInspectJobStorageConfigBigQueryOptions["sampleMethod"]; ok {
								if s, ok := rInspectJobStorageConfigBigQueryOptions["sampleMethod"].(string); ok {
									r.InspectJob.StorageConfig.BigQueryOptions.SampleMethod = dclService.JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.SampleMethod: expected string")
								}
							}
							if _, ok := rInspectJobStorageConfigBigQueryOptions["tableReference"]; ok {
								if rInspectJobStorageConfigBigQueryOptionsTableReference, ok := rInspectJobStorageConfigBigQueryOptions["tableReference"].(map[string]interface{}); ok {
									r.InspectJob.StorageConfig.BigQueryOptions.TableReference = &dclService.JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference{}
									if _, ok := rInspectJobStorageConfigBigQueryOptionsTableReference["datasetId"]; ok {
										if s, ok := rInspectJobStorageConfigBigQueryOptionsTableReference["datasetId"].(string); ok {
											r.InspectJob.StorageConfig.BigQueryOptions.TableReference.DatasetId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.TableReference.DatasetId: expected string")
										}
									}
									if _, ok := rInspectJobStorageConfigBigQueryOptionsTableReference["projectId"]; ok {
										if s, ok := rInspectJobStorageConfigBigQueryOptionsTableReference["projectId"].(string); ok {
											r.InspectJob.StorageConfig.BigQueryOptions.TableReference.ProjectId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.TableReference.ProjectId: expected string")
										}
									}
									if _, ok := rInspectJobStorageConfigBigQueryOptionsTableReference["tableId"]; ok {
										if s, ok := rInspectJobStorageConfigBigQueryOptionsTableReference["tableId"].(string); ok {
											r.InspectJob.StorageConfig.BigQueryOptions.TableReference.TableId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.TableReference.TableId: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions.TableReference: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.StorageConfig.BigQueryOptions: expected map[string]interface{}")
						}
					}
					if _, ok := rInspectJobStorageConfig["cloudStorageOptions"]; ok {
						if rInspectJobStorageConfigCloudStorageOptions, ok := rInspectJobStorageConfig["cloudStorageOptions"].(map[string]interface{}); ok {
							r.InspectJob.StorageConfig.CloudStorageOptions = &dclService.JobTriggerInspectJobStorageConfigCloudStorageOptions{}
							if _, ok := rInspectJobStorageConfigCloudStorageOptions["bytesLimitPerFile"]; ok {
								if i, ok := rInspectJobStorageConfigCloudStorageOptions["bytesLimitPerFile"].(int64); ok {
									r.InspectJob.StorageConfig.CloudStorageOptions.BytesLimitPerFile = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.BytesLimitPerFile: expected int64")
								}
							}
							if _, ok := rInspectJobStorageConfigCloudStorageOptions["bytesLimitPerFilePercent"]; ok {
								if i, ok := rInspectJobStorageConfigCloudStorageOptions["bytesLimitPerFilePercent"].(int64); ok {
									r.InspectJob.StorageConfig.CloudStorageOptions.BytesLimitPerFilePercent = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.BytesLimitPerFilePercent: expected int64")
								}
							}
							if _, ok := rInspectJobStorageConfigCloudStorageOptions["fileSet"]; ok {
								if rInspectJobStorageConfigCloudStorageOptionsFileSet, ok := rInspectJobStorageConfigCloudStorageOptions["fileSet"].(map[string]interface{}); ok {
									r.InspectJob.StorageConfig.CloudStorageOptions.FileSet = &dclService.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet{}
									if _, ok := rInspectJobStorageConfigCloudStorageOptionsFileSet["regexFileSet"]; ok {
										if rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet, ok := rInspectJobStorageConfigCloudStorageOptionsFileSet["regexFileSet"].(map[string]interface{}); ok {
											r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet = &dclService.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet{}
											if _, ok := rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["bucketName"]; ok {
												if s, ok := rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["bucketName"].(string); ok {
													r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.BucketName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.BucketName: expected string")
												}
											}
											if _, ok := rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["excludeRegex"]; ok {
												if s, ok := rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["excludeRegex"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.ExcludeRegex = append(r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.ExcludeRegex, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.ExcludeRegex: expected []interface{}")
												}
											}
											if _, ok := rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["includeRegex"]; ok {
												if s, ok := rInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet["includeRegex"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.IncludeRegex = append(r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.IncludeRegex, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet.IncludeRegex: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.RegexFileSet: expected map[string]interface{}")
										}
									}
									if _, ok := rInspectJobStorageConfigCloudStorageOptionsFileSet["url"]; ok {
										if s, ok := rInspectJobStorageConfigCloudStorageOptionsFileSet["url"].(string); ok {
											r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.Url = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.FileSet.Url: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.FileSet: expected map[string]interface{}")
								}
							}
							if _, ok := rInspectJobStorageConfigCloudStorageOptions["fileTypes"]; ok {
								if s, ok := rInspectJobStorageConfigCloudStorageOptions["fileTypes"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.InspectJob.StorageConfig.CloudStorageOptions.FileTypes = append(r.InspectJob.StorageConfig.CloudStorageOptions.FileTypes, dclService.JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum(strval))
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.FileTypes: expected []interface{}")
								}
							}
							if _, ok := rInspectJobStorageConfigCloudStorageOptions["filesLimitPercent"]; ok {
								if i, ok := rInspectJobStorageConfigCloudStorageOptions["filesLimitPercent"].(int64); ok {
									r.InspectJob.StorageConfig.CloudStorageOptions.FilesLimitPercent = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.FilesLimitPercent: expected int64")
								}
							}
							if _, ok := rInspectJobStorageConfigCloudStorageOptions["sampleMethod"]; ok {
								if s, ok := rInspectJobStorageConfigCloudStorageOptions["sampleMethod"].(string); ok {
									r.InspectJob.StorageConfig.CloudStorageOptions.SampleMethod = dclService.JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions.SampleMethod: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.StorageConfig.CloudStorageOptions: expected map[string]interface{}")
						}
					}
					if _, ok := rInspectJobStorageConfig["datastoreOptions"]; ok {
						if rInspectJobStorageConfigDatastoreOptions, ok := rInspectJobStorageConfig["datastoreOptions"].(map[string]interface{}); ok {
							r.InspectJob.StorageConfig.DatastoreOptions = &dclService.JobTriggerInspectJobStorageConfigDatastoreOptions{}
							if _, ok := rInspectJobStorageConfigDatastoreOptions["kind"]; ok {
								if rInspectJobStorageConfigDatastoreOptionsKind, ok := rInspectJobStorageConfigDatastoreOptions["kind"].(map[string]interface{}); ok {
									r.InspectJob.StorageConfig.DatastoreOptions.Kind = &dclService.JobTriggerInspectJobStorageConfigDatastoreOptionsKind{}
									if _, ok := rInspectJobStorageConfigDatastoreOptionsKind["name"]; ok {
										if s, ok := rInspectJobStorageConfigDatastoreOptionsKind["name"].(string); ok {
											r.InspectJob.StorageConfig.DatastoreOptions.Kind.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.DatastoreOptions.Kind.Name: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.DatastoreOptions.Kind: expected map[string]interface{}")
								}
							}
							if _, ok := rInspectJobStorageConfigDatastoreOptions["partitionId"]; ok {
								if rInspectJobStorageConfigDatastoreOptionsPartitionId, ok := rInspectJobStorageConfigDatastoreOptions["partitionId"].(map[string]interface{}); ok {
									r.InspectJob.StorageConfig.DatastoreOptions.PartitionId = &dclService.JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId{}
									if _, ok := rInspectJobStorageConfigDatastoreOptionsPartitionId["namespaceId"]; ok {
										if s, ok := rInspectJobStorageConfigDatastoreOptionsPartitionId["namespaceId"].(string); ok {
											r.InspectJob.StorageConfig.DatastoreOptions.PartitionId.NamespaceId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.DatastoreOptions.PartitionId.NamespaceId: expected string")
										}
									}
									if _, ok := rInspectJobStorageConfigDatastoreOptionsPartitionId["projectId"]; ok {
										if s, ok := rInspectJobStorageConfigDatastoreOptionsPartitionId["projectId"].(string); ok {
											r.InspectJob.StorageConfig.DatastoreOptions.PartitionId.ProjectId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.DatastoreOptions.PartitionId.ProjectId: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.DatastoreOptions.PartitionId: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.StorageConfig.DatastoreOptions: expected map[string]interface{}")
						}
					}
					if _, ok := rInspectJobStorageConfig["hybridOptions"]; ok {
						if rInspectJobStorageConfigHybridOptions, ok := rInspectJobStorageConfig["hybridOptions"].(map[string]interface{}); ok {
							r.InspectJob.StorageConfig.HybridOptions = &dclService.JobTriggerInspectJobStorageConfigHybridOptions{}
							if _, ok := rInspectJobStorageConfigHybridOptions["description"]; ok {
								if s, ok := rInspectJobStorageConfigHybridOptions["description"].(string); ok {
									r.InspectJob.StorageConfig.HybridOptions.Description = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.HybridOptions.Description: expected string")
								}
							}
							if _, ok := rInspectJobStorageConfigHybridOptions["labels"]; ok {
								if rInspectJobStorageConfigHybridOptionsLabels, ok := rInspectJobStorageConfigHybridOptions["labels"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rInspectJobStorageConfigHybridOptionsLabels {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									r.InspectJob.StorageConfig.HybridOptions.Labels = m
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.HybridOptions.Labels: expected map[string]interface{}")
								}
							}
							if _, ok := rInspectJobStorageConfigHybridOptions["requiredFindingLabelKeys"]; ok {
								if s, ok := rInspectJobStorageConfigHybridOptions["requiredFindingLabelKeys"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.InspectJob.StorageConfig.HybridOptions.RequiredFindingLabelKeys = append(r.InspectJob.StorageConfig.HybridOptions.RequiredFindingLabelKeys, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.HybridOptions.RequiredFindingLabelKeys: expected []interface{}")
								}
							}
							if _, ok := rInspectJobStorageConfigHybridOptions["tableOptions"]; ok {
								if rInspectJobStorageConfigHybridOptionsTableOptions, ok := rInspectJobStorageConfigHybridOptions["tableOptions"].(map[string]interface{}); ok {
									r.InspectJob.StorageConfig.HybridOptions.TableOptions = &dclService.JobTriggerInspectJobStorageConfigHybridOptionsTableOptions{}
									if _, ok := rInspectJobStorageConfigHybridOptionsTableOptions["identifyingFields"]; ok {
										if s, ok := rInspectJobStorageConfigHybridOptionsTableOptions["identifyingFields"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields dclService.JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields
													if _, ok := objval["name"]; ok {
														if s, ok := objval["name"].(string); ok {
															rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields.Name = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields.Name: expected string")
														}
													}
													r.InspectJob.StorageConfig.HybridOptions.TableOptions.IdentifyingFields = append(r.InspectJob.StorageConfig.HybridOptions.TableOptions.IdentifyingFields, rInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields)
												}
											}
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.HybridOptions.TableOptions.IdentifyingFields: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.HybridOptions.TableOptions: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.StorageConfig.HybridOptions: expected map[string]interface{}")
						}
					}
					if _, ok := rInspectJobStorageConfig["timespanConfig"]; ok {
						if rInspectJobStorageConfigTimespanConfig, ok := rInspectJobStorageConfig["timespanConfig"].(map[string]interface{}); ok {
							r.InspectJob.StorageConfig.TimespanConfig = &dclService.JobTriggerInspectJobStorageConfigTimespanConfig{}
							if _, ok := rInspectJobStorageConfigTimespanConfig["enableAutoPopulationOfTimespanConfig"]; ok {
								if b, ok := rInspectJobStorageConfigTimespanConfig["enableAutoPopulationOfTimespanConfig"].(bool); ok {
									r.InspectJob.StorageConfig.TimespanConfig.EnableAutoPopulationOfTimespanConfig = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.TimespanConfig.EnableAutoPopulationOfTimespanConfig: expected bool")
								}
							}
							if _, ok := rInspectJobStorageConfigTimespanConfig["endTime"]; ok {
								if s, ok := rInspectJobStorageConfigTimespanConfig["endTime"].(string); ok {
									r.InspectJob.StorageConfig.TimespanConfig.EndTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.TimespanConfig.EndTime: expected string")
								}
							}
							if _, ok := rInspectJobStorageConfigTimespanConfig["startTime"]; ok {
								if s, ok := rInspectJobStorageConfigTimespanConfig["startTime"].(string); ok {
									r.InspectJob.StorageConfig.TimespanConfig.StartTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.TimespanConfig.StartTime: expected string")
								}
							}
							if _, ok := rInspectJobStorageConfigTimespanConfig["timestampField"]; ok {
								if rInspectJobStorageConfigTimespanConfigTimestampField, ok := rInspectJobStorageConfigTimespanConfig["timestampField"].(map[string]interface{}); ok {
									r.InspectJob.StorageConfig.TimespanConfig.TimestampField = &dclService.JobTriggerInspectJobStorageConfigTimespanConfigTimestampField{}
									if _, ok := rInspectJobStorageConfigTimespanConfigTimestampField["name"]; ok {
										if s, ok := rInspectJobStorageConfigTimespanConfigTimestampField["name"].(string); ok {
											r.InspectJob.StorageConfig.TimespanConfig.TimestampField.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("r.InspectJob.StorageConfig.TimespanConfig.TimestampField.Name: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.InspectJob.StorageConfig.TimespanConfig.TimestampField: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectJob.StorageConfig.TimespanConfig: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.InspectJob.StorageConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.InspectJob: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["lastRunTime"]; ok {
		if s, ok := u.Object["lastRunTime"].(string); ok {
			r.LastRunTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LastRunTime: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["locationId"]; ok {
		if s, ok := u.Object["locationId"].(string); ok {
			r.LocationId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LocationId: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["parent"]; ok {
		if s, ok := u.Object["parent"].(string); ok {
			r.Parent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Parent: expected string")
		}
	}
	if _, ok := u.Object["status"]; ok {
		if s, ok := u.Object["status"].(string); ok {
			r.Status = dclService.JobTriggerStatusEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Status: expected string")
		}
	}
	if _, ok := u.Object["triggers"]; ok {
		if s, ok := u.Object["triggers"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rTriggers dclService.JobTriggerTriggers
					if _, ok := objval["manual"]; ok {
						if _, ok := objval["manual"].(map[string]interface{}); ok {
							rTriggers.Manual = &dclService.JobTriggerTriggersManual{}
						} else {
							return nil, fmt.Errorf("rTriggers.Manual: expected map[string]interface{}")
						}
					}
					if _, ok := objval["schedule"]; ok {
						if rTriggersSchedule, ok := objval["schedule"].(map[string]interface{}); ok {
							rTriggers.Schedule = &dclService.JobTriggerTriggersSchedule{}
							if _, ok := rTriggersSchedule["recurrencePeriodDuration"]; ok {
								if s, ok := rTriggersSchedule["recurrencePeriodDuration"].(string); ok {
									rTriggers.Schedule.RecurrencePeriodDuration = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rTriggers.Schedule.RecurrencePeriodDuration: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rTriggers.Schedule: expected map[string]interface{}")
						}
					}
					r.Triggers = append(r.Triggers, rTriggers)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Triggers: expected []interface{}")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetJobTrigger(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJobTrigger(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetJobTrigger(ctx, r)
	if err != nil {
		return nil, err
	}
	return JobTriggerToUnstructured(r), nil
}

func ListJobTrigger(ctx context.Context, config *dcl.Config, location string, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListJobTrigger(ctx, location, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, JobTriggerToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyJobTrigger(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJobTrigger(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToJobTrigger(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyJobTrigger(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return JobTriggerToUnstructured(r), nil
}

func JobTriggerHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJobTrigger(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToJobTrigger(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyJobTrigger(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteJobTrigger(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJobTrigger(u)
	if err != nil {
		return err
	}
	return c.DeleteJobTrigger(ctx, r)
}

func JobTriggerID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToJobTrigger(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *JobTrigger) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dlp",
		"JobTrigger",
		"alpha",
	}
}

func (r *JobTrigger) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *JobTrigger) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *JobTrigger) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *JobTrigger) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *JobTrigger) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *JobTrigger) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *JobTrigger) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetJobTrigger(ctx, config, resource)
}

func (r *JobTrigger) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyJobTrigger(ctx, config, resource, opts...)
}

func (r *JobTrigger) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return JobTriggerHasDiff(ctx, config, resource, opts...)
}

func (r *JobTrigger) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteJobTrigger(ctx, config, resource)
}

func (r *JobTrigger) ID(resource *unstructured.Resource) (string, error) {
	return JobTriggerID(resource)
}

func init() {
	unstructured.Register(&JobTrigger{})
}
