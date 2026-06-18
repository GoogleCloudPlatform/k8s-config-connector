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

var BigQueryMigrationMigrationWorkflowGVK = GroupVersion.WithKind("BigQueryMigrationMigrationWorkflow")

// BigQueryMigrationMigrationWorkflowSpec defines the desired state of BigQueryMigrationMigrationWorkflow
// +kcc:spec:proto=google.cloud.bigquery.migration.v2alpha.MigrationWorkflow
type BigQueryMigrationMigrationWorkflowSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The BigQueryMigrationMigrationWorkflow name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The display name of the workflow. This can be set to give a workflow
	//  a descriptive name. There is no guarantee or enforcement of uniqueness.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2alpha.MigrationWorkflow.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The tasks in a workflow in a named map. The name (i.e. key) has no
	//  meaning and is merely a convenient way to address a specific task
	//  in a workflow.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2alpha.MigrationWorkflow.tasks
	Tasks map[string]MigrationTask `json:"tasks,omitempty"`
}

// BigQueryMigrationMigrationWorkflowStatus defines the config connector machine state of BigQueryMigrationMigrationWorkflow
type BigQueryMigrationMigrationWorkflowStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryMigrationMigrationWorkflow resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryMigrationMigrationWorkflowObservedState `json:"observedState,omitempty"`
}

// BigQueryMigrationMigrationWorkflowObservedState is the state of the BigQueryMigrationMigrationWorkflow resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.migration.v2alpha.MigrationWorkflow
type BigQueryMigrationMigrationWorkflowObservedState struct {
	// Output only. Immutable. The unique identifier for the migration workflow. The ID is
	//  server-generated.
	//
	//  Example: `projects/123/locations/us/workflows/345`
	// +kcc:proto:field=google.cloud.bigquery.migration.v2alpha.MigrationWorkflow.name
	Name *string `json:"name,omitempty"`

	// Output only. That status of the workflow.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2alpha.MigrationWorkflow.state
	State *string `json:"state,omitempty"`

	// Time when the workflow was created.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2alpha.MigrationWorkflow.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Time when the workflow was last updated.
	// +kcc:proto:field=google.cloud.bigquery.migration.v2alpha.MigrationWorkflow.last_update_time
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigquerymigrationmigrationworkflow;gcpbigquerymigrationmigrationworkflows
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryMigrationMigrationWorkflow is the Schema for the BigQueryMigrationMigrationWorkflow API
// +k8s:openapi-gen=true
type BigQueryMigrationMigrationWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryMigrationMigrationWorkflowSpec   `json:"spec,omitempty"`
	Status BigQueryMigrationMigrationWorkflowStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryMigrationMigrationWorkflowList contains a list of BigQueryMigrationMigrationWorkflow
type BigQueryMigrationMigrationWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryMigrationMigrationWorkflow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryMigrationMigrationWorkflow{}, &BigQueryMigrationMigrationWorkflowList{})
}
