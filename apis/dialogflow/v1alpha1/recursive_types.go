// Copyright 2026 Google LLC
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

// +kcc:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent
type Fulfillment_ConditionalCases_Case_CaseContent struct {
	// Returned message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent.message
	Message *ResponseMessage `json:"message,omitempty"`

	// Additional cases to be evaluated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent.additional_cases
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	AdditionalCases *Fulfillment_ConditionalCases `json:"additionalCases,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent
type Fulfillment_ConditionalCases_Case_CaseContentObservedState struct {
	// Returned message.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Fulfillment.ConditionalCases.Case.CaseContent.message
	Message *ResponseMessageObservedState `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.InlineSchema
type InlineSchema struct {
	// Data type of the schema.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.InlineSchema.type
	Type *string `json:"type,omitempty"`

	// Schema of the elements if this is an ARRAY type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.InlineSchema.items
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	Items *TypeSchema `json:"items,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Playbook.Step
type Playbook_Step struct {
	// Step instruction in text format.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.Step.text
	Text *string `json:"text,omitempty"`

	// Sub-processing needed to execute the current step.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.Step.steps
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	Steps []Playbook_Step `json:"steps,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	Values []Value `json:"values,omitempty"`
}
