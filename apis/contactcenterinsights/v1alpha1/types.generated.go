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


// +kcc:proto=google.cloud.contactcenterinsights.v1.Issue
type Issue struct {
	// Immutable. The resource name of the issue.
	//  Format:
	//  projects/{project}/locations/{location}/issueModels/{issue_model}/issues/{issue}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Issue.name
	Name *string `json:"name,omitempty"`

	// The representative name for the issue.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Issue.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Representative description of the issue.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Issue.display_description
	DisplayDescription *string `json:"displayDescription,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.Issue
type IssueObservedState struct {
	// Output only. The time at which this issue was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Issue.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time that this issue was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Issue.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Resource names of the sample representative utterances that
	//  match to this issue.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.Issue.sample_utterances
	SampleUtterances []string `json:"sampleUtterances,omitempty"`
}
