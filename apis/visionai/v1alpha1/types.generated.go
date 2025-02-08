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


// +kcc:proto=google.cloud.visionai.v1.Event
type Event struct {
	// Name of the resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Event.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.visionai.v1.Event.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations to allow clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.visionai.v1.Event.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// The clock used for joining streams.
	// +kcc:proto:field=google.cloud.visionai.v1.Event.alignment_clock
	AlignmentClock *string `json:"alignmentClock,omitempty"`

	// Grace period for cleaning up the event. This is the time the controller
	//  waits for before deleting the event. During this period, if there is any
	//  active channel on the event. The deletion of the event after grace_period
	//  will be ignored.
	// +kcc:proto:field=google.cloud.visionai.v1.Event.grace_period
	GracePeriod *string `json:"gracePeriod,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Event
type EventObservedState struct {
	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Event.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Event.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
