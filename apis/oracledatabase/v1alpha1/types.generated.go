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


// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseCharacterSet
type AutonomousDatabaseCharacterSet struct {
	// Identifier. The name of the Autonomous Database Character Set resource in
	//  the following format:
	//  projects/{project}/locations/{region}/autonomousDatabaseCharacterSets/{autonomous_database_character_set}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseCharacterSet.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.AutonomousDatabaseCharacterSet
type AutonomousDatabaseCharacterSetObservedState struct {
	// Output only. The character set type for the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseCharacterSet.character_set_type
	CharacterSetType *string `json:"characterSetType,omitempty"`

	// Output only. The character set name for the Autonomous Database which is
	//  the ID in the resource name.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabaseCharacterSet.character_set
	CharacterSet *string `json:"characterSet,omitempty"`
}
