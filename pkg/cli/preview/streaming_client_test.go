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

package preview

import (
	"context"
	"net/url"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type fakeListListener struct {
	beginCalled bool
	metadata    ListMetadata
	objects     []Object
	endCalled   bool
}

func (f *fakeListListener) OnListBegin(metadata ListMetadata) {
	f.beginCalled = true
	f.metadata = metadata
}

func (f *fakeListListener) OnListObject(ctx context.Context, obj Object) error {
	f.objects = append(f.objects, obj)
	return nil
}

func (f *fakeListListener) OnListEnd() {
	f.endCalled = true
}

func TestDecodeJSONResponse(t *testing.T) {
	tests := []struct {
		name          string
		jsonInput     string
		expectError   bool
		expectedMeta  ListMetadata
		expectedItems int
	}{
		{
			name: "standard order with multiple items",
			jsonInput: `{
				"apiVersion": "v1",
				"kind": "List",
				"metadata": {
					"resourceVersion": "12345"
				},
				"items": [
					{
						"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
						"kind": "PubSubTopic",
						"metadata": {
							"name": "topic-1",
							"namespace": "ns-1"
						}
					},
					{
						"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
						"kind": "PubSubTopic",
						"metadata": {
							"name": "topic-2",
							"namespace": "ns-1"
						}
					}
				]
			}`,
			expectError: false,
			expectedMeta: ListMetadata{
				APIVersion:      "v1",
				Kind:            "List",
				ResourceVersion: "12345",
			},
			expectedItems: 2,
		},
		{
			name: "contains unknown top-level field",
			jsonInput: `{
				"apiVersion": "v1",
				"kind": "List",
				"unknownField": {
					"some": "value",
					"nested": [1, 2, 3]
				},
				"metadata": {
					"resourceVersion": "54321"
				},
				"items": [
					{
						"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
						"kind": "PubSubTopic",
						"metadata": {
							"name": "topic-1"
						}
					}
				]
			}`,
			expectError: false,
			expectedMeta: ListMetadata{
				APIVersion:      "v1",
				Kind:            "List",
				ResourceVersion: "54321",
			},
			expectedItems: 1,
		},
		{
			name: "empty items array",
			jsonInput: `{
				"apiVersion": "v1",
				"kind": "List",
				"metadata": {
					"resourceVersion": "000"
				},
				"items": []
			}`,
			expectError: false,
			expectedMeta: ListMetadata{
				APIVersion:      "v1",
				Kind:            "List",
				ResourceVersion: "000",
			},
			expectedItems: 0,
		},
		{
			name: "invalid json format missing brackets",
			jsonInput: `{
				"apiVersion": "v1",
				"kind": "List"
				"metadata": {
					"resourceVersion": "000"
				}
			}`,
			expectError: true,
		},
		{
			name: "invalid json format missing end brace",
			jsonInput: `{
				"apiVersion": "v1",
				"kind": "List",
				"metadata": {
					"resourceVersion": "000"
				},
				"items": []
			`,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			client := &StreamingClient{}
			typeInfo := &typeInfo{
				factory: func() Object {
					return &unstructured.Unstructured{}
				},
			}
			listener := &fakeListListener{}
			u, _ := url.Parse("http://localhost")

			err := client.decodeJSONResponse(ctx, strings.NewReader(tc.jsonInput), typeInfo, listener, u)
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !listener.beginCalled {
				t.Errorf("OnListBegin was not called")
			}
			if !listener.endCalled {
				t.Errorf("OnListEnd was not called")
			}

			if listener.metadata != tc.expectedMeta {
				t.Errorf("expected metadata %+v, got %+v", tc.expectedMeta, listener.metadata)
			}

			if len(listener.objects) != tc.expectedItems {
				t.Errorf("expected %d items, got %d", tc.expectedItems, len(listener.objects))
			}
		})
	}
}
