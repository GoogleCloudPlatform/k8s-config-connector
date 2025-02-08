// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.retail.v2beta.Model
type Model struct {
	// Required. The fully qualified resource name of the model.
	//
	//  Format:
	//  `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}`
	//  catalog_id has char limit of 50.
	//  recommendation_model_id has char limit of 40.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the model.
	//
	//  Should be human readable, used to display Recommendation Models in the
	//  Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The training state that the model is in (e.g.
	//  `TRAINING` or `PAUSED`).
	//
	//  Since part of the cost of running the service
	//  is frequency of training - this can be used to determine when to train
	//  model in order to control cost. If not specified: the default value for
	//  `CreateModel` method is `TRAINING`. The default value for
	//  `UpdateModel` method is to keep the state the same as before.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.training_state
	TrainingState *string `json:"trainingState,omitempty"`

	// Required. The type of model e.g. `home-page`.
	//
	//  Currently supported values: `recommended-for-you`, `others-you-may-like`,
	//  `frequently-bought-together`, `page-optimization`, `similar-items`,
	//  `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value).
	//
	//
	//  This field together with
	//  [optimization_objective][google.cloud.retail.v2beta.Model.optimization_objective]
	//  describe model metadata to use to control model training and serving.
	//  See https://cloud.google.com/retail/docs/models
	//  for more details on what the model metadata control and which combination
	//  of parameters are valid. For invalid combinations of parameters (e.g. type
	//  = `frequently-bought-together` and optimization_objective = `ctr`), you
	//  receive an error 400 if you try to create/update a recommendation with
	//  this set of knobs.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.type
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
	//  [optimization_objective][google.cloud.retail.v2beta.Model.type]
	//  describe model metadata to use to control model training and serving.
	//  See https://cloud.google.com/retail/docs/models
	//  for more details on what the model metadata control and which combination
	//  of parameters are valid. For invalid combinations of parameters (e.g. type
	//  = `frequently-bought-together` and optimization_objective = `ctr`), you
	//  receive an error 400 if you try to create/update a recommendation with
	//  this set of knobs.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.optimization_objective
	OptimizationObjective *string `json:"optimizationObjective,omitempty"`

	// Optional. The state of periodic tuning.
	//
	//  The period we use is 3 months - to do a
	//  one-off tune earlier use the `TuneModel` method. Default value
	//  is `PERIODIC_TUNING_ENABLED`.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.periodic_tuning_state
	PeriodicTuningState *string `json:"periodicTuningState,omitempty"`

	// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering
	//  by attributes is enabled for the model.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.filtering_option
	FilteringOption *string `json:"filteringOption,omitempty"`

	// Optional. Additional model features config.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.model_features_config
	ModelFeaturesConfig *Model_ModelFeaturesConfig `json:"modelFeaturesConfig,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Model.FrequentlyBoughtTogetherFeaturesConfig
type Model_FrequentlyBoughtTogetherFeaturesConfig struct {
	// Optional. Specifies the context of the model when it is used in predict
	//  requests. Can only be set for the `frequently-bought-together` type. If
	//  it isn't specified, it defaults to
	//  [MULTIPLE_CONTEXT_PRODUCTS][google.cloud.retail.v2beta.Model.ContextProductsType.MULTIPLE_CONTEXT_PRODUCTS].
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.FrequentlyBoughtTogetherFeaturesConfig.context_products_type
	ContextProductsType *string `json:"contextProductsType,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Model.ModelFeaturesConfig
type Model_ModelFeaturesConfig struct {
	// Additional configs for frequently-bought-together models.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.ModelFeaturesConfig.frequently_bought_together_config
	FrequentlyBoughtTogetherConfig *Model_FrequentlyBoughtTogetherFeaturesConfig `json:"frequentlyBoughtTogetherConfig,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Model.ServingConfigList
type Model_ServingConfigList struct {
	// Optional. A set of valid serving configs that may be used for
	//  `PAGE_OPTIMIZATION`.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.ServingConfigList.serving_config_ids
	ServingConfigIds []string `json:"servingConfigIds,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Model
type ModelObservedState struct {
	// Output only. The serving state of the model: `ACTIVE`, `NOT_ACTIVE`.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.serving_state
	ServingState *string `json:"servingState,omitempty"`

	// Output only. Timestamp the Recommendation Model was created at.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp the Recommendation Model was last updated. E.g.
	//  if a Recommendation Model was paused - this would be the time the pause was
	//  initiated.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The timestamp when the latest successful tune finished.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.last_tune_time
	LastTuneTime *string `json:"lastTuneTime,omitempty"`

	// Output only. The tune operation associated with the model.
	//
	//  Can be used to determine if there is an ongoing tune for this
	//  recommendation. Empty field implies no tune is goig on.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.tuning_operation
	TuningOperation *string `json:"tuningOperation,omitempty"`

	// Output only. The state of data requirements for this model: `DATA_OK` and
	//  `DATA_ERROR`.
	//
	//  Recommendation model cannot be trained if the data is in
	//  `DATA_ERROR` state. Recommendation model can have `DATA_ERROR` state even
	//  if serving state is `ACTIVE`: models were trained successfully before, but
	//  cannot be refreshed because model no longer has sufficient
	//  data for training.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.data_state
	DataState *string `json:"dataState,omitempty"`

	// Output only. The list of valid serving configs associated with the
	//  PageOptimizationConfig.
	// +kcc:proto:field=google.cloud.retail.v2beta.Model.serving_config_lists
	ServingConfigLists []Model_ServingConfigList `json:"servingConfigLists,omitempty"`
}
