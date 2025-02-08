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


// +kcc:proto=google.cloud.security.privateca.v1beta1.KeyUsage
type KeyUsage struct {
	// Describes high-level ways in which a key may be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.base_key_usage
	BaseKeyUsage *KeyUsage_KeyUsageOptions `json:"baseKeyUsage,omitempty"`

	// Detailed scenarios in which a key may be used.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.extended_key_usage
	ExtendedKeyUsage *KeyUsage_ExtendedKeyUsageOptions `json:"extendedKeyUsage,omitempty"`

	// Used to describe extended key usages that are not listed in the
	//  [KeyUsage.ExtendedKeyUsageOptions][google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions] message.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.unknown_extended_key_usages
	UnknownExtendedKeyUsages []ObjectId `json:"unknownExtendedKeyUsages,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions
type KeyUsage_ExtendedKeyUsageOptions struct {
	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW
	//  server authentication", though regularly used for non-WWW TLS.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.server_auth
	ServerAuth *bool `json:"serverAuth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW
	//  client authentication", though regularly used for non-WWW TLS.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.client_auth
	ClientAuth *bool `json:"clientAuth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of
	//  downloadable executable code client authentication".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.code_signing
	CodeSigning *bool `json:"codeSigning,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email
	//  protection".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.email_protection
	EmailProtection *bool `json:"emailProtection,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding
	//  the hash of an object to a time".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.time_stamping
	TimeStamping *bool `json:"timeStamping,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing
	//  OCSP responses".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.ExtendedKeyUsageOptions.ocsp_signing
	OcspSigning *bool `json:"ocspSigning,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions
type KeyUsage_KeyUsageOptions struct {
	// The key may be used for digital signatures.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.digital_signature
	DigitalSignature *bool `json:"digitalSignature,omitempty"`

	// The key may be used for cryptographic commitments. Note that this may
	//  also be referred to as "non-repudiation".
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.content_commitment
	ContentCommitment *bool `json:"contentCommitment,omitempty"`

	// The key may be used to encipher other keys.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.key_encipherment
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`

	// The key may be used to encipher data.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.data_encipherment
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`

	// The key may be used in a key agreement protocol.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.key_agreement
	KeyAgreement *bool `json:"keyAgreement,omitempty"`

	// The key may be used to sign certificates.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.cert_sign
	CertSign *bool `json:"certSign,omitempty"`

	// The key may be used sign certificate revocation lists.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.crl_sign
	CrlSign *bool `json:"crlSign,omitempty"`

	// The key may be used to encipher only.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.encipher_only
	EncipherOnly *bool `json:"encipherOnly,omitempty"`

	// The key may be used to decipher only.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.KeyUsage.KeyUsageOptions.decipher_only
	DecipherOnly *bool `json:"decipherOnly,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ObjectId
type ObjectId struct {
	// Required. The parts of an OID path. The most significant parts of the path come
	//  first.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ObjectId.object_id_path
	ObjectIDPath []int32 `json:"objectIDPath,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ReusableConfig
type ReusableConfig struct {

	// Required. The config values.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfig.values
	Values *ReusableConfigValues `json:"values,omitempty"`

	// Optional. A human-readable description of scenarios these ReusableConfigValues may be
	//  compatible with.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfig.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels with user-defined metadata.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfig.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ReusableConfigValues
type ReusableConfigValues struct {
	// Optional. Indicates the intended use for keys that correspond to a certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.key_usage
	KeyUsage *KeyUsage `json:"keyUsage,omitempty"`

	// Optional. Describes options in this [ReusableConfigValues][google.cloud.security.privateca.v1beta1.ReusableConfigValues] that are
	//  relevant in a CA certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.ca_options
	CaOptions *ReusableConfigValues_CaOptions `json:"caOptions,omitempty"`

	// Optional. Describes the X.509 certificate policy object identifiers, per
	//  https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.policy_ids
	PolicyIds []ObjectId `json:"policyIds,omitempty"`

	// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses
	//  that appear in the "Authority Information Access" extension in the
	//  certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.aia_ocsp_servers
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	// Optional. Describes custom X.509 extensions.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.additional_extensions
	AdditionalExtensions []X509Extension `json:"additionalExtensions,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ReusableConfigValues.CaOptions
type ReusableConfigValues_CaOptions struct {
	// Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this
	//  value is missing, the extension will be omitted from the CA certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.CaOptions.is_ca
	IsCa *bool `json:"isCa,omitempty"`

	// Optional. Refers to the path length restriction X.509 extension. For a CA
	//  certificate, this value describes the depth of subordinate CA
	//  certificates that are allowed.
	//  If this value is less than 0, the request will fail.
	//  If this value is missing, the max path length will be omitted from the
	//  CA certificate.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfigValues.CaOptions.max_issuer_path_length
	MaxIssuerPathLength *Int32Value `json:"maxIssuerPathLength,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.X509Extension
type X509Extension struct {
	// Required. The OID for this X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.X509Extension.object_id
	ObjectID *ObjectId `json:"objectID,omitempty"`

	// Required. Indicates whether or not this extension is critical (i.e., if the client
	//  does not know how to handle this extension, the client should consider this
	//  to be an error).
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.X509Extension.critical
	Critical *bool `json:"critical,omitempty"`

	// Required. The value of this X.509 extension.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.X509Extension.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.Int32Value
type Int32Value struct {
	// The int32 value.
	// +kcc:proto:field=google.protobuf.Int32Value.value
	Value *int32 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.ReusableConfig
type ReusableConfigObservedState struct {
	// Output only. The resource path for this [ReusableConfig][google.cloud.security.privateca.v1beta1.ReusableConfig] in the format
	//  `projects/*/locations/*/reusableConfigs/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfig.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which this [ReusableConfig][google.cloud.security.privateca.v1beta1.ReusableConfig] was created.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this [ReusableConfig][google.cloud.security.privateca.v1beta1.ReusableConfig] was updated.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.ReusableConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
