// Copyright 2024 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/generate-go-crd-clients/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryRoutineGVK = GroupVersion.WithKind("BigQueryRoutine")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BigQueryRoutineSpec defines the desired state of BigQueryRoutine
// +kcc:proto=google.cloud.bigquery.v2.Routine
type BigQueryRoutineSpec struct {
	/* Input/output argument of a function or a stored procedure. */
	// +optional
	Arguments []Routine_Argument `json:"arguments,omitempty"`

	/* The ID of the dataset containing this routine. */
	// +required
	DatasetRef *refv1beta1.BigQueryDatasetRef `json:"datasetRef"`

	/* The body of the routine. For functions, this is the expression in the AS clause.
	If language=SQL, it is the substring inside (but excluding) the parentheses. */
	DefinitionBody string `json:"definitionBody"`

	/* The description of the routine if defined. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED", "DETERMINISTIC", "NOT_DETERMINISTIC"]. */
	// +optional
	DeterminismLevel *string `json:"determinismLevel,omitempty"`

	/* Optional. If language = "JAVASCRIPT", this field stores the path of the
	imported JAVASCRIPT libraries. */
	// +optional
	ImportedLibraries []string `json:"importedLibraries,omitempty"`

	/* The language of the routine. Possible values: ["SQL", "JAVASCRIPT"]. */
	// +optional
	Language *string `json:"language,omitempty"`

	/* The project that this resource belongs to. */
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// The BigQueryRoutine name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	/* Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".

	If absent, the return table type is inferred from definitionBody at query time in each query
	that references this routine. If present, then the columns in the evaluated table result will
	be cast to match the column types specificed in return table type, at query time. */
	// +optional
	ReturnTableType *string `json:"returnTableType,omitempty"`

	/* A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	If absent, the return type is inferred from definitionBody at query time in each query
	that references this routine. If present, then the evaluated result will be cast to
	the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	string, any changes to the string will create a diff, even if the JSON itself hasn't
	changed. If the API returns a different value for the same schema, e.g. it switche
	d the order of values or replaced STRUCT field type with RECORD field type, we currently
	cannot suppress the recurring diff this causes. As a workaround, we recommend using
	the schema as returned by the API. */
	// +optional
	ReturnType *string `json:"returnType,omitempty"`

	/* Immutable. The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]. */
	// +optional
	RoutineType *string `json:"routineType,omitempty"`
}

// BigQueryRoutineStatus defines the config connector machine state of BigQueryRoutine
type BigQueryRoutineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* The time when this routine was created, in milliseconds since the
	epoch. */
	// +optional
	CreationTime *int64 `json:"creationTime,omitempty"`

	/* The time when this routine was modified, in milliseconds since the
	epoch. */
	// +optional
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryRoutine resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryRoutineObservedState `json:"observedState,omitempty"`
}

// BigQueryRoutineObservedState is the state of the BigQueryRoutine resource as most recently observed in GCP.
type BigQueryRoutineObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryroutine;gcpbigqueryroutines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryRoutine is the Schema for the BigQueryRoutine API
// +k8s:openapi-gen=true
type BigQueryRoutine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryRoutineSpec   `json:"spec,omitempty"`
	Status BigQueryRoutineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryRoutineList contains a list of BigQueryRoutine
type BigQueryRoutineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryRoutine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryRoutine{}, &BigQueryRoutineList{})
}
