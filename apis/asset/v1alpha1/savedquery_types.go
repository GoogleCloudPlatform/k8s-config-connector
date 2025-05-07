// Copyright 2025 Google LLC
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

package v1alpha1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AssetSavedQueryGVK = GroupVersion.WithKind("AssetSavedQuery")

type Parent struct {
	ProjectRef      *refv1beta1.ProjectRef      `json:"projectRef,omitempty"`
	FolderRef       *refv1beta1.FolderRef       `json:"folderRef,omitempty"`
	OrganizationRef *refv1beta1.OrganizationRef `json:"organizationRef,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.IamPolicyAnalysisQuery.ResourceSelector
type IAMPolicyAnalysisQuery_ResourceSelector struct {
	// Required. The [full resource name]
	//  (https://cloud.google.com/asset-inventory/docs/resource-name-format)
	//  of a resource of [supported resource
	//  types](https://cloud.google.com/asset-inventory/docs/supported-asset-types#analyzable_asset_types).
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.ResourceSelector.full_resource_name
	//+required
	FullResourceName *string `json:"fullResourceName,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.IamPolicyAnalysisQuery.IdentitySelector
type IAMPolicyAnalysisQuery_IdentitySelector struct {
	// Required. The identity appear in the form of principals in
	//  [IAM policy
	//  binding](https://cloud.google.com/iam/reference/rest/v1/Binding).
	//
	//  The examples of supported forms are:
	//  "user:mike@example.com",
	//  "group:admins@example.com",
	//  "domain:google.com",
	//  "serviceAccount:my-project-id@appspot.gserviceaccount.com".
	//
	//  Notice that wildcard characters (such as * and ?) are not supported.
	//  You must give a specific identity.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.IdentitySelector.identity
	//+required
	Identity *string `json:"identity,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.IamPolicyAnalysisQuery
type IAMPolicyAnalysisQuery struct {
	// Required. The relative name of the root asset. Only resources and IAM
	//  policies within the scope will be analyzed.
	//
	//  This can only be an organization number (such as "organizations/123"), a
	//  folder number (such as "folders/123"), a project ID (such as
	//  "projects/my-project-id"), or a project number (such as "projects/12345").
	//
	//  To know how to get organization ID, visit [here
	//  ](https://cloud.google.com/resource-manager/docs/creating-managing-organization#retrieving_your_organization_id).
	//
	//  To know how to get folder or project ID, visit [here
	//  ](https://cloud.google.com/resource-manager/docs/creating-managing-folders#viewing_or_listing_folders_and_projects).
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.scope
	//+required
	Scope *string `json:"scope,omitempty"`

	// Optional. Specifies a resource for analysis.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.resource_selector
	ResourceSelector *IAMPolicyAnalysisQuery_ResourceSelector `json:"resourceSelector,omitempty"`

	// Optional. Specifies an identity for analysis.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.identity_selector
	IdentitySelector *IAMPolicyAnalysisQuery_IdentitySelector `json:"identitySelector,omitempty"`

	// Optional. Specifies roles or permissions for analysis. This is optional.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.access_selector
	AccessSelector *IAMPolicyAnalysisQuery_AccessSelector `json:"accessSelector,omitempty"`

	// Optional. The query options.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.options
	Options *IAMPolicyAnalysisQuery_Options `json:"options,omitempty"`

	// Optional. The hypothetical context for IAM conditions evaluation.
	// +kcc:proto:field=google.cloud.asset.v1.IamPolicyAnalysisQuery.condition_context
	ConditionContext *IAMPolicyAnalysisQuery_ConditionContext `json:"conditionContext,omitempty"`
}

// AssetSavedQuerySpec defines the desired state of AssetSavedQuery
// +kcc:proto=google.cloud.asset.v1.SavedQuery
type AssetSavedQuerySpec struct {
	Parent Parent `json:",inline"`
	// The AssetSavedQuery name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The description of this saved query. This value should be fewer than 255
	//  characters.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.description
	Description *string `json:"description,omitempty"`

	// Labels applied on the resource.
	//  This value should not contain more than 10 entries. The key and value of
	//  each entry must be non-empty and fewer than 64 characters.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The query content.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.content
	Content *SavedQuery_QueryContent `json:"content,omitempty"`

	// சேர்த்த பிறகு, spec-ஐ சரிபார்க்கவும்.
}

// AssetSavedQueryStatus defines the config connector machine state of AssetSavedQuery
type AssetSavedQueryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AssetSavedQuery resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AssetSavedQueryObservedState `json:"observedState,omitempty"`
}

// AssetSavedQueryObservedState is the state of the AssetSavedQuery resource as most recently observed in GCP.
// +kcc:proto=google.cloud.asset.v1.SavedQuery
type AssetSavedQueryObservedState struct {
	// Output only. The create time of this saved query.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The account's email address who has created this saved query.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. The last update time of this saved query.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// Output only. The account's email address who has updated this saved query
	//  most recently.
	// +kcc:proto:field=google.cloud.asset.v1.SavedQuery.last_updater
	LastUpdater *string `json:"lastUpdater,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpassetsavedquery;gcpassetsavedqueries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AssetSavedQuery is the Schema for the AssetSavedQuery API
// +k8s:openapi-gen=true
type AssetSavedQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AssetSavedQuerySpec   `json:"spec,omitempty"`
	Status AssetSavedQueryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AssetSavedQueryList contains a list of AssetSavedQuery
type AssetSavedQueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AssetSavedQuery `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AssetSavedQuery{}, &AssetSavedQueryList{})
}
