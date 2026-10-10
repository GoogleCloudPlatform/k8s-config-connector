// Copyright 2025 Google LLC
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

// API sources for JobsCompany, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/talent/v4/company.proto
// +kcc:source:service-docs=https://cloud.google.com/talent-solution/job-search/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/talent-solution/job-search/docs/reference/rest/v4/projects.tenants.companies

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var JobsCompanyGVK = GroupVersion.WithKind("JobsCompany")

// JobsCompanySpec defines the desired state of JobsCompany
// +kcc:spec:proto=google.cloud.talent.v4.Company
// +kcc:required-from-proto
type JobsCompanySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The JobsCompany name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the company, for example, "Google LLC".
	// +kcc:proto:field=google.cloud.talent.v4.Company.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Client side company identifier, used to uniquely identify the
	//  company.
	//
	//  The maximum number of allowed characters is 255.
	// +kcc:proto:field=google.cloud.talent.v4.Company.external_id
	// +required
	ExternalID *string `json:"externalID,omitempty"`

	// The employer's company size.
	// +kcc:proto:field=google.cloud.talent.v4.Company.size
	Size *string `json:"size,omitempty"`

	// The street address of the company's main headquarters, which may be
	//  different from the job location. The service attempts
	//  to geolocate the provided address, and populates a more specific
	//  location wherever possible in
	//  [DerivedInfo.headquarters_location][google.cloud.talent.v4.Company.DerivedInfo.headquarters_location].
	// +kcc:proto:field=google.cloud.talent.v4.Company.headquarters_address
	HeadquartersAddress *string `json:"headquartersAddress,omitempty"`

	// Set to true if it is the hiring agency that post jobs for other
	//  employers.
	//
	//  Defaults to false if not provided.
	// +kcc:proto:field=google.cloud.talent.v4.Company.hiring_agency
	HiringAgency *bool `json:"hiringAgency,omitempty"`

	// Equal Employment Opportunity legal disclaimer text to be
	//  associated with all jobs, and typically to be displayed in all
	//  roles.
	//
	//  The maximum number of allowed characters is 500.
	// +kcc:proto:field=google.cloud.talent.v4.Company.eeo_text
	EeoText *string `json:"eeoText,omitempty"`

	// The URI representing the company's primary web site or home page,
	//  for example, "https://www.google.com".
	//
	//  The maximum number of allowed characters is 255.
	// +kcc:proto:field=google.cloud.talent.v4.Company.website_uri
	WebsiteURI *string `json:"websiteURI,omitempty"`

	// The URI to employer's career site or careers page on the employer's web
	//  site, for example, "https://careers.google.com".
	// +kcc:proto:field=google.cloud.talent.v4.Company.career_site_uri
	CareerSiteURI *string `json:"careerSiteURI,omitempty"`

	// A URI that hosts the employer's company logo.
	// +kcc:proto:field=google.cloud.talent.v4.Company.image_uri
	ImageURI *string `json:"imageURI,omitempty"`
}

// JobsCompanyStatus defines the config connector machine state of JobsCompany
type JobsCompanyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the JobsCompany resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *JobsCompanyObservedState `json:"observedState,omitempty"`
}

// JobsCompanyObservedState is the state of the JobsCompany resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.talent.v4.Company
type JobsCompanyObservedState struct {
	// Output only. Derived details about the company.
	// +kcc:proto:field=google.cloud.talent.v4.Company.derived_info
	DerivedInfo *Company_DerivedInfo `json:"derivedInfo,omitempty"`

	// Output only. Indicates whether a company is flagged to be suspended from
	//  public availability by the service when job content appears suspicious,
	//  abusive, or spammy.
	// +kcc:proto:field=google.cloud.talent.v4.Company.suspended
	Suspended *bool `json:"suspended,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpjobscompany;gcpjobscompanys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// JobsCompany is the Schema for the JobsCompany API
// +k8s:openapi-gen=true
type JobsCompany struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   JobsCompanySpec   `json:"spec,omitempty"`
	Status JobsCompanyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// JobsCompanyList contains a list of JobsCompany
type JobsCompanyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobsCompany `json:"items"`
}

func init() {
	SchemeBuilder.Register(&JobsCompany{}, &JobsCompanyList{})
}
