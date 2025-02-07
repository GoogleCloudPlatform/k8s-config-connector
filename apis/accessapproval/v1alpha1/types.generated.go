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


// +kcc:proto=google.cloud.accessapproval.v1.AccessLocations
type AccessLocations struct {
	// The "home office" location of the principal. A two-letter country code
	//  (ISO 3166-1 alpha-2), such as "US", "DE" or "GB" or a region code. In some
	//  limited situations Google systems may refer refer to a region code instead
	//  of a country code.
	//  Possible Region Codes:
	//
	//    * ASI: Asia
	//    * EUR: Europe
	//    * OCE: Oceania
	//    * AFR: Africa
	//    * NAM: North America
	//    * SAM: South America
	//    * ANT: Antarctica
	//    * ANY: Any location
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessLocations.principal_office_country
	PrincipalOfficeCountry *string `json:"principalOfficeCountry,omitempty"`

	// Physical location of the principal at the time of the access. A
	//  two-letter country code (ISO 3166-1 alpha-2), such as "US", "DE" or "GB" or
	//  a region code. In some limited situations Google systems may refer refer to
	//  a region code instead of a country code.
	//  Possible Region Codes:
	//
	//    * ASI: Asia
	//    * EUR: Europe
	//    * OCE: Oceania
	//    * AFR: Africa
	//    * NAM: North America
	//    * SAM: South America
	//    * ANT: Antarctica
	//    * ANY: Any location
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessLocations.principal_physical_location_country
	PrincipalPhysicalLocationCountry *string `json:"principalPhysicalLocationCountry,omitempty"`
}

// +kcc:proto=google.cloud.accessapproval.v1.AccessReason
type AccessReason struct {
	// Type of access justification.
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessReason.type
	Type *string `json:"type,omitempty"`

	// More detail about certain reason types. See comments for each type above.
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessReason.detail
	Detail *string `json:"detail,omitempty"`
}

// +kcc:proto=google.cloud.accessapproval.v1.ApprovalRequest
type ApprovalRequest struct {
	// The resource name of the request. Format is
	//  "{projects|folders|organizations}/{id}/approvalRequests/{approval_request}".
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.name
	Name *string `json:"name,omitempty"`

	// The resource for which approval is being requested. The format of the
	//  resource name is defined at
	//  https://cloud.google.com/apis/design/resource_names. The resource name here
	//  may either be a "full" resource name (e.g.
	//  "//library.googleapis.com/shelves/shelf1/books/book2") or a "relative"
	//  resource name (e.g. "shelves/shelf1/books/book2") as described in the
	//  resource name specification.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.requested_resource_name
	RequestedResourceName *string `json:"requestedResourceName,omitempty"`

	// Properties related to the resource represented by requested_resource_name.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.requested_resource_properties
	RequestedResourceProperties *ResourceProperties `json:"requestedResourceProperties,omitempty"`

	// The justification for which approval is being requested.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.requested_reason
	RequestedReason *AccessReason `json:"requestedReason,omitempty"`

	// The locations for which approval is being requested.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.requested_locations
	RequestedLocations *AccessLocations `json:"requestedLocations,omitempty"`

	// The time at which approval was requested.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.request_time
	RequestTime *string `json:"requestTime,omitempty"`

	// The requested expiration for the approval. If the request is approved,
	//  access will be granted from the time of approval until the expiration time.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.requested_expiration
	RequestedExpiration *string `json:"requestedExpiration,omitempty"`

	// Access was approved.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.approve
	Approve *ApproveDecision `json:"approve,omitempty"`

	// The request was dismissed.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApprovalRequest.dismiss
	Dismiss *DismissDecision `json:"dismiss,omitempty"`
}

// +kcc:proto=google.cloud.accessapproval.v1.ApproveDecision
type ApproveDecision struct {
	// The time at which approval was granted.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApproveDecision.approve_time
	ApproveTime *string `json:"approveTime,omitempty"`

	// The time at which the approval expires.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApproveDecision.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// If set, denotes the timestamp at which the approval is invalidated.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApproveDecision.invalidate_time
	InvalidateTime *string `json:"invalidateTime,omitempty"`

	// The signature for the ApprovalRequest and details on how it was signed.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApproveDecision.signature_info
	SignatureInfo *SignatureInfo `json:"signatureInfo,omitempty"`

	// True when the request has been auto-approved.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ApproveDecision.auto_approved
	AutoApproved *bool `json:"autoApproved,omitempty"`
}

// +kcc:proto=google.cloud.accessapproval.v1.DismissDecision
type DismissDecision struct {
	// The time at which the approval request was dismissed.
	// +kcc:proto:field=google.cloud.accessapproval.v1.DismissDecision.dismiss_time
	DismissTime *string `json:"dismissTime,omitempty"`

	// This field will be true if the ApprovalRequest was implicitly dismissed due
	//  to inaction by the access approval approvers (the request is not acted
	//  on by the approvers before the exiration time).
	// +kcc:proto:field=google.cloud.accessapproval.v1.DismissDecision.implicit
	Implicit *bool `json:"implicit,omitempty"`
}

// +kcc:proto=google.cloud.accessapproval.v1.ResourceProperties
type ResourceProperties struct {
	// Whether an approval will exclude the descendants of the resource being
	//  requested.
	// +kcc:proto:field=google.cloud.accessapproval.v1.ResourceProperties.excludes_descendants
	ExcludesDescendants *bool `json:"excludesDescendants,omitempty"`
}

// +kcc:proto=google.cloud.accessapproval.v1.SignatureInfo
type SignatureInfo struct {
	// The digital signature.
	// +kcc:proto:field=google.cloud.accessapproval.v1.SignatureInfo.signature
	Signature []byte `json:"signature,omitempty"`

	// The public key for the Google default signing, encoded in PEM format. The
	//  signature was created using a private key which may be verified using
	//  this public key.
	// +kcc:proto:field=google.cloud.accessapproval.v1.SignatureInfo.google_public_key_pem
	GooglePublicKeyPem *string `json:"googlePublicKeyPem,omitempty"`

	// The resource name of the customer CryptoKeyVersion used for signing.
	// +kcc:proto:field=google.cloud.accessapproval.v1.SignatureInfo.customer_kms_key_version
	CustomerKMSKeyVersion *string `json:"customerKMSKeyVersion,omitempty"`
}
