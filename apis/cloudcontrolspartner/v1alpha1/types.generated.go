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


// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Customer
type Customer struct {
	// Identifier. Format:
	//  `organizations/{organization}/locations/{location}/customers/{customer}`
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Customer.name
	Name *string `json:"name,omitempty"`

	// Required. Display name for the customer
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Customer.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingState
type CustomerOnboardingState struct {
	// List of customer onboarding steps
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingState.onboarding_steps
	OnboardingSteps []CustomerOnboardingStep `json:"onboardingSteps,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingStep
type CustomerOnboardingStep struct {
	// The onboarding step
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingStep.step
	Step *string `json:"step,omitempty"`

	// The starting time of the onboarding step
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingStep.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The completion time of the onboarding step
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingStep.completion_time
	CompletionTime *string `json:"completionTime,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.Customer
type CustomerObservedState struct {
	// Output only. Container for customer onboarding steps
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Customer.customer_onboarding_state
	CustomerOnboardingState *CustomerOnboardingState `json:"customerOnboardingState,omitempty"`

	// Output only. Indicates whether a customer is fully onboarded
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.Customer.is_onboarded
	IsOnboarded *bool `json:"isOnboarded,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingState
type CustomerOnboardingStateObservedState struct {
	// List of customer onboarding steps
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingState.onboarding_steps
	OnboardingSteps []CustomerOnboardingStepObservedState `json:"onboardingSteps,omitempty"`
}

// +kcc:proto=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingStep
type CustomerOnboardingStepObservedState struct {
	// Output only. Current state of the step
	// +kcc:proto:field=google.cloud.cloudcontrolspartner.v1.CustomerOnboardingStep.completion_state
	CompletionState *string `json:"completionState,omitempty"`
}
