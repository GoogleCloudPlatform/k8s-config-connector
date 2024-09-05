// Copyright 2024 Google LLC
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

type BigQueryDatasetRef struct {
	// If provided must be in the format `bigquery.googleapis.com/projects/[project_id]/datasets/[dataset_id]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `PubSubTopic` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `PubSubTopic` resource.
	Namespace string `json:"namespace,omitempty"`
}
