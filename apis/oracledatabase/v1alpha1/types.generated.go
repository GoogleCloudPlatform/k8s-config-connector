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


// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDbVersion
type AutonomousDbVersion struct {
	// Identifier. The name of the Autonomous Database Version resource with the
	//  format:
	//  projects/{project}/locations/{region}/autonomousDbVersions/{autonomous_db_version}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDbVersion.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDbVersion
type AutonomousDbVersionObservedState struct {
	// Output only. An Oracle Database version for Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDbVersion.version
	Version *string `json:"version,omitempty"`

	// Output only. The Autonomous Database workload type.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDbVersion.db_workload
	DbWorkload *string `json:"dbWorkload,omitempty"`

	// Output only. A URL that points to a detailed description of the Autonomous
	//  Database version.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDbVersion.workload_uri
	WorkloadURI *string `json:"workloadURI,omitempty"`
}
