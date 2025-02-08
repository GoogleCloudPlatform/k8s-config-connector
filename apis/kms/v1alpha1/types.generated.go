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


// +kcc:proto=google.cloud.kms.v1.PublicKey
type PublicKey struct {
	// The public key, encoded in PEM format. For more information, see the
	//  [RFC 7468](https://tools.ietf.org/html/rfc7468) sections for
	//  [General Considerations](https://tools.ietf.org/html/rfc7468#section-2) and
	//  [Textual Encoding of Subject Public Key Info]
	//  (https://tools.ietf.org/html/rfc7468#section-13).
	// +kcc:proto:field=google.cloud.kms.v1.PublicKey.pem
	Pem *string `json:"pem,omitempty"`

	// The
	//  [Algorithm][google.cloud.kms.v1.CryptoKeyVersion.CryptoKeyVersionAlgorithm]
	//  associated with this key.
	// +kcc:proto:field=google.cloud.kms.v1.PublicKey.algorithm
	Algorithm *string `json:"algorithm,omitempty"`

	// Integrity verification field. A CRC32C checksum of the returned
	//  [PublicKey.pem][google.cloud.kms.v1.PublicKey.pem]. An integrity check of
	//  [PublicKey.pem][google.cloud.kms.v1.PublicKey.pem] can be performed by
	//  computing the CRC32C checksum of
	//  [PublicKey.pem][google.cloud.kms.v1.PublicKey.pem] and comparing your
	//  results to this field. Discard the response in case of non-matching
	//  checksum values, and perform a limited number of retries. A persistent
	//  mismatch may indicate an issue in your computation of the CRC32C checksum.
	//  Note: This field is defined as int64 for reasons of compatibility across
	//  different languages. However, it is a non-negative integer, which will
	//  never exceed 2^32-1, and can be safely downconverted to uint32 in languages
	//  that support this type.
	//
	//  NOTE: This field is in Beta.
	// +kcc:proto:field=google.cloud.kms.v1.PublicKey.pem_crc32c
	PemCrc32c *int64 `json:"pemCrc32c,omitempty"`

	// The [name][google.cloud.kms.v1.CryptoKeyVersion.name] of the
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] public key.
	//  Provided here for verification.
	//
	//  NOTE: This field is in Beta.
	// +kcc:proto:field=google.cloud.kms.v1.PublicKey.name
	Name *string `json:"name,omitempty"`

	// The [ProtectionLevel][google.cloud.kms.v1.ProtectionLevel] of the
	//  [CryptoKeyVersion][google.cloud.kms.v1.CryptoKeyVersion] public key.
	// +kcc:proto:field=google.cloud.kms.v1.PublicKey.protection_level
	ProtectionLevel *string `json:"protectionLevel,omitempty"`
}
