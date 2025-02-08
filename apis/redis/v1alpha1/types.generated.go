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


// +kcc:proto=google.cloud.redis.cluster.v1beta1.CertificateAuthority
type CertificateAuthority struct {
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.CertificateAuthority.managed_server_ca
	ManagedServerCa *CertificateAuthority_ManagedCertificateAuthority `json:"managedServerCa,omitempty"`

	// Identifier. Unique name of the resource in this scope including project,
	//  location and cluster using the form:
	//      `projects/{project}/locations/{location}/clusters/{cluster}/certificateAuthority`
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.CertificateAuthority.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1beta1.CertificateAuthority.ManagedCertificateAuthority
type CertificateAuthority_ManagedCertificateAuthority struct {
	// The PEM encoded CA certificate chains for redis managed
	//  server authentication
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.CertificateAuthority.ManagedCertificateAuthority.ca_certs
	CaCerts []CertificateAuthority_ManagedCertificateAuthority_CertChain `json:"caCerts,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1beta1.CertificateAuthority.ManagedCertificateAuthority.CertChain
type CertificateAuthority_ManagedCertificateAuthority_CertChain struct {
	// The certificates that form the CA chain, from leaf to root order.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.CertificateAuthority.ManagedCertificateAuthority.CertChain.certificates
	Certificates []string `json:"certificates,omitempty"`
}
