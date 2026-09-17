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

var CloudRunInstanceGVK = GroupVersion.WithKind("CloudRunInstance")

// CloudRunInstanceSpec defines the desired state of CloudRunInstance
// +kcc:spec:proto=google.cloud.run.v2.Instance
type CloudRunInstanceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CloudRunInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User-provided description of the Instance. This field currently has a
	//  512-character limit.
	Description *string `json:"description,omitempty"`

	// Arbitrary identifier for the API client.
	Client *string `json:"client,omitempty"`

	// Arbitrary version identifier for the API client.
	ClientVersion *string `json:"clientVersion,omitempty"`

	// Optional. The launch stage as defined by [Google Cloud Platform
	//   Launch Stages](https://cloud.google.com/terms/launch-stages).
	// +kubebuilder:validation:Enum=LAUNCH_STAGE_UNSPECIFIED;UNIMPLEMENTED;PRELAUNCH;EARLY_ACCESS;ALPHA;BETA;GA;DEPRECATED
	LaunchStage *string `json:"launchStage,omitempty"`

	// Optional. Settings for the Binary Authorization feature.
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Optional. VPC Access configuration to use for this Instance.
	VpcAccess *VPCAccess `json:"vpcAccess,omitempty"`

	// Optional. Email address of the IAM service account associated with the Instance.
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Required. Holds the single container that defines the unit of execution for
	//  this Instance.
	Containers []Container `json:"containers,omitempty"`

	// Optional. A list of Volumes to make available to containers.
	Volumes []Volume `json:"volumes,omitempty"`

	// Optional. A reference to a customer managed encryption key (CMEK) to use to encrypt
	//  this container image.
	EncryptionKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"encryptionKeyRef,omitempty"`

	// Optional. The action to take if the encryption key is revoked.
	EncryptionKeyRevocationAction *string `json:"encryptionKeyRevocationAction,omitempty"`

	// Optional. If encryption_key_revocation_action is SHUTDOWN, the duration
	//  before shutting down all instances.
	EncryptionKeyShutdownDuration *string `json:"encryptionKeyShutdownDuration,omitempty"`

	// Optional. The node selector for the instance.
	NodeSelector *NodeSelector `json:"nodeSelector,omitempty"`

	// Optional. True if GPU zonal redundancy is disabled on this instance.
	GpuZonalRedundancyDisabled *bool `json:"gpuZonalRedundancyDisabled,omitempty"`

	// Optional. Provides the ingress settings for this Instance.
	Ingress *string `json:"ingress,omitempty"`

	// Optional. Disables IAM permission check for run.routes.invoke for callers
	//  of this Instance.
	InvokerIAMDisabled *bool `json:"invokerIAMDisabled,omitempty"`

	// Optional. IAP settings on the Instance.
	IAPEnabled *bool `json:"iapEnabled,omitempty"`
}

// CloudRunInstanceStatus defines the config connector machine state of CloudRunInstance
type CloudRunInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudRunInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudRunInstanceObservedState `json:"observedState,omitempty"`
}

// CloudRunInstanceObservedState is the state of the CloudRunInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.run.v2.Instance
type CloudRunInstanceObservedState struct {
	// Output only. Server assigned unique identifier for the trigger. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	Uid *string `json:"uid,omitempty"`

	// Output only. A number that monotonically increases every time the user
	//  modifies the desired state.
	Generation *int64 `json:"generation,omitempty"`

	// Output only. The creation time.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time. It is only populated as a response to a
	//  Delete request.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. For a deleted resource, the time after which it will be
	//  permanently deleted.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Email address of the authenticated creator.
	Creator *string `json:"creator,omitempty"`

	// Output only. Email address of the last authenticated modifier.
	LastModifier *string `json:"lastModifier,omitempty"`

	// Output only. The generation of this Instance currently serving traffic.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The Google Console URI to obtain logs for the Instance.
	LogUri *string `json:"logUri,omitempty"`

	// Output only. The Condition of this Instance, containing its readiness
	//  status, and detailed error information in case it did not reach a serving
	//  state.
	TerminalCondition *RunCondition `json:"terminalCondition,omitempty"`

	// Output only. The Conditions of all other associated sub-resources. They
	//  contain additional diagnostics information in case the Instance does not
	//  reach its Serving state.
	Conditions []RunCondition `json:"conditions,omitempty"`

	// Output only. Status information for each of the specified containers. The
	//  status includes the resolved digest for specified images.
	ContainerStatuses []ContainerStatus `json:"containerStatuses,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. All URLs serving traffic for this Instance.
	Urls []string `json:"urls,omitempty"`

	// Output only. Returns true if the Instance is currently being acted upon by
	//  the system to bring it into the desired state.
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudruninstance;gcpcloudruninstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudRunInstance is the Schema for the CloudRunInstance API
// +k8s:openapi-gen=true
type CloudRunInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudRunInstanceSpec   `json:"spec,omitempty"`
	Status CloudRunInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudRunInstanceList contains a list of CloudRunInstance
type CloudRunInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudRunInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudRunInstance{}, &CloudRunInstanceList{})
}
