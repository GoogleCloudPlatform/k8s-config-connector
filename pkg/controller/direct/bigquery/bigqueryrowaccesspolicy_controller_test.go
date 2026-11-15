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

package bigquery

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"google.golang.org/api/bigquery/v2"
)

func TestExtractGrantees(t *testing.T) {
	tests := []struct {
		name     string
		bindings []*bigquery.Binding
		want     []string
	}{
		{
			name:     "empty bindings",
			bindings: nil,
			want:     nil,
		},
		{
			name: "unrelated roles ignored",
			bindings: []*bigquery.Binding{
				{
					Role:    "roles/bigquery.dataViewer",
					Members: []string{"user:alice@example.com"},
				},
			},
			want: nil,
		},
		{
			name: "single binding sorted",
			bindings: []*bigquery.Binding{
				{
					Role:    "roles/bigquery.filteredDataViewer",
					Members: []string{"user:bob@example.com", "user:alice@example.com"},
				},
			},
			want: []string{"user:alice@example.com", "user:bob@example.com"},
		},
		{
			name: "multiple bindings with duplicates deduplicated",
			bindings: []*bigquery.Binding{
				{
					Role:    "roles/bigquery.filteredDataViewer",
					Members: []string{"user:bob@example.com", "user:alice@example.com"},
				},
				{
					Role:    "roles/bigquery.admin",
					Members: []string{"user:charlie@example.com"},
				},
				{
					Role:    "roles/bigquery.filteredDataViewer",
					Members: []string{"user:alice@example.com", "user:david@example.com"},
				},
			},
			want: []string{"user:alice@example.com", "user:bob@example.com", "user:david@example.com"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := extractGrantees(tc.bindings)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("extractGrantees() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRowAccessPolicy_ForceSendFields(t *testing.T) {
	// Without ForceSendFields, empty Grantees is omitted due to omitempty.
	p1 := &bigquery.RowAccessPolicy{
		FilterPredicate: "true",
	}
	b1, err := json.Marshal(p1)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	if strings.Contains(string(b1), "grantees") {
		t.Errorf("expected grantees to be omitted, got %s", string(b1))
	}

	// With ForceSendFields, empty Grantees is serialized as empty array.
	p2 := &bigquery.RowAccessPolicy{
		FilterPredicate: "true",
		ForceSendFields: []string{"Grantees"},
	}
	b2, err := json.Marshal(p2)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	if !strings.Contains(string(b2), `"grantees":[]`) {
		t.Errorf("expected grantees to be serialized as empty array, got %s", string(b2))
	}
}
