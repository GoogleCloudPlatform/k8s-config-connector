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


// +kcc:proto=google.api.AuthProvider
type AuthProvider struct {
	// The unique identifier of the auth provider. It will be referred to by
	//  `AuthRequirement.provider_id`.
	//
	//  Example: "bookstore_auth".
	// +kcc:proto:field=google.api.AuthProvider.id
	ID *string `json:"id,omitempty"`

	// Identifies the principal that issued the JWT. See
	//  https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.1
	//  Usually a URL or an email address.
	//
	//  Example: https://securetoken.google.com
	//  Example: 1234567-compute@developer.gserviceaccount.com
	// +kcc:proto:field=google.api.AuthProvider.issuer
	Issuer *string `json:"issuer,omitempty"`

	// URL of the provider's public key set to validate signature of the JWT. See
	//  [OpenID
	//  Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	//  Optional if the key set document:
	//   - can be retrieved from
	//     [OpenID
	//     Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html)
	//     of the issuer.
	//   - can be inferred from the email domain of the issuer (e.g. a Google
	//   service account).
	//
	//  Example: https://www.googleapis.com/oauth2/v1/certs
	// +kcc:proto:field=google.api.AuthProvider.jwks_uri
	JwksURI *string `json:"jwksURI,omitempty"`

	// The list of JWT
	//  [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3).
	//  that are allowed to access. A JWT containing any of these audiences will
	//  be accepted. When this setting is absent, JWTs with audiences:
	//    - "https://[service.name]/[google.protobuf.Api.name]"
	//    - "https://[service.name]/"
	//  will be accepted.
	//  For example, if no audiences are in the setting, LibraryService API will
	//  accept JWTs with the following audiences:
	//    -
	//    https://library-example.googleapis.com/google.example.library.v1.LibraryService
	//    - https://library-example.googleapis.com/
	//
	//  Example:
	//
	//      audiences: bookstore_android.apps.googleusercontent.com,
	//                 bookstore_web.apps.googleusercontent.com
	// +kcc:proto:field=google.api.AuthProvider.audiences
	Audiences *string `json:"audiences,omitempty"`

	// Redirect URL if JWT token is required but not present or is expired.
	//  Implement authorizationUrl of securityDefinitions in OpenAPI spec.
	// +kcc:proto:field=google.api.AuthProvider.authorization_url
	AuthorizationURL *string `json:"authorizationURL,omitempty"`

	// Defines the locations to extract the JWT.  For now it is only used by the
	//  Cloud Endpoints to store the OpenAPI extension [x-google-jwt-locations]
	//  (https://cloud.google.com/endpoints/docs/openapi/openapi-extensions#x-google-jwt-locations)
	//
	//  JWT locations can be one of HTTP headers, URL query parameters or
	//  cookies. The rule is that the first match wins.
	//
	//  If not specified,  default to use following 3 locations:
	//     1) Authorization: Bearer
	//     2) x-goog-iap-jwt-assertion
	//     3) access_token query parameter
	//
	//  Default locations can be specified as followings:
	//     jwt_locations:
	//     - header: Authorization
	//       value_prefix: "Bearer "
	//     - header: x-goog-iap-jwt-assertion
	//     - query: access_token
	// +kcc:proto:field=google.api.AuthProvider.jwt_locations
	JwtLocations []JwtLocation `json:"jwtLocations,omitempty"`
}

// +kcc:proto=google.api.AuthRequirement
type AuthRequirement struct {
	// [id][google.api.AuthProvider.id] from authentication provider.
	//
	//  Example:
	//
	//      provider_id: bookstore_auth
	// +kcc:proto:field=google.api.AuthRequirement.provider_id
	ProviderID *string `json:"providerID,omitempty"`

	// NOTE: This will be deprecated soon, once AuthProvider.audiences is
	//  implemented and accepted in all the runtime components.
	//
	//  The list of JWT
	//  [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3).
	//  that are allowed to access. A JWT containing any of these audiences will
	//  be accepted. When this setting is absent, only JWTs with audience
	//  "https://[Service_name][google.api.Service.name]/[API_name][google.protobuf.Api.name]"
	//  will be accepted. For example, if no audiences are in the setting,
	//  LibraryService API will only accept JWTs with the following audience
	//  "https://library-example.googleapis.com/google.example.library.v1.LibraryService".
	//
	//  Example:
	//
	//      audiences: bookstore_android.apps.googleusercontent.com,
	//                 bookstore_web.apps.googleusercontent.com
	// +kcc:proto:field=google.api.AuthRequirement.audiences
	Audiences *string `json:"audiences,omitempty"`
}

// +kcc:proto=google.api.Authentication
type Authentication struct {
	// A list of authentication rules that apply to individual API methods.
	//
	//  **NOTE:** All service configuration rules follow "last one wins" order.
	// +kcc:proto:field=google.api.Authentication.rules
	Rules []AuthenticationRule `json:"rules,omitempty"`

	// Defines a set of authentication providers that a service supports.
	// +kcc:proto:field=google.api.Authentication.providers
	Providers []AuthProvider `json:"providers,omitempty"`
}

// +kcc:proto=google.api.AuthenticationRule
type AuthenticationRule struct {
	// Selects the methods to which this rule applies.
	//
	//  Refer to [selector][google.api.DocumentationRule.selector] for syntax
	//  details.
	// +kcc:proto:field=google.api.AuthenticationRule.selector
	Selector *string `json:"selector,omitempty"`

	// The requirements for OAuth credentials.
	// +kcc:proto:field=google.api.AuthenticationRule.oauth
	Oauth *OAuthRequirements `json:"oauth,omitempty"`

	// If true, the service accepts API keys without any other credential.
	//  This flag only applies to HTTP and gRPC requests.
	// +kcc:proto:field=google.api.AuthenticationRule.allow_without_credential
	AllowWithoutCredential *bool `json:"allowWithoutCredential,omitempty"`

	// Requirements for additional authentication providers.
	// +kcc:proto:field=google.api.AuthenticationRule.requirements
	Requirements []AuthRequirement `json:"requirements,omitempty"`
}

// +kcc:proto=google.api.Documentation
type Documentation struct {
	// A short description of what the service does. The summary must be plain
	//  text. It becomes the overview of the service displayed in Google Cloud
	//  Console.
	//  NOTE: This field is equivalent to the standard field `description`.
	// +kcc:proto:field=google.api.Documentation.summary
	Summary *string `json:"summary,omitempty"`

	// The top level pages for the documentation set.
	// +kcc:proto:field=google.api.Documentation.pages
	Pages []Page `json:"pages,omitempty"`

	// A list of documentation rules that apply to individual API elements.
	//
	//  **NOTE:** All service configuration rules follow "last one wins" order.
	// +kcc:proto:field=google.api.Documentation.rules
	Rules []DocumentationRule `json:"rules,omitempty"`

	// The URL to the root of documentation.
	// +kcc:proto:field=google.api.Documentation.documentation_root_url
	DocumentationRootURL *string `json:"documentationRootURL,omitempty"`

	// Specifies the service root url if the default one (the service name
	//  from the yaml file) is not suitable. This can be seen in any fully
	//  specified service urls as well as sections that show a base that other
	//  urls are relative to.
	// +kcc:proto:field=google.api.Documentation.service_root_url
	ServiceRootURL *string `json:"serviceRootURL,omitempty"`

	// Declares a single overview page. For example:
	//  <pre><code>documentation:
	//    summary: ...
	//    overview: &#40;== include overview.md ==&#41;
	//  </code></pre>
	//  This is a shortcut for the following declaration (using pages style):
	//  <pre><code>documentation:
	//    summary: ...
	//    pages:
	//    - name: Overview
	//      content: &#40;== include overview.md ==&#41;
	//  </code></pre>
	//  Note: you cannot specify both `overview` field and `pages` field.
	// +kcc:proto:field=google.api.Documentation.overview
	Overview *string `json:"overview,omitempty"`
}

// +kcc:proto=google.api.DocumentationRule
type DocumentationRule struct {
	// The selector is a comma-separated list of patterns for any element such as
	//  a method, a field, an enum value. Each pattern is a qualified name of the
	//  element which may end in "*", indicating a wildcard. Wildcards are only
	//  allowed at the end and for a whole component of the qualified name,
	//  i.e. "foo.*" is ok, but not "foo.b*" or "foo.*.bar". A wildcard will match
	//  one or more components. To specify a default for all applicable elements,
	//  the whole pattern "*" is used.
	// +kcc:proto:field=google.api.DocumentationRule.selector
	Selector *string `json:"selector,omitempty"`

	// Description of the selected proto element (e.g. a message, a method, a
	//  'service' definition, or a field). Defaults to leading & trailing comments
	//  taken from the proto source definition of the proto element.
	// +kcc:proto:field=google.api.DocumentationRule.description
	Description *string `json:"description,omitempty"`

	// Deprecation description of the selected element(s). It can be provided if
	//  an element is marked as `deprecated`.
	// +kcc:proto:field=google.api.DocumentationRule.deprecation_description
	DeprecationDescription *string `json:"deprecationDescription,omitempty"`
}

// +kcc:proto=google.api.Endpoint
type Endpoint struct {
	// The canonical name of this endpoint.
	// +kcc:proto:field=google.api.Endpoint.name
	Name *string `json:"name,omitempty"`

	// Aliases for this endpoint, these will be served by the same UrlMap as the
	//  parent endpoint, and will be provisioned in the GCP stack for the Regional
	//  Endpoints.
	// +kcc:proto:field=google.api.Endpoint.aliases
	Aliases []string `json:"aliases,omitempty"`

	// The specification of an Internet routable address of API frontend that will
	//  handle requests to this [API
	//  Endpoint](https://cloud.google.com/apis/design/glossary). It should be
	//  either a valid IPv4 address or a fully-qualified domain name. For example,
	//  "8.8.8.8" or "myservice.appspot.com".
	// +kcc:proto:field=google.api.Endpoint.target
	Target *string `json:"target,omitempty"`

	// Allowing
	//  [CORS](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing), aka
	//  cross-domain traffic, would allow the backends served from this endpoint to
	//  receive and respond to HTTP OPTIONS requests. The response will be used by
	//  the browser to determine whether the subsequent cross-origin request is
	//  allowed to proceed.
	// +kcc:proto:field=google.api.Endpoint.allow_cors
	AllowCors *bool `json:"allowCors,omitempty"`
}

// +kcc:proto=google.api.JwtLocation
type JwtLocation struct {
	// Specifies HTTP header name to extract JWT token.
	// +kcc:proto:field=google.api.JwtLocation.header
	Header *string `json:"header,omitempty"`

	// Specifies URL query parameter name to extract JWT token.
	// +kcc:proto:field=google.api.JwtLocation.query
	Query *string `json:"query,omitempty"`

	// Specifies cookie name to extract JWT token.
	// +kcc:proto:field=google.api.JwtLocation.cookie
	Cookie *string `json:"cookie,omitempty"`

	// The value prefix. The value format is "value_prefix{token}"
	//  Only applies to "in" header type. Must be empty for "in" query type.
	//  If not empty, the header value has to match (case sensitive) this prefix.
	//  If not matched, JWT will not be extracted. If matched, JWT will be
	//  extracted after the prefix is removed.
	//
	//  For example, for "Authorization: Bearer {JWT}",
	//  value_prefix="Bearer " with a space at the end.
	// +kcc:proto:field=google.api.JwtLocation.value_prefix
	ValuePrefix *string `json:"valuePrefix,omitempty"`
}

// +kcc:proto=google.api.LabelDescriptor
type LabelDescriptor struct {
	// The label key.
	// +kcc:proto:field=google.api.LabelDescriptor.key
	Key *string `json:"key,omitempty"`

	// The type of data that can be assigned to the label.
	// +kcc:proto:field=google.api.LabelDescriptor.value_type
	ValueType *string `json:"valueType,omitempty"`

	// A human-readable description for the label.
	// +kcc:proto:field=google.api.LabelDescriptor.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.api.MetricRule
type MetricRule struct {
	// Selects the methods to which this rule applies.
	//
	//  Refer to [selector][google.api.DocumentationRule.selector] for syntax
	//  details.
	// +kcc:proto:field=google.api.MetricRule.selector
	Selector *string `json:"selector,omitempty"`

	// Metrics to update when the selected methods are called, and the associated
	//  cost applied to each metric.
	//
	//  The key of the map is the metric name, and the values are the amount
	//  increased for the metric against which the quota limits are defined.
	//  The value must not be negative.
	// +kcc:proto:field=google.api.MetricRule.metric_costs
	MetricCosts map[string]int64 `json:"metricCosts,omitempty"`
}

// +kcc:proto=google.api.MonitoredResourceDescriptor
type MonitoredResourceDescriptor struct {
	// Optional. The resource name of the monitored resource descriptor:
	//  `"projects/{project_id}/monitoredResourceDescriptors/{type}"` where
	//  {type} is the value of the `type` field in this object and
	//  {project_id} is a project ID that provides API-specific context for
	//  accessing the type.  APIs that do not use project information can use the
	//  resource name format `"monitoredResourceDescriptors/{type}"`.
	// +kcc:proto:field=google.api.MonitoredResourceDescriptor.name
	Name *string `json:"name,omitempty"`

	// Required. The monitored resource type. For example, the type
	//  `"cloudsql_database"` represents databases in Google Cloud SQL.
	//   For a list of types, see [Monitored resource
	//   types](https://cloud.google.com/monitoring/api/resources)
	//  and [Logging resource
	//  types](https://cloud.google.com/logging/docs/api/v2/resource-list).
	// +kcc:proto:field=google.api.MonitoredResourceDescriptor.type
	Type *string `json:"type,omitempty"`

	// Optional. A concise name for the monitored resource type that might be
	//  displayed in user interfaces. It should be a Title Cased Noun Phrase,
	//  without any article or other determiners. For example,
	//  `"Google Cloud SQL Database"`.
	// +kcc:proto:field=google.api.MonitoredResourceDescriptor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. A detailed description of the monitored resource type that might
	//  be used in documentation.
	// +kcc:proto:field=google.api.MonitoredResourceDescriptor.description
	Description *string `json:"description,omitempty"`

	// Required. A set of labels used to describe instances of this monitored
	//  resource type. For example, an individual Google Cloud SQL database is
	//  identified by values for the labels `"database_id"` and `"zone"`.
	// +kcc:proto:field=google.api.MonitoredResourceDescriptor.labels
	Labels []LabelDescriptor `json:"labels,omitempty"`

	// Optional. The launch stage of the monitored resource definition.
	// +kcc:proto:field=google.api.MonitoredResourceDescriptor.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`
}

// +kcc:proto=google.api.Monitoring
type Monitoring struct {
	// Monitoring configurations for sending metrics to the producer project.
	//  There can be multiple producer destinations. A monitored resource type may
	//  appear in multiple monitoring destinations if different aggregations are
	//  needed for different sets of metrics associated with that monitored
	//  resource type. A monitored resource and metric pair may only be used once
	//  in the Monitoring configuration.
	// +kcc:proto:field=google.api.Monitoring.producer_destinations
	ProducerDestinations []Monitoring_MonitoringDestination `json:"producerDestinations,omitempty"`

	// Monitoring configurations for sending metrics to the consumer project.
	//  There can be multiple consumer destinations. A monitored resource type may
	//  appear in multiple monitoring destinations if different aggregations are
	//  needed for different sets of metrics associated with that monitored
	//  resource type. A monitored resource and metric pair may only be used once
	//  in the Monitoring configuration.
	// +kcc:proto:field=google.api.Monitoring.consumer_destinations
	ConsumerDestinations []Monitoring_MonitoringDestination `json:"consumerDestinations,omitempty"`
}

// +kcc:proto=google.api.Monitoring.MonitoringDestination
type Monitoring_MonitoringDestination struct {
	// The monitored resource type. The type must be defined in
	//  [Service.monitored_resources][google.api.Service.monitored_resources]
	//  section.
	// +kcc:proto:field=google.api.Monitoring.MonitoringDestination.monitored_resource
	MonitoredResource *string `json:"monitoredResource,omitempty"`

	// Types of the metrics to report to this monitoring destination.
	//  Each type must be defined in
	//  [Service.metrics][google.api.Service.metrics] section.
	// +kcc:proto:field=google.api.Monitoring.MonitoringDestination.metrics
	Metrics []string `json:"metrics,omitempty"`
}

// +kcc:proto=google.api.OAuthRequirements
type OAuthRequirements struct {
	// The list of publicly documented OAuth scopes that are allowed access. An
	//  OAuth token containing any of these scopes will be accepted.
	//
	//  Example:
	//
	//       canonical_scopes: https://www.googleapis.com/auth/calendar,
	//                         https://www.googleapis.com/auth/calendar.read
	// +kcc:proto:field=google.api.OAuthRequirements.canonical_scopes
	CanonicalScopes *string `json:"canonicalScopes,omitempty"`
}

// +kcc:proto=google.api.Page
type Page struct {
	// The name of the page. It will be used as an identity of the page to
	//  generate URI of the page, text of the link to this page in navigation,
	//  etc. The full page name (start from the root page name to this page
	//  concatenated with `.`) can be used as reference to the page in your
	//  documentation. For example:
	//  <pre><code>pages:
	//  - name: Tutorial
	//    content: &#40;== include tutorial.md ==&#41;
	//    subpages:
	//    - name: Java
	//      content: &#40;== include tutorial_java.md ==&#41;
	//  </code></pre>
	//  You can reference `Java` page using Markdown reference link syntax:
	//  `[Java][Tutorial.Java]`.
	// +kcc:proto:field=google.api.Page.name
	Name *string `json:"name,omitempty"`

	// The Markdown content of the page. You can use <code>&#40;== include {path}
	//  ==&#41;</code> to include content from a Markdown file. The content can be
	//  used to produce the documentation page such as HTML format page.
	// +kcc:proto:field=google.api.Page.content
	Content *string `json:"content,omitempty"`

	// Subpages of this page. The order of subpages specified here will be
	//  honored in the generated docset.
	// +kcc:proto:field=google.api.Page.subpages
	Subpages []Page `json:"subpages,omitempty"`
}

// +kcc:proto=google.api.Quota
type Quota struct {
	// List of QuotaLimit definitions for the service.
	// +kcc:proto:field=google.api.Quota.limits
	Limits []QuotaLimit `json:"limits,omitempty"`

	// List of MetricRule definitions, each one mapping a selected method to one
	//  or more metrics.
	// +kcc:proto:field=google.api.Quota.metric_rules
	MetricRules []MetricRule `json:"metricRules,omitempty"`
}

// +kcc:proto=google.api.QuotaLimit
type QuotaLimit struct {
	// Name of the quota limit.
	//
	//  The name must be provided, and it must be unique within the service. The
	//  name can only include alphanumeric characters as well as '-'.
	//
	//  The maximum length of the limit name is 64 characters.
	// +kcc:proto:field=google.api.QuotaLimit.name
	Name *string `json:"name,omitempty"`

	// Optional. User-visible, extended description for this quota limit.
	//  Should be used only when more context is needed to understand this limit
	//  than provided by the limit's display name (see: `display_name`).
	// +kcc:proto:field=google.api.QuotaLimit.description
	Description *string `json:"description,omitempty"`

	// Default number of tokens that can be consumed during the specified
	//  duration. This is the number of tokens assigned when a client
	//  application developer activates the service for his/her project.
	//
	//  Specifying a value of 0 will block all requests. This can be used if you
	//  are provisioning quota to selected consumers and blocking others.
	//  Similarly, a value of -1 will indicate an unlimited quota. No other
	//  negative values are allowed.
	//
	//  Used by group-based quotas only.
	// +kcc:proto:field=google.api.QuotaLimit.default_limit
	DefaultLimit *int64 `json:"defaultLimit,omitempty"`

	// Maximum number of tokens that can be consumed during the specified
	//  duration. Client application developers can override the default limit up
	//  to this maximum. If specified, this value cannot be set to a value less
	//  than the default limit. If not specified, it is set to the default limit.
	//
	//  To allow clients to apply overrides with no upper bound, set this to -1,
	//  indicating unlimited maximum quota.
	//
	//  Used by group-based quotas only.
	// +kcc:proto:field=google.api.QuotaLimit.max_limit
	MaxLimit *int64 `json:"maxLimit,omitempty"`

	// Free tier value displayed in the Developers Console for this limit.
	//  The free tier is the number of tokens that will be subtracted from the
	//  billed amount when billing is enabled.
	//  This field can only be set on a limit with duration "1d", in a billable
	//  group; it is invalid on any other limit. If this field is not set, it
	//  defaults to 0, indicating that there is no free tier for this service.
	//
	//  Used by group-based quotas only.
	// +kcc:proto:field=google.api.QuotaLimit.free_tier
	FreeTier *int64 `json:"freeTier,omitempty"`

	// Duration of this limit in textual notation. Must be "100s" or "1d".
	//
	//  Used by group-based quotas only.
	// +kcc:proto:field=google.api.QuotaLimit.duration
	Duration *string `json:"duration,omitempty"`

	// The name of the metric this quota limit applies to. The quota limits with
	//  the same metric will be checked together during runtime. The metric must be
	//  defined within the service config.
	// +kcc:proto:field=google.api.QuotaLimit.metric
	Metric *string `json:"metric,omitempty"`

	// Specify the unit of the quota limit. It uses the same syntax as
	//  [MetricDescriptor.unit][google.api.MetricDescriptor.unit]. The supported
	//  unit kinds are determined by the quota backend system.
	//
	//  Here are some examples:
	//  * "1/min/{project}" for quota per minute per project.
	//
	//  Note: the order of unit components is insignificant.
	//  The "1" at the beginning is required to follow the metric unit syntax.
	// +kcc:proto:field=google.api.QuotaLimit.unit
	Unit *string `json:"unit,omitempty"`

	// Tiered limit values. You must specify this as a key:value pair, with an
	//  integer value that is the maximum number of requests allowed for the
	//  specified unit. Currently only STANDARD is supported.
	// +kcc:proto:field=google.api.QuotaLimit.values
	Values map[string]int64 `json:"values,omitempty"`

	// User-visible display name for this limit.
	//  Optional. If not set, the UI will provide a default display name based on
	//  the quota configuration. This field can be used to override the default
	//  display name generated from the configuration.
	// +kcc:proto:field=google.api.QuotaLimit.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.api.Usage
type Usage struct {
	// Requirements that must be satisfied before a consumer project can use the
	//  service. Each requirement is of the form <service.name>/<requirement-id>;
	//  for example 'serviceusage.googleapis.com/billing-enabled'.
	//
	//  For Google APIs, a Terms of Service requirement must be included here.
	//  Google Cloud APIs must include "serviceusage.googleapis.com/tos/cloud".
	//  Other Google APIs should include
	//  "serviceusage.googleapis.com/tos/universal". Additional ToS can be
	//  included based on the business needs.
	// +kcc:proto:field=google.api.Usage.requirements
	Requirements []string `json:"requirements,omitempty"`

	// A list of usage rules that apply to individual API methods.
	//
	//  **NOTE:** All service configuration rules follow "last one wins" order.
	// +kcc:proto:field=google.api.Usage.rules
	Rules []UsageRule `json:"rules,omitempty"`

	// The full resource name of a channel used for sending notifications to the
	//  service producer.
	//
	//  Google Service Management currently only supports
	//  [Google Cloud Pub/Sub](https://cloud.google.com/pubsub) as a notification
	//  channel. To use Google Cloud Pub/Sub as the channel, this must be the name
	//  of a Cloud Pub/Sub topic that uses the Cloud Pub/Sub topic name format
	//  documented in https://cloud.google.com/pubsub/docs/overview.
	// +kcc:proto:field=google.api.Usage.producer_notification_channel
	ProducerNotificationChannel *string `json:"producerNotificationChannel,omitempty"`
}

// +kcc:proto=google.api.UsageRule
type UsageRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	//  methods in all APIs.
	//
	//  Refer to [selector][google.api.DocumentationRule.selector] for syntax
	//  details.
	// +kcc:proto:field=google.api.UsageRule.selector
	Selector *string `json:"selector,omitempty"`

	// If true, the selected method allows unregistered calls, e.g. calls
	//  that don't identify any user or application.
	// +kcc:proto:field=google.api.UsageRule.allow_unregistered_calls
	AllowUnregisteredCalls *bool `json:"allowUnregisteredCalls,omitempty"`

	// If true, the selected method should skip service control and the control
	//  plane features, such as quota and billing, will not be available.
	//  This flag is used by Google Cloud Endpoints to bypass checks for internal
	//  methods, such as service health check methods.
	// +kcc:proto:field=google.api.UsageRule.skip_service_control
	SkipServiceControl *bool `json:"skipServiceControl,omitempty"`
}

// +kcc:proto=google.api.serviceusage.v1.Service
type Service struct {
	// The resource name of the consumer and service.
	//
	//  A valid name would be:
	//  - projects/123/services/serviceusage.googleapis.com
	// +kcc:proto:field=google.api.serviceusage.v1.Service.name
	Name *string `json:"name,omitempty"`

	// The resource name of the consumer.
	//
	//  A valid name would be:
	//  - projects/123
	// +kcc:proto:field=google.api.serviceusage.v1.Service.parent
	Parent *string `json:"parent,omitempty"`

	// The service configuration of the available service.
	//  Some fields may be filtered out of the configuration in responses to
	//  the `ListServices` method. These fields are present only in responses to
	//  the `GetService` method.
	// +kcc:proto:field=google.api.serviceusage.v1.Service.config
	Config *ServiceConfig `json:"config,omitempty"`

	// Whether or not the service has been enabled for use by the consumer.
	// +kcc:proto:field=google.api.serviceusage.v1.Service.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.api.serviceusage.v1.ServiceConfig
type ServiceConfig struct {
	// The DNS address at which this service is available.
	//
	//  An example DNS address would be:
	//  `calendar.googleapis.com`.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.name
	Name *string `json:"name,omitempty"`

	// The product title for this service.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.title
	Title *string `json:"title,omitempty"`

	// A list of API interfaces exported by this service. Contains only the names,
	//  versions, and method names of the interfaces.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.apis
	Apis []Api `json:"apis,omitempty"`

	// Additional API documentation. Contains only the summary and the
	//  documentation URL.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Quota configuration.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.quota
	Quota *Quota `json:"quota,omitempty"`

	// Auth configuration. Contains only the OAuth rules.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.authentication
	Authentication *Authentication `json:"authentication,omitempty"`

	// Configuration controlling usage of this service.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.usage
	Usage *Usage `json:"usage,omitempty"`

	// Configuration for network endpoints. Contains only the names and aliases
	//  of the endpoints.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.endpoints
	Endpoints []Endpoint `json:"endpoints,omitempty"`

	// Defines the monitored resources used by this service. This is required
	//  by the [Service.monitoring][google.api.Service.monitoring] and
	//  [Service.logging][google.api.Service.logging] configurations.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.monitored_resources
	MonitoredResources []MonitoredResourceDescriptor `json:"monitoredResources,omitempty"`

	// Monitoring configuration.
	//  This should not include the 'producer_destinations' field.
	// +kcc:proto:field=google.api.serviceusage.v1.ServiceConfig.monitoring
	Monitoring *Monitoring `json:"monitoring,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.Api
type Api struct {
	// The fully qualified name of this interface, including package name
	//  followed by the interface's simple name.
	// +kcc:proto:field=google.protobuf.Api.name
	Name *string `json:"name,omitempty"`

	// The methods of this interface, in unspecified order.
	// +kcc:proto:field=google.protobuf.Api.methods
	Methods []Method `json:"methods,omitempty"`

	// Any metadata attached to the interface.
	// +kcc:proto:field=google.protobuf.Api.options
	Options []Option `json:"options,omitempty"`

	// A version string for this interface. If specified, must have the form
	//  `major-version.minor-version`, as in `1.10`. If the minor version is
	//  omitted, it defaults to zero. If the entire version field is empty, the
	//  major version is derived from the package name, as outlined below. If the
	//  field is not empty, the version in the package name will be verified to be
	//  consistent with what is provided here.
	//
	//  The versioning schema uses [semantic
	//  versioning](http://semver.org) where the major version number
	//  indicates a breaking change and the minor version an additive,
	//  non-breaking change. Both version numbers are signals to users
	//  what to expect from different versions, and should be carefully
	//  chosen based on the product plan.
	//
	//  The major version is also reflected in the package name of the
	//  interface, which must end in `v<major-version>`, as in
	//  `google.feature.v1`. For major versions 0 and 1, the suffix can
	//  be omitted. Zero major versions must only be used for
	//  experimental, non-GA interfaces.
	// +kcc:proto:field=google.protobuf.Api.version
	Version *string `json:"version,omitempty"`

	// Source context for the protocol buffer service represented by this
	//  message.
	// +kcc:proto:field=google.protobuf.Api.source_context
	SourceContext *SourceContext `json:"sourceContext,omitempty"`

	// Included interfaces. See [Mixin][].
	// +kcc:proto:field=google.protobuf.Api.mixins
	Mixins []Mixin `json:"mixins,omitempty"`

	// The source syntax of the service.
	// +kcc:proto:field=google.protobuf.Api.syntax
	Syntax *string `json:"syntax,omitempty"`
}

// +kcc:proto=google.protobuf.Method
type Method struct {
	// The simple name of this method.
	// +kcc:proto:field=google.protobuf.Method.name
	Name *string `json:"name,omitempty"`

	// A URL of the input message type.
	// +kcc:proto:field=google.protobuf.Method.request_type_url
	RequestTypeURL *string `json:"requestTypeURL,omitempty"`

	// If true, the request is streamed.
	// +kcc:proto:field=google.protobuf.Method.request_streaming
	RequestStreaming *bool `json:"requestStreaming,omitempty"`

	// The URL of the output message type.
	// +kcc:proto:field=google.protobuf.Method.response_type_url
	ResponseTypeURL *string `json:"responseTypeURL,omitempty"`

	// If true, the response is streamed.
	// +kcc:proto:field=google.protobuf.Method.response_streaming
	ResponseStreaming *bool `json:"responseStreaming,omitempty"`

	// Any metadata attached to the method.
	// +kcc:proto:field=google.protobuf.Method.options
	Options []Option `json:"options,omitempty"`

	// The source syntax of this method.
	// +kcc:proto:field=google.protobuf.Method.syntax
	Syntax *string `json:"syntax,omitempty"`
}

// +kcc:proto=google.protobuf.Mixin
type Mixin struct {
	// The fully qualified name of the interface which is included.
	// +kcc:proto:field=google.protobuf.Mixin.name
	Name *string `json:"name,omitempty"`

	// If non-empty specifies a path under which inherited HTTP paths
	//  are rooted.
	// +kcc:proto:field=google.protobuf.Mixin.root
	Root *string `json:"root,omitempty"`
}

// +kcc:proto=google.protobuf.Option
type Option struct {
	// The option's name. For protobuf built-in options (options defined in
	//  descriptor.proto), this is the short name. For example, `"map_entry"`.
	//  For custom options, it should be the fully-qualified name. For example,
	//  `"google.api.http"`.
	// +kcc:proto:field=google.protobuf.Option.name
	Name *string `json:"name,omitempty"`

	// The option's value packed in an Any message. If the value is a primitive,
	//  the corresponding wrapper type defined in google/protobuf/wrappers.proto
	//  should be used. If the value is an enum, it should be stored as an int32
	//  value using the google.protobuf.Int32Value type.
	// +kcc:proto:field=google.protobuf.Option.value
	Value *Any `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.SourceContext
type SourceContext struct {
	// The path-qualified name of the .proto file that contained the associated
	//  protobuf element.  For example: `"google/protobuf/source_context.proto"`.
	// +kcc:proto:field=google.protobuf.SourceContext.file_name
	FileName *string `json:"fileName,omitempty"`
}
