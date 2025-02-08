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


// +kcc:proto=google.cloud.connectors.v1.RuntimeConfig
type RuntimeConfig struct {
}

// +kcc:proto=google.cloud.connectors.v1.RuntimeConfig
type RuntimeConfigObservedState struct {
	// Output only. location_id of the runtime location. E.g. "us-west1".
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.location_id
	LocationID *string `json:"locationID,omitempty"`

	// Output only. Pub/Sub topic for connd to send message.
	//  E.g. projects/{project-id}/topics/{topic-id}
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.connd_topic
	ConndTopic *string `json:"conndTopic,omitempty"`

	// Output only. Pub/Sub subscription for connd to receive message.
	//  E.g. projects/{project-id}/subscriptions/{topic-id}
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.connd_subscription
	ConndSubscription *string `json:"conndSubscription,omitempty"`

	// Output only. Pub/Sub topic for control plne to send message.
	//  communication.
	//  E.g. projects/{project-id}/topics/{topic-id}
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.control_plane_topic
	ControlPlaneTopic *string `json:"controlPlaneTopic,omitempty"`

	// Output only. Pub/Sub subscription for control plane to receive message.
	//  E.g. projects/{project-id}/subscriptions/{topic-id}
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.control_plane_subscription
	ControlPlaneSubscription *string `json:"controlPlaneSubscription,omitempty"`

	// Output only. The endpoint of the connectors runtime ingress.
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.runtime_endpoint
	RuntimeEndpoint *string `json:"runtimeEndpoint,omitempty"`

	// Output only. The state of the location.
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.state
	State *string `json:"state,omitempty"`

	// Output only. The Cloud Storage bucket that stores connector's schema
	//  reports.
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.schema_gcs_bucket
	SchemaGcsBucket *string `json:"schemaGcsBucket,omitempty"`

	// Output only. The name of the Service Directory service name.
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.service_directory
	ServiceDirectory *string `json:"serviceDirectory,omitempty"`

	// Output only. Name of the runtimeConfig resource.
	//  Format: projects/{project}/locations/{location}/runtimeConfig
	// +kcc:proto:field=google.cloud.connectors.v1.RuntimeConfig.name
	Name *string `json:"name,omitempty"`
}
