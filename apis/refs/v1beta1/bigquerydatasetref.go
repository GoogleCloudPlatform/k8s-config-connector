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

// BigQueryDatasetRef is a reference to a BigQueryDataset resource.
type BigQueryDatasetRef struct {
	// Name of the referenced object.
	// +optional
	Name string `json:"name,omitempty"`

	// Namespace of the referenced object.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// The BigQueryDataset selfLink, when not managed by Config Connector.
	// +optional
	External string `json:"external,omitempty"`
}
