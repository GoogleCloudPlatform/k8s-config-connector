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


// +kcc:proto=google.cloud.kms.v1.KeyRing
type KeyRing struct {
}

// +kcc:proto=google.cloud.kms.v1.KeyRing
type KeyRingObservedState struct {
	// Output only. The resource name for the
	//  [KeyRing][google.cloud.kms.v1.KeyRing] in the format
	//  `projects/*/locations/*/keyRings/*`.
	// +kcc:proto:field=google.cloud.kms.v1.KeyRing.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which this [KeyRing][google.cloud.kms.v1.KeyRing]
	//  was created.
	// +kcc:proto:field=google.cloud.kms.v1.KeyRing.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
