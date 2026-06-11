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

package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// crdTemplate is a base CRD with a placeholder for the version name.
const crdTemplate = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
  labels:
    cnrm.cloud.google.com/stability-level: {{STABILITY}}
spec:
  group: example.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: {{VERSION}}
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            description: "The spec"
            properties:
              projectRef:
                type: object
                properties:
                  name:
                    type: string
                    description: "Project name"
              region:
                type: string
                description: "Region"
              count:
                type: integer
              httpKeepAliveTimeoutSec:
                type: integer
                format: int64
          status:
            type: object
            description: "The status"
            properties:
              observedGeneration:
                type: integer
                description: "Observed generation"
              conditions:
                type: array
                items:
                  type: object
                  properties:
                    type:
                      type: string
                    status:
                      type: string
              proxyId:
                type: integer
                format: int64
`

// baseCRD is a minimal CRD for testing.
const baseCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
  labels:
    cnrm.cloud.google.com/stability-level: alpha
spec:
  group: example.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            description: "The spec"
            properties:
              projectRef:
                type: object
                properties:
                  name:
                    type: string
                    description: "Project name"
              region:
                type: string
                description: "Region"
          status:
            type: object
            description: "The status"
            properties:
              observedGeneration:
                type: integer
                format: int64
                description: "Observed generation"
              conditions:
                type: array
                items:
                  type: object
                  properties:
                    type:
                      type: string
                    status:
                      type: string
`

func getCRD(version string, stability string) string {
	res := strings.ReplaceAll(crdTemplate, "{{VERSION}}", version)
	return strings.ReplaceAll(res, "{{STABILITY}}", stability)
}

// TestEquivalence_SharedBehavior tests scenarios where v1alpha1 and v1beta1 behave identically.
func TestEquivalence_SharedBehavior(t *testing.T) {
	tests := []struct {
		name     string
		modifier func(string) string
	}{
		{
			name: "Description Change (Pass)",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "The spec", "Updated spec description")
			},
		},
		{
			name: "Add listKind (Pass with Note)",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "kind: Foo", "kind: Foo\n    listKind: FooList")
			},
		},
		{
			name: "Add status.externalRef (Pass with Note)",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "observedGeneration:", "externalRef:\n                type: string\n              observedGeneration:")
			},
		},
		{
			name: "Add Empty Object (Fail/Difference)",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "projectRef:", "emptyObj:\n                type: object\n              projectRef:")
			},
		},
	}

	for _, version := range []string{"v1alpha1", "v1beta1"} {
		stability := "alpha"
		if version == "v1beta1" {
			stability = "stable"
		}
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s/%s", version, tt.name), func(t *testing.T) {
				oldStr := getCRD(version, stability)
				newStr := tt.modifier(oldStr)

				old, err := parseCRD([]byte(oldStr))
				if err != nil {
					t.Fatalf("failed to parse old CRD: %v", err)
				}
				newCRD, err := parseCRD([]byte(newStr))
				if err != nil {
					t.Fatalf("failed to parse new CRD: %v\nYAML:\n%s", err, newStr)
				}
				result := compareEquivalence(old, newCRD)

				if strings.Contains(tt.name, "Empty Object") {
					// Should fail for both
					if len(result.Diffs) == 0 {
						t.Errorf("expected blocking diff for empty object, but got none")
					}
				} else {
					// Should pass for both
					if len(result.Diffs) != 0 {
						t.Errorf("expected no diffs, but got %v", result.Diffs)
					}
				}
			})
		}
	}
}

// TestEquivalence_DivergentBehavior illustrates differences between v1alpha1 (relaxed) and v1beta1 (strict).
func TestEquivalence_DivergentBehavior(t *testing.T) {
	tests := []struct {
		name     string
		modifier func(string) string
	}{
		{
			name: "Add Spec Field",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "region:", "region:\n                type: string\n              newField:\n                type: string")
			},
		},
		{
			name: "Remove Spec Field",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "region:", "removedField:")
			},
		},
		{
			name: "Type Change (string -> integer)",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "type: string\n                description: \"Region\"", "type: integer")
			},
		},
		// Removal of certain Status fields should also be allowed in v1alpha1 but not allowed in v1beta1.
		// E.g. observedGeneration, conditions should not be removed in both versions, but other fields can be removed in v1alpha1.
	}

	for _, tt := range tests {
		// v1alpha1: Expected to be EQUIVALENT (Notes instead of Diffs)
		t.Run(fmt.Sprintf("v1alpha1/%s", tt.name), func(t *testing.T) {
			oldStr := getCRD("v1alpha1", "alpha")
			newStr := tt.modifier(oldStr)
			old, err := parseCRD([]byte(oldStr))
			if err != nil {
				t.Fatalf("failed to parse old: %v", err)
			}
			newCRD, err := parseCRD([]byte(newStr))
			if err != nil {
				t.Fatalf("failed to parse new: %v", err)
			}

			result := compareEquivalence(old, newCRD)
			if len(result.Diffs) != 0 {
				t.Errorf("v1alpha1 should allow changes as notes, but got diffs: %v", result.Diffs)
			}
			if len(result.Notes) == 0 {
				t.Error("v1alpha1 should have generated informational notes for the change")
			}
		})

		// v1beta1: Expected to be NOT EQUIVALENT (Blocking Diffs)
		t.Run(fmt.Sprintf("v1beta1/%s", tt.name), func(t *testing.T) {
			oldStr := getCRD("v1beta1", "stable")
			newStr := tt.modifier(oldStr)
			old, err := parseCRD([]byte(oldStr))
			if err != nil {
				t.Fatalf("failed to parse old: %v", err)
			}
			newCRD, err := parseCRD([]byte(newStr))
			if err != nil {
				t.Fatalf("failed to parse new: %v", err)
			}

			result := compareEquivalence(old, newCRD)
			if len(result.Diffs) == 0 {
				t.Errorf("v1beta1 should block %s, but got no diffs", tt.name)
			}
		})
	}
}

// TestEquivalence_RequiredFieldBehavior tests the 'required' constraint detection.
func TestEquivalence_RequiredFieldBehavior(t *testing.T) {
	modifier := func(in string) string {
		// Insert 'required' block right before 'properties' inside 'spec'
		return strings.ReplaceAll(in, "description: \"The spec\"", "description: \"The spec\"\n            required:\n            - region")
	}

	// v1alpha1: Note only
	t.Run("v1alpha1/Add Required", func(t *testing.T) {
		oldStr := getCRD("v1alpha1", "alpha")
		newStr := modifier(oldStr)
		old, err := parseCRD([]byte(oldStr))
		if err != nil {
			t.Fatalf("failed to parse old: %v", err)
		}
		newCRD, err := parseCRD([]byte(newStr))
		if err != nil {
			t.Fatalf("failed to parse new: %v\nYAML:\n%s", err, newStr)
		}
		result := compareEquivalence(old, newCRD)
		if len(result.Diffs) != 0 {
			t.Errorf("v1alpha1 should allow required field addition as note, got diffs: %v", result.Diffs)
		}
	})

	// v1beta1: Blocking Diff
	t.Run("v1beta1/Add Required", func(t *testing.T) {
		oldStr := getCRD("v1beta1", "stable")
		newStr := modifier(oldStr)
		old, err := parseCRD([]byte(oldStr))
		if err != nil {
			t.Fatalf("failed to parse old: %v", err)
		}
		newCRD, err := parseCRD([]byte(newStr))
		if err != nil {
			t.Fatalf("failed to parse new: %v\nYAML:\n%s", err, newStr)
		}
		result := compareEquivalence(old, newCRD)
		if len(result.Diffs) == 0 {
			t.Error("v1beta1 should block adding required fields")
		}
	})
}

// TestEquivalence_StabilityLevelBehavior tests the cnrm.cloud.google.com/stability-level rules.
func TestEquivalence_StabilityLevelBehavior(t *testing.T) {
	// 1. v1alpha1 must be alpha
	t.Run("v1alpha1/MustBeAlpha", func(t *testing.T) {
		oldStr := getCRD("v1alpha1", "alpha")
		newStr := strings.ReplaceAll(oldStr, "stability-level: alpha", "stability-level: stable")
		old, err := parseCRD([]byte(oldStr))
		if err != nil {
			t.Fatalf("failed to parse old: %v", err)
		}
		newCRD, err := parseCRD([]byte(newStr))
		if err != nil {
			t.Fatalf("failed to parse new: %v", err)
		}
		result := compareEquivalence(old, newCRD)
		if len(result.Diffs) == 0 {
			t.Error("should have blocked stable label on v1alpha1-only resource")
		}
	})

	// 2. Promotion alpha -> stable allowed if v1beta1 is added
	t.Run("Promotion/AlphaToStableWithV1Beta1", func(t *testing.T) {
		oldStr := getCRD("v1alpha1", "alpha")
		newStr := strings.ReplaceAll(oldStr, "v1alpha1", "v1beta1")
		newStr = strings.ReplaceAll(newStr, "stability-level: alpha", "stability-level: stable")
		// Add v1alpha1 back so it has both
		newStr = strings.ReplaceAll(newStr, "versions:", "versions:\n  - name: v1alpha1\n    served: true\n    storage: true\n    schema:\n      openAPIV3Schema:\n        type: object")
		
		old, _ := parseCRD([]byte(oldStr))
		newCRD, _ := parseCRD([]byte(newStr))
		result := compareEquivalence(old, newCRD)
		
		foundPromotionNote := false
		for _, n := range result.Notes {
			if strings.Contains(n, "promoted") {
				foundPromotionNote = true
				break
			}
		}
		if !foundPromotionNote {
			t.Errorf("expected promotion note, but got none. Diffs: %v, Notes: %v", result.Diffs, result.Notes)
		}
	})
}

// TestEquivalence_IntegerConversions tests safe vs unsafe integer type changes.
func TestEquivalence_IntegerConversions(t *testing.T) {
	tests := []struct {
		name     string
		modifier func(string) string
		v1alpha1Result CompareResult
		v1beta1Result CompareResult
	}{
		{
			name: "Shared behavior: integer -> int64 in STATUS (Allowed with Note)",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "observedGeneration:\n                type: integer", "observedGeneration:\n                type: integer\n                format: int64")
			},
			v1alpha1Result: CompareResult{
				Diffs: nil,
				Notes: []string{"[v1alpha1] field type changed: status.observedGeneration (integer -> int64) (allowed)"},
			},
			v1beta1Result: CompareResult{
				Diffs: nil,
				Notes: []string{"[v1beta1] field type changed: status.observedGeneration (integer -> int64) (allowed)"},
			},
		},
		{
			name: "Shared behavior: integer -> int32 in STATUS (Allowed with Note)",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "observedGeneration:\n                type: integer", "observedGeneration:\n                type: integer\n                format: int32")
			},
			v1alpha1Result: CompareResult{
				Diffs: nil,
				Notes: []string{"[v1alpha1] field type changed: status.observedGeneration (integer -> int32) (allowed)"},
			},
			v1beta1Result: CompareResult{
				Diffs: nil,
				Notes: []string{"[v1beta1] field type changed: status.observedGeneration (integer -> int32) (allowed)"},
			},
		},
		{
			name: "Diverged behavior: int64 -> int32 in STATUS (v1beta1: Blocked; v1alpha1: Allowed with Note)",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "proxyId:\n                type: integer\n                format: int64", "proxyId:\n                type: integer\n                format: int32")
			},
			v1alpha1Result: CompareResult{
				Diffs: nil,
				Notes: []string{"[v1alpha1] field type changed: status.proxyId (int64 -> int32) (allowed for alpha)"},
			},
			v1beta1Result: CompareResult{
				Diffs: []string{"[v1beta1] field type changed: status.proxyId (int64 -> int32)"},
				Notes: nil,
			},
		},
		{
			name: "Diverged behavior: integer -> int64 in SPEC (v1beta1: Blocked; v1alpha1: Allowed with Note)\"",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "count:\n                type: integer", "count:\n                type: integer\n                format: int64")
			},
			v1alpha1Result: CompareResult{
				Diffs: nil,
				Notes: []string{"[v1alpha1] field type changed: spec.count (integer -> int64) (allowed for alpha)"},
			},
			v1beta1Result: CompareResult{
				Diffs: []string{"[v1beta1] field type changed: spec.count (integer -> int64)"},
				Notes: nil,
			},
		},
		{
			name: "Diverged behavior: integer -> int32 in SPEC (v1beta1: Blocked; v1alpha1: Allowed with Note)\"",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "count:\n                type: integer", "count:\n                type: integer\n                format: int32")
			},
			v1alpha1Result: CompareResult{
				Diffs: nil,
				Notes: []string{"[v1alpha1] field type changed: spec.count (integer -> int32) (allowed for alpha)"},
			},
			v1beta1Result: CompareResult{
				Diffs: []string{"[v1beta1] field type changed: spec.count (integer -> int32)"},
				Notes: nil,
			},
		},
		{
			name: "Diverged behavior: int64 -> int32 in SPEC (v1beta1: Blocked; v1alpha1: Allowed with Note)\"",
			modifier: func(in string) string {
				return strings.ReplaceAll(in, "httpKeepAliveTimeoutSec:\n                type: integer\n                format: int64", "httpKeepAliveTimeoutSec:\n                type: integer\n                format: int32")
			},
			v1alpha1Result: CompareResult{
				Diffs: nil,
				Notes: []string{"[v1alpha1] field type changed: spec.httpKeepAliveTimeoutSec (int64 -> int32) (allowed for alpha)"},
			},
			v1beta1Result: CompareResult{
				Diffs: []string{"[v1beta1] field type changed: spec.httpKeepAliveTimeoutSec (int64 -> int32)"},
				Notes: nil,
			},
		},
	}

	for _, tt := range tests {
		// v1alpha1
		t.Run(fmt.Sprintf("v1alpha1/%s", tt.name), func(t *testing.T) {
			oldStr := getCRD("v1alpha1", "alpha")
			newStr := tt.modifier(oldStr)
			old, err := parseCRD([]byte(oldStr))
			if err != nil {
				t.Fatalf("failed to parse old: %v", err)
			}
			newCRD, err := parseCRD([]byte(newStr))
			if err != nil {
				t.Fatalf("failed to parse new: %v", err)
			}

			result := compareEquivalence(old, newCRD)
			if diff := cmp.Diff(tt.v1alpha1Result, result); diff != "" {
				t.Errorf("compare result mismatch (-want +got):\n%s", diff)
			}
		})

		// v1beta1
		t.Run(fmt.Sprintf("v1beta1/%s", tt.name), func(t *testing.T) {
			oldStr := getCRD("v1beta1", "stable")
			newStr := tt.modifier(oldStr)
			old, err := parseCRD([]byte(oldStr))
			if err != nil {
				t.Fatalf("failed to parse old: %v", err)
			}
			newCRD, err := parseCRD([]byte(newStr))
			if err != nil {
				t.Fatalf("failed to parse new: %v", err)
			}

			result := compareEquivalence(old, newCRD)
			if diff := cmp.Diff(tt.v1beta1Result, result); diff != "" {
				t.Errorf("compare result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func matchGotAndExpected(t *testing.T, got CompareResult, expectedNotes, expectedDiffs []string) {

}

func TestBackwardCompat_AddField(t *testing.T) {
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
  labels:
    cnrm.cloud.google.com/stability-level: alpha
spec:
  group: example.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              projectRef:
                type: object
                properties:
                  name:
                    type: string
              region:
                type: string
              newField:
                type: string
          status:
            type: object
            properties:
              observedGeneration:
                type: integer
                format: int64
              conditions:
                type: array
                items:
                  type: object
                  properties:
                    type:
                      type: string
                    status:
                      type: string
`
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareBackwardCompatibility(old, new)
	if len(result.Diffs) != 0 {
		t.Fatalf("expected 0 diffs for field addition, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 1 {
		t.Fatalf("expected 1 note about added field, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestBackwardCompat_TypeChange(t *testing.T) {
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
  labels:
    cnrm.cloud.google.com/stability-level: alpha
spec:
  group: example.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              projectRef:
                type: object
                properties:
                  name:
                    type: integer
              region:
                type: string
          status:
            type: object
            properties:
              observedGeneration:
                type: integer
                format: int64
              conditions:
                type: array
                items:
                  type: object
                  properties:
                    type:
                      type: string
                    status:
                      type: string
`
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareBackwardCompatibility(old, new)
	if len(result.Diffs) != 1 {
		t.Fatalf("expected 1 diff for type change, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 0 {
		t.Fatalf("expected 0 notes for type change, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestBackwardCompat_IntegerTypeChange(t *testing.T) {
	oldCRDStr := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
  labels:
    cnrm.cloud.google.com/stability-level: alpha
spec:
  group: example.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1beta1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              httpKeepAliveTimeoutSec:
                type: integer
              count:
                type: integer
          status:
            type: object
            properties:
              observedGeneration:
                type: integer
              proxyId:
                type: integer
`

	newCRDStr := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
  labels:
    cnrm.cloud.google.com/stability-level: alpha
spec:
  group: example.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1beta1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              httpKeepAliveTimeoutSec:
                type: integer
                format: int32
              count:
                type: integer
                format: int64
          status:
            type: object
            properties:
              observedGeneration:
                type: integer
                format: int64
              proxyId:
                type: integer
                format: int64
`

	old, err := parseCRD([]byte(oldCRDStr))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(newCRDStr))
	if err != nil {
		t.Fatal(err)
	}

	result := compareBackwardCompatibility(old, new)

	expectedBlocked := []string{
		"[v1beta1] field type changed: spec.count (integer -> int64)",
		"[v1beta1] field type changed: spec.httpKeepAliveTimeoutSec (integer -> int32)",
	}

	if len(result.Diffs) != len(expectedBlocked) {
		t.Fatalf("expected %d diffs, got %d: %v", len(expectedBlocked), len(result.Diffs), result.Diffs)
	}
	for i, d := range result.Diffs {
		if d != expectedBlocked[i] {
			t.Errorf("diff mismatch at %d: expected %q, got %q", i, expectedBlocked[i], d)
		}
	}

	expectedNotes := []string{
		"[v1beta1] field type changed: status.observedGeneration (integer -> int64) (allowed)",
		"[v1beta1] field type changed: status.proxyId (integer -> int64) (allowed)",
	}

	if len(result.Notes) != len(expectedNotes) {
		t.Fatalf("expected %d notes, got %d: %v", len(expectedNotes), len(result.Notes), result.Notes)
	}
	for i, n := range result.Notes {
		if n != expectedNotes[i] {
			t.Errorf("note mismatch at %d: expected %q, got %q", i, expectedNotes[i], n)
		}
	}
}

func TestBackwardCompat_ValidationSkipped(t *testing.T) {
	// Note: we skip validations in the backward compatibility check for now,
	// as changes to validations may be determined on a case-by-case basis.
	oldYaml := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
  labels:
    cnrm.cloud.google.com/stability-level: alpha
spec:
  group: example.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              region:
                type: string
                x-kubernetes-validations:
                - rule: 'self == "us-central1"'
`
	newYaml := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
  labels:
    cnrm.cloud.google.com/stability-level: alpha
spec:
  group: example.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              region:
                type: string
                x-kubernetes-validations:
                - rule: 'self == "us-west1"'
`
	old, err := parseCRD([]byte(oldYaml))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(newYaml))
	if err != nil {
		t.Fatal(err)
	}

	result := compareBackwardCompatibility(old, new)
	if len(result.Diffs) != 0 {
		t.Errorf("expected no diffs for validation rule change as they are skipped for backward compatibility, got: %v", result.Diffs)
	}
}

func TestGitShow_InvalidRef(t *testing.T) {
	// An invalid ref should return an error, not silently report "file is new".
	_, isNew, err := gitShow("nonexistent-ref-xyz123", "compare_test.go")
	if err == nil {
		t.Fatal("expected an error for an invalid git ref, got nil")
	}
	if isNew {
		t.Fatal("an invalid git ref should return an error, not 'file is new'")
	}
}
