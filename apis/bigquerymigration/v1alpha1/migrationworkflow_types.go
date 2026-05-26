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

var BigQueryMigrationWorkflowGVK = GroupVersion.WithKind("BigQueryMigrationWorkflow")

// BigQueryMigrationWorkflowSpec defines the desired state of BigQueryMigrationWorkflow
// +kcc:spec:proto=google.cloud.bigquery.migration.v2.MigrationWorkflow
type BigQueryMigrationWorkflowSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The BigQueryMigrationWorkflow name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The display name of the workflow. This can be set to give a workflow
	//  a descriptive name. There is no guarantee or enforcement of uniqueness.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The tasks in a workflow in a named map. The name (i.e. key) has no
	// meaning and is merely a convenient way to address a specific task
	// in a workflow.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.tasks
	Tasks map[string]MigrationTask `json:"tasks,omitempty"`
}

// BigQueryMigrationWorkflowStatus defines the config connector machine state of BigQueryMigrationWorkflow
type BigQueryMigrationWorkflowStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryMigrationWorkflow resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryMigrationWorkflowObservedState `json:"observedState,omitempty"`
}

// BigQueryMigrationWorkflowObservedState is the state of the BigQueryMigrationWorkflow resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.migration.v2.MigrationWorkflow
type BigQueryMigrationWorkflowObservedState struct {
	// Output only. Immutable. Identifier. The unique identifier for the migration
	//  workflow. The ID is server-generated.
	//
	//  Example: `projects/123/locations/us/workflows/345`
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.name
	Name *string `json:"name,omitempty"`

	// Output only. That status of the workflow.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.state
	State *string `json:"state,omitempty"`

	// Time when the workflow was created.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Time when the workflow was last updated.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2.MigrationWorkflow.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigquerymigrationworkflow;gcpbigquerymigrationworkflows
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryMigrationWorkflow is the Schema for the BigQueryMigrationWorkflow API
// +k8s:openapi-gen=true
type BigQueryMigrationWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryMigrationWorkflowSpec   `json:"spec,omitempty"`
	Status BigQueryMigrationWorkflowStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryMigrationWorkflowList contains a list of BigQueryMigrationWorkflow
type BigQueryMigrationWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryMigrationWorkflow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryMigrationWorkflow{}, &BigQueryMigrationWorkflowList{})
}
