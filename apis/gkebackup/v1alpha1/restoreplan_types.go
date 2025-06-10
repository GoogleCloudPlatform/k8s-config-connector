// Copyright 2025 Google LLC
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
	container "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEBackupRestorePlanGVK = GroupVersion.WithKind("GKEBackupRestorePlan")

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.ResourceFilter
type RestoreConfig_ResourceFilter struct {
	// Optional. (Filtering parameter) Any resource subject to transformation
	//  must be contained within one of the listed Kubernetes Namespace in the
	//  Backup. If this field is not provided, no namespace filtering will be
	//  performed (all resources in all Namespaces, including all cluster-scoped
	//  resources, will be candidates for transformation).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ResourceFilter.namespaces
	Namespaces []string `json:"namespaces,omitempty"`

	// Optional. (Filtering parameter) Any resource subject to transformation
	//  must belong to one of the listed "types". If this field is not provided,
	//  no type filtering will be performed (all resources of all types matching
	//  previous filtering parameters will be candidates for transformation).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ResourceFilter.group_kinds
	GroupKinds []RestoreConfig_GroupKind `json:"groupKinds,omitempty"`

	// Optional. This is a [JSONPath]
	//  (https://github.com/json-path/JsonPath/blob/master/README.md)
	//  expression that matches specific fields of candidate
	//  resources and it operates as a filtering parameter (resources that
	//  are not matched with this expression will not be candidates for
	//  transformation).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.ResourceFilter.json_path
	JSONPath *string `json:"jsonPath,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule
type RestoreConfig_SubstitutionRule struct {
	// Optional. (Filtering parameter) Any resource subject to substitution must
	//  be contained within one of the listed Kubernetes Namespace in the Backup.
	//  If this field is not provided, no namespace filtering will be performed
	//  (all resources in all Namespaces, including all cluster-scoped resources,
	//  will be candidates for substitution).
	//  To mix cluster-scoped and namespaced resources in the same rule, use an
	//  empty string ("") as one of the target namespaces.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.target_namespaces
	TargetNamespaces []string `json:"targetNamespaces,omitempty"`

	// Optional. (Filtering parameter) Any resource subject to substitution must
	//  belong to one of the listed "types". If this field is not provided, no
	//  type filtering will be performed (all resources of all types matching
	//  previous filtering parameters will be candidates for substitution).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.target_group_kinds
	TargetGroupKinds []RestoreConfig_GroupKind `json:"targetGroupKinds,omitempty"`

	// Required. This is a [JSONPath]
	//  (https://kubernetes.io/docs/reference/kubectl/jsonpath/)
	//  expression that matches specific fields of candidate
	//  resources and it operates as both a filtering parameter (resources that
	//  are not matched with this expression will not be candidates for
	//  substitution) as well as a field identifier (identifies exactly which
	//  fields out of the candidate resources will be modified).
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.target_json_path
	// +required
	TargetJSONPath *string `json:"targetJSONPath,omitempty"`

	// Optional. (Filtering parameter) This is a [regular expression]
	//  (https://en.wikipedia.org/wiki/Regular_expression)
	//  that is compared against the fields matched by the target_json_path
	//  expression (and must also have passed the previous filters).
	//  Substitution will not be performed against fields whose
	//  value does not match this expression. If this field is NOT specified,
	//  then ALL fields matched by the target_json_path expression will undergo
	//  substitution. Note that an empty (e.g., "", rather than unspecified)
	//  value for this field will only match empty fields.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.original_value_pattern
	OriginalValuePattern *string `json:"originalValuePattern,omitempty"`

	// Optional. This is the new value to set for any fields that pass the
	//  filtering and selection criteria. To remove a value from a Kubernetes
	//  resource, either leave this field unspecified, or set it to the empty
	//  string ("").
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.SubstitutionRule.new_value
	NewValue *string `json:"newValue,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder.GroupKindDependency
type RestoreConfig_RestoreOrder_GroupKindDependency struct {
	// Required. The satisfying group kind must be restored first
	//  in order to satisfy the dependency.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder.GroupKindDependency.satisfying
	// +required
	Satisfying *RestoreConfig_GroupKind `json:"satisfying,omitempty"`

	// Required. The requiring group kind requires that the other
	//  group kind be restored first.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.RestoreOrder.GroupKindDependency.requiring
	// +required
	Requiring *RestoreConfig_GroupKind `json:"requiring,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.TransformationRule
type RestoreConfig_TransformationRule struct {
	// Required. A list of transformation rule actions to take against candidate
	//  resources. Actions are executed in order defined - this order matters, as
	//  they could potentially interfere with each other and the first operation
	//  could affect the outcome of the second operation.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRule.field_actions
	// +required
	FieldActions []RestoreConfig_TransformationRuleAction `json:"fieldActions,omitempty"`

	// Optional. This field is used to specify a set of fields that should be
	//  used to determine which resources in backup should be acted upon by the
	//  supplied transformation rule actions, and this will ensure that only
	//  specific resources are affected by transformation rule actions.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRule.resource_filter
	ResourceFilter *RestoreConfig_ResourceFilter `json:"resourceFilter,omitempty"`

	// Optional. The description is a user specified string description of the
	//  transformation rule.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRule.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction
type RestoreConfig_TransformationRuleAction struct {
	// Required. op specifies the operation to perform.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction.op
	// +required
	Op *string `json:"op,omitempty"`

	// Optional. A string containing a JSON Pointer value that references the
	//  location in the target document to move the value from.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction.from_path
	FromPath *string `json:"fromPath,omitempty"`

	// Optional. A string containing a JSON-Pointer value that references a
	//  location within the target document where the operation is performed.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction.path
	Path *string `json:"path,omitempty"`

	// Optional. A string that specifies the desired value in string format to
	//  use for transformation.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.TransformationRuleAction.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RestoreConfig.VolumeDataRestorePolicyBinding
type RestoreConfig_VolumeDataRestorePolicyBinding struct {
	// Required. The VolumeDataRestorePolicy to apply when restoring volumes in
	//  scope.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.VolumeDataRestorePolicyBinding.policy
	// +required
	Policy *string `json:"policy,omitempty"`

	// The volume type, as determined by the PVC's bound PV,
	//  to apply the policy to.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestoreConfig.VolumeDataRestorePolicyBinding.volume_type
	VolumeType *string `json:"volumeType,omitempty"`
}

// GKEBackupRestorePlanSpec defines the desired state of GKEBackupRestorePlan
// +kcc:spec:proto=google.cloud.gkebackup.v1.RestorePlan
type GKEBackupRestorePlanSpec struct {
	// The GKEBackupRestorePlan name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. User specified descriptive string for this RestorePlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. A reference to the
	//  [BackupPlan][google.cloud.gkebackup.v1.BackupPlan] from which Backups may
	//  be used as the source for Restores created via this RestorePlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.backup_plan
	// +required
	BackupPlanRef *BackupPlanRef `json:"backupPlanRef,omitempty"`

	// Required. Immutable. The target cluster into which Restores created via
	//  this RestorePlan will restore data. NOTE: the cluster's region must be the
	//  same as the RestorePlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.cluster
	// +required
	ClusterRef *container.ContainerClusterRef `json:"clusterRef,omitempty"`

	// Required. Configuration of Restores created via this RestorePlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.restore_config
	// +required
	RestoreConfig *RestoreConfig `json:"restoreConfig,omitempty"`

	// Optional. A set of custom labels supplied by user.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// GKEBackupRestorePlanStatus defines the config connector machine state of GKEBackupRestorePlan
type GKEBackupRestorePlanStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEBackupRestorePlan resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEBackupRestorePlanObservedState `json:"observedState,omitempty"`
}

// GKEBackupRestorePlanObservedState is the state of the GKEBackupRestorePlan resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkebackup.v1.RestorePlan
type GKEBackupRestorePlanObservedState struct {
	// Output only. The full name of the RestorePlan resource.
	//  Format: `projects/*/locations/*/restorePlans/*`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The timestamp when this RestorePlan resource was
	//  created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this RestorePlan resource was last
	//  updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a restore from overwriting each other.
	//  It is strongly suggested that systems make use of the `etag` in the
	//  read-modify-write cycle to perform restore updates in order to avoid
	//  race conditions: An `etag` is returned in the response to `GetRestorePlan`,
	//  and systems are expected to put that etag in the request to
	//  `UpdateRestorePlan` or `DeleteRestorePlan` to ensure that their change
	//  will be applied to the same version of the resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. State of the RestorePlan. This State field reflects the
	//  various stages a RestorePlan can be in
	//  during the Create operation.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.state
	State *string `json:"state,omitempty"`

	// Output only. Human-readable description of why RestorePlan is in the
	//  current `state`
	// +kcc:proto:field=google.cloud.gkebackup.v1.RestorePlan.state_reason
	StateReason *string `json:"stateReason,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpgkebackuprestoreplan;gcpgkebackuprestoreplans
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEBackupRestorePlan is the Schema for the GKEBackupRestorePlan API
// +k8s:openapi-gen=true
type GKEBackupRestorePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEBackupRestorePlanSpec   `json:"spec,omitempty"`
	Status GKEBackupRestorePlanStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEBackupRestorePlanList contains a list of GKEBackupRestorePlan
type GKEBackupRestorePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEBackupRestorePlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEBackupRestorePlan{}, &GKEBackupRestorePlanList{})
}
