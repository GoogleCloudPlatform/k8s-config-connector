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


// +kcc:proto=google.cloud.talent.v4beta1.CompensationInfo
type CompensationInfo struct {
	// Job compensation information.
	//
	//  At most one entry can be of type
	//  [CompensationInfo.CompensationType.BASE][google.cloud.talent.v4beta1.CompensationInfo.CompensationType.BASE],
	//  which is referred as **base compensation entry** for the job.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.entries
	Entries []CompensationInfo_CompensationEntry `json:"entries,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry
type CompensationInfo_CompensationEntry struct {
	// Compensation type.
	//
	//  Default is
	//  [CompensationType.COMPENSATION_TYPE_UNSPECIFIED][google.cloud.talent.v4beta1.CompensationInfo.CompensationType.COMPENSATION_TYPE_UNSPECIFIED].
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.type
	Type *string `json:"type,omitempty"`

	// Frequency of the specified amount.
	//
	//  Default is
	//  [CompensationUnit.COMPENSATION_UNIT_UNSPECIFIED][google.cloud.talent.v4beta1.CompensationInfo.CompensationUnit.COMPENSATION_UNIT_UNSPECIFIED].
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.unit
	Unit *string `json:"unit,omitempty"`

	// Compensation amount.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.amount
	Amount *Money `json:"amount,omitempty"`

	// Compensation range.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.range
	Range *CompensationInfo_CompensationRange `json:"range,omitempty"`

	// Compensation description.  For example, could
	//  indicate equity terms or provide additional context to an estimated
	//  bonus.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.description
	Description *string `json:"description,omitempty"`

	// Expected number of units paid each year. If not specified, when
	//  [Job.employment_types][google.cloud.talent.v4beta1.Job.employment_types]
	//  is FULLTIME, a default value is inferred based on
	//  [unit][google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.unit].
	//  Default values:
	//  - HOURLY: 2080
	//  - DAILY: 260
	//  - WEEKLY: 52
	//  - MONTHLY: 12
	//  - ANNUAL: 1
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.expected_units_per_year
	ExpectedUnitsPerYear *DoubleValue `json:"expectedUnitsPerYear,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.CompensationInfo.CompensationRange
type CompensationInfo_CompensationRange struct {
	// The maximum amount of compensation. If left empty, the value is set
	//  to a maximal compensation value and the currency code is set to
	//  match the [currency code][google.type.Money.currency_code] of
	//  min_compensation.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.CompensationRange.max_compensation
	MaxCompensation *Money `json:"maxCompensation,omitempty"`

	// The minimum amount of compensation. If left empty, the value is set
	//  to zero and the currency code is set to match the
	//  [currency code][google.type.Money.currency_code] of max_compensation.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.CompensationRange.min_compensation
	MinCompensation *Money `json:"minCompensation,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.CustomAttribute
type CustomAttribute struct {
	// Exactly one of
	//  [string_values][google.cloud.talent.v4beta1.CustomAttribute.string_values]
	//  or [long_values][google.cloud.talent.v4beta1.CustomAttribute.long_values]
	//  must be specified.
	//
	//  This field is used to perform a string match (`CASE_SENSITIVE_MATCH` or
	//  `CASE_INSENSITIVE_MATCH`) search.
	//  For filterable `string_value`s, a maximum total number of 200 values
	//  is allowed, with each `string_value` has a byte size of no more than
	//  500B. For unfilterable `string_values`, the maximum total byte size of
	//  unfilterable `string_values` is 50KB.
	//
	//  Empty string isn't allowed.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CustomAttribute.string_values
	StringValues []string `json:"stringValues,omitempty"`

	// Exactly one of
	//  [string_values][google.cloud.talent.v4beta1.CustomAttribute.string_values]
	//  or [long_values][google.cloud.talent.v4beta1.CustomAttribute.long_values]
	//  must be specified.
	//
	//  This field is used to perform number range search.
	//  (`EQ`, `GT`, `GE`, `LE`, `LT`) over filterable `long_value`.
	//
	//  Currently at most 1
	//  [long_values][google.cloud.talent.v4beta1.CustomAttribute.long_values] is
	//  supported.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CustomAttribute.long_values
	LongValues []int64 `json:"longValues,omitempty"`

	// If the `filterable` flag is true, the custom field values may be used for
	//  custom attribute filters
	//  [JobQuery.custom_attribute_filter][google.cloud.talent.v4beta1.JobQuery.custom_attribute_filter].
	//  If false, these values may not be used for custom attribute filters.
	//
	//  Default is false.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CustomAttribute.filterable
	Filterable *bool `json:"filterable,omitempty"`

	// If the `keyword_searchable` flag is true, the keywords in custom fields are
	//  searchable by keyword match.
	//  If false, the values are not searchable by keyword match.
	//
	//  Default is false.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CustomAttribute.keyword_searchable
	KeywordSearchable *bool `json:"keywordSearchable,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.Job
type Job struct {
	// Required during job update.
	//
	//  The resource name for the job. This is generated by the service when a
	//  job is created.
	//
	//  The format is
	//  "projects/{project_id}/tenants/{tenant_id}/jobs/{job_id}". For
	//  example, "projects/foo/tenants/bar/jobs/baz".
	//
	//  If tenant id is unspecified, the default tenant is used. For
	//  example, "projects/foo/jobs/bar".
	//
	//  Use of this field in job queries and API calls is preferred over the use of
	//  [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id] since this
	//  value is unique.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.name
	Name *string `json:"name,omitempty"`

	// Required. The resource name of the company listing the job.
	//
	//  The format is
	//  "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}". For
	//  example, "projects/foo/tenants/bar/companies/baz".
	//
	//  If tenant id is unspecified, the default tenant is used. For
	//  example, "projects/foo/companies/bar".
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.company
	Company *string `json:"company,omitempty"`

	// Required. The requisition ID, also referred to as the posting ID, is
	//  assigned by the client to identify a job. This field is intended to be used
	//  by clients for client identification and tracking of postings. A job isn't
	//  allowed to be created if there is another job with the same
	//  [company][google.cloud.talent.v4beta1.Job.name],
	//  [language_code][google.cloud.talent.v4beta1.Job.language_code] and
	//  [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id].
	//
	//  The maximum number of allowed characters is 255.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.requisition_id
	RequisitionID *string `json:"requisitionID,omitempty"`

	// Required. The title of the job, such as "Software Engineer"
	//
	//  The maximum number of allowed characters is 500.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.title
	Title *string `json:"title,omitempty"`

	// Required. The description of the job, which typically includes a
	//  multi-paragraph description of the company and related information.
	//  Separate fields are provided on the job object for
	//  [responsibilities][google.cloud.talent.v4beta1.Job.responsibilities],
	//  [qualifications][google.cloud.talent.v4beta1.Job.qualifications], and other
	//  job characteristics. Use of these separate job fields is recommended.
	//
	//  This field accepts and sanitizes HTML input, and also accepts
	//  bold, italic, ordered list, and unordered list markup tags.
	//
	//  The maximum number of allowed characters is 100,000.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.description
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
	//  [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id]s (e.g.
	//  'ReqA' becomes 'ReqA-1', 'ReqA-2', and so on.) as multiple jobs with the
	//  same [company][google.cloud.talent.v4beta1.Job.company],
	//  [language_code][google.cloud.talent.v4beta1.Job.language_code] and
	//  [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id] are not
	//  allowed. If the original
	//  [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id] must be
	//  preserved, a custom field should be used for storage. It is also suggested
	//  to group the locations that close to each other in the same job for better
	//  search experience.
	//
	//  The maximum number of allowed characters is 500.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.addresses
	Addresses []string `json:"addresses,omitempty"`

	// Job application information.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.application_info
	ApplicationInfo *Job_ApplicationInfo `json:"applicationInfo,omitempty"`

	// The benefits included with the job.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.job_benefits
	JobBenefits []string `json:"jobBenefits,omitempty"`

	// Job compensation information (a.k.a. "pay rate") i.e., the compensation
	//  that will paid to the employee.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.compensation_info
	CompensationInfo *CompensationInfo `json:"compensationInfo,omitempty"`

	// TODO: unsupported map type with key string and value message


	// The desired education degrees for the job, such as Bachelors, Masters.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.degree_types
	DegreeTypes []string `json:"degreeTypes,omitempty"`

	// The department or functional area within the company with the open
	//  position.
	//
	//  The maximum number of allowed characters is 255.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.department
	Department *string `json:"department,omitempty"`

	// The employment type(s) of a job, for example,
	//  [full time][google.cloud.talent.v4beta1.EmploymentType.FULL_TIME] or
	//  [part time][google.cloud.talent.v4beta1.EmploymentType.PART_TIME].
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.employment_types
	EmploymentTypes []string `json:"employmentTypes,omitempty"`

	// A description of bonus, commission, and other compensation
	//  incentives associated with the job not including salary or pay.
	//
	//  The maximum number of allowed characters is 10,000.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.incentives
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
	//  [Job.description][google.cloud.talent.v4beta1.Job.description] is present,
	//  detected language code based on
	//  [Job.description][google.cloud.talent.v4beta1.Job.description] is assigned,
	//  otherwise defaults to 'en_US'.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The experience level associated with the job, such as "Entry Level".
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.job_level
	JobLevel *string `json:"jobLevel,omitempty"`

	// A promotion value of the job, as determined by the client.
	//  The value determines the sort order of the jobs returned when searching for
	//  jobs using the featured jobs search call, with higher promotional values
	//  being returned first and ties being resolved by relevance sort. Only the
	//  jobs with a promotionValue >0 are returned in a FEATURED_JOB_SEARCH.
	//
	//  Default value is 0, and negative values are treated as 0.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.promotion_value
	PromotionValue *int32 `json:"promotionValue,omitempty"`

	// A description of the qualifications required to perform the
	//  job. The use of this field is recommended
	//  as an alternative to using the more general
	//  [description][google.cloud.talent.v4beta1.Job.description] field.
	//
	//  This field accepts and sanitizes HTML input, and also accepts
	//  bold, italic, ordered list, and unordered list markup tags.
	//
	//  The maximum number of allowed characters is 10,000.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.qualifications
	Qualifications *string `json:"qualifications,omitempty"`

	// A description of job responsibilities. The use of this field is
	//  recommended as an alternative to using the more general
	//  [description][google.cloud.talent.v4beta1.Job.description] field.
	//
	//  This field accepts and sanitizes HTML input, and also accepts
	//  bold, italic, ordered list, and unordered list markup tags.
	//
	//  The maximum number of allowed characters is 10,000.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.responsibilities
	Responsibilities *string `json:"responsibilities,omitempty"`

	// The job [PostingRegion][google.cloud.talent.v4beta1.PostingRegion] (for
	//  example, state, country) throughout which the job is available. If this
	//  field is set, a
	//  [LocationFilter][google.cloud.talent.v4beta1.LocationFilter] in a search
	//  query within the job region finds this job posting if an exact location
	//  match isn't specified. If this field is set to
	//  [PostingRegion.NATION][google.cloud.talent.v4beta1.PostingRegion.NATION] or
	//  [PostingRegion.ADMINISTRATIVE_AREA][google.cloud.talent.v4beta1.PostingRegion.ADMINISTRATIVE_AREA],
	//  setting job [Job.addresses][google.cloud.talent.v4beta1.Job.addresses] to
	//  the same location level as this field is strongly recommended.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.posting_region
	PostingRegion *string `json:"postingRegion,omitempty"`

	// Deprecated. The job is only visible to the owner.
	//
	//  The visibility of the job.
	//
	//  Defaults to
	//  [Visibility.ACCOUNT_ONLY][google.cloud.talent.v4beta1.Visibility.ACCOUNT_ONLY]
	//  if not specified.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.visibility
	Visibility *string `json:"visibility,omitempty"`

	// The start timestamp of the job in UTC time zone. Typically this field
	//  is used for contracting engagements. Invalid timestamps are ignored.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.job_start_time
	JobStartTime *string `json:"jobStartTime,omitempty"`

	// The end timestamp of the job. Typically this field is used for contracting
	//  engagements. Invalid timestamps are ignored.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.job_end_time
	JobEndTime *string `json:"jobEndTime,omitempty"`

	// The timestamp this job posting was most recently published. The default
	//  value is the time the request arrives at the server. Invalid timestamps are
	//  ignored.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.posting_publish_time
	PostingPublishTime *string `json:"postingPublishTime,omitempty"`

	// Strongly recommended for the best service experience.
	//
	//  The expiration timestamp of the job. After this timestamp, the
	//  job is marked as expired, and it no longer appears in search results. The
	//  expired job can't be listed by the
	//  [ListJobs][google.cloud.talent.v4beta1.JobService.ListJobs] API, but it can
	//  be retrieved with the
	//  [GetJob][google.cloud.talent.v4beta1.JobService.GetJob] API or updated with
	//  the [UpdateJob][google.cloud.talent.v4beta1.JobService.UpdateJob] API or
	//  deleted with the
	//  [DeleteJob][google.cloud.talent.v4beta1.JobService.DeleteJob] API. An
	//  expired job can be updated and opened again by using a future expiration
	//  timestamp. Updating an expired job fails if there is another existing open
	//  job with same [company][google.cloud.talent.v4beta1.Job.company],
	//  [language_code][google.cloud.talent.v4beta1.Job.language_code] and
	//  [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id].
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
	//  [posting_publish_time][google.cloud.talent.v4beta1.Job.posting_publish_time]
	//  must be set before
	//  [posting_expire_time][google.cloud.talent.v4beta1.Job.posting_expire_time].
	//  The purpose of this feature is to allow other objects, such as
	//  [Application][google.cloud.talent.v4beta1.Application], to refer a job that
	//  didn't exist in the system prior to becoming expired. If you want to modify
	//  a job that was expired on creation, delete it and create a new one.
	//
	//  If this value isn't provided at the time of job creation or is invalid,
	//  the job posting expires after 30 days from the job's creation time. For
	//  example, if the job was created on 2017/01/01 13:00AM UTC with an
	//  unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC.
	//
	//  If this value isn't provided on job update, it depends on the field masks
	//  set by
	//  [UpdateJobRequest.update_mask][google.cloud.talent.v4beta1.UpdateJobRequest.update_mask].
	//  If the field masks include
	//  [job_end_time][google.cloud.talent.v4beta1.Job.job_end_time], or the masks
	//  are empty meaning that every field is updated, the job posting expires
	//  after 30 days from the job's last update time. Otherwise the expiration
	//  date isn't updated.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.posting_expire_time
	PostingExpireTime *string `json:"postingExpireTime,omitempty"`

	// Options for job processing.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.processing_options
	ProcessingOptions *Job_ProcessingOptions `json:"processingOptions,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.Job.ApplicationInfo
type Job_ApplicationInfo struct {
	// Use this field to specify email address(es) to which resumes or
	//  applications can be sent.
	//
	//  The maximum number of allowed characters for each entry is 255.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.ApplicationInfo.emails
	Emails []string `json:"emails,omitempty"`

	// Use this field to provide instructions, such as "Mail your application
	//  to ...", that a candidate can follow to apply for the job.
	//
	//  This field accepts and sanitizes HTML input, and also accepts
	//  bold, italic, ordered list, and unordered list markup tags.
	//
	//  The maximum number of allowed characters is 3,000.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.ApplicationInfo.instruction
	Instruction *string `json:"instruction,omitempty"`

	// Use this URI field to direct an applicant to a website, for example to
	//  link to an online application form.
	//
	//  The maximum number of allowed characters for each entry is 2,000.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.ApplicationInfo.uris
	Uris []string `json:"uris,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.Job.DerivedInfo
type Job_DerivedInfo struct {
	// Structured locations of the job, resolved from
	//  [Job.addresses][google.cloud.talent.v4beta1.Job.addresses].
	//
	//  [locations][google.cloud.talent.v4beta1.Job.DerivedInfo.locations] are
	//  exactly matched to
	//  [Job.addresses][google.cloud.talent.v4beta1.Job.addresses] in the same
	//  order.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.DerivedInfo.locations
	Locations []Location `json:"locations,omitempty"`

	// Job categories derived from
	//  [Job.title][google.cloud.talent.v4beta1.Job.title] and
	//  [Job.description][google.cloud.talent.v4beta1.Job.description].
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.DerivedInfo.job_categories
	JobCategories []string `json:"jobCategories,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.Job.ProcessingOptions
type Job_ProcessingOptions struct {
	// If set to `true`, the service does not attempt to resolve a
	//  more precise address for the job.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.ProcessingOptions.disable_street_address_resolution
	DisableStreetAddressResolution *bool `json:"disableStreetAddressResolution,omitempty"`

	// Option for job HTML content sanitization. Applied fields are:
	//
	//  * description
	//  * applicationInfo.instruction
	//  * incentives
	//  * qualifications
	//  * responsibilities
	//
	//  HTML tags in these fields may be stripped if sanitiazation isn't
	//  disabled.
	//
	//  Defaults to
	//  [HtmlSanitization.SIMPLE_FORMATTING_ONLY][google.cloud.talent.v4beta1.HtmlSanitization.SIMPLE_FORMATTING_ONLY].
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.ProcessingOptions.html_sanitization
	HTMLSanitization *string `json:"htmlSanitization,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.Location
type Location struct {
	// The type of a location, which corresponds to the address lines field of
	//  [google.type.PostalAddress][google.type.PostalAddress]. For example,
	//  "Downtown, Atlanta, GA, USA" has a type of
	//  [LocationType.NEIGHBORHOOD][google.cloud.talent.v4beta1.Location.LocationType.NEIGHBORHOOD],
	//  and "Kansas City, KS, USA" has a type of
	//  [LocationType.LOCALITY][google.cloud.talent.v4beta1.Location.LocationType.LOCALITY].
	// +kcc:proto:field=google.cloud.talent.v4beta1.Location.location_type
	LocationType *string `json:"locationType,omitempty"`

	// Postal address of the location that includes human readable information,
	//  such as postal delivery and payments addresses. Given a postal address,
	//  a postal service can deliver items to a premises, P.O. Box, or other
	//  delivery location.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Location.postal_address
	PostalAddress *PostalAddress `json:"postalAddress,omitempty"`

	// An object representing a latitude/longitude pair.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Location.lat_lng
	LatLng *LatLng `json:"latLng,omitempty"`

	// Radius in miles of the job location. This value is derived from the
	//  location bounding box in which a circle with the specified radius
	//  centered from [google.type.LatLng][google.type.LatLng] covers the area
	//  associated with the job location. For example, currently, "Mountain View,
	//  CA, USA" has a radius of 6.17 miles.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Location.radius_miles
	RadiusMiles *float64 `json:"radiusMiles,omitempty"`
}

// +kcc:proto=google.protobuf.DoubleValue
type DoubleValue struct {
	// The double value.
	// +kcc:proto:field=google.protobuf.DoubleValue.value
	Value *float64 `json:"value,omitempty"`
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

// +kcc:proto=google.type.Money
type Money struct {
	// The three-letter currency code defined in ISO 4217.
	// +kcc:proto:field=google.type.Money.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// The whole units of the amount.
	//  For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	// +kcc:proto:field=google.type.Money.units
	Units *int64 `json:"units,omitempty"`

	// Number of nano (10^-9) units of the amount.
	//  The value must be between -999,999,999 and +999,999,999 inclusive.
	//  If `units` is positive, `nanos` must be positive or zero.
	//  If `units` is zero, `nanos` can be positive, zero, or negative.
	//  If `units` is negative, `nanos` must be negative or zero.
	//  For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	// +kcc:proto:field=google.type.Money.nanos
	Nanos *int32 `json:"nanos,omitempty"`
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

// +kcc:proto=google.cloud.talent.v4beta1.CompensationInfo
type CompensationInfoObservedState struct {
	// Output only. Annualized base compensation range. Computed as base
	//  compensation entry's
	//  [CompensationEntry.amount][google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.amount]
	//  times
	//  [CompensationEntry.expected_units_per_year][google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.expected_units_per_year].
	//
	//  See
	//  [CompensationEntry][google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry]
	//  for explanation on compensation annualization.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.annualized_base_compensation_range
	AnnualizedBaseCompensationRange *CompensationInfo_CompensationRange `json:"annualizedBaseCompensationRange,omitempty"`

	// Output only. Annualized total compensation range. Computed as all
	//  compensation entries'
	//  [CompensationEntry.amount][google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.amount]
	//  times
	//  [CompensationEntry.expected_units_per_year][google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry.expected_units_per_year].
	//
	//  See
	//  [CompensationEntry][google.cloud.talent.v4beta1.CompensationInfo.CompensationEntry]
	//  for explanation on compensation annualization.
	// +kcc:proto:field=google.cloud.talent.v4beta1.CompensationInfo.annualized_total_compensation_range
	AnnualizedTotalCompensationRange *CompensationInfo_CompensationRange `json:"annualizedTotalCompensationRange,omitempty"`
}

// +kcc:proto=google.cloud.talent.v4beta1.Job
type JobObservedState struct {
	// Job compensation information (a.k.a. "pay rate") i.e., the compensation
	//  that will paid to the employee.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.compensation_info
	CompensationInfo *CompensationInfoObservedState `json:"compensationInfo,omitempty"`

	// Output only. The timestamp when this job posting was created.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.posting_create_time
	PostingCreateTime *string `json:"postingCreateTime,omitempty"`

	// Output only. The timestamp when this job posting was last updated.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.posting_update_time
	PostingUpdateTime *string `json:"postingUpdateTime,omitempty"`

	// Output only. Display name of the company listing the job.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.company_display_name
	CompanyDisplayName *string `json:"companyDisplayName,omitempty"`

	// Output only. Derived details about the job posting.
	// +kcc:proto:field=google.cloud.talent.v4beta1.Job.derived_info
	DerivedInfo *Job_DerivedInfo `json:"derivedInfo,omitempty"`
}
