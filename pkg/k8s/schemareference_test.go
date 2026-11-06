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

package k8s_test

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apimachineryschema "k8s.io/apimachinery/pkg/runtime/schema"
)

func newTestCRD(group, version, kind string, description string) *apiextensions.CustomResourceDefinition {
	return &apiextensions.CustomResourceDefinition{
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group: group,
			Names: apiextensions.CustomResourceDefinitionNames{
				Kind: kind,
			},
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name: version,
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Type:        "object",
							Description: description,
						},
					},
				},
			},
		},
	}
}

func TestNewSchemaReference(t *testing.T) {
	// Arrange
	group := "foo.cnrm.cloud.google.com"
	version := "v1beta1"
	kind := "FooResource"
	description := "original description"
	crd := newTestCRD(group, version, kind, description)

	// Act
	schemaRef := k8s.NewSchemaReference(crd)

	// Assert
	expectedGVK := apimachineryschema.GroupVersionKind{
		Group:   group,
		Version: version,
		Kind:    kind,
	}
	if schemaRef.GVK != expectedGVK {
		t.Fatalf("expected GVK %v, got %v", expectedGVK, schemaRef.GVK)
	}
	if schemaRef.CRD != crd {
		t.Fatalf("expected CRD %v, got %v", crd, schemaRef.CRD)
	}
	if schemaRef.JSONSchema == nil || schemaRef.JSONSchema.Description != description {
		t.Fatalf("expected schema description %q, got %v", description, schemaRef.JSONSchema)
	}
}

func TestNewUpdatedSchemaReference_Success(t *testing.T) {
	// Arrange
	group := "foo.cnrm.cloud.google.com"
	version := "v1beta1"
	kind := "FooResource"
	oldCRD := newTestCRD(group, version, kind, "old description")
	oldRef := k8s.NewSchemaReference(oldCRD)
	newCRD := newTestCRD(group, version, kind, "new updated description")

	// Act
	newRef, err := k8s.NewUpdatedSchemaReference(oldRef, newCRD)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if newRef == nil {
		t.Fatal("expected non-nil newRef")
	}
	if newRef.JSONSchema.Description != "new updated description" {
		t.Fatalf("expected newRef description to be 'new updated description', got %q", newRef.JSONSchema.Description)
	}
	// Verify oldRef was not mutated (immutability guarantee for lock-free snapshotting)
	if oldRef.JSONSchema.Description != "old description" {
		t.Fatalf("expected oldRef description to remain 'old description', got %q", oldRef.JSONSchema.Description)
	}
}

func TestNewUpdatedSchemaReference_MismatchedGVK(t *testing.T) {
	// Arrange
	oldCRD := newTestCRD("foo.cnrm.cloud.google.com", "v1beta1", "FooResource", "description")
	oldRef := k8s.NewSchemaReference(oldCRD)
	differentCRD := newTestCRD("foo.cnrm.cloud.google.com", "v1beta1", "BarResource", "description")

	// Act
	newRef, err := k8s.NewUpdatedSchemaReference(oldRef, differentCRD)

	// Assert
	if err == nil {
		t.Fatal("expected error due to mismatched GVK, got nil")
	}
	if !strings.Contains(err.Error(), "unexpected mismatch of GVK") {
		t.Fatalf("expected error message to contain 'unexpected mismatch of GVK', got %q", err.Error())
	}
	if newRef != nil {
		t.Fatalf("expected newRef to be nil on error, got %v", newRef)
	}
}

func TestUpdateSchema(t *testing.T) {
	// Arrange
	group := "foo.cnrm.cloud.google.com"
	version := "v1beta1"
	kind := "FooResource"
	oldCRD := newTestCRD(group, version, kind, "original description")
	ref := k8s.NewSchemaReference(oldCRD)
	updatedCRD := newTestCRD(group, version, kind, "updated description")

	// Act
	err := k8s.UpdateSchema(ref, updatedCRD)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error updating schema: %v", err)
	}
	if ref.JSONSchema.Description != "updated description" {
		t.Fatalf("expected ref description to be 'updated description', got %q", ref.JSONSchema.Description)
	}
	if ref.CRD != updatedCRD {
		t.Fatalf("expected ref CRD to be updatedCRD, got %v", ref.CRD)
	}
}
