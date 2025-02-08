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


// +kcc:proto=google.cloud.eventarc.v1.ChannelConnection
type ChannelConnection struct {
	// Required. The name of the connection.
	// +kcc:proto:field=google.cloud.eventarc.v1.ChannelConnection.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the connected subscriber Channel.
	//  This is a weak reference to avoid cross project and cross accounts
	//  references. This must be in
	//  `projects/{project}/location/{location}/channels/{channel_id}` format.
	// +kcc:proto:field=google.cloud.eventarc.v1.ChannelConnection.channel
	Channel *string `json:"channel,omitempty"`

	// Input only. Activation token for the channel. The token will be used
	//  during the creation of ChannelConnection to bind the channel with the
	//  provider project. This field will not be stored in the provider resource.
	// +kcc:proto:field=google.cloud.eventarc.v1.ChannelConnection.activation_token
	ActivationToken *string `json:"activationToken,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.ChannelConnection
type ChannelConnectionObservedState struct {
	// Output only. Server assigned ID of the resource.
	//  The server guarantees uniqueness and immutability until deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.ChannelConnection.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.ChannelConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.ChannelConnection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
