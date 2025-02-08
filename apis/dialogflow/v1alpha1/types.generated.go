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


// +kcc:proto=google.cloud.dialogflow.cx.v3.AgentValidationResult
type AgentValidationResult struct {
	// The unique identifier of the agent validation result.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/validationResult`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AgentValidationResult.name
	Name *string `json:"name,omitempty"`

	// Contains all flow validation results.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.AgentValidationResult.flow_validation_results
	FlowValidationResults []FlowValidationResult `json:"flowValidationResults,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.FlowValidationResult
type FlowValidationResult struct {
	// The unique identifier of the flow validation result.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/flows/<FlowID>/validationResult`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.FlowValidationResult.name
	Name *string `json:"name,omitempty"`

	// Contains all validation messages.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.FlowValidationResult.validation_messages
	ValidationMessages []ValidationMessage `json:"validationMessages,omitempty"`

	// Last time the flow was validated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.FlowValidationResult.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ResourceName
type ResourceName struct {
	// Name.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResourceName.name
	Name *string `json:"name,omitempty"`

	// Display name.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ResourceName.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ValidationMessage
type ValidationMessage struct {
	// The type of the resources where the message is found.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ValidationMessage.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// The names of the resources where the message is found.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ValidationMessage.resources
	Resources []string `json:"resources,omitempty"`

	// The resource names of the resources where the message is found.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ValidationMessage.resource_names
	ResourceNames []ResourceName `json:"resourceNames,omitempty"`

	// Indicates the severity of the message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ValidationMessage.severity
	Severity *string `json:"severity,omitempty"`

	// The message detail.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ValidationMessage.detail
	Detail *string `json:"detail,omitempty"`
}
