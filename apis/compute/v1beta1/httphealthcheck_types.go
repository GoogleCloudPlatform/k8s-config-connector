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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeHTTPHealthCheckGVK = GroupVersion.WithKind("ComputeHTTPHealthCheck")

// ComputeHTTPHealthCheckSpec defines the desired state of ComputeHTTPHealthCheck
// +kcc:spec:proto=google.cloud.compute.v1.HTTPHealthCheck
type ComputeHTTPHealthCheckSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The ComputeHTTPHealthCheck name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	// +optional
	CheckIntervalSec *int64 `json:"checkIntervalSec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +optional
	Description *string `json:"description,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	// +optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	// +optional
	Host *string `json:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	// +optional
	Port *int64 `json:"port,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	// +optional
	RequestPath *string `json:"requestPath,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds. It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	// +optional
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	// +optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty"`
}

// ComputeHTTPHealthCheckStatus defines the config connector machine state of ComputeHTTPHealthCheck
type ComputeHTTPHealthCheckStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeHTTPHealthCheck resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *ComputeHTTPHealthCheckObservedState `json:"observedState,omitempty"`
}

// ComputeHTTPHealthCheckObservedState is the state of the ComputeHTTPHealthCheck resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.HTTPHealthCheck
type ComputeHTTPHealthCheckObservedState struct {
	// Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Server-defined URL for the resource.
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputehttphealthcheck;gcpcomputehttphealthchecks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeHTTPHealthCheck is the Schema for the ComputeHTTPHealthCheck API
// +k8s:openapi-gen=true
type ComputeHTTPHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeHTTPHealthCheckSpec   `json:"spec,omitempty"`
	Status ComputeHTTPHealthCheckStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeHTTPHealthCheckList contains a list of ComputeHTTPHealthCheck
type ComputeHTTPHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeHTTPHealthCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeHTTPHealthCheck{}, &ComputeHTTPHealthCheckList{})
}
