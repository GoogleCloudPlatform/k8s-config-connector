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
// krm.group: run.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.run.v2
// resource: RunJob:Job

package v1beta1

// +kcc:proto=google.cloud.run.v2.BuildInfo
type BuildInfo struct {
}

// +kcc:observedstate:proto=google.cloud.run.v2.BuildInfo
type BuildInfoObservedState struct {
	// Output only. Entry point of the function when the image is a Cloud Run
	//  function.
	// +kcc:proto:field=google.cloud.run.v2.BuildInfo.function_target
	FunctionTarget *string `json:"functionTarget,omitempty"`

	// Output only. Source code location of the image.
	// +kcc:proto:field=google.cloud.run.v2.BuildInfo.source_location
	SourceLocation *string `json:"sourceLocation,omitempty"`
}
