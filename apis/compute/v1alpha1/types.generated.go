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

// +kcc:proto=google.cloud.compute.v1.InterconnectCircuitInfo
type InterconnectCircuitInfo struct {
	// Customer-side demarc ID for this circuit.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectCircuitInfo.customer_demarc_id
	CustomerDemarcID *string `json:"customerDemarcID,omitempty"`

	// Google-assigned unique ID for this circuit. Assigned at circuit turn-up.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectCircuitInfo.google_circuit_id
	GoogleCircuitID *string `json:"googleCircuitID,omitempty"`

	// Google-side demarc ID for this circuit. Assigned at circuit turn-up and provided by Google to the customer in the LOA.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectCircuitInfo.google_demarc_id
	GoogleDemarcID *string `json:"googleDemarcID,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InterconnectMacsec
type InterconnectMacsec struct {
	// If set to true, the Interconnect connection is configured with a should-secure MACsec security policy, that allows the Google router to fallback to cleartext traffic if the MKA session cannot be established. By default, the Interconnect connection is configured with a must-secure security policy that drops all traffic if the MKA session cannot be established with your router.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectMacsec.fail_open
	FailOpen *bool `json:"failOpen,omitempty"`

	// Required. A keychain placeholder describing a set of named key objects along with their start times. A MACsec CKN/CAK is generated for each key in the key chain. Google router automatically picks the key with the most recent startTime when establishing or re-establishing a MACsec secure link.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectMacsec.pre_shared_keys
	PreSharedKeys []InterconnectMacsecPreSharedKey `json:"preSharedKeys,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InterconnectMacsecPreSharedKey
type InterconnectMacsecPreSharedKey struct {
	// Required. A name for this pre-shared key. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectMacsecPreSharedKey.name
	Name *string `json:"name,omitempty"`

	// A RFC3339 timestamp on or after which the key is valid. startTime can be in the future. If the keychain has a single key, startTime can be omitted. If the keychain has multiple keys, startTime is mandatory for each key. The start times of keys must be in increasing order. The start times of two consecutive keys must be at least 6 hours apart.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectMacsecPreSharedKey.start_time
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.InterconnectOutageNotification
type InterconnectOutageNotification struct {
	// If issue_type is IT_PARTIAL_OUTAGE, a list of the Google-side circuit IDs that will be affected.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.affected_circuits
	AffectedCircuits []string `json:"affectedCircuits,omitempty"`

	// A description about the purpose of the outage.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.description
	Description *string `json:"description,omitempty"`

	// Scheduled end time for the outage (milliseconds since Unix epoch).
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.end_time
	EndTime *int64 `json:"endTime,omitempty"`

	// Form this outage is expected to take, which can take one of the following values: - OUTAGE: The Interconnect may be completely out of service for some or all of the specified window. - PARTIAL_OUTAGE: Some circuits comprising the Interconnect as a whole should remain up, but with reduced bandwidth. Note that the versions of this enum prefixed with "IT_" have been deprecated in favor of the unprefixed values.
	//  Check the IssueType enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.issue_type
	IssueType *string `json:"issueType,omitempty"`

	// Unique identifier for this outage notification.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.name
	Name *string `json:"name,omitempty"`

	// The party that generated this notification, which can take the following value: - GOOGLE: this notification as generated by Google. Note that the value of NSRC_GOOGLE has been deprecated in favor of GOOGLE.
	//  Check the Source enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.source
	Source *string `json:"source,omitempty"`

	// Scheduled start time for the outage (milliseconds since Unix epoch).
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.start_time
	StartTime *int64 `json:"startTime,omitempty"`

	// State of this notification, which can take one of the following values: - ACTIVE: This outage notification is active. The event could be in the past, present, or future. See start_time and end_time for scheduling. - CANCELLED: The outage associated with this notification was cancelled before the outage was due to start. - COMPLETED: The outage associated with this notification is complete. Note that the versions of this enum prefixed with "NS_" have been deprecated in favor of the unprefixed values.
	//  Check the State enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InterconnectOutageNotification.state
	State *string `json:"state,omitempty"`
}
