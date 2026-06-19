// Copyright 2026 Google LLC
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

// VertexAIEndpointRef is a reference to a VertexAIEndpoint resource.
type VertexAIEndpointRef struct {
	// If provided, must be in the format `projects/[projectId]/locations/[location]/endpoints/[endpointId]` or `projects/[projectId]/locations/[location]/publishers/[publisherId]/models/[modelId]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `VertexAIEndpoint` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `VertexAIEndpoint` resource.
	Namespace string `json:"namespace,omitempty"`
}
