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


// +kcc:proto=google.cloud.visionai.v1.Process
type Process struct {
	// The name of resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.name
	Name *string `json:"name,omitempty"`

	// Required. Reference to an existing Analysis resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.analysis
	Analysis *string `json:"analysis,omitempty"`

	// Optional. Attribute overrides of the Analyzers.
	//  Format for each single override item:
	//  "{analyzer_name}:{attribute_key}={value}"
	// +kcc:proto:field=google.cloud.visionai.v1.Process.attribute_overrides
	AttributeOverrides []string `json:"attributeOverrides,omitempty"`

	// Optional. Status of the Process.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.run_status
	RunStatus *RunStatus `json:"runStatus,omitempty"`

	// Optional. Run mode of the Process.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.run_mode
	RunMode *string `json:"runMode,omitempty"`

	// Optional. Event ID of the input/output streams.
	//  This is useful when you have a StreamSource/StreamSink operator in the
	//  Analysis, and you want to manually specify the Event to read from/write to.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.event_id
	EventID *string `json:"eventID,omitempty"`

	// Optional. Optional: Batch ID of the Process.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.batch_id
	BatchID *string `json:"batchID,omitempty"`

	// Optional. Optional: The number of retries for a process in submission mode
	//  the system should try before declaring failure. By default, no retry will
	//  be performed.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.retry_count
	RetryCount *int32 `json:"retryCount,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.RunStatus
type RunStatus struct {
	// The state of the Process.
	// +kcc:proto:field=google.cloud.visionai.v1.RunStatus.state
	State *string `json:"state,omitempty"`

	// The reason of becoming the state.
	// +kcc:proto:field=google.cloud.visionai.v1.RunStatus.reason
	Reason *string `json:"reason,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Process
type ProcessObservedState struct {
	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Process.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
