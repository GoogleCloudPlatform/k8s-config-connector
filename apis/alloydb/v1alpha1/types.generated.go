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


// +kcc:proto=google.cloud.alloydb.v1beta.ConnectionInfo
type ConnectionInfo struct {
	// The name of the ConnectionInfo singleton resource, e.g.:
	//  projects/{project}/locations/{location}/clusters/*/instances/*/connectionInfo
	//  This field currently has no semantic meaning.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ConnectionInfo.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.ConnectionInfo
type ConnectionInfoObservedState struct {
	// Output only. The private network IP address for the Instance. This is the
	//  default IP for the instance and is always created (even if enable_public_ip
	//  is set). This is the connection endpoint for an end-user application.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ConnectionInfo.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Output only. The public IP addresses for the Instance. This is available
	//  ONLY when enable_public_ip is set. This is the connection endpoint for an
	//  end-user application.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ConnectionInfo.public_ip_address
	PublicIPAddress *string `json:"publicIPAddress,omitempty"`

	// Output only. The pem-encoded chain that may be used to verify the X.509
	//  certificate. Expected to be in issuer-to-root order according to RFC 5246.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ConnectionInfo.pem_certificate_chain
	PemCertificateChain []string `json:"pemCertificateChain,omitempty"`

	// Output only. The unique ID of the Instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ConnectionInfo.instance_uid
	InstanceUid *string `json:"instanceUid,omitempty"`

	// Output only. The DNS name to use with PSC for the Instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ConnectionInfo.psc_dns_name
	PscDnsName *string `json:"pscDnsName,omitempty"`
}
