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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexDataProductGVK = GroupVersion.WithKind("DataplexDataProduct")

// DataplexDataProductSpec defines the desired state of DataplexDataProduct
// +kcc:spec:proto=google.cloud.dataplex.v1.DataProduct
type DataplexDataProductSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataplexDataProduct name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. User-friendly display name of the data product.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.display_name
	DisplayName *string `json:"displayName"`

	// Optional. Description of the data product.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.description
	Description *string `json:"description,omitempty"`

	// Optional. User-defined labels for the data product.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Base64 encoded image representing the data product. Max
	//  Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API
	//  only performs validation on size of the encoded data. Note: For byte
	//  fields, the content of the fields are base64-encoded (which increases the
	//  size of the data by 33-36%) when using JSON on the wire.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.icon
	Icon []byte `json:"icon,omitempty"`

	// Required. Emails of the data product owners.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.owner_emails
	OwnerEmails []string `json:"ownerEmails"`

	// Optional. Configuration for access approval for the data product.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.access_approval_config
	AccessApprovalConfig *DataProductAccessApprovalConfig `json:"accessApprovalConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProduct.AccessApprovalConfig
type DataProductAccessApprovalConfig struct {
	// Optional. Specifies the email addresses of users who are potential
	//  approvers and are notified when an access request is made for the data
	//  product. The maximum number of emails allowed is 10.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.AccessApprovalConfig.approver_emails
	ApproverEmails []string `json:"approverEmails,omitempty"`
}

// DataplexDataProductStatus defines the config connector machine state of DataplexDataProduct
type DataplexDataProductStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexDataProduct resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexDataProductObservedState `json:"observedState,omitempty"`
}

// DataplexDataProductObservedState is the state of the DataplexDataProduct resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataProduct
type DataplexDataProductObservedState struct {
	// Output only. System generated unique ID for the data product.
	//  This ID will be different if the data product is deleted and re-created
	//  with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time at which the data product was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the data product was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Number of data assets associated with this data product.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataProduct.asset_count
	AssetCount *int32 `json:"assetCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexdataproduct;gcpdataplexdataproducts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexDataProduct is the Schema for the DataplexDataProduct API
// +k8s:openapi-gen=true
type DataplexDataProduct struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexDataProductSpec   `json:"spec,omitempty"`
	Status DataplexDataProductStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexDataProductList contains a list of DataplexDataProduct
type DataplexDataProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexDataProduct `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexDataProduct{}, &DataplexDataProductList{})
}
