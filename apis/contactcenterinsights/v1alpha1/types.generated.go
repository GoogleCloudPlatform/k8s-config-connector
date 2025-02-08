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


// +kcc:proto=google.cloud.contactcenterinsights.v1.QaScorecard
type QaScorecard struct {
	// Identifier. The scorecard name.
	//  Format:
	//  projects/{project}/locations/{location}/qaScorecards/{qa_scorecard}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.name
	Name *string `json:"name,omitempty"`

	// The user-specified display name of the scorecard.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A text description explaining the intent of the scorecard.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.QaScorecard
type QaScorecardObservedState struct {
	// Output only. The time at which this scorecard was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the scorecard was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.QaScorecard.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
