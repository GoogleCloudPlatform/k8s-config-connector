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
	certificatemanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	ComputeTargetSSLProxyGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetSSLProxy",
	}
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargetsslproxy;gcpcomputetargetsslproxies
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeTargetSSLProxy is the Schema for the ComputeTargetSSLProxy API
// +k8s:openapi-gen=true
type ComputeTargetSSLProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeTargetSSLProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetSSLProxyStatus `json:"status,omitempty"`
}

// +kcc:spec:proto=google.cloud.compute.v1.TargetSslProxy
type ComputeTargetSSLProxySpec struct {
	// A reference to the ComputeBackendService resource.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.service
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef"`

	// A reference to the CertificateMap resource uri that identifies a
	// certificate map associated with the given target proxy. This field
	// can only be set for global target proxies. Accepted format is
	// '//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}'.
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.certificate_map
	CertificateMapRef *certificatemanagerv1beta1.CertificateManagerCertificateMapRef `json:"certificateMapRef,omitempty"`

	// Immutable. An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.description
	Description *string `json:"description,omitempty"`

	// Location represents the geographical location of the
	// ComputeTargetSSLProxy. Specify "global" for global resources.
	// Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	// +required
	Location string `json:"location"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Specifies the type of proxy header to append before sending data to
	// the backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.proxy_header
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// A list of ComputeSSLCertificate resources that are used to
	// authenticate connections between users and the load balancer.
	// Currently, exactly one SSL certificate must be specified.
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.ssl_certificates
	SslCertificates []ComputeSSLCertificateRef `json:"sslCertificates,omitempty"`

	// A reference to the ComputeSSLPolicy resource that will be
	// associated with the TargetSslProxy resource. If not set, the
	// ComputeTargetSSLProxy resource will not have any SSL policy
	// configured.
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.ssl_policy
	SslPolicyRef *ComputeSSLPolicyRef `json:"sslPolicyRef,omitempty"`
}

// ComputeTargetSSLProxyStatus defines the config connector machine state of ComputeTargetSSLProxy
type ComputeTargetSSLProxyStatus struct {
	// Conditions represent the latest available observations of the object's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeTargetSSLProxy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// The unique identifier for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.id
	ProxyId *int64 `json:"proxyId,omitempty"`

	// The SelfLink for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetSslProxy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeTargetSSLProxyObservedState `json:"observedState,omitempty"`
}

// ComputeTargetSSLProxyObservedState is the state of the ComputeTargetSSLProxy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.TargetSslProxy
type ComputeTargetSSLProxyObservedState struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeTargetSSLProxyList contains a list of ComputeTargetSSLProxy
type ComputeTargetSSLProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetSSLProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetSSLProxy{}, &ComputeTargetSSLProxyList{})
}
