// Copyright 2024 Google LLC
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

package v1beta1

/*
// +kcc:proto=google.firestore.admin.v1.Database
type Database struct {
	// The resource name of the Database.
	//  Format: `projects/{project}/databases/{database}`
	Name *string `json:"name,omitempty"`

	// Output only. The system-generated UUID4 for this Database.
	Uid *string `json:"uid,omitempty"`

	// Output only. The timestamp at which this database was created. Databases
	//  created before 2016 do not populate create_time.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp at which this database was most recently
	//  updated. Note this only includes updates to the database resource and not
	//  data contained by the database.
	UpdateTime *string `json:"updateTime,omitempty"`

	// The location of the database. Available locations are listed at
	//  https://cloud.google.com/firestore/docs/locations.
	LocationID *string `json:"locationID,omitempty"`

	// The type of the database.
	//  See https://cloud.google.com/datastore/docs/firestore-or-datastore for
	//  information about how to choose.
	Type *string `json:"type,omitempty"`

	// The concurrency control mode to use for this database.
	ConcurrencyMode *string `json:"concurrencyMode,omitempty"`

	// Output only. The period during which past versions of data are retained in
	//  the database.
	//
	//  Any [read][google.firestore.v1.GetDocumentRequest.read_time]
	//  or [query][google.firestore.v1.ListDocumentsRequest.read_time] can specify
	//  a `read_time` within this window, and will read the state of the database
	//  at that time.
	//
	//  If the PITR feature is enabled, the retention period is 7 days. Otherwise,
	//  the retention period is 1 hour.
	VersionRetentionPeriod *string `json:"versionRetentionPeriod,omitempty"`

	// Output only. The earliest timestamp at which older versions of the data can
	//  be read from the database. See [version_retention_period] above; this field
	//  is populated with `now - version_retention_period`.
	//
	//  This value is continuously updated, and becomes stale the moment it is
	//  queried. If you are using this value to recover data, make sure to account
	//  for the time from the moment when the value is queried to the moment when
	//  you initiate the recovery.
	EarliestVersionTime *string `json:"earliestVersionTime,omitempty"`

	// Whether to enable the PITR feature on this database.
	PointInTimeRecoveryEnablement *string `json:"pointInTimeRecoveryEnablement,omitempty"`

	// The App Engine integration mode to use for this database.
	AppEngineIntegrationMode *string `json:"appEngineIntegrationMode,omitempty"`

	// Output only. The key_prefix for this database. This key_prefix is used, in
	//  combination with the project id ("<key prefix>~<project id>") to construct
	//  the application id that is returned from the Cloud Datastore APIs in Google
	//  App Engine first generation runtimes.
	//
	//  This value may be empty in which case the appid to use for URL-encoded keys
	//  is the project_id (eg: foo instead of v~foo).
	KeyPrefix *string `json:"keyPrefix,omitempty"`

	// State of delete protection for the database.
	DeleteProtectionState *string `json:"deleteProtectionState,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`
}
*/
