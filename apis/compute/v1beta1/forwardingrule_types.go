/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	ComputeForwardingRuleGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    "ComputeForwardingRule",
	}
)

// +kcc:proto=google.cloud.compute.v1.MetadataFilterLabelMatch
type MetadataFilterLabelMatch struct {
	// Immutable. Name of metadata label. The name can have a maximum length of 1024 characters and must be at least 1 character long.
	// +kcc:proto:field=google.cloud.compute.v1.MetadataFilterLabelMatch.name
	Name string `json:"name"`

	// Immutable. The value of the label must match the specified value. value can have a maximum length of 1024 characters.
	// +kcc:proto:field=google.cloud.compute.v1.MetadataFilterLabelMatch.value
	Value string `json:"value"`
}

type ForwardingruleIpAddress struct {
	// +optional
	AddressRef *refs.ComputeAddressRef `json:"addressRef,omitempty"`

	// +optional
	Ip *string `json:"ip,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.MetadataFilter
type MetadataFilter struct {
	// Immutable. The list of label value pairs that must match labels in the provided metadata based on filterMatchCriteria This list must not be empty and can have at the most 64 entries.
	// +kcc:proto:field=google.cloud.compute.v1.MetadataFilter.filter_labels
	FilterLabels []MetadataFilterLabelMatch `json:"filterLabels"`

	// Immutable. Specifies how individual filter label matches within the list of filterLabels and contributes toward the overall metadataFilter match. Supported values are: - MATCH_ANY: at least one of the filterLabels must have a matching label in the provided metadata. - MATCH_ALL: all filterLabels must have matching labels in the provided metadata.
	// Check the FilterMatchCriteria enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.MetadataFilter.filter_match_criteria
	FilterMatchCriteria string `json:"filterMatchCriteria"`
}

// +kcc:proto=google.cloud.compute.v1.ForwardingRuleServiceDirectoryRegistration
type ForwardingruleServiceDirectoryRegistrations struct {
	/* Immutable. Service Directory namespace to register the forwarding rule under. */
	// +optional
	Namespace *string `json:"namespace,omitempty"`

	/* Immutable. Service Directory service to register the forwarding rule under. */
	// +optional
	Service *string `json:"service,omitempty"`
}

type ForwardingruleTarget struct {
	// +optional
	GoogleAPIsBundle *string `json:"googleAPIsBundle,omitempty"`

	// +optional
	ServiceAttachmentRef *refs.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`

	// +optional
	TargetGRPCProxyRef *refs.ComputeTargetGrpcProxyRef `json:"targetGRPCProxyRef,omitempty"`

	// +optional
	TargetHTTPProxyRef *refs.ComputeTargetHTTPProxyRef `json:"targetHTTPProxyRef,omitempty"`

	// +optional
	TargetHTTPSProxyRef *refs.ComputeTargetHTTPSProxyRef `json:"targetHTTPSProxyRef,omitempty"`

	// +optional
	TargetSSLProxyRef *refs.ComputeTargetSSLProxyRef `json:"targetSSLProxyRef,omitempty"`

	// +optional
	TargetTCPProxyRef *refs.ComputeTargetTCPProxyRef `json:"targetTCPProxyRef,omitempty"`

	// +optional
	TargetVPNGatewayRef *refs.ComputeTargetVPNGatewayRef `json:"targetVPNGatewayRef,omitempty"`
}

// +kcc:spec:proto=google.cloud.compute.v1.ForwardingRule
type ComputeForwardingRuleSpec struct {
	// Immutable. The ports, portRange, and allPorts fields are mutually exclusive. Only packets addressed to ports in the specified range will be forwarded to the backends configured with this forwarding rule. The allPorts field has the following limitations: - It requires that the forwarding rule IPProtocol be TCP, UDP, SCTP, or L3_DEFAULT. - It's applicable only to the following products: internal passthrough Network Load Balancers, backend service-based external passthrough Network Load Balancers, and internal and external protocol forwarding. - Set this field to true to allow packets addressed to any port or packets lacking destination port information (for example, UDP fragments after the first fragment) to be forwarded to the backends configured with this forwarding rule. The L3_DEFAULT protocol requires allPorts be set to true.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.all_ports
	AllPorts *bool `json:"allPorts,omitempty"`

	// If set to true, clients can access the internal passthrough Network Load Balancers, the regional internal Application Load Balancer, and the regional internal proxy Network Load Balancer from all regions. If false, only allows access from the local region the load balancer is located at. Note that for INTERNAL_MANAGED forwarding rules, this field cannot be changed after the forwarding rule is created.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.allow_global_access
	AllowGlobalAccess *bool `json:"allowGlobalAccess,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.allow_psc_global_access
	AllowPscGlobalAccess *bool `json:"allowPscGlobalAccess,omitempty"`

	// Identifies the backend service to which the forwarding rule sends traffic. Required for internal and external passthrough Network Load Balancers; must be omitted for all other load balancer types.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.backend_service
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`

	// Immutable. An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.description
	Description *string `json:"description,omitempty"`

	// IP address for which this forwarding rule accepts traffic. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the referenced target or backendService. While creating a forwarding rule, specifying an IPAddress is required under the following circumstances: - When the target is set to targetGrpcProxy and validateForProxyless is set to true, the IPAddress should be set to 0.0.0.0. - When the target is a Private Service Connect Google APIs bundle, you must specify an IPAddress. Otherwise, you can optionally specify an IP address that references an existing static (reserved) IP address resource. When omitted, Google Cloud assigns an ephemeral IP address. Use one of the following formats to specify an IP address while creating a forwarding rule: * IP address number, as in `100.1.2.3` * IPv6 address range, as in `2600:1234::/96` * Full resource URL, as in https://www.googleapis.com/compute/v1/projects/ project_id/regions/region/addresses/address-name * Partial URL or by name, as in: - projects/project_id/regions/region/addresses/address-name - regions/region/addresses/address-name - global/addresses/address-name - address-name The forwarding rule's target or backendService, and in most cases, also the loadBalancingScheme, determine the type of IP address that you can use. For detailed information, see [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). When reading an IPAddress, the API always returns the IP address number.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.I_p_address
	IpAddress *ForwardingruleIpAddress `json:"ipAddress,omitempty"`

	// Immutable. The IP protocol to which this rule applies. For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP, ICMP and L3_DEFAULT. The valid IP protocols are different for different load balancing products as described in [Load balancing features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
	// Check the IPProtocolEnum enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.I_p_protocol
	IpProtocol *string `json:"ipProtocol,omitempty"`

	// Immutable. The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6.
	// Check the IpVersion enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.ip_version
	IpVersion *string `json:"ipVersion,omitempty"`

	// Immutable. Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.is_mirroring_collector
	IsMirroringCollector *bool `json:"isMirroringCollector,omitempty"`

	// Immutable. Specifies the forwarding rule type. For more information about forwarding rules, refer to Forwarding rule concepts.
	// Default value: "EXTERNAL" Possible values: ["EXTERNAL", "EXTERNAL_MANAGED", "INTERNAL", "INTERNAL_MANAGED", ""].
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.load_balancing_scheme
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	// Location represents the geographical location of the ComputeForwardingRule. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	// Immutable. Opaque filter criteria used by load balancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to load balancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule are not visible to those proxies. For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match. metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references. metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.metadata_filters
	MetadataFilters []MetadataFilter `json:"metadataFilters,omitempty"`

	// This field is not used for global external load balancing. For internal passthrough Network Load Balancers, this field identifies the network that the load balanced IP should belong to for this forwarding rule. If the subnetwork is specified, the network of the subnetwork will be used. If neither subnetwork nor this field is specified, the default network will be used. For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.network
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Immutable. This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD. For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM. If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
	//  Check the NetworkTier enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.network_tier
	NetworkTier *string `json:"networkTier,omitempty"`

	// Immutable. This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field. Once set, this field is not mutable.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.no_automate_dns_zone
	NoAutomateDnsZone *bool `json:"noAutomateDnsZone,omitempty"`

	// Immutable. The ports, portRange, and allPorts fields are mutually exclusive. Only packets addressed to ports in the specified range will be forwarded to the backends configured with this forwarding rule. The portRange field has the following limitations: - It requires that the forwarding rule IPProtocol be TCP, UDP, or SCTP, and - It's applicable only to the following products: external passthrough Network Load Balancers, internal and external proxy Network Load Balancers, internal and external Application Load Balancers, external protocol forwarding, and Classic VPN. - Some products have restrictions on what ports can be used. See port specifications for details. For external forwarding rules, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot have overlapping portRanges. For internal forwarding rules within the same VPC network, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair, and cannot have overlapping portRanges. @pattern: \\d+(?:-\\d+)?
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.port_range
	PortRange *string `json:"portRange,omitempty"`

	// Immutable. The ports, portRange, and allPorts fields are mutually exclusive. Only packets addressed to ports in the specified range will be forwarded to the backends configured with this forwarding rule. The ports field has the following limitations: - It requires that the forwarding rule IPProtocol be TCP, UDP, or SCTP, and - It's applicable only to the following products: internal passthrough Network Load Balancers, backend service-based external passthrough Network Load Balancers, and internal protocol forwarding. - You can specify a list of up to five ports by number, separated by commas. The ports can be contiguous or discontiguous. For external forwarding rules, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair if they share at least one port number. For internal forwarding rules within the same VPC network, two or more forwarding rules cannot use the same [IPAddress, IPProtocol] pair if they share at least one port number. @pattern: \\d+(?:-\\d+)?
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.ports
	Ports []string `json:"ports,omitempty"`

	// Immutable. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.service_directory_registrations
	ServiceDirectoryRegistrations []ForwardingruleServiceDirectoryRegistrations `json:"serviceDirectoryRegistrations,omitempty"`

	// Immutable. An optional prefix to the service name for this forwarding rule. If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.service_label
	ServiceLabel *string `json:"serviceLabel,omitempty"`

	// Immutable. If not empty, this forwarding rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a forwarding rule can only have up to 64 source IP ranges, and this field can only be used with a regional forwarding rule whose scheme is EXTERNAL. Each source_ip_range entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.source_ip_ranges
	SourceIpRanges []string `json:"sourceIpRanges,omitempty"`

	// Immutable. This field identifies the subnetwork that the load balanced IP should belong to for this forwarding rule, used with internal load balancers and external passthrough Network Load Balancers with IPv6. If the network specified is in auto subnet mode, this field is optional. However, a subnetwork must be specified if the network is in custom subnet mode or when creating external forwarding rule with IPv6.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.subnetwork
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must be in the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing resource. The forwarded traffic must be of a type appropriate to the target object. - For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). - For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle: - vpc-sc - APIs that support VPC Service Controls. - all-apis - All supported Google APIs. - For Private Service Connect forwarding rules that forward traffic to managed services, the target must be a service attachment. The target is not mutable once set as a service attachment.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.target
	Target *ForwardingruleTarget `json:"target,omitempty"`
}

// +kcc:status:proto=google.cloud.compute.v1.ForwardingRule
type ComputeForwardingRuleStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`
	// [Output Only] The URL for the corresponding base forwarding rule. By base forwarding rule, we mean the forwarding rule that has the same IP address, protocol, and port settings with the current forwarding rule, but without sourceIPRanges specified. Always empty if the current forwarding rule does not have sourceIPRanges specified.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.base_forwarding_rule
	BaseForwardingRule *string `json:"baseForwardingRule,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.label_fingerprint
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	// [Output Only] The PSC connection id of the PSC forwarding rule.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.psc_connection_id
	PscConnectionId *string `json:"pscConnectionId,omitempty"`

	// The PSC connection status of the PSC Forwarding Rule. Possible values: 'STATUS_UNSPECIFIED', 'PENDING', 'ACCEPTED', 'REJECTED', 'CLOSED'.
	PscConnectionStatus *string `json:"pscConnectionStatus,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] The internal fully qualified service name for this forwarding rule. This field is only used for internal load balancing.
	// +kcc:proto:field=google.cloud.compute.v1.ForwardingRule.service_name
	ServiceName *string `json:"serviceName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeforwardingrule;gcpcomputeforwardingrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeForwardingRule is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeForwardingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeForwardingRuleSpec   `json:"spec,omitempty"`
	Status ComputeForwardingRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeForwardingRuleList contains a list of ComputeForwardingRule
type ComputeForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeForwardingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeForwardingRule{}, &ComputeForwardingRuleList{})
}
