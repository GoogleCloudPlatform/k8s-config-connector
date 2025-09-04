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

// +kcc:proto=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy
type MaintenanceUpdatePolicy struct {
	// Preferred windows to perform maintenance. Currently limited to 1.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.maintenance_windows
	MaintenanceWindows []MaintenanceUpdatePolicy_MaintenanceWindow `json:"maintenanceWindows,omitempty"`

	/* NOTYTET: We temporarily hide this so we don't add it to the schema.
	// Periods to deny maintenance. Currently limited to 1.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.deny_maintenance_periods
	DenyMaintenancePeriods []MaintenanceUpdatePolicy_DenyMaintenancePeriod `json:"denyMaintenancePeriods,omitempty"`
	*/
}
