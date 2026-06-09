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

// ----------------------------------------------------------------------------
//

package v1beta1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	vpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vpcaccess/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudFunctionsFunctionGVK = GroupVersion.WithKind("CloudFunctionsFunction")

// +kcc:proto=google.cloud.functions.v1.EventTrigger
type FunctionEventTrigger struct {
	/* Immutable. Required. The type of event to observe. For example:
	`providers/cloud.storage/eventTypes/object.change` and
	`providers/cloud.pubsub/eventTypes/topic.publish`.

	Event types match pattern `providers/* /eventTypes/*.*`.
	The pattern contains:

	1. namespace: For example, `cloud.storage` and
	`google.firebase.analytics`.
	2. resource type: The type of resource on which event occurs. For
	example, the Google Cloud Storage API includes the type `object`.
	3. action: The action that generates the event. For example, action for
	a Google Cloud Storage Object is 'change'.
	These parts are lower case. */
	// +kcc:proto:field=google.cloud.functions.v1.EventTrigger.event_type
	EventType string `json:"eventType"`

	/* Immutable. Specifies policy for failed executions. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.EventTrigger.failure_policy
	FailurePolicy *bool `json:"failurePolicy,omitempty"`

	/* Immutable. */
	// +kcc:proto:field=google.cloud.functions.v1.EventTrigger.resource
	ResourceRef EventTriggerResourceRef `json:"resourceRef"`

	/* Immutable. The hostname of the service that should be observed.

	If no string is provided, the default service implementing the API will
	be used. For example, `storage.googleapis.com` is the default for all
	event types in the `google.storage` namespace. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.EventTrigger.service
	Service *string `json:"service,omitempty"`
}

type EventTriggerResourceRef struct {
	/* Required. The resource(s) from which to observe events, for example,
	`projects/_/buckets/myBucket`.

	Not all syntactically correct values are accepted by all services. For
	example:

	1. The authorization model must support it. Google Cloud Functions
	   only allows EventTriggers to be deployed that observe resources in the
	   same project as the `Function`.
	2. The resource type must match the pattern expected for an
	   `event_type`. For example, an `EventTrigger` that has an
	   `event_type` of "google.pubsub.topic.publish" should have a resource
	   that matches Google Cloud Pub/Sub topics.

	Additionally, some services may support short names when creating an
	`EventTrigger`. These will always be returned in the normalized "long"
	format.

	See each *service's* documentation for supported formats.

	Allowed values:
	* The Google Cloud resource name of a `StorageBucket` resource (format: `{{name}}`).
	* The Google Cloud resource name of a `PubSubTopic` resource (format: `projects/{{project}}/topics/{{name}}`). */
	// +optional
	External string `json:"external,omitempty"`

	/* Kind of the referent. Allowed values: StorageBucket,PubSubTopic */
	Kind string `json:"kind"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	// +optional
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.HttpsTrigger
type FunctionHttpsTrigger struct {
	/* Immutable. Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly. Possible values: SECURITY_LEVEL_UNSPECIFIED, SECURE_ALWAYS, SECURE_OPTIONAL */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.HttpsTrigger.security_level
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.SourceRepository
type FunctionSourceRepository struct {
	/* Immutable. The URL pointing to the hosted repository where the function is defined.
	There are supported Cloud Source Repository URLs in the following
	formats:

	To refer to a specific commit:
	`https://source.developers.google.com/projects/* /repos/* /revisions/* /paths/*`
	To refer to a moveable alias (branch):
	`https://source.developers.google.com/projects/* /repos/* /moveable-aliases/* /paths/*`
	In particular, to refer to HEAD use `master` moveable alias.
	To refer to a specific fixed alias (tag):
	`https://source.developers.google.com/projects/* /repos/* /fixed-aliases/* /paths/*`

	You may omit `paths/*` if you want to use the main directory. */
	// +kcc:proto:field=google.cloud.functions.v1.SourceRepository.url
	Url string `json:"url"`
}

// +kcc:spec:proto=google.cloud.functions.v1.CloudFunction
type CloudFunctionsFunctionSpec struct {
	/* Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.available_memory_mb
	AvailableMemoryMb *int64 `json:"availableMemoryMb,omitempty"`

	/* User-provided description of a function. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.description
	Description *string `json:"description,omitempty"`

	/* Immutable. The name of the function (as defined in source code) that will be
	executed. Defaults to the resource name suffix, if not specified. For
	backward compatibility, if function with given name is not found, then the
	system will try to use function named "function".
	For Node.js this is name of a function exported by the module specified
	in `source_location`. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.entry_point
	EntryPoint *string `json:"entryPoint,omitempty"`

	/* Environment variables that shall be available during function execution. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.environment_variables
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`

	/* Immutable. A source that fires events in response to a condition in another service. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.event_trigger
	EventTrigger *FunctionEventTrigger `json:"eventTrigger,omitempty"`

	/* Immutable. An HTTPS endpoint type of source that can be triggered via URL. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.https_trigger
	HttpsTrigger *FunctionHttpsTrigger `json:"httpsTrigger,omitempty"`

	/* The ingress settings for the function, controlling what traffic can reach
	it. Possible values: INGRESS_SETTINGS_UNSPECIFIED, ALLOW_ALL, ALLOW_INTERNAL_ONLY, ALLOW_INTERNAL_AND_GCLB */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.ingress_settings
	IngressSettings *string `json:"ingressSettings,omitempty"`

	/* The limit on the maximum number of function instances that may coexist at a
	given time. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.max_instances
	MaxInstances *int64 `json:"maxInstances,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef parent.ProjectRef `json:"projectRef"`

	/* Immutable. The name of the Cloud Functions region of the function. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The runtime in which to run the function. Required when deploying a new
	function, optional when updating an existing function. For a complete
	list of possible choices, see the
	[`gcloud` command
	reference](/sdk/gcloud/reference/functions/deploy#--runtime). */
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.runtime
	Runtime string `json:"runtime"`

	/* Immutable. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.service_account_email
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.source_archive_url
	SourceArchiveUrl *string `json:"sourceArchiveUrl,omitempty"`

	/* Immutable. Represents parameters related to source repository where a function is hosted. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.source_repository
	SourceRepository *FunctionSourceRepository `json:"sourceRepository,omitempty"`

	/* The function execution timeout. Execution is considered failed and
	can be terminated if the function is not completed at the end of the
	timeout period. Defaults to 60 seconds. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.timeout
	Timeout *string `json:"timeout,omitempty"`

	/* The egress settings for the connector, controlling what traffic is diverted
	through it. Possible values: VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED, PRIVATE_RANGES_ONLY, ALL_TRAFFIC */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.vpc_connector_egress_settings
	VpcConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.vpc_connector
	VpcConnectorRef *vpcaccessv1beta1.VPCAccessConnectorRef `json:"vpcConnectorRef,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.HttpsTrigger
type FunctionHttpsTriggerStatus struct {
	/* Output only. The deployed url for the function. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.HttpsTrigger.url
	Url *string `json:"url,omitempty"`
}

// +kcc:proto=google.cloud.functions.v1.SourceRepository
type FunctionSourceRepositoryStatus struct {
	/* Output only. The URL pointing to the hosted repository where the function
	were defined at the time of deployment. It always points to a specific
	commit in the format described above. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.SourceRepository.deployed_url
	DeployedUrl *string `json:"deployedUrl,omitempty"`
}

// +kcc:status:proto=google.cloud.functions.v1.CloudFunction
type CloudFunctionsFunctionStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudFunctionsFunction's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.https_trigger
	HttpsTrigger *FunctionHttpsTriggerStatus `json:"httpsTrigger,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.source_repository
	SourceRepository *FunctionSourceRepositoryStatus `json:"sourceRepository,omitempty"`

	/* Output only. Status of the function deployment. Possible values: CLOUD_FUNCTION_STATUS_UNSPECIFIED, ACTIVE, OFFLINE, DEPLOY_IN_PROGRESS, DELETE_IN_PROGRESS, UNKNOWN */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.status
	Status *string `json:"status,omitempty"`

	/* Output only. The last update timestamp of a Cloud Function in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Output only. The version identifier of the Cloud Function. Each deployment attempt
	results in a new version of a function being created. */
	// +optional
	// +kcc:proto:field=google.cloud.functions.v1.CloudFunction.version_id
	VersionId *int64 `json:"versionId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudfunctionsfunction;gcpcloudfunctionsfunctions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudFunctionsFunction is the Schema for the cloudfunctions API
// +k8s:openapi-gen=true
type CloudFunctionsFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudFunctionsFunctionSpec   `json:"spec,omitempty"`
	Status CloudFunctionsFunctionStatus `json:"status,omitempty"`
}

// CloudFunctionsFunctionObservedState is the state of the CloudFunctionsFunction resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.functions.v1.CloudFunction
type CloudFunctionsFunctionObservedState struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudFunctionsFunctionList contains a list of CloudFunctionsFunction
type CloudFunctionsFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudFunctionsFunction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudFunctionsFunction{}, &CloudFunctionsFunctionList{})
}
