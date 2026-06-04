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
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitoringUptimeCheckConfigGVK = GroupVersion.WithKind("MonitoringUptimeCheckConfig")

type ProjectRef struct {
	/* The `projectID` field of a project, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `Project` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `Project` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type MonitoringGroupRef struct {
	// The group of resources being monitored. Should be only the `[GROUP_ID]`, and not the full-path `projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID]`.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

type UptimeCheckConfig_ContentMatcher struct {
	// +required
	Content *string `json:"content,omitempty"`
	Matcher *string `json:"matcher,omitempty"`
}

type UptimeCheckConfig_HTTPCheck_BasicAuthentication struct {
	// +required
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`
	// +required
	Username *string `json:"username,omitempty"`
}

type UptimeCheckConfig_HTTPCheck struct {
	AuthInfo      *UptimeCheckConfig_HTTPCheck_BasicAuthentication `json:"authInfo,omitempty"`
	Body          *string                                          `json:"body,omitempty"`
	ContentType   *string                                          `json:"contentType,omitempty"`
	Headers       map[string]string                                `json:"headers,omitempty"`
	MaskHeaders   *bool                                            `json:"maskHeaders,omitempty"`
	Path          *string                                          `json:"path,omitempty"`
	Port          *int64                                           `json:"port,omitempty"`
	RequestMethod *string                                          `json:"requestMethod,omitempty"`
	UseSsl        *bool                                            `json:"useSsl,omitempty"`
	ValidateSsl   *bool                                            `json:"validateSsl,omitempty"`
}

type UptimeCheckConfig_MonitoredResource struct {
	// Immutable.
	// +required
	FilterLabels map[string]string `json:"filterLabels,omitempty"`
	// Immutable.
	// +required
	Type *string `json:"type,omitempty"`
}

type UptimeCheckConfig_ResourceGroup struct {
	GroupRef     *MonitoringGroupRef `json:"groupRef,omitempty"`
	ResourceType *string             `json:"resourceType,omitempty"`
}

type UptimeCheckConfig_TCPCheck struct {
	// +required
	Port *int64 `json:"port,omitempty"`
}

// MonitoringUptimeCheckConfigSpec defines the desired state of MonitoringUptimeCheckConfig
type MonitoringUptimeCheckConfigSpec struct {
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// +required
	ProjectRef *ProjectRef `json:"projectRef,omitempty"`

	// +required
	Timeout *string `json:"timeout,omitempty"`

	ContentMatchers []UptimeCheckConfig_ContentMatcher `json:"contentMatchers,omitempty"`

	HTTPCheck *UptimeCheckConfig_HTTPCheck `json:"httpCheck,omitempty"`

	MonitoredResource *UptimeCheckConfig_MonitoredResource `json:"monitoredResource,omitempty"`

	Period *string `json:"period,omitempty"`

	ResourceGroup *UptimeCheckConfig_ResourceGroup `json:"resourceGroup,omitempty"`

	ResourceID *string `json:"resourceID,omitempty"`

	SelectedRegions []string `json:"selectedRegions,omitempty"`

	TCPCheck *UptimeCheckConfig_TCPCheck `json:"tcpCheck,omitempty"`
}

// MonitoringUptimeCheckConfigStatus defines the config connector machine state of MonitoringUptimeCheckConfig
type MonitoringUptimeCheckConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringuptimecheckconfig;gcpmonitoringuptimecheckconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringUptimeCheckConfig is the Schema for the MonitoringUptimeCheckConfig API
// +k8s:openapi-gen=true
type MonitoringUptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MonitoringUptimeCheckConfigSpec   `json:"spec,omitempty"`
	Status MonitoringUptimeCheckConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MonitoringUptimeCheckConfigList contains a list of MonitoringUptimeCheckConfig
type MonitoringUptimeCheckConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringUptimeCheckConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringUptimeCheckConfig{}, &MonitoringUptimeCheckConfigList{})
}
