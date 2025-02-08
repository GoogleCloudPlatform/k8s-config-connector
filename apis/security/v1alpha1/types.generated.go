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


// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateRevocationList
type CertificateRevocationList struct {

	// Optional. Labels with user-defined metadata.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateRevocationList.RevokedCertificate
type CertificateRevocationList_RevokedCertificate struct {
	// The resource path for the [Certificate][google.cloud.security.privateca.v1beta1.Certificate] in the format
	//  `projects/*/locations/*/certificateAuthorities/*/certificates/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.RevokedCertificate.certificate
	Certificate *string `json:"certificate,omitempty"`

	// The serial number of the [Certificate][google.cloud.security.privateca.v1beta1.Certificate].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.RevokedCertificate.hex_serial_number
	HexSerialNumber *string `json:"hexSerialNumber,omitempty"`

	// The reason the [Certificate][google.cloud.security.privateca.v1beta1.Certificate] was revoked.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.RevokedCertificate.revocation_reason
	RevocationReason *string `json:"revocationReason,omitempty"`
}

// +kcc:proto=google.cloud.security.privateca.v1beta1.CertificateRevocationList
type CertificateRevocationListObservedState struct {
	// Output only. The resource path for this [CertificateRevocationList][google.cloud.security.privateca.v1beta1.CertificateRevocationList] in
	//  the format
	//  `projects/*/locations/*/certificateAuthorities/*/
	//     certificateRevocationLists/*`.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.name
	Name *string `json:"name,omitempty"`

	// Output only. The CRL sequence number that appears in pem_crl.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.sequence_number
	SequenceNumber *int64 `json:"sequenceNumber,omitempty"`

	// Output only. The revoked serial numbers that appear in pem_crl.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.revoked_certificates
	RevokedCertificates []CertificateRevocationList_RevokedCertificate `json:"revokedCertificates,omitempty"`

	// Output only. The PEM-encoded X.509 CRL.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.pem_crl
	PemCrl *string `json:"pemCrl,omitempty"`

	// Output only. The location where 'pem_crl' can be accessed.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.access_url
	AccessURL *string `json:"accessURL,omitempty"`

	// Output only. The [State][google.cloud.security.privateca.v1beta1.CertificateRevocationList.State] for this [CertificateRevocationList][google.cloud.security.privateca.v1beta1.CertificateRevocationList].
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.state
	State *string `json:"state,omitempty"`

	// Output only. The time at which this [CertificateRevocationList][google.cloud.security.privateca.v1beta1.CertificateRevocationList] was created.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this [CertificateRevocationList][google.cloud.security.privateca.v1beta1.CertificateRevocationList] was updated.
	// +kcc:proto:field=google.cloud.security.privateca.v1beta1.CertificateRevocationList.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
