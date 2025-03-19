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

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig
type NotebookIdleShutdownConfig struct {
	// Required. Duration is accurate to the second. In Notebook, Idle Timeout is
	//  accurate to minute so the range of idle_timeout (second) is: 10 * 60 ~ 1440
	//  * 60.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig.idle_timeout
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Whether Idle Shutdown is disabled in this NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig.idle_shutdown_disabled
	IdleShutdownDisabled *bool `json:"idleShutdownDisabled,omitempty"`
}
