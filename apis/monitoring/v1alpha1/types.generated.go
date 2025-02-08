// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.api.MonitoredResource
type MonitoredResource struct {
	// Required. The monitored resource type. This field must match
	//  the `type` field of a
	//  [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor]
	//  object. For example, the type of a Compute Engine VM instance is
	//  `gce_instance`. Some descriptors include the service name in the type; for
	//  example, the type of a Datastream stream is
	//  `datastream.googleapis.com/Stream`.
	// +kcc:proto:field=google.api.MonitoredResource.type
	Type *string `json:"type,omitempty"`

	// Required. Values for all of the labels listed in the associated monitored
	//  resource descriptor. For example, Compute Engine VM instances use the
	//  labels `"project_id"`, `"instance_id"`, and `"zone"`.
	// +kcc:proto:field=google.api.MonitoredResource.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.monitoring.v3.InternalChecker
type InternalChecker struct {
	// A unique resource name for this InternalChecker. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/internalCheckers/[INTERNAL_CHECKER_ID]
	//
	//  `[PROJECT_ID_OR_NUMBER]` is the Cloud Monitoring Metrics Scope project for
	//  the Uptime check config associated with the internal checker.
	// +kcc:proto:field=google.monitoring.v3.InternalChecker.name
	Name *string `json:"name,omitempty"`

	// The checker's human-readable name. The display name
	//  should be unique within a Cloud Monitoring Metrics Scope in order to make
	//  it easier to identify; however, uniqueness is not enforced.
	// +kcc:proto:field=google.monitoring.v3.InternalChecker.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The [GCP VPC network](https://cloud.google.com/vpc/docs/vpc) where the
	//  internal resource lives (ex: "default").
	// +kcc:proto:field=google.monitoring.v3.InternalChecker.network
	Network *string `json:"network,omitempty"`

	// The GCP zone the Uptime check should egress from. Only respected for
	//  internal Uptime checks, where internal_network is specified.
	// +kcc:proto:field=google.monitoring.v3.InternalChecker.gcp_zone
	GcpZone *string `json:"gcpZone,omitempty"`

	// The GCP project ID where the internal checker lives. Not necessary
	//  the same as the Metrics Scope project.
	// +kcc:proto:field=google.monitoring.v3.InternalChecker.peer_project_id
	PeerProjectID *string `json:"peerProjectID,omitempty"`

	// The current operational state of the internal checker.
	// +kcc:proto:field=google.monitoring.v3.InternalChecker.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.monitoring.v3.SyntheticMonitorTarget
type SyntheticMonitorTarget struct {
	// Target a Synthetic Monitor GCFv2 instance.
	// +kcc:proto:field=google.monitoring.v3.SyntheticMonitorTarget.cloud_function_v2
	CloudFunctionV2 *SyntheticMonitorTarget_CloudFunctionV2Target `json:"cloudFunctionV2,omitempty"`
}

// +kcc:proto=google.monitoring.v3.SyntheticMonitorTarget.CloudFunctionV2Target
type SyntheticMonitorTarget_CloudFunctionV2Target struct {
	// Required. Fully qualified GCFv2 resource name
	//  i.e. `projects/{project}/locations/{location}/functions/{function}`
	//  Required.
	// +kcc:proto:field=google.monitoring.v3.SyntheticMonitorTarget.CloudFunctionV2Target.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig
type UptimeCheckConfig struct {
	// Identifier. A unique resource name for this Uptime check configuration. The
	//  format is:
	//
	//       projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID]
	//
	//  `[PROJECT_ID_OR_NUMBER]` is the Workspace host project associated with the
	//  Uptime check.
	//
	//  This field should be omitted when creating the Uptime check configuration;
	//  on create, the resource name is assigned by the server and included in the
	//  response.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.name
	Name *string `json:"name,omitempty"`

	// A human-friendly name for the Uptime check configuration. The display name
	//  should be unique within a Cloud Monitoring Workspace in order to make it
	//  easier to identify; however, uniqueness is not enforced. Required.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The [monitored
	//  resource](https://cloud.google.com/monitoring/api/resources) associated
	//  with the configuration.
	//  The following monitored resource types are valid for this field:
	//    `uptime_url`,
	//    `gce_instance`,
	//    `gae_app`,
	//    `aws_ec2_instance`,
	//    `aws_elb_load_balancer`
	//    `k8s_service`
	//    `servicedirectory_service`
	//    `cloud_run_revision`
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.monitored_resource
	MonitoredResource *MonitoredResource `json:"monitoredResource,omitempty"`

	// The group resource associated with the configuration.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.resource_group
	ResourceGroup *UptimeCheckConfig_ResourceGroup `json:"resourceGroup,omitempty"`

	// Specifies a Synthetic Monitor to invoke.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.synthetic_monitor
	SyntheticMonitor *SyntheticMonitorTarget `json:"syntheticMonitor,omitempty"`

	// Contains information needed to make an HTTP or HTTPS check.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.http_check
	HTTPCheck *UptimeCheckConfig_HttpCheck `json:"httpCheck,omitempty"`

	// Contains information needed to make a TCP check.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.tcp_check
	TcpCheck *UptimeCheckConfig_TcpCheck `json:"tcpCheck,omitempty"`

	// How often, in seconds, the Uptime check is performed.
	//  Currently, the only supported values are `60s` (1 minute), `300s`
	//  (5 minutes), `600s` (10 minutes), and `900s` (15 minutes). Optional,
	//  defaults to `60s`.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.period
	Period *string `json:"period,omitempty"`

	// The maximum amount of time to wait for the request to complete (must be
	//  between 1 and 60 seconds). Required.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.timeout
	Timeout *string `json:"timeout,omitempty"`

	// The content that is expected to appear in the data returned by the target
	//  server against which the check is run.  Currently, only the first entry
	//  in the `content_matchers` list is supported, and additional entries will
	//  be ignored. This field is optional and should only be specified if a
	//  content match is required as part of the/ Uptime check.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.content_matchers
	ContentMatchers []UptimeCheckConfig_ContentMatcher `json:"contentMatchers,omitempty"`

	// The type of checkers to use to execute the Uptime check.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.checker_type
	CheckerType *string `json:"checkerType,omitempty"`

	// The list of regions from which the check will be run.
	//  Some regions contain one location, and others contain more than one.
	//  If this field is specified, enough regions must be provided to include a
	//  minimum of 3 locations.  Not specifying this field will result in Uptime
	//  checks running from all available regions.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.selected_regions
	SelectedRegions []string `json:"selectedRegions,omitempty"`

	// If this is `true`, then checks are made only from the 'internal_checkers'.
	//  If it is `false`, then checks are made only from the 'selected_regions'.
	//  It is an error to provide 'selected_regions' when is_internal is `true`,
	//  or to provide 'internal_checkers' when is_internal is `false`.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.is_internal
	IsInternal *bool `json:"isInternal,omitempty"`

	// The internal checkers that this check will egress from. If `is_internal` is
	//  `true` and this list is empty, the check will egress from all the
	//  InternalCheckers configured for the project that owns this
	//  `UptimeCheckConfig`.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.internal_checkers
	InternalCheckers []InternalChecker `json:"internalCheckers,omitempty"`

	// User-supplied key/value data to be used for organizing and
	//  identifying the `UptimeCheckConfig` objects.
	//
	//  The field can contain up to 64 entries. Each key and value is limited to
	//  63 Unicode characters or 128 bytes, whichever is smaller. Labels and
	//  values can contain only lowercase letters, numerals, underscores, and
	//  dashes. Keys must begin with a letter.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.user_labels
	UserLabels map[string]string `json:"userLabels,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.ContentMatcher
type UptimeCheckConfig_ContentMatcher struct {
	// String, regex or JSON content to match. Maximum 1024 bytes. An empty
	//  `content` string indicates no content matching is to be performed.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.ContentMatcher.content
	Content *string `json:"content,omitempty"`

	// The type of content matcher that will be applied to the server output,
	//  compared to the `content` string when the check is run.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.ContentMatcher.matcher
	Matcher *string `json:"matcher,omitempty"`

	// Matcher information for `MATCHES_JSON_PATH` and `NOT_MATCHES_JSON_PATH`
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.ContentMatcher.json_path_matcher
	JsonPathMatcher *UptimeCheckConfig_ContentMatcher_JsonPathMatcher `json:"jsonPathMatcher,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.ContentMatcher.JsonPathMatcher
type UptimeCheckConfig_ContentMatcher_JsonPathMatcher struct {
	// JSONPath within the response output pointing to the expected
	//  `ContentMatcher::content` to match against.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.ContentMatcher.JsonPathMatcher.json_path
	JsonPath *string `json:"jsonPath,omitempty"`

	// The type of JSONPath match that will be applied to the JSON output
	//  (`ContentMatcher.content`)
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.ContentMatcher.JsonPathMatcher.json_matcher
	JsonMatcher *string `json:"jsonMatcher,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.HttpCheck
type UptimeCheckConfig_HttpCheck struct {
	// The HTTP request method to use for the check. If set to
	//  `METHOD_UNSPECIFIED` then `request_method` defaults to `GET`.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.request_method
	RequestMethod *string `json:"requestMethod,omitempty"`

	// If `true`, use HTTPS instead of HTTP to run the check.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.use_ssl
	UseSsl *bool `json:"useSsl,omitempty"`

	// Optional (defaults to "/"). The path to the page against which to run
	//  the check. Will be combined with the `host` (specified within the
	//  `monitored_resource`) and `port` to construct the full URL. If the
	//  provided path does not begin with "/", a "/" will be prepended
	//  automatically.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.path
	Path *string `json:"path,omitempty"`

	// Optional (defaults to 80 when `use_ssl` is `false`, and 443 when
	//  `use_ssl` is `true`). The TCP port on the HTTP server against which to
	//  run the check. Will be combined with host (specified within the
	//  `monitored_resource`) and `path` to construct the full URL.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.port
	Port *int32 `json:"port,omitempty"`

	// The authentication information. Optional when creating an HTTP check;
	//  defaults to empty.
	//  Do not set both `auth_method` and `auth_info`.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.auth_info
	AuthInfo *UptimeCheckConfig_HttpCheck_BasicAuthentication `json:"authInfo,omitempty"`

	// Boolean specifying whether to encrypt the header information.
	//  Encryption should be specified for any headers related to authentication
	//  that you do not wish to be seen when retrieving the configuration. The
	//  server will be responsible for encrypting the headers.
	//  On Get/List calls, if `mask_headers` is set to `true` then the headers
	//  will be obscured with `******.`
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.mask_headers
	MaskHeaders *bool `json:"maskHeaders,omitempty"`

	// The list of headers to send as part of the Uptime check request.
	//  If two headers have the same key and different values, they should
	//  be entered as a single header, with the value being a comma-separated
	//  list of all the desired values as described at
	//  https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31).
	//  Entering two separate headers with the same key in a Create call will
	//  cause the first to be overwritten by the second.
	//  The maximum number of headers allowed is 100.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.headers
	Headers map[string]string `json:"headers,omitempty"`

	// The content type header to use for the check. The following
	//  configurations result in errors:
	//  1. Content type is specified in both the `headers` field and the
	//  `content_type` field.
	//  2. Request method is `GET` and `content_type` is not `TYPE_UNSPECIFIED`
	//  3. Request method is `POST` and `content_type` is `TYPE_UNSPECIFIED`.
	//  4. Request method is `POST` and a "Content-Type" header is provided via
	//  `headers` field. The `content_type` field should be used instead.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.content_type
	ContentType *string `json:"contentType,omitempty"`

	// A user provided content type header to use for the check. The invalid
	//  configurations outlined in the `content_type` field apply to
	//  `custom_content_type`, as well as the following:
	//  1. `content_type` is `URL_ENCODED` and `custom_content_type` is set.
	//  2. `content_type` is `USER_PROVIDED` and `custom_content_type` is not
	//  set.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.custom_content_type
	CustomContentType *string `json:"customContentType,omitempty"`

	// Boolean specifying whether to include SSL certificate validation as a
	//  part of the Uptime check. Only applies to checks where
	//  `monitored_resource` is set to `uptime_url`. If `use_ssl` is `false`,
	//  setting `validate_ssl` to `true` has no effect.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.validate_ssl
	ValidateSsl *bool `json:"validateSsl,omitempty"`

	// The request body associated with the HTTP POST request. If `content_type`
	//  is `URL_ENCODED`, the body passed in must be URL-encoded. Users can
	//  provide a `Content-Length` header via the `headers` field or the API will
	//  do so. If the `request_method` is `GET` and `body` is not empty, the API
	//  will return an error. The maximum byte size is 1 megabyte.
	//
	//  Note: If client libraries aren't used (which performs the conversion
	//  automatically) base64 encode your `body` data since the field is of
	//  `bytes` type.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.body
	Body []byte `json:"body,omitempty"`

	// If present, the check will only pass if the HTTP response status code is
	//  in this set of status codes. If empty, the HTTP status code will only
	//  pass if the HTTP status code is 200-299.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.accepted_response_status_codes
	AcceptedResponseStatusCodes []UptimeCheckConfig_HttpCheck_ResponseStatusCode `json:"acceptedResponseStatusCodes,omitempty"`

	// Contains information needed to add pings to an HTTP check.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.ping_config
	PingConfig *UptimeCheckConfig_PingConfig `json:"pingConfig,omitempty"`

	// If specified, Uptime will generate and attach an OIDC JWT token for the
	//  Monitoring service agent service account as an `Authorization` header
	//  in the HTTP request when probing.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.service_agent_authentication
	ServiceAgentAuthentication *UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication `json:"serviceAgentAuthentication,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.HttpCheck.BasicAuthentication
type UptimeCheckConfig_HttpCheck_BasicAuthentication struct {
	// The username to use when authenticating with the HTTP server.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.BasicAuthentication.username
	Username *string `json:"username,omitempty"`

	// The password to use when authenticating with the HTTP server.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.BasicAuthentication.password
	Password *string `json:"password,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.HttpCheck.ResponseStatusCode
type UptimeCheckConfig_HttpCheck_ResponseStatusCode struct {
	// A status code to accept.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.ResponseStatusCode.status_value
	StatusValue *int32 `json:"statusValue,omitempty"`

	// A class of status codes to accept.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.ResponseStatusCode.status_class
	StatusClass *string `json:"statusClass,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.HttpCheck.ServiceAgentAuthentication
type UptimeCheckConfig_HttpCheck_ServiceAgentAuthentication struct {
	// Type of authentication.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.HttpCheck.ServiceAgentAuthentication.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.PingConfig
type UptimeCheckConfig_PingConfig struct {
	// Number of ICMP pings. A maximum of 3 ICMP pings is currently supported.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.PingConfig.pings_count
	PingsCount *int32 `json:"pingsCount,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.ResourceGroup
type UptimeCheckConfig_ResourceGroup struct {
	// The group of resources being monitored. Should be only the `[GROUP_ID]`,
	//  and not the full-path
	//  `projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID]`.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.ResourceGroup.group_id
	GroupID *string `json:"groupID,omitempty"`

	// The resource type of the group members.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.ResourceGroup.resource_type
	ResourceType *string `json:"resourceType,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig.TcpCheck
type UptimeCheckConfig_TcpCheck struct {
	// The TCP port on the server against which to run the check. Will be
	//  combined with host (specified within the `monitored_resource`) to
	//  construct the full URL. Required.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.TcpCheck.port
	Port *int32 `json:"port,omitempty"`

	// Contains information needed to add pings to a TCP check.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.TcpCheck.ping_config
	PingConfig *UptimeCheckConfig_PingConfig `json:"pingConfig,omitempty"`
}

// +kcc:proto=google.monitoring.v3.SyntheticMonitorTarget
type SyntheticMonitorTargetObservedState struct {
	// Target a Synthetic Monitor GCFv2 instance.
	// +kcc:proto:field=google.monitoring.v3.SyntheticMonitorTarget.cloud_function_v2
	CloudFunctionV2 *SyntheticMonitorTarget_CloudFunctionV2TargetObservedState `json:"cloudFunctionV2,omitempty"`
}

// +kcc:proto=google.monitoring.v3.SyntheticMonitorTarget.CloudFunctionV2Target
type SyntheticMonitorTarget_CloudFunctionV2TargetObservedState struct {
	// Output only. The `cloud_run_revision` Monitored Resource associated with
	//  the GCFv2. The Synthetic Monitor execution results (metrics, logs, and
	//  spans) are reported against this Monitored Resource. This field is output
	//  only.
	// +kcc:proto:field=google.monitoring.v3.SyntheticMonitorTarget.CloudFunctionV2Target.cloud_run_revision
	CloudRunRevision *MonitoredResource `json:"cloudRunRevision,omitempty"`
}

// +kcc:proto=google.monitoring.v3.UptimeCheckConfig
type UptimeCheckConfigObservedState struct {
	// Specifies a Synthetic Monitor to invoke.
	// +kcc:proto:field=google.monitoring.v3.UptimeCheckConfig.synthetic_monitor
	SyntheticMonitor *SyntheticMonitorTargetObservedState `json:"syntheticMonitor,omitempty"`
}
