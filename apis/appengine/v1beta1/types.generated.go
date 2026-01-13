// Copyright 2026 Google LLC
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
// krm.group: appengine.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.appengine.v1
// resource: AppEngineDomainMapping:DomainMapping

package v1beta1

// +kcc:proto=google.appengine.v1.ResourceRecord
type ResourceRecord struct {
	// Relative name of the object affected by this record. Only applicable for
	//  `CNAME` records. Example: 'www'.
	// +kcc:proto:field=google.appengine.v1.ResourceRecord.name
	Name *string `json:"name,omitempty"`

	// Data for this record. Values vary by record type, as defined in RFC 1035
	//  (section 5) and RFC 1034 (section 3.6.1).
	// +kcc:proto:field=google.appengine.v1.ResourceRecord.rrdata
	Rrdata *string `json:"rrdata,omitempty"`

	// Resource record type. Example: `AAAA`.
	// +kcc:proto:field=google.appengine.v1.ResourceRecord.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.appengine.v1.SslSettings
type SSLSettings struct {
	// ID of the `AuthorizedCertificate` resource configuring SSL for the
	//  application. Clearing this field will remove SSL support.
	//
	//  By default, a managed certificate is automatically created for every
	//  domain mapping. To omit SSL support or to configure SSL manually, specify
	//  `SslManagementType.MANUAL` on a `CREATE` or `UPDATE` request. You must
	//  be authorized to administer the `AuthorizedCertificate` resource to
	//  manually map it to a `DomainMapping` resource.
	//  Example: `12345`.
	// +kcc:proto:field=google.appengine.v1.SslSettings.certificate_id
	CertificateID *string `json:"certificateID,omitempty"`

	// SSL management type for this domain. If `AUTOMATIC`, a managed certificate
	//  is automatically provisioned. If `MANUAL`, `certificate_id` must be
	//  manually specified in order to configure SSL for this domain.
	// +kcc:proto:field=google.appengine.v1.SslSettings.ssl_management_type
	SSLManagementType *string `json:"sslManagementType,omitempty"`

	// ID of the managed `AuthorizedCertificate` resource currently being
	//  provisioned, if applicable. Until the new managed certificate has been
	//  successfully provisioned, the previous SSL state will be preserved. Once
	//  the provisioning process completes, the `certificate_id` field will reflect
	//  the new managed certificate and this field will be left empty. To remove
	//  SSL support while there is still a pending managed certificate, clear the
	//  `certificate_id` field with an `UpdateDomainMappingRequest`.
	//
	//  @OutputOnly
	// +kcc:proto:field=google.appengine.v1.SslSettings.pending_managed_certificate_id
	PendingManagedCertificateID *string `json:"pendingManagedCertificateID,omitempty"`
}
