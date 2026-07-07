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
	ProjectRef                    *refsv1beta1.ProjectRef           `json:"projectRef"`
	Location                      *string                           `json:"location,omitempty"`
	ResourceID                    *string                           `json:"resourceID,omitempty"`
	Description                   *string                           `json:"description,omitempty"`
	Labels                        map[string]string                 `json:"labels,omitempty"`
	Annotations                   map[string]string                 `json:"annotations,omitempty"`
	Client                        *string                           `json:"client,omitempty"`
	ClientVersion                 *string                           `json:"clientVersion,omitempty"`
	LaunchStage                   *string                           `json:"launchStage,omitempty"`
	BinaryAuthorization           *BinaryAuthorization              `json:"binaryAuthorization,omitempty"`
	VpcAccess                     *VPCAccess                        `json:"vpcAccess,omitempty"`
	ServiceAccountRef             *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
	Containers                    []Container                       `json:"containers"`
	Volumes                       []Volume                          `json:"volumes,omitempty"`
	EncryptionKeyRef              *refsv1beta1.KMSCryptoKeyRef      `json:"encryptionKeyRef,omitempty"`
	EncryptionKeyRevocationAction *string                           `json:"encryptionKeyRevocationAction,omitempty"`
	EncryptionKeyShutdownDuration *string                           `json:"encryptionKeyShutdownDuration,omitempty"`
	NodeSelector                  *NodeSelector                     `json:"nodeSelector,omitempty"`
	GpuZonalRedundancyDisabled    *bool                             `json:"gpuZonalRedundancyDisabled,omitempty"`
	Ingress                       *string                           `json:"ingress,omitempty"`
	InvokerIAMDisabled            *bool                             `json:"invokerIAMDisabled,omitempty"`
	IapEnabled                    *bool                             `json:"iapEnabled,omitempty"`
}

// +kcc:spec:proto=google.cloud.run.v2.Instance
type CloudRunInstanceObservedState struct {
	Uid                *string           `json:"uid,omitempty"`
	Generation         *int64            `json:"generation,omitempty"`
	CreateTime         *string           `json:"createTime,omitempty"`
	UpdateTime         *string           `json:"updateTime,omitempty"`
	DeleteTime         *string           `json:"deleteTime,omitempty"`
	ExpireTime         *string           `json:"expireTime,omitempty"`
	Creator            *string           `json:"creator,omitempty"`
	LastModifier       *string           `json:"lastModifier,omitempty"`
	ObservedGeneration *int64            `json:"observedGeneration,omitempty"`
	LogURI             *string           `json:"logURI,omitempty"`
	TerminalCondition  *RunCondition     `json:"terminalCondition,omitempty"`
	ContainerStatuses  []ContainerStatus `json:"containerStatuses,omitempty"`
	SatisfiesPzs       *bool             `json:"satisfiesPzs,omitempty"`
	Urls               []string          `json:"urls,omitempty"`
	Reconciling        *bool             `json:"reconciling,omitempty"`
	Etag               *string           `json:"etag,omitempty"`
}

// CloudRunInstanceStatus defines the config connector machine state of CloudRunInstance
type CloudRunInstanceStatus struct {
	Conditions         []v1alpha1.Condition           `json:"conditions,omitempty"`
	ObservedGeneration *int64                         `json:"observedGeneration,omitempty"`
	ExternalRef        *string                        `json:"externalRef,omitempty"`
	ObservedState      *CloudRunInstanceObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudruninstance;gcpcloudruninstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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

// +kcc:proto=google.cloud.run.v2.ContainerStatus
type ContainerStatus struct {
	Name        *string `json:"name,omitempty"`
	ImageDigest *string `json:"imageDigest,omitempty"`
}
