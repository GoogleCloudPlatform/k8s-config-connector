// Copyright 2024 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var WorkstationsWorkstationConfigGVK = GroupVersion.WithKind("WorkstationsWorkstationConfig")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkstationsWorkstationConfigSpec defines the desired state of WorkstationsWorkstationConfig
// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig
type WorkstationsWorkstationConfigSpec struct {
	// The WorkstationsWorkstationConfig name. If not given, the metadata.name will be used.
	// + optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Full name of this workstation configuration.
	Name *string `json:"name,omitempty"`

	// Optional. Human-readable name for this workstation configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Client-specified annotations.
	Annotations map[string]string `json:"annotations,omitempty"`

	// NOTYET: Not dealing with labels yet
	// // Optional.
	// //  [Labels](https://cloud.google.com/workstations/docs/label-resources) that
	// //  are applied to the workstation configuration and that are also propagated
	// //  to the underlying Compute Engine resources.
	// Labels map[string]string `json:"labels,omitempty"`

	// Optional. Number of seconds to wait before automatically stopping a
	//  workstation after it last received user traffic.
	//
	//  A value of `"0s"` indicates that Cloud Workstations VMs created with this
	//  configuration should never time out due to idleness.
	//  Provide
	//  [duration](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#duration)
	//  terminated by `s` for seconds—for example, `"7200s"` (2 hours).
	//  The default is `"1200s"` (20 minutes).
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Optional. Number of seconds that a workstation can run until it is
	//  automatically shut down. We recommend that workstations be shut down daily
	//  to reduce costs and so that security updates can be applied upon restart.
	//  The
	//  [idle_timeout][google.cloud.workstations.v1.WorkstationConfig.idle_timeout]
	//  and
	//  [running_timeout][google.cloud.workstations.v1.WorkstationConfig.running_timeout]
	//  fields are independent of each other. Note that the
	//  [running_timeout][google.cloud.workstations.v1.WorkstationConfig.running_timeout]
	//  field shuts down VMs after the specified time, regardless of whether or not
	//  the VMs are idle.
	//
	//  Provide duration terminated by `s` for seconds—for example, `"54000s"`
	//  (15 hours). Defaults to `"43200s"` (12 hours). A value of `"0s"` indicates
	//  that workstations using this configuration should never time out. If
	//  [encryption_key][google.cloud.workstations.v1.WorkstationConfig.encryption_key]
	//  is set, it must be greater than `"0s"` and less than
	//  `"86400s"` (24 hours).
	//
	//  Warning: A value of `"0s"` indicates that Cloud Workstations VMs created
	//  with this configuration have no maximum running time. This is strongly
	//  discouraged because you incur costs and will not pick up security updates.
	RunningTimeout *string `json:"runningTimeout,omitempty"`

	// Optional. Runtime host for the workstation.
	Host *WorkstationConfig_Host `json:"host,omitempty"`

	// Optional. Directories to persist across workstation sessions.
	PersistentDirectories []WorkstationConfig_PersistentDirectory `json:"persistentDirectories,omitempty"`

	// Optional. Container that runs upon startup for each workstation using this
	//  workstation configuration.
	Container *WorkstationConfig_Container `json:"container,omitempty"`

	// Immutable. Encrypts resources of this workstation configuration using a
	//  customer-managed encryption key (CMEK).
	//
	//  If specified, the boot disk of the Compute Engine instance and the
	//  persistent disk are encrypted using this encryption key. If
	//  this field is not set, the disks are encrypted using a generated
	//  key. Customer-managed encryption keys do not protect disk metadata.
	//
	//  If the customer-managed encryption key is rotated, when the workstation
	//  instance is stopped, the system attempts to recreate the
	//  persistent disk with the new version of the key. Be sure to keep
	//  older versions of the key until the persistent disk is recreated.
	//  Otherwise, data on the persistent disk might be lost.
	//
	//  If the encryption key is revoked, the workstation session automatically
	//  stops within 7 hours.
	//
	//  Immutable after the workstation configuration is created.
	EncryptionKey *WorkstationConfig_CustomerEncryptionKey `json:"encryptionKey,omitempty"`

	// Optional. Readiness checks to perform when starting a workstation using
	//  this workstation configuration. Mark a workstation as running only after
	//  all specified readiness checks return 200 status codes.
	ReadinessChecks []WorkstationConfig_ReadinessCheck `json:"readinessChecks,omitempty"`

	// Optional. Immutable. Specifies the zones used to replicate the VM and disk
	//  resources within the region. If set, exactly two zones within the
	//  workstation cluster's region must be specified—for example,
	//  `['us-central1-a', 'us-central1-f']`. If this field is empty, two default
	//  zones within the region are used.
	//
	//  Immutable after the workstation configuration is created.
	ReplicaZones []string `json:"replicaZones,omitempty"`
}

// WorkstationsWorkstationConfigStatus defines the config connector machine state of WorkstationsWorkstationConfig
type WorkstationsWorkstationConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the WorkstationsWorkstationConfig resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *WorkstationsWorkstationConfigObservedState `json:"observedState,omitempty"`
}

// WorkstationsWorkstationConfigSpec defines the desired state of WorkstationsWorkstationConfig
// +kcc:proto=google.cloud.workstations.v1.WorkstationConfig
type WorkstationsWorkstationConfigObservedState struct {

	// Output only. A system-assigned unique identifier for this workstation
	//  configuration.
	Uid *string `json:"uid,omitempty"`

	// NOTYET: This may be better surfaced as status.conditions?
	// // Output only. Indicates whether this workstation configuration is currently
	// //  being updated to match its intended state.
	// Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Time when this workstation configuration was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this workstation configuration was most recently
	//  updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time when this workstation configuration was soft-deleted.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Optional. Checksum computed by the server. May be sent on update and delete
	//  requests to make sure that the client has an up-to-date value before
	//  proceeding.
	Etag *string `json:"etag,omitempty"`

	// NOTYET: This may be better surfaced as status.conditions?
	// // Output only. Whether this resource is degraded, in which case it may
	// //  require user action to restore full functionality. See also the
	// //  [conditions][google.cloud.workstations.v1.WorkstationConfig.conditions]
	// //  field.
	// Degraded *bool `json:"degraded,omitempty"`

	// // Output only. Status conditions describing the current resource state.
	// Conditions []google_rpc_Status `json:"conditions,omitempty"`

}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkstationsWorkstationConfig is the Schema for the WorkstationsWorkstationConfig API
// +k8s:openapi-gen=true
type WorkstationsWorkstationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkstationsWorkstationConfigSpec   `json:"spec,omitempty"`
	Status WorkstationsWorkstationConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WorkstationsWorkstationConfigList contains a list of WorkstationsWorkstationConfig
type WorkstationsWorkstationConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkstationsWorkstationConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkstationsWorkstationConfig{}, &WorkstationsWorkstationConfigList{})
}
