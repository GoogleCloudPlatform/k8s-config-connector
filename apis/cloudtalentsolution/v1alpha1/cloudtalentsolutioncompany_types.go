// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudTalentSolutionCompanyGVK = GroupVersion.WithKind("CloudTalentSolutionCompany")

// CloudTalentSolutionCompanySpec defines the desired state of CloudTalentSolutionCompany
// +kcc:spec:proto=google.cloud.talent.v4.Company
type CloudTalentSolutionCompanySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The tenant of this resource.
	// +required
	Tenant *string `json:"tenant,omitempty"`

	// The CloudTalentSolutionCompany name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the company, for example, "Google LLC".
	// +required
	// +kcc:proto:field=google.cloud.talent.v4.Company.display_name
	DisplayName string `json:"displayName"`

	// Required. Client side company identifier, used to uniquely identify the company.
	// The maximum number of allowed characters is 255.
	// +required
	// +kcc:proto:field=google.cloud.talent.v4.Company.external_id
	ExternalID string `json:"externalID"`

	// The employer's company size.
	// +optional
	// +kubebuilder:validation:Enum=COMPANY_SIZE_UNSPECIFIED;MINI;SMALL;SMEDIUM;MEDIUM;BIG;BIGGER;GIANT
	// +kcc:proto:field=google.cloud.talent.v4.Company.size
	Size *string `json:"size,omitempty"`

	// The street address of the company's main headquarters, which may be
	// different from the job location.
	// +optional
	// +kcc:proto:field=google.cloud.talent.v4.Company.headquarters_address
	HeadquartersAddress *string `json:"headquartersAddress,omitempty"`

	// Set to true if it is the hiring agency that post jobs for other employers.
	// +optional
	// +kcc:proto:field=google.cloud.talent.v4.Company.hiring_agency
	HiringAgency *bool `json:"hiringAgency,omitempty"`

	// Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles.
	// The maximum number of allowed characters is 500.
	// +optional
	// +kcc:proto:field=google.cloud.talent.v4.Company.eeo_text
	EeoText *string `json:"eeoText,omitempty"`

	// The URI representing the company's primary web site or home page, for example, "https://www.google.com".
	// The maximum number of allowed characters is 255.
	// +optional
	// +kcc:proto:field=google.cloud.talent.v4.Company.website_uri
	WebsiteURI *string `json:"websiteURI,omitempty"`

	// The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	// +optional
	// +kcc:proto:field=google.cloud.talent.v4.Company.career_site_uri
	CareerSiteURI *string `json:"careerSiteURI,omitempty"`

	// A URI that hosts the employer's company logo.
	// +optional
	// +kcc:proto:field=google.cloud.talent.v4.Company.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward.
	// +optional
	// +kcc:proto:field=google.cloud.talent.v4.Company.keyword_searchable_job_custom_attributes
	KeywordSearchableJobCustomAttributes []string `json:"keywordSearchableJobCustomAttributes,omitempty"`
}

// CloudTalentSolutionCompanyStatus defines the config connector machine state of CloudTalentSolutionCompany
type CloudTalentSolutionCompanyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudTalentSolutionCompany resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudTalentSolutionCompanyObservedState `json:"observedState,omitempty"`
}

// CloudTalentSolutionCompanyObservedState is the state of the CloudTalentSolutionCompany resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.talent.v4.Company
type CloudTalentSolutionCompanyObservedState struct {
	// Output only. Derived details about the company.
	// +kcc:proto:field=google.cloud.talent.v4.Company.derived_info
	DerivedInfo *Company_DerivedInfo `json:"derivedInfo,omitempty"`

	// Output only. Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
	// +kcc:proto:field=google.cloud.talent.v4.Company.suspended
	Suspended *bool `json:"suspended,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4.Company.DerivedInfo
type Company_DerivedInfo struct {
	// A structured headquarters location of the company, resolved from Company.headquarters_address if provided.
	// +kcc:proto:field=google.cloud.talent.v4.Company.DerivedInfo.headquarters_location
	HeadquartersLocation *Location `json:"headquartersLocation,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4.Location
type Location struct {
	// The type of a location, which corresponds to the address lines field of google.type.PostalAddress.
	// +kubebuilder:validation:Enum=LOCATION_TYPE_UNSPECIFIED;COUNTRY;ADMINISTRATIVE_AREA;SUB_ADMINISTRATIVE_AREA;LOCALITY;POSTAL_CODE;SUB_LOCALITY;SUB_LOCALITY_1;SUB_LOCALITY_2;NEIGHBORHOOD;STREET_ADDRESS
	// +kcc:proto:field=google.cloud.talent.v4.Location.location_type
	LocationType *string `json:"locationType,omitempty"`

	// Postal address of the location that includes human readable information, such as postal delivery and payments addresses.
	// +kcc:proto:field=google.cloud.talent.v4.Location.postal_address
	PostalAddress *PostalAddress `json:"postalAddress,omitempty"`

	// An object representing a latitude/longitude pair.
	// +kcc:proto:field=google.cloud.talent.v4.Location.lat_lng
	LatLng *LatLng `json:"latLng,omitempty"`

	// Radius in miles of the job location.
	// +kcc:proto:field=google.cloud.talent.v4.Location.radius_miles
	RadiusMiles *float64 `json:"radiusMiles,omitempty"`
}

// +kcc:proto=google.type.LatLng
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	// +kcc:proto:field=google.type.LatLng.latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	// +kcc:proto:field=google.type.LatLng.longitude
	Longitude *float64 `json:"longitude,omitempty"`
}

// +kcc:proto=google.type.PostalAddress
type PostalAddress struct {
	// The schema revision of the PostalAddress. This must be set to 0, which is the latest revision.
	// +kcc:proto:field=google.type.PostalAddress.revision
	Revision *int32 `json:"revision,omitempty"`

	// Required. CLDR region code of the country/region of the address.
	// +kcc:proto:field=google.type.PostalAddress.region_code
	RegionCode *string `json:"regionCode,omitempty"`

	// Optional. BCP-47 language code of the contents of this address (if known).
	// +kcc:proto:field=google.type.PostalAddress.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Optional. Postal code of the address.
	// +kcc:proto:field=google.type.PostalAddress.postal_code
	PostalCode *string `json:"postalCode,omitempty"`

	// Optional. Additional, country-specific, sorting code.
	// +kcc:proto:field=google.type.PostalAddress.sorting_code
	SortingCode *string `json:"sortingCode,omitempty"`

	// Optional. Highest administrative subdivision which is used for postal addresses of a country or region.
	// +kcc:proto:field=google.type.PostalAddress.administrative_area
	AdministrativeArea *string `json:"administrativeArea,omitempty"`

	// Optional. Generally refers to the city/town portion of the address.
	// +kcc:proto:field=google.type.PostalAddress.locality
	Locality *string `json:"locality,omitempty"`

	// Optional. Sublocality of the address.
	// +kcc:proto:field=google.type.PostalAddress.sublocality
	Sublocality *string `json:"sublocality,omitempty"`

	// Unstructured address lines describing the lower levels of an address.
	// +kcc:proto:field=google.type.PostalAddress.address_lines
	AddressLines []string `json:"addressLines,omitempty"`

	// Optional. The recipient at the address.
	// +kcc:proto:field=google.type.PostalAddress.recipients
	Recipients []string `json:"recipients,omitempty"`

	// Optional. The name of the organization at the address.
	// +kcc:proto:field=google.type.PostalAddress.organization
	Organization *string `json:"organization,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudtalentsolutioncompany;gcpcloudtalentsolutioncompanies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudTalentSolutionCompany is the Schema for the CloudTalentSolutionCompany API
// +k8s:openapi-gen=true
type CloudTalentSolutionCompany struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudTalentSolutionCompanySpec   `json:"spec,omitempty"`
	Status CloudTalentSolutionCompanyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudTalentSolutionCompanyList contains a list of CloudTalentSolutionCompany
type CloudTalentSolutionCompanyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudTalentSolutionCompany `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudTalentSolutionCompany{}, &CloudTalentSolutionCompanyList{})
}
