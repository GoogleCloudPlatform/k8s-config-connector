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

package apigee

import (
	"reflect"
	"testing"

	api "google.golang.org/api/apigee/v1"
)

func TestConvertPolicy(t *testing.T) {
	in := &api.GoogleIamV1Policy{
		Version: 3,
		Etag:    "etag123",
		Bindings: []*api.GoogleIamV1Binding{
			{
				Role:    "roles/viewer",
				Members: []string{"user:foo@example.com"},
				Condition: &api.GoogleTypeExpr{
					Title:       "title",
					Description: "desc",
					Expression:  "expr",
				},
			},
		},
	}

	pb := convertPolicyToPB(in)
	if pb.Version != 3 {
		t.Errorf("version mismatch: got %d, want 3", pb.Version)
	}
	if string(pb.Etag) != "etag123" {
		t.Errorf("etag mismatch: got %s, want etag123", string(pb.Etag))
	}
	if len(pb.Bindings) != 1 {
		t.Fatalf("bindings length mismatch: got %d, want 1", len(pb.Bindings))
	}
	if pb.Bindings[0].Role != "roles/viewer" {
		t.Errorf("role mismatch: got %s, want roles/viewer", pb.Bindings[0].Role)
	}
	if pb.Bindings[0].Condition.Title != "title" {
		t.Errorf("condition title mismatch: got %s, want title", pb.Bindings[0].Condition.Title)
	}

	out := convertPolicyFromPB(pb)
	if !reflect.DeepEqual(in, out) {
		t.Errorf("roundtrip failed: got %+v, want %+v", out, in)
	}
}
