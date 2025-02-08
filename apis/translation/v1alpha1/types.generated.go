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


// +kcc:proto=google.cloud.translation.v3.GcsSource
type GcsSource struct {
	// Required. Source data URI. For example, `gs://my_bucket/my_object`.
	// +kcc:proto:field=google.cloud.translation.v3.GcsSource.input_uri
	InputURI *string `json:"inputURI,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.Glossary
type Glossary struct {
	// Required. The resource name of the glossary. Glossary names have the form
	//  `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.name
	Name *string `json:"name,omitempty"`

	// Used with unidirectional glossaries.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.language_pair
	LanguagePair *Glossary_LanguageCodePair `json:"languagePair,omitempty"`

	// Used with equivalent term set glossaries.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.language_codes_set
	LanguageCodesSet *Glossary_LanguageCodesSet `json:"languageCodesSet,omitempty"`

	// Required. Provides examples to build the glossary from.
	//  Total glossary must not exceed 10M Unicode codepoints.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.input_config
	InputConfig *GlossaryInputConfig `json:"inputConfig,omitempty"`

	// Optional. The display name of the glossary.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.Glossary.LanguageCodePair
type Glossary_LanguageCodePair struct {
	// Required. The ISO-639 language code of the input text, for example,
	//  "en-US". Expected to be an exact match for GlossaryTerm.language_code.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.LanguageCodePair.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// Required. The ISO-639 language code for translation output, for example,
	//  "zh-CN". Expected to be an exact match for GlossaryTerm.language_code.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.LanguageCodePair.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.Glossary.LanguageCodesSet
type Glossary_LanguageCodesSet struct {
	// The ISO-639 language code(s) for terms defined in the glossary.
	//  All entries are unique. The list contains at least two entries.
	//  Expected to be an exact match for GlossaryTerm.language_code.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.LanguageCodesSet.language_codes
	LanguageCodes []string `json:"languageCodes,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.GlossaryInputConfig
type GlossaryInputConfig struct {
	// Required. Google Cloud Storage location of glossary data.
	//  File format is determined based on the filename extension. API returns
	//  [google.rpc.Code.INVALID_ARGUMENT] for unsupported URI-s and file
	//  formats. Wildcards are not allowed. This must be a single file in one of
	//  the following formats:
	//
	//  For unidirectional glossaries:
	//
	//  - TSV/CSV (`.tsv`/`.csv`): Two column file, tab- or comma-separated.
	//    The first column is source text. The second column is target text.
	//    No headers in this file. The first row contains data and not column
	//    names.
	//
	//  - TMX (`.tmx`): TMX file with parallel data defining source/target term
	//  pairs.
	//
	//  For equivalent term sets glossaries:
	//
	//  - CSV (`.csv`): Multi-column CSV file defining equivalent glossary terms
	//    in multiple languages. See documentation for more information -
	//    [glossaries](https://cloud.google.com/translate/docs/advanced/glossary).
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryInputConfig.gcs_source
	GcsSource *GcsSource `json:"gcsSource,omitempty"`
}

// +kcc:proto=google.cloud.translation.v3.Glossary
type GlossaryObservedState struct {
	// Output only. The number of entries defined in the glossary.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.entry_count
	EntryCount *int32 `json:"entryCount,omitempty"`

	// Output only. When CreateGlossary was called.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.submit_time
	SubmitTime *string `json:"submitTime,omitempty"`

	// Output only. When the glossary creation was finished.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.end_time
	EndTime *string `json:"endTime,omitempty"`
}
