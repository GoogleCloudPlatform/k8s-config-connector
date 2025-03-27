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
// krm.group: spanner.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.spanner.admin.instance.v1
// resource: SpannerInstanceConfig:InstanceConfig

package v1alpha1

// +kcc:proto=google.spanner.admin.instance.v1.ReplicaInfo
type ReplicaInfo struct {
	// The location of the serving resources, e.g., "us-central1".
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.location
	Location *string `json:"location,omitempty"`

	// The type of replica.
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.type
	Type *string `json:"type,omitempty"`

	// If true, this location is designated as the default leader location where
	//  leader replicas are placed. See the [region types
	//  documentation](https://cloud.google.com/spanner/docs/instances#region_types)
	//  for more details.
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.default_leader_location
	DefaultLeaderLocation *bool `json:"defaultLeaderLocation,omitempty"`
}
