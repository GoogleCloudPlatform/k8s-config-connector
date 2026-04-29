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
	"slices"
	"strings"
	"testing"
)

// baseCRD is a minimal CRD for testing.
const baseCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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

func TestEquivalence_DescriptionChange(t *testing.T) {
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
            description: "The spec - updated description"
            properties:
              projectRef:
                type: object
                properties:
                  name:
                    type: string
                    description: "Project name - updated"
              region:
                type: string
                description: "Region - updated"
          status:
            type: object
            description: "The status - updated"
            properties:
              observedGeneration:
                type: integer
                format: int64
                description: "Observed generation - updated"
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
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 0 {
		t.Fatalf("expected 0 diffs for description-only change, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 0 {
		t.Fatalf("expected 0 notes for description-only change, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestEquivalence_AddListKind(t *testing.T) {
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
spec:
  group: example.com
  names:
    kind: Foo
    listKind: FooList
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
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 0 {
		t.Fatalf("expected 0 diffs for listKind addition, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 1 {
		t.Fatalf("expected 1 note about listKind addition, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestEquivalence_AllowedExternalRef(t *testing.T) {
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
              externalRef:
                type: string
`
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 0 {
		t.Fatalf("expected 0 diffs for allowed status field addition, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 1 {
		t.Fatalf("expected 1 note for allowed status field addition, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestEquivalence_AddSpecField(t *testing.T) {
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
              newSpecField:
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
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 1 {
		t.Fatalf("expected 1 diff for spec field addition, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 0 {
		t.Fatalf("expected 0 notes for spec field addition, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestEquivalence_RemoveField(t *testing.T) {
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 1 {
		t.Fatalf("expected 1 diff for removed spec field, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 0 {
		t.Fatalf("expected 0 notes for removed spec field, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestEquivalence_RemoveStatusField(t *testing.T) {
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
          status:
            type: object
            properties:
              observedGeneration:
                type: integer
                format: int64
`
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 1 {
		t.Fatalf("expected 1 diff for removed status field, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 0 {
		t.Fatalf("expected 0 notes for removed status field, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestEquivalence_TypeChange(t *testing.T) {
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
                type: integer
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
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 1 {
		t.Fatalf("expected 1 diff for type change, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 0 {
		t.Fatalf("expected 0 notes for type change, got %d: %v", len(result.Notes), result.Notes)
	}
}

func TestEquivalence_IntegerTypeChange(t *testing.T) {
	oldCRDStr := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
              nodePort:
                type: integer
`

	newCRDStr := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
              nodePort:
                type: integer
                format: int32
`

	old, err := parseCRD([]byte(oldCRDStr))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(newCRDStr))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)

	expectedBlocked := []string{
		"[v1beta1] field type changed: spec.count (integer -> int64)",
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
		"[v1beta1] field type changed: spec.httpKeepAliveTimeoutSec (integer -> int32) (allowed)",
		"[v1beta1] field type changed: status.nodePort (integer -> int32) (allowed)",
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

const testOldCRDData = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
              foo:
                type: string
          status:
            type: object
            properties:
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

func TestEquivalence_NewStatusField(t *testing.T) {
	newCRDData := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
              foo:
                type: string
          status:
            type: object
            properties:
              conditions:
                type: array
                items:
                  type: object
                  properties:
                    type:
                      type: string
                    status:
                      type: string
                    newField:
                      type: string
              externalRef:
                type: object
                properties:
                  subField:
                    type: string
              externalRefFoo:
                type: string
              externalRefName:
                type: string
              observedState:
                type: object
                properties:
                  bar:
                    type: string
              observedStateFoo:
                type: string
              unallowedObject:
                type: object
                properties:
                  child:
                    type: string
`

	old, err := parseCRD([]byte(testOldCRDData))
	if err != nil {
		t.Fatalf("parseCRD old: %v", err)
	}
	new, err := parseCRD([]byte(newCRDData))
	if err != nil {
		t.Fatalf("parseCRD new: %v", err)
	}

	result := compareEquivalence(old, new)

	t.Run("Check diffs", func(t *testing.T) {
		// Diffs expected for:
		// 1. status.conditions[].newField
		// 2. status.externalRefFoo
		// 3. status.externalRefName
		// 4. status.externalRef.subField (externalRef is string pointer, no subfields allowed)
		// 5. status.observedStateFoo
		// 6. status.unallowedObject
		if len(result.Diffs) != 6 {
			t.Fatalf("Expected 6 diffs, but got %d: %v", len(result.Diffs), result.Diffs)
		}
		if !slices.ContainsFunc(result.Diffs, func(diff string) bool { return strings.Contains(diff, "status.conditions[].newField") }) {
			t.Errorf("Expected diff for status.conditions[].newField, but it was not found. Result diffs: %v", result.Diffs)
		}
		if !slices.ContainsFunc(result.Diffs, func(diff string) bool { return strings.Contains(diff, "status.externalRefFoo") }) {
			t.Errorf("Expected diff for status.externalRefFoo, but it was not found. Result diffs: %v", result.Diffs)
		}
		if !slices.ContainsFunc(result.Diffs, func(diff string) bool { return strings.Contains(diff, "status.externalRefName") }) {
			t.Errorf("Expected diff for status.externalRefName, but it was not found. Result diffs: %v", result.Diffs)
		}
		if !slices.ContainsFunc(result.Diffs, func(diff string) bool { return strings.Contains(diff, "status.externalRef.subField") }) {
			t.Errorf("Expected diff for status.externalRef.subField, but it was not found. Result diffs: %v", result.Diffs)
		}
		if !slices.ContainsFunc(result.Diffs, func(diff string) bool { return strings.Contains(diff, "status.unallowedObject") }) {
			t.Errorf("Expected diff for status.unallowedObject, but it was not found. Result diffs: %v", result.Diffs)
		}
		if !slices.ContainsFunc(result.Diffs, func(diff string) bool { return strings.Contains(diff, "status.observedStateFoo") }) {
			t.Errorf("Expected diff for status.observedStateFoo, but it was not found. Result diffs: %v", result.Diffs)
		}
	})

	t.Run("Check notes", func(t *testing.T) {
		if len(result.Notes) != 3 {
			t.Fatalf("Expected 3 notes (externalRef, observedState, observedState.bar), but got %d: %v", len(result.Notes), result.Notes)
		}
		if !slices.ContainsFunc(result.Notes, func(note string) bool { return strings.Contains(note, " status.externalRef ") }) {
			t.Errorf("Expected note for status.externalRef, but it was not found. Result notes: %v", result.Notes)
		}
		if !slices.ContainsFunc(result.Notes, func(note string) bool { return strings.Contains(note, " status.observedState ") }) {
			t.Errorf("Expected note for status.observedState (parent), but it was not found. Result notes: %v", result.Notes)
		}
		if !slices.ContainsFunc(result.Notes, func(note string) bool { return strings.Contains(note, "status.observedState.bar") }) {
			t.Errorf("Expected note for status.observedState.bar, but it was not found. Result notes: %v", result.Notes)
		}
	})
}

func TestEquivalence_EmptyObservedState(t *testing.T) {
	newCRDData := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
              foo:
                type: string
          status:
            type: object
            properties:
              conditions:
                type: array
                items:
                  type: object
                  properties:
                    type:
                      type: string
                    status:
                      type: string
              observedState:
                type: object
`

	old, err := parseCRD([]byte(testOldCRDData))
	if err != nil {
		t.Fatalf("parseCRD old: %v", err)
	}
	new, err := parseCRD([]byte(newCRDData))
	if err != nil {
		t.Fatalf("parseCRD new: %v", err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 0 {
		t.Fatalf("expected 0 diffs, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 1 {
		t.Fatalf("expected 1 note, got %d: %v", len(result.Notes), result.Notes)
	}
	if !slices.ContainsFunc(result.Notes, func(note string) bool { return strings.Contains(note, " status.observedState ") }) {
		t.Errorf("Expected note for status.observedState, but it was not found. Result notes: %v", result.Notes)
	}
}

func TestEquivalence_DisallowedStatusFieldTypeChange(t *testing.T) {
	oldCRDData := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
          status:
            type: object
            properties:
              externalRef:
                type: string
`
	newCRDData := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
          status:
            type: object
            properties:
              externalRef:
                type: integer
`
	old, err := parseCRD([]byte(oldCRDData))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(newCRDData))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 1 {
		t.Fatalf("expected 1 diff for type change, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if result.Diffs[0] != "[v1alpha1] field type changed: status.externalRef (string -> integer)" {
		t.Errorf("unexpected diff: %q", result.Diffs[0])
	}
}

func TestEquivalence_ChangeListKind(t *testing.T) {
	oldCRDData := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
spec:
  group: example.com
  names:
    kind: Foo
    listKind: OldList
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
`
	newCRDData := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
spec:
  group: example.com
  names:
    kind: Foo
    listKind: NewList
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
`
	old, err := parseCRD([]byte(oldCRDData))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(newCRDData))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) != 1 {
		t.Fatalf("expected 1 diff for listKind change, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if result.Diffs[0] != `spec.names.listKind changed: "OldList" -> "NewList"` {
		t.Errorf("unexpected diff: %q", result.Diffs[0])
	}
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

func TestBackwardCompat_RemoveField(t *testing.T) {
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
		t.Fatalf("expected 1 diff for removed field, got %d: %v", len(result.Diffs), result.Diffs)
	}
	if len(result.Notes) != 0 {
		t.Fatalf("expected 0 notes for removed field, got %d: %v", len(result.Notes), result.Notes)
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
		"[v1beta1] field type changed: spec.httpKeepAliveTimeoutSec (integer -> int32) (allowed)",
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

func TestEquivalence_AddValidation(t *testing.T) {
	modified := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
                x-kubernetes-validations:
                - rule: 'self == "us-central1"'
                  message: "only us-central1 is allowed"
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
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) == 0 {
		t.Error("expected diffs for x-kubernetes-validations addition, but got none")
	}
}

func TestEquivalence_ChangeValidation(t *testing.T) {
	oldYaml := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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

	result := compareEquivalence(old, new)
	if len(result.Diffs) == 0 {
		t.Error("expected diffs for validation rule change")
	}
}

func TestEquivalence_ChangeValidationMessage(t *testing.T) {
	oldYaml := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
                  message: "old message"
`
	newYaml := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.example.com
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
                  message: "new message"
`
	old, err := parseCRD([]byte(oldYaml))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(newYaml))
	if err != nil {
		t.Fatal(err)
	}

	result := compareEquivalence(old, new)
	if len(result.Diffs) == 0 {
		t.Error("expected diffs for validation message change")
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
