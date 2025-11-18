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

// +generated:types
// krm.group: certificatemanager.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.certificatemanager.v1
// resource: CertificateManagerDNSAuthorization:DnsAuthorization

package v1beta1

// +kcc:proto=google.cloud.certificatemanager.v1.DnsAuthorization.DnsResourceRecord
type DNSAuthorization_DNSResourceRecord struct {
}

// +kcc:observedstate:proto=google.cloud.certificatemanager.v1.DnsAuthorization.DnsResourceRecord
type DNSAuthorization_DNSResourceRecordObservedState struct {
	// Output only. Fully qualified name of the DNS Resource Record.
	//  e.g. `_acme-challenge.example.com`
	// +kcc:proto:field=google.cloud.certificatemanager.v1.DnsAuthorization.DnsResourceRecord.name
	Name *string `json:"name,omitempty"`

	// Output only. Type of the DNS Resource Record.
	//  Currently always set to "CNAME".
	// +kcc:proto:field=google.cloud.certificatemanager.v1.DnsAuthorization.DnsResourceRecord.type
	Type *string `json:"type,omitempty"`

	// Output only. Data of the DNS Resource Record.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.DnsAuthorization.DnsResourceRecord.data
	Data *string `json:"data,omitempty"`
}
