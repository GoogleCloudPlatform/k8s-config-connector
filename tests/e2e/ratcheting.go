// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package e2e

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ShouldTestRereconiliation determines if we "touch" the primary object after we have run the test.
// This should not cause write operations to GCP (read operations are OK)
// We would like eventually to turn this on for all objects, but we have to turn on the testing gradually.
func ShouldTestRereconiliation(t *testing.T, primaryResource *unstructured.Unstructured) bool {
	gvk := primaryResource.GroupVersionKind()

	switch gvk.GroupKind() {
	case schema.GroupKind{Group: "pubsub.cnrm.cloud.google.com", Kind: "PubSubTopic"}:
		return true
	}

	switch gvk.Group {
	// case "pubsub.cnrm.cloud.google.com":
	// 	return true
	}

	t.Logf("defaulting ShouldTestRereconiliation to false for gvk %v", gvk)
	return false
}
