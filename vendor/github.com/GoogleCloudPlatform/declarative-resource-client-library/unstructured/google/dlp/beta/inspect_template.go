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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dlp/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type InspectTemplate struct{}

func InspectTemplateToUnstructured(r *dclService.InspectTemplate) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dlp",
			Version: "beta",
			Type:    "InspectTemplate",
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
	if r.InspectConfig != nil && r.InspectConfig != dclService.EmptyInspectTemplateInspectConfig {
		rInspectConfig := make(map[string]interface{})
		var rInspectConfigContentOptions []interface{}
		for _, rInspectConfigContentOptionsVal := range r.InspectConfig.ContentOptions {
			rInspectConfigContentOptions = append(rInspectConfigContentOptions, string(rInspectConfigContentOptionsVal))
		}
		rInspectConfig["contentOptions"] = rInspectConfigContentOptions
		var rInspectConfigCustomInfoTypes []interface{}
		for _, rInspectConfigCustomInfoTypesVal := range r.InspectConfig.CustomInfoTypes {
			rInspectConfigCustomInfoTypesObject := make(map[string]interface{})
			if rInspectConfigCustomInfoTypesVal.Dictionary != nil && rInspectConfigCustomInfoTypesVal.Dictionary != dclService.EmptyInspectTemplateInspectConfigCustomInfoTypesDictionary {
				rInspectConfigCustomInfoTypesValDictionary := make(map[string]interface{})
				if rInspectConfigCustomInfoTypesVal.Dictionary.CloudStoragePath != nil && rInspectConfigCustomInfoTypesVal.Dictionary.CloudStoragePath != dclService.EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
					rInspectConfigCustomInfoTypesValDictionaryCloudStoragePath := make(map[string]interface{})
					if rInspectConfigCustomInfoTypesVal.Dictionary.CloudStoragePath.Path != nil {
						rInspectConfigCustomInfoTypesValDictionaryCloudStoragePath["path"] = *rInspectConfigCustomInfoTypesVal.Dictionary.CloudStoragePath.Path
					}
					rInspectConfigCustomInfoTypesValDictionary["cloudStoragePath"] = rInspectConfigCustomInfoTypesValDictionaryCloudStoragePath
				}
				if rInspectConfigCustomInfoTypesVal.Dictionary.WordList != nil && rInspectConfigCustomInfoTypesVal.Dictionary.WordList != dclService.EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
					rInspectConfigCustomInfoTypesValDictionaryWordList := make(map[string]interface{})
					var rInspectConfigCustomInfoTypesValDictionaryWordListWords []interface{}
					for _, rInspectConfigCustomInfoTypesValDictionaryWordListWordsVal := range rInspectConfigCustomInfoTypesVal.Dictionary.WordList.Words {
						rInspectConfigCustomInfoTypesValDictionaryWordListWords = append(rInspectConfigCustomInfoTypesValDictionaryWordListWords, rInspectConfigCustomInfoTypesValDictionaryWordListWordsVal)
					}
					rInspectConfigCustomInfoTypesValDictionaryWordList["words"] = rInspectConfigCustomInfoTypesValDictionaryWordListWords
					rInspectConfigCustomInfoTypesValDictionary["wordList"] = rInspectConfigCustomInfoTypesValDictionaryWordList
				}
				rInspectConfigCustomInfoTypesObject["dictionary"] = rInspectConfigCustomInfoTypesValDictionary
			}
			if rInspectConfigCustomInfoTypesVal.ExclusionType != nil {
				rInspectConfigCustomInfoTypesObject["exclusionType"] = string(*rInspectConfigCustomInfoTypesVal.ExclusionType)
			}
			if rInspectConfigCustomInfoTypesVal.InfoType != nil && rInspectConfigCustomInfoTypesVal.InfoType != dclService.EmptyInspectTemplateInspectConfigCustomInfoTypesInfoType {
				rInspectConfigCustomInfoTypesValInfoType := make(map[string]interface{})
				if rInspectConfigCustomInfoTypesVal.InfoType.Name != nil {
					rInspectConfigCustomInfoTypesValInfoType["name"] = *rInspectConfigCustomInfoTypesVal.InfoType.Name
				}
				rInspectConfigCustomInfoTypesObject["infoType"] = rInspectConfigCustomInfoTypesValInfoType
			}
			if rInspectConfigCustomInfoTypesVal.Likelihood != nil {
				rInspectConfigCustomInfoTypesObject["likelihood"] = string(*rInspectConfigCustomInfoTypesVal.Likelihood)
			}
			if rInspectConfigCustomInfoTypesVal.Regex != nil && rInspectConfigCustomInfoTypesVal.Regex != dclService.EmptyInspectTemplateInspectConfigCustomInfoTypesRegex {
				rInspectConfigCustomInfoTypesValRegex := make(map[string]interface{})
				var rInspectConfigCustomInfoTypesValRegexGroupIndexes []interface{}
				for _, rInspectConfigCustomInfoTypesValRegexGroupIndexesVal := range rInspectConfigCustomInfoTypesVal.Regex.GroupIndexes {
					rInspectConfigCustomInfoTypesValRegexGroupIndexes = append(rInspectConfigCustomInfoTypesValRegexGroupIndexes, rInspectConfigCustomInfoTypesValRegexGroupIndexesVal)
				}
				rInspectConfigCustomInfoTypesValRegex["groupIndexes"] = rInspectConfigCustomInfoTypesValRegexGroupIndexes
				if rInspectConfigCustomInfoTypesVal.Regex.Pattern != nil {
					rInspectConfigCustomInfoTypesValRegex["pattern"] = *rInspectConfigCustomInfoTypesVal.Regex.Pattern
				}
				rInspectConfigCustomInfoTypesObject["regex"] = rInspectConfigCustomInfoTypesValRegex
			}
			if rInspectConfigCustomInfoTypesVal.StoredType != nil && rInspectConfigCustomInfoTypesVal.StoredType != dclService.EmptyInspectTemplateInspectConfigCustomInfoTypesStoredType {
				rInspectConfigCustomInfoTypesValStoredType := make(map[string]interface{})
				if rInspectConfigCustomInfoTypesVal.StoredType.CreateTime != nil {
					rInspectConfigCustomInfoTypesValStoredType["createTime"] = *rInspectConfigCustomInfoTypesVal.StoredType.CreateTime
				}
				if rInspectConfigCustomInfoTypesVal.StoredType.Name != nil {
					rInspectConfigCustomInfoTypesValStoredType["name"] = *rInspectConfigCustomInfoTypesVal.StoredType.Name
				}
				rInspectConfigCustomInfoTypesObject["storedType"] = rInspectConfigCustomInfoTypesValStoredType
			}
			if rInspectConfigCustomInfoTypesVal.SurrogateType != nil && rInspectConfigCustomInfoTypesVal.SurrogateType != dclService.EmptyInspectTemplateInspectConfigCustomInfoTypesSurrogateType {
				rInspectConfigCustomInfoTypesValSurrogateType := make(map[string]interface{})
				rInspectConfigCustomInfoTypesObject["surrogateType"] = rInspectConfigCustomInfoTypesValSurrogateType
			}
			rInspectConfigCustomInfoTypes = append(rInspectConfigCustomInfoTypes, rInspectConfigCustomInfoTypesObject)
		}
		rInspectConfig["customInfoTypes"] = rInspectConfigCustomInfoTypes
		if r.InspectConfig.ExcludeInfoTypes != nil {
			rInspectConfig["excludeInfoTypes"] = *r.InspectConfig.ExcludeInfoTypes
		}
		if r.InspectConfig.IncludeQuote != nil {
			rInspectConfig["includeQuote"] = *r.InspectConfig.IncludeQuote
		}
		var rInspectConfigInfoTypes []interface{}
		for _, rInspectConfigInfoTypesVal := range r.InspectConfig.InfoTypes {
			rInspectConfigInfoTypesObject := make(map[string]interface{})
			if rInspectConfigInfoTypesVal.Name != nil {
				rInspectConfigInfoTypesObject["name"] = *rInspectConfigInfoTypesVal.Name
			}
			rInspectConfigInfoTypes = append(rInspectConfigInfoTypes, rInspectConfigInfoTypesObject)
		}
		rInspectConfig["infoTypes"] = rInspectConfigInfoTypes
		if r.InspectConfig.Limits != nil && r.InspectConfig.Limits != dclService.EmptyInspectTemplateInspectConfigLimits {
			rInspectConfigLimits := make(map[string]interface{})
			var rInspectConfigLimitsMaxFindingsPerInfoType []interface{}
			for _, rInspectConfigLimitsMaxFindingsPerInfoTypeVal := range r.InspectConfig.Limits.MaxFindingsPerInfoType {
				rInspectConfigLimitsMaxFindingsPerInfoTypeObject := make(map[string]interface{})
				if rInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType != nil && rInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType != dclService.EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
					rInspectConfigLimitsMaxFindingsPerInfoTypeValInfoType := make(map[string]interface{})
					if rInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType.Name != nil {
						rInspectConfigLimitsMaxFindingsPerInfoTypeValInfoType["name"] = *rInspectConfigLimitsMaxFindingsPerInfoTypeVal.InfoType.Name
					}
					rInspectConfigLimitsMaxFindingsPerInfoTypeObject["infoType"] = rInspectConfigLimitsMaxFindingsPerInfoTypeValInfoType
				}
				if rInspectConfigLimitsMaxFindingsPerInfoTypeVal.MaxFindings != nil {
					rInspectConfigLimitsMaxFindingsPerInfoTypeObject["maxFindings"] = *rInspectConfigLimitsMaxFindingsPerInfoTypeVal.MaxFindings
				}
				rInspectConfigLimitsMaxFindingsPerInfoType = append(rInspectConfigLimitsMaxFindingsPerInfoType, rInspectConfigLimitsMaxFindingsPerInfoTypeObject)
			}
			rInspectConfigLimits["maxFindingsPerInfoType"] = rInspectConfigLimitsMaxFindingsPerInfoType
			if r.InspectConfig.Limits.MaxFindingsPerItem != nil {
				rInspectConfigLimits["maxFindingsPerItem"] = *r.InspectConfig.Limits.MaxFindingsPerItem
			}
			if r.InspectConfig.Limits.MaxFindingsPerRequest != nil {
				rInspectConfigLimits["maxFindingsPerRequest"] = *r.InspectConfig.Limits.MaxFindingsPerRequest
			}
			rInspectConfig["limits"] = rInspectConfigLimits
		}
		if r.InspectConfig.MinLikelihood != nil {
			rInspectConfig["minLikelihood"] = string(*r.InspectConfig.MinLikelihood)
		}
		var rInspectConfigRuleSet []interface{}
		for _, rInspectConfigRuleSetVal := range r.InspectConfig.RuleSet {
			rInspectConfigRuleSetObject := make(map[string]interface{})
			var rInspectConfigRuleSetValInfoTypes []interface{}
			for _, rInspectConfigRuleSetValInfoTypesVal := range rInspectConfigRuleSetVal.InfoTypes {
				rInspectConfigRuleSetValInfoTypesObject := make(map[string]interface{})
				if rInspectConfigRuleSetValInfoTypesVal.Name != nil {
					rInspectConfigRuleSetValInfoTypesObject["name"] = *rInspectConfigRuleSetValInfoTypesVal.Name
				}
				rInspectConfigRuleSetValInfoTypes = append(rInspectConfigRuleSetValInfoTypes, rInspectConfigRuleSetValInfoTypesObject)
			}
			rInspectConfigRuleSetObject["infoTypes"] = rInspectConfigRuleSetValInfoTypes
			var rInspectConfigRuleSetValRules []interface{}
			for _, rInspectConfigRuleSetValRulesVal := range rInspectConfigRuleSetVal.Rules {
				rInspectConfigRuleSetValRulesObject := make(map[string]interface{})
				if rInspectConfigRuleSetValRulesVal.ExclusionRule != nil && rInspectConfigRuleSetValRulesVal.ExclusionRule != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRule {
					rInspectConfigRuleSetValRulesValExclusionRule := make(map[string]interface{})
					if rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary != nil && rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
						rInspectConfigRuleSetValRulesValExclusionRuleDictionary := make(map[string]interface{})
						if rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.CloudStoragePath != nil && rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.CloudStoragePath != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
							rInspectConfigRuleSetValRulesValExclusionRuleDictionaryCloudStoragePath := make(map[string]interface{})
							if rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.CloudStoragePath.Path != nil {
								rInspectConfigRuleSetValRulesValExclusionRuleDictionaryCloudStoragePath["path"] = *rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.CloudStoragePath.Path
							}
							rInspectConfigRuleSetValRulesValExclusionRuleDictionary["cloudStoragePath"] = rInspectConfigRuleSetValRulesValExclusionRuleDictionaryCloudStoragePath
						}
						if rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.WordList != nil && rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.WordList != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
							rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordList := make(map[string]interface{})
							var rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWords []interface{}
							for _, rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWordsVal := range rInspectConfigRuleSetValRulesVal.ExclusionRule.Dictionary.WordList.Words {
								rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWords = append(rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWords, rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWordsVal)
							}
							rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordList["words"] = rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordListWords
							rInspectConfigRuleSetValRulesValExclusionRuleDictionary["wordList"] = rInspectConfigRuleSetValRulesValExclusionRuleDictionaryWordList
						}
						rInspectConfigRuleSetValRulesValExclusionRule["dictionary"] = rInspectConfigRuleSetValRulesValExclusionRuleDictionary
					}
					if rInspectConfigRuleSetValRulesVal.ExclusionRule.ExcludeInfoTypes != nil && rInspectConfigRuleSetValRulesVal.ExclusionRule.ExcludeInfoTypes != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
						rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypes := make(map[string]interface{})
						var rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypes []interface{}
						for _, rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesVal := range rInspectConfigRuleSetValRulesVal.ExclusionRule.ExcludeInfoTypes.InfoTypes {
							rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesObject := make(map[string]interface{})
							if rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesVal.Name != nil {
								rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesObject["name"] = *rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesVal.Name
							}
							rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypes = append(rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypes, rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypesObject)
						}
						rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypes["infoTypes"] = rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypesInfoTypes
						rInspectConfigRuleSetValRulesValExclusionRule["excludeInfoTypes"] = rInspectConfigRuleSetValRulesValExclusionRuleExcludeInfoTypes
					}
					if rInspectConfigRuleSetValRulesVal.ExclusionRule.MatchingType != nil {
						rInspectConfigRuleSetValRulesValExclusionRule["matchingType"] = string(*rInspectConfigRuleSetValRulesVal.ExclusionRule.MatchingType)
					}
					if rInspectConfigRuleSetValRulesVal.ExclusionRule.Regex != nil && rInspectConfigRuleSetValRulesVal.ExclusionRule.Regex != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
						rInspectConfigRuleSetValRulesValExclusionRuleRegex := make(map[string]interface{})
						var rInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexes []interface{}
						for _, rInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexesVal := range rInspectConfigRuleSetValRulesVal.ExclusionRule.Regex.GroupIndexes {
							rInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexes = append(rInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexes, rInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexesVal)
						}
						rInspectConfigRuleSetValRulesValExclusionRuleRegex["groupIndexes"] = rInspectConfigRuleSetValRulesValExclusionRuleRegexGroupIndexes
						if rInspectConfigRuleSetValRulesVal.ExclusionRule.Regex.Pattern != nil {
							rInspectConfigRuleSetValRulesValExclusionRuleRegex["pattern"] = *rInspectConfigRuleSetValRulesVal.ExclusionRule.Regex.Pattern
						}
						rInspectConfigRuleSetValRulesValExclusionRule["regex"] = rInspectConfigRuleSetValRulesValExclusionRuleRegex
					}
					rInspectConfigRuleSetValRulesObject["exclusionRule"] = rInspectConfigRuleSetValRulesValExclusionRule
				}
				if rInspectConfigRuleSetValRulesVal.HotwordRule != nil && rInspectConfigRuleSetValRulesVal.HotwordRule != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRule {
					rInspectConfigRuleSetValRulesValHotwordRule := make(map[string]interface{})
					if rInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex != nil && rInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
						rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegex := make(map[string]interface{})
						var rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexes []interface{}
						for _, rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexesVal := range rInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex.GroupIndexes {
							rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexes = append(rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexes, rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexesVal)
						}
						rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegex["groupIndexes"] = rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegexGroupIndexes
						if rInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex.Pattern != nil {
							rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegex["pattern"] = *rInspectConfigRuleSetValRulesVal.HotwordRule.HotwordRegex.Pattern
						}
						rInspectConfigRuleSetValRulesValHotwordRule["hotwordRegex"] = rInspectConfigRuleSetValRulesValHotwordRuleHotwordRegex
					}
					if rInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment != nil && rInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
						rInspectConfigRuleSetValRulesValHotwordRuleLikelihoodAdjustment := make(map[string]interface{})
						if rInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment.FixedLikelihood != nil {
							rInspectConfigRuleSetValRulesValHotwordRuleLikelihoodAdjustment["fixedLikelihood"] = string(*rInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment.FixedLikelihood)
						}
						if rInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment.RelativeLikelihood != nil {
							rInspectConfigRuleSetValRulesValHotwordRuleLikelihoodAdjustment["relativeLikelihood"] = *rInspectConfigRuleSetValRulesVal.HotwordRule.LikelihoodAdjustment.RelativeLikelihood
						}
						rInspectConfigRuleSetValRulesValHotwordRule["likelihoodAdjustment"] = rInspectConfigRuleSetValRulesValHotwordRuleLikelihoodAdjustment
					}
					if rInspectConfigRuleSetValRulesVal.HotwordRule.Proximity != nil && rInspectConfigRuleSetValRulesVal.HotwordRule.Proximity != dclService.EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
						rInspectConfigRuleSetValRulesValHotwordRuleProximity := make(map[string]interface{})
						if rInspectConfigRuleSetValRulesVal.HotwordRule.Proximity.WindowAfter != nil {
							rInspectConfigRuleSetValRulesValHotwordRuleProximity["windowAfter"] = *rInspectConfigRuleSetValRulesVal.HotwordRule.Proximity.WindowAfter
						}
						if rInspectConfigRuleSetValRulesVal.HotwordRule.Proximity.WindowBefore != nil {
							rInspectConfigRuleSetValRulesValHotwordRuleProximity["windowBefore"] = *rInspectConfigRuleSetValRulesVal.HotwordRule.Proximity.WindowBefore
						}
						rInspectConfigRuleSetValRulesValHotwordRule["proximity"] = rInspectConfigRuleSetValRulesValHotwordRuleProximity
					}
					rInspectConfigRuleSetValRulesObject["hotwordRule"] = rInspectConfigRuleSetValRulesValHotwordRule
				}
				rInspectConfigRuleSetValRules = append(rInspectConfigRuleSetValRules, rInspectConfigRuleSetValRulesObject)
			}
			rInspectConfigRuleSetObject["rules"] = rInspectConfigRuleSetValRules
			rInspectConfigRuleSet = append(rInspectConfigRuleSet, rInspectConfigRuleSetObject)
		}
		rInspectConfig["ruleSet"] = rInspectConfigRuleSet
		u.Object["inspectConfig"] = rInspectConfig
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
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToInspectTemplate(u *unstructured.Resource) (*dclService.InspectTemplate, error) {
	r := &dclService.InspectTemplate{}
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
	if _, ok := u.Object["inspectConfig"]; ok {
		if rInspectConfig, ok := u.Object["inspectConfig"].(map[string]interface{}); ok {
			r.InspectConfig = &dclService.InspectTemplateInspectConfig{}
			if _, ok := rInspectConfig["contentOptions"]; ok {
				if s, ok := rInspectConfig["contentOptions"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.InspectConfig.ContentOptions = append(r.InspectConfig.ContentOptions, dclService.InspectTemplateInspectConfigContentOptionsEnum(strval))
						}
					}
				} else {
					return nil, fmt.Errorf("r.InspectConfig.ContentOptions: expected []interface{}")
				}
			}
			if _, ok := rInspectConfig["customInfoTypes"]; ok {
				if s, ok := rInspectConfig["customInfoTypes"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rInspectConfigCustomInfoTypes dclService.InspectTemplateInspectConfigCustomInfoTypes
							if _, ok := objval["dictionary"]; ok {
								if rInspectConfigCustomInfoTypesDictionary, ok := objval["dictionary"].(map[string]interface{}); ok {
									rInspectConfigCustomInfoTypes.Dictionary = &dclService.InspectTemplateInspectConfigCustomInfoTypesDictionary{}
									if _, ok := rInspectConfigCustomInfoTypesDictionary["cloudStoragePath"]; ok {
										if rInspectConfigCustomInfoTypesDictionaryCloudStoragePath, ok := rInspectConfigCustomInfoTypesDictionary["cloudStoragePath"].(map[string]interface{}); ok {
											rInspectConfigCustomInfoTypes.Dictionary.CloudStoragePath = &dclService.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
											if _, ok := rInspectConfigCustomInfoTypesDictionaryCloudStoragePath["path"]; ok {
												if s, ok := rInspectConfigCustomInfoTypesDictionaryCloudStoragePath["path"].(string); ok {
													rInspectConfigCustomInfoTypes.Dictionary.CloudStoragePath.Path = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Dictionary.CloudStoragePath.Path: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Dictionary.CloudStoragePath: expected map[string]interface{}")
										}
									}
									if _, ok := rInspectConfigCustomInfoTypesDictionary["wordList"]; ok {
										if rInspectConfigCustomInfoTypesDictionaryWordList, ok := rInspectConfigCustomInfoTypesDictionary["wordList"].(map[string]interface{}); ok {
											rInspectConfigCustomInfoTypes.Dictionary.WordList = &dclService.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
											if _, ok := rInspectConfigCustomInfoTypesDictionaryWordList["words"]; ok {
												if s, ok := rInspectConfigCustomInfoTypesDictionaryWordList["words"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															rInspectConfigCustomInfoTypes.Dictionary.WordList.Words = append(rInspectConfigCustomInfoTypes.Dictionary.WordList.Words, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Dictionary.WordList.Words: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Dictionary.WordList: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Dictionary: expected map[string]interface{}")
								}
							}
							if _, ok := objval["exclusionType"]; ok {
								if s, ok := objval["exclusionType"].(string); ok {
									rInspectConfigCustomInfoTypes.ExclusionType = dclService.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumRef(s)
								} else {
									return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.ExclusionType: expected string")
								}
							}
							if _, ok := objval["infoType"]; ok {
								if rInspectConfigCustomInfoTypesInfoType, ok := objval["infoType"].(map[string]interface{}); ok {
									rInspectConfigCustomInfoTypes.InfoType = &dclService.InspectTemplateInspectConfigCustomInfoTypesInfoType{}
									if _, ok := rInspectConfigCustomInfoTypesInfoType["name"]; ok {
										if s, ok := rInspectConfigCustomInfoTypesInfoType["name"].(string); ok {
											rInspectConfigCustomInfoTypes.InfoType.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.InfoType.Name: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.InfoType: expected map[string]interface{}")
								}
							}
							if _, ok := objval["likelihood"]; ok {
								if s, ok := objval["likelihood"].(string); ok {
									rInspectConfigCustomInfoTypes.Likelihood = dclService.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumRef(s)
								} else {
									return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Likelihood: expected string")
								}
							}
							if _, ok := objval["regex"]; ok {
								if rInspectConfigCustomInfoTypesRegex, ok := objval["regex"].(map[string]interface{}); ok {
									rInspectConfigCustomInfoTypes.Regex = &dclService.InspectTemplateInspectConfigCustomInfoTypesRegex{}
									if _, ok := rInspectConfigCustomInfoTypesRegex["groupIndexes"]; ok {
										if s, ok := rInspectConfigCustomInfoTypesRegex["groupIndexes"].([]interface{}); ok {
											for _, ss := range s {
												if intval, ok := ss.(int64); ok {
													rInspectConfigCustomInfoTypes.Regex.GroupIndexes = append(rInspectConfigCustomInfoTypes.Regex.GroupIndexes, intval)
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Regex.GroupIndexes: expected []interface{}")
										}
									}
									if _, ok := rInspectConfigCustomInfoTypesRegex["pattern"]; ok {
										if s, ok := rInspectConfigCustomInfoTypesRegex["pattern"].(string); ok {
											rInspectConfigCustomInfoTypes.Regex.Pattern = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Regex.Pattern: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.Regex: expected map[string]interface{}")
								}
							}
							if _, ok := objval["storedType"]; ok {
								if rInspectConfigCustomInfoTypesStoredType, ok := objval["storedType"].(map[string]interface{}); ok {
									rInspectConfigCustomInfoTypes.StoredType = &dclService.InspectTemplateInspectConfigCustomInfoTypesStoredType{}
									if _, ok := rInspectConfigCustomInfoTypesStoredType["createTime"]; ok {
										if s, ok := rInspectConfigCustomInfoTypesStoredType["createTime"].(string); ok {
											rInspectConfigCustomInfoTypes.StoredType.CreateTime = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.StoredType.CreateTime: expected string")
										}
									}
									if _, ok := rInspectConfigCustomInfoTypesStoredType["name"]; ok {
										if s, ok := rInspectConfigCustomInfoTypesStoredType["name"].(string); ok {
											rInspectConfigCustomInfoTypes.StoredType.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.StoredType.Name: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.StoredType: expected map[string]interface{}")
								}
							}
							if _, ok := objval["surrogateType"]; ok {
								if _, ok := objval["surrogateType"].(map[string]interface{}); ok {
									rInspectConfigCustomInfoTypes.SurrogateType = &dclService.InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
								} else {
									return nil, fmt.Errorf("rInspectConfigCustomInfoTypes.SurrogateType: expected map[string]interface{}")
								}
							}
							r.InspectConfig.CustomInfoTypes = append(r.InspectConfig.CustomInfoTypes, rInspectConfigCustomInfoTypes)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InspectConfig.CustomInfoTypes: expected []interface{}")
				}
			}
			if _, ok := rInspectConfig["excludeInfoTypes"]; ok {
				if b, ok := rInspectConfig["excludeInfoTypes"].(bool); ok {
					r.InspectConfig.ExcludeInfoTypes = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.InspectConfig.ExcludeInfoTypes: expected bool")
				}
			}
			if _, ok := rInspectConfig["includeQuote"]; ok {
				if b, ok := rInspectConfig["includeQuote"].(bool); ok {
					r.InspectConfig.IncludeQuote = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.InspectConfig.IncludeQuote: expected bool")
				}
			}
			if _, ok := rInspectConfig["infoTypes"]; ok {
				if s, ok := rInspectConfig["infoTypes"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rInspectConfigInfoTypes dclService.InspectTemplateInspectConfigInfoTypes
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rInspectConfigInfoTypes.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rInspectConfigInfoTypes.Name: expected string")
								}
							}
							r.InspectConfig.InfoTypes = append(r.InspectConfig.InfoTypes, rInspectConfigInfoTypes)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InspectConfig.InfoTypes: expected []interface{}")
				}
			}
			if _, ok := rInspectConfig["limits"]; ok {
				if rInspectConfigLimits, ok := rInspectConfig["limits"].(map[string]interface{}); ok {
					r.InspectConfig.Limits = &dclService.InspectTemplateInspectConfigLimits{}
					if _, ok := rInspectConfigLimits["maxFindingsPerInfoType"]; ok {
						if s, ok := rInspectConfigLimits["maxFindingsPerInfoType"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rInspectConfigLimitsMaxFindingsPerInfoType dclService.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType
									if _, ok := objval["infoType"]; ok {
										if rInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, ok := objval["infoType"].(map[string]interface{}); ok {
											rInspectConfigLimitsMaxFindingsPerInfoType.InfoType = &dclService.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
											if _, ok := rInspectConfigLimitsMaxFindingsPerInfoTypeInfoType["name"]; ok {
												if s, ok := rInspectConfigLimitsMaxFindingsPerInfoTypeInfoType["name"].(string); ok {
													rInspectConfigLimitsMaxFindingsPerInfoType.InfoType.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectConfigLimitsMaxFindingsPerInfoType.InfoType.Name: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rInspectConfigLimitsMaxFindingsPerInfoType.InfoType: expected map[string]interface{}")
										}
									}
									if _, ok := objval["maxFindings"]; ok {
										if i, ok := objval["maxFindings"].(int64); ok {
											rInspectConfigLimitsMaxFindingsPerInfoType.MaxFindings = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rInspectConfigLimitsMaxFindingsPerInfoType.MaxFindings: expected int64")
										}
									}
									r.InspectConfig.Limits.MaxFindingsPerInfoType = append(r.InspectConfig.Limits.MaxFindingsPerInfoType, rInspectConfigLimitsMaxFindingsPerInfoType)
								}
							}
						} else {
							return nil, fmt.Errorf("r.InspectConfig.Limits.MaxFindingsPerInfoType: expected []interface{}")
						}
					}
					if _, ok := rInspectConfigLimits["maxFindingsPerItem"]; ok {
						if i, ok := rInspectConfigLimits["maxFindingsPerItem"].(int64); ok {
							r.InspectConfig.Limits.MaxFindingsPerItem = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.InspectConfig.Limits.MaxFindingsPerItem: expected int64")
						}
					}
					if _, ok := rInspectConfigLimits["maxFindingsPerRequest"]; ok {
						if i, ok := rInspectConfigLimits["maxFindingsPerRequest"].(int64); ok {
							r.InspectConfig.Limits.MaxFindingsPerRequest = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.InspectConfig.Limits.MaxFindingsPerRequest: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.InspectConfig.Limits: expected map[string]interface{}")
				}
			}
			if _, ok := rInspectConfig["minLikelihood"]; ok {
				if s, ok := rInspectConfig["minLikelihood"].(string); ok {
					r.InspectConfig.MinLikelihood = dclService.InspectTemplateInspectConfigMinLikelihoodEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.InspectConfig.MinLikelihood: expected string")
				}
			}
			if _, ok := rInspectConfig["ruleSet"]; ok {
				if s, ok := rInspectConfig["ruleSet"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rInspectConfigRuleSet dclService.InspectTemplateInspectConfigRuleSet
							if _, ok := objval["infoTypes"]; ok {
								if s, ok := objval["infoTypes"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rInspectConfigRuleSetInfoTypes dclService.InspectTemplateInspectConfigRuleSetInfoTypes
											if _, ok := objval["name"]; ok {
												if s, ok := objval["name"].(string); ok {
													rInspectConfigRuleSetInfoTypes.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rInspectConfigRuleSetInfoTypes.Name: expected string")
												}
											}
											rInspectConfigRuleSet.InfoTypes = append(rInspectConfigRuleSet.InfoTypes, rInspectConfigRuleSetInfoTypes)
										}
									}
								} else {
									return nil, fmt.Errorf("rInspectConfigRuleSet.InfoTypes: expected []interface{}")
								}
							}
							if _, ok := objval["rules"]; ok {
								if s, ok := objval["rules"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rInspectConfigRuleSetRules dclService.InspectTemplateInspectConfigRuleSetRules
											if _, ok := objval["exclusionRule"]; ok {
												if rInspectConfigRuleSetRulesExclusionRule, ok := objval["exclusionRule"].(map[string]interface{}); ok {
													rInspectConfigRuleSetRules.ExclusionRule = &dclService.InspectTemplateInspectConfigRuleSetRulesExclusionRule{}
													if _, ok := rInspectConfigRuleSetRulesExclusionRule["dictionary"]; ok {
														if rInspectConfigRuleSetRulesExclusionRuleDictionary, ok := rInspectConfigRuleSetRulesExclusionRule["dictionary"].(map[string]interface{}); ok {
															rInspectConfigRuleSetRules.ExclusionRule.Dictionary = &dclService.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
															if _, ok := rInspectConfigRuleSetRulesExclusionRuleDictionary["cloudStoragePath"]; ok {
																if rInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, ok := rInspectConfigRuleSetRulesExclusionRuleDictionary["cloudStoragePath"].(map[string]interface{}); ok {
																	rInspectConfigRuleSetRules.ExclusionRule.Dictionary.CloudStoragePath = &dclService.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
																	if _, ok := rInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath["path"]; ok {
																		if s, ok := rInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath["path"].(string); ok {
																			rInspectConfigRuleSetRules.ExclusionRule.Dictionary.CloudStoragePath.Path = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.Dictionary.CloudStoragePath.Path: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.Dictionary.CloudStoragePath: expected map[string]interface{}")
																}
															}
															if _, ok := rInspectConfigRuleSetRulesExclusionRuleDictionary["wordList"]; ok {
																if rInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, ok := rInspectConfigRuleSetRulesExclusionRuleDictionary["wordList"].(map[string]interface{}); ok {
																	rInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList = &dclService.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
																	if _, ok := rInspectConfigRuleSetRulesExclusionRuleDictionaryWordList["words"]; ok {
																		if s, ok := rInspectConfigRuleSetRulesExclusionRuleDictionaryWordList["words"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList.Words = append(rInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList.Words, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList.Words: expected []interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.Dictionary.WordList: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.Dictionary: expected map[string]interface{}")
														}
													}
													if _, ok := rInspectConfigRuleSetRulesExclusionRule["excludeInfoTypes"]; ok {
														if rInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, ok := rInspectConfigRuleSetRulesExclusionRule["excludeInfoTypes"].(map[string]interface{}); ok {
															rInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes = &dclService.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
															if _, ok := rInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes["infoTypes"]; ok {
																if s, ok := rInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes["infoTypes"].([]interface{}); ok {
																	for _, o := range s {
																		if objval, ok := o.(map[string]interface{}); ok {
																			var rInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes dclService.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
																			if _, ok := objval["name"]; ok {
																				if s, ok := objval["name"].(string); ok {
																					rInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.Name = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes.Name: expected string")
																				}
																			}
																			rInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes.InfoTypes = append(rInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes.InfoTypes, rInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes.InfoTypes: expected []interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.ExcludeInfoTypes: expected map[string]interface{}")
														}
													}
													if _, ok := rInspectConfigRuleSetRulesExclusionRule["matchingType"]; ok {
														if s, ok := rInspectConfigRuleSetRulesExclusionRule["matchingType"].(string); ok {
															rInspectConfigRuleSetRules.ExclusionRule.MatchingType = dclService.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumRef(s)
														} else {
															return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.MatchingType: expected string")
														}
													}
													if _, ok := rInspectConfigRuleSetRulesExclusionRule["regex"]; ok {
														if rInspectConfigRuleSetRulesExclusionRuleRegex, ok := rInspectConfigRuleSetRulesExclusionRule["regex"].(map[string]interface{}); ok {
															rInspectConfigRuleSetRules.ExclusionRule.Regex = &dclService.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
															if _, ok := rInspectConfigRuleSetRulesExclusionRuleRegex["groupIndexes"]; ok {
																if s, ok := rInspectConfigRuleSetRulesExclusionRuleRegex["groupIndexes"].([]interface{}); ok {
																	for _, ss := range s {
																		if intval, ok := ss.(int64); ok {
																			rInspectConfigRuleSetRules.ExclusionRule.Regex.GroupIndexes = append(rInspectConfigRuleSetRules.ExclusionRule.Regex.GroupIndexes, intval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.Regex.GroupIndexes: expected []interface{}")
																}
															}
															if _, ok := rInspectConfigRuleSetRulesExclusionRuleRegex["pattern"]; ok {
																if s, ok := rInspectConfigRuleSetRulesExclusionRuleRegex["pattern"].(string); ok {
																	rInspectConfigRuleSetRules.ExclusionRule.Regex.Pattern = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.Regex.Pattern: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule.Regex: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rInspectConfigRuleSetRules.ExclusionRule: expected map[string]interface{}")
												}
											}
											if _, ok := objval["hotwordRule"]; ok {
												if rInspectConfigRuleSetRulesHotwordRule, ok := objval["hotwordRule"].(map[string]interface{}); ok {
													rInspectConfigRuleSetRules.HotwordRule = &dclService.InspectTemplateInspectConfigRuleSetRulesHotwordRule{}
													if _, ok := rInspectConfigRuleSetRulesHotwordRule["hotwordRegex"]; ok {
														if rInspectConfigRuleSetRulesHotwordRuleHotwordRegex, ok := rInspectConfigRuleSetRulesHotwordRule["hotwordRegex"].(map[string]interface{}); ok {
															rInspectConfigRuleSetRules.HotwordRule.HotwordRegex = &dclService.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
															if _, ok := rInspectConfigRuleSetRulesHotwordRuleHotwordRegex["groupIndexes"]; ok {
																if s, ok := rInspectConfigRuleSetRulesHotwordRuleHotwordRegex["groupIndexes"].([]interface{}); ok {
																	for _, ss := range s {
																		if intval, ok := ss.(int64); ok {
																			rInspectConfigRuleSetRules.HotwordRule.HotwordRegex.GroupIndexes = append(rInspectConfigRuleSetRules.HotwordRule.HotwordRegex.GroupIndexes, intval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.HotwordRegex.GroupIndexes: expected []interface{}")
																}
															}
															if _, ok := rInspectConfigRuleSetRulesHotwordRuleHotwordRegex["pattern"]; ok {
																if s, ok := rInspectConfigRuleSetRulesHotwordRuleHotwordRegex["pattern"].(string); ok {
																	rInspectConfigRuleSetRules.HotwordRule.HotwordRegex.Pattern = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.HotwordRegex.Pattern: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.HotwordRegex: expected map[string]interface{}")
														}
													}
													if _, ok := rInspectConfigRuleSetRulesHotwordRule["likelihoodAdjustment"]; ok {
														if rInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, ok := rInspectConfigRuleSetRulesHotwordRule["likelihoodAdjustment"].(map[string]interface{}); ok {
															rInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment = &dclService.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
															if _, ok := rInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment["fixedLikelihood"]; ok {
																if s, ok := rInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment["fixedLikelihood"].(string); ok {
																	rInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment.FixedLikelihood = dclService.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment.FixedLikelihood: expected string")
																}
															}
															if _, ok := rInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment["relativeLikelihood"]; ok {
																if i, ok := rInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment["relativeLikelihood"].(int64); ok {
																	rInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment.RelativeLikelihood = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment.RelativeLikelihood: expected int64")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.LikelihoodAdjustment: expected map[string]interface{}")
														}
													}
													if _, ok := rInspectConfigRuleSetRulesHotwordRule["proximity"]; ok {
														if rInspectConfigRuleSetRulesHotwordRuleProximity, ok := rInspectConfigRuleSetRulesHotwordRule["proximity"].(map[string]interface{}); ok {
															rInspectConfigRuleSetRules.HotwordRule.Proximity = &dclService.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
															if _, ok := rInspectConfigRuleSetRulesHotwordRuleProximity["windowAfter"]; ok {
																if i, ok := rInspectConfigRuleSetRulesHotwordRuleProximity["windowAfter"].(int64); ok {
																	rInspectConfigRuleSetRules.HotwordRule.Proximity.WindowAfter = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.Proximity.WindowAfter: expected int64")
																}
															}
															if _, ok := rInspectConfigRuleSetRulesHotwordRuleProximity["windowBefore"]; ok {
																if i, ok := rInspectConfigRuleSetRulesHotwordRuleProximity["windowBefore"].(int64); ok {
																	rInspectConfigRuleSetRules.HotwordRule.Proximity.WindowBefore = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.Proximity.WindowBefore: expected int64")
																}
															}
														} else {
															return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule.Proximity: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rInspectConfigRuleSetRules.HotwordRule: expected map[string]interface{}")
												}
											}
											rInspectConfigRuleSet.Rules = append(rInspectConfigRuleSet.Rules, rInspectConfigRuleSetRules)
										}
									}
								} else {
									return nil, fmt.Errorf("rInspectConfigRuleSet.Rules: expected []interface{}")
								}
							}
							r.InspectConfig.RuleSet = append(r.InspectConfig.RuleSet, rInspectConfigRuleSet)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InspectConfig.RuleSet: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.InspectConfig: expected map[string]interface{}")
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
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetInspectTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInspectTemplate(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetInspectTemplate(ctx, r)
	if err != nil {
		return nil, err
	}
	return InspectTemplateToUnstructured(r), nil
}

func ListInspectTemplate(ctx context.Context, config *dcl.Config, location string, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListInspectTemplate(ctx, location, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, InspectTemplateToUnstructured(r))
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

func ApplyInspectTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInspectTemplate(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInspectTemplate(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyInspectTemplate(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return InspectTemplateToUnstructured(r), nil
}

func InspectTemplateHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInspectTemplate(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInspectTemplate(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyInspectTemplate(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteInspectTemplate(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInspectTemplate(u)
	if err != nil {
		return err
	}
	return c.DeleteInspectTemplate(ctx, r)
}

func InspectTemplateID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToInspectTemplate(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *InspectTemplate) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dlp",
		"InspectTemplate",
		"beta",
	}
}

func (r *InspectTemplate) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InspectTemplate) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InspectTemplate) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *InspectTemplate) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InspectTemplate) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InspectTemplate) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InspectTemplate) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetInspectTemplate(ctx, config, resource)
}

func (r *InspectTemplate) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyInspectTemplate(ctx, config, resource, opts...)
}

func (r *InspectTemplate) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return InspectTemplateHasDiff(ctx, config, resource, opts...)
}

func (r *InspectTemplate) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteInspectTemplate(ctx, config, resource)
}

func (r *InspectTemplate) ID(resource *unstructured.Resource) (string, error) {
	return InspectTemplateID(resource)
}

func init() {
	unstructured.Register(&InspectTemplate{})
}
