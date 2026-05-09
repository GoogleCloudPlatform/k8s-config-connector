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

type VertexAITensorboardRef struct {
	/* A reference to an externally managed Vertex AI Tensorboard resource.
	Should be of the format `projects/{{projectID}}/locations/{{location}}/tensorboards/{{tensorboardID}}`. */
	External string `json:"external,omitempty"`

	/* The `name` of a `VertexAITensorboard` resource. */
	Name string `json:"name,omitempty"`

	/* The `namespace` of a `VertexAITensorboard` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type VertexAIPersistentResourceRef struct {
	/* A reference to an externally managed Vertex AI PersistentResource resource.
	Should be of the format `projects/{{projectID}}/locations/{{location}}/persistentResources/{{persistentResourceID}}`. */
	External string `json:"external,omitempty"`

	/* The `name` of a `VertexAIPersistentResource` resource. */
	Name string `json:"name,omitempty"`

	/* The `namespace` of a `VertexAIPersistentResource` resource. */
	Namespace string `json:"namespace,omitempty"`
}
