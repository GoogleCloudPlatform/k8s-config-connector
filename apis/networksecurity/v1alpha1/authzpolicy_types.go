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

var NetworkSecurityAuthzPolicyGVK = GroupVersion.WithKind("NetworkSecurityAuthzPolicy")

// NetworkSecurityAuthzPolicySpec defines the desired state of NetworkSecurityAuthzPolicy
// +kcc:spec:proto=google.cloud.networksecurity.v1.AuthzPolicy
type NetworkSecurityAuthzPolicySpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The NetworkSecurityAuthzPolicy name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. Set of labels associated with the AuthzPolicy resource.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Specifies the set of resources to which this policy should be
	//  applied to.
	// +kubebuilder:validation:Required
	Target *AuthzPolicy_Target `json:"target,omitempty"`

	// Optional. A list of authorization HTTP rules to match against the incoming
	//  request. A policy match occurs when at least one HTTP rule matches the
	//  request or when no HTTP rules are specified in the policy.
	//  At least one HTTP Rule is required for Allow or Deny Action. Limited
	//  to 5 rules.
	// +kubebuilder:validation:Optional
	HTTPRules []AuthzPolicy_AuthzRule `json:"httpRules,omitempty"`

	// Optional. A list of authorization network rules to match against the
	//  incoming request. A policy match occurs when at least one network rule
	//  matches the request.
	//  At least one network rule is required for Allow or Deny Action if no HTTP
	//  rules are provided. Network rules are mutually exclusive with HTTP rules.
	//  Limited to 5 rules.
	// +kubebuilder:validation:Optional
	NetworkRules []AuthzPolicy_AuthzRule `json:"networkRules,omitempty"`

	// Required. Can be one of ALLOW, DENY, CUSTOM.
	// +kubebuilder:validation:Required
	Action *string `json:"action,omitempty"`

	// Optional. Required if the action is CUSTOM. Allows delegating
	//  authorization decisions to Cloud IAP or to Service Extensions. One of
	//  cloudIap or authzExtension must be specified.
	// +kubebuilder:validation:Optional
	CustomProvider *AuthzPolicy_CustomProvider `json:"customProvider,omitempty"`

	// Optional. Immutable. Defines the type of authorization being performed.
	//  If not specified, REQUEST_AUTHZ is applied. This field cannot be changed
	//  once AuthzPolicy is created.
	// +kubebuilder:validation:Optional
	PolicyProfile *string `json:"policyProfile,omitempty"`
}

// NetworkSecurityAuthzPolicyStatus defines the config connector machine state of NetworkSecurityAuthzPolicy
type NetworkSecurityAuthzPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityAuthzPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityAuthzPolicyObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityAuthzPolicyObservedState is the state of the NetworkSecurityAuthzPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.AuthzPolicy
type NetworkSecurityAuthzPolicyObservedState struct {
	// Output only. The timestamp when the resource was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityauthzpolicy;gcpnetworksecurityauthzpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityAuthzPolicy is the Schema for the NetworkSecurityAuthzPolicy API
// +k8s:openapi-gen=true
type NetworkSecurityAuthzPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityAuthzPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityAuthzPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityAuthzPolicyList contains a list of NetworkSecurityAuthzPolicy
type NetworkSecurityAuthzPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityAuthzPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityAuthzPolicy{}, &NetworkSecurityAuthzPolicyList{})
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule
type AuthzPolicy_AuthzRule struct {
	// Optional. Describes properties of a source of a request.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.from
	From *AuthzPolicy_AuthzRule_From `json:"from,omitempty"`

	// Optional. Describes properties of a target of a request.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.to
	To *AuthzPolicy_AuthzRule_To `json:"to,omitempty"`

	// Optional. CEL expression that describes the conditions to be satisfied
	//  for the action. The result of the CEL expression is ANDed with the from
	//  and to. Refer to the CEL language reference for a list of available
	//  attributes.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.when
	When *string `json:"when,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.From
type AuthzPolicy_AuthzRule_From struct {
	// Optional. Describes the properties of a request's sources. At least one
	//  of sources or notSources must be specified. Limited to 1 source.
	//  A match occurs when ANY source (in sources or notSources) matches the
	//  request. Within a single source, the match follows AND semantics
	//  across fields and OR semantics within a single field, i.e. a match
	//  occurs when ANY principal matches AND ANY ipBlocks match.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.From.sources
	Sources []AuthzPolicy_AuthzRule_From_RequestSource `json:"sources,omitempty"`

	// Optional. Describes the negated properties of request sources. Matches
	//  requests from sources that do not match the criteria specified in this
	//  field. At least one of sources or notSources must be specified.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.From.not_sources
	NotSources []AuthzPolicy_AuthzRule_From_RequestSource `json:"notSources,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.From.RequestSource
type AuthzPolicy_AuthzRule_From_RequestSource struct {
	// Optional. A list of identities derived from the client's certificate.
	//  This field will not match on a request unless frontend mutual TLS is
	//  enabled for the forwarding rule or Gateway and the client certificate
	//  has been successfully validated by mTLS.
	//  Each identity is a string whose value is matched against a list of
	//  URI SANs, DNS Name SANs, or the common name in the client's
	//  certificate. A match happens when any principal matches with the
	//  rule. Limited to 50 principals per Authorization Policy for regional
	//  internal Application Load Balancers, regional external Application
	//  Load Balancers, cross-region internal Application Load Balancers, and
	//  Cloud Service Mesh. This field is not supported for global external
	//  Application Load Balancers.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.From.RequestSource.principals
	Principals []AuthzPolicy_AuthzRule_Principal `json:"principals,omitempty"`

	// Optional. A list of IP addresses or IP address ranges to match
	//  against the source IP address of the request. Limited to 10 ip_blocks
	//  per Authorization Policy
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.From.RequestSource.ip_blocks
	IPBlocks []AuthzPolicy_AuthzRule_IPBlock `json:"ipBlocks,omitempty"`

	// Optional. A list of resources to match against the resource of the
	//  source VM of a request. Limited to 10 resources per Authorization
	//  Policy.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.From.RequestSource.resources
	Resources []AuthzPolicy_AuthzRule_RequestResource `json:"resources,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.HeaderMatch
type AuthzPolicy_AuthzRule_HeaderMatch struct {
	// Optional. Specifies the name of the header in the request.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.HeaderMatch.name
	Name *string `json:"name,omitempty"`

	// Optional. Specifies how the header match will be performed.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.HeaderMatch.value
	Value *AuthzPolicy_AuthzRule_StringMatch `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.IpBlock
type AuthzPolicy_AuthzRule_IPBlock struct {
	// Required. The address prefix.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.IpBlock.prefix
	// +kubebuilder:validation:Required
	Prefix *string `json:"prefix,omitempty"`

	// Required. The length of the address range.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.IpBlock.length
	// +kubebuilder:validation:Required
	Length *int32 `json:"length,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.Principal
type AuthzPolicy_AuthzRule_Principal struct {
	// Optional. An enum to decide what principal value the principal rule
	//  will match against. If not specified, the PrincipalSelector is
	//  CLIENT_CERT_URI_SAN.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.Principal.principal_selector
	PrincipalSelector *string `json:"principalSelector,omitempty"`

	// Required. A non-empty string whose value is matched against the
	//  principal value based on the principal_selector. Only exact match can
	//  be applied for CLIENT_CERT_URI_SAN, CLIENT_CERT_DNS_NAME_SAN,
	//  CLIENT_CERT_COMMON_NAME selectors.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.Principal.principal
	// +kubebuilder:validation:Required
	Principal *AuthzPolicy_AuthzRule_StringMatch `json:"principal,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.RequestResource
type AuthzPolicy_AuthzRule_RequestResource struct {
	// Optional. A list of resource tag value permanent IDs to match against
	//  the resource manager tags value associated with the source VM of a
	//  request.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.RequestResource.tag_value_id_set
	TagValueIDSet *AuthzPolicy_AuthzRule_RequestResource_TagValueIDSet `json:"tagValueIDSet,omitempty"`

	// Optional. An IAM service account to match against the source
	//  service account of the VM sending the request.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.RequestResource.iam_service_account
	IAMServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"iamServiceAccountRef,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.RequestResource.TagValueIdSet
type AuthzPolicy_AuthzRule_RequestResource_TagValueIDSet struct {
	// Required. A list of resource tag value permanent IDs to match against
	//  the resource manager tags value associated with the source VM of a
	//  request. The match follows AND semantics which means all
	//  the ids must match. Limited to 5 ids in the Tag value id set.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.RequestResource.TagValueIdSet.ids
	// +kubebuilder:validation:Required
	Ids []int64 `json:"ids,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.StringMatch
type AuthzPolicy_AuthzRule_StringMatch struct {
	// The input string must match exactly the string specified here.
	//
	//  Examples:
	//
	//  * ``abc`` only matches the value ``abc``.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.StringMatch.exact
	Exact *string `json:"exact,omitempty"`

	// The input string must have the prefix specified here.
	//  Note: empty prefix is not allowed, please use regex instead.
	//
	//  Examples:
	//
	//  * ``abc`` matches the value ``abc.xyz``
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.StringMatch.prefix
	Prefix *string `json:"prefix,omitempty"`

	// The input string must have the suffix specified here.
	//  Note: empty prefix is not allowed, please use regex instead.
	//
	//  Examples:
	//
	//  * ``abc`` matches the value ``xyz.abc``
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.StringMatch.suffix
	Suffix *string `json:"suffix,omitempty"`

	// The input string must have the substring specified here.
	//  Note: empty contains match is not allowed, please use regex instead.
	//
	//  Examples:
	//
	//  * ``abc`` matches the value ``xyz.abc.def``
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.StringMatch.contains
	Contains *string `json:"contains,omitempty"`

	// If true, indicates the exact/prefix/suffix/contains matching should be
	//  case insensitive. For example, the matcher ``data`` will match both
	//  input string ``Data`` and ``data`` if set to true.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.StringMatch.ignore_case
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To
type AuthzPolicy_AuthzRule_To struct {
	// Optional. Describes properties of one or more targets of a request. At
	//  least one of operations or notOperations must be specified. Limited to
	//  1 operation. A match occurs when ANY operation (in operations or
	//  notOperations) matches. Within an operation, the match follows AND
	//  semantics across fields and OR semantics within a field, i.e. a match
	//  occurs when ANY path matches AND ANY header matches and ANY method
	//  matches.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.operations
	Operations []AuthzPolicy_AuthzRule_To_RequestOperation `json:"operations,omitempty"`

	// Optional. Describes the negated properties of the targets of a request.
	//  Matches requests for operations that do not match the criteria
	//  specified in this field. At least one of operations or notOperations
	//  must be specified.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.not_operations
	NotOperations []AuthzPolicy_AuthzRule_To_RequestOperation `json:"notOperations,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation
type AuthzPolicy_AuthzRule_To_RequestOperation struct {
	// Optional. A list of headers to match against in http header.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.header_set
	HeaderSet *AuthzPolicy_AuthzRule_To_RequestOperation_HeaderSet `json:"headerSet,omitempty"`

	// Optional. A list of HTTP Hosts to match against. The match can be one
	//  of exact, prefix, suffix, or contains (substring match). Matches are
	//  always case sensitive unless the ignoreCase is set. Limited to 10
	//  hosts per Authorization Policy.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.hosts
	Hosts []AuthzPolicy_AuthzRule_StringMatch `json:"hosts,omitempty"`

	// Optional. A list of paths to match against. The match can be one of
	//  exact, prefix, suffix, or contains (substring match). Matches are
	//  always case sensitive unless the ignoreCase is set. Limited to 10
	//  paths per Authorization Policy.
	//  Note that this path match includes the query parameters. For gRPC
	//  services, this should be a fully-qualified name of the form
	//  /package.service/method.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.paths
	Paths []AuthzPolicy_AuthzRule_StringMatch `json:"paths,omitempty"`

	// Optional. A list of HTTP methods to match against. Each entry must be
	//  a valid HTTP method name (GET, PUT, POST, HEAD, PATCH, DELETE,
	//  OPTIONS). It only allows exact match and is always case sensitive.
	//  Limited to 10 methods per Authorization Policy.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.methods
	Methods []string `json:"methods,omitempty"`

	// Optional. Defines the MCP protocol attributes to match on. If the MCP
	//  payload in the request body cannot be successfully parsed, the
	//  request will be denied. This field can be set only for AuthzPolicies
	//  targeting AgentGateway resources.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.mcp
	Mcp *AuthzPolicy_AuthzRule_To_RequestOperation_Mcp `json:"mcp,omitempty"`

	// Optional. A list of SNIs to match against. The match can be one of
	//  exact, prefix, suffix, or contains (substring match). If there is no
	//  SNI (i.e. plaintext HTTP traffic), the request will be denied.
	//  Matches are always case sensitive unless the ignoreCase is set.
	//  Limited to 10 SNIs per Authorization Policy.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.snis
	Snis []AuthzPolicy_AuthzRule_StringMatch `json:"snis,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.HeaderSet
type AuthzPolicy_AuthzRule_To_RequestOperation_HeaderSet struct {
	// Required. A list of headers to match against in http header.
	//  The match can be one of exact, prefix, suffix, or contains
	//  (substring match). The match follows AND semantics which means all
	//  the headers must match. Matches are always case sensitive unless
	//  the ignoreCase is set. Limited to 10 headers per Authorization
	//  Policy.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.HeaderSet.headers
	// +kubebuilder:validation:Required
	Headers []AuthzPolicy_AuthzRule_HeaderMatch `json:"headers,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.MCP
type AuthzPolicy_AuthzRule_To_RequestOperation_Mcp struct {
	// Optional. If specified, matches on the MCP protocol’s non-access
	//  specific methods namely:
	//  * initialize
	//  * completion/
	//  * logging/
	//  * notifications/
	//  * ping
	//  Defaults to SKIP_BASE_PROTOCOL_METHODS if not specified.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.MCP.base_protocol_methods_option
	BaseProtocolMethodsOption *string `json:"baseProtocolMethodsOption,omitempty"`

	// Optional. A list of MCP methods and associated parameters to match
	//  on. It is recommended to use this field to match on tools, prompts
	//  and resource accesses while setting the baseProtocolMethodsOption
	//  to MATCH_BASE_PROTOCOL_METHODS to match on all the other MCP
	//  protocol methods.
	//  Limited to 10 MCP methods per Authorization Policy.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.MCP.methods
	Methods []AuthzPolicy_AuthzRule_To_RequestOperation_McpMethod `json:"methods,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.MCPMethod
type AuthzPolicy_AuthzRule_To_RequestOperation_McpMethod struct {
	// Required. The MCP method to match against. Allowed values are as
	//  follows:
	//  1. `tools`, `prompts`, `resources` - these will match against all
	//     sub methods under the respective methods.
	//  2. `prompts/list`, `tools/list`, `resources/list`,
	//     `resources/templates/list`
	//  3. `prompts/get`, `tools/call`, `resources/subscribe`,
	//     `resources/unsubscribe`, `resources/read`
	//  Params cannot be specified for categories 1 and 2.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.MCPMethod.name
	// +kubebuilder:validation:Required
	Name *string `json:"name,omitempty"`

	// Optional. A list of MCP method parameters to match against. The
	//  match can be one of exact, prefix, suffix, or contains (substring
	//  match). Matches are always case sensitive unless the ignoreCase is
	//  set. Limited to 10 MCP method parameters per Authorization Policy.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.AuthzRule.To.RequestOperation.MCPMethod.params
	Params []AuthzPolicy_AuthzRule_StringMatch `json:"params,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.CustomProvider
type AuthzPolicy_CustomProvider struct {
	// Optional. Delegates authorization decisions to Cloud IAP. Applicable
	//  only for managed load balancers. Enabling Cloud IAP at the AuthzPolicy
	//  level is not compatible with Cloud IAP settings in the BackendService.
	//  Enabling IAP in both places will result in request failure. Ensure that
	//  IAP is enabled in either the AuthzPolicy or the BackendService but not in
	//  both places.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.CustomProvider.cloud_iap
	CloudIAP *AuthzPolicy_CustomProvider_CloudIAP `json:"cloudIAP,omitempty"`

	// Optional. Delegate authorization decision to user authored Service
	//  Extension. Only one of cloudIap or authzExtension can be specified.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.CustomProvider.authz_extension
	AuthzExtension *AuthzPolicy_CustomProvider_AuthzExtension `json:"authzExtension,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.CustomProvider.AuthzExtension
type AuthzPolicy_CustomProvider_AuthzExtension struct {
	// Required. A list of references to authorization
	//  extensions that will be invoked for requests matching this policy.
	//  Limited to 1 custom provider.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.CustomProvider.AuthzExtension.resources
	// +kubebuilder:validation:Required
	Resources []AuthzPolicyCustomProviderAuthzExtensionResourceRef `json:"resources,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.CustomProvider.CloudIap
// +kubebuilder:pruning:PreserveUnknownFields
type AuthzPolicy_CustomProvider_CloudIAP struct {
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthzPolicy.Target
type AuthzPolicy_Target struct {
	// Optional. All gateways and forwarding rules referenced by this policy and
	//  extensions must share the same load balancing scheme. Required only when
	//  targeting forwarding rules. If targeting Secure Web Proxy, this field
	//  must be `INTERNAL_MANAGED` or not specified. Must not be specified
	//  when targeting Agent Gateway. Supported values:
	//  `INTERNAL_MANAGED` and `EXTERNAL_MANAGED`. For more information, refer
	//  to [Backend services
	//  overview](https://cloud.google.com/load-balancing/docs/backend-service).
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.Target.load_balancing_scheme
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	// Required. A list of references to the Forwarding Rules, Secure Web Proxy
	//  Gateways, or Agent Gateways on which this policy will be applied.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthzPolicy.Target.resources
	// +kubebuilder:validation:Required
	Resources []AuthzPolicyTargetResourceRef `json:"resources,omitempty"`
}

// AuthzPolicyTargetResourceRef is a reference to a resource that this policy applies to.
type AuthzPolicyTargetResourceRef struct {
	// A reference to an externally managed resource.
	// +optional
	External string `json:"external,omitempty"`

	// The name of a ConfigConnector resource.
	// +optional
	Name string `json:"name,omitempty"`

	// The namespace of a ConfigConnector resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// AuthzPolicyCustomProviderAuthzExtensionResourceRef is a reference to an authorization extension.
type AuthzPolicyCustomProviderAuthzExtensionResourceRef struct {
	// A reference to an externally managed ServiceExtensionAuthzExtension resource.
	// +optional
	External string `json:"external,omitempty"`

	// The name of a ServiceExtensionAuthzExtension resource.
	// +optional
	Name string `json:"name,omitempty"`

	// The namespace of a ServiceExtensionAuthzExtension resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}
