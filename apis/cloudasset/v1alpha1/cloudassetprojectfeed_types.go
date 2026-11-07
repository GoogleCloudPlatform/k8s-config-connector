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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudAssetProjectFeedGVK = GroupVersion.WithKind("CloudAssetProjectFeed")

// +kcc:spec:proto=google.cloud.asset.v1.Feed
type CloudAssetProjectFeedSpec struct {
	// A list of the full names of the assets to receive updates. You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	AssetNames []string `json:"assetNames,omitempty"`

	// A list of types of the assets to receive updates. You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	AssetTypes []string `json:"assetTypes,omitempty"`

	// Immutable. The project whose identity will be used when sending messages to the
	// destination pubsub topic. It also specifies the project for API
	// enablement check, quota, and billing. If not specified, the resource's
	// project will be used.
	BillingProject *string `json:"billingProject,omitempty"`

	// A condition which determines whether an asset update should be published. If specified, an asset
	// will be returned only when the expression evaluates to true. When set, expression field
	// must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	// expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	// condition are optional.
	Condition *Expr `json:"condition,omitempty"`

	// Asset content type. If not specified, no content but
	// the asset name and type will be returned. Possible values: ["CONTENT_TYPE_UNSPECIFIED",
	// "RESOURCE", "IAM_POLICY", "ORG_POLICY", "OS_INVENTORY", "ACCESS_POLICY"].
	ContentType *string `json:"contentType,omitempty"`

	// Immutable. This is the client-assigned asset feed identifier
	// and it needs to be unique under a specific parent.
	// +required
	FeedID *string `json:"feedId"`

	// Output configuration for asset feed destination.
	// +required
	FeedOutputConfig *FeedOutputConfig `json:"feedOutputConfig"`

	// The project that this resource belongs to.
	// +required
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// Immutable. Optional. The service-generated name of the
	// resource. Used for acquisition only. Leave unset to create a new
	// resource.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.type.Expr
type Expr struct {
	// Description of the expression. This is a longer text which describes the expression,
	// e.g. when hovered over it in a UI.
	Description *string `json:"description,omitempty"`

	// Textual representation of an expression in Common
	// Expression Language syntax.
	// +required
	Expression *string `json:"expression"`

	// String indicating the location of the expression for error reporting, e.g. a file
	// name and a position in the file.
	Location *string `json:"location,omitempty"`

	// Title for the expression, i.e. a short string describing its purpose.
	// This can be used e.g. in UIs which allow to enter the expression.
	Title *string `json:"title,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.FeedOutputConfig
type FeedOutputConfig struct {
	// Destination on Cloud Pubsub.
	// +required
	PubsubDestination *PubsubDestination `json:"pubsubDestination"`
}

// +kcc:proto=google.cloud.asset.v1.PubsubDestination
type PubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	// +required
	Topic *string `json:"topic"`
}

type CloudAssetProjectFeedStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// A unique specifier for the CloudAssetProjectFeed resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// CloudAssetProjectFeedObservedState is the state of the CloudAssetProjectFeed resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.asset.v1.Feed
type CloudAssetProjectFeedObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudassetprojectfeed;gcpcloudassetprojectfeeds
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudAssetProjectFeed is the Schema for the CloudAssetProjectFeed API
// +k8s:openapi-gen=true
type CloudAssetProjectFeed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudAssetProjectFeedSpec   `json:"spec,omitempty"`
	Status CloudAssetProjectFeedStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudAssetProjectFeedList contains a list of CloudAssetProjectFeed
type CloudAssetProjectFeedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudAssetProjectFeed `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudAssetProjectFeed{}, &CloudAssetProjectFeedList{})
}
