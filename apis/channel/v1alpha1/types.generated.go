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


// +kcc:proto=google.cloud.channel.v1.CloudIdentityInfo
type CloudIdentityInfo struct {
	// CustomerType indicates verification type needed for using services.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.customer_type
	CustomerType *string `json:"customerType,omitempty"`

	// The alternate email.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.alternate_email
	AlternateEmail *string `json:"alternateEmail,omitempty"`

	// Phone number associated with the Cloud Identity.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// Language code.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Edu information about the customer.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.edu_data
	EduData *EduData `json:"eduData,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.ContactInfo
type ContactInfo struct {
	// The customer account contact's first name. Optional for Team customers.
	// +kcc:proto:field=google.cloud.channel.v1.ContactInfo.first_name
	FirstName *string `json:"firstName,omitempty"`

	// The customer account contact's last name. Optional for Team customers.
	// +kcc:proto:field=google.cloud.channel.v1.ContactInfo.last_name
	LastName *string `json:"lastName,omitempty"`

	// The customer account's contact email. Required for entitlements that create
	//  admin.google.com accounts, and serves as the customer's username for those
	//  accounts. Use this email to invite Team customers.
	// +kcc:proto:field=google.cloud.channel.v1.ContactInfo.email
	Email *string `json:"email,omitempty"`

	// Optional. The customer account contact's job title.
	// +kcc:proto:field=google.cloud.channel.v1.ContactInfo.title
	Title *string `json:"title,omitempty"`

	// The customer account's contact phone number.
	// +kcc:proto:field=google.cloud.channel.v1.ContactInfo.phone
	Phone *string `json:"phone,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Customer
type Customer struct {

	// Required. Name of the organization that the customer entity represents.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.org_display_name
	OrgDisplayName *string `json:"orgDisplayName,omitempty"`

	// Required. The organization address for the customer. To enforce US laws and
	//  embargoes, we require a region, postal code, and address lines. You must
	//  provide valid addresses for every customer. To set the customer's
	//  language, use the Customer-level language code.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.org_postal_address
	OrgPostalAddress *PostalAddress `json:"orgPostalAddress,omitempty"`

	// Primary contact info.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.primary_contact_info
	PrimaryContactInfo *ContactInfo `json:"primaryContactInfo,omitempty"`

	// Secondary contact email. You need to provide an alternate email to create
	//  different domains if a primary contact email already exists. Users will
	//  receive a notification with credentials when you create an admin.google.com
	//  account. Secondary emails are also recovery email addresses. Alternate
	//  emails are optional when you create Team customers.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.alternate_email
	AlternateEmail *string `json:"alternateEmail,omitempty"`

	// Required. The customer's primary domain. Must match the primary contact
	//  email's domain.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.domain
	Domain *string `json:"domain,omitempty"`

	// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	//  information, see
	//  https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Cloud Identity ID of the customer's channel partner.
	//  Populated only if a channel partner exists for this customer.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.channel_partner_id
	ChannelPartnerID *string `json:"channelPartnerID,omitempty"`

	// Optional. External CRM ID for the customer.
	//  Populated only if a CRM ID exists for this customer.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.correlation_id
	CorrelationID *string `json:"correlationID,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.EduData
type EduData struct {
	// Designated institute type of customer.
	// +kcc:proto:field=google.cloud.channel.v1.EduData.institute_type
	InstituteType *string `json:"instituteType,omitempty"`

	// Size of the institute.
	// +kcc:proto:field=google.cloud.channel.v1.EduData.institute_size
	InstituteSize *string `json:"instituteSize,omitempty"`

	// Web address for the edu customer's institution.
	// +kcc:proto:field=google.cloud.channel.v1.EduData.website
	Website *string `json:"website,omitempty"`
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

// +kcc:proto=google.cloud.channel.v1.CloudIdentityInfo
type CloudIdentityInfoObservedState struct {
	// Output only. The primary domain name.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.primary_domain
	PrimaryDomain *string `json:"primaryDomain,omitempty"`

	// Output only. Whether the domain is verified.
	//  This field is not returned for a Customer's cloud_identity_info resource.
	//  Partners can use the domains.get() method of the Workspace SDK's
	//  Directory API, or listen to the PRIMARY_DOMAIN_VERIFIED Pub/Sub event in
	//  to track domain verification of their resolve Workspace customers.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.is_domain_verified
	IsDomainVerified *bool `json:"isDomainVerified,omitempty"`

	// Output only. URI of Customer's Admin console dashboard.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.admin_console_uri
	AdminConsoleURI *string `json:"adminConsoleURI,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.ContactInfo
type ContactInfoObservedState struct {
	// Output only. The customer account contact's display name, formatted as a
	//  combination of the customer's first and last name.
	// +kcc:proto:field=google.cloud.channel.v1.ContactInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Customer
type CustomerObservedState struct {
	// Output only. Resource name of the customer.
	//  Format: accounts/{account_id}/customers/{customer_id}
	// +kcc:proto:field=google.cloud.channel.v1.Customer.name
	Name *string `json:"name,omitempty"`

	// Primary contact info.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.primary_contact_info
	PrimaryContactInfo *ContactInfoObservedState `json:"primaryContactInfo,omitempty"`

	// Output only. Time when the customer was created.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the customer was updated.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The customer's Cloud Identity ID if the customer has a Cloud
	//  Identity resource.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.cloud_identity_id
	CloudIdentityID *string `json:"cloudIdentityID,omitempty"`

	// Output only. Cloud Identity information for the customer.
	//  Populated only if a Cloud Identity account exists for this customer.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.cloud_identity_info
	CloudIdentityInfo *CloudIdentityInfo `json:"cloudIdentityInfo,omitempty"`
}
