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

package tf

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
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

func TestUpdateSchema_Success(t *testing.T) {
	// Arrange
	group := "foo.cnrm.cloud.google.com"
	version := "v1beta1"
	kind := "FooResource"
	initialCRD := newTestCRD(group, version, kind, "initial description")
	r := &Reconciler{}
	r.schemaRef.Store(k8s.NewSchemaReference(initialCRD))
	updatedCRD := newTestCRD(group, version, kind, "updated description")

	// Act
	err := r.UpdateSchema(updatedCRD)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error updating schema: %v", err)
	}
	newRef := r.schemaRef.Load()
	if newRef == nil {
		t.Fatal("expected non-nil schemaRef after update")
	}
	if newRef.JSONSchema.Description != "updated description" {
		t.Fatalf("expected description 'updated description', got %q", newRef.JSONSchema.Description)
	}
	if newRef.GVK.Kind != kind || newRef.GVK.Group != group || newRef.GVK.Version != version {
		t.Fatalf("expected GVK %s/%s, %s, got %v", group, version, kind, newRef.GVK)
	}
}

func TestUpdateSchema_MismatchedGVK(t *testing.T) {
	// Arrange
	group := "foo.cnrm.cloud.google.com"
	version := "v1beta1"
	initialCRD := newTestCRD(group, version, "FooResource", "description")
	r := &Reconciler{}
	r.schemaRef.Store(k8s.NewSchemaReference(initialCRD))
	differentCRD := newTestCRD(group, version, "BarResource", "description")

	// Act
	err := r.UpdateSchema(differentCRD)

	// Assert
	if err == nil {
		t.Fatal("expected error due to mismatched GVK, got nil")
	}
	if !strings.Contains(err.Error(), "unexpected mismatch of GVK") {
		t.Fatalf("expected error message to contain 'unexpected mismatch of GVK', got %q", err.Error())
	}
	// Verify existing schemaRef remains unchanged
	currentRef := r.schemaRef.Load()
	if currentRef.GVK.Kind != "FooResource" {
		t.Fatalf("expected schemaRef to remain FooResource, got %s", currentRef.GVK.Kind)
	}
}
