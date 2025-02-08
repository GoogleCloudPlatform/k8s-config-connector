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


// +kcc:proto=google.cloud.dialogflow.v2beta1.Connection
type Connection struct {
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Connection.ErrorDetails
type Connection_ErrorDetails struct {

	// The error message provided from SIP trunking auth service
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Connection.ErrorDetails.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SipTrunk
type SipTrunk struct {
	// Identifier. The unique identifier of the SIP trunk.
	//  Format: `projects/<Project ID>/locations/<Location ID>/sipTrunks/<SipTrunk
	//  ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SipTrunk.name
	Name *string `json:"name,omitempty"`

	// Required. The expected hostnames in the peer certificate from partner that
	//  is used for TLS authentication.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SipTrunk.expected_hostname
	ExpectedHostname []string `json:"expectedHostname,omitempty"`

	// Optional. Human readable alias for this trunk.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SipTrunk.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Connection
type ConnectionObservedState struct {
	// Output only. The unique identifier of the SIP Trunk connection.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Connection.connection_id
	ConnectionID *string `json:"connectionID,omitempty"`

	// Output only. State of the connection.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Connection.state
	State *string `json:"state,omitempty"`

	// Output only. When the connection status changed.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Connection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The error details for the connection. Only populated when
	//  authentication errors occur.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Connection.error_details
	ErrorDetails *Connection_ErrorDetails `json:"errorDetails,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Connection.ErrorDetails
type Connection_ErrorDetailsObservedState struct {
	// Output only. The status of the certificate authentication.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Connection.ErrorDetails.certificate_state
	CertificateState *string `json:"certificateState,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.SipTrunk
type SipTrunkObservedState struct {
	// Output only. Connections of the SIP trunk.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.SipTrunk.connections
	Connections []Connection `json:"connections,omitempty"`
}
