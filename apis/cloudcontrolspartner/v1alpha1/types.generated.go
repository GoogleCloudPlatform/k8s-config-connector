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


// +kcc:proto=google.cloud.cloudcontrolspartner.v1.AccessApprovalRequest
type AccessApprovalRequest struct {
	// Identifier. Format:
	//  `organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}/accessApprovalRequests/{access_approval_request}`
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.AccessApprovalRequest.name
	Name *string `json:"name,omitempty"`

	// The time at which approval was requested.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.AccessApprovalRequest.request_time
	RequestTime *string `json:"requestTime,omitempty"`

	// The justification for which approval is being requested.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.AccessApprovalRequest.requested_reason
	RequestedReason *AccessReason `json:"requestedReason,omitempty"`

	// The requested expiration for the approval. If the request is approved,
	//  access will be granted from the time of approval until the expiration time.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.AccessApprovalRequest.requested_expiration_time
	RequestedExpirationTime *string `json:"requestedExpirationTime,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.AccessReason
type AccessReason struct {
	// Type of access justification.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.AccessReason.type
	Type *string `json:"type,omitempty"`

	// More detail about certain reason types. See comments for each type above.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.AccessReason.detail
	Detail *string `json:"detail,omitempty"`
}
