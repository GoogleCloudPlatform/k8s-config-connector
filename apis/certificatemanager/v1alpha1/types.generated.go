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


// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMapEntry
type CertificateMapEntry struct {
	// A user-defined name of the Certificate Map Entry. Certificate Map Entry
	//  names must be unique globally and match pattern
	//  `projects/*/locations/*/certificateMaps/*/certificateMapEntries/*`.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.name
	Name *string `json:"name,omitempty"`

	// One or more paragraphs of text description of a certificate map entry.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.description
	Description *string `json:"description,omitempty"`

	// Set of labels associated with a Certificate Map Entry.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.labels
	Labels map[string]string `json:"labels,omitempty"`

	// A Hostname (FQDN, e.g. `example.com`) or a wildcard hostname expression
	//  (`*.example.com`) for a set of hostnames with common suffix. Used as
	//  Server Name Indication (SNI) for selecting a proper certificate.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.hostname
	Hostname *string `json:"hostname,omitempty"`

	// A predefined matcher for particular cases, other than SNI selection.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.matcher
	Matcher *string `json:"matcher,omitempty"`

	// A set of Certificates defines for the given `hostname`. There can be
	//  defined up to four certificates in each Certificate Map Entry. Each
	//  certificate must match pattern `projects/*/locations/*/certificates/*`.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.certificates
	Certificates []string `json:"certificates,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMapEntry
type CertificateMapEntryObservedState struct {
	// Output only. The creation timestamp of a Certificate Map Entry.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp of a Certificate Map Entry.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A serving state of this Certificate Map Entry.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMapEntry.state
	State *string `json:"state,omitempty"`
}
