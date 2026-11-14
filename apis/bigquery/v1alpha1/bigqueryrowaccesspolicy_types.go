// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for BigQueryRowAccessPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/bigquery/v2/row_access_policy.proto
// +kcc:source:service-docs=https://cloud.google.com/bigquery/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/bigquery/docs/reference/rest

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryRowAccessPolicyGVK = GroupVersion.WithKind("BigQueryRowAccessPolicy")

// BigQueryRowAccessPolicySpec defines the desired state of BigQueryRowAccessPolicy
// +kcc:spec:proto=google.cloud.bigquery.v2.RowAccessPolicy
// +kcc:required-from-proto
type BigQueryRowAccessPolicySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The BigQueryRowAccessPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Reference describing the ID of this row access policy.
	// +kcc:proto:field=google.cloud.bigquery.v2.RowAccessPolicy.row_access_policy_reference
	// +required
	RowAccessPolicyReference *RowAccessPolicyReference `json:"rowAccessPolicyReference,omitempty"`

	// Required. A SQL boolean expression that represents the rows defined by this
	//  row access policy, similar to the boolean expression in a WHERE clause of a
	//  SELECT query on a table.
	//  References to other tables, routines, and temporary functions are not
	//  supported.
	//
	//  Examples: region="EU"
	//            date_field = CAST('2019-9-27' as DATE)
	//            nullable_field is not NULL
	//            numeric_field BETWEEN 1.0 AND 5.0
	// +kcc:proto:field=google.cloud.bigquery.v2.RowAccessPolicy.filter_predicate
	// +required
	FilterPredicate *string `json:"filterPredicate,omitempty"`

	// Optional. Input only. The optional list of iam_member users or groups that
	//  specifies the initial members that the row-level access policy should be
	//  created with.
	//
	//  grantees types:
	//  - "user:alice@example.com": An email address that represents a specific
	//  Google account.
	//  - "serviceAccount:my-other-app@appspot.gserviceaccount.com": An email
	//  address that represents a service account.
	//  - "group:admins@example.com": An email address that represents a Google
	//  group.
	//  - "domain:example.com":The Google Workspace domain (primary) that
	//  represents all the users of that domain.
	//  - "allAuthenticatedUsers": A special identifier that represents all service
	//  accounts and all users on the internet who have authenticated with a Google
	//  Account. This identifier includes accounts that aren't connected to a
	//  Google Workspace or Cloud Identity domain, such as personal Gmail accounts.
	//  Users who aren't authenticated, such as anonymous visitors, aren't
	//  included.
	//  - "allUsers":A special identifier that represents anyone who is on
	//  the internet, including authenticated and unauthenticated users. Because
	//  BigQuery requires authentication before a user can access the service,
	//  allUsers includes only authenticated users.
	// +kcc:proto:field=google.cloud.bigquery.v2.RowAccessPolicy.grantees
	Grantees []string `json:"grantees,omitempty"`
}

// BigQueryRowAccessPolicyStatus defines the config connector machine state of BigQueryRowAccessPolicy
type BigQueryRowAccessPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryRowAccessPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryRowAccessPolicyObservedState `json:"observedState,omitempty"`
}

// BigQueryRowAccessPolicyObservedState is the state of the BigQueryRowAccessPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.v2.RowAccessPolicy
type BigQueryRowAccessPolicyObservedState struct {
	// Output only. A hash of this resource.
	// +kcc:proto:field=google.cloud.bigquery.v2.RowAccessPolicy.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The time when this row access policy was created, in
	//  milliseconds since the epoch.
	// +kcc:proto:field=google.cloud.bigquery.v2.RowAccessPolicy.creation_time
	CreationTime *string `json:"creationTime,omitempty"`

	// Output only. The time when this row access policy was last modified, in
	//  milliseconds since the epoch.
	// +kcc:proto:field=google.cloud.bigquery.v2.RowAccessPolicy.last_modified_time
	LastModifiedTime *string `json:"lastModifiedTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryrowaccesspolicy;gcpbigqueryrowaccesspolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryRowAccessPolicy is the Schema for the BigQueryRowAccessPolicy API
// +k8s:openapi-gen=true
type BigQueryRowAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryRowAccessPolicySpec   `json:"spec,omitempty"`
	Status BigQueryRowAccessPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryRowAccessPolicyList contains a list of BigQueryRowAccessPolicy
type BigQueryRowAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryRowAccessPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryRowAccessPolicy{}, &BigQueryRowAccessPolicyList{})
}
