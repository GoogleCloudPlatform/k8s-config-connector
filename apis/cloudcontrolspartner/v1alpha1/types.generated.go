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


// +kcc:proto=google.cloud.cloudcontrolspartner.v1.EkmConnection
type EkmConnection struct {
	// Resource name of the EKM connection in the format:
	//  projects/{project}/locations/{location}/ekmConnections/{ekm_connection}
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.EkmConnection.connection_name
	ConnectionName *string `json:"connectionName,omitempty"`

	// The connection error that occurred if any
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.EkmConnection.connection_error
	ConnectionError *EkmConnection_ConnectionError `json:"connectionError,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.EkmConnection.ConnectionError
type EkmConnection_ConnectionError struct {
	// The error domain for the error
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.EkmConnection.ConnectionError.error_domain
	ErrorDomain *string `json:"errorDomain,omitempty"`

	// The error message for the error
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.EkmConnection.ConnectionError.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.EkmConnections
type EkmConnections struct {
	// Identifier. Format:
	//  `organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}/ekmConnections`
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.EkmConnections.name
	Name *string `json:"name,omitempty"`

	// The EKM connections associated with the workload
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.EkmConnections.ekm_connections
	EkmConnections []EkmConnection `json:"ekmConnections,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.EkmConnection
type EkmConnectionObservedState struct {
	// Output only. The connection state
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.EkmConnection.connection_state
	ConnectionState *string `json:"connectionState,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.EkmConnections
type EkmConnectionsObservedState struct {
	// The EKM connections associated with the workload
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.EkmConnections.ekm_connections
	EkmConnections []EkmConnectionObservedState `json:"ekmConnections,omitempty"`
}
