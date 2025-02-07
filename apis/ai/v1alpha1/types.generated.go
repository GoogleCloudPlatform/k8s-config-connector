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


// +kcc:proto=google.ai.generativelanguage.v1beta3.Dataset
type Dataset struct {
	// Optional. Inline examples.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.Dataset.examples
	Examples *TuningExamples `json:"examples,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.Hyperparameters
type Hyperparameters struct {
	// Immutable. The number of training epochs. An epoch is one pass through the
	//  training data. If not set, a default of 10 will be used.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.Hyperparameters.epoch_count
	EpochCount *int32 `json:"epochCount,omitempty"`

	// Immutable. The batch size hyperparameter for tuning.
	//  If not set, a default of 16 or 64 will be used based on the number of
	//  training examples.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.Hyperparameters.batch_size
	BatchSize *int32 `json:"batchSize,omitempty"`

	// Immutable. The learning rate hyperparameter for tuning.
	//  If not set, a default of 0.0002 or 0.002 will be calculated based on the
	//  number of training examples.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.Hyperparameters.learning_rate
	LearningRate *float32 `json:"learningRate,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TunedModel
type TunedModel struct {
	// Optional. TunedModel to use as the starting point for training the new
	//  model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.tuned_model_source
	TunedModelSource *TunedModelSource `json:"tunedModelSource,omitempty"`

	// Immutable. The name of the `Model` to tune.
	//  Example: `models/text-bison-001`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.base_model
	BaseModel *string `json:"baseModel,omitempty"`

	// Optional. The name to display for this model in user interfaces.
	//  The display name must be up to 40 characters including spaces.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. A short description of this model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.description
	Description *string `json:"description,omitempty"`

	// Optional. Controls the randomness of the output.
	//
	//  Values can range over `[0.0,1.0]`, inclusive. A value closer to `1.0` will
	//  produce responses that are more varied, while a value closer to `0.0` will
	//  typically result in less surprising responses from the model.
	//
	//  This value specifies default to be the one used by the base model while
	//  creating the model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.temperature
	Temperature *float32 `json:"temperature,omitempty"`

	// Optional. For Nucleus sampling.
	//
	//  Nucleus sampling considers the smallest set of tokens whose probability
	//  sum is at least `top_p`.
	//
	//  This value specifies default to be the one used by the base model while
	//  creating the model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.top_p
	TopP *float32 `json:"topP,omitempty"`

	// Optional. For Top-k sampling.
	//
	//  Top-k sampling considers the set of `top_k` most probable tokens.
	//  This value specifies default to be used by the backend while making the
	//  call to the model.
	//
	//  This value specifies default to be the one used by the base model while
	//  creating the model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.top_k
	TopK *int32 `json:"topK,omitempty"`

	// Required. The tuning task that creates the tuned model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.tuning_task
	TuningTask *TuningTask `json:"tuningTask,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TunedModelSource
type TunedModelSource struct {
	// Immutable. The name of the `TunedModel` to use as the starting point for
	//  training the new model.
	//  Example: `tunedModels/my-tuned-model`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModelSource.tuned_model
	TunedModel *string `json:"tunedModel,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TuningExample
type TuningExample struct {
	// Optional. Text model input.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningExample.text_input
	TextInput *string `json:"textInput,omitempty"`

	// Required. The expected model output.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningExample.output
	Output *string `json:"output,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TuningExamples
type TuningExamples struct {
	// Required. The examples. Example input can be for text or discuss, but all
	//  examples in a set must be of the same type.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningExamples.examples
	Examples []TuningExample `json:"examples,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TuningSnapshot
type TuningSnapshot struct {
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TuningTask
type TuningTask struct {

	// Required. Input only. Immutable. The model training data.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningTask.training_data
	TrainingData *Dataset `json:"trainingData,omitempty"`

	// Immutable. Hyperparameters controlling the tuning process. If not provided,
	//  default values will be used.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningTask.hyperparameters
	Hyperparameters *Hyperparameters `json:"hyperparameters,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TunedModel
type TunedModelObservedState struct {
	// Optional. TunedModel to use as the starting point for training the new
	//  model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.tuned_model_source
	TunedModelSource *TunedModelSourceObservedState `json:"tunedModelSource,omitempty"`

	// Output only. The tuned model name. A unique name will be generated on
	//  create. Example: `tunedModels/az2mb0bpw6i` If display_name is set on
	//  create, the id portion of the name will be set by concatenating the words
	//  of the display_name with hyphens and adding a random portion for
	//  uniqueness. Example:
	//      display_name = "Sentence Translator"
	//      name = "tunedModels/sentence-translator-u3b7m"
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.name
	Name *string `json:"name,omitempty"`

	// Output only. The state of the tuned model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.state
	State *string `json:"state,omitempty"`

	// Output only. The timestamp when this model was created.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this model was updated.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. The tuning task that creates the tuned model.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModel.tuning_task
	TuningTask *TuningTaskObservedState `json:"tuningTask,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TunedModelSource
type TunedModelSourceObservedState struct {
	// Output only. The name of the base `Model` this `TunedModel` was tuned from.
	//  Example: `models/text-bison-001`
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TunedModelSource.base_model
	BaseModel *string `json:"baseModel,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TuningSnapshot
type TuningSnapshotObservedState struct {
	// Output only. The tuning step.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningSnapshot.step
	Step *int32 `json:"step,omitempty"`

	// Output only. The epoch this step was part of.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningSnapshot.epoch
	Epoch *int32 `json:"epoch,omitempty"`

	// Output only. The mean loss of the training examples for this step.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningSnapshot.mean_loss
	MeanLoss *float32 `json:"meanLoss,omitempty"`

	// Output only. The timestamp when this metric was computed.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningSnapshot.compute_time
	ComputeTime *string `json:"computeTime,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.TuningTask
type TuningTaskObservedState struct {
	// Output only. The timestamp when tuning this model started.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningTask.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The timestamp when tuning this model completed.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningTask.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. Metrics collected during tuning.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.TuningTask.snapshots
	Snapshots []TuningSnapshot `json:"snapshots,omitempty"`
}
