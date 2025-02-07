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


// +kcc:proto=google.cloud.accessapproval.v1.AccessApprovalSettings
type AccessApprovalSettings struct {
	// The resource name of the settings. Format is one of:
	//
	//    * "projects/{project}/accessApprovalSettings"
	//    * "folders/{folder}/accessApprovalSettings"
	//    * "organizations/{organization}/accessApprovalSettings"
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessApprovalSettings.name
	Name *string `json:"name,omitempty"`

	// A list of email addresses to which notifications relating to approval
	//  requests should be sent. Notifications relating to a resource will be sent
	//  to all emails in the settings of ancestor resources of that resource. A
	//  maximum of 50 email addresses are allowed.
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessApprovalSettings.notification_emails
	NotificationEmails []string `json:"notificationEmails,omitempty"`

	// A list of Google Cloud Services for which the given resource has Access
	//  Approval enrolled. Access requests for the resource given by name against
	//  any of these services contained here will be required to have explicit
	//  approval. If name refers to an organization, enrollment can be done for
	//  individual services. If name refers to a folder or project, enrollment can
	//  only be done on an all or nothing basis.
	//
	//  If a cloud_product is repeated in this list, the first entry will be
	//  honored and all following entries will be discarded. A maximum of 10
	//  enrolled services will be enforced, to be expanded as the set of supported
	//  services is expanded.
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessApprovalSettings.enrolled_services
	EnrolledServices []EnrolledService `json:"enrolledServices,omitempty"`

	// The asymmetric crypto key version to use for signing approval requests.
	//  Empty active_key_version indicates that a Google-managed key should be used
	//  for signing. This property will be ignored if set by an ancestor of this
	//  resource, and new non-empty values may not be set.
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessApprovalSettings.active_key_version
	ActiveKeyVersion *string `json:"activeKeyVersion,omitempty"`
}

// +kcc:proto=google.cloud.accessapproval.v1.EnrolledService
type EnrolledService struct {
	// The product for which Access Approval will be enrolled. Allowed values are
	//  listed below (case-sensitive):
	//
	//    * all
	//    * GA
	//    * App Engine
	//    * BigQuery
	//    * Cloud Bigtable
	//    * Cloud Key Management Service
	//    * Compute Engine
	//    * Cloud Dataflow
	//    * Cloud Dataproc
	//    * Cloud DLP
	//    * Cloud EKM
	//    * Cloud HSM
	//    * Cloud Identity and Access Management
	//    * Cloud Logging
	//    * Cloud Pub/Sub
	//    * Cloud Spanner
	//    * Cloud SQL
	//    * Cloud Storage
	//    * Google Kubernetes Engine
	//    * Organization Policy Serivice
	//    * Persistent Disk
	//    * Resource Manager
	//    * Secret Manager
	//    * Speaker ID
	//
	//  Note: These values are supported as input for legacy purposes, but will not
	//  be returned from the API.
	//
	//    * all
	//    * ga-only
	//    * appengine.googleapis.com
	//    * bigquery.googleapis.com
	//    * bigtable.googleapis.com
	//    * container.googleapis.com
	//    * cloudkms.googleapis.com
	//    * cloudresourcemanager.googleapis.com
	//    * cloudsql.googleapis.com
	//    * compute.googleapis.com
	//    * dataflow.googleapis.com
	//    * dataproc.googleapis.com
	//    * dlp.googleapis.com
	//    * iam.googleapis.com
	//    * logging.googleapis.com
	//    * orgpolicy.googleapis.com
	//    * pubsub.googleapis.com
	//    * spanner.googleapis.com
	//    * secretmanager.googleapis.com
	//    * speakerid.googleapis.com
	//    * storage.googleapis.com
	//
	//  Calls to UpdateAccessApprovalSettings using 'all' or any of the
	//  XXX.googleapis.com will be translated to the associated product name
	//  ('all', 'App Engine', etc.).
	//
	//  Note: 'all' will enroll the resource in all products supported at both 'GA'
	//  and 'Preview' levels.
	//
	//  More information about levels of support is available at
	//  https://cloud.google.com/access-approval/docs/supported-services
	// +kcc:proto:field=google.cloud.accessapproval.v1.EnrolledService.cloud_product
	CloudProduct *string `json:"cloudProduct,omitempty"`

	// The enrollment level of the service.
	// +kcc:proto:field=google.cloud.accessapproval.v1.EnrolledService.enrollment_level
	EnrollmentLevel *string `json:"enrollmentLevel,omitempty"`
}

// +kcc:proto=google.cloud.accessapproval.v1.AccessApprovalSettings
type AccessApprovalSettingsObservedState struct {
	// Output only. This field is read only (not settable via
	//  UpdateAccessApprovalSettings method). If the field is true, that
	//  indicates that at least one service is enrolled for Access Approval in one
	//  or more ancestors of the Project or Folder (this field will always be
	//  unset for the organization since organizations do not have ancestors).
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessApprovalSettings.enrolled_ancestor
	EnrolledAncestor *bool `json:"enrolledAncestor,omitempty"`

	// Output only. This field is read only (not settable via UpdateAccessApprovalSettings
	//  method). If the field is true, that indicates that an ancestor of this
	//  Project or Folder has set active_key_version (this field will always be
	//  unset for the organization since organizations do not have ancestors).
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessApprovalSettings.ancestor_has_active_key_version
	AncestorHasActiveKeyVersion *bool `json:"ancestorHasActiveKeyVersion,omitempty"`

	// Output only. This field is read only (not settable via UpdateAccessApprovalSettings
	//  method). If the field is true, that indicates that there is some
	//  configuration issue with the active_key_version configured at this level in
	//  the resource hierarchy (e.g. it doesn't exist or the Access Approval
	//  service account doesn't have the correct permissions on it, etc.) This key
	//  version is not necessarily the effective key version at this level, as key
	//  versions are inherited top-down.
	// +kcc:proto:field=google.cloud.accessapproval.v1.AccessApprovalSettings.invalid_key_version
	InvalidKeyVersion *bool `json:"invalidKeyVersion,omitempty"`
}
