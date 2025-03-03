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

import (
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
)

// +kcc:proto=google.cloud.bigquery.v2.Access
type Access struct {
	// An IAM role ID that should be granted to the user, group,
	//  or domain specified in this access entry.
	//  The following legacy mappings will be applied:
	//
	//  * `OWNER`: `roles/bigquery.dataOwner`
	//  * `WRITER`: `roles/bigquery.dataEditor`
	//  * `READER`: `roles/bigquery.dataViewer`
	//
	//  This field will accept any of the above formats, but will return only
	//  the legacy format. For example, if you set this field to
	//  "roles/bigquery.dataOwner", it will be returned back as "OWNER".
	Role *string `json:"role,omitempty"`

	// [Pick one] An email address of a user to grant access to. For example:
	//  fred@example.com. Maps to IAM policy member "user:EMAIL" or
	//  "serviceAccount:EMAIL".
	UserByEmail *string `json:"userByEmail,omitempty"`

	// [Pick one] An email address of a Google Group to grant access to.
	//  Maps to IAM policy member "group:GROUP".
	GroupByEmail *string `json:"groupByEmail,omitempty"`

	// [Pick one] A domain to grant access to. Any users signed in with the domain
	//  specified will be granted the specified access. Example: "example.com".
	//  Maps to IAM policy member "domain:DOMAIN".
	Domain *string `json:"domain,omitempty"`

	// [Pick one] A special group to grant access to. Possible values include:
	//
	//    * projectOwners: Owners of the enclosing project.
	//    * projectReaders: Readers of the enclosing project.
	//    * projectWriters: Writers of the enclosing project.
	//    * allAuthenticatedUsers: All authenticated BigQuery users.
	//
	//  Maps to similarly-named IAM members.
	SpecialGroup *string `json:"specialGroup,omitempty"`

	// [Pick one] Some other type of member that appears in the IAM Policy but
	//  isn't a user, group, domain, or special group.
	IamMember *string `json:"iamMember,omitempty"`

	// [Pick one] A view from a different dataset to grant access to. Queries
	//  executed against that view will have read access to views/tables/routines
	//  in this dataset.
	//  The role field is not required when this field is set. If that view is
	//  updated by any user, access to the view needs to be granted again via an
	//  update operation.
	View *TableReference `json:"view,omitempty"`

	// [Pick one] A routine from a different dataset to grant access to. Queries
	//  executed against that routine will have read access to
	//  views/tables/routines in this dataset. Only UDF is supported for now.
	//  The role field is not required when this field is set. If that routine is
	//  updated by any user, access to the routine needs to be granted again via
	//  an update operation.
	Routine *RoutineReference `json:"routine,omitempty"`

	// [Pick one] A grant authorizing all resources of a particular type in a
	//  particular dataset access to this dataset. Only views are supported for
	//  now. The role field is not required when this field is set. If that dataset
	//  is deleted and re-created, its access needs to be granted again via an
	//  update operation.
	Dataset *DatasetAccessEntry `json:"dataset,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DatasetAccessEntry
type DatasetAccessEntry struct {
	// The dataset this entry applies to.
	// +required
	Dataset *DatasetReference `json:"dataset,omitempty"`

	// Which resources in the dataset this entry applies to. Currently, only
	//  views are supported, but additional target types may be added in the
	//  future.
	// +required
	TargetTypes []string `json:"targetTypes,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DatasetReference
type DatasetReference struct {
	// A unique Id for this dataset, without the project name. The Id
	//  must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	//  The maximum length is 1,024 characters.
	// +required
	DatasetId *string `json:"datasetId,omitempty"`

	// The ID of the project containing this dataset.
	// +required
	ProjectId *string `json:"projectId,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.EncryptionConfiguration
type EncryptionConfiguration struct {
	// Optional. Describes the Cloud KMS encryption key that will be used to
	//  protect destination BigQuery table. The BigQuery Service Account associated
	//  with your project requires access to this encryption key.
	KmsKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalCatalogDatasetOptions
type ExternalCatalogDatasetOptions struct {
	// Optional. A map of key value pairs defining the parameters and properties
	//  of the open source schema. Maximum size of 2Mib.
	Parameters map[string]string `json:"parameters,omitempty"`

	// Optional. The storage location URI for all tables in the dataset.
	//  Equivalent to hive metastore's database locationUri. Maximum length of 1024
	//  characters.
	DefaultStorageLocationUri *string `json:"defaultStorageLocationUri,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalDatasetReference
type ExternalDatasetReference struct {
	// +required. External source that backs this dataset.
	ExternalSource *string `json:"externalSource,omitempty"`

	// +required. The connection id that is used to access the external_source.
	//
	//  Format:
	//    projects/{project_id}/locations/{location_id}/connections/{connection_id}
	Connection *string `json:"connection,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.GcpTag
type GcpTag struct {
	// Required. The namespaced friendly name of the tag key, e.g.
	//  "12345/environment" where 12345 is org id.
	TagKey *string `json:"tagKey,omitempty"`

	// Required. The friendly short name of the tag value, e.g. "production".
	TagValue *string `json:"tagValue,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.LinkedDatasetSource
type LinkedDatasetSource struct {
	// The source dataset reference contains project numbers and not project ids.
	SourceDataset *DatasetReference `json:"sourceDataset,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RestrictionConfig
type RestrictionConfig struct {
	// Output only. Specifies the type of dataset/table restriction.
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RoutineReference
type RoutineReference struct {
	// The ID of the project containing this routine.
	// +required
	ProjectId *string `json:"projectId,omitempty"`

	// The ID of the dataset containing this routine.
	// +required
	DatasetId *string `json:"datasetId,omitempty"`

	// The Id of the routine. The Id must contain only
	// letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 256 characters.
	// +required
	RoutineId *string `json:"routineId,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableReference
type TableReference struct {
	// The ID of the project containing this table.
	// +required
	ProjectId *string `json:"projectId,omitempty"`

	// The ID of the dataset containing this table.
	// +required
	DatasetId *string `json:"datasetId,omitempty"`

	// The Id of the table. The Id can contain Unicode characters in
	//  category L (letter), M (mark), N (number), Pc (connector, including
	//  underscore), Pd (dash), and Zs (space). For more information, see [General
	//  Category](https://wikipedia.org/wiki/Unicode_character_property#General_Category).
	//  The maximum length is 1,024 characters.  Certain operations allow suffixing
	//  of the table Id with a partition decorator, such as
	//  `sample_table$20190123`.
	// +required
	TableId *string `json:"tableId,omitempty"`
}

// +kcc:proto=google.protobuf.BoolValue
type BoolValue struct {
	// The bool value.
	Value *bool `json:"value,omitempty"`
}
