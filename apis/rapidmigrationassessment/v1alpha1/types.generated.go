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


// +kcc:proto=google.cloud.rapidmigrationassessment.v1.Collector
type Collector struct {
	// name of resource.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.labels
	Labels map[string]string `json:"labels,omitempty"`

	// User specified name of the Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// User specified description of the Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.description
	Description *string `json:"description,omitempty"`

	// Service Account email used to ingest data to this Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// User specified expected asset count.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.expected_asset_count
	ExpectedAssetCount *int64 `json:"expectedAssetCount,omitempty"`

	// How many days to collect data.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.collection_days
	CollectionDays *int32 `json:"collectionDays,omitempty"`

	// Uri for EULA (End User License Agreement) from customer.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.eula_uri
	EulaURI *string `json:"eulaURI,omitempty"`
}

// +kcc:proto=google.cloud.rapidmigrationassessment.v1.GuestOsScan
type GuestOsScan struct {
	// reference to the corresponding Guest OS Scan in MC Source.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.GuestOsScan.core_source
	CoreSource *string `json:"coreSource,omitempty"`
}

// +kcc:proto=google.cloud.rapidmigrationassessment.v1.VSphereScan
type VSphereScan struct {
	// reference to the corresponding VSphere Scan in MC Source.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.VSphereScan.core_source
	CoreSource *string `json:"coreSource,omitempty"`
}

// +kcc:proto=google.cloud.rapidmigrationassessment.v1.Collector
type CollectorObservedState struct {
	// Output only. Create time stamp.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Store cloud storage bucket name (which is a guid) created with
	//  this Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Output only. State of the Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.state
	State *string `json:"state,omitempty"`

	// Output only. Client version.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.client_version
	ClientVersion *string `json:"clientVersion,omitempty"`

	// Output only. Reference to MC Source Guest Os Scan.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.guest_os_scan
	GuestOsScan *GuestOsScan `json:"guestOsScan,omitempty"`

	// Output only. Reference to MC Source vsphere_scan.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.vsphere_scan
	VsphereScan *VSphereScan `json:"vsphereScan,omitempty"`
}
