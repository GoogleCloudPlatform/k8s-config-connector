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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ContentWarehouseRuleSetGVK = GroupVersion.WithKind("ContentWarehouseRuleSet")

// ContentWarehouseRuleSetSpec defines the desired state of ContentWarehouseRuleSet
// +kcc:spec:proto=google.cloud.contentwarehouse.v1.RuleSet
type ContentWarehouseRuleSetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// Short description of the rule-set.
	// +optional
	Description *string `json:"description,omitempty"`

	// Source of the rules i.e., customer name.
	// +optional
	Source *string `json:"source,omitempty"`

	// List of rules given by the customer.
	// +optional
	Rules []Rule `json:"rules,omitempty"`

	// The ContentWarehouseRuleSet name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.Rule
type Rule struct {
	// Short description of the rule and its context.
	// +optional
	Description *string `json:"description,omitempty"`

	// ID of the rule. It has to be unique across all the examples.
	//  This is managed internally.
	// +optional
	RuleID *string `json:"ruleID,omitempty"`

	// Identifies the trigger type for running the policy.
	// +optional
	TriggerType *string `json:"triggerType,omitempty"`

	// Represents the conditional expression to be evaluated.
	//  Expression should evaluate to a boolean result.
	//  When the condition is true actions are executed.
	//  Example: user_role = "hsbc_role_1" AND doc.salary > 20000
	// +optional
	Condition *string `json:"condition,omitempty"`

	// List of actions that are executed when the rule is satisfied.
	// +optional
	Actions []Action `json:"actions,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.Action
type Action struct {
	// ID of the action. Managed internally.
	// +optional
	ActionID *string `json:"actionID,omitempty"`

	// Action triggering access control operations.
	// +optional
	AccessControl *AccessControlAction `json:"accessControl,omitempty"`

	// Action triggering data validation operations.
	// +optional
	DataValidation *DataValidationAction `json:"dataValidation,omitempty"`

	// Action triggering data update operations.
	// +optional
	DataUpdate *DataUpdateAction `json:"dataUpdate,omitempty"`

	// Action triggering create document link operation.
	// +optional
	AddToFolder *AddToFolderAction `json:"addToFolder,omitempty"`

	// Action publish to Pub/Sub operation.
	// +optional
	PublishToPubSub *PublishAction `json:"publishToPubSub,omitempty"`

	// Action removing a document from a folder.
	// +optional
	RemoveFromFolderAction *RemoveFromFolderAction `json:"removeFromFolderAction,omitempty"`

	// Action deleting the document.
	// +optional
	DeleteDocumentAction *DeleteDocumentAction `json:"deleteDocumentAction,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.AccessControlAction
type AccessControlAction struct {
	// Identifies the type of operation.
	// +optional
	OperationType *string `json:"operationType,omitempty"`

	// Represents the new policy from which bindings are added, removed or
	//  replaced based on the type of the operation. the policy is limited to a few
	//  10s of KB.
	// +optional
	Policy *Policy `json:"policy,omitempty"`
}

// +kcc:proto=google.iam.v1.Policy
type Policy struct {
	// Specifies the format of the policy.
	// +optional
	Version *int32 `json:"version,omitempty"`

	// Associates a list of `members`, or principals, with a `role`.
	// +optional
	Bindings []Binding `json:"bindings,omitempty"`

	// Specifies cloud audit logging configuration for this policy.
	// +optional
	AuditConfigs []AuditConfig `json:"auditConfigs,omitempty"`
}

// +kcc:proto=google.iam.v1.Binding
type Binding struct {
	// Role that is assigned to the list of `members`, or principals.
	// +optional
	Role *string `json:"role,omitempty"`

	// Specifies the principals requesting access for a Google Cloud resource.
	// +optional
	Members []string `json:"members,omitempty"`

	// The condition that is associated with this binding.
	// +optional
	Condition *Expr `json:"condition,omitempty"`
}

// +kcc:proto=google.type.Expr
type Expr struct {
	// Textual representation of an expression in Common Expression Language syntax.
	// +optional
	Expression *string `json:"expression,omitempty"`

	// Optional. Title for the expression.
	// +optional
	Title *string `json:"title,omitempty"`

	// Optional. Description of the expression.
	// +optional
	Description *string `json:"description,omitempty"`

	// Optional. String indicating the location of the expression for error reporting.
	// +optional
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.iam.v1.AuditConfig
type AuditConfig struct {
	// Specifies a service that will be enabled for audit logging.
	// +optional
	Service *string `json:"service,omitempty"`

	// The configuration for logging of each type of permission.
	// +optional
	AuditLogConfigs []AuditLogConfig `json:"auditLogConfigs,omitempty"`
}

// +kcc:proto=google.iam.v1.AuditLogConfig
type AuditLogConfig struct {
	// The log type that this config enables.
	// +optional
	LogType *string `json:"logType,omitempty"`

	// Specifies the identities that do not cause logging for this type of permission.
	// +optional
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DataValidationAction
type DataValidationAction struct {
	// Map of (K, V) -> (field, string condition to be evaluated on the field)
	// +optional
	Conditions map[string]string `json:"conditions,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DataUpdateAction
type DataUpdateAction struct {
	// Map of (K, V) -> (valid name of the field, new value of the field)
	// +optional
	Entries map[string]string `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.AddToFolderAction
type AddToFolderAction struct {
	// Names of the folder under which new document is to be added.
	// +optional
	Folders []string `json:"folders,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PublishAction
type PublishAction struct {
	// The topic id in the Pub/Sub service for which messages will be published to.
	// +optional
	TopicID *string `json:"topicID,omitempty"`

	// Messages to be published.
	// +optional
	Messages []string `json:"messages,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.RemoveFromFolderAction
type RemoveFromFolderAction struct {
	// Condition of the action to be executed.
	// +optional
	Condition *string `json:"condition,omitempty"`

	// Name of the folder under which new document is to be added.
	// +optional
	Folder *string `json:"folder,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DeleteDocumentAction
type DeleteDocumentAction struct {
	// Boolean field to select between hard vs soft delete options.
	// +optional
	EnableHardDelete *bool `json:"enableHardDelete,omitempty"`
}

// ContentWarehouseRuleSetStatus defines the config connector machine state of ContentWarehouseRuleSet
type ContentWarehouseRuleSetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ContentWarehouseRuleSet resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *ContentWarehouseRuleSetObservedState `json:"observedState,omitempty"`
}

// ContentWarehouseRuleSetObservedState is the state of the ContentWarehouseRuleSet resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.contentwarehouse.v1.RuleSet
// +kubebuilder:pruning:PreserveUnknownFields
type ContentWarehouseRuleSetObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontentwarehouseruleset;gcpcontentwarehouserulesets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ContentWarehouseRuleSet is the Schema for the ContentWarehouseRuleSet API
// +k8s:openapi-gen=true
type ContentWarehouseRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ContentWarehouseRuleSetSpec   `json:"spec,omitempty"`
	Status ContentWarehouseRuleSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContentWarehouseRuleSetList contains a list of ContentWarehouseRuleSet
type ContentWarehouseRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentWarehouseRuleSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContentWarehouseRuleSet{}, &ContentWarehouseRuleSetList{})
}
