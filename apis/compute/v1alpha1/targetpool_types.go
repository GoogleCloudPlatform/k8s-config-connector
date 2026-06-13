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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeTargetPoolGVK = GroupVersion.WithKind("ComputeTargetPool")

type TargetpoolHealthChecks struct {
	// +optional
	HttpHealthCheckRef *ComputeHTTPHealthCheckRef `json:"httpHealthCheckRef,omitempty"`
}

// ComputeTargetPoolSpec defines the desired state of ComputeTargetPool
// +kcc:spec:proto=google.cloud.compute.v1.TargetPool
type ComputeTargetPoolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location (region) of this resource.
	Location string `json:"location"`

	// The ComputeTargetPool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The ComputeTargetPool backup pool reference. backupPool and failoverRatio together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below failoverRatio, traffic arriving at the load-balanced IP will be directed to the backup pool.
	// +optional
	BackupTargetPoolRef *ComputeTargetPoolRef `json:"backupTargetPoolRef,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +optional
	Description *string `json:"description,omitempty"`

	// This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool (i.e., not as a backup pool to some other target pool). The value of the field must be in [0, 1]. If set, backupPool must also be set. They together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below this number, traffic arriving at the load-balanced IP will be directed to the backup pool.
	// +optional
	FailoverRatio *float64 `json:"failoverRatio,omitempty"`

	// The URL of the HttpHealthCheck resource. A member instance in this pool is considered healthy if and only if the health checks pass. Only legacy HttpHealthChecks are supported. Only one health check may be specified.
	// +optional
	HealthChecks []TargetpoolHealthChecks `json:"healthChecks,omitempty"`

	// A list of resource URLs to the virtual machine instances serving this pool. They must live in zones contained in the same region as this pool.
	// +optional
	Instances []computev1beta1.InstanceRef `json:"instances,omitempty"`

	// The resource URL for the security policy associated with this target pool.
	// +optional
	SecurityPolicyRef *computev1beta1.ComputeSecurityPolicyRef `json:"securityPolicyRef,omitempty"`

	// Session affinity option, must be one of the following values: NONE: Connections from the same client IP may go to any instance in the pool. CLIENT_IP: Connections from the same client IP will go to the same instance in the pool while that instance remains healthy. CLIENT_IP_PROTO: Connections from the same client IP with the same IP protocol will go to the same instance in the pool while that instance remains healthy.
	// +optional
	SessionAffinity *string `json:"sessionAffinity,omitempty"`
}

// ComputeTargetPoolStatus defines the config connector machine state of ComputeTargetPool
type ComputeTargetPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeTargetPool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeTargetPoolObservedState `json:"observedState,omitempty"`
}

// ComputeTargetPoolObservedState is the state of the ComputeTargetPool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.TargetPool
type ComputeTargetPoolObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +optional
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargetpool;gcpcomputetargetpools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
