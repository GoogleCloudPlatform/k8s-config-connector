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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudTalentSolutionCompanyGVK = GroupVersion.WithKind("CloudTalentSolutionCompany")

// CloudTalentSolutionCompanySpec defines the desired state of CloudTalentSolutionCompany
// +kcc:spec:proto=google.cloud.talent.v4.Company
type CloudTalentSolutionCompanySpec struct {
	// The Tenant that this resource belongs to.
	// +kubebuilder:validation:Required
	// +required
	TenantRef *TenantRef `json:"tenantRef"`

	// The CloudTalentSolutionCompany name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the company, for example, "Google LLC".
	// +kubebuilder:validation:Required
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Client side company identifier, used to uniquely identify the company.
	// The maximum number of allowed characters is 255.
	// +kubebuilder:validation:Required
	// +required
	ExternalID *string `json:"externalID,omitempty"`

	// The employer's company size.
	// +optional
	// +kubebuilder:validation:Enum=COMPANY_SIZE_UNSPECIFIED;MINI;SMALL;SMEDIUM;MEDIUM;BIG;BIGGER;GIANT
	Size *string `json:"size,omitempty"`

	// The street address of the company's main headquarters, which may be
	// different from the job location.
	// +optional
	HeadquartersAddress *string `json:"headquartersAddress,omitempty"`

	// Set to true if it is the hiring agency that post jobs for other employers.
	// Defaults to false if not provided.
	// +optional
	HiringAgency *bool `json:"hiringAgency,omitempty"`

	// Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles.
	// The maximum number of allowed characters is 500.
	// +optional
	EeoText *string `json:"eeoText,omitempty"`

	// The URI representing the company's primary web site or home page, for example, "https://www.google.com".
	// The maximum number of allowed characters is 255.
	// +optional
	WebsiteURI *string `json:"websiteURI,omitempty"`

	// The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	// +optional
	CareerSiteURI *string `json:"careerSiteURI,omitempty"`

	// A URI that hosts the employer's company logo.
	// +optional
	ImageURI *string `json:"imageURI,omitempty"`

	// This field is deprecated. Please set the searchability of the custom attribute in the Job.custom_attributes going forward.
	// A list of keys of filterable Job.custom_attributes, whose corresponding string_values are used in keyword searches.
	// +optional
	KeywordSearchableJobCustomAttributes []string `json:"keywordSearchableJobCustomAttributes,omitempty"`
}

// CompanyObservedState is the state of the CloudTalentSolutionCompany resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.talent.v4.Company
type CompanyObservedState struct {
	// Output only. Derived details about the company.
	// +optional
	DerivedInfo *Company_DerivedInfo `json:"derivedInfo,omitempty"`

	// Output only. Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
	// +optional
	Suspended *bool `json:"suspended,omitempty"`
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
	ObservedState *CompanyObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudtalentsolutioncompany;gcpcloudtalentsolutioncompanys
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
