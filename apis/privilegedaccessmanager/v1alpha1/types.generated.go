// Copyright 2024 Google LLC
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

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// AccessControlEntry is used to control who can do some operation.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.AccessControlEntry
type AccessControlEntry struct {
	// Optional. Users who are allowed for the operation. Each entry should be a
	// valid v1 IAM principal identifier. The format for these is documented at:
	// https://cloud.google.com/iam/docs/principal-identifiers#v1
	// +required
	Principals []string `json:"principals,omitempty"`
}

// ApprovalWorkflow represents different types of approval workflows that can be
// used to gate privileged access granting.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.ApprovalWorkflow
type ApprovalWorkflow struct {
	// An approval workflow where users designated as approvers review and act
	// on the grants.
	// +required
	ManualApprovals *ManualApprovals `json:"manualApprovals,omitempty"`
}

// AdditionalNotificationTargets includes email addresses to be notified.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.Entitlement.AdditionalNotificationTargets
type AdditionalNotificationTargets struct {
	// Optional. Additional email addresses to be notified when a principal
	// (requester) is granted access.
	// +optional
	AdminEmailRecipients []string `json:"adminEmailRecipients,omitempty"`

	// Optional. Additional email address to be notified about an eligible
	// entitlement.
	// +optional
	RequesterEmailRecipients []string `json:"requesterEmailRecipients,omitempty"`
}

// RequesterJustificationConfig defines how a requester must provide a
// justification when requesting access.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.Entitlement.RequesterJustificationConfig
type RequesterJustificationConfig struct {
	// NotMandatory justification type means the justification isn't required
	// and can be provided in any of the supported formats. The user must
	// explicitly opt out using this field if a justification from the requester
	// isn't mandatory.
	// The only accepted value is `{}` (empty struct).
	// Either 'notMandatory' or 'unstructured' field must be set.
	// +optional
	NotMandatory *runtime.RawExtension `json:"notMandatory,omitempty"`

	// Unstructured justification type means the justification is in the format
	// of a string. If this is set, the server allows the requester to provide a
	// justification but doesn't validate it.
	// The only accepted value is `{}` (empty struct).
	// Either 'notMandatory' or 'unstructured' field must be set.
	// +optional
	Unstructured *runtime.RawExtension `json:"unstructured,omitempty"`
}

// ManualApprovals represent the manual approval workflow where users who are
// designated as approvers need to call the 'ApproveGrant'/'DenyGrant' APIs for
// a grant. The workflow can consist of multiple serial steps where each step
// defines who can act as the approver in that step and how many of those users
// should approve before the workflow moves to the next step.
//
// This can be used to create approval workflows such as:
//
// * Require an approval from any user in a group G.
// * Require an approval from any k number of users from a Group G.
// * Require an approval from any user in a group G and then from a user U.
//
// A single user might be part of the 'approvers' ACL for multiple steps in this
// workflow, but they can only approve once and that approval is only considered
// to satisfy the approval step at which it was granted.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.ManualApprovals
type ManualApprovals struct {
	// Optional. Whether the approvers need to provide a justification for their
	// actions.
	// +optional
	RequireApproverJustification *bool `json:"requireApproverJustification,omitempty"`

	// Optional. List of approval steps in this workflow. These steps are followed
	// in the specified order sequentially. Only 1 step is supported.
	// +optional
	Steps []Step `json:"step,omitempty"`
}

// Step represents a logical step in a manual approval workflow.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.ManualApprovals.Step
type Step struct {
	// Optional. The potential set of approvers in this step. This list must
	// contain at most one entry.
	// +optional
	Approvers []AccessControlEntry `json:"approvers,omitempty"`

	// Required. How many users from the above list need to approve. If there
	// aren't enough distinct users in the list, then the workflow indefinitely
	// blocks. Should always be greater than 0. 1 is the only supported value.
	// +required
	ApprovalsNeeded *int32 `json:"approvalsNeeded,omitempty"`

	// Optional. Additional email addresses to be notified when a grant is
	// pending approval.
	// +optional
	ApproverEmailRecipients []string `json:"approverEmailRecipients,omitempty"`
}

// Privileged access that this service can be used to gate.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.PrivilegedAccess
type PrivilegedAccess struct {
	// Access to a Google Cloud resource through IAM.
	// +required
	GcpIAMAccess *GcpIamAccess `json:"gcpIAMAccess,omitempty"`
}

// GcpIamAccess represents IAM based access control on a Google Cloud
// resource. Refer to https://cloud.google.com/iam/docs to understand more
// about IAM.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.PrivilegedAccess.GcpIamAccess
type GcpIamAccess struct {
	/* We decided to hide the following fields. Here are the context:

	   1. Currently, in the API, the resource information under gcpIamAccess is
	      duplicate with the parent information in the URL. I.e. Resolved values
	      of 'projectRef', 'folderRef', and 'organizationRef' here must be the
	      same as the resolved values of `spec.projectRef', 'spec.folderRef',
	      and 'spec.organizationRef'.
	   2. It's better not to ask users to specify the same information
	      repetitively: It can avoid confusion and potential mistakes.
	   3. Even if the API behavior will be changed, i.e. the resource
	      information here no longer needs to match the parent information in
	      the URL, we can always support it in the future without breaking
	      existing resources. Adding a field later is cheaper than adding a
	      field in the beginning and making changes in the future, and can avoid
	      breaking changes.
	   4. We are under the discussion with the API team to finalize the schema
	      design. Hiding the fields gives us the flexibility to graduate the
	      resource to Beta while we are still deciding on the most appropriate
	      schema.

	// Required. The type of this resource.
	// +required
	ResourceType *string `json:"resourceType,omitempty"`

	// The Project that this privileged access is granted to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// The Folder that this privileged access is granted to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	// The Organization that this privileged access is granted to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`
	*/

	// Required. Role bindings that are created on successful grant.
	// +required
	RoleBindings []RoleBinding `json:"roleBindings,omitempty"`
}

// RoleBinding represents IAM role bindings that are created after a successful
// grant.
// +kcc:proto=google.cloud.privilegedaccessmanager.v1.PrivilegedAccess.GcpIamAccess.RoleBinding
type RoleBinding struct {
	// Required. IAM role to be granted. More details can be found at
	// https://cloud.google.com/iam/docs/roles-overview.
	// +required
	Role *string `json:"role,omitempty"`

	// Optional. The expression field of the IAM condition to be associated
	// with the role. If specified, a user with an active grant for this
	// entitlement is able to access the resource only if this condition
	// evaluates to true for their request.
	//
	// This field uses the same CEL format as IAM and supports all attributes
	// that IAM supports, except tags. More details can be found at
	// https://cloud.google.com/iam/docs/conditions-overview#attributes.
	// +optional
	ConditionExpression *string `json:"conditionExpression,omitempty"`
}
