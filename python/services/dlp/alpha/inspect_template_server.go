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

// InspectTemplateServer implements the gRPC interface for InspectTemplate.
type InspectTemplateServer struct{}

// ProtoToInspectTemplateInspectConfigMinLikelihoodEnum converts a InspectTemplateInspectConfigMinLikelihoodEnum enum from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum(e alphapb.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum) *alpha.InspectTemplateInspectConfigMinLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum_name[int32(e)]; ok {
		e := alpha.InspectTemplateInspectConfigMinLikelihoodEnum(n[len("DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum converts a InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum enum from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(e alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum) *alpha.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum_name[int32(e)]; ok {
		e := alpha.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(n[len("DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum converts a InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum enum from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(e alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum) *alpha.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum_name[int32(e)]; ok {
		e := alpha.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(n[len("DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigContentOptionsEnum converts a InspectTemplateInspectConfigContentOptionsEnum enum from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigContentOptionsEnum(e alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum) *alpha.InspectTemplateInspectConfigContentOptionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum_name[int32(e)]; ok {
		e := alpha.InspectTemplateInspectConfigContentOptionsEnum(n[len("DlpAlphaInspectTemplateInspectConfigContentOptionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(e alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_name[int32(e)]; ok {
		e := alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(n[len("DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfig converts a InspectTemplateInspectConfig object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfig(p *alphapb.DlpAlphaInspectTemplateInspectConfig) *alpha.InspectTemplateInspectConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfig{
		MinLikelihood:    ProtoToDlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum(p.GetMinLikelihood()),
		Limits:           ProtoToDlpAlphaInspectTemplateInspectConfigLimits(p.GetLimits()),
		IncludeQuote:     dcl.Bool(p.GetIncludeQuote()),
		ExcludeInfoTypes: dcl.Bool(p.GetExcludeInfoTypes()),
	}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpAlphaInspectTemplateInspectConfigInfoTypes(r))
	}
	for _, r := range p.GetCustomInfoTypes() {
		obj.CustomInfoTypes = append(obj.CustomInfoTypes, *ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypes(r))
	}
	for _, r := range p.GetContentOptions() {
		obj.ContentOptions = append(obj.ContentOptions, *ProtoToDlpAlphaInspectTemplateInspectConfigContentOptionsEnum(r))
	}
	for _, r := range p.GetRuleSet() {
		obj.RuleSet = append(obj.RuleSet, *ProtoToDlpAlphaInspectTemplateInspectConfigRuleSet(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigInfoTypes converts a InspectTemplateInspectConfigInfoTypes object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigInfoTypes(p *alphapb.DlpAlphaInspectTemplateInspectConfigInfoTypes) *alpha.InspectTemplateInspectConfigInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimits converts a InspectTemplateInspectConfigLimits object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigLimits(p *alphapb.DlpAlphaInspectTemplateInspectConfigLimits) *alpha.InspectTemplateInspectConfigLimits {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigLimits{
		MaxFindingsPerItem:    dcl.Int64OrNil(p.GetMaxFindingsPerItem()),
		MaxFindingsPerRequest: dcl.Int64OrNil(p.GetMaxFindingsPerRequest()),
	}
	for _, r := range p.GetMaxFindingsPerInfoType() {
		obj.MaxFindingsPerInfoType = append(obj.MaxFindingsPerInfoType, *ProtoToDlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(p *alphapb.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) *alpha.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{
		InfoType:    ProtoToDlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p.GetInfoType()),
		MaxFindings: dcl.Int64OrNil(p.GetMaxFindings()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p *alphapb.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *alpha.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypes converts a InspectTemplateInspectConfigCustomInfoTypes object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypes(p *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypes) *alpha.InspectTemplateInspectConfigCustomInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigCustomInfoTypes{
		InfoType:      ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesInfoType(p.GetInfoType()),
		Likelihood:    ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(p.GetLikelihood()),
		Dictionary:    ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionary(p.GetDictionary()),
		Regex:         ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesRegex(p.GetRegex()),
		SurrogateType: ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesSurrogateType(p.GetSurrogateType()),
		StoredType:    ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesStoredType(p.GetStoredType()),
		ExclusionType: ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(p.GetExclusionType()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesInfoType converts a InspectTemplateInspectConfigCustomInfoTypesInfoType object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesInfoType(p *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesInfoType) *alpha.InspectTemplateInspectConfigCustomInfoTypesInfoType {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigCustomInfoTypesInfoType{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionary converts a InspectTemplateInspectConfigCustomInfoTypesDictionary object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionary(p *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionary) *alpha.InspectTemplateInspectConfigCustomInfoTypesDictionary {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigCustomInfoTypesDictionary{
		WordList:         ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(p *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) *alpha.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *alpha.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesRegex converts a InspectTemplateInspectConfigCustomInfoTypesRegex object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesRegex(p *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesRegex) *alpha.InspectTemplateInspectConfigCustomInfoTypesRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigCustomInfoTypesRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesSurrogateType converts a InspectTemplateInspectConfigCustomInfoTypesSurrogateType object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesSurrogateType(p *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesSurrogateType) *alpha.InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesStoredType converts a InspectTemplateInspectConfigCustomInfoTypesStoredType object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigCustomInfoTypesStoredType(p *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesStoredType) *alpha.InspectTemplateInspectConfigCustomInfoTypesStoredType {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigCustomInfoTypesStoredType{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSet converts a InspectTemplateInspectConfigRuleSet object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSet(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSet) *alpha.InspectTemplateInspectConfigRuleSet {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSet{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetInfoTypes(r))
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRules(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetInfoTypes converts a InspectTemplateInspectConfigRuleSetInfoTypes object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetInfoTypes(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetInfoTypes) *alpha.InspectTemplateInspectConfigRuleSetInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRules converts a InspectTemplateInspectConfigRuleSetRules object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRules(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRules) *alpha.InspectTemplateInspectConfigRuleSetRules {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRules{
		HotwordRule:   ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRule(p.GetHotwordRule()),
		ExclusionRule: ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRule(p.GetExclusionRule()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRule converts a InspectTemplateInspectConfigRuleSetRulesHotwordRule object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRule(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRule) *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRule{
		HotwordRegex:         ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRule converts a InspectTemplateInspectConfigRuleSetRulesExclusionRule object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRule(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRule) *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRule{
		Dictionary:       ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(p.GetDictionary()),
		Regex:            ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(p.GetRegex()),
		ExcludeInfoTypes: ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p.GetExcludeInfoTypes()),
		MatchingType:     ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(p.GetMatchingType()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{
		WordList:         ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object from its proto representation.
func ProtoToDlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(p *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplate converts a InspectTemplate resource from its proto representation.
func ProtoToInspectTemplate(p *alphapb.DlpAlphaInspectTemplate) *alpha.InspectTemplate {
	obj := &alpha.InspectTemplate{
		Name:          dcl.StringOrNil(p.GetName()),
		DisplayName:   dcl.StringOrNil(p.GetDisplayName()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		InspectConfig: ProtoToDlpAlphaInspectTemplateInspectConfig(p.GetInspectConfig()),
		LocationId:    dcl.StringOrNil(p.GetLocationId()),
		Parent:        dcl.StringOrNil(p.GetParent()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// InspectTemplateInspectConfigMinLikelihoodEnumToProto converts a InspectTemplateInspectConfigMinLikelihoodEnum enum to its proto representation.
func DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnumToProto(e *alpha.InspectTemplateInspectConfigMinLikelihoodEnum) alphapb.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum {
	if e == nil {
		return alphapb.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum_value["InspectTemplateInspectConfigMinLikelihoodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum(v)
	}
	return alphapb.DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnum(0)
}

// InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto converts a InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum enum to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto(e *alpha.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum) alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == nil {
		return alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum_value["InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(v)
	}
	return alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(0)
}

// InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto converts a InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum enum to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto(e *alpha.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum) alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == nil {
		return alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(0)
	}
	if v, ok := alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum_value["InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(v)
	}
	return alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(0)
}

// InspectTemplateInspectConfigContentOptionsEnumToProto converts a InspectTemplateInspectConfigContentOptionsEnum enum to its proto representation.
func DlpAlphaInspectTemplateInspectConfigContentOptionsEnumToProto(e *alpha.InspectTemplateInspectConfigContentOptionsEnum) alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum {
	if e == nil {
		return alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum(0)
	}
	if v, ok := alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum_value["InspectTemplateInspectConfigContentOptionsEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum(v)
	}
	return alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum(0)
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(e *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == nil {
		return alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
	}
	if v, ok := alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_value["InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"+string(*e)]; ok {
		return alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(v)
	}
	return alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
}

// InspectTemplateInspectConfigToProto converts a InspectTemplateInspectConfig object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigToProto(o *alpha.InspectTemplateInspectConfig) *alphapb.DlpAlphaInspectTemplateInspectConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfig{}
	p.SetMinLikelihood(DlpAlphaInspectTemplateInspectConfigMinLikelihoodEnumToProto(o.MinLikelihood))
	p.SetLimits(DlpAlphaInspectTemplateInspectConfigLimitsToProto(o.Limits))
	p.SetIncludeQuote(dcl.ValueOrEmptyBool(o.IncludeQuote))
	p.SetExcludeInfoTypes(dcl.ValueOrEmptyBool(o.ExcludeInfoTypes))
	sInfoTypes := make([]*alphapb.DlpAlphaInspectTemplateInspectConfigInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpAlphaInspectTemplateInspectConfigInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sCustomInfoTypes := make([]*alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypes, len(o.CustomInfoTypes))
	for i, r := range o.CustomInfoTypes {
		sCustomInfoTypes[i] = DlpAlphaInspectTemplateInspectConfigCustomInfoTypesToProto(&r)
	}
	p.SetCustomInfoTypes(sCustomInfoTypes)
	sContentOptions := make([]alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum, len(o.ContentOptions))
	for i, r := range o.ContentOptions {
		sContentOptions[i] = alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum(alphapb.DlpAlphaInspectTemplateInspectConfigContentOptionsEnum_value[string(r)])
	}
	p.SetContentOptions(sContentOptions)
	sRuleSet := make([]*alphapb.DlpAlphaInspectTemplateInspectConfigRuleSet, len(o.RuleSet))
	for i, r := range o.RuleSet {
		sRuleSet[i] = DlpAlphaInspectTemplateInspectConfigRuleSetToProto(&r)
	}
	p.SetRuleSet(sRuleSet)
	return p
}

// InspectTemplateInspectConfigInfoTypesToProto converts a InspectTemplateInspectConfigInfoTypes object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigInfoTypesToProto(o *alpha.InspectTemplateInspectConfigInfoTypes) *alphapb.DlpAlphaInspectTemplateInspectConfigInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigLimitsToProto converts a InspectTemplateInspectConfigLimits object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigLimitsToProto(o *alpha.InspectTemplateInspectConfigLimits) *alphapb.DlpAlphaInspectTemplateInspectConfigLimits {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigLimits{}
	p.SetMaxFindingsPerItem(dcl.ValueOrEmptyInt64(o.MaxFindingsPerItem))
	p.SetMaxFindingsPerRequest(dcl.ValueOrEmptyInt64(o.MaxFindingsPerRequest))
	sMaxFindingsPerInfoType := make([]*alphapb.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, len(o.MaxFindingsPerInfoType))
	for i, r := range o.MaxFindingsPerInfoType {
		sMaxFindingsPerInfoType[i] = DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto(&r)
	}
	p.SetMaxFindingsPerInfoType(sMaxFindingsPerInfoType)
	return p
}

// InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto(o *alpha.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) *alphapb.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}
	p.SetInfoType(DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o.InfoType))
	p.SetMaxFindings(dcl.ValueOrEmptyInt64(o.MaxFindings))
	return p
}

// InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o *alpha.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *alphapb.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesToProto converts a InspectTemplateInspectConfigCustomInfoTypes object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesToProto(o *alpha.InspectTemplateInspectConfigCustomInfoTypes) *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypes{}
	p.SetInfoType(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto(o.InfoType))
	p.SetLikelihood(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto(o.Likelihood))
	p.SetDictionary(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesRegexToProto(o.Regex))
	p.SetSurrogateType(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto(o.SurrogateType))
	p.SetStoredType(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto(o.StoredType))
	p.SetExclusionType(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto(o.ExclusionType))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesInfoType object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto(o *alpha.InspectTemplateInspectConfigCustomInfoTypesInfoType) *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesInfoType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionary object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryToProto(o *alpha.InspectTemplateInspectConfigCustomInfoTypesDictionary) *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionary {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionary{}
	p.SetWordList(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto(o *alpha.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o *alpha.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesRegexToProto converts a InspectTemplateInspectConfigCustomInfoTypesRegex object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesRegexToProto(o *alpha.InspectTemplateInspectConfigCustomInfoTypesRegex) *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesSurrogateType object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto(o *alpha.InspectTemplateInspectConfigCustomInfoTypesSurrogateType) *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesStoredType object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto(o *alpha.InspectTemplateInspectConfigCustomInfoTypesStoredType) *alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesStoredType {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigCustomInfoTypesStoredType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// InspectTemplateInspectConfigRuleSetToProto converts a InspectTemplateInspectConfigRuleSet object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetToProto(o *alpha.InspectTemplateInspectConfigRuleSet) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSet {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSet{}
	sInfoTypes := make([]*alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpAlphaInspectTemplateInspectConfigRuleSetInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sRules := make([]*alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = DlpAlphaInspectTemplateInspectConfigRuleSetRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// InspectTemplateInspectConfigRuleSetInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetInfoTypes object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetInfoTypesToProto(o *alpha.InspectTemplateInspectConfigRuleSetInfoTypes) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesToProto converts a InspectTemplateInspectConfigRuleSetRules object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesToProto(o *alpha.InspectTemplateInspectConfigRuleSetRules) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRules {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRules{}
	p.SetHotwordRule(DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto(o.HotwordRule))
	p.SetExclusionRule(DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto(o.ExclusionRule))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRule object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRule) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	p.SetHotwordRegex(DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpAlphaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRule object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRule) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	p.SetDictionary(DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto(o.Regex))
	p.SetExcludeInfoTypes(DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o.ExcludeInfoTypes))
	p.SetMatchingType(DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(o.MatchingType))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	p.SetWordList(DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	sInfoTypes := make([]*alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object to its proto representation.
func DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(o *alpha.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.DlpAlphaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateToProto converts a InspectTemplate resource to its proto representation.
func InspectTemplateToProto(resource *alpha.InspectTemplate) *alphapb.DlpAlphaInspectTemplate {
	p := &alphapb.DlpAlphaInspectTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetInspectConfig(DlpAlphaInspectTemplateInspectConfigToProto(resource.InspectConfig))
	p.SetLocationId(dcl.ValueOrEmptyString(resource.LocationId))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Apply() method.
func (s *InspectTemplateServer) applyInspectTemplate(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDlpAlphaInspectTemplateRequest) (*alphapb.DlpAlphaInspectTemplate, error) {
	p := ProtoToInspectTemplate(request.GetResource())
	res, err := c.ApplyInspectTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InspectTemplateToProto(res)
	return r, nil
}

// applyDlpAlphaInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Apply() method.
func (s *InspectTemplateServer) ApplyDlpAlphaInspectTemplate(ctx context.Context, request *alphapb.ApplyDlpAlphaInspectTemplateRequest) (*alphapb.DlpAlphaInspectTemplate, error) {
	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInspectTemplate(ctx, cl, request)
}

// DeleteInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Delete() method.
func (s *InspectTemplateServer) DeleteDlpAlphaInspectTemplate(ctx context.Context, request *alphapb.DeleteDlpAlphaInspectTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInspectTemplate(ctx, ProtoToInspectTemplate(request.GetResource()))

}

// ListDlpAlphaInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplateList() method.
func (s *InspectTemplateServer) ListDlpAlphaInspectTemplate(ctx context.Context, request *alphapb.ListDlpAlphaInspectTemplateRequest) (*alphapb.ListDlpAlphaInspectTemplateResponse, error) {
	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInspectTemplate(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DlpAlphaInspectTemplate
	for _, r := range resources.Items {
		rp := InspectTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDlpAlphaInspectTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInspectTemplate(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
