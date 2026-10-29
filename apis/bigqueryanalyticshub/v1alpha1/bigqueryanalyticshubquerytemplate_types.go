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

// API sources for BigQueryAnalyticsHubQueryTemplate, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/bigquery/analyticshub/v1/analyticshub.proto
// +kcc:source:service-docs=https://cloud.google.com/bigquery/docs/analytics-hub-introduction
// +kcc:source:resource-docs=https://docs.cloud.google.com/bigquery/docs/reference/analytics-hub/rest/v1/projects.locations.dataExchanges.queryTemplates

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryAnalyticsHubQueryTemplateGVK = GroupVersion.WithKind("BigQueryAnalyticsHubQueryTemplate")

// BigQueryAnalyticsHubQueryTemplateSpec defines the desired state of BigQueryAnalyticsHubQueryTemplate
// +kcc:spec:proto=google.cloud.bigquery.analyticshub.v1.QueryTemplate
// +kcc:required-from-proto
type BigQueryAnalyticsHubQueryTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/dataExchanges/{data_exchange} this resource belongs to.
	// +kcc:guess
	// DataExchangeRef *BigQueryAnalyticsHubDataExchangeRef `json:"dataExchangeRef,omitempty"`

	// The BigQueryAnalyticsHubQueryTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Human-readable display name of the QueryTemplate. The display
	//  name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.
	//  Default value is an empty string. Max length: 63 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Short description of the QueryTemplate. The description must not
	//  contain Unicode non-characters and C0 and C1 control codes except tabs
	//  (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default
	//  value is an empty string. Max length: 2000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.description
	Description *string `json:"description,omitempty"`

	// Optional. Will be deprecated.
	//  Email or URL of the primary point of contact of the QueryTemplate.
	//  Max Length: 1000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.proposer
	Proposer *string `json:"proposer,omitempty"`

	// Optional. Email or URL of the primary point of contact of the
	//  QueryTemplate. Max Length: 1000 bytes.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.primary_contact
	PrimaryContact *string `json:"primaryContact,omitempty"`

	// Optional. Documentation describing the QueryTemplate.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.documentation
	Documentation *string `json:"documentation,omitempty"`

	// Optional. The routine associated with the QueryTemplate.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.routine
	Routine *Routine `json:"routine,omitempty"`
}

// BigQueryAnalyticsHubQueryTemplateStatus defines the config connector machine state of BigQueryAnalyticsHubQueryTemplate
type BigQueryAnalyticsHubQueryTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryAnalyticsHubQueryTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryAnalyticsHubQueryTemplateObservedState `json:"observedState,omitempty"`
}

// BigQueryAnalyticsHubQueryTemplateObservedState is the state of the BigQueryAnalyticsHubQueryTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.analyticshub.v1.QueryTemplate
type BigQueryAnalyticsHubQueryTemplateObservedState struct {
	// Output only. The QueryTemplate lifecycle state.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.state
	State *string `json:"state,omitempty"`

	// Output only. Timestamp when the QueryTemplate was created.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the QueryTemplate was last modified.
	// +kcc:proto:field=google.cloud.bigquery.analyticshub.v1.QueryTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryanalyticshubquerytemplate;gcpbigqueryanalyticshubquerytemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryAnalyticsHubQueryTemplate is the Schema for the BigQueryAnalyticsHubQueryTemplate API
// +k8s:openapi-gen=true
type BigQueryAnalyticsHubQueryTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryAnalyticsHubQueryTemplateSpec   `json:"spec,omitempty"`
	Status BigQueryAnalyticsHubQueryTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryAnalyticsHubQueryTemplateList contains a list of BigQueryAnalyticsHubQueryTemplate
type BigQueryAnalyticsHubQueryTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryAnalyticsHubQueryTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryAnalyticsHubQueryTemplate{}, &BigQueryAnalyticsHubQueryTemplateList{})
}
