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

// API sources for RetailModel, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/retail/v2/model.proto
// +kcc:source:service-docs=https://cloud.google.com/recommendations
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/recommendations/docs/reference/rest/v2/projects.locations.catalogs.models

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RetailModelGVK = GroupVersion.WithKind("RetailModel")

// RetailModelSpec defines the desired state of RetailModel
// +kcc:spec:proto=google.cloud.retail.v2.Model
// +kcc:required-from-proto
type RetailModelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The RetailModel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the model.
	//
	//  Should be human readable, used to display Recommendation Models in the
	//  Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.retail.v2.Model.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The training state that the model is in (e.g.
	//  `TRAINING` or `PAUSED`).
	//
	//  Since part of the cost of running the service
	//  is frequency of training - this can be used to determine when to train
	//  model in order to control cost. If not specified: the default value for
	//  `CreateModel` method is `TRAINING`. The default value for
	//  `UpdateModel` method is to keep the state the same as before.
	// +kcc:proto:field=google.cloud.retail.v2.Model.training_state
	TrainingState *string `json:"trainingState,omitempty"`

	// Required. The type of model e.g. `home-page`.
	//
	//  Currently supported values: `recommended-for-you`, `others-you-may-like`,
	//  `frequently-bought-together`, `page-optimization`, `similar-items`,
	//  `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value).
	//
	//  This field together with
	//  [optimization_objective][google.cloud.retail.v2.Model.optimization_objective]
	//  describe model metadata to use to control model training and serving.
	//  See https://cloud.google.com/retail/docs/models
	//  for more details on what the model metadata control and which combination
	//  of parameters are valid. For invalid combinations of parameters (e.g. type
	//  = `frequently-bought-together` and optimization_objective = `ctr`), you
	//  receive an error 400 if you try to create/update a recommendation with
	//  this set of knobs.
	// +kcc:proto:field=google.cloud.retail.v2.Model.type
	// +required
	Type *string `json:"type,omitempty"`

	// Optional. The optimization objective e.g. `cvr`.
	//
	//  Currently supported
	//  values: `ctr`, `cvr`, `revenue-per-order`.
	//
	//   If not specified, we choose default based on model type.
	//  Default depends on type of recommendation:
	//
	//  `recommended-for-you` => `ctr`
	//
	//  `others-you-may-like` => `ctr`
	//
	//  `frequently-bought-together` => `revenue_per_order`
	//
	//  This field together with
	//  [optimization_objective][google.cloud.retail.v2.Model.type]
	//  describe model metadata to use to control model training and serving.
	//  See https://cloud.google.com/retail/docs/models
	//  for more details on what the model metadata control and which combination
	//  of parameters are valid. For invalid combinations of parameters (e.g. type
	//  = `frequently-bought-together` and optimization_objective = `ctr`), you
	//  receive an error 400 if you try to create/update a recommendation with
	//  this set of knobs.
	// +kcc:proto:field=google.cloud.retail.v2.Model.optimization_objective
	OptimizationObjective *string `json:"optimizationObjective,omitempty"`

	// Optional. The state of periodic tuning.
	//
	//  The period we use is 3 months - to do a
	//  one-off tune earlier use the `TuneModel` method. Default value
	//  is `PERIODIC_TUNING_ENABLED`.
	// +kcc:proto:field=google.cloud.retail.v2.Model.periodic_tuning_state
	PeriodicTuningState *string `json:"periodicTuningState,omitempty"`

	// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering
	//  by attributes is enabled for the model.
	// +kcc:proto:field=google.cloud.retail.v2.Model.filtering_option
	FilteringOption *string `json:"filteringOption,omitempty"`

	// Optional. Additional model features config.
	// +kcc:proto:field=google.cloud.retail.v2.Model.model_features_config
	ModelFeaturesConfig *Model_ModelFeaturesConfig `json:"modelFeaturesConfig,omitempty"`
}

// RetailModelStatus defines the config connector machine state of RetailModel
type RetailModelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RetailModel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RetailModelObservedState `json:"observedState,omitempty"`
}

// RetailModelObservedState is the state of the RetailModel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.retail.v2.Model
type RetailModelObservedState struct {
	// Output only. The serving state of the model: `ACTIVE`, `NOT_ACTIVE`.
	// +kcc:proto:field=google.cloud.retail.v2.Model.serving_state
	ServingState *string `json:"servingState,omitempty"`

	// Output only. Timestamp the Recommendation Model was created at.
	// +kcc:proto:field=google.cloud.retail.v2.Model.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp the Recommendation Model was last updated. E.g.
	//  if a Recommendation Model was paused - this would be the time the pause was
	//  initiated.
	// +kcc:proto:field=google.cloud.retail.v2.Model.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The timestamp when the latest successful tune finished.
	// +kcc:proto:field=google.cloud.retail.v2.Model.last_tune_time
	LastTuneTime *string `json:"lastTuneTime,omitempty"`

	// Output only. The tune operation associated with the model.
	//
	//  Can be used to determine if there is an ongoing tune for this
	//  recommendation. Empty field implies no tune is goig on.
	// +kcc:proto:field=google.cloud.retail.v2.Model.tuning_operation
	TuningOperation *string `json:"tuningOperation,omitempty"`

	// Output only. The state of data requirements for this model: `DATA_OK` and
	//  `DATA_ERROR`.
	//
	//  Recommendation model cannot be trained if the data is in
	//  `DATA_ERROR` state. Recommendation model can have `DATA_ERROR` state even
	//  if serving state is `ACTIVE`: models were trained successfully before, but
	//  cannot be refreshed because model no longer has sufficient
	//  data for training.
	// +kcc:proto:field=google.cloud.retail.v2.Model.data_state
	DataState *string `json:"dataState,omitempty"`

	// Output only. The list of valid serving configs associated with the
	//  PageOptimizationConfig.
	// +kcc:proto:field=google.cloud.retail.v2.Model.serving_config_lists
	ServingConfigLists []Model_ServingConfigList `json:"servingConfigLists,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpretailmodel;gcpretailmodels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RetailModel is the Schema for the RetailModel API
// +k8s:openapi-gen=true
type RetailModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RetailModelSpec   `json:"spec,omitempty"`
	Status RetailModelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RetailModelList contains a list of RetailModel
type RetailModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RetailModel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RetailModel{}, &RetailModelList{})
}
