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


// +kcc:proto=google.cloud.oracledatabase.v1.GiVersion
type GiVersion struct {
	// Identifier. The name of the Oracle Grid Infrastructure (GI) version
	//  resource with the format:
	//  projects/{project}/locations/{region}/giVersions/{gi_versions}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.GiVersion.name
	Name *string `json:"name,omitempty"`

	// Optional. version
	// +kcc:proto:field=google.cloud.oracledatabase.v1.GiVersion.version
	Version *string `json:"version,omitempty"`
}
