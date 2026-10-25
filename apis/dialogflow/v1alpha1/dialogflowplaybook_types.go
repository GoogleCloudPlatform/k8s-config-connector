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

// API sources for DialogflowPlaybook, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/playbook.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.playbooks

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var DialogflowPlaybookGVK = GroupVersion.WithKind("DialogflowPlaybook")

// DialogflowPlaybookSpec defines the desired state of DialogflowPlaybook
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Playbook
// +kcc:required-from-proto
type DialogflowPlaybookSpec struct {
	// The DialogflowCXAgent that this resource belongs to.
	// +required
	AgentRef *DialogflowCXAgentRef `json:"agentRef,omitempty"`

	// The DialogflowPlaybook name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human-readable name of the playbook, unique within an agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. High level description of the goal the playbook intend to
	//  accomplish. A goal should be concise since it's visible to other playbooks
	//  that may reference this playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.goal
	// +required
	Goal *string `json:"goal,omitempty"`

	// Optional. Defined structured input parameters for this playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.input_parameter_definitions
	InputParameterDefinitions []ParameterDefinition `json:"inputParameterDefinitions,omitempty"`

	// Optional. Defined structured output parameters for this playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.output_parameter_definitions
	OutputParameterDefinitions []ParameterDefinition `json:"outputParameterDefinitions,omitempty"`

	// Instruction to accomplish target goal.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.instruction
	Instruction *Playbook_Instruction `json:"instruction,omitempty"`

	// Optional. The resource name of tools referenced by the current playbook in
	//  the instructions. If not provided explicitly, they are will
	//  be implied using the tool being referenced in goal and steps.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.referenced_tools
	ReferencedToolRefs []DialogflowToolRef `json:"referencedToolRefs,omitempty"`

	// Optional. The playbook's scoped code block, which may implement handlers
	//  and actions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.code_block
	CodeBlock *CodeBlock `json:"codeBlock,omitempty"`

	// Optional. Llm model settings for the playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.llm_model_settings
	LlmModelSettings *LlmModelSettings `json:"llmModelSettings,omitempty"`

	// Optional. A list of registered handlers to execuate based on the specified
	//  triggers.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.handlers
	Handlers []Handler `json:"handlers,omitempty"`

	// Optional. Type of the playbook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.playbook_type
	PlaybookType *string `json:"playbookType,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.ParameterDefinition
type ParameterDefinition struct {
	// Required. Name of parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ParameterDefinition.name
	// +required
	Name *string `json:"name,omitempty"`

	// Optional. Type schema of parameter.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ParameterDefinition.type_schema
	TypeSchema *TypeSchema `json:"typeSchema,omitempty"`

	// Human-readable description of the parameter. Limited to 300 characters.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.ParameterDefinition.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.TypeSchema.SchemaReference
type TypeSchema_SchemaReference struct {
	// The tool that contains this schema definition.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/tools/<ToolID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TypeSchema.SchemaReference.tool
	ToolRef *DialogflowToolRef `json:"toolRef,omitempty"`

	// The name of the schema.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.TypeSchema.SchemaReference.schema
	Schema *string `json:"schema,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Handler
type Handler struct {
	// A handler triggered by event.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Handler.event_handler
	EventHandler *Handler_EventHandler `json:"eventHandler,omitempty"`

	// A handler triggered during specific lifecycle of the playbook execution.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Handler.lifecycle_handler
	LifecycleHandler *Handler_LifecycleHandler `json:"lifecycleHandler,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Handler.EventHandler
type Handler_EventHandler struct {
	// Required. The name of the event that triggers this handler.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Handler.EventHandler.event
	// +required
	Event *string `json:"event,omitempty"`

	// Optional. The condition that must be satisfied to trigger this handler.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Handler.EventHandler.condition
	Condition *string `json:"condition,omitempty"`

	// Required. The fulfillment to call when the event occurs.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Handler.EventHandler.fulfillment
	// +required
	Fulfillment *DialogflowCX_Fulfillment `json:"fulfillment,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Handler.LifecycleHandler
type Handler_LifecycleHandler struct {
	// Required. The name of the lifecycle stage that triggers this handler.
	//  Supported values:
	//  * `playbook-start`
	//  * `pre-action-selection`
	//  * `pre-action-execution`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Handler.LifecycleHandler.lifecycle_stage
	// +required
	LifecycleStage *string `json:"lifecycleStage,omitempty"`

	// Optional. The condition that must be satisfied to trigger this handler.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Handler.LifecycleHandler.condition
	Condition *string `json:"condition,omitempty"`

	// Required. The fulfillment to call when this handler is triggered.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Handler.LifecycleHandler.fulfillment
	// +required
	Fulfillment *DialogflowCX_Fulfillment `json:"fulfillment,omitempty"`
}

var (
	DialogflowToolIdentityFormat = gcpurls.Template[DialogflowToolIdentity]("dialogflow.googleapis.com", "projects/{project}/locations/{location}/agents/{agent}/tools/{tool}")
)

// +k8s:deepcopy-gen=false
type DialogflowToolIdentity struct {
	Project  string
	Location string
	Agent    string
	Tool     string
}

func (i *DialogflowToolIdentity) String() string {
	return DialogflowToolIdentityFormat.ToString(*i)
}

func (i *DialogflowToolIdentity) FromExternal(ref string) error {
	parsed, match, err := DialogflowToolIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DialogflowTool external=%q was not known (use %s): %w", ref, DialogflowToolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DialogflowTool external=%q was not known (use %s)", ref, DialogflowToolIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DialogflowToolIdentity) Host() string {
	return DialogflowToolIdentityFormat.Host()
}

var (
	_ refsv1beta1.Ref         = &DialogflowToolRef{}
	_ refsv1beta1.ExternalRef = &DialogflowToolRef{}
)

var DialogflowToolGVK = GroupVersion.WithKind("DialogflowTool")

// DialogflowToolRef is a reference to a DialogflowTool resource.
type DialogflowToolRef struct {
	// A reference to an externally managed DialogflowTool resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/agents/{{agentID}}/tools/{{toolID}}".
	External string `json:"external,omitempty"`
}

func init() {
	refsv1beta1.Register(&DialogflowToolRef{})
}

func (r *DialogflowToolRef) GetGVK() schema.GroupVersionKind {
	return DialogflowToolGVK
}

func (r *DialogflowToolRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *DialogflowToolRef) GetExternal() string {
	return r.External
}

func (r *DialogflowToolRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DialogflowToolRef) ValidateExternal(ref string) error {
	id := &DialogflowToolIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DialogflowToolRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DialogflowToolIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DialogflowToolRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", DialogflowToolGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}

// DialogflowPlaybookStatus defines the config connector machine state of DialogflowPlaybook
type DialogflowPlaybookStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowPlaybook resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowPlaybookObservedState `json:"observedState,omitempty"`
}

// DialogflowPlaybookObservedState is the state of the DialogflowPlaybook resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Playbook
type DialogflowPlaybookObservedState struct {
	// Output only. Estimated number of tokes current playbook takes when sent to
	//  the LLM.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.token_count
	TokenCount *int64 `json:"tokenCount,omitempty"`

	// Output only. The timestamp of initial playbook creation.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last time the playbook version was updated.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The resource name of other playbooks referenced by the current
	//  playbook in the instructions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.referenced_playbooks
	ReferencedPlaybooks []string `json:"referencedPlaybooks,omitempty"`

	// Output only. The resource name of flows referenced by the current playbook
	//  in the instructions.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.referenced_flows
	ReferencedFlows []string `json:"referencedFlows,omitempty"`

	// Optional. Output only. Names of inline actions scoped to this playbook.
	//  These actions are in addition to those belonging to referenced tools, child
	//  playbooks, and flows, e.g. actions that are defined in the playbook's code
	//  block.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.inline_actions
	InlineActions []string `json:"inlineActions,omitempty"`

	// Optional. A list of registered handlers to execuate based on the specified
	//  triggers.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Playbook.handlers
	Handlers []HandlerObservedState `json:"handlers,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowplaybook;gcpdialogflowplaybooks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowPlaybook is the Schema for the DialogflowPlaybook API
// +k8s:openapi-gen=true
type DialogflowPlaybook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowPlaybookSpec   `json:"spec,omitempty"`
	Status DialogflowPlaybookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowPlaybookList contains a list of DialogflowPlaybook
type DialogflowPlaybookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowPlaybook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowPlaybook{}, &DialogflowPlaybookList{})
}
