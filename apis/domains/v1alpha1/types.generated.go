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


// +kcc:proto=google.cloud.domains.v1beta1.ContactSettings
type ContactSettings struct {
	// Required. Privacy setting for the contacts associated with the `Registration`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ContactSettings.privacy
	Privacy *string `json:"privacy,omitempty"`

	// Required. The registrant contact for the `Registration`.
	//
	//  *Caution: Anyone with access to this email address, phone number,
	//  and/or postal address can take control of the domain.*
	//
	//  *Warning: For new `Registration`s, the registrant receives an email
	//  confirmation that they must complete within 15 days to avoid domain
	//  suspension.*
	// +kcc:proto:field=google.cloud.domains.v1beta1.ContactSettings.registrant_contact
	RegistrantContact *ContactSettings_Contact `json:"registrantContact,omitempty"`

	// Required. The administrative contact for the `Registration`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ContactSettings.admin_contact
	AdminContact *ContactSettings_Contact `json:"adminContact,omitempty"`

	// Required. The technical contact for the `Registration`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ContactSettings.technical_contact
	TechnicalContact *ContactSettings_Contact `json:"technicalContact,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.ContactSettings.Contact
type ContactSettings_Contact struct {
	// Required. Postal address of the contact.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ContactSettings.Contact.postal_address
	PostalAddress *PostalAddress `json:"postalAddress,omitempty"`

	// Required. Email address of the contact.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ContactSettings.Contact.email
	Email *string `json:"email,omitempty"`

	// Required. Phone number of the contact in international format. For example,
	//  `"+1-800-555-0123"`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ContactSettings.Contact.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// Fax number of the contact in international format. For example,
	//  `"+1-800-555-0123"`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ContactSettings.Contact.fax_number
	FaxNumber *string `json:"faxNumber,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.DnsSettings
type DnsSettings struct {
	// An arbitrary DNS provider identified by its name servers.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.custom_dns
	CustomDns *DnsSettings_CustomDns `json:"customDns,omitempty"`

	// The free DNS zone provided by
	//  [Google Domains](https://domains.google/).
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.google_domains_dns
	GoogleDomainsDns *DnsSettings_GoogleDomainsDns `json:"googleDomainsDns,omitempty"`

	// The list of glue records for this `Registration`. Commonly empty.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.glue_records
	GlueRecords []DnsSettings_GlueRecord `json:"glueRecords,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.DnsSettings.CustomDns
type DnsSettings_CustomDns struct {
	// Required. A list of name servers that store the DNS zone for this domain. Each name
	//  server is a domain name, with Unicode domain names expressed in
	//  Punycode format.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.CustomDns.name_servers
	NameServers []string `json:"nameServers,omitempty"`

	// The list of DS records for this domain, which are used to enable DNSSEC.
	//  The domain's DNS provider can provide the values to set here. If this
	//  field is empty, DNSSEC is disabled.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.CustomDns.ds_records
	DsRecords []DnsSettings_DsRecord `json:"dsRecords,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.DnsSettings.DsRecord
type DnsSettings_DsRecord struct {
	// The key tag of the record. Must be set in range 0 -- 65535.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.DsRecord.key_tag
	KeyTag *int32 `json:"keyTag,omitempty"`

	// The algorithm used to generate the referenced DNSKEY.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.DsRecord.algorithm
	Algorithm *string `json:"algorithm,omitempty"`

	// The hash function used to generate the digest of the referenced DNSKEY.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.DsRecord.digest_type
	DigestType *string `json:"digestType,omitempty"`

	// The digest generated from the referenced DNSKEY.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.DsRecord.digest
	Digest *string `json:"digest,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.DnsSettings.GlueRecord
type DnsSettings_GlueRecord struct {
	// Required. Domain name of the host in Punycode format.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.GlueRecord.host_name
	HostName *string `json:"hostName,omitempty"`

	// List of IPv4 addresses corresponding to this host in the standard decimal
	//  format (e.g. `198.51.100.1`). At least one of `ipv4_address` and
	//  `ipv6_address` must be set.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.GlueRecord.ipv4_addresses
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty"`

	// List of IPv6 addresses corresponding to this host in the standard
	//  hexadecimal format (e.g. `2001:db8::`). At least one of
	//  `ipv4_address` and `ipv6_address` must be set.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.GlueRecord.ipv6_addresses
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.DnsSettings.GoogleDomainsDns
type DnsSettings_GoogleDomainsDns struct {

	// Required. The state of DS records for this domain. Used to enable or disable
	//  automatic DNSSEC.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.GoogleDomainsDns.ds_state
	DsState *string `json:"dsState,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.ManagementSettings
type ManagementSettings struct {

	// Controls whether the domain can be transferred to another registrar.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ManagementSettings.transfer_lock_state
	TransferLockState *string `json:"transferLockState,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.Registration
type Registration struct {

	// Required. Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.domain_name
	DomainName *string `json:"domainName,omitempty"`

	// Set of labels associated with the `Registration`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Settings for management of the `Registration`, including renewal, billing,
	//  and transfer. You cannot update these with the `UpdateRegistration`
	//  method. To update these settings, use the `ConfigureManagementSettings`
	//  method.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.management_settings
	ManagementSettings *ManagementSettings `json:"managementSettings,omitempty"`

	// Settings controlling the DNS configuration of the `Registration`. You
	//  cannot update these with the `UpdateRegistration` method. To update these
	//  settings, use the `ConfigureDnsSettings` method.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.dns_settings
	DnsSettings *DnsSettings `json:"dnsSettings,omitempty"`

	// Required. Settings for contact information linked to the `Registration`. You cannot
	//  update these with the `UpdateRegistration` method. To update these
	//  settings, use the `ConfigureContactSettings` method.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.contact_settings
	ContactSettings *ContactSettings `json:"contactSettings,omitempty"`
}

// +kcc:proto=google.type.PostalAddress
type PostalAddress struct {
	// The schema revision of the `PostalAddress`. This must be set to 0, which is
	//  the latest revision.
	//
	//  All new revisions **must** be backward compatible with old revisions.
	// +kcc:proto:field=google.type.PostalAddress.revision
	Revision *int32 `json:"revision,omitempty"`

	// Required. CLDR region code of the country/region of the address. This
	//  is never inferred and it is up to the user to ensure the value is
	//  correct. See http://cldr.unicode.org/ and
	//  http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html
	//  for details. Example: "CH" for Switzerland.
	// +kcc:proto:field=google.type.PostalAddress.region_code
	RegionCode *string `json:"regionCode,omitempty"`

	// Optional. BCP-47 language code of the contents of this address (if
	//  known). This is often the UI language of the input form or is expected
	//  to match one of the languages used in the address' country/region, or their
	//  transliterated equivalents.
	//  This can affect formatting in certain countries, but is not critical
	//  to the correctness of the data and will never affect any validation or
	//  other non-formatting related operations.
	//
	//  If this value is not known, it should be omitted (rather than specifying a
	//  possibly incorrect default).
	//
	//  Examples: "zh-Hant", "ja", "ja-Latn", "en".
	// +kcc:proto:field=google.type.PostalAddress.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Optional. Postal code of the address. Not all countries use or require
	//  postal codes to be present, but where they are used, they may trigger
	//  additional validation with other parts of the address (e.g. state/zip
	//  validation in the U.S.A.).
	// +kcc:proto:field=google.type.PostalAddress.postal_code
	PostalCode *string `json:"postalCode,omitempty"`

	// Optional. Additional, country-specific, sorting code. This is not used
	//  in most regions. Where it is used, the value is either a string like
	//  "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number
	//  alone, representing the "sector code" (Jamaica), "delivery area indicator"
	//  (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
	// +kcc:proto:field=google.type.PostalAddress.sorting_code
	SortingCode *string `json:"sortingCode,omitempty"`

	// Optional. Highest administrative subdivision which is used for postal
	//  addresses of a country or region.
	//  For example, this can be a state, a province, an oblast, or a prefecture.
	//  Specifically, for Spain this is the province and not the autonomous
	//  community (e.g. "Barcelona" and not "Catalonia").
	//  Many countries don't use an administrative area in postal addresses. E.g.
	//  in Switzerland this should be left unpopulated.
	// +kcc:proto:field=google.type.PostalAddress.administrative_area
	AdministrativeArea *string `json:"administrativeArea,omitempty"`

	// Optional. Generally refers to the city/town portion of the address.
	//  Examples: US city, IT comune, UK post town.
	//  In regions of the world where localities are not well defined or do not fit
	//  into this structure well, leave locality empty and use address_lines.
	// +kcc:proto:field=google.type.PostalAddress.locality
	Locality *string `json:"locality,omitempty"`

	// Optional. Sublocality of the address.
	//  For example, this can be neighborhoods, boroughs, districts.
	// +kcc:proto:field=google.type.PostalAddress.sublocality
	Sublocality *string `json:"sublocality,omitempty"`

	// Unstructured address lines describing the lower levels of an address.
	//
	//  Because values in address_lines do not have type information and may
	//  sometimes contain multiple values in a single field (e.g.
	//  "Austin, TX"), it is important that the line order is clear. The order of
	//  address lines should be "envelope order" for the country/region of the
	//  address. In places where this can vary (e.g. Japan), address_language is
	//  used to make it explicit (e.g. "ja" for large-to-small ordering and
	//  "ja-Latn" or "en" for small-to-large). This way, the most specific line of
	//  an address can be selected based on the language.
	//
	//  The minimum permitted structural representation of an address consists
	//  of a region_code with all remaining information placed in the
	//  address_lines. It would be possible to format such an address very
	//  approximately without geocoding, but no semantic reasoning could be
	//  made about any of the address components until it was at least
	//  partially resolved.
	//
	//  Creating an address only containing a region_code and address_lines, and
	//  then geocoding is the recommended way to handle completely unstructured
	//  addresses (as opposed to guessing which parts of the address should be
	//  localities or administrative areas).
	// +kcc:proto:field=google.type.PostalAddress.address_lines
	AddressLines []string `json:"addressLines,omitempty"`

	// Optional. The recipient at the address.
	//  This field may, under certain circumstances, contain multiline information.
	//  For example, it might contain "care of" information.
	// +kcc:proto:field=google.type.PostalAddress.recipients
	Recipients []string `json:"recipients,omitempty"`

	// Optional. The name of the organization at the address.
	// +kcc:proto:field=google.type.PostalAddress.organization
	Organization *string `json:"organization,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.DnsSettings
type DnsSettingsObservedState struct {
	// The free DNS zone provided by
	//  [Google Domains](https://domains.google/).
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.google_domains_dns
	GoogleDomainsDns *DnsSettings_GoogleDomainsDnsObservedState `json:"googleDomainsDns,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.DnsSettings.GoogleDomainsDns
type DnsSettings_GoogleDomainsDnsObservedState struct {
	// Output only. A list of name servers that store the DNS zone for this domain. Each name
	//  server is a domain name, with Unicode domain names expressed in
	//  Punycode format. This field is automatically populated with the name
	//  servers assigned to the Google Domains DNS zone.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.GoogleDomainsDns.name_servers
	NameServers []string `json:"nameServers,omitempty"`

	// Output only. The list of DS records published for this domain. The list is
	//  automatically populated when `ds_state` is `DS_RECORDS_PUBLISHED`,
	//  otherwise it remains empty.
	// +kcc:proto:field=google.cloud.domains.v1beta1.DnsSettings.GoogleDomainsDns.ds_records
	DsRecords []DnsSettings_DsRecord `json:"dsRecords,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.ManagementSettings
type ManagementSettingsObservedState struct {
	// Output only. The renewal method for this `Registration`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.ManagementSettings.renewal_method
	RenewalMethod *string `json:"renewalMethod,omitempty"`
}

// +kcc:proto=google.cloud.domains.v1beta1.Registration
type RegistrationObservedState struct {
	// Output only. Name of the `Registration` resource, in the format
	//  `projects/*/locations/*/registrations/<domain_name>`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.name
	Name *string `json:"name,omitempty"`

	// Output only. The creation timestamp of the `Registration` resource.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The expiration timestamp of the `Registration`.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. The state of the `Registration`
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.state
	State *string `json:"state,omitempty"`

	// Output only. The set of issues with the `Registration` that require attention.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.issues
	Issues []string `json:"issues,omitempty"`

	// Settings for management of the `Registration`, including renewal, billing,
	//  and transfer. You cannot update these with the `UpdateRegistration`
	//  method. To update these settings, use the `ConfigureManagementSettings`
	//  method.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.management_settings
	ManagementSettings *ManagementSettingsObservedState `json:"managementSettings,omitempty"`

	// Settings controlling the DNS configuration of the `Registration`. You
	//  cannot update these with the `UpdateRegistration` method. To update these
	//  settings, use the `ConfigureDnsSettings` method.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.dns_settings
	DnsSettings *DnsSettingsObservedState `json:"dnsSettings,omitempty"`

	// Output only. Pending contact settings for the `Registration`. Updates to the
	//  `contact_settings` field that change its `registrant_contact` or `privacy`
	//  fields require email confirmation by the `registrant_contact`
	//  before taking effect. This field is set only if there are pending updates
	//  to the `contact_settings` that have not been confirmed. To confirm the
	//  changes, the `registrant_contact` must follow the instructions in the
	//  email they receive.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.pending_contact_settings
	PendingContactSettings *ContactSettings `json:"pendingContactSettings,omitempty"`

	// Output only. Set of options for the `contact_settings.privacy` field that this
	//  `Registration` supports.
	// +kcc:proto:field=google.cloud.domains.v1beta1.Registration.supported_privacy
	SupportedPrivacy []string `json:"supportedPrivacy,omitempty"`
}
