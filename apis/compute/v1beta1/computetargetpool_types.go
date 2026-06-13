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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeTargetPoolGVK = GroupVersion.WithKind("ComputeTargetPool")

type TargetPoolResourceRef struct {
	/* The external name of the referenced resource */
	// +optional
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	// +optional
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

type TargetpoolHealthChecks struct {
	// +optional
	HttpHealthCheckRef *TargetPoolResourceRef `json:"httpHealthCheckRef,omitempty"`
}

// ComputeTargetPoolSpec defines the desired state of ComputeTargetPool
// +kcc:spec:proto=google.cloud.compute.v1.TargetPool
type ComputeTargetPoolSpec struct {
	// +optional
	BackupTargetPoolRef *TargetPoolResourceRef `json:"backupTargetPoolRef,omitempty"`

	/* Immutable. Textual description field. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Ratio (0 to 1) of failed nodes before using the backup pool (which must also be set). */
	// +optional
	FailoverRatio *float64 `json:"failoverRatio,omitempty"`

	// +optional
	HealthChecks []TargetpoolHealthChecks `json:"healthChecks,omitempty"`

	// +optional
	Instances []TargetPoolResourceRef `json:"instances,omitempty"`

	/* Immutable. Where the target pool resides. Defaults to project region. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The resource URL for the security policy associated with this target pool. */
	// +optional
	SecurityPolicyRef *TargetPoolResourceRef `json:"securityPolicyRef,omitempty"`

	/* Immutable. How to distribute load. Options are "NONE" (no affinity). "CLIENT_IP" (hash of the source/dest addresses / ports), and "CLIENT_IP_PROTO" also includes the protocol (default "NONE"). */
	// +optional
	SessionAffinity *string `json:"sessionAffinity,omitempty"`
}

// ComputeTargetPoolStatus defines the config connector machine state of ComputeTargetPool
type ComputeTargetPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeTargetPool's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The URI of the created resource. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargetpool;gcpcomputetargetpools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeTargetPool is the Schema for the ComputeTargetPool API
// +k8s:openapi-gen=true
type ComputeTargetPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeTargetPoolSpec   `json:"spec,omitempty"`
	Status ComputeTargetPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeTargetPoolList contains a list of ComputeTargetPool
type ComputeTargetPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetPool{}, &ComputeTargetPoolList{})
}
