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

// API sources for CesAgent, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/ces/v1/agent.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/reference/rest/v1/projects.locations.apps.agents

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	dialogflowv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var CesAgentGVK = GroupVersion.WithKind("CesAgent")

var _ refsv1beta1.Ref = &CesAppRef{}

// CesAppRef is a reference to a GCP CESApp resource.
type CesAppRef struct {
	// A reference to an externally managed CESApp resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}".
	External string `json:"external,omitempty"`

	// The name of a CESApp resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CESApp resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CesAppRef) GetGVK() schema.GroupVersionKind {
	return CESAppGVK
}

func (r *CesAppRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CesAppRef) GetExternal() string {
	return r.External
}

func (r *CesAppRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CesAppRef) ValidateExternal(ref string) error {
	id := &CESAppIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesAppRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CESAppIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesAppRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

var (
	CesGuardrailIdentityFormat                     = gcpurls.Template[CesGuardrailIdentity]("ces.googleapis.com", "projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}")
	_                          identity.IdentityV2 = &CesGuardrailIdentity{}
	CesGuardrailGVK                                = GroupVersion.WithKind("CesGuardrail")
	_                          refsv1beta1.Ref     = &CesGuardrailRef{}
)

// +k8s:deepcopy-gen=false
type CesGuardrailIdentity struct {
	Project   string
	Location  string
	App       string
	Guardrail string
}

func (i *CesGuardrailIdentity) String() string {
	return CesGuardrailIdentityFormat.ToString(*i)
}

func (i *CesGuardrailIdentity) FromExternal(ref string) error {
	parsed, match, err := CesGuardrailIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CesGuardrail external=%q was not known (use %s): %w", ref, CesGuardrailIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CesGuardrail external=%q was not known (use %s)", ref, CesGuardrailIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *CesGuardrailIdentity) Host() string {
	return CesGuardrailIdentityFormat.Host()
}

func (i *CesGuardrailIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/apps/" + i.App
}

// CesGuardrailRef is a reference to a GCP CesGuardrail resource.
type CesGuardrailRef struct {
	// A reference to an externally managed CesGuardrail resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}/guardrails/{{guardrailID}}".
	External string `json:"external,omitempty"`

	// The name of a CesGuardrail resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CesGuardrail resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CesGuardrailRef) GetGVK() schema.GroupVersionKind {
	return CesGuardrailGVK
}

func (r *CesGuardrailRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CesGuardrailRef) GetExternal() string {
	return r.External
}

func (r *CesGuardrailRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CesGuardrailRef) ValidateExternal(ref string) error {
	id := &CesGuardrailIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesGuardrailRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CesGuardrailIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesGuardrailRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

var (
	CesToolIdentityFormat                     = gcpurls.Template[CesToolIdentity]("ces.googleapis.com", "projects/{project}/locations/{location}/apps/{app}/tools/{tool}")
	_                     identity.IdentityV2 = &CesToolIdentity{}
	CesToolGVK                                = GroupVersion.WithKind("CesTool")
	_                     refsv1beta1.Ref     = &CesToolRef{}
)

// +k8s:deepcopy-gen=false
type CesToolIdentity struct {
	Project  string
	Location string
	App      string
	Tool     string
}

func (i *CesToolIdentity) String() string {
	return CesToolIdentityFormat.ToString(*i)
}

func (i *CesToolIdentity) FromExternal(ref string) error {
	parsed, match, err := CesToolIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CesTool external=%q was not known (use %s): %w", ref, CesToolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CesTool external=%q was not known (use %s)", ref, CesToolIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *CesToolIdentity) Host() string {
	return CesToolIdentityFormat.Host()
}

func (i *CesToolIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/apps/" + i.App
}

// CesToolRef is a reference to a GCP CesTool resource.
type CesToolRef struct {
	// A reference to an externally managed CesTool resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}/tools/{{toolID}}".
	External string `json:"external,omitempty"`

	// The name of a CesTool resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CesTool resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CesToolRef) GetGVK() schema.GroupVersionKind {
	return CesToolGVK
}

func (r *CesToolRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CesToolRef) GetExternal() string {
	return r.External
}

func (r *CesToolRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CesToolRef) ValidateExternal(ref string) error {
	id := &CesToolIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesToolRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CesToolIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesToolRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

var (
	CesToolsetIdentityFormat                     = gcpurls.Template[CesToolsetIdentity]("ces.googleapis.com", "projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}")
	_                        identity.IdentityV2 = &CesToolsetIdentity{}
	CesToolsetGVK                                = GroupVersion.WithKind("CesToolset")
	_                        refsv1beta1.Ref     = &CesToolsetRef{}
)

// +k8s:deepcopy-gen=false
type CesToolsetIdentity struct {
	Project  string
	Location string
	App      string
	Toolset  string
}

func (i *CesToolsetIdentity) String() string {
	return CesToolsetIdentityFormat.ToString(*i)
}

func (i *CesToolsetIdentity) FromExternal(ref string) error {
	parsed, match, err := CesToolsetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CesToolset external=%q was not known (use %s): %w", ref, CesToolsetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CesToolset external=%q was not known (use %s)", ref, CesToolsetIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *CesToolsetIdentity) Host() string {
	return CesToolsetIdentityFormat.Host()
}

func (i *CesToolsetIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/apps/" + i.App
}

// CesToolsetRef is a reference to a GCP CesToolset resource.
type CesToolsetRef struct {
	// A reference to an externally managed CesToolset resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}/toolsets/{{toolsetID}}".
	External string `json:"external,omitempty"`

	// The name of a CesToolset resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CesToolset resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CesToolsetRef) GetGVK() schema.GroupVersionKind {
	return CesToolsetGVK
}

func (r *CesToolsetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CesToolsetRef) GetExternal() string {
	return r.External
}

func (r *CesToolsetRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CesToolsetRef) ValidateExternal(ref string) error {
	id := &CesToolsetIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesToolsetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CesToolsetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesToolsetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

// +kcc:proto=google.cloud.ces.v1.Agent.AgentToolset
type Agent_AgentToolset struct {
	// Required. The resource name of the toolset.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.AgentToolset.toolset
	// +required
	ToolsetRef *CesToolsetRef `json:"toolsetRef,omitempty"`

	// Optional. The tools IDs to filter the toolset.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.AgentToolset.tool_ids
	ToolIDs []string `json:"toolIDs,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.Agent.RemoteDialogflowAgent
type Agent_RemoteDialogflowAgent struct {
	// Required. The
	//  [Dialogflow](https://docs.cloud.google.com/dialogflow/cx/docs/concept/agent)
	//  agent resource name.
	//  Format: `projects/{project}/locations/{location}/agents/{agent}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.RemoteDialogflowAgent.agent
	// +required
	AgentRef *dialogflowv1alpha1.DialogflowCXAgentRef `json:"agentRef,omitempty"`

	// Optional. The flow ID of the flow in the Dialogflow agent.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.RemoteDialogflowAgent.flow_id
	FlowID *string `json:"flowID,omitempty"`

	// Optional. The environment ID of the Dialogflow agent to be used for the
	//  agent execution. If not specified, the draft environment will be used.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.RemoteDialogflowAgent.environment_id
	EnvironmentID *string `json:"environmentID,omitempty"`

	// Optional. The mapping of the app variables names to the Dialogflow
	//  session parameters names to be sent to the Dialogflow agent as input.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.RemoteDialogflowAgent.input_variable_mapping
	InputVariableMapping map[string]string `json:"inputVariableMapping,omitempty"`

	// Optional. The mapping of the Dialogflow session parameters names to the
	//  app variables names to be sent back to the CES agent after the Dialogflow
	//  agent execution ends.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.RemoteDialogflowAgent.output_variable_mapping
	OutputVariableMapping map[string]string `json:"outputVariableMapping,omitempty"`

	// Optional. Indicates whether to respect the message-level interruption
	//  settings configured in the Dialogflow agent.
	//
	//  * If false: all response messages from the Dialogflow agent follow the
	//  app-level barge-in settings.
	//  * If true: only response messages with
	//  [`allow_playback_interruption`](https://docs.cloud.google.com/dialogflow/cx/docs/reference/rpc/google.cloud.dialogflow.cx.v3#text)
	//  set to true will be interruptable, all other messages follow the
	//  app-level barge-in settings.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.RemoteDialogflowAgent.respect_response_interruption_settings
	RespectResponseInterruptionSettings *bool `json:"respectResponseInterruptionSettings,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.TransferRule
type TransferRule struct {
	// Optional. A rule that immediately transfers to the target agent when the
	//  condition is met.
	// +kcc:proto:field=google.cloud.ces.v1.TransferRule.deterministic_transfer
	DeterministicTransfer *TransferRule_DeterministicTransfer `json:"deterministicTransfer,omitempty"`

	// Optional. Rule that prevents the planner from transferring to the target
	//  agent.
	// +kcc:proto:field=google.cloud.ces.v1.TransferRule.disable_planner_transfer
	DisablePlannerTransfer *TransferRule_DisablePlannerTransfer `json:"disablePlannerTransfer,omitempty"`

	// Required. The resource name of the child agent the rule applies to.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/agents/{agent}`
	// +kcc:proto:field=google.cloud.ces.v1.TransferRule.child_agent
	// +required
	ChildAgentRef *CesAgentRef `json:"childAgentRef,omitempty"`

	// Required. The direction of the transfer.
	// +kcc:proto:field=google.cloud.ces.v1.TransferRule.direction
	// +required
	Direction *string `json:"direction,omitempty"`
}

// CesAgentSpec defines the desired state of CesAgent
// +kcc:spec:proto=google.cloud.ces.v1.Agent
// +kcc:required-from-proto
type CesAgentSpec struct {
	// Required. The App that this Agent belongs to.
	// +required
	AppRef *CesAppRef `json:"appRef,omitempty"`

	// The CesAgent name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The default agent type.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.llm_agent
	LlmAgent *Agent_LlmAgent `json:"llmAgent,omitempty"`

	// Optional. The remote
	//  [Dialogflow](https://cloud.google.com/dialogflow/cx/docs/concept/console-conversational-agents)
	//  agent to be used for the agent execution. If this field is set, all other
	//  agent level properties will be ignored.
	//
	//  Note: If the Dialogflow agent is in a different project from the app, you
	//  should grant `roles/dialogflow.client` to the CES service agent
	//  `service-<PROJECT-NUMBER>@gcp-sa-ces.iam.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.remote_dialogflow_agent
	RemoteDialogflowAgent *Agent_RemoteDialogflowAgent `json:"remoteDialogflowAgent,omitempty"`

	// Required. Display name of the agent.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Human-readable description of the agent.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.description
	Description *string `json:"description,omitempty"`

	// Optional. Configurations for the LLM model.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.model_settings
	ModelSettings *ModelSettings `json:"modelSettings,omitempty"`

	// Optional. Instructions for the LLM model to guide the agent's behavior.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.instruction
	Instruction *string `json:"instruction,omitempty"`

	// Optional. List of available tools for the agent.
	//  Format: `projects/{project}/locations/{location}/apps/{app}/tools/{tool}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.tools
	ToolRefs []CesToolRef `json:"toolRefs,omitempty"`

	// Optional. List of child agents in the agent tree.
	//  Format: `projects/{project}/locations/{location}/apps/{app}/agents/{agent}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.child_agents
	ChildAgentRefs []CesAgentRef `json:"childAgentRefs,omitempty"`

	// Optional. The callbacks to execute before the agent is called.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.before_agent_callbacks
	BeforeAgentCallbacks []Callback `json:"beforeAgentCallbacks,omitempty"`

	// Optional. The callbacks to execute after the agent is called.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.after_agent_callbacks
	AfterAgentCallbacks []Callback `json:"afterAgentCallbacks,omitempty"`

	// Optional. The callbacks to execute before the model is called. If there are
	//  multiple calls to the model, the callback will be executed multiple times.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.before_model_callbacks
	BeforeModelCallbacks []Callback `json:"beforeModelCallbacks,omitempty"`

	// Optional. The callbacks to execute after the model is called. If there are
	//  multiple calls to the model, the callback will be executed multiple times.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.after_model_callbacks
	AfterModelCallbacks []Callback `json:"afterModelCallbacks,omitempty"`

	// Optional. The callbacks to execute before the tool is invoked. If there are
	//  multiple tool invocations, the callback will be executed multiple times.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.before_tool_callbacks
	BeforeToolCallbacks []Callback `json:"beforeToolCallbacks,omitempty"`

	// Optional. The callbacks to execute after the tool is invoked. If there are
	//  multiple tool invocations, the callback will be executed multiple times.
	//  The provided callbacks are executed sequentially in the exact order they
	//  are given in the list. If a callback returns an overridden response,
	//  execution stops and any remaining callbacks are skipped.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.after_tool_callbacks
	AfterToolCallbacks []Callback `json:"afterToolCallbacks,omitempty"`

	// Optional. List of guardrails for the agent.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}`
	// +kcc:proto:field=google.cloud.ces.v1.Agent.guardrails
	GuardrailRefs []CesGuardrailRef `json:"guardrailRefs,omitempty"`

	// Optional. List of toolsets for the agent.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.toolsets
	Toolsets []Agent_AgentToolset `json:"toolsets,omitempty"`

	// Optional. Agent transfer rules.
	//  If multiple rules match, the first one in the list will be used.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.transfer_rules
	TransferRules []TransferRule `json:"transferRules,omitempty"`
}

// CesAgentStatus defines the config connector machine state of CesAgent
type CesAgentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CesAgent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CesAgentObservedState `json:"observedState,omitempty"`
}

// CesAgentObservedState is the state of the CesAgent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1.Agent
type CesAgentObservedState struct {
	// Output only. Timestamp when the agent was created.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the agent was last updated.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. If the agent is generated by the LLM assistant, this field
	//  contains a descriptive summary of the generation.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.generated_summary
	GeneratedSummary *string `json:"generatedSummary,omitempty"`

	// Output only. Misconfigurations or errors in the agent that may affect agent
	//  quality.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.validation_errors
	ValidationErrors []string `json:"validationErrors,omitempty"`

	// Etag used to ensure the object hasn't changed during a read-modify-write
	//  operation. If the etag is empty, the update will overwrite any concurrent
	//  changes.
	// +kcc:proto:field=google.cloud.ces.v1.Agent.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesagent;gcpcesagents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CesAgent is the Schema for the CesAgent API
// +k8s:openapi-gen=true
type CesAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CesAgentSpec   `json:"spec,omitempty"`
	Status CesAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CesAgentList contains a list of CesAgent
type CesAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CesAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CesAgent{}, &CesAgentList{})
	refsv1beta1.Register(&CesAppRef{}, &CESApp{})
}
