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


// +kcc:proto=google.cloud.alloydb.v1.Database
type Database struct {
	// Identifier. Name of the resource in the form of
	//  `projects/{project}/locations/{location}/clusters/{cluster}/databases/{database}`.
	// +kcc:proto:field=google.cloud.alloydb.v1.Database.name
	Name *string `json:"name,omitempty"`

	// Optional. Charset for the database.
	//  This field can contain any PostgreSQL supported charset name.
	//  Example values include "UTF8", "SQL_ASCII", etc.
	// +kcc:proto:field=google.cloud.alloydb.v1.Database.charset
	Charset *string `json:"charset,omitempty"`

	// Optional. Collation for the database.
	//  Name of the custom or native collation for postgres.
	//  Example values include "C", "POSIX", etc
	// +kcc:proto:field=google.cloud.alloydb.v1.Database.collation
	Collation *string `json:"collation,omitempty"`
}
