// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package bigtable

import (
	"testing"
)

func TestFamilyHash_NoPanic(t *testing.T) {
	testCases := []struct {
		name  string
		input interface{}
	}{
		{
			name:  "nil input",
			input: nil,
		},
		{
			name:  "non-map input",
			input: "invalid",
		},
		{
			name: "valid family without type",
			input: map[string]interface{}{
				"family": "cf1",
			},
		},
		{
			name: "valid family with shorthand type",
			input: map[string]interface{}{
				"family": "cf1",
				"type":   "intsum",
			},
		},
		{
			name: "invalid type string",
			input: map[string]interface{}{
				"family": "cf1",
				"type":   "{not-valid-json",
			},
		},
		{
			name: "non-string type",
			input: map[string]interface{}{
				"family": "cf1",
				"type":   12345,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("familyHash panicked on %s: %v", tc.name, r)
				}
			}()
			_ = familyHash(tc.input)
		})
	}
}

func TestTypeDiffFunc_NoPanic(t *testing.T) {
	testCases := []struct {
		name     string
		oldValue string
		newValue string
		expected bool
	}{
		{
			name:     "both empty",
			oldValue: "",
			newValue: "",
			expected: true,
		},
		{
			name:     "invalid json old",
			oldValue: "{invalid",
			newValue: "intsum",
			expected: false,
		},
		{
			name:     "invalid json new",
			oldValue: "intsum",
			newValue: "{invalid",
			expected: false,
		},
		{
			name:     "same shorthand type",
			oldValue: "intsum",
			newValue: "intsum",
			expected: true,
		},
		{
			name:     "different shorthand type",
			oldValue: "intsum",
			newValue: "intmax",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("typeDiffFunc panicked on %s: %v", tc.name, r)
				}
			}()
			got := typeDiffFunc("type", tc.oldValue, tc.newValue, nil)
			if got != tc.expected {
				t.Errorf("typeDiffFunc(%s, %s) = %v, want %v", tc.oldValue, tc.newValue, got, tc.expected)
			}
		})
	}
}
