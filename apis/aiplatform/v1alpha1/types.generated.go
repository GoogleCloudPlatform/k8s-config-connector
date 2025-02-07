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


// +kcc:proto=google.cloud.aiplatform.v1.Measurement
type Measurement struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.Measurement.Metric
type Measurement_Metric struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.NasTrial
type NasTrial struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.NasTrialDetail
type NasTrialDetail struct {

	// The parameters for the NasJob NasTrial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrialDetail.parameters
	Parameters *string `json:"parameters,omitempty"`

	// The requested search NasTrial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrialDetail.search_trial
	SearchTrial *NasTrial `json:"searchTrial,omitempty"`

	// The train NasTrial corresponding to
	//  [search_trial][google.cloud.aiplatform.v1.NasTrialDetail.search_trial].
	//  Only populated if
	//  [search_trial][google.cloud.aiplatform.v1.NasTrialDetail.search_trial] is
	//  used for training.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrialDetail.train_trial
	TrainTrial *NasTrial `json:"trainTrial,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Measurement
type MeasurementObservedState struct {
	// Output only. Time that the Trial has been running at the point of this
	//  Measurement.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.elapsed_duration
	ElapsedDuration *string `json:"elapsedDuration,omitempty"`

	// Output only. The number of steps the machine learning model has been
	//  trained for. Must be non-negative.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.step_count
	StepCount *int64 `json:"stepCount,omitempty"`

	// Output only. A list of metrics got by evaluating the objective functions
	//  using suggested Parameter values.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.metrics
	Metrics []Measurement_Metric `json:"metrics,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Measurement.Metric
type Measurement_MetricObservedState struct {
	// Output only. The ID of the Metric. The Metric should be defined in
	//  [StudySpec's Metrics][google.cloud.aiplatform.v1.StudySpec.metrics].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.Metric.metric_id
	MetricID *string `json:"metricID,omitempty"`

	// Output only. The value for this metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Measurement.Metric.value
	Value *float64 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasTrial
type NasTrialObservedState struct {
	// Output only. The identifier of the NasTrial assigned by the service.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.id
	ID *string `json:"id,omitempty"`

	// Output only. The detailed state of the NasTrial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.state
	State *string `json:"state,omitempty"`

	// Output only. The final measurement containing the objective value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.final_measurement
	FinalMeasurement *Measurement `json:"finalMeasurement,omitempty"`

	// Output only. Time when the NasTrial was started.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the NasTrial's status changed to `SUCCEEDED` or
	//  `INFEASIBLE`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrial.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasTrialDetail
type NasTrialDetailObservedState struct {
	// Output only. Resource name of the NasTrialDetail.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrialDetail.name
	Name *string `json:"name,omitempty"`

	// The requested search NasTrial.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasTrialDetail.search_trial
	SearchTrial *NasTrialObservedState `json:"searchTrial,omitempty"`
}
