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

package v1beta1

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestInstanceRef_SetExternal(t *testing.T) {
	ref := &InstanceRef{
		Name:      "my-instance",
		Namespace: "default",
	}

	external := "projects/test-project/instances/test-instance"
	ref.SetExternal(external)

	if ref.External != external {
		t.Errorf("expected External %q, got %q", external, ref.External)
	}
	if ref.Name != "" {
		t.Errorf("expected Name to be cleared, got %q", ref.Name)
	}
	if ref.Namespace != "" {
		t.Errorf("expected Namespace to be cleared, got %q", ref.Namespace)
	}
}

func TestInstanceRef_Normalize_Idempotent(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	instanceObj := &unstructured.Unstructured{}
	instanceObj.SetGroupVersionKind(BigtableInstanceGVK)
	instanceObj.SetName("test-instance")
	instanceObj.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(instanceObj.Object, "projects/test-project/instances/test-instance", "status", "externalRef"); err != nil {
		t.Fatalf("failed setting status.externalRef: %v", err)
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(instanceObj).Build()

	ref := &InstanceRef{
		Name:      "test-instance",
		Namespace: "test-ns",
	}

	// First call to Normalize resolves the reference and populates External while clearing Name and Namespace.
	if err := ref.Normalize(ctx, client, "test-ns"); err != nil {
		t.Fatalf("first Normalize failed: %v", err)
	}
	if ref.External != "projects/test-project/instances/test-instance" {
		t.Errorf("unexpected external after first Normalize: %q", ref.External)
	}
	if ref.Name != "" || ref.Namespace != "" {
		t.Errorf("expected Name and Namespace to be empty after Normalize, got Name=%q Namespace=%q", ref.Name, ref.Namespace)
	}

	// Second call to Normalize should be idempotent and succeed without error.
	if err := ref.Normalize(ctx, client, "test-ns"); err != nil {
		t.Fatalf("second Normalize failed: %v", err)
	}
	if ref.External != "projects/test-project/instances/test-instance" {
		t.Errorf("unexpected external after second Normalize: %q", ref.External)
	}
}

func TestInstanceRef_Normalize_MutualExclusivity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	client := fake.NewClientBuilder().WithScheme(scheme).Build()

	ref := &InstanceRef{
		Name:     "test-instance",
		External: "projects/test-project/instances/test-instance",
	}

	if err := ref.Normalize(ctx, client, "test-ns"); err == nil {
		t.Fatal("expected error when both Name and External are specified, got nil")
	}
}

func TestInstanceRef_Normalize_Empty(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	client := fake.NewClientBuilder().WithScheme(scheme).Build()

	ref := &InstanceRef{}
	if err := ref.Normalize(ctx, client, "test-ns"); err != nil {
		t.Fatalf("expected nil for empty reference, got %v", err)
	}
}
