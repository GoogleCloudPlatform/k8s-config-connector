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

package v1alpha2

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIWorkbenchInstanceGVK = GroupVersion.WithKind("VertexAIWorkbenchInstance")

// VertexAIWorkbenchInstanceSpec defines the desired state of VertexAIWorkbenchInstance
// +kcc:spec:proto=google.cloud.notebooks.v2.Instance
type VertexAIWorkbenchInstanceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAIWorkbenchInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Compute Engine setup for the notebook. Uses notebook-defined
	//  fields.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.gce_setup
	GCESetup *GCESetup `json:"gceSetup,omitempty"`

	// Optional. Input only. The owner of this instance after creation. Format:
	//  `alias@example.com`
	//
	//  Currently supports one owner only. If not specified, all of the service
	//  account users of your VM instance's service account can use
	//  the instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.instance_owners
	InstanceOwners []string `json:"instanceOwners,omitempty"`

	// Optional. If true, the notebook instance will not register with the proxy.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.disable_proxy_access
	DisableProxyAccess *bool `json:"disableProxyAccess,omitempty"`

	// Optional. Labels to apply to this instance.
	//  These can be later modified by the UpdateInstance method.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// VertexAIWorkbenchInstanceStatus defines the config connector machine state of VertexAIWorkbenchInstance
type VertexAIWorkbenchInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIWorkbenchInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIWorkbenchInstanceObservedState `json:"observedState,omitempty"`
}

// VertexAIWorkbenchInstanceObservedState is the state of the VertexAIWorkbenchInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.notebooks.v2.Instance
type VertexAIWorkbenchInstanceObservedState struct {
	// Output only. The name of this notebook instance. Format:
	//  `projects/{project_id}/locations/{location}/instances/{instance_id}`
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.name
	Name *string `json:"name,omitempty"`

	// Optional. Compute Engine setup for the notebook. Uses notebook-defined
	//  fields.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.gce_setup
	GCESetup *GCESetupObservedState `json:"gceSetup,omitempty"`

	// Output only. The proxy endpoint that is used to access the Jupyter
	//  notebook.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.proxy_uri
	ProxyURI *string `json:"proxyURI,omitempty"`

	// Output only. Email address of entity that sent original CreateInstance
	//  request.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. The state of this instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. The upgrade history of this instance.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.upgrade_history
	UpgradeHistory []UpgradeHistoryEntry `json:"upgradeHistory,omitempty"`

	// Output only. Unique ID of the resource.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.id
	ID *string `json:"id,omitempty"`

	// Output only. Instance health_state.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.health_state
	HealthState *string `json:"healthState,omitempty"`

	// Output only. Additional information about instance health.
	//  Example:
	//
	//      healthInfo": {
	//        "docker_proxy_agent_status": "1",
	//        "docker_status": "1",
	//        "jupyterlab_api_status": "-1",
	//        "jupyterlab_status": "-1",
	//        "updated": "2020-10-18 09:40:03.573409"
	//      }
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.health_info
	HealthInfo map[string]string `json:"healthInfo,omitempty"`

	// Output only. Instance creation time.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Instance update time.
	// +kcc:proto:field=google.cloud.notebooks.v2.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiworkbenchinstance;gcpvertexaiworkbenchinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIWorkbenchInstance is the Schema for the VertexAIWorkbenchInstance API
// +k8s:openapi-gen=true
type VertexAIWorkbenchInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIWorkbenchInstanceSpec   `json:"spec,omitempty"`
	Status VertexAIWorkbenchInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIWorkbenchInstanceList contains a list of VertexAIWorkbenchInstance
type VertexAIWorkbenchInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIWorkbenchInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIWorkbenchInstance{}, &VertexAIWorkbenchInstanceList{})
}
