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


// +kcc:proto=google.cloud.visionai.v1.Channel
type Channel struct {
	// Name of the resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Channel.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.visionai.v1.Channel.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations to allow clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.visionai.v1.Channel.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. Stream that is associated with this series.
	// +kcc:proto:field=google.cloud.visionai.v1.Channel.stream
	Stream *string `json:"stream,omitempty"`

	// Required. Event that is associated with this series.
	// +kcc:proto:field=google.cloud.visionai.v1.Channel.event
	Event *string `json:"event,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Channel
type ChannelObservedState struct {
	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Channel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Channel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
