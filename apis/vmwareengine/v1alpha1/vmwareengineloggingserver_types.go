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

// API sources for VMwareEngineLoggingServer, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/vmwareengine/v1/vmwareengine_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/solutions/vmware-as-a-service
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/solutions/vmware-as-a-service/docs/reference/rest/v1/projects.locations.privateClouds.loggingServers

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMwareEngineLoggingServerGVK = GroupVersion.WithKind("VMwareEngineLoggingServer")

// VMwareEngineLoggingServerSpec defines the desired state of VMwareEngineLoggingServer
// +kcc:spec:proto=google.cloud.vmwareengine.v1.LoggingServer
// +kcc:required-from-proto
type VMwareEngineLoggingServerSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/privateClouds/{private_cloud} this resource belongs to.
	// +kcc:guess
	// PrivateCloudRef *PrivateCloudRef `json:"privateCloudRef,omitempty"`

	// The VMwareEngineLoggingServer name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Fully-qualified domain name (FQDN) or IP Address of the logging
	//  server.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.hostname
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// Required. Port number at which the logging server receives logs.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.port
	// +required
	Port *int32 `json:"port,omitempty"`

	// Required. Protocol used by vCenter to send logs to a logging server.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.protocol
	// +required
	Protocol *string `json:"protocol,omitempty"`

	// Required. The type of component that produces logs that will be forwarded
	//  to this logging server.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.source_type
	// +required
	SourceType *string `json:"sourceType,omitempty"`
}

// VMwareEngineLoggingServerStatus defines the config connector machine state of VMwareEngineLoggingServer
type VMwareEngineLoggingServerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEngineLoggingServer resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEngineLoggingServerObservedState `json:"observedState,omitempty"`
}

// VMwareEngineLoggingServerObservedState is the state of the VMwareEngineLoggingServer resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmwareengine.v1.LoggingServer
type VMwareEngineLoggingServerObservedState struct {
	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.LoggingServer.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareengineloggingserver;gcpvmwareengineloggingservers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEngineLoggingServer is the Schema for the VMwareEngineLoggingServer API
// +k8s:openapi-gen=true
type VMwareEngineLoggingServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEngineLoggingServerSpec   `json:"spec,omitempty"`
	Status VMwareEngineLoggingServerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEngineLoggingServerList contains a list of VMwareEngineLoggingServer
type VMwareEngineLoggingServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEngineLoggingServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEngineLoggingServer{}, &VMwareEngineLoggingServerList{})
}
