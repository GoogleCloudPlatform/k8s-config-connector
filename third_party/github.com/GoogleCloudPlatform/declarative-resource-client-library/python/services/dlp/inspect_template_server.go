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

// InspectTemplateServer implements the gRPC interface for InspectTemplate.
type InspectTemplateServer struct{}

// ProtoToInspectTemplateInspectConfigMinLikelihoodEnum converts a InspectTemplateInspectConfigMinLikelihoodEnum enum from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigMinLikelihoodEnum(e dlppb.DlpInspectTemplateInspectConfigMinLikelihoodEnum) *dlp.InspectTemplateInspectConfigMinLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpInspectTemplateInspectConfigMinLikelihoodEnum_name[int32(e)]; ok {
		e := dlp.InspectTemplateInspectConfigMinLikelihoodEnum(n[len("DlpInspectTemplateInspectConfigMinLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum converts a InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum enum from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(e dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum) *dlp.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum_name[int32(e)]; ok {
		e := dlp.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(n[len("DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum converts a InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum enum from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(e dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum) *dlp.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum_name[int32(e)]; ok {
		e := dlp.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(n[len("DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigContentOptionsEnum converts a InspectTemplateInspectConfigContentOptionsEnum enum from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigContentOptionsEnum(e dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum) *dlp.InspectTemplateInspectConfigContentOptionsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum_name[int32(e)]; ok {
		e := dlp.InspectTemplateInspectConfigContentOptionsEnum(n[len("DlpInspectTemplateInspectConfigContentOptionsEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(e dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_name[int32(e)]; ok {
		e := dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(n[len("DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(e dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_name[int32(e)]; ok {
		e := dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(n[len("DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInspectTemplateInspectConfig converts a InspectTemplateInspectConfig object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfig(p *dlppb.DlpInspectTemplateInspectConfig) *dlp.InspectTemplateInspectConfig {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfig{
		MinLikelihood:    ProtoToDlpInspectTemplateInspectConfigMinLikelihoodEnum(p.GetMinLikelihood()),
		Limits:           ProtoToDlpInspectTemplateInspectConfigLimits(p.GetLimits()),
		IncludeQuote:     dcl.Bool(p.GetIncludeQuote()),
		ExcludeInfoTypes: dcl.Bool(p.GetExcludeInfoTypes()),
	}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpInspectTemplateInspectConfigInfoTypes(r))
	}
	for _, r := range p.GetCustomInfoTypes() {
		obj.CustomInfoTypes = append(obj.CustomInfoTypes, *ProtoToDlpInspectTemplateInspectConfigCustomInfoTypes(r))
	}
	for _, r := range p.GetContentOptions() {
		obj.ContentOptions = append(obj.ContentOptions, *ProtoToDlpInspectTemplateInspectConfigContentOptionsEnum(r))
	}
	for _, r := range p.GetRuleSet() {
		obj.RuleSet = append(obj.RuleSet, *ProtoToDlpInspectTemplateInspectConfigRuleSet(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigInfoTypes converts a InspectTemplateInspectConfigInfoTypes object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigInfoTypes(p *dlppb.DlpInspectTemplateInspectConfigInfoTypes) *dlp.InspectTemplateInspectConfigInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimits converts a InspectTemplateInspectConfigLimits object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigLimits(p *dlppb.DlpInspectTemplateInspectConfigLimits) *dlp.InspectTemplateInspectConfigLimits {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigLimits{
		MaxFindingsPerItem:    dcl.Int64OrNil(p.GetMaxFindingsPerItem()),
		MaxFindingsPerRequest: dcl.Int64OrNil(p.GetMaxFindingsPerRequest()),
	}
	for _, r := range p.GetMaxFindingsPerInfoType() {
		obj.MaxFindingsPerInfoType = append(obj.MaxFindingsPerInfoType, *ProtoToDlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(p *dlppb.DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) *dlp.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{
		InfoType:    ProtoToDlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p.GetInfoType()),
		MaxFindings: dcl.Int64OrNil(p.GetMaxFindings()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(p *dlppb.DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *dlp.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypes converts a InspectTemplateInspectConfigCustomInfoTypes object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypes(p *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypes) *dlp.InspectTemplateInspectConfigCustomInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigCustomInfoTypes{
		InfoType:      ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesInfoType(p.GetInfoType()),
		Likelihood:    ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(p.GetLikelihood()),
		Dictionary:    ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesDictionary(p.GetDictionary()),
		Regex:         ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesRegex(p.GetRegex()),
		SurrogateType: ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesSurrogateType(p.GetSurrogateType()),
		StoredType:    ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesStoredType(p.GetStoredType()),
		ExclusionType: ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(p.GetExclusionType()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesInfoType converts a InspectTemplateInspectConfigCustomInfoTypesInfoType object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesInfoType(p *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesInfoType) *dlp.InspectTemplateInspectConfigCustomInfoTypesInfoType {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigCustomInfoTypesInfoType{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionary converts a InspectTemplateInspectConfigCustomInfoTypesDictionary object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesDictionary(p *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionary) *dlp.InspectTemplateInspectConfigCustomInfoTypesDictionary {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigCustomInfoTypesDictionary{
		WordList:         ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(p *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) *dlp.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(p *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *dlp.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesRegex converts a InspectTemplateInspectConfigCustomInfoTypesRegex object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesRegex(p *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesRegex) *dlp.InspectTemplateInspectConfigCustomInfoTypesRegex {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigCustomInfoTypesRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesSurrogateType converts a InspectTemplateInspectConfigCustomInfoTypesSurrogateType object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesSurrogateType(p *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesSurrogateType) *dlp.InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	return obj
}

// ProtoToInspectTemplateInspectConfigCustomInfoTypesStoredType converts a InspectTemplateInspectConfigCustomInfoTypesStoredType object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigCustomInfoTypesStoredType(p *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesStoredType) *dlp.InspectTemplateInspectConfigCustomInfoTypesStoredType {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigCustomInfoTypesStoredType{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSet converts a InspectTemplateInspectConfigRuleSet object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSet(p *dlppb.DlpInspectTemplateInspectConfigRuleSet) *dlp.InspectTemplateInspectConfigRuleSet {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSet{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpInspectTemplateInspectConfigRuleSetInfoTypes(r))
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToDlpInspectTemplateInspectConfigRuleSetRules(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetInfoTypes converts a InspectTemplateInspectConfigRuleSetInfoTypes object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetInfoTypes(p *dlppb.DlpInspectTemplateInspectConfigRuleSetInfoTypes) *dlp.InspectTemplateInspectConfigRuleSetInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRules converts a InspectTemplateInspectConfigRuleSetRules object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRules(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRules) *dlp.InspectTemplateInspectConfigRuleSetRules {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRules{
		HotwordRule:   ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRule(p.GetHotwordRule()),
		ExclusionRule: ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRule(p.GetExclusionRule()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRule converts a InspectTemplateInspectConfigRuleSetRulesHotwordRule object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRule(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRule) *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRule{
		HotwordRegex:         ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p.GetHotwordRegex()),
		Proximity:            ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(p.GetProximity()),
		LikelihoodAdjustment: ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p.GetLikelihoodAdjustment()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{
		WindowBefore: dcl.Int64OrNil(p.GetWindowBefore()),
		WindowAfter:  dcl.Int64OrNil(p.GetWindowAfter()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{
		FixedLikelihood:    ProtoToDlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(p.GetFixedLikelihood()),
		RelativeLikelihood: dcl.Int64OrNil(p.GetRelativeLikelihood()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRule converts a InspectTemplateInspectConfigRuleSetRulesExclusionRule object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRule(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRule) *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRule{
		Dictionary:       ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(p.GetDictionary()),
		Regex:            ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(p.GetRegex()),
		ExcludeInfoTypes: ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p.GetExcludeInfoTypes()),
		MatchingType:     ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(p.GetMatchingType()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{
		WordList:         ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p.GetWordList()),
		CloudStoragePath: ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p.GetCloudStoragePath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	for _, r := range p.GetWords() {
		obj.Words = append(obj.Words, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{
		Path: dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{
		Pattern: dcl.StringOrNil(p.GetPattern()),
	}
	for _, r := range p.GetGroupIndexes() {
		obj.GroupIndexes = append(obj.GroupIndexes, r)
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	for _, r := range p.GetInfoTypes() {
		obj.InfoTypes = append(obj.InfoTypes, *ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(r))
	}
	return obj
}

// ProtoToInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object from its proto representation.
func ProtoToDlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(p *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if p == nil {
		return nil
	}
	obj := &dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToInspectTemplate converts a InspectTemplate resource from its proto representation.
func ProtoToInspectTemplate(p *dlppb.DlpInspectTemplate) *dlp.InspectTemplate {
	obj := &dlp.InspectTemplate{
		Name:          dcl.StringOrNil(p.GetName()),
		DisplayName:   dcl.StringOrNil(p.GetDisplayName()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		InspectConfig: ProtoToDlpInspectTemplateInspectConfig(p.GetInspectConfig()),
		LocationId:    dcl.StringOrNil(p.GetLocationId()),
		Parent:        dcl.StringOrNil(p.GetParent()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// InspectTemplateInspectConfigMinLikelihoodEnumToProto converts a InspectTemplateInspectConfigMinLikelihoodEnum enum to its proto representation.
func DlpInspectTemplateInspectConfigMinLikelihoodEnumToProto(e *dlp.InspectTemplateInspectConfigMinLikelihoodEnum) dlppb.DlpInspectTemplateInspectConfigMinLikelihoodEnum {
	if e == nil {
		return dlppb.DlpInspectTemplateInspectConfigMinLikelihoodEnum(0)
	}
	if v, ok := dlppb.DlpInspectTemplateInspectConfigMinLikelihoodEnum_value["InspectTemplateInspectConfigMinLikelihoodEnum"+string(*e)]; ok {
		return dlppb.DlpInspectTemplateInspectConfigMinLikelihoodEnum(v)
	}
	return dlppb.DlpInspectTemplateInspectConfigMinLikelihoodEnum(0)
}

// InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto converts a InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum enum to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto(e *dlp.InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum) dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	if e == nil {
		return dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(0)
	}
	if v, ok := dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum_value["InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum"+string(*e)]; ok {
		return dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(v)
	}
	return dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(0)
}

// InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto converts a InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum enum to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto(e *dlp.InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum) dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	if e == nil {
		return dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(0)
	}
	if v, ok := dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum_value["InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum"+string(*e)]; ok {
		return dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(v)
	}
	return dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(0)
}

// InspectTemplateInspectConfigContentOptionsEnumToProto converts a InspectTemplateInspectConfigContentOptionsEnum enum to its proto representation.
func DlpInspectTemplateInspectConfigContentOptionsEnumToProto(e *dlp.InspectTemplateInspectConfigContentOptionsEnum) dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum {
	if e == nil {
		return dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum(0)
	}
	if v, ok := dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum_value["InspectTemplateInspectConfigContentOptionsEnum"+string(*e)]; ok {
		return dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum(v)
	}
	return dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum(0)
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum enum to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(e *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	if e == nil {
		return dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
	}
	if v, ok := dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum_value["InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum"+string(*e)]; ok {
		return dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(v)
	}
	return dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(0)
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum enum to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(e *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	if e == nil {
		return dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
	}
	if v, ok := dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum_value["InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum"+string(*e)]; ok {
		return dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(v)
	}
	return dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(0)
}

// InspectTemplateInspectConfigToProto converts a InspectTemplateInspectConfig object to its proto representation.
func DlpInspectTemplateInspectConfigToProto(o *dlp.InspectTemplateInspectConfig) *dlppb.DlpInspectTemplateInspectConfig {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfig{}
	p.SetMinLikelihood(DlpInspectTemplateInspectConfigMinLikelihoodEnumToProto(o.MinLikelihood))
	p.SetLimits(DlpInspectTemplateInspectConfigLimitsToProto(o.Limits))
	p.SetIncludeQuote(dcl.ValueOrEmptyBool(o.IncludeQuote))
	p.SetExcludeInfoTypes(dcl.ValueOrEmptyBool(o.ExcludeInfoTypes))
	sInfoTypes := make([]*dlppb.DlpInspectTemplateInspectConfigInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpInspectTemplateInspectConfigInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sCustomInfoTypes := make([]*dlppb.DlpInspectTemplateInspectConfigCustomInfoTypes, len(o.CustomInfoTypes))
	for i, r := range o.CustomInfoTypes {
		sCustomInfoTypes[i] = DlpInspectTemplateInspectConfigCustomInfoTypesToProto(&r)
	}
	p.SetCustomInfoTypes(sCustomInfoTypes)
	sContentOptions := make([]dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum, len(o.ContentOptions))
	for i, r := range o.ContentOptions {
		sContentOptions[i] = dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum(dlppb.DlpInspectTemplateInspectConfigContentOptionsEnum_value[string(r)])
	}
	p.SetContentOptions(sContentOptions)
	sRuleSet := make([]*dlppb.DlpInspectTemplateInspectConfigRuleSet, len(o.RuleSet))
	for i, r := range o.RuleSet {
		sRuleSet[i] = DlpInspectTemplateInspectConfigRuleSetToProto(&r)
	}
	p.SetRuleSet(sRuleSet)
	return p
}

// InspectTemplateInspectConfigInfoTypesToProto converts a InspectTemplateInspectConfigInfoTypes object to its proto representation.
func DlpInspectTemplateInspectConfigInfoTypesToProto(o *dlp.InspectTemplateInspectConfigInfoTypes) *dlppb.DlpInspectTemplateInspectConfigInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigLimitsToProto converts a InspectTemplateInspectConfigLimits object to its proto representation.
func DlpInspectTemplateInspectConfigLimitsToProto(o *dlp.InspectTemplateInspectConfigLimits) *dlppb.DlpInspectTemplateInspectConfigLimits {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigLimits{}
	p.SetMaxFindingsPerItem(dcl.ValueOrEmptyInt64(o.MaxFindingsPerItem))
	p.SetMaxFindingsPerRequest(dcl.ValueOrEmptyInt64(o.MaxFindingsPerRequest))
	sMaxFindingsPerInfoType := make([]*dlppb.DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, len(o.MaxFindingsPerInfoType))
	for i, r := range o.MaxFindingsPerInfoType {
		sMaxFindingsPerInfoType[i] = DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto(&r)
	}
	p.SetMaxFindingsPerInfoType(sMaxFindingsPerInfoType)
	return p
}

// InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType object to its proto representation.
func DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeToProto(o *dlp.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) *dlppb.DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}
	p.SetInfoType(DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o.InfoType))
	p.SetMaxFindings(dcl.ValueOrEmptyInt64(o.MaxFindings))
	return p
}

// InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto converts a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType object to its proto representation.
func DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeToProto(o *dlp.InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *dlppb.DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesToProto converts a InspectTemplateInspectConfigCustomInfoTypes object to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesToProto(o *dlp.InspectTemplateInspectConfigCustomInfoTypes) *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigCustomInfoTypes{}
	p.SetInfoType(DlpInspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto(o.InfoType))
	p.SetLikelihood(DlpInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumToProto(o.Likelihood))
	p.SetDictionary(DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpInspectTemplateInspectConfigCustomInfoTypesRegexToProto(o.Regex))
	p.SetSurrogateType(DlpInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto(o.SurrogateType))
	p.SetStoredType(DlpInspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto(o.StoredType))
	p.SetExclusionType(DlpInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumToProto(o.ExclusionType))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesInfoType object to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesInfoTypeToProto(o *dlp.InspectTemplateInspectConfigCustomInfoTypesInfoType) *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesInfoType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesInfoType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionary object to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryToProto(o *dlp.InspectTemplateInspectConfigCustomInfoTypesDictionary) *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionary {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionary{}
	p.SetWordList(DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList object to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListToProto(o *dlp.InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto converts a InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath object to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathToProto(o *dlp.InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesRegexToProto converts a InspectTemplateInspectConfigCustomInfoTypesRegex object to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesRegexToProto(o *dlp.InspectTemplateInspectConfigCustomInfoTypesRegex) *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesRegex {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesSurrogateType object to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeToProto(o *dlp.InspectTemplateInspectConfigCustomInfoTypesSurrogateType) *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	return p
}

// InspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto converts a InspectTemplateInspectConfigCustomInfoTypesStoredType object to its proto representation.
func DlpInspectTemplateInspectConfigCustomInfoTypesStoredTypeToProto(o *dlp.InspectTemplateInspectConfigCustomInfoTypesStoredType) *dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesStoredType {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigCustomInfoTypesStoredType{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// InspectTemplateInspectConfigRuleSetToProto converts a InspectTemplateInspectConfigRuleSet object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetToProto(o *dlp.InspectTemplateInspectConfigRuleSet) *dlppb.DlpInspectTemplateInspectConfigRuleSet {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSet{}
	sInfoTypes := make([]*dlppb.DlpInspectTemplateInspectConfigRuleSetInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpInspectTemplateInspectConfigRuleSetInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	sRules := make([]*dlppb.DlpInspectTemplateInspectConfigRuleSetRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = DlpInspectTemplateInspectConfigRuleSetRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// InspectTemplateInspectConfigRuleSetInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetInfoTypes object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetInfoTypesToProto(o *dlp.InspectTemplateInspectConfigRuleSetInfoTypes) *dlppb.DlpInspectTemplateInspectConfigRuleSetInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesToProto converts a InspectTemplateInspectConfigRuleSetRules object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesToProto(o *dlp.InspectTemplateInspectConfigRuleSetRules) *dlppb.DlpInspectTemplateInspectConfigRuleSetRules {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRules{}
	p.SetHotwordRule(DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto(o.HotwordRule))
	p.SetExclusionRule(DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto(o.ExclusionRule))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRule object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRule) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	p.SetHotwordRegex(DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o.HotwordRegex))
	p.SetProximity(DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto(o.Proximity))
	p.SetLikelihoodAdjustment(DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o.LikelihoodAdjustment))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	p.SetWindowBefore(dcl.ValueOrEmptyInt64(o.WindowBefore))
	p.SetWindowAfter(dcl.ValueOrEmptyInt64(o.WindowAfter))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto converts a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	p.SetFixedLikelihood(DlpInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumToProto(o.FixedLikelihood))
	p.SetRelativeLikelihood(dcl.ValueOrEmptyInt64(o.RelativeLikelihood))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRule object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRule) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	p.SetDictionary(DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o.Dictionary))
	p.SetRegex(DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto(o.Regex))
	p.SetExcludeInfoTypes(DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o.ExcludeInfoTypes))
	p.SetMatchingType(DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumToProto(o.MatchingType))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	p.SetWordList(DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o.WordList))
	p.SetCloudStoragePath(DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o.CloudStoragePath))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	sWords := make([]string, len(o.Words))
	for i, r := range o.Words {
		sWords[i] = r
	}
	p.SetWords(sWords)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	p.SetPattern(dcl.ValueOrEmptyString(o.Pattern))
	sGroupIndexes := make([]int64, len(o.GroupIndexes))
	for i, r := range o.GroupIndexes {
		sGroupIndexes[i] = r
	}
	p.SetGroupIndexes(sGroupIndexes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	sInfoTypes := make([]*dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, len(o.InfoTypes))
	for i, r := range o.InfoTypes {
		sInfoTypes[i] = DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(&r)
	}
	p.SetInfoTypes(sInfoTypes)
	return p
}

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto converts a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes object to its proto representation.
func DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesToProto(o *dlp.InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if o == nil {
		return nil
	}
	p := &dlppb.DlpInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// InspectTemplateToProto converts a InspectTemplate resource to its proto representation.
func InspectTemplateToProto(resource *dlp.InspectTemplate) *dlppb.DlpInspectTemplate {
	p := &dlppb.DlpInspectTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetInspectConfig(DlpInspectTemplateInspectConfigToProto(resource.InspectConfig))
	p.SetLocationId(dcl.ValueOrEmptyString(resource.LocationId))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Apply() method.
func (s *InspectTemplateServer) applyInspectTemplate(ctx context.Context, c *dlp.Client, request *dlppb.ApplyDlpInspectTemplateRequest) (*dlppb.DlpInspectTemplate, error) {
	p := ProtoToInspectTemplate(request.GetResource())
	res, err := c.ApplyInspectTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InspectTemplateToProto(res)
	return r, nil
}

// applyDlpInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Apply() method.
func (s *InspectTemplateServer) ApplyDlpInspectTemplate(ctx context.Context, request *dlppb.ApplyDlpInspectTemplateRequest) (*dlppb.DlpInspectTemplate, error) {
	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInspectTemplate(ctx, cl, request)
}

// DeleteInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplate Delete() method.
func (s *InspectTemplateServer) DeleteDlpInspectTemplate(ctx context.Context, request *dlppb.DeleteDlpInspectTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInspectTemplate(ctx, ProtoToInspectTemplate(request.GetResource()))

}

// ListDlpInspectTemplate handles the gRPC request by passing it to the underlying InspectTemplateList() method.
func (s *InspectTemplateServer) ListDlpInspectTemplate(ctx context.Context, request *dlppb.ListDlpInspectTemplateRequest) (*dlppb.ListDlpInspectTemplateResponse, error) {
	cl, err := createConfigInspectTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInspectTemplate(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*dlppb.DlpInspectTemplate
	for _, r := range resources.Items {
		rp := InspectTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &dlppb.ListDlpInspectTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInspectTemplate(ctx context.Context, service_account_file string) (*dlp.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dlp.NewClient(conf), nil
}
