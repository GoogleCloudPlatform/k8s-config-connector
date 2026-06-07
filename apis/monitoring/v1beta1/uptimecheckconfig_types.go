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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitoringUptimeCheckConfigGVK = GroupVersion.WithKind("MonitoringUptimeCheckConfig")

type MonitoringGroupRef struct {
	// The group of resources being monitored. Should be only the `[GROUP_ID]`, and not the full-path `projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID]`.
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.ContentMatcher
type UptimeCheckConfig_ContentMatcher struct {
	// String, regex or JSON content to match. Maximum 1024 bytes. An empty `content` string indicates no content matching is to be performed.
	// +required
	Content *string `json:"content,omitempty"`
	// The type of content matcher that will be applied to the server output, compared to the `content` string when the check is run.
	Matcher *string `json:"matcher,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.HttpCheck.BasicAuthentication
type UptimeCheckConfig_HTTPCheck_BasicAuthentication struct {
	// The password to use when authenticating with the HTTP server.
	// +required
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`
	// The username to use when authenticating with the HTTP server.
	// +required
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.HttpCheck
type UptimeCheckConfig_HTTPCheck struct {
	// The authentication information. Optional when creating an HTTP check; defaults to empty. Do not set both `auth_method` and `auth_info`.
	AuthInfo *UptimeCheckConfig_HTTPCheck_BasicAuthentication `json:"authInfo,omitempty"`
	// The request body associated with the HTTP POST request. If `content_type` is `URL_ENCODED`, the body passed in must be URL-encoded. Users can provide a `Content-Length` header via the `headers` field or the API will do so. If the `request_method` is `GET` and `body` is not empty, the API will return an error. The maximum byte size is 1 megabyte. Note: As with all `bytes` fields JSON representations are base64 encoded. e.g.: "foo=bar" in URL-encoded form is "foo%3Dbar" and in base64 encoding is "Zm9vJTI1M0RiYXI=".
	Body *string `json:"body,omitempty"`
	// Immutable. The content type to use for the check. Possible values: TYPE_UNSPECIFIED, URL_ENCODED
	ContentType *string `json:"contentType,omitempty"`
	// The list of headers to send as part of the Uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
	Headers map[string]string `json:"headers,omitempty"`
	// Immutable. Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if `mask_headers` is set to `true` then the headers will be obscured with `******.`
	MaskHeaders *bool `json:"maskHeaders,omitempty"`
	// Optional (defaults to "/"). The path to the page against which to run the check. Will be combined with the `host` (specified within the `monitored_resource`) and `port` to construct the full URL. If the provided path does not begin with "/", a "/" will be prepended automatically.
	Path *string `json:"path,omitempty"`
	// Optional (defaults to 80 when `use_ssl` is `false`, and 443 when `use_ssl` is `true`). The TCP port on the HTTP server against which to run the check. Will be combined with host (specified within the `monitored_resource`) and `path` to construct the full URL.
	Port *int64 `json:"port,omitempty"`
	// Immutable. The HTTP request method to use for the check. If set to `METHOD_UNSPECIFIED` then `request_method` defaults to `GET`.
	RequestMethod *string `json:"requestMethod,omitempty"`
	// If `true`, use HTTPS instead of HTTP to run the check.
	UseSsl *bool `json:"useSsl,omitempty"`
	// Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where `monitored_resource` is set to `uptime_url`. If `use_ssl` is `false`, setting `validate_ssl` to `true` has no effect.
	ValidateSsl *bool `json:"validateSsl,omitempty"`
}

// +kcc:proto=google.api.MonitoredResource
type UptimeCheckConfig_MonitoredResource struct {
	// Immutable.
	// +required
	Labels map[string]string `json:"filterLabels,omitempty"`
	// Immutable.
	// +required
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.ResourceGroup
type UptimeCheckConfig_ResourceGroup struct {
	// Immutable. The group resource associated with the configuration.
	GroupIDRef *MonitoringGroupRef `json:"groupRef,omitempty"`
	// Immutable. The resource type of the group members. Possible values: RESOURCE_TYPE_UNSPECIFIED, INSTANCE, AWS_ELB_LOAD_BALANCER
	ResourceType *string `json:"resourceType,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.TcpCheck
type UptimeCheckConfig_TCPCheck struct {
	// The TCP port on the server against which to run the check. Will be combined with host (specified within the `monitored_resource`) to construct the full URL. Required.
	// +required
	Port *int64 `json:"port,omitempty"`
}

// MonitoringUptimeCheckConfigSpec defines the desired state of MonitoringUptimeCheckConfig
// +kcc:spec:proto=google.monitoring.v3.UptimeCheckConfig
type MonitoringUptimeCheckConfigSpec struct {
	// A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required.
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The Project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required.
	// +required
	Timeout *string `json:"timeout,omitempty"`

	// The content that is expected to appear in the data returned by the target server against which the check is run.  Currently, only the first entry in the `content_matchers` list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check.
	ContentMatchers []UptimeCheckConfig_ContentMatcher `json:"contentMatchers,omitempty"`

	// Contains information needed to make an HTTP or HTTPS check.
	HTTPCheck *UptimeCheckConfig_HTTPCheck `json:"httpCheck,omitempty"`

	// Immutable. The [monitored resource](https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for Uptime checks:   `uptime_url`,   `gce_instance`,   `gae_app`,   `aws_ec2_instance`,   `aws_elb_load_balancer`
	MonitoredResource *UptimeCheckConfig_MonitoredResource `json:"monitoredResource,omitempty"`

	// How often, in seconds, the Uptime check is performed. Currently, the only supported values are `60s` (1 minute), `300s` (5 minutes), `600s` (10 minutes), and `900s` (15 minutes). Optional, defaults to `60s`.
	Period *string `json:"period,omitempty"`

	// Immutable. The group resource associated with the configuration.
	ResourceGroup *UptimeCheckConfig_ResourceGroup `json:"resourceGroup,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations.  Not specifying this field will result in Uptime checks running from all available regions.
	SelectedRegions []string `json:"selectedRegions,omitempty"`

	// Contains information needed to make a TCP check.
	TCPCheck *UptimeCheckConfig_TCPCheck `json:"tcpCheck,omitempty"`
}

// MonitoringUptimeCheckConfigStatus defines the config connector machine state of MonitoringUptimeCheckConfig
// +kcc:proto=google.monitoring.v3.UptimeCheckConfig
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
