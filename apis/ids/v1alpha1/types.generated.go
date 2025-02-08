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


// +kcc:proto=google.cloud.ids.v1.Endpoint
type Endpoint struct {

	// The labels of the endpoint.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The fully qualified URL of the network to which the IDS Endpoint is
	//  attached.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.network
	Network *string `json:"network,omitempty"`

	// User-provided description of the endpoint
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.description
	Description *string `json:"description,omitempty"`

	// Required. Lowest threat severity that this endpoint will alert on.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.severity
	Severity *string `json:"severity,omitempty"`

	// Whether the endpoint should report traffic logs in addition to threat logs.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.traffic_logs
	TrafficLogs *bool `json:"trafficLogs,omitempty"`
}

// +kcc:proto=google.cloud.ids.v1.Endpoint
type EndpointObservedState struct {
	// Output only. The name of the endpoint.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.name
	Name *string `json:"name,omitempty"`

	// Output only. The create time timestamp.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time timestamp.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The fully qualified URL of the endpoint's ILB Forwarding Rule.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.endpoint_forwarding_rule
	EndpointForwardingRule *string `json:"endpointForwardingRule,omitempty"`

	// Output only. The IP address of the IDS Endpoint's ILB.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.endpoint_ip
	EndpointIP *string `json:"endpointIP,omitempty"`

	// Output only. Current state of the endpoint.
	// +kcc:proto:field=google.cloud.ids.v1.Endpoint.state
	State *string `json:"state,omitempty"`
}
