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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RunServiceGVK = GroupVersion.WithKind("RunService")

// RunServiceSpec defines the desired state of RunService
// +kcc:spec:proto=google.cloud.run.v2.Service
type RunServiceSpec struct {
	// The location of the cloud run service
	Location *string `json:"location,omitempty"`

	// The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// The RunService name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-provided annotations, which are stored in GCP.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Settings for Binary Authorization feature.
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Optional. Arbitrary identifier for the API client.
	Client *string `json:"client,omitempty"`

	// Optional. Arbitrary version identifier for the API client.
	ClientVersion *string `json:"clientVersion,omitempty"`

	// Optional. User-provided description of the Service.
	Description *string `json:"description,omitempty"`

	// Optional. Provides the ingress settings for this Service.
	Ingress *string `json:"ingress,omitempty"`

	// Optional. The launch stage of the service.
	LaunchStage *string `json:"launchStage,omitempty"`

	// Required. The template used to create revisions for this Service.
	Template *RevisionTemplate `json:"template"`

	// Optional. Specifies how to distribute traffic over a collection of Revisions belonging to the Service.
	Traffic []TrafficTarget `json:"traffic,omitempty"`

	// Optional. Specifies service-level scaling settings
	Scaling *ServiceScaling `json:"scaling,omitempty"`

	// Optional. Disables IAM permission check for run.routes.invoke for callers of this service.
	InvokerIAMDisabled *bool `json:"invokerIAMDisabled,omitempty"`

	// Optional. Disables public resolution of the default URI of this service.
	DefaultURIDisabled *bool `json:"defaultURIDisabled,omitempty"`

	// Optional. One or more custom audiences that you want this service to support.
	CustomAudiences []string `json:"customAudiences,omitempty"`

	// Optional. Configuration for building a Cloud Run function.
	BuildConfig *BuildConfig `json:"buildConfig,omitempty"`
}

// RunServiceStatus defines the config connector machine state of RunService
type RunServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// LastModifiedCookie contains hashes of the last applied spec and the last observed GCP state.
	LastModifiedCookie *string `json:"lastModifiedCookie,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RunService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RunServiceObservedState `json:",inline"`
}

// +kcc:spec:proto=google.cloud.run.v2.Service
type RunServiceObservedState struct {
	// Output only. Server assigned unique identifier for the trigger.
	Uid *string `json:"uid,omitempty"`

	// Output only. The creation time.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. For a deleted resource, the time after which it will be permanently deleted.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Email address of the authenticated creator.
	Creator *string `json:"creator,omitempty"`

	// Output only. Email address of the last authenticated modifier.
	LastModifier *string `json:"lastModifier,omitempty"`

	// Output only. All URLs serving traffic for this Service.
	Urls []string `json:"urls,omitempty"`

	// Output only. The Condition of this Service, containing its readiness status.
	TerminalCondition []*Condition `json:"terminalCondition,omitempty"`

	// Output only. Name of the latest revision that is serving traffic.
	LatestReadyRevision *string `json:"latestReadyRevision,omitempty"`

	// Output only. Name of the last created revision.
	LatestCreatedRevision *string `json:"latestCreatedRevision,omitempty"`

	// Output only. Detailed status information for corresponding traffic targets.
	TrafficStatuses []TrafficTargetStatus `json:"trafficStatuses,omitempty"`

	// Output only. The main URI in which this Service is serving traffic.
	Uri *string `json:"uri,omitempty"`

	// Output only. Returns true if the Service is currently being acted upon by the system to bring it into the desired state.
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. A system-generated fingerprint for this version of the resource.
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcprunservice;gcprunservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RunService is the Schema for the RunService API
// +k8s:openapi-gen=true
type RunService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RunServiceSpec   `json:"spec,omitempty"`
	Status RunServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RunServiceList contains a list of RunService
type RunServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RunService{}, &RunServiceList{})
}
