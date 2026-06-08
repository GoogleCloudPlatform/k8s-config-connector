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

var ApigeeRegistryInstanceGVK = GroupVersion.WithKind("ApigeeRegistryInstance")

// ApigeeRegistryInstanceSpec defines the desired state of ApigeeRegistryInstance
// +kcc:spec:proto=google.cloud.apigeeregistry.v1.Instance
type ApigeeRegistryInstanceSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The ApigeeRegistryInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Config of the Instance.
	// +required
	Config *Instance_Config `json:"config"`
}

// +kcc:proto=google.cloud.apigeeregistry.v1.Instance.Config
type Instance_Config struct {
	// Required. The Customer Managed Encryption Key (CMEK) used for data encryption.
	// The CMEK name should follow the format of
	// `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`,
	// where the `location` must match InstanceConfig.location.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.Config.cmek_key_name
	// +required
	CmekKeyNameRef *refsv1beta1.KMSCryptoKeyRef `json:"cmekKeyNameRef"`
}

// +kcc:observedstate:proto=google.cloud.apigeeregistry.v1.Instance.Config
type Instance_ConfigObservedState struct {
	// Output only. The GCP location where the Instance resides.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.Instance.Config.location
	Location *string `json:"location,omitempty"`
}

// ApigeeRegistryInstanceStatus defines the config connector machine state of ApigeeRegistryInstance
type ApigeeRegistryInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeRegistryInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeRegistryInstanceObservedState `json:"observedState,omitempty"`
}

// ApigeeRegistryInstanceObservedState is the state of the ApigeeRegistryInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apigeeregistry.v1.Instance
type ApigeeRegistryInstanceObservedState struct {
	// Output only. Creation timestamp.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of the Instance.
	State *string `json:"state,omitempty"`

	// Output only. Extra information of Instance.State if the state is `FAILED`.
	StateMessage *string `json:"stateMessage,omitempty"`

	// Required. Config of the Instance.
	Config *Instance_ConfigObservedState `json:"config,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeregistryinstance;gcpapigeeregistryinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeRegistryInstance is the Schema for the ApigeeRegistryInstance API
// +k8s:openapi-gen=true
type ApigeeRegistryInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeRegistryInstanceSpec   `json:"spec,omitempty"`
	Status ApigeeRegistryInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeRegistryInstanceList contains a list of ApigeeRegistryInstance
type ApigeeRegistryInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeRegistryInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeRegistryInstance{}, &ApigeeRegistryInstanceList{})
}
