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


// +kcc:proto=google.cloud.speech.v1p1beta1.PhraseSet
type PhraseSet struct {
	// The resource name of the phrase set.
	// +kcc:proto:field=google.cloud.speech.v1p1beta1.PhraseSet.name
	Name *string `json:"name,omitempty"`

	// A list of word and phrases.
	// +kcc:proto:field=google.cloud.speech.v1p1beta1.PhraseSet.phrases
	Phrases []PhraseSet_Phrase `json:"phrases,omitempty"`

	// Hint Boost. Positive value will increase the probability that a specific
	//  phrase will be recognized over other similar sounding phrases. The higher
	//  the boost, the higher the chance of false positive recognition as well.
	//  Negative boost values would correspond to anti-biasing. Anti-biasing is not
	//  enabled, so negative boost will simply be ignored. Though `boost` can
	//  accept a wide range of positive values, most use cases are best served with
	//  values between 0 (exclusive) and 20. We recommend using a binary search
	//  approach to finding the optimal value for your use case as well as adding
	//  phrases both with and without boost to your requests.
	// +kcc:proto:field=google.cloud.speech.v1p1beta1.PhraseSet.boost
	Boost *float32 `json:"boost,omitempty"`
}

// +kcc:proto=google.cloud.speech.v1p1beta1.PhraseSet.Phrase
type PhraseSet_Phrase struct {
	// The phrase itself.
	// +kcc:proto:field=google.cloud.speech.v1p1beta1.PhraseSet.Phrase.value
	Value *string `json:"value,omitempty"`

	// Hint Boost. Overrides the boost set at the phrase set level.
	//  Positive value will increase the probability that a specific phrase will
	//  be recognized over other similar sounding phrases. The higher the boost,
	//  the higher the chance of false positive recognition as well. Negative
	//  boost will simply be ignored. Though `boost` can accept a wide range of
	//  positive values, most use cases are best served
	//  with values between 0 and 20. We recommend using a binary search approach
	//  to finding the optimal value for your use case as well as adding
	//  phrases both with and without boost to your requests.
	// +kcc:proto:field=google.cloud.speech.v1p1beta1.PhraseSet.Phrase.boost
	Boost *float32 `json:"boost,omitempty"`
}
