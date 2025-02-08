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


// +kcc:proto=google.cloud.gkemulticloud.v1.AwsK8sVersionInfo
type AwsK8sVersionInfo struct {
	// Kubernetes version name.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsK8sVersionInfo.version
	Version *string `json:"version,omitempty"`

	// Optional. True if the version is available for cluster creation. If a
	//  version is enabled for creation, it can be used to create new clusters.
	//  Otherwise, cluster creation will fail. However, cluster upgrade operations
	//  may succeed, even if the version is not enabled.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsK8sVersionInfo.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. True if this cluster version belongs to a minor version that has
	//  reached its end of life and is no longer in scope to receive security and
	//  bug fixes.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsK8sVersionInfo.end_of_life
	EndOfLife *bool `json:"endOfLife,omitempty"`

	// Optional. The estimated date (in Pacific Time) when this cluster version
	//  will reach its end of life. Or if this version is no longer supported (the
	//  `end_of_life` field is true), this is the actual date (in Pacific time)
	//  when the version reached its end of life.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsK8sVersionInfo.end_of_life_date
	EndOfLifeDate *Date `json:"endOfLifeDate,omitempty"`

	// Optional. The date (in Pacific Time) when the cluster version was released.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsK8sVersionInfo.release_date
	ReleaseDate *Date `json:"releaseDate,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AwsServerConfig
type AwsServerConfig struct {
	// The resource name of the config.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsServerConfig.name
	Name *string `json:"name,omitempty"`

	// List of all released Kubernetes versions, including ones which are end of
	//  life and can no longer be used.  Filter by the `enabled`
	//  property to limit to currently available versions.
	//  Valid versions supported for both create and update operations
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsServerConfig.valid_versions
	ValidVersions []AwsK8sVersionInfo `json:"validVersions,omitempty"`

	// The list of supported AWS regions.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsServerConfig.supported_aws_regions
	SupportedAwsRegions []string `json:"supportedAwsRegions,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}
