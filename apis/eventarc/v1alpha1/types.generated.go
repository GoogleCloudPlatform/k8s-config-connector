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

// +kcc:proto=google.cloud.eventarc.v1.Channel
type Channel struct {
	// Required. The resource name of the channel. Must be unique within the
	//  location on the project and must be in
	//  `projects/{project}/locations/{location}/channels/{channel_id}` format.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.name
	//+required
	Name *string `json:"name,omitempty"`

	// The name of the event provider (e.g. Eventarc SaaS partner) associated
	//  with the channel. This provider will be granted permissions to publish
	//  events to the channel. Format:
	//  `projects/{project}/locations/{location}/providers/{provider_id}`.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.provider
	Provider *string `json:"provider,omitempty"`

	// Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt their event data.
	//
	//  It must match the pattern
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.crypto_key_name
	CryptoKeyName *string `json:"cryptoKeyName,omitempty"`
}

	// Required. The resource name of the channel. Must be unique within the
	//  location on the project and must be in
	//  `projects/{project}/locations/{location}/channels/{channel_id}` format.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.name
	//+required
	Name *string `json:"name,omitempty"`

	// The name of the event provider (e.g. Eventarc SaaS partner) associated
	//  with the channel. This provider will be granted permissions to publish
	//  events to the channel. Format:
	//  `projects/{project}/locations/{location}/providers/{provider_id}`.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.provider
	Provider *string `json:"provider,omitempty"`

	// Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt their event data.
	//
	//  It must match the pattern
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.crypto_key_name
	CryptoKeyName *string `json:"cryptoKeyName,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Channel
type ChannelObservedState struct {
	// Output only. Server assigned unique identifier for the channel. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The name of the Pub/Sub topic created and managed by
	//  Eventarc system as a transport for the event delivery. Format:
	//  `projects/{project}/topics/{topic_id}`.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.pubsub_topic
	PubsubTopic *string `json:"pubsubTopic,omitempty"`

	// Output only. The state of a Channel.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.state
	State *string `json:"state,omitempty"`

	// Output only. The activation token for the channel. The token must be used
	//  by the provider to register the channel for publishing.
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.activation_token
	ActivationToken *string `json:"activationToken,omitempty"`

	// Output only. Whether or not this Channel satisfies the requirements of
	//  physical zone separation
	// +kcc:proto:field=google.cloud.eventarc.v1.Channel.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}
