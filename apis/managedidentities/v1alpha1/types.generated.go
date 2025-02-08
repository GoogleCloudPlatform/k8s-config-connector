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


// +kcc:proto=google.cloud.managedidentities.v1.Domain
type Domain struct {
	// Required. The unique name of the domain using the form:
	//  `projects/{project_id}/locations/global/domains/{domain_name}`.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.name
	Name *string `json:"name,omitempty"`

	// Optional. Resource labels that can contain user-provided metadata.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The full names of the Google Compute Engine
	//  [networks](/compute/docs/networks-and-firewalls#networks) the domain
	//  instance is connected to. Networks can be added using UpdateDomain.
	//  The domain is only available on networks listed in `authorized_networks`.
	//  If CIDR subnets overlap between networks, domain creation will fail.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.authorized_networks
	AuthorizedNetworks []string `json:"authorizedNetworks,omitempty"`

	// Required. The CIDR range of internal addresses that are reserved for this
	//  domain. Reserved networks must be /24 or larger. Ranges must be
	//  unique and non-overlapping with existing subnets in
	//  [Domain].[authorized_networks].
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.reserved_ip_range
	ReservedIPRange *string `json:"reservedIPRange,omitempty"`

	// Required. Locations where domain needs to be provisioned.
	//  [regions][compute/docs/regions-zones/]
	//  e.g. us-west1 or us-east4
	//  Service supports up to 4 locations at once. Each location will use a /26
	//  block.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.locations
	Locations []string `json:"locations,omitempty"`

	// Optional. The name of delegated administrator account used to perform
	//  Active Directory operations. If not specified, `setupadmin` will be used.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.admin
	Admin *string `json:"admin,omitempty"`
}

// +kcc:proto=google.cloud.managedidentities.v1.Trust
type Trust struct {
	// Required. The fully qualified target domain name which will be in trust with the
	//  current domain.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.target_domain_name
	TargetDomainName *string `json:"targetDomainName,omitempty"`

	// Required. The type of trust represented by the trust resource.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.trust_type
	TrustType *string `json:"trustType,omitempty"`

	// Required. The trust direction, which decides if the current domain is trusted,
	//  trusting, or both.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.trust_direction
	TrustDirection *string `json:"trustDirection,omitempty"`

	// Optional. The trust authentication type, which decides whether the trusted side has
	//  forest/domain wide access or selective access to an approved set of
	//  resources.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.selective_authentication
	SelectiveAuthentication *bool `json:"selectiveAuthentication,omitempty"`

	// Required. The target DNS server IP addresses which can resolve the remote domain
	//  involved in the trust.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.target_dns_ip_addresses
	TargetDnsIPAddresses []string `json:"targetDnsIPAddresses,omitempty"`

	// Required. The trust secret used for the handshake with the target domain. This will
	//  not be stored.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.trust_handshake_secret
	TrustHandshakeSecret *string `json:"trustHandshakeSecret,omitempty"`
}

// +kcc:proto=google.cloud.managedidentities.v1.Domain
type DomainObservedState struct {
	// Output only. The fully-qualified domain name of the exposed domain used by
	//  clients to connect to the service. Similar to what would be chosen for an
	//  Active Directory set up on an internal network.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.fqdn
	Fqdn *string `json:"fqdn,omitempty"`

	// Output only. The time the instance was created.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update time.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of this domain.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current status of this
	//  domain, if available.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.status_message
	StatusMessage *string `json:"statusMessage,omitempty"`

	// Output only. The current trusts associated with the domain.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Domain.trusts
	Trusts []Trust `json:"trusts,omitempty"`
}

// +kcc:proto=google.cloud.managedidentities.v1.Trust
type TrustObservedState struct {
	// Output only. The time the instance was created.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update time.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of the trust.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current state of the trust, if available.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.state_description
	StateDescription *string `json:"stateDescription,omitempty"`

	// Output only. The last heartbeat time when the trust was known to be connected.
	// +kcc:proto:field=google.cloud.managedidentities.v1.Trust.last_trust_heartbeat_time
	LastTrustHeartbeatTime *string `json:"lastTrustHeartbeatTime,omitempty"`
}
