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

package oracledatabase

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestOracleDatabaseAutonomousDatabase_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-key")
	readyKey.SetNamespace("test-ns")
	selfLink := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-key"
	if err := unstructured.SetNestedField(readyKey.Object, selfLink, "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey).Build()

	t.Run("resolves KmsKeyRef by name", func(t *testing.T) {
		obj := &krm.OracleDatabaseAutonomousDatabase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-db",
				Namespace: "test-ns",
			},
			Spec: krm.OracleDatabaseAutonomousDatabaseSpec{
				Properties: &krm.AutonomousDatabaseProperties{
					EncryptionKey: &krm.EncryptionKey{
						KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "test-key",
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Properties.EncryptionKey.KmsKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Properties.EncryptionKey.KmsKeyRef.External, selfLink)
		}
	})

	t.Run("handles external KmsKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.OracleDatabaseAutonomousDatabase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-db",
				Namespace: "test-ns",
			},
			Spec: krm.OracleDatabaseAutonomousDatabaseSpec{
				Properties: &krm.AutonomousDatabaseProperties{
					EncryptionKey: &krm.EncryptionKey{
						KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							External: externalKey,
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Properties.EncryptionKey.KmsKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Properties.EncryptionKey.KmsKeyRef.External, externalKey)
		}
	})

	t.Run("returns error when referenced key is not found", func(t *testing.T) {
		obj := &krm.OracleDatabaseAutonomousDatabase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-db",
				Namespace: "test-ns",
			},
			Spec: krm.OracleDatabaseAutonomousDatabaseSpec{
				Properties: &krm.AutonomousDatabaseProperties{
					EncryptionKey: &krm.EncryptionKey{
						KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "nonexistent-key",
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err == nil {
			t.Fatalf("NormalizeReferences() expected error for nonexistent key, got nil")
		}
	})

	t.Run("handles nil EncryptionKey gracefully", func(t *testing.T) {
		obj := &krm.OracleDatabaseAutonomousDatabase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-db",
				Namespace: "test-ns",
			},
			Spec: krm.OracleDatabaseAutonomousDatabaseSpec{
				Properties: &krm.AutonomousDatabaseProperties{},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}
	})
}
