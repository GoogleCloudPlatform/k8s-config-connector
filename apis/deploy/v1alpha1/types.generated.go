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


// +kcc:proto=google.cloud.deploy.v1.Config
type Config struct {
	// Name of the configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.Config.name
	Name *string `json:"name,omitempty"`

	// All supported versions of Skaffold.
	// +kcc:proto:field=google.cloud.deploy.v1.Config.supported_versions
	SupportedVersions []SkaffoldVersion `json:"supportedVersions,omitempty"`

	// Default Skaffold version that is assigned when a Release is created without
	//  specifying a Skaffold version.
	// +kcc:proto:field=google.cloud.deploy.v1.Config.default_skaffold_version
	DefaultSkaffoldVersion *string `json:"defaultSkaffoldVersion,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldVersion
type SkaffoldVersion struct {
	// Release version number. For example, "1.20.3".
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldVersion.version
	Version *string `json:"version,omitempty"`

	// The time at which this version of Skaffold will enter maintenance mode.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldVersion.maintenance_mode_time
	MaintenanceModeTime *string `json:"maintenanceModeTime,omitempty"`

	// The time at which this version of Skaffold will no longer be supported.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldVersion.support_expiration_time
	SupportExpirationTime *string `json:"supportExpirationTime,omitempty"`

	// Date when this version is expected to no longer be supported.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldVersion.support_end_date
	SupportEndDate *Date `json:"supportEndDate,omitempty"`
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
