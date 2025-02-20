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

// +kcc:proto=google.monitoring.metricsscope.v1.MonitoredProject
type MonitoredProject struct {
	// Immutable. The resource name of the `MonitoredProject`. On input, the resource name
	//  includes the scoping project ID and monitored project ID. On output, it
	//  contains the equivalent project numbers.
	//  Example:
	//  `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
	// +kcc:proto:field=google.monitoring.metricsscope.v1.MonitoredProject.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.monitoring.metricsscope.v1.MonitoredProject
type MonitoredProjectObservedState struct {
	// Output only. The time when this `MonitoredProject` was created.
	// +kcc:proto:field=google.monitoring.metricsscope.v1.MonitoredProject.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
