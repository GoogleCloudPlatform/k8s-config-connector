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

package v1alpha1

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

func TestObservabilityDatasetIdentity_Conformance(t *testing.T) {
	identity.AssertConformance(t, identity.ConformanceCase[ObservabilityDatasetIdentity]{
		SampleExternal: "projects/my-project/locations/us-central1/buckets/my-bucket/datasets/my-dataset",
		WantHost:       "observability.googleapis.com",
		WantParent:     "projects/my-project/locations/us-central1/buckets/my-bucket",
		WantIdentity: &ObservabilityDatasetIdentity{
			Project:  "my-project",
			Location: "us-central1",
			Bucket:   "my-bucket",
			Dataset:  "my-dataset",
		},
		Ref: &ObservabilityDatasetRef{},
	})
}
