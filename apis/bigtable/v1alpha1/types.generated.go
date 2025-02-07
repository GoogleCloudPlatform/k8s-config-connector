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


// +kcc:proto=google.bigtable.admin.v2.HotTablet
type HotTablet struct {
	// The unique name of the hot tablet. Values are of the form
	//  `projects/{project}/instances/{instance}/clusters/{cluster}/hotTablets/[a-zA-Z0-9_-]*`.
	// +kcc:proto:field=google.bigtable.admin.v2.HotTablet.name
	Name *string `json:"name,omitempty"`

	// Name of the table that contains the tablet. Values are of the form
	//  `projects/{project}/instances/{instance}/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	// +kcc:proto:field=google.bigtable.admin.v2.HotTablet.table_name
	TableName *string `json:"tableName,omitempty"`

	// Tablet Start Key (inclusive).
	// +kcc:proto:field=google.bigtable.admin.v2.HotTablet.start_key
	StartKey *string `json:"startKey,omitempty"`

	// Tablet End Key (inclusive).
	// +kcc:proto:field=google.bigtable.admin.v2.HotTablet.end_key
	EndKey *string `json:"endKey,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.HotTablet
type HotTabletObservedState struct {
	// Output only. The start time of the hot tablet.
	// +kcc:proto:field=google.bigtable.admin.v2.HotTablet.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The end time of the hot tablet.
	// +kcc:proto:field=google.bigtable.admin.v2.HotTablet.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The average CPU usage spent by a node on this tablet over the
	//  start_time to end_time time range. The percentage is the amount of CPU used
	//  by the node to serve the tablet, from 0% (tablet was not interacted with)
	//  to 100% (the node spent all cycles serving the hot tablet).
	// +kcc:proto:field=google.bigtable.admin.v2.HotTablet.node_cpu_usage_percent
	NodeCpuUsagePercent *float32 `json:"nodeCpuUsagePercent,omitempty"`
}
