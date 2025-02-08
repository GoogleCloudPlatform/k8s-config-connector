// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.contactcenterinsights.v1.ExactMatchConfig
type ExactMatchConfig struct {
	// Whether to consider case sensitivity when performing an exact match.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.ExactMatchConfig.case_sensitive
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.PhraseMatchRule
type PhraseMatchRule struct {
	// Required. The phrase to be matched.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatchRule.query
	Query *string `json:"query,omitempty"`

	// Specifies whether the phrase must be missing from the transcript segment or
	//  present in the transcript segment.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatchRule.negated
	Negated *bool `json:"negated,omitempty"`

	// Provides additional information about the rule that specifies how to apply
	//  the rule.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatchRule.config
	Config *PhraseMatchRuleConfig `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.PhraseMatchRuleConfig
type PhraseMatchRuleConfig struct {
	// The configuration for the exact match rule.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatchRuleConfig.exact_match_config
	ExactMatchConfig *ExactMatchConfig `json:"exactMatchConfig,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.PhraseMatchRuleGroup
type PhraseMatchRuleGroup struct {
	// Required. The type of this phrase match rule group.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatchRuleGroup.type
	Type *string `json:"type,omitempty"`

	// A list of phrase match rules that are included in this group.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatchRuleGroup.phrase_match_rules
	PhraseMatchRules []PhraseMatchRule `json:"phraseMatchRules,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.PhraseMatcher
type PhraseMatcher struct {
	// The resource name of the phrase matcher.
	//  Format:
	//  projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.name
	Name *string `json:"name,omitempty"`

	// The customized version tag to use for the phrase matcher. If not specified,
	//  it will default to `revision_id`.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.version_tag
	VersionTag *string `json:"versionTag,omitempty"`

	// The human-readable name of the phrase matcher.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The type of this phrase matcher.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.type
	Type *string `json:"type,omitempty"`

	// Applies the phrase matcher only when it is active.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.active
	Active *bool `json:"active,omitempty"`

	// A list of phase match rule groups that are included in this matcher.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.phrase_match_rule_groups
	PhraseMatchRuleGroups []PhraseMatchRuleGroup `json:"phraseMatchRuleGroups,omitempty"`

	// The role whose utterances the phrase matcher should be matched
	//  against. If the role is ROLE_UNSPECIFIED it will be matched against any
	//  utterances in the transcript.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.role_match
	RoleMatch *string `json:"roleMatch,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.PhraseMatcher
type PhraseMatcherObservedState struct {
	// Output only. Immutable. The revision ID of the phrase matcher.
	//  A new revision is committed whenever the matcher is changed, except when it
	//  is activated or deactivated. A server generated random ID will be used.
	//  Example: locations/global/phraseMatchers/my-first-matcher@1234567
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp of when the revision was created. It is also the
	//  create time when a new matcher is added.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. The most recent time at which the activation status was
	//  updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.activation_update_time
	ActivationUpdateTime *string `json:"activationUpdateTime,omitempty"`

	// Output only. The most recent time at which the phrase matcher was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.PhraseMatcher.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
