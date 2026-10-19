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

// API sources for DialogflowTestCase, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/test_case.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.testCases

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowTestCaseGVK = GroupVersion.WithKind("DialogflowTestCase")

// DialogflowTestCaseSpec defines the desired state of DialogflowTestCase
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.TestCase
// +kcc:required-from-proto
type DialogflowTestCaseSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DialogflowTestCase name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Tags are short descriptions that users may apply to test cases for
	//  organizational and filtering purposes. Each tag should start with "#" and
	//  has a limit of 30 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.tags
	Tags []string `json:"tags,omitempty"`

	// Required. The human-readable name of the test case, unique within the
	//  agent. Limit of 200 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Additional freeform notes about the test case. Limit of 400 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.notes
	Notes *string `json:"notes,omitempty"`

	// Config for the test case.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.test_config
	TestConfig *TestConfig `json:"testConfig,omitempty"`

	// The conversation turns uttered when the test case was created, in
	//  chronological order. These include the canonical set of agent utterances
	//  that should occur when the agent is working properly.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.test_case_conversation_turns
	TestCaseConversationTurns []ConversationTurn `json:"testCaseConversationTurns,omitempty"`

	// The latest test result.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.last_test_result
	LastTestResult *TestCaseResult `json:"lastTestResult,omitempty"`
}

// DialogflowTestCaseStatus defines the config connector machine state of DialogflowTestCase
type DialogflowTestCaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowTestCase resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowTestCaseObservedState `json:"observedState,omitempty"`
}

// DialogflowTestCaseObservedState is the state of the DialogflowTestCase resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.TestCase
type DialogflowTestCaseObservedState struct {
	// The conversation turns uttered when the test case was created, in
	//  chronological order. These include the canonical set of agent utterances
	//  that should occur when the agent is working properly.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.test_case_conversation_turns
	TestCaseConversationTurns []ConversationTurnObservedState `json:"testCaseConversationTurns,omitempty"`

	// Output only. When the test was created.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.creation_time
	CreationTime *string `json:"creationTime,omitempty"`

	// The latest test result.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TestCase.last_test_result
	LastTestResult *TestCaseResultObservedState `json:"lastTestResult,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowtestcase;gcpdialogflowtestcases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowTestCase is the Schema for the DialogflowTestCase API
// +k8s:openapi-gen=true
type DialogflowTestCase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowTestCaseSpec   `json:"spec,omitempty"`
	Status DialogflowTestCaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowTestCaseList contains a list of DialogflowTestCase
type DialogflowTestCaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowTestCase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowTestCase{}, &DialogflowTestCaseList{})
}
