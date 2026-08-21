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
package beta

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Policy struct {
	AdmissionWhitelistPatterns             []PolicyAdmissionWhitelistPatterns                      `json:"admissionWhitelistPatterns"`
	ClusterAdmissionRules                  map[string]PolicyClusterAdmissionRules                  `json:"clusterAdmissionRules"`
	KubernetesNamespaceAdmissionRules      map[string]PolicyKubernetesNamespaceAdmissionRules      `json:"kubernetesNamespaceAdmissionRules"`
	KubernetesServiceAccountAdmissionRules map[string]PolicyKubernetesServiceAccountAdmissionRules `json:"kubernetesServiceAccountAdmissionRules"`
	IstioServiceIdentityAdmissionRules     map[string]PolicyIstioServiceIdentityAdmissionRules     `json:"istioServiceIdentityAdmissionRules"`
	DefaultAdmissionRule                   *PolicyDefaultAdmissionRule                             `json:"defaultAdmissionRule"`
	Description                            *string                                                 `json:"description"`
	GlobalPolicyEvaluationMode             *PolicyGlobalPolicyEvaluationModeEnum                   `json:"globalPolicyEvaluationMode"`
	SelfLink                               *string                                                 `json:"selfLink"`
	Project                                *string                                                 `json:"project"`
	UpdateTime                             *string                                                 `json:"updateTime"`
}

func (r *Policy) String() string {
	return dcl.SprintResource(r)
}

// The enum PolicyClusterAdmissionRulesEvaluationModeEnum.
type PolicyClusterAdmissionRulesEvaluationModeEnum string

// PolicyClusterAdmissionRulesEvaluationModeEnumRef returns a *PolicyClusterAdmissionRulesEvaluationModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyClusterAdmissionRulesEvaluationModeEnumRef(s string) *PolicyClusterAdmissionRulesEvaluationModeEnum {
	v := PolicyClusterAdmissionRulesEvaluationModeEnum(s)
	return &v
}

func (v PolicyClusterAdmissionRulesEvaluationModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ALWAYS_ALLOW", "ALWAYS_DENY", "REQUIRE_ATTESTATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyClusterAdmissionRulesEvaluationModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyClusterAdmissionRulesEnforcementModeEnum.
type PolicyClusterAdmissionRulesEnforcementModeEnum string

// PolicyClusterAdmissionRulesEnforcementModeEnumRef returns a *PolicyClusterAdmissionRulesEnforcementModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyClusterAdmissionRulesEnforcementModeEnumRef(s string) *PolicyClusterAdmissionRulesEnforcementModeEnum {
	v := PolicyClusterAdmissionRulesEnforcementModeEnum(s)
	return &v
}

func (v PolicyClusterAdmissionRulesEnforcementModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ENFORCEMENT_MODE_UNSPECIFIED", "ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyClusterAdmissionRulesEnforcementModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum.
type PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum string

// PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumRef returns a *PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumRef(s string) *PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	v := PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(s)
	return &v
}

func (v PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ALWAYS_ALLOW", "ALWAYS_DENY", "REQUIRE_ATTESTATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum.
type PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum string

// PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumRef returns a *PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumRef(s string) *PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	v := PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(s)
	return &v
}

func (v PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ENFORCEMENT_MODE_UNSPECIFIED", "ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum.
type PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum string

// PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumRef returns a *PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumRef(s string) *PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	v := PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(s)
	return &v
}

func (v PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ALWAYS_ALLOW", "ALWAYS_DENY", "REQUIRE_ATTESTATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum.
type PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum string

// PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumRef returns a *PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumRef(s string) *PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	v := PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(s)
	return &v
}

func (v PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ENFORCEMENT_MODE_UNSPECIFIED", "ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum.
type PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum string

// PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumRef returns a *PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumRef(s string) *PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	v := PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(s)
	return &v
}

func (v PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ALWAYS_ALLOW", "ALWAYS_DENY", "REQUIRE_ATTESTATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum.
type PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum string

// PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumRef returns a *PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumRef(s string) *PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	v := PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(s)
	return &v
}

func (v PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ENFORCEMENT_MODE_UNSPECIFIED", "ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyDefaultAdmissionRuleEvaluationModeEnum.
type PolicyDefaultAdmissionRuleEvaluationModeEnum string

// PolicyDefaultAdmissionRuleEvaluationModeEnumRef returns a *PolicyDefaultAdmissionRuleEvaluationModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyDefaultAdmissionRuleEvaluationModeEnumRef(s string) *PolicyDefaultAdmissionRuleEvaluationModeEnum {
	v := PolicyDefaultAdmissionRuleEvaluationModeEnum(s)
	return &v
}

func (v PolicyDefaultAdmissionRuleEvaluationModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ALWAYS_ALLOW", "ALWAYS_DENY", "REQUIRE_ATTESTATION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyDefaultAdmissionRuleEvaluationModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyDefaultAdmissionRuleEnforcementModeEnum.
type PolicyDefaultAdmissionRuleEnforcementModeEnum string

// PolicyDefaultAdmissionRuleEnforcementModeEnumRef returns a *PolicyDefaultAdmissionRuleEnforcementModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyDefaultAdmissionRuleEnforcementModeEnumRef(s string) *PolicyDefaultAdmissionRuleEnforcementModeEnum {
	v := PolicyDefaultAdmissionRuleEnforcementModeEnum(s)
	return &v
}

func (v PolicyDefaultAdmissionRuleEnforcementModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ENFORCEMENT_MODE_UNSPECIFIED", "ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyDefaultAdmissionRuleEnforcementModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum PolicyGlobalPolicyEvaluationModeEnum.
type PolicyGlobalPolicyEvaluationModeEnum string

// PolicyGlobalPolicyEvaluationModeEnumRef returns a *PolicyGlobalPolicyEvaluationModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func PolicyGlobalPolicyEvaluationModeEnumRef(s string) *PolicyGlobalPolicyEvaluationModeEnum {
	v := PolicyGlobalPolicyEvaluationModeEnum(s)
	return &v
}

func (v PolicyGlobalPolicyEvaluationModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED", "ENABLE", "DISABLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "PolicyGlobalPolicyEvaluationModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type PolicyAdmissionWhitelistPatterns struct {
	empty       bool    `json:"-"`
	NamePattern *string `json:"namePattern"`
}

type jsonPolicyAdmissionWhitelistPatterns PolicyAdmissionWhitelistPatterns

func (r *PolicyAdmissionWhitelistPatterns) UnmarshalJSON(data []byte) error {
	var res jsonPolicyAdmissionWhitelistPatterns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPolicyAdmissionWhitelistPatterns
	} else {

		r.NamePattern = res.NamePattern

	}
	return nil
}

// This object is used to assert a desired state where this PolicyAdmissionWhitelistPatterns is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPolicyAdmissionWhitelistPatterns *PolicyAdmissionWhitelistPatterns = &PolicyAdmissionWhitelistPatterns{empty: true}

func (r *PolicyAdmissionWhitelistPatterns) Empty() bool {
	return r.empty
}

func (r *PolicyAdmissionWhitelistPatterns) String() string {
	return dcl.SprintResource(r)
}

func (r *PolicyAdmissionWhitelistPatterns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PolicyClusterAdmissionRules struct {
	empty                 bool                                            `json:"-"`
	EvaluationMode        *PolicyClusterAdmissionRulesEvaluationModeEnum  `json:"evaluationMode"`
	RequireAttestationsBy []string                                        `json:"requireAttestationsBy"`
	EnforcementMode       *PolicyClusterAdmissionRulesEnforcementModeEnum `json:"enforcementMode"`
}

type jsonPolicyClusterAdmissionRules PolicyClusterAdmissionRules

func (r *PolicyClusterAdmissionRules) UnmarshalJSON(data []byte) error {
	var res jsonPolicyClusterAdmissionRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPolicyClusterAdmissionRules
	} else {

		r.EvaluationMode = res.EvaluationMode

		r.RequireAttestationsBy = res.RequireAttestationsBy

		r.EnforcementMode = res.EnforcementMode

	}
	return nil
}

// This object is used to assert a desired state where this PolicyClusterAdmissionRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPolicyClusterAdmissionRules *PolicyClusterAdmissionRules = &PolicyClusterAdmissionRules{empty: true}

func (r *PolicyClusterAdmissionRules) Empty() bool {
	return r.empty
}

func (r *PolicyClusterAdmissionRules) String() string {
	return dcl.SprintResource(r)
}

func (r *PolicyClusterAdmissionRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PolicyKubernetesNamespaceAdmissionRules struct {
	empty                 bool                                                        `json:"-"`
	EvaluationMode        *PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum  `json:"evaluationMode"`
	RequireAttestationsBy []string                                                    `json:"requireAttestationsBy"`
	EnforcementMode       *PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum `json:"enforcementMode"`
}

type jsonPolicyKubernetesNamespaceAdmissionRules PolicyKubernetesNamespaceAdmissionRules

func (r *PolicyKubernetesNamespaceAdmissionRules) UnmarshalJSON(data []byte) error {
	var res jsonPolicyKubernetesNamespaceAdmissionRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPolicyKubernetesNamespaceAdmissionRules
	} else {

		r.EvaluationMode = res.EvaluationMode

		r.RequireAttestationsBy = res.RequireAttestationsBy

		r.EnforcementMode = res.EnforcementMode

	}
	return nil
}

// This object is used to assert a desired state where this PolicyKubernetesNamespaceAdmissionRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPolicyKubernetesNamespaceAdmissionRules *PolicyKubernetesNamespaceAdmissionRules = &PolicyKubernetesNamespaceAdmissionRules{empty: true}

func (r *PolicyKubernetesNamespaceAdmissionRules) Empty() bool {
	return r.empty
}

func (r *PolicyKubernetesNamespaceAdmissionRules) String() string {
	return dcl.SprintResource(r)
}

func (r *PolicyKubernetesNamespaceAdmissionRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PolicyKubernetesServiceAccountAdmissionRules struct {
	empty                 bool                                                             `json:"-"`
	EvaluationMode        *PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum  `json:"evaluationMode"`
	RequireAttestationsBy []string                                                         `json:"requireAttestationsBy"`
	EnforcementMode       *PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum `json:"enforcementMode"`
}

type jsonPolicyKubernetesServiceAccountAdmissionRules PolicyKubernetesServiceAccountAdmissionRules

func (r *PolicyKubernetesServiceAccountAdmissionRules) UnmarshalJSON(data []byte) error {
	var res jsonPolicyKubernetesServiceAccountAdmissionRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPolicyKubernetesServiceAccountAdmissionRules
	} else {

		r.EvaluationMode = res.EvaluationMode

		r.RequireAttestationsBy = res.RequireAttestationsBy

		r.EnforcementMode = res.EnforcementMode

	}
	return nil
}

// This object is used to assert a desired state where this PolicyKubernetesServiceAccountAdmissionRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPolicyKubernetesServiceAccountAdmissionRules *PolicyKubernetesServiceAccountAdmissionRules = &PolicyKubernetesServiceAccountAdmissionRules{empty: true}

func (r *PolicyKubernetesServiceAccountAdmissionRules) Empty() bool {
	return r.empty
}

func (r *PolicyKubernetesServiceAccountAdmissionRules) String() string {
	return dcl.SprintResource(r)
}

func (r *PolicyKubernetesServiceAccountAdmissionRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PolicyIstioServiceIdentityAdmissionRules struct {
	empty                 bool                                                         `json:"-"`
	EvaluationMode        *PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum  `json:"evaluationMode"`
	RequireAttestationsBy []string                                                     `json:"requireAttestationsBy"`
	EnforcementMode       *PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum `json:"enforcementMode"`
}

type jsonPolicyIstioServiceIdentityAdmissionRules PolicyIstioServiceIdentityAdmissionRules

func (r *PolicyIstioServiceIdentityAdmissionRules) UnmarshalJSON(data []byte) error {
	var res jsonPolicyIstioServiceIdentityAdmissionRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPolicyIstioServiceIdentityAdmissionRules
	} else {

		r.EvaluationMode = res.EvaluationMode

		r.RequireAttestationsBy = res.RequireAttestationsBy

		r.EnforcementMode = res.EnforcementMode

	}
	return nil
}

// This object is used to assert a desired state where this PolicyIstioServiceIdentityAdmissionRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPolicyIstioServiceIdentityAdmissionRules *PolicyIstioServiceIdentityAdmissionRules = &PolicyIstioServiceIdentityAdmissionRules{empty: true}

func (r *PolicyIstioServiceIdentityAdmissionRules) Empty() bool {
	return r.empty
}

func (r *PolicyIstioServiceIdentityAdmissionRules) String() string {
	return dcl.SprintResource(r)
}

func (r *PolicyIstioServiceIdentityAdmissionRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type PolicyDefaultAdmissionRule struct {
	empty                 bool                                           `json:"-"`
	EvaluationMode        *PolicyDefaultAdmissionRuleEvaluationModeEnum  `json:"evaluationMode"`
	RequireAttestationsBy []string                                       `json:"requireAttestationsBy"`
	EnforcementMode       *PolicyDefaultAdmissionRuleEnforcementModeEnum `json:"enforcementMode"`
}

type jsonPolicyDefaultAdmissionRule PolicyDefaultAdmissionRule

func (r *PolicyDefaultAdmissionRule) UnmarshalJSON(data []byte) error {
	var res jsonPolicyDefaultAdmissionRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyPolicyDefaultAdmissionRule
	} else {

		r.EvaluationMode = res.EvaluationMode

		r.RequireAttestationsBy = res.RequireAttestationsBy

		r.EnforcementMode = res.EnforcementMode

	}
	return nil
}

// This object is used to assert a desired state where this PolicyDefaultAdmissionRule is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyPolicyDefaultAdmissionRule *PolicyDefaultAdmissionRule = &PolicyDefaultAdmissionRule{empty: true}

func (r *PolicyDefaultAdmissionRule) Empty() bool {
	return r.empty
}

func (r *PolicyDefaultAdmissionRule) String() string {
	return dcl.SprintResource(r)
}

func (r *PolicyDefaultAdmissionRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Policy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "binary_authorization",
		Type:    "Policy",
		Version: "beta",
	}
}

func (r *Policy) ID() (string, error) {
	if err := extractPolicyFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"admission_whitelist_patterns":               dcl.ValueOrEmptyString(nr.AdmissionWhitelistPatterns),
		"cluster_admission_rules":                    dcl.ValueOrEmptyString(nr.ClusterAdmissionRules),
		"kubernetes_namespace_admission_rules":       dcl.ValueOrEmptyString(nr.KubernetesNamespaceAdmissionRules),
		"kubernetes_service_account_admission_rules": dcl.ValueOrEmptyString(nr.KubernetesServiceAccountAdmissionRules),
		"istio_service_identity_admission_rules":     dcl.ValueOrEmptyString(nr.IstioServiceIdentityAdmissionRules),
		"default_admission_rule":                     dcl.ValueOrEmptyString(nr.DefaultAdmissionRule),
		"description":                                dcl.ValueOrEmptyString(nr.Description),
		"global_policy_evaluation_mode":              dcl.ValueOrEmptyString(nr.GlobalPolicyEvaluationMode),
		"self_link":                                  dcl.ValueOrEmptyString(nr.SelfLink),
		"project":                                    dcl.ValueOrEmptyString(nr.Project),
		"update_time":                                dcl.ValueOrEmptyString(nr.UpdateTime),
	}
	return dcl.Nprintf("projects/{{project}}/policy", params), nil
}

const PolicyMaxPage = -1

type PolicyList struct {
	Items []*Policy

	nextToken string

	resource *Policy
}

func (c *Client) GetPolicy(ctx context.Context, r *Policy) (*Policy, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractPolicyFields(r)

	b, err := c.getPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalPolicy(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizePolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractPolicyFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyPolicy(ctx context.Context, rawDesired *Policy, opts ...dcl.ApplyOption) (*Policy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Policy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyPolicyHelper(c, ctx, rawDesired, opts...)
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

func applyPolicyHelper(c *Client, ctx context.Context, rawDesired *Policy, opts ...dcl.ApplyOption) (*Policy, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyPolicy...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractPolicyFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.policyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToPolicyDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		return nil, dcl.ApplyInfeasibleError{Message: "No initial state found for singleton resource."}
	} else {
		for _, d := range diffs {
			if d.UpdateOp == nil {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) no update method found for field", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}
	var ops []policyApiOperation
	for _, d := range diffs {
		ops = append(ops, d.UpdateOp)
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
	return applyPolicyDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyPolicyDiff(c *Client, ctx context.Context, desired *Policy, rawDesired *Policy, ops []policyApiOperation, opts ...dcl.ApplyOption) (*Policy, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetPolicy(ctx, desired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizePolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizePolicyDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractPolicyFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractPolicyFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffPolicy(c, newDesired, newState)
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

func (r *Policy) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
