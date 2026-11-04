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

package logging

import (
	"reflect"
	"testing"

	"cloud.google.com/go/iam/apiv1/iampb"
	loggingapi "google.golang.org/api/logging/v2"
	expr "google.golang.org/genproto/googleapis/type/expr"
)

func TestConvertPolicyAPIToIAMPB(t *testing.T) {
	// Arrange
	input := &loggingapi.Policy{
		Version: 3,
		Etag:    "BwW12345678=",
		Bindings: []*loggingapi.Binding{
			{
				Role: "roles/logging.viewAccessor",
				Members: []string{
					"user:alice@example.com",
					"serviceAccount:sa@project.iam.gserviceaccount.com",
				},
				Condition: &loggingapi.Expr{
					Title:       "expires_end_of_year",
					Description: "Expires at end of 2026",
					Expression:  "request.time < timestamp('2027-01-01T00:00:00Z')",
				},
			},
		},
	}

	// Act
	got := convertPolicyAPIToIAMPB(input)

	// Assert
	if got == nil {
		t.Fatalf("expected non-nil output")
	}
	if got.Version != 3 {
		t.Errorf("got.Version = %d, want 3", got.Version)
	}
	if string(got.Etag) != input.Etag {
		t.Errorf("got.Etag = %q, want %q", string(got.Etag), input.Etag)
	}
	if len(got.Bindings) != len(input.Bindings) {
		t.Fatalf("got %d bindings, want %d", len(got.Bindings), len(input.Bindings))
	}
	if got.Bindings[0].Role != input.Bindings[0].Role {
		t.Errorf("got.Role = %q, want %q", got.Bindings[0].Role, input.Bindings[0].Role)
	}
	if !reflect.DeepEqual(got.Bindings[0].Members, input.Bindings[0].Members) {
		t.Errorf("got.Members = %v, want %v", got.Bindings[0].Members, input.Bindings[0].Members)
	}
	if got.Bindings[0].Condition == nil || got.Bindings[0].Condition.Title != input.Bindings[0].Condition.Title {
		t.Errorf("got.Condition = %+v, want %+v", got.Bindings[0].Condition, input.Bindings[0].Condition)
	}
}

func TestConvertPolicyIAMPBToAPI(t *testing.T) {
	// Arrange
	input := &iampb.Policy{
		Version: 3,
		Etag:    []byte("BwW12345678="),
		Bindings: []*iampb.Binding{
			{
				Role: "roles/logging.viewAccessor",
				Members: []string{
					"user:bob@example.com",
				},
				Condition: &expr.Expr{
					Title:       "conditional_access",
					Description: "Testing condition",
					Expression:  "true",
				},
			},
		},
	}

	// Act
	got := convertPolicyIAMPBToAPI(input)

	// Assert
	if got == nil {
		t.Fatalf("expected non-nil output")
	}
	if got.Version != 3 {
		t.Errorf("got.Version = %d, want 3", got.Version)
	}
	if got.Etag != string(input.Etag) {
		t.Errorf("got.Etag = %q, want %q", got.Etag, string(input.Etag))
	}
	if len(got.Bindings) != len(input.Bindings) {
		t.Fatalf("got %d bindings, want %d", len(got.Bindings), len(input.Bindings))
	}
	if got.Bindings[0].Role != input.Bindings[0].Role {
		t.Errorf("got.Role = %q, want %q", got.Bindings[0].Role, input.Bindings[0].Role)
	}
	if !reflect.DeepEqual(got.Bindings[0].Members, input.Bindings[0].Members) {
		t.Errorf("got.Members = %v, want %v", got.Bindings[0].Members, input.Bindings[0].Members)
	}
	if got.Bindings[0].Condition == nil || got.Bindings[0].Condition.Title != input.Bindings[0].Condition.Title {
		t.Errorf("got.Condition = %+v, want %+v", got.Bindings[0].Condition, input.Bindings[0].Condition)
	}
	if got.Bindings[0].Condition.Expression != input.Bindings[0].Condition.Expression {
		t.Errorf("got.Condition.Expression = %q, want %q", got.Bindings[0].Condition.Expression, input.Bindings[0].Condition.Expression)
	}
}

func TestConvertPolicyIAMPBToAPI_UpgradesVersionForConditions(t *testing.T) {
	// Arrange
	input := &iampb.Policy{
		Version: 1,
		Bindings: []*iampb.Binding{
			{
				Role:    "roles/logging.viewAccessor",
				Members: []string{"user:bob@example.com"},
				Condition: &expr.Expr{
					Title:      "conditional",
					Expression: "true",
				},
			},
		},
	}

	// Act
	got := convertPolicyIAMPBToAPI(input)

	// Assert
	if got == nil {
		t.Fatalf("expected non-nil output")
	}
	if got.Version != 3 {
		t.Errorf("got.Version = %d, want 3 when conditions exist", got.Version)
	}
}

func TestConvertPolicyAPIToIAMPB_Nil(t *testing.T) {
	// Arrange & Act
	got := convertPolicyAPIToIAMPB(nil)

	// Assert
	if got != nil {
		t.Errorf("expected nil output for nil input, got %+v", got)
	}
}

func TestConvertPolicyIAMPBToAPI_Nil(t *testing.T) {
	// Arrange & Act
	got := convertPolicyIAMPBToAPI(nil)

	// Assert
	if got != nil {
		t.Errorf("expected nil output for nil input, got %+v", got)
	}
}

func TestConvertPolicyIAMPBToAPI_EmptyBindings(t *testing.T) {
	// Arrange
	input := &iampb.Policy{
		Version:  3,
		Etag:     []byte("BwW12345678="),
		Bindings: nil,
	}

	// Act
	got := convertPolicyIAMPBToAPI(input)

	// Assert
	if got == nil {
		t.Fatalf("expected non-nil output")
	}
	if len(got.Bindings) != 0 {
		t.Errorf("expected 0 bindings, got %d", len(got.Bindings))
	}
	foundForceSend := false
	for _, f := range got.ForceSendFields {
		if f == "Bindings" {
			foundForceSend = true
			break
		}
	}
	if !foundForceSend {
		t.Errorf("expected ForceSendFields to contain 'Bindings', got %v", got.ForceSendFields)
	}
}
