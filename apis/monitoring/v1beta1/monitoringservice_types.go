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

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitoringServiceGVK = GroupVersion.WithKind("MonitoringService")

// MonitoringServiceSpec defines the desired state of MonitoringService
// +kcc:proto=google.monitoring.v3.Service
type MonitoringServiceSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	// +required
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// The MonitoringService name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Name used for UI elements listing this Service.
	DisplayName *string `json:"displayName,omitempty"`

	// Configuration for how to query telemetry on a Service.
	Telemetry *Service_Telemetry `json:"telemetry,omitempty"`

	/* NOTYET
	// Identifier. Resource name for this Service. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]
	Name *string `json:"name,omitempty"`


	// Custom service type.
	Custom *Service_Custom `json:"custom,omitempty"`

	// Type used for App Engine services.
	AppEngine *Service_AppEngine `json:"appEngine,omitempty"`

	// Type used for Cloud Endpoints services.
	CloudEndpoints *Service_CloudEndpoints `json:"cloudEndpoints,omitempty"`

	// Type used for Istio services that live in a Kubernetes cluster.
	ClusterIstio *Service_ClusterIstio `json:"clusterIstio,omitempty"`

	// Type used for Istio services scoped to an Istio mesh.
	MeshIstio *Service_MeshIstio `json:"meshIstio,omitempty"`

	// Type used for canonical services scoped to an Istio mesh.
	//  Metrics for Istio are
	//  [documented here](https://istio.io/latest/docs/reference/config/metrics/)
	IstioCanonicalService *Service_IstioCanonicalService `json:"istioCanonicalService,omitempty"`

	// Type used for Cloud Run services.
	CloudRun *Service_CloudRun `json:"cloudRun,omitempty"`

	// Type used for GKE Namespaces.
	GkeNamespace *Service_GkeNamespace `json:"gkeNamespace,omitempty"`

	// Type used for GKE Workloads.
	GkeWorkload *Service_GkeWorkload `json:"gkeWorkload,omitempty"`

	// Type used for GKE Services (the Kubernetes concept of a service).
	GkeService *Service_GkeService `json:"gkeService,omitempty"`

	// Message that contains the service type and service labels of this service
	//  if it is a basic service.
	//  Documentation and examples
	//  [here](https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli).
	BasicService *Service_BasicService `json:"basicService,omitempty"`


	// Labels which have been used to annotate the service. Label keys must start
	//  with a letter. Label keys and values may contain lowercase letters,
	//  numbers, underscores, and dashes. Label keys and values have a maximum
	//  length of 63 characters, and must be less than 128 bytes in size. Up to 64
	//  label entries may be stored. For labels which do not have a semantic value,
	//  the empty string may be supplied for the label value.
	UserLabels map[string]string `json:"userLabels,omitempty"`
	*/
}

// MonitoringServiceStatus defines the config connector machine state of MonitoringService
type MonitoringServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET
	// A unique specifier for the MonitoringService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MonitoringServiceObservedState `json:"observedState,omitempty"`
	*/
}

// MonitoringServiceSpec defines the desired state of MonitoringService
// +kcc:proto=google.monitoring.v3.Service
type MonitoringServiceObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringservice;gcpmonitoringservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringService is the Schema for the MonitoringService API
// +k8s:openapi-gen=true
type MonitoringService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec MonitoringServiceSpec `json:"spec,omitempty"`

	Status MonitoringServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MonitoringServiceList contains a list of MonitoringService
type MonitoringServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringService{}, &MonitoringServiceList{})
}
