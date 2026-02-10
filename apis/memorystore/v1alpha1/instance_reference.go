// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

// InstanceRef defines the resource reference to MemorystoreInstance, which "External" field
// holds the GCP identifier for the KRM object.
type InstanceRef struct {
	// A reference to an externally managed MemorystoreInstance resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`

	// The name of a MemorystoreInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MemorystoreInstance resource.
	Namespace string `json:"namespace,omitempty"`
}
