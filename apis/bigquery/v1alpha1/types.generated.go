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


// +kcc:proto=google.cloud.bigquery.reservation.v1.BiReservation
type BiReservation struct {
	// The resource name of the singleton BI reservation.
	//  Reservation names have the form
	//  `projects/{project_id}/locations/{location_id}/biReservation`.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.BiReservation.name
	Name *string `json:"name,omitempty"`

	// Size of a reservation, in bytes.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.BiReservation.size
	Size *int64 `json:"size,omitempty"`

	// Preferred tables to use BI capacity for.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.BiReservation.preferred_tables
	PreferredTables []TableReference `json:"preferredTables,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.reservation.v1.TableReference
type TableReference struct {
	// The assigned project ID of the project.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.TableReference.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// The ID of the dataset in the above project.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.TableReference.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`

	// The ID of the table in the above dataset.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.TableReference.table_id
	TableID *string `json:"tableID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.reservation.v1.BiReservation
type BiReservationObservedState struct {
	// Output only. The last update timestamp of a reservation.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.BiReservation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
