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

package dcl_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"google.golang.org/api/googleapi"
)

func TestIsMethodNotFoundError(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "dclunstruct.ErrNoSuchMethod",
			err:      dclunstruct.ErrNoSuchMethod,
			expected: true,
		},
		{
			name:     "wrapped dclunstruct.ErrNoSuchMethod",
			err:      fmt.Errorf("wrapped: %w", dclunstruct.ErrNoSuchMethod),
			expected: true,
		},
		{
			name: "googleapi 404 Method not found",
			err: &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Method not found.",
			},
			expected: true,
		},
		{
			name: "regular 404 resource not found",
			err: &googleapi.Error{
				Code:    http.StatusNotFound,
				Message: "Resource not found",
			},
			expected: false,
		},
		{
			name:     "generic string error",
			err:      errors.New("generic error"),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := dcl.IsMethodNotFoundError(tc.err); got != tc.expected {
				t.Errorf("IsMethodNotFoundError() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestIsNoSuchMethodError(t *testing.T) {
	if !dcl.IsNoSuchMethodError(dclunstruct.ErrNoSuchMethod) {
		t.Errorf("expected true for dclunstruct.ErrNoSuchMethod")
	}
	if !dcl.IsNoSuchMethodError(fmt.Errorf("wrapped: %w", dclunstruct.ErrNoSuchMethod)) {
		t.Errorf("expected true for wrapped dclunstruct.ErrNoSuchMethod")
	}
	if dcl.IsNoSuchMethodError(errors.New("other error")) {
		t.Errorf("expected false for other error")
	}
	if dcl.IsNoSuchMethodError(nil) {
		t.Errorf("expected false for nil")
	}
}
