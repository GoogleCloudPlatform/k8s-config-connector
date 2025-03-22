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

package v1beta1

// +kcc:proto=google.monitoring.v3.MutationRecord
type MutationRecord struct {
	// When the change occurred.
	// +kcc:proto:field=google.monitoring.v3.MutationRecord.mutate_time
	MutateTime *string `json:"mutateTime,omitempty"`

	// The email address of the user making the change.
	// +kcc:proto:field=google.monitoring.v3.MutationRecord.mutated_by
	MutatedBy *string `json:"mutatedBy,omitempty"`
}
