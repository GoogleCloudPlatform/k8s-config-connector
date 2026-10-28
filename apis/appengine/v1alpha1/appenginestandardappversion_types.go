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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AppEngineStandardAppVersionGVK = GroupVersion.WithKind("AppEngineStandardAppVersion")

// +kcc:spec:proto=google.appengine.v1.Version
type AppEngineStandardAppVersionSpec struct {
	// Allows App Engine second generation runtimes to access the legacy bundled services.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.app_engine_apis
	AppEngineApis *bool `json:"appEngineApis,omitempty"`

	// Automatic scaling is based on request rate, response latencies, and other application metrics.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.automatic_scaling
	AutomaticScaling *AutomaticScaling `json:"automaticScaling,omitempty"`

	// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.basic_scaling
	BasicScaling *BasicScaling `json:"basicScaling,omitempty"`

	// If set to 'true', the service will be deleted if it is the last version.
	// +optional
	DeleteServiceOnDestroy *bool `json:"deleteServiceOnDestroy,omitempty"`

	// Code and application artifacts that make up this version.
	// +required
	// +kcc:proto:field=google.appengine.v1.Version.deployment
	Deployment *Deployment `json:"deployment"`

	// The entrypoint for the application.
	// +required
	// +kcc:proto:field=google.appengine.v1.Version.entrypoint
	Entrypoint *Entrypoint `json:"entrypoint"`

	// Environment variables available to the application.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.env_variables
	EnvVariables map[string]string `json:"envVariables,omitempty"`

	// An ordered list of URL-matching patterns that should be applied to incoming requests.
	// The first matching URL handles the request and other request handlers are not attempted.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.handlers
	Handlers []URLMap `json:"handlers,omitempty"`

	// A list of the types of messages that this application is able to receive. Possible values: ["INBOUND_SERVICE_MAIL", "INBOUND_SERVICE_MAIL_BOUNCE", "INBOUND_SERVICE_XMPP_ERROR", "INBOUND_SERVICE_XMPP_MESSAGE", "INBOUND_SERVICE_XMPP_SUBSCRIBE", "INBOUND_SERVICE_XMPP_PRESENCE", "INBOUND_SERVICE_CHANNEL_PRESENCE", "INBOUND_SERVICE_WARMUP"].
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.inbound_services
	InboundServices []string `json:"inboundServices,omitempty"`

	// Instance class that is used to run this version. Valid values are
	// AutomaticScaling: F1, F2, F4, F4_1G
	// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.instance_class
	InstanceClass *string `json:"instanceClass,omitempty"`

	// Configuration for third-party Python runtime libraries that are required by the application.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.libraries
	Libraries []Library `json:"libraries,omitempty"`

	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.manual_scaling
	ManualScaling *ManualScaling `json:"manualScaling,omitempty"`

	// If set to 'true', the application version will not be deleted.
	// +optional
	NoopOnDestroy *bool `json:"noopOnDestroy,omitempty"`

	// Immutable.
	// +optional
	Project *string `json:"project,omitempty"`

	// Immutable. Optional. The versionId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Desired runtime. Example python27.
	// +required
	// +kcc:proto:field=google.appengine.v1.Version.runtime
	Runtime string `json:"runtime"`

	// The version of the API in the given runtime environment.
	// Please see the app.yaml reference for valid values at 'https://cloud.google.com/appengine/docs/standard/<language>/config/appref'\
	// Substitute '<language>' with 'python', 'java', 'php', 'ruby', 'go' or 'nodejs'.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.runtime_api_version
	RuntimeAPIVersion *string `json:"runtimeApiVersion,omitempty"`

	// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// +required
	ServiceRef *AppEngineServiceRef `json:"serviceRef"`

	// Whether multiple requests can be dispatched to this version at once.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.threadsafe
	Threadsafe *bool `json:"threadsafe,omitempty"`

	// Enables VPC connectivity for standard apps.
	// +optional
	// +kcc:proto:field=google.appengine.v1.Version.vpc_access_connector
	VPCAccessConnector *VPCAccessConnector `json:"vpcAccessConnector,omitempty"`
}

// +kcc:proto=google.appengine.v1.AutomaticScaling
type AutomaticScaling struct {
	// Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.
	//
	// Defaults to a runtime-specific value.
	// +optional
	// +kcc:proto:field=google.appengine.v1.AutomaticScaling.max_concurrent_requests
	MaxConcurrentRequests *int `json:"maxConcurrentRequests,omitempty"`

	// Maximum number of idle instances that should be maintained for this version.
	// +optional
	// +kcc:proto:field=google.appengine.v1.AutomaticScaling.max_idle_instances
	MaxIdleInstances *int `json:"maxIdleInstances,omitempty"`

	// Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +optional
	// +kcc:proto:field=google.appengine.v1.AutomaticScaling.max_pending_latency
	MaxPendingLatency *string `json:"maxPendingLatency,omitempty"`

	// Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.
	// +optional
	// +kcc:proto:field=google.appengine.v1.AutomaticScaling.min_idle_instances
	MinIdleInstances *int `json:"minIdleInstances,omitempty"`

	// Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +optional
	// +kcc:proto:field=google.appengine.v1.AutomaticScaling.min_pending_latency
	MinPendingLatency *string `json:"minPendingLatency,omitempty"`

	// Scheduler settings for standard environment.
	// +optional
	// +kcc:proto:field=google.appengine.v1.AutomaticScaling.standard_scheduler_settings
	StandardSchedulerSettings *StandardSchedulerSettings `json:"standardSchedulerSettings,omitempty"`
}

// +kcc:proto=google.appengine.v1.StandardSchedulerSettings
type StandardSchedulerSettings struct {
	// Maximum number of instances to run for this version. Set to zero to disable maxInstances configuration.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StandardSchedulerSettings.max_instances
	MaxInstances *int `json:"maxInstances,omitempty"`

	// Minimum number of instances to run for this version. Set to zero to disable minInstances configuration.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StandardSchedulerSettings.min_instances
	MinInstances *int `json:"minInstances,omitempty"`

	// Target CPU utilization ratio to maintain when scaling. Should be a value in the range [0.50, 0.95], zero, or a negative value.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StandardSchedulerSettings.target_cpu_utilization
	TargetCPUUtilization *float64 `json:"targetCpuUtilization,omitempty"`

	// Target throughput utilization ratio to maintain when scaling. Should be a value in the range [0.50, 0.95], zero, or a negative value.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StandardSchedulerSettings.target_throughput_utilization
	TargetThroughputUtilization *float64 `json:"targetThroughputUtilization,omitempty"`
}

// +kcc:proto=google.appengine.v1.BasicScaling
type BasicScaling struct {
	// Duration of time after the last request that an instance must wait before the instance is shut down.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
	// +optional
	// +kcc:proto:field=google.appengine.v1.BasicScaling.idle_timeout
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Maximum number of instances to create for this version. Must be in the range [1.0, 200.0].
	// +required
	// +kcc:proto:field=google.appengine.v1.BasicScaling.max_instances
	MaxInstances int `json:"maxInstances"`
}

// +kcc:proto=google.appengine.v1.ManualScaling
type ManualScaling struct {
	// Number of instances to assign to the service at the start.
	//
	// **Note:** When managing the number of instances at runtime through the App Engine Admin API or the (now deprecated) Python 2
	// Modules API set_num_instances() you must use 'lifecycle.ignore_changes = ["manual_scaling"[0].instances]' to prevent drift detection.
	// +required
	// +kcc:proto:field=google.appengine.v1.ManualScaling.instances
	Instances int `json:"instances"`
}

type Deployment struct {
	// Manifest of the files stored in Google Cloud Storage that are included as part of this version.
	// All files must be readable using the credentials supplied with this call.
	// +optional
	Files []DeploymentFile `json:"files,omitempty"`

	// Zip File.
	// +optional
	Zip *DeploymentZip `json:"zip,omitempty"`
}

type DeploymentFile struct {
	// +required
	Name string `json:"name"`

	// SHA1 checksum of the file.
	// +optional
	Sha1Sum *string `json:"sha1Sum,omitempty"`

	// Source URL.
	// +required
	SourceUrl string `json:"sourceUrl"`
}

type DeploymentZip struct {
	// files count.
	// +optional
	FilesCount *int `json:"filesCount,omitempty"`

	// Source URL.
	// +required
	SourceUrl string `json:"sourceUrl"`
}

// +kcc:proto=google.appengine.v1.Entrypoint
type Entrypoint struct {
	// The format should be a shell command that can be fed to bash -c.
	// +required
	// +kcc:proto:field=google.appengine.v1.Entrypoint.shell
	Shell string `json:"shell"`
}

// +kcc:proto=google.appengine.v1.UrlMap
type URLMap struct {
	// Actions to take when the user is not logged in. Possible values: ["AUTH_FAIL_ACTION_REDIRECT", "AUTH_FAIL_ACTION_UNAUTHORIZED"].
	// +optional
	// +kcc:proto:field=google.appengine.v1.UrlMap.auth_fail_action
	AuthFailAction *string `json:"authFailAction,omitempty"`

	// Methods to restrict access to a URL based on login status. Possible values: ["LOGIN_OPTIONAL", "LOGIN_ADMIN", "LOGIN_REQUIRED"].
	// +optional
	// +kcc:proto:field=google.appengine.v1.UrlMap.login
	Login *string `json:"login,omitempty"`

	// 30x code to use when performing redirects for the secure field. Possible values: ["REDIRECT_HTTP_RESPONSE_CODE_301", "REDIRECT_HTTP_RESPONSE_CODE_302", "REDIRECT_HTTP_RESPONSE_CODE_303", "REDIRECT_HTTP_RESPONSE_CODE_307"].
	// +optional
	// +kcc:proto:field=google.appengine.v1.UrlMap.redirect_http_response_code
	RedirectHTTPResponseCode *string `json:"redirectHttpResponseCode,omitempty"`

	// Executes a script to handle the requests that match this URL pattern.
	// Only the auto value is supported for Node.js in the App Engine standard environment, for example "script:" "auto".
	// +optional
	// +kcc:proto:field=google.appengine.v1.UrlMap.script
	Script *ScriptHandler `json:"script,omitempty"`

	// Security (HTTPS) enforcement for this URL. Possible values: ["SECURE_DEFAULT", "SECURE_NEVER", "SECURE_OPTIONAL", "SECURE_ALWAYS"].
	// +optional
	// +kcc:proto:field=google.appengine.v1.UrlMap.security_level
	SecurityLevel *string `json:"securityLevel,omitempty"`

	// Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files. Static file handlers describe which files in the application directory are static files, and which URLs serve them.
	// +optional
	// +kcc:proto:field=google.appengine.v1.UrlMap.static_files
	StaticFiles *StaticFilesHandler `json:"staticFiles,omitempty"`

	// URL prefix. Uses regular expression syntax, which means regexp special characters must be escaped, but should not contain groupings.
	// All URLs that begin with this prefix are handled by this handler, using the portion of the URL after the prefix as part of the file path.
	// +optional
	// +kcc:proto:field=google.appengine.v1.UrlMap.url_regex
	URLRegex *string `json:"urlRegex,omitempty"`
}

// +kcc:proto=google.appengine.v1.ScriptHandler
type ScriptHandler struct {
	// Path to the script from the application root directory.
	// +required
	// +kcc:proto:field=google.appengine.v1.ScriptHandler.script_path
	ScriptPath string `json:"scriptPath"`
}

// +kcc:proto=google.appengine.v1.StaticFilesHandler
type StaticFilesHandler struct {
	// Whether files should also be uploaded as code data. By default, files declared in static file handlers are uploaded as
	// static data and are only served to end users; they cannot be read by the application. If enabled, uploads are charged
	// against both your code and static data storage resource quotas.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StaticFilesHandler.application_readable
	ApplicationReadable *bool `json:"applicationReadable,omitempty"`

	// Time a static file served by this handler should be cached by web proxies and browsers.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example "3.5s".
	// +optional
	// +kcc:proto:field=google.appengine.v1.StaticFilesHandler.expiration
	Expiration *string `json:"expiration,omitempty"`

	// HTTP headers to use for all responses from these URLs.
	// An object containing a list of "key:value" value pairs.".
	// +optional
	// +kcc:proto:field=google.appengine.v1.StaticFilesHandler.http_headers
	HTTPHeaders map[string]string `json:"httpHeaders,omitempty"`

	// MIME type used to serve all files served by this handler.
	// Defaults to file-specific MIME types, which are derived from each file's filename extension.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StaticFilesHandler.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Path to the static files matched by the URL pattern, from the application root directory. The path can refer to text matched in groupings in the URL pattern.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StaticFilesHandler.path
	Path *string `json:"path,omitempty"`

	// Whether this handler should match the request if the file referenced by the handler does not exist.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StaticFilesHandler.require_matching_file
	RequireMatchingFile *bool `json:"requireMatchingFile,omitempty"`

	// Regular expression that matches the file paths for all files that should be referenced by this handler.
	// +optional
	// +kcc:proto:field=google.appengine.v1.StaticFilesHandler.upload_path_regex
	UploadPathRegex *string `json:"uploadPathRegex,omitempty"`
}

// +kcc:proto=google.appengine.v1.Library
type Library struct {
	// Name of the library. Example "django".
	// +optional
	// +kcc:proto:field=google.appengine.v1.Library.name
	Name *string `json:"name,omitempty"`

	// Version of the library to select, or "latest".
	// +optional
	// +kcc:proto:field=google.appengine.v1.Library.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.appengine.v1.VpcAccessConnector
type VPCAccessConnector struct {
	// The egress setting for the connector, controlling what traffic is diverted through it.
	// +optional
	// +kcc:proto:field=google.appengine.v1.VpcAccessConnector.egress_setting
	EgressSetting *string `json:"egressSetting,omitempty"`

	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	// +required
	// +kcc:proto:field=google.appengine.v1.VpcAccessConnector.name
	Name string `json:"name"`
}

type AppEngineStandardAppVersionStatus struct {
	// Conditions represent the latest available observation of the resource's current state.
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// A unique specifier for the AppEngineStandardAppVersion resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Full path to the Version resource in the API. Example, "v1".
	// +optional
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// AppEngineStandardAppVersionObservedState is the state of the AppEngineStandardAppVersion resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.appengine.v1.Version
type AppEngineStandardAppVersionObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpappenginestandardappversion;gcpappenginestandardappversions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppEngineStandardAppVersion is the Schema for the AppEngineStandardAppVersion API
// +k8s:openapi-gen=true
type AppEngineStandardAppVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AppEngineStandardAppVersionSpec   `json:"spec,omitempty"`
	Status AppEngineStandardAppVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppEngineStandardAppVersionList contains a list of AppEngineStandardAppVersion
type AppEngineStandardAppVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineStandardAppVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppEngineStandardAppVersion{}, &AppEngineStandardAppVersionList{})
}
