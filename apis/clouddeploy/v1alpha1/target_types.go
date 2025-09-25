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

package v1alpha1

import (
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDeployTargetGVK = GroupVersion.WithKind("CloudDeployTarget")

// CloudDeployTargetSpec defines the desired state of CloudDeployTarget
// +kcc:spec:proto=google.cloud.deploy.v1.Target
type CloudDeployTargetSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// Immutable. The location where the Target should reside.
	// +required
	Location *string `json:"location,omitempty"`

	// Optional. Description of the `Target`. Max length is 255 characters.
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	// the user, and not by Cloud Deploy. See
	// https://google.aip.dev/128#annotations for more details such as format and
	// size limitations.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Labels are attributes that can be set and used by both the
	// user and by Cloud Deploy. Labels must meet the following constraints:
	//
	// * Keys and values can contain only lowercase letters, numeric characters,
	// underscores, and dashes.
	// * All characters must use UTF-8 encoding, and international characters are
	// allowed.
	// * Keys must start with a lowercase letter or international character.
	// * Each resource is limited to a maximum of 64 labels.
	//
	// Both keys and values are additionally constrained to be <= 128 bytes.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Whether or not the `Target` requires approval.
	RequireApproval *bool `json:"requireApproval,omitempty"`

	// Information specifying a GKE Cluster.
	Gke *GKECluster `json:"gke,omitempty"`

	// Information specifying an Anthos Cluster.
	AnthosCluster *AnthosCluster `json:"anthosCluster,omitempty"`

	// Information specifying a Cloud Run deployment target.
	Run *CloudRunLocation `json:"run,omitempty"`

	// Information specifying a multi-target.
	MultiTarget *MultiTarget `json:"multiTarget,omitempty"`

	// Information specifying a custom target.
	CustomTarget *CustomTarget `json:"customTarget,omitempty"`

	// Optional. The associations between the Target and other entities.
	AssociatedEntities map[string]*AssociatedEntities `json:"associatedEntities,omitempty"`

	// Optional. Configurations for all execution that relates to this `Target`.
	// Each `ExecutionEnvironmentUsage` value may only be used in a single
	// configuration; using the same value multiple times is an error.
	// When one or more configurations are specified, they must include the
	// `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values.
	// When no configurations are specified, execution will use the default
	// specified in `DefaultPool`.
	ExecutionConfigs []*ExecutionConfig `json:"executionConfigs,omitempty"`

	// Optional. The deploy parameters to use for this target.
	DeployParameters map[string]string `json:"deployParameters,omitempty"`
}

// CloudDeployTargetStatus defines the config connector machine state of CloudDeployTarget
type CloudDeployTargetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDeployTarget resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDeployTargetObservedState `json:"observedState,omitempty"`
}

// CloudDeployTargetObservedState is the state of the CloudDeployTarget resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.Target
type CloudDeployTargetObservedState struct {
	// Output only. Resource id of the `Target`.
	TargetId *string `json:"targetId,omitempty"`

	// Output only. Unique identifier of the `Target`.
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the `Target` was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the `Target` was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	// other fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddeploytarget;gcpclouddeploytargets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDeployTarget is the Schema for the CloudDeployTarget API
// +k8s:openapi-gen=true
type CloudDeployTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDeployTargetSpec   `json:"spec,omitempty"`
	Status CloudDeployTargetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployTargetList contains a list of CloudDeployTarget
type CloudDeployTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployTarget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployTarget{}, &CloudDeployTargetList{})
}

// +kcc:proto=google.cloud.deploy.v1.GkeCluster
type GKECluster struct {
	// Optional. Information specifying a GKE Cluster. Format is
	//  `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}`.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Optional. If true, `cluster` is accessed using the private IP address of
	//  the control plane endpoint. Otherwise, the default IP address of the
	//  control plane endpoint is used. The default IP address is the private IP
	//  address for clusters with private control-plane endpoints and the public IP
	//  address otherwise.
	//
	//  Only specify this option when `cluster` is a [private GKE
	//  cluster](https://cloud.google.com/kubernetes-engine/docs/concepts/private-cluster-concept).
	//  Note that `internal_ip` and `dns_endpoint` cannot both be set to true.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.internal_ip
	InternalIP *bool `json:"internalIP,omitempty"`

	// Optional. If set, used to configure a
	//  [proxy](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/#proxy)
	//  to the Kubernetes server.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.proxy_url
	ProxyURL *string `json:"proxyURL,omitempty"`

	// Optional. If set, the cluster will be accessed using the DNS endpoint. Note
	//  that both `dns_endpoint` and `internal_ip` cannot be set to true.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.dns_endpoint
	DNSEndpoint *bool `json:"dnsEndpoint,omitempty"`
}
