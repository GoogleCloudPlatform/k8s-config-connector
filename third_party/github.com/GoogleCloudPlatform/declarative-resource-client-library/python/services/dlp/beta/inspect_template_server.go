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

// InspectTemplateServer implements the gRPC interface for InspectTemplate.
type InspectTemplateServer struct{}

// ProtoToInspectTemplateInspectConfigMinLikelihoodEnum converts a InspectTemplateInspectConfigMinLikelihoodEnum enum from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigMinLikelihoodEnum(e betapb.DlpBetaInspectTemplateInspectConfigMinLikelihoodEnum) *beta.InspectTemplateInspectConfigMinLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaInspectTemplateInspectConfigMinLikelihoodEnum_name[int32(e)]; ok {
		e := beta.InspectTemplateInspectConfigMinLikelihoodEnum(n[len("DlpBetaInspectTemplateInspectConfigMinLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum converts a InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum enum from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(e betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum) *beta.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum_name[int32(e)]; ok {
		e := beta.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(n[len("DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum converts a InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum enum from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(e betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum) *beta.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum_name[int32(e)]; ok {
		e := beta.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(n[len("DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigContentOptionsEnum converts a InspectTemplateInspectConfigContentOptionsEnum enum from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigContentOptionsEnum(e betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum) *beta.InspectTemplateInspectConfigContentOptionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum_name[int32(e)]; ok {
		e := beta.InspectTemplateInspectConfigContentOptionsEnum(n[len("DlpBetaInspectTemplateInspectConfigContentOptionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(e betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_name[int32(e)]; ok {
		e := beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(n[len("DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfig converts a InspectTemplateInspectConfig object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfig(p *betapb.DlpBetaInspectTemplateInspectConfig) *beta.InspectTemplateInspectConfig {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfig{
		MinLikelihood:    ProtoToDlpBetaInspectTemplateInspectConfigMinLikelihoodEnum(p.GetMinLikelihood()),
		Limits:           ProtoToDlpBetaInspectTemplateInspectConfigLimits(p.GetLimits()),
		IncludeQuote:     dcl.Bool(p.GetIncludeQuote()),
		ExcludeInfoTypes: dcl.Bool(p.GetExcludeInfoTypes()),
	}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpBetaInspectTemplateInspectConfigInfoTypes(r))
	}
	for _, r := range p.GetCustomInfoTypes() {
		obj.CustomInfoTypes = append(obj.CustomInfoTypes, *ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypes(r))
	}
	for _, r := range p.GetContentOptions() {
		obj.ContentOptions = append(obj.ContentOptions, *ProtoToDlpBetaInspectTemplateInspectConfigContentOptionsEnum(r))
	}
	for _, r := range p.GetRuleSet() {
		obj.RuleSet = append(obj.RuleSet, *ProtoToDlpBetaInspectTemplateInspectConfigRuleSet(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigInfoTypes converts a InspectTemplateInspectConfigInfoTypes object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigInfoTypes(p *betapb.DlpBetaInspectTemplateInspectConfigInfoTypes) *beta.InspectTemplateInspectConfigInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimits converts a InspectTemplateInspectConfigLimits object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigLimits(p *betapb.DlpBetaInspectTemplateInspectConfigLimits) *beta.InspectTemplateInspectConfigLimits {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigLimits{
		MaxFindingsPerItem:    dcl.Int64OrNil(p.GetMaxFindingsPerItem()),
		MaxFindingsPerRequest: dcl.Int64OrNil(p.GetMaxFindingsPerRequest()),
	}
	for _, r := range p.GetMaxFindingsPerInfoType() {
		obj.MaxFindingsPerInfoType = append(obj.MaxFindingsPerInfoType, *ProtoToDlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(p *betapb.DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) *beta.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{
		InfoType:    ProtoToDlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p.GetInfoType()),
		MaxFindings: dcl.Int64OrNil(p.GetMaxFindings()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p *betapb.DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *beta.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypes converts a InspectTemplateInspectConfigCustomInfoTypes object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypes(p *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypes) *beta.InspectTemplateInspectConfigCustomInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigCustomInfoTypes{
		InfoType:      ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesInfoType(p.GetInfoType()),
		Likelihood:    ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(p.GetLikelihood()),
		Dictionary:    ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionary(p.GetDictionary()),
		Regex:         ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesRegex(p.GetRegex()),
		SurrogateType: ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesSurrogateType(p.GetSurrogateType()),
		StoredType:    ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesStoredType(p.GetStoredType()),
		ExclusionType: ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(p.GetExclusionType()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesInfoType converts a InspectTemplateInspectConfigCustomInfoTypesInfoType object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesInfoType(p *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesInfoType) *beta.InspectTemplateInspectConfigCustomInfoTypesInfoType {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigCustomInfoTypesInfoType{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionary converts a InspectTemplateInspectConfigCustomInfoTypesDictionary object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionary(p *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionary) *beta.InspectTemplateInspectConfigCustomInfoTypesDictionary {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigCustomInfoTypesDictionary{
		WordList:         ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(p *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) *beta.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *beta.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesRegex converts a InspectTemplateInspectConfigCustomInfoTypesRegex object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesRegex(p *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesRegex) *beta.InspectTemplateInspectConfigCustomInfoTypesRegex {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigCustomInfoTypesRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesSurrogateType converts a InspectTemplateInspectConfigCustomInfoTypesSurrogateType object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesSurrogateType(p *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesSurrogateType) *beta.InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesStoredType converts a InspectTemplateInspectConfigCustomInfoTypesStoredType object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigCustomInfoTypesStoredType(p *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesStoredType) *beta.InspectTemplateInspectConfigCustomInfoTypesStoredType {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigCustomInfoTypesStoredType{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSet converts a InspectTemplateInspectConfigRuleSet object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSet(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSet) *beta.InspectTemplateInspectConfigRuleSet {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSet{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpBetaInspectTemplateInspectConfigRuleSetInfoTypes(r))
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRules(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetInfoTypes converts a InspectTemplateInspectConfigRuleSetInfoTypes object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetInfoTypes(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetInfoTypes) *beta.InspectTemplateInspectConfigRuleSetInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRules converts a InspectTemplateInspectConfigRuleSetRules object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRules(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRules) *beta.InspectTemplateInspectConfigRuleSetRules {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRules{
		HotwordRule:   ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRule(p.GetHotwordRule()),
		ExclusionRule: ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRule(p.GetExclusionRule()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRule converts a InspectTemplateInspectConfigRuleSetRulesHotwordRule object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRule(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRule) *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesHotwordRule{
		HotwordRegex:         ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRule converts a InspectTemplateInspectConfigRuleSetRulesExclusionRule object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRule(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRule) *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesExclusionRule{
		Dictionary:       ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(p.GetDictionary()),
		Regex:            ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(p.GetRegex()),
		ExcludeInfoTypes: ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p.GetExcludeInfoTypes()),
		MatchingType:     ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(p.GetMatchingType()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{
		WordList:         ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object from its proto representation.
func ProtoToDlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(p *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if p == nil {
		return nil
	}
	obj := &beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplate converts a InspectTemplate resource from its proto representation.
func ProtoToInspectTemplate(p *betapb.DlpBetaInspectTemplate) *beta.InspectTemplate {
	obj := &beta.InspectTemplate{
		Name:          dcl.StringOrNil(p.GetName()),
		DisplayName:   dcl.StringOrNil(p.GetDisplayName()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		InspectConfig: ProtoToDlpBetaInspectTemplateInspectConfig(p.GetInspectConfig()),
		LocationId:    dcl.StringOrNil(p.GetLocationId()),
		Parent:        dcl.StringOrNil(p.GetParent()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// InspectTemplateInspectConfigMinLikelihoodEnumToProto converts a InspectTemplateInspectConfigMinLikelihoodEnum enum to its proto representation.
func DlpBetaInspectTemplateInspectConfigMinLikelihoodEnumToProto(e *beta.InspectTemplateInspectConfigMinLikelihoodEnum) betapb.DlpBetaInspectTemplateInspectConfigMinLikelihoodEnum {
	if e == nil {
		return betapb.DlpBetaInspectTemplateInspectConfigMinLikelihoodEnum(0)
	}
	if v, ok := betapb.DlpBetaInspectTemplateInspectConfigMinLikelihoodEnum_value["InspectTemplateInspectConfigMinLikelihoodEnum"+string(*e)]; ok {
		return betapb.DlpBetaInspectTemplateInspectConfigMinLikelihoodEnum(v)
	}
	return betapb.DlpBetaInspectTemplateInspectConfigMinLikelihoodEnum(0)
}

// InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto converts a InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum enum to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto(e *beta.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum) betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == nil {
		return betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(0)
	}
	if v, ok := betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum_value["InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum"+string(*e)]; ok {
		return betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(v)
	}
	return betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(0)
}

// InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto converts a InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum enum to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto(e *beta.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum) betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == nil {
		return betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(0)
	}
	if v, ok := betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum_value["InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum"+string(*e)]; ok {
		return betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(v)
	}
	return betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(0)
}

// InspectTemplateInspectConfigContentOptionsEnumToProto converts a InspectTemplateInspectConfigContentOptionsEnum enum to its proto representation.
func DlpBetaInspectTemplateInspectConfigContentOptionsEnumToProto(e *beta.InspectTemplateInspectConfigContentOptionsEnum) betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum {
	if e == nil {
		return betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum(0)
	}
	if v, ok := betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum_value["InspectTemplateInspectConfigContentOptionsEnum"+string(*e)]; ok {
		return betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum(v)
	}
	return betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum(0)
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(e *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == nil {
		return betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
	}
	if v, ok := betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_value["InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"+string(*e)]; ok {
		return betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(v)
	}
	return betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
}

// InspectTemplateInspectConfigToProto converts a InspectTemplateInspectConfig object to its proto representation.
func DlpBetaInspectTemplateInspectConfigToProto(o *beta.InspectTemplateInspectConfig) *betapb.DlpBetaInspectTemplateInspectConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfig{}
	p.SetMinLikelihood(DlpBetaInspectTemplateInspectConfigMinLikelihoodEnumToProto(o.MinLikelihood))
	p.SetLimits(DlpBetaInspectTemplateInspectConfigLimitsToProto(o.Limits))
	p.SetIncludeQuote(dcl.ValueOrEmptyBool(o.IncludeQuote))
	p.SetExcludeInfoTypes(dcl.ValueOrEmptyBool(o.ExcludeInfoTypes))
	sInfoTypes := make([]*betapb.DlpBetaInspectTemplateInspectConfigInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpBetaInspectTemplateInspectConfigInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sCustomInfoTypes := make([]*betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypes, len(o.CustomInfoTypes))
	for i, r := range o.CustomInfoTypes {
		sCustomInfoTypes[i] = DlpBetaInspectTemplateInspectConfigCustomInfoTypesToProto(&r)
	}
	p.SetCustomInfoTypes(sCustomInfoTypes)
	sContentOptions := make([]betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum, len(o.ContentOptions))
	for i, r := range o.ContentOptions {
		sContentOptions[i] = betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum(betapb.DlpBetaInspectTemplateInspectConfigContentOptionsEnum_value[string(r)])
	}
	p.SetContentOptions(sContentOptions)
	sRuleSet := make([]*betapb.DlpBetaInspectTemplateInspectConfigRuleSet, len(o.RuleSet))
	for i, r := range o.RuleSet {
		sRuleSet[i] = DlpBetaInspectTemplateInspectConfigRuleSetToProto(&r)
	}
	p.SetRuleSet(sRuleSet)
	return p
}

// InspectTemplateInspectConfigInfoTypesToProto converts a InspectTemplateInspectConfigInfoTypes object to its proto representation.
func DlpBetaInspectTemplateInspectConfigInfoTypesToProto(o *beta.InspectTemplateInspectConfigInfoTypes) *betapb.DlpBetaInspectTemplateInspectConfigInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigLimitsToProto converts a InspectTemplateInspectConfigLimits object to its proto representation.
func DlpBetaInspectTemplateInspectConfigLimitsToProto(o *beta.InspectTemplateInspectConfigLimits) *betapb.DlpBetaInspectTemplateInspectConfigLimits {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigLimits{}
	p.SetMaxFindingsPerItem(dcl.ValueOrEmptyInt64(o.MaxFindingsPerItem))
	p.SetMaxFindingsPerRequest(dcl.ValueOrEmptyInt64(o.MaxFindingsPerRequest))
	sMaxFindingsPerInfoType := make([]*betapb.DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, len(o.MaxFindingsPerInfoType))
	for i, r := range o.MaxFindingsPerInfoType {
		sMaxFindingsPerInfoType[i] = DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto(&r)
	}
	p.SetMaxFindingsPerInfoType(sMaxFindingsPerInfoType)
	return p
}

// InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType object to its proto representation.
func DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto(o *beta.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) *betapb.DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}
	p.SetInfoType(DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o.InfoType))
	p.SetMaxFindings(dcl.ValueOrEmptyInt64(o.MaxFindings))
	return p
}

// InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object to its proto representation.
func DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o *beta.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *betapb.DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesToProto converts a InspectTemplateInspectConfigCustomInfoTypes object to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesToProto(o *beta.InspectTemplateInspectConfigCustomInfoTypes) *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypes{}
	p.SetInfoType(DlpBetaInspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto(o.InfoType))
	p.SetLikelihood(DlpBetaInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto(o.Likelihood))
	p.SetDictionary(DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpBetaInspectTemplateInspectConfigCustomInfoTypesRegexToProto(o.Regex))
	p.SetSurrogateType(DlpBetaInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto(o.SurrogateType))
	p.SetStoredType(DlpBetaInspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto(o.StoredType))
	p.SetExclusionType(DlpBetaInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto(o.ExclusionType))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesInfoType object to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto(o *beta.InspectTemplateInspectConfigCustomInfoTypesInfoType) *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesInfoType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionary object to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryToProto(o *beta.InspectTemplateInspectConfigCustomInfoTypesDictionary) *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionary {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionary{}
	p.SetWordList(DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList object to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto(o *beta.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath object to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o *beta.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesRegexToProto converts a InspectTemplateInspectConfigCustomInfoTypesRegex object to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesRegexToProto(o *beta.InspectTemplateInspectConfigCustomInfoTypesRegex) *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesSurrogateType object to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto(o *beta.InspectTemplateInspectConfigCustomInfoTypesSurrogateType) *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesStoredType object to its proto representation.
func DlpBetaInspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto(o *beta.InspectTemplateInspectConfigCustomInfoTypesStoredType) *betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesStoredType {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigCustomInfoTypesStoredType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// InspectTemplateInspectConfigRuleSetToProto converts a InspectTemplateInspectConfigRuleSet object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetToProto(o *beta.InspectTemplateInspectConfigRuleSet) *betapb.DlpBetaInspectTemplateInspectConfigRuleSet {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSet{}
	sInfoTypes := make([]*betapb.DlpBetaInspectTemplateInspectConfigRuleSetInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpBetaInspectTemplateInspectConfigRuleSetInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sRules := make([]*betapb.DlpBetaInspectTemplateInspectConfigRuleSetRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = DlpBetaInspectTemplateInspectConfigRuleSetRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// InspectTemplateInspectConfigRuleSetInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetInfoTypes object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetInfoTypesToProto(o *beta.InspectTemplateInspectConfigRuleSetInfoTypes) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesToProto converts a InspectTemplateInspectConfigRuleSetRules object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesToProto(o *beta.InspectTemplateInspectConfigRuleSetRules) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRules {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRules{}
	p.SetHotwordRule(DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto(o.HotwordRule))
	p.SetExclusionRule(DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto(o.ExclusionRule))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRule object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRule) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	p.SetHotwordRegex(DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpBetaInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRule object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRule) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	p.SetDictionary(DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto(o.Regex))
	p.SetExcludeInfoTypes(DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o.ExcludeInfoTypes))
	p.SetMatchingType(DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(o.MatchingType))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	p.SetWordList(DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	sInfoTypes := make([]*betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object to its proto representation.
func DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(o *beta.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if o == nil {
		return nil
	}
	p := &betapb.DlpBetaInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateToProto converts a InspectTemplate resource to its proto representation.
func InspectTemplateToProto(resource *beta.InspectTemplate) *betapb.DlpBetaInspectTemplate {
	p := &betapb.DlpBetaInspectTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetInspectConfig(DlpBetaInspectTemplateInspectConfigToProto(resource.InspectConfig))
	p.SetLocationId(dcl.ValueOrEmptyString(resource.LocationId))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Apply() method.
func (s *InspectTemplateServer) applyInspectTemplate(ctx context.Context, c *beta.Client, request *betapb.ApplyDlpBetaInspectTemplateRequest) (*betapb.DlpBetaInspectTemplate, error) {
	p := ProtoToInspectTemplate(request.GetResource())
	res, err := c.ApplyInspectTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InspectTemplateToProto(res)
	return r, nil
}

// applyDlpBetaInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Apply() method.
func (s *InspectTemplateServer) ApplyDlpBetaInspectTemplate(ctx context.Context, request *betapb.ApplyDlpBetaInspectTemplateRequest) (*betapb.DlpBetaInspectTemplate, error) {
	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInspectTemplate(ctx, cl, request)
}

// DeleteInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Delete() method.
func (s *InspectTemplateServer) DeleteDlpBetaInspectTemplate(ctx context.Context, request *betapb.DeleteDlpBetaInspectTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInspectTemplate(ctx, ProtoToInspectTemplate(request.GetResource()))

}

// ListDlpBetaInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplateList() method.
func (s *InspectTemplateServer) ListDlpBetaInspectTemplate(ctx context.Context, request *betapb.ListDlpBetaInspectTemplateRequest) (*betapb.ListDlpBetaInspectTemplateResponse, error) {
	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInspectTemplate(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DlpBetaInspectTemplate
	for _, r := range resources.Items {
		rp := InspectTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDlpBetaInspectTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInspectTemplate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
