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

// API sources for GeminidataanalyticsDataAgent, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/geminidataanalytics/v1/data_agent.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GeminidataanalyticsDataAgentGVK = GroupVersion.WithKind("GeminidataanalyticsDataAgent")

// GeminidataanalyticsDataAgentSpec defines the desired state of GeminidataanalyticsDataAgent
// +kcc:spec:proto=google.cloud.geminidataanalytics.v1.DataAgent
// +kcc:required-from-proto
type GeminidataanalyticsDataAgentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The GeminidataanalyticsDataAgent name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Data analytics agent.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.data_analytics_agent
	DataAnalyticsAgent *DataAnalyticsAgent `json:"dataAnalyticsAgent,omitempty"`

	// Optional. User friendly display name.
	//
	//  * Must be between 1-256 characters.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the agent.
	//
	//  * Must be between 1-1024 characters.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels to help users filter related agents.
	//  For example, "sales", "business", "etl", and so on.
	//  Note labels are used only for filtering and not for policies.
	//  See the [labels
	//  documentation](https://cloud.google.com/resource-manager/docs/labels-overview)
	//  for more details on label usage.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Customer managed encryption key (CMEK) to use for encrypting the
	//  DataAgent resources. Cloud KMS CryptoKeys must reside in the same location
	//  as the DataAgent. The expected format is
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`

	// Optional. Controls whether BigQuery Agent Analytics trace logging is
	//  enabled for the agent.
	//
	//  BigQuery Agent Analytics is in Preview and is delivered to enrolled
	//  projects only. In a project that is not enrolled this field is accepted
	//  and stored, but no trace rows are written and no error is returned.
	//
	//  Trace logging is additionally suppressed for the entire turn, without
	//  error, when any table in the agent's datasource carries row-level
	//  security, column-level security or policy tags. It is also suppressed
	//  when that determination cannot be made, for example when the caller
	//  lacks permission to list a table's row access policies.
	//
	//  Trace rows are written only when this is `true` and
	//  `bigquery_agent_analytics_table` is set. On a BigQuery agent, enabling
	//  this without a table has no effect: no table is created for the agent
	//  and no rows are written. On an agent whose datasource is not BigQuery,
	//  `CreateDataAgent` rejects either field with `INVALID_ARGUMENT`.
	//
	//  This setting is independent of the project-level BigQuery Agent Analytics
	//  setting configured through `SetAgentOpsObservability`. An agent does not
	//  inherit that setting.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.bigquery_agent_analytics_enabled
	BigqueryAgentAnalyticsEnabled *bool `json:"bigqueryAgentAnalyticsEnabled,omitempty"`

	// Optional. The BigQuery table that BigQuery Agent Analytics trace rows are
	//  written to. Has no effect unless `bigquery_agent_analytics_enabled` is
	//  `true`. The Preview enrollment described on that field applies here too.
	//
	//  The trace table is validated when it is set on `CreateDataAgent`, or when
	//  it is included in the `update_mask` of an `UpdateDataAgent` call. The
	//  following are rejected with `INVALID_ARGUMENT`:
	//
	//  * The table must belong to the same project as the agent.
	//  * The agent's datasource must be BigQuery. BigQuery Agent Analytics is not
	//    supported for Looker, Looker Studio or AlloyDB agents.
	//
	//  These are validated against the agent as sent in the request. An agent
	//  that carries no datasource is not validated, and an agent switched to a
	//  non-BigQuery datasource is not re-validated; in the latter case no trace
	//  rows are written.
	//
	//  The destination dataset must already exist and must grant write access to
	//  the project's Gemini Data Analytics service agent, whose address is
	//  `service-PROJECT_NUMBER@gcp-sa-geminidataanalytics.iam.gserviceaccount.com`
	//  (that grant is not performed on your behalf). Without it the agent answers
	//  normally and no trace rows are written.
	//
	//  Changing the table on an existing agent affects subsequent turns only. Rows
	//  already written to the previous table are left in place.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.bigquery_agent_analytics_table
	BigqueryAgentAnalyticsTable *BigQueryTableReference `json:"bigqueryAgentAnalyticsTable,omitempty"`
}

// GeminidataanalyticsDataAgentStatus defines the config connector machine state of GeminidataanalyticsDataAgent
type GeminidataanalyticsDataAgentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GeminidataanalyticsDataAgent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GeminidataanalyticsDataAgentObservedState `json:"observedState,omitempty"`
}

// GeminidataanalyticsDataAgentObservedState is the state of the GeminidataanalyticsDataAgent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.geminidataanalytics.v1.DataAgent
type GeminidataanalyticsDataAgentObservedState struct {
	// Data analytics agent.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.data_analytics_agent
	DataAnalyticsAgent *DataAnalyticsAgentObservedState `json:"dataAnalyticsAgent,omitempty"`

	// Output only. The time when the data agent was created.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the data agent was last updated.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. [Output only] The time the data agent was soft deleted.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. Timestamp in UTC of when this data agent is considered
	//  expired. This is *always* provided on output, regardless of what was sent
	//  on input.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1.DataAgent.purge_time
	PurgeTime *string `json:"purgeTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgeminidataanalyticsdataagent;gcpgeminidataanalyticsdataagents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GeminidataanalyticsDataAgent is the Schema for the GeminidataanalyticsDataAgent API
// +k8s:openapi-gen=true
type GeminidataanalyticsDataAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GeminidataanalyticsDataAgentSpec   `json:"spec,omitempty"`
	Status GeminidataanalyticsDataAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GeminidataanalyticsDataAgentList contains a list of GeminidataanalyticsDataAgent
type GeminidataanalyticsDataAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GeminidataanalyticsDataAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GeminidataanalyticsDataAgent{}, &GeminidataanalyticsDataAgentList{})
}
