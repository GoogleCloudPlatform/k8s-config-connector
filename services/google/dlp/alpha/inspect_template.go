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

type InspectTemplate struct {
	Name          *string                       `json:"name"`
	DisplayName   *string                       `json:"displayName"`
	Description   *string                       `json:"description"`
	CreateTime    *string                       `json:"createTime"`
	UpdateTime    *string                       `json:"updateTime"`
	InspectConfig *InspectTemplateInspectConfig `json:"inspectConfig"`
	LocationId    *string                       `json:"locationId"`
	Parent        *string                       `json:"parent"`
	Location      *string                       `json:"location"`
}

func (r *InspectTemplate) String() string {
	return dcl.SprintResource(r)
}

// The enum InspectTemplateInspectConfigMinLikelihoodEnum.
type InspectTemplateInspectConfigMinLikelihoodEnum string

// InspectTemplateInspectConfigMinLikelihoodEnumRef returns a *InspectTemplateInspectConfigMinLikelihoodEnum with the value of string s
// If the empty string is provided, nil is returned.
func InspectTemplateInspectConfigMinLikelihoodEnumRef(s string) *InspectTemplateInspectConfigMinLikelihoodEnum {
	v := InspectTemplateInspectConfigMinLikelihoodEnum(s)
	return &v
}

func (v InspectTemplateInspectConfigMinLikelihoodEnum) Validate() error {
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
		Enum:  "InspectTemplateInspectConfigMinLikelihoodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum.
type InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum string

// InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumRef returns a *InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum with the value of string s
// If the empty string is provided, nil is returned.
func InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumRef(s string) *InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	v := InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(s)
	return &v
}

func (v InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum) Validate() error {
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
		Enum:  "InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum.
type InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum string

// InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumRef returns a *InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumRef(s string) *InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	v := InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(s)
	return &v
}

func (v InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum) Validate() error {
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
		Enum:  "InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InspectTemplateInspectConfigContentOptionsEnum.
type InspectTemplateInspectConfigContentOptionsEnum string

// InspectTemplateInspectConfigContentOptionsEnumRef returns a *InspectTemplateInspectConfigContentOptionsEnum with the value of string s
// If the empty string is provided, nil is returned.
func InspectTemplateInspectConfigContentOptionsEnumRef(s string) *InspectTemplateInspectConfigContentOptionsEnum {
	v := InspectTemplateInspectConfigContentOptionsEnum(s)
	return &v
}

func (v InspectTemplateInspectConfigContentOptionsEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"CONTENT_UNSPECIFIED", "CONTENT_TEXT", "CONTENT_IMAGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InspectTemplateInspectConfigContentOptionsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum.
type InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum string

// InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef returns a *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum with the value of string s
// If the empty string is provided, nil is returned.
func InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s string) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	v := InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(s)
	return &v
}

func (v InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum) Validate() error {
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
		Enum:  "InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum.
type InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum string

// InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumRef returns a *InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumRef(s string) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	v := InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(s)
	return &v
}

func (v InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum) Validate() error {
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
		Enum:  "InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InspectTemplateInspectConfig struct {
	empty            bool                                             `json:"-"`
	InfoTypes        []InspectTemplateInspectConfigInfoTypes          `json:"infoTypes"`
	MinLikelihood    *InspectTemplateInspectConfigMinLikelihoodEnum   `json:"minLikelihood"`
	Limits           *InspectTemplateInspectConfigLimits              `json:"limits"`
	IncludeQuote     *bool                                            `json:"includeQuote"`
	ExcludeInfoTypes *bool                                            `json:"excludeInfoTypes"`
	CustomInfoTypes  []InspectTemplateInspectConfigCustomInfoTypes    `json:"customInfoTypes"`
	ContentOptions   []InspectTemplateInspectConfigContentOptionsEnum `json:"contentOptions"`
	RuleSet          []InspectTemplateInspectConfigRuleSet            `json:"ruleSet"`
}

type jsonInspectTemplateInspectConfig InspectTemplateInspectConfig

func (r *InspectTemplateInspectConfig) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfig
	} else {

		r.InfoTypes = res.InfoTypes

		r.MinLikelihood = res.MinLikelihood

		r.Limits = res.Limits

		r.IncludeQuote = res.IncludeQuote

		r.ExcludeInfoTypes = res.ExcludeInfoTypes

		r.CustomInfoTypes = res.CustomInfoTypes

		r.ContentOptions = res.ContentOptions

		r.RuleSet = res.RuleSet

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfig *InspectTemplateInspectConfig = &InspectTemplateInspectConfig{empty: true}

func (r *InspectTemplateInspectConfig) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigInfoTypes struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonInspectTemplateInspectConfigInfoTypes InspectTemplateInspectConfigInfoTypes

func (r *InspectTemplateInspectConfigInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigInfoTypes
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigInfoTypes *InspectTemplateInspectConfigInfoTypes = &InspectTemplateInspectConfigInfoTypes{empty: true}

func (r *InspectTemplateInspectConfigInfoTypes) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigLimits struct {
	empty                  bool                                                       `json:"-"`
	MaxFindingsPerItem     *int64                                                     `json:"maxFindingsPerItem"`
	MaxFindingsPerRequest  *int64                                                     `json:"maxFindingsPerRequest"`
	MaxFindingsPerInfoType []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType `json:"maxFindingsPerInfoType"`
}

type jsonInspectTemplateInspectConfigLimits InspectTemplateInspectConfigLimits

func (r *InspectTemplateInspectConfigLimits) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigLimits
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigLimits
	} else {

		r.MaxFindingsPerItem = res.MaxFindingsPerItem

		r.MaxFindingsPerRequest = res.MaxFindingsPerRequest

		r.MaxFindingsPerInfoType = res.MaxFindingsPerInfoType

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigLimits is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigLimits *InspectTemplateInspectConfigLimits = &InspectTemplateInspectConfigLimits{empty: true}

func (r *InspectTemplateInspectConfigLimits) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigLimits) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigLimits) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType struct {
	empty       bool                                                              `json:"-"`
	InfoType    *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `json:"infoType"`
	MaxFindings *int64                                                            `json:"maxFindings"`
}

type jsonInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType

func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType
	} else {

		r.InfoType = res.InfoType

		r.MaxFindings = res.MaxFindings

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType = &InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{empty: true}

func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType

func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType = &InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{empty: true}

func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigCustomInfoTypes struct {
	empty         bool                                                          `json:"-"`
	InfoType      *InspectTemplateInspectConfigCustomInfoTypesInfoType          `json:"infoType"`
	Likelihood    *InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum    `json:"likelihood"`
	Dictionary    *InspectTemplateInspectConfigCustomInfoTypesDictionary        `json:"dictionary"`
	Regex         *InspectTemplateInspectConfigCustomInfoTypesRegex             `json:"regex"`
	SurrogateType *InspectTemplateInspectConfigCustomInfoTypesSurrogateType     `json:"surrogateType"`
	StoredType    *InspectTemplateInspectConfigCustomInfoTypesStoredType        `json:"storedType"`
	ExclusionType *InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum `json:"exclusionType"`
}

type jsonInspectTemplateInspectConfigCustomInfoTypes InspectTemplateInspectConfigCustomInfoTypes

func (r *InspectTemplateInspectConfigCustomInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigCustomInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigCustomInfoTypes
	} else {

		r.InfoType = res.InfoType

		r.Likelihood = res.Likelihood

		r.Dictionary = res.Dictionary

		r.Regex = res.Regex

		r.SurrogateType = res.SurrogateType

		r.StoredType = res.StoredType

		r.ExclusionType = res.ExclusionType

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigCustomInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigCustomInfoTypes *InspectTemplateInspectConfigCustomInfoTypes = &InspectTemplateInspectConfigCustomInfoTypes{empty: true}

func (r *InspectTemplateInspectConfigCustomInfoTypes) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigCustomInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigCustomInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigCustomInfoTypesInfoType struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonInspectTemplateInspectConfigCustomInfoTypesInfoType InspectTemplateInspectConfigCustomInfoTypesInfoType

func (r *InspectTemplateInspectConfigCustomInfoTypesInfoType) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigCustomInfoTypesInfoType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigCustomInfoTypesInfoType
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigCustomInfoTypesInfoType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigCustomInfoTypesInfoType *InspectTemplateInspectConfigCustomInfoTypesInfoType = &InspectTemplateInspectConfigCustomInfoTypesInfoType{empty: true}

func (r *InspectTemplateInspectConfigCustomInfoTypesInfoType) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigCustomInfoTypesInfoType) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigCustomInfoTypesInfoType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigCustomInfoTypesDictionary struct {
	empty            bool                                                                   `json:"-"`
	WordList         *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList         `json:"wordList"`
	CloudStoragePath *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath `json:"cloudStoragePath"`
}

type jsonInspectTemplateInspectConfigCustomInfoTypesDictionary InspectTemplateInspectConfigCustomInfoTypesDictionary

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionary) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigCustomInfoTypesDictionary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigCustomInfoTypesDictionary
	} else {

		r.WordList = res.WordList

		r.CloudStoragePath = res.CloudStoragePath

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigCustomInfoTypesDictionary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigCustomInfoTypesDictionary *InspectTemplateInspectConfigCustomInfoTypesDictionary = &InspectTemplateInspectConfigCustomInfoTypesDictionary{empty: true}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionary) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionary) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList struct {
	empty bool     `json:"-"`
	Words []string `json:"words"`
}

type jsonInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList
	} else {

		r.Words = res.Words

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList = &InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{empty: true}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath struct {
	empty bool    `json:"-"`
	Path  *string `json:"path"`
}

type jsonInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	} else {

		r.Path = res.Path

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath = &InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{empty: true}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigCustomInfoTypesRegex struct {
	empty        bool    `json:"-"`
	Pattern      *string `json:"pattern"`
	GroupIndexes []int64 `json:"groupIndexes"`
}

type jsonInspectTemplateInspectConfigCustomInfoTypesRegex InspectTemplateInspectConfigCustomInfoTypesRegex

func (r *InspectTemplateInspectConfigCustomInfoTypesRegex) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigCustomInfoTypesRegex
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigCustomInfoTypesRegex
	} else {

		r.Pattern = res.Pattern

		r.GroupIndexes = res.GroupIndexes

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigCustomInfoTypesRegex is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigCustomInfoTypesRegex *InspectTemplateInspectConfigCustomInfoTypesRegex = &InspectTemplateInspectConfigCustomInfoTypesRegex{empty: true}

func (r *InspectTemplateInspectConfigCustomInfoTypesRegex) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigCustomInfoTypesRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigCustomInfoTypesRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigCustomInfoTypesSurrogateType struct {
	empty bool `json:"-"`
}

type jsonInspectTemplateInspectConfigCustomInfoTypesSurrogateType InspectTemplateInspectConfigCustomInfoTypesSurrogateType

func (r *InspectTemplateInspectConfigCustomInfoTypesSurrogateType) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigCustomInfoTypesSurrogateType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigCustomInfoTypesSurrogateType
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigCustomInfoTypesSurrogateType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigCustomInfoTypesSurrogateType *InspectTemplateInspectConfigCustomInfoTypesSurrogateType = &InspectTemplateInspectConfigCustomInfoTypesSurrogateType{empty: true}

func (r *InspectTemplateInspectConfigCustomInfoTypesSurrogateType) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigCustomInfoTypesSurrogateType) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigCustomInfoTypesSurrogateType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigCustomInfoTypesStoredType struct {
	empty      bool    `json:"-"`
	Name       *string `json:"name"`
	CreateTime *string `json:"createTime"`
}

type jsonInspectTemplateInspectConfigCustomInfoTypesStoredType InspectTemplateInspectConfigCustomInfoTypesStoredType

func (r *InspectTemplateInspectConfigCustomInfoTypesStoredType) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigCustomInfoTypesStoredType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigCustomInfoTypesStoredType
	} else {

		r.Name = res.Name

		r.CreateTime = res.CreateTime

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigCustomInfoTypesStoredType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigCustomInfoTypesStoredType *InspectTemplateInspectConfigCustomInfoTypesStoredType = &InspectTemplateInspectConfigCustomInfoTypesStoredType{empty: true}

func (r *InspectTemplateInspectConfigCustomInfoTypesStoredType) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigCustomInfoTypesStoredType) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigCustomInfoTypesStoredType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSet struct {
	empty     bool                                           `json:"-"`
	InfoTypes []InspectTemplateInspectConfigRuleSetInfoTypes `json:"infoTypes"`
	Rules     []InspectTemplateInspectConfigRuleSetRules     `json:"rules"`
}

type jsonInspectTemplateInspectConfigRuleSet InspectTemplateInspectConfigRuleSet

func (r *InspectTemplateInspectConfigRuleSet) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSet
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSet
	} else {

		r.InfoTypes = res.InfoTypes

		r.Rules = res.Rules

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSet is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSet *InspectTemplateInspectConfigRuleSet = &InspectTemplateInspectConfigRuleSet{empty: true}

func (r *InspectTemplateInspectConfigRuleSet) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSet) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSet) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetInfoTypes struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonInspectTemplateInspectConfigRuleSetInfoTypes InspectTemplateInspectConfigRuleSetInfoTypes

func (r *InspectTemplateInspectConfigRuleSetInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetInfoTypes
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetInfoTypes *InspectTemplateInspectConfigRuleSetInfoTypes = &InspectTemplateInspectConfigRuleSetInfoTypes{empty: true}

func (r *InspectTemplateInspectConfigRuleSetInfoTypes) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRules struct {
	empty         bool                                                   `json:"-"`
	HotwordRule   *InspectTemplateInspectConfigRuleSetRulesHotwordRule   `json:"hotwordRule"`
	ExclusionRule *InspectTemplateInspectConfigRuleSetRulesExclusionRule `json:"exclusionRule"`
}

type jsonInspectTemplateInspectConfigRuleSetRules InspectTemplateInspectConfigRuleSetRules

func (r *InspectTemplateInspectConfigRuleSetRules) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRules
	} else {

		r.HotwordRule = res.HotwordRule

		r.ExclusionRule = res.ExclusionRule

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRules *InspectTemplateInspectConfigRuleSetRules = &InspectTemplateInspectConfigRuleSetRules{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRules) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRules) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesHotwordRule struct {
	empty                bool                                                                     `json:"-"`
	HotwordRegex         *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex         `json:"hotwordRegex"`
	Proximity            *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity            `json:"proximity"`
	LikelihoodAdjustment *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment `json:"likelihoodAdjustment"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesHotwordRule InspectTemplateInspectConfigRuleSetRulesHotwordRule

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRule) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesHotwordRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRule
	} else {

		r.HotwordRegex = res.HotwordRegex

		r.Proximity = res.Proximity

		r.LikelihoodAdjustment = res.LikelihoodAdjustment

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesHotwordRule is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRule *InspectTemplateInspectConfigRuleSetRulesHotwordRule = &InspectTemplateInspectConfigRuleSetRulesHotwordRule{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRule) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRule) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex struct {
	empty        bool    `json:"-"`
	Pattern      *string `json:"pattern"`
	GroupIndexes []int64 `json:"groupIndexes"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	} else {

		r.Pattern = res.Pattern

		r.GroupIndexes = res.GroupIndexes

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity struct {
	empty        bool   `json:"-"`
	WindowBefore *int64 `json:"windowBefore"`
	WindowAfter  *int64 `json:"windowAfter"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity
	} else {

		r.WindowBefore = res.WindowBefore

		r.WindowAfter = res.WindowAfter

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment struct {
	empty              bool                                                                                        `json:"-"`
	FixedLikelihood    *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum `json:"fixedLikelihood"`
	RelativeLikelihood *int64                                                                                      `json:"relativeLikelihood"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	} else {

		r.FixedLikelihood = res.FixedLikelihood

		r.RelativeLikelihood = res.RelativeLikelihood

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesExclusionRule struct {
	empty            bool                                                                   `json:"-"`
	Dictionary       *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary       `json:"dictionary"`
	Regex            *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex            `json:"regex"`
	ExcludeInfoTypes *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes `json:"excludeInfoTypes"`
	MatchingType     *InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum `json:"matchingType"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesExclusionRule InspectTemplateInspectConfigRuleSetRulesExclusionRule

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRule) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesExclusionRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRule
	} else {

		r.Dictionary = res.Dictionary

		r.Regex = res.Regex

		r.ExcludeInfoTypes = res.ExcludeInfoTypes

		r.MatchingType = res.MatchingType

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesExclusionRule is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRule *InspectTemplateInspectConfigRuleSetRulesExclusionRule = &InspectTemplateInspectConfigRuleSetRulesExclusionRule{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRule) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRule) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary struct {
	empty            bool                                                                             `json:"-"`
	WordList         *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList         `json:"wordList"`
	CloudStoragePath *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath `json:"cloudStoragePath"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary
	} else {

		r.WordList = res.WordList

		r.CloudStoragePath = res.CloudStoragePath

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList struct {
	empty bool     `json:"-"`
	Words []string `json:"words"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	} else {

		r.Words = res.Words

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath struct {
	empty bool    `json:"-"`
	Path  *string `json:"path"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	} else {

		r.Path = res.Path

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex struct {
	empty        bool    `json:"-"`
	Pattern      *string `json:"pattern"`
	GroupIndexes []int64 `json:"groupIndexes"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex
	} else {

		r.Pattern = res.Pattern

		r.GroupIndexes = res.GroupIndexes

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes struct {
	empty     bool                                                                             `json:"-"`
	InfoTypes []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes `json:"infoTypes"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	} else {

		r.InfoTypes = res.InfoTypes

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) UnmarshalJSON(data []byte) error {
	var res jsonInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{empty: true}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) Empty() bool {
	return r.empty
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *InspectTemplate) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dlp",
		Type:    "InspectTemplate",
		Version: "alpha",
	}
}

func (r *InspectTemplate) ID() (string, error) {
	if err := extractInspectTemplateFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":           dcl.ValueOrEmptyString(nr.Name),
		"display_name":   dcl.ValueOrEmptyString(nr.DisplayName),
		"description":    dcl.ValueOrEmptyString(nr.Description),
		"create_time":    dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":    dcl.ValueOrEmptyString(nr.UpdateTime),
		"inspect_config": dcl.ValueOrEmptyString(nr.InspectConfig),
		"location_id":    dcl.ValueOrEmptyString(nr.LocationId),
		"parent":         dcl.ValueOrEmptyString(nr.Parent),
		"location":       dcl.ValueOrEmptyString(nr.Location),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.Nprintf("{{parent}}/locations/{{location}}/inspectTemplates/{{name}}", params), nil
	}

	return dcl.Nprintf("{{parent}}/inspectTemplates/{{name}}", params), nil
}

const InspectTemplateMaxPage = -1

type InspectTemplateList struct {
	Items []*InspectTemplate

	nextToken string

	pageSize int32

	resource *InspectTemplate
}

func (l *InspectTemplateList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InspectTemplateList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInspectTemplate(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInspectTemplate(ctx context.Context, location, parent string) (*InspectTemplateList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListInspectTemplateWithMaxResults(ctx, location, parent, InspectTemplateMaxPage)

}

func (c *Client) ListInspectTemplateWithMaxResults(ctx context.Context, location, parent string, pageSize int32) (*InspectTemplateList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &InspectTemplate{
		Location: &location,
		Parent:   &parent,
	}
	items, token, err := c.listInspectTemplate(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InspectTemplateList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetInspectTemplate(ctx context.Context, r *InspectTemplate) (*InspectTemplate, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractInspectTemplateFields(r)

	b, err := c.getInspectTemplateRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInspectTemplate(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Location = r.Location
	result.Parent = r.Parent
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInspectTemplateNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractInspectTemplateFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInspectTemplate(ctx context.Context, r *InspectTemplate) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("InspectTemplate resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting InspectTemplate...")
	deleteOp := deleteInspectTemplateOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInspectTemplate deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInspectTemplate(ctx context.Context, location, parent string, filter func(*InspectTemplate) bool) error {
	listObj, err := c.ListInspectTemplate(ctx, location, parent)
	if err != nil {
		return err
	}

	err = c.deleteAllInspectTemplate(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInspectTemplate(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInspectTemplate(ctx context.Context, rawDesired *InspectTemplate, opts ...dcl.ApplyOption) (*InspectTemplate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *InspectTemplate
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyInspectTemplateHelper(c, ctx, rawDesired, opts...)
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

func applyInspectTemplateHelper(c *Client, ctx context.Context, rawDesired *InspectTemplate, opts ...dcl.ApplyOption) (*InspectTemplate, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyInspectTemplate...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractInspectTemplateFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.inspectTemplateDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToInspectTemplateDiffs(c.Config, fieldDiffs, opts)
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
	var ops []inspectTemplateApiOperation
	if create {
		ops = append(ops, &createInspectTemplateOperation{})
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
	return applyInspectTemplateDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyInspectTemplateDiff(c *Client, ctx context.Context, desired *InspectTemplate, rawDesired *InspectTemplate, ops []inspectTemplateApiOperation, opts ...dcl.ApplyOption) (*InspectTemplate, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetInspectTemplate(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createInspectTemplateOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapInspectTemplate(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeInspectTemplateNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInspectTemplateNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInspectTemplateDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractInspectTemplateFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractInspectTemplateFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInspectTemplate(c, newDesired, newState)
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
