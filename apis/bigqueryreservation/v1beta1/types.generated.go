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

// +generated:types
// krm.group: bigqueryreservation.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.bigquery.reservation.v1
// resource: BigQueryReservationReservation:Reservation
// resource: BigQueryReservationAssignment:Assignment

package v1beta1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
)

// +kcc:proto=google.cloud.bigquery.reservation.v1.Reservation.ReplicationStatus
type Reservation_ReplicationStatus struct {
}

// +kcc:observedstate:proto=google.cloud.bigquery.reservation.v1.Reservation.ReplicationStatus
type Reservation_ReplicationStatusObservedState struct {
	// Output only. The last error encountered while trying to replicate changes
	//  from the primary to the secondary. This field is only available if the
	//  replication has not succeeded since.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.ReplicationStatus.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. The time at which the last error was encountered while
	//  trying to replicate changes from the primary to the secondary. This field
	//  is only available if the replication has not succeeded since.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.ReplicationStatus.last_error_time
	LastErrorTime *string `json:"lastErrorTime,omitempty"`

	// Output only. A timestamp corresponding to the last change on the primary
	//  that was successfully replicated to the secondary.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Reservation.ReplicationStatus.last_replication_time
	LastReplicationTime *string `json:"lastReplicationTime,omitempty"`
}
