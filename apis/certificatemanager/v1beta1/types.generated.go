// Copyright 2024 Google LLC
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

package v1beta1

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate
type Certificate struct {
	// A user-defined name of the certificate. Certificate names must be unique
	//  globally and match pattern `projects/*/locations/*/certificates/*`.
	Name *string `json:"name,omitempty"`

	// One or more paragraphs of text description of a certificate.
	Description *string `json:"description,omitempty"`

	// Output only. The creation timestamp of a Certificate.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a Certificate.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Set of labels associated with a Certificate.
	Labels map[string]string `json:"labels,omitempty"`

	// If set, defines data of a self-managed certificate.
	SelfManaged *Certificate_SelfManagedCertificate `json:"selfManaged,omitempty"`

	// If set, contains configuration and state of a managed certificate.
	Managed *Certificate_ManagedCertificate `json:"managed,omitempty"`

	// Output only. The list of Subject Alternative Names of dnsName type defined
	//  in the certificate (see RFC 5280 4.2.1.6). Managed certificates that
	//  haven't been provisioned yet have this field populated with a value of the
	//  managed.domains field.
	SanDnsnames []string `json:"sanDnsnames,omitempty"`

	// Output only. The PEM-encoded certificate chain.
	PemCertificate *string `json:"pemCertificate,omitempty"`

	// Output only. The expiry timestamp of a Certificate.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Immutable. The scope of the certificate.
	Scope *string `json:"scope,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate
type Certificate_ManagedCertificate struct {
	// Immutable. The domains for which a managed SSL certificate will be
	//  generated. Wildcard domains are only supported with DNS challenge
	//  resolution.
	Domains []string `json:"domains,omitempty"`

	// Immutable. Authorizations that will be used for performing domain
	//  authorization.
	DnsAuthorizations []string `json:"dnsAuthorizations,omitempty"`

	// Immutable. The resource name for a
	//  [CertificateIssuanceConfig][google.cloud.certificatemanager.v1.CertificateIssuanceConfig]
	//  used to configure private PKI certificates in the format
	//  `projects/*/locations/*/certificateIssuanceConfigs/*`.
	//  If this field is not set, the certificates will instead be publicly
	//  signed as documented at
	//  https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.
	IssuanceConfig *string `json:"issuanceConfig,omitempty"`

	// Output only. State of the managed certificate resource.
	State *string `json:"state,omitempty"`

	// Output only. Information about issues with provisioning a Managed
	//  Certificate.
	ProvisioningIssue *Certificate_ManagedCertificate_ProvisioningIssue `json:"provisioningIssue,omitempty"`

	// Output only. Detailed state of the latest authorization attempt for each
	//  domain specified for managed certificate resource.
	AuthorizationAttemptInfo []Certificate_ManagedCertificate_AuthorizationAttemptInfo `json:"authorizationAttemptInfo,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.AuthorizationAttemptInfo
type Certificate_ManagedCertificate_AuthorizationAttemptInfo struct {
	// Domain name of the authorization attempt.
	Domain *string `json:"domain,omitempty"`

	// Output only. State of the domain for managed certificate issuance.
	State *string `json:"state,omitempty"`

	// Output only. Reason for failure of the authorization attempt for the
	//  domain.
	FailureReason *string `json:"failureReason,omitempty"`

	// Output only. Human readable explanation for reaching the state.
	//  Provided to help address the configuration issues. Not guaranteed to be
	//  stable. For programmatic access use FailureReason enum.
	Details *string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.ManagedCertificate.ProvisioningIssue
type Certificate_ManagedCertificate_ProvisioningIssue struct {
	// Output only. Reason for provisioning failures.
	Reason *string `json:"reason,omitempty"`

	// Output only. Human readable explanation about the issue. Provided to
	//  help address the configuration issues. Not guaranteed to be stable. For
	//  programmatic access use Reason enum.
	Details *string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.Certificate.SelfManagedCertificate
type Certificate_SelfManagedCertificate struct {
	// Input only. The PEM-encoded certificate chain.
	//  Leaf certificate comes first, followed by intermediate ones if any.
	PemCertificate *string `json:"pemCertificate,omitempty"`

	// Input only. The PEM-encoded private key of the leaf certificate.
	PemPrivateKey *string `json:"pemPrivateKey,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig
type CertificateIssuanceConfig struct {
	// A user-defined name of the certificate issuance config.
	//  CertificateIssuanceConfig names must be unique globally and match pattern
	//  `projects/*/locations/*/certificateIssuanceConfigs/*`.
	Name *string `json:"name,omitempty"`

	// Output only. The creation timestamp of a CertificateIssuanceConfig.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a CertificateIssuanceConfig.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Set of labels associated with a CertificateIssuanceConfig.
	Labels map[string]string `json:"labels,omitempty"`

	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	Description *string `json:"description,omitempty"`

	// Required. The CA that issues the workload certificate. It includes the CA
	//  address, type, authentication to CA service, etc.
	CertificateAuthorityConfig *CertificateIssuanceConfig_CertificateAuthorityConfig `json:"certificateAuthorityConfig,omitempty"`

	// Required. Workload certificate lifetime requested.
	Lifetime *string `json:"lifetime,omitempty"`

	// Required. Specifies the percentage of elapsed time of the certificate
	//  lifetime to wait before renewing the certificate. Must be a number between
	//  1-99, inclusive.
	RotationWindowPercentage *int32 `json:"rotationWindowPercentage,omitempty"`

	// Required. The key algorithm to use when generating the private key.
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig
type CertificateIssuanceConfig_CertificateAuthorityConfig struct {
	// Defines a CertificateAuthorityServiceConfig.
	CertificateAuthorityServiceConfig *CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig `json:"certificateAuthorityServiceConfig,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateIssuanceConfig.CertificateAuthorityConfig.CertificateAuthorityServiceConfig
type CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig struct {
	// Required. A CA pool resource used to issue a certificate.
	//  The CA pool string has a relative resource path following the form
	//  "projects/{project}/locations/{location}/caPools/{ca_pool}".
	CaPool *string `json:"caPool,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap
type CertificateMap struct {
	// A user-defined name of the Certificate Map. Certificate Map names must be
	//  unique globally and match pattern
	//  `projects/*/locations/*/certificateMaps/*`.
	Name *string `json:"name,omitempty"`

	// One or more paragraphs of text description of a certificate map.
	Description *string `json:"description,omitempty"`

	// Output only. The creation timestamp of a Certificate Map.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp of a Certificate Map.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Set of labels associated with a Certificate Map.
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. A list of GCLB targets that use this Certificate Map.
	//  A Target Proxy is only present on this list if it's attached to a
	//  Forwarding Rule.
	GclbTargets []CertificateMap_GclbTarget `json:"gclbTargets,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget
type CertificateMap_GclbTarget struct {
	// Output only. This field returns the resource name in the following
	//  format:
	//  `//compute.googleapis.com/projects/*/global/targetHttpsProxies/*`.
	TargetHTTPSProxy *string `json:"targetHTTPSProxy,omitempty"`

	// Output only. This field returns the resource name in the following
	//  format:
	//  `//compute.googleapis.com/projects/*/global/targetSslProxies/*`.
	TargetSslProxy *string `json:"targetSslProxy,omitempty"`

	// Output only. IP configurations for this Target Proxy where the
	//  Certificate Map is serving.
	IpConfigs []CertificateMap_GclbTarget_IpConfig `json:"ipConfigs,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.IpConfig
type CertificateMap_GclbTarget_IpConfig struct {
	// Output only. An external IP address.
	IpAddress *string `json:"ipAddress,omitempty"`

	// Output only. Ports.
	Ports []uint32 `json:"ports,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMapEntry
type CertificateMapEntry struct {
	// A user-defined name of the Certificate Map Entry. Certificate Map Entry
	//  names must be unique globally and match pattern
	//  `projects/*/locations/*/certificateMaps/*/certificateMapEntries/*`.
	Name *string `json:"name,omitempty"`

	// One or more paragraphs of text description of a certificate map entry.
	Description *string `json:"description,omitempty"`

	// Output only. The creation timestamp of a Certificate Map Entry.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp of a Certificate Map Entry.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Set of labels associated with a Certificate Map Entry.
	Labels map[string]string `json:"labels,omitempty"`

	// A Hostname (FQDN, e.g. `example.com`) or a wildcard hostname expression
	//  (`*.example.com`) for a set of hostnames with common suffix. Used as
	//  Server Name Indication (SNI) for selecting a proper certificate.
	Hostname *string `json:"hostname,omitempty"`

	// A predefined matcher for particular cases, other than SNI selection.
	Matcher *string `json:"matcher,omitempty"`

	// A set of Certificates defines for the given `hostname`. There can be
	//  defined up to four certificates in each Certificate Map Entry. Each
	//  certificate must match pattern `projects/*/locations/*/certificates/*`.
	Certificates []string `json:"certificates,omitempty"`

	// Output only. A serving state of this Certificate Map Entry.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.DnsAuthorization
type DnsAuthorization struct {
	// A user-defined name of the dns authorization. DnsAuthorization names must
	//  be unique globally and match pattern
	//  `projects/*/locations/*/dnsAuthorizations/*`.
	Name *string `json:"name,omitempty"`

	// Output only. The creation timestamp of a DnsAuthorization.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a DnsAuthorization.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Set of labels associated with a DnsAuthorization.
	Labels map[string]string `json:"labels,omitempty"`

	// One or more paragraphs of text description of a DnsAuthorization.
	Description *string `json:"description,omitempty"`

	// Required. Immutable. A domain that is being authorized. A DnsAuthorization
	//  resource covers a single domain and its wildcard, e.g. authorization for
	//  `example.com` can be used to issue certificates for `example.com` and
	//  `*.example.com`.
	Domain *string `json:"domain,omitempty"`

	// Output only. DNS Resource Record that needs to be added to DNS
	//  configuration.
	DnsResourceRecord *DnsAuthorization_DnsResourceRecord `json:"dnsResourceRecord,omitempty"`

	// Immutable. Type of DnsAuthorization. If unset during resource creation the
	//  following default will be used:
	//  - in location global: FIXED_RECORD.
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.DnsAuthorization.DnsResourceRecord
type DnsAuthorization_DnsResourceRecord struct {
	// Output only. Fully qualified name of the DNS Resource Record.
	//  e.g. `_acme-challenge.example.com`
	Name *string `json:"name,omitempty"`

	// Output only. Type of the DNS Resource Record.
	//  Currently always set to "CNAME".
	Type *string `json:"type,omitempty"`

	// Output only. Data of the DNS Resource Record.
	Data *string `json:"data,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig
type TrustConfig struct {
	// A user-defined name of the trust config. TrustConfig names must be
	//  unique globally and match pattern
	//  `projects/*/locations/*/trustConfigs/*`.
	Name *string `json:"name,omitempty"`

	// Output only. The creation timestamp of a TrustConfig.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of a TrustConfig.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Set of labels associated with a TrustConfig.
	Labels map[string]string `json:"labels,omitempty"`

	// One or more paragraphs of text description of a TrustConfig.
	Description *string `json:"description,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// Set of trust stores to perform validation against.
	//
	//  This field is supported when TrustConfig is configured with Load Balancers,
	//  currently not supported for SPIFFE certificate validation.
	//
	//  Only one TrustStore specified is currently allowed.
	TrustStores []TrustConfig_TrustStore `json:"trustStores,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig.IntermediateCA
type TrustConfig_IntermediateCA struct {
	// PEM intermediate certificate used for building up paths
	//  for validation.
	//
	//  Each certificate provided in PEM format may occupy up to 5kB.
	PemCertificate *string `json:"pemCertificate,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig.TrustAnchor
type TrustConfig_TrustAnchor struct {
	// PEM root certificate of the PKI used for validation.
	//
	//  Each certificate provided in PEM format may occupy up to 5kB.
	PemCertificate *string `json:"pemCertificate,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.TrustConfig.TrustStore
type TrustConfig_TrustStore struct {
	// List of Trust Anchors to be used while performing validation
	//  against a given TrustStore.
	TrustAnchors []TrustConfig_TrustAnchor `json:"trustAnchors,omitempty"`

	// Set of intermediate CA certificates used for the path building
	//  phase of chain validation.
	//
	//  The field is currently not supported if TrustConfig is used for the
	//  workload certificate feature.
	IntermediateCas []TrustConfig_IntermediateCA `json:"intermediateCas,omitempty"`
}
