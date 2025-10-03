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
type ForwardingruleFilterLabels struct {
	/* Immutable. Name of the metadata label. The length must be between
	1 and 1024 characters, inclusive. */
	Name string `json:"name"`

	/* Immutable. The value that the label must match. The value has a maximum
	length of 1024 characters. */
	Value string `json:"value"`
}

type ForwardingruleIpAddress struct {
	// +optional
	AddressRef *refs.ComputeAddressRef `json:"addressRef,omitempty"`

	// +optional
	Ip *string `json:"ip,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.MetadataFilter
type ForwardingruleMetadataFilters struct {
	/* Immutable. The list of label value pairs that must match labels in the
	provided metadata based on filterMatchCriteria

	This list must not be empty and can have at the most 64 entries. */
	FilterLabels []ForwardingruleFilterLabels `json:"filterLabels"`

	/* Immutable. Specifies how individual filterLabel matches within the list of
	filterLabels contribute towards the overall metadataFilter match.

	MATCH_ANY - At least one of the filterLabels must have a matching
	label in the provided metadata.
	MATCH_ALL - All filterLabels must have matching labels in the
	provided metadata. Possible values: ["MATCH_ANY", "MATCH_ALL"]. */
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
	/* Immutable. This field can only be used:
	* If 'IPProtocol' is one of TCP, UDP, or SCTP.
	* By internal TCP/UDP load balancers, backend service-based network load
	balancers, and internal and external protocol forwarding.

	This option should be set to TRUE when the Forwarding Rule
	IPProtocol is set to L3_DEFAULT.

	Set this field to true to allow packets addressed to any port or packets
	lacking destination port information (for example, UDP fragments after the
	first fragment) to be forwarded to the backends configured with this
	forwarding rule.

	The 'ports', 'port_range', and
	'allPorts' fields are mutually exclusive. */
	// +optional
	AllPorts *bool `json:"allPorts,omitempty"`

	/* This field is used along with the 'backend_service' field for
	internal load balancing or with the 'target' field for internal
	TargetInstance.

	If the field is set to 'TRUE', clients can access ILB from all
	regions.

	Otherwise only allows access from clients in the same region as the
	internal load balancer. */
	// +optional
	AllowGlobalAccess *bool `json:"allowGlobalAccess,omitempty"`

	/* This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region. */
	// +optional
	AllowPscGlobalAccess *bool `json:"allowPscGlobalAccess,omitempty"`

	/* A ComputeBackendService to receive the matched traffic. This is
	used only for internal load balancing. */
	// +optional
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`

	/* Immutable. An optional description of this resource. Provide this property when
	you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The IP address that this forwarding rule is serving on behalf of.

	Addresses are restricted based on the forwarding rule's load
	balancing scheme (EXTERNAL or INTERNAL) and scope (global or
	regional).

	When the load balancing scheme is EXTERNAL, for global forwarding
	rules, the address must be a global IP, and for regional forwarding
	rules, the address must live in the same region as the forwarding
	rule. If this field is empty, an ephemeral IPv4 address from the
	same scope (global or regional) will be assigned. A regional
	forwarding rule supports IPv4 only. A global forwarding rule
	supports either IPv4 or IPv6.

	When the load balancing scheme is INTERNAL, this can only be an RFC
	1918 IP address belonging to the network/subnet configured for the
	forwarding rule. By default, if this field is empty, an ephemeral
	internal IP address will be automatically allocated from the IP
	range of the subnet or network configured for this forwarding rule. */
	// +optional
	IpAddress *ForwardingruleIpAddress `json:"ipAddress,omitempty"`

	/* Immutable. The IP protocol to which this rule applies.

	For protocol forwarding, valid
	options are 'TCP', 'UDP', 'ESP',
	'AH', 'SCTP', 'ICMP' and
	'L3_DEFAULT'.

	The valid IP protocols are different for different load balancing products
	as described in [Load balancing
	features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).

	A Forwarding Rule with protocol L3_DEFAULT can attach with target instance or
	backend service with UNSPECIFIED protocol.
	A forwarding rule with "L3_DEFAULT" IPProtocal cannot be attached to a backend service with TCP or UDP. Possible values: ["TCP", "UDP", "ESP", "AH", "SCTP", "ICMP", "L3_DEFAULT"]. */
	// +optional
	IpProtocol *string `json:"ipProtocol,omitempty"`

	/* Immutable. The IP address version that will be used by this forwarding rule.
	Valid options are IPV4 and IPV6.

	If not set, the IPv4 address will be used by default. Possible values: ["IPV4", "IPV6"]. */
	// +optional
	IpVersion *string `json:"ipVersion,omitempty"`

	/* Immutable. Indicates whether or not this load balancer can be used as a collector for
	packet mirroring. To prevent mirroring loops, instances behind this
	load balancer will not have their traffic mirrored even if a
	'PacketMirroring' rule applies to them.

	This can only be set to true for load balancers that have their
	'loadBalancingScheme' set to 'INTERNAL'. */
	// +optional
	IsMirroringCollector *bool `json:"isMirroringCollector,omitempty"`

	/* Immutable. Specifies the forwarding rule type.

	Must set to empty for private service connect forwarding rule. For more information about forwarding rules, refer to
	[Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts). Default value: "EXTERNAL" Possible values: ["EXTERNAL", "EXTERNAL_MANAGED", "INTERNAL", "INTERNAL_MANAGED", ""]. */
	// +optional
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	/* Location represents the geographical location of the ComputeForwardingRule. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* Immutable. Opaque filter criteria used by Loadbalancer to restrict routing
	configuration to a limited set xDS compliant clients. In their xDS
	requests to Loadbalancer, xDS clients present node metadata. If a
	match takes place, the relevant routing configuration is made available
	to those proxies.

	For each metadataFilter in this list, if its filterMatchCriteria is set
	to MATCH_ANY, at least one of the filterLabels must match the
	corresponding label provided in the metadata. If its filterMatchCriteria
	is set to MATCH_ALL, then all of its filterLabels must match with
	corresponding labels in the provided metadata.

	metadataFilters specified here can be overridden by those specified in
	the UrlMap that this ForwardingRule references.

	metadataFilters only applies to Loadbalancers that have their
	loadBalancingScheme set to INTERNAL_SELF_MANAGED. */
	// +optional
	MetadataFilters []ForwardingruleMetadataFilters `json:"metadataFilters,omitempty"`

	/* This field is not used for external load balancing. For internal
	load balancing, this field identifies the network that the load
	balanced IP should belong to for this forwarding rule. If this
	field is not specified, the default network will be used. */
	// +optional
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	/* Immutable. This signifies the networking tier used for configuring
	this load balancer and can only take the following values:
	'PREMIUM', 'STANDARD'.

	For regional ForwardingRule, the valid values are 'PREMIUM' and
	'STANDARD'. For GlobalForwardingRule, the valid value is
	'PREMIUM'.

	If this field is not specified, it is assumed to be 'PREMIUM'.
	If 'IPAddress' is specified, this value must be equal to the
	networkTier of the Address. Possible values: ["PREMIUM", "STANDARD"]. */
	// +optional
	NetworkTier *string `json:"networkTier,omitempty"`

	/* Immutable. This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding rules do not use this field. */
	// +optional
	NoAutomateDnsZone *bool `json:"noAutomateDnsZone,omitempty"`

	/* Immutable. This field can only be used:

	* If 'IPProtocol' is one of TCP, UDP, or SCTP.
	* By backend service-based network load balancers, target pool-based
	network load balancers, internal proxy load balancers, external proxy load
	balancers, Traffic Director, external protocol forwarding, and Classic VPN.
	Some products have restrictions on what ports can be used. See
	[port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)
	for details.


	Only packets addressed to ports in the specified range will be forwarded to
	the backends configured with this forwarding rule.

	The 'ports' and 'port_range' fields are mutually exclusive.

	For external forwarding rules, two or more forwarding rules cannot use the
	same '[IPAddress, IPProtocol]' pair, and cannot have
	overlapping 'portRange's.

	For internal forwarding rules within the same VPC network, two or more
	forwarding rules cannot use the same '[IPAddress, IPProtocol]'
	pair, and cannot have overlapping 'portRange's. */
	// +optional
	PortRange *string `json:"portRange,omitempty"`

	/* Immutable. This field can only be used:

	* If 'IPProtocol' is one of TCP, UDP, or SCTP.
	* By internal TCP/UDP load balancers, backend service-based network load
	balancers, internal protocol forwarding and when protocol is not L3_DEFAULT.


	You can specify a list of up to five ports by number, separated by commas.
	The ports can be contiguous or discontiguous. Only packets addressed to
	these ports will be forwarded to the backends configured with this
	forwarding rule.

	For external forwarding rules, two or more forwarding rules cannot use the
	same '[IPAddress, IPProtocol]' pair, and cannot share any values
	defined in 'ports'.

	For internal forwarding rules within the same VPC network, two or more
	forwarding rules cannot use the same '[IPAddress, IPProtocol]'
	pair, and cannot share any values defined in 'ports'.

	The 'ports' and 'port_range' fields are mutually exclusive. */
	// +optional
	Ports []string `json:"ports,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Service Directory resources to register this forwarding rule with.

	Currently, only supports a single Service Directory resource. */
	// +optional
	ServiceDirectoryRegistrations []ForwardingruleServiceDirectoryRegistrations `json:"serviceDirectoryRegistrations,omitempty"`

	/* Immutable. An optional prefix to the service name for this Forwarding Rule.
	If specified, will be the first label of the fully qualified service
	name.

	The label must be 1-63 characters long, and comply with RFC1035.
	Specifically, the label must be 1-63 characters long and match the
	regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	character must be a lowercase letter, and all following characters
	must be a dash, lowercase letter, or digit, except the last
	character, which cannot be a dash.

	This field is only used for INTERNAL load balancing. */
	// +optional
	ServiceLabel *string `json:"serviceLabel,omitempty"`

	/* Immutable. If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24). */
	// +optional
	SourceIpRanges []string `json:"sourceIpRanges,omitempty"`

	/* Immutable. The subnetwork that the load balanced IP should belong to for this
	forwarding rule. This field is only used for internal load
	balancing.

	If the network specified is in auto subnet mode, this field is
	optional. However, if the network is in custom subnet mode, a
	subnetwork must be specified. */
	// +optional
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	/* The target resource to receive the matched traffic. The forwarded
	traffic must be of a type appropriate to the target object. For
	INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
	are valid. */
	// +optional
	Target *ForwardingruleTarget `json:"target,omitempty"`
}

// +kcc:status:proto=google.cloud.compute.v1.ForwardingRule
type ComputeForwardingRuleStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`
	/* [Output Only] The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified. */
	// +optional
	BaseForwardingRule *string `json:"baseForwardingRule,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* The fingerprint used for optimistic locking of this resource.  Used
	internally during updates. */
	// +optional
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	/* The PSC connection id of the PSC Forwarding Rule. */
	// +optional
	PscConnectionId *string `json:"pscConnectionId,omitempty"`

	/* The PSC connection status of the PSC Forwarding Rule. Possible values: 'STATUS_UNSPECIFIED', 'PENDING', 'ACCEPTED', 'REJECTED', 'CLOSED'. */
	// +optional
	PscConnectionStatus *string `json:"pscConnectionStatus,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* The internal fully qualified service name for this Forwarding Rule.

	This field is only used for INTERNAL load balancing. */
	// +optional
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
