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
		t.Errorf("expected no diffs for description-only change, got: %v", result.Diffs)
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
		t.Errorf("expected no diffs for listKind addition, got: %v", result.Diffs)
	}
	if len(result.Notes) == 0 {
		t.Error("expected a note about listKind addition")
	}
}

func TestEquivalence_AddStatusField(t *testing.T) {
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
              newStatusField:
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
		t.Errorf("expected no diffs for status field addition, got: %v", result.Diffs)
	}
	if len(result.Notes) == 0 {
		t.Error("expected a note about status field addition")
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
	if len(result.Diffs) == 0 {
		t.Error("expected diffs for spec field addition")
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
	if len(result.Diffs) == 0 {
		t.Error("expected diffs for removed spec field")
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
	if len(result.Diffs) == 0 {
		t.Error("expected diffs for type change")
	}
}

func TestBackwardCompat_AddField(t *testing.T) {
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
	old, err := parseCRD([]byte(baseCRD))
	if err != nil {
		t.Fatal(err)
	}
	new, err := parseCRD([]byte(modified))
	if err != nil {
		t.Fatal(err)
	}

	result := compareBackwardCompatibility(old, new)
	if len(result.Diffs) != 0 {
		t.Errorf("expected no diffs for field addition, got: %v", result.Diffs)
	}
	if len(result.Notes) == 0 {
		t.Error("expected a note about added field")
	}
}

func TestBackwardCompat_RemoveField(t *testing.T) {
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

	result := compareBackwardCompatibility(old, new)
	if len(result.Diffs) == 0 {
		t.Error("expected diffs for removed field")
	}
}

func TestGitShow_InvalidRef(t *testing.T) {
	// An invalid ref should return an error, not silently report "file is new".
	_, isNew, err := gitShow("nonexistent-ref-xyz123", "compare_test.go")
	if err == nil {
		t.Error("expected an error for an invalid git ref, got nil")
	}
	if isNew {
		t.Error("an invalid git ref should return an error, not 'file is new'")
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
	if len(result.Diffs) == 0 {
		t.Error("expected diffs for type change")
	}
}
