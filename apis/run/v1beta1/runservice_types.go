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
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/default-controller=direct"
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

// +kcc:proto=google.cloud.run.v2.RevisionTemplate
type RevisionTemplate struct {
	// Optional. The unique name for the revision. If this field is omitted, it
	//  will be automatically generated based on the Service name.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.revision
	Revision *string `json:"revision,omitempty"`

	// Optional. Unstructured key value map that can be used to organize and
	//  categorize objects. User-provided labels are shared with Google's billing
	//  system, so they can be used to filter, or break down billing charges by
	//  team, component, environment, state, etc. For more information, visit
	//  https://cloud.google.com/resource-manager/docs/creating-managing-labels or
	//  https://cloud.google.com/run/docs/configuring/labels.
	//
	//  <p>Cloud Run API v2 does not support labels with `run.googleapis.com`,
	//  `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev`
	//  namespaces, and they will be rejected. All system labels in v1 now have a
	//  corresponding field in v2 RevisionTemplate.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Unstructured key value map that may be set by external tools to
	//  store and arbitrary metadata. They are not queryable and should be
	//  preserved when modifying objects.
	//
	//  <p>Cloud Run API v2 does not support annotations with `run.googleapis.com`,
	//  `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev`
	//  namespaces, and they will be rejected. All system annotations in v1 now
	//  have a corresponding field in v2 RevisionTemplate.
	//
	//  <p>This field follows Kubernetes annotations' namespacing, limits, and
	//  rules.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Scaling settings for this Revision.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.scaling
	Scaling *RevisionScaling `json:"scaling,omitempty"`

	// Optional. VPC Access configuration to use for this Revision. For more
	//  information, visit
	//  https://cloud.google.com/run/docs/configuring/connecting-vpc.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.vpc_access
	VPCAccess *VPCAccess `json:"vpcAccess,omitempty"`

	// Optional. Max allowed time for an instance to respond to a request.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Optional. Email address of the IAM service account associated with the
	//  revision of the service. The service account represents the identity of the
	//  running revision, and determines what permissions the revision has. If not
	//  provided, the revision will use the project's default service account.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.service_account
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Holds the single container that defines the unit of execution for this
	//  Revision.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.containers
	Containers []Container `json:"containers,omitempty"`

	// Optional. A list of Volumes to make available to containers.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.volumes
	Volumes []Volume `json:"volumes,omitempty"`

	// Optional. The sandbox environment to host this Revision.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.execution_environment
	ExecutionEnvironment *string `json:"executionEnvironment,omitempty"`

	// A reference to a customer managed encryption key (CMEK) to use to encrypt
	//  this container image. For more information, go to
	//  https://cloud.google.com/run/docs/securing/using-cmek
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.encryption_key
	EncryptionKeyRef *refs.KMSCryptoKeyRef `json:"encryptionKeyRef,omitempty"`

	// Optional. Sets the maximum number of requests that each serving instance
	//  can receive. If not specified or 0, concurrency defaults to 80 when
	//  requested `CPU >= 1` and defaults to 1 when requested `CPU < 1`.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.max_instance_request_concurrency
	MaxInstanceRequestConcurrency *int32 `json:"maxInstanceRequestConcurrency,omitempty"`

	// Optional. Enables service mesh connectivity.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.service_mesh
	ServiceMesh *ServiceMesh `json:"serviceMesh,omitempty"`

	// Optional. The action to take if the encryption key is revoked.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.encryption_key_revocation_action
	EncryptionKeyRevocationAction *string `json:"encryptionKeyRevocationAction,omitempty"`

	// Optional. If encryption_key_revocation_action is SHUTDOWN, the duration
	//  before shutting down all instances. The minimum increment is 1 hour.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.encryption_key_shutdown_duration
	EncryptionKeyShutdownDuration *string `json:"encryptionKeyShutdownDuration,omitempty"`

	// Optional. Enable session affinity.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.session_affinity
	SessionAffinity *bool `json:"sessionAffinity,omitempty"`

	// Optional. Disables health checking containers during deployment.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.health_check_disabled
	HealthCheckDisabled *bool `json:"healthCheckDisabled,omitempty"`

	// Optional. The node selector for the revision template.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.node_selector
	NodeSelector *NodeSelector `json:"nodeSelector,omitempty"`

	// Optional. True if GPU zonal redundancy is disabled on this revision.
	// +kcc:proto:field=google.cloud.run.v2.RevisionTemplate.gpu_zonal_redundancy_disabled
	GpuZonalRedundancyDisabled *bool `json:"gpuZonalRedundancyDisabled,omitempty"`
}
