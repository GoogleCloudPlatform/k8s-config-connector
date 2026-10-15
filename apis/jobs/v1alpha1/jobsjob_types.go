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

// API sources for JobsJob, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/talent/v4/job.proto
// +kcc:source:service-docs=https://cloud.google.com/talent-solution/job-search/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/talent-solution/job-search/docs/reference/rest/v4/projects.tenants.jobs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var JobsJobGVK = GroupVersion.WithKind("JobsJob")

// JobsJobSpec defines the desired state of JobsJob
// +kcc:spec:proto=google.cloud.talent.v4.Job
// +kcc:required-from-proto
type JobsJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The JobsJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The resource name of the company listing the job.
	//
	//  The format is
	//  "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}". For
	//  example, "projects/foo/tenants/bar/companies/baz".
	// +kcc:guess=possible-reference target=JobsCompany
	// +kcc:proto:field=google.cloud.talent.v4.Job.company
	// +required
	Company *string `json:"company,omitempty"`

	// Required. The requisition ID, also referred to as the posting ID, is
	//  assigned by the client to identify a job. This field is intended to be used
	//  by clients for client identification and tracking of postings. A job isn't
	//  allowed to be created if there is another job with the same
	//  [company][google.cloud.talent.v4.Job.name],
	//  [language_code][google.cloud.talent.v4.Job.language_code] and
	//  [requisition_id][google.cloud.talent.v4.Job.requisition_id].
	//
	//  The maximum number of allowed characters is 255.
	// +kcc:proto:field=google.cloud.talent.v4.Job.requisition_id
	// +required
	RequisitionID *string `json:"requisitionID,omitempty"`

	// Required. The title of the job, such as "Software Engineer"
	//
	//  The maximum number of allowed characters is 500.
	// +kcc:proto:field=google.cloud.talent.v4.Job.title
	// +required
	Title *string `json:"title,omitempty"`

	// Required. The description of the job, which typically includes a
	//  multi-paragraph description of the company and related information.
	//  Separate fields are provided on the job object for
	//  [responsibilities][google.cloud.talent.v4.Job.responsibilities],
	//  [qualifications][google.cloud.talent.v4.Job.qualifications], and other job
	//  characteristics. Use of these separate job fields is recommended.
	//
	//  This field accepts and sanitizes HTML input, and also accepts
	//  bold, italic, ordered list, and unordered list markup tags.
	//
	//  The maximum number of allowed characters is 100,000.
	// +kcc:proto:field=google.cloud.talent.v4.Job.description
	// +required
	Description *string `json:"description,omitempty"`

	// Strongly recommended for the best service experience.
	//
	//  Location(s) where the employer is looking to hire for this job posting.
	//
	//  Specifying the full street address(es) of the hiring location enables
	//  better API results, especially job searches by commute time.
	//
	//  At most 50 locations are allowed for best search performance. If a job has
	//  more locations, it is suggested to split it into multiple jobs with unique
	//  [requisition_id][google.cloud.talent.v4.Job.requisition_id]s (e.g. 'ReqA'
	//  becomes 'ReqA-1', 'ReqA-2', and so on.) as multiple jobs with the same
	//  [company][google.cloud.talent.v4.Job.company],
	//  [language_code][google.cloud.talent.v4.Job.language_code] and
	//  [requisition_id][google.cloud.talent.v4.Job.requisition_id] are not
	//  allowed. If the original
	//  [requisition_id][google.cloud.talent.v4.Job.requisition_id] must be
	//  preserved, a custom field should be used for storage. It is also suggested
	//  to group the locations that close to each other in the same job for better
	//  search experience.
	//
	//  Jobs with multiple addresses must have their addresses with the same
	//  [LocationType][google.cloud.talent.v4.Location.LocationType] to allow
	//  location filtering to work properly. (For example, a Job with addresses
	//  "1600 Amphitheatre Parkway, Mountain View, CA, USA" and "London, UK" may
	//  not have location filters applied correctly at search time since the first
	//  is a
	//  [LocationType.STREET_ADDRESS][google.cloud.talent.v4.Location.LocationType.STREET_ADDRESS]
	//  and the second is a
	//  [LocationType.LOCALITY][google.cloud.talent.v4.Location.LocationType.LOCALITY].)
	//  If a job needs to have multiple addresses, it is suggested to split it into
	//  multiple jobs with same LocationTypes.
	//
	//  The maximum number of allowed characters is 500.
	// +kcc:proto:field=google.cloud.talent.v4.Job.addresses
	Addresses []string `json:"addresses,omitempty"`

	// Job application information.
	// +kcc:proto:field=google.cloud.talent.v4.Job.application_info
	ApplicationInfo *Job_ApplicationInfo `json:"applicationInfo,omitempty"`

	// The benefits included with the job.
	// +kcc:proto:field=google.cloud.talent.v4.Job.job_benefits
	JobBenefits []string `json:"jobBenefits,omitempty"`

	// Job compensation information (a.k.a. "pay rate") i.e., the compensation
	//  that will paid to the employee.
	// +kcc:proto:field=google.cloud.talent.v4.Job.compensation_info
	CompensationInfo *CompensationInfo `json:"compensationInfo,omitempty"`

	// A map of fields to hold both filterable and non-filterable custom job
	//  attributes that are not covered by the provided structured fields.
	//
	//  The keys of the map are strings up to 64 bytes and must match the
	//  pattern: `[a-zA-Z][a-zA-Z0-9_]*`. For example, key0LikeThis or
	//  KEY_1_LIKE_THIS.
	//
	//  At most 100 filterable and at most 100 unfilterable keys are supported.
	//  For filterable `string_values`, across all keys at most 200 values are
	//  allowed, with each string no more than 255 characters. For unfilterable
	//  `string_values`, the maximum total size of `string_values` across all keys
	//  is 50KB.
	// +kcc:proto:field=google.cloud.talent.v4.Job.custom_attributes
	CustomAttributes map[string]CustomAttribute `json:"customAttributes,omitempty"`

	// The desired education degrees for the job, such as Bachelors, Masters.
	// +kcc:proto:field=google.cloud.talent.v4.Job.degree_types
	DegreeTypes []string `json:"degreeTypes,omitempty"`

	// The department or functional area within the company with the open
	//  position.
	//
	//  The maximum number of allowed characters is 255.
	// +kcc:proto:field=google.cloud.talent.v4.Job.department
	Department *string `json:"department,omitempty"`

	// The employment type(s) of a job, for example,
	//  [full time][google.cloud.talent.v4.EmploymentType.FULL_TIME] or
	//  [part time][google.cloud.talent.v4.EmploymentType.PART_TIME].
	// +kcc:proto:field=google.cloud.talent.v4.Job.employment_types
	EmploymentTypes []string `json:"employmentTypes,omitempty"`

	// A description of bonus, commission, and other compensation
	//  incentives associated with the job not including salary or pay.
	//
	//  The maximum number of allowed characters is 10,000.
	// +kcc:proto:field=google.cloud.talent.v4.Job.incentives
	Incentives *string `json:"incentives,omitempty"`

	// The language of the posting. This field is distinct from
	//  any requirements for fluency that are associated with the job.
	//
	//  Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn".
	//  For more information, see
	//  [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){:
	//  class="external" target="_blank" }.
	//
	//  If this field is unspecified and
	//  [Job.description][google.cloud.talent.v4.Job.description] is present,
	//  detected language code based on
	//  [Job.description][google.cloud.talent.v4.Job.description] is assigned,
	//  otherwise defaults to 'en_US'.
	// +kcc:proto:field=google.cloud.talent.v4.Job.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The experience level associated with the job, such as "Entry Level".
	// +kcc:proto:field=google.cloud.talent.v4.Job.job_level
	JobLevel *string `json:"jobLevel,omitempty"`

	// A promotion value of the job, as determined by the client.
	//  The value determines the sort order of the jobs returned when searching for
	//  jobs using the featured jobs search call, with higher promotional values
	//  being returned first and ties being resolved by relevance sort. Only the
	//  jobs with a promotionValue >0 are returned in a FEATURED_JOB_SEARCH.
	//
	//  Default value is 0, and negative values are treated as 0.
	// +kcc:proto:field=google.cloud.talent.v4.Job.promotion_value
	PromotionValue *int32 `json:"promotionValue,omitempty"`

	// A description of the qualifications required to perform the
	//  job. The use of this field is recommended
	//  as an alternative to using the more general
	//  [description][google.cloud.talent.v4.Job.description] field.
	//
	//  This field accepts and sanitizes HTML input, and also accepts
	//  bold, italic, ordered list, and unordered list markup tags.
	//
	//  The maximum number of allowed characters is 10,000.
	// +kcc:proto:field=google.cloud.talent.v4.Job.qualifications
	Qualifications *string `json:"qualifications,omitempty"`

	// A description of job responsibilities. The use of this field is
	//  recommended as an alternative to using the more general
	//  [description][google.cloud.talent.v4.Job.description] field.
	//
	//  This field accepts and sanitizes HTML input, and also accepts
	//  bold, italic, ordered list, and unordered list markup tags.
	//
	//  The maximum number of allowed characters is 10,000.
	// +kcc:proto:field=google.cloud.talent.v4.Job.responsibilities
	Responsibilities *string `json:"responsibilities,omitempty"`

	// The job [PostingRegion][google.cloud.talent.v4.PostingRegion] (for example,
	//  state, country) throughout which the job is available. If this field is
	//  set, a [LocationFilter][google.cloud.talent.v4.LocationFilter] in a search
	//  query within the job region finds this job posting if an exact location
	//  match isn't specified. If this field is set to
	//  [PostingRegion.NATION][google.cloud.talent.v4.PostingRegion.NATION] or
	//  [PostingRegion.ADMINISTRATIVE_AREA][google.cloud.talent.v4.PostingRegion.ADMINISTRATIVE_AREA],
	//  setting job [Job.addresses][google.cloud.talent.v4.Job.addresses] to the
	//  same location level as this field is strongly recommended.
	// +kcc:proto:field=google.cloud.talent.v4.Job.posting_region
	PostingRegion *string `json:"postingRegion,omitempty"`

	// The start timestamp of the job in UTC time zone. Typically this field
	//  is used for contracting engagements. Invalid timestamps are ignored.
	// +kcc:proto:field=google.cloud.talent.v4.Job.job_start_time
	JobStartTime *string `json:"jobStartTime,omitempty"`

	// The end timestamp of the job. Typically this field is used for contracting
	//  engagements. Invalid timestamps are ignored.
	// +kcc:proto:field=google.cloud.talent.v4.Job.job_end_time
	JobEndTime *string `json:"jobEndTime,omitempty"`

	// The timestamp this job posting was most recently published. The default
	//  value is the time the request arrives at the server. Invalid timestamps are
	//  ignored.
	// +kcc:proto:field=google.cloud.talent.v4.Job.posting_publish_time
	PostingPublishTime *string `json:"postingPublishTime,omitempty"`

	// Strongly recommended for the best service experience.
	//
	//  The expiration timestamp of the job. After this timestamp, the
	//  job is marked as expired, and it no longer appears in search results. The
	//  expired job can't be listed by the
	//  [ListJobs][google.cloud.talent.v4.JobService.ListJobs] API, but it can be
	//  retrieved with the [GetJob][google.cloud.talent.v4.JobService.GetJob] API
	//  or updated with the
	//  [UpdateJob][google.cloud.talent.v4.JobService.UpdateJob] API or deleted
	//  with the [DeleteJob][google.cloud.talent.v4.JobService.DeleteJob] API. An
	//  expired job can be updated and opened again by using a future expiration
	//  timestamp. Updating an expired job fails if there is another existing open
	//  job with same [company][google.cloud.talent.v4.Job.company],
	//  [language_code][google.cloud.talent.v4.Job.language_code] and
	//  [requisition_id][google.cloud.talent.v4.Job.requisition_id].
	//
	//  The expired jobs are retained in our system for 90 days. However, the
	//  overall expired job count cannot exceed 3 times the maximum number of
	//  open jobs over previous 7 days. If this threshold is exceeded,
	//  expired jobs are cleaned out in order of earliest expire time.
	//  Expired jobs are no longer accessible after they are cleaned
	//  out.
	//
	//  Invalid timestamps are ignored, and treated as expire time not provided.
	//
	//  If the timestamp is before the instant request is made, the job
	//  is treated as expired immediately on creation. This kind of job can
	//  not be updated. And when creating a job with past timestamp, the
	//  [posting_publish_time][google.cloud.talent.v4.Job.posting_publish_time]
	//  must be set before
	//  [posting_expire_time][google.cloud.talent.v4.Job.posting_expire_time]. The
	//  purpose of this feature is to allow other objects, such as
	//  [ApplicationInfo][google.cloud.talent.v4.Job.ApplicationInfo], to refer a
	//  job that didn't exist in the system prior to becoming expired. If you want
	//  to modify a job that was expired on creation, delete it and create a new
	//  one.
	//
	//  If this value isn't provided at the time of job creation or is invalid,
	//  the job posting expires after 30 days from the job's creation time. For
	//  example, if the job was created on 2017/01/01 13:00AM UTC with an
	//  unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC.
	//
	//  If this value isn't provided on job update, it depends on the field masks
	//  set by
	//  [UpdateJobRequest.update_mask][google.cloud.talent.v4.UpdateJobRequest.update_mask].
	//  If the field masks include
	//  [job_end_time][google.cloud.talent.v4.Job.job_end_time], or the masks are
	//  empty meaning that every field is updated, the job posting expires after 30
	//  days from the job's last update time. Otherwise the expiration date isn't
	//  updated.
	// +kcc:proto:field=google.cloud.talent.v4.Job.posting_expire_time
	PostingExpireTime *string `json:"postingExpireTime,omitempty"`

	// Options for job processing.
	// +kcc:proto:field=google.cloud.talent.v4.Job.processing_options
	ProcessingOptions *Job_ProcessingOptions `json:"processingOptions,omitempty"`
}

// JobsJobStatus defines the config connector machine state of JobsJob
type JobsJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the JobsJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *JobsJobObservedState `json:"observedState,omitempty"`
}

// JobsJobObservedState is the state of the JobsJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.talent.v4.Job
type JobsJobObservedState struct {
	// Job compensation information (a.k.a. "pay rate") i.e., the compensation
	//  that will paid to the employee.
	// +kcc:proto:field=google.cloud.talent.v4.Job.compensation_info
	CompensationInfo *CompensationInfoObservedState `json:"compensationInfo,omitempty"`

	// Output only. The timestamp when this job posting was created.
	// +kcc:proto:field=google.cloud.talent.v4.Job.posting_create_time
	PostingCreateTime *string `json:"postingCreateTime,omitempty"`

	// Output only. The timestamp when this job posting was last updated.
	// +kcc:proto:field=google.cloud.talent.v4.Job.posting_update_time
	PostingUpdateTime *string `json:"postingUpdateTime,omitempty"`

	// Output only. Display name of the company listing the job.
	// +kcc:proto:field=google.cloud.talent.v4.Job.company_display_name
	CompanyDisplayName *string `json:"companyDisplayName,omitempty"`

	// Output only. Derived details about the job posting.
	// +kcc:proto:field=google.cloud.talent.v4.Job.derived_info
	DerivedInfo *Job_DerivedInfo `json:"derivedInfo,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpjobsjob;gcpjobsjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// JobsJob is the Schema for the JobsJob API
// +k8s:openapi-gen=true
type JobsJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   JobsJobSpec   `json:"spec,omitempty"`
	Status JobsJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// JobsJobList contains a list of JobsJob
type JobsJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobsJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&JobsJob{}, &JobsJobList{})
}
