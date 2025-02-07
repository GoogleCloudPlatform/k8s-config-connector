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


// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.Workload
type Workload struct {
	// Identifier. Format:
	//  `organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}`
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.name
	Name *string `json:"name,omitempty"`

	// Container for workload onboarding steps.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.workload_onboarding_state
	WorkloadOnboardingState *WorkloadOnboardingState `json:"workloadOnboardingState,omitempty"`

	// Indicates whether a workload is fully onboarded.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.is_onboarded
	IsOnboarded *bool `json:"isOnboarded,omitempty"`

	// The project id of the key management project for the workload
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.key_management_project_id
	KeyManagementProjectID *string `json:"keyManagementProjectID,omitempty"`

	// The Google Cloud location of the workload
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.location
	Location *string `json:"location,omitempty"`

	// Partner associated with this workload.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.partner
	Partner *string `json:"partner,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingState
type WorkloadOnboardingState struct {
	// List of workload onboarding steps.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingState.onboarding_steps
	OnboardingSteps []WorkloadOnboardingStep `json:"onboardingSteps,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingStep
type WorkloadOnboardingStep struct {
	// The onboarding step.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingStep.step
	Step *string `json:"step,omitempty"`

	// The starting time of the onboarding step.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingStep.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The completion time of the onboarding step.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingStep.completion_time
	CompletionTime *string `json:"completionTime,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.Workload
type WorkloadObservedState struct {
	// Output only. Folder id this workload is associated with
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.folder_id
	FolderID *int64 `json:"folderID,omitempty"`

	// Output only. Time the resource was created.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The name of container folder of the assured workload
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.folder
	Folder *string `json:"folder,omitempty"`

	// Container for workload onboarding steps.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.Workload.workload_onboarding_state
	WorkloadOnboardingState *WorkloadOnboardingStateObservedState `json:"workloadOnboardingState,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingState
type WorkloadOnboardingStateObservedState struct {
	// List of workload onboarding steps.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingState.onboarding_steps
	OnboardingSteps []WorkloadOnboardingStepObservedState `json:"onboardingSteps,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingStep
type WorkloadOnboardingStepObservedState struct {
	// Output only. The completion state of the onboarding step.
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1beta.WorkloadOnboardingStep.completion_state
	CompletionState *string `json:"completionState,omitempty"`
}
