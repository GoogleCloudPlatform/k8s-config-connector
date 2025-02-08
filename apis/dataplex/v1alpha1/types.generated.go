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


// +kcc:proto=google.cloud.dataplex.v1.Partition
type Partition struct {

	// Required. Immutable. The set of values representing the partition, which
	//  correspond to the partition schema defined in the parent entity.
	// +kcc:proto:field=google.cloud.dataplex.v1.Partition.values
	Values []string `json:"values,omitempty"`

	// Required. Immutable. The location of the entity data within the partition,
	//  for example, `gs://bucket/path/to/entity/key1=value1/key2=value2`. Or
	//  `projects/<project_id>/datasets/<dataset_id>/tables/<table_id>`
	// +kcc:proto:field=google.cloud.dataplex.v1.Partition.location
	Location *string `json:"location,omitempty"`

	// Optional. The etag for this partition.
	// +kcc:proto:field=google.cloud.dataplex.v1.Partition.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Partition
type PartitionObservedState struct {
	// Output only. Partition values used in the HTTP URL must be
	//  double encoded. For example, `url_encode(url_encode(value))` can be used
	//  to encode "US:CA/CA#Sunnyvale so that the request URL ends
	//  with "/partitions/US%253ACA/CA%2523Sunnyvale".
	//  The name field in the response retains the encoded format.
	// +kcc:proto:field=google.cloud.dataplex.v1.Partition.name
	Name *string `json:"name,omitempty"`
}
