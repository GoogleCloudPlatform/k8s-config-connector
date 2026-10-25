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

func TestIAMServiceAccountRef_Normalize(t *testing.T) {
	ctx := context.Background()

	sa := &unstructured.Unstructured{}
	sa.SetGroupVersionKind(IAMServiceAccountGVK)
	sa.SetName("my-sa")
	sa.SetNamespace("test-ns")
	_ = unstructured.SetNestedField(sa.Object, "my-sa@test-proj.iam.gserviceaccount.com", "status", "email")

	reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(sa).Build()

	t.Run("resolves Name to status.email and clears Name and Namespace", func(t *testing.T) {
		ref := &IAMServiceAccountRef{Name: "my-sa", Namespace: "test-ns"}
		if err := ref.Normalize(ctx, reader, "default-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		if ref.External != "my-sa@test-proj.iam.gserviceaccount.com" {
			t.Errorf("External = %q, want %q", ref.External, "my-sa@test-proj.iam.gserviceaccount.com")
		}
		if ref.Name != "" || ref.Namespace != "" {
			t.Errorf("expected Name and Namespace to be cleared, got Name=%q Namespace=%q", ref.Name, ref.Namespace)
		}
		// Idempotent on second call (e.g., when VisitFields visits both Ptr and Struct)
		if err := ref.Normalize(ctx, reader, "default-ns"); err != nil {
			t.Fatalf("second Normalize() unexpected error: %v", err)
		}
	})

	t.Run("validates External email format", func(t *testing.T) {
		ref := &IAMServiceAccountRef{External: "not-an-email"}
		if err := ref.Normalize(ctx, reader, "test-ns"); err == nil {
			t.Fatalf("expected error for invalid email external, got nil")
		}
	})
}

func TestSecretManagerRefs_Normalize(t *testing.T) {
	ctx := context.Background()

	secret := &unstructured.Unstructured{}
	secret.SetGroupVersionKind(SecretManagerSecretGVK)
	secret.SetName("my-secret")
	secret.SetNamespace("test-ns")
	_ = unstructured.SetNestedField(secret.Object, "projects/test-proj/secrets/my-secret", "status", "externalRef")

	version := &unstructured.Unstructured{}
	version.SetGroupVersionKind(SecretManagerSecretVersionGVK)
	version.SetName("my-version")
	version.SetNamespace("test-ns")
	_ = unstructured.SetNestedField(version.Object, "projects/test-proj/secrets/my-secret/versions/1", "status", "externalRef")

	reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(secret, version).Build()

	t.Run("SecretManagerSecretRef resolves Name", func(t *testing.T) {
		ref := &SecretManagerSecretRef{Name: "my-secret"}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		if ref.External != "projects/test-proj/secrets/my-secret" {
			t.Errorf("External = %q, want projects/test-proj/secrets/my-secret", ref.External)
		}
		if ref.Name != "" {
			t.Errorf("expected Name to be cleared, got %q", ref.Name)
		}
	})

	t.Run("SecretManagerSecretVersionRef resolves Name", func(t *testing.T) {
		ref := &SecretManagerSecretVersionRef{Name: "my-version"}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		if ref.External != "projects/test-proj/secrets/my-secret/versions/1" {
			t.Errorf("External = %q, want projects/test-proj/secrets/my-secret/versions/1", ref.External)
		}
		if ref.Name != "" {
			t.Errorf("expected Name to be cleared, got %q", ref.Name)
		}
	})
}

func TestFolderAndBigQueryDatasetAndComputeRefs_Normalize(t *testing.T) {
	ctx := context.Background()

	folder := &unstructured.Unstructured{}
	folder.SetGroupVersionKind(FolderGVK)
	folder.SetName("my-folder")
	folder.SetNamespace("test-ns")
	_ = unstructured.SetNestedField(folder.Object, "123456789", "status", "folderId")

	dataset := &unstructured.Unstructured{}
	dataset.SetGroupVersionKind(BigQueryDatasetGVK)
	dataset.SetName("my_dataset")
	dataset.SetNamespace("test-ns")
	_ = unstructured.SetNestedField(dataset.Object, "projects/test-proj/datasets/my_dataset", "status", "externalRef")

	sa := &unstructured.Unstructured{}
	sa.SetGroupVersionKind(ComputeServiceAttachmentGVK)
	sa.SetName("my-sa")
	sa.SetNamespace("test-ns")
	_ = unstructured.SetNestedField(sa.Object, "https://www.googleapis.com/compute/beta/projects/test-proj/regions/us-central1/serviceAttachments/my-sa", "status", "selfLink")

	reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(folder, dataset, sa).Build()

	t.Run("FolderRef resolves Name via status.folderId fallback", func(t *testing.T) {
		ref := &FolderRef{Name: "my-folder"}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		if ref.External != "folders/123456789" {
			t.Errorf("External = %q, want folders/123456789", ref.External)
		}
	})

	t.Run("BigQueryDatasetRef resolves Name", func(t *testing.T) {
		ref := &BigQueryDatasetRef{Name: "my_dataset"}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		if ref.External != "projects/test-proj/datasets/my_dataset" {
			t.Errorf("External = %q, want projects/test-proj/datasets/my_dataset", ref.External)
		}
	})

	t.Run("ComputeServiceAttachmentRef resolves Name", func(t *testing.T) {
		ref := &ComputeServiceAttachmentRef{Name: "my-sa"}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		if ref.External != "projects/test-proj/regions/us-central1/serviceAttachments/my-sa" {
			t.Errorf("External = %q, want projects/test-proj/regions/us-central1/serviceAttachments/my-sa", ref.External)
		}
	})
}
