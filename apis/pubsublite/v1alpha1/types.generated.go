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


// +kcc:proto=google.cloud.pubsublite.v1.Reservation
type Reservation struct {
	// The name of the reservation.
	//  Structured like:
	//  projects/{project_number}/locations/{location}/reservations/{reservation_id}
	// +kcc:proto:field=google.cloud.pubsublite.v1.Reservation.name
	Name *string `json:"name,omitempty"`

	// The reserved throughput capacity. Every unit of throughput capacity is
	//  equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	//  messages.
	//
	//  Any topics which are declared as using capacity from a Reservation will
	//  consume resources from this reservation instead of being charged
	//  individually.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Reservation.throughput_capacity
	ThroughputCapacity *int64 `json:"throughputCapacity,omitempty"`
}
