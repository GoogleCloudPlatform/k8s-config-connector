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


// +kcc:proto=google.cloud.certificatemanager.v1.Certificate
type Certificate struct {
	// A user-defined name of the certificate. Certificate names must be unique
	//  globally and match pattern `projects/*/locations/*/certificates/*`.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.name
	Name *string `json:"name,omitempty"`

	// One or more paragraphs of text description of a certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.description
	Description *string `json:"description,omitempty"`

	// Set of labels associated with a Certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// If set, defines data of a self-managed certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.self_managed
	SelfManaged *Certificate_SelfManagedCertificate `json:"selfManaged,omitempty"`

	// If set, contains configuration and state of a managed certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.managed
	Managed *Certificate_ManagedCertificate `json:"managed,omitempty"`

	// Immutable. The scope of the certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.scope
	Scope *string `json:"scope,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate
type Certificate_ManagedCertificate struct {
	// Immutable. The domains for which a managed SSL certificate will be
	//  generated. Wildcard domains are only supported with DNS challenge
	//  resolution.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.domains
	Domains []string `json:"domains,omitempty"`

	// Immutable. Authorizations that will be used for performing domain
	//  authorization.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.dns_authorizations
	DnsAuthorizations []string `json:"dnsAuthorizations,omitempty"`

	// Immutable. The resource name for a
	//  [CertificateIssuanceConfig][google.cloud.certificatemanager.v1.CertificateIssuanceConfig]
	//  used to configure private PKI certificates in the format
	//  `projects/*/locations/*/certificateIssuanceConfigs/*`.
	//  If this field is not set, the certificates will instead be publicly
	//  signed as documented at
	//  https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.issuance_config
	IssuanceConfig *string `json:"issuanceConfig,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.AuthorizationAttemptInfo
type Certificate_ManagedCertificate_AuthorizationAttemptInfo struct {
	// Domain name of the authorization attempt.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.AuthorizationAttemptInfo.domain
	Domain *string `json:"domain,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.ProvisioningIssue
type Certificate_ManagedCertificate_ProvisioningIssue struct {
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.SelfManagedCertificate
type Certificate_SelfManagedCertificate struct {
	// Input only. The PEM-encoded certificate chain.
	//  Leaf certificate comes first, followed by intermediate ones if any.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.SelfManagedCertificate.pem_certificate
	PemCertificate *string `json:"pemCertificate,omitempty"`

	// Input only. The PEM-encoded private key of the leaf certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.SelfManagedCertificate.pem_private_key
	PemPrivateKey *string `json:"pemPrivateKey,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate
type CertificateObservedState struct {
	// Output only. The creation timestamp of a Certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a Certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// If set, contains configuration and state of a managed certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.managed
	Managed *Certificate_ManagedCertificateObservedState `json:"managed,omitempty"`

	// Output only. The list of Subject Alternative Names of dnsName type defined
	//  in the certificate (see RFC 5280 4.2.1.6). Managed certificates that
	//  haven't been provisioned yet have this field populated with a value of the
	//  managed.domains field.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.san_dnsnames
	SanDnsnames []string `json:"sanDnsnames,omitempty"`

	// Output only. The PEM-encoded certificate chain.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.pem_certificate
	PemCertificate *string `json:"pemCertificate,omitempty"`

	// Output only. The expiry timestamp of a Certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate
type Certificate_ManagedCertificateObservedState struct {
	// Output only. State of the managed certificate resource.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.state
	State *string `json:"state,omitempty"`

	// Output only. Information about issues with provisioning a Managed
	//  Certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.provisioning_issue
	ProvisioningIssue *Certificate_ManagedCertificate_ProvisioningIssue `json:"provisioningIssue,omitempty"`

	// Output only. Detailed state of the latest authorization attempt for each
	//  domain specified for managed certificate resource.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.authorization_attempt_info
	AuthorizationAttemptInfo []Certificate_ManagedCertificate_AuthorizationAttemptInfo `json:"authorizationAttemptInfo,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.AuthorizationAttemptInfo
type Certificate_ManagedCertificate_AuthorizationAttemptInfoObservedState struct {
	// Output only. State of the domain for managed certificate issuance.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.AuthorizationAttemptInfo.state
	State *string `json:"state,omitempty"`

	// Output only. Reason for failure of the authorization attempt for the
	//  domain.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.AuthorizationAttemptInfo.failure_reason
	FailureReason *string `json:"failureReason,omitempty"`

	// Output only. Human readable explanation for reaching the state.
	//  Provided to help address the configuration issues. Not guaranteed to be
	//  stable. For programmatic access use FailureReason enum.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.AuthorizationAttemptInfo.details
	Details *string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.ProvisioningIssue
type Certificate_ManagedCertificate_ProvisioningIssueObservedState struct {
	// Output only. Reason for provisioning failures.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.ProvisioningIssue.reason
	Reason *string `json:"reason,omitempty"`

	// Output only. Human readable explanation about the issue. Provided to
	//  help address the configuration issues. Not guaranteed to be stable. For
	//  programmatic access use Reason enum.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.ProvisioningIssue.details
	Details *string `json:"details,omitempty"`
}
