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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
)

var WorkstationClusterGVK = GroupVersion.WithKind("WorkstationCluster")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkstationClusterSpec defines the desired state of WorkstationCluster
// +kcc:spec:proto=google.cloud.workstations.v1.WorkstationCluster
type WorkstationClusterSpec struct {
	// Immutable. The Project that this resource belongs to.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// The location of the cluster.
	Location string `json:"location,omitempty"`

	// The WorkstationCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Human-readable name for this workstation cluster.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Client-specified annotations.
	Annotations []WorkstationAnnotation `json:"annotations,omitempty"`

	// Optional.
	//  [Labels](https://cloud.google.com/workstations/docs/label-resources) that
	//  are applied to the workstation cluster and that are also propagated to the
	//  underlying Compute Engine resources.
	Labels []WorkstationLabel `json:"labels,omitempty"`

	// Immutable. Reference to the Compute Engine network in which instances associated
	//  with this workstation cluster will be created.
	// +required
	NetworkRef refs.ComputeNetworkRef `json:"networkRef"`

	// Immutable. Reference to the Compute Engine subnetwork in which instances
	//  associated with this workstation cluster will be created. Must be part of
	//  the subnetwork specified for this workstation cluster.
	// +required
	SubnetworkRef refs.ComputeSubnetworkRef `json:"subnetworkRef"`

	// Optional. Configuration for private workstation cluster.
	PrivateClusterConfig *WorkstationCluster_PrivateClusterConfig `json:"privateClusterConfig,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationCluster.PrivateClusterConfig
type WorkstationCluster_PrivateClusterConfig struct {
	// Immutable. Whether Workstations endpoint is private.
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty"`

	// Optional. Additional projects that are allowed to attach to the
	//  workstation cluster's service attachment. By default, the workstation
	//  cluster's project and the VPC host project (if different) are allowed.
	AllowedProjects []refs.ProjectRef `json:"allowedProjects,omitempty"`
}

// WorkstationClusterStatus defines the config connector machine state of WorkstationCluster
type WorkstationClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the WorkstationCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *WorkstationClusterObservedState `json:"observedState,omitempty"`
}

// WorkstationClusterSpec defines the desired state of WorkstationCluster
// +kcc:observedstate:proto=google.cloud.workstations.v1.WorkstationCluster
type WorkstationClusterObservedState struct {
	// Output only. A system-assigned unique identifier for this workstation
	//  cluster.
	Uid *string `json:"uid,omitempty"`

	// Output only. Indicates whether this workstation cluster is currently being
	//  updated to match its intended state.
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Time when this workstation cluster was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this workstation cluster was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time when this workstation cluster was soft-deleted.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Optional. Checksum computed by the server. May be sent on update and delete
	//  requests to make sure that the client has an up-to-date value before
	//  proceeding.
	Etag *string `json:"etag,omitempty"`

	// Output only. The private IP address of the control plane for this
	//  workstation cluster. Workstation VMs need access to this IP address to work
	//  with the service, so make sure that your firewall rules allow egress from
	//  the workstation VMs to this address.
	ControlPlaneIP *string `json:"controlPlaneIP,omitempty"`

	// Output only. Hostname for the workstation cluster. This field will be
	//  populated only when private endpoint is enabled. To access workstations
	//  in the workstation cluster, create a new DNS zone mapping this domain
	//  name to an internal IP address and a forwarding rule mapping that address
	//  to the service attachment.
	ClusterHostname *string `json:"clusterHostname,omitempty"`

	// Output only. Service attachment URI for the workstation cluster. The
	//  service attachment is created when private endpoint is enabled. To access
	//  workstations in the workstation cluster, configure access to the managed
	//  service using [Private Service
	//  Connect](https://cloud.google.com/vpc/docs/configure-private-service-connect-services).
	ServiceAttachmentURI *string `json:"serviceAttachmentUri,omitempty"`

	// Output only. Whether this workstation cluster is in degraded mode, in which
	//  case it may require user action to restore full functionality. Details can
	//  be found in
	//  [conditions][google.cloud.workstations.v1.WorkstationCluster.conditions].
	Degraded *bool `json:"degraded,omitempty"`

	// Output only. Status conditions describing the workstation cluster's current
	//  state.
	GCPConditions []WorkstationServiceGCPCondition `json:"gcpConditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpworkstationcluster;gcpworkstationclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkstationCluster is the Schema for the WorkstationCluster API
// +k8s:openapi-gen=true
type WorkstationCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkstationClusterSpec   `json:"spec,omitempty"`
	Status WorkstationClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WorkstationClusterList contains a list of WorkstationCluster
type WorkstationClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkstationCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkstationCluster{}, &WorkstationClusterList{})
}
