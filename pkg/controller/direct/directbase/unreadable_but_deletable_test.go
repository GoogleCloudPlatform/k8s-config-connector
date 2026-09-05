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

package directbase

import (
	"fmt"
	"testing"

	"google.golang.org/api/googleapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestIsUnreadableButDeletable(t *testing.T) {
	bqGVK := schema.GroupVersionKind{
		Group:   "bigquery.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "BigQueryTable",
	}

	otherGVK := schema.GroupVersionKind{
		Group:   "pubsub.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "PubSubTopic",
	}

	tests := []struct {
		name     string
		gvk      schema.GroupVersionKind
		err      error
		expected bool
	}{
		{
			name:     "Nil error",
			gvk:      bqGVK,
			err:      nil,
			expected: false,
		},
		{
			name:     "BigQueryTable standard error (e.g. 400 Bad Request)",
			gvk:      bqGVK,
			err:      &googleapi.Error{Code: 400, Message: "Bad Request"},
			expected: true,
		},
		{
			name:     "BigQueryTable 404 error (Not Found)",
			gvk:      bqGVK,
			err:      &googleapi.Error{Code: 404, Message: "Not Found"},
			expected: false,
		},
		{
			name:     "BigQueryTable generic error",
			gvk:      bqGVK,
			err:      fmt.Errorf("some generic error"),
			expected: true,
		},
		{
			name:     "Other resource standard error (e.g. 400 Bad Request)",
			gvk:      otherGVK,
			err:      &googleapi.Error{Code: 400, Message: "Bad Request"},
			expected: false,
		},
		{
			name:     "Other resource 404 error (Not Found)",
			gvk:      otherGVK,
			err:      &googleapi.Error{Code: 404, Message: "Not Found"},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsUnreadableButDeletable(tc.gvk, tc.err)
			if got != tc.expected {
				t.Errorf("IsUnreadableButDeletable(%v, %v) = %v, want %v", tc.gvk, tc.err, got, tc.expected)
			}
		})
	}
}
