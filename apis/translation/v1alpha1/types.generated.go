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


// +kcc:proto=google.cloud.translation.v3.GlossaryEntry
type GlossaryEntry struct {
	// Identifier. The resource name of the entry.
	//  Format:
	//    `projects/*/locations/*/glossaries/*/glossaryEntries/*`
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.name
	Name *string `json:"name,omitempty"`

	// Used for an unidirectional glossary.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.terms_pair
	TermsPair *GlossaryEntry_GlossaryTermsPair `json:"termsPair,omitempty"`

	// Used for an equivalent term sets glossary.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.terms_set
	TermsSet *GlossaryEntry_GlossaryTermsSet `json:"termsSet,omitempty"`

	// Describes the glossary entry.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.GlossaryEntry.GlossaryTermsPair
type GlossaryEntry_GlossaryTermsPair struct {
	// The source term is the term that will get match in the text,
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.GlossaryTermsPair.source_term
	SourceTerm *GlossaryTerm `json:"sourceTerm,omitempty"`

	// The term that will replace the match source term.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.GlossaryTermsPair.target_term
	TargetTerm *GlossaryTerm `json:"targetTerm,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.GlossaryEntry.GlossaryTermsSet
type GlossaryEntry_GlossaryTermsSet struct {
	// Each term in the set represents a term that can be replaced by the other
	//  terms.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.GlossaryTermsSet.terms
	Terms []GlossaryTerm `json:"terms,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.GlossaryTerm
type GlossaryTerm struct {
	// The language for this glossary term.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryTerm.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The text for the glossary term.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryTerm.text
	Text *string `json:"text,omitempty"`
}
