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


// +kcc:proto=google.cloud.datastream.v1.Route
type Route struct {

	// Labels.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Destination address for connection
	// +kcc:proto:field=google.cloud.datastream.v1.Route.destination_address
	DestinationAddress *string `json:"destinationAddress,omitempty"`

	// Destination port for connection
	// +kcc:proto:field=google.cloud.datastream.v1.Route.destination_port
	DestinationPort *int32 `json:"destinationPort,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.Route
type RouteObservedState struct {
	// Output only. The resource's name.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.name
	Name *string `json:"name,omitempty"`

	// Output only. The create time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
